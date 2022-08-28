package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.POST("/abc", callback)
	e.Logger.Fatal(e.Start(":9487"))
}

func callback(c echo.Context) error {
	// 驗證是否有登陸，分別轉入 login 或者 first
	fmt.Println("hello world")
	return c.String(http.StatusOK, "callback")
}

type CallbackReq struct {
	CallbackCommand string          `json:"CallbackCommand"`  // 回调命令
	OperatorAccount string          `json:"Operator_Account"` //发起创建群组请求的操作者 UserID
	OwnerAccount    string          `json:"Owner_Account"`    //请求创建的群的群主 UserID
	Type            string          `json:"Type"`             //	产生群消息的 群组类型介绍，例如 Public
	Name            string          `json:"Name"`             // 请求创建的群组的名称
	CreateGroupNum  int             `json:"CreateGroupNum"`   // 该用户已创建的同类的群组个数
	MemberList      []MemberAccount // 请求的成员列表
}

type MemberAccount struct {
	MemberAccount string `json:"Member_Account"` // 群成员 UserID
}

type CallbackRes struct {
	ActionStatus string `json:"ActionStatus"` // 回调命令执行结果，OK 表示执行成功，FAIL 表示执行失败
	ErrorInfo    string `json:"ErrorInfo"`    // 错误信息
	ErrorCode    int    `json:"ErrorCode"`    // 错误码
}
