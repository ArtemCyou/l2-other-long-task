package db

import (
	"dev11/internals/app"
	"errors"
	"sync"
	"time"
)

var (
	mtxInstance sync.Mutex
	instance    *Cache = nil
)

type Cache struct {
	data map[int64]map[time.Time]app.Event
	mtx  sync.RWMutex
}

//инициализирует структуру Cache
func initCacheInstance() {
	mtxInstance.Lock()
	defer mtxInstance.Unlock()
	if instance == nil {
		cache := new(Cache)
		cache.data = make(map[int64]map[time.Time]app.Event)
		instance = cache
	}
}

// GetCacheInstance возвращает указатель на структуру Cache
func GetCacheInstance() *Cache {
	if instance == nil {
		initCacheInstance()
	}
	return instance
}

// Get возвращает данные из кэша по id.
// Если не находит данные, то возвращает ошибку.
func (c *Cache) Get(id int64) (map[time.Time]app.Event, error) {

	if c.data == nil {
		err := errors.New("Cache.Get кеш не инициализирован")
		return nil, err
	}

	// используем блокировку для чтения (не блокирует чтение для остальных, но блокирует запись)
	c.mtx.RLock()
	defer c.mtx.RUnlock()


	val, ok := c.data[id]

	if !ok {
		err := errors.New("объект с таким id не найден")
		return nil, err
	}
	// если данные есть, то возвращаем их
	return val, nil
}

func (c *Cache) Insert(data *app.Event) error {
	id := data.Id
	date := data.Data

	//проверяем, инициализирован ли кеш
	if c.data == nil {
		err := errors.New("Cache.Get кеш не инициализирован")
		return err
	}

	c.mtx.Lock()
	defer c.mtx.Unlock()

	// проверка, существуют ли в кэше данные с таким же id.
	// если не существуют, то инициализируем
	if _, ok := c.data[id]; !ok {
		c.data[id] = make(map[time.Time]app.Event)
	}

	// записываем данные в кэш
	c.data[id][date] = *data

	return nil
}

func (c *Cache) Update(data *app.Event) error {
	id := data.Id
	date := data.Data

	if c.data == nil {
		err := errors.New("Cache.Get кеш не инициализирован")
		return err
	}

	c.mtx.Lock()
	defer c.mtx.Unlock()

	// проверка, существуют ли в кэше данные с таким же id.
	// если не существуют, то возвращаем ошибку
	if _, ok := c.data[id]; !ok {
		err := errors.New("объект с таким id не найден")
		return err
	}
	if _, ok := c.data[id][date]; !ok {
		err := errors.New("объект с таким date не найден")
		return err
	}

	// обновляем данные в кэше
	c.data[id][date] = *data

	return nil

}

func (c *Cache) Delete(data *app.Event) error {
	id := data.Id
	date := data.Data

	if c.data == nil {
		err := errors.New("Cache.Get кеш не инициализирован")
		return err
	}

	c.mtx.Lock()
	defer c.mtx.Unlock()

	// считываем данные и их наличие
	// если данных нет, то возвращаем ошибку
	if _, ok := c.data[id]; !ok {
		err := errors.New("объект с таким id не найден")
		return err
	}
	if _, ok := c.data[id][date]; !ok {
		err := errors.New("объект с таким date не найден")
		return err
	}

	delete(c.data[id], date)

	return nil
}
