package baseint

import (
	"bytes"
	"errors"
	"math"
)

func init() {
	charMapInt, _ = mapToInt(baseStr)
}

var (
	charMapInt map[byte]uint64
	baseStr    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	ErrRepeatChar = errors.New("baseint: error set repeat char for base")
	ErrChar       = errors.New("baseint: error char for convert")
)

func mapToInt(base64Str string) (m map[byte]uint64, err error) {
	m = make(map[byte]uint64)
	for i, b := range []byte(base64Str) {
		if _, ok := m[b]; ok {
			err = ErrRepeatChar
			return
		}
		m[b] = uint64(i)
	}
	return
}

//SetBaseString set base str for baseint, set a random str
//so that can make different convert
func SetBaseString(base string) (err error) {
	m, err := mapToInt(base)
	if err != nil {
		return
	}
	charMapInt = m
	baseStr = base
	return
}

//Atoi convert string to int
func Atoi(b string) (i uint64, err error) {
	var length = float64(len(baseStr))
	for index, c := range []byte(b) {
		if mc, ok := charMapInt[c]; ok {
			i += uint64(math.Pow(length, float64(index))) * mc
		} else {
			err = ErrChar
			return
		}
	}
	return
}

//Itoa convert int to string
func Itoa(i uint64) (b string) {
	var length = uint64(len(baseStr))
	buf := &bytes.Buffer{}
	for i != 0 {
		buf.WriteByte(baseStr[i%length])
		i = i / length
	}
	b = buf.String()
	return
}
