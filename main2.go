package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func isPowerOfTwo(n int) bool {
	a := 1
	for i := 0; i <= n; i++ {
		if a == n {
			return true
		} else {
			a *= 2
		}
	}
	return false
}

func main() {
	solveProblemHanlder := http.HandlerFunc(solveProblem)
	http.Handle("/", solveProblemHanlder)
	http.ListenAndServe(":8080", nil)
}

func solveProblem(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("request.Form::")
	for key, value := range r.Form {
		fmt.Printf("Key:%s, Value:%s\n", key, value)
	}
	w.Write([]byte(fmt.Sprintf("%v", r.FormValue("age"))))
	age, _ := strconv.Atoi(r.FormValue("age"))
	if isPowerOfTwo(age) == true {
		w.Write([]byte(fmt.Sprintf("\nAge is power of two")))
	} else {
		w.Write([]byte(fmt.Sprintf("\nAge is not power of two")))
	}
	w.WriteHeader(200)
	return
}
