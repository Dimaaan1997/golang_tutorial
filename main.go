package main

import (
	"fmt"
)

/*
func funcName (params ...) (returned values) {
	function's body
}
*/

/*
Базовый for
for init; condition; post statement {
// body
}

For как while
for condition {
// body
}

Бесконечный цикл:

for {
// body
}
*/

const (
	child_treshold int8 = 18
	min                 = 1
	max                 = 5
)

type OurString string
type OurInt int

type Person struct {
	Name string
	Age  int
}

func main() {
	//__________________________________________________________________________________________________________________________________________
	// первые 3 урока
	/*
		This comment for something
	*/
	// first, second := 1, 2

	// Greet()
	// personGreet("Dima")
	// FioGreet("Dima", "Batishchev")
	// sum_number := Sum(first, second)
	// fmt.Println(sum_number)
	// sum, mult := SumAndMult(first, second)
	// sum, mult := namedSumAndMult(first, second)
	// fmt.Println(sum, mult)

	// var num1 uint8 = 1
	// var num2 uint64 = 1

	// name := "valya"
	// hello := fmt.Sprintf("hello %s", name)
	// _ = hello
	// fmt.Println(hello)

	// fmt.Printf("Type: %T Value %v\n", name, name)

	// var num1 int64 = 15
	// var num2 int = 15

	// fmt.Println(num1 + int64(num2))

	// fmt.Println(unsafe.Sizeof(num1))
	// fmt.Println(unsafe.Sizeof(num2))

	//__________________________________________________________________________________________________________________________________________
	// УРОК №4 Функции (advanced):

	// var multiplier func(x, y int) int
	// multiplier = func(x, y int) int { return x * y }
	// fmt.Println(multiplier(first, second))

	// divider := func(x, y int) int { return x / y }
	// fmt.Println(divider(second, first))

	// cal := calculate(second, first, divider)
	// fmt.Println(cal)

	// createdDivider := createDivider(2)

	// fmt.Println(createdDivider(22))
	//__________________________________________________________________________________________________________________________________________

	// УРОК №4 Замыкания (advanced):

	// dollar := 30

	// getDollarValue := func() int {
	// 	return dollar
	// }

	// fmt.Println(getDollarValue())

	// dollar = 70

	// fmt.Println(getDollarValue())

	//__________________________________________________________________________________________________________________________________________
	// УРОК №5 Условный оператор (if/else). Логические операторы (advanced):

	// var age int8 = 19

	// if age < 18 {
	// 	fmt.Println("You are too young")
	// }

	// if isChild := isChild(age); isChild == true {
	// 	fmt.Println("You are too young")
	// 	// область видимости переменной, только в блоке логического оператора
	// 	fmt.Println(isChild)
	// } else {
	// 	fmt.Println("You are not too young")
	// 	fmt.Println(isChild)
	// }

	// if age >= 7 && age <= 18 {
	// 	fmt.Println("you are student in school.")
	// }

	// if !(age == 7 || age == 15) && (age <= 20) {
	// 	fmt.Println("your age is 7 or 15")
	// }

	//__________________________________________________________________________________________________________________________________________
	//Урок №6. Циклы (for, for как while). Инкремент. Декремент

	// i := 0

	// три способа увеличить переменнную на еденицу
	// i = i + 1
	// i += 1
	// i++
	// fmt.Println(i)

	// три способа вычесть еденицу
	// i = i - 1
	// i -= 1
	// i--
	// fmt.Println(i)

	/*
		Базовый for
		for init; condition; post statement {
		// body
	*/

	// i := 0 инициализация цикла (init)
	// i < 10 условие до которого будет выполнятся цикл (condition)
	// i++ то, что выполняется после цикла (post statement)

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// for i < 21 {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 		i += 1
	// 	} else {
	// 		i += 1
	// 		continue
	// 	}
	// }

	// myVar := 1
	// for {
	// 	fmt.Println(myVar)
	// 	if myVar == 20 {
	// 		break
	// 	}
	// 	myVar++
	// }
	// Label1:
	// 	for i := 1; i <= 20; i++ {
	// 		for j := 1; j <= 10; j++ {
	// 			fmt.Println("I:", i, "J:", j)
	// 			if i >= 10 {
	// 				continue Label1
	// 			}
	// 		}
	// 		fmt.Print("Hello")
	// 	}

	// Label2:
	// 	for i := 1; i <= 20; i++ {
	// 		for j := 1; j <= 10; j++ {
	// 			fmt.Println("I:", i, "J:", j)
	// 			if i >= 10 {
	// 				break Label2
	// 			}
	// 		}
	// 		fmt.Print("Hello")
	// 	}

	//__________________________________________________________________________________________________________________________________________
	//  Урок №8. Switch .. case

	// rand.Seed(time.Now().UnixNano())
	// value := rand.Intn(max-min) + min

	// if value == 1 {
	// 	fmt.Println("Number is one")
	// } else if 2 == value || value == 3 {
	// 	fmt.Println("2 or 3")
	// } else {
	// 	fmt.Println("Other number")
	// }

	// switch value {
	// case 1:
	// 	fmt.Println("one")
	// case 2, 3:
	// 	fmt.Println("2 or 3")
	// default:
	// 	fmt.Println("other")
	// }

	// switch num := rand.Intn(max-min) + min; num {
	// case 1:
	// 	fmt.Println("one")
	// case 2, 3:
	// 	fmt.Println("2 or 3")
	// 	fallthrough
	// default:
	// 	fmt.Println("other")
	// }

	// switch {
	// case value == 1:
	// 	fmt.Println("one")
	// case value == 2 || value == 3:
	// 	fmt.Println("2 or 3")
	// 	fallthrough
	// default:
	// 	fmt.Println("other")

	// }

	//__________________________________________________________________________________________________________________________________________
	// Урок №9. Указатели (pointers)

	// Что такое указатели (pointers):
	// Указатели - это тип данных, которые в качестве значения хранят адрес ячейки памяти значения,
	//  либо другого указателя (может быть nil)

	// var intPointer *int

	// fmt.Printf("1 %T %#v \n", intPointer, intPointer)

	// var a int64 = 7

	// fmt.Printf("2 %T %#v \n", a, a)

	// var pointerA *int64 = &a

	// // * переуказателем , называется разименованием указателя

	// fmt.Printf("3 %T %#v %#v\n", pointerA, pointerA, *pointerA)

	// var newPointer = new(float32)
	// fmt.Printf("4 %T %#v %#v\n", newPointer, newPointer, *newPointer)
	// *newPointer = 3
	// fmt.Printf("5 %T %#v %#v\n", newPointer, newPointer, *newPointer)

	// __________________________________________________________________________________________________________________________________________
	// Урок №10. Указатели (2). Usecases

	// num := 3
	// num_pointer := &num
	// fmt.Printf("1 %#v\n", num_pointer)

	// unname_pointer := *num_pointer
	// fmt.Printf("2 %#v\n", unname_pointer)

	// unname_pointer += 1
	// fmt.Printf("3 %#v\n", unname_pointer)
	// fmt.Printf("4 %#v\n", *num_pointer)

	// *num_pointer += 2
	// fmt.Printf("5 %#v\n", unname_pointer)
	// fmt.Printf("6 %#v\n", *num_pointer)

	// squarPointer(&num)
	// fmt.Printf("7 %#v\n", num)

	// var wallet1 *int

	// fmt.Printf("8 %#v\n", hasWallet(wallet1))

	// wallet2 := 0
	// fmt.Printf("9 %#v\n", hasWallet(&wallet2))

	// wallet3 := 100
	// fmt.Printf("10 %#v\n", hasWallet(&wallet3))

	//__________________________________________________________________________________________________________________________________________
	// Урок №11. Кастомные типы. Структуры(1)

	// Структура это сборка полей

	// var customString OurString
	// fmt.Printf("1 %T %#v \n\n", customString, customString)

	// customString = "Hello dudes"

	// fmt.Printf("2 %T %#v \n\n", customString, customString)

	// customInt := OurInt(5)

	// fmt.Printf("3 %T %#v \n\n", customInt, customInt)

	// fmt.Printf("4 %#v\n\n", int(customInt))

	// var Jonh Person

	// fmt.Printf("5 %T, %#v \n\n", Jonh, Jonh)

	// Jonh = Person{}
	// Jonh.Name = "Dima"
	// Jonh.Age = 28
	// fmt.Printf("6 %#v\n\n", Jonh)

	// Brad := Person{Name: "Brad", Age: 22}

	// fmt.Printf("7 %#v\n\n", Brad)

	// Vlad := Person{"Vlad", 40}
	// fmt.Printf("8 %#v\n\n", Vlad)

	// pVlad := &Vlad
	// fmt.Printf("9 %#v\n\n", (*pVlad).Age)
	// fmt.Printf("10 %#v\n\n", pVlad.Age)

	// pIvan := &Person{"Ivan", 90}

	// fmt.Printf("11 %#v\n\n", pIvan)

	// unnamedStrunct := struct {
	// 	Name, LastName, BirthDate string
	// 	Number1, Number2          int
	// 	BigNumber                 int64
	// }{Name: "Hello", LastName: "Gooo", BirthDate: fmt.Sprintf("%s", time.Now())}

	// fmt.Printf("12 %#v\n\n", unnamedStrunct)

	//__________________________________________________________________________________________________________________________________________
	// Урок №12. Методы
	// Метод - это функция, которая принадлежит определенному типу или указателю на тип (receiver)

	// 1. Создание метода
	// type T struct {}

	// func (receiver T) methodName() {}
	// func (receiver *T) methodName() {}
	// value receiver это метот, которй закреплен за типом или непосредтвенно объектом
	// pointer receiver это метод, который закреплен за указателем
	//value reciver просто называется так, потому го просто копирует значение

	// 1) Если нужно поменять значение, то pointer receiver
	// 2) Если копия значения большого размера, то лучше использовать укзаатели
	// У указателей фиксированный весь (строка с адресом)
	// 3) Много указателей заставляют GarbageCollector (GC) больше работать
	// 4) Значения пакета sync передавать только по указателю, иначе пропадает смысл в их использовании
	// пакет sync предназначен для безопасной работы в несоколько потоков
	
	// Ограничения для создания методов

	// Мы можем добавлять метод для типа T или *T при условиях:

	// 1. T находится в данном пакете
	// 2. T не является указателем
	// 3. T - конкретный, не интерфейсный тип
	// 4. T не является builtin(втроенным) типом
	// 5. T является объявленным типом

	// метод объявлен с маленькой буквы или с большой
	// Имя метода - модификатор доступа: private, Public 

	// definition()

	//__________________________________________________________________________________________________________________________________________
	// Урок №13. Интерфейсы
	
	// Интерфейс - спец тип в го, представляющий из себя набор сигнатур методов, которые нужно реализовать в его иплементации.



	//__________________________________________________________________________________________________________________________________________
	

}

