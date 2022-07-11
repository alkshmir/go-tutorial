package main

import (
	"fmt"
	"net/url"
)

// urlもじれつを処理する

func main() {
	// urlをパース
	u, _ := url.Parse("http://example.com/search?a=1&b=2#top")
	fmt.Println(u.Scheme)   // http
	fmt.Println(u.Host)     // example.com
	fmt.Println(u.RawQuery) // a=1&b=2
	fmt.Println(u.Path)     // /search
	fmt.Println(u.Fragment) // top

	fmt.Println(u.Query()) //mapでqueryを取得

	// urlを生成
	url := &url.URL{}
	url.Scheme = "https"
	url.Host = "google.com"
	q := url.Query()
	q.Set("q", "Golang")

	url.RawQuery = q.Encode() //日本語はBASE64 エンコードしてくれる

	fmt.Println(url)
}
