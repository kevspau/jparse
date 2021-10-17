package main

import (
	"encoding/json"
	"os"
)

type parse interface {
	GetValue(name string) string
	SetValue(name string, value string) error
	Decode() *map[string]interface{}
}

func (j *jsonFile) GetValue(name string) string {
	var decoded map[string]interface{}
	json.Unmarshal(j.body, &decoded)
	return decoded[name].(string)
}
func (j *jsonFile) SetValue(name string, value string) error {
	var decoded map[string]interface{}
	json.Unmarshal(j.body, &decoded)
	decoded[name] = value
	b, e := json.Marshal(decoded)
	if e != nil {
		return e
	}
	e = os.WriteFile(j.name, b, 0777)
	return e
}
func (j *jsonFile) Decode() *map[string]interface{} {
	var decoded map[string]interface{}
	json.Unmarshal(j.body, &decoded)
	return &decoded
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
	json := jsonFile{name: file, body: f}
	return &json, nil
}
