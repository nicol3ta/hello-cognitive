package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//FaceResponse is a JSON structure converted into a Go type and represents the detected face within a picture
type FaceResponse []struct {
	FaceID        string `json:"faceId"`
	FaceRectangle struct {
		Top    int `json:"top"`
		Left   int `json:"left"`
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"faceRectangle"`
	FaceAttributes struct {
		Gender  string  `json:"gender"`
		Age     float64 `json:"age"`
		Glasses string  `json:"glasses"`
	} `json:"faceAttributes"`
}

//EmotionResponse is a JSON structure converted into a Go type and represents the detected emotions within a picture
type EmotionResponse []struct {
	FaceRectangle struct {
		Height int `json:"height"`
		Left   int `json:"left"`
		Top    int `json:"top"`
		Width  int `json:"width"`
	} `json:"faceRectangle"`
	Scores struct {
		Anger     float64 `json:"anger"`
		Contempt  float64 `json:"contempt"`
		Disgust   float64 `json:"disgust"`
		Fear      float64 `json:"fear"`
		Happiness float64 `json:"happiness"`
		Neutral   float64 `json:"neutral"`
		Sadness   float64 `json:"sadness"`
		Surprise  float64 `json:"surprise"`
	} `json:"scores"`
}

func detectEmotions() {
	var jsonStr = []byte(`{"url":"https://upload.wikimedia.org/wikipedia/commons/thumb/1/19/Bill_Gates_June_2015.jpg/339px-Bill_Gates_June_2015.jpg"}`)
	emotionURL := fmt.Sprintf("https://westus.api.cognitive.microsoft.com/emotion/v1.0/recognize")

	req, err := http.NewRequest("POST", emotionURL, bytes.NewBuffer(jsonStr))

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Ocp-Apim-Subscription-Key", "Your key goes here")

	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)

	if err != nil {

		log.Fatal("Do: ", err)
		return

	}

	defer resp.Body.Close()
	var record EmotionResponse

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Fatal("Decode: ", err)
		return
	}

	fmt.Printf("\nThe anger score of the identified person is: %v\n", record[0].Scores.Anger)
	fmt.Printf("The happiness score of the identified person is: %v\n", record[0].Scores.Happiness)
	fmt.Printf("The neutral score of the identified person is: %v\n", record[0].Scores.Neutral)
	fmt.Printf("The contempt score of the identified person is: %v\n", record[0].Scores.Contempt)
	fmt.Printf("The disgust score of the identified person is: %v\n", record[0].Scores.Disgust)
	fmt.Printf("The fear score of the identified person is: %v\n", record[0].Scores.Fear)
	fmt.Printf("The sadness score of the identified person is: %v\n", record[0].Scores.Sadness)
	fmt.Printf("The surprise score of the identified person is: %v\n", record[0].Scores.Surprise)
}

func detectFaces() {
	var jsonStr = []byte(`{"url":"https://upload.wikimedia.org/wikipedia/commons/thumb/1/19/Bill_Gates_June_2015.jpg/339px-Bill_Gates_June_2015.jpg"}`)
	faceURL := fmt.Sprintf("https://westus.api.cognitive.microsoft.com/face/v1.0/detect?returnFaceId=true&returnFaceLandmarks=false&returnFaceAttributes=age,gender,glasses")

	req, err := http.NewRequest("POST", faceURL, bytes.NewBuffer(jsonStr))

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Ocp-Apim-Subscription-Key", "Your key goes here")

	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)

	if err != nil {

		log.Fatal("Do: ", err)
		return

	}

	defer resp.Body.Close()
	var record FaceResponse

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Fatal("Decode: ", err)
		return
	}

	fmt.Printf("The gender of the identified person is: %v\n", record[0].FaceAttributes.Gender)
	fmt.Printf("Estimated age is: %v years.\n", record[0].FaceAttributes.Age)
	fmt.Printf("Does the person wear glasses? %v", record[0].FaceAttributes.Glasses)
}

func main() {

	// Concurrency using goroutines
	go detectFaces()
	go detectEmotions()

	var input string
	fmt.Scanln(&input)
}
