package model

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"strconv"
	"time"
)

type (
	// Handler 回话结构体
	Handler struct {
		DB *mgo.Session
	}
)

const (
	Key = "secret"
)

func (h *Handler) Signup(c echo.Context) (err error) {
	// 新建和绑定
	user := &User{
		ID: bson.NewObjectId(),
	}
	err = c.Bind(user)
	if err != nil {
		return
	}

	// 校验信息：邮箱和输入的密码
	if user.Email == "" || user.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	// 保存
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("twitter").C("users").Insert(user); err != nil {
		return
	}

	// 响应json
	return c.JSON(http.StatusCreated, user)
}

func (h *Handler) Login(c echo.Context) (err error) {
	// 绑定
	u := new(User)
	err = c.Bind(u)
	if err != nil {
		return err
	}

	//查询
	clone := h.DB.Clone()
	defer clone.Close()
	if err = clone.DB("twitter").C("users").Find(
		bson.M{"email": u.Email, "password": u.Password}).One(u); err != nil {
		// 判断错误
		if errors.Is(err, mgo.ErrNotFound) {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
		}
		return
	}

	//生成token jwt
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// 配合token秘钥转为token字符串
	u.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		return err
	}

	// 置空密码
	u.Password = ""

	// 响应
	return c.JSON(http.StatusOK, u)
}

func userIDFromToken(c echo.Context) string {
	// 上下文中获取key为user的token
	token := c.Get("user").(*jwt.Token)
	// 获取token中的claims
	claims := token.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}

// Follow  关注某人
func (h *Handler) Follow(c echo.Context) (err error) {
	thisId := userIDFromToken(c)
	fId := c.Param("id")

	clone := h.DB.Clone()
	defer clone.Close()

	// 更新
	if err = clone.DB("twitter").C("users").UpdateId(
		bson.ObjectIdHex(fId),
		bson.M{"$addToSet": bson.M{"followers": thisId}}); err != nil {
		if errors.Is(err, mgo.ErrNotFound) {
			return echo.ErrNotFound
		}
	}

	return
}

// CreatePost

func (h *Handler) CreatePost(c echo.Context) (err error) {
	// 新建一个from
	user := &User{ID: bson.ObjectIdHex(userIDFromToken(c))}
	// 新建一个Post
	post := &Post{
		ID:   bson.NewObjectId(),
		From: user.ID.Hex(),
	}
	// post和上下文bind
	if err = c.Bind(post); err != nil {
		return
	}

	// 发消息前校验
	if post.To == "" || post.Message == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
			//Internal: nil,
		}
	}

	// 获取连接 查询from用户是否存在
	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("twitter").C("users").FindId(user.ID).One(user); err != nil {
		if errors.Is(err, mgo.ErrNotFound) {
			return echo.ErrNotFound
		}
		return
	}

	// 存在执行保存发送消息
	if err = db.DB("twitter").C("posts").Insert(post); err != nil {
		return
	}

	// 响应
	return c.JSON(http.StatusCreated, post)
}

// FetchPost 查询分页历史消息
func (h *Handler) FetchPost(c echo.Context) (err error) {
	// 获取token中的用户ID
	uid := userIDFromToken(c)
	// 获取page和limit
	page, _ := strconv.Atoi(c.Param("page"))
	limit, _ := strconv.Atoi(c.Param("limit"))

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}

	// 结果集数组
	var posts []*Post
	db := h.DB.Clone()
	if err = db.DB("twitter").C("posts").
		Find(bson.M{"to": uid}).
		Skip((page - 1) * limit).
		Limit(limit).
		All(&posts); err != nil {
		return
	}

	defer db.Close()

	return c.JSON(200, posts)

}

// Main 主函数
func Main() {
	// 创建一个新的Echo实例
	e := echo.New()

	// 中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// 连接到MongoDB
	session, err := mgo.Dial("mongodb://210.22.22.150:1055")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer session.Close()

	// 创建处理程序实例
	handler := &Handler{DB: session}

	// 路由
	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)
	e.POST("/follow/:id", handler.Follow, middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(Key),
	}))
	e.POST("/post", handler.CreatePost, middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(Key),
	}))
	e.GET("/posts/:page/:limit", handler.FetchPost, middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(Key),
	}))

	// 启动服务器
	e.Logger.Fatal(e.Start(":9999"))
}
