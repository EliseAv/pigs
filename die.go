package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math"
)

type Threshold struct {
	amount int
	value  uint64
}

type Die struct {
	sides         uint64
	state         uint64
	threshold     Threshold
	buffer, queue []int
}

func NewDie(sides int) *Die {
	// Generate a cryptographically secure seed for the RNG
	buffer := make([]byte, 8)
	_, err := rand.Read(buffer)
	if err != nil {
		log.Panic(err)
	}
	state := binary.LittleEndian.Uint64(buffer)
	die := &Die{sides: uint64(sides), state: state}

	// Store the threshold to avoid bias
	die.threshold = die.getThreshold(die.sides)
	die.buffer = make([]int, die.threshold.amount)

	return die
}

func (Die) getThreshold(sides uint64) Threshold {
	var t Threshold
	// ln(2**64)~==44.36, I'm rounding it down a bit to reduce RNG re-rolls.
	t.amount = int(40 / math.Log(float64(sides)))

	var modulus uint64 = 1
	for i := t.amount; i > 0; i-- {
		modulus *= sides
	}
	t.value = math.MaxUint64 - math.MaxUint64%modulus
	return t
}

func (die *Die) Roll() int {
	for len(die.queue) == 0 {
		die.queue = die.generateRolls()
	}
	roll := die.queue[0]
	die.queue = die.queue[1:]
	return roll + 1
}

func (die *Die) generateRolls() []int {
	random := die.rngNext()
	for random >= die.threshold.value {
		random = die.rngNext()
	}
	// Store roll results in buffer, queue will be a slice from it
	for i := range die.buffer {
		die.buffer[i] = int(random % die.sides)
		random /= die.sides
	}
	return die.buffer
}

func (die *Die) rngNext() uint64 {
	// A 64 bit linear congruential RNG by Donald Knuth
	die.state = die.state*6364136223846793005 + 1442695040888963407
	return die.state
}

func (die Die) String() string {
	return fmt.Sprintf("D%d", die.sides)
}
