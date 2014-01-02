package utils

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"log"
	//"os"
	//"encoding/json"
	"net/http"
)

type HttpResponse struct {
	Url      string
	ByteStr  []byte
	Response *http.Response
	Err      error
}

func Enumerate(x interface{}) {
	val := reflect.ValueOf(x).Elem()

	i := 0
	for {
		if(i >= val.NumField()){
			break
		}
		//valueField := val.Field(i)
		typeField := val.Type().Field(i)

		fmt.Printf("Field Name: %s,\t Field Value: ,\t \n",
			typeField.Name)
		i++
	}
}

func Get(url string, pid string ) (chan *HttpResponse) {

	channel  := make(chan *HttpResponse)
	client   := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	CheckError(err)

	req.Header.Set("pid",pid)
	req.Header.Set("fp","gormn")

	go func(){
		resp, err := client.Do(req)
		
		CheckError(err)

		defer resp.Body.Close()

		bs, _ := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		channel <- &HttpResponse{url, bs, resp, err}
	}()

	return channel
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err)
		log.Fatal(err)
		//os.Exit(1)
	}
}
