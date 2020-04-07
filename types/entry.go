package types

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Entry struct {
	URL       string `yaml:"url" json:"url"`
	Bucket    string `yaml:"bucket" json:"bucket"`
	KeyPrefix string `yaml:"key_prefix" json:"key_prefix"`
	Timezone  string `yaml:"timezone" json:"timezone"`
}

type Entries []*Entry

func DecodeEntriesYAML(body []byte) (Entries, error) {
	es := Entries{}

	if err := yaml.Unmarshal(body, &es); err != nil {
		return Entries{}, fmt.Errorf("cannot decode YAML: %w", err)
	}

	return es, nil
}
