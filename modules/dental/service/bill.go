package service

import (
	"dilu/common/codes"
	"dilu/common/consts"
	"dilu/common/utils"
	"dilu/modules/dental/enums"
	"dilu/modules/dental/models"
	"dilu/modules/dental/service/dto"
	smodels "dilu/modules/sys/models"

	senums "dilu/modules/sys/enums"
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

	"github.com/xuri/excelize/v2"
)

type BillService struct {
	*base.BaseService
}

var SerBill = BillService{
	base.NewService(consts.DB_CRM),
}

func (s *BillService) Page(teamId int, userId int, req dto.BillGetPageReq, list *[]dto.BillDto, total *int64) error {
	var tu smodels.SysMember
	service.SerSysMember.GetMember(teamId, userId, &tu)
	uid := 0
	var deptPath string
	if tu.PostId == senums.Staff.Id {
		uid = userId
	} else if tu.PostId > senums.Admin.Id {
		deptPath = tu.DeptPath
	}

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
	if uid > 0 {
		db.Where("user_id =?", uid)
	} else if deptPath != "" {
		db.Where("dept_path like?", deptPath+"%")
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
	if err := service.SerSysMember.GetMembers(req.TeamId, 0, "", bill.Name, 0, &members); err != nil {
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
func (s *BillService) StDay(teamId, userId int, deptPath string, day time.Time, reqId string) (string, error) {

	if teamId < 1 {
		return "", codes.ErrInvalidParameter(reqId, "teamId is nil")
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
		return "", err
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
		return "", err
	}
	for _, ed := range edList {
		firstCnt += ed.FirstDiagnosis
		dealCnt += ed.Deal
	}
	todayPaid := paid.Add(debt)
	tPaid := totalPaid.Add(totalDebt).Sub(totalrRefund)
	var texts utils.StringBuilder
	texts.Append(fmt.Sprintf("今日初诊数：%d\n", firstCnt))
	texts.Append(fmt.Sprintf("今日成交患者：%d\n", dealCnt))
	texts.Append(fmt.Sprintf("今日成交流水：%s\n", deal.StringFixedBank(0)))
	texts.Append(fmt.Sprintf("今日实收流水：%s\n", todayPaid.StringFixedBank(0)))
	texts.Append(fmt.Sprintf("今日欠款：%s\n", arrear.StringFixedBank(0)))
	texts.Append(fmt.Sprintf("总成交：%s\n", totalDeal.StringFixedBank(0)))
	texts.Append(fmt.Sprintf("总实收：%s\n", tPaid.StringFixedBank(0)))
	return texts.String(), nil
}

/*
* 时间段内每人统计
 */
func (s *BillService) StQuery(teamId, userId int, deptPath string, begin, end *time.Time, reqId string) ([]dto.BillUserStDto, error) {
	if teamId < 1 {
		return nil, codes.ErrInvalidParameter(reqId, "teamId is nil")
	}

	// var curM smodels.SysMember
	// if deptPath == "" {
	// 	service.SerSysMember.GetMember(teamId, userId, &curM)
	// 	deptPath = curM.DeptPath
	// }

	var members []smodels.SysMember

	if err := service.SerSysMember.GetMembers(teamId, 0, deptPath, "", 0, &members); err != nil {
		return nil, err
	}
	if end.IsZero() {
		*end = time.Now()
	}
	if begin.IsZero() {
		*begin = utils.GetMonthFirstDayLocal(*end)
	}

	bt := utils.GetZoreTimeLocal(*begin)
	et := utils.GetZoreTimeLocal(*end).Add(24 * time.Hour)

	var taskList []models.TargetTask
	if err := SerTargetTask.GetTasks(enums.Month, bt.Year()*100+int(bt.Month()), teamId, 0, deptPath, &taskList); err != nil {
		return nil, err
	}
	m := make(map[int]dto.BillUserStDto, len(taskList))
	for _, task := range taskList {
		br, ok := m[task.UserId]
		if !ok {
			br = dto.BillUserStDto{
				UserId:      task.UserId,
				Target:      decimal.NewFromInt(int64(task.Deal)),
				TargetNew:   task.NewCustomerCnt,
				TargetFirst: task.FirstDiagnosis,
				//TargetDealCnt: task.De
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

func (s *BillService) ExportBill(teamId, userId int, name string, deptPath string, begin, end *time.Time, reqId string) (*excelize.File, string, error) {
	db := s.DB().Order("user_id asc,trade_at asc,id asc")
	if teamId > 0 {
		db.Where("team_id = ?", teamId)
	}

	if end.IsZero() {
		*end = time.Now()
	}
	if begin.IsZero() {
		*begin = utils.GetMonthFirstDayLocal(*end)
	}

	bt := utils.GetZoreTimeLocal(*begin)
	et := utils.GetZoreTimeLocal(*end).Add(24 * time.Hour)

	db.Where("trade_at > ?", bt).Where("trade_at < ?", et)

	if userId > 0 {
		db.Where("user_id =?", userId)
	} else if deptPath != "" {
		db.Where("dept_path like?", deptPath+"%")
	}
	var ds []models.Bill
	db.Find(&ds)

	var mids []int
	var cids []int
	for _, v := range ds {
		mids = append(mids, v.UserId)
		cids = append(cids, v.CustomerId)
	}

	var members []smodels.SysMember
	if err := service.SerSysMember.GetMembersByUids(teamId, mids, &members); err != nil {
		return nil, "", err
	}

	var customers []models.Customer
	if err := SerCustomer.GetByIds(teamId, cids, &customers); err != nil {
		return nil, "", err
	}

	month := begin.Month()
	return s.BillExcel(int(month), name, ds, members, customers)
}

var billTitleClolumns = []string{"咨询师", "成交日期", "患者", "成交金额", "实收金额", "欠款金额", "奥齿泰", "皓圣", "雅定", "ITI", "诺贝尔", "延期", "备注"}

func (s *BillService) BillExcel(month int, name string, list []models.Bill, members []smodels.SysMember, customers []models.Customer) (*excelize.File, string, error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	borderColor := "696969"
	fontColor := "696969"

	err := f.SetColWidth("Sheet1", "A", "M", 13)
	if err != nil {
		fmt.Println(err)
	}

	err = f.MergeCell("Sheet1", "A1", "M1")
	if err != nil {
		fmt.Println(err)
	}
	title := fmt.Sprintf("%s组%d月份小组成交明细", name, month)
	f.SetCellValue("Sheet1", "A1", title)

	f.SetSheetRow("Sheet1", "A2", &billTitleClolumns)
	//f.SetCellStyle("Sheet1", "A2", fmt.Sprintf("A%d", len(titleClolumns)+1), style)

	for i, v := range list {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+3), v.UserId)
		for _, m := range members {
			if v.UserId == m.UserId {
				f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+3), m.Name)
				break
			}
		}
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+3), v.TradeAt.Format("01月02日"))
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+3), v.CustomerId)
		for _, c := range customers {
			if v.CustomerId == c.Id {
				f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+3), c.Name)
				break
			}
		}
		ar, _ := v.RealAmount.Float64()
		f.SetCellFloat("Sheet1", fmt.Sprintf("D%d", i+3), ar, 2, 32)
		pa, _ := v.PaidAmount.Float64()
		f.SetCellFloat("Sheet1", fmt.Sprintf("E%d", i+3), pa, 2, 32)
		var remark string
		if v.TradeType == 1 {
			d, _ := v.RealAmount.Sub(v.PaidAmount).Float64()
			f.SetCellFloat("Sheet1", fmt.Sprintf("F%d", i+3), d, 2, 32)
		} else if v.TradeType == 2 {
			d, _ := v.PaidAmount.Mul(decimal.NewFromInt(-1)).Float64()
			f.SetCellFloat("Sheet1", fmt.Sprintf("F%d", i+3), d, 2, 32)
			remark = "补当月款;"
		} else if v.TradeType == 3 {
			d, _ := v.DebtAmount.Float64()
			f.SetCellFloat("Sheet1", fmt.Sprintf("E%d", i+3), d, 2, 32)
			remark = "补上月款;"
		} else {
			d, _ := v.RefundAmount.Mul(decimal.NewFromInt(-1)).Float64()
			f.SetCellFloat("Sheet1", fmt.Sprintf("E%d", i+3), d, 2, 32)
			remark = "退款;"
		}
		if v.Brand == enums.DentalBrands[0].Id {
			f.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+3), v.DentalCount)
			f.SetCellValue("Sheet1", fmt.Sprintf("H%d", i+3), 0)
			f.SetCellValue("Sheet1", fmt.Sprintf("I%d", i+3), 0)
			f.SetCellValue("Sheet1", fmt.Sprintf("J%d", i+3), 0)
			f.SetCellValue("Sheet1", fmt.Sprintf("K%d", i+3), 0)
		} else if v.Brand == enums.DentalBrands[1].Id {
			f.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+3), 0)
			f.SetCellValue("Sheet1", fmt.Sprintf("H%d", i+3), v.DentalCount)
			f.SetCellValue("Sheet1", fmt.Sprintf("I%d", i+3), 0)
			f.SetCellValue("Sheet1", fmt.Sprintf("J%d", i+3), 0)
			f.SetCellValue("Sheet1", fmt.Sprintf("K%d", i+3), 0)
		} else if v.Brand == enums.DentalBrands[2].Id {
			f.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+3), 0)
			f.SetCellValue("Sheet1", fmt.Sprintf("H%d", i+3), 0)
			f.SetCellValue("Sheet1", fmt.Sprintf("I%d", i+3), v.DentalCount)
			f.SetCellValue("Sheet1", fmt.Sprintf("J%d", i+3), 0)
			f.SetCellValue("Sheet1", fmt.Sprintf("K%d", i+3), 0)
		} else if v.Brand == enums.DentalBrands[3].Id {
			f.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+3), 0)
			f.SetCellValue("Sheet1", fmt.Sprintf("H%d", i+3), 0)
			f.SetCellValue("Sheet1", fmt.Sprintf("I%d", i+3), 0)
			f.SetCellValue("Sheet1", fmt.Sprintf("J%d", i+3), v.DentalCount)
			f.SetCellValue("Sheet1", fmt.Sprintf("K%d", i+3), 0)
		} else if v.Brand == enums.DentalBrands[4].Id {
			f.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+3), 0)
			f.SetCellValue("Sheet1", fmt.Sprintf("H%d", i+3), 0)
			f.SetCellValue("Sheet1", fmt.Sprintf("I%d", i+3), 0)
			f.SetCellValue("Sheet1", fmt.Sprintf("J%d", i+3), 0)
			f.SetCellValue("Sheet1", fmt.Sprintf("K%d", i+3), v.DentalCount)
		}
		if v.Pack == int(enums.PackHalf) {
			remark += "半口;"
		} else if v.Pack == int(enums.PackFull) {
			remark += "全口;"
		}
		extension := v.DentalCount - v.ImplantedCount
		if extension < 0 {
			extension = 0
		}

		f.SetCellValue("Sheet1", fmt.Sprintf("L%d", i+3), extension)
		f.SetCellValue("Sheet1", fmt.Sprintf("M%d", i+3), remark+v.OtherPrj+" "+v.Remark)
	}

	last := len(list) + 3
	f.SetCellValue("Sheet1", fmt.Sprintf("A%d", last), "合计")
	for i := 4; i < len(billTitleClolumns); i++ {
		cell, err := excelize.CoordinatesToCellName(i, last)
		if err != nil {
			fmt.Println(err)
			return nil, "", err
		}
		f.SetCellFormula("Sheet1", cell, fmt.Sprintf("=SUM(%s%d:%s%d)", BASE_CLOUMN[i-1], 3, BASE_CLOUMN[i-1], last-1))
	}
	f.SetCellFormula("Sheet1", fmt.Sprintf("M%d", last), fmt.Sprintf("=SUM(G%d:K%d)", last, last))

	titleS, err2 := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: borderColor, Style: 1},
			{Type: "top", Color: borderColor, Style: 1},
			{Type: "bottom", Color: borderColor, Style: 1},
			{Type: "right", Color: borderColor, Style: 1},
		},
		Font: &excelize.Font{
			Bold:   true,
			Family: "微软雅黑",
			Size:   12,
			Color:  fontColor,
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: true,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     true,
			Vertical:        "",
			WrapText:        true,
		},
		Fill: excelize.Fill{
			Color:   []string{"8FBC8F"},
			Type:    "pattern",
			Pattern: 1,
		},
	})
	if err2 != nil {
		fmt.Println(err2)
	}
	f.SetCellStyle("Sheet1", "A1", "A1", titleS)

	style, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: borderColor, Style: 1},
			{Type: "top", Color: borderColor, Style: 1},
			{Type: "bottom", Color: borderColor, Style: 1},
			{Type: "right", Color: borderColor, Style: 1},
		},
		Font: &excelize.Font{
			//Bold:   true,
			Family: "微软雅黑",
			Size:   12,
			Color:  fontColor,
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: true,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     true,
			Vertical:        "",
			WrapText:        true,
		},
		Fill: excelize.Fill{
			Color:   []string{"F5FFFA"},
			Type:    "pattern",
			Pattern: 1,
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	f.SetCellStyle("Sheet1", "A2", fmt.Sprintf("M%d", last), style)

	titleC, err3 := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: borderColor, Style: 1},
			{Type: "top", Color: borderColor, Style: 1},
			{Type: "bottom", Color: borderColor, Style: 1},
			{Type: "right", Color: borderColor, Style: 1},
		},
		Font: &excelize.Font{
			//Bold:   true,
			Family: "微软雅黑",
			Size:   11,
			Color:  fontColor,
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: true,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     true,
			Vertical:        "",
			WrapText:        true,
		},
		Fill: excelize.Fill{
			Color:   []string{"FFFFFF"},
			Type:    "pattern",
			Pattern: 1,
		},
	})
	if err3 != nil {
		fmt.Println(err3)
	}
	f.SetCellStyle("Sheet1", "B3", fmt.Sprintf("M%d", last), titleC)
	// if err := f.SaveAs("Book1.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }

	return f, title, nil
}

