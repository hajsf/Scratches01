package vendor

import (
	"DigitalAssistance/global"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func SendContact(sender string) {
	displayName := "مشتريات قطوف و حلى"
	tel := "966138117381"

	msg := waProto.ContactMessage{
		DisplayName: proto.String(displayName),
		Vcard:       proto.String(createVCard(displayName, tel)),
	}

	m := &waProto.Message{ContactMessage: &msg}

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
