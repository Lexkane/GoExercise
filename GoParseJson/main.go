package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RouteStruct  struct {
	Index  string   `json:"index"`
	Routes []string `json:"routes"`
}


type RouteStructs [] RouteStruct // `json:"routes"`

//var 	AllRoutes map [string] [] string



var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target *RouteStruct) error {

	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
func main() {

	//var result map[string]interface{}
	//var result []interface{}

	//route := new(RouteStruct)
	AllRoutes:=make(map[string][]string,100)

	//getJson("https://api.myjson.com/bins/x7rfk", route)
	r, err := myClient.Get("https://api.myjson.com/bins/x7rfk")

	if err != nil {
		fmt.Printf("Error while get responce ")

	}

	respBytes, err := ioutil.ReadAll(r.Body)
	if err !=nil{
		fmt.Println("Error while reading http request")
	}
	//json.Unmarshal(respBytes,&allRoutes)
	//jsonBytes := respBytes[bytes.Index(respBytes, []byte("{")):bytes.LastIndex(respBytes, []byte("}"))+1]
	routes := new(RouteStructs)

	//allRoutes:=new(AllRoutes)

	json.Unmarshal(respBytes,&routes)


	for _,v:=range *routes{
		AllRoutes[v.Index]=v.Routes
	}

	fmt.Println(AllRoutes)

	//fmt.Println("%s", r.Header.Get("https://api.myjson.com/bins/x7rfk"))


}
