package entity

type IEntity interface {
	TableName() string
	PrimarySeted() bool
	String() string
}
