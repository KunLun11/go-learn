package main

import (
	"fmt"
	"errors"
)

const MIN_EXPENSE = 0.01
var categories = [4]string{"Еда", "Транспорт", "Развелечения", "Прочее"}
var totals = [4]float64{}
var expenseCount int 

func addExpense(total *float64, amount float64) error {
	if amount < MIN_EXPENSE {
		return errors.New("сумма не может быть меньше расхода")
	}
	*total += amount
	return nil
}

func printCategories() {
	for i, cat := range categories {
		fmt.Println("Номер: ", i+1, "Название: ", cat, "Сумма:", totals[i])
	}
}

func totalSum() float64 {
	var sumV float64 = 0
	for i := 0; i < 4; i++ {
		sumV += totals[i]
	}
	return sumV
}

func biggestCategory() (string, float64, error) {
	var maxV = totals[0]
	var maxIdx = 0
	for i := 0;  i < 4; i++ {
		if totals[i] > maxV {
			maxV = totals[i]
			maxIdx = i
		}
	}
	if maxV == 0 {
		return "", 0, errors.New("нет расходов")
	}
	return categories[maxIdx], maxV, nil
}


func resetAll() {
	for i := 0; i < 4; i++ {
		totals[i] = 0
	}
	expenseCount = 0
}

func safeRun(name string, fn func()) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(name, r)
		}
	}()
	fn()
}


func main() {
	Loop:
		for {
			fmt.Println("----------------------------")
			fmt.Println("----------------------------")
			fmt.Println("1. Добавить расход")
			fmt.Println("2. Показать по категориям")
			fmt.Println("3. Показать итого")
			fmt.Println("4. Самая затратная категория")
			fmt.Println("5. Сбросить все")
			fmt.Println("0. Выход")
			fmt.Println("----------------------------")
			fmt.Println("----------------------------")
		
		var choice int
			if _, err := fmt.Scan(&choice); err != nil {
				fmt.Println("Неверный ввод")
				continue
			}
			switch choice {
			case 1: safeRun("расходы", func() {
				fmt.Println("Номер категории (1-4): ")
				var number_cat int
				fmt.Scan(&number_cat)
				fmt.Println("Введие сумму: ")
				var amount float64
				fmt.Scan(&amount)
				if number_cat > 4 || number_cat < 1 {
					fmt.Println("Нет такой категории")
					return
				}
				addExpense(&totals[number_cat-1], amount)
				expenseCount++
				fmt.Print("Добавлено")
			})
			case 2: safeRun("показ по категориям", func() {
				printCategories()
			})
			case 3: safeRun("итого", func() {
				fmt.Println("Сумма: ", totalSum())
			})
			case 4: safeRun("затратная категория", func() {
				category, maxV, err := biggestCategory()
				if err != nil {
					fmt.Println("Ошибка", err)
					return
				}
				fmt.Println("category | maxV")
				fmt.Printf(" %s | %.2f\n", category, maxV)
			})
			case 5: safeRun("сброс", func() {
				resetAll()
				fmt.Println("Все сброшено")
			})
			case 0: {
				fmt.Println("До свидания")
				break Loop
			}
			default: {
				fmt.Println("Нет такого пункта")
			} 
		}
}
}