package main

import ("fmt"
	"math/rand"
	"strconv"
	"time"
	)

func generate(len int) string {
	result := ""
	for i:=1;i<= len;i++ {
		result += strconv.Itoa(rand.Intn(10))
	}
	return result
}

func evaluate(guess string, answer string) (int, int, int) {
	exact, less, greater := 0, 0, 0
	lofa := len(answer)
	for i:=1;i<=lofa;i++ {
		fmt.Println(i)
	}
	return exact, less, greater
}

func bruteForce(answer string) int {
	guesses, lofa := 0, len(answer)
	
}

func main() {
	rand.Seed(time.Now().Unix())
	answer, guess := generate(20), generate(20)
	fmt.Println(answer)
	fmt.Println(guess)
	evaluate(guess, answer)
    
    
	fmt.Println("Brute force")

}

