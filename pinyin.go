package pinyin4go

import (
	"strconv"
	"strings"
)

type RomanType int

const (
	Hanyu    RomanType = iota //Hanyu Pinyin system
	Wade                      //Wade-Giles Pinyin system
	MPSII                     //Mandarin Phonetic Symbols 2 (MPS2) Pinyin system
	Yale                      //Yale Pinyin system
	Tongyong                  //Tongyong Pinyin system
	Gwoyeu                    //Gwoyeu Romatzyh system
)

var none = "(none0)"

func Pinyin(ch rune) []string {
	return getPinyinRecord(ch)
}

func GwoyeuRomatzyh(ch rune) []string {
	hypy := Pinyin(ch)
	if hypy == nil {
		return nil
	}
	r := make([]string, len(hypy))

	for i := 0; i < len(hypy); i++ {
		r[i] = convertToGwoyeuRomatzyh(hypy[i])
	}
	return r
}
func WadeGilesPinyin(ch rune) []string {
	return romanPinyin(ch, Wade)
}

func MPS2Pinyin(ch rune) []string {
	return romanPinyin(ch, MPSII)
}

func YalePinyin(ch rune) []string {
	return romanPinyin(ch, Yale)
}

func TongyongPinyin(ch rune) []string {
	return romanPinyin(ch, Tongyong)
}

func romanPinyin(ch rune, t RomanType) []string {
	hypy := Pinyin(ch)
	if hypy == nil {
		return nil
	}
	r := make([]string, len(hypy))

	for i := 0; i < len(hypy); i++ {
		r[i] = convertRomanizationSystem(hypy[i], t)
	}
	return r
}

func getPinyinRecord(ch rune) []string {
	record := UnicodeToPinyin[ch]
	if record == "" || record == none {
		return nil
	}
	if strings.HasPrefix(record, "(") && strings.HasSuffix(record, ")") {
		return strings.Split(record[1:len(record)-1], ",")
	}
	return nil
}

func convertToGwoyeuRomatzyh(wholePy string) string {
	p := len(wholePy) - 1
	pinyin, tone := wholePy[:p], wholePy[p:]

	gwoteus, ok := GwoyeuMapping[pinyin]
	if !ok {
		return ""
	}

	i, _ := strconv.Atoi(tone)
	i = i - 1
	if i > len(gwoteus) {
		return ""
	}
	return gwoteus[i]
}

func convertRomanizationSystem(input string, t RomanType) string {
	if t == Hanyu {
		return input
	}
	p := len(input) - 1
	pinyin, tone := input[:p], input[p:]

	romans, ok := RomanMapping[pinyin]
	if !ok {
		return ""
	}

	var i int = int(t) - 1
	if i > len(romans) {
		return ""
	}
	return romans[i] + tone
}
