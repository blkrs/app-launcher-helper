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
package cc

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cloudfoundry/gosteno"

	"github.com/trustedanalytics/app-launcher-helper/service"
)

type RestCloudController struct {
	client      *http.Client
	apiUrl      string
	accessToken string
	logger      *gosteno.Logger
}

func NewRestCloudController(apiUrl string, accessToken string) service.CloudController {
	return &RestCloudController{
		client:      &http.Client{},
		apiUrl:      apiUrl,
		accessToken: accessToken,
		logger:      gosteno.NewLogger("cc_client"),
	}
}

func (c *RestCloudController) Spaces(organization string) (*service.ResourceList, error) {
	var resList service.ResourceList
	return &resList, c.doGet("/v2/organizations/"+organization+"/spaces", &resList)
}

func (c *RestCloudController) SpaceSummary(space string) (*service.SpaceSummary, error) {
	var summary service.SpaceSummary
	return &summary, c.doGet("/v2/spaces/"+space+"/summary", &summary)
}

func (c *RestCloudController) Services() (*service.ResourceList, error) {
	var services service.ResourceList
	return &services, c.doGet("/v2/services", &services)
}

func (c *RestCloudController) ServicePlans(servicePlansUrl string) (*service.ResourceList, error) {
	var services service.ResourceList
	return &services, c.doGet(servicePlansUrl, &services)
}

func (c *RestCloudController) doGet(path string, target interface{}) error {
	c.logger.Debug("GET " + path)
	req, err := http.NewRequest("GET", c.apiUrl+path, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "bearer "+c.accessToken)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	dec := json.NewDecoder(resp.Body)

	return dec.Decode(&target)
}
