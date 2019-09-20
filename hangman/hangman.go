package hangman

/**
 * @type : hangman => Gerer le déroulement dujeux
 * @author : AB alias Venon2Fork
 * @version : 1.0
 */

import (
	"fmt"
	"regexp"
	"strings"
)

//Creer le point d'entrer du jeux
type Game struct {
	// definition jeux , perdu ou non, on definit l 'etat courant
	State string
	// definit la liste des lettres que l'on doit trouver
	Letters []string
	// les lettres que l'on a trouver en tant que joueur
	FoundLetters []string
	// les lettres que l'on a deja utliser
	UseLetters []string
	// afficher le nombre de tours restant
	TurnLeft int
}

//fonction qui va initialiser la partie pour commencer le jeu
func New(turns int, words string) (*Game, error) {
	// test si le mot est inferieur a 2 caractere
	if len(words) < 2 {
		return nil, fmt.Errorf("Words '%s' must be at least 2 characters. got=%v", words, len(words))
	}
	// Ajouter pour eviter les chiffres
	matched, _ := regexp.MatchString(`[0-9]*`, words)
	if matched != true {
		return nil, fmt.Errorf("Words '%s' contains number", words)
	}
	// Creer un tableau des lettres a trouver
	letter := strings.Split(strings.ToUpper(words), "")
	// lettre trouver dans un tableau
	found := make([]string, len(letter))
	//initialiser le tableau avec "__"
	for i := 0; i < len(letter); i++ {
		found[i] = "__"
	}
	// initaliser le game pour commencer la partie
	g := &Game{
		State:        "",
		Letters:      letter,
		FoundLetters: found,
		UseLetters:   []string{},
		TurnLeft:     turns,
	}
	// renvois  un nouveau Jeux
	return g, nil
}

//Partie logique du jeu

// function MakeAGuess proposition une hypothese en params la lettre proposer par le joueur
func (g *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)
	// on utliser la fonction letterInword pour verifier UseLetter
	if letterInWord(guess, g.UseLetters) {
		//on change l'etat de la state
		g.State = "alreadyGuessed"
	} else if letterInWord(guess, g.Letters) {
		//on change l'etat de la state
		g.State = "goodGuess"
		//on revele la lettre
		g.RevealLetter(guess)

		//puis on verifie si on a trouvées toutes les lettres
		if hasWon(g.Letters, g.FoundLetters) {
			// on change l'etat si on a trouvées toutes les letttres
			g.State = "won"
		}
	} else {
		// si la lettre proposer n'est pas bonne
		g.State = "badGuess"
		g.LoseTurn(guess)

		// on verifie le nombre de tour restant
		if g.TurnLeft <= 0 {
			g.State = "lost"
		}
	}

}

//creer une fonction qui va chercher la lettre dans les differents tableaux (proposer ou trouver)
func letterInWord(guess string, letters []string) bool {
	matched, _ := regexp.MatchString(`[0-9]*`, guess)
	if matched != true {
		return false
	}
	for _, l := range letters {
		if guess == l {
			return true
		}
	}
	return false
}

// fonction qui permet de reveler les lettres deja trouver
func (g *Game) RevealLetter(guess string) {
	// on ajoute guess dans les lettres deja utiliser
	g.UseLetters = append(g.UseLetters, guess)
	//Puis on revele dans le tableau des lettres a trouvées
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}

}

// fonction qui permet de verifier si on a trouver toutes les lettres
// on compare les lettres qu'on a trouvée avec les lettres deja trouver
func hasWon(letter []string, FoundLetters []string) bool {
	for i := range letter {
		if letter[i] != FoundLetters[i] {
			return false
		}
	}
	return true
}

// LoseTurn Permet de decrementer le nombre de tour restant jusqu a zero
func (g *Game) LoseTurn(guess string) {
	g.TurnLeft--
	// on ajoute la lettre au lettre deja utiliser
	g.UseLetters = append(g.UseLetters, guess)
}
