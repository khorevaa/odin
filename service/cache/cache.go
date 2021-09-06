package cache

type Cache interface {
	Connect()
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
	Clear(key string)
	HealthCheck() (bool, error)
}
