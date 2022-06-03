package service

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gofrs/uuid"
	"github.com/xiaoriri-team/gin-skeleton/global"
	"github.com/xiaoriri-team/gin-skeleton/internal/model"
	"github.com/xiaoriri-team/gin-skeleton/pkg/convert"
	"github.com/xiaoriri-team/gin-skeleton/pkg/errcode"
	"github.com/xiaoriri-team/gin-skeleton/pkg/util"
)

const MAX_CAPTCHA_TIMES = 2

type PhoneCaptchaReq struct {
	Phone        string `json:"phone" form:"phone" binding:"required"`
	ImgCaptcha   string `json:"img_captcha" form:"img_captcha" binding:"required"`
	ImgCaptchaID string `json:"img_captcha_id" form:"img_captcha_id" binding:"required"`
}

type UserPhoneBindReq struct {
	Phone   string `json:"phone" form:"phone" binding:"required"`
	Captcha string `json:"captcha" form:"captcha" binding:"required"`
}

type AuthRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
type RegisterRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type ChangePasswordReq struct {
	Password    string `json:"password" form:"password" binding:"required"`
	OldPassword string `json:"old_password" form:"old_password" binding:"required"`
}
type ChangeNicknameReq struct {
	Nickname string `json:"nickname" form:"nickname" binding:"required"`
}
type ChangeAvatarReq struct {
	Avatar string `json:"avatar" form:"avatar" binding:"required"`
}

const LOGIN_ERR_KEY = "PaoPaoUserLoginErr"
const MAX_LOGIN_ERR_TIMES = 10

// 用户认证
func (svc *Service) DoLogin(param *AuthRequest) (*model.User, error) {
	user, err := svc.dao.GetUserByUsername(param.Username)
	if err != nil {
		return nil, errcode.UnauthorizedAuthNotExist
	}

	if user.Model != nil && user.ID > 0 {
		if errTimes, err := global.Redis.Get(svc.ctx, fmt.Sprintf("%s:%d", LOGIN_ERR_KEY, user.ID)).Result(); err == nil {
			if convert.StrTo(errTimes).MustInt() >= MAX_LOGIN_ERR_TIMES {
				return nil, errcode.TooManyLoginError
			}
		}

		// 对比密码是否正确
		if svc.ValidPassword(user.Password, param.Password, user.Salt) {

			if user.Status == model.UserStatusClosed {
				return nil, errcode.UserHasBeenBanned
			}

			// 清空登录计数
			global.Redis.Del(svc.ctx, fmt.Sprintf("%s:%d", LOGIN_ERR_KEY, user.ID))
			return user, nil
		}

		// 登录错误计数
		_, err = global.Redis.Incr(svc.ctx, fmt.Sprintf("%s:%d", LOGIN_ERR_KEY, user.ID)).Result()
		if err == nil {
			global.Redis.Expire(svc.ctx, fmt.Sprintf("%s:%d", LOGIN_ERR_KEY, user.ID), time.Hour).Result()
		}

		return nil, errcode.UnauthorizedAuthFailed
	}

	return nil, errcode.UnauthorizedAuthNotExist
}

// 检查密码是否一致
func (svc *Service) ValidPassword(dbPassword, password, salt string) bool {
	return strings.Compare(dbPassword, util.EncodeMD5(util.EncodeMD5(password)+salt)) == 0
}

// 检测用户权限
func (svc *Service) CheckStatus(user *model.User) bool {
	return user.Status == model.UserStatusNormal
}

// 验证用户
func (svc *Service) ValidUsername(username string) error {
	// 检测用户是否合规
	if utf8.RuneCountInString(username) < 3 || utf8.RuneCountInString(username) > 12 {
		return errcode.UsernameLengthLimit
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(username) {
		return errcode.UsernameCharLimit
	}

	// 重复检查
	user, _ := svc.dao.GetUserByUsername(username)

	if user.Model != nil && user.ID > 0 {
		return errcode.UsernameHasExisted
	}

	return nil
}

// 密码检查
func (svc *Service) CheckPassword(password string) error {
	// 检测用户是否合规
	if utf8.RuneCountInString(password) < 6 || utf8.RuneCountInString(password) > 16 {
		return errcode.PasswordLengthLimit
	}

	return nil
}

// 检测手机号是否存在
func (svc *Service) CheckPhoneExist(uid int64, phone string) bool {
	u, err := svc.dao.GetUserByPhone(phone)
	if err != nil {
		return false
	}

	if u.Model == nil || u.ID == 0 {
		return false
	}

	if u.ID == uid {
		return false
	}

	return true
}

// 密码加密&生成salt
func (svc *Service) EncryptPasswordAndSalt(password string) (string, string) {
	salt := uuid.Must(uuid.NewV4()).String()[:8]
	password = util.EncodeMD5(util.EncodeMD5(password) + salt)

	return password, salt
}

// 用户注册
func (svc *Service) Register(username, password string) (*model.User, error) {
	password, salt := svc.EncryptPasswordAndSalt(password)

	user := &model.User{
		Nickname: username,
		Username: username,
		Password: password,
		Avatar:   svc.GetRandomAvatar(),
		Salt:     salt,
		Status:   model.UserStatusNormal,
	}

	user, err := svc.dao.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// 获取用户信息
func (svc *Service) GetUserInfo(param *AuthRequest) (*model.User, error) {
	user, err := svc.dao.GetUserByUsername(param.Username)

	if err != nil {
		return nil, err
	}

	if user.Model != nil && user.ID > 0 {
		return user, nil
	}

	return nil, errcode.UnauthorizedAuthNotExist
}

func (svc *Service) GetUserByUsername(username string) (*model.User, error) {
	user, err := svc.dao.GetUserByUsername(username)

	if err != nil {
		return nil, err
	}

	if user.Model != nil && user.ID > 0 {
		return user, nil
	}

	return nil, errcode.NoExistUsername
}

// 更新用户信息
func (svc *Service) UpdateUserInfo(user *model.User) error {
	return svc.dao.UpdateUser(user)
}

// 根据关键词获取用户推荐
func (svc *Service) GetSuggestUsers(keyword string) ([]string, error) {
	users, err := svc.dao.GetUsersByKeyword(keyword)
	if err != nil {
		return nil, err
	}

	usernames := []string{}
	for _, user := range users {
		usernames = append(usernames, user.Username)
	}

	return usernames, nil
}
