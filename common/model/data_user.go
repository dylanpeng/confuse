package model

var User = &userModel{
	DbBase: &DbBase{
		readDbName:  "slave",
		writeDbName: "master",
	},
}

type userModel struct {
	*DbBase
}
