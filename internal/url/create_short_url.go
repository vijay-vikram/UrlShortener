package url

import (
	"errors"
	"github.com/UrlShortener/internal/keyprovider"
	"github.com/UrlShortener/logger"
	"github.com/UrlShortener/storage/redis"
	"github.com/UrlShortener/utils"
	"strings"
)

func CreateShortUrl(actualUrl string) (string, error) {
	uniqueInteger := keyprovider.GetUniqueKeyProvider().GetUniqueKey()

	var shortenURLBuilder strings.Builder
	for uniqueInteger > 0 {
		_, err := shortenURLBuilder.WriteString(string(possibleCharacters[uniqueInteger%62]))
		if err != nil {
			logger.Handler.Error(tag, "Error when trying to create short url %v", err)
			return "", errors.New("error while creating short url")
		}
		uniqueInteger = uniqueInteger / 62
	}

	shortUrl := utils.Reverse(shortenURLBuilder.String())

	err := redis.GetRedisStorage().SaveFullUrl(shortUrl, actualUrl)
	if err != nil {
		return "", err
	}

	return domainName + shortUrl, nil
}
