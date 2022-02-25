/*
File Name:  Main_test.go
Copyright:  2022
Author:     Emir Skopljak
*/

package main

import (
	"fmt"
	"strconv"
	"testing"
)

// assertEqual is helper function
// Task 2
func assertEqual(t *testing.T, a interface{}, b interface{}, errorMessage string) {
	if a == b {
		return
	}
	if len(errorMessage) == 0 {
		errorMessage = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(errorMessage)
}

// TestToCheckValidity tests testValidity method
// estimated time: 5 minutes
// used time: ~5 minutes
// Task 2
func TestToCheckValidity(t *testing.T) {
	testFalse, _ := testValidity("ajdg23h9-123bjb1-nb")
	assertEqual(t, false, testFalse, "format of the input is actually okay")

	testTrue, _ := testValidity("23-123bjb1-62-asdasf")
	assertEqual(t, true, testTrue, "invalid input format")
}

// TestToCheckAverageNumber tests averageNumber method
// estimated time: 5 minutes
// used time: ~5 minutes
// Task 2
func TestToCheckAverageNumber(t *testing.T) {
	averageFalse, _ := averageNumber("ajdg23h9-123bjb1-nb")
	assertEqual(t, -1.0, averageFalse, "format of the input is actually okay")

	averageTrue, _ := averageNumber("23-123bjb1-62-asdasf")
	assertEqual(t, 42.5, averageTrue, "invalid input format or result is not 42.5: " + fmt.Sprintf("%f", averageTrue))
}

// TestToCheckWholeStory tests wholeStory method
// estimated time: 5 minutes
// used time: ~5 minutes
// Task 2
func TestToCheckWholeStory(t *testing.T) {
	wholeStoryFalse, _ := wholeStory("ajdg23h9-123bjb1-nb")
	assertEqual(t, "-", wholeStoryFalse, "format of the input is actually okay")

	wholeStoryTrue, _ := wholeStory("23-123bjb1-62-asdasf")
	assertEqual(t, "123bjb1 asdasf", wholeStoryTrue, "invalid input format or result is not '123bjb1 asdasf': " + wholeStoryTrue)
}

// TestToCheckStoryStats tests storyStats method
// estimated time: 10 minutes
// used time: ~8 minutes
// Task 2
func TestToCheckStoryStats(t *testing.T) {
	storyStatsFalse, _ := storyStats("ajdg23h9-123bjb1-nb")
	assertEqual(t, 0, len(storyStatsFalse.AverageWords), "format of the input is actually okay")
	assertEqual(t, "", storyStatsFalse.LongestWord, "format of the input is actually okay")
	assertEqual(t, "", storyStatsFalse.ShortestWord, "format of the input is actually okay")
	assertEqual(t, 0.0, storyStatsFalse.AverageWordLength, "format of the input is actually okay")

	storyStatsTrue, _ := storyStats("23-123bjb1-62-asdasf")
	assertEqual(t, 2, len(storyStatsTrue.AverageWords), "invalid input format or length of avarage words is not 2: " + strconv.Itoa(len(storyStatsTrue.AverageWords)))
	assertEqual(t, "123bjb1", storyStatsTrue.LongestWord, "invalid input format or longest word is not '123bjb1': " + storyStatsTrue.LongestWord)
	assertEqual(t, "asdasf", storyStatsTrue.ShortestWord, "invalid input format or shortest word is not 'asdasf': " + storyStatsTrue.ShortestWord)
	assertEqual(t, 6.5, storyStatsTrue.AverageWordLength, "invalid input format or average word length is not 6.5: " + fmt.Sprintf("%f", storyStatsTrue.AverageWordLength))
}