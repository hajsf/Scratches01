package responces

import (
	"DigitalAssistance/Enum"
	"DigitalAssistance/bots/branches"
	"DigitalAssistance/bots/franchisee"
	"DigitalAssistance/bots/vendor"
	"DigitalAssistance/global"
)

func ButtonResponses(id int, sender string) {
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
		vendor.SendContract(sender)
		vendor.VendorRegistration(sender)
		vendor.ItemRegistration(sender)
		vendor.SendLocation(sender)
		vendor.FillForms(sender)
		AnotherQuestion(sender)
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
		Bye(sender)

	}
}
