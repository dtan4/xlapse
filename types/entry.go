package types

import (
	"fmt"

	"sigs.k8s.io/yaml"

	v1 "github.com/dtan4/xlapse/types/v1"
)

type Entries []*v1.Entry

func DecodeEntriesYAML(body []byte) (Entries, error) {
	es := Entries{}

	if err := yaml.Unmarshal(body, &es); err != nil {
		return Entries{}, fmt.Errorf("cannot decode YAML: %w", err)
	}

	return es, nil
}
