package go_basics

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"reflect"
	"sort"
	"time"
	"unicode/utf8"
)

type Weekday int

const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func HelloWorld() {
	fmt.Println("Hello, Go!")
}

func FizzBuzz() {
	for i := 1; i < 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}

func TestingErrorProcessing() {
	result, err := Division(1, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func Division(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("кто-то передал запрещенный символ")
	}

	return a / b, nil
}

func DoPanic() {
	panic("Some text")
}

func DeferTest() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
}

func TestingInteger() {
	var a, b int

	a, b = 32, 64
	b, a = a, b

	fmt.Println(a + b)
}

func TestingFloatingPointNumbers() {
	var a, b float32
	var c, d complex64

	a, b = 32.01, 64.0
	c, d = 32.01, 64.0

	fmt.Println(a + b)
	fmt.Println(c + d)
}

func TestingBoolean() {
	var a bool

	a = false

	fmt.Println(a)
}

func TestingString() {
	var a string

	a = "Привет"

	fmt.Println(len(a))
	fmt.Println(a[5]) //Вернет номер байта символа из таблицы ASCII

	fmt.Println(utf8.RuneCountInString(a))
}

func TestingBackticks() {
	var a, b string

	a = "Какая-то строка\nна несколько строк"
	b = `Какая-то строка
на несколько строк`

	fmt.Println(a)
	fmt.Println(b)
}

func TestingUserTypes() {
	type Name string
	type Fruit string

	var name Name
	var fruit Fruit
	var str string

	name = "Adam"
	fruit = Fruit(name)
	str = "Hello world!"

	fmt.Println(name)
	fmt.Println(fruit)
	fmt.Println(string(str[0]))
}

func TestingAliases() {
	type String = string

	var a String
	var b string

	a = "1"
	b = a

	fmt.Println(a == b)
}

func TestingAlternativeWaysOfVariablesDeclaration() {
	var (
		height, length int8 = 52, 52
	)

	someValue := uint16(1)
	ver := "v0.0.1"
	id := 0
	pi := 3.1415

	fmt.Println(height, length, someValue)
	fmt.Println("ver =", ver, "id =", id, "pi =", pi)
}

func TestingUntypedConstants() {
	const (
		number         = 100
		floatingNumber = 5.01
		andOneMore
	)

	var (
		integer         int64
		floating        float64
		anotherFloating float32
	)

	integer, floating = number, number
	anotherFloating = floatingNumber

	fmt.Println(integer, floating, 5+anotherFloating+andOneMore)
}

func TestingTypedConstants() {
	const (
		constantName int = 10
	)

	var variableName = constantName

	fmt.Println(variableName)
}

func TestingEnums() {
	const (
		White = 1
		Black = 2
		Gray  = 3
	)

	const (
		Zero = iota
		One
		Two
		Three
	)

	fmt.Println(White != Gray, Zero, One, Two, Three)
}

func UserTypesInConstants() {
	today := Saturday

	fmt.Println("Today: ", today, "Tomorrow: ", NextDay(today))
}

func NextDay(day Weekday) Weekday {
	return (day % 7) + 1
}

func TestingLiterals() {
	var (
		a = 1000
		b = 1_000
		c = 01750          // x8
		d = 0x3e8          // x16
		e = 0b001111101000 //binary
	)

	fmt.Println(a, b, c, d, e)
}

func TestingIota() {
	const (
		one = 2*iota + 1
		three
		five
		seven
		nine
		eleven
	)

	fmt.Println(one, three, five, seven, nine, eleven)
}

func TestingConditionalOperators() {
	var a int

	if i := 1; i == 1 {
		fmt.Println("Testing")
	}

	switch b := a % 5; {
	case b == 0:
		fmt.Println("Equal 5")
	case b > 0:
		fmt.Println("Two")
	case b < 0:
		fmt.Println("Three")
	default:
		if a < 1 {
			break
		}
		fmt.Println("Idk that number")
	}

	var (
		birthYear = 1996
		message   string
	)

	switch {
	case birthYear >= 1946 && birthYear <= 1964:
		message = "Привет, бумер!"
	case birthYear >= 1965 && birthYear <= 1980:
		message = "Привет, представитель X!"
	case birthYear >= 1981 && birthYear <= 1996:
		message = "Привет, миллениал!"
	case birthYear >= 1997 && birthYear <= 2012:
		message = "Привет, зумер!"
	case birthYear >= 2013:
		message = "Привет, альфа!"
	default:
		message = "Привет, динозавр!"
	}

	fmt.Println(message)
}

func TestingCycles() {
	for {
		fmt.Println("Infinity `while` cycle")

		break
	}

	for a, b := 1, 2; a < 2 && b < 4; a, b = a+1, b+2 {
		fmt.Println("Regular 'for' cycle")
	}

	i := 1

	for i < 2 {
		fmt.Println("`While` cycle with condition")

		i++
	}

	array := [1]int{1}

	for value, index := range array {
		fmt.Println("Range (foreach) cycle", value, index)
	}

	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(i)
		}
	}

}

func Pointers() {
	var (
		a             = 1
		intPointer    *int
		b             = "string"
		stringPointer *string
	)

	intPointer = &a
	stringPointer = &b

	fmt.Println(intPointer)
	fmt.Println(*intPointer)
	fmt.Println(stringPointer)
	fmt.Println(*stringPointer)
}

