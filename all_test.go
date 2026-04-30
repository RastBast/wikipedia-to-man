package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestConverter(t *testing.T) {
	input := []byte(`<page><title>Тест</title><text>{{Box}} == Заголовок == [[Файл:1.jpg]] Текст.</text></page>`)
	res := ConvertXmlToMan(input)

	if !strings.Contains(res, ".TH") || !strings.Contains(strings.ToUpper(res), "ТЕСТ") {
		t.Errorf("Ошибка заголовка. Получено:\n%s", res)
	}
	if strings.Contains(res, "{{") || strings.Contains(res, "Файл:") {
		t.Error("Мусор не удален")
	}
	if !strings.Contains(res, "\x1b[34m") {
		t.Error("Синий цвет ссылки отсутствует")
	}
}

func TestSearchInBz2(t *testing.T) {
	tmpName := "test_data.xml.bz2"
	xmlContent := "<page><title>Go</title><text>Fast</text></page>"

	cmd := exec.Command("bzip2", "-c")
	cmd.Stdin = strings.NewReader(xmlContent)
	out, err := cmd.Output()
	if err != nil { t.Skip("bzip2 не найден"); return }

	os.WriteFile(tmpName, out, 0644)
	defer os.Remove(tmpName)

	res, _ := FindArticleInBz2(tmpName, "Go")
	if !bytes.Contains(res, []byte("Fast")) {
		t.Error("Статья не найдена в bzip2")
	}
}
