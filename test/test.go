package main

import (
	"fmt"
	"zhanst"
)

type Message struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func GetMessage(c *zhanst.Context) {
	if id, ok := c.Params["id"]; !ok {
		c.JSON(500, zhanst.Res{
			"code": 500,
			"msg":  "id not exist",
		})
		return
	} else {
		c.JSON(200, zhanst.Res{
			"code": 200,
			"id":   id,
		})
		return
	}
}

func PostMessage(c *zhanst.Context) {
	msg := &Message{}
	if err := c.Bind(msg); err != nil {
		c.JSON(500, zhanst.Res{
			"code":500,
			"msg": err.Error(),
		})
		return
	} else {
		c.JSON(200, zhanst.Res{
			"code":200,
			"msg":msg,
		})
		return
	}
}

func test1(c *zhanst.Context) {
	fmt.Println(c.Params)
	fmt.Println("test1")
}

func test2(c *zhanst.Context) {
	fmt.Println(c.Params)
	fmt.Println("test2")
}

func test3(c *zhanst.Context) {
	fmt.Println(c.Params)
	fmt.Println("test3")
}

func test4(c *zhanst.Context) {
	c.Next()
	fmt.Println(c.Params)
	fmt.Println("test4")

}

func test5(c *zhanst.Context) {
	msg := &Message{}
	err := c.Bind(msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c.Params)
	fmt.Println("test5")
	c.JSON(200, map[string]interface{}{
		"code":   200,
		"params": c.Params,
		"msg":    msg,
	})
}

func test6(c *zhanst.Context) {
	fmt.Println(c.Params)
	fmt.Println("test6")
}

func test7(c *zhanst.Context) {
	fmt.Println(c.Params)
	fmt.Println("test7")
}

func BeforeRequest(c *zhanst.Context) {
	fmt.Println("before request")
	c.Next()
}

func AfterRequest(c *zhanst.Context) {
	c.Next()
	fmt.Println("before request")
}

func main() {
	engine := zhanst.Default()
	group := engine.Group("/api")
	group.Use(BeforeRequest)
	group.Use(AfterRequest)
	group.GET("/messages/:id", GetMessage)
	group.POST("/messages", PostMessage)

	engine.Run("0.0.0.0:8084")

}
