# Leapyear

Return false or true if provided year is a leapyear or not.

## Example

```go
package main

import "github.com/henriklundgren/leapyear"

func main() {
  year := 2012 // Year to test
  isLeapYearOrNot, err := leapyear.IsLeapYear(year) // Returns false or true
}
```

