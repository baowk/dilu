package service

import (
	"dilu/common/codes"
	"dilu/common/consts"
	"dilu/common/utils"
	"dilu/modules/dental/enums"
	"dilu/modules/dental/models"
	"dilu/modules/dental/service/dto"
	smodels "dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/baowk/dilu-core/core/errs"
	"github.com/jinzhu/copier"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type BillService struct {
	*base.BaseService
}

var SerBill = BillService{
	base.NewService(consts.DB_CRM),
}

func (s *BillService) Page(teamId int, req dto.BillGetPageReq, list *[]dto.BillDto, total *int64) error {
	db := s.DB().Offset(req.GetOffset()).Limit(req.GetSize()).Order("id desc")
	if teamId > 0 {
		db.Where("team_id = ?", teamId)
	}
	if req.TradeType != 0 {
		db.Where("trade_type = ?", req.TradeType)
	}
	if !req.Begin.IsZero() {
		db.Where("trade_at > ?", req.Begin)
	}
	if !req.End.IsZero() {
		db.Where("trade_at < ?", req.End)
	}
	if req.UserId > 0 {
		db.Where("user_id = ?", req.UserId)
	}
	var ds []models.Bill
	db.Find(&ds).Offset(-1).Limit(-1).Count(total)
	var cids []int
	for _, b := range ds {
		cids = append(cids, b.CustomerId)

	}
	var cs []models.Customer
	if len(cids) > 0 {
		if err := SerCustomer.GetByIds(teamId, cids, &cs); err != nil {
			return err
		}
	}
	for _, b := range ds {
		for _, c := range cs {
			if c.Id == b.CustomerId {
				var bt dto.BillDto
				copier.Copy(&bt, b)
				bt.CustomerName = c.Name
				*list = append(*list, bt)
				break
			}
		}
	}
	return nil
}

func (s *BillService) CreateBill(reqId string, bill dto.IdentifyBillDto, dbill *models.Bill, createBy int) errs.IError {
	if bill.TeamId < 1 {
		return codes.ErrInvalidParameter(reqId, "teamId is nil")
	}
	if bill.UserId < 1 {
		return codes.ErrInvalidParameter(reqId, "userId is nil")
	}
	var team smodels.SysTeam
	if err := service.SerSysTeam.Get(bill.TeamId, &team); err != nil {
		return codes.ErrNotFound(strconv.Itoa(bill.TeamId), "team", reqId, err)
	}
	var teamM smodels.SysMember
	if err := service.SerSysMember.GetMember(bill.TeamId, bill.UserId, &teamM); err != nil {
		return codes.ErrNotFound(fmt.Sprintf("%d-%d", bill.TeamId, bill.UserId), "teamMember", reqId, err)
	}
	if bill.CustomerId < 1 {
		var customers []models.Customer
		if err := SerCustomer.GetByUserIdAndName(bill.UserId, 0, bill.CustomerName, &customers); err != nil {
			core.Log.Error("获取客户错误", zap.Error(err))
		}
		if len(customers) > 0 {
			bill.CustomerId = customers[0].Id
		} else {
			customer := models.Customer{
				Name:        bill.CustomerName,
				UserId:      bill.UserId,
				TeamId:      bill.TeamId,
				InviterName: bill.InviterName,
				Inviter:     bill.Inviter,
				CreateBy:    createBy,
				DeptPath:    teamM.DeptPath,
			}
			if err := SerCustomer.Create(&customer); err != nil {
				return codes.ErrSys(err)
			}
			bill.CustomerId = customer.Id
		}
	} else {
		var customer models.Customer
		if err := SerCustomer.Get(bill.CustomerId, &customer); err != nil {
			return codes.ErrNotFound(strconv.Itoa(bill.CustomerId), "customer", reqId, err)
		}
	}
	if err := copier.Copy(dbill, bill); err != nil {
		return codes.ErrSys(err)
	}

	dbill.DeptPath = teamM.DeptPath

	if dbill.TradeType == int(enums.TradeDebt) {
		dbill.DebtAmount = dbill.PaidAmount
		dbill.PaidAmount = decimal.Zero
		dbill.RefundAmount = decimal.Zero
	} else if dbill.TradeType == int(enums.TradeRefund) {
		dbill.RefundAmount = dbill.PaidAmount
		dbill.PaidAmount = decimal.Zero
	}
	dbill.CreatedAt = time.Now()
	dbill.UpdatedAt = dbill.CreatedAt

	if bill.TradeAt != "" {
		if d, err := time.Parse("2006-01-02", bill.TradeAt); err != nil {
			dbill.TradeAt = dbill.CreatedAt
		} else {
			dbill.TradeAt = d
		}
	}

	if bill.ImplantedCount < 1 {
		dbill.Implant = 1
	} else {
		if bill.ImplantDate != "" {
			if d, err := time.Parse("2006-01-02", bill.ImplantDate); err == nil {
				dbill.ImplantDate = d
			}
		} else {
			dbill.ImplantDate = dbill.TradeAt
		}
		if bill.ImplantedCount < bill.DentalCount {
			dbill.Implant = 2
		} else {
			dbill.Implant = 3
		}
	}

	if bill.PaybackDate != "" {
		if d, err := time.Parse("2006-01-02", bill.PaybackDate); err == nil {
			dbill.PaybackDate = d
		}
	}

	dbill.No = strings.Replace(dbill.CreatedAt.Format("20060102150405.000000"), ".", "", -1)
	dbill.CreateBy = createBy

	if dbill.Pack == 0 {
		if dbill.TradeType == int(enums.TradeDeal) && dbill.DentalCount == 0 {
			dbill.Pack = int(enums.General)
		}
	}

	if err := s.Create(dbill); err != nil {
		return codes.ErrSys(err)
	}
	return nil
}

func (s *BillService) UpdateBill(reqId string, bill dto.IdentifyBillDto, dbill *models.Bill, createBy int) errs.IError {
	if bill.TeamId < 1 {
		return codes.ErrInvalidParameter(reqId, "teamId is nil")
	}
	if bill.UserId < 1 {
		return codes.ErrInvalidParameter(reqId, "userId is nil")
	}
	var team smodels.SysTeam
	if err := service.SerSysTeam.Get(bill.TeamId, &team); err != nil {
		return codes.ErrNotFound(strconv.Itoa(bill.TeamId), "team", reqId, err)
	}
	var teamM smodels.SysMember
	if err := service.SerSysMember.GetMember(bill.TeamId, bill.UserId, &teamM); err != nil {
		return codes.ErrNotFound(fmt.Sprintf("%d-%d", bill.TeamId, bill.UserId), "teamMember", reqId, err)
	}
	if bill.CustomerId < 1 {
		var customers []models.Customer
		if err := SerCustomer.GetByUserIdAndName(bill.UserId, 0, bill.CustomerName, &customers); err != nil {
			core.Log.Error("获取客户错误", zap.Error(err))
		}
		if len(customers) > 0 {
			bill.CustomerId = customers[0].Id
		} else {
			customer := models.Customer{
				Name:        bill.CustomerName,
				UserId:      bill.UserId,
				TeamId:      bill.TeamId,
				InviterName: bill.InviterName,
				Inviter:     bill.Inviter,
				CreateBy:    createBy,
				DeptPath:    teamM.DeptPath,
			}
			if err := SerCustomer.Create(&customer); err != nil {
				return codes.ErrSys(err)
			}
			bill.CustomerId = customer.Id
		}
	} else {
		var customer models.Customer
		if err := SerCustomer.Get(bill.CustomerId, &customer); err != nil {
			return codes.ErrNotFound(strconv.Itoa(bill.CustomerId), "customer", reqId, err)
		}
	}
	if err := copier.Copy(dbill, bill); err != nil {
		return codes.ErrSys(err)
	}

	dbill.DeptPath = teamM.DeptPath

	if dbill.TradeType == int(enums.TradeDebt) {
		dbill.DebtAmount = dbill.PaidAmount
		dbill.PaidAmount = decimal.Zero
		dbill.RefundAmount = decimal.Zero
	} else if dbill.TradeType == int(enums.TradeRefund) {
		dbill.RefundAmount = dbill.PaidAmount
		dbill.PaidAmount = decimal.Zero
	}
	dbill.UpdatedAt = time.Now()

	if bill.TradeAt != "" {
		if d, err := time.Parse("2006-01-02", bill.TradeAt); err != nil {
			dbill.TradeAt = dbill.CreatedAt
		} else {
			dbill.TradeAt = d
		}
	}

	if bill.ImplantedCount < 1 {
		dbill.Implant = 1
	} else {
		if bill.ImplantDate != "" {
			if d, err := time.Parse("2006-01-02", bill.ImplantDate); err == nil {
				dbill.ImplantDate = d
			}
		} else {
			dbill.ImplantDate = dbill.TradeAt
		}
		if bill.ImplantedCount < bill.DentalCount {
			dbill.Implant = 2
		} else {
			dbill.Implant = 3
		}
	}

	if bill.PaybackDate != "" {
		if d, err := time.Parse("2006-01-02", bill.PaybackDate); err == nil {
			dbill.PaybackDate = d
		}
	}

	//dbill.No = strings.Replace(dbill.CreatedAt.Format("20060102150405.000000"), ".", "", -1)
	dbill.UpdateBy = createBy

	if err := s.UpdateById(dbill); err != nil {
		return codes.ErrSys(err)
	}
	return nil
}

func (s *BillService) LinkBill(reqId string, bill dto.LinkBillDto) errs.IError {
	var old models.Bill
	if err := s.Get(bill.LinkId, &old); err != nil {
		return codes.ErrSys(err)
	}

	curBill := models.Bill{
		LinkId:         bill.LinkId,
		TeamId:         old.TeamId,
		CustomerId:     old.CustomerId,
		UserId:         old.UserId,
		DeptPath:       old.DeptPath,
		Amount:         decimal.Zero,
		TradeType:      bill.TradeType,
		DentalCount:    0,
		Brand:          old.Brand,
		ImplantedCount: old.ImplantedCount,
		Doctor:         old.Doctor,
		Pack:           old.Pack,
		PaybackDate:    old.PaybackDate,
		Tags:           old.Tags,
		PrjName:        old.PrjName,
		OtherPrj:       old.OtherPrj,
		Remark:         bill.Remark,
	}

	if bill.TradeType == int(enums.TradeDeal) {
		return errs.ErrWithCode(codes.InvalidParameter)
	} else if bill.TradeType == int(enums.TradeBalance) {
		d, err := decimal.NewFromString(bill.PaidAmount)
		if err != nil {
			return codes.ErrSys(err)
		}
		curBill.PaidAmount = d

		d, err = decimal.NewFromString(bill.RealAmount)
		if err != nil {
			return codes.ErrSys(err)
		}
		curBill.RealAmount = d
	} else if bill.TradeType == int(enums.TradeDebt) {
		d, err := decimal.NewFromString(bill.PaidAmount)
		if err != nil {
			return codes.ErrSys(err)
		}
		curBill.DebtAmount = d

		d2, err := decimal.NewFromString(bill.RealAmount)
		if err != nil {
			return codes.ErrSys(err)
		}
		if d2.Cmp(decimal.Zero) > 0 {
			curBill.RealAmount = d2
			curBill.DebtAmount = curBill.DebtAmount.Sub(d2)
			curBill.PaidAmount = d2
		}

	} else if bill.TradeType == int(enums.TradeRefund) {
		d, err := decimal.NewFromString(bill.PaidAmount)
		if err != nil {
			return codes.ErrSys(err)
		}
		curBill.RefundAmount = d
	}

	if bill.TradeAt != "" {
		if d, err := time.Parse("2006-01-02", bill.TradeAt); err != nil {
			curBill.TradeAt = curBill.CreatedAt
		} else {
			curBill.TradeAt = d
		}
	}

	if bill.ImplantedCount < 1 {
		curBill.Implant = old.Implant
	} else {
		if bill.ImplantDate != "" {
			if d, err := time.Parse("2006-01-02", bill.ImplantDate); err == nil {
				curBill.ImplantDate = d
			}
		} else {
			curBill.ImplantDate = curBill.TradeAt
		}
		if bill.ImplantedCount+old.ImplantedCount < old.DentalCount {
			curBill.Implant = 2
		} else {
			curBill.Implant = 3
		}
	}
	curBill.No = strings.Replace(curBill.CreatedAt.Format("20060102150405.000000"), ".", "", -1)
	if err := s.Create(curBill); err != nil {
		return codes.ErrSys(err)
	}
	return nil
}

func (s *BillService) Identify(req dto.BillTmplReq, bill *dto.IdentifyBillDto) errs.IError {
	(*bill).TeamId = req.TeamId
	arr := strings.Split(req.Text, "\n")
	var custName string
	for _, v := range arr {
		if strings.Trim(v, " ") == "" {
			continue
		}
		for _, key := range enums.Counselor {
			if strings.Contains(v, key) {
				(*bill).Name = getVal(v)
				break
			}
		}
		for _, key := range enums.TradeAt {
			if strings.Contains(v, key) {
				tmpD := getVal(v)
				if tmpD != "" {
					(*bill).TradeAt = getDate(tmpD)
				}
				break
			}
		}

		for _, key := range enums.CustomerName {
			if strings.Contains(v, key) {
				custName = getVal(v)
				break
			}
		}
		for _, key := range enums.Doctor {
			if strings.Contains(v, key) {
				(*bill).Doctor = getVal(v)
				break
			}
		}
		for _, key := range enums.Project {
			if strings.Contains(v, key) {
				(*bill).PrjName = getVal(v)
				break
			}
		}
		for _, key := range enums.Brand {
			if strings.Contains(v, key) {
				(*bill).BrandName = getVal(v)
				break
			}
		}
		for _, key := range enums.Cnt {
			if strings.Contains(v, key) {
				cntStr := getVal(v)
				if cnt, err := strconv.Atoi(cntStr); err == nil {
					(*bill).DentalCount = cnt
				}
				break
			}
		}
		for _, key := range enums.Others {
			if strings.Contains(v, key) {
				(*bill).OtherPrj = getVal(v)
				break
			}
		}
		for _, key := range enums.Total {
			if strings.Contains(v, key) {
				(*bill).RealAmount = getVal(v)
				break
			}
		}
		for _, key := range enums.Paid {
			if strings.Contains(v, key) {
				(*bill).PaidAmount = getVal(v)
				break
			}
		}
		for _, key := range enums.Debts {
			if strings.Contains(v, key) {
				(*bill).Debts = getVal(v)
				break
			}
		}
		for _, key := range enums.PaybackDate {
			if strings.Contains(v, key) {
				tmpD := getVal(v)
				if tmpD != "" {
					(*bill).PaybackDate = getDate(tmpD)
				}
				break
			}
		}
		for _, key := range enums.Implant {
			if strings.Contains(v, key) {
				imp := getVal(v)
				var flag bool
				for _, iv := range enums.ImplantVals {
					if imp == iv {
						(*bill).Implant = int(enums.ImplantFull)
						flag = true
						break
					}
				}
				if !flag {
					(*bill).Implant = int(enums.ImplantHalf)
				}
				break
			}
		}
		for _, key := range enums.Extensions {
			if strings.Contains(v, key) {
				(*bill).Extensions = getVal(v)
				break
			}
		}
		for _, key := range enums.Remark {
			if strings.Contains(v, key) {
				(*bill).Remark = getVal(v)
				break
			}
		}
	}
	if bill.PrjName == "半口" {
		bill.Pack = int(enums.PackHalf)
	} else if bill.PrjName == "全口" {
		bill.Pack = int(enums.PackFull)
	} else if bill.DentalCount > 0 {
		bill.Pack = int(enums.PackCnt)
	} else {
		bill.Pack = int(enums.General)
	}

	var members []smodels.SysMember
	if err := service.SerSysMember.GetMembers(req.TeamId, 0, "", bill.Name, &members); err != nil {
		core.Log.Error("获取咨询师错误", zap.Error(err))
		return errs.Err(codes.FAILURE, "", err)
	}
	if len(members) > 0 {
		(*bill).UserId = members[0].UserId
	}

	var customers []models.Customer
	if err := SerCustomer.GetByUserIdAndName(bill.UserId, 0, custName, &customers); err != nil {
		core.Log.Error("获取客户错误", zap.Error(err))
	}
	for _, c := range customers {
		op := dto.Option{
			Value: c.Id,
			Label: c.Name,
		}
		(*bill).Customers = append((*bill).Customers, op)
	}
	if len(customers) > 0 {
		(*bill).CustomerId = customers[0].Id
		(*bill).CustomerName = customers[0].Name
	} else {
		(*bill).CustomerName = custName
	}
	if bill.Implant == 3 {
		(*bill).ImplantedCount = bill.DentalCount
		(*bill).ImplantDate = bill.TradeAt
	}
	if bill.BrandName != "" {
		bn := strings.ToUpper(bill.BrandName)
		for _, v := range enums.DentalBrands {
			if strings.Contains(bn, v.Name) {
				bill.Brand = v.Id
				break
			}
			for _, a := range v.Alias {
				if strings.Contains(bn, a) {
					bill.Brand = v.Id
					break
				}
			}
			if bill.Brand > 0 {
				break
			}
		}
	}
	(*bill).TradeType = int(enums.TradeDeal)
	if bill.RealAmount == "" && bill.PaidAmount != "" {
		(*bill).TradeType = int(enums.TradeBalance)
	}
	if bill.PaidAmount == "" && bill.RealAmount != "" {
		(*bill).PaidAmount = bill.RealAmount
	}
	return nil
}

func getVal(data string) string {
	data = strings.ReplaceAll(data, "：", ":")
	arr := strings.Split(data, ":")
	if len(arr) == 2 {
		return strings.Trim(arr[1], " ")
	}
	return ""
}

func getDate(tmpD string) string {
	var sep string
	if strings.Contains(tmpD, ".") {
		sep = "."
	} else if strings.Contains(tmpD, "-") {
		sep = "-"
	} else if strings.Contains(tmpD, "/") {
		sep = "/"
	} else if strings.Contains(tmpD, ":") {
		sep = ":"
	} else if strings.Contains(tmpD, ",") {
		sep = ","
	} else if strings.Contains(tmpD, "。") {
		sep = "。"
	}
	arr := strings.Split(tmpD, sep)
	if len(arr) == 3 {
		for idx, d := range arr {
			d = strings.Trim(d, " ")
			if len(d) == 1 {
				d = "0" + d
			}
			arr[idx] = d
		}
		return strings.Join(arr, "-")
	}
	return ""
}

/*
*	日统计文字版
 */
func (s *BillService) StDay(teamId, userId int, deptPath string, day time.Time, reqId string) ([]string, error) {
	var texts []string
	if teamId < 1 {
		return texts, codes.ErrInvalidParameter(reqId, "teamId is nil")
	}
	today := utils.GetZoreTimeLocal(day)
	end := today.Add(24 * time.Hour)
	begin := utils.GetMonthFirstDayLocal(day)
	unixToday := today.Unix()

	var curM smodels.SysMember
	if deptPath == "" {
		service.SerSysMember.GetMember(teamId, userId, &curM)
		deptPath = curM.DeptPath
	}

	db := s.DB().Where("team_id = ?", teamId).Where("trade_at >=?", begin).
		Where("trade_at < ?", end).Where("dept_path like ?", deptPath+"%")
	var list []models.Bill
	if err := db.Find(&list).Error; err != nil {
		return texts, err
	}
	var totalDeal, totalPaid, totalDebt, totalrRefund, deal, paid, debt, refund, arrear decimal.Decimal
	var firstCnt, dealCnt int
	for _, b := range list {
		totalDeal = totalDeal.Add(b.RealAmount)
		totalPaid = totalPaid.Add(b.PaidAmount)
		totalDebt = totalDebt.Add(b.DebtAmount)
		totalrRefund = totalrRefund.Add(b.RefundAmount)
		if b.TradeAt.Unix() >= unixToday {
			deal = deal.Add(b.RealAmount)
			paid = paid.Add(b.PaidAmount)
			debt = debt.Add(b.DebtAmount)
			refund = refund.Add(b.RefundAmount)
			if b.TradeType == int(enums.TradeDeal) {
				arrear = arrear.Add(b.RealAmount.Sub(b.PaidAmount))
				//dealCnt += 1
			}
		}
	}

	var edList []models.EventDaySt
	if err := SerEventDaySt.GetList(teamId, 0, deptPath, today, end, &edList); err != nil {
		return texts, err
	}
	for _, ed := range edList {
		firstCnt += ed.FirstDiagnosis
		dealCnt += ed.Deal
	}
	todayPaid := paid.Add(debt).Sub(refund)
	tPaid := totalPaid.Add(totalDebt).Sub(totalrRefund)
	texts = append(texts, fmt.Sprintf("今日初诊数：%d", firstCnt))
	texts = append(texts, fmt.Sprintf("今日成交患者：%d", dealCnt))
	texts = append(texts, fmt.Sprintf("今日成交流水：%s", deal.StringFixedBank(0)))
	texts = append(texts, fmt.Sprintf("今日实收流水：%s", todayPaid.StringFixedBank(0)))
	texts = append(texts, fmt.Sprintf("今日欠款：%s", arrear.StringFixedBank(0)))
	texts = append(texts, fmt.Sprintf("总成交：%s", totalDeal.StringFixedBank(0)))
	texts = append(texts, fmt.Sprintf("总实收：%s", tPaid.StringFixedBank(0)))
	return texts, nil
}

/*
* 时间段内每人统计
 */
func (s *BillService) StQuery(teamId, userId int, deptPath string, begin, end time.Time, reqId string) ([]dto.BillUserStDto, error) {
	if teamId < 1 {
		return nil, codes.ErrInvalidParameter(reqId, "teamId is nil")
	}

	var curM smodels.SysMember
	if deptPath == "" {
		service.SerSysMember.GetMember(teamId, userId, &curM)
		deptPath = curM.DeptPath
	}

	var members []smodels.SysMember

	if err := service.SerSysMember.GetMembers(teamId, 0, deptPath, "", &members); err != nil {
		return nil, err
	}
	if end.IsZero() {
		end = time.Now()
	}
	if begin.IsZero() {
		begin = utils.GetMonthFirstDayLocal(end)
	}

	bt := utils.GetZoreTimeLocal(begin)
	et := utils.GetZoreTimeLocal(end).Add(24 * time.Hour)

	var taskList []models.TargetTask
	if err := SerTargetTask.GetTasks(enums.Month, bt.Year()*100+int(bt.Month()), teamId, 0, deptPath, &taskList); err != nil {
		return nil, err
	}
	m := make(map[int]dto.BillUserStDto, len(taskList))
	for _, task := range taskList {
		br, ok := m[task.UserId]
		if !ok {
			br = dto.BillUserStDto{
				UserId: task.UserId,
				Target: decimal.NewFromInt(int64(task.Deal)),
			}
		}
		for _, member := range members {
			if member.UserId == task.UserId {
				br.Name = member.Name
			}
		}
		m[br.UserId] = br
	}

	db := s.DB().Where("team_id = ?", teamId).Where("trade_at >=?", bt).
		Where("trade_at < ?", et)
	if userId > 0 {
		db.Where("user_id = ?", userId)
	} else if deptPath != "" {
		db.Where("dept_path like ?", deptPath+"%")
	}
	var list []models.Bill
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}

	for _, b := range list {
		br, ok := m[b.UserId]
		if !ok {
			br = dto.BillUserStDto{
				UserId: userId,
			}
		}
		br.Deal = br.Deal.Add(b.RealAmount)
		br.Paid = br.Paid.Add(b.PaidAmount)
		br.Debt = br.Debt.Add(b.DebtAmount)
		br.Refund = br.Refund.Add(b.RefundAmount)
		m[b.UserId] = br
	}

	var edList []models.EventDaySt
	if err := SerEventDaySt.GetList(teamId, 0, deptPath, bt, et, &edList); err != nil {
		return nil, err
	}

	for _, ed := range edList {
		br, ok := m[ed.UserId]
		if !ok {
			br = dto.BillUserStDto{
				UserId: userId,
			}
		}
		br.FirstDiagnosis += ed.FirstDiagnosis
		br.NewCustomerCnt += ed.NewCustomerCnt
		br.FurtherDiagnosis += ed.FurtherDiagnosis
		br.DealCnt += ed.Deal
		m[ed.UserId] = br
	}

	res := make([]dto.BillUserStDto, 0)
	for _, v := range m {
		v.CurDebt = v.Deal.Sub(v.Paid)
		v.Total = v.Paid.Add(v.Debt).Sub(v.Refund)
		res = append(res, v)
	}
	return res, nil
}

