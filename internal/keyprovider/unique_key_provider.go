package keyprovider

import "github.com/UrlShortener/logger"

type uniqueKeyProviderImpl struct {
	RangeStart int64
	CurrentKey int64
	RangeEnd   int64
}

func Initialize(rangeStart int64, rangeEnd int64) {
	// TODO:// use sync.once
	uniqueKeyProvider = &uniqueKeyProviderImpl{
		RangeStart: rangeStart,
		CurrentKey: rangeStart,
		RangeEnd:   rangeEnd,
	}
	logger.Handler.Info(logTag, " initialized Unique key provider")
}

func (k *uniqueKeyProviderImpl) GetUniqueKey() int64 {
	if k.CurrentKey < k.RangeEnd {
		//	TODO :// get a new range and assign it
	}
	uniqueKey := k.CurrentKey
	k.CurrentKey++
	return uniqueKey
}

func GetUniqueKeyProvider() UniqueKeyProvider {
	return uniqueKeyProvider
}
