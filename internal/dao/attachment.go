package dao

import "github.com/xiaoriri-team/gin-skeleton/internal/model"

func (d *Dao) CreateAttachment(attachment *model.Attachment) (*model.Attachment, error) {
	return attachment.Create(d.engine)
}
