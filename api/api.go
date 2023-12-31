package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/moalf/passgen/rndstr"
)

const (
	defaultNumberOfPasswords = 1
	defaultMinPasswordLength = rndstr.MinPasswordLength
	maxNumberOfPasswords     = 10
	maxPasswordLength        = rndstr.MaxPasswordLength
)

type Data struct {
	Passwords []string `json:"passwords"`
	Details   struct {
		Specs struct {
			Number string `json:"number,omitempty"`
			Length string `json:"length,omitempty"`
		} `json:"specs,omitempty"`
		Error string `json:"error,omitempty"`
	} `json:"details,omitempty"`
}

func GetPassword(w http.ResponseWriter, req *http.Request) {
	var (
		requestedNumberOfPasswords int
		requestedPasswordLength    int
		err                        error
		data                       = Data{}
		NumberOfPasswords          = defaultNumberOfPasswords
		PasswordLength             = defaultMinPasswordLength
	)

	requestedNumberOfPasswords = NumberOfPasswords

	userRequestedPreferences := strings.Split(strings.Trim(req.URL.Path, "/"), "/")

	// no URL path, that means no user requested specs, keep defaults
	if len(userRequestedPreferences) == 1 && userRequestedPreferences[0] == "" {
		data.Details.Specs.Number = strconv.Itoa(defaultNumberOfPasswords)
		data.Details.Specs.Length = strconv.Itoa(PasswordLength)
	} else if len(userRequestedPreferences) == 1 {
		requestedNumberOfPasswords, err = strconv.Atoi(userRequestedPreferences[0])
		if err != nil {
			data.Details.Error = err.Error()
		}
		requestedPasswordLength = PasswordLength
	} else if len(userRequestedPreferences) == 2 {
		requestedNumberOfPasswords, err = strconv.Atoi(userRequestedPreferences[0])
		if err != nil {
			data.Details.Error = err.Error()
		}
		requestedPasswordLength, err = strconv.Atoi(userRequestedPreferences[1])
		if err != nil {
			data.Details.Error = err.Error()
		}
	}

	if requestedNumberOfPasswords <= maxNumberOfPasswords {
		NumberOfPasswords = requestedNumberOfPasswords
	}

	if requestedPasswordLength >= defaultMinPasswordLength && requestedPasswordLength <= maxPasswordLength {
		PasswordLength = requestedPasswordLength
	}

	data.Details.Specs.Length = strconv.Itoa(PasswordLength)
	data.Details.Specs.Number = strconv.Itoa(NumberOfPasswords)

	var passwords []string
	for i := 0; i < NumberOfPasswords; i++ {
		var password string
		for {
			password = rndstr.RandString(PasswordLength)
			if rndstr.IsComplex(password) {
				passwords = append(passwords, password)
				break
			}
		}
	}

	data.Passwords = passwords

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func Status(w http.ResponseWriter, req *http.Request) {
	var status map[string]string
	json.Unmarshal([]byte("{\"status\":\"OK\"}"), &status)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}
