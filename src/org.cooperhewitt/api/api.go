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

func Client(token string) *API {

     return &API{
     	    token: token,
	    endpoint: "https://api.collection.cooperhewitt.org/rest/",
     }
}

func (api *API) ExecuteMethod (method string, params *url.Values) ([]byte, error) {

	params.Set("method",  method)
	params.Set("access_token", api.token)

	req, err := http.NewRequest("POST", api.endpoint, nil)
	req.URL.RawQuery = (*params).Encode()

	client := &http.Client{}
	rsp, err := client.Do(req)

	if err != nil {
        	panic(err)
	}

	defer rsp.Body.Close()

	fmt.Println("response Status:", rsp.Status)
	fmt.Println("response Headers:", rsp.Header)

	body, _ := ioutil.ReadAll(rsp.Body)
	return body, nil
}


