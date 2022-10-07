package csvformat

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// Movie structor is used to save parsed csv
type Movie struct {
	Title string
	Director string
	Year int
}

// ReadCSV function is manifasted example to do csv conveied in io.Reader
func ReadCSV(b io.Reader) ([]Movie, error){
	r := csv.NewReader(b)

	// this is optional work for setting option
	r.Comma = ';' // this is delimeter
	r.Comment = '-' // this is annoation

	var movies []Movie

	// just now, get header and ignore that
	// header can be used for dictionary key or lookup for another shape
	_, err := r.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	// loop until done all csv process
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}else if err != nil {
			return nil, err
		}
		year ,err := strconv.ParseInt(record[2],10,64);
		
		if err != nil {
			return nil, err
		}
		m := Movie{record[0],record[1],int(year)}
		movies = append(movies, m)
	}
	return movies, nil

 }

 // AddMoviesFromText func  use CSV Parser
 func AddMoviesFromText() error {
	// this example show example to read converted buffer from string by csv package
	in := `
movie title;director;year released

Guardians of the Galaxy Vol. 2;James Gunn;2017
Star Wars: Episode VIII;Rian Johnson;2017
`
	b := bytes.NewBufferString(in)
	m, err := ReadCSV(b)
	if err != nil {
		return err
	}

	fmt.Printf("%#vn",m)
	return nil
}