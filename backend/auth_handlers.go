package main

import (
	"bytes"
	"context"
	"coursework/functools"
	"crypto/sha512"
	"github.com/jackc/pgx/v4"
	"github.com/valyala/fasthttp"
	"log"
	"strconv"
	"strings"
)

func authPageHandler(ctx *fasthttp.RequestCtx){
	ctx.SendFile("frontend/html/auth.html")
}


func AuthMiddleware(next fasthttp.RequestHandler)fasthttp.RequestHandler{
	return func(ctx *fasthttp.RequestCtx){
		redisAccessToken, err := Redis.Get(functools.ByteSliceToString(ctx.Request.Header.Cookie("userId"))).Result()
		if err == nil && bytes.Compare(ctx.Request.Header.Cookie("accessToken"),
			functools.StringToByteSlice(redisAccessToken)) == 0{
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
		Email:    functools.ByteSliceToString(ctx.FormValue("email")),
		Password: functools.ByteSliceToString(ctx.FormValue("password")),
	}


	err:= Validator.Struct(ls)

	if err!=nil{
		ctx.Error("not validated", 403)
		return
	}

	var dbToken []byte
	var userId int
	var firstName string
	var lastName string

	Postgres.Conn.QueryRow(context.Background(),"select user_id, first_name, last_name, token from registration where email = $1 limit 1", ls.Email).Scan(
		&userId,
		&firstName,
		&lastName,
		&dbToken)

	if userToken:=sha512.Sum512(append(functools.StringToByteSlice(ls.Password), Salt...)); bytes.Compare(dbToken, userToken[:]) != 0{
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

	if err:= Postgres.Conn.QueryRow(context.Background(),"select 1 from registration where email = $1 limit 1", email).Scan();err != pgx.ErrNoRows{
		ctx.Error("User already exists",402)
	}

	token:=sha512.Sum512(append(password, Salt...))
	var userId int
	if err := Postgres.Conn.QueryRow(context.Background(), "insert into registration (first_name,last_name,email,token) values($1,$2,$3,$4) returning user_id",
		functools.ByteSliceToString(firstName),
		functools.ByteSliceToString(lastName),
		functools.ByteSliceToString(email),
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
	var accessToken string
	for {
		accessToken = functools.RandomStringGenerator(128)
		if _,err:= Redis.Get(accessToken).Result();err!=nil{
			break
		}
	}

	accessTokenCookie :=CreateCookie("accessToken", accessToken,36000000000)
	idCookie :=CreateCookie("userId",userId,36000000000)
	Redis.Set(userId, accessToken,360000000000)
	ctx.Response.Header.SetCookie(accessTokenCookie)
	ctx.Response.Header.SetCookie(idCookie)
}