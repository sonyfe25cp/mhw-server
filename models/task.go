package models

type Task struct {
	ID     int64
	Title  string
	First  string
	Second string
	Map    string

	Level int
	Key   bool
}
