package confformat

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

// TOMLDATA struct is normal data structure having TOML Structure tag
type TOMLDATA struct{
	Name string `toml:"name"`
	Age  int    `toml:"age"`
}

// ToTOML function return TOML format bytes.Buffer from  TOML structure
func (t *TOMLDATA) ToTOML() (* bytes.Buffer, error) {
	b := &bytes.Buffer{}
	encoder := toml.NewEncoder(b)
	if err := encoder.Encode(t); err != nil {
		return nil, err
	}
	return b,nil
}

func (t* TOMLDATA) Decode(data []byte) (toml.MetaData, error) {
	return toml.Decode(string(data), t)
}