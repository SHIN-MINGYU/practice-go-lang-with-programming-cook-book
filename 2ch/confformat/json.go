package confformat

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JSONDATA structure is noraml structure that hava JSON structure tag
type JSONDATA struct {
	Name string `json:"name"`	
	Age int `json:"age"`
}

// ToJSON function return Json format bytes.Buffer have JSON structure tag
func (j *JSONDATA) ToJSON() (*bytes.Buffer, error) {
    d, err := json.Marshal(j)
	if err!= nil {
		return nil, err
	}
	return bytes.NewBuffer(d), nil
}

// Decode function decode to JSON structure
func (j *JSONDATA) Decode(data []byte) error {
    return json.Unmarshal(data, j)
}

// OtherJSONExamples function examines How to use JSON in another structure or another useful function
func OtherJSONExamples() error {
	res := make(map[string]string)
	err := json.Unmarshal([]byte(`{"key":"value"}`),&res)
	if err!= nil {
        return err
    }
	fmt.Println("We can unmarshal into a map instead of a struct : ", res)
	
	b := bytes.NewReader([]byte(`{"key2":"value2"}`))
	decoder := json.NewDecoder(b)

	if err := decoder.Decode(&res); err != nil {
		return err 
	}
	fmt.Println("we can also use decoders/encoders to work with stream : ",res)
	return nil
}