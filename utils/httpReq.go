package utils

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/imroc/req"
	newReq "github.com/imroc/req/V3"
	"github.com/rs/zerolog/log"
)

func HttpPost(url string, header req.Header, requestBody interface{}) (*req.Resp, error) {

	var response *req.Resp
	var err error
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	response, err = req.Post(url, req.BodyJSON(&requestBody), header, client)
	if err != nil {
		return nil, errors.New("Func: utils.HttpPost. Message: " + err.Error())
	}
	return response, nil
}

func HttpGet(url string, header req.Header, param interface{}) (*req.Resp, error) {

	var response *req.Resp
	var err error

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	response, err = req.Get(url, param, header, client)
	if err != nil {
		return nil, errors.New("Func: utils.HttpGet. Message: " + err.Error())
	}

	return response, nil
}

func NewHttpPost(url string, header req.Header, requestBody interface{}) (*req.Resp, error) {

	client := newReq.C().DevMode()

	resp, err := client.R().
		Post(url)
	if err != nil {
		log.Fatal(err)
	}

	if !resp.IsSuccessState() {
		fmt.Println("bad response status:", resp.Status)
		return nil, errors.New("Func: utils.NewHttpPost. Message: " + resp.Status)
	}

	return resp, nil
}
