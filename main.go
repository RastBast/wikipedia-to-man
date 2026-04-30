package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Имя файла архива Википедии
	fileName := "ruwiki-20260201-pages-articles-multistream(2).xml.bz2"

	fmt.Print("Введите точное название статьи: ")
	target, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	target = strings.TrimSpace(target)
	
	if target == "" {
		fmt.Println("Название не может быть пустым")
		return
	}

	//Что бы первая буква становилась заглавной
	if len(target) > 0 {
    runes := []rune(target)
    runes[0] = []rune(strings.ToUpper(string(runes[0])))[0]
    target = string(runes)
	}


	fmt.Printf("Начинаю поиск статьи [%s] в файле %s...\n", target, fileName)
	startTime := time.Now()

	// Вызываем функцию поиска из второго файла
	// Она вернет нам "сырой" XML кусок статьи
	rawXML, err := FindArticleInBz2(fileName, target)
	
	if err != nil {
		fmt.Printf("Ошибка при поиске: %v\n", err)
		return
	}

	if rawXML == nil {
		fmt.Printf("Статья не найдена. Время поиска: %v\n", time.Since(startTime))
		return
	}

	fmt.Printf("[Найдено за %v]. Конвертирую в man...\n", time.Since(startTime))

	// Вызываем функцию конвертации из третьего файла
	manOutput := ConvertXmlToMan(rawXML)

	// Сохраняем результат в файл для просмотра через 'man ./result.1'
	resultFile := target
	err = os.WriteFile(resultFile, []byte(manOutput), 0644)
	if err != nil {
		fmt.Printf("Ошибка сохранения файла: %v\n", err)
		return
	}

	fmt.Printf("\nГотово! Проверьте результат командой:\nman ./%s\n", resultFile)
}
