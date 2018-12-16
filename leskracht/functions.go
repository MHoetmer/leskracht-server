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
		birthDate string
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
	user.BirthDate = birthDate

	return user
}

func GetAllUsers() []User {
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
			*
		FROM users
		`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var (
		id        int
		firstName string
		lastName  string
		email     string
		birthDate string
	)
	var users []User
	for rows.Next() {

		if err := rows.Scan(&id, &firstName, &lastName, &email, &birthDate); err != nil {
			log.Fatal(err)
		}
		var user User
		user.ID = id
		user.FirstName = firstName
		user.LastName = lastName
		user.Email = email
		user.BirthDate = birthDate

		fmt.Printf("%d, %s, %s, %s\n", id, firstName, lastName, email)
		users = append(users, user)
	}

	// Check for errors after we are done iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users
}

func CreateUser(firstName string, secondName string, email string, birthDate string) User {
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

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	var user User

	return user
}

func DeleteUser(id int) {
	fmt.Println("delete id", id)
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
		DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// Check for errors after we are done iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func GetMessage(id int) Message {
	var message Message
	return message
}
