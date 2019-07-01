package main

import (
	"bytes"
	"fmt"
	"testing"

	"gopkg.in/yaml.v3"
)

func Test_Something(t *testing.T) {
	data := `a: 1
b: 2
---
{

}`
	dataReader := bytes.NewReader([]byte(data))

	fmt.Println(dataReader.)
}
