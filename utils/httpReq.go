package utils

// import (
// 	"errors"
// 	"fmt"

// 	"github.com/imroc/req/v3"
// 	"github.com/rs/zerolog/log"
// )

// func NewHttpPost(url string, requestBody interface{}) (*req.Response, error) {

// 	client := req.C().DevMode()

// 	resp, err := client.R().
// 		Post(url)
// 	if err != nil {
// 		log.Err(err).Msg("Func: utils.NewHttpPost. Message: " + err.Error())
// 	}

// 	if !resp.IsSuccessState() {
// 		fmt.Println("bad response status:", resp.Status)
// 		return nil, errors.New("Func: utils.NewHttpPost. Message: " + resp.Status)
// 	}

// 	return resp, nil
// }
