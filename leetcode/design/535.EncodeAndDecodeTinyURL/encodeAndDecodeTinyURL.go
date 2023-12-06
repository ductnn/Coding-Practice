package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

const (
	TinyURLLength = 6
)

type Codec struct {
	m     map[string]string
	randz string
}

func Constructor() Codec {
	m := map[string]string{}
	return Codec{m, generateRandomString(TinyURLLength)}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	prefix := "http://tinyurl.com/"
	this.m[this.randz] = longUrl
	return prefix + this.randz
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	i := strings.Split(shortUrl, "/")
	encodedPart := i[len(i)-1]
	longUrl, _ := this.m[encodedPart]

	return longUrl
}

func generateRandomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */

func main() {
	longUrl := "https://www.example.com"
	obj := Constructor()
	url := obj.encode(longUrl)
	ans := obj.decode(url)

	fmt.Println(url, ans)
}
