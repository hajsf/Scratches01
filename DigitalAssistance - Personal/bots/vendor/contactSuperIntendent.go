package vendor

import (
	"DigitalAssistance/global"
	"fmt"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func ContactSuperIntendent(sender string) {
	msg := &waProto.Message{
		Conversation: proto.String("سيقوم مدير المشتريات بالتواصل معك قريبا، أو يمكنك التواصل مع مشرف القسم")}

	jid, ok := global.ParseJID(sender)
	if !ok {
		fmt.Println(ok, "can not Parse:", sender)
		return
	}
	send, err := global.Cli.SendMessage(jid, "", msg) // jid = recipient

	if err != nil {
		global.Log.Errorf("Error sending message: %v", err)
	} else {
		global.Log.Infof("Message sent (server timestamp: %s)", send)
	}
}
