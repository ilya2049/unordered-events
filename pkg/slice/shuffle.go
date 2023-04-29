package slice

import (
	"math/rand"
	"time"
)

func Shuffle[T any](slc []T) {
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(slc), func(i, j int) {
		slc[i], slc[j] = slc[j], slc[i]
	})
}
