# go-cooperhewitt-api

Go language API library for the Cooper Hewitt API.

## Example

Assume this is a file called [echo.go](https://github.com/cooperhewitt/go-cooperhewitt-api/blob/master/bin/echo.go):

	import (
		"flag"
		"fmt"
		"net/url"
		"org.cooperhewitt/api"
		"strings"
	)

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

This would yield:

	$> echo -token ACCESS_TOKEN wub wub wub
	map[echo:wub wub wub method:api.test.echo stat:ok]

## To do

* Setting host and endpoint in constructor
* Support for multipart-mime uploads (just because, you can't actually do those in the API)
* Better error handling (currently freaks out and dies)
* Better internal logging
* Proper documentation

## See also

* https://collection.cooperhewitt.org/api/
