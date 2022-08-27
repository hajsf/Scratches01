package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jchv/go-webview2"
)

func main() {

	http.HandleFunc("/", MapPage)
	go http.ListenAndServe(":8080", nil)

	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Window:    nil,
		Debug:     true,
		DataPath:  "",
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title: "Minimal webview example",
			Width: 800, Height: 600, IconId: 2, Center: true,
		},
	})

	w.SetSize(800, 600, webview2.HintMax)

	if w == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer w.Destroy()
	w.SetSize(800, 600, webview2.HintFixed)
	w.Navigate("http://localhost:8080/")
	w.Run()
}

// https://github.com/heremaps/maps-api-for-javascript-examples

func MapPage(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "HERE Map")

	t, err := template.ParseFiles("map.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
