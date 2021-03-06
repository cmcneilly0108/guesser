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

func genSplit(length int, charS string, charE string) string {
	result := ""
	for i:=0;i<length;i++ {
		if (i*2 < length) {
			result += charS
		} else {
			result += charE
		}
	}
	return result
}

func incGuess (g string, pos int) string {
	res := g[0:pos]
	nnum, err := strconv.Atoi(g[pos:pos+1])
	if err != nil {
		fmt.Println(err)
        }
	nnum++
	res += strconv.Itoa(nnum)
	res += g[pos+1:len(g)]
	return res
}

func modGuess (g string, pos int, rep int) string {
	res := g[0:pos]
	res += strconv.Itoa(rep)
	res += g[pos+1:len(g)]
	return res
}

func decGuess (g string, pos int) string {
	res := g[0:pos]
	nnum, err := strconv.Atoi(g[pos:pos+1])
	if err != nil {
		fmt.Println(err)
        }
	nnum--
	res += strconv.Itoa(nnum)
	res += g[pos+1:len(g)]
	return res
}

func inc2Guess (g string, pos int) string {
	res := g[0:pos]
	nnum, err := strconv.Atoi(g[pos:pos+1])
	if err != nil {
		fmt.Println(err)
        }
        if (nnum < 8) {
        	nnum = nnum + 2
        } else {
        	nnum++
        }
	res += strconv.Itoa(nnum)
	res += g[pos+1:len(g)]
	return res
}

func bruteForce(answer string) int {
	/* This only ever uses the exact data from the evaluate function.
	   We can do better. */
	guesses, lofa := 0, len(answer)
	cur := 0
	previous := guess{genAll(lofa,"0"), 0,0,0}
	previous.exact, previous.less, previous.greater = evaluate(previous.gstring,answer)
	guesses++
	current := guess{"",0,0,0}

	for cur < lofa { 
		current = guess{incGuess(previous.gstring,cur),0,0,0}
		current.exact, current.less, current.greater = evaluate(current.gstring,answer)
		//fmt.Println(current)
		guesses++
		if (current.exact > previous.exact) {
			cur++
			previous = current
		} else {
			if (current.exact < previous.exact) {
				cur++
			} else {
				previous = current
			}
		}
	}
	//fmt.Print("The correct answer is =")
	//fmt.Println(previous)
	return guesses
}

func binarySearch(answer string) int {
	guesses, lofa := 0, len(answer)
	cur, glo, ghi, gcur := 0,0,9,0
	previous := guess{genAll(lofa,"0"), 0,0,0}
	previous.exact, previous.less, previous.greater = evaluate(previous.gstring,answer)
	guesses++
	current := guess{"",0,0,0}
	//fmt.Println("The answer is = "+answer)

	for cur < lofa {
		gcur = (glo +ghi + 1)/2
		//fmt.Print(glo)
		//fmt.Print(" and ")
		//fmt.Println(ghi)
		current = guess{modGuess(previous.gstring,cur,gcur),0,0,0}
		current.exact, current.less, current.greater = evaluate(current.gstring,answer)
		//fmt.Println(current)
		guesses++
		if (current.exact > previous.exact) {
			cur++
			previous = current
			glo, ghi = 0,9
		} else {
			if (current.exact < previous.exact) {
				cur++
				glo, ghi = 0,9
			} else {
				if (current.greater > previous.greater) {
					ghi = gcur - 1
				} else {
					glo = gcur + 1
					previous = current
				}
			}
		}
	}
	//fmt.Print("The correct answer is =")
	//fmt.Println(previous)
	return guesses
}

