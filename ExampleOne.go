/*
File Name:  Main.go
Copyright:  2022
Author:     Emir Skopljak
*/

package main

import (
	"errors"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// StoryStats struct object
type StoryStats struct {
	ShortestWord      string
	LongestWord       string
	AverageWordLength float64
	AverageWords      []string
}

// testValidity is function to check validity of the input format string
// @var string Input
// @return (bool, []string)
// estimated time: 10 minutes
// used time: ~15 minutes
// Task 1
func testValidity(Input string) (validity bool, parts []string) {
	parts = strings.Split(Input, "-")

	if len(parts) < 2 && len(parts)%2 != 0 {
		return false, parts
	}

	for i := 0; i < len(parts); i++ {
		_, err := strconv.ParseUint(parts[i], 10, 32)
		if (i%2 == 0 && err != nil) || (i%2 != 0 && err == nil) {
			return false, parts
		}
	}
	return true, parts
}

// averageNumber is function to calculate average number of the input format string
// @var string Input
// @return (float, error)
// estimated time: 10 minutes
// used time: ~15 minutes
// Task 1
func averageNumber(Input string) (averageNum float64, err error) {
	val, parts := testValidity(Input)
	if !val {
		return -1, errors.New("invalid format")
	}

	average := 0
	for i := 0; i < len(parts); i++ {
		if i%2 == 0 {
			num, _ := strconv.Atoi(parts[i])
			average += num
		}
	}

	result := float64(average) / (float64(len(parts)) / 2.0)
	return result, nil
}

// wholeStory is function to create new string from the input format string that only has string parts
// @var string Input
// @return (string, error)
// estimated time: 20 minutes
// used time: ~20 minutes
// Task 1
func wholeStory(Input string) (story string, err error) {
	val, parts := testValidity(Input)
	if !val {
		return "-", errors.New("invalid format")
	}

	for i := 0; i < len(parts); i++ {
		if i%2 != 0 {
			if i > 1 {
				story += " "
			}
			story += parts[i]
		}
	}

	return story, nil
}

// storyStats is function to return StoryStats object of the input format string
// @var string Input
// @return (StoryStats, error)
// estimated time: 20 minutes
// used time: ~25 minutes
// Task 1
func storyStats(Input string) (story StoryStats, err error) {
	val, parts := testValidity(Input)
	if !val {
		return story, errors.New("invalid format")
	}
	average := 0
	for i := 0; i < len(parts); i++ {
		if i%2 != 0 {
			if i == 1 {
				story.LongestWord = parts[i]
				story.ShortestWord = parts[i]
			} else if len(story.ShortestWord) > len(parts[i]) {
				story.ShortestWord = parts[i]
			} else if len(story.LongestWord) < len(parts[i]) {
				story.LongestWord = parts[i]
			}
			average += len(parts[i])
		}
	}
	story.AverageWordLength = float64(average) / (float64(len(parts)) / 2)

	for i := 0; i < len(parts); i++ {
		if i%2 != 0 && (len(parts[i]) == int(math.Ceil(story.AverageWordLength)) || len(parts[i]) == int(math.Floor(story.AverageWordLength))) {
			story.AverageWords = append(story.AverageWords, parts[i])
		}
	}
	return story, nil
}

// generateRandomString is function that will generate correct or wrong string format based on flag
// @var bool flagValidity
// @return string
// estimated time: 25 minutes
// used time: ~20 minutes
// Task 2
func generateRandomString(flagValidity bool) string {
	if flagValidity {
		return generateRandomCorrectString()
	} else {
		return generateRandomInvalidString()
	}
}

// generateRandomCorrectString generates random correct string
// @var
// @return string
// Task 2
func generateRandomCorrectString() (result string) {
	limitation := rand.Intn(50) + 1
	for i := 0; i < limitation; i++ {
		if i > 0 {
			result += "-"
		}
		result += strconv.Itoa(randomNumber(rand.Intn(10)+1)) + "-" + randomString(rand.Intn(10)+1)
	}
	return result
}

// generateRandomInvalidString generates random invalid string
// @var
// @return string
// Task 2
func generateRandomInvalidString() (result string) {
	lengthRandom := rand.Intn(100)
	result = randomString(lengthRandom)
	return result
}

// randomString generate random string of predefined length
// @var int n
// @return string
// Task 2
func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

// randomNumber generate random number of predefined length
// @var int n
// @return int
// Task 2
func randomNumber(n int) int {
	return int(math.Pow10(n)) + rand.Intn(int(math.Pow10(n+1))-int(math.Pow10(n))-1)
}
