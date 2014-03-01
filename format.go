package pinyin4go

import (
	"errors"
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	a        = 'a'
	e        = 'e'
	ou       = "ou"
	unmarked = "aeiouv"
	marked   = "āáăàaēéĕèeīíĭìiōóŏòoūúŭùuǖǘǚǜü"
)

type (
	CaseType  int
	ToneType  int
	VCharType int
)

const (
	TONE_NUMBER ToneType = iota
	TONE_MARK
	NO_TONE
)

const (
	UPPER_CASE CaseType = iota
	LOWER_CASE
)

const (
	U_COLON VCharType = iota
	V
	U_UNICODE
)

type Format struct {
	VCharType
	CaseType
	ToneType
}

var defaultFormat = Format{U_COLON, LOWER_CASE, TONE_NUMBER}

func formatPinyin(s string, f Format) (string, error) {
	if f.ToneType == TONE_MARK && (f.VCharType == U_COLON || f.VCharType == V) {
		return "", errors.New("tone marks cannot be added to v or u:")
	}

	if f.ToneType == NO_TONE {
		s = regexp.MustCompile("[1-5]").ReplaceAllString(s, "")
	} else if f.ToneType == TONE_MARK {
		s = strings.Replace(s, "u:", "v", -1)
		s = convertToneNumber2ToneMark(s)
	}

	if f.VCharType == V {
		s = strings.Replace(s, "u:", "v", -1)
	} else if f.VCharType == U_UNICODE {
		s = strings.Replace(s, "u:", "ü", -1)
	}

	if f.CaseType == UPPER_CASE {
		s = strings.ToUpper(s)
	}
	return s, nil
}

func convertToneNumber2ToneMark(pinyin string) string {
	lowerpy := strings.ToLower(pinyin)

	// bad format
	if ok, _ := regexp.MatchString(`[a-z]+[1-5]?`, lowerpy); !ok {
		return lowerpy
	}

	// input string has no any tune number
	// only replace v with ü (umlat) character
	if ok, _ := regexp.MatchString(`[a-z]+[1-5]`, lowerpy); !ok {
		return strings.Replace(lowerpy, "v", "ü", -1)
	}

	char := '$'
	index := -1

	unmarkedVowel := char
	indexOfUnmarkedVowel := index

	l := len(lowerpy)
	tuneN := lowerpy[l-1] - '0'

	indexOfA := strings.IndexByte(lowerpy, a)
	indexOfE := strings.IndexByte(lowerpy, e)
	indexOfOu := strings.Index(lowerpy, ou)

	if indexOfA >= 0 {
		unmarkedVowel = a
		indexOfUnmarkedVowel = indexOfA
	} else if indexOfE >= 0 {
		unmarkedVowel = e
		indexOfUnmarkedVowel = indexOfE
	} else if indexOfOu >= 0 {
		unmarkedVowel = rune(ou[0])
		indexOfUnmarkedVowel = indexOfOu
	} else {
		for i := l - 1; i >= 0; i-- {
			if strings.IndexByte(unmarked, lowerpy[i]) >= 0 {
				indexOfUnmarkedVowel = i
				unmarkedVowel = rune(lowerpy[i])
				break
			}
		}
	}
	if char != unmarkedVowel && index != indexOfUnmarkedVowel {
		row := strings.IndexRune(unmarked, unmarkedVowel)
		col := tuneN - 1
		c := getMarked(row*5 + int(col))
		var buf []byte
		b1 := []byte(strings.Replace(lowerpy[:indexOfUnmarkedVowel], "v", "ü", -1))
		b2 := []byte(lowerpy[indexOfUnmarkedVowel+1 : l-1])
		buf = append(buf, b1...)
		buf = append(buf, c...)
		buf = append(buf, b2...)
		return string(buf)
	} else {
		return lowerpy
	}
}

func getMarked(i int) []byte {
	b := []byte(marked)
	j := 0
	for len(b) > 0 {
		_, n := utf8.DecodeRune(b)
		if j == i {
			return b[:n]
		}
		j++
		b = b[n:]
	}
	return nil
}
