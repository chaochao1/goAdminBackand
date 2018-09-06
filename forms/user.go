package forms

type Login struct {
	Username	string	`form:"username" json:"username" binding:"required"`
	Password	string	`form:"password" json:"password" binding:"required"`
}

type User struct {
	UserName 	string	`form:"username" json:"username" binding:"required"`
	FirstName	string	`form:"first_name" json:"first_name" binding:"required"`
	LastName 	string	`form:"last_name" json:"last_name"`
}

type Register struct {
	Username  			string		`form:"username" json:"username" binding:"required"`
	RealName  			string		`form:"real_name" json:"real_name"`
	Email 				string		`form:"email" json:"email"`
	Password			string		`form:"password" json:"password"`
}