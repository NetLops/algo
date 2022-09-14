package main

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maximumSwap(rand.Int())
	}
}

func BenchmarkName2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maximumSwap1(rand.Int())
	}
}
