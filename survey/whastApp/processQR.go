package whastapp

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"survey/config"
	"survey/stream"
	"survey/structs"
	"time"

	"github.com/skip2/go-qrcode"
	"go.mau.fi/whatsmeow"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func processQR() (response structs.LoginResponse, err error) {
	deviceStore := setup()
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(&deviceStore, clientLog)
	// create the image stream
	qr := stream.NewStream()
	content, _ := ioutil.ReadFile("./bots/documents/ContractTerms.pdf")

	http.Handle("/qr", qr)
	chImage := make(chan string)

	if client.Store.ID == nil {
		// No ID stored, new login
		//	global.LogginStatus <- global.NewLoggin

		qrChan, err := client.GetQRChannel(context.Background())
		if err != nil {
			// This error means that we're already logged in, so ignore it.
			if errors.Is(err, whatsmeow.ErrQRStoreContainsID) {
				_ = client.Connect() // just connect to websocket
				if client.IsLoggedIn() {
					return response, errors.New("you already logged in :)")
				}
				return response, errors.New("your session have been saved, please wait to connect 2 second and refresh again")
			} else {
				return response, errors.New("Error when GetQRChannel:" + err.Error())
			}
		}
		err = client.Connect()
		if err != nil {
			panic(err)
		}

		for evt := range qrChan {
			//	response.Code = evt.Code
			duration := evt.Timeout / time.Second // / 2
			qrPath := fmt.Sprintf("%s/scan-qr.png", config.PathQrCode)
			if evt.Event == "code" {
				go func() {
					fmt.Println("********************")
					fmt.Println("new QR received")
					fmt.Println("********************")
					// Render the QR code here
					// e.g. qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
					// or just manually `echo 2@... | qrencode -t ansiutf8` in a terminal
					fmt.Println("QR code:", evt.Code)
					//qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
					/**/
					//	"http://localhost:8080/static/" or "./static/qrcode"

					err = qrcode.WriteFile(evt.Code, qrcode.Medium, 512, qrPath)
					if err != nil {
						fmt.Println("error when write qrImage file", err.Error())
					}

					qr.UpdateIMG(content)
					chImage <- qrPath
					fmt.Println(qrPath)
					//		global.LogginStatus <- global.QRSaved
				}()
				go func() {
					time.Sleep(duration * time.Second)
					err := os.Remove(qrPath)
					if err != nil {
						fmt.Println("Failed to remove qrPath " + qrPath)
					}
					//			global.LogginStatus <- global.GetNewQR
					fmt.Println("********************")
					fmt.Println("QR deleted")
					fmt.Println("********************")

				}()

				/**/
			} else {
				fmt.Println("Login event:", evt.Event)
				fmt.Println("********************")
				fmt.Println("signed in")
				fmt.Println("********************")
			}
		}
	} else {
		// Already logged in, just connect
		//	global.LogginStatus <- global.LoggedIn
		err := client.Connect()
		if err != nil {
			return response, errors.New("Failed to connect bro " + err.Error())
		}
	}
	response.ImagePath = <-chImage
	response.Client = *client
	return response, nil
}
