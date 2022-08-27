package branches

import (
	"context"
	"fmt"
	"io/ioutil"

	"DigitalAssistance/global"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func SendCard(sender, barcode string) {
	card, ipart, dpart := GetCard(barcode)
	//	fmt.Println(card, "\n", ipart, "\n", dpart)
	if len(card.BarCode) == 0 {
		msg := &waProto.Message{
			Conversation: proto.String("نعتذر كرت الصنف غير متوفر حاليا")}

		jid, ok := global.ParseJID(sender)
		if !ok {
			return
		}
		send, err := global.Cli.SendMessage(jid, "", msg) // jid = recipient

		if err != nil {
			global.Log.Errorf("Error sending message: %v", err)
		} else {
			global.Log.Infof("Message sent (server timestamp: %s)", send)
		}
	} else {
		fmt.Println("******")
		fmt.Println(card.ReturnPolicy)
		r := NewRequestPdf("")
		//html template path
		templatePath := "./templates/card.html"
		//path for download pdf
		outputPath := "./storage/" + barcode + "Card.pdf"
		//html template data
		templateData := struct {
			Card         global.SKUcard[float64]
			Ipart, Dpart string
		}{
			Card:  card,
			Ipart: ipart,
			Dpart: dpart,
		}
		fmt.Println(templateData.Card)
		if err := r.ParseTemplate(templatePath, templateData); err == nil {
			ok, err := r.GeneratePDF(outputPath)
			if err != nil {
				global.Log.Errorf("Error generating pdf file: %v", err)
			} else {
				fmt.Println(ok, "pdf generated successfully")
			}
		} else {
			global.Log.Errorf("Error parsing template: %v", err)
		}
		fmt.Println("sending file")
		content, err := ioutil.ReadFile(outputPath)
		if err != nil {
			global.Log.Errorf("Error reading file: %v", err)
		}
		resp, err := global.Cli.Upload(context.Background(), content, whatsmeow.MediaDocument)
		if err != nil {
			global.Log.Errorf("Error uploading document to whatsapp server: %v", err)
		}

		msg := &waProto.DocumentMessage{
			FileName:      proto.String(fmt.Sprintf("SKU card: %v", barcode)),
			Mimetype:      proto.String("application/pdf"),
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
		send, err := global.Cli.SendMessage(targetJID, "", &waProto.Message{
			DocumentMessage: msg,
		})

		if err != nil {
			global.Log.Errorf("Error sending message: %v", err)
		} else {
			global.Log.Infof("Message sent (server timestamp: %s)", send)
		}
	}
}
