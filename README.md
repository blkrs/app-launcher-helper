App Launcher Helper
==================

App Launcher Helper is a service that provides a list of services created by a specific instance of [App Launching Service Broker](https://github.com/trustedanalytics/app-launching-service-broker).

Usage
=====

The problem with App Launching SB is that there's no direct connection between a service and a related application - they're bound by naming convention only. App Launcher Helper is trying to fill this gap by providing a list of entries on a REST call:

```
http://hostname/rest/orgs/:orgId/atkinstances
```

Example response body:

```
{  
  "instances": [
    {
	  "name": "Name of a service instance",
  	  "url": "Url of an application related to a service instance",
	  "state": "current state of an app - STARTED, STOPPED, etc ...."
    }
  ]
}
```

AL Helper finds all service instances started in a specific organization for a service with a name configured by env variable. 
It's worth to mention that the application is an OAuth2 Resource Server, which means that there's access token in Authorization header needed. When deployed on Cloud Foundry, the application can be queried this way:

```
curl -H "Authorization: \`cf oauth-token|grep bearer\`" http://applauncher-helper.54.154.194.181.xip.io/rest/orgs/:orgId/atkinstances
```

Development
===========

To locally develop this application you'll need `godep` tool to manage dependencies and build the project:

```
$ godep go build
$ godep go test ./...
```

This might not work if you clone the project outside of GOPATH directory, because of absolute subpackages imports. Recommended way of cloning the project is:
```
$ go get github.com/trustedanalytics/app-launcher-helper
``` 


Deployment
==========

Before pushing the app to the Cloud Foundry, there're three env variables to be set:

* `TOKEN_KEY_URL` - an address of a key, to validate user's access token;
* `API_URL` - Cloud Foundry API address;
* `SERVICE_NAME` - a service name provided by App Launching Service Broker;
* `SE_SERVICE_NAME` - a Scoring Engine service name.

They are defined in manifest.yml, but they can be set by a `cf set-env` command as well.
When environment is ready, there's only one command needed:

```
$ cf push
```

Versioning
==========
`Bumpversion` tools is used to manage project version number, which is kept in two places: .bumpversion.cfg and manifest.yml. The first one is for bumpversion itself,
while the second one helps to identify the version of an application deployed in Cloud Foundry.

There's no need to use bumpversion manually - it's being used by CI.                                                                                             
