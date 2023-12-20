package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Data struct {
	Passwords []string `json:"passwords"`
}

const (
	lower    = `abcdefghijklmnopqrstuvwxyz`
	upper    = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	numbers  = `0123456789`
	specials = `!@#$%^*()_-+={}[]|:;,.?/`
	port     = "8080"
)

var (
	minPasswordLength    = 12
	maxNumberOfPasswords = 10
	numOfPasswords       = 1
)

func main() {

	http.HandleFunc("/", getPassword)

	fmt.Printf("Starting server at port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func isComplex(s string) bool {
	for idx, c := range specials {
		if strings.Contains(s, string(c)) {
			break
		}

		if idx == len(s) {
			return false
		}
	}

	rules := []string{".{12,}", "[a-z]", "[A-Z]", "[0-9]"}
	for _, rule := range rules {
		r := regexp.MustCompile(rule)
		if !r.MatchString(s) {
			return false
		}
	}

	return true
}

func RandString(n int) string {
	sb := strings.Builder{}
	all := []string{lower, upper, numbers, specials}

	for i, g := 0, 0; i < n; i, g = i+1, rand.Intn(4) {
		x := rand.Intn(len(all[g]))
		sb.WriteByte(all[g][x])
	}

	return sb.String()
}

func getPassword(w http.ResponseWriter, req *http.Request) {
	data := Data{}
	numOfPasswords = 1

	r := regexp.MustCompile(`\/[0-9]{1}$`)
	if r.MatchString(req.URL.Path) {
		requestedNumOfPasswords, err := strconv.Atoi(strings.Split(strings.TrimSpace(req.URL.Path), "/")[1])
		if err != nil {
			return
		}
		if requestedNumOfPasswords <= maxNumberOfPasswords {
			numOfPasswords = requestedNumOfPasswords
		}
	}

	// Fix when to return empty
	// if req.URL.Path != "/" {
	// 	json.NewEncoder(w).Encode(data)
	// 	return
	// }

	// TODO: Iterate over numOfPasswords if greater than 1
	// TODO: Collect all passwords in an []string if numOfPasswords is greater than 1

	var passwords []string
	for i := 0; i < numOfPasswords; i++ {
		var password string
		for {
			password = RandString(minPasswordLength)
			if isComplex(password) {
				passwords = append(passwords, password)
				break
			}
		}
	}

	data.Passwords = passwords

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
