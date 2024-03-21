/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	col     = flag.Int("k", 1, "номер колонки для сортировки")
	numeric = flag.Bool("n", false, "сортировать по числовому значению")
	reverse = flag.Bool("r", false, "сортировать в обратном порядке")
	unique  = flag.Bool("u", false, "не выводить повторяющиеся строки")
)

func main() {
	flag.Parse()

	// Считываем строки из stdin
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Функция для сравнения строк в зависимости от флагов
	comparator := func(i, j int) bool {
		line1 := strings.Fields(lines[i])
		line2 := strings.Fields(lines[j])

		// Обработка ключа -k
		k := *col - 1
		if k >= len(line1) || k >= len(line2) {
			return false
		}

		// Преобразование к числовому значению, если указан ключ -n
		if *numeric {
			num1, err1 := strconv.Atoi(line1[k])
			num2, err2 := strconv.Atoi(line2[k])
			if err1 == nil && err2 == nil {
				if num1 != num2 {
					return num1 < num2
				}
			} else {
				// Если не удалось преобразовать к числу, то сравниваем как строки
				return lines[i] < lines[j]
			}
		}

		// Сравнение строк
		if lines[i] != lines[j] {
			if *reverse {
				return lines[i] > lines[j] // Инвертированный порядок
			}
			return lines[i] < lines[j]
		}
		return false
	}

	// Сортировка строк
	sort.SliceStable(lines, comparator)

	// Удаление повторяющихся строк, если указан флаг -u
	if *unique {
		lines = removeDuplicates(lines)
	}

	// Вывод отсортированных строк
	for _, line := range lines {
		fmt.Println(line)
	}
}

// Функция для удаления повторяющихся строк
func removeDuplicates(lines []string) []string {
	uniqueLines := make(map[string]bool)
	var result []string
	for _, line := range lines {
		if !uniqueLines[line] {
			uniqueLines[line] = true
			result = append(result, line)
		}
	}
	return result
}
