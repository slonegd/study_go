package main

import "fmt"

// определим пару интерфейсов

 // Comparable сравнивает пару структур
type Comparable interface {
	Less(interface{}) bool
}

// Doubler удваевает что-то в структуре
type Doubler interface {
	Double()
}

// GetMax возвращает максимальное из двух сравниваемых структур
func GetMax (_1, _2 Comparable) Comparable {
	if (_1).Less(_2) {
		return _2
	} else {
		return _1
	}
}

// ComparableAndDoubler объединение двух интерфейсов
type ComparableAndDoubler interface {
	Comparable
	Doubler
}

// FindMaxAndDouble ...
func FindMaxAndDouble (_1, _2 ComparableAndDoubler) {
	GetMax(_1,_2).(Doubler).Double()
}




// S некая структура, которая реализует интерфейсы
type S struct{
	a int
}

// Less метод по значению, поскольку в самой структуре ничего не изменяется
func (s S) Less(other interface{}) bool {
	var aOther int
	// передать интерфейс можно как по значению, так и по указателю (sic)
	// а узнать о типе можно только в switch
	switch other.(type) {
	case S:
		aOther = other.(S).a;
	case *S:
		aOther = other.(*S).a;
	}
	return s.a < aOther
}

// Double метод по указателю, так как в нём изменяются поля структуры
// нельзя вызвать этот метод, если интерфейс передан по значению
func (s* S) Double() {
	s.a *= 2
}

func main() {

	a1 := S{1}
	a2 := S{2}

	// если передавать по значению, то внутри ничего изменять нельзя
	// , а возвращается копия структуры, которую ещё и привести надо
	// , чтобы был доступ к полям
	maxCopy := GetMax(a1, a2).(S)

	// если передавать по указателю, то и возвращается по указателю
	maxPointer := GetMax(&a1, &a2).(*S)

	fmt.Printf ("a2.a before maxPointer: %v\n", a2.a)
	maxPointer.a = 10 // поле меняется как в maxPointer, так и в a2
	maxCopy.a = 11
	fmt.Printf ("maxPointer.a: %v\n", maxPointer.a)
	fmt.Printf ("a2.a after maxPointer: %v\n", a2.a)
	fmt.Printf ("maxCopy.a: %v\n", maxCopy.a)
	a1.Double()
	

	// в эту функцию нельзя передать по значению, так как она меняет поля структуры
	FindMaxAndDouble(&a1,&a2)
	fmt.Printf ("FindMaxAndDouble: %v\n", a2.a)
}
