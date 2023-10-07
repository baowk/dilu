package enums

type RoleType struct {
	Id   int
	Name string
	Key  string
}

var SuperAdmin = RoleType{
	Id:   1,
	Name: "超级管理",
	Key:  "SuperAdmin",
}

var TeamOwner = RoleType{
	Id:   2,
	Name: "团队创始人",
	Key:  "TeamOwner",
}

var DeptHead = RoleType{
	Id:   4,
	Name: "部门主管",
	Key:  "DeptHead",
}

var DeptDeputy = RoleType{
	Id:   8,
	Name: "部门副管",
	Key:  "DeptDeputy",
}

var Staff = RoleType{
	Id:   16,
	Name: "职员",
	Key:  "Staff",
}
