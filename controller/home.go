package controller

import (
	"net/http"
	"html/template"
	"my_blog/models"
)

//控制主页的显示。。。
type HomeController struct{}


func (this *HomeController)ServeHTTP(w http.ResponseWriter,r *http.Request){

	data:=models.GetAllTopic()
	templateVar :=make(map[string]models.Data)			//设置模板里的模板标签，models.Data是个空interface,存放所有类型
	templateVar["Topics"]=data
	templateVar["IsLogin"]=CheckAccount(r)
	t,_:=template.ParseFiles("statics/home.html")
	t.Execute(w,templateVar)

}


func CheckAccount(r *http.Request)bool{			//根据cookie用于判断页面上的登录，退出状态
	ck,e:= r.Cookie("uname")
	if e!=nil{
		return false
	}
	uname:=ck.Value

	ck,e=r.Cookie("pwd")
	if e!=nil{
		return false
	}
	pwd:=ck.Value
	return ("账户"==uname)&&("密码"==pwd)
}