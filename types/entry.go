package types

import (
	"fmt"

	"sigs.k8s.io/yaml"
)

type Entry struct {
	URL       string `json:"url"`
	Bucket    string `json:"bucket"`
	KeyPrefix string `json:"key_prefix"`
	Timezone  string `json:"timezone"`
}

type Entries []*Entry

func DecodeEntriesYAML(body []byte) (Entries, error) {
	es := Entries{}

	if err := yaml.Unmarshal(body, &es); err != nil {
		return Entries{}, fmt.Errorf("cannot decode YAML: %w", err)
	}

	return es, nil
}
