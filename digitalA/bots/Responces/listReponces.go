package responces

import (
	"DigitalAssistance/Enum"
	"DigitalAssistance/bots/vendor"
	"DigitalAssistance/communication"
	"fmt"

	"go.mau.fi/whatsmeow/types"
)

func ListResponces(id int, sender, pushname string, chat types.JID) {
	fmt.Println(id, sender)
	switch id {
	case Enum.ContractTerms:
		vendor.SendContract(sender)
		AnotherQuestion(sender)
	case Enum.VenderRegistration:
		vendor.VendorRegistration(sender, chat)
		AnotherQuestion(sender)
	case Enum.ItemRegistration:
		vendor.ItemRegistration(sender, chat)
		AnotherQuestion(sender)
	case Enum.Location:
		vendor.SendLocation(sender)
		AnotherQuestion(sender)
	case Enum.CallArrangement:
		fmt.Println("request call")
		go communication.SendEmail(sender, pushname, "Telephone call request")
		vendor.ContactSuperIntendent(sender)
		vendor.SendContact(sender)
		AnotherQuestion(sender)
	case Enum.VisitArrangement:
		go communication.SendEmail(sender, pushname, "Visit/Meeting request")
		//vendor.ContactSuperIntendent(sender)
		vendor.SendContact(sender)
		AnotherQuestion(sender)
	case Enum.VendorComplain:
		vendor.VendorIssue(sender)
		AnotherQuestion(sender)
	case Enum.Discount:
		vendor.VendorDiscount(sender)
		AnotherQuestion(sender)
	case Enum.VAT:
		vendor.SendVAT(sender)
		AnotherQuestion(sender)
	case Enum.PriceList:
		vendor.UploadPriceList(sender)
		AnotherQuestion(sender)
	case Enum.PriceChange:
		vendor.PriceChange(sender)
		AnotherQuestion(sender)
	case Enum.AdDocumentation:
		vendor.VendorAd(sender)
		AnotherQuestion(sender)
	case Enum.Email:
		vendor.SendMyEmail(sender)
		AnotherQuestion(sender)
	}

}
