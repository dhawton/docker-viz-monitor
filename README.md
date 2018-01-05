# docker-viz-monitor

[![Build Status](https://travis-ci.org/dhawton/docker-viz-monitor.svg?branch=master)](https://travis-ci.org/dhawton/docker-viz-monitor)
[![GitHub issues](https://img.shields.io/github/issues/dhawton/docker-viz-monitor.svg)](https://github.com/dhawton/docker-viz-monitor/issues)
[![GitHub license](https://img.shields.io/github/license/dhawton/docker-viz-monitor.svg)](https://github.com/dhawton/docker-viz-monitor/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/dhawton/docker-viz-monitor.svg)](https://github.com/dhawton/docker-viz-monitor/releases)

Generates JSON files for monitor a Docker Swarm
---

* 3 environment variables must be defined
  * JSON_NODES - location to dump the nodes json
  * JSON_SERVICES
  * JSON_TASKS

Must have access to /var/run/docker.sock