package apis

// type GenColumnsApi struct {
// 	base.BaseApi
// }

// var ApiGenColumns = GenColumnsApi{}

// // QueryPage 获取GenColumns列表
// // @Summary 获取GenColumns列表
// // @Tags sys-GenColumns
// // @Accept application/json
// // @Product application/json
// // @Param teamId header int false "团队id"
// // @Param data body dto.GenColumnsGetPageReq true "body"
// // @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.GenColumns}} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sys/gen-columns/page [post]
// // @Security Bearer
// func (e *GenColumnsApi) QueryPage(c *gin.Context) {
// 	var req dto.GenColumnsGetPageReq
// 	if err := c.ShouldBind(&req); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	list := make([]models.GenColumns, 10)
// 	var total int64

// 	var model models.GenColumns
// 	if err := copier.Copy(&model, req); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	if err := service.SerGenColumns.Page(model, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	e.Page(c, list, total, req.GetPage(), req.GetSize())
// }

// // Get 获取GenColumns
// // @Summary 获取GenColumns
// // @Tags sys-GenColumns
// // @Accept application/json
// // @Product application/json
// // @Param teamId header int false "团队id"
// // @Param data body base.ReqId true "body"
// // @Success 200 {object} base.Resp{data=models.GenColumns} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sys/gen-columns/get [post]
// // @Security Bearer
// func (e *GenColumnsApi) Get(c *gin.Context) {
// 	var req base.ReqId
// 	if err := c.ShouldBind(&req); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	var data models.GenColumns
// 	if err := service.SerGenColumns.Get(req.Id, &data); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	e.Ok(c, data)
// }

// // Create 创建GenColumns
// // @Summary 创建GenColumns
// // @Tags sys-GenColumns
// // @Accept application/json
// // @Product application/json
// // @Param teamId header int false "团队id"
// // @Param data body dto.GenColumnsDto true "body"
// // @Success 200 {object} base.Resp{data=models.GenColumns} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sys/gen-columns/create [post]
// // @Security Bearer
// func (e *GenColumnsApi) Create(c *gin.Context) {
// 	var req dto.GenColumnsDto
// 	if err := c.ShouldBind(&req); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	var data models.GenColumns
// 	copier.Copy(&data, req)
// 	if err := service.SerGenColumns.Create(&data); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	e.Ok(c, data)
// }

// // Update 更新GenColumns
// // @Summary 更新GenColumns
// // @Tags sys-GenColumns
// // @Accept application/json
// // @Product application/json
// // @Param teamId header int false "团队id"
// // @Param data body dto.GenColumnsDto true "body"
// // @Success 200 {object} base.Resp{data=models.GenColumns} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sys/gen-columns/update [post]
// // @Security Bearer
// func (e *GenColumnsApi) Update(c *gin.Context) {
// 	var req dto.GenColumnsDto
// 	if err := c.ShouldBind(&req); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	var data models.GenColumns
// 	copier.Copy(&data, req)
// 	if err := service.SerGenColumns.UpdateById(&data); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	e.Ok(c, data)
// }

// // Del 删除GenColumns
// // @Summary 删除GenColumns
// // @Tags sys-GenColumns
// // @Accept application/json
// // @Product application/json
// // @Param teamId header int false "团队id"
// // @Param data body base.ReqIds true "body"
// // @Success 200 {object} base.Resp{data=models.GenColumns} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sys/gen-columns/del [post]
// // @Security Bearer
// func (e *GenColumnsApi) Del(c *gin.Context) {
// 	var req base.ReqIds
// 	if err := c.ShouldBind(&req); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	if err := service.SerGenColumns.DelIds(&models.GenColumns{}, req.Ids); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	e.Ok(c)
// }