func (s *BillService) ExportSt(teamId, userId int, name string, deptPath string, begin, end *time.Time, reqId string) (*excelize.File, string, error) {
	list, err := s.StQuery(teamId, userId, deptPath, begin, end, reqId)
	if err != nil {
		return nil, "", err
	}
	month := begin.Month()
	return s.StExcel(int(month), name, list)
}

var titleClolumns = []string{"", "姓名", "信息留存任务", "信息留存数量", "留存达成率", "到诊任务", "实际到诊", "到诊达成率", "到诊成交数量", "成交率", "成交金额", "待收欠款", "当月补欠款", "实收任务", "当月实收", "完成度", "备注"}

var BASE_CLOUMN = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func (s *BillService) StExcel(month int, name string, list []dto.BillUserStDto) (*excelize.File, string, error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	borderColor := "696969"
	fontColor := "696969"

	err := f.SetColWidth("Sheet1", "A", "Q", 17)

	if err != nil {
		fmt.Println(err)
		//return nil, "", err
	}

	err = f.MergeCell("Sheet1", "A1", "Q1")
	if err != nil {
		fmt.Println(err)
		return nil, "", err
	}
	title := fmt.Sprintf("%s组%d月份小组进度数据表", name, month)
	f.SetCellValue("Sheet1", "A1", title)

	f.SetSheetRow("Sheet1", "A2", &titleClolumns)
	//f.SetCellStyle("Sheet1", "A2", fmt.Sprintf("A%d", len(titleClolumns)+1), style)
	var targetNew, newCustomerCnt, targetFirst, firstDiagnosis, dealCnt int
	var total, target decimal.Decimal
	for i, v := range list {
		targetNew += v.TargetNew
		newCustomerCnt += v.NewCustomerCnt
		targetFirst += v.TargetFirst
		firstDiagnosis += v.FirstDiagnosis
		dealCnt += v.DealCnt
		total = total.Add(v.Total)
		target = target.Add(v.Target)

		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+3), i+1)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+3), v.Name)
		f.SetCellInt("Sheet1", fmt.Sprintf("C%d", i+3), v.TargetNew)
		f.SetCellInt("Sheet1", fmt.Sprintf("D%d", i+3), v.NewCustomerCnt)
		if v.TargetNew == 0 {
			f.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+3), "0%")
		} else {
			f.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+3), fmt.Sprintf("%.f%%", float64(v.NewCustomerCnt)*100.0/float64(v.TargetNew)))
		}
		f.SetCellInt("Sheet1", fmt.Sprintf("F%d", i+3), v.TargetFirst)
		f.SetCellInt("Sheet1", fmt.Sprintf("G%d", i+3), v.FirstDiagnosis)
		if v.TargetFirst == 0 {
			f.SetCellValue("Sheet1", fmt.Sprintf("H%d", i+3), "0%")
		} else {
			f.SetCellValue("Sheet1", fmt.Sprintf("H%d", i+3), fmt.Sprintf("%.f%%", float64(v.FirstDiagnosis)*100.0/float64(v.TargetFirst)))
		}
		f.SetCellInt("Sheet1", fmt.Sprintf("I%d", i+3), v.DealCnt)
		if v.FirstDiagnosis == 0 {
			f.SetCellValue("Sheet1", fmt.Sprintf("J%d", i+3), "0%")
		} else {
			f.SetCellValue("Sheet1", fmt.Sprintf("J%d", i+3), fmt.Sprintf("%.f%%", float64(v.DealCnt)*100.0/float64(v.FirstDiagnosis)))
		}

		d, _ := v.Deal.Float64()
		f.SetCellFloat("Sheet1", fmt.Sprintf("K%d", i+3), d, 2, 32)
		cd, _ := v.CurDebt.Float64()
		f.SetCellFloat("Sheet1", fmt.Sprintf("L%d", i+3), cd, 2, 32)
		debt, _ := v.Debt.Float64()
		f.SetCellFloat("Sheet1", fmt.Sprintf("M%d", i+3), debt, 2, 32)
		t, _ := v.Target.Float64()
		f.SetCellFloat("Sheet1", fmt.Sprintf("N%d", i+3), t, 2, 32)
		total, _ := v.Total.Float64()
		f.SetCellFloat("Sheet1", fmt.Sprintf("O%d", i+3), total, 2, 32)
		if v.Target.IsZero() {
			f.SetCellValue("Sheet1", fmt.Sprintf("P%d", i+3), "0%")
		} else {
			f.SetCellValue("Sheet1", fmt.Sprintf("P%d", i+3), fmt.Sprintf("%s%%", v.Total.Mul(decimal.New(100, 0)).Div(v.Target).StringFixed(0)))
		}
		f.SetCellValue("Sheet1", fmt.Sprintf("Q%d", i+3), "")
	}

	last := len(list) + 3
	f.SetCellValue("Sheet1", fmt.Sprintf("A%d", last), "合计")
	for i := 3; i < len(titleClolumns); i++ {
		cell, err := excelize.CoordinatesToCellName(i, last)
		if err != nil {
			fmt.Println(err)
			return nil, "", err
		}
		if i == 5 {
			if targetNew == 0 {
				f.SetCellValue("Sheet1", cell, "0%")
			} else {
				f.SetCellValue("Sheet1", cell, fmt.Sprintf("%.f%%", float64(newCustomerCnt)*100.0/float64(targetNew)))
			}
		} else if i == 8 {
			if targetFirst == 0 {
				f.SetCellValue("Sheet1", cell, "0%")
			} else {
				f.SetCellValue("Sheet1", cell, fmt.Sprintf("%.f%%", float64(firstDiagnosis)*100.0/float64(targetFirst)))
			}
		} else if i == 16 {
			if target.IsZero() {
				f.SetCellValue("Sheet1", cell, "0%")
			} else {
				f.SetCellValue("Sheet1", cell, fmt.Sprintf("%s%%", total.Mul(decimal.New(100, 0)).Div(target).StringFixedBank(0)))
			}
		} else if i == 10 {
			if firstDiagnosis == 0 {
				f.SetCellValue("Sheet1", cell, "0%")
			} else {
				f.SetCellValue("Sheet1", cell, fmt.Sprintf("%.f%%", float64(dealCnt)*100.0/float64(firstDiagnosis)))
			}
		} else {
			f.SetCellFormula("Sheet1", cell, fmt.Sprintf("=SUM(%s%d:%s%d)", BASE_CLOUMN[i-1], 3, BASE_CLOUMN[i-1], last-1))
		}

	}

	titleS, err2 := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: borderColor, Style: 1},
			{Type: "top", Color: borderColor, Style: 1},
			{Type: "bottom", Color: borderColor, Style: 1},
			{Type: "right", Color: borderColor, Style: 1},
		},
		Font: &excelize.Font{
			Bold:   true,
			Family: "微软雅黑",
			Size:   12,
			Color:  fontColor,
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: true,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     true,
			Vertical:        "",
			WrapText:        true,
		},
		Fill: excelize.Fill{
			Color:   []string{"8FBC8F"},
			Type:    "pattern",
			Pattern: 1,
		},
	})
	if err2 != nil {
		fmt.Println(err2)
	}
	f.SetCellStyle("Sheet1", "A1", "A1", titleS)

	style, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: borderColor, Style: 1},
			{Type: "top", Color: borderColor, Style: 1},
			{Type: "bottom", Color: borderColor, Style: 1},
			{Type: "right", Color: borderColor, Style: 1},
		},
		Font: &excelize.Font{
			//Bold:   true,
			Family: "微软雅黑",
			Size:   12,
			Color:  fontColor,
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: true,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     true,
			Vertical:        "",
			WrapText:        true,
		},
		Fill: excelize.Fill{
			Color:   []string{"F5FFFA"},
			Type:    "pattern",
			Pattern: 1,
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	f.SetCellStyle("Sheet1", "A2", fmt.Sprintf("Q%d", last), style)

	titleC, err3 := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: borderColor, Style: 1},
			{Type: "top", Color: borderColor, Style: 1},
			{Type: "bottom", Color: borderColor, Style: 1},
			{Type: "right", Color: borderColor, Style: 1},
		},
		Font: &excelize.Font{
			//Bold:   true,
			Family: "微软雅黑",
			Size:   11,
			Color:  fontColor,
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: true,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     true,
			Vertical:        "",
			WrapText:        true,
		},
		Fill: excelize.Fill{
			Color:   []string{"FFFFFF"},
			Type:    "pattern",
			Pattern: 1,
		},
	})
	if err3 != nil {
		fmt.Println(err3)
	}
	f.SetCellStyle("Sheet1", "B3", fmt.Sprintf("Q%d", last), titleC)

	// if err := f.SaveAs("Book1.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }

	return f, title, nil
}

