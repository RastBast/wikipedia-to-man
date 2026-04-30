package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func ConvertXmlToMan(xml []byte) string {
	tRe := regexp.MustCompile(`(?s)<title>(.*?)</title>`)
	cRe := regexp.MustCompile(`(?s)<text.*?>(.*?)</text>`)

	tMatch := tRe.FindSubmatch(xml)
	cMatch := cRe.FindSubmatch(xml)

	title := "Manual"
	if len(tMatch) > 1 {
		title = string(tMatch[1])
	}

	text := ""
	if len(cMatch) > 1 {
		text = string(cMatch[1])

		// 1. Чистка шаблонов {{...}}
		reT := regexp.MustCompile(`(?s)\{\{[^{}]*\}\}`)
		for reT.MatchString(text) {
			text = reT.ReplaceAllString(text, "")
		}

		// 2. Удаление таблиц
		text = regexp.MustCompile(`(?s)\{\|.*?\|\}`).ReplaceAllString(text, "\n[Таблица]\n")

		// 3. Удаление медиа и категорий
		text = regexp.MustCompile(`(?i)\[\[(Файл|File|Image|Категория|Category):.*?\]\]`).ReplaceAllString(text, "")

		// 4. ФИКС ЗАГОЛОВКОВ: схлопываем разрывы строк внутри знаков =
		text = regexp.MustCompile(`(?m)^\s*(=+)\s*\n+\s*(.*?)\s*\n+\s*(=+)`).ReplaceAllString(text, "$1 $2 $3")

		// 5. ФОРМАТИРОВАНИЕ ТЕКСТА (Жирный/Курсив)
		text = regexp.MustCompile(`'''''(.*?)'''''`).ReplaceAllString(text, `\f(BI$1\fP`)
		text = regexp.MustCompile(`'''(.*?)'''`).ReplaceAllString(text, `\fB$1\fP`)
		text = regexp.MustCompile(`''(.*?)''`).ReplaceAllString(text, `\fI$1\fP`)

		// 6. ПРЕВРАЩЕНИЕ В МАКРОСЫ MAN
		text = regexp.MustCompile(`(?m)^\s*===\s*(.*?)\s*===\s*`).ReplaceAllString(text, "\n.SS $1\n")
		text = regexp.MustCompile(`(?m)^\s*==\s*(.*?)\s*==\s*`).ReplaceAllString(text, "\n.SH $1\n")
		text = regexp.MustCompile(`(?m)^\s*=\s*(.*?)\s*=\s*`).ReplaceAllString(text, "\n.SH $1\n")

		// 7. ОЧИСТКА ХВОСТОВ
		text = regexp.MustCompile(`(?s)&lt;ref.*?&gt;.*?&lt;/ref&gt;`).ReplaceAllString(text, "")
		text = regexp.MustCompile(`(?s)&lt;.*?&gt;`).ReplaceAllString(text, "")
		text = regexp.MustCompile(`\[\[(?:[^|\]]*\|)?([^\]]+)\]\]`).ReplaceAllString(text, "$1")
		text = strings.ReplaceAll(text, "]]", "")

		// Списки: заменяем "*" на буллит (двойной слэш для Go)
		text = regexp.MustCompile(`(?m)^\s*\*\s*`).ReplaceAllString(text, "\n\\(bu ")

		// 8. ПАРАГРАФЫ
		text = regexp.MustCompile(`\n{3,}`).ReplaceAllString(text, "\n\n")
		text = strings.ReplaceAll(text, "\n\n", "\n.PP\n")
	}

	// Синяя ссылка
	url := fmt.Sprintf("\x1b[34mhttps:\\&//ru.wikipedia.org/wiki/%s\x1b[0m", strings.ReplaceAll(title, " ", "_"))

	return fmt.Sprintf(".TH \"%s\" 1 \"%s\"\n.SH NAME\n%s\n.SH DESCRIPTION\n%s\n.SH URL\n%s",
		strings.ToUpper(title), time.Now().Format("2006-01-02"), title, text, url)
}
