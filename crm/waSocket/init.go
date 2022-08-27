package main

import (
	"fmt"

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

var err error

const maxClients = 1

type global struct {
	client     whatsmeow.Client
	maxClients int
	parser     DataPasser
}

type DataPasser struct {
	logs chan string
	sema chan struct{} // To control maximum allowed clients connections
}

func main() {
	g := global{
		client:     *whatsmeow.Client,
		maxClients: 1,
		parser: DataPasser{
			logs: make(chan string),
			sema: make(chan struct{}, maxClients),
		},
	}

	g.client
	//	var Client *whatsmeow.Client
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
}
