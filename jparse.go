package jparse

import (
	"encoding/json"
	"os"
)

type parse interface {
	GetValue(name string) interface{}
	SetValue(name string, value string) error
}

func (j *jsonFile) GetValue(name string) interface{} {
	return j.decode[name]
}
func (j *jsonFile) SetValue(name string, value string) error {
	j.decode[name] = value
	b, e := json.Marshal(decode)
	if e != nil {
		return e
	}
	e = os.WriteFile(j.name, b, 0777)
	if e != nil {
		return e
	}
	return nil
}

type jsonFile struct {
	name string
	body []byte
	decode *map[string]interface{}
	parse
}

func New(file string) (*jsonFile, error) {
	f, e := os.ReadFile(file)
	if e != nil {
		return nil, e
	}
	json := jsonFile{name: file, body: f, decoded: json.Unmarshal(f, &map[string]interface{})}
	return &json, nil
}
