package services

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"rustdesk-api-server/app/models"
	"rustdesk-api-server/utils"
	"rustdesk-api-server/utils/gmd5"
	"time"
)

var Login = new(LoginService)

type LoginService struct {
}

func (s *LoginService) UserLogin(username, password, clientId, uuid string, ctx *context.Context) (token string, err error) {

	// Query whether the user exists
	var user models.User
	err = orm.NewOrm().QueryTable(new(models.User)).
		Filter("username", username).One(&user)
	if err != nil {
		return "", errors.New("Incorrect username or password")
	}

	// 生成密码
	pwd := User.GenPwd(password)
	// 检测密码是否正确
	if user.Password != pwd {
		return "", errors.New("Incorrect username or password")
	}
	// Determine if a user is disabled
	if user.Status != 1 {
		return "", errors.New("current user is disabled")
	}

	m := orm.NewOrm()
	entity := models.User{Id: user.Id}
	entity.LastLoginTime = time.Now().Unix()
	entity.LastLoginIp = ctx.Input.IP()
	entity.UpdateTime = time.Now().Unix()
	m.Update(&entity, "LastLoginTime", "LastLoginIp", "UpdateTime")

	// 生成登录token
	token2 := gmd5.EncryptNE(user.Password + clientId + uuid)

	// 返回jwt
	token, _ = utils.GenerateJwtToken(int(user.Id), user.Username, token2, clientId, uuid)

	// 保存当前电脑登录信息
	Token.Login(&user, clientId, uuid, token2)
	return token, nil
}
