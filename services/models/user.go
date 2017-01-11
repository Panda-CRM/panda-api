package models

type User struct {
	UUID			string	`sql:"type:uuid; primary_key; default:uuid_generate_v4();unique"`
	PersonUUID 		string	`sql:"type:uuid"`
	Active			bool	`sql:"type:boolean"`
	Username		string	`sql:"type:varchar(50); unique"`
	Password		string	`sql:"type:varchar(50)"`
}

type Users []User

func (User) TableName() string {
    return "users"
}