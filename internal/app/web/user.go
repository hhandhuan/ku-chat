package web

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/o1egl/govatar"
	"gorm.io/gorm"
	ew "ku-chat/internal/entity/web"
	"ku-chat/internal/model"
	"ku-chat/internal/service"
	"ku-chat/pkg/config"
	"ku-chat/pkg/utils/encrypt"
	"log"
	"net/http"
	"os"
	"time"
)

var User = cUser{}

type cUser struct{}

// Register 用户注册
func (c *cUser) Register(ctx *gin.Context) {
	s := service.Context(ctx)
	if ctx.Request.Method == http.MethodGet {
		s.View("register", nil)
		return
	}

	var req ew.RegisterReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}
	if err := g.Validator().Data(req).Run(context.Background()); err != nil {
		s.Back().WithError(err.FirstError()).Redirect()
		return
	}

	var user *model.Users
	err := model.User().M.Where("name = ?", req.Name).Find(&user).Error
	if err != nil {
		s.Back().WithError(err).Redirect()
		return
	}
	if user.ID > 0 {
		s.Back().WithError("用户名已被注册，请更换用户名继续尝试").Redirect()
		return
	}

	avatar, err := c.genAvatar(req.Name)
	if err != nil {
		s.Back().WithError("用户默认头像生成失败").Redirect()
		return
	}

	res := model.User().M.Create(&model.Users{
		Name:     req.Name,
		Avatar:   avatar,
		Password: encrypt.GenerateFromPassword(req.Password),
	})
	if res.Error != nil || res.RowsAffected <= 0 {
		s.Back().WithError("用户注册失败，请稍后在试").Redirect()
	} else {
		s.To("/login").WithMsg("注册成功，请继续登录").Redirect()
	}
}

func (*cUser) genAvatar(name string) (string, error) {
	path := fmt.Sprintf("%s/users/", config.Conf.Upload.Path)

	// 检查目录是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, os.ModePerm)
		_ = os.Chmod(path, os.ModePerm)
	}

	avatarName := encrypt.Md5(gconv.String(time.Now().UnixMicro()))
	avatarPath := fmt.Sprintf("users/%s.png", avatarName)
	uploadPath := fmt.Sprintf("%s/%s", config.Conf.Upload.Path, avatarPath)

	if err := govatar.GenerateFileForUsername(1, name, uploadPath); err != nil {
		log.Println(err)
		return "", err
	} else {
		return "/assets/upload/" + avatarPath, nil
	}
}

// Login 用户登录
func (*cUser) Login(ctx *gin.Context) {
	s := service.Context(ctx)
	if ctx.Request.Method == http.MethodGet {
		s.View("login", nil)
		return
	}

	var req ew.LoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	var user model.Users
	err := model.User().M.Where("name = ?", req.Name).Find(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.Back().WithError(err).Redirect()
		return
	}
	if user.ID <= 0 || !encrypt.CompareHashAndPassword(user.Password, req.Password) {
		s.Back().WithError("用户名或密码错误").Redirect()
	} else {
		s.SetAuth(user)
		s.To("/").WithMsg("登录成功").Redirect()
	}
}

// Logout 退出登录
func (*cUser) Logout(ctx *gin.Context) {
	s := service.Context(ctx)

	s.Forget()

	s.To("/login").WithMsg("退出成功").Redirect()
}

// Search 用户搜索
func (*cUser) Search(ctx *gin.Context) {
	s := service.Context(ctx)
	k := ctx.Query("user")
	if len(k) <= 0 {
		s.Json(gin.H{"code": 1, "msg": "请输入关键词"})
		return
	}

	var user *model.Users
	res := model.User().M.Where("name like ?", fmt.Sprintf("%%%s%%", k)).Or("id = ?", k).Find(&user)
	if res.Error != nil {
		s.Json(gin.H{"code": 1, "msg": res.Error})
	} else {
		s.Json(gin.H{"code": 0, "msg": "ok", "data": user})
	}
}
