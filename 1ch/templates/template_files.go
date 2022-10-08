package templates

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

// CreateTemplate fucntion create template files including the data
func CreateTemplate(path string, data string)  error {
	return ioutil.WriteFile(path, []byte(data),os.FileMode(0755))
}

// InitTemplates function set template from directory
func InitTemplates() error {
	tempdir, err := ioutil.TempDir("","temp")
	if err!= nil {
        return err
    }
	defer os.RemoveAll(tempdir)

	err = CreateTemplate(filepath.Join(tempdir, "t1.tmpl"), `
        Template 1! {{ .Var1 }}
        {{ block "template2" .}} {{end}}
        {{ block "template3" .}} {{end}}
    `)
    if err!= nil {
        return err
    }
	err = CreateTemplate(filepath.Join(tempdir, "t2.tmpl"),
	`{{ define "template2"}}Template 2! {{ .Var2 }}{{end}}`)
	if err!= nil {
        return err
    }

	err = CreateTemplate(filepath.Join(tempdir, "t3.tmpl"),
	`{{ define "template3"}}Template 3! {{ .Var3 }}{{end}}`)
	if err!= nil {
        return err
    }

	pattern := filepath.Join(tempdir,"*.tmpl")

	//ParseGlob function combine one template to gather all files what coincide pattern
	tmpl, err := template.ParseGlob(pattern)
    if err!= nil {
        return err
    }

	// Execute function can use in map type instead of struct type
	tmpl.Execute(os.Stdout, map[string]string{
		"Var1": "Template 1!",
        "Var2": "Template 2!",
        "Var3": "Template 3!",
	})
	
	return nil
 }