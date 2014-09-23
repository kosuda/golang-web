package model

// User structure
type User struct {
	Name    string   ",omitempty"
	Age     int      ",omitempty"
	Hobbies []string ",omitempty"
	Kids    *User    ",omitempty"
}
