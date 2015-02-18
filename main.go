package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/kennygrant/sanitize"
	"github.com/maxmcd/writeprint/lexical"
)

type JsonBody struct {
	Messages []JsonMessage
}
type JsonMessage struct {
	Timestamp        int64
	Content          string
	Sender_id        int64
	Sender_email     string
	Sender_full_name string
	Id               int64
}

type User struct {
	Id       int64
	Email    string
	Messages []Message
	Name     string
	A        lexical.Analysis
}

type Message struct {
	Content string
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	userMap := readData()
	testString := "This is a string that I'm writing, how cool is this, I wonder what is going to come of all this."
	userMap = runAnalysis(userMap)
	fmt.Println(guessUser(testString, userMap).Name)

}

func runAnalysis(userMap map[string]User) map[string]User {
	for key, value := range userMap {
		var corpus string
		for _, value := range value.Messages {
			corpus = corpus + value.Content
		}
		value.A.New(corpus)
		userMap[key] = value
	}
	var RatioOfDigitsToNTot float64
	var RatioOfLettersToNTot float64
	var RatioOfUppercaseLettersToNTot float64
	var RatioOfSpacesToNTot float64
	var AverageTokenLengthTot float64
	for _, value := range userMap {
		RatioOfDigitsToNTot += value.A.RatioOfDigitsToN
		RatioOfLettersToNTot += value.A.RatioOfLettersToN
		RatioOfUppercaseLettersToNTot += value.A.RatioOfUppercaseLettersToN
		RatioOfSpacesToNTot += value.A.RatioOfSpacesToN
		AverageTokenLengthTot += value.A.AverageTokenLength
	}
	lexical.RatioOfDigitsToNAvg = RatioOfDigitsToNTot / float64(len(userMap))
	lexical.RatioOfLettersToNAvg = RatioOfLettersToNTot / float64(len(userMap))
	lexical.RatioOfUppercaseLettersToNAvg = RatioOfUppercaseLettersToNTot / float64(len(userMap))
	lexical.RatioOfSpacesToNAvg = RatioOfSpacesToNTot / float64(len(userMap))
	lexical.AverageTokenLengthAvg = AverageTokenLengthTot / float64(len(userMap))

	fmt.Println(lexical.RatioOfDigitsToNAvg)
	fmt.Println(lexical.RatioOfLettersToNAvg)
	fmt.Println(lexical.RatioOfUppercaseLettersToNAvg)
	fmt.Println(lexical.RatioOfSpacesToNAvg)
	fmt.Println(lexical.AverageTokenLengthAvg)
	return userMap
}

func readData() map[string]User {
	userMap := make(map[string]User)
	contents, err := ioutil.ReadFile("data/map.json")
	handle(err)
	json.Unmarshal(contents, &userMap)
	return userMap
}

func guessUser(testString string, userMap map[string]User) User {
	var test lexical.Analysis
	test.New(testString)

	var lowestUser User
	var lowest float64
	lowest = 1000
	for _, user := range userMap {
		score := lexical.Similarities(user.A, test)
		if score < lowest {
			lowest = score
			lowestUser = user
		}
	}
	return lowestUser
}

func formatData() {
	contents, _ := ioutil.ReadFile("data/data.json")
	var jsonBody JsonBody

	userMap := make(map[string]User)

	json.Unmarshal(contents, &jsonBody)
	for _, value := range jsonBody.Messages {
		thing, exists := userMap[value.Sender_email]
		if exists == false {
			userMap[value.Sender_email] = User{
				Id:    value.Sender_id,
				Email: value.Sender_email,
				Name:  value.Sender_full_name,
				Messages: []Message{
					Message{
						Content: sanitize.HTML(value.Content),
					},
				},
			}
		} else {
			thing.Messages = append(
				userMap[value.Sender_email].Messages,
				Message{
					Content: sanitize.HTML(value.Content),
				},
			)
			userMap[value.Sender_email] = thing
		}
	}
	bytes, err := json.Marshal(userMap)
	handle(err)
	// doesn't work when using lexical.Analysis as type
	// in user
	ioutil.WriteFile("data/map.json", bytes, 0644)
}
