package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func SearchMovie(c *gin.Context) {
	apiEndpoint := "https://api.themoviedb.org/3/search/movie"
	params := url.Values{}
	params.Add("query", c.Query("moviename"))

	url := apiEndpoint + "?" + params.Encode()
	fmt.Println("FULL URL: ", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJlYjNlNWNkNjI5MzIxYjc2NjVkMGY0NWIwMTM5NWY2OCIsIm5iZiI6MTczOTAxNTc1NS44NCwic3ViIjoiNjdhNzQ2NGI1ZmE0MmQ3ZTc2ZjEwZTY3Iiwic2NvcGVzIjpbImFwaV9yZWFkIl0sInZlcnNpb24iOjF9.pZ9O3p5eGkyjlqra7RNp7OY5hsvjwvxb1GxBz2Pxf-k")
	req.Header.Add("accept", "application/json")
	fmt.Println("HEADERS: ", req.Header)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	var jsonObject interface{}

	err = json.Unmarshal(body, &jsonObject)
	fmt.Println("Response Status:", resp.Status)

	c.JSON(http.StatusOK, jsonObject)
}
