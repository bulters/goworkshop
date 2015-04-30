package main

import (
	"fmt"
)

// BEGIN STRUCT OMIT
type Template struct {
	Content string
}

// END STRUCT OMIT

// BEGIN SFUN OMIT
func (t Template) Render() string {
	return t.Content
}

// END SFUN OMIT

// BEGIN FUN OMIT
func main() {
	template := Template{"Dit is een test"}
	fmt.Println(RenderTemplate(template))
}

func RenderTemplate(t Template) string {
	return t.Render()
}

// END FUN OMIT
