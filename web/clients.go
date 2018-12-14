package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"monitor/util/config"
	"net/http"
	"net/url"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func request(body *ApiBody, w http.ResponseWriter, r *http.Request) {

	url, _ := url.Parse(body.Url)
	url.Host = config.GetHostAddr() + ":" + url.Port()
	newUrl := url.String()

	switch body.Method {
	case http.MethodGet:
		req, _ := http.NewRequest("GET", newUrl, nil)
		req.Header = r.Header
		response, err := httpClient.Do(req)
		if err != nil {
			log.Errorf("httpClient error: %v", err)
			return
		}
		normalResponse(w, response)

	case http.MethodPost:
		req, _ := http.NewRequest("POST", newUrl, bytes.NewBuffer([]byte(body.ReqBody)))
		req.Header = r.Header
		response, err := httpClient.Do(req)
		if err != nil {
			log.Errorf("httpclient error: %v", err)
			return
		}
		normalResponse(w, response)
	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "bad api request")
	}
}

func normalResponse(w http.ResponseWriter, r *http.Response) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response, _ := json.Marshal(ErrInternalFults)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, string(response))
		return
	}

	w.WriteHeader(r.StatusCode)
	io.WriteString(w, string(res))
}
