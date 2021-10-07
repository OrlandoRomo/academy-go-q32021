package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/repository"
)

const (
	UrbanDictionaryURL = "https://mashape-community-urban-dictionary.p.rapidapi.com/define"
	RapidapiKeyName    = "x-rapidapi-key"
	CSVPath            = "data/definitions.csv"
)

type UrbanDictionary struct {
	ApiURL  string
	Headers map[string]string
	CSVPath string
	repository.UrbanReaderWriter
}

func NewUrbanDictionary(apiKey string) *UrbanDictionary {
	return &UrbanDictionary{
		ApiURL: UrbanDictionaryURL,
		Headers: map[string]string{
			RapidapiKeyName: apiKey,
		},
		CSVPath: CSVPath,
	}
}

func (u *UrbanDictionary) GetDefinitions(term string) (*model.List, error) {
	var list *model.List
	url := fmt.Sprintf("%s?term=%s", u.ApiURL, term)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add(RapidapiKeyName, u.Headers[RapidapiKeyName])

	response, err := http.DefaultClient.Do(request)
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
	return list, nil
}

func (u *UrbanDictionary) GetDefinitionsCSV(id string) (*model.List, error) {
	list := new(model.List)
	definitions := make([]*model.Definition, 0)

	file, err := u.Read()
	if err != nil {
		return nil, err
	}
	csvReader, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}
	for _, line := range csvReader {
		if id == line[0] {
			idCsv, _ := strconv.Atoi(line[0])
			definition := model.Definition{
				Defid:      int64(idCsv),
				Word:       line[1],
				Definition: line[2],
				Permalink:  line[3],
				Example:    line[4],
			}
			definitions = append(definitions, &definition)
			break
		}
	}

	list.Definitions = definitions

	return list, nil
}

func (u *UrbanDictionary) Read() (*os.File, error) {
	file, err := os.OpenFile(u.CSVPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return file, err
}

func (u *UrbanDictionary) Write(definitionsList *model.List) error {
	file, err := os.OpenFile(u.CSVPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	csvWriter := csv.NewWriter(file)
	for _, definition := range definitionsList.Definitions {
		id := strconv.Itoa(int(definition.Defid))
		_ = csvWriter.Write([]string{id, definition.Word, definition.Definition, definition.Permalink, definition.Example})
	}
	csvWriter.Flush()
	file.Close()
	return err
}
