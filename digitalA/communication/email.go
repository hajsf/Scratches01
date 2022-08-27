package communication

import (
	"DigitalAssistance/global"
	"crypto/tls"
	"fmt"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
	"gopkg.in/mail.v2"
)

func SendEmail(sender, pushname, subject string) {
	user := "Kottouf.Procurement@googlemail.com" // "O365 logging name"
	password := "kvuyhakkojwjwvlx"               //"O365 logging pasword"
	// Generate App password
	// https://myaccount.google.com/u/3/security?rapt=AEjHL4Mi_H31usxNz7LtUye3Ao4XEAdxHYf-YTPJQdqh7bNwllrbfbNKnQT1f3P7Zo9nyLXQHkEff6TG7gnoFOjFUaIf92DvbQ
	// https://myaccount.google.com/u/3/apppasswords?rapt=AEjHL4OMZmrC7jubo9TOKRKEN3nqhgZNXqbkyudhBaum7vl4pqs8jsrev2pGQSjDfhiW3_omqrvSjJ9QIgcLFEgsNgDs2GvEuw
	smtpHost := "smtp.gmail.com" //"smtp.office365.com" // mail.kottouf.sa
	smtpPort := 587              //587 // 465 993 incoming

	d := mail.NewDialer(smtpHost, smtpPort, user, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := mail.NewMessage()
	m.SetHeader("From", "Kottouf.Procurement@gmail.com") //
	m.SetHeader("To", "hasan.y@kottouf.net")
	m.SetHeader("Subject", fmt.Sprintf("%s: from %s - %s", subject, pushname, sender))
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("failed to send email:", err)
		// panic(err)
	}
	fmt.Println("Email sent, pls check")

	/*	msg := &waProto.Message{
			Conversation: proto.String(fmt.Sprintf("Call: %s", sender))}

		jid, ok := global.ParseJID("966138117381")
		if !ok {
			return
		}
		send, err := global.Cli.SendMessage(jid, "", msg) // jid = recipient

		if err != nil {
			global.Log.Errorf("Error sending message: %v", err)
		} else {
			global.Log.Infof("Message sent (server timestamp: %s)", send)
		}
	*/

	// Creating template message
	msg := &waProto.Message{
		TemplateMessage: &waProto.TemplateMessage{
			HydratedTemplate: &waProto.HydratedFourRowTemplate{
				Title: &waProto.HydratedFourRowTemplate_HydratedTitleText{
					HydratedTitleText: "Telephonic call request from:",
				},
				TemplateId:          proto.String("template-id"),
				HydratedContentText: proto.String(pushname),
				HydratedFooterText:  proto.String(""),
				HydratedButtons: []*waProto.HydratedTemplateButton{
					{
						Index: proto.Uint32(1),
						HydratedButton: &waProto.HydratedTemplateButton_CallButton{
							CallButton: &waProto.HydratedCallButton{
								DisplayText: proto.String("call"), // Can not be empty
								PhoneNumber: proto.String("00" + sender),
							},
						},
					},
				},
			},
		},
	}

	jid, ok := global.ParseJID("966138117381")
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
