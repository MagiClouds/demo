package util

import (
	"fmt"
	"strings"
	"time"
)

var lastTimestamp = time.Now().UnixNano()/1000000
var sequence int64

const MaxSequence = 4096

func CreateOrder(nodeId int64) string {
	currentTimestamp := getCurrentTimestamp()
	if currentTimestamp == lastTimestamp {
		sequence = (sequence + 1) % MaxSequence
		if sequence == 0 {
			currentTimestamp = waitNextMillis(currentTimestamp)
		}
	}else {
			sequence = 0
	}

	orderId := currentTimestamp << 22
	orderId |= nodeId << 10
	orderId |= sequence

	return strings.ToUpper(fmt.Sprintf("%x", orderId) )
}

func getCurrentTimestamp() int64  {
	return time.Now().UnixNano() / 1000000
}

func waitNextMillis(currentTimestamp int64) int64  {
	for currentTimestamp == lastTimestamp {
		currentTimestamp = getCurrentTimestamp()
	}
	return currentTimestamp
}

