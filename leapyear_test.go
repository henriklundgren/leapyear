package leapyear

import "testing"

// Leapyears
var leaps = [...]int{1804,1808,1812,1816,1820,1824,1828,1832,1836,1840,1844,1848,1852,1856,1860,1864,1868,1872,1876,1880,1884,1888,1892,1896,1904,1908,1912,1916,1920,1924,1928,1932,1936,1940,1944,1948,1952,1956,1960,1964,1968,1972,1976,1980,1984,1988,1992,1996,2000,2004,2008,2012}


func Test_IsFloat(t *testing.T) {
  year := 2014
  divide := 4
  float := isFloat(float64(year), float64(divide))
  if float != true {
    t.Error("Expected true, but got ", float)
  }
}

// Test if (known beforehand) leapyears return true
func Test_IsLeapYear_1(t *testing.T) {

  i := 0
  for i <= len(leaps)-1 {
    if IsLeapYear(float64(leaps[i])) == false {
      t.Error("Known leapyear returns error.")
    }
    i = i + 1
  }
}

// Test if commonyears return false
// TODO: Have a better range for testing.
func Test_IsLeapYear_2(t *testing.T) {

  // Known commonyear
  commonyear := 2014

  if IsLeapYear(float64(commonyear)) == true {
    t.Error("Known commonyear returns error.")
  }
}

