package projectory

import (
	"fmt"
	"math/rand"

	"github.com/moul/number-to-words"
)

type Color struct {
	R int
	G int
	B int
}

type Card struct {
	Color  Color
	Number int
	Lines  []string
	Result int
}

type Project struct {
	Color Color
	Cards []Card
	Total int
}

var cardsPerProject int = 48

func generateColors() []Color {
	return []Color{
		Color{0xff, 0x00, 0x00},
		Color{0x00, 0x80, 0x00},
		Color{0x00, 0x00, 0xFF},
		Color{0x80, 0x80, 0x00},
		Color{0x80, 0x00, 0x80}}
}

func GenerateProjects() []Project {
	LOG.Println("generating projects...")
	var colors []Color = generateColors()
	var projects []Project = make([]Project, len(colors))
	for index, color := range colors {
		projects[index] = generateProject(color)
	}
	return projects
}

func generateProject(color Color) Project {
	total, cards := generateCards(color)
	return Project{color, cards, total}
}

func generateCards(color Color) (int, []Card) {
	var total int
	var cards []Card = make([]Card, cardsPerProject)
	for i := 0; i < cardsPerProject; i++ {
		result, lines := generateCard()
		cards[i] = Card{color, i + 1, lines, result}
		total += result
	}
	return total, cards
}

func generateCard() (int, []string) {
	switch rand.Intn(9) {
	case 0:
		return additionOperation()
	case 1:
		return substractionOperation()
	case 2:
		return multiplicationOperation()
	case 3:
		return divisionOperation()
	case 4:
		return romanAdditionOperation()
	case 5:
		return romanSubstractionOperation()
	case 6:
		return englishAdditionOperation()
	case 7:
		return englishMultiplicationOperation()
	case 8:
		return englishSubstractionOperation()
	}
	LOG.Panicln("Unhandled switch case!")
	return 0, nil
}

func additionOperation() (int, []string) {
	result := rand.Intn(12) + 3
	x := rand.Intn(result-2) + 2
	y := result - x
	return result, []string{fmt.Sprintf("%d + %d", x, y)}
}

func substractionOperation() (int, []string) {
	result := rand.Intn(30) + 1
	x := rand.Intn(150-result) + 1 + result
	y := x - result
	return result, []string{fmt.Sprintf("%d - %d", x, y)}
}

func multiplicationOperation() (int, []string) {
	r := rand.Intn(95) + 5
	xr := r/3 - 2
	if xr < 1 {
		xr = 1
	}
	x := rand.Intn(xr) + 2
	y := r / x
	result := x * y
	return result, []string{fmt.Sprintf("%d * %d", x, y)}
}

func divisionOperation() (int, []string) {
	result := rand.Intn(100) + 1
	y := rand.Intn(50) + 1
	x := result * y
	return result, []string{fmt.Sprintf("%d / %d", x, y)}
}

func romanAdditionOperation() (int, []string) {
	result := rand.Intn(24) + 5
	x := rand.Intn(result-4) + 4
	y := result - x
	return result, []string{ntw.IntegerToRoman(x), "plus", ntw.IntegerToRoman(y)}
}

func romanSubstractionOperation() (int, []string) {
	result := rand.Intn(50) + 1
	x := rand.Intn(200-result) + 1 + result
	y := x - result
	return result, []string{ntw.IntegerToRoman(x), "minus", ntw.IntegerToRoman(y)}
}

func englishAdditionOperation() (int, []string) {
	result := rand.Intn(25) + 5
	x := rand.Intn(result-4) + 4
	y := result - x
	return result, []string{ntw.IntegerToEnUs(x), "plus", ntw.IntegerToEnUs(y)}
}

func englishMultiplicationOperation() (int, []string) {
	r := rand.Intn(97) + 3
	x := rand.Intn(r/2) + 2
	y := r / x
	result := x * y
	return result, []string{ntw.IntegerToEnUs(x), "times", ntw.IntegerToEnUs(y)}
}

func englishSubstractionOperation() (int, []string) {
	result := rand.Intn(30) + 1
	x := rand.Intn(99-result) + 1 + result
	y := x - result
	return result, []string{ntw.IntegerToEnUs(x), "minus", ntw.IntegerToEnUs(y)}
}
