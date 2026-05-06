package ui

import (
	"fmt"
	"io"
)

func RenderPrompt(w io.Writer) {

	fmt.Fprint(w, Prompt())
}

func RenderOutput(w io.Writer, text string) {

	fmt.Fprintln(w, Output(text))
}

func RenderError(w io.Writer, text string) {

	fmt.Fprintln(w, ErrorText(text))
}
