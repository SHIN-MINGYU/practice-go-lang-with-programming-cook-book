package confformat

import (
	"bytes"

	"github.com/go-yaml/yaml"
)

// YAML structure have YARML structure tag
type YAMLDATA struct {
	Name string `yaml:"name"`
	Age int `yaml:"age"`
}

// ToYAML RETURN YAML format bytes.Buffer from  YAMLDATA structure
func (y *YAMLDATA) ToYAML() (*bytes.Buffer, error) {
	d, err :=yaml.Marshal(y)
	if err!= nil {
        return nil, err
    }
	return bytes.NewBuffer(d), nil
}

// Decode function decode to YAMLDATA structure

func (y *YAMLDATA)Decode(data []byte) ( error) {
    return yaml.Unmarshal(data, &y)
}

