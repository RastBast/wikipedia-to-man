package main

import (
	"bytes"
	"context"
	"io"
	"os"

	"github.com/cosnicolaou/pbzip2"
)

// FindArticleInBz2 ищет статью по названию и возвращает полный XML блок <page>...</page>
func FindArticleInBz2(fileName, target string) ([]byte, error) {
	searchTag := []byte("<title>" + target + "</title>")
	startTag := []byte("<page>")
	endTag := []byte("</page>")

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Используем контекст и многопоточный bzip2 ридер
	ctx := context.Background()
	bzReader := pbzip2.NewReader(ctx, file)

	// Буфер 8МБ для баланса между скоростью и потреблением памяти
	buffer := make([]byte, 8*1024*1024)
	var leftover []byte

	for {
		n, err := bzReader.Read(buffer)
		if n > 0 {
			// Объединяем остаток с прошлого шага и новые данные
			data := append(leftover, buffer[:n]...)
			
			// Ищем заголовок в текущем куске данных
			titleIdx := bytes.Index(data, searchTag)
			if titleIdx != -1 {
				// Нашли заголовок, ищем границы <page> вокруг него
				pStart := bytes.LastIndex(data[:titleIdx], startTag)
				pEnd := bytes.Index(data[titleIdx:], endTag)

				if pStart != -1 && pEnd != -1 {
					totalEnd := titleIdx + pEnd + len(endTag)
					return data[pStart:totalEnd], nil
				}
			}

			// Оставляем в leftover всё, что идет после последнего <page>
			// Это гарантирует, что мы не разрежем тег пополам между чтениями
			lastStart := bytes.LastIndex(data, startTag)
			if lastStart != -1 {
				leftover = data[lastStart:]
			} else {
				leftover = nil
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}

	return nil, nil // Статья не найдена
}
