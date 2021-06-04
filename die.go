package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
)

type Die struct {
	sides                     uint64
	state                     uint64
	thresholds, buffer, queue []int
}

var sidesToThresholds = map[uint64][]int{}

func NewDie(sides int) *Die {
	// Generate a cryptographically secure seed for the RNG
	buffer := make([]byte, 8)
	_, err := rand.Read(buffer)
	if err != nil {
		log.Panic(err)
	}
	state := binary.LittleEndian.Uint64(buffer)
	die := &Die{sides: uint64(sides), state: state}

	// Store the thresholds to avoid bias when our state is a very large number
	die.thresholds = getThresholds(die.sides)
	die.buffer = make([]int, len(die.thresholds))

	return die
}

func getThresholds(sides uint64) []int {
	result, found := sidesToThresholds[sides]
	if found {
		return result
	}
	for i := ^uint64(0); i > 0; i /= sides {
		result = append(result, int(i%sides))
	}
	trimmed := make([]int, len(result))
	copy(trimmed, result)
	sidesToThresholds[sides] = trimmed
	return trimmed
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
	// A 64 bit linear congruential RNG by Donald Knuth
	die.state = die.state*6364136223846793005 + 1442695040888963407
	// Store roll results in buffer
	random := die.state
	for i := range die.thresholds {
		die.buffer[i] = int(random % die.sides)
		random /= die.sides
	}

	// Check thresholds for bias
	for i := len(die.thresholds) - 1; i > 0; i-- {
		if die.buffer[i] < die.thresholds[i] {
			return die.buffer[:i]
		}
	}

	// Wow, we really did roll a -1
	return nil
}

func (die Die) String() string {
	return fmt.Sprintf("d%d", die.sides)
}
