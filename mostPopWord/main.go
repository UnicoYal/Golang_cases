package mostpopword

func MostPopularWord(words []string) string {
	max_counter := 0
	res_word := "" 
	my_map := make(map[string]int)
	for _, word := range words {
		my_map[word] += 1
		if my_map[word] > max_counter {
			max_counter = my_map[word]
			res_word = word
		}
	}
	return res_word
}