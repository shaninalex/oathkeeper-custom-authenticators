package main

import (
	"net/http"
	"net/url"

	"github.com/golang-jwt/jwt/v5"
)

type AuthenticationSession struct {
	Subject      string
	Extra        map[string]interface{}
	Header       http.Header
	MatchContext MatchContext
}

type MatchContext struct {
	RegexpCaptureGroups []string
	URL                 *url.URL
	Method              string
	Header              http.Header
}

type MyCustomClaims struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Exp   string `json:"exp"`
	jwt.RegisteredClaims
}
