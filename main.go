package main

import "fmt"

func main() {
	template()
}

// running go main.go > index.html will dump the template html into an index file
func template() {
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
