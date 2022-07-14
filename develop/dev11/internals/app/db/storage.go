package db

import (
	"dev11/internals/app"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

var cache *Cache

func init() {
	//инициализируем кеш
	cache = GetCacheInstance()
}

//добавляет данные в память
func InsertData(data *app.Event) error {
	// добавляем данные в кеш
	if err := cache.Insert(data); err != nil {
		return err
	}
	return nil
}

//обновляет данные в памяти
func UpdateData(data *app.Event) error {
	//обновляем данные в кеше
	if err := cache.Update(data); err != nil {
		return err
	}
	return nil

}

//удаляет данные из кеша
func DeleteData(data *app.Event) error {
	// удаляем данные из кеша
	if err := cache.Delete(data); err != nil {
		return err
	}
	return nil
}

//возвращает список событий за указанный период
func GetData(data *app.Event, days int) (string, error) {
	id := data.Id
	date := data.Data

	//получаем мапку событий
	events, err := cache.Get(id)
	if err != nil {
		return "", err
	}

	//отбираем нужные нам события
	var targetEvents []app.Event
	for i := 0; i < days; i++ {
		if event, ok := events[date.Add(time.Hour*24*time.Duration(i))]; ok {
			targetEvents = append(targetEvents, event)
		}
	}
	//если события не найдены, возвращаем ошибку
	if len(targetEvents) == 0 {
		err := errors.New("Совпадений не найдено")
		return "", err
	}

	// формируем возвращаемый список событий
	var result []string
	for _, event := range targetEvents {
		result = append(result, fmt.Sprintf("%v: %v", event.Data.Format("2006-01-02"), event.Events))
	}

	// сортируем список событий
	sort.Strings(result)

	return strings.Join(result, ",\n"), nil
}
