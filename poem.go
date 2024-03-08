package main

import (
	"math/rand"
	"time"
)

var poems = []string{
    "Poem 1",
    "Poem 2",
    "Poem 3",
    // Add more Rumi poems here
}

func getRandomPoem() string {
    rand.Seed(time.Now().UnixNano())
    index := rand.Intn(len(poems))
    return poems[index]
}
