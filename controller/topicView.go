package controller

import (
	"net/http"
	"my_blog/models"
	"text/template"
)

type TopicView struct{}


func (this *TopicView)ServeHTTP(w http.ResponseWriter,r *http.Request){
	data:=models.GetTheTopic(r.FormValue("TopicID"))
	templateVar :=make(map[string]models.Data)
	templateVar["Title"]=data.Title
	templateVar["Content"]=data.Content
	templateVar["Views"]=data.Views
	templateVar["IsTopic"]=true
	templateVar["IsLogin"]=CheckAccount(r)
	t,_:=template.ParseFiles("statics/article.html")
	t.Execute(w,templateVar)
}
