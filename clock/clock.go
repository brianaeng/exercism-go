// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import (
    "fmt"
)

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

//Clock struct 
type Clock struct { // Complete the type definition.  Pick a suitable data type.
    hours, minutes int
    time string
}

//New function for Clock struct
func New(hour, minute int) Clock {  
    clock := Clock{hours: hour, minutes: minute, time: ""}
    clock.time = clock.String()
    return clock
}

func (c *Clock) String() string {
    min := c.minutes
    additionalHour := 0

    if min >= 60 {
        for min >= 60 {
            additionalHour++
            min -= 60
        }
    } 

    if min < 0 {
        for min < 0 {
            min += 60
            additionalHour--
        }
    }

    hr := c.hours

    if additionalHour != 0 {
        hr += additionalHour
    }

    if hr >= 24 {
        hr = hr % 12
    }

    for hr < 0 {
        hr = 24 + hr
    }

    finalTime := ""

    c.hours = hr
    c.minutes = min

    if min < 10 && hr < 10 {
        finalTime = fmt.Sprintf("0%v:0%v", hr, min)
    } else if min < 10 {
        finalTime = fmt.Sprintf("%v:0%v", hr, min)
    } else if hr < 10 {
        finalTime = fmt.Sprintf("0%v:%v", hr, min)
    } else {
        finalTime = fmt.Sprintf("%v:%v", hr, min)
    }

    return finalTime
}

func (c Clock) Add(minutes int) Clock {
    c.minutes = c.minutes + minutes

    return c
}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
