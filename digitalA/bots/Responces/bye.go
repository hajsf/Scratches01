package responces

import (
	"DigitalAssistance/global"
	"context"
	"fmt"
	"io/ioutil"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func Bye(sender string) {

	content, err := ioutil.ReadFile("./documents/kottouf.png")
	if err != nil {
		fmt.Println(err)
	}
	resp, err := global.Cli.Upload(context.Background(), content, whatsmeow.MediaImage)
	if err != nil {
		fmt.Println(err)
	}
	/*
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

		imageMessage := waProto.ImageMessage{
		//	Caption:  proto.String("شكرا و عيدكم مبارك"),
			Mimetype: proto.String("image/png"), // replace this with the actual mime type
			// you can also optionally add other fields like ContextInfo and JpegThumbnail here

			Url:           &resp.URL,
			DirectPath:    &resp.DirectPath,
			MediaKey:      resp.MediaKey,
			FileEncSha256: resp.FileEncSHA256,
			FileSha256:    resp.FileSHA256,
			FileLength:    &resp.FileLength,
		}

		hydratedImage := waProto.HydratedFourRowTemplate_ImageMessage{
			ImageMessage: &imageMessage,
		}

		hydratedFourRowTemplate := waProto.HydratedFourRowTemplate{
		//	HydratedContentText: proto.String("الآن العرض الجديد"),
		//	HydratedFooterText:  proto.String("تطبّق الشروط والأحكام"),
			HydratedButtons: []*waProto.HydratedTemplateButton{
				{
					Index:          proto.Uint32(1),
					HydratedButton: &hydratedUrlButton,
				},
				{
					Index:          proto.Uint32(2),
					HydratedButton: &hydratedCallButton,
				},
				{
					Index:          proto.Uint32(3),
					HydratedButton: &hydratedQuickReplyButton,
				},
			},
			TemplateId: proto.String("id1"),
			Title:      &hydratedImage,
		}

		_ = &hydratedImage
		_ = &hydratedUrlButton
		_ = &hydratedQuickReplyButton

		templateMessage := waProto.TemplateMessage{
			// ContextInfo:      &waProto.ContextInfo{},
			HydratedTemplate: &hydratedFourRowTemplate,
			Format: nil, * &waProto.TemplateMessage_HydratedFourRowTemplate{
				HydratedFourRowTemplate: &hydratedFourRowTemplate,
			},
		}

		msg := waProto.Message{
		//	Conversation:    proto.String("Hello"),
			TemplateMessage: &templateMessage,
		} */

	//	msg := waProto.Message{
	//		Conversation: proto.String("*مع السلامه*"),
	//	}

	// Creating template message
	msg := waProto.Message{
		TemplateMessage: &waProto.TemplateMessage{
			HydratedTemplate: &waProto.HydratedFourRowTemplate{
				Title: &waProto.HydratedFourRowTemplate_ImageMessage{
					ImageMessage: &waProto.ImageMessage{
						//	Caption:  proto.String("شكرا و عيدكم مبارك"),
						Mimetype: proto.String("image/png"), // replace this with the actual mime type
						// you can also optionally add other fields like ContextInfo and JpegThumbnail here

						Url:           &resp.URL,
						DirectPath:    &resp.DirectPath,
						MediaKey:      resp.MediaKey,
						FileEncSha256: resp.FileEncSHA256,
						FileSha256:    resp.FileSHA256,
						FileLength:    &resp.FileLength,
						Height:        proto.Uint32(410),
						Width:         proto.Uint32(1200),
					},
				},

				/*	&waProto.HydratedFourRowTemplate_HydratedTitleText{
					HydratedTitleText: "مع السلامة",
				}, */
				TemplateId:          proto.String("template-id"),
				HydratedContentText: proto.String("سيقوم قسم المشتريات بالتواصل معكم قريبا"),
				HydratedFooterText:  proto.String("لطلب منتجات من قطوف و حلا، يمكنكم زيارة السوق الإلكتروني أو التواصل هاتفيا"),
				HydratedButtons: []*waProto.HydratedTemplateButton{

					// This for URL button
					{
						Index: proto.Uint32(1),
						HydratedButton: &waProto.HydratedTemplateButton_UrlButton{
							UrlButton: &waProto.HydratedURLButton{
								DisplayText: proto.String("للتسوق إضغط الزر"),
								Url:         proto.String("https://kottouf.co/"),
							},
						},
					},

					// This for call button
					{
						Index: proto.Uint32(2),
						HydratedButton: &waProto.HydratedTemplateButton_CallButton{
							CallButton: &waProto.HydratedCallButton{
								DisplayText: proto.String("للطلبات إتصل على"),
								PhoneNumber: proto.String("00966559528883"),
							},
						},
					},

					// This is just a quick reply
					/*	{
						Index: proto.Uint32(3),
						HydratedButton: &waProto.HydratedTemplateButton_QuickReplyButton{
							QuickReplyButton: &waProto.HydratedQuickReplyButton{
								DisplayText: proto.String("Quick reply"),
								Id:          proto.String("quick-id"),
							},
						},
					},*/
				},
			},
		},
	}

	// Sending message
	// WaClient.SendMessage(event.Info.Chat, "", this_message)

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
	/*
		send, err = global.Cli.SendMessage(jid, "", &waProto.Message{
			ImageMessage: &imageMessage,
		})

		if err != nil {
			global.Log.Errorf("Error sending message: %v", err)
		} else {
			global.Log.Infof("Message sent (server timestamp: %s)", send)
		}
	*/
}
