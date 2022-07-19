package tokenservice

import (
	"errors"
	"time"
)

var Tokens []Token

func CreateToken(userID int64) (Token, error) {
	token := Token{
		ID:          1,
		UserID:      userID,
		Token:       "023a246d-274d-4a84-8a9d-dd12b406d6c7",
		IsValid:     true,
		CreatedDate: time.Now().UTC(),
	}

	Tokens = append(Tokens, token)

	return token, nil
}

func GetTokenByToken(token string) (Token, error) {
	for _, t := range Tokens {
		if t.Token == token {
			return t, nil
		}
	}

	return Token{}, errors.New("Unable to find token")
}
