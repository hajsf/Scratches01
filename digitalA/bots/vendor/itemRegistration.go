package vendor

import (
	"DigitalAssistance/global"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func ItemRegistration(sender string, chat types.JID) {
	//	msg := &waProto.Message{
	//		Conversation: proto.String("*مع السلامه*")}

	/*	msg := &waProto.Message{
		ExtendedTextMessage: &waProto.ExtendedTextMessage{
			Title: proto.String("نموذج تسجل صنف"),
			Text:  proto.String("https://forms.gle/eV1G8uzM5jbPAoAeA"),
			//	CanonicalUrl: proto.String("https://forms.gle/7Qxvr5RUzcL8KpWT7"),
			// MatchedText: proto.String("https://forms.gle/7Qxvr5RUzcL8KpWT7"),
			//		JpegThumbnail: thumb,
			//	Description: proto.String("https://forms.gle/7Qxvr5RUzcL8KpWT7"),
		},
	} */

	// Creating template message
	msg := &waProto.Message{
		TemplateMessage: &waProto.TemplateMessage{
			HydratedTemplate: &waProto.HydratedFourRowTemplate{
				Title: &waProto.HydratedFourRowTemplate_HydratedTitleText{
					HydratedTitleText: "نموذج تسجيل صنف جديد",
				},
				TemplateId:          proto.String("template-id"),
				HydratedContentText: proto.String(""),
				HydratedFooterText:  proto.String(""),
				HydratedButtons: []*waProto.HydratedTemplateButton{
					{
						Index: proto.Uint32(1),
						HydratedButton: &waProto.HydratedTemplateButton_UrlButton{
							UrlButton: &waProto.HydratedURLButton{
								DisplayText: proto.String("👉 أنقر هنا"),
								Url:         proto.String("https://forms.gle/QL5tA8oWgpPBPYzq9"),
							},
						},
					},
				},
			},
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

	send2, err := global.Cli.SendMessage(chat, "", msg) // jid = recipient

	if err != nil {
		global.Log.Errorf("Error sending message: %v", err)
	} else {
		global.Log.Infof("Message sent (server timestamp: %s)", send2)
	}
}
