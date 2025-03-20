package main

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Field1 MaybeStringArray `yaml:"field1,omitempty" json:"field1,omitempty"`
	Field2 MaybeStringArray `yaml:"field2,omitempty" json:"field2,omitempty"`
	Field3 MaybeStringArray `yaml:"field3,omitempty" json:"field3,omitempty"`
	Field4 MaybeStringArray `yaml:"field4,omitempty" json:"field4,omitempty"`
}

type MaybeStringArray []string

var (
	_ yaml.Unmarshaler = &MaybeStringArray{}
	_ yaml.Marshaler   = MaybeStringArray{}
	_ json.Unmarshaler = &MaybeStringArray{}
	_ json.Marshaler   = MaybeStringArray{}
)

func (a *MaybeStringArray) UnmarshalYAML(value *yaml.Node) error {
	var slice []string
	if err := value.Decode(&slice); err == nil {
		*a = slice
		return nil
	}

	var single string
	if err := value.Decode(&single); err != nil {
		return err
	}
	*a = []string{single}
	return nil
}

func (a MaybeStringArray) MarshalYAML() (interface{}, error) {
	switch len(a) {
	case 0:
		return nil, nil
	case 1:
		return a[0], nil
	default:
		return []string(a), nil
	}
}

func (a *MaybeStringArray) UnmarshalJSON(data []byte) error {
	var slice []string
	if err := json.Unmarshal(data, &slice); err == nil {
		*a = slice
		return nil
	}

	var single string
	if err := json.Unmarshal(data, &single); err != nil {
		return err
	}
	*a = []string{single}
	return nil
}

func (a MaybeStringArray) MarshalJSON() ([]byte, error) {
	switch len(a) {
	case 0:
		return json.Marshal(nil)
	case 1:
		return json.Marshal(a[0])
	default:
		return json.Marshal([]string(a))
	}
}
