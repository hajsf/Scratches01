package check

import (
	"DigitalAssistance/global"
	"context"
	"fmt"
	"io/ioutil"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func SendDisposable(sender string) {

	content, err := ioutil.ReadFile("./documents/Eid.png")
	if err != nil {
		fmt.Println(err)
	}
	resp, err := global.Cli.Upload(context.Background(), content, whatsmeow.MediaImage)
	if err != nil {
		fmt.Println(err)
	}

	msg := waProto.ImageMessage{
		Caption:  proto.String("شكرا و عيدكم مبارك يمكنك الآن الحصول على *خصم 10%* من منتجات قطوف في زيارتك القادمة"),
		Mimetype: proto.String("image/png"), // replace this with the actual mime type
		// you can also optionally add other fields like ContextInfo and JpegThumbnail here

		Url:           &resp.URL,
		DirectPath:    &resp.DirectPath,
		MediaKey:      resp.MediaKey,
		FileEncSha256: resp.FileEncSHA256,
		FileSha256:    resp.FileSHA256,
		FileLength:    &resp.FileLength,
		ViewOnce:      proto.Bool(true),
	}

	m := &waProto.Message{ImageMessage: &msg}

	jid, ok := global.ParseJID(sender)
	if !ok {
		return
	}
	send, err := global.Cli.SendMessage(jid, "", m) // jid = recipient

	if err != nil {
		global.Log.Errorf("Error sending message: %v", err)
	} else {
		global.Log.Infof("Message sent (server timestamp: %s)", send)
	}

}
