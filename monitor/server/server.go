package server

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"monitor/common"
	"monitor/model"
	"monitor/msg"

	jsoniter "github.com/json-iterator/go"
	logger "github.com/sirupsen/logrus"
)

var MSG,URL string

func GetJenkinsData(method, urlVal,data string) {

	client := &http.Client{}
	var req *http.Request

	if data == "" {
		urlArr := strings.Split(urlVal,"?")
		if len(urlArr)  == 2 { urlVal = urlArr[0] + "?" + (urlArr[1])
		}
		req, _ = http.NewRequest(method, urlVal, nil)
	}else {
		req, _ = http.NewRequest(method, urlVal, strings.NewReader(data))
	}

	cookie1 := &http.Cookie{
		Name: "remember-me",
		Value: "eWlubG9uZzoxNjI5ODc5NTI4MTc5OjhlOTM1NTEyMjU4NjlhODgzZjQyYTlkNDE5NTA2MTg0NjZmNWVhNzkzZGI4ODMyMDc5OTdkYzZhNjM5MWY3MTY",
	}

	req.AddCookie(cookie1)

	resp, err := client.Do(req)
	if err != nil{
		log.Fatal(err)
	}

	jksdata, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(jksdata))

	Jenkinsmsg := &model.Jenkins{}
	err = jsoniter.UnmarshalFromString(string(jksdata), &Jenkinsmsg)
	if err != nil {
		fmt.Println(err)
		return
	}

	feishumsg, err := msg.Init_Msg()
	if err != nil {
		logger.Println(err)
		return
	}

	if Jenkinsmsg.Description == "" {
		Jenkinsmsg.Description = Jenkinsmsg.Actions[0].Causes[0].UserName
	}

	Jenkinsmsg.BuiltOn = MSG

	feishumsg.Card.Elements[0].Text.Content =  "构建总数:" + Jenkinsmsg.ID + "\n" + "构建结果:" + Jenkinsmsg.Result+ "\n" + "构建人:" + Jenkinsmsg.Description + "\n" +"URL:" + Jenkinsmsg.URL + "\n" + "Msg:" + Jenkinsmsg.BuiltOn

	toString, err := jsoniter.MarshalToString(&feishumsg)
	if err != nil {
		logger.Errorf("err :%v", err)
		return
	}

	fmt.Println(bytes.NewBufferString(toString))

	request, err := http.NewRequest("POST", common.WEBHOOK, bytes.NewBufferString(toString))
	if err != nil {
		logger.Errorf("err:%v", err)
	}

	request.Header.Add("Content-Type", "application/json")

	rep, err := client.Do(request)
	if err != nil {
		logger.Errorf("request err:%v", err)
		return
	}
	defer rep.Body.Close()
}

func InputArg() {

	name := flag.String("name", "", "地址名")
	msg := flag.String("msg", "", "信息")

	flag.Parse()

	fmt.Println(*name)
	fmt.Println(*msg)

	URL = *name
	MSG = *msg
}


