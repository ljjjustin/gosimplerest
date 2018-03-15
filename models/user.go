package models

type User struct {
	Id       int    `json:"id" xorm:"pk autoincr"`
	Name     string `json:"name" xorm:"unique notnull varchar(64)"`
	NickName string `json:"nick_name" xorm:"notnull varchar(128)"`
	Age      int    `json:"age" xorm:"notnull default 0"`
	Email    string `json:"email" xorm:"not null default ''"`
}

type Users []User

func UserCreate(user *User) error {
	if _, err := db.Insert(user); err != nil {
		return err
	}
	return nil
}

func UserGet(user *User) (bool, error) {
	return db.Where("id=?", user.Id).Get(user)
}

func UserGetAll() ([]*User, error) {
	var users []*User
	err := db.Iterate(new(User), func(i int, bean interface{}) error {
		user := bean.(*User)
		users = append(users, user)
		return nil
	})
	return users, err
}

func UserUpdate(user *User) (int64, error) {
	return db.Where("id=?", user.Id).Update(user)
}

func UserDelete(user *User) error {
	_, err := db.Where("id=?", user.Id).Delete(user)
	return err
}
