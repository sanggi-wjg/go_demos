package sorts

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
	loc  *location
}

type location struct {
	nation string
	city   string
}

func newPerson(name string, age int, nation, city string) *person {
	return &person{
		name: name,
		age:  age,
		loc: &location{
			nation: nation,
			city:   city,
		},
	}
}

func SortStructMain() {
	// Integer 정렬
	numbers := []int{1, 3, 4, 2, 5}
	slice := sort.IntSlice(numbers)
	sort.Sort(slice)
	fmt.Println(numbers)

	// Integer 역정렬
	sort.Sort(sort.Reverse(slice))
	fmt.Println(numbers)

	// String 정렬
	strings := []string{"apple", "pineapple", "banana", "kiwi", "grape"}
	sort.Sort(sort.StringSlice(strings))
	fmt.Println(strings)

	// Struct 정렬
	persons := make([]person, 0)
	persons = append(persons, *newPerson("User1", 10, "Korea", "Seoul"))
	persons = append(persons, *newPerson("User2", 30, "USA", "LA"))
	persons = append(persons, *newPerson("User3", 40, "Japan", "Tokyo"))
	persons = append(persons, *newPerson("User4", 20, "Austraila", "Sideny"))
	persons = append(persons, *newPerson("User5", 50, "China", "SangHai"))
	fmt.Println(persons)

	// Person age 정렬
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].age < persons[j].age
	})
	fmt.Println(persons)

	// Person loc.nation 정렬
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].loc.nation < persons[j].loc.nation
	})
	fmt.Println(persons)
}

/*
[1 2 3 4 5]
[5 4 3 2 1]
[apple banana grape kiwi pineapple]
[{User1 10 {Korea Seoul}} {User2 30 {USA LA}} {User3 40 {Japan Tokyo}} {User4 20 {Austraila Sideny}} {User5 50 {China SangHai}}]
[{User1 10 {Korea Seoul}} {User4 20 {Austraila Sideny}} {User2 30 {USA LA}} {User3 40 {Japan Tokyo}} {User5 50 {China SangHai}}]
[{User4 20 {Austraila Sideny}} {User5 50 {China SangHai}} {User3 40 {Japan Tokyo}} {User1 10 {Korea Seoul}} {User2 30 {USA LA}}]
*/
