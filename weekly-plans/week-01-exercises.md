# Задачник недели 1 — переменные, типы, условия

[← Неделя 1](week-01.md) · [Оглавление](README.md)

Тренажёр к дню 2 недели 1. 30 задач: от объявления переменных до маленьких программ
с ветвлениями. Решения спрятаны под спойлером — сначала пиши сам, потом сверяйся.

## Как работать

1. Каждая задача — отдельная папка с `package main`:
   ```
   go-labs/week-01/task01/main.go
   go-labs/week-01/task02/main.go
   ...
   ```
2. Запуск из корня модуля:
   ```bash
   go run ./go-labs/week-01/task01
   ```
3. Правила:
   - без тернарного оператора (его в Go нет);
   - без `while` (только `for`);
   - все конверсии типов — явные;
   - обработай крайние случаи: ноль, отрицательные, пустую строку, неверный диапазон;
   - `gofmt -w .` после каждой задачи.
4. Если задача кажется слишком простой — добавь валидацию ввода и тесты.
5. Решил задачу — отметь её в `learning-log.md` одной строкой.

Уровни: ⭐ база, ⭐⭐ нужно подумать, ⭐⭐⭐ комбинированная.

---

## Блок A. Переменные и типы

### 1. ⭐ Все базовые типы

**Задача.** Объяви по одной переменной типов `int`, `int64`, `float64`, `string`, `bool`,
`byte`, `rune`. Выведи значение и тип каждой через `%T`.

**Ожидаемый вывод (пример).**
```
int = 42
int64 = 9000000000
float64 = 3.14
string = go
bool = true
uint8 = 65
int32 = 1071
```

**Подсказка.** `fmt.Printf("%T = %v\n", v, v)`. Тип `byte` — это алиас `uint8`,
`rune` — алиас `int32`, поэтому `%T` покажет именно их.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	var i int = 42
	var i64 int64 = 9_000_000_000
	var f float64 = 3.14
	var s string = "go"
	var b bool = true
	var by byte = 'A'
	var r rune = 'Я'

	fmt.Printf("%T = %v\n", i, i)
	fmt.Printf("%T = %v\n", i64, i64)
	fmt.Printf("%T = %v\n", f, f)
	fmt.Printf("%T = %v\n", s, s)
	fmt.Printf("%T = %v\n", b, b)
	fmt.Printf("%T = %v\n", by, by)
	fmt.Printf("%T = %v\n", r, r)
}
```

</details>

---

### 2. ⭐ Нулевые значения

**Задача.** Объяви без инициализации `int`, `float64`, `string`, `bool`, `*int`,
`[]int`, `map[string]int`. Выведи каждое значение.

**Ожидаемый вывод.**
```
0 0  false <nil> [] map[]
```

**Подсказка.** В Go нет неинициализированных переменных — есть нулевые значения.
Для указателя `%v` печатает `<nil>`.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	var i int
	var f float64
	var s string
	var b bool
	var p *int
	var sl []int
	var m map[string]int

	fmt.Println(i, f, s, b, p, sl, m)
	// 0 0  false <nil> [] map[]
}
```

</details>

---

### 3. ⭐ Градусы Цельсия → Фаренгейт

**Задача.** Дано `c := 36.6`. Переведи в фаренгейты по формуле $F = C \cdot 9/5 + 32$
и выведи с двумя знаками после запятой.

**Ожидаемый вывод.**
```
97.88 F
```

**Подсказка.** `c` должен быть `float64`, иначе деление `9/5` даст `1`.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	c := 36.6
	f := c*9/5 + 32
	fmt.Printf("%.2f F\n", f)
}
```

</details>

---

### 4. ⭐ Обмен значениями

**Задача.** Поменяй местами `a` и `b` без третьей переменной и выведи результат.

**Ожидаемый вывод.**
```
b=1 a=2
```

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	a, b := 1, 2
	a, b = b, a
	fmt.Printf("b=%d a=%d\n", b, a)
}
```

