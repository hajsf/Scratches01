package vendor

import "fmt"

func createVCard(display, number string) string {
	card := fmt.Sprintf("BEGIN:VCARD\nVERSION:3.0\nN:;%v;;;\nFN:%v\nTEL;type=CELL;waid=%v:+%v\nEND:VCARD", display, display, number, number)
	return card
}
