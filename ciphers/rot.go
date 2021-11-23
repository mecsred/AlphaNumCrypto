package ciphers

import (
	lang "github.com/mecsred/AlphaNumCrypto/language"
	"github.com/mecsred/AlphaNumCrypto/scoring"
)

const (

)

func RotEncrypt(l lang.Language, pText []rune, key int) []rune {
	cText := make([]rune, len(pText))
	for i, v := range pText {
		cText[i] = l.RotateRune(v, key)
	}
	return cText
}

func RotDecrypt(l lang.Language, cText []rune, key int) []rune {
	pText := make([]rune, len(cText))
	for i, v := range cText {
		pText[i] = l.RotateRune(v, l.KeyModulus() - key)
	}
	return pText
}

func RotKeyfindSearch(l lang.Language, cText []rune, numResults int) []int {
	var sl scoring.ScoredList
	for i := 0; i < l.KeyModulus(); i += 1 {
		newScore := l.ScoreChi2(RotDecrypt(l, cText, i))
		sl.AddItem(i, newScore)
	}
	out := make([]int, numResults)
	results := sl.TopN(numResults)
	for i := 0; i < numResults; i += 1 {
		out[i] = results[i].Item.(int)
	}
	return out
}
