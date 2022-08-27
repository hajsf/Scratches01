package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/protobuf/proto"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())
	}
}

var client *whatsmeow.Client
var err error
var passer *DataPasser

const maxClients = 1

func init() {
	store.DeviceProps.Os = proto.String("WhatsApp GO")
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New("sqlite3", "file:datastore.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client = whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(eventHandler)

	passer = &DataPasser{
		logs: make(chan string),
		sema: make(chan struct{}, maxClients),
	}
}

//go:embed static
var embededFiles embed.FS

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	fsys, err := fs.Sub(embededFiles, "static")
	if err != nil {
		panic(err)
	}
	http.Handle("/", http.FileServer(http.FS(fsys)))
	http.HandleFunc("/upload", upload)

	http.HandleFunc("/sse/dashboard", passer.handleHello)
	go http.ListenAndServe(":8888", nil)
	/*
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		done := make(chan bool)
		go func() {
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					//	fmt.Println("Tick at", t)
					// passer.logs <- buffer.String()
				}
			}
		}()
	*/

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Println("shutting down")
	passer.logs <- "panic"
	fmt.Println("byw")
	if client.IsConnected() {
		client.Disconnect()
	}

}
