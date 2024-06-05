package enums

// 查询排序
type SubscriptionPageOrderBy int32

const (
	OrderByCreateTimeDesc SubscriptionPageOrderBy = 0
	OrderByCreateTimeAsc  SubscriptionPageOrderBy = 1
	OrderByIdAsc          SubscriptionPageOrderBy = 2
	OrderByIdDesc         SubscriptionPageOrderBy = 3
)
