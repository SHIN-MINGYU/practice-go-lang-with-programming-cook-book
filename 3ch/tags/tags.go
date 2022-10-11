package tags

import "fmt"

// Person structure have properties about name, city , state, and anohter property

type Person struct {
    Name string `serialize:"name"`
    City string `serialize:"city"`
    State string
    Misc string `serialize:"-"`
	Year int `serialize:"year"`
}

//EmptyStruct function show how to serialize and deserialize of empty structure had tag

func EmptyStruct() error {
	p := Person{}

	res, err := SerializeStructStrings(&p)

	if err!= nil {
        return err
    }

	fmt.Printf("Empty struct : %#v\n", p)

	fmt.Println("Serialize Results : ", res)

	newP := Person{}
	if err := DeserializeStructString(res, &newP); err != nil {
		return err 
	}

	fmt.Printf("Derialize Result : %#v\n", newP)
	return nil
}


// FullStruct function show how to deserialize and serialize a full filled structure had tag
func FullStruct() error {
	p := Person {
		Name: "John",
        City: "London",
        State: "Texas",
        Misc: "Misc",
        Year: 2019,
	}
	res, err := SerializeStructStrings(&p)
	if err != nil {
		return err
	}

	fmt.Printf("Full struct : %#v\n", p)
	fmt.Println("Serialize Results : ", res)

	newP := Person{}
	if err := DeserializeStructString(res, &newP); err!= nil {
		return err
	}
	fmt.Printf("Deserialize results : %#v\n", newP)
	return nil
}