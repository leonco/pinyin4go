package main

import (
	"fmt"
	pinyin "github.com/leonco/pinyin4go"
)

func main() {
	f := pinyin.Format{pinyin.U_UNICODE, pinyin.LOWER_CASE, pinyin.TONE_MARK}
	r := '和'
	ss, err := pinyin.PinyinF(r, f)
	if err != nil {
		panic(err)
	}
	fmt.Println(ss)
}
