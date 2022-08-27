package vendor

import (
	"DigitalAssistance/global"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func DutyTiming(sender string) {
	msg := &waProto.Message{
		Conversation: proto.String("مواعيد دوام قسم المشتريات من *الأحد* إلى *الخميس* من ال *9 صباحا* إلى ال *5 مساءا*")}

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
