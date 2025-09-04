package main

import (
	"fmt"
	"net/url"
)

func main() {
	//url parsing
	//[Scheme ://][userinfo@]host[:port][/path][?query][#fragment]
	rawurl := "https://example.com:8080/path?query=param#fragment"
	parsedurl, err := url.Parse(rawurl)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(parsedurl.Scheme, parsedurl.Fragment, parsedurl.Host, parsedurl.Fragment, parsedurl.Port(), parsedurl.RawQuery)


	rawurl1 :="https://example.com/path?name=john&age=30"
	parserawurl1,err:=url.Parse(rawurl1); if err !=nil{
		fmt.Println("Error parsing url",err)
		return
	}
	queryParams :=parserawurl1.Query()
    fmt.Println("fckk",queryParams)
	fmt.Println("Name :",queryParams.Get("name"))
    	fmt.Println("Age :",queryParams.Get("age"))

		//building url

		basedURL :=&url.URL{
			Scheme: "https",
			Host: "example.com",
			Path: "/path",
		}
		query :=basedURL.Query()
		query.Set("name","John")
		query.Set("age","30")
		basedURL.RawQuery=query.Encode()

		fmt.Println("built url",basedURL.String())
		val :=url.Values{}

		//add some key val pairs to value object
		val.Add("name","Abhi")
		val.Add("age","30")
		val.Add("city","Hetimpur")
		val.Add("country","INdia")

		//encoded query
		encodedquery  :=val.Encode()

		fmt.Println(encodedquery)

		//built a url 
		basedURL, _ = url.Parse("https://example@.com/search")
		fullurl := basedURL.String() + "?" + encodedquery

		fmt.Println(fullurl)

	
}
