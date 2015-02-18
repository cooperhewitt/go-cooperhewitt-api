package api

import (
       "io/ioutil"
       "net/http"
       "net/url"
       "fmt"
)

type API struct {
     token string
     endpoint string
}

func New(token string) *API {

     return &API{
     	    token: token,
	    endpoint: "https://api.collection.cooperhewitt.org/rest/",
     }
}

function ExecuteMethod (api *API, method string, params *url.Values) ([]byte error) {

	params.Set("method",  method)
	params.Set("access_token", api.token)

	req, err := http.NewRequest("POST", api.endpoint, params)

	client := &http.Client{}
	rsp, err := client.Do(req)

	if err != nil {
        	panic(err)
	}

	defer rsp.Body.Close()

	fmt.Println("response Status:", rsp.Status)
	fmt.Println("response Headers:", rsp.Header)

	body, _ := ioutil.ReadAll(resp.Body)
	return body
}


