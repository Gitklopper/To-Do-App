package main

import "fmt"

type Describable interface {
	Describe() string
}

type Book struct {
	Title string
	Pages int
	IsRead bool
}
type Movie struct {
	Title string
	Length int
}

func (b *Book) MarkAsRead() {
	b.IsRead = true
}

func (b Book) Describe() string {
	status := "Unread"
	if b.IsRead {
		status = "Read"
	}
	return fmt.Sprintf("%s (%d pages) - %s", b.Title, b.Pages, status)
}

func (m Movie) Describe() string {
	return fmt.Sprintf("%s (%d min)", m.Title, m.Length)
}

func PrintDescription(s Describable) {
	fmt.Println(s.Describe())
}

func CountUnread(myBooks []Book) int {
	count := 0

	for _, book := range myBooks {
		if !book.IsRead {
			count++
		}
	}
	return count
}

func IncrementPageCount(pages *int) {
	*pages++
}

func SwapPageCount(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func main() {
	myBooks := []Book{ 
		{Title: "Moby Dick", Pages: 150, IsRead: true},
		{Title: "Raytracer Challange", Pages: 361, IsRead: false},
		{Title: "How to get over a Breakup", Pages: 211, IsRead: false},
	}
	myMovie := Movie{Title: "Idiocracy", Length: 220}
	PrintDescription(myBooks[0])
	PrintDescription(myMovie)
}