func biBrute(answer string) int {
	guesses, lofa := 0, len(answer)
	fcur, bcur := 0, len(answer)-1
	fdone, bdone, fhit, bhit := false, false, true, true
	previous := guess{genSplit(lofa,"0","9"), 0,0,0}
	previous.exact, previous.less, previous.greater = evaluate(previous.gstring,answer)
	//fmt.Println(previous)
	guesses++
	current := guess{"",0,0,0}

	//for current.exact < len(answer) { 
	//for z:=0;z<30;z++ {
	for (!fdone || !bdone) {
		// Try - if I have just gotten one, don't change the other
		current = guess{previous.gstring,0,0,0}
		if (!fdone && (fhit || !bhit)) {
			current.gstring = incGuess(current.gstring,fcur)
			if (!bdone && (!bhit && !fhit)) {
				current.gstring  = decGuess(current.gstring,bcur)
				bhit = false
			}
			fhit = false
		} else {
			current.gstring  = decGuess(current.gstring,bcur)
			bhit = false
		}

		current.exact, current.less, current.greater = evaluate(current.gstring,answer)
		//fmt.Println(current)
		guesses++
		switch {
		case (current.exact == previous.exact+2):
			fhit = true
			fcur++
			bhit = true
			bcur--
			previous = current
		case (current.exact == previous.exact):
			previous = current
		case (current.exact == previous.exact-2):
			fhit = true
			fcur++
			bhit = true
			bcur--
		case (current.exact == previous.exact+1):
			if (current.less < previous.less) {
				fhit = true
				fcur++
			} else {
				bhit = true
				bcur--
			}
			previous = current
		case (current.exact == previous.exact-1):
			if (current.greater > previous.greater) {
				fhit = true
				fcur++
			} else {
				bhit = true
				bcur--
			}
		default:
			previous = current
		}
		if (fcur*2 >= lofa) {
			fdone, fhit = true, false
			//fmt.Print("fdone!")
		}
		if (bcur*2 < lofa) {
			bdone, bhit = true, false
			//fmt.Print("bdone!")
		}
	}
	//fmt.Print("The correct answer is =")
	//fmt.Println(previous)
	return guesses
}

func skipOne(answer string) int {
	guesses, lofa := 0, len(answer)
	cur := 0
	previous := guess{genAll(lofa,"0"), 0,0,0}
	previous.exact, previous.less, previous.greater = evaluate(previous.gstring,answer)
	guesses++
	current := guess{"",0,0,0}

	for cur < lofa { 
		current = guess{inc2Guess(previous.gstring,cur),0,0,0}
		current.exact, current.less, current.greater = evaluate(current.gstring,answer)
		//fmt.Print(current)
		guesses++
		//fmt.Print(" guess #")
		//fmt.Println(guesses)
		if (current.exact > previous.exact) {
			cur++
			previous = current
		} else {
			if (current.exact < previous.exact) {
				cur++
			} else {
				if (current.greater > previous.greater) {
					previous = guess{decGuess(current.gstring,cur),0,0,0}
					previous.exact, previous.less, previous.greater = evaluate(previous.gstring,answer)
					cur++
				} else {
					previous = current
				}
			}
		}
	}
	//fmt.Print("The correct answer is =")
	//fmt.Println(previous)
	return guesses
}

// TBD - binary guessing
// TBD - biSkip - is possible?
// TBD - does finding how many of each num help?
// Am I making use of all 3 pieces of information with each guess?

func main() {
	rand.Seed(time.Now().Unix())
	const tests = 1000
	numSize := 50
	total := 0
	var agd float64
	var scores [tests] int

	fmt.Println("Brute force")
	for j:=0;j<tests;j++ {
		answer := generate(numSize)
		//fmt.Println("The answer is = "+answer)
		scores[j] = bruteForce(answer)
		//fmt.Println("It took "+strconv.Itoa(scores[j]) + " guesses")
		total += scores[j]
	}
	fmt.Print("Average guesses/digit = ")
	agd = float64(total)/(float64(tests)*float64(numSize))
	fmt.Println(agd)

	fmt.Println("Skip One")
	total = 0
	for j:=0;j<tests;j++ {
		answer := generate(numSize)
		//fmt.Println("The answer is = "+answer)
		scores[j] = skipOne(answer)
		//fmt.Println("It took "+strconv.Itoa(scores[j]) + " guesses")
		total += scores[j]
	}
	fmt.Print("Average guesses/digit = ")
	agd = float64(total)/(float64(tests)*float64(numSize))
	fmt.Println(agd)

	fmt.Println("Bi Brute")
	total = 0
	for j:=0;j<tests;j++ {
		answer := generate(numSize)
		//answer := "25192812505289738406"
		//fmt.Println("The answer is = "+answer)
		scores[j] = biBrute(answer)
		//fmt.Println("It took "+strconv.Itoa(scores[j]) + " guesses")
		total += scores[j]
	}
	fmt.Print("Average guesses/digit = ")
	agd = float64(total)/(float64(tests)*float64(numSize))
	fmt.Println(agd)

	fmt.Println("Binary Search")
	total = 0
	for j:=0;j<tests;j++ {
		answer := generate(numSize)
		//fmt.Println("The answer is = "+answer)
		scores[j] = binarySearch(answer)
		//fmt.Println("It took "+strconv.Itoa(scores[j]) + " guesses")
		total += scores[j]
	}
	fmt.Print("Average guesses/digit = ")
	agd = float64(total)/(float64(tests)*float64(numSize))
	fmt.Println(agd)
}