</details>

---

### 5. ⭐⭐ Среднее трёх чисел

**Задача.** Даны `a, b, c := 4, 7, 10`. Посчитай среднее арифметическое и выведи
с двумя знаками. Затем поменяй `c` на `9` и проверь, что вывод изменился.

**Ожидаемый вывод.**
```
7.00
```

**Подсказка.** `(a+b+c)/3` — целочисленное деление. Приведи сумму к `float64`
до деления. Сравни оба варианта и объясни разницу.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	a, b, c := 4, 7, 10
	avg := float64(a+b+c) / 3
	fmt.Printf("%.2f\n", avg) // 7.00
}
```

</details>

---

### 6. ⭐ Часы и минуты

**Задача.** Дано `total := 137`. Переведи в часы и минуты и выведи в формате `02:17`.

**Подсказка.** `/` и `%`. Формат — `%02d`.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	total := 137
	h := total / 60
	m := total % 60
	fmt.Printf("%02d:%02d\n", h, m)
}
```

</details>

---

### 7. ⭐⭐ Константы и iota

**Задача.** Объяви через `iota` константы `KB`, `MB`, `GB` типа `int64`,
где `KB = 1024`, `MB = 1024*1024`, `GB = 1024*1024*1024`. Выведи их.

**Ожидаемый вывод.**
```
1024 1048576 1073741824
```

**Подсказка.** `1 << (10 * iota)` плюс пустая константа `_` на нулевой позиции.

<details><summary>Решение</summary>

```go
package main

import "fmt"

const (
	_  = iota
	KB int64 = 1 << (10 * iota)
	MB
	GB
)

func main() {
	fmt.Println(KB, MB, GB)
}
```

</details>

---

### 8. ⭐⭐ Переполнение типа

**Задача.** Объяви `var x int8 = 127`, увеличь на 1 и выведи результат.
Затем попробуй `const c int8 = 128` и объясни, почему это не компилируется.

**Ожидаемый вывод.**
```
-128
```

**Подсказка.** Переполнение переменной — runtime-поведение (wrap-around),
переполнение константы — ошибка компиляции.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	var x int8 = 127
	x++
	fmt.Println(x) // -128

	// const c int8 = 128 // compile error: constant 128 overflows int8
}
```

</details>

---

## Блок B. Условия

### 9. ⭐ Чётное или нечётное

**Задача.** Для `n := -7` выведи `чётное` или `нечётное`. Проверь на `0`, `-7`, `8`.

**Подсказка.** `n%2 == 0` работает и для отрицательных.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	n := -7
	if n%2 == 0 {
		fmt.Println("чётное")
	} else {
		fmt.Println("нечётное")
	}
}
```

</details>

---

### 10. ⭐ Знак числа

**Задача.** Для `n := -3` выведи `положительное`, `отрицательное` или `ноль`.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	n := -3
	if n > 0 {
		fmt.Println("положительное")
	} else if n < 0 {
		fmt.Println("отрицательное")
	} else {
		fmt.Println("ноль")
	}
}
```

</details>

---

### 11. ⭐⭐ Високосный год

**Задача.** Для `year := 1900` определи, високосный ли год.
Проверь на `1900` (нет), `2000` (да), `2024` (да), `2023` (нет).

**Подсказка.** Год високосный, если делится на 4 и не делится на 100, либо делится на 400.
Обрати внимание на порядок условий и скобки.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	year := 1900
	leap := (year%4 == 0 && year%100 != 0) || year%400 == 0
	if leap {
		fmt.Println("високосный")
	} else {
		fmt.Println("обычный")
	}
}
```

</details>

---

### 12. ⭐⭐ Оценка по баллам

**Задача.** Для `score := 87` выведи оценку: `90–100 → A`, `80–89 → B`,
`70–79 → C`, `60–69 → D`, иначе `F`. Баллы вне `0..100` — `неверный балл`.

