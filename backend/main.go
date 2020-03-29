package main

import (
	"fmt"
	"github.com/lab259/cors"
	"github.com/valyala/fasthttp"
	"log"
)

func main(){
	if err := Initializer(); err!= nil{
		log.Fatal("Провалена инициализация: ", err)
		return
	}

	Router.GET("/auth", CORSHandler(authPageHandler))
	Router.POST("/registration",CORSHandler(registrationHandler))
	Router.POST("/login",CORSHandler(loginHandler))

	Router.GET("/static/*filepath", CORSHandler(fasthttp.FSHandler("./frontend",1)))
	Router.GET("/frontend/*filepath", CORSHandler(fasthttp.FSHandler("./frontend", 1)))


	Router.GET("/testpage", CORSHandler(testPageHandler))
	Router.GET("/posts", CORSHandler(PostTestHandler))

	Router.POST("/update_basic_info/text_data", CORSHandler(UpdateBasicInfoTextHandler))
	Router.POST("/update_basic_info/profile_avatar", CORSHandler(UpdateProfileAvatar))
	Router.POST("/update_basic_info/profile_bg", CORSHandler(UpdateProfileBg))


	Router.NotFound = CORSHandler(testPageHandler)

	fmt.Println("LISTENING ON PORT " + ServePort)

	if err:= fasthttp.ListenAndServe(":" + ServePort, Router.Handler); err!= nil{
		log.Println("error when starting server: " + err.Error())
	}

	if err := Postgres.Close(); err!=nil{
		log.Println("error when closing Postgres conn: " + err.Error())
	}
	if err := Redis.Close(); err!=nil{
		log.Println("error when closing Redis conn: " + err.Error())
	}
}

func redirectHandler(ctx *fasthttp.RequestCtx){
	ctx.Redirect("/auth",302)
}

func testPageHandler(ctx *fasthttp.RequestCtx){
	ctx.SendFile("frontend/build/index.html")
}

func CORSHandler (h fasthttp.RequestHandler) fasthttp.RequestHandler{
	return cors.AllowAll().Handler(h)
}