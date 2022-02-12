package main

func main() {
	name := "Alex"
	pName := &name

	var emptyPointer *int
	print(emptyPointer)

	result := randomName(pName)
	print(result)
	print(*pName)
}

func randomName(name *string) string {
	newName := *name
	newName += "Surname"

	return newName
}
