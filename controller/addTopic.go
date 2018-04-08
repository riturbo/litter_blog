package controller

import (
	"net/http"
	"text/template"
	"my_blog/models"
	"github.com/astaxie/beego/logs"
)

type AddTpoicCon struct{}

func (this *AddTpoicCon)ServeHTTP(w http.ResponseWriter,r *http.Request){
	if r.Method=="GET"{
		if !CheckAccount(r){
			http.Redirect(w,r,"/login",302)
			return
		}
		templateVar :=make(map[string]models.Data)
		templateVar["IsTopic"]=true
		templateVar["IsLogin"]=CheckAccount(r)
		t,_:=template.ParseFiles("statics/edit.html")
		t.Execute(w,templateVar)
	}else if r.Method=="POST"{
		if !CheckAccount(r){
			http.Redirect(w,r,"/login",302)
			return
		}
		title:=r.FormValue("title")
		content:=r.FormValue("content")

		var err error

		err=models.AddTopic(title,content)
		if err!=nil{
			logs.Debug(err)
		}
		http.Redirect(w,r,"/topic",302)


	}
}