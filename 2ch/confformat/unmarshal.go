package confformat

import "fmt"

const (
	exampleTOML = ` name ="Example1"
age=99
	`
	exampleJSON = `{"name": "Example2", "age":99}`
	exampleYAML = ` name: Example3
age: 99
	`
)

// UnmarshalAll function return structure from variable format
func UnmarshalAll() error {
	t := TOMLDATA{}
	j := JSONDATA{}
	y := YAMLDATA{}
	if _, err := t.Decode([]byte(exampleTOML)); err != nil {
		{
			return err
		}
	}

	fmt.Println("TOML Unmarshal = ", t)
	if err := j.Decode([]byte(exampleJSON)); err != nil {
		return err
	}
	fmt.Println("JSON Unmarshal = ", j)

	if err := y.Decode([]byte(exampleYAML)); err!= nil {
		return err
	}
	fmt.Println("YAML Unmarshal = ", y)
    return nil
}