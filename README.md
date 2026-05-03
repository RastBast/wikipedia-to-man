# Wiki2Man
**Преобразует статьи из XML-дампов Википедии (ru/en/любые языки, 30+ ГБ) в man-страницы.** Работает на **Linux/macOS/Windows** — оффлайн-справка в стиле Unix везде! [habr](https://habr.com/ru/companies/mvideo/articles/780776/)

## Зачем это
- Локальный доступ к Википедии без сети (Go, Linux, алгоритмы — как `man bash`).
- Поиск в гигантском дампе за секунды (потоковый bz2).
- Man-формат с заголовками, форматированием. [ru.wikipedia](https://ru.wikipedia.org/wiki/Man)

## Поддержка платформ
| Платформа | Статус | Примечания |
|-----------|--------|------------|
| **Linux** | ✅ Полная | Нативно, `man` встроен. |
| **macOS** | ✅ Полная | `man` встроен, Go native. |
| **Windows** | ✅ С WSL/WSL2 | Установите Go+WSL, `man` через Ubuntu. Альтернатива: Git Bash + less. | [github](https://github.com/containers/podman/issues/8452)

## Требования и скачивание
1. **Go 1.21+**: [golang.org/dl](https://golang.org/dl) (Windows/macOS/Linux).
2. **Дамп Википедии** (~20–50 ГБ разархивировано):
   - Русский: [dumps.wikimedia.org/ruwiki/latest/ruwiki-latest-pages-articles-multistream.xml.bz2](https://dumps.wikimedia.org/ruwiki/latest/) (~30 ГБ). [academictorrents](https://academictorrents.com/details/a808c21cd92487fa9a28c238b88a13dcc7ebaf4e)
   - Английский: [dumps.wikimedia.org/enwiki/latest/enwiki-latest-pages-articles-multistream.xml.bz2](https://dumps.wikimedia.org/enwiki/latest/) (~200 ГБ!).
   - Любой язык: Замените `ruwiki`/`enwiki` на `dewiki` и т.д.
3. **Зависимости Go** (авто): Только `github.com/cosnicolaou/pbzip2` для bz2.
4. **На Windows**: Установите [WSL2](https://learn.microsoft.com/ru-ru/windows/wsl/install) + Ubuntu (`sudo apt install man-db`).
5. **Диск**: 50+ ГБ свободно (дамп + temp).

## Установка (2 минуты)
```bash
# 1. Клонируйте/создайте проект
git clone <repo> && cd wiki2man
go mod init wiki2man

# 2. Зависимости (1 команда)
go mod tidy  # Авто скачает pbzip2

# 3. Скачайте дамп в корень (ru/en/...)
wget https://dumps.wikimedia.org/ruwiki/latest/ruwiki-latest-pages-articles-multistream.xml.bz2

# 4. Соберите
go build -o wiki2man

# Windows (WSL): то же самое
```
Готово! Бинарник ~5 МБ, кросс-платформенный. [praxiscode](https://praxiscode.io/knowledge-base/golang-mod-guide)

## Использование
```bash
./wiki2man
# Название статьи: "Go (programming language)" или "Go"
# Файл: ./Go_programming_language  (man ./Go_programming_language)
```
- **Любой дамп**: Код универсален (ru/en/de).
- **Точное название**: Первая буква uppercase.
- Время: 5–60 сек (зависит от размера/позиции). [github](https://github.com/dustin/go-wikiparse)

## Пример (Linux/macOS)
```
$ ./wiki2man
Введите название: Linux kernel
[Найдено за 15s]
$ man ./Linux_kernel
```

## Windows (WSL)
```
wsl ./wiki2man  # В Ubuntu WSL
man ./Article
```

## Troubleshooting
- **bz2 ошибка**: `go get github.com/cosnicolaou/pbzip2`.
- **man не видит**: `chmod +x Article.1; man ./Article.1`.
- **Большой дамп**: Тестируйте на малом (pages-articles.xml.bz2 ~1 ГБ).
- **Другие языки**: Просто скачайте дамп и запустите.

## Лицензия
MIT. PR welcome!

***
*Кросс-платформенный инструмент для локальной Вики-man. Работает с любыми dumps.wikimedia.org.* [desile.github](https://desile.github.io/2017/04/%D0%B4%D0%B0%D0%BC%D0%BF%D1%8B-ru.wikipedia.org-%D0%B8%D0%BF-0/)
