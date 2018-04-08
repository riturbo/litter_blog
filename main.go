package main

import (
	"net/http"
	"log"
	"my_blog/controller"
	"os"
)




func main(){
	mux:=http.NewServeMux()
	wd,err:=os.Getwd()
	mux.Handle(	"/statics",http.StripPrefix("/statics",http.FileServer(http.Dir(wd))))		//这个路由用来传输静态资源的
	mux.Handle("/categroy",&controller.Categroy{})
	mux.Handle("/login",&controller.LoginController{})
	mux.Handle("/",&controller.HomeController{})
	mux.Handle("/topic",&controller.TopicController{})
	mux.Handle("/topic/add",&controller.AddTpoicCon{})
	mux.Handle("/topic/view",&controller.TopicView{})
	err=http.ListenAndServe(":8888",mux)
	cheackErr(err)
}


func cheackErr(err error){
	if err!=nil{
		log.Fatal(err)
	}
}

