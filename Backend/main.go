package main

import "fmt"

type Book struct {
	Title string
	Pages int
	IsRead bool
}

func (b *Book) MarkAsRead() {
	b.IsRead = true
}

func (b Book) Summary() string {
	status := "Unread"
	if b.IsRead {
		status = "Read"
	}
	return fmt.Sprintf("%s (%d pages) - %s", b.Title, b.Pages, status)
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

func main() {
	myBooks := []Book{ 
		{Title: "Moby Dick", Pages: 150, IsRead: true},
		{Title: "Raytracer Challange", Pages: 361, IsRead: false},
		{Title: "How to get over a Breakup", Pages: 211, IsRead: false},
	}
	pageCount := map[string]int{
		"Moby Dick": 150,
		"Raytracer Challange": 361,
		"How to get over a Breakup": 211,
	}
	fmt.Println(CountUnread(myBooks))
	val, exists := pageCount["Moby Dick"]
	fmt.Println(val, exists)
	val, exists = pageCount["Easy"]
	fmt.Println(val, exists)
}
