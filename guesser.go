package main

import ("fmt"
	"math/rand"
	"strconv"
	"time"
	)

type guess struct {
	gstring string
	exact int
	less int
	greater int
}

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

func genAll(length int, char string) string {
	result := ""
	for i:=0;i<length;i++ {
		result += char
	}
	return result
}

func incGuess (g string, pos int) string {
	res := g[0:5]
	//int nnum
	nnum := strconv.Atoi(g[5:6])
	nnum++
	res += strconv.Itoa(nnum)
	res += g[6:len(g)]
	return res
}

func bruteForce(answer string) int {
	guesses, lofa := 0, len(answer)
	cur := 0
	previous := guess{genAll(lofa,"0"), 0,0,0}
	previous.exact, previous.less, previous.greater = evaluate(previous.gstring,answer)
	guesses++
	current := guess{incGuess(previous.gstring,cur),0,0,0}

	/* gen++, evaluate, then compare results
	   if exact, cur++ and gen++, if lower, last guess, cur++ and g++
	   else gen++
	for cur < lofa {
		
	} */
	   	   
	fmt.Println(previous)
	fmt.Println(current)
	return guesses
}

func main() {
	rand.Seed(time.Now().Unix())
	answer := generate(20)
	fmt.Println("The answer is = "+answer)

	fmt.Println("Brute force")
	guesses := bruteForce(answer)
	fmt.Println("It took "+strconv.Itoa(guesses) + " guesses")
    
    

}

