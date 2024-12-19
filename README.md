# calc_go

```Сервер для подсчёта арифметических выражений```

## Table of contents
 - [Usage](#usage)
 - [Documentation](#documentation)
 - [Errors](#errors)
 - [Status_codes](#statuscodes)
 - [Examples](#examples)

## Usage

```
В данном разделе описана модель использования сервера
```
1. Скачайте проект
2. Разместите файлы проекта в одной папке
3. В терминале перейдите в эту папку (cd ..., где вместо ... вставьте путь)
4. go run ./cmd/main.go
5. Можете обращаться к серверу запросами с методом POST, данными типа json вида '{"expression":"example_expression"}'

## Documentation

```
В данном разделе описана структура проекта
```
Проект состоит из трёх папок: cmd, internal, pkg. В папке cmd содержится единственный файл main.go (package main), который содержит функцию main, запускающей в теле application.  
  
internal содержит в себе подпапку apllication, в которой содержаться два файла: application.go и application_test.go (package application). application.go содержит функции для работы с сервером. application_test.go предназначен для тестирования сервера.

pkg содержит в себе подпапку calculation, в которой содержится три файла (package calculation): calculation.go, calculation_test.go, errors.go. В calculation.go описаны функции для подсчёта арифметических выражений, в calculation_test.go содержатся тесты для отладки кода calculation.go, в errors.go - описание ошибок, возникающих при подсчёте результата.


## Errors
```
В данном разделе описаны ошибки, возникающие при вызове функции func Calc(expression) float64 пакета calculation. 
```
| Ошибка | Описание | Примеры выражений |
| :---:  | :---:    | :---:            |
| ErrInvalidExpression | Неправильный формат выражения: несоответсиве скобок,  несоотвествие операций | "+1+2"; "/1*2", "(1+2", "1+2(3*4)"|
| ErrExternalSymbols | Использование недопустимых символов | "2!", "3?", "4A" |
| ErrDivisionByZero | Ошибка деления на ноль | "1/0", "100/(2-2)" |
| ErrEmptyExpression | Пустое выражение | "", "   " |


## Status_codes
```
В данном разделе описаны статус-коды, возвращаемые сервером при обращении к нему и примеры запросов.
```
| Статус-код | Описание | Примеры запросов |
| :---: | :---: | :---: |
| 200 | Успешный запрос | |
| 422 | Ошибка при вычислении | |
| 500 | Иная ошибка | |

## Examples

```
В данном разделе представлены примеры запросов и ответом серверов
```
| Запрос | Ответ сервера |
| :---: | :---: |
|1|1|
|2|2|
|3|3|
|4|4|
|5|5|
