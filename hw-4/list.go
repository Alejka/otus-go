package hw4

/*
List // тип контейнер
Len() // длинна списка
First() // первый Item
Last() // последний Item
PushFront(v interface{}) // добавить значение в начало
PushBack(v interface{}) // добавить значение в конец
Remove(i Item) // удалить элемент

Item // элемент списка
Value() interface{} // возвращает значение
Next() *Item // следующий Item
Prev() *Item // предыдущий
*/

type Item struct {
    value interface{}
    prev *Item
    next *Item
}

func (i Item) Value() interface{} {
}

func (i Item) Next() *Item {
}

func (i Item) Prev() *Item {
}

type List struct {
    data []Item
    len int
}

func (l List) Len() int {
}

func (l List) First() *Item {
}

func (l List) Last() *Item {
}

func (l *List) PushFront(v interface{}) {
}

func (l *List) PushBack(v interface{}) {
}

func (l *List) Remove(i Item) {
}
