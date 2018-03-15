package models

type Group struct {
	Id   uint   `json:"id" xorm:"pk autoincr"`
	Name string `json:"name" xorm:"notnull varchar(64)"`
}

type Groups []Group

func GroupCreate(t Group) {
}

func GroupGet(id int) {
}

func GroupUpdate(id int) {
}

func GroupDelete(id int) {
}
