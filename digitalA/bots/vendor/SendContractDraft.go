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

func SendContractDraf(sender string) {
	fmt.Println("sending file")
	content, err := ioutil.ReadFile("./documents/ContractDraft.docx")
	//content, err := ioutil.ReadFile("./bots/documents/Kottouf.png")
	if err != nil {
		fmt.Println(err)
	}

	resp, err := global.Cli.Upload(context.Background(), content, whatsmeow.MediaDocument)
	if err != nil {
		fmt.Println(err)
	}

	msg := &waProto.DocumentMessage{
		//	Title:    proto.String("Kottouf Contract Terms"),
		FileName:      proto.String("مسودة العقد الرسمي"),
		Mimetype:      proto.String("application/vnd.openxmlformats-officedocument.wordprocessingml.document"), // replace this with the actual mime type
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
}
