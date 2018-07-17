package goroutines

func Process(input chan string) chan string {
	output := make(chan string)
	done := make(chan bool)

	go func() {
		output <- "(" + <-input + ")"
		done <- true
	}()

	go func() {
		<-done
		close(output)
	}()
	return output
}
