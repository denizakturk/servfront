package config

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
)

func NewSecurityConfig(key, iv string) *Crypt {
	return &Crypt{Key: key, IV: iv}
}

type Crypt struct {
	Key string
	IV  string
}

func (c *Crypt) GetKeyBytes() []byte {
	return c.sha256([]byte(c.Key))
}

func (c *Crypt) GetIVBytes() []byte {
	return c.sha256([]byte(c.IV))[:16]
}

func (c Crypt) sha256(data []byte) []byte {
	hashed := sha256.Sum256(data)
	return []byte(hex.EncodeToString(hashed[:]))
}

func (c Crypt) randString() string {
	rrk := []rune{}
	for i := 0; i < 16; i++ {
		rrk = append(rrk, int32(rand.Intn(127)))
	}

	return string(rrk)
}
