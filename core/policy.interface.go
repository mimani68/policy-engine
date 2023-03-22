package core

type IPolicy interface {
	SetPolicy(realm, title, value string) bool
	CheckPolicy(realm, title, value string) bool
	LoadPolicy(db string, table string)
}
