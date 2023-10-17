package service

import (
	"dilu/common/codes"
	"dilu/common/consts"
	"dilu/common/utils"
	"dilu/modules/dental/enums"
	"dilu/modules/dental/models"
	"dilu/modules/dental/service/dto"
	smodles "dilu/modules/sys/models"
	"dilu/modules/sys/service"
	sdto "dilu/modules/sys/service/dto"
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

var (
	DAY_TMPL = `今日初诊数：%d
	今日成交患者：%d
	今日成交流水：%s
	今日实收流水：%s
	今日欠款:%s
	总成交：%s
	总实收：%s`

	DAY_TOTAL_TMPL = `{day}
	汇报人:{username}
	本月团队任务:%s
	人员数量:%d
	今日留存信息:%d
	今日邀约到诊:%d
	今日成交患者:%d
	今日种植颗数:%d
	明日邀约患者:%d
	本月留存患者数:%d
	本月初诊患者数:%d
	本月成交患者数:%d
	本月患者成交率:%s
	种植颗数:%d
	延期颗数:%d
	成交总流水:%s
	总欠款金额:%s
	本月实收:%s
	今日实收:%s
	实收率:%s
	团队人效:%s
	收回上月欠款:%s
	上月延期种植:%d
	本月时间进度:%s
	本月任务达成率:%s
	
	今日工作汇报
	%s
	
	今日留存:%d
	%s
	
	明日工作计划
	%s`

	DAY_MEMBER_TMPL = `%s:留存%d初诊%d复诊%d成交%d`
)

func (s *BillService) StDay(teamId, userId int, deptPath string, day time.Time, reqId string) (string, error) {
	if teamId < 1 {
		return "", codes.ErrInvalidParameter(reqId, "teamId is nil")
	}
	today := utils.GetZoreTime(day)
	end := today.Add(24 * time.Hour)
	begin := utils.GetMonthFirstDay(day)

	db := s.DB().Where("team_id = ?", teamId).Where("trade_at >=?", begin).
		Where("trade_at < ?", end)
	if userId > 0 {
		db.Where("user_id = ?", userId)
	} else if deptPath != "" {
		db.Where("dept_path like ?", deptPath+"%s")
	}
	var list []models.Bill
	if err := db.Find(&list).Error; err != nil {
		return "", err
	}
	var tDeal, tPaid, deal, paid, debt decimal.Decimal
	var firstCnt, dealCnt int
	for _, b := range list {
		if b.TradeType == 1 {
			tDeal = tDeal.Add(b.RealTotal)
			tPaid = tPaid.Add(b.PaidTotal)
			if b.TradeAt.After(today) {
				deal = deal.Add(b.RealTotal)
				paid = paid.Add(b.PaidTotal)
				debt = debt.Add(b.RealTotal.Sub(b.PaidTotal))
				dealCnt += 1
			}
		} else if b.TradeType == 2 {
			tPaid = tPaid.Add(b.PaidTotal)
			if b.TradeAt.After(today) {
				paid = paid.Add(b.PaidTotal)
			}
		} else if b.TradeType == 3 {
			tPaid = tPaid.Add(b.PaidTotal)
			if b.TradeAt.After(today) {
				paid = paid.Add(b.PaidTotal)
			}
		} else if b.TradeType == 10 {
			tPaid = tPaid.Sub(b.PaidTotal)
			if b.TradeAt.After(today) {
				paid = paid.Sub(b.PaidTotal)
			}
		}
	}
	var edList []models.EventDaySt
	if err := SerEventDaySt.GetList(teamId, userId, deptPath, today, end, &edList); err != nil {
		return "", err
	}
	for _, ed := range edList {
		firstCnt += ed.FirstDiagnosis
		dealCnt += ed.Deal
	}
	return fmt.Sprintf(DAY_TMPL, firstCnt, dealCnt, deal.StringFixedBank(0), paid.StringFixedBank(0), debt.StringFixedBank(0), tDeal.StringFixedBank(0), tPaid.StringFixedBank(0)), nil
}

func (s *BillService) StMonth(teamId, userId int, deptPath string, day time.Time, reqId string) (string, error) {
	if teamId < 1 {
		return "", codes.ErrInvalidParameter(reqId, "teamId is nil")
	}

	var members []smodles.SysMember

	if err := service.SerSysMember.GetMembers(teamId, userId, deptPath, "", &members); err != nil {
		return "", err
	}

	today := utils.GetZoreTime(day)
	end := today.Add(24 * time.Hour)
	begin := utils.GetMonthFirstDay(day)

	db := s.DB().Where("team_id = ?", teamId).Where("trade_at >=?", begin).
		Where("trade_at < ?", end)
	if userId > 0 {
		db.Where("user_id = ?", userId)
	} else if deptPath != "" {
		db.Where("dept_path like ?", deptPath+"%s")
	}
	var list []models.Bill
	if err := db.Find(&list).Error; err != nil {
		return "", err
	}

	var tmDeal, tPaid, tDebt, deal, paid, debt, befPaid decimal.Decimal
	var dealCnt, dCnt, iCnt, tdCnt, tiCnt int
	for _, b := range list {
		dCnt += b.DentalCount
		iCnt += b.ImplantedCount
		if b.TradeType == 1 {
			tmDeal = tmDeal.Add(b.RealTotal)
			tPaid = tPaid.Add(b.PaidTotal)
			tDebt = tDebt.Add(b.RealTotal.Sub(b.PaidTotal))
			if b.TradeAt.After(today) {
				deal = deal.Add(b.RealTotal)
				paid = paid.Add(b.PaidTotal)
				debt = debt.Add(b.RealTotal.Sub(b.PaidTotal))
				dealCnt += 1
				tdCnt += b.DentalCount
				tiCnt += b.ImplantedCount
			}
		} else if b.TradeType == 2 {
			tPaid = tPaid.Add(b.PaidTotal)
			tDebt = tDebt.Sub(b.PaidTotal)
			if b.TradeAt.After(today) {
				paid = paid.Add(b.PaidTotal)
			}
		} else if b.TradeType == 3 {
			tPaid = tPaid.Add(b.PaidTotal)
			befPaid = befPaid.Add(b.PaidTotal)
			if b.TradeAt.After(today) {
				paid = paid.Add(b.PaidTotal)
			}
		} else if b.TradeType == 10 {
			tPaid = tPaid.Sub(b.PaidTotal)
			if b.TradeAt.After(today) {
				paid = paid.Sub(b.PaidTotal)
			}
		}
	}

	dayFmt := day.Format("2006年01月03日")
	build := utils.NewSB()
	build.Append(dayFmt).Append("\n\t")
	var tu sdto.TeamMemberResp
	service.SerSysMember.GetTeamUser(teamId, userId, &tu)
	build.Append("汇报人:").Append(tu.Name).Append("\n\t")

	var taskList []models.TargetTask
	if err := SerTargetTask.GetTasks(enums.Month, today.Year()*100+int(today.Month()), teamId, userId, deptPath, &taskList); err != nil {
		return "", err
	}
	var memberLen int
	var totalDeal int
	for _, task := range taskList {
		if task.Deal > 0 {
			totalDeal += task.Deal
			memberLen++
		}
	}
	build.Append("本月团队任务:").Append(utils.MoneyFmt(float64(totalDeal))).Append("\n\t").
		Append("人员数量:").Append(strconv.Itoa(memberLen)).Append("\n\t")

	var edList []models.EventDaySt
	if err := SerEventDaySt.GetList(teamId, userId, deptPath, begin, end, &edList); err != nil {
		return "", err
	}

	stDay := utils.NewSB()

	var tNc, tFirD, tFuD, tDeal int
	var dayNc, dayFirD, dayFuD, dayDeal, dayIv int
	for _, ed := range edList {
		if ed.Day.After(today) { //今日
			dayNc += ed.NewCustomerCnt
			dayFirD += ed.FirstDiagnosis
			dayFuD += ed.FurtherDiagnosis
			dayDeal += ed.Deal
			dayIv += ed.Invitation
			for _, m := range members {
				if m.UserId == ed.UserId {
					stDay.Append(m.Name).Append(":")
					if ed.Rest == 2 {
						stDay.Append("0休息\n\t")
					} else {
						stDay.Append("留存").Append(strconv.Itoa(ed.NewCustomerCnt)).
							Append("初诊").Append(strconv.Itoa(ed.FirstDiagnosis)).
							Append("复诊").Append(strconv.Itoa(ed.FurtherDiagnosis)).
							Append("成交").Append(strconv.Itoa(ed.Deal)).Append("\n\t")
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
	build.Append("今日留存信息:").Append(strconv.Itoa(dayNc)).Append("\n\t")
	build.Append("今日邀约到诊:").Append(strconv.Itoa(dayNc)).Append("\n\t")
	build.Append("今日成交患者:").Append(strconv.Itoa(dayNc)).Append("\n\t")
	build.Append("今日种植颗数:").Append(strconv.Itoa(tdCnt)).Append("\n\t")
	build.Append("明日邀约患者:").Append(strconv.Itoa(dayIv)).Append("\n\t")
	build.Append("本月留存患者数:").Append(strconv.Itoa(tNc)).Append("\n\t")
	build.Append("本月初诊患者数:").Append(strconv.Itoa(tFirD)).Append("\n\t")
	build.Append("本月成交患者数:").Append(strconv.Itoa(tDeal)).Append("\n\t")
	f := fmt.Sprintf("%d%%", tDeal*100/tNc)
	build.Append("本月患者成交率:").Append(f).Append("\n\t")

	build.Append("种植颗数:").Append(strconv.Itoa(dCnt)).Append("\n\t")
	build.Append("延期颗数:").Append(strconv.Itoa(dCnt - iCnt)).Append("\n\t")
	build.Append("成交总流水:").Append(tmDeal.StringFixedBank(0)).Append("\n\t")
	build.Append("总欠款金额:").Append(tDebt.StringFixedBank(0)).Append("\n\t")
	build.Append("本月实收:").Append(tPaid.StringFixedBank(0)).Append("\n\t")
	build.Append("今日实收:").Append(paid.StringFixedBank(0)).Append("\n\t")
	build.Append("实收率:").Append(fmt.Sprintf("%.2f%%", tPaid.Div(tmDeal).InexactFloat64())).Append("\n\t")
	build.Append("团队人效:").Append(tPaid.Div(decimal.NewFromInt(int64(memberLen))).StringFixedBank(0)).Append("\n\t")
	build.Append("收回上月欠款:").Append(befPaid.StringFixedBank(0)).Append("\n\t")
	build.Append("上月延期种植:").Append(strconv.Itoa(dayNc)).Append("\n\t") //TODO
	build.Append("本月时间进度:").Append(strconv.Itoa(dayNc)).Append("\n\t") //TODO
	tp := fmt.Sprintf("%f%%", tPaid.Div(tmDeal).InexactFloat64())
	build.Append("本月任务达成率:").Append(tp).Append("\n\t")

	build.Append("\n\t").Append("今日工作汇报").Append("\n\t")
	build.Append("今日工作汇报").Append("\n\t") //TODO

	build.Append("\n\t").Append("今日留存:").Append(strconv.Itoa(dayNc)).Append("\n\t") //TODO
	build.Append(stDay.String()).Append("\n\t")

	build.Append("明日工作计划").Append("\n\t")
	build.Append("zzz") //TODO

	return build.String(), nil
}

func (s *BillService) CreateBill(reqId string, bill dto.IdentifyBillDto, dbill *models.Bill) errs.IError {
	if bill.TeamId < 1 {
		return codes.ErrInvalidParameter(reqId, "teamId is nil")
	}
	if bill.UserId < 1 {
		return codes.ErrInvalidParameter(reqId, "userId is nil")
	}
	var team smodles.SysTeam
	if err := service.SerSysTeam.Get(bill.TeamId, &team); err != nil {
		return codes.ErrNotFound(strconv.Itoa(bill.TeamId), "team", reqId, err)
	}
	var teamM smodles.SysMember
	tmWhere := smodles.SysMember{
		TeamId: bill.TeamId,
		UserId: bill.UserId,
	}
	if err := service.SerSysMember.GetByWhere(tmWhere, &teamM); err != nil {
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
				Birthday:    0,
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
	if dbill.Total.IsZero() {
		dbill.Total = dbill.RealTotal
	}
	dbill.CreatedAt = time.Now()
	dbill.UpdatedAt = dbill.CreatedAt
	if bill.ImplantDate != "" {
		if d, err := time.Parse("2006-01-02", bill.ImplantDate); err == nil {
			dbill.ImplantDate = d
		}
	}

	if bill.PaybackDate != "" {
		if d, err := time.Parse("2006-01-02", bill.PaybackDate); err == nil {
			dbill.PaybackDate = d
		}
	}

	if bill.TradeAt != "" {
		if d, err := time.Parse("2006-01-02", bill.TradeAt); err != nil {
			dbill.TradeAt = dbill.CreatedAt
		} else {
			dbill.TradeAt = d
		}
	}
	dbill.No = strings.Replace(dbill.CreatedAt.Format("20060102150405.000000"), ".", "", -1)

	if err := s.Create(dbill); err != nil {
		return codes.ErrSys(err)
	}
	return nil
}

func (s *BillService) Identify(req dto.BillTmplReq, bill *dto.IdentifyBillDto) errs.IError {
	(*bill).TeamId = req.TeamId
	arr := strings.Split(req.Text, "\n")
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
				(*bill).CustomerName = getVal(v)
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
				(*bill).RealTotal = getVal(v)
				break
			}
		}
		for _, key := range enums.Paid {
			if strings.Contains(v, key) {
				(*bill).PaidTotal = getVal(v)
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
				if imp == "是" {
					(*bill).Implant = 1
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
		bill.Pack = 2
	} else if bill.PrjName == "全口" {
		bill.Pack = 3
	} else {
		bill.Pack = 1
	}

	var members []smodles.SysMember
	if err := service.SerSysMember.GetMembers(req.TeamId, 0, "", bill.Name, &members); err != nil {
		core.Log.Error("获取咨询师错误", zap.Error(err))
		return errs.Err(codes.FAILURE, "", err)
	}
	if len(members) > 0 {
		(*bill).UserId = members[0].UserId
	}

	var customers []models.Customer
	if err := SerCustomer.GetByUserIdAndName(bill.UserId, 0, bill.CustomerName, &customers); err != nil {
		core.Log.Error("获取客户错误", zap.Error(err))
	}
	if len(customers) > 0 {
		(*bill).CustomerId = customers[0].Id
	}
	if bill.Implant == 1 {
		(*bill).ImplantedCount = bill.DentalCount
		(*bill).ImplantDate = bill.TradeAt
	}
	if bill.BrandName != "" {
		bn := strings.ToUpper(bill.BrandName)
		for _, v := range enums.DentalBrands {
			if bn == v.Name {
				bill.Brand = v.Id
				break
			}
			for _, a := range v.Alias {
				if a == bn {
					bill.Brand = v.Id
					break
				}
			}
			if bill.Brand > 0 {
				break
			}
		}
	}
	if bill.RealTotal == "" && bill.PaidTotal != "" {
		(*bill).RealTotal = bill.PaidTotal
	}
	if bill.PaidTotal == "" && bill.RealTotal != "" {
		(*bill).PaidTotal = bill.RealTotal
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
