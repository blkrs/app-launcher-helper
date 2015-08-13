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
package service

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestUuidToAppName(t *testing.T) {
	RegisterTestingT(t)

	app := UuidToAppName("7e587e45-08a6-46d5-a412-52d5eb897299","atk")
	Expect(app).To(Equal("atk-7e587e45-08a6-46d5-a412"))
}
