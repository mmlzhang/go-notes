// 语言结构体
package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	// 创建一个新的结构体
	fmt.Println(Books{"GO 语言", "www.runoob.com", "go learning", 10001})

	// key-value形式
	fmt.Println(Books{title: "go title", author: "runoob", subject: "golang", book_id: 111})

	// 忽略部分字段
	fmt.Println(Books{title: "Go", author: "mlzhang"})

	book := Books{title: "Go", author: "mlzhang"}
	fmt.Println(book.title)
	fmt.Println(book.author)

	ptrPrint(&book)
}

func printBook(book Books) {
	fmt.Printf("titile: %s\n", book.title)
	fmt.Printf("author: %s\n", book.author)
	fmt.Printf("subject: %s\n", book.subject)
	fmt.Printf("book_id: %d\n", book.title)
}

func ptrPrint(book *Books) {
	fmt.Println("prt Print")
	fmt.Printf("titile: %s\n", book.title)
	fmt.Printf("author: %s\n", book.author)
	fmt.Printf("subject: %s\n", book.subject)
	fmt.Printf("book_id: %d\n", book.book_id)
}
