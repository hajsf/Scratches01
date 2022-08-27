package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

//go:embed static templates/*
var embededFiles embed.FS

const (
	layoutsDir   = "templates/layouts"
	viewsDir     = "templates/views"
	templatesDir = "templates"
	extension    = "/*.html"
)

func main() {

	useOS := len(os.Args) > 1 && os.Args[1] == "live"
	http.Handle("/", http.FileServer(getFileSystem(useOS)))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(getFileSystem(useOS))))

	// compile single template if required
	// tmpl := template.Must(template.ParseFiles("layout.html"))
	// tmpl := template.Must(template.ParseFS(embededFiles, layoutsDir+"/layout.html"))

	// compile all templates and cache them
	var templates = template.Must(template.ParseFS(embededFiles, // template.Must(template.ParseGlob("YOURTEMPLATEDIR/*"))
		layoutsDir+extension,
		viewsDir+extension))
	http.HandleFunc("/layout", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: true},
				{Title: "Task 2", Done: false},
				{Title: "Task 3", Done: true},
			},
		}
		// tmpl.Execute(w, data)
		err := templates.ExecuteTemplate(w, "indexPage", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8888", nil)
}

// Files started with . or _ are not embeded

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		log.Print("using live mode")
		return http.FS(os.DirFS("static"))
	}

	log.Print("using embed mode, open: http://localhost:8888/")
	fsys, err := fs.Sub(embededFiles, "static")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

// $ go run github.io/hajsf/erp live // live so you can change static files, without live you worl with rmbeded copy
// open: http://localhost:8888/index.html
