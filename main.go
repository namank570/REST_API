package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{}

func main() {

	var book Book

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println("Error occured")
	}

	defer db.Close()
	r := gin.Default()

	r.GET("/books", func(c *gin.Context) {

		rows, err := db.Query("select * from book")

		if err != nil {
			panic(err)
		}
		for rows.Next() {

			err := rows.Scan(&book.ID, &book.Title, &book.Author)

			if err != nil {
				panic(err)
			}
			books = append(books, book)

			// fmt.Printf("ID: '%s', Title: '%s', Author: '%s''\n'", book.ID, book.Title, book.Author)
		}

		c.JSON(200, books)

		books = nil
		defer rows.Close()
	})

	r.GET("/books/:id", func(c *gin.Context) {

		var book Book

		id := c.Param("id")
		row := db.QueryRow("select * from book where ID = ?;", id)
		err = row.Scan(&book.ID, &book.Title, &book.Author)
		// fmt.Printf("ID: '%s', Title: '%s', Author: '%s''\n'", book1.ID, book1.Title, book1.Author)
		c.JSON(200, book)

	})

	r.POST("/books", func(c *gin.Context) {

	})

	r.POST("/bookss", func(c *gin.Context) {
		var book Book
		err := c.BindJSON(&book)
		// fmt.Println(book.ID, book.Title, book.Author)

		if err != nil {
			panic(err)
		}

		stmt, err := db.Prepare("insert into book values(?,?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}

		_, err = stmt.Exec(book.ID, book.Title, book.Author)
		if err != nil {
			fmt.Print(err.Error())
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Inserted successfully",
		})

		c.JSON(200, book)
	})

	r.DELETE("/books/:id", func(c *gin.Context) {

		ID := c.Param("id")

		stmt, err := db.Prepare("delete from book where ID= ?;")

		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(ID)

		if err != nil {
			fmt.Print(err.Error())
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Deletion Done ",
		})

	})

	r.PUT("/book/:id", func(c *gin.Context) {

		var book Book
		err := c.BindJSON(&book)
		if err != nil {
			fmt.Print(err.Error())
		}

		ID := c.Param("id")

		stmt, err := db.Prepare("update book set Author= ? where ID = ?;")
		if err != nil {
			fmt.Print(err.Error())
		}

		_, err = stmt.Exec("Naman Kumar", ID)
		if err != nil {
			fmt.Print(err.Error())
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "PUT successfully",
		})

	})
	r.Run()

}
