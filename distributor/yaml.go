package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Entry struct {
	URL       string `yaml:"url"`
	Bucket    string `yaml:"bucket"`
	KeyPrefix string `yaml:"key_prefix"`
	Timezone  string `yaml:"timezone"`
}

type Entries []*Entry

func decodeYAML(body []byte) (Entries, error) {
	es := Entries{}

	if err := yaml.Unmarshal(body, &es); err != nil {
		return Entries{}, fmt.Errorf("cannot decode YAML: %w", err)
	}

	return es, nil
}
