package utils

import (
	"strings"

	"github.com/emirpasic/gods/sets/hashset"
)

func ToStringSet(s string) *hashset.Set {
	ret := hashset.New()
	if s == "" {
		return ret
	}
	for _, s2 := range strings.Split(s, ",") {
		ret.Add(s2)
	}
	return ret
}

func SliceToStringSet(s []string) *hashset.Set {
	ret := hashset.New()
	if len(s) == 0 {
		return ret
	}
	for _, s2 := range s {
		ret.Add(s2)
	}
	return ret
}
