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

func main() {
	myBook := Book {Title: "Moby Dick", Pages : 150, IsRead: false}
	fmt.Println(myBook.Summary())
	myBook.MarkAsRead()
	fmt.Println(myBook.Summary())
}
