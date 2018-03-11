package projectory

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

var cardWidth float64 = 45
var cardHeight float64 = 67
var pageOffsetX float64 = 15
var pageOffsetY float64 = 14

func PrintProjects(projects []Project, pdf *gofpdf.Fpdf) {
	LOG.Println("printing project cards...")
	for _, proj := range projects {
		paintProject(proj, pdf)
	}
	LOG.Println("printing answers page...")
	pdf.AddPage()
	paintAnswersPage(projects, pdf)
}

func paintProject(project Project, pdf *gofpdf.Fpdf) {
	for i := 0; i < 3; i++ {
		pdf.AddPage()
		pageOfCards := project.Cards[i*16 : (i+1)*16]
		painPageOfCards(pageOfCards, pdf)
	}
}

func painPageOfCards(cards []Card, pdf *gofpdf.Fpdf) {
	paintGrid(pdf)
	paintCardsContents(cards, pdf)
}

func paintGrid(pdf *gofpdf.Fpdf) {
	pdf.SetLineWidth(0.1)
	pdf.SetLineCapStyle("square")
	var x1 float64 = pageOffsetX
	var y1 float64 = pageOffsetY
	var x2 float64 = 195
	var y2 float64 = 282

	for i := 0; i < 5; i++ {
		y := y1 + float64(i)*cardHeight
		pdf.Line(x1, y, x2, y)
	}
	for i := 0; i < 5; i++ {
		x := x1 + float64(i)*cardWidth
		pdf.Line(x, y1, x, y2)
	}
}

func paintCardsContents(cards []Card, pdf *gofpdf.Fpdf) {
	index := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			card := cards[index]
			paintCardContents(card, pageOffsetX+float64(j)*cardWidth, pageOffsetY+float64(i)*cardHeight, pdf)
			index++
		}
	}
}

func paintCardContents(card Card, offsetX float64, offsetY float64, pdf *gofpdf.Fpdf) {
	var colorMarkerOffsetX float64 = 3
	var colorMarkerOffsetY float64 = 59
	var cardNumberOffsetX float64 = 10
	var cardNumberOffsetY float64 = 63
	var textOffsetY float64 = 13
	var textLineHeight float64 = 10
	// color marker:
	pdf.SetFillColor(card.Color.R, card.Color.G, card.Color.B)
	pdf.Rect(offsetX+colorMarkerOffsetX, offsetY+colorMarkerOffsetY, 5, 5, "F")
	// card number:
	pdf.SetFont("Courier", "", 14)
	pdf.Text(offsetX+cardNumberOffsetX, offsetY+cardNumberOffsetY, fmt.Sprintf("%d", card.Number))
	// card text:
	pdf.SetFont("Helvetica", "", 14)
	for index, line := range card.Lines {
		lineWidth := pdf.GetStringWidth(line)
		lineOffset := (cardWidth - lineWidth) / 2
		pdf.Text(offsetX+lineOffset, offsetY+textOffsetY+(textLineHeight*float64(index)), line)
	}

}

func paintAnswersPage(projects []Project, pdf *gofpdf.Fpdf) {
	var answersOffsetX float64 = 8
	var answersOffsetY float64 = 10
	var projectOffsetX float64 = 38
	pdf.SetLineWidth(0.5)
	pdf.SetFont("Courier", "", 12)
	for projIndex, proj := range projects {
		lineX := answersOffsetX + float64(projIndex)*projectOffsetX
		textX := lineX + 3
		pdf.Line(lineX, answersOffsetY, lineX, answersOffsetY+280)
		pdf.SetFillColor(proj.Color.R, proj.Color.G, proj.Color.B)
		pdf.Rect(textX, answersOffsetY, 5, 5, "F")
		for cardIndex, card := range proj.Cards {
			cardNumberStr := fmt.Sprintf("%d: ", card.Number)
			cardResultStr := fmt.Sprintf("%d", card.Result)
			cardNumberStrLength := pdf.GetStringWidth(cardNumberStr)
			cardAnswerOffsetY := answersOffsetY + 10 + 5*float64(cardIndex)
			pdf.SetTextColor(card.Color.R, card.Color.G, card.Color.B)
			pdf.Text(1+textX, cardAnswerOffsetY, cardNumberStr)
			pdf.SetTextColor(0, 0, 0)
			pdf.Text(1+textX+cardNumberStrLength, cardAnswerOffsetY, cardResultStr)
		}
		totalOffsetY := answersOffsetY + 15 + 5*float64(len(proj.Cards))
		pdf.SetTextColor(0, 0, 0)
		pdf.Text(textX, totalOffsetY, "TOTAL:")
		pdf.Text(textX+5, totalOffsetY+5, fmt.Sprintf("%d", proj.Total))
	}
}

func PrintIncidents(incidents [][]string, pdf *gofpdf.Fpdf) {
	LOG.Println("printing incidents...")
	pdf.SetFont("Helvetica", "", 10)
	pdf.SetTextColor(0, 0, 0)
	fullPages := len(incidents) / 16
	for i := 0; i < fullPages; i++ {
		paintIncidentsPage(incidents[i*16:(i+1)*16], pdf)
	}
	partialPages := len(incidents) % 16
	if partialPages > 0 {
		paintIncidentsPage(incidents[fullPages*16:], pdf)
	}
}

func paintIncidentsPage(incidents [][]string, pdf *gofpdf.Fpdf) {
	pdf.AddPage()
	paintGrid(pdf)
	count := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			x := pageOffsetX + float64(j)*cardWidth
			y := pageOffsetY + float64(i)*cardHeight + 7
			lines := incidents[count]
			for lineIndex, line := range lines {
				lineWidth := pdf.GetStringWidth(line)
				lineOffsetX := (cardWidth - lineWidth) / 2
				lineOffsetY := float64(lineIndex) * 5
				pdf.Text(x+lineOffsetX, y+lineOffsetY, line)
			}
			count++
			if count >= len(incidents) {
				return
			}
		}
	}
}
