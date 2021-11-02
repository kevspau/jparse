package jparse

import (
  "encoding/json"
  "os"
)

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
