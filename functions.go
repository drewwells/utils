package utils

import (
	"fmt"
	"io/ioutil"
	"reflect"
	//"encoding/json"
	"net/http"
)

var channel = make(chan *HttpResponse)

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

	client := &http.Client{
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("pid",pid)
	req.Header.Set("fp","gormn")

	go func(){
		resp, errr := client.Do(req)

		//var data interface{}
		if errr != nil {
			fmt.Println(errr)
		}

		defer resp.Body.Close()

		bs, _ := ioutil.ReadAll(resp.Body)
		//err := json.Unmarshal(bs, &data)
		if err != nil {
			fmt.Println(err)
		}

		//fmt.Printf("%+v", string(bs))

		channel <- &HttpResponse{url, bs, resp, err}
	}()

	return channel
}
