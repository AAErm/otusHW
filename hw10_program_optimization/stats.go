package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"

	"github.com/mailru/easyjson"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	scanner := bufio.NewScanner(r)
	stat := make(DomainStat)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		if err := appendStat(domain, bytes, &stat); err != nil {
			return stat, err
		}
	}

	if err := scanner.Err(); err != nil {
		return stat, err
	}

	return stat, nil
}

func appendStat(domain string, bytes []byte, stat *DomainStat) error {
	user := &User{}

	if err := easyjson.Unmarshal(bytes, user); err != nil {
		return err
	}
	email := user.Email

	if strings.HasSuffix(email, "."+domain) {
		(*stat)[strings.ToLower(strings.SplitN(email, "@", 2)[1])]++
	}

	return nil
}
