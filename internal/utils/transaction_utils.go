package utils

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/google/uuid"
)

func GenerateTransactionIdString() string {

	id := uuid.New().String()
	id = fmt.Sprintf("TXN%s", id)

	return id
}

func ProcessTransaction(txnID string) string {

	// stimulate delay
	delay := rand.IntN(15) + 2
	// delay := 15
	time.Sleep(time.Duration(delay) * time.Second)

	randInt := rand.IntN(100)

	if randInt < 80 {
		return "success"
	}

	return "failed"

}
