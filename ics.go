package goics

import (
	"errors"
	"io"
)

/*
ICS .ics file struct
 */
type ICS struct {
	RootComponent *Component
}

/*
Component a component .ics file contains
 */
type Component struct {
	Name          string
	Values        map[string]string
	SubComponents []*Component
}

func (ics *ICS) Write(w io.Writer) error {

	if ics.RootComponent.Name != VCALENDAR {
		return errors.New("Invalid root ics component: " + ics.RootComponent.Name)
	}

	writeComponent(w, ics.RootComponent)

	return nil
}

func writeComponent(w io.Writer, c *Component) {

	if c == nil {
		return
	}

	writeln(w, BEGIN + c.Name)
	defer func() {
		writeln(w, END + c.Name)
	}()

	for k, v := range c.Values {
		writeln(w, k + v)
	}

	for _, subc := range c.SubComponents {
		writeComponent(w, subc)
	}
}

func writeln(w io.Writer, line string) error {
	_, err := w.Write([]byte(line + "\n"))
	return err
}