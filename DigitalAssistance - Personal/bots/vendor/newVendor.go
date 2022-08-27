package vendor

import (
	"strconv"

	"DigitalAssistance/Enum"
	"DigitalAssistance/global"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func NewVender(sender string) {
	msg := &waProto.Message{
		ListMessage: &waProto.ListMessage{
			Description: proto.String("نرحب دائما ببناء شراكات جديدة، يرجى تحديد سبب التواصل"),
			ButtonText:  proto.String("أنقر هنا  👈"),
			ListType:    waProto.ListMessage_SINGLE_SELECT.Enum(),
			Sections: []*waProto.Section{
				{
					Rows: []*waProto.Row{
						{
							RowId: proto.String(strconv.Itoa(Enum.ContractTerms)),
							Title: proto.String("معرفة آلية التعاقد"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.VenderRegistration)),
							Title: proto.String("تسجيل مورد لدى قطوف و حلا"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.ItemRegistration)),
							Title: proto.String("تسجيل صنف لدى قطوف و حلا"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.Location)),
							Title: proto.String("طلب الموقع لإرسال عينات"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.VAT)),
							Title: proto.String("طلب الرقم الضريبي و السجل التجاري"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.CallArrangement)),
							Title: proto.String("ترتيب إتصال مرئي أو هاتفي"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.VisitArrangement)),
							Title: proto.String("جدولة زيارة و لقاء شخصي"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.Email)),
							Title: proto.String("البريد الإلكتروني لمدير القسم"),
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

}
