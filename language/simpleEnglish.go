package language

import (
	"unicode/utf8"
)

const(
	cSEAlphabetic = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	cSEAlphabeticLower = "abcdefghijklmnopqrstuvwxyz"
	cSEPunctuation = ",.!?"
	cLowerStart = 0x61
	cLowerEnd = 0x7A
	cUpperStart = 0x41
	cUpperEnd = 0x5A
	cNumCharacters = 26
	cInvalidChi2 = -1000000.0
)

/*-----------------------------------------------------------------------------/
* SimpleEnglish implementation
/-----------------------------------------------------------------------------*/

type SimpleEnglish struct {
	Stats [cNumCharacters]float64
}

func NewSimpleEnglish(sample string) *SimpleEnglish {
	var lang SimpleEnglish
	rs := []rune(sample)
	var total float64 = 0.0
	for _, v := range rs {
		i := lang.RuneToIndex(v)
		if i >= 0 {
			lang.Stats[i] += 1.0
			total += 1.0
		}
	}
	return &lang
}

func (se *SimpleEnglish) FilterAlphabetic(input []rune) []rune {
	var out []rune
	for _, r := range input {
		if (r >= cLowerStart) && (r <= cLowerEnd) ||
			(r >= cUpperStart) && (r <= cUpperEnd) {
			out = append(out, r)
		}
	}
	return out
}

func (se *SimpleEnglish) FormatAlphabetic(alpha, format []rune) []rune {
	out := make([]rune, len(format))
	alphaIdx := 0
	for i,r := range format {
		if ((r >= cLowerStart) && (r <= cLowerEnd)) ||
			((r >= cUpperStart) && (r <= cUpperEnd)) {
			out[i] = alpha[alphaIdx]
			alphaIdx += 1
		} else {
			out[i] = format[i]
		}
	}
	return out
}

func (se *SimpleEnglish) RotateRune(r rune, key int) rune {
	if (r >= cLowerStart) && (r <= cLowerEnd) {
		idx := r - cLowerStart + rune(key)
		if idx < 0 {
			idx = cNumCharacters - idx
		}
		if idx >= cNumCharacters {
			idx = idx % cNumCharacters
		}
		return idx + cLowerStart
	}
	if (r >= cUpperStart) && (r <= cUpperEnd) {
		idx := r - cUpperStart + rune(key)
		if idx < 0 {
			idx = cNumCharacters - idx
		}
		if idx >= cNumCharacters {
			idx = idx % cNumCharacters
		}
		return idx + cUpperStart
	}
	return r
}

func (se *SimpleEnglish) KeyModulus() int {
	return cNumCharacters
}

func (se *SimpleEnglish) IndexToRune(idx int) rune {
	if idx > 0 {
		return rune((idx % cNumCharacters) + cUpperStart)
	}
	return utf8.RuneError
}

func (se *SimpleEnglish) RuneToIndex(r rune) int {
	if (r >= cLowerStart) && (r <= cLowerEnd) {
		return int((r - cLowerStart) % cNumCharacters)
	}
	if (r >= cUpperStart) && (r <= cUpperEnd) {
		return int((r - cUpperStart) % cNumCharacters)
	}
	return -1
}

func (se *SimpleEnglish) ScoreChi2(text []rune) float64 {
	var sampleStats [cNumCharacters]float64
	var total float64 = 0.0
	for _, r := range text {
		if (r >= cLowerStart) && (r <= cLowerEnd) {
			sampleStats[int((r - cLowerStart) % cNumCharacters)] += 1.0
			total += 1.0
		}
		if (r >= cUpperStart) && (r <= cUpperEnd) {
			sampleStats[int((r - cUpperStart) % cNumCharacters)] += 1.0
			total += 1.0
		}
	}
	if total == 0 {
		return cInvalidChi2
	}
	var chi2 float64 = 0.0
	for i := range sampleStats {
		sampleStats[i] /= total
		chi2 += ((sampleStats[i] - se.Stats[i]) * (sampleStats[i] - se.Stats[i])) / se.Stats[i]
	}
	return -1.0*chi2
}

/*-----------------------------------------------------------------------------/
* Default implementations
/-----------------------------------------------------------------------------*/

// Basic SimpleEnglish implementation
var BSE SimpleEnglish

/*-----------------------------------------------------------------------------/
* Init
/-----------------------------------------------------------------------------*/

func init() {
	// Stats from a random webpage on Cornell's site
	BSE.Stats = [cNumCharacters]float64{
		0.0812, //A
		0.0149, //B
		0.0271, //C
		0.0432, //D
		0.1202, //E
		0.0230, //F
		0.0203, //G
		0.0592, //H
		0.0731, //I
		0.0010, //J
		0.0069, //K
		0.0398, //L
		0.0261, //M
		0.0695, //N
		0.0768, //O
		0.0182, //P
		0.0011, //Q
		0.0602,	//R
		0.0628, //S
		0.0910, //T
		0.0288, //U
		0.0111, //V
		0.0209, //W
		0.0017, //X
		0.0211, //Y
		0.0007, //Z
	}
}
