package utils

import (
	"bytes"
	"fmt"
	"gin-demo/defs"
	"gin-demo/logger"
	"io/ioutil"
	"net/http"
)

// var HTTPPost

var client = &http.Client{}

//HTTPGet send http request
func HTTPGet(url string) (string, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(defs.HTTPRequestErr, err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(defs.ReadRespBodyErr, err)
		return "", err
	}

	return string(body), err
}

// HTTPPost send post request
func HTTPPost(postURL string, paramJSON string) (string, error) {

	var jsonStr = []byte(paramJSON)
	req, err := http.NewRequest("POST", postURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(defs.HTTPRequestErr, err)
		return "", err
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(defs.ReadRespBodyErr, err)
		return "", err
	}
	fmt.Println("response Body:", string(body))
	return string(body), err
}
