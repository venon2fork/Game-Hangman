package dictionary

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

// on stock le mot recuperer dans un slice prive
var words = make([]string, 0, 50)

// func qui permet le chargement de notre dictionnaire
func Load(filename string) error {
	//on ouvre le fichier et nous renvois un fichier et une erreur
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	// on ferme le fichier
	defer f.Close()

	// on creer un scanner
	scanner := bufio.NewScanner(f)
	////lire  le fichier ligne par ligne et ajouter au slice words
	for scanner.Scan() {
		// on ajoute a words un mot du fichier
		words = append(words, scanner.Text())
	}
	// s'il y aune lors de lecture du fichier
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// func qui recuperer un mot aleatoire dans liste de words
func PickWord() string {
	//seed une racine qui genere de l'aleatoire
	// time.now pour tirer un nombre different
	// unix pour le nombre de milliseconde ecoule depuis le debut
	rand.Seed(time.Now().Unix())
	//pour tirer un nombre a l'interieur de l intervalle(la longueur de words)
	i := rand.Intn(len(words))
	//et on retourne un mots avec le nombre entier comme index
	return words[i]
}
