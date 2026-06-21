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

func IncrementPageCount(pages *int) {
	*pages = *pages + 1
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
	pageCount := map[int]*int{
		0: &myBooks[0].Pages,
		1: &myBooks[1].Pages,
		2: &myBooks[2].Pages,
	}
	val, exists := pageCount[0]
	fmt.Println(*val, exists)
	IncrementPageCount(&myBooks[0].Pages)
	val, exists = pageCount[0]
	fmt.Println(*val, exists)
	fmt.Println(myBooks[0].Pages)
	fmt.Println(myBooks[1].Pages)
	SwapPageCount(&myBooks[0].Pages, &myBooks[1].Pages)
	fmt.Println(myBooks[0].Pages)
	fmt.Println(myBooks[1].Pages) 
}
