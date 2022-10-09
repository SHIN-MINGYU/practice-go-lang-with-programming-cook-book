package confformat

import "fmt"

// MarshalAll function return variable foramt from saved data at structure
func MarshalAll() error {
	t := TOMLDATA {
		Name : "Name1",
		Age : 20,
	}
	j := JSONDATA {
		Name : "Name2",
		Age: 30,
	}
	y := YAMLDATA {
		Name : "Name3",
        Age : 40,
	}

	tomlRes, err := t.ToTOML()
	if err != nil {
		return err
	}

	fmt.Println("TOML Mashaled = ", tomlRes.String())

	jsonRes, err := j.ToJSON()
    if err!= nil {
		return err
	}

	fmt.Println("JSON Mashaled = ", jsonRes.String())

	yamlRes, err := y.ToYAML()
    if err!= nil {
		return err
	}
	fmt.Println("YAML Mashaled = ", yamlRes.String())
	return nil
}