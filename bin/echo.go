package main

import (
       "flag"
       "fmt"
       "org.cooperhewitt/api"
       "net/url"
)

func main() {

     token := flag.String("token", "", "token")
     flag.Parse()

     client := api.OAuth2Client(*token)

     method := "api.test.echo"
     args := url.Values{}
     args.Set("foo", "bar")

     rsp, _ := client.ExecuteMethod(method, &args)
     fmt.Printf("%v", rsp)

}
