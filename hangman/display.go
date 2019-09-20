package hangman

/**
 * @type : Display=> Faire l'affichage sur le terminal
 * @author : AB alias Venon2Fork
 * @version : 1.0
 */

import (
	"fmt"
)

// DrawWelcome Affiche le Contenu au demarage du jeu
func DrawWelcome() {
	fmt.Println(`
	----------------------------------------------------------------
	|\     /|(  ___  )( (    /|(  ____ \(       )(  ___  )( (    /|  	
	| )   ( || (   ) ||  \  ( || (    \/| () () || (   ) ||  \  ( |  
	| (___) || (___) ||   \ | || |      | || || || (___) ||   \ | |  
	|  ___  ||  ___  || (\ \) || | ____ | |(_)| ||  ___  || (\ \) |     
	| (   ) || (   ) || | \   || | \_  )| |   | || (   ) || | \   |   
	| )   ( || )   ( || )  \  || (___) || )   ( || )   ( || )  \  |  
	|/     \||/     \||/    )_)(_______)|/     \||/     \||/    )_)  	

	----------------------------------------------------------------

	------------------------------

        / ___|   / \   |  \/  | ____|
       | |  _   / _ \  | |\/| |  _| 
       | |_| | / ___ \ | |  | | |___
        \____/_/    \_\|_|  |_|_____|

       ------------------------------

	----------------------------------------------------------------
	`)
}

// Draw dessine l'etat de la partie
func Draw(g *Game, guess string) {
	// on appelle l'etat courant du pendu
	drawTurns(g.TurnLeft) //TurnLeft afficher le nombre de tours restant
	// on appelle l 'etat de la partie au sens game(lettre trouver , ...etc)
	drawState(g, guess)
	fmt.Println()
}

// drawTurns affiche l'etat du pendu la fonction est privé
func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
    ____
   |    |      
   |    |
   |    o _      
   |   / \      
   |      \__
   |      / 
  _|_
 |   |______
 |GAME OVER | 
 |__________|
		`
	case 1:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |    
  _|_
 |   |______
 |          |
 |__________|
		`
	case 2:
		draw = `
    ____
   |    |      
   |    o      
   |    |
   |    |
   |     
  _|_
 |   |______
 |          |
 |__________|
		`
	case 3:
		draw = `
    ____
   |    |      
   |    o      
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 4:
		draw = `
    ____
   |    |      
   |      
   |      
   |  
   |  
  _|_
 |   |______
 |          |
 |__________|
		`
	case 5:
		draw = `
    ____
   |        
   |        
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 6:
		draw = `
    
   |     
   |     
   |     
   |
   |
  _|_
 |   |______
 |          |
 |__________|
		`
	case 7:
		draw = `
  _ _
 |   |______
 |          |
 |__________|
		`
	case 8:
		draw = `

		`
	}
	fmt.Println(draw)

}

func drawState(g *Game, guess string) {
	// afficher les lettres qui ont été trouver
	fmt.Print("Guessed : ")
	drawLetters(g.FoundLetters)

	//Pour les lettres qui ont deja ete utiliser
	fmt.Print("Used :  ")
	drawLetters(g.UseLetters)

	// En fonction de l'etat du jeu afficher quelque chose de different
	switch g.State {
	//si on trouve une lettre
	case "goodGuess":
		fmt.Println("Good guess!(｡◕‿◕｡)  👍")
		fmt.Printf("||YOU HAVE %v TRY || ", g.TurnLeft)
	// si on propose une lettres deja soumise
	case "alreadyGuessed":
		fmt.Printf("(҂◡_◡)👎  Letter '%s' was already used", guess)
		fmt.Printf("||YOU HAVE %v TRY || ", g.TurnLeft)
	// si on a proposer une mauvaise
	case "badGuess":
		fmt.Printf("Bad guess ¿ⓧ_ⓧﮌ, '%s' is not in the word \n", guess)
		fmt.Printf("||YOU HAVE %v TRY || ", g.TurnLeft)
	// si on a perdu
	case "lost":
		fmt.Print("You lost :(! 💀 The word was: ")
		// drawLetters(g.Letters): les lettres du mot qu'il  fallait trouver
		drawLetters(g.Letters)
	// si on a gagne
	case "won":
		fmt.Println("YOU 👏👏👏  WON! 👍 \nThe word was: ")
		// les lettres du mot qu'il  fallait trouver drawLetters(g.Letters)
		fmt.Print("________________\n")
		fmt.Print("|")
		drawLetters(g.Letters)
		fmt.Print("|")
		fmt.Print("________________|\n")
		fmt.Printf("| YOU HAD %v TRY  | \n", g.TurnLeft)
		fmt.Print("________________\n")
	}
	fmt.Println()

}

// les lettres qui ont deja ete utiliser
func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}
