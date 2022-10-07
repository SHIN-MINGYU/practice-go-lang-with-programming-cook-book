package csvformat

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

// Book structor have Author and Title
type Book struct{
	Author string
	Title string
}

// Books is array of Book structor type
// using for save a lot of books info
type Books []Book


// ToCSV function get input what a lot of Book.
// And , Atfer write to io.Writer, The case of error return error
 func (books *Books) ToCSV(w io.Writer) error {
	n := csv.NewWriter(w)

	err := n.Write([]string{"Author","Title"})
	if err != nil {
		return err
	}

	for _, book := range *books {
		err := n.Write([]string{book.Author, book.Title})
		if err != nil {
			return err
		}
	}
	defer n.Flush()
	return n.Error()
 }

 // WriteCSVOutput function initailize many info of books
 // and record this in os.Stdout

func WriteCSVOutput() error {
	b := Books{
		Book {
			Author : "F Scott Fitzgerald",
			Title : "위대한 게츠비",
		},
		Book{
			Author : "J D Salinger",
			Title : "The Catcher in the Rye",
		},
	}
	return b.ToCSV(os.Stdout)
}

// WriteCSVBuffer function return buffer csv what include many info of books
func WriteCSVBuffer() (*bytes.Buffer, error){
	b := Books{
		Book {
			Author : "F Scott Fitzgerald",
			Title : "위대한 게츠비",
		},
		Book{
			Author : "J D Salinger",
			Title : "The Catcher in the Rye",
		},
	}
	w := &bytes.Buffer{}
	err := b.ToCSV(w)
	return w, err
}