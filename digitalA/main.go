package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"DigitalAssistance/global"
	"DigitalAssistance/handlers"
	"DigitalAssistance/structs"

	"github.com/BurntSushi/toml"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal/v3"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"google.golang.org/protobuf/proto"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/appstate"
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
	http.HandleFunc("/sse/signal", global.Passer.HandleSignal)

	/* Set language translation */

	// Create a new i18n bundle with default language.
	global.Bundle = i18n.NewBundle(language.English)

	// Register a toml unmarshal function for i18n bundle.
	global.Bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// Load translations from toml files for non-default languages.
	global.Bundle.MustLoadMessageFile("./lang/active.ar.toml")
	global.Bundle.MustLoadMessageFile("./lang/active.es.toml")

}

type Contact struct {
	FirstName, FullName, PushName, BusinessName string
	Found                                       bool
}

func main() {

	//global.ExceptionedNumbers, err := readLines("ExceptionedNumbers.txt")
	/*	lines, err := readLines("foo.in.txt")
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		for i, line := range lines {
			fmt.Println(i, line)
		}

		if err := writeLines(lines, "foo.out.txt"); err != nil {
			log.Fatalf("writeLines: %s", err)
		}
	*/
	go http.ListenAndServe(":1234", nil)
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

	global.Cli.FetchAppState(appstate.WAPatchCriticalUnblockLow, true, false)
	contacts, err := global.Cli.Store.Contacts.GetAllContacts()
	if err != nil {
		fmt.Println("Error reading contacts:", err)
	}
	fmt.Println("You have:", len(contacts), "contacts")

	jid, ok := global.ParseJID("966538382888")
	if !ok {
		fmt.Println("errror")
	}

	fmt.Println(contacts[jid])

	file, err := os.Create("contacts2.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	//	defer file.Close()

	w := csv.NewWriter(file)
	// defer w.Flush()

	// Using Write
	for jid, info := range contacts {
		_ = Contact{
			FirstName:    info.FirstName,
			FullName:     info.FullName,
			PushName:     info.PushName,
			BusinessName: info.BusinessName,
			Found:        info.Found,
		}
		row := []string{
			jid.User,
			info.FirstName,
			info.FullName,
			info.PushName,
			info.BusinessName,
			strconv.FormatBool(info.Found),
		}
		if err := w.Write(row); err != nil {
			log.Fatal("error writing record to file", err)
		}
	}

	// Using WriteAll
	/*	var data [][]string
		for jid, info := range contacts {
			row := []string{
				jid.User,
				info.FirstName,
				info.FullName,
				info.PushName,
				info.BusinessName,
				strconv.FormatBool(info.Found),
			}
			data = append(data, row)
		}
		if err := w.WriteAll(data); err != nil {
			log.Fatal("error writing record to file", err)
		}
	*/

	w.Flush()
	file.Close()
	fmt.Println("************")

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
