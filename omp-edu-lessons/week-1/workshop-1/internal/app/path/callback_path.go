package path

import (
	"errors"
	"fmt"
	"strings"
)

type CallbackPath struct {
	Domain       string
	Subdomain    string
	CallbackName string
	CallbackData string
}

var ErrUnknownCallback = errors.New("unknown callback")

func ParseCallback(callbackData string) (CallbackPath, error) {
	callbackParts := strings.SplitN(callbackData, "__", 4)
	if len(callbackParts) != 4 {
		return CallbackPath{}, ErrUnknownCallback
	}

	return CallbackPath{
		Domain:       callbackParts[0],
		Subdomain:    callbackParts[1],
		CallbackName: callbackParts[2],
		CallbackData: callbackParts[3],
	}, nil
}

func (p CallbackPath) String() string {
	return fmt.Sprintf("%s__%s__%s__%s", p.Domain, p.Subdomain, p.CallbackName, p.CallbackData)
}
