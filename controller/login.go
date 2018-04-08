package controller

import (
	"net/http"
	"text/template"
)

type LoginController struct{}


func (this *LoginController)ServeHTTP(w http.ResponseWriter,r *http.Request){
	if r.Method=="GET"{
		isExit:=r.FormValue("exit")=="yes"		//如果点登录退出按钮则删除cookie
		if isExit{
			cookie_u:=http.Cookie{
				Name:"uname",
				Value:"",
				MaxAge:-1,
				Path:"/",
			}
			cookie_p:=http.Cookie{
				Name:"pwd",
				Value:"",
				MaxAge:-1,
				Path:"/",
			}
			http.SetCookie(w,&cookie_u)
			http.SetCookie(w,&cookie_p)
		}
		templateVar :=make(map[string]bool)		//定义网页中留下的几个模板变量，返回登录页面
		t,_:=template.ParseFiles("statics/login.html")
		t.Execute(w,templateVar)					//解析网页

	}else {										//处理登录请求的表单
		uname:=r.FormValue("uname")
		pwd:=r.FormValue("pwd")
		check:=r.FormValue("auto_login")=="on"
		if "wang"==uname&&"yong"==pwd{			//如果成功且勾选’记住密码‘就保存cookie，登录失败就刷新登录页面
			maxAge:=0
			if check{
				maxAge=1<<30
			}
			http.SetCookie(w,&http.Cookie{Name:"uname",Value:uname,MaxAge:maxAge,Path:"/"})
			http.SetCookie(w,&http.Cookie{Name:"pwd",Value:pwd,MaxAge:maxAge,Path:"/"})
			http.Redirect(w,r,"/",301)
		}else {
			templateVar :=make(map[string]bool)		//定义网页中留下的几个模板变量，返回账号密码错误
			templateVar["IsError"]=true
			t,_:=template.ParseFiles("statics/login.html")
			t.Execute(w,templateVar)					//解析网页
			//http.Redirect(w,r,"/login",301)
		}
		return
	}
}


