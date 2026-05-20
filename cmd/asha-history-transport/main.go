package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bagherbal/asha-engine/pkg/historytransport"
)

func main() {
	out := flag.String("out", "history_transport/asha_history_transport_end_calculation_v1", "directory for 01_inputs.yaml ... 07_summary.md")
	flag.Parse()
	bundle, err := historytransport.BuildDefault()
	if err != nil {
		log.Fatal(err)
	}
	if err := historytransport.WriteBundle(*out, bundle); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("wrote %s to %s\n", historytransport.TaskName, *out)
}
