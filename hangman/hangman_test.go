package hangman

/**
 * @type : test
 * @dest: test file Hangman.go
 * @author : AB alias Venon2Fork
 * @version : 1.0
 */

import (
	"regexp"
	"testing"
)

// test si la valeur est contenu dans le mot avec la fonction letterInWord
func TestLetterInWord(t *testing.T) {
	// on genere un mot
	word := []string{"b", "o", "b"}
	// on genere une entrer par le user
	guess := "b"
	// on appelle la fonction pour verifier ca presence dans le mot
	hasLetter := letterInWord(guess, word)
	// on gere si il y a une erreur avec un code erreur
	if !hasLetter {
		t.Errorf("word %s contains letter %s, got=%v ", word, guess, hasLetter)
	}
}

// test que la lettre ne soit pas dans le mot
func TestLetterNotWord(t *testing.T) {
	// on genere un mot
	word := []string{"b", "o", "b"}
	// on genere une entrer par le user
	guess := "c"
	// on appelle la fonction pour verifier ca presence dans le mot
	nothasLetter := letterInWord(guess, word)
	// on gere s'il y a une erreur avec un code erreur
	if nothasLetter {
		t.Errorf("word %s contains letter %s, got=%v ", word, guess, nothasLetter)
	}
}

// TEST QUE LA VALEUR ENTRER NE SOIT PAS UN CHIFFRE A CORRIGER ENCORE
func TestNotNumberInWord(t *testing.T) {
	word := []string{"b", "o", "b"}
	guess := "4"
	matched, _ := regexp.MatchString(`[0-9]*`, guess)
	invalidletter := letterInWord(guess, word)
	if matched != true {
		t.Errorf("returned invalid word got=%v", invalidletter)
	}

}

//TEST SI L ENTRER N EST PAS PLUS PETIT QUE 2
func TestInvalidWord(t *testing.T) {
	invalidWords, err := New(8, " ")
	if err == nil {
		t.Errorf("returned invalid word got=%v", invalidWords)
	}

}

// TEST QUE LE MOT NE CONTIENT PAS DES CHIFFRES
func TestInvalidWordNumeric(t *testing.T) {
	invalidWords, err := New(8, "Aba3456")
	if err != nil {
		t.Errorf("returned invalid word got=%v", invalidWords)
	}
}

// TEST LA FONCTION MakeAGuess AVEC COMME STATE: goodGuess
func TestGameGoodGuess(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	ValidState(t, "goodGuess", g.State)
}

// TEST LA FONCTION MakeAGuess AVEC COMME STATE: alreadyGuessed
func TestGamealreadyGuessed(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	g.MakeAGuess("o")
	g.MakeAGuess("b")
	ValidState(t, "alreadyGuessed", g.State)
}

// TEST LA FONCTION MakeAGuess AVEC COMME STATE: WON
func TestGameWon(t *testing.T) {
	g, _ := New(2, "bob")
	g.MakeAGuess("b")
	g.MakeAGuess("o")
	g.MakeAGuess("b")
	ValidState(t, "Won", g.State)
}

// TEST LA FONCTION MakeAGuess AVEC COMME STATE: badGuess
func TestGameBadGuess(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("c")
	g.MakeAGuess("o")
	g.MakeAGuess("f")
	ValidState(t, "badGuess", g.State)
}

// TEST LA FONCTION MakeAGuess AVEC COMME STATE: lost
func TestGameLost(t *testing.T) {
	g, _ := New(2, "bob")
	g.MakeAGuess("c")
	g.MakeAGuess("o")
	g.MakeAGuess("f")
	ValidState(t, "lost", g.State)
}

// func qui test l'etat actuel et l'etat  attendu
func ValidState(t *testing.T, exceptedState, actualState string) bool {
	if exceptedState != actualState {
		t.Errorf("State should be '%v' , got=%v", exceptedState, actualState)
		return false
	}
	return true
}
