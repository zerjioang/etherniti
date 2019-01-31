// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

package wrapper

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

/*
a wrapper functions to call external services via wrapper
*/
func PostInterface(url string, data interface{}, c echo.Context) ([]byte, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	} else {
		return Post(url, dataBytes, c)
	}
}

func Post(url string, json []byte, c echo.Context) ([]byte, error) {
	if url == "" {
		return nil, errors.New("please, provide a valid endpoint url to make a post request")
	}
	if json == nil || len(json) == 0 {
		return nil, errors.New("there is no valid json payload in the POST request")
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Client", "MaPi")
	req.Header.Set("Accept", c.Request().Header.Get("Accept"))
	req.Header.Set("Authorization", c.Request().Header.Get("Authorization"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	} else {
		//fmt.Println("response Status:", resp.Status)
		//fmt.Println("response Headers:", resp.Header)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		} else {
			resp.Body.Close()
			//fmt.Println("response Body:", string(body))
			return body, nil
		}
	}
}

func Get(url string) ([]byte, error) {
	if url == "" {
		return nil, errors.New("please, provide a valid endpoint url to make a post request")
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Client", "MaPi")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	} else {
		//fmt.Println("response Status:", resp.Status)
		//fmt.Println("response Headers:", resp.Header)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		} else {
			resp.Body.Close()
			//fmt.Println("response Body:", string(body))
			return body, nil
		}
	}
}
