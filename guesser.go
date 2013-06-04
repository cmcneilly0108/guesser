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
	g, a := "", ""
	for i:=0;i<lofa;i++ {
		g = guess[i:i+1]
		a = answer[i:i+1]
		if (g < a) {
			less++
		} else {
			if (g == a) {
				exact++
			} else {
 				greater++
 			}
		}
	}
	return exact, less, greater
}

func bruteForce(answer string) int {
	guesses, lofa := 0, len(answer)

	guess := generate(lofa)
	evaluate(guess,answer)
	fmt.Println(guess)
	return guesses
}

func main() {
	rand.Seed(time.Now().Unix())
	answer := generate(20)
	fmt.Println("The answer is ="+answer)

	fmt.Println("Brute force")
	guesses := bruteForce(answer)
	fmt.Println("It took "+strconv.Itoa(guesses) + " guesses")
    
    

}

