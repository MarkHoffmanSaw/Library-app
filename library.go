package main

import (
	"fmt"
	"time"
)

type Title string
type Name string

type LendAudit struct {
	checkOut time.Time
	checkIn time.Time
}

type Member struct {
	name Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total int
	lended int
}

type Library struct {
	members map[Name]Member
	books map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit :=range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name,":",title,":", audit.checkOut.String(), "through", returnTime)
	}
}

func printMemberAudits(library *Library) {
	for _,member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book :=range library.books {
		fmt.Println(title,"/ total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println()
}

func checkOutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not found!")
		return false
	}

	if book.lended == book.total {
		fmt.Println("No more books available to lend")
		return false
	}

	book.lended += 1
	library.books[title] = book
	member.books[title] = LendAudit{checkOut: time.Now()}

	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not check out this book")
		return false
	}

	book.lended -= 1
	library.books[title] = book
	audit.checkIn = time.Now()
	member.books[title] = audit

	return true
}
 
func main() {
	library := Library{
		books: make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["Investments"] = BookEntry{
		total: 5,
		lended: 0,
	}
	library.books["Finances"] = BookEntry{
		total: 3,
		lended: 0,
	}
	library.books["Personal Development"] = BookEntry{
		total: 6,
		lended: 0,
	}
	library.books["Psychology"] = BookEntry{
		total: 2,
		lended: 0,
	}

	library.members["Arthur"] = Member{"Arthur", make(map[Title]LendAudit)}
	library.members["Lera"] = Member{"Lera", make(map[Title]LendAudit)}

	fmt.Println("Initial:")
	printLibraryBooks(&library)
	printMemberAudits(&library)

	member := library.members["Arthur"]
	checkedOut := checkOutBook(&library, "Investments", &member)
	fmt.Println("Checked out a book:")
	if checkedOut {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}

	returned := returnBook(&library, "Investments", &member)
	fmt.Println("Checked in a book:")
	if returned {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}
}
