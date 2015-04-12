package goics

import (
	"bytes"
	"testing"
	"log"
)

const nilContent = "<nil>"

func TestSimpleICSWrite(t *testing.T) {

	buf := bytes.NewBuffer([]byte{})

	prodid := "test-prodid"

	rc  := &Component{
		Name: VCALENDAR,
		Values: map[string]string{
			PRODID: prodid,
		},
	}
	ics := &ICS{
		RootComponent: rc,
	}

	ics.Write(buf)

	icsContent := buf.String()

	if icsContent == nilContent {
		t.Errorf("Failed to write ics file")
	}

	log.Printf("buf:\n%s\n", buf.String())

}

func TestComplexICSWrite(t *testing.T) {

	buf := bytes.NewBuffer([]byte{})

	prodid := "test-prodid"

	sc  := &Component{
		Name: EVENT,
		Values: map[string]string{
			ORGANIZER: `MAILTO:mk2@mk2.red`,
		},
	}

	rc  := &Component{
		Name: VCALENDAR,
		Values: map[string]string{
			PRODID: prodid,
		},
		SubComponents: []*Component{GetJapanTimezoneTemplate(), sc},
	}
	ics := &ICS{
		RootComponent: rc,
	}

	ics.Write(buf)

	icsContent := buf.String()

	if icsContent == nilContent {
		t.Errorf("Failed to write ics file")
	}

	log.Printf("buf:\n%s\n", buf.String())

}