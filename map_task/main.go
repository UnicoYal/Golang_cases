package maptask

func UniqueUserIDs(userIDs []int64) []int64 {
	// пустая структура struct{} — это тип данных, который занимает 0 байт
	// используется, когда нужно проверять в мапе только наличие ключа
	processed := make(map[int64]struct{})

	uniqUserIDs := make([]int64, 0)
	for _, uid := range userIDs {
		if _, ok := processed[uid]; ok {
			continue
		}

		uniqUserIDs = append(uniqUserIDs, uid)
		processed[uid] = struct{}{}
	}

	return uniqUserIDs
}