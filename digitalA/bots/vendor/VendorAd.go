package vendor

import (
	"DigitalAssistance/global"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func VendorAd(sender string) {
	//	msg := &waProto.Message{
	//		Conversation: proto.String("*مع السلامه*")}

	msg := &waProto.Message{
		ExtendedTextMessage: &waProto.ExtendedTextMessage{
			Title: proto.String("توثيق حملة دعائية و تسويقية"),
			Text:  proto.String("https://forms.gle/dFv4b3dGQ7YuTN7Q9"),
			//	CanonicalUrl: proto.String("https://forms.gle/7Qxvr5RUzcL8KpWT7"),
			// MatchedText: proto.String("https://forms.gle/7Qxvr5RUzcL8KpWT7"),
			//		JpegThumbnail: thumb,
			//	Description: proto.String("https://forms.gle/7Qxvr5RUzcL8KpWT7"),
		},
	}

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
}
