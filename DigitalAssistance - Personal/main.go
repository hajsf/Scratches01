package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"HasanAssistance/global"
	"HasanAssistance/handlers"
	"HasanAssistance/structs"

	"github.com/BurntSushi/toml"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal/v3"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"google.golang.org/protobuf/proto"

	"go.mau.fi/whatsmeow"
	waBinary "go.mau.fi/whatsmeow/binary"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var logLevel = "INFO"
var debugLogs = flag.Bool("debug", false, "Enable debug logs?")
var dbDialect = flag.String("db-dialect", "sqlite3", "Database dialect (sqlite3 or postgres)")
var dbAddress = flag.String("db-address", "file:mdtest.db?_foreign_keys=on", "Database address")

func init() {
	/* Set language translation */

	// Create a new i18n bundle with default language.
	global.Bundle = i18n.NewBundle(language.English)

	// Register a toml unmarshal function for i18n bundle.
	global.Bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// Load translations from toml files for non-default languages.
	global.Bundle.MustLoadMessageFile("./lang/active.ar.toml")
	global.Bundle.MustLoadMessageFile("./lang/active.es.toml")

}
func main() {

	//	pdf.Generator()
	//Create a folder/directory at a full qualified path
	err := os.MkdirAll("./temp", 0755)
	if err != nil {
		global.Log.Errorf(fmt.Sprint(err))
	}
	store.CompanionProps.Os = proto.String("DigitalAssistance") // App name to appear at WhatsApp linked devices
	waBinary.IndentXML = true
	flag.Parse()

	if *debugLogs {
		logLevel = "DEBUG"
	}
	global.Log = waLog.Stdout("Main", logLevel, true)

	dbLog := waLog.Stdout("Database", logLevel, true)
	storeContainer, err := sqlstore.New(*dbDialect, *dbAddress, dbLog)
	if err != nil {
		global.Log.Errorf("Failed to connect to database: %v", err)
		return
	}
	device, err := storeContainer.GetFirstDevice()
	if err != nil {
		global.Log.Errorf("Failed to get device: %v", err)
		return
	}
	global.Db, err = sql.Open("sqlite3", *global.DbMessages)
	if err != nil {
		log.Fatal(err)
	}
	defer global.Db.Close()

	// auto rowid is created, use WITHOUT ROWID if not required, and use select rowid or oid to see it
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS locations (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, name text, jid text,
		longitude real, latitude real, address text,
		created DATETIME DEFAULT (STRFTIME('%d-%m-%Y', 'NOW','localtime')),
		created_at DATETIME DEFAULT (STRFTIME('%d-%m-%Y   %H:%M', 'NOW','localtime')));
	`
	_, err = global.Db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	var response structs.LoginResponse
	global.Cli = whatsmeow.NewClient(device, nil) // cli = whatsmeow.NewClient(device, waLog.Stdout("Client", logLevel, true))
	ch, err := global.Cli.GetQRChannel(context.Background())
	if err != nil {
		// This error means that we're already logged in, so ignore it.
		if !errors.Is(err, whatsmeow.ErrQRStoreContainsID) {
			global.Log.Errorf("Failed to get QR channel: %v", err)
		}
	} else {
		go func() {
			for evt := range ch {
				response.Code = evt.Code
				response.Duration = evt.Timeout / time.Second / 2
				if evt.Event == "code" {
					qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				} else {
					global.Log.Infof("QR channel result: %s", evt.Event)
				}
			}
		}()
	}

	global.Cli.AddEventHandler(handlers.Handler)
	err = global.Cli.Connect()
	if err != nil {
		global.Log.Errorf("Failed to connect: %v", err)
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	global.Log.Infof("Interrupt received, exiting")
	global.Cli.Disconnect()
}
