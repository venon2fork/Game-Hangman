package hangman

/**
 * @type : Input => Fichier permettant de recuperer la saisie utilisateur
 * @author : AB alias Venon2Fork
 * @version : 1.0
 */
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//1. Declaration du reader qui permet de recuperer la saisie utilsateur clavier
var reader = bufio.NewReader(os.Stdin)

//Fonction ReadGuess Lecture d'une proposition faite au clavier..
//Renvois la proposition et une erreur a la saisie du clavier
func ReadGuess() (guess string, err error) {
	// Faire une proposition tant que la proposition n'est pas valide ou qu'on a eu une erreur
	// definir une variable valid
	valid := false
	// On creer une boucle tant que l 'on est pas valid
	for !valid {
		// afficher la question qu'lle est votre lettre
		fmt.Print("What is your letter ? ")
		// 1. pour recuperer la saisie clavier on utilise un reader qu on declare
		// a l'exterieur de la fonction.
		// Pour recuperer l'enterer standart
		guess, err = reader.ReadString('\n') //jusqu a ou il fait sa lecture (\n == retour a la ligne "Enter")
		//renvois deux guess = la saisie clavier ou une erreur
		if err != nil {
			return guess, err
		}
		// s'il y a pas d'erreur dans ce cas la on renvois guess sans les espaces
		guess = strings.TrimSpace(guess)
		// Si jamais la longueur depasse  un,  continue.
		if len(guess) != 1 {
			//renvois  la saisie clavier et ca longueur
			fmt.Printf("Invalid Letter .... Letter =>  %v , len =>  %v  \n", guess, len(guess))
			continue
		}

		//si la valeur est valid on arrete la boucle
		valid = true
	}
	return guess, nil
}
