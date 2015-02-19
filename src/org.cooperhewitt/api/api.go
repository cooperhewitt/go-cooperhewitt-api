package api

import (
       "io/ioutil"
       "net/http"
       "net/url"
       "encoding/json"
       _ "fmt"
)

type APIClient struct {
     token string
     scheme string
     host string
     endpoint string
     isa string
}

// http://blog.golang.org/json-and-go
type APIResponse interface {}

func OAuth2Client(token string) *APIClient {

     return &APIClient{
     	    token: token,
	    scheme: "https",
	    host: "api.collection.cooperhewitt.org",
	    endpoint: "rest/",
	    isa: "oauth2",
     }
}

func (client *APIClient) ExecuteMethod (method string, params *url.Values) (APIResponse, error) {

        url := "https://" + client.host + "/" + client.endpoint

	params.Set("method",  method)
	params.Set("access_token", client.token)

	http_req, err := http.NewRequest("POST", url, nil)
	http_req.URL.RawQuery = (*params).Encode()

        http_req.Header.Add("Accept-Encoding", "gzip")

	http_client := &http.Client{}
	http_rsp, err := http_client.Do(http_req)

	if err != nil {
        	panic(err)
	}

	defer http_rsp.Body.Close()

	/*
	fmt.Println("response Status:", rsp.Status)
	fmt.Println("response Headers:", rsp.Header)
	*/

	http_body, _ := ioutil.ReadAll(http_rsp.Body)

	var rsp APIResponse
	json.Unmarshal(http_body, &rsp)

	return rsp, nil
}


