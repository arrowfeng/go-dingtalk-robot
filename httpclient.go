/*
 * File : httpclient.go
 * Author : arrowfeng
 * Date : 2019/11/13
 */
package main

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"log"
)


func send(data map[string]interface{}) error {

	request := gorequest.New()

	log.Println(data)

	resp, body, _ := request.Post(url + "?access_token=" + access_token).
		Send(data).
		End()

	log.Println(resp)

	var response Response

	if err := json.Unmarshal([]byte(body), &response); err != nil {
		return err
	}

	if response.errcode == 0 {
		log.Println("send successfully!")
	} else {
		log.Printf("send unsuccessfully, msg: %s", response.errmsg)
	}

	return nil
}
