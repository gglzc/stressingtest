package handler

import (
	
)
const RegisterAPI = "localhost:8080/bet/register"
const LoginAPI = "localhost:8080/bet/login"
const BetAPI = "localhost:8080/bet/result"
//記User
var User map[string]string=make(map[string]string)
//記token  key:token value:account
var Token map[string]string=make(map[string]string)

type UserRequest struct{
	Account   string  `json:"account"`
	Password  string  `json:"password"`
}
type GameRequest struct{
	Bet 	string 	  `json:"bet"`
}