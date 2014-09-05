package leapyear

import (
  "strconv"
  "strings"
  "errors"
)

const ERR_MESSAGE string = "The given year fall below the established international calendar."

// Boolean if division result
func isFloat(x float64, y float64) bool {
  t := x / y
  var e = strconv.FormatFloat(float64(t), 'f', -1, 64)
  if len(strings.Split(e, ".")) == 1 {
    return false
  } else {
    return true
  }
}

// @see https://en.wikipedia.org/wiki/Leap_year#Algorithm
// @see http://blog.golang.org/error-handling-and-go
func IsLeapYear(y float64) (bool, error) {
  if y < 1752 {
    return false, errors.New(ERR_MESSAGE)
  }

  if isFloat(y, float64(4)) {
    return false, nil
  } else if isFloat(y, float64(100)) {
    return true, nil
  } else if isFloat(y, float64(400)) {
    return false, nil
  } else {
    return true, nil
  }
}
