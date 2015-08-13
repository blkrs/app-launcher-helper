/**
 * Copyright (c) 2015 Intel Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package oauth2

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-martini/martini"
	. "github.com/onsi/gomega"
)

var testData = []struct {
	accessToken string
	httpCode    int
}{
	{
		validAccessToken(),
		http.StatusOK,
	},
	{
		"",
		http.StatusUnauthorized,
	},
	{
		"invalid token",
		http.StatusUnauthorized,
	},
}

func Test_ResourceServerHandler(t *testing.T) {
	RegisterTestingT(t)

	tokenKey, _ := ioutil.ReadFile("test/sample_key.pub")
	handler := ResourceServer(tokenKey)

	m := martini.Classic()
	m.Use(handler)
	m.Get("/valid-token", func(t *jwt.Token) (int, string) {
		if t.Valid {
			return http.StatusOK, "OK"
		}
		return http.StatusInternalServerError, "Error"
	})

	for _, test := range testData {
		r, _ := http.NewRequest("GET", "/valid-token", nil)
		if test.accessToken != "" {
			r.Header.Add("Authorization", "bearer "+test.accessToken)
		}
		recorder := httptest.NewRecorder()
		m.ServeHTTP(recorder, r)

		Expect(recorder.Code).To(Equal(test.httpCode))
	}
}

func validAccessToken() string {
	key, _ := ioutil.ReadFile("test/sample_key")
	parsedKey, _ := jwt.ParseRSAPrivateKeyFromPEM(key)

	t := jwt.New(jwt.SigningMethodRS256)
	signed, _ := t.SignedString(parsedKey)

	return signed
}
