package service

import (
	"dilu/modules/notice/models"
	"dilu/modules/notice/service/dto"
	"fmt"
	"time"

	"github.com/baowk/dilu-core/core/base"
	"github.com/jinzhu/copier"
)

type UserNoticeService struct {
	*base.BaseService
}

var SerUserNotice = UserNoticeService{
	base.NewService("notice"),
}

func (s *UserNoticeService) UserNotices(req *dto.UserNoticeGetPageReq, list *[]models.UserNotice, total *int64, unReadCnt *int64) error {
	var pns []models.PubNotice
	if err := s.DB().Where("expired > ?", time.Now()).Where("status = ?", 1).
		Where(fmt.Sprintf("team_id in (0,%d)", req.TeamId)).Limit(100).Order("id desc").Find(&pns).Error; err != nil {
		return err
	}
	var nu models.UserNotice
	s.DB().Where("team_id = ?", req.TeamId).Where("user_id = ?", req.UserId).Where("pub_id > 0").
		Order("pub_id desc").Limit(1).First(&nu)

	var is []models.UserNotice
	for _, v := range pns {
		if v.Id <= nu.PubId {
			break
		}
		var nuT models.UserNotice
		copier.Copy(&nuT, v)
		nuT.Id = 0
		nuT.PubId = v.Id
		nuT.Status = 1
		nuT.TeamId = req.TeamId
		nuT.UserId = req.UserId
		is = append(is, nuT)
	}
	if len(is) > 0 {
		s.DB().Create(is)
	}
	db := s.DB().Where("team_id = ?", req.TeamId).Where("user_id = ?", req.UserId)
	if req.Status != 0 {
		db.Where("status = ?", req.Status)
	}
	if err := db.Order("status asc ,created_at desc").Offset(req.GetOffset()).Limit(req.GetSize()).
		Find(list).Offset(-1).Limit(-1).Count(total).Error; err != nil {
		return err
	}
	if req.Status == 0 {
		if err := db.Where("status = ?", 1).Count(unReadCnt).Error; err != nil {
			return err
		}
	} else {
		*unReadCnt = *total
	}
	return nil
}

func (s *UserNoticeService) ReadUserNotice(req *dto.ReadNoticeDto, reqId string, teamId int, userId int) error {
	var nu models.UserNotice
	if err := s.DB().First(&nu, req.Id).Error; err != nil {
		return err
	}
	if nu.UserId != userId {
		return fmt.Errorf("无权限")
	}
	if nu.Status == 1 {
		nu.Status = 2
		nu.UpdatedAt = time.Now()
		if err := s.DB().Updates(&nu).Error; err != nil {
			return err
		}
	}
	return nil
}
