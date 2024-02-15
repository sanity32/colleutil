package colleselector

import "strings"

type TraitMap[T comparable] map[T][]string

func (v TraitMap[T]) findAny(s string) (found bool, key T, token string) {
	for k, needles := range v {
		for _, needle := range needles {
			if strings.Contains(s, needle) {
				return true, k, needle
			}
		}
	}
	return
}

func (v TraitMap[T]) findLastOccurance(s string) (found bool, idx int, key T, token string) {
	idxs := map[T]int{}
	values := map[T]string{}
	for k, needles := range v {
	l:
		for _, needle := range needles {
			if lastIdx := strings.LastIndex(s, needle); lastIdx != -1 {
				if !found {
					found = true
				}
				idxs[k] = lastIdx
				values[k] = needle
				// return true, lastIdx, k, needle
				break l
			}
		}
	}
	findMaxIdx := func(m map[T]int) (maxI int, r T) {
		for k, idx := range m {
			if idx > maxI {
				maxI = idx
				r = k
			}
		}
		return
	}
	idx, key = findMaxIdx(idxs)
	token = values[key]
	return
}
