package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Choose Your Own Adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    {{range .Story}}
        <p>{{.}}</p>
    {{end}}
    <ul>
    {{range .Options}}
        <li><a href="/{{.Arc}}">{{.Text}}</a></li>
    {{end}}
    </ul>
</body>
</html>
`

// NewHandler will return a new http.Handler given a Story
func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

// ServeHTTP ...
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("").Parse(defaultHandlerTmpl))
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		log.Fatalln(err)
	}
}

// JSONStory will return a Story given a io.Reader
func JSONStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

// Story represents a story
type Story map[string]Chapter

// Chapter represents a part of a story.
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

// Option of a section
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
