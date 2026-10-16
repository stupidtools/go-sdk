package hpa

import (
	"math/rand"
	"strconv"
	"time"
)

type RandomIdProvider struct {
	random *rand.Rand
}

func NewRandomIdProvider() *RandomIdProvider {
	return &RandomIdProvider{random: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func (r *RandomIdProvider) GetRequestId() int {
	return 100000 + r.random.Intn(999999)
}

func (r *RandomIdProvider) GetRequestIdString() string {
	return strconv.Itoa(r.GetRequestId())
}
