package main

func main() {
	setupErr := SetupServer()

	if setupErr != nil {
		print(setupErr.Error())
	}
}
