package main

import (
	"bytes"
	"context"
	"coursework/postgres"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v7"
	"github.com/jackc/pgx/v4"
	"github.com/lab259/cors"
	"github.com/valyala/fasthttp"
	"log"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)



var (
	rc = postgres.RegistrationConn{}
	rdb *redis.Client
	err error
	salt = []byte("Ilya Bychkov")
	r  = router.New()
	v = validator.New()
)

func ValidateSex(fl validator.FieldLevel) bool {
	str:= fl.Field().String()
	return str == "Мужской" || str == "Женский"
}

func main() {
	if err:= rc.CreateConnection("user=postgres dbname=my_coursework_db");err!=nil{
		log.Fatal(err)
	}

	if err := rc.InitDatabasesIfNotExist(); err!=nil{
		log.Fatal(err)
	}

	rdb = redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	if err = rdb.Ping().Err();err!=nil{
		log.Fatal(err)
		return
	}

	if err = v.RegisterValidation("sex", ValidateSex);err!=nil{
		log.Fatal(err)
		return
	}



	r.GET("/auth", CORSHandler(authPageHandler))
	r.POST("/registration",CORSHandler(registrationHandler))
	r.POST("/login",CORSHandler(loginHandler))

	r.GET("/static/*filepath", CORSHandler(fasthttp.FSHandler("./frontend",1)))
	r.GET("/frontend/*filepath", CORSHandler(fasthttp.FSHandler("./frontend", 1)))


	r.GET("/testpage", CORSHandler(testPageHandler))
	r.GET("/posts", CORSHandler(PostTestHandler))

	r.POST("/update_basic_info/text_data", CORSHandler(UpdateBasicInfoTextHandler))

	r.POST("/update_basic_info/profile_avatar", CORSHandler(UpdateProfileAvatar))
	r.POST("/update_basic_info/profile_bg", CORSHandler(UpdateProfileBg))

	r.NotFound = CORSHandler(testPageHandler)
        
	clientPort:="8090"
	fmt.Println("LISTENING ON PORT " + clientPort)


	if err:= fasthttp.ListenAndServe(":" + clientPort, r.Handler); err!= nil{
		log.Println("error when starting server: " + err.Error())
	}
	if err := rc.Close(); err!=nil{
		log.Println("error when closing regDb conn: " + err.Error())
	}
	if err := rdb.Close(); err!=nil{
		log.Println("error when closing regRedis conn: " + err.Error())
	}
}

type BasicInfoStruct struct{
	Sex      string `json:"sex" validate:"sex"`
	Status   int `json:"status" validate:"required,min=0,max=5"`
	Birthday string `json:"birthday"`
	Tel      int `json:"tel"`
	Country  string `json:"country"`
	City     string `json:"city"`
}

