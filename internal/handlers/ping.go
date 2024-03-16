package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type responseBody struct {
	URL  string `json:url`
	Data pong   `json:data`
}

type pong struct {
	Pong string `json:ok`
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{ "pong": "ok" }`))
}

func PongHandler(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("PONG_URL")
	if url == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("env PONG_URL is empty"))
		return
	}
	w.Header().Add("content-type", "application/json")
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var p pong
	if err := json.Unmarshal(responseData, &p); err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	resp := responseBody{
		URL:  url,
		Data: p,
	}

	var respBody []byte
	if respBody, err = json.Marshal(resp); err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}
