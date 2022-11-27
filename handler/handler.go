package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gglzc/stressingtest/pkg"
)



func AutoRegister(){
	var  account string =pkg.RandString(5) 
	var	 password string =pkg.RandString(6)
	User[account]=password
	//存入全域變數
	values :=map[string]string{"account": account ,"password": password}
	//request 把資料成json type
	jsonValue, _ := json.Marshal(values)
	
	req,err := http.Post(RegisterAPI,"application/json",bytes.NewBuffer(jsonValue))
	if err!=nil{
		return 
	}
	//prevent memory leak
	 req.Body.Close()
	
	//get login and start bet 
	AutoBet(AutoLogin(account,User[account]))
	
}


func AutoLogin(acconut string,pwd string)string{
	var account  string
	var password string

	var token  	string
	values :=map[string]string{"account": account ,"password": password}
	jsonValue, _ := json.Marshal(values)
	//request登入
	resp,err:= http.Post(LoginAPI,"application/json",bytes.NewBuffer(jsonValue))
	if err!=nil{
		return err.Error()
	}
	defer resp.Body.Close()
	//response 可以拿到token 
	json.NewDecoder(resp.Body).Decode(token)
	//token儲存
	Token[token]=acconut
	
	return token
}

func AutoBet(token string){
	var win int
	var bet int
	//隨機5-35的bet
	bet=pkg.RandNumber()
	win=bet
}