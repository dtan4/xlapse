package types

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	"sigs.k8s.io/yaml"

	v1 "github.com/dtan4/xlapse/types/v1"
)

func DecodeEntriesYAML(body []byte) (*v1.Entries, error) {
	j, err := yaml.YAMLToJSON(body)
	if err != nil {
		return nil, fmt.Errorf("failed to convert entries YAML to JSON: %w", err)
	}

	var es v1.Entries

	if err := protojson.Unmarshal(j, &es); err != nil {
		return nil, fmt.Errorf("failed to unmarshal entries JSON to object: %w", err)
	}

	return &es, nil
}
