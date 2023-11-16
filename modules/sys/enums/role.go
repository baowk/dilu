package enums

type RoleType int

const PlatformRoleSuper RoleType = 1

type PostType struct {
	Id   int
	Name string
	Key  string
}

var Admin = PostType{
	Id:   -1,
	Name: "超管",
	Key:  "Admin",
}

var DeptHead = PostType{
	Id:   2,
	Name: "部门主管",
	Key:  "DeptHead",
}

var DeptDeputy = PostType{
	Id:   4,
	Name: "部门副管",
	Key:  "DeptDeputy",
}

var Staff = PostType{
	Id:   8,
	Name: "职员",
	Key:  "Staff",
}
