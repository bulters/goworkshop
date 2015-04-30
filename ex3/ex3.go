package main

import (
	"fmt"
)

// BEGIN IF OMIT
type Renderer interface {
	Render() string
}

// END IF OMIT

type Template struct {
	Content string
}

// BEGIN COND OMIT
func (t Template) Render() string {
	if t.Content == "ERROR" {
		return "An error ha..."
	} else {
		return t.Content
	}
}

// END COND OMIT

// BEGIN IFI OMIT
type Log struct {
	Message string
}

func (l Log) Render() string {
	return l.Message
}

// END IFI OMIT

func WrapRender(r Renderer) string {
	return fmt.Sprintf("%s\t%s", "MSG", r.Render())
}

func main() {
	// BEGIN ARRAY OMIT
	var messages []Renderer = []Renderer{
		Log{"Log bericht"},
		Template{"Vanuit een template"},
		Template{"ERROR"}}
	// END ARRAY OMIT

	// BEGIN FOR OMIT
	for _, msg := range messages {
		rendered := WrapRender(msg)
		fmt.Println(rendered)
	}
	// END FOR OMIT
}