/*
если мы хотим, чтобы функция была доступна из другого покета, то называть нужно с заглавной буквы,
*/

/**/

//__________________________________________________________________________________________________________________________________________
// Урок №12. Методы
type Square struct {
	Side int
}

func (s Square) Perimeter() {
	fmt.Printf("%T, %#v \n", s, s)
	fmt.Printf("Периметр фигуры: %d \n", s.Side*4)
}

func (s *Square) Scale(multiplier int) {
	fmt.Printf("%T, %#v \n", s, s)
	s.Side *= multiplier
	fmt.Printf("%T, %#v \n", s, s)
}
func (s Square) WrongScale(multiplier int) {
	fmt.Printf("1 %T, %#v \n", s, s)
	s.Side *= multiplier
	fmt.Printf("2 %T, %#v \n", s, s)
}

func definition() {
	square := Square{Side: 4}
	pSquare := &square
	// square2 := Square{Side: 2}
	square.Perimeter()
	// square2.Perimeter()
	pSquare.Scale(2)
	// го понимает, что метод относится к value и поэтому он разименовывает указатель и передает объект
	pSquare.Perimeter() // под капотом (*pSquare.Perimeter()) Square{Slide: 8}
	// и наоборот го получает указатель, когда понимает, что метод относится к указателю
	square.Scale(2)
	square.WrongScale(2)
	square.Perimeter()
}
//__________________________________________________________________________________________________________________________________________

type Runner interface {
	Run() string
}

type Swimmer interface {
	Swim() string
}


type Flyer interface {
	Fly() string
}





//__________________________________________________________________________________________________________________________________________

func Greet() {
	fmt.Println("Hello guys")
}

func personGreet(name string) {
	fmt.Printf("Hello %s\n", name)
}

// компонуем параметры по типу
func FioGreet(name, surname string) {
	fmt.Printf("Hi %s %s\n", name, surname)
}

func Sum(first, second int) int {
	sum := first + second
	return sum
}

func SumAndMult(first, second int) (int, int) {
	return first + second, first * second
}

func namedSumAndMult(first, second int) (sum int64, mult int64) {
	sum = int64(first + second)
	mult = int64(first * second)
	// return // или return sum, mult
	return sum, mult
}

func calculate(x, y int, action func(x, y int) int) int {
	return action(x, y)
}

func createDivider(divider int) func(x int) int {
	dividerFunc := func(x int) int {
		return x / divider
	}
	return dividerFunc
}

func isChild(age int8) bool {
	return age < child_treshold
}

func square(number int) int {
	number *= number
	return number
}

func squarPointer(number *int) {
	*number *= *number
}

func hasWallet(money *int) bool {
	return money != nil
}
