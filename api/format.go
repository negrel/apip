package api

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
)

// Response define common response methods.
type Response interface {
	ToJSON() string
	ToTXT() string
	ToXML() string
	string() string
}

var _ Response = response{}

type response map[string]string

// ToJSON convert the response to JSON.
func (r response) ToJSON() string {
	json, err := json.Marshal(r)

	if err != nil {
		log.Error().Msg("The following JSON encoding failed :")
		log.Error().Msg(r.string())
	}

	return string(json)
}

// ToTXT convert the response to plain text.
func (r response) ToTXT() string {
	return r.string()
}

// ToXML convert the response to XML.
func (r response) ToXML() string {
	var str strings.Builder

	str.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	str.WriteString("<query>")

	for key, value := range r {
		str.WriteString(fmt.Sprintf("<%v>%v</%v>", key, value, key))
	}

	str.WriteString("</query>")

	return str.String()
}

func (r response) string() string {
	var str strings.Builder

	for key, value := range r {
		str.WriteString(fmt.Sprintf("%v: %v\n", key, value))
	}

	return str.String()
}
