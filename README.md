# ZHANST

### 一个web api框架 使用 Trie树作为路由存储

```golang
import (
    "zhanst"
)

type Message struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
}

func GetMessage(c *zhanst.Context) {
    if id, ok := c.Params["id"]; !ok {
        c.JSON(500, zhanst.Res{
            "code":500,
            "msg":"id not exist",
            })
        return
    } else {
        c.JSON(200, zhanst.Res{
            "code":200, 
            "id":id,
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
```
