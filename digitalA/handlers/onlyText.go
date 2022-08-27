package handlers

import (
	"DigitalAssistance/global"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func OnlyText(sender string) {
	// Create a new localizer.
	localizer := i18n.NewLocalizer(global.Bundle, global.Lang)
	// Set title message.
	onlyText := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "OnlyText",                                                                 // set translation ID
			Other: "Sorry, can not handle audio/video/images, kindly send text message only.", // set default translation
		},
		TemplateData: nil,
	})

	msg := &waProto.Message{
		Conversation: proto.String(onlyText)}

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
