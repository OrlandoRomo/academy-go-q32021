package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
	"github.com/OrlandoRomo/academy-go-q32021/workerpool"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

const (
	Odd  string = "odd"
	Even string = "even"
)

var (
	Client       HTTPClient
	WorkerPool   *workerpool.WorkerPool
	mutex        sync.Mutex
	itemsCounter int32 = 0
)

const (
	urbanDictionaryURL = "https://mashape-community-urban-dictionary.p.rapidapi.com/define"
	rapidapiKeyName    = "x-rapidapi-key"
	csvPath            = "data/definitions.csv"
)

type UrbanDictionary struct {
	ApiURL  string
	Headers map[string]string
	CSVPath string
}

func init() {
	Client = &http.Client{}
}

// NewUrbanDictionary returns a new instance of the UrbanDictionary client
func NewUrbanDictionary(apiKey string) *UrbanDictionary {
	return &UrbanDictionary{
		ApiURL: urbanDictionaryURL,
		Headers: map[string]string{
			rapidapiKeyName: apiKey,
		},
		CSVPath: csvPath,
	}
}

func (u *UrbanDictionary) UpdateCSVPath(path string) {
	u.CSVPath = path
}

// GetDefinitions reach out urban dictionary API by term
func (u *UrbanDictionary) GetDefinitions(term string) (*model.List, error) {
	var list *model.List
	url := fmt.Sprintf("%s?term=%s", u.ApiURL, term)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add(rapidapiKeyName, u.Headers[rapidapiKeyName])

	response, err := Client.Do(request)
	if err != nil {
		return nil, err
	}
	err = errorStatus(response.StatusCode)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}
	if len(list.Definitions) == 0 {
		return nil, model.ErrNotFound{Term: term}
	}
	return list, nil
}

// GetDefinitionById reads a local csv to find the definition by id paramater
func (u *UrbanDictionary) GetDefinitionById(id string) (*model.List, error) {
	list := new(model.List)

	definitions, err := u.Read(id)
	if err != nil {
		return nil, err
	}
	if len(definitions) == 0 {
		return nil, model.ErrNotFoundInCSV{Id: id}
	}
	list.Definitions = definitions

	return list, nil
}

// GetConcurrentDefinitions reads concurrently the local csv file
func (u *UrbanDictionary) GetConcurrentDefinitions(idType string, items, itemsWorker int) (*model.List, error) {
	list := new(model.List)
	workers := items / itemsWorker
	results := make(chan model.Definition, items)
	if WorkerPool == nil {
		WorkerPool = workerpool.NewWorkerPool()
	}
	WorkerPool.InitChan()
	WorkerPool.AddWorkers(workers)
	itemsCounter = 0
	wg := sync.WaitGroup{}
	file, err := u.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()
	csvReader := csv.NewReader(file)

	go func() {
		for definition := range results {
			list.Definitions = append(list.Definitions, definition)
		}
	}()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			WorkerPool.Add(worker(csvReader, idType, items, itemsWorker, results))
			wg.Done()
		}()
	}

	wg.Wait()

	WorkerPool.ShutDown()

	return list, nil
}
func worker(reader *csv.Reader, idType string, total, itemsWork int, results chan<- model.Definition) func() {
	return func() {
		counter := 0
		for {
			if int(itemsCounter) == total {
				break
			}
			if counter == itemsWork {
				break
			}
			mutex.Lock()
			line, err := reader.Read()
			mutex.Unlock()
			if err == io.EOF {
				break
			}
			if err != nil {
				break
			}
			idCsv, err := strconv.Atoi(line[0])
			if err != nil {
				break
			}
			if includeDefinition(idType, idCsv) {

				definition, err := parseDefinition(line)
				if err != nil {
					break
				}
				results <- definition
				counter++
				atomic.AddInt32(&itemsCounter, 1)
			}
		}
	}
}

// Open returns a pointer of the local csv file
func (u *UrbanDictionary) Open() (*os.File, error) {
	file, err := os.OpenFile(u.CSVPath, os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Read takes every definition record from the csv file into a []Definition
func (u *UrbanDictionary) Read(id string) ([]model.Definition, error) {
	definitions := make([]model.Definition, 0)
	file, err := u.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		if id == line[0] {
			definition, err := parseDefinition(line)
			if err != nil {
				return nil, err
			}
			definitions = append(definitions, definition)
			break
		}
	}

	return definitions, nil
}

// Write updates the local csv file with incoming definitions
func (u *UrbanDictionary) Write(definitionsList *model.List) error {
	file, err := u.Open()
	if err != nil {
		return err
	}

	defer file.Close()

	csvWriter := csv.NewWriter(file)
	for _, definition := range definitionsList.Definitions {
		id := strconv.Itoa(definition.Defid)
		err = csvWriter.Write([]string{
			id,
			definition.Word,
			definition.WrittenOn,
			definition.Definition,
			definition.Permalink,
			definition.Example})
		if err != nil {
			return err
		}
	}
	csvWriter.Flush()
	if csvWriter.Error() != nil {
		return csvWriter.Error()
	}
	return nil
}

// errorStatus returns the possible errors from the External Urban dictionary API
func errorStatus(code int) error {
	switch code {
	case http.StatusForbidden:
		return model.ErrMissingApiKey{}
	case http.StatusBadRequest:
		return model.ErrInvalidData{Field: "term"}
	default:
		return nil
	}
}

// includeDefinition returns a bool if the definitionIs is even or odd based on idType
func includeDefinition(idType string, definitionId int) bool {
	if idType == Even {
		return definitionId%2 == 0
	}
	return definitionId%2 != 0
}

func parseDefinition(str []string) (model.Definition, error) {
	idCsv, err := strconv.Atoi(str[0])
	if err != nil {
		return model.Definition{}, err
	}
	return model.Definition{
		Defid:      idCsv,
		Word:       str[1],
		WrittenOn:  str[2],
		Definition: str[3],
		Permalink:  str[4],
		Example:    str[5],
	}, nil
}
