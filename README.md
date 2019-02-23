# modeapi

a command line client tool for interacting with Mode Analytics API

## Mode REST API

`http://developer.modeanalytics.com`

### Supported* Endpoints

- [x] GET /api/account
- [x] GET /api/${org}

- [x] GET /api/${org}/spaces/
- [ ] POST /api/${org}/spaces
- [x] GET /api/${org}/spaces/${space_token}
- [x] DELETE /api/${org}/spaces/${space_token}

- [ ] GET /api/${org}/spaces/${space_token}/reports

- [ ] GET /api/${org}/reports/${report_token}
- [ ] DELETE /api/${org}/reports/${report_token}

- [ ] GET /api/${org}/reports/${report_token}/runs
- [ ] POST /api/${org}/reports/${report_token}/runs
- [ ] GET /api/${org}/reports/${report_token}/runs/${run_token}

- [ ] GET /api/${org}/reports/${report_token}/schedules
- [ ] GET /api/${org}/reports/${report_token}/schedules/${schedule_token}
- [ ] DELETE /api/${org}/reports/${report_token}/schedules/${subscription_token}
- [ ] PATCH /api/${org}/reports/${report_token}/schedules/${schedule_token}

- [ ] GET /api/${org}/reports/${report_token}/subscriptions
- [ ] POST /api/${org}/reports/${report_token}/subscription
- [ ] GET /api/${org}/reports/${report_token}/subscriptions/${subscription_token}
- [ ] DELETE /api/${org}/reports/${report_token}/subscriptions/${subscription_token}

- [ ] GET /api/${org}/reports/${report_token}/subscriptions/${subscription_token}/email_memberships
- [ ] POST /api/${org}/reports/${report_token}/subscriptions/${subscription_token}/email_memberships
- [ ] GET /api/${org}/reports/${report_token}/subscriptions/${subscription_token}/email_memberships/${email_token}
- [ ] DELETE /api/${org}/reports/${report_token}/subscriptions/${subscription_token}/email_memberships/${email_token}

- [ ] GET /api/${org}/reports/${report_token}/subscriptions/${subscription_token}/slack_memberships
- [ ] POST /api/${org}/reports/${report_token}/subscriptions/${subscription_token}/slack_memberships
- [ ] GET /api/${org}/reports/${report_token}/subscriptions/${subscription_token}/slack_memberships/${slack_token}
- [ ] DELETE /api/${org}/reports/${report_token}/subscriptions/${subscription_token}/slack_memberships/${slack_token}

- [ ] PATCH /api/${org}/reports/${report_token}/sync_to_github

## environment config

create a config file in your home directory `~/.config/.modeanalytics.json`

```json
{
  "default": {
    "api_url": "https://modeanalytics.com",
    "mode_org": "union_pos",
    "api_token": "",
    "api_secret": ""
  }
}
```

for support of multiple configurations create additional top level elements in your config file and pass a matching `--env` flag with your commands

if no `--env` flag is passed the cli will load the `default` configuration

## load build harness

```sh
make init
```

## load dependencies

```sh
GO111MODULE=on go mod vendor
```
