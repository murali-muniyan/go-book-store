package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseRequestBody(req *http.Request, x interface{}) error {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, x)

	return err
}

func WiteToResponse(wtr http.ResponseWriter, statusCode int, x interface{}) {
	wtr.WriteHeader(statusCode)

	if x == nil {
		return
	}

	respBody, _ := json.Marshal(x)

	wtr.Write(respBody)
}

func GetID(req *http.Request) int64 {
	vars := mux.Vars(req)

	id, _ := strconv.Atoi(vars["id"])

	return int64(id)
}

func HandleErr(err error, wtr http.ResponseWriter) {
	wtr.WriteHeader(http.StatusBadRequest)

	if err == nil {
		return
	}

	resp, _ := json.Marshal(&errResp{Err: err.Error()})

	wtr.Write(resp)
}

type errResp struct {
	Err string `json:"error"`
}
