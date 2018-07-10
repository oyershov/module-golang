package cipher

import (
	"strings"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

/* Caesar Cipher */

type shift int

type caesarCipher struct{}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(str int) Cipher {
	c := shift(str)
	if c > 0 && c < 26 {
		return c
	} else if c > -26 && c < 0 {
		c += 26
		return c
	}
	return nil
}

func (c shift) Encode(str string) string {
	return strings.Map(func(r rune) rune {
		return encode_char(r, int(c))
	}, str)
}

func (c shift) Decode(str string) string {
	return strings.Map(func(r rune) rune {
		return decode_char(r, int(c))
	}, str)
}

func encode_char(r rune, c int) rune {
	if r >= 'a' && r <= 'z' {
		return (r-'a'+rune(c))%26 + 'a'
	} else if r >= 'A' && r <= 'Z' {
		return (r-'A'+rune(c))%26 + 'a'
	}
	return -1
}

func decode_char(r rune, c int) rune {
	if r >= 'a' && r <= 'z' {
		return (r-'a'+rune(26-c))%26 + 'a'
	}
	return -1
}

/* Vigenere Cipher */

type vigenere string

func NewVigenere(key string) Cipher {
	flag := false
	for _, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		} else if r > 'a' {
			flag = true
		}
	}
	if !flag {
		return nil
	}

	return vigenere(key)
}

func (v vigenere) Encode(str string) string {
	i := 0
	return strings.Map(func(r rune) rune {
		if r = encode_char(r, int(v[i]-'a')); r >= 0 {
			i = (i + 1) % len(v)
		}
		return r
	}, str)
}

func (v vigenere) Decode(str string) string {
	i := 0
	return strings.Map(func(r rune) rune {
		if r = decode_char(r, int(v[i]-'a')); r >= 0 {
			i = (i + 1) % len(v)
		}
		return r
	}, str)
}
