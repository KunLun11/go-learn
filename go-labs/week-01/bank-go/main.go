package main

import (
	"errors"
	"fmt"
)

const MAX_ACCOUNTS = 10
const MIN_DEPOSIT = 1.0
const MAX_WITHDRAW = 10000.0
const MIN_BALANCE = 0.0
const START_BALANCE = 0

var ids [MAX_ACCOUNTS]int 
var names [MAX_ACCOUNTS]string 
var balances [MAX_ACCOUNTS]float64
var count_of_score int 
var nextID int = 1

func createAccount(name string) (id int, err error) {
	if count_of_score >= MAX_ACCOUNTS {
		return 0, errors.New("нет свободных мест")
	}
	if len(name) == 0 {
		return 0, errors.New("имя не может быть пустым")
	}
	pos := count_of_score
	newID := nextID
	ids[pos] = newID
	names[pos] = name
	balances[pos] = START_BALANCE
	count_of_score++
	nextID++
	return newID, nil
}

// пополнение
func deposit(balance *float64, amount float64) error {
	if amount < MIN_DEPOSIT {
		return  errors.New("сумма не должна быть меньше 1")
	}
	*balance += amount
	return nil
}

// снятие
func withdraw(balance *float64, amount float64) error {
	if amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}
	if amount > MAX_WITHDRAW {
		return errors.New("превышен лимит на одну операцию")
	}
	if amount > *balance {
		return  errors.New("недостаточно средств")
	}
	*balance -= amount
	return nil
}

func findAccount(id int) int {
	for i := 0; i < count_of_score; i++{
		if ids[i] == id {
			return i
		}		
	}
	return -1
}

// перевод
func transfer(from, to *float64, amount float64) error {
	if from == to {
		return errors.New("перевод на один и тот же кошелек недоступен")
	}
	if amount < MIN_DEPOSIT {
		return errors.New("сумма не может быть меньше 1")
	}
	err1 := withdraw(from, amount)
	if err1 != nil {
		return err1
	}
	err2 := deposit(to, amount) 
	if err2 != nil {
		return err2
	}
	return nil
}

func printAccounts() {
	if count_of_score == 0 {
		fmt.Println("счетов пока нет")
		return
	}
	fmt.Println("ID | Имя | Баланс")
	for i := 0; i < count_of_score; i++{
		fmt.Printf("%d | %s | %.2f\n", ids[i], names[i], balances[i])
	}
}

func stats() (minV, maxV, sumV float64, err error) {
	if count_of_score == 0 {
		return 0, 0, 0, errors.New("нет счетов для статистики")
	}
	minV = balances[0]
	maxV = balances[0]
	for i := 0; i < count_of_score; i++ {
		if balances[i] < minV {
			minV = balances[i]
		}
		if balances[i] > maxV {
			maxV = balances[i]
		}
		sumV += balances[i]
	}
	return minV, maxV, sumV, nil
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
			fmt.Println("1. Создать счет")
			fmt.Println("2. Пополнить")
			fmt.Println("3. Снять")
			fmt.Println("4. Сделать перевод")
			fmt.Println("5. Показать все счета")
			fmt.Println("6. Просмотреть статистику")
			fmt.Println("----------------------------")
			fmt.Println("----------------------------")

			var choice int
			if _, err := fmt.Scan(&choice); err != nil {
				fmt.Println("неверный ввод")
				continue
			}
			switch choice {
			case 0: 
				fmt.Println("До свидания")
				break Loop
			case 1: safeRun("создание счета", func() {
				fmt.Print("Имя:")
				var name string
				fmt.Scan(&name)
				id, err := createAccount(name)
				if err != nil {
					fmt.Println("ошибка", err)
					return
				}
				fmt.Println("создан счет с id=", id)
			})
			case 2: safeRun("пополнение", func() {
				fmt.Print("ID счета: ")
				var id int
				fmt.Scan(&id)
				fmt.Print("Сумма: ")
				var amount float64
				fmt.Scan(&amount)
				pos := findAccount(id)
				if pos == -1 {
					fmt.Println("счет не найден")
					return
				}
				err := deposit(&balances[pos], amount)
				if err != nil {
					fmt.Println("ошибка:", err)
					return
				}
				fmt.Println("пополнено, новый баланс: ", balances[pos])
			})
			case 3: safeRun("снять", func() {
				fmt.Print("ID счета: ")
				var id int
				fmt.Scan(&id)
				fmt.Print("Сумма: ")
				var amount float64
				fmt.Scan(&amount)
				pos := findAccount(id)
				if pos == -1 {
					fmt.Println("счет не найден")
					return
				}
				err := withdraw(&balances[pos], amount)
				if err != nil {
					fmt.Println("ошибка:", err)
					return
				} 
				fmt.Println("сняли, новый баланс: ", balances[pos])
			})
			case 4: safeRun("перевести", func() {
				fmt.Print("ID счета отправителя: ")
				var fromID int
				fmt.Scan(&fromID)
				fmt.Print("ID счета получателя: ")
				var toID int
				fmt.Scan(&toID)
				fmt.Print("Сумма: ")
				var amount float64
				fmt.Scan(&amount)
				fromPos := findAccount(fromID)
				toPos := findAccount(toID)
				if fromPos == -1 {
					fmt.Println("счет отправителя не найден")
					return
				}
				if toPos == -1 {
					fmt.Println("счет получателя не найден")
					return
				}
				err := transfer(&balances[fromPos], &balances[toPos], amount)
				if err != nil {
					fmt.Println("ошибка:", err)
					return 
				} 
				fmt.Println("перевели, баланс отправителя:", balances[fromPos])
				fmt.Println("баланс получателя:", balances[toPos])
			})
			case 5: safeRun("показать счета", func() {
				// panic("тест")
				printAccounts()
			})
			case 6: safeRun("статистика", func() {
				minV, maxV, sumV, err := stats()
				if err != nil {
					fmt.Println("ошибка", err)
					return
				} 
				avgV := sumV / float64(count_of_score)
				fmt.Println("minV | maxV | sumV | avgV")
				fmt.Printf("%.2f | %.2f | %.2f | %.2f\n", minV, maxV, sumV, avgV)
			})	
			default:
				fmt.Println("нет такого пункта")
			}
		}
}