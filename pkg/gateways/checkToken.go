package gateways

import "regexp"

func CheckToken(expectedToken, Token string) bool {

	var re = regexp.MustCompile(`(?m)^\s*(Bearer)\s+(` + regexp.QuoteMeta(Token) + `)\s*$`)

	if re.MatchString(expectedToken) {
		return true
	}
	return false
}
