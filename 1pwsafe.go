package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

var (
	version  string
	fVersion = flag.Bool("v", false, "print program version")
)

func main() {
	flag.Parse()
	if *fVersion {
		fmt.Fprintf(os.Stdout, "1pwsafe %s\n", version)
		os.Exit(0)
	}

	w := csv.NewWriter(os.Stdout)
	r := csv.NewReader(os.Stdin)
	r.Comma = '\t'

	// Skip header.
	r.Read()

	// Write the header.
	columns := []string{
		"title",
		"URL",
		"username",
		"password",
		"notes",
	}
	w.Write(columns)
	w.Flush()

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		entry := newEntry(record)
		w.Write(entry.Record())
		w.Flush()
	}

	err := w.Error()
	if err != nil {
		log.Fatal(err)
	}
}

type entry struct {
	title                string
	username             string
	password             string
	url                  string
	createdTime          string
	passwordModifiedTime string
	recordModifiedTime   string
	passwordPolicy       string
	passwordPolicyName   string
	history              string
	email                string
	symbols              string
	notes                string
}

func newEntry(record []string) *entry {
	return &entry{
		title:                record[0],
		username:             record[1],
		password:             record[2],
		url:                  record[3],
		createdTime:          record[4],
		passwordModifiedTime: record[5],
		recordModifiedTime:   record[6],
		passwordPolicy:       record[7],
		passwordPolicyName:   record[8],
		history:              record[9],
		email:                record[10],
		symbols:              record[11],
		notes:                record[12],
	}
}

func (e *entry) Title() string {
	title := e.title
	fields := strings.SplitN(e.title, ".", 2)
	if len(fields) == 2 {
		title = fields[1]
	}
	return strings.Replace(title, "»", ".", -1)
}

func (e *entry) Username() string {
	username := e.username
	if username == "" {
		username = e.email
	}
	return username
}

func (e *entry) URL() string {
	return e.url
}

func (e *entry) Password() string {
	return e.password
}

func (e *entry) Email() string {
	return e.email
}

func (e *entry) Notes() string {
	var buf bytes.Buffer
	tNotes.Execute(&buf, e)
	return buf.String()
}

func (e *entry) Record() []string {
	return []string{
		e.Title(),
		e.URL(),
		e.Username(),
		e.Password(),
		e.Notes(),
	}
}

var tNotes = template.Must(
	template.New("notes").
		Parse(`Notes:
Title: {{.Title}}
Username: {{.Username}}
URL: {{.URL}}
Email: {{.Email}}
`))
