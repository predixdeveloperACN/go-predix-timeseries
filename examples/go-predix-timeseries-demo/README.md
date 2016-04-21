Demo app consists of two parts:
- "Sensor" pushing data Time Series
- Web app reading from Time Series and visualizing it

##Prerequisites
- Predix account
- Cloudfoundry CLI
- CloudFoundry UAA Command Line Client

##Creating application and setting up services

First of all you need to create an application in Predix:
```
cf push APP_NAME
```
Then you need to create UAA service and bind it to your application
```
cf create-service predix-uaa Tiered go-predix-timeseries-demo-uaa -c '{"adminClientSecret":"<secret>"}'
cf bind-service APP_NAME go-predix-timeseries-demo-uaa
```
Create and bind Time Series service:
```
cf create-service predix-timeseries Bronze go-predix-timeseries-demo-ts -c '{"trustedIssuerIds":["<issuerId of UAA service>"]}'
cf bind-service APP_NAME go-predix-timeseries-demo-ts
```
You can find `issuerId` in `cf env APP_NAME` output:
```JSON
cf env APP_NAME
...
  "predix-uaa": [
    {
      "credentials": {
        "issuerId": "https://36f83803-6e60-4ad7-8635-cac7b140f149.predix-uaa.run.aws-usw02-pr.ice.predix.io/oauth/token",
        "uri": "https://36f83803-6e60-4ad7-8635-cac7b140f149.predix-uaa.run.aws-usw02-pr.ice.predix.io",
        "zone": {
          "http-header-name": "X-Identity-Zone-Id",
          "http-header-value": "36f83803-6e60-4ad7-8635-cac7b140f149"
        }
      },
      "label": "predix-uaa",
      "name": "go-predix-timeseries-demo-uaa",
      "plan": "Tiered",
      "provider": null,
      "syslog_drain_url": null,
      "tags": []
    }
  ]
...
```
Now we need to create ingest and query clients in our UAA service with apporiate authorities. To find out requeried authorities type `cf env APP_NAME`
and look for `zone-token-scopes` fields in `credentials` for predix-timeseries service:
```JSON
  "predix-timeseries": [
    {
      "credentials": {
        "ingest": {
          "uri": "wss://gateway-predix-data-services.run.aws-usw02-pr.ice.predix.io/v1/stream/messages",
          "zone-http-header-name": "Predix-Zone-Id",
          "zone-http-header-value": "bfdd9bb0-0b0f-4667-8338-18b048c8ce31",
          "zone-token-scopes": [
            "timeseries.zones.bfdd9bb0-0b0f-4667-8338-18b048c8ce31.user",
            "timeseries.zones.bfdd9bb0-0b0f-4667-8338-18b048c8ce31.ingest"
          ]
        },
        "query": {
          "uri": "https://time-series-store-predix.run.aws-usw02-pr.ice.predix.io/v1/datapoints",
          "zone-http-header-name": "Predix-Zone-Id",
          "zone-http-header-value": "bfdd9bb0-0b0f-4667-8338-18b048c8ce31",
          "zone-token-scopes": [
            "timeseries.zones.bfdd9bb0-0b0f-4667-8338-18b048c8ce31.user",
            "timeseries.zones.bfdd9bb0-0b0f-4667-8338-18b048c8ce31.query"
          ]
        }
      },
      "label": "predix-timeseries",
      "name": "go-predix-timeseries-demo-ts",
      "plan": "Bronze",
      "provider": null,
      "syslog_drain_url": null,
      "tags": [
        "timeseries",
        "time-series",
        "time series"
      ]
    }
  ],
```
To create client we will use [uaac](https://github.com/cloudfoundry/cf-uaac):
```
uaac target UAA_uri
uaac token client get admin
uaac client add ingest --authorized_grant_types client_credentials --authorities "timeseries.zones.bfdd9bb0-0b0f-4667-8338-18b048c8ce31.user,timeseries.zones.bfdd9bb0-0b0f-4667-8338-18b048c8ce31.ingest"
uaac client add query --authorized_grant_types client_credentials --authorities "timeseries.zones.bfdd9bb0-0b0f-4667-8338-18b048c8ce31.user,timeseries.zones.bfdd9bb0-0b0f-4667-8338-18b048c8ce31.query"
```
Now set CLIENT_ID and CLIENT_SECRET environment variables and restage application  to ensure the environment variable changes take effect:
```
cf set-env go-predix-timeseries-demo CLIENT_ID query
cf set-env go-predix-timeseries-demo CLIENT_SECRET <secret_for_query_client>
cf restage APP_NAME
```
Run "sensor" to start pushing data to Time Series service:
```
cd sensor
go run sensor.go -clientId=ingest -clientSecret=<secrer_for_ingest_client> -ingestUrl="ingest_url" -uaaIssuerId="uaa_IssuerId" -zoneId="zone_id"
```
Look for ingest url and zone id(`uri` and `zone-http-header-value` fields) in Time Series credentials:
```JSON
  "ingest": {
    "uri": "wss://gateway-predix-data-services.run.aws-usw02-pr.ice.predix.io/v1/stream/messages",
    "zone-http-header-name": "Predix-Zone-Id",
    "zone-http-header-value": "bfdd9bb0-0b0f-4667-8338-18b048c8ce31",
    "zone-token-scopes": [
      "timeseries.zones.bfdd9bb0-0b0f-4667-8338-18b048c8ce31.user",
      "timeseries.zones.bfdd9bb0-0b0f-4667-8338-18b048c8ce31.ingest"
    ]
  }
```
