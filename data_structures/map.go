package data_structures

// this is a map also known ad dict
func Map_Dict() {
	var my_map = make(map[string]int) // "key" -> string : value -> int

	// Let's do some operations
	my_map["one"] = 1
	my_map["two"] = 2
	my_map["three"] = 3

	// Print the map
	for key, value := range my_map {
		println("Key:", key, "Value:", value)
	}

	// Delete the keys
	delete(my_map, "two")
	println("After deletion:")
	for key, value := range my_map {
		println("Key:", key, "Value:", value)
	}

	// Find a value by key
	value, found := my_map["one"]

	if found {
		println("Value for key 'one':", value)
	} else {
		println("Key 'one' not found")
	}
}
