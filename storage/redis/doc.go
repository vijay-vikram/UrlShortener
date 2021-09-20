package redis

var getUniqueRangeLuaScript = `local key = KEYS[1]

local startRange = 100000

local result = redis.call('GET', key)
if false == result then
   redis.call('SET', key, 200000)
   return startRange
end

local currentRange = tonumber(result)
local newRange  =  currentRange + startRange
redis.call('SET', key, newRange)

return currentRange`



var (
	uniqueKeyRangeLength int64 = 100000
)

const (
	uniqueRangeKey = "unique_keys_range"
	logTag = "storage.redis"
)
