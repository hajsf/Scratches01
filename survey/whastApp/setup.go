package whastapp

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	waBinary "go.mau.fi/whatsmeow/binary"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

func setup() store.Device {
	fmt.Println("WhatsApp")
	//Create a folder/directory at a full qualified path
	err := os.MkdirAll("./temp", 0755)
	if err != nil {
		fmt.Println(err)
	}
	store.CompanionProps.Os = proto.String("Kottouf") // App name to appear at WhatsApp linked devices
	waBinary.IndentXML = true

	dbLog := waLog.Stdout("Database", "DEBUG", true)
	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New("sqlite3", "file:examplestore.db?_foreign_keys=on", dbLog)
	//	container, err := sqlstore.New(*dbDialect, *dbAddress, dbLog)
	if err != nil {
		panic(err)
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	device, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}

	return *device
}
