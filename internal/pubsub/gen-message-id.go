package pubsub

import (
	"math/rand/v2"
	"strconv"
)

func genMessageID() string {
	min := 100000000000000
	max := 999999999999999
	return strconv.Itoa((rand.IntN(max-min) + min))
}
