package dao

import (
	"github.com/xiaoriri-team/gin-skeleton/internal/model"
	"strings"
)

type JuhePhoneCaptchaRsp struct {
	ErrorCode int    `json:"error_code"`
	Reason    string `json:"reason"`
}

// 根据用户ID获取用户
func (d *Dao) GetUserByID(id int64) (*model.User, error) {
	user := &model.User{
		Model: &model.Model{
			ID: id,
		},
	}

	return user.Get(d.engine)
}

// 根据用户名获取用户
func (d *Dao) GetUserByUsername(username string) (*model.User, error) {
	user := &model.User{
		Username: username,
	}

	return user.Get(d.engine)
}

// 根据手机号获取用户
func (d *Dao) GetUserByPhone(phone string) (*model.User, error) {
	user := &model.User{
		Phone: phone,
	}

	return user.Get(d.engine)
}

// 根据IDs获取用户列表
func (d *Dao) GetUsersByIDs(ids []int64) ([]*model.User, error) {
	user := &model.User{}

	return user.List(d.engine, &model.ConditionsT{
		"id IN ?": ids,
	}, 0, 0)
}

// 根据关键词模糊获取用户列表
func (d *Dao) GetUsersByKeyword(keyword string) ([]*model.User, error) {
	user := &model.User{}

	if strings.Trim(keyword, "") == "" {
		return user.List(d.engine, &model.ConditionsT{
			"ORDER": "id ASC",
		}, 0, 6)
	} else {

		return user.List(d.engine, &model.ConditionsT{
			"username LIKE ?": strings.Trim(keyword, "") + "%",
		}, 0, 6)
	}
}

// 创建用户
func (d *Dao) CreateUser(user *model.User) (*model.User, error) {
	return user.Create(d.engine)
}

// 更新用户
func (d *Dao) UpdateUser(user *model.User) error {
	return user.Update(d.engine)
}
