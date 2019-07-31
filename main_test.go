package main_test

import (
	"testing"

	. "github.com/reaandrew/reqit"
	"github.com/stretchr/testify/assert"
)

var simpleRequest = `---
request: 
  type: http
  method: POST
  url: https://somewhere
  headers:
    X-SOMETHING: Boom
    Content-Type: application/json
  verify: false
  pretty: true
  before:
    - ./get-reference-data-badges.yml
---
{
  "name":"barney",
}
`

func TestSimpleRequest(t *testing.T) {
	httpClient := CreateFakeHTTPClient()
	reqit := CreateReqitClient(httpClient)
	reader := CreateFakeRequestReader(simpleRequest)
	result := reqit.Execute(reader)

	t.Run("result is not nil", func(t *testing.T) {
		assert.NotNil(t, result)
	})

	t.Run("request type is http", func(t *testing.T) {
		assert.Equal(t, httpClient.Request().RequestObject.Type, "http")
	})

	t.Run("request method is POST", func(t *testing.T) {
		assert.Equal(t, httpClient.Request().RequestObject.Method, "POST")
	})

	t.Run("request URL is https://somewhere", func(t *testing.T) {
		assert.Equal(t, httpClient.Request().RequestObject.URL, "https://somewhere")
	})
}
