package customError

import (
	"context"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"inverse.so/structure"
)

type JSONError struct {
	Message    string                 `json:"message"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

func ErrToGraphQLError(err structure.ErrorCode, errMessage string, ctx context.Context) error {
	message := ProperTitle(strings.ToLower(errMessage))

	return &gqlerror.Error{
		Message:   message,
		Path:      graphql.GetPath(ctx),
		Locations: []gqlerror.Location{},
		Extensions: map[string]interface{}{
			"code":      err,
			"timestamp": time.Now().Format("2006-01-02T15:04:05.000Z"),
		},
	}
}

func ProperTitle(input string) string {
	words := strings.Split(input, " ")
	smallwords := " a an and as at in is of on or so the to "
	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") && word != string(words[0]) {
			words[index] = word
		} else {
			words[index] = cases.Title(language.English).String(word)
		}
	}
	modifiedInput := strings.Join(words, " ")
	modifiedInput = strings.ReplaceAll(modifiedInput, "'T ", "'t ")
	modifiedInput = strings.ReplaceAll(modifiedInput, "'S ", "'s ")
	return modifiedInput
}

func ErrToJsonError(err structure.ErrorCode, errMessage string, ctx context.Context) JSONError {
	message := ProperTitle(strings.ToLower(errMessage))

	return JSONError{
		Message: message,
		Extensions: map[string]interface{}{
			"code":      err,
			"timestamp": time.Now().Format("2006-01-02T15:04:05.000Z"),
		},
	}
}
