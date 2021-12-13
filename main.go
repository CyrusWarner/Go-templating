package main

// Think of templates as a containernholding all of your templates
import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// CreateTemplateWithDumpingToIndexFile()
	// CreateTemplateByCreatingAFileUsingTheOsPackage
}

// running go main.go > index.html will dump the template html into an index file
func CreateTemplateWithDumpingToIndexFile() {
	tpl, err := template.ParseFiles("tpl.gohtml") // Parses the files and returns a pointer to the template and an error
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil) // executes the template
	if err != nil {
		log.Fatalln(err)
	}
}

func CreateTemplateByCreatingAFileUsingTheOsPackage() {
	name := "Cyrus Warner"
	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="TUF-8>
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`)

	nf, err := os.Create("index.html") // Creates a file using the standard os library
	if err != nil {
		log.Fatal("Error creating file", err)
	}
	defer nf.Close() // we defer the nf.Close function to run after the template method is complete

	io.Copy(nf, strings.NewReader(str))
}
