package service

import "github.com/xiaoriri-team/gin-skeleton/internal/model"

func (svc *Service) CreateAttachment(attachment *model.Attachment) (*model.Attachment, error) {
	return svc.dao.CreateAttachment(attachment)
}
