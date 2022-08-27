package vendor

import (
	"context"
	"fmt"
	"io/ioutil"

	"DigitalAssistance/global"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func SendVAT(sender string) {
	fmt.Println("sending VAT file")
	content, err := ioutil.ReadFile("./documents/VATcertificate.pdf")
	//content, err := ioutil.ReadFile("./bots/documents/Kottouf.png")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("File contents: %s", content)
	// resp, err := global.Cli.Upload(context.Background(), content, whatsmeow.MediaImage)
	resp, err := global.Cli.Upload(context.Background(), content, whatsmeow.MediaDocument)
	if err != nil {
		fmt.Println(err)
	}
	// handle error

	/*	msg := &waProto.ImageMessage{
			Caption:  proto.String("Hello, world!"),
			Mimetype: proto.String("image/png"), // replace this with the actual mime type
			// you can also optionally add other fields like ContextInfo and JpegThumbnail here

			Url:           &resp.URL,
			DirectPath:    &resp.DirectPath,
			MediaKey:      resp.MediaKey,
			FileEncSha256: resp.FileEncSHA256,
			FileSha256:    resp.FileSHA256,
			FileLength:    &resp.FileLength,
		}
	*/

	msg := &waProto.DocumentMessage{
		//	Title:         proto.String("الشهادة الضريبية VAT"),
		FileName:      proto.String("الشهادة الضريبية VAT"),
		Mimetype:      proto.String("application/pdf"), // replace this with the actual mime type
		Url:           &resp.URL,
		DirectPath:    &resp.DirectPath,
		MediaKey:      resp.MediaKey,
		FileEncSha256: resp.FileEncSHA256,
		FileSha256:    resp.FileSHA256,
		FileLength:    &resp.FileLength,
	}

	targetJID, ok := global.ParseJID(sender)
	if !ok {
		return
	}
	//	send, err := global.Cli.SendMessage(jid, "", msg) // jid = recipient
	send, err := global.Cli.SendMessage(targetJID, "", &waProto.Message{
		DocumentMessage: msg,
		//ImageMessage: msg,
	})

	if err != nil {
		global.Log.Errorf("Error sending message: %v", err)
	} else {
		global.Log.Infof("Message sent (server timestamp: %s)", send)
	}

	fmt.Println("sending CR file")
	content2, err := ioutil.ReadFile("./documents/CR.pdf")
	//content, err := ioutil.ReadFile("./bots/documents/Kottouf.png")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("File contents: %s", content)
	// resp, err := global.Cli.Upload(context.Background(), content, whatsmeow.MediaImage)
	resp2, err := global.Cli.Upload(context.Background(), content2, whatsmeow.MediaDocument)
	if err != nil {
		fmt.Println(err)
	}
	msg2 := &waProto.DocumentMessage{
		//	Title:         proto.String("الشهادة الضريبية VAT"),
		FileName:      proto.String("السجل التجاري"),
		Mimetype:      proto.String("application/pdf"), // replace this with the actual mime type
		Url:           &resp2.URL,
		DirectPath:    &resp2.DirectPath,
		MediaKey:      resp2.MediaKey,
		FileEncSha256: resp2.FileEncSHA256,
		FileSha256:    resp2.FileSHA256,
		FileLength:    &resp2.FileLength,
	}

	//	send, err := global.Cli.SendMessage(jid, "", msg) // jid = recipient
	send2, err := global.Cli.SendMessage(targetJID, "", &waProto.Message{
		DocumentMessage: msg2,
		//ImageMessage: msg,
	})

	if err != nil {
		global.Log.Errorf("Error sending message: %v", err)
	} else {
		global.Log.Infof("Message sent (server timestamp: %s)", send2)
	}
}
