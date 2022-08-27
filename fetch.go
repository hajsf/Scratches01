package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Number interface {
	int64 | float64
}
type AnyString string

func (s *AnyString) UnmarshalJSON(data []byte) error {
	//const layout = "\"2006-01-02T03:04:05.999Z\""
	const layout = "2006-01-02T03:04:05.999Z"
	if data[0] != '"' { // not string, so probably int, make it a string by wrapping it in double quotes
		data = []byte(`"` + string(data) + `"`)
	} else if len(data) > 4 {
		timestamp := fmt.Sprintf("%s", data)
		stamp, err := strconv.Unquote(timestamp)
		if err != nil {
			fmt.Println(err)
		}
		t, error := time.Parse(layout, stamp)

		if error != nil {
			fmt.Println(error)
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

type SKUcard[number Number] struct {
	BarCode, SKUCode, VendorCode, RegistrationDate                                AnyString
	VendorName, BrandName, ContactPerson                                          string
	ContactNumber                                                                 AnyString
	ItemName, ItemImage                                                           string
	NetWeight, CartoonPack, StorageTemperature, ShelfLife, ShelfPrice, KottofCost number
	SupplyType, CoveredAreas, MinimumOrderQty, ContractDate, ReturnPolicy, Notes  string
	InActive                                                                      string
}

func main() {
	req, _ := http.NewRequest("GET", "https://script.google.com/macros/s/AKfycbzw0TKWycxeB5sx1wIefAiEHeYQt2mVuM-NAZTccxedhyntdv8FvcUteOZ2k03wRHGE/exec?", nil)

	q := req.URL.Query()
	q.Add("barcode", "6287029390129")
	//q.Add("another_thing", "foo & bar")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	// url := "https://script.google.com/macros/s/AKfycbzw0TKWycxeB5sx1wIefAiEHeYQt2mVuM-NAZTccxedhyntdv8FvcUteOZ2k03wRHGE/exec?barcode=6287029390129"
	// resp, err := http.Get(req.URL.String())
	resp, err := http.DefaultClient.Do(req)
	_ = req.URL.RawQuery
	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	/* body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb) */

	//	var target interface{}
	var target = new(SKUcard[float64])
	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		fmt.Println("err:", err)
	}

	/*	if s, err := strconv.ParseFloat(target.BarCode[int64], 32); err == nil {
		fmt.Println(s) // 3.1415927410125732
	} */

	/*	str := fmt.Sprintf("%v", target.BarCode)
		fmt.Println(str)

		fmt.Println(reflect.TypeOf(target.BarCode))

		fmt.Println(fmt.Sprintf("%.0f", target.BarCode)) // 6287029390129

		fmt.Println(target) */
	fmt.Println(target.BarCode, target.RegistrationDate)

	ipart := int64(target.ShelfPrice)
	decpart := fmt.Sprintf("%.3g", target.ShelfPrice-float64(ipart))[2:]
	fmt.Println(target.ShelfPrice, ipart, decpart)
}
