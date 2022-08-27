package branches

import (
	"DigitalAssistance/global"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

type Policy struct {
	SKUCode, Description, Vendor, SupplyType, ReturnPolicy string
}

func GetReturnPolicy(sender, SKU string) {
	webApp := "AKfycbyzdRbXIackgJHtTnzSwqpIyLpPBXzae1qncK_zf1hPb9idooNPIVX5WgHhcC0yU16Vvg"
	req, _ := http.NewRequest("GET", "https://script.google.com/macros/s/"+webApp+"/exec?", nil)
	q := req.URL.Query()
	q.Add("sku", SKU)
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	resp, err := http.DefaultClient.Do(req)
	_ = req.URL.RawQuery
	if err != nil {
		global.Log.Errorf("Error fetching data: %v", err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	// fmt.Println(buf.String())
	if buf.String() == "NotFound" {
		msg := &waProto.Message{
			Conversation: proto.String("الصنف ليس له مرتجع، يرجى مراجعة قسم المشتريات")}

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
	} else {
		target := Policy{}
		err = json.NewDecoder(buf).Decode(&target)
		if err != nil {
			global.Log.Errorf("Error parsing response body: %v", err)
		}

		var sb strings.Builder

		sb.WriteString(fmt.Sprintf("رقم الصنف      : *%v*\n", target.SKUCode))
		sb.WriteString(fmt.Sprintf("إسم الصنف      : *%v*\n", target.Description))
		sb.WriteString(fmt.Sprintf("المورد              : *%v*\n", target.Vendor))
		sb.WriteString(fmt.Sprintf("نوع التوريد      : *%v*\n", target.SupplyType))
		sb.WriteString(fmt.Sprintf("سياسة المرتجع : *%v*\n", target.ReturnPolicy))

		msg := &waProto.Message{
			Conversation: proto.String(sb.String())}

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
}
