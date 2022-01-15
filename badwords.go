
package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	var file_name string
	var word string
	var words = []string{}

	

	fmt.Scan(&file_name)
	for {
		_, userWordErr := fmt.Scanln(&word)
		if userWordErr != nil {
			
		} 
		if (word == "exit"){
			fmt.Println("Bye!")
			return
		}

		file, err := os.Open(file_name)
	
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	
		scanner := bufio.NewScanner(file)
		for scanner.Scan(){
			words = append(words, scanner.Text())
		}

		//word = strings.ToLower(word)
		if (Contains(words, word) == true){
			createStar(word)
		} else {
			fmt.Println(word)
		}
		
	
		if err := scanner.Err(); err != nil{
			log.Fatal(err)
		}
	}

}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			  return true
		}
	}
	return false
}

func createStar(word string) {
	star := ""
	for  i := 0; i < len(word); i++ {
		star += "*"
	}
	fmt.Println(star)
}
