package datastruct

func removeArrayByIndex(slice []any, index int) []any {
	// fmt.Println(slice[:index], slice[index+1:])
	return append(slice[:index], slice[index+1:]...)
}