func (s *BillService) StMonth(teamId, userId int, deptPath string, day time.Time, reqId string) (string, error) {

	if teamId < 1 {
		return "", codes.ErrInvalidParameter(reqId, "teamId is nil")
	}

	var tu smodels.SysMember
	if err := service.SerSysMember.GetMember(teamId, userId, &tu); err != nil {
		return "", err
	}

	if deptPath == "" {
		deptPath = tu.DeptPath
	}

	var members []smodels.SysMember
	if err := service.SerSysMember.GetMembers(teamId, 0, deptPath, "", 0, &members); err != nil {
		return "", err
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
		return "", err
	}

	var tmDeal, tPaid, tbDebt, tRefund, deal, paid, bdebt, refund decimal.Decimal
	var dealCnt, dCnt, iCnt, tdCnt, tiCnt int
	for _, b := range list {
		if b.TradeType == int(enums.TradeDeal) || b.TradeType == int(enums.TradeBalance) {
			dCnt += b.DentalCount
			iCnt += b.ImplantedCount
		}
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
	todayPaid := paid.Add(bdebt)
	totalDebt := tmDeal.Sub(tPaid) //欠款

	dayFmt := day.Format("2006年01月02日\n")

	var texts utils.StringBuilder
	texts.Append(dayFmt)

	texts.Append(fmt.Sprintf("汇报人：%s\n", tu.Name))

	var taskList []models.TargetTask
	if err := SerTargetTask.GetTasks(enums.Month, today.Year()*100+int(today.Month()), teamId, 0, deptPath, &taskList); err != nil {
		return "", err
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
	texts.Append(fmt.Sprintf("本月团队任务：%s\n", utils.MoneyFmt(float64(totalDeal))))
	if tmpLen > 0 {
		texts.Append(fmt.Sprintf("未完成任务：%d+%d\n", memberLen, tmpLen))
	} else {
		texts.Append(fmt.Sprintf("人员数量：%d\n", memberLen))
	}

	var edList []models.EventDaySt
	if err := SerEventDaySt.GetList(teamId, 0, deptPath, begin, end, &edList); err != nil {
		return "", err
	}

	var stDay utils.StringBuilder

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
						stDay.Append(fmt.Sprintf("%s：0休息\n", m.Name))
					} else {
						stDay.Append(fmt.Sprintf("%s：留存%d初诊%d复诊%d成交%d\n", m.Name, ed.NewCustomerCnt, ed.FirstDiagnosis, ed.FurtherDiagnosis, ed.Deal))
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
		return "", err
	}

	texts.Append(fmt.Sprintf("今日留存信息：%d\n", dayNc))
	texts.Append(fmt.Sprintf("今日邀约到诊：%d\n", dayFirD))
	texts.Append(fmt.Sprintf("今日成交患者：%d\n", dayDeal))
	texts.Append(fmt.Sprintf("今日种植颗数：%d\n", tdCnt))
	texts.Append(fmt.Sprintf("明日邀约患者：%d\n", dayIv))
	texts.Append(fmt.Sprintf("本月留存患者数：%d\n", tNc))
	texts.Append(fmt.Sprintf("本月初诊患者数：%d\n", tFirD))
	texts.Append(fmt.Sprintf("本月成交患者数：%d\n", tDeal))

	if tFirD == 0 {
		texts.Append("本月患者成交率：0%\n")
	} else {
		f := fmt.Sprintf("%d%%", tDeal*100/tFirD)
		texts.Append(fmt.Sprintf("本月患者成交率：%s\n", f))
	}

	texts.Append(fmt.Sprintf("种植颗数：%d\n", dCnt))
	texts.Append(fmt.Sprintf("延期颗数：%d\n", dCnt-iCnt))
	texts.Append(fmt.Sprintf("成交总流水：%s\n", tmDeal.StringFixedBank(0)))

	texts.Append(fmt.Sprintf("总欠款金额：%s\n", totalDebt.StringFixedBank(0)))
	texts.Append(fmt.Sprintf("本月实收：%s\n", totalPaid.StringFixedBank(0)))
	texts.Append(fmt.Sprintf("今日实收：%s\n", todayPaid.StringFixedBank(0)))
	if tmDeal.IsZero() {
		texts.Append("实收率：0%\n")
	} else {
		texts.Append(fmt.Sprintf("实收率：%s\n", fmt.Sprintf("%s%%", totalPaid.Div(tmDeal).Mul(decimal.NewFromInt(100)).StringFixedBank(0))))
	}
	if memberLen == 0 {
		texts.Append("团队人效：0%%\n")
	} else {
		texts.Append(fmt.Sprintf("团队人效：%s\n", totalPaid.Div(decimal.NewFromInt(int64(memberLen))).StringFixedBank(0)))
	}
	texts.Append(fmt.Sprintf("收回上月欠款：%s\n", tbDebt.StringFixedBank(0)))

	befMonth := begin.AddDate(0, -1, 0)
	deferDental := 0
	if m, err := s.StDental(teamId, 0, deptPath, befMonth, begin); err == nil {
		deferDental = m.Total - m.Implanted
	}

	texts.Append(fmt.Sprintf("上月延期种植：%d\n", deferDental)) //TODO

	dp := fmt.Sprintf("%d%%", today.Day()*100/utils.GetMonthLen(today))
	texts.Append(fmt.Sprintf("本月时间进度：%s\n", dp))

	if tmDeal.IsZero() {
		texts.Append("本月任务达成率：0%\n")
	} else {
		tp := fmt.Sprintf("%s%%", totalPaid.Div(decimal.NewFromInt(int64(totalDeal))).Mul(decimal.NewFromInt(100)).StringFixedBank(0))
		texts.Append(fmt.Sprintf("本月任务达成率：%s\n", tp))
	}

	texts.Append("\n")
	texts.Append("今日工作汇报：\n")
	texts.Append(spday.Summary)
	texts.Append("\n\n")

	texts.Append(fmt.Sprintf("今日留存：%s\n", strconv.Itoa(dayNc)))

	texts.Append(stDay.String())

	texts.Append("\n")
	texts.Append("明日工作计划：\n")
	texts.Append(spday.Plan)
	texts.Append("\n")
	return texts.String(), nil
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
