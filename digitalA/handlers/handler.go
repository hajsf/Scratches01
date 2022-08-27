package handlers

import (
	responces "DigitalAssistance/bots/Responces"
	"DigitalAssistance/bots/branches"
	"DigitalAssistance/global"
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/abadojack/whatlanggo"
	"go.mau.fi/whatsmeow/types/events"
)

func Handler(rawEvt interface{}) {
	go func() {
		switch evt := rawEvt.(type) {
		case *events.Message:
			//fmt.Println("**********")
			//fmt.Println(evt.Message)
			//fmt.Println("**********")
			sender := evt.Info.Chat.User
			if sender != "status" {
				pushName := evt.Info.PushName

				if sender == "966598840555" {
					fmt.Println("**********")
					fmt.Println(evt.Message)
					/*	StanzaId := evt.Info.ID
						var caption string = ""
						var url = ""
						if evt.Message.ImageMessage != nil {
							caption = *evt.Message.ImageMessage.Caption
							url = *evt.Message.ImageMessage.Url
						}

						JID := evt.Info.Chat

						msg := &waProto.Message{
							ExtendedTextMessage: &waProto.ExtendedTextMessage{
								Text: proto.String("fine"),
								ContextInfo: &waProto.ContextInfo{
									StanzaId: proto.String(StanzaId),
									//	Participant:                      new(string),
									QuotedMessage: &waProto.Message{
										//	Conversation: proto.String(text),
										ImageMessage: &waProto.ImageMessage{
											//	Caption:  proto.String("شكرا و عيدكم مبارك"),
											Mimetype: proto.String("image/png"), // replace this with the actual mime type
											// you can also optionally add other fields like ContextInfo and JpegThumbnail here
											Caption: proto.String(caption),
											Url:     &url,
											Height:  proto.Uint32(410),
											Width:   proto.Uint32(1200),
										},
									},
									//	IsForwarded:                      new(bool),

								},
							},
						}

						//	jid, ok := global.ParseJID(JID)
						//	if !ok {
						//		return
						//	}
						send, err := global.Cli.SendMessage(JID, "", msg) // jid = recipient

						if err != nil {
							global.Log.Errorf("Error sending message: %v", err)
						} else {
							global.Log.Infof("Message sent (server timestamp: %s)", send)
						}
					*/
					fmt.Println("**********")

				}

				switch {
				case evt.Message.Conversation != nil:
					go func(evt *events.Message) {
						var msgReceived string
						received := evt.Message.GetConversation()
						// convert numbers in Arabic scrtip to numbers in latin script
						for _, e := range received {
							if e >= 48 && e <= 57 {
								//	fmt.Println("Number in english script number")
								msgReceived = fmt.Sprintf("%s%v", msgReceived, string(e))
							} else if e >= 1632 && e <= 1641 {
								//	fmt.Println("It is Arabic script")
								msgReceived = fmt.Sprintf("%s%v", msgReceived, global.NormalizeNumber(e))
							} else {
								//	fmt.Println("Dose not looks to be a number")
								msgReceived = fmt.Sprintf("%s%v", msgReceived, string(e))
							}
						}

						re := regexp.MustCompile(`(?i)^(card|كرت|بطاقه|بطاقة)\s(?P<barcode>\w*)`)
						barcodeIndex := re.SubexpIndex("barcode")
						matches := re.FindStringSubmatch(msgReceived)

						re2 := regexp.MustCompile(`(?i)^(return|مرتجع|رجيع)\s(?P<SKU>\w*)`)
						returnIndex := re2.SubexpIndex("SKU")
						returnRequest := re2.FindStringSubmatch(msgReceived)
						//	if sender == "966138117381" {
						//		check.SendDisposable(sender)
						//	}
						//_ = global.ExceptionedNumbers.Contains("1")
						if len(matches) > 0 {
							//	fmt.Println("searching for barcode:", matches[barcodeIndex])
							go branches.SendCard(sender, matches[barcodeIndex])
						} else if len(returnRequest) > 0 {
							go branches.GetReturnPolicy(sender, returnRequest[returnIndex])
						} else if !evt.Info.IsGroup && !evt.Info.IsFromMe && (sender != "966556888145" && // && !evt.Info.IsFromMe
							sender != "966505148268" && sender != "966531041222" && sender != "966577942979" &&
							sender != "966506888972" && sender != "966557776097" && sender != "966505360700" && sender != "966555786616" &&
							sender != "966508884337" && sender != "966508899479" && sender != "966530052201" && sender != "966558936645" &&
							sender != "966502887935" && sender != "971563451686" && sender != "966570936334" && sender != "966543345488" &&
							sender != "966533026947" && sender != "966920006666" && sender != "966532984790" && sender != "966546313131" &&
							sender != "966549700035" && sender != "966580392145" && sender != "966504501968" && sender != "966599522201" &&
							sender != "966551177066" && sender != "966504676300" && sender != "966560080826" && sender != "966554520552" &&
							sender != "966505930535" && sender != "966551775959" && sender != "966561144475") {

							info := whatlanggo.Detect(evt.Message.GetConversation())
							fmt.Println("Language:", info.Lang.String(), " Script:", whatlanggo.Scripts[info.Script], " Confidence: ", info.Confidence)

							switch whatlanggo.Scripts[info.Script] {
							case "Arabic":
								//	go WelcomeMessage(sender, pushName)
								global.Lang = "ar"
							case "Latin":
								global.Lang = "en"
								//	go WelcomeMessageLatin(sender, pushName)
							}
							go WelcomeMessage(sender, strings.TrimSpace(pushName))
						}
						data, _ := PrepareModel(evt.Info.Chat.User,
							sender, pushName, evt.Info.Timestamp.Local().Format("Mon 02-Jan-2006 15:04"),
							evt.Info.ID, "text", msgReceived, "", "")
						global.Passer.Logs <- data
					}(evt)
				case evt.Message.ExtendedTextMessage != nil:
					go func() {
						info, err := json.MarshalIndent(evt.Message.ExtendedTextMessage.GetText(), "", "\t")
						if err != nil {
							fmt.Println(err)
						}
						msgReceived := string(info)
						data, _ := PrepareModel(evt.Info.Chat.User,
							sender, pushName, evt.Info.Timestamp.Local().Format("Mon 02-Jan-2006 15:04"),
							evt.Info.ID, "text", msgReceived, "", "")
						global.Passer.Logs <- data
					}()

				case evt.Message.DeviceSentMessage != nil:
					go func() {
						data, _ := PrepareModel(evt.Info.Chat.User,
							sender, pushName, evt.Info.Timestamp.Local().Format("Mon 02-Jan-2006 15:04"),
							evt.Info.ID, "text", evt.RawMessage.String(), "", "")
						global.Passer.Logs <- data
					}()

				case evt.Message.Chat != nil:
					go func() {
						msgReceived := evt.Message.GetChat()
						data, _ := PrepareModel(evt.Info.Chat.User,
							sender, pushName, evt.Info.Timestamp.Local().Format("Mon 02-Jan-2006 15:04"),
							evt.Info.ID, "text", fmt.Sprintf("%v", msgReceived), "", "")
						global.Passer.Logs <- data
					}()
				//	global.Passer.Logs <- fmt.Sprintf("Chat: ", chat)

				case evt.Message.ImageMessage != nil:
					go func() {
						img := evt.Message.GetImageMessage()
						var caption string = ""
						if evt.Message.ImageMessage.Caption != nil {
							caption = *evt.Message.ImageMessage.Caption
						}

						if img != nil {
							file, err := global.Cli.Download(img)
							if err != nil {
								log.Printf("Failed to download image: %v", err)
								return
							}
							exts, _ := mime.ExtensionsByType(img.GetMimetype())
							path := fmt.Sprintf("D:/Deployment/DigitalAssistance/Downloads/Image/%s-%s%s", sender, evt.Info.ID, exts[0])
							err = os.WriteFile(path, file, 0600)
							if err != nil {
								log.Printf("Failed to save image: %v", err)
								return
							}
							log.Printf("Saved image in message to %s", path)
							data, _ := PrepareModel(evt.Info.Chat.User,
								sender, pushName, evt.Info.Timestamp.Local().Format("Mon 02-Jan-2006 15:04"),
								evt.Info.ID, "image", caption, "", path)
							global.Passer.Logs <- data
							//	global.Passer.Logs <- fmt.Sprintf("Image: <a href='%v' target='_blank'>Open</a>", path)
						}
					}()

				case evt.Message.StickerMessage != nil:
					go func() {
						sticker := evt.Message.GetStickerMessage()
						if sticker.Url != nil {
							fmt.Println(sticker)
							/*	data, err := global.Cli.Download(audio)
								if err != nil {
									log.Printf("Failed to download audio: %v", err)
									return
								} */
							//	global.Passer.Logs <- fmt.Sprintf("Sticker: <a href='%v' target='_blank'>Open</a>", sticker.GetUrl())
						}
					}()
					data, _ := PrepareModel(evt.Info.Chat.User,
						sender, pushName, evt.Info.Timestamp.Local().Format("Mon 02-Jan-2006 15:04"),
						evt.Info.ID, "sticker", "Sticker recieved", "", "")
					global.Passer.Logs <- data

				case evt.Message.AudioMessage != nil:
					go func() {
						audio := evt.Message.GetAudioMessage()
						if audio != nil {
							file, err := global.Cli.Download(audio)
							if err != nil {
								log.Printf("Failed to download audio: %v", err)
								return
							}
							exts, _ := mime.ExtensionsByType(audio.GetMimetype())
							path := fmt.Sprintf("D:/Deployment/DigitalAssistance/Downloads/Audio/%s-%s%s", sender, evt.Info.ID, exts[0])
							err = os.WriteFile(path, file, 0600)
							if err != nil {
								log.Printf("Failed to save audio: %v", err)
								return
							}
							log.Printf("Saved audio in message to %s", path)
							data, _ := PrepareModel(evt.Info.Chat.User,
								sender, pushName, evt.Info.Timestamp.Local().Format("Mon 02-Jan-2006 15:04"),
								evt.Info.ID, "audio", "", "", path)
							global.Passer.Logs <- data
							//	global.Passer.Logs <- fmt.Sprintf("Audio: <a href='%v' target='_blank'>Open</a>", path)
						}
					}()

				case evt.Message.VideoMessage != nil:
					go func() {
						video := evt.Message.GetVideoMessage()
						var caption string = ""
						if evt.Message.VideoMessage.Caption != nil {
							caption = *evt.Message.VideoMessage.Caption
						}

						if video != nil {
							file, err := global.Cli.Download(video)
							if err != nil {
								log.Printf("Failed to download video: %v", err)
								return
							}
							exts, _ := mime.ExtensionsByType(video.GetMimetype())
							path := fmt.Sprintf("D:/Deployment/DigitalAssistance/Downloads/Video/%s-%s%s", sender, evt.Info.ID, exts[0])
							err = os.WriteFile(path, file, 0600)
							if err != nil {
								log.Printf("Failed to save video: %v", err)
								return
							}
							log.Printf("Saved video in message to %s", path)
							//	global.Passer.Logs <- fmt.Sprintf("Video: <a href='%v' target='_blank'>Open</a>", path)
							data, _ := PrepareModel(evt.Info.Chat.User,
								sender, pushName, evt.Info.Timestamp.Local().Format("Mon 02-Jan-2006 15:04"),
								evt.Info.ID, "video", caption, "", path)
							global.Passer.Logs <- data
						}
					}()
				case evt.Message.DocumentMessage != nil:
					go func() {
						document := evt.Message.GetDocumentMessage()
						if document != nil {
							file, err := global.Cli.Download(document)
							if err != nil {
								log.Printf("Failed to download audio: %v", err)
								return
							}
							exts, _ := mime.ExtensionsByType(document.GetMimetype())
							path := fmt.Sprintf("D:/Deployment/DigitalAssistance/Downloads/Documents/%s-%s%s", sender, evt.Info.ID, exts[0])
							err = os.WriteFile(path, file, 0600)
							if err != nil {
								log.Printf("Failed to save document: %v", err)
								return
							}
							log.Printf("Saved document in message to %s", path)
							data, _ := PrepareModel(evt.Info.Chat.User,
								sender, pushName, evt.Info.Timestamp.Local().Format("Mon 02-Jan-2006 15:04"),
								evt.Info.ID, "document", "", "", path)
							global.Passer.Logs <- data
							//	global.Passer.Logs <- fmt.Sprintf("Document: <a href='%v' target='_blank'>Open</a>", path)
						}
					}()
					/*	case evt.Message.ImageMessage != nil,
						evt.Message.AudioMessage != nil,
						evt.Message.VideoMessage != nil,
						evt.Message.DocumentMessage != nil:
						if !evt.Info.IsGroup && !evt.Info.IsFromMe && (sender != "966556888145" && // && !evt.Info.IsFromMe
							sender != "966505148268" && sender != "966531041222" && sender != "966577942979" &&
							sender != "966506888972" && sender != "966557776097" && sender != "966505360700" && sender != "966555786616" &&
							sender != "966508884337" && sender != "966508899479" && sender != "966530052201" && sender != "966558936645" &&
							sender != "966502887935" && sender != "971563451686") {
							OnlyText(sender)
						} */
				case evt.Message.ButtonsResponseMessage != nil:
					fmt.Println("Button responce pressed")
					ButtonResponse := evt.Message.GetButtonsResponseMessage()
					id, _ := strconv.Atoi(ButtonResponse.GetSelectedButtonId())
					responces.ButtonResponses(id, sender, evt.Info.Chat)

				case evt.Message.ListResponseMessage != nil:
					fmt.Println("List responce pressed")
					ListResponse := evt.Message.GetListResponseMessage()
					id, _ := strconv.Atoi(ListResponse.SingleSelectReply.GetSelectedRowId())
					fmt.Println(id, sender)
					responces.ListResponces(id, sender, pushName, evt.Info.Chat)
				case evt.Message.LocationMessage != nil:

					Location := evt.Message.GetLocationMessage()
					fmt.Println(Location.GetDegreesLatitude())
					fmt.Println(Location.GetDegreesLongitude())
					fmt.Println(Location.GetAddress())

					latitude := Location.GetDegreesLatitude()
					longitud := Location.GetDegreesLongitude()
					address := Location.GetAddress()
					link := fmt.Sprintf("<a href='https://www.google.com/maps/@%f,%f,15z' target='_blank'>Open map</a>", latitude, longitud)
					go func() {
						data, _ := PrepareModel(evt.Info.Chat.User,
							sender, pushName, evt.Info.Timestamp.Local().Format("Mon 02-Jan-2006 15:04"),
							evt.Info.ID, "location", "Location: "+link, "", "")
						global.Passer.Logs <- data
					}()

					global.Locations = append(global.Locations, global.Location{
						Sender:    sender,
						PushName:  pushName,
						Latitude:  Location.GetDegreesLatitude(),
						Longitude: Location.GetDegreesLongitude(),
						Address:   Location.GetAddress(),
					})

					tx, err := global.Db.Begin()
					if err != nil {
						log.Fatal(err)
					}
					stmt, err := tx.Prepare("insert into locations (name, jid, longitude, latitude, address) values(?, ?, ?, ?, ?)")
					if err != nil {
						log.Fatal(err)
					}
					defer stmt.Close()

					_, err = stmt.Exec(pushName, sender, longitud, latitude, address)
					if err != nil {
						log.Fatal(err)
					}

					tx.Commit()

					fmt.Println(global.Locations)

				case evt.Message.ContactMessage != nil:
					Contact := evt.Message.GetContactMessage()
					fmt.Println(Contact.GetDisplayName())
					fmt.Println(Contact.GetVcard())
				} // End of switch
			}
		}
	}()
}
