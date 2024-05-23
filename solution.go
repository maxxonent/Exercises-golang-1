package kata

import(
  "fmt"
  "strings"
  "strconv"
  "sort"
)

type Time struct {
  h, m, s int
}

func (t Time) String() string {
	return fmt.Sprintf("%02d|%02d|%02d", t.h, t.m, t.s)
}

func parseTime(s string) Time {
  fields := strings.Split(s, "|")
  var numbers [3]int
  for i:= 0; i < 3; i++ {
    numbers[i], _ = strconv.Atoi(fields[i])
  }
  return Time{numbers[0], numbers[1], numbers[2]}
}

func (t Time) Secs() int {
  return t.h * 3600 + t.m * 60 + t.s
}

func FromSecs(secs int) (result Time) {
  result.h = secs / 3600; secs -= result.h * 3600
  result.m = secs / 60; secs -= result.m * 60
  result.s = secs
  return
}

func (t Time) add(t2 Time) Time {
  sumSecs := t.Secs() + t2.Secs()
  return FromSecs(sumSecs)
}

func Sum(times []Time) (result Time) {
  for _, t := range times {
    result = result.add(t)
  }
  return
}

func (t Time) Divide(divider int) Time {
  secs := t.Secs()
  secs /= divider
  return FromSecs(secs)
}

func Average(times []Time) Time {
  sum := Sum(times)
  return sum.Divide( len(times) )
}

func Range(times []Time) Time {
  sort.Slice(times, func (i, j int) bool {return times[i].Secs() < times[j].Secs()})
  min := times[0].Secs()
  max := times[len(times)-1].Secs()
  return FromSecs(max-min)
}

func Median(times []Time) Time {
  sort.Slice(times, func (i, j int) bool {return times[i].Secs() < times[j].Secs()})
  length := len(times)
  if length % 2 == 1 {
    return times[length / 2]
  } else {
    return Average( times[length/2-1 : length/2] )
  }
}

func Stati(strg string) string { 
  if strg == "" {return ""}

  strg = strings.ReplaceAll(strg, ",", "")
  
  fields := strings.Fields(strg)
  
  times := make([]Time, len(fields))
  for i, field := range fields {
    times[i] = parseTime(field)
  }  
  
  return fmt.Sprintf("Range: %v Average: %v Median: %v", Range(times), Average(times), Median(times))
}
