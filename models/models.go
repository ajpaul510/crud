package models

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Models struct{}

type Person struct {
	Id   string
	Name string
	Age  int
}

var (
	CONNECTION_STRING string
)

func (m Models) Start(cs string, p string) {
	CONNECTION_STRING = cs
	db, err := sql.Open("postgres", cs)
	defer db.Close()

	if err != nil {
		s := fmt.Sprintf("Could not open connection to %v", cs)
		log.Fatal(s)
	}

}

func GetUser(c *gin.Context) {
	query := ""
	if id := c.Query("id"); id != "" {
		id = string(id)
		query = fmt.Sprintf("SELECT * FROM person WHERE p_id = %v Limit 1", id)
	} else {
		query = fmt.Sprintf("SELECT * FROM person LIMIT 100")
	}

	db, err := sql.Open("postgres", CONNECTION_STRING)
	if err != nil {
		s := fmt.Sprintf("Could not open connection to %v", CONNECTION_STRING)
		log.Fatal(s)
	}

	defer db.Close()

	rows, err := db.Query(query)

	var results []Person

	for rows.Next() {
		var result Person
		err := rows.Scan(&result.Id, &result.Name, &result.Age)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	c.IndentedJSON(http.StatusOK, results)
}

func PostUser(c *gin.Context) {
	var buffer Person

	if err := c.BindJSON(&buffer); err != nil {
		fmt.Printf("Error reading record %v\n", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
	}

	db, err := sql.Open("postgres", CONNECTION_STRING)
	if err != nil {
		s := fmt.Sprintf("Could not open connection to %v", CONNECTION_STRING)
		log.Fatal(s)
	}
	defer db.Close()

	res, err := db.Exec(
		fmt.Sprintf("INSERT INTO person(p_name, p_age) VALUES ('%v', %v);", buffer.Name, buffer.Age),
	)

	if err != nil {
		fmt.Printf("Error writting record %v\n", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
	}
	_ = res

	c.IndentedJSON(http.StatusOK, nil)
}
