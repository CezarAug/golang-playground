package models

type Person struct {
	Id    int    `json:id`
	Name  string `json:"name"`
	About string `json:"about"`
}

var Persons []Person

// When running gorm, Person was translated into "people".
func (Person) TableName() string {
	return "persons"
}
