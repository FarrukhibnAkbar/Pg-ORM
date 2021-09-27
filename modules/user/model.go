package user

import(
	// "fmt"
	// "log"
	"app/config"
) 

type User struct{
	Id        uint64 	`pg:"user_id" 	  json:"id"`
	Firstname string 	`pg:"firstname"   json:"firstName"`
	Lastname  string 	`pg:"lastname"    json:"lastName"`
	CreatedAt string	`pg:"create_at"   json:"createdAt"`
}

var pg = config.DBConnection()

func NewUser(user []User) error {

   _, err := pg.Model(&user).Insert()
   return err
}


func GetById(id interface{}) (User, error) {

	user := User {}

	err := pg.Model(&user).Where("user_id = ?0", id).Select()

	return user, err
	
}
