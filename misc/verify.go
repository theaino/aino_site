package misc

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"time"
)

func GenerateVerificationKey(email string, salt string) string {
	year, week := time.Now().UTC().ISOWeek()
	week = (year * 53) + week
	message := fmt.Sprintf("%s%s%d", email, salt, week)
	hasher := sha256.New()
	hasher.Write([]byte(message))
	value := binary.BigEndian.Uint64(hasher.Sum(nil))
	return fmt.Sprintf("%d", value)
}
