package jparse

import (
	"encoding/json"
	"os"
)

var decoded map[string]interface{}

type parse interface {
	GetValue(name string) interface{}
	SetValue(name string, value string) error
	Decode() *map[string]interface{}
}

func (j *jsonFile) GetValue(name string) interface{} {
	return decoded[name]
}
func (j *jsonFile) SetValue(name string, value string) error {
	decoded[name] = value
	b, e := json.Marshal(decoded)
	if e != nil {
		return e
	}
	e = os.WriteFile(j.name, b, 0777)
	j.body = b
	if e != nil {
		return e
	}
	return nil
}
func (j *jsonFile) Decode() map[string]interface{} {
	return decoded
}

type jsonFile struct {
	name string
	body []byte
	parse
}

func New(file string) (*jsonFile, error) {
	f, e := os.ReadFile(file)
	if e != nil {
		return nil, e
	}
	j := jsonFile{name: file, body: f}
	json.Unmarshal(f, &decoded)
	return &j, nil
}
/*
..\..\..\go\pkg\mod\github.com\csharpdf\jparse@v0.0.0-20211017041748-c4870ae2b8f5\jparse.go:41:40: cannot use json.Unmarshal(f, map[string]interface {}) (type error) as type map[string]interface {} in field value
..\..\..\go\pkg\mod\github.com\csharpdf\jparse@v0.0.0-20211017041748-c4870ae2b8f5\jparse.go:41:66: type map[string]interface {} is not an expression
*/