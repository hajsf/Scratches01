package responces

import (
	"strconv"

	"DigitalAssistance/Enum"
	"DigitalAssistance/global"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func AnotherQuestion(sender string) {
	msg := &waProto.Message{ButtonsMessage: &waProto.ButtonsMessage{
		HeaderType:  waProto.ButtonsMessage_EMPTY.Enum(),
		ContentText: proto.String("هل أنت بحاجه إلى مساعده أخرى"),
		Buttons: []*waProto.Button{
			{
				ButtonId: proto.String(strconv.Itoa(Enum.Yes)),
				ButtonText: &waProto.ButtonText{
					DisplayText: proto.String("نعم"),
				},
				Type: waProto.Button_RESPONSE.Enum(),
			},
			{
				ButtonId: proto.String(strconv.Itoa(Enum.No)),
				ButtonText: &waProto.ButtonText{
					DisplayText: proto.String("لا"),
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
