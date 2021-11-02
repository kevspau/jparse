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

func New(file string) (*jsonFile, error) {
	f, e := os.ReadFile(file)
	if e != nil {
		return nil, e
	}
	j := jsonFile{name: file, body: f}
	json.Unmarshal(f, &decoded)
	return &j, nil
}
