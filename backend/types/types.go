package types

type User struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}
type ReceivedUser struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
}
type EmailCheck struct {
	Email string `db:"email"`
	Count string `db:"count"`
}
type Error struct {
	Error string
}
type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Date  string `json:"date"`
	Email string `json:"email"`
}
