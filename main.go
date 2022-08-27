package main

import (
	"io"
	"log"
	"strings"

	"github.com/emersion/go-vcard"
)

func main() {

	var content strings.Builder
	content.WriteString("BEGIN:VCARD\n")
	content.WriteString("VERSION:3.0\n")
	content.WriteString("FN:Hasan Yousef\n")
	content.WriteString("TEL:999\n")
	content.WriteString("Fitem1.TEL;waid=966000000000\n")
	content.WriteString("item1.X-ABLabel:Mobile\n")
	content.WriteString("END:VCARD")

	r := strings.NewReader(content.String())
	dec := vcard.NewDecoder(r)
	var card vcard.Card
	var err error
	for {
		card, err = dec.Decode()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		log.Println(card.PreferredValue(vcard.FieldFormattedName))
		log.Println(card.PreferredValue(vcard.FieldTelephone))
	}

	var enc vcard.Encoder

	// enc = vcard.NewEncoder(destFile)
	err = enc.Encode(card)
	if err != nil {
		log.Fatal(err)
	}

	// How can I convert enc to string!

	/*	io.ReadWriter(enc)
		io.WriteString(enc.w, begin)
	*/
	//x := enc.String()
	log.Println(enc.String())

}
