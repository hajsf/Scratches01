package responces

import (
	"DigitalAssistance/Enum"
	"DigitalAssistance/bots/branches"
	"DigitalAssistance/bots/franchisee"
	"DigitalAssistance/bots/vendor"
	"DigitalAssistance/global"
	"fmt"

	"go.mau.fi/whatsmeow/types"
)

func ButtonResponses(id int, sender string, chat types.JID) {
	switch id {

	case Enum.Vendor:
		go vendor.WhichVender(sender)
	case Enum.ExistingVendor:
		global.Users = append(global.Users, global.Communicator{
			Sender:     sender,
			UserType:   Enum.ExistingVendor,
			UserScript: "Arabic",
		})
		go vendor.CurrentVender(sender)
	//	franchisee.ContactSuperIntendent(sender)
	//	franchisee.SendContact(sender)
	//	Bye(sender)
	case Enum.NewVendor:
		global.Users = append(global.Users, global.Communicator{
			Sender:     sender,
			UserType:   Enum.NewVendor,
			UserScript: "Arabic",
		})
		//	vendor.NewVender(sender)
		vendor.VendorRegistration(sender, chat)
		vendor.ItemRegistration(sender, chat)
		vendor.SendLocation(sender)
		vendor.FillForms(sender)
		go vendor.SendContract(sender)
		go vendor.SendCoGS(sender)
		go vendor.SendCR(sender)
		go vendor.SendContractDraf(sender)
		go vendor.SendBranches(sender)
		go vendor.SendAnnex(sender)
		go vendor.SendChamper(sender)
		go vendor.SendVAT(sender)

		//AnotherQuestion(sender)
		vendor.CurrentVender(sender)
	case Enum.Supervisor:
		global.Users = append(global.Users, global.Communicator{
			Sender:     sender,
			UserType:   Enum.Supervisor,
			UserScript: "Arabic",
		})
		branches.BranchIssue(sender)
		//	franchisee.ContactSuperIntendent(sender)
		//	franchisee.SendContact(sender)
		Bye(sender)
	case Enum.Franchisee:
		global.Users = append(global.Users, global.Communicator{
			Sender:     sender,
			UserType:   Enum.Franchisee,
			UserScript: "Arabic",
		})
		franchisee.ContactSuperIntendent(sender)
		franchisee.SendContact(sender)
		Bye(sender)
	case Enum.Yes: // yes another service is required
		vendor.CurrentVender(sender)
	/*	var userType int
		for _, v := range global.Users {
			if v.Sender == sender {
				userType = v.UserType
				break
			}
			//fmt.Println(k, "is:", v[sender])
			/*	for key, value := range v.UserType {
				if key == sender {
					userType = v.UserType
					break
				}
			} *
		}

		switch userType {
		case Enum.ExistingVendor:
			vendor.CurrentVender(sender)
		case Enum.NewVendor:
			vendor.NewVender(sender)
		case Enum.Franchisee:
		case Enum.Supervisor:
		} */
	case Enum.No:
		fmt.Println("bye bye")
		Bye(sender)

	}
}
