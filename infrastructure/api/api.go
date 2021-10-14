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

	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var Client HTTPClient

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

// Open returns a pointer of the local csv file
func (u *UrbanDictionary) Open() (*os.File, error) {
	file, err := os.OpenFile(u.CSVPath, os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Read takes every definition record from the csv file into a []Definition
func (u *UrbanDictionary) Read(id string) ([]*model.Definition, error) {
	definitions := make([]*model.Definition, 0)
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
			idCsv, err := strconv.Atoi(line[0])
			if err != nil {
				return nil, err
			}

			definition := model.Definition{
				Defid:      idCsv,
				Word:       line[1],
				WrittenOn:  line[2],
				Definition: line[3],
				Permalink:  line[4],
				Example:    line[5],
			}
			definitions = append(definitions, &definition)
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
