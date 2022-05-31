package config

import (
	"basic-api/entity"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// AppResponse is for response config show to Frontend side
func AppResponse(data interface{}) *entity.SuccessResponse {
	return &entity.SuccessResponse{
		Success: true,
		Data:    data,
		Error:   nil,
	}
}

//RequestURL is to call endpoint from another source
func RequestURL(url string) (string, error) {
	// start the new request
	req, errGet := http.NewRequest("GET", url, nil)
	// add header
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:81.0) Gecko/20100101 Firefox/81.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	if errGet != nil {
		log.Fatalln(errGet)
		return "", errGet
	}
	// start get client
	client := &http.Client{}
	// start the request
	resp, errDo := client.Do(req)
	if errDo != nil {
		fmt.Println(errDo)
		return "", errDo
	}
	defer resp.Body.Close()
	// read the response body
	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		log.Fatalln(errRead)
		return "", errRead
	}
	//Convert the body to type string
	return string(body), nil
}

// Contains is the function find element in array slice
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
