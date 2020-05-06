package main

import (
	"context"
	"fmt"
	"github.com/lab259/cors"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	if err := Initializer(); err != nil {
		log.Fatal("Провалена инициализация: ", err)
		return
	}

	//Router.GET("/auth", CORSHandler(authPageHandler))
	Router.POST("/registration", CORSHandler(RegistrationHandler))
	Router.POST("/login", CORSHandler(loginHandler))

	Router.GET("/static/*filepath", fasthttp.FSHandler("../frontend", 1))
	Router.GET("/frontend/*filepath", fasthttp.FSHandler("../frontend", 1))

	Router.GET("/music/*filepath", CORSHandler(fasthttp.FSHandler("../music", 1)))

	Router.GET("/posts", CORSHandler(GetPostsHandler))
	Router.GET("/settings/hobbies", CORSHandler(HobbiesHandler))
	Router.GET("/settings/privacy", CORSHandler(PrivacyHandler))
	Router.GET("/settings/edu_emp", CORSHandler(EduEmpHandler))
	Router.GET("/get_comments/", CORSHandler(CommentsTestHandler))
	Router.POST("/leave_comment/", CORSHandler(AddCommentHandler))

	Router.GET("/get_user_music", CORSHandler(GetUserMusicHandler))
	Router.GET("/get_all_music", CORSHandler(GetAllMusicHandler))
	Router.POST("/post_music", CORSHandler(PostMusicHandler))

	Router.GET("/remove_music", CORSHandler(DeleteMusicHandler))

	Router.POST("/like/", CORSHandler(LikeHandler))
	Router.POST("/settings/update_basic_info/text_data", CORSHandler(UpdateBasicInfoTextHandler))
	Router.POST("/settings/update_basic_info/profile_avatar", CORSHandler(AuthMiddleware(UpdateProfileAvatar)))
	Router.POST("/settings/update_basic_info/profile_bg", CORSHandler(AuthMiddleware(UpdateProfileBg)))

	fmt.Println("LISTENING ON PORT " + ServePort)

	server:=fasthttp.Server{MaxRequestBodySize: 1024*1024*1024, Handler: Router.Handler}

	if err := server.ListenAndServe(":"+ServePort); err != nil {
		log.Println("error when starting server: " + err.Error())
	}

	if err := Postgres.Conn.Close(context.Background()); err != nil {
		log.Println("error when closing Postgres conn: " + err.Error())
	}
	if err := Redis.Close(); err != nil {
		log.Println("error when closing Redis conn: " + err.Error())
	}
}

func CORSHandler(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return cors.AllowAll().Handler(h)
}