**Подсказка.** Сначала проверь диапазон, потом сравнивай по убыванию.
Порядок веток здесь критичен.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	score := 87
	switch {
	case score < 0 || score > 100:
		fmt.Println("неверный балл")
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}
}
```

</details>

---

### 13. ⭐ Максимум из трёх

**Задача.** Для `a, b, c := 12, 45, 7` найди максимум без встроенной `max`.
Затем перепиши через встроенную `max(a, b, c)` (Go 1.21+) и сравни.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	a, b, c := 12, 45, 7

	m := a
	if b > m {
		m = b
	}
	if c > m {
		m = c
	}
	fmt.Println(m)          // 45
	fmt.Println(max(a, b, c)) // 45
}
```

</details>

---

### 14. ⭐⭐⭐ Треугольник

**Задача.** Даны стороны `a, b, c := 5, 5, 8`. Проверь, что треугольник существует
(каждая сторона меньше суммы двух других и больше нуля), и определи тип:
равносторонний, равнобедренный или разносторонний.

**Подсказка.** Сначала проверка существования, затем `switch` по равенствам.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	a, b, c := 5, 5, 8

	if a <= 0 || b <= 0 || c <= 0 || a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("треугольник не существует")
		return
	}

	switch {
	case a == b && b == c:
		fmt.Println("равносторонний")
	case a == b || b == c || a == c:
		fmt.Println("равнобедренный")
	default:
		fmt.Println("разносторонний")
	}
}
```

</details>

---

### 15. ⭐⭐⭐ Счастливый билет

**Задача.** Для шестизначного номера `n := 123321` проверь, счастливый ли он
(сумма первых трёх цифр равна сумме последних трёх). Разбери номер на цифры
арифметикой. Проверь также `123456` и `000000`.

**Подсказка.** `n/100000%10`, `n/10000%10` и так далее.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	n := 123321
	if n < 0 || n > 999999 {
		fmt.Println("номер должен быть от 000000 до 999999")
		return
	}

	d1 := n / 100000 % 10
	d2 := n / 10000 % 10
	d3 := n / 1000 % 10
	d4 := n / 100 % 10
	d5 := n / 10 % 10
	d6 := n % 10

	if d1+d2+d3 == d4+d5+d6 {
		fmt.Println("счастливый")
	} else {
		fmt.Println("обычный")
	}
}
```

</details>

---

### 16. ⭐⭐ Логин и пароль

**Задача.** Даны `username`, `password`. Если логин совпадает с `admin`
без учёта регистра и пароль равен `"golang2026"` — выведи `доступ разрешён`,
иначе `доступ запрещён`.

**Подсказка.** `strings.EqualFold` для логина, обычное `==` для пароля.
В реальном коде пароль сравнивают только по хешу — это учебный пример.

<details><summary>Решение</summary>

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	username := "Admin"
	password := "golang2026"

	if strings.EqualFold(username, "admin") && password == "golang2026" {
		fmt.Println("доступ разрешён")
	} else {
		fmt.Println("доступ запрещён")
	}
}
```

</details>

---

### 17. ⭐⭐ Короткое замыкание

**Задача.** Даны `a := 10`, `b := 0`. Напиши условие, которое проверяет,
что `b != 0` и `a/b > 2`, не вызывая панику. Затем поменяй порядок условий
и объясни, что произойдёт.

**Подсказка.** `&&` вычисляется слева направо и останавливается на первом `false`.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	a, b := 10, 0

	if b != 0 && a/b > 2 {
		fmt.Println("больше двух")
	} else {
		fmt.Println("условие не выполнено")
	}

	// if a/b > 2 && b != 0 { ... } // panic: integer divide by zero
}
```

</details>

---

## Блок C. switch

### 18. ⭐ День недели