func UpdateBasicInfoTextHandler(ctx *fasthttp.RequestCtx){
	obj := &BasicInfoStruct{}

	if err := json.Unmarshal(ctx.PostBody(), obj); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

	if err := v.Struct(obj); err!=nil{
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

	fmt.Printf("%+v\n",obj)
}

func UpdateProfileAvatar(ctx *fasthttp.RequestCtx){
	f, err :=ctx.FormFile("photo")
	if err != nil{
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	err = fasthttp.SaveMultipartFile(f, "./profile_avatars/" + f.Filename)
	if err != nil{
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}

func UpdateProfileBg(ctx *fasthttp.RequestCtx){
	f, err :=ctx.FormFile("photo")
	if err != nil{
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	err = fasthttp.SaveMultipartFile(f, "./profile_bgs/" + f.Filename)
	if err != nil{
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}

func PostTestHandler(ctx *fasthttp.RequestCtx){
	ctx.Response.Header.Set("Content-Type", "application/json")
	_, _ = ctx.Write([]byte(`{"sex":"Женский",
							  "status": 2,
						      "birthday":"2000-02-05", 
							  "tel":88005553535, 
 							  "country": "Россия",
							  "city":"Нижний Новгород"
							}`))
}

func redirectHandler(ctx *fasthttp.RequestCtx){
	ctx.Redirect("/auth",302)
}

func testPageHandler(ctx *fasthttp.RequestCtx){
	ctx.SendFile("frontend/build/index.html")
}

func authPageHandler(ctx *fasthttp.RequestCtx){
	ctx.SendFile("frontend/html/auth.html")
}

func AuthMiddleware(next fasthttp.RequestHandler)fasthttp.RequestHandler{
	return func(ctx *fasthttp.RequestCtx){
		redisAccessToken, err := rdb.Get(ByteSliceToString(ctx.Request.Header.Cookie("userId"))).Result()
		if err == nil && bytes.Compare(ctx.Request.Header.Cookie("accessToken"),
										StringToByteSlice(redisAccessToken)) == 0{
			next(ctx)
		}else{
			ctx.Redirect("/auth",401)
		}
	}
}


type loginStruct struct{
	Email     string `validate:"required,email"`
	Password  string `validate:"required,min=7"`
}

func loginHandler(ctx *fasthttp.RequestCtx){
	ls:= &loginStruct{
		Email:    ByteSliceToString(ctx.FormValue("email")),
		Password: ByteSliceToString(ctx.FormValue("password")),
	}


	err:= v.Struct(ls)

	if err!=nil{
		ctx.Error("not validated", 403)
		return
	}

	var dbToken []byte
	var userId int
	var firstName string
	var lastName string

	rc.Conn.QueryRow(context.Background(),"select user_id, first_name, last_name, token from registration where email = $1 limit 1", ls.Email).Scan(
		&userId,
		&firstName,
		&lastName,
		&dbToken)

	if userToken:=sha512.Sum512(append(StringToByteSlice(ls.Password), salt...)); bytes.Compare(dbToken, userToken[:]) != 0{
		ctx.Error("Incorrect email/pass combination",402)
		return
	}
	successfulAuth(ctx,strconv.Itoa(userId))
	ctx.Redirect("/secretpage",200)

}

func registrationHandler(ctx *fasthttp.RequestCtx){
	firstName:=ctx.FormValue("first_name")
	lastName:=ctx.FormValue("last_name")
	email:= ctx.FormValue("email")
	password:= ctx.FormValue("password")

	if len(email)==0 || len(password)==0{
		ctx.Error("Поля не заполнены",402)
	}

	if err:= rc.Conn.QueryRow(context.Background(),"select 1 from registration where email = $1 limit 1", email).Scan();err != pgx.ErrNoRows{
		ctx.Error("User already exists",402)
	}

	token:=sha512.Sum512(append(password,salt...))
	var userId int
	if err := rc.Conn.QueryRow(context.Background(), "insert into registration (first_name,last_name,email,token) values($1,$2,$3,$4) returning user_id",
		ByteSliceToString(firstName),
		ByteSliceToString(lastName),
		ByteSliceToString(email),
		token[:]).Scan(&userId);
	err!=nil{
		log.Println(err)
		ctx.Error("Unhandled error",404)
	}else{
		successfulAuth(ctx,strconv.Itoa(userId))
		ctx.Redirect("/secretpage",200)
	}
}

func CreateCookie(key string, value string, expire int) *fasthttp.Cookie {
	if strings.Compare(key, "") == 0 {
		key = "unhandled cookie"
	}
	authCookie := fasthttp.Cookie{}
	authCookie.SetKey(key)
	authCookie.SetValue(value)
	authCookie.SetMaxAge(expire)
	authCookie.SetHTTPOnly(true)
	authCookie.SetSameSite(fasthttp.CookieSameSiteLaxMode)
	return &authCookie
}

func successfulAuth(ctx *fasthttp.RequestCtx, userId string){
	var access_token string
	for {
		access_token = Hasher(128)
		if _,err:= rdb.Get(access_token).Result();err!=nil{
			break
		}
	}

	accessTokenCookie :=CreateCookie("accessToken",access_token,36000000000)
	idCookie :=CreateCookie("userId",userId,36000000000)
	rdb.Set(userId,access_token,360000000000)
	ctx.Response.Header.SetCookie(accessTokenCookie)
	ctx.Response.Header.SetCookie(idCookie)
}

func ByteSliceToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func StringToByteSlice(str string) []byte {
	hdr := *(*reflect.StringHeader)(unsafe.Pointer(&str))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: hdr.Data,
		Len:  hdr.Len,
		Cap:  hdr.Len,
	}))
}

func CORSHandler (h fasthttp.RequestHandler) fasthttp.RequestHandler{
	return cors.AllowAll().Handler(h)
}