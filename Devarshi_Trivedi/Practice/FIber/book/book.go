package book

import (
	"main/database"

	"github.com/gofiber/fiber"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating string `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	var bookList []Book
	rec, err := database.Dbconn.Query("SELECT Title, Author, Rating from BookTable")
	if err != nil {
		panic(err)
	}
	for rec.Next() {
		var singleBook Book
		rec.Scan(&singleBook.Title, &singleBook.Author, &singleBook.Rating)
		bookList = append(bookList, singleBook)
	}
	c.JSON(bookList)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	rec := database.Dbconn.QueryRow("SELECT Title, Author, Rating FROM BookTable WHERE id=?", id)
	var singleBook Book
	rec.Scan(&singleBook.Title, &singleBook.Author, &singleBook.Rating)
	c.JSON(singleBook)
}

func NewBook(c *fiber.Ctx) {
	// var singleBook Book
	// singleBook.Title = "Jorja"
	// singleBook.Author = "jorjajafioya"
	// singleBook.Rating = "5"
	singleBook := new(Book)
	if err := c.BodyParser(singleBook); err != nil {
		c.Status(503).Send(err)
		return
	}
	_, err := database.Dbconn.Exec("INSERT INTO BookTable (Title,Author,Rating) VALUES (?,?,?)", singleBook.Title, singleBook.Author, singleBook.Rating)
	if err != nil {
		panic(err)
	}
	c.Send("Book Successfully Added")
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")

	_, err := database.Dbconn.Exec("DELETE FROM BookTable where id=?", id)
	if err != nil {
		panic(err)
	}
}

func UpdateBook(c *fiber.Ctx) {
	id := c.Params("id")
	singleBook := new(Book)
	err := c.BodyParser(singleBook)
	if err != nil {
		c.Status(503).Send(err)
		return
	}
	_, err = database.Dbconn.Exec("UPDATE BookTable SET Title=?, Author=?, Rating=? WHERE id=?", singleBook.Title, singleBook.Author, singleBook.Rating, id)
	if err != nil {
		c.Status(503).Send(err)
		panic(err)
	}

}
