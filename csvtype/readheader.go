package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {
	filePath := "./file.csv"
	headerNames := make(map[int]string)
	headerTypes := make(map[int]string)
	// Load a csv file.
	f, _ := os.Open(filePath)
	// Create a new reader.
	r := csv.NewReader(f)
	// Read first row only
	header, err := r.Read()
	checkError("Some other error occurred", err)

	// Add mapping: Column/property name --> record index
	for i, v := range header {
		headerNames[i] = v
	}

	// Read second row
	record, err := r.Read()
	checkError("Some other error occurred", err)

	// layout := "3/17/2022 10:18:54 AM"
	// Check record fields types
	for i, v := range record {
		var value interface{}
		if value, err = strconv.Atoi(v); err != nil {
			if value, err = strconv.ParseFloat(v, 64); err != nil {
				if value, err = strconv.ParseBool(v); err != nil {
					if value, err = strconv.ParseBool(v); err != nil { // <== How to do this with unknown layout
						// Value is a string
						headerTypes[i] = "string"
						value = v
						fmt.Println(reflect.TypeOf(value), reflect.ValueOf(value))
					} else {
						// Value is a timestamp
						headerTypes[i] = "time"
						fmt.Println(reflect.TypeOf(value), reflect.ValueOf(value))
					}
				} else {
					// Value is a bool
					headerTypes[i] = "bool"
					fmt.Println(reflect.TypeOf(value), reflect.ValueOf(value))
				}
			} else {
				// Value is a float
				headerTypes[i] = "float"
				fmt.Println(reflect.TypeOf(value), reflect.ValueOf(value))
			}
		} else {
			// Value is an int
			headerTypes[i] = "int"
			fmt.Println(reflect.TypeOf(value), reflect.ValueOf(value))
		}
	}

	for i := range header {
		fmt.Printf("Header: %v \tis\t %v\n", headerNames[i], headerTypes[i])
	}

	e := CreateStruct() // reflect.Value
	e.FieldByName("Area").SetInt(1234)
	fmt.Printf("value: %+#v\n", e)
	fmt.Printf("value: %+#v %+#v\n", e.FieldByName("Myfield"), e.FieldByName("Area"))

	s := e.Addr().Interface()

	w := new(bytes.Buffer)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}

	fmt.Printf("value: %+v\n", s)
	fmt.Printf("json:  %s", w.Bytes())

	rx := bytes.NewReader([]byte(`{"Myfield":0,"Area":1234,"Size":20}`))
	if err := json.NewDecoder(rx).Decode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)

}

/*
	// The map will be sorted by header name alphapatically, we need to sort it based on header index, i.e. by value
	temp := map[int][]string{}
	var a []int
	for k, v := range headerMap {
		temp[v] = append(temp[v], k)
	}
	for k := range temp {
		a = append(a, k)
	}

	// sort in increasing order, if required to sort in decreasing order use: sort.Sort(sort.Reverse(sort.IntSlice(a))) if need reverse sorting
	sort.Ints(a)
	/*	for _, k := range a {
		for _, s := range temp[k] {
			fmt.Printf("%s, %d\n", s, k)
		}
	} *

	jsonByte, err := json.Marshal(headerMap)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println("csv file headers had been read")
		// fmt.Println(string(jsonByte))
	}

	// Save the header to text file
	out, err := os.Create("headers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err2 := out.WriteString(string(jsonByte))

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Headers file created")
}
*/
