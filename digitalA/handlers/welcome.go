package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"DigitalAssistance/Enum"
	"DigitalAssistance/global"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func WelcomeMessage(sender, name string) {

	// Create a new localizer.
	localizer := i18n.NewLocalizer(global.Bundle, global.Lang)
	// Set title message.
	helloPerson := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "HelloPerson", // set translation ID
			Other: "Hello *{{.Name}}* \n" +
				"This is the digital assistant of Kottouf's procurement manager", // set default translation
		},
		TemplateData: map[string]string{
			"Name": name,
		},
		PluralCount: nil,
	})

	whoIsThis := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "WhoIsThis",    // set translation ID
			Other: "Who are you?", // set default translation
		},
		TemplateData: map[string]string{
			"Name": name,
		},
	})

	fmt.Println(helloPerson)

	var content strings.Builder
	//	content.WriteString(fmt.Sprintf("مرحبا *%v* \n", name))
	content.WriteString(helloPerson)
	//content.WriteString("\n")
	//content.WriteString("معك المساعد الرقمي لمدير المشتريات في شركة قطوف و حلا")
	msg := &waProto.Message{ButtonsMessage: &waProto.ButtonsMessage{
		HeaderType:  waProto.ButtonsMessage_EMPTY.Enum(),
		ContentText: proto.String(content.String()),
		FooterText:  proto.String(whoIsThis),
		Buttons: []*waProto.Button{
			{
				ButtonId: proto.String(strconv.Itoa(Enum.Vendor)),
				ButtonText: &waProto.ButtonText{
					DisplayText: proto.String("مورد"),
				},
				Type: waProto.Button_RESPONSE.Enum(),
			},
			{
				ButtonId: proto.String(strconv.Itoa(Enum.Supervisor)),
				ButtonText: &waProto.ButtonText{
					DisplayText: proto.String("مشرف فرع"),
				},
				Type: waProto.Button_RESPONSE.Enum(),
			},
			/*	{
				ButtonId: proto.String(strconv.Itoa(Enum.Franchisee)),
				ButtonText: &waProto.ButtonText{
					DisplayText: proto.String("حامل إمتياز"),
				},
				Type: waProto.Button_RESPONSE.Enum(),
			}, */
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
}
