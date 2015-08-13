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
	"strings"

	"github.com/cloudfoundry/gosteno"
)

type SpaceSummaryHelper struct {
	logger          *gosteno.Logger
}

func NewSpaceSummaryHelper() SpaceSummaryHelper {
	return SpaceSummaryHelper{
		logger:          gosteno.NewLogger("space_summary_helper"),
	}
}

func (p *SpaceSummaryHelper) getMainGuidPart(guid string ) string {
	split := strings.Split(guid,"-")
	mainPart := split[0] + "-" + split[1] + "-" + split[2] + "-" + split[3]
	return mainPart
}

func (p *SpaceSummaryHelper) getMapOfAppsByService(planLabel string, serviceSearchString string, summary *SpaceSummary, apps map[string]Application ) map[string]AtkInstance{
	seInstancesMap := make(map[string]AtkInstance)
	for _, s := range summary.Services {
		if s.ServicePlan.Service.Label == planLabel {
			if a, ok := apps[UuidToAppName(s.Guid,planLabel)]; ok {
				p.logger.Debug("App name: " + s.Name)
				serviceName := p.FindRelatedService(*summary, serviceSearchString, s.Guid)
				seInstancesMap[serviceName] = AtkInstance{s.Name, a.Urls[0], a.Guid, s.Guid, a.State, nil}
			} else {
				p.logger.Warn("App not found for service: " + s.Guid)
			}
		}
	}
	return seInstancesMap
}

func (p *SpaceSummaryHelper) FindRelatedService(summary SpaceSummary, serviceSearchString string, guid string) string {
	var serviceName string
	a := p.FindAppBoundToService(summary, guid)
	for _, ss := range a.ServiceNames {
		if strings.Contains(ss,serviceSearchString) {
			serviceName = ss
			break
		}
	}
	return serviceName
}

func (p *SpaceSummaryHelper) FindAppBoundToService(summary SpaceSummary, guid string) Application {
	var app Application
	commonPartOfId := p.getMainGuidPart(guid)
	for _, a := range summary.Apps {
		if strings.Contains(a.Name,commonPartOfId) {
			return a
		}
	}
	return app
}


