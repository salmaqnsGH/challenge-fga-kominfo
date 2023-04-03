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

// ●Jika water dibawah 5 maka status aman
// ●jika water antara 6 - 8 maka status siaga
// ●jika water diatas 8 maka status bahaya
// ●jika wind dibawah 6 maka status aman
// ●jika wind antara 7 - 15 maka status siaga
// ●jika wind diatas 15 maka status bahaya
// ●value water dalam satuan meter
// ●value wind dalam satuan meter per detik

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

		result := CreatePost(post)
		fmt.Println(result)

		statusWaterWind(statusWater, statusWind)

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

func CreatePost(post Post) string {
	requestJSON, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Failed to marshal request data:", err)
		return err.Error()
	}

	res, err := http.Post("https://jsonplaceholder.typicode.com/posts", "Application/json", bytes.NewBuffer(requestJSON))
	if err != nil {
		fmt.Println("Failed to initialize create post:", err)
		return err.Error()
	}

	defer res.Body.Close()

	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return err.Error()
	}

	result := fmt.Sprint("POST Request:", string(requestJSON), "\n", "Response:\n", string(bodyByte), "\n", "Response code:", res.Status)

	return result
}
