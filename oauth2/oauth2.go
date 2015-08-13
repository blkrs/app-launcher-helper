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

	"github.com/cloudfoundry/gosteno"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-martini/martini"
)

func ResourceServer(tokenKey []byte) martini.Handler {
	logger := gosteno.NewLogger("resource_server_handler")
	return func(c martini.Context, w http.ResponseWriter, r *http.Request) {
		jwtToken, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
			return tokenKey, nil
		})

		if err == nil && jwtToken.Valid {
			logger.Debug("Token valid")
			c.Map(jwtToken)
		} else {
			logger.Warnf("Token invalid, err=%s", err.Error())
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}

func TokenKey(tokenKeyUrl string) ([]byte, error) {
	resp, err := http.Get(tokenKeyUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return ParseTokenKey(body)
}
