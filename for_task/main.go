package for_task

func Map(strs []string, mapFunc func(s string) string) []string {
	mapped := make([]string, len(strs))
	for i, s := range strs {
		mapped[i] = mapFunc(s)
	}

	return mapped
}