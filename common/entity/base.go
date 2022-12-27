package entity

type IEntity interface {
	TableName() string
	PrimaryPairs() []interface{}
	PrimarySeted() bool
	String() string
}
