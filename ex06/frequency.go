package letter

type FreqMap map[rune]int

func Frequency(text string) FreqMap {
	res := FreqMap{}
	for _, i := range text {
		res[i]++
	}
	return res
}

func ConcurrentFrequency(text []string) FreqMap {
	// map freq over text
	ch := make(chan FreqMap, len(text))
	for _, wordList := range text {
		go func(word string) {
			ch <- Frequency(word)
		}(wordList)
	}
	// all results into one map
	res := FreqMap{}
	for range text {
		for key, val := range <-ch {
			res[key] += val
		}
	}
	return res
}
