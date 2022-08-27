package vendor

import (
	"strconv"

	"DigitalAssistance/Enum"
	"DigitalAssistance/global"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func WhichVender(sender string) {
	msg := &waProto.Message{ButtonsMessage: &waProto.ButtonsMessage{
		HeaderType:  waProto.ButtonsMessage_EMPTY.Enum(),
		ContentText: proto.String("حياك، هل أنت  مورد حالي أم جديد؟"),
		//FooterText:  proto.String("من معي؟"),
		Buttons: []*waProto.Button{
			{
				ButtonId: proto.String(strconv.Itoa(Enum.ExistingVendor)),
				ButtonText: &waProto.ButtonText{
					DisplayText: proto.String("حالي"),
				},
				Type: waProto.Button_RESPONSE.Enum(),
			},
			{
				ButtonId: proto.String(strconv.Itoa(Enum.NewVendor)),
				ButtonText: &waProto.ButtonText{
					DisplayText: proto.String("جديد"),
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
