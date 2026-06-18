package storage

const TokenBucketLua = `
local tokens_key = KEYS[1]
local timestamp_key = KEYS[2]

local capacity = tonumber(ARGV[1])
local refill_rate = tonumber(ARGV[2])
local current_time = tonumber(ARGV[3])

local tokens = tonumber(redis.call("GET", tokens_key))
if tokens == nil then
	tokens = capacity
end

local last_refill = tonumber(redis.call("GET", timestamp_key))
if last_refill == nil then
	last_refill = current_time
end

local elapsed = current_time - last_refill
local refill = math.floor(elapsed * refill_rate)

tokens = math.min(capacity, tokens + refill)

local allowed = 0

if tokens >= 1 then
	allowed = 1
	tokens = tokens - 1
end

if refill > 0 then
    tokens = math.min(capacity, tokens + refill)
    last_refill = current_time
end


redis.call("SET", tokens_key, tokens)
redis.call("SET", timestamp_key, last_refill)

return {allowed, tokens}
`
