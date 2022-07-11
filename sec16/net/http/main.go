package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//net/http
func main() {
	// GET method
	res, _ := http.Get("https://example.com")

	fmt.Println(res.StatusCode)
	fmt.Println(res.Proto) //プロトコル

	fmt.Println(res.Header["Date"])
	fmt.Println(res.Header["Content-Type"])

	fmt.Println(res.Request.Method)
	fmt.Println(res.Request.URL)
	defer res.Body.Close() // 最後にレスポンスをcloseする

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	// POST method
	vs := url.Values{}
	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	fmt.Println(vs.Encode())

	res, err := http.PostForm("https://example.com", vs)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ = ioutil.ReadAll(res.Body)
	fmt.Print(string(body))
}
