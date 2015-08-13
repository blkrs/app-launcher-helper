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
package main

import (
	"os"

	"github.com/cloudfoundry/gosteno"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-martini/martini"
	"github.com/trustedanalytics/app-launcher-helper/cc"
	"github.com/trustedanalytics/app-launcher-helper/config"
	atkoauth2 "github.com/trustedanalytics/app-launcher-helper/oauth2"
	"github.com/trustedanalytics/app-launcher-helper/service"
	"github.com/martini-contrib/render"
)

func main() {
	c := &gosteno.Config{
		Sinks: []gosteno.Sink{
			gosteno.NewIOSink(os.Stdout),
		},
		Level:     gosteno.LOG_DEBUG,
		Codec:     gosteno.NewJsonPrettifier(0),
		EnableLOC: true,
	}
	gosteno.Init(c)

	logger := gosteno.NewLogger("atk_instances")

	conf := config.NewConfig()

	m := martini.Classic()

	key, err := atkoauth2.TokenKey(conf.TokenKeyUrl)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	m.Handlers(
		atkoauth2.ResourceServer(key),
		martini.Logger(),
		render.Renderer(render.Options{IndentJSON: true}),
	)

	m.Get("/rest/orgs/:id/atkinstances", func(params martini.Params, t *jwt.Token, r render.Render) {

		cloudController := cc.NewRestCloudController(conf.ApiUrl, t.Raw)

		spaceSummaryHelper := service.NewSpaceSummaryHelper()

		srv := service.NewAtkListService(cloudController, spaceSummaryHelper)

		if (len(conf.CommonService) == 0) {
			conf.CommonService = "postgres"
		}

		//TODO: check id param

		instances, err := srv.GetAllInstances(conf.ServiceLabel,
			                                  conf.ScoringEngineLabel,
			                                  conf.CommonService,
									params["id"])
		if err != nil {
			r.JSON(500, err.Error())
		}

		r.JSON(200, instances)
	})

	m.Run()
}
