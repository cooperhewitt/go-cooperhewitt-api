package main

import (
	"flag"
	"fmt"
	"net/url"
	"org.cooperhewitt/api"
	"strings"
)

func main() {

	token := flag.String("token", "", "token")
	flag.Parse()

	echo := flag.Args()
	str := strings.Join(echo, " ")

	client := api.OAuth2Client(*token)

	method := "api.test.echo"
	args := url.Values{}
	args.Set("echo", str)

	rsp, _ := client.ExecuteMethod(method, &args)
	fmt.Printf("%v", rsp)

}
