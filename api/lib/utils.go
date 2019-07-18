package lib

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
)

func GenerateSignatureV1(ts string) (s string) {
	ctx := md5.New()
	bs := generateSalt(ts)
	ctx.Write(bs)
	return hex.EncodeToString(ctx.Sum(nil))
}

func generateSalt(ts string) (bs []byte) {
	// salt 为 ts 的倒序字符串拼接上固定字符串
	HARDCODE1 := "915683c585654984"
	HARDCODE2 := "828e3196301c43a1"
	_ts := reverseString(ts)
	return joinString(HARDCODE1, _ts, HARDCODE2)
}

func joinString(code1, ts, code2 string) (bs []byte) {
	var buffer bytes.Buffer
	buffer.WriteString(code1)
	buffer.WriteString(ts)
	buffer.WriteString(code2)
	return buffer.Bytes()
}

func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}
