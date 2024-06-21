package main

func main() {
	_, err := func1()

	if err != nil {
		println(err)
		return
	}

	func2()
}

func func1() (interface{}, error) {
	err := func3()

	if err != nil {
		println(err)
		return nil, err
	}

	return nil, nil
}

func func2() {
	err := func3()
	if err != nil {
		println(err)
		return
	}
}

func func3() error {
	return nil
}
