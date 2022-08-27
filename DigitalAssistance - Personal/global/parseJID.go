package global

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var Bundle *i18n.Bundle
var Lang string

type Communicator struct {
	Sender     string
	UserType   int
	UserScript string
}

var Cli *whatsmeow.Client
var Log waLog.Logger
var Users []Communicator
var DbMessages = flag.String("db-locations", "file:locations.db?loc=auto", "Database address")
var Db *sql.DB

type Location struct {
	Sender, PushName, Address string
	Latitude, Longitude       float64
}

var Locations []Location

type Number interface {
	int64 | float64
}

type AnyString string

// "1WPWH9zkQ25CASuh3uXs7T3i5PvsEfIKU"
type SKUcard[number Number] struct {
	BarCode, SKUCode, VendorCode, RegistrationDate                                AnyString
	VendorName, BrandName, ContactPerson                                          string
	ContactNumber, ItemName, ItemImage                                            AnyString
	NetWeight, CartoonPack, StorageTemperature, ShelfLife, ShelfPrice, KottofCost number
	SupplyType, CoveredAreas, MinimumOrderQty, ContractDate, ReturnPolicy, Notes  string
	InActive                                                                      bool
}

func ParseJID(arg string) (types.JID, bool) {
	if arg[0] == '+' {
		arg = arg[1:]
	}
	if !strings.ContainsRune(arg, '@') {
		return types.NewJID(arg, types.DefaultUserServer), true
	} else {
		recipient, err := types.ParseJID(arg)
		if err != nil {
			Log.Errorf("Invalid JID %s: %v", arg, err)
			return recipient, false
		} else if recipient.User == "" {
			Log.Errorf("Invalid JID %s: no server specified", arg)
			return recipient, false
		}
		return recipient, true
	}
}

func (s *AnyString) UnmarshalJSON(data []byte) error {
	//	fmt.Println(data, "length: ", len(data))
	var prefex string
	var link string
	if len(data) == 68 {
		link, _ = strconv.Unquote(string(data))
		prefex = link[:32] //	https://drive.google.com/open?id=1WPWH9zkQ25CASuh3uXs7T3i5PvsEfIKU
	}
	//const layout = "\"2006-01-02T03:04:05.999Z\""
	const layout = "2006-01-02T15:04:05.999Z"
	if data[0] != '"' { // not string, so probably int, make it a string by wrapping it in double quotes
		data = []byte(`"` + string(data) + `"`)
	} else if prefex == "https://drive.google.com/open?id" {
		id := string(data)[34:67] //	https://drive.google.com/open?id=1WPWH9zkQ25CASuh3uXs7T3i5PvsEfIKU
		url := "https://drive.google.com/uc?export=view&id="
		data = []byte(`"` + fmt.Sprintf("%v%v", url, string(id)) + `"`)
	} else if len(data) == 26 {
		timestamp := string(data)
		stamp, err := strconv.Unquote(timestamp)
		if err != nil {
			fmt.Println(err)
		}
		t, error := time.Parse(layout, stamp)

		if error != nil {
			Log.Errorf("Error parsing data in template: %v", err)
		} else {
			// fmt.Println("valid!:", t)
			y, m, d := t.Date()
			// fmt.Printf("%v/%v/%v", d, m, y)
			data = []byte(`"` + fmt.Sprintf("%v/%v/%v", d, m, y) + `"`)
		}
	}

	// unmarshal as plain string
	return json.Unmarshal(data, (*string)(s))
}
