package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"DigitalAssistance/Enum"
	"DigitalAssistance/global"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func WelcomeMessageLatin(sender, name string) {
	var content strings.Builder
	content.WriteString(fmt.Sprintf("Hi *%v* \n", name))
	content.WriteString("Here is the Digital assistance of Kottouf's Procurement Manager")
	msg := &waProto.Message{ButtonsMessage: &waProto.ButtonsMessage{
		HeaderType:  waProto.ButtonsMessage_EMPTY.Enum(),
		ContentText: proto.String(content.String()),
		FooterText:  proto.String("Are you existing vendor or new?"),
		Buttons: []*waProto.Button{
			{
				ButtonId: proto.String(strconv.Itoa(Enum.ExistingVendor)),
				ButtonText: &waProto.ButtonText{
					DisplayText: proto.String("Existing"),
				},
				Type: waProto.Button_RESPONSE.Enum(),
			},
			{
				ButtonId: proto.String(strconv.Itoa(Enum.NewVendor)),
				ButtonText: &waProto.ButtonText{
					DisplayText: proto.String("New"),
				},
				Type: waProto.Button_RESPONSE.Enum(),
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
}
