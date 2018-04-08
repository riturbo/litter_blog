package controller

import (
	"net/http"
	"text/template"
	"my_blog/models"
)

type Categroy struct{}

func (this *Categroy)ServeHTTP(w http.ResponseWriter,r *http.Request){
	if r.Method=="GET"{
		op:=r.FormValue("op")				//根据key判断操作类型
		switch op {
		case "add":
			name:=r.FormValue("name")
			if len(name)==0{
				break
			}

			models.AddCategroy(name)
		case "del":
			id:=r.FormValue("id")
			if len(id)==0{
				break
			}
			models.DelCategroy(id)
		}

		data:=models.GetAllCategroy()
		templateVar :=make(map[string]models.Data)
		templateVar["Categroies"]=data
		templateVar["IsCategroy"]=true
		templateVar["IsLogin"]=CheckAccount(r)
		t,_:=template.ParseFiles("statics/categroy.html")
		t.Execute(w,templateVar)
	}
}
