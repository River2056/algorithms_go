package timeconversion

import (
	"fmt"
	"strconv"
)

func TimeConversion(s string) string {
	// Write your code here
	var resultString string
	indicator := s[len(s)-2:]
	hour := s[:2]
	mmss := s[3 : len(s)-2]
	if indicator == "PM" {
		h, _ := strconv.Atoi(hour)
		if h == 12 {
			resultString = fmt.Sprintf("%v:%s\n", h, mmss)
		} else {
		    h += 12
			if h < 10 {
				resultString = fmt.Sprintf("%v%v:%s\n", "0", h, mmss)
			} else {
				resultString = fmt.Sprintf("%v:%s\n", h, mmss)
			}
		}
	} else if indicator == "AM" {
		h, _ := strconv.Atoi(hour)

		if h == 12 {
			resultString = fmt.Sprintf("%v:%s\n", "00", mmss)
		} else {

			if h < 10 {
				resultString = fmt.Sprintf("%v%v:%s\n", "0", h, mmss)
			} else {
				resultString = fmt.Sprintf("%v:%s\n", h, mmss)
			}
		}
	}
	fmt.Printf("result: %v\n", resultString)
	return resultString
}
