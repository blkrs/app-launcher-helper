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

type ResourceMetadata struct {
	Id  string `json:"guid"`
	Url string `json:"url"`
}

type Resource struct {
	Metadata ResourceMetadata `json:"metadata"`
	Entity   ResourceEntity   `json:"entity"`
}

type ResourceEntity struct {
	Label           string `json:"label"`
	ServicePlansUrl string `json:"service_plans_url"`
}

type ResourceList struct {
	Count     int        `json:"total_results"`
	Resources []Resource `json:"resources"`
}

type SpaceSummary struct {
	Apps     []Application `json:"apps"`
	Services []Service     `json:"services"`
}

type Application struct {
	Name  string   `json:"name"`
	Urls  []string `json:"urls"`
	Guid  string   `json:"guid"`
	State string   `json:"state"`
	ServiceNames []string `json:"service_names"`
}

type Service struct {
	Name        string      `json:"name"`
	Guid        string      `json:"guid"`
	ServicePlan ServicePlan `json:"service_plan"`
}

type ServicePlan struct {
	Guid    string             `json:"guid"`
	Service ServicePlanService `json:"service"`
}
type ServicePlanService struct {
	Label string `json:"label"`
}

func (rl *ResourceList) Contains(Id string) bool {
	for _, r := range rl.Resources {
		if r.Metadata.Id == Id {
			return true
		}
	}

	return false
}

func (rl *ResourceList) IdList() []string {
	ids := make([]string, rl.Count)
	for i, r := range rl.Resources {
		ids[i] = r.Metadata.Id
	}

	return ids
}

type CloudController interface {
	Spaces(organization string) (*ResourceList, error)
	SpaceSummary(space string) (*SpaceSummary, error)
	Services() (*ResourceList, error)
	ServicePlans(Name string) (*ResourceList, error)
}
