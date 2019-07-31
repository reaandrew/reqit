package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type ReqitRequest struct {
	RequestObject ReqitRequestData `yaml:"request"`
}

type ReqitRequestData struct {
	Data   []byte `yaml:"data"`
	Type   string `yaml:"type"`
	Method string `yaml:"method"`
	URL    string `yaml:"url"`
}

type ReqitResult struct {
}

type HttpClient interface {
	Execute(request ReqitRequest) ReqitResult
}

type FakeHTTPClient struct {
	lastRequest ReqitRequest
}

func (self *FakeHTTPClient) Execute(request ReqitRequest) ReqitResult {
	self.lastRequest = request
	return ReqitResult{}
}

func (self *FakeHTTPClient) Request() ReqitRequest {
	return self.lastRequest
}

type RequestWithAssertions struct {
	request ReqitRequest
}

func (self RequestWithAssertions) IsOfType(requestType string) bool {
	result := self.request.RequestObject.Type == requestType
	if !result {
		log.Println(fmt.Sprintf("type = %s", self.request.RequestObject.Type))
	}
	return result
}

func CreateFakeHTTPClient() *FakeHTTPClient {
	return &FakeHTTPClient{}
}

type FakeRequestReader struct {
	data string
}

func (self FakeRequestReader) Read() string {
	return self.data
}

func CreateFakeRequestReader(data string) FakeRequestReader {
	return FakeRequestReader{
		data: data,
	}
}

type ReqitRequestReader interface {
	Read() string
}

type ReqitClient struct {
	httpClient HttpClient
}

func (self ReqitClient) Execute(reader ReqitRequestReader) ReqitResult {
	reqitData := reader.Read()
	stringReader := strings.NewReader(reqitData)
	scanner := bufio.NewScanner(stringReader)
	request := []string{}
	data := []string{}
	line := 0
	setData := false
	for scanner.Scan() {
		lineContent := scanner.Text()
		if line > 0 && lineContent == "---" {
			setData = true
		}

		if !setData {
			request = append(request, lineContent)
		} else {
			if lineContent != "---" {
				data = append(data, lineContent)
			}
		}
		line++
	}

	requestObject := ReqitRequest{}
	dataToDecode := strings.Join(request, "\n")
	err := yaml.Unmarshal([]byte(dataToDecode), &requestObject)

	if err != nil {
		panic(err)
	}
	requestObject.RequestObject.Data = []byte(strings.Join(data, "\n"))

	return self.httpClient.Execute(requestObject)
}

func CreateReqitClient(httpClient HttpClient) ReqitClient {
	return ReqitClient{
		httpClient: httpClient,
	}
}

func main() {
	f, _ := os.Open("sample2.yml")
	scanner := bufio.NewScanner(f)
	request := []string{}
	data := []string{}
	line := 0
	setData := false
	for scanner.Scan() {
		lineContent := scanner.Text()
		if line > 0 && lineContent == "---" {
			setData = true
		}

		if !setData {
			request = append(request, lineContent)
		} else {
			if lineContent != "---" {
				data = append(data, lineContent)
			}
		}
		line++
	}

	fmt.Println(strings.Join(request, "\n"))
	fmt.Println(data)
}
