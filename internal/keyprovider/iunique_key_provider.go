package keyprovider

type UniqueKeyProvider interface {
	GetUniqueKey() int64
}
