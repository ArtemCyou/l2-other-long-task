package main

import (
	"reflect"
	"testing"
)

func TestAnagram(t *testing.T) {
	tests := struct {
		args []string
		want map[string]*[]string
	}{
		args: []string{"Пятак", "пятка", "Тяпка", "листок", "слиток", "столик", "отвар", "автор", "автор", "товар"},
		want: map[string]*[]string{
			"листок": {"листок", "слиток", "столик"},
			"пятак":  {"пятак", "пятка", "тяпка"},
			"отвар": {"автор", "отвар", "товар"},
		},
	}


if got:= findAnagramm(&tests.args);  !reflect.DeepEqual(*got, tests.want) {
	t.Errorf("Ожидали = %v, получили %v",  tests.want, got)
}
}
