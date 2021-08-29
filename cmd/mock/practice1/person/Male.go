package person

//这里的Male可以假设为与数据库交互的部分
type Male interface {
	Get(id int64) string
}
