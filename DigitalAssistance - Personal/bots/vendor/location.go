package vendor

import (
	"fmt"

	"DigitalAssistance/global"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func SendLocation(sender string) {
	fmt.Println("sending file")
	msg := &waProto.Message{
		LocationMessage: &waProto.LocationMessage{
			DegreesLatitude:  proto.Float64(26.3723116),
			DegreesLongitude: proto.Float64(50.0410585),
			Name:             proto.String("موقع إدارة شركة قطوف و حلا - الدمام"),
			/*	Address:                           new(string),
				Url:                               new(string),
				IsLive:                            new(bool),
				AccuracyInMeters:                  new(uint32),
				SpeedInMps:                        new(float32),
				DegreesClockwiseFromMagneticNorth: new(uint32),
				Comment:                           new(string),
				JpegThumbnail:                     []byte{},
				ContextInfo:                       &waProto.ContextInfo{}, */
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
