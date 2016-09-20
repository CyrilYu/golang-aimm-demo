package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/kataras/iris"
)

var apiURL = "https://aimm.diuit.net"

func main() {
	iris.Post("/login", login)
	iris.Post("/pair", pair)
	iris.Post("/leave", leave)
	iris.Listen(":8080")
}

func getAPIURL(host, path string) string {
	u, _ := url.ParseRequestURI(host)
	u.Path = path
	urlStr := fmt.Sprintf("%v", u)
	return urlStr
}

func login(ctx *iris.Context) {
	// get nonce
	userID := ctx.PostValue("userId")
	deviceID := ctx.PostValue("deviceId")
	platform := ctx.PostValue("platform")

	log.Println(userID)
	var result map[string]interface{}
	loginAPI := getAPIURL(apiURL, "/users/login")
	data := url.Values{}
	data.Add("userId", userID)
	data.Add("deviceId", deviceID)
	data.Add("platform", platform)
	req, _ := http.NewRequest("POST", loginAPI, bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	log.Println(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		ctx.Text(200, "GG")
		return
	}
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}
	ctx.JSON(resp.StatusCode, result)

}

func pair(ctx *iris.Context) {
	userID := ctx.PostValue("userId")
	sessionToken := ctx.PostValue("sessionToken")
	var result map[string]interface{}
	pairAPI := getAPIURL(apiURL, "/users/pair")
	data := url.Values{}
	data.Add("userId", userID)
	data.Add("sessionToken", sessionToken)
	req, _ := http.NewRequest("POST", pairAPI, bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}
	ctx.JSON(resp.StatusCode, result)
}

func leave(ctx *iris.Context) {
	chatID := ctx.PostValue("chatId")
	sessionToken := ctx.PostValue("sessionToken")
	var result map[string]interface{}
	leaveAPI := getAPIURL(apiURL, "/users/leave")
	data := url.Values{}
	data.Add("chatId", chatID)
	data.Add("sessionToken", sessionToken)
	req, _ := http.NewRequest("POST", leaveAPI, bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}
	ctx.JSON(resp.StatusCode, result)
}
