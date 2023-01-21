package main

import "fmt"

func main() {
	jamesBond := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jamesBond)
	missPenny := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(missPenny)

	multidimansionalSlice := [][]string{jamesBond, missPenny}
	fmt.Println(multidimansionalSlice)
}
