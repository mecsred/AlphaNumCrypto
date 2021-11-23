package ciphers

import (
	lang "github.com/mecsred/AlphaNumCrypto/language"
	"github.com/mecsred/AlphaNumCrypto/scoring"
)

const (
	DEFAULT_MAX_KEYSIZE = 30
)

func RunesToKey(l lang.Language, key []rune) []int {
	intKey := make([]int, len(key))
	for i := range(key) {
		intKey[i] = l.RuneToIndex(key[i])
	}
	return intKey
}

func KeyToRunes(l lang.Language, key []int) []rune {
	runeKey := make([]rune, len(key))
	for i := range key {
		runeKey[i] = l.IndexToRune(key[i])
	}
	return runeKey
}

func splitInterval(l lang.Language, text []rune, ival int) [][]rune {
	splits := make([][]rune, ival)
	for i,r := range text {
		splits[i%ival] = append(splits[i%ival], r)
	}
	return splits
}

func GetIndexOfCoincidence(l lang.Language, text []rune) float64 {
	stats := make([]int, l.KeyModulus())
	for _,r := range text {
		stats[l.RuneToIndex(r)] += 1
	}
	var IC float64
	for i := range stats {
		IC += float64(stats[i]*(stats[i]-1))
	}
	IC /= float64(len(text)*(len(text) - 1))
	return IC
}

func EstimateKeysize(l lang.Language, pText []rune, maxKeysize, results int) []int {
	body := l.FilterAlphabetic(pText)
	var sl scoring.ScoredList
	for i := 2; i < maxKeysize; i += 1 {
		splits := splitInterval(l, body, i)
		var avgIC float64
		for j := 0; j < i; j += 1 {
			avgIC += GetIndexOfCoincidence(l, splits[j])
		}
		avgIC /= float64(i)
		sl.AddItem(i, avgIC)
	}
	items := sl.TopN(results)
	out := make([]int, len(items))
	for i := range items {
		out[i] = items[i].Item.(int)
	}
	sl.Print()
	return out
}

func VigEncrypt(l lang.Language, pText []rune, key []int) []rune {
	keyLen := len(key)
	input := l.FilterAlphabetic(pText)
	cText := make([]rune, len(input))
	for i := range input {
		cText[i] = l.RotateRune(input[i], key[i % keyLen])
	}
	return l.FormatAlphabetic(cText, pText)
}

func VigDecrypt(l lang.Language, cText []rune, key []int) []rune {
	keyLen := len(key)
	input := l.FilterAlphabetic(cText)
	pText := make([]rune, len(input))
	for i := range input {
		pText[i] = l.RotateRune(input[i], l.KeyModulus() - key[i % keyLen])
	}
	return l.FormatAlphabetic(pText, cText)
}

func VigKeyfindSearch(l lang.Language, cText []rune, depth int) []int {
	body := l.FilterAlphabetic(cText)
	keysize := EstimateKeysize(l, body, DEFAULT_MAX_KEYSIZE, depth)
	var sl scoring.ScoredList
	for _,size := range keysize {
		splits := splitInterval(l, body, size)
		key := make([]int, size)
		for j := range splits {
			key[j] = RotKeyfindSearch(l, splits[j], 1)[0]
		}
		score := l.ScoreChi2(VigDecrypt(l, body, key))
		sl.AddItem(key, score)
	}
	keys := sl.TopN(depth)
	out := make([][]int, depth)
	sl.Print()
	for i := range keys {
		out[i] = keys[i].Item.([]int)
	}
	return out[0]
}
