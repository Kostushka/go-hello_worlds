package db

import "sync"

// условная бд с методами
type DB struct {
	db map[string]string
	mu sync.RWMutex
}

// добавить в бд
func (p *DB) Set(key, value string) {
	// во время записи в хеш-таблицу, в созданную структуру записывается указатель на весь связанный список, сама же структура записывается в массив по индексу
	// при одновременной записи структур по одному индексу можно потерять какую-либо из новых структур
	// блокировка для всех: никому больше нельзя читать и писать
	p.mu.Lock()
	p.db[key] = value
	p.mu.Unlock()
}

// проверить, записано ли уже в бд значение по этому ключу
func (p *DB) IsExist(key string) bool {
	// блокировка на чтение: никому нельзя писать, но можно читать одновременно нескольким клиентам
	p.mu.RLock()
	defer p.mu.RUnlock()
	
	if _, ok := p.db[key]; ok {
		return true
	}
	return false
}

// получить значение по ключу
func (p *DB) Get(key string) string {
	// блокировка на чтение: никому нельзя писать, но можно читать одновременно нескольким клиентам
	p.mu.RLock()
	defer p.mu.RUnlock()
	
	return p.db[key]
}

// функция-конструктор
func NewDB() *DB {
	return &DB{
		db: make(map[string]string, 10),
	}
}
