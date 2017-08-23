package main

func main() {

	http.HandleFunc("/", errorHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}
}
