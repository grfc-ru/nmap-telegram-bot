package main

func RemoveChar(word string) (result string) {
	if len(word) > 2 {
		result = word[0 : len(word)-1]
	}
	return
}
