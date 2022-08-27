package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"survey/web"
	whastapp "survey/whastApp"
	"sync"
)

//go:embed static/*.html
var embededFiles embed.FS

// Files started with . or _ are not embeded

var staticServer http.Handler

func init() {
	useOS := len(os.Args) > 1 && os.Args[1] == "live"
	if useOS {
		log.Print("using live mode")
		staticServer = http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	} else {
		log.Print("using embed mode")
		staticServer = http.FileServer(http.FS(embededFiles))
	}
}

func main() {
	http.Handle("/static/", staticServer)
	http.Handle("/qrcode/", http.StripPrefix("/qrcode/", http.FileServer(http.Dir("./qrcode"))))
	if err := web.LoadTemplates(); err != nil {
		// TODO handle error
		fmt.Println(err)
	}
	web.Handlers()

	var wg sync.WaitGroup
	wg.Add(1) // to keep the server goroutine running
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			// TODO handle error
			fmt.Println(err)
		}
	}()

	go whastapp.Define()
	wg.Wait() // to keep the server goroutine running
}
