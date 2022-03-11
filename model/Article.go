package model

import (
	"diy_ginHello/initDB"
	"log"
)

type Article struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

func (article Article) FindById(id int) Article {
	row := initDB.Db.QueryRow("select * from article where id =?", id)
	err := row.Scan(&article.Id, &article.Type, &article.Content)
	if err != nil {
		log.Panicln("绑定发生错误", err.Error())
	}
	return article
}

func (article Article) FindAll() []Article {
	rows, err := initDB.Db.Query("select * from article")
	if err != nil {
		log.Panicln("查询数据失败")
	}
	var articles []Article
	for rows.Next() {
		var a Article
		if e := rows.Scan(&a.Id, &a.Type, &a.Content); e != nil {
			articles = append(articles, a)
		}
	}
	return articles
}

func (article Article) Insert() int {
	res, err := initDB.Db.Exec("insert into article (type,content) values(?,?);", article.Type, article.Content)
	if err != nil {
		log.Panicln("插入失败", err.Error())
	}
	id, _ := res.LastInsertId()
	return int(id)
}

func (article *Article) DeleteOne() {
	if _, err := initDB.Db.Exec("delete from article where id =?;", article.Id); err != nil {
		log.Panicln("无法删除")
	}
}
