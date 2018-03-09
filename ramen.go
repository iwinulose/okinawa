// Copyright (C) 2018 Charles Duyk. All rights reserved.
// 
// Available under BSD 3 clause license. See 'LICENSE'.


// Package okinawa implements a simple struct pretty printer.
//
// Why okinawa? Why Ramen? Perhaps I was hungry at the time, or was just sick of boring names.
// The go maintainers are adamant that packages not be called anything so mundane as "utils"
package okinawa

import (
	"fmt"
	"strings"
)

// A Ramen is a pretty printer. An empty Ramen is ready to use, though won't print anything
// particularly interesting. Set the Name, call Append() repeatedly, and finally call String()
// to pretty print.
type Ramen struct {
	// The Name is used in the pretty printed struct name.
	Name string

	ingredients []ingredient
}

type ingredient struct {
	Name  string
	Value interface{}
}

// Append adds the key-value pair to the Ramen.
func (r *Ramen) Append(name string, value interface{}) {
	i := ingredient{name, value}
	r.ingredients = append(r.ingredients, i)
}

// String implements the fmt.Stringer interface.
func (r *Ramen) String() string {
	b := strings.Builder{}

	appendFormatted := func(i ingredient) {
		s := fmt.Sprintf("\t%s: %v\n", i.Name, i.Value)
		b.WriteString(s)
	}

	opening := fmt.Sprintf("%s{\n", r.Name)
	b.WriteString(opening)
	for _, v := range r.ingredients {
		appendFormatted(v)
	}
	b.WriteString("}")
	return b.String()
}
