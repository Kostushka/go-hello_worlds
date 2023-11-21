package main

import "time"
import "log"

func main() {
	bigSlowOp()
}

func bigSlowOp() {
	// здесь происходит вызов функции trace, которая возвращает анонимную функцию
	// вызов анонимной функции, благодаря defer, выполняется отложено
	defer trace("bigSlowOp")()
	// здесь отложено выполняется вызов функции trace, которая возвращает анонимную функцию
	// анонимная функция нигде не вызывается, потому мы не видим лог "выход из %s (%s)"
	// defer trace("bigSlowOp")
	
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("вход в %s", msg)
	return func() {
		log.Printf("выход из %s (%s)", msg, time.Since(start))
	}
}
