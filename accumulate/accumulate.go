package accumulate

func accumulate(collection []string, operation func(string) string) []int {
  final_results := make([]int, len(collection))

  for _, val := range(collection) {
    item := operation(val)
    final_results = append(final_results, item)
  }

  return final_results
}
