package service

import (
	"dilu/modules/notice/models"
	"dilu/modules/notice/service/dto"

	"github.com/baowk/dilu-core/core/base"
)

type TaskService struct {
	*base.BaseService
}

var SerTask = TaskService{
	base.NewService("notice"),
}

func (s *TaskService) UserTasks(req *dto.TaskGetPageReq, list *[]models.Task, total *int64, unReadCnt *int64) error {
	db := s.DB().Where("team_id = ?", req.TeamId).Where("user_id = ?", req.UserId)
	if req.Status != 0 {
		db.Where("status = ?", req.Status)
	}
	if err := db.Order("status asc ,created_at desc").Offset(req.GetOffset()).Limit(req.GetSize()).
		Find(list).Offset(-1).Limit(-1).Count(total).Error; err != nil {
		return err
	}

	if err := db.Where("reminder_status = ?", 1).Count(unReadCnt).Error; err != nil {
		return err
	}

	return nil
}
