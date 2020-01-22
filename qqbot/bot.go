package qqbot

//QQ user info
type User struct{
	//private info
	Id 			int64		`json:"user_id"`
	NickName	string		`json:"nickname"`
	Age 		int64		`json:"age"`
	Sex 		string 		`json:"sex"`

	//group member info
	Area		string 		`json:"area"`
	Card		string 		`json:"card"`
	Level		string		`json:"level"`
	Role 		string 		`json:"role"`
	Title		string 		`json:"title"`

	//anonymous user info in group
	AnonymousId	int64		`json:"anonymousId"`




}

//Message data from cqhttp
type CQMsg struct{




}