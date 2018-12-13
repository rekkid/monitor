package web

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
)

func apiHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != http.MethodPost {
		res, _ := json.Marshal(ErrRequestNotRe)
		io.WriteString(w, string(res))
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	apiBody := &ApiBody{}
	if err := json.Unmarshal(res, apiBody); err != nil {
		res, _ = json.Marshal(ErrRequestBodyParseFailed)
		io.WriteString(w, string(res))
		return
	}

	request(apiBody, w, r)
	defer r.Body.Close()
}
