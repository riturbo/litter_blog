package models

import (
"time"
_"github.com/go-sql-driver/mysql"
"database/sql"
	"strconv"
)

const(
	 db_name string="mysql"
	 data_source string="root:mima@/test?charset=utf8"
)
type Categroy struct{
	Id int64
	Title string
	Count int
}

type Topic struct{
	Id int64
	Title string
	Content string
	Updated string
	Views int
	ReplyCount int
}


func AddCategroy(name string)bool{
	db,err:=sql.Open(db_name,data_source)			//判断name是否被使用
	checkError(err)
	row,err:=db.Query("SELECT id FROM categroy where title=?",name)
	if row.Next(){
		return false
	}
	stmt,err:=db.Prepare("INSERT categroy SET title=?,created=?,count=?")
	checkError(err)
	result,err:=stmt.Exec(name,time.Now(),0)
	checkError(err)
	id,err:= result.LastInsertId()
	checkError(err)
	return id>0
}

func DelCategroy(id string)bool{	//判断手否数据错误
	db,err:=sql.Open(db_name,data_source)
	checkError(err)
	tid, err := strconv.ParseInt(id, 10, 64)
	checkError(err)
	row,err:=db.Query("SELECT id FROM categroy where id=?",tid)
	checkError(err)
	if !row.Next(){
		return false
	}
	stmt,err:=db.Prepare("DELETE FROM categroy WHERE id=?")
	checkError(err)
	result,err:=stmt.Exec(tid)
	checkError(err)
	affect,err:=result.RowsAffected()
	return affect>0
}

func GetAllCategroy()[]Categroy{
	db,err:=sql.Open(db_name,data_source)
	checkError(err)
	row,err:=db.Query("SELECT id,title,count FROM categroy")
	checkError(err)
	var categroies []Categroy
	for row.Next(){
		var temp Categroy
		err=row.Scan(&temp.Id,&temp.Title,&temp.Count)
		checkError(err)
		categroies=append(categroies,temp)
	}
	return categroies
}

func AddTopic(title ,content string) error{
	db,err:=sql.Open(db_name,data_source)			//判断name是否被使用
	checkError(err)

	stmt,err:=db.Prepare("INSERT topic SET title=?,content=?,created=?,updated=?")
	checkError(err)
	result,err:=stmt.Exec(title,content,time.Now(),time.Now())
	checkError(err)
	_,err= result.LastInsertId()
	return err
}

func GetAllTopic()[]Topic{								//
	db,err:=sql.Open(db_name,data_source)
	checkError(err)
	row,err:=db.Query("SELECT id,title,LEFT(content,70),updated,views,reply_count FROM topic")
	checkError(err)
	var topics []Topic
	for row.Next(){
		var temp Topic
		var x []uint8
		err=row.Scan(&temp.Id,&temp.Title,&temp.Content,&x,&temp.Views,&temp.ReplyCount)
		temp.Updated=string(x)
		checkError(err)
		topics=append(topics,temp)
	}
	return topics
}

func GetTheTopic(id string)Topic{				//没啥好说的...
	db,err:=sql.Open(db_name,data_source)
	checkError(err)
	str:="SELECT title,content,updated,views,reply_count FROM topic where id="+id
	row,err:=db.Query(str)
	checkError(err)
	var topicView Topic
	for row.Next(){
		var x []uint8
		err=row.Scan(&topicView.Title,&topicView.Content,&x,&topicView.Views,&topicView.ReplyCount)
		topicView.Updated=string(x)
		checkError(err)
	}

	stmt,err:=db.Prepare("update topic set views=? where id=?")
	checkError(err)
	_,err=stmt.Exec(topicView.Views+1,id)
	checkError(err)

	return topicView
}


func checkError(err error){
	if err!=nil{
		print(err)
	}
}

