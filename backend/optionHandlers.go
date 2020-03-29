package main

import (
	"coursework/functools"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
	"time"
)

type BasicInfoStruct struct{
	Sex      string `json:"sex" validate:"sex"`
	Status   uint `json:"status" validate:"required,min=0,max=5"`
	Birthday string `json:"birthday"`
	Tel      uint `json:"tel"`
	Country  string `json:"country"`
	City     string `json:"city"`
}

func UpdateBasicInfoTextHandler(ctx *fasthttp.RequestCtx){
	obj := &BasicInfoStruct{}

	if err := json.Unmarshal(ctx.PostBody(), obj); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

	if err := Validator.Struct(obj); err!=nil{
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

}

func UpdateProfileAvatar(ctx *fasthttp.RequestCtx){
	f, err :=ctx.FormFile("photo")
	if err != nil{
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	err = fasthttp.SaveMultipartFile(
		f,
		"./profile_avatars/" + functools.RandomStringGenerator(16) + strconv.FormatInt(time.Now().UnixNano(),10) + ".jpg")
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
	err = fasthttp.SaveMultipartFile(f, "./profile_bgs/" + strconv.FormatInt(time.Now().UnixNano(),10) + ".jpg")
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