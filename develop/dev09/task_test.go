package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestWget(t *testing.T) {
	// Создаем временный файл для сохранения данных
	tmpfile, err := ioutil.TempFile("", "testdata")
	if err != nil {
		t.Fatal("Ошибка при создании временного файла:", err)
	}
	defer os.Remove(tmpfile.Name())

	// URL-адрес для скачивания
	url := "https://mail.ru"

	// Запускаем утилиту wget с помощью команды exec.Command
	cmd := exec.Command("go", "run", "task.go", "-url", url, "-output", tmpfile.Name())
	err = cmd.Run()
	if err != nil {
		t.Fatal("Ошибка при выполнении утилиты wget:", err)
	}

	// Проверяем, что файл был создан и содержит данные
	fileInfo, err := os.Stat(tmpfile.Name())
	if err != nil {
		t.Fatal("Ошибка при получении информации о файле:", err)
	}
	if fileInfo.Size() == 0 {
		t.Error("Файл пуст")
	}

	// Проверяем, что файл содержит ожидаемые данные (например, содержимое сайта)
	expectedContent := "Example Domain"
	content, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatal("Ошибка при чтении файла:", err)
	}
	if !strings.Contains(string(content), expectedContent) {
		t.Errorf("Файл не содержит ожидаемое содержимое: %s", expectedContent)
	}
}
