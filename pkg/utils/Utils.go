package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

type Dob struct {
	Year  int
	Month int
	Day   int
}

func ParseDate(date string) (Dob, error) {
	separator := regexp.MustCompile(`/`)
	dateArray := separator.Split(date, 3)

	if len(dateArray) == 3 {
		var newDob Dob
		newDob.Day, _ = strconv.Atoi(dateArray[0])
		newDob.Month, _ = strconv.Atoi(dateArray[1])
		newDob.Year, _ = strconv.Atoi(dateArray[2])

		return newDob, nil
	} else {
		return Dob{}, errors.New("Error parsing date")
	}

}
