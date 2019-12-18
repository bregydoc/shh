package main

func main() {
	wizard, err := NewDefaultWizard()
	if err != nil {
		panic(err)
	}

	errChan := wizard.Run()

	<- errChan
}
