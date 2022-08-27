package vendor

import (
	"DigitalAssistance/global"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func FillForms(sender string) {
	msg := &waProto.Message{
		Conversation: proto.String("يرجى الإطلاع على آلية التعاقد و تعبأة نموذجي تسجيل مورد و تسجيل الصنف، و إرسال عينات، و سيتم التواصل معكم خلال 3 أيام عمل بإذن الله")}

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
