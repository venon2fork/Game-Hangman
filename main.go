package main

/**
 * @type : main
 * @dest: main
 * @author : AB alias Venon2Fork
 * @version : 1.0
 */

import (
	//"FAIRE L INPORT DU => dictionary"
	//"FAIRE L IMPORT DU => hangman"
	"fmt"
	"os"
	"time"
)

func main() {

	//TEST L'INITIALISATION DU JEU AVEC DES VALEURS(fichier hangman.go)
	// //IL Y A 8 ETAT DANS LE PENDU
	// g := hangman.New(8, "GOLANG")

	// Chargement avec un dictionnaire
	// on charge le dictionnaire
	err := dictionary.Load("words.txt")
	if err != nil {
		fmt.Printf("Could not load Dictionnary %v\n", err)
		os.Exit(1)
	}
	g, err := hangman.New(8, dictionary.PickWord())
	if err != nil {
		fmt.Printf("Could not create game: %v\n", err)
		os.Exit(1)
	}

	//fmt.Println(c)
	hangman.DrawWelcome()

	// creer une vraible qui contient notre hypothese la lettre qui a été trouver
	guess := ""
	for {
		//a chaque tours on affiche notre etat de jeu avec au debut un guess vide
		hangman.Draw(g, guess)

		//Pour sortir de la partie soit lost ou Won
		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}

		// si on a un probleme sur la saisie utilisateur on utlise le bloc test
		//TEST LA VALEUR EN ENTRER clavier
		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal %v", err)
			//sortie avec code erreur
			os.Exit(1)
		}
		// on assigne a guess la lettre proposer
		guess = l
		//pour afficher l'etat du pendu : bon ou  mauvaise lettre
		g.MakeAGuess(guess)
		time.Sleep(time.Second)
	}

}
