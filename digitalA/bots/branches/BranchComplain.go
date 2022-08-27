package branches

import (
	"DigitalAssistance/global"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func BranchIssue(sender string) {
	//	msg := &waProto.Message{
	//		Conversation: proto.String("*مع السلامه*")}

/*	msg1 := &waProto.ButtonsMessage{
	TemplateData: waProto.TemplateData {
		{
			index: 1,
			proto.CanonicalUrl
		}
	}	
	}
 */
	

	msg := &waProto.Message{
		ExtendedTextMessage: &waProto.ExtendedTextMessage{
			Title: proto.String("ملاحظات و شكاوى الفروع"),
			Text:  proto.String("https://forms.gle/uUDcxPgVit5nCnVC6"),
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
