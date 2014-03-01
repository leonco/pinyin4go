Overview
=========

pinyin4go is a golang library supporting convertion between Chinese characters and most popular Pinyin systems. The output format of pinyin could be customized. It's a clone of [pinyin4j](http://pinyin4j.sourceforge.net/).


Features
==================

1. Convert [Chinese](http://en.wikipedia.org/wiki/Chinese_language) (both Simplified and Tranditional) to most popular [pinyin](http://en.wikipedia.org/wiki/Pinyin) systems. Supporting pinyin system are listed below. 
   * [Hanyu Pinyin](http://en.wikipedia.org/wiki/Hanyu_Pinyin)
   * [Tongyong Pinyin](http://en.wikipedia.org/wiki/Tongyong_Pinyin)
   * [Wade-Giles](http://en.wikipedia.org/wiki/Wade-Giles)
   * [MPS2 (Mandarin Phonetic Symbols II)](http://en.wikipedia.org/wiki/MPS2)
   * [Yale Romanization](http://en.wikipedia.org/wiki/Yale_Romanization)
   * [Gwoyeu Romatzyh](http://en.wikipedia.org/wiki/Gwoyeu_Romatzyh)

2. Support multiple pronounciations of a single Chinese character. For example, taking the Chinese character `'和'` as input, you will get eight Hanyu Pinyin results `(hé hè huó huò huo hāi he hú)`, depeding on different contexts.

3. Multiple options for output format
   * All uppercase or lowercase
   * Can output Unicode `ü` or `v` or `u`:
   * With tone numbers (lü3) or tone marks `(lǚ)` or without tone `(lü)`


Combination of output format options
=======================================
Three types of output format options can be customized.
* `VCharType`: output format of character 'ü', which has three options.
   - `U_COLON` (default)
   - `V`
   - `U_UNICODE`
*  `ToneType`: output format of Chinese tones, which has three options.
   - `TONE_NUMBER` (default)
   - `NO_TONE`
   - `TONE_MARK`
*  `CaseType`: cases of letters in outputted string, which has two options.
   - `LOWER_CASE` (default)
   - `UPPER_CASE`
Some combinations of these three output formats are forbidden. 


Example
====================
```go
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
```
Output: `[hé hè huó huò huo hāi he hú]`


