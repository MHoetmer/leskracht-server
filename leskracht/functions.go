package leskracht

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "mel"
	dbname = "leskracht"
)

func ConnectToDb() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Query the database.
	rows, err := db.Query(`
		SELECT
			id,
			firstName,
			lastName,
			email
		FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {

		var (
			id        int
			firstName string
			lastName  string
			email     string
		)

		if err := rows.Scan(&id, &firstName, &lastName, &email); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d, %s, %s, %s\n", id, firstName, lastName, email)
	}

	// Check for errors after we are done iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func GetUser(name string) User {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Query the database.
	rows, err := db.Query(`
		SELECT
			id,
			firstName,
			lastName,
			email
		FROM users
		WHERE firstName = $1`, name)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var (
		id        int
		firstName string
		lastName  string
		email     string
	)
	for rows.Next() {

		if err := rows.Scan(&id, &firstName, &lastName, &email); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d, %s, %s, %s\n", id, firstName, lastName, email)
	}

	// Check for errors after we are done iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	var user User
	user.ID = id
	user.FirstName = firstName
	user.LastName = lastName
	user.Email = email

	return user
}

func CreateUser(firstName string, secondName string, email string, birthDate int) User {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Query the database.
	rows, err := db.Query(`
		INSERT INTO users (firstName, lastName, email, birthDate)
		VALUES($1, $2, $3, $4)`, firstName, secondName, email, birthDate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mutation", rows)
	defer rows.Close()
	// var (
	// 	id        int
	// 	firstName string
	// 	lastName  string
	// 	email     string
	// )
	// for rows.Next() {

	// 	if err := rows.Scan(&id, &firstName, &lastName, &email); err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Printf("%d, %s, %s, %s\n", id, firstName, lastName, email)
	// }

	// Check for errors after we are done iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	var user User
	// user.ID = id
	// user.FirstName = firstName
	// user.LastName = lastName
	// user.Email = email

	return user
}

func GetMessage(id int) Message {
	var message Message
	return message
}
