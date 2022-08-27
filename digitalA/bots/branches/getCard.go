package branches

import (
	"DigitalAssistance/global"
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

func GetCard(barcode string) (target global.SKUcard[float64], ipartString, decpartString string) {
	webApp := "AKfycbyHSUv6Y9CXrdXrsdPFY6E5cxXMWkKQm71UORYtKHcT6Rr58OxmklAc1fQFIK8rnBTD"
	// webApp := "AKfycbxfc7e9V9rbDpX8eeoqlsNOunjs4VM496ibPEmdwhaTdJhmHmThI-Hj7_RxrUgpwLz0"
	req, _ := http.NewRequest("GET", "https://script.google.com/macros/s/"+webApp+"/exec?", nil)

	q := req.URL.Query()
	q.Add("barcode", barcode)

	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	// resp, err := http.Get(req.URL.String())
	resp, err := http.DefaultClient.Do(req)
	_ = req.URL.RawQuery
	if err != nil {
		global.Log.Errorf("Error fetching data: %v", err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	// fmt.Println(buf.String())
	if buf.String() == "NotFound" {
		return
	} else {
		fmt.Println(buf.String())
		//	var target = new(global.SKUcard[float64])
		err = json.NewDecoder(buf).Decode(&target)
		if err != nil {
			global.Log.Errorf("Error parsing response body: %v", err)
		}
		//	fmt.Println("barcode:", target.BarCode, "price:", target.ShelfPrice)
		ipart := math.Floor(target.ShelfPrice)
		/*	var intPointer *float64 = new(float64)
			*intPointer = ipart
			if intPointer == nil {
				ipart2 := int64(target.ShelfPrice)
				fmt.Println("ipart2", ipart2)
			} else {
				fmt.Println("ipart", ipart)
			} */
		//	fmt.Println("ipart", ipart, "type", reflect.TypeOf(ipart), "balance:", target.ShelfPrice-ipart)
		balance := target.ShelfPrice - ipart
		//	fmt.Println("balance:", balance)
		decpart := fmt.Sprint(balance) // float64( // "%.3g"
		fmt.Println("len of dec:", len(decpart))
		//	fmt.Println("decpart:", decpart, "length:", len(decpart))
		if len(decpart) == 1 {
			decpartString = "00"
		} else if len(decpart) == 3 {
			decpartString = decpart[2:] + "0"
		} else if len(decpart) > 3 {
			decpartString = decpart[2:4]
		} else {
			fmt.Println("check decemials")
		}
		ipartString = fmt.Sprint(ipart)
		//	fmt.Println("shelf price", target.ShelfPrice)
		//	fmt.Println("ipartString", ipartString)
		//	fmt.Println("decpartString", decpartString)
		return
	}
}
