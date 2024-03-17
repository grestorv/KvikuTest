package stores

type Store interface {
	Init()
	Get(key string) any
	Set(key string, value any)
	SetWithTTL(key string, value any, ttl int)
	Delete(key string)
}
