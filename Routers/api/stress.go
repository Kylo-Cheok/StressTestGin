package api

import (
	"ginDemoProject/Controllers"
	"ginDemoProject/Models"
	"github.com/gin-gonic/gin"
	"log"
)

/**
 * @Author: kylo_cheok
 * @Email:  maggic0816@gmail.com
 * @Date:   2022/7/2 14:05
 * @Desc:   Grace under pressure
 */

type Stress struct {
	Num    int    `json:"Num"`
	Count  int    `json:"Count"`
	Url    string `json:"Url"`
	Method string `json:"Method"`
	Body   string `json:"Body"`
}

func Start(c *gin.Context) {
	json := Stress{}
	c.BindJSON(&json)
	log.Printf("request body: %v", &json)
	r := &Models.Request{
		URL:    json.Url,
		Method: json.Method,
		Headers: map[string]string{
			"Content-Type": "application/json;charset=utf-8",
		},
		Body: json.Body,
	}
	controller.StartStress(json.Num, json.Count, r, c)
}
