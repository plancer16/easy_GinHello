package model

import (
	"diy_ginHello/initDB"
	"log"
)

type User struct {
	Id int `form:"id"`
	Email string `form:"email" binding:"email"`
	Password string `form:"password"`
}

func (user *User) Save() int64 {
	res, err := initDB.Db.Exec("insert into user (email, password) values (?,?);", user.Email, user.Password)
	if err != nil {
		log.Panicln("user insert error", err.Error())
	}
	id, e := res.LastInsertId()
	if e != nil {
		log.Panicln("user insert id error", err.Error())
	}
	return id
}

func (user *User) QueryByEmail() User{
	u := User{}
	row := initDB.Db.QueryRow("select * from user where email=?;", user.Email)
	err := row.Scan(&u.Id, &u.Email, &u.Password)
	if err != nil {
		log.Panicln(err)
	}
	return u
}

func (user *User) QueryById(id int) (User,error){
	u := User{}
	row := initDB.Db.QueryRow("select * from user where id = ?;", id)
	err := row.Scan(&u.Id, &u.Email, &u.Password)
	if err != nil {
		log.Panicln(err)
	}
	return u,err
}

func (user *User) Update(id int) error {
	_, e := initDB.Db.Exec("update user set password =? where id=?", user.Password, id)
	if e != nil {
		log.Panicln("更新错误",e.Error())
	}
	return e
}