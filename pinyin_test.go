package pinyin4go

import (
	"testing"
)

func TestPinyin(t *testing.T) {
	exp := []string{"li3"}
	if !eq(exp, Pinyin('李')) {
		t.Error("func Pinyin() got error")
	}
}

func TestGwoyeuRomatzyh(t *testing.T) {
	if !eq([]string{"lii"}, GwoyeuRomatzyh('李')) {
		t.Error("func GwoyeuRomatzyh() got error")
	}
}

func TestWadeGilesPinyin(t *testing.T) {
	exp := []string{"li3"}
	if !eq(exp, WadeGilesPinyin('李')) {
		t.Error("func WadeGilesPinyin() got error")
	}
}

func TestMPS2Pinyin(t *testing.T) {
	exp := []string{"li3"}
	if !eq(exp, MPS2Pinyin('李')) {
		t.Error("func MPS2Pinyin() got error")
	}
}

func TestYalePinyin(t *testing.T) {
	exp := []string{"li3"}
	if !eq(exp, YalePinyin('李')) {
		t.Error("func YalePinyin() got error")
	}
}

func TestTongyongPinyin(t *testing.T) {
	exp := []string{"li3"}
	if !eq(exp, TongyongPinyin('李')) {
		t.Error("func TongyongPinyin() got error")
	}
}

func eq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
