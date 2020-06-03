package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/reaandrew/reqit/core"
	"github.com/reaandrew/reqit/fileio"
	"github.com/reaandrew/reqit/http"
	"github.com/reaandrew/reqit/reporters"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var (
	CommitHash string
	Version    string
	BuildTime  string
)

func existing(client core.HTTPClient, reader core.RequestReader) core.Result {
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

	requestObject := core.Request{}
	dataToDecode := strings.Join(request, "\n")
	err := yaml.Unmarshal([]byte(dataToDecode), &requestObject)

	if err != nil {
		panic(err)
	}
	requestObject.RequestObject.Data = []byte(strings.Join(data, "\n"))

	return client.Execute(requestObject)
}

func Execute(args []string) {
	app := cli.NewApp()
	app.Name = "boom"
	app.Version = Version
	app.Metadata = map[string]interface{}{}
	app.Metadata["CommitHash"] = CommitHash
	app.Metadata["BuildTime"] = BuildTime
	app.Usage = "Schmokin"
	app.Action = func(c *cli.Context) error {
		filepath := c.Args().Get(0)
		if filepath == "" {
			return errors.New("Filepath required")
		}
		client := http.DefaultHTTPClient{}
		result := existing(client, fileio.FileRequestReader{
			Path: filepath,
		})
		reporters.CliReporter{}.Execute(result)
		return nil
	}

	err := app.Run(args)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	Execute(os.Args)
}
