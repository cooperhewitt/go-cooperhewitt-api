package api

import (
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type APIClient struct {
	scheme   string
	isa      string
	Host     string
	Endpoint string
	Token    string
}

// http://blog.golang.org/json-and-go
type APIResponse interface{}

func OAuth2Client(token string) *APIClient {

	return &APIClient{
		isa:      "oauth2",
		scheme:   "https",
		Token:    token,
		Host:     "api.collection.cooperhewitt.org",
		Endpoint: "rest/",
	}
}

func (client *APIClient) ExecuteMethod(method string, params *url.Values) (APIResponse, error) {

	url := "https://" + client.Host + "/" + client.Endpoint
	// fmt.Println(url)

	params.Set("method", method)
	params.Set("access_token", client.Token)

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
