package main

import "testing"

func TestAnagram(t *testing.T) {
	tests := struct {
		args []string
		want map[string]*[]string
	}{
		args: []string{"Пятак", "пятка", "Тяпка", "листок", "слиток", "столик", "отвар", "автор", "автор", "товар"},
		want: map[string]*[]string{
			"листок": {"листок", "слиток", "столик"},
			"пятка":  {"пятак", "пятка", "тяпка"},
			"столик": {"автор", "отвар", "товар"},
		},
	}



}