func LiteralPointer() {
	type A struct {
		intField int
	}

	var (
		a = 1
		b *int
		c **int
	)

	b = &a
	c = &b

	//p := &A{intField: 1}
	p := new(A)

	fmt.Println(p)
	fmt.Println(*p)

	fmt.Println(c)
	fmt.Println(*c)
	fmt.Println(**c)
}

func StructPointer() {
	type Person struct {
		Name        string
		Age         int
		lastVisited time.Time
	}

	updateLastVisitedField := func(person *Person) {
		person.lastVisited = time.Now()
	}

	person := Person{
		Name: "Name",
		Age:  12,
	}

	updateLastVisitedField(&person)

	fmt.Println(person)
}

func UserInputCounter() {
	// Получаем читателя пользовательского ввода
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interaction counter")

	f := func(counter *int) {
		*counter++
	}

	cnt := 0
	for {
		fmt.Print("-> ")
		// Считываем введённую пользователем строку. Программа ждёт, пока пользователь введёт строку
		_, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		f(&cnt)

		fmt.Printf("User input %d lines\n", cnt)
	}
}

func Arrays() {
	var (
		lastWeekTemp = [7]int{1, 2, 3}
		dynamicArray = [...]int{10: 1, 11: 2}
	)

	fmt.Println(lastWeekTemp, dynamicArray)
}

func MultidimensionalArray() {
	var (
		a            [2][2]int
		lastWeekTemp = [7]int{1, 2, 3, 4, 5, 6, 7}
		sum          = 0
	)

	for i := 0; i < len(lastWeekTemp); i++ {
		sum += lastWeekTemp[i]
	}

	for _, temp := range &lastWeekTemp {
		sum += temp
	}

	for i := range &lastWeekTemp {
		lastWeekTemp[i] = 0
	}

	fmt.Println(a, lastWeekTemp, sum/len(lastWeekTemp))
}

func Slices() {
	var slice = []int{1, 2, 3, 4, 5, 6}

	newSlice := make([]int, 6)
	fullySlicedSlice := slice[:]
	firstHalfSlice := slice[:4]
	lastHalfSlice := slice[3:]
	newSliceWithNewBasisArray := slice[3:4:6]

	slice[3] = 7

	fmt.Println(slice, newSlice, fullySlicedSlice, firstHalfSlice, lastHalfSlice, newSliceWithNewBasisArray)
}

func ChangingSliceSize() {
	var (
		s = []int{1, 2, 3}
		a = []int{4, 5, 6, 7}
		b = a[2:3]
	)

	b[0] = 12

	b = append(b, 9)

	fmt.Println(s, a, b, len(b), cap(b))

	b = append(b, 10)

	b[0] = 13

	s = s[:len(s)-1]

	fmt.Println(s, a, b, len(b), cap(b))

	ms := append(s, a...)
	ms = append(ms, b...)

	fmt.Println(ms)
}

func WorkingWithSlicesInFunc() {
	sn := []int{1, 3, 5, 4, 2}
	sb := []byte("\n\t Some string")

	sort.Ints(sn)

	fmt.Println(sn) //Изменяет базовый слайл потому что не добавляются/удаляются элементы
	fmt.Println(sb)
	fmt.Println(bytes.TrimSpace(sb)) //Не изменяет и возвращает новый т.к. было изменено количество элементов
}

func ChangingSliceSizeInFunc() {
	s := []int{1, 2, 3, 4}

	fn := func(s *[]int) {
		*s = append(*s, 5)
	}

	fn(&s)
	fmt.Println(s)
}

func OperationsOnSlices() {
	var (
		dest  []int
		dest2 = make([]int, 3)
		dest3 = make([]int, 5)
		array = [4]int{1, 2, 3, 4}
		src   = array[:]
		dest4 = make([]int, 5)
	)

	copy(dest, src)
	copy(dest2, src)
	copy(dest3, src)
	copy(dest4, src)

	//Удаление последнего элемента слайса
	if len(dest3) != 0 {
		dest3 = dest3[:len(dest3)-1]
	}
	if len(dest4) != 0 {
		dest4 = dest4[:len(dest4)-1]
	}

	//Удаление первого элемента слайса
	if len(dest2) != 0 {
		dest2 = dest2[1:]
	}

	//Удаление по индексу
	i := 1
	if len(src) > 0 && i < len(src) {
		src = append(src[:i], src[i+1:]...)
	}

	//Сравнение двух слайсов
	fmt.Println(reflect.DeepEqual(dest3, dest4))
	fmt.Println(reflect.DeepEqual(dest3, dest2))

	fmt.Println(src, dest, dest2, dest3)

	for _, value := range src {
		fmt.Println(value)
	}
}

func TestTaskWithSlices() {
	s := make([]int, 100)
	b := []byte("ABCDFG")

	fmt.Println(s, string(b))

	for i := 0; i < 100; i++ {
		s[i] = i + 1
	}

	fmt.Println(s)

	a := 10
	if len(s) > 0 && a < len(s) {
		s = append(s[:a], s[len(s)-a:]...)
	}

	fmt.Println(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}

	fmt.Println(s, string(b))
}
