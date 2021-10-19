package jparse

import (
	"encoding/json"
	"os"
)

var decoded map[string]interface{}

type parse interface {
	GetValue(name string) interface{}
	SetValue(name string, value interface{}) error
	Decode() map[string]interface{}
}

func (j *jsonFile) GetValue(name string) interface{} {
	return decoded[name]
}
func (j *jsonFile) SetValue(name string, value interface{}) error {
	decoded[name] = value
	e := j.updateFile(decoded)
	if e != nil {
		return e
	}
	return nil
}
func (j *jsonFile) updateFile(decode map[string]interface{}) error {
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
