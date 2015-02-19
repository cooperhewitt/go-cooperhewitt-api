package main

import (
	"flag"
	"fmt"
	"net/url"
	"org.cooperhewitt/api"
)

func main() {

	token := flag.String("token", "", "token")
	flag.Parse()

	client := api.OAuth2Client(*token)

	method := "cooperhewitt.labs.whatWouldMicahSay"
	args := url.Values{}

	rsp, _ := client.ExecuteMethod(method, &args)

	// Go is weird...

	data := rsp.(map[string]interface{})
	micah := data["micah"].(map[string]interface{})

	fmt.Println(micah["says"])
}
