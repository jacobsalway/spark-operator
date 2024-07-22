# Changelog

All notable changes to the Spark operator Helm chart will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Deploy webhook server with a separated deployment and enable webhook server by default.
- The webhook secret will be created and updated by operator, and will not be deleted after uninstalling the chart.
- The certificates will not change during chart upgrading/rollback process.
- Leader election for controller/webhook is enabled by default to ensure only an controller instance is working when upgrading the chart.

### Changed

- Change value `imagePullSecrets` to `image.pullSecrets`.
- All controller configurations are prefixed with controller e.g. `controller.replicaCount` and `controller.resyncInterval`.
- All webhook configurations are prefixed with webhook e.g. `webhook.replicaCount` and `webhook.failurePolicy`.
- All monitoring configurations are prefixed with prometheus e.g. `prometheus.metrics` and `prometheus.podMonitor`.
- Change the update strategy of controller/webhook deployment from `recreate` to the default (`rollingUpdate`).
- Change the default of `webhook.timeoutSeconds` from `30` to `10`.
- Change the default of `webhook.failurePolicy` from `Ignore` to `Fail`.
- Change the default `spark.jobNamespaces` from `[]` to `["default]`.
- Service accounts are configured with `controller.serviceAccount`, `webhook.serviceAccount` and `spark.serviceAccount` respectively.
- RBAC resources are configured with `controller.rbac`, `webhook.rbac` and `spark.rbac` respectively.
- The available options of `logLevel` will be one of `info`, `debug` and `error`.
