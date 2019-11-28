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
    return i.value
}

func (i Item) Next() *Item {
    return i.next
}

func (i Item) Prev() *Item {
    return i.prev
}

type List struct {
    len int
    first *Item
    last *Item
}

func (l List) Len() int {
    return l.len
}

func (l List) First() *Item {
    return l.first
}

func (l List) Last() *Item {
    return l.last
}

func (l *List) PushFront(v interface{}) {
    i := Item{
        value: v,
    }
    
    if l.first == nil {
        l.last = &i
    } else {
        l.first.prev = &i
    }
    
    l.first = &i
    
    l.len++
}

func (l *List) PushBack(v interface{}) {
    i := Item{
        value: v,
    }
    
    if l.last == nil {
        l.first = &i
    } else {
        l.last.next = &i
    }
    
    l.last = &i
    
    l.len++
}

//func (l *List) Remove(i Item) {
//}
