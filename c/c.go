package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


type Coupons struct {
	Coupon []string
}

type Cupons struct{
	Codes []string
}

func (c Coupons) Check(code string) string {
	for _, item := range c.Coupon {
		if code == item {
			return "valid"
		}		
	}
	return "invalid"
}

type Result struct {
	Status string
}

var coupons Coupons

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9092", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	getCoupons()
	 coupon := r.PostFormValue("coupon")
	 valid := coupons.Check(coupon)
	fmt.Printf("%s \n", valid)

	 result := Result{Status: valid}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}
	fmt.Fprintf(w, string(jsonResult))

}


func getCoupons(){
	response, err := http.Get("http://localhost:9094/coupons")
	
    if err != nil {
        fmt.Printf("%s", err)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
		fmt.Printf("%s\n", string(contents))
		cupons := Cupons{}

		json.Unmarshal([]byte(contents), &cupons)

		for _, item := range cupons.Codes {
			coupons.Coupon = append(coupons.Coupon, item)
		}
    }
}