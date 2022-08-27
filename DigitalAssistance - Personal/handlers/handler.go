package handlers

import (
	"HasanAssistance/global"
	"fmt"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

func Handler(rawEvt interface{}) {

	switch evt := rawEvt.(type) {
	case *events.Message:
		sender := evt.Info.Chat.User

		if sender == "966138117381" {
			fmt.Println(evt.Message)
			fmt.Println("**********")
			//		var x, y, z uint32 = 1, 2, 3
			/*
				templateButtons := []*waProto.HydratedTemplateButton{
					{Index: &x,
						HydratedButton: &waProto.HydratedTemplateButton_UrlButton{
							UrlButton: &waProto.HydratedURLButton{
								DisplayText: proto.String("⭐ Star Baileys on GitHub!"),
								Url:         proto.String("https://github.com/adiwajshing/Baileys"),
							},
						},
					},
					{Index: &y,
						HydratedButton: &waProto.HydratedTemplateButton_CallButton{
							CallButton: &waProto.HydratedCallButton{
								DisplayText: proto.String("call me"),
								PhoneNumber: proto.String("+1 123 467"),
							},
						},
					},
					{Index: &z,
						HydratedButton: &waProto.HydratedTemplateButton_QuickReplyButton{
							QuickReplyButton: &waProto.HydratedQuickReplyButton{
								DisplayText: proto.String("call me"),
								Id:          proto.String("test"),
							},
						},
					},
				}

				m2 := &waProto.Message{
					Conversation: new(string),

					TemplateMessage: &waProto.TemplateMessage{
						ContextInfo:      &waProto.ContextInfo{},
						HydratedTemplate: &waProto.HydratedFourRowTemplate{},
						Format:           nil,
					},
				}

			*/
			/*	buttonMessage := &waProto.Button{
				ButtonId: "Hi it's a template message",
				footer: 'Hello World',
				templateButtons: templateButttons,
				image: { url: 'https://example.com/image.jpeg' } //it can be array { url: 'https://example.com/image.jpeg' } or buffer
			} */

			/*  {index: &x,
				HydratedButton: &waProto.HydratedTemplateButton_CallButton{
					UrlButton: &waProto.HydratedURLButton{
						DisplayText: proto.String("Call me!"),
						:         proto.String("+1 (234) 5678-901"),
					},
			  {index: &x,
				HydratedButton: &waProto.HydratedTemplateButton_UrlButton{
					UrlButton: &waProto.HydratedURLButton{
						DisplayText: proto.String("⭐ Star Baileys on GitHub!"),
						Url:         proto.String("https://github.com/adiwajshing/Baileys"),
					},

				 quickReplyButton: {displayText: 'This is a reply, just like normal buttons!', id: 'id-like-buttons-message'}},
						}
			*/

			hydratedUrlButton := waProto.HydratedTemplateButton_UrlButton{
				UrlButton: &waProto.HydratedURLButton{
					DisplayText: proto.String("زر الرابط"),
					Url:         proto.String("https://www.bankaljazira.com/FFL"),
				},
			}

			hydratedCallButton := waProto.HydratedTemplateButton_CallButton{
				CallButton: &waProto.HydratedCallButton{
					DisplayText: proto.String("Call me"),
					PhoneNumber: proto.String("966138117381"),
				},
			}

			hydratedQuickReplyButton := waProto.HydratedTemplateButton_QuickReplyButton{
				QuickReplyButton: &waProto.HydratedQuickReplyButton{
					DisplayText: proto.String("Quick Reply 1"),
					Id:          proto.String("QuickReply1"),
				},
			}

			imageMessage := waProto.ImageMessage{
				Url:      proto.String("https://mmg.whatsapp.net/d/f/AlAF9hFHR7336ScKIqHneB2NpVYsiKWsT05RKTdEUSbl.enc "),
				Mimetype: proto.String("image/jpeg"),
				Caption:  proto.String("الآن بطاقة العملات مجاناً مدى الحياة تقدّم بالحصول عليها من خلال القنوات الالكترونية لبنك الجزيرة أو عبر الفرع"),
			}

			hydratedImage := waProto.HydratedFourRowTemplate_ImageMessage{
				ImageMessage: &imageMessage,
			}

			hydratedFourRowTemplate := waProto.HydratedFourRowTemplate{
				HydratedContentText: proto.String("الآن العرض الجديد"),
				HydratedFooterText:  proto.String("تطبّق الشروط والأحكام"),
				HydratedButtons: []*waProto.HydratedTemplateButton{
					{
						Index:          proto.Uint32(1),
						HydratedButton: &hydratedUrlButton,
					},
					{
						Index:          proto.Uint32(uint32(2)),
						HydratedButton: &hydratedCallButton,
					},
					{
						Index:          proto.Uint32(uint32(3)),
						HydratedButton: &hydratedQuickReplyButton,
					},
				},
				TemplateId: proto.String("id1"),
				Title:      &hydratedImage,
			}

			templateMessage := waProto.TemplateMessage{
				//	ContextInfo:      &waProto.ContextInfo{},
				HydratedTemplate: &hydratedFourRowTemplate,
				Format: &waProto.TemplateMessage_HydratedFourRowTemplate{
					HydratedFourRowTemplate: &hydratedFourRowTemplate,
				},
			}

			msg := waProto.Message{
				//	Conversation:    proto.String("Hello"),
				TemplateMessage: &templateMessage,
			}

			jid, ok := global.ParseJID(sender)
			if !ok {
				return
			}
			send, err := global.Cli.SendMessage(jid, "", &msg) // jid = recipient

			if err != nil {
				global.Log.Errorf("Error sending message: %v", err)
			} else {
				global.Log.Infof("Message sent (server timestamp: %s)", send)
			}
		}
		// End of switch
	}
}
