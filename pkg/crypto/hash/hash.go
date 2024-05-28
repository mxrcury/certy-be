package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

type Hasher struct {
	salt string
}

func NewHasher(salt string) *Hasher {
	return &Hasher{salt}
}

func (h *Hasher) Hash(input []byte) string {
	hashReader := sha256.New()

	hashReader.Write(input)

	return hex.EncodeToString(hashReader.Sum([]byte(h.salt)))
}

func (h *Hasher) Verify(input []byte, hashedValue string) bool {
	incomingHashedVal := h.Hash(input)

	return incomingHashedVal == hashedValue
}
