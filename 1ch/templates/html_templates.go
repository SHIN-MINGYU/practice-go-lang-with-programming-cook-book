package templates

import (
	"fmt"
	"html/template"
	"os"
)

// HTMLDifferences function emphasize the differences between text/template and html/template
func HTMLDifferences() error {
	t := template.New("html")
	t, err := t.Parse("<h1>Hello {{.Name}}</h1>")
	if err!= nil {
        return err
    }

	// html/template process automatically escapes dangerous action like injecting javascript
	// this be recognized by the situation, work different according to variables location
	err = t.Execute(os.Stdout, map[string]string{"Name":"<script>alert('Hello {{.Name}} ')</script>"})
	if err!= nil {
        return err
    }

	// can call escaper manually
	fmt.Println(template.JSEscaper(`example <example@example.com>`))
	
	fmt.Println(template.HTMLEscaper(`example <example@example.com>`))
	fmt.Println(template.URLQueryEscaper(`example <example@example.com>`))
	return nil
}