**Задача.** Для `day := 3` выведи название дня. Для значений вне `1..7` — `неверный день`.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	day := 3
	switch day {
	case 1:
		fmt.Println("понедельник")
	case 2:
		fmt.Println("вторник")
	case 3:
		fmt.Println("среда")
	case 4:
		fmt.Println("четверг")
	case 5:
		fmt.Println("пятница")
	case 6:
		fmt.Println("суббота")
	case 7:
		fmt.Println("воскресенье")
	default:
		fmt.Println("неверный день")
	}
}
```

</details>

---

### 19. ⭐⭐ Сезон по месяцу

**Задача.** Для `month := 10` выведи сезон. Используй несколько значений в одном `case`.

**Подсказка.** `case 12, 1, 2:` — несколько значений через запятую.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	month := 10
	switch month {
	case 12, 1, 2:
		fmt.Println("зима")
	case 3, 4, 5:
		fmt.Println("весна")
	case 6, 7, 8:
		fmt.Println("лето")
	case 9, 10, 11:
		fmt.Println("осень")
	default:
		fmt.Println("неверный месяц")
	}
}
```

</details>

---

### 20. ⭐⭐ Калькулятор

**Задача.** Даны `a, b := 10.0, 4.0` и `op := "/"`. Выведи результат операции
`+`, `-`, `*`, `/`. Для деления на ноль и неизвестного оператора — понятное сообщение.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	a, b := 10.0, 4.0
	op := "/"

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("деление на ноль")
			return
		}
		fmt.Println(a / b)
	default:
		fmt.Println("неизвестная операция")
	}
}
```

</details>

---

### 21. ⭐⭐ switch без выражения

**Задача.** Для `t := 21.5` выведи `мороз` (< 0), `холодно` (< 15),
`нормально` (< 25), `жарко` (иначе). Используй `switch` без выражения.

**Подсказка.** `switch { case t < 0: ... }` — это `switch true`.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	t := 21.5
	switch {
	case t < 0:
		fmt.Println("мороз")
	case t < 15:
		fmt.Println("холодно")
	case t < 25:
		fmt.Println("нормально")
	default:
		fmt.Println("жарко")
	}
}
```

</details>

---

### 22. ⭐⭐ Классификатор символа

**Задача.** Для строки `input := "Я"` возьми первый символ и определи:
буква, цифра, пробельный символ или что-то ещё. Проверь на `"7"`, `" "`, `"!"`.

**Подсказка.** `[]rune(input)`, `unicode.IsLetter`, `unicode.IsDigit`, `unicode.IsSpace`.
Не забудь проверить пустую строку.

<details><summary>Решение</summary>

```go
package main

import (
	"fmt"
	"unicode"
)

func main() {
	input := "Я"
	if input == "" {
		fmt.Println("пустая строка")
		return
	}

	r := []rune(input)[0]
	switch {
	case unicode.IsLetter(r):
		fmt.Println("буква")
	case unicode.IsDigit(r):
		fmt.Println("цифра")
	case unicode.IsSpace(r):
		fmt.Println("пробельный символ")
	default:
		fmt.Println("другое")
	}
}
```

</details>

---

### 23. ⭐⭐ switch с init-выражением

**Задача.** Для `password := "golang"` классифицируй длину: меньше 8 — `слабый`,
меньше 12 — `средний`, иначе `сильный`. Используй `switch` с инициализацией
переменной прямо в заголовке.

**Подсказка.** `switch n := len(password); { ... }`.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	password := "golang"
	switch n := len(password); {
	case n < 8:
		fmt.Println("слабый")
	case n < 12:
		fmt.Println("средний")
	default:
		fmt.Println("сильный")
	}
}
```

</details>

---

## Блок D. Комбинированные задачи

### 24. ⭐⭐ FizzBuzz

**Задача.** Выведи числа от 1 до 20: кратные 3 — `Fizz`, кратные 5 — `Buzz`,
кратные 15 — `FizzBuzz`, остальные — само число.

**Подсказка.** Проверяй делимость на 15 первой.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
```

</details>

---

### 25. ⭐⭐⭐ Простые числа до N

**Задача.** Выведи все простые числа от 2 до `n := 30` через пробел.

**Ожидаемый вывод.**
```
2 3 5 7 11 13 17 19 23 29
```

**Подсказка.** Внутренний цикл до `j*j <= i`, флаг `isPrime`, `break`.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	n := 30
	for i := 2; i <= n; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
```

</details>

---

### 26. ⭐ Таблица умножения

**Задача.** Выведи таблицу умножения от 1 до 5 ровными колонками шириной 4 символа.

**Подсказка.** Вложенные `for` и `fmt.Printf("%4d", ...)`, после внутреннего цикла — `fmt.Println()`.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}
```

</details>

---

### 27. ⭐⭐ Индекс массы тела

**Задача.** Даны `weight := 70.0` (кг) и `height := 1.75` (м). Посчитай ИМТ
и выведи категорию: `< 18.5` — недостаточный вес, `< 25` — норма,
`< 30` — избыточный вес, иначе — ожирение. ИМТ выведи с одним знаком.

**Формула.** $\text{ИМТ} = \frac{\text{вес}}{\text{рост}^2}$.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	weight, height := 70.0, 1.75
	bmi := weight / (height * height)

	fmt.Printf("ИМТ: %.1f\n", bmi)
	switch {
	case bmi < 18.5:
		fmt.Println("недостаточный вес")
	case bmi < 25:
		fmt.Println("норма")
	case bmi < 30:
		fmt.Println("избыточный вес")
	default:
		fmt.Println("ожирение")
	}
}
```

</details>

---

### 28. ⭐⭐⭐ Валидация пароля

**Задача.** Проверь пароль `"golang2026"` по правилам: длина не меньше 8,
нет пробельных символов, есть хотя бы одна буква и хотя бы одна цифра.
Выведи все нарушения или `пароль надёжный`.

**Подсказка.** `len`, `strings.ContainsAny`, цикл по `rune` с `unicode.IsLetter` /
`unicode.IsDigit`. Отдельно проверь `""`.

<details><summary>Решение</summary>

```go
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	password := "golang2026"

	ok := true
	if len(password) < 8 {
		fmt.Println("- короче 8 символов")
		ok = false
	}
	if strings.ContainsAny(password, " \t\n") {
		fmt.Println("- содержит пробельные символы")
		ok = false
	}

	hasLetter, hasDigit := false, false
	for _, r := range password {
		if unicode.IsLetter(r) {
			hasLetter = true
		}
		if unicode.IsDigit(r) {
			hasDigit = true
		}
	}
	if !hasLetter {
		fmt.Println("- нет ни одной буквы")
		ok = false
	}
	if !hasDigit {
		fmt.Println("- нет ни одной цифры")
		ok = false
	}

	if ok {
		fmt.Println("пароль надёжный")
	}
}
```

</details>

---

### 29. ⭐⭐ Скидка по сумме

**Задача.** Для `amount := 7300` посчитай скидку: от 10 000 — 15%, от 5 000 — 10%,
от 1 000 — 5%, иначе 0%. Выведи сумму скидки и итог.

**Ожидаемый вывод.**
```
скидка: 730, итого: 6570
```

**Подсказка.** Проценты считай целочисленно: `amount * percent / 100`.
Порядок `case` — по убыванию суммы.

<details><summary>Решение</summary>

```go
package main

import "fmt"

func main() {
	amount := 7300

	var percent int
	switch {
	case amount >= 10000:
		percent = 15
	case amount >= 5000:
		percent = 10
	case amount >= 1000:
		percent = 5
	}

	discount := amount * percent / 100
	fmt.Printf("скидка: %d, итого: %d\n", discount, amount-discount)
}
```

</details>

---

### 30. ⭐⭐⭐ Билетная касса (с flags)

**Задача.** Напиши программу с флагами `-age` и `-day` (`mon`…`sun`).
Базовая цена 1000. Возраст: до 7 — бесплатно, 7–17 — 50%, 18–64 — полная,
65+ — 70% от цены. Выходные (`sat`, `sun`) — наценка 20%.
Неверный возраст или день — ошибка с кодом выхода `1`.

**Запуск.**
```bash
go run ./go-labs/week-01/task30 -age 25 -day sat
```

**Ожидаемый вывод.**
```
1200
```

**Подсказка.** `flag.Int`, `flag.String`, `flag.Parse`; для ошибки — `fmt.Fprintln(os.Stderr, ...)`
и `os.Exit(1)`. Сначала посчитай базу, потом наценку.

<details><summary>Решение</summary>

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	age := flag.Int("age", 0, "возраст посетителя")
	day := flag.String("day", "mon", "день недели: mon..sun")
	flag.Parse()

	if *age < 0 || *age > 120 {
		fmt.Fprintln(os.Stderr, "неверный возраст")
		os.Exit(1)
	}
	validDay := map[string]bool{
		"mon": true, "tue": true, "wed": true, "thu": true,
		"fri": true, "sat": true, "sun": true,
	}
	if !validDay[*day] {
		fmt.Fprintln(os.Stderr, "неверный день недели")
		os.Exit(1)
	}

	price := 1000
	switch {
	case *age < 7:
		price = 0
	case *age < 18:
		price = price * 50 / 100
	case *age >= 65:
		price = price * 70 / 100
	}

	if *day == "sat" || *day == "sun" {
		price = price * 120 / 100
	}

	fmt.Println(price)
}
```

</details>

---

## Где ещё потренироваться

Сначала закрой 30 задач выше, потом добирай объём на внешних тренажёрах.

| Ресурс | Что решать | Зачем |
|---|---|---|
| [A Tour of Go](https://go.dev/tour/) | «Flow control statements», «More types» | короткие интерактивные упражнения на синтаксис |
| [Go by Example](https://gobyexample.com/) | For, If/Else, Switch, Variables, Constants | эталонные примеры, с которыми можно сверяться |
| [Exercism Go](https://exercism.org/tracks/go) | *Basics*, *Numbers, strings, booleans*, *Conditionals* (Blackjack) | автопроверка и менторские решения |
| [Codewars](https://www.codewars.com/) | Go, 8–7 kyu | много коротких задач на условия и арифметику |
| [LeetCode](https://leetcode.com/) | Easy, простые задачи на массивы и условия | привычка разбирать крайние случаи |

**Как считать прогресс:** 2–3 внешние задачи в день достаточно. Если задача решается
за 5 минут без ошибок — переходи к следующей теме, объём важнее.

## Чеклист темы

- [ ] Объявляю переменные через `var` и `:=`, понимаю разницу.
- [ ] Знаю нулевые значения для основных типов.
- [ ] Все конверсии делаю явно и понимаю, где теряется точность.
- [ ] Помню, что в Go нет `while` и тернарника.
- [ ] `if`/`else if`/`else` и порядок условий — под контролем.
- [ ] `switch` умею в трёх видах: по значению, без выражения, с init.
- [ ] Понимаю `&&`/`||` и короткое замыкание.
- [ ] Обрабатываю крайние случаи, а не только «счастливый путь».

## Самопроверка

1. Чем `var x int` отличается от `x := 0`?
2. Что напечатает `%v` для nil-среза, nil-мапы и nil-указателя?
3. Почему `float64(a+b)/3` и `(a+b)/3` дают разный результат?
4. Что произойдёт с `int8(127) + 1` и почему константа `128` в `int8` не компилируется?
5. Почему `b != 0 && a/b > 2` безопасно, а обратный порядок — нет?
6. Чем `switch x { case 1, 2: }` отличается от `fallthrough`?
7. Когда `switch { case cond: }` удобнее цепочки `if`?
8. Чем `byte` отличается от `rune` и почему `len("Я")` равно 2?

## Куда дальше

- Закрепить циклы и функции: день 3 в [week-01.md](week-01.md).
- Записать ловушки (переполнение, деление, `len` vs `rune`) в `../go-gotchas.md`.
- Отметить решённые задачи в `../learning-log.md`.
