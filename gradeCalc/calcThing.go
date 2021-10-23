package main

import "fmt"

var firstScore float32 = 66  // the score you got
var firstPart float32 = 30   // percentage of first module
var targetScore float32 = 70 // minimum to pass

func main() {
	otherPart := 100 - firstPart

	overallFirstScore := firstScore * firstPart
	overallFirstScore = overallFirstScore / 100
	fmt.Println("You have ", overallFirstScore, "% already down")

	newNeeded := targetScore - overallFirstScore
	fmt.Println("Which means you need ", newNeeded, "% to go")

	otherPartNeeded := newNeeded / otherPart
	otherPartNeeded = otherPartNeeded * 100

	fmt.Println("Which means you need ", otherPartNeeded, "% on the other part to get ", targetScore, "% overall")
}
