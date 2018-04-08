package controller

import (
	"net/http"
	"text/template"
	"my_blog/models"
)

type TopicController struct{}


func  (this *TopicController)ServeHTTP(w http.ResponseWriter,r *http.Request){
	if r.Method=="GET"{
		data:=models.GetAllTopic()
		templateVar :=make(map[string]models.Data)
		templateVar["Topics"]=data
		templateVar["IsTopic"]=true
		templateVar["IsLogin"]=CheckAccount(r)
		t,_:=template.ParseFiles("statics/topic.html")
		t.Execute(w,templateVar)

	}
}