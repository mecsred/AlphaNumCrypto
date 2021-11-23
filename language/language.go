package language

import (
)

type Language interface {
	RotateRune(r rune, key int) rune
	KeyModulus() int

	IndexToRune(idx int) rune
	RuneToIndex(r rune) int

	FilterAlphabetic([]rune) []rune
	FormatAlphabetic([]rune,[]rune) []rune

	ScoreChi2([]rune) float64
}
