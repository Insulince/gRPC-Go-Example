package util

import (
	"time"
	"math/rand"
)

type Void struct{} // Used in the bidirectional streaming rpc in the client as a type supplied to a channel to make the channel purely signal-driven. No data is sent through the channel, just an empty struct as a means to signal an event.

const Port = "9000"

const MinProcessingDuration = 1000          // Processing is guaranteed to take AT LEAST this long in milliseconds.
const RandomProcessingDurationWindow = 1000 // Processing will take as long as MinProcessingDuration + a random number of milliseconds less than or equal to this value.

func SimulateProcessing() () {
	time.Sleep(time.Duration(MinProcessingDuration+rand.Intn(RandomProcessingDurationWindow)) * time.Millisecond)
}
