package accumulate

func Accumulate(collection []string, operation func(string) string) []string {
  var final_results []string

  for _, val := range(collection) {
    item := operation(val)
    final_results = append(final_results, item)
  }

  return final_results
}