func (s *BillService) StMonth(teamId, userId int, deptPath string, day time.Time, reqId string) ([]string, error) {
	var texts []string
	if teamId < 1 {
		return texts, codes.ErrInvalidParameter(reqId, "teamId is nil")
	}

	var tu smodels.SysMember
	if err := service.SerSysMember.GetMember(teamId, userId, &tu); err != nil {
		return texts, err
	}

	if deptPath == "" {
		deptPath = tu.DeptPath
	}

	var members []smodels.SysMember
	if err := service.SerSysMember.GetMembers(teamId, 0, deptPath, "", &members); err != nil {
		return texts, err
	}

	today := utils.GetZoreTimeLocal(day)
	end := today.Add(24 * time.Hour)
	begin := utils.GetMonthFirstDayLocal(day)
	unixToday := today.Unix()

	db := s.DB().Where("team_id = ?", teamId).Where("trade_at >=?", begin).
		Where("trade_at < ?", end)
	db.Where("dept_path like ?", deptPath+"%")
	var list []models.Bill
	if err := db.Find(&list).Error; err != nil {
		return texts, err
	}

	var tmDeal, tPaid, tbDebt, tRefund, deal, paid, bdebt, refund decimal.Decimal
	var dealCnt, dCnt, iCnt, tdCnt, tiCnt int
	for _, b := range list {
		dCnt += b.DentalCount
		iCnt += b.ImplantedCount

		tmDeal = tmDeal.Add(b.RealAmount)
		tPaid = tPaid.Add(b.PaidAmount)
		tbDebt = tbDebt.Add(b.DebtAmount)
		tRefund = tRefund.Add(b.RefundAmount)
		if b.TradeAt.Unix() >= unixToday {
			deal = deal.Add(b.RealAmount)
			paid = paid.Add(b.PaidAmount)
			bdebt = bdebt.Add(b.DebtAmount)
			refund = refund.Add(b.RefundAmount)
			dealCnt += 1
			tdCnt += b.DentalCount
			tiCnt += b.ImplantedCount
		}
	}

	totalPaid := tPaid.Add(tbDebt).Sub(tRefund)
	todayPaid := paid.Add(bdebt).Sub(refund)
	totalDebt := tmDeal.Sub(tPaid) //欠款

	dayFmt := day.Format("2006年01月02日")

	texts = append(texts, dayFmt)

	texts = append(texts, fmt.Sprintf("汇报人：%s", tu.Name))

	var taskList []models.TargetTask
	if err := SerTargetTask.GetTasks(enums.Month, today.Year()*100+int(today.Month()), teamId, 0, deptPath, &taskList); err != nil {
		return texts, err
	}
	var tmpLen int
	var memberLen int
	var totalDeal int
	for _, task := range taskList {
		if task.Deal > 0 {
			totalDeal += task.Deal
			memberLen++
		} else {
			tmpLen++
		}
	}
	texts = append(texts, fmt.Sprintf("本月团队任务：%s", utils.MoneyFmt(float64(totalDeal))))
	if tmpLen > 0 {
		texts = append(texts, fmt.Sprintf("未完成任务：%d+%d", memberLen, tmpLen))
	} else {
		texts = append(texts, fmt.Sprintf("人员数量：%d", memberLen))
	}

	var edList []models.EventDaySt
	if err := SerEventDaySt.GetList(teamId, 0, deptPath, begin, end, &edList); err != nil {
		return texts, err
	}

	var stDay []string

	var tNc, tFirD, tFuD, tDeal int
	var dayNc, dayFirD, dayFuD, dayDeal, dayIv int
	for _, ed := range edList {
		if ed.Day.Unix() >= unixToday { //今日
			dayNc += ed.NewCustomerCnt
			dayFirD += ed.FirstDiagnosis
			dayFuD += ed.FurtherDiagnosis
			dayDeal += ed.Deal
			dayIv += ed.Invitation
			for _, m := range members {
				if m.UserId == ed.UserId {
					if ed.Rest == 2 {
						stDay = append(stDay, fmt.Sprintf("%s：0休息", m.Name))
					} else {
						stDay = append(stDay, fmt.Sprintf("%s：留存%d初诊%d复诊%d成交%d", m.Name, ed.NewCustomerCnt, ed.FirstDiagnosis, ed.FurtherDiagnosis, ed.Deal))
					}
					break
				}
			}

		}
		tNc += ed.NewCustomerCnt
		tFirD += ed.FirstDiagnosis
		tFuD += ed.FurtherDiagnosis
		tDeal += ed.Deal
	}
	var spday models.SummaryPlanDay

	if err := SerSummaryPlanDay.GetByDay(teamId, userId, today, &spday); err != nil {
		return texts, err
	}

	texts = append(texts, fmt.Sprintf("今日留存信息：%d", dayNc))
	texts = append(texts, fmt.Sprintf("今日邀约到诊：%d", dayFirD))
	texts = append(texts, fmt.Sprintf("今日成交患者：%d", dayDeal))
	texts = append(texts, fmt.Sprintf("今日种植颗数：%d", tdCnt))
	texts = append(texts, fmt.Sprintf("明日邀约患者：%d", dayIv))
	texts = append(texts, fmt.Sprintf("本月留存患者数：%d", tNc))
	texts = append(texts, fmt.Sprintf("本月初诊患者数：%d", tFirD))
	texts = append(texts, fmt.Sprintf("本月成交患者数：%d", tDeal))

	if tFirD == 0 {
		texts = append(texts, "本月患者成交率：0%")
	} else {
		f := fmt.Sprintf("%d%%", tDeal*100/tFirD)
		texts = append(texts, fmt.Sprintf("本月患者成交率：%s", f))
	}

	texts = append(texts, fmt.Sprintf("种植颗数：%d", dCnt))
	texts = append(texts, fmt.Sprintf("延期颗数：%d", dCnt-iCnt))
	texts = append(texts, fmt.Sprintf("成交总流水：%s", tmDeal.StringFixedBank(0)))

	texts = append(texts, fmt.Sprintf("总欠款金额：%s", totalDebt.StringFixedBank(0)))
	texts = append(texts, fmt.Sprintf("本月实收：%s", totalPaid.StringFixedBank(0)))
	texts = append(texts, fmt.Sprintf("今日实收：%s", todayPaid.StringFixedBank(0)))
	if tmDeal.IsZero() {
		texts = append(texts, "实收率：0%")
	} else {
		texts = append(texts, fmt.Sprintf("实收率：%s", fmt.Sprintf("%s%%", totalPaid.Div(tmDeal).Mul(decimal.NewFromInt(100)).StringFixedBank(0))))
	}
	if memberLen == 0 {
		texts = append(texts, "团队人效：0%%")
	} else {
		texts = append(texts, fmt.Sprintf("团队人效：%s", totalPaid.Div(decimal.NewFromInt(int64(memberLen))).StringFixedBank(0)))
	}
	texts = append(texts, fmt.Sprintf("收回上月欠款：%s", tbDebt.StringFixedBank(0)))

	befMonth := begin.AddDate(0, -1, 0)
	deferDental := 0
	if m, err := s.StDental(teamId, 0, deptPath, befMonth, begin); err == nil {
		deferDental = m.Total - m.Implanted
	}

	texts = append(texts, fmt.Sprintf("上月延期种植：%d", deferDental)) //TODO

	dp := fmt.Sprintf("%d%%", today.Day()*100/utils.GetMonthLen(today))
	texts = append(texts, fmt.Sprintf("本月时间进度：%s", dp))

	if tmDeal.IsZero() {
		texts = append(texts, "本月任务达成率：0%")
	} else {
		tp := fmt.Sprintf("%s%%", totalPaid.Div(decimal.NewFromInt(int64(totalDeal))).Mul(decimal.NewFromInt(100)).StringFixedBank(0))
		texts = append(texts, fmt.Sprintf("本月任务达成率：%s", tp))
	}

	texts = append(texts, "")
	texts = append(texts, "今日工作汇报：")
	texts = append(texts, spday.Summary)
	texts = append(texts, "")

	texts = append(texts, fmt.Sprintf("今日留存：%s", strconv.Itoa(dayNc)))

	texts = append(texts, stDay...)

	texts = append(texts, "")
	texts = append(texts, "明日工作计划：")
	texts = append(texts, spday.Plan)

	return texts, nil
}

func (s *BillService) StDental(teamId, userId int, deptPath string, begin, end time.Time) (dto.DentalStDto, error) {
	var m dto.DentalStDto
	db := s.DB().Model(&models.Bill{}).Select("sum(dental_count) as total", "sum(implanted_count) as implanted").
		Where("team_id = ?", teamId).Where("trade_at >=?", begin).Where("trade_at < ?", end)
	if userId != 0 {
		db.Where("user_id = ?", userId)
	} else if deptPath != "" {
		db.Where("dept_path like ?", deptPath+"%")
	}
	err := db.First(&m).Error
	return m, err
}
