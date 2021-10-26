package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type templateData struct {
	String          map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated int
	API             string
	CSSVersion      string
}

var functions = template.FuncMap{}

// this will compile all files in "templates" directory to the final binary
//go:embed templates
var templateFS embed.FS

// Not implemented yet
func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	return td
}

// renderTemplate renders a page stored in a file called "templates/#{page}.page.tmpl". After parsing the template file,
// it writes (t.Execute) the output to the "w http.ResponseWriter" parameter.
// For performance reasons, renderTemplate keeps a local cache of rendered pages for production use.
// It accepts templateData pointer to pass additional values into the template, as well as "partials".
// It returns a potential error.
func (app *application) renderTemplate(
	w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error
	templateToRender := fmt.Sprintf("templates/%s.page.gohtml", page)

	_, templateInMap := app.templateCache[templateToRender]

	if app.config.env == "production" && templateInMap {
		t = app.templateCache[templateToRender]
	} else {
		t, err = app.parseTemplate(partials, page, templateToRender)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}
	if td == nil {
		td = &templateData{}
	}
	td = app.addDefaultData(td, r)

	err = t.Execute(w, td)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}
	return nil
}

// parseTemplate does the parsing of a file (page) specified by a base name of a Go template file, it must be stored
// inside the "web/templates" directory.
// It returns a *template.Template object with the parsed file.
func (app *application) parseTemplate(
	partials []string, page string, templateToRender string,
) (*template.Template, error) {
	var t *template.Template
	var err error

	// build partials
	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partial.gohtml", x)
		}
	}

	if len(partials) > 0 {
		t, err = template.New(
			fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(
			templateFS, "templates/base.layout.gohtml",
			strings.Join(partials, ","), templateToRender,
		)
	} else {
		t, err = template.New(
			fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(
			templateFS, "templates/base.layout.gohtml", templateToRender,
		)
	}
	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}
	app.templateCache[templateToRender] = t
	return t, nil
}
