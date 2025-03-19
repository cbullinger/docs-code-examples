// :snippet-start: script-structs
package main

import "time"

type ListAtlasProcessesParams struct {
	GroupID      string `json:"groupId"`
	IncludeCount *bool  `json:"includeCount,omitempty"`
	ItemsPerPage *int   `json:"itemsPerPage,omitempty"`
	PageNum      *int   `json:"pageNum,omitempty"`
}

type Z_GetHostLogsParams struct {
	GroupID   string `json:"groupId"`
	HostName  string `json:"hostName"`
	LogName   string `json:"logName"`
	EndDate   *int64 `json:"endDate,omitempty"`
	StartDate *int64 `json:"startDate,omitempty"`
}

type ListProjectsParams struct {
	GroupID      string `json:"groupId"`
	ItemsPerPage *int   `json:"itemsPerPage,omitempty"`
	IncludeCount *bool  `json:"includeCount,omitempty"`
	PageNum      *int   `json:"pageNum,omitempty"`
}

type Z_HostMetricParams struct {
	GroupID     string     `json:"groupId"`
	ProcessID   string     `json:"processId"`
	Granularity *string    `json:"granularity"`
	M           *[]string  `json:"metrics"`
	Period      *string    `json:"period"`
	Start       *time.Time `json:"start,omitempty"`
	End         *time.Time `json:"end,omitempty"`
}
type Z_ClusterMetricParams struct {
	GroupID       string     `json:"groupId"`
	ProcessID     string     `json:"processId"`
	PartitionName string     `json:"partitionName"`
	M             *[]string  `json:"metrics,omitempty"`
	Period        *string    `json:"period,omitempty"`
	Start         *time.Time `json:"start,omitempty"`
	End           *time.Time `json:"end,omitempty"`
}

type Measurement struct {
	Name string                 `json:"name"`
	Data []MetricDataPointAtlas `json:"data"`
}

type MetricDataPointAtlas struct {
	Timestamp string  `json:"timestamp"`
	Value     float64 `json:"value"`
}
