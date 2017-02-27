package gigasecond

import (
  "time"
  "fmt"
)

const testVersion = 4

func AddGigasecond(t time.Time) time.Time {
  new_time := t.Add(time.Second * 1000000000)
  return new_time
}
