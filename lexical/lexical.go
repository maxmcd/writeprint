package lexical

// http://www.ncfta.ca/papers/emailforensics.pdf
import (
	"math"
	"strings"
)

type Analysis struct {
	corpus                          string
	tokenArray                      []string
	CharacterCount                  int
	RatioOfDigitsToN                float64
	RatioOfLettersToN               float64
	RatioOfUppercaseLettersToN      float64
	RatioOfSpacesToN                float64
	RatioOfTabsToN                  float64
	OccurrencesOfAlphabets          map[rune]float64
	OccurrencesOfSpecialCharachters map[rune]float64
	TokenCount                      int
	AverageSentenceLength           int
	AverageTokenLength              float64
	RationOfCharachtersInWordsToN   float64
	RatioOfShortWordsToT            float64
}

var RatioOfDigitsToNAvg float64
var RatioOfLettersToNAvg float64
var RatioOfUppercaseLettersToNAvg float64
var RatioOfSpacesToNAvg float64
var RatioOfTabsToNAvg float64
var AverageTokenLengthAvg float64

func square(difference float64) float64 {
	return difference * difference
}
func squareInt(difference int) float64 {
	return math.Pow(float64(difference), float64(difference))
}

func Similarities(first, second Analysis) float64 {
	bignumber := square((first.RatioOfDigitsToN-second.RatioOfDigitsToN)/RatioOfDigitsToNAvg) +
		square((first.RatioOfLettersToN-second.RatioOfLettersToN)/RatioOfLettersToNAvg) +
		square((first.RatioOfUppercaseLettersToN-second.RatioOfUppercaseLettersToN)/RatioOfUppercaseLettersToNAvg) +
		square((first.RatioOfSpacesToN-second.RatioOfSpacesToN)/RatioOfSpacesToNAvg) +
		square((first.AverageTokenLength-second.AverageTokenLength)/AverageTokenLengthAvg)
	return bignumber
}

func (a *Analysis) New(corpus string) {
	a.corpus = corpus

	a.characterCount()
	a.ratioOfDigitsToN()
	a.ratioOfLettersToN()
	a.ratioOfUppercaseLettersToN()
	a.ratioOfSpacesToN()
	a.ratioOfTabsToN()
	a.occurrencesOfAlphabets()
	a.occurrencesOfSpecialCharachters()
	a.tokenCount()
	a.averageSentenceLength()
	a.averageTokenLength()
	a.rationOfCharachtersInWordsToN()
	a.ratioOfShortWordsToT()

}

func (a *Analysis) characterCount() {
	a.CharacterCount = len(a.corpus)
}

func (a *Analysis) ratioOfDigitsToN() {
	a.RatioOfDigitsToN = ratioOfThingToN(
		a.corpus,
		"0123456789",
		a.CharacterCount,
	)
}
func (a *Analysis) ratioOfLettersToN() {
	a.RatioOfLettersToN = ratioOfThingToN(
		a.corpus,
		"qwertyuiopasdfghjklzxcvbnm",
		a.CharacterCount,
	)
}
func (a *Analysis) ratioOfUppercaseLettersToN() {
	a.RatioOfUppercaseLettersToN = ratioOfThingToN(
		a.corpus,
		"QWERTYUIOPASDFGHJKLZXCVBNM",
		a.CharacterCount,
	)
}
func (a *Analysis) ratioOfSpacesToN() {
	a.RatioOfSpacesToN = ratioOfThingToN(
		a.corpus,
		" ",
		a.CharacterCount,
	)
}
func (a *Analysis) ratioOfTabsToN() {
	a.RatioOfTabsToN = ratioOfThingToN(
		a.corpus,
		"\t",
		a.CharacterCount,
	)
}
func (a *Analysis) occurrencesOfAlphabets() {

}

func (a *Analysis) occurrencesOfSpecialCharachters() {

}

func (a *Analysis) tokenCount() {
	a.tokenArray = strings.Split(a.corpus, " ")
	a.TokenCount = len(a.tokenArray)
}

func (a *Analysis) averageSentenceLength() {

}

func (a *Analysis) averageTokenLength() {
	var total float64 = 0
	for _, value := range a.tokenArray {
		total += float64(len(value))
	}
	a.AverageTokenLength = total / float64(len(a.tokenArray))
}

func (a *Analysis) rationOfCharachtersInWordsToN() {

}

func (a *Analysis) ratioOfShortWordsToT() {

}

func ratioOfThingToN(corpus, thing string, nLength int) (ratio float64) {
	var count float64
	digitMap := make(map[rune]bool)
	for _, value := range thing {
		digitMap[value] = true
	}
	for _, value := range corpus {
		if digitMap[value] == true {
			count = count + 1
		}
	}
	f := float64(nLength)
	ratio = count / f
	return
}
