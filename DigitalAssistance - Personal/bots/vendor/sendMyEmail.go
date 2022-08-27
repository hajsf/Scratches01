package vendor

import (
	"DigitalAssistance/global"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func SendMyEmail(sender string) {
	//	displayName := "مدير مشتريات قطوف و حلى"
	//	tel := "966506889946"
	email := "hasan.y@kottouf.net"

	/*	msg := waProto.ContactMessage{
			DisplayName: proto.String(displayName),
			Vcard:       proto.String(createVCard(displayName, tel)),
		}
	*/
	msg := waProto.ExtendedTextMessage{
		Title: proto.String("إيميل مدير المشتريات"),
		Text:  proto.String(email),
		//	CanonicalUrl: proto.String("https://forms.gle/7Qxvr5RUzcL8KpWT7"),
		// MatchedText: proto.String("https://forms.gle/7Qxvr5RUzcL8KpWT7"),
		//		JpegThumbnail: thumb,
		//	Description: proto.String("https://forms.gle/7Qxvr5RUzcL8KpWT7"),
	}

	m := &waProto.Message{ExtendedTextMessage: &msg}

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
