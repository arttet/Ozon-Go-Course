package storage

// Cache интерфейс кеша который нужно имплементировать
type Cache interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

// CacheWithMetrics объединяет два интерфейса в один
// Теперь те кто его имплементируют должны реализовать методы двух интерфейсов
type CacheWithMetrics interface {
	Cache
	Metrics
}
