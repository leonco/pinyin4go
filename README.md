Work in progress.

Overview
=========

pinyin4go is a golang library supporting convertion between Chinese characters and most popular Pinyin systems. The output format of pinyin could be customized.


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



