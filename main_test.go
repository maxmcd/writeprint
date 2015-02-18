package main

import (
	"fmt"
	"testing"
)

func TestMatchingStrength(t *testing.T) {
	var wins float64
	var losses float64
	userMap := readData()
	userMap = runAnalysis(userMap)

	for _, user := range userMap {
		for _, message := range user.Messages {
			if len(message.Content) > 0 {
				guessedUser := guessUser(message.Content, userMap)
				if guessedUser.Email == user.Email {
					wins = wins + 1
				} else {
					losses = losses + 1
				}
			}
		}
	}
	percentage := (wins / (wins + losses)) * float64(100)
	fmt.Println("There were", wins, "wins and", losses, "losses")
	fmt.Println(percentage, "percent correct")
	if percentage < 30 {
		t.Fail()
	}

}
