// :snippet-start: get-metrics-prod
// :state-remove-start: copy
// See entire project at https://github.com/mongodb/atlas-architecture-go-sdk
// :state-remove-end: [copy]
package main

import (
	"atlas-sdk-go/internal/auth"
	"atlas-sdk-go/internal/config"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/atlas-sdk/v20250219001/admin"

	"atlas-sdk-go/internal/metrics"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not loaded: %v", err)
	}
	secrets, cfg, err := config.LoadAll("configs/config.json")
	if err != nil {
		log.Fatalf("config: load config file: %v", err)
	}

	sdk, err := auth.NewClient(cfg, secrets)
	if err != nil {
		log.Fatalf("auth: client init: %v", err)
	}

	ctx := context.Background()
	p := &admin.GetHostMeasurementsApiParams{
		GroupId:   cfg.ProjectID,
		ProcessId: cfg.ProcessID,
		M: &[]string{
			"OPCOUNTER_INSERT", "OPCOUNTER_QUERY", "OPCOUNTER_UPDATE", "TICKETS_AVAILABLE_READS",
			"TICKETS_AVAILABLE_WRITE", "CONNECTIONS", "QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED",
			"QUERY_TARGETING_SCANNED_PER_RETURNED", "SYSTEM_CPU_GUEST", "SYSTEM_CPU_IOWAIT",
			"SYSTEM_CPU_IRQ", "SYSTEM_CPU_KERNEL", "SYSTEM_CPU_NICE", "SYSTEM_CPU_SOFTIRQ",
			"SYSTEM_CPU_STEAL", "SYSTEM_CPU_USER",
		},
		Granularity: admin.PtrString("PT1H"),
		Period:      admin.PtrString("P7D"),
	}

	view, err := metrics.FetchProcessMetrics(ctx, sdk.MonitoringAndLogsApi, p)
	if err != nil {
		log.Fatalf("metrics: fetch process metrics: %v", err)
	}

	out, _ := json.MarshalIndent(view, "", "  ")
	fmt.Println(string(out))
}

// :snippet-end: [get-metrics-prod]
