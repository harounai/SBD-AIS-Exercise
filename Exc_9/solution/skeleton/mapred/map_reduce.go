package mapred

import (
	"regexp"
	"strings"
	"sync"
)

type MapReduce struct {
}

// todo implement mapreduce

// the map function
func (mr *MapReduce) wordCountMapper(text string) []KeyValue {
	reg := regexp.MustCompile(`[^a-zA-Z]+`)
	clean := reg.ReplaceAllString(strings.ToLower(text), " ")
	words := strings.Fields(clean)

	out := make([]KeyValue, 0, len(words))
	for _, w := range words {
		out = append(out, KeyValue{Key: w, Value: 1})
	}
	return out
}

// reduce function
func (mr *MapReduce) wordCountReducer(key string, values []int) KeyValue {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return KeyValue{Key: key, Value: sum}
}

// shuffle function
func shuffle(mapped [][]KeyValue) map[string][]int {
	groups := make(map[string][]int)
	for _, kvList := range mapped {
		for _, kv := range kvList {
			groups[kv.Key] = append(groups[kv.Key], kv.Value)
		}
	}
	return groups
}

// Run: execute map, shuffle, reduce concurrently
func (mr *MapReduce) Run(input []string) map[string]int {
	var wg sync.WaitGroup
	mapped := make([][]KeyValue, len(input))

	// concurrent map
	for i, line := range input {
		wg.Add(1)
		go func(idx int, text string) {
			defer wg.Done()
			mapped[idx] = mr.wordCountMapper(text)
		}(i, line)
	}

	wg.Wait()

	// shuffle
	groups := shuffle(mapped)

	// reduce
	result := make(map[string]int)
	for key, values := range groups {
		reduced := mr.wordCountReducer(key, values)
		result[reduced.Key] = reduced.Value
	}

	return result

}
