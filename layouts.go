package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	layoutDir   string = "layouts/"
	templateDir string = "layouts/"
	templateExt string = ".gohtml"
)

func layoutFiles() []string {
	files, err := filepath.Glob(layoutDir + "*" + templateExt)
	if err != nil {
		panic(err)
	}
	return files
}

func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = templateDir + f
	}
}

func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + templateExt
	}
}

// Page is a controller that builds the layouts together into a View.
type Page struct {
	Home *View
}

// NewPage generates the view based on the arguments provided to NewView.
func newPage() *Page {
	return &Page{
		Home: newView(
			"bootstrap", "home"),
	}
}

// View type defines a layout and a template
type View struct {
	Template *template.Template
	Layout   string
}

// Render let's us affirm the content-type and assists with wrapping the layouts together.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

// NewView includes all additional layout files for the site
func newView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}
