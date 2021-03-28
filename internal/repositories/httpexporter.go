package repositories

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/brandoncate-personal/monitor-cli/internal/core/domain"
)

type HttpExporter struct {
	Url string
}

func (exporter HttpExporter) Export(domain.TestResults) error {
	postBody, _ := json.Marshal(exporter)
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(exporter.Url, "application/json", responseBody)
	//Handle Error
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	sb := string(body)
	log.Println(sb)

	return nil
}
