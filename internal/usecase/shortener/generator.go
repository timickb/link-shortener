package shortener

import (
	"crypto/sha256"
	"github.com/speps/go-hashids"
	"strconv"
	"strings"
)

const (
	lower    = "abcdefghijklmnopqrstuvwxyz"
	upper    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers  = "0123456789"
	special  = "_"
	alphabet = lower + upper + numbers + special
)

func generateShortening(url string) (string, error) {
	url = strings.ToLower(url)
	var salt int
	for i := 0; i < len(url); i++ {
		salt += int(url[i])
	}

	sum := sha256.Sum256([]byte(url))

	intSlice := make([]int, len(sum))
	for i := 0; i < len(sum); i++ {
		intSlice[i] = int(sum[i])
	}

	hd := hashids.NewData()
	hd.MinLength = 10
	hd.Alphabet = alphabet
	hd.Salt = strconv.Itoa(salt)

	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}

	encoded, err := h.Encode(intSlice)
	if err != nil {
		return "", err
	}

	return encoded[:hd.MinLength], nil
}
