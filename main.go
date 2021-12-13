package main

import (
	"fmt"
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
	name := "Cyrus Warner"
	template := `
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
	`

	fmt.Println(template)
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
