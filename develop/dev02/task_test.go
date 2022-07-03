package main

import "testing"

func TestStringUnpack(t *testing.T) {
	//- "a4bc2d5e" => "aaaabccddddde"
	//- "abcd" => "abcd"
	//- "45" => "" (некорректная строка)
	//- "" => ""
	strTestTable := []struct {
		strTest  string
		expected string
	}{
		{
			strTest:  "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			strTest:  "abcd",
			expected: "abcd",
		},
		{
			strTest:  "45",
			expected: "",
		},
		{
			strTest:  "",
			expected: "",
		},
	}

	for _, testCase := range strTestTable {
		result, _ := Unpack(testCase.strTest)

		if result != testCase.expected {
			t.Errorf("Некорректный результат. Ожидалось: %s, получили: %s", testCase.expected, result)
		}
	}
}

func TestEscapeUnpack(t *testing.T) {
	//Дополнительное задание: поддержка escape - последовательностей
	//- qwe\4\5 => qwe45 (*)
	//- qwe\45 => qwe44444 (*)
	//- qwe\\5 => qwe\\\\\ (*)
	escTestTable := []struct {
		escTest  string
		expected string
	}{
		{
			escTest:  "qwe\\4\\5",
			expected: "qwe45",
		},
		{
			escTest:  "qwe\\45",
			expected: "qwe44444",
		},
		{
			escTest:  "qwe\\\\5",
			expected: "qwe\\\\\\\\\\",
		},
	}

	for _, testCase := range escTestTable {
		result, _ := Unpack(testCase.escTest)
		if testCase.expected != result {
			t.Errorf("Некорректный результат. Ожидалось: %s, получили: %s", testCase.expected, result)
		}
	}
}
