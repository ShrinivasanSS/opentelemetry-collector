// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build darwin || freebsd
// +build darwin freebsd

package processesscraper

import (
	"github.com/shirou/gopsutil/load"

	"go.opentelemetry.io/collector/model/pdata"
)

const unixSystemSpecificMetricsLen = 0

func appendUnixSystemSpecificProcessesMetrics(metrics pdata.MetricSlice, startTime pdata.Timestamp, now pdata.Timestamp, misc *load.MiscStat) error {
	return nil
}
