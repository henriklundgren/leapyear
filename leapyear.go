package leapyear

import (
  "strconv"
  "strings"
)

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
func IsLeapYear(y float64) bool {
  if isFloat(y, float64(4)) {
    return false
  } else if isFloat(y, float64(100)) {
    return true
  } else if isFloat(y, float64(400)) {
    return false
  } else {
    return true
  }
}
