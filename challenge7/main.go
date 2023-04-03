package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const url = "https://jsonplaceholder.typicode.com"

type Post struct {
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type WaterWind struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		userID := rand.Intn(100) + 1
		statusWater := rand.Intn(100) + 1
		statusWind := rand.Intn(100) + 1

		post := Post{
			UserID: userID,
			Title:  "Test Post",
			Body:   "This is a test",
		}

		CreatePost(post)

		statusWaterWind(statusWater, statusWind)

		fmt.Println()

		time.Sleep(15 * time.Second)
	}
}

func statusWaterWind(statusWater, statusWind int) {
	var water string
	var wind string

	if statusWater < 5 {
		water = "aman"
	} else if statusWater >= 6 && statusWater <= 8 {
		water = "siaga"
	} else {
		water = "bahaya"
	}

	if statusWind < 6 {
		wind = "aman"
	} else if statusWind >= 7 && statusWind <= 15 {
		wind = "siaga"
	} else {
		wind = "bahaya"
	}

	waterWind := WaterWind{
		Water: statusWater,
		Wind:  statusWind,
	}
	requestJSON, err := json.Marshal(waterWind)
	if err != nil {
		fmt.Println("Failed to marshal request data:", err)
		return
	}

	result := fmt.Sprint(string(requestJSON), "\n", "status water : ", water, "\n", "status wind : ", wind)

	fmt.Println(result)
}

func CreatePost(post Post) {
	requestJSON, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Failed to marshal request data:", err)
		return
	}

	res, err := http.Post("https://jsonplaceholder.typicode.com/posts", "Application/json", bytes.NewBuffer(requestJSON))
	if err != nil {
		fmt.Println("Failed to initialize create post:", err)
		return
	}

	defer res.Body.Close()

	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	result := fmt.Sprint("POST Request:", string(requestJSON), "\n", "Response:\n", string(bodyByte), "\n", "Response code:", res.Status)

	fmt.Println(result)
}
