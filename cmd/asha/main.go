package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bagherbal/asha-engine/pkg/asha"
)

func main() {
	var (
		formatFlag   = flag.String("format", "text", "output format: text, markdown, json")
		scenarioFlag = flag.String("scenario", "all", "scenario: all, native, higgs, family, dark-stable-thermal, cosmology, ci")
		betaFlag     = flag.Float64("beta", asha.DefaultBeta, "family KMS beta for quarantined K_gen scenario")
		planckFlag   = flag.Float64("planck-gev", asha.DefaultPlanckMassGeV, "Planck mass in GeV for Pfaffian bridge convention")
		outFlag      = flag.String("out", "", "optional output path; stdout if empty")
		strictFlag   = flag.Bool("strict", true, "exit non-zero if any runtime consistency check fails")
	)
	flag.Parse()

	format, err := asha.ParseFormat(*formatFlag)
	fatalIf(err)
	scenario, err := asha.ParseScenario(*scenarioFlag)
	fatalIf(err)

	engine := asha.New(asha.WithBeta(*betaFlag), asha.WithPlanckMassGeV(*planckFlag))
	report, err := engine.Report(scenario)
	fatalIf(err)

	var f *os.File
	if *outFlag != "" {
		f, err = os.Create(*outFlag)
		fatalIf(err)
		defer f.Close()
	} else {
		f = os.Stdout
	}
	fatalIf(asha.WriteReport(f, report, format))

	if *strictFlag {
		for _, check := range report.Checks {
			if !check.Passed {
				fmt.Fprintf(os.Stderr, "asha runtime check failed: %s: %s\n", check.Name, check.Detail)
				os.Exit(2)
			}
		}
	}
}

func fatalIf(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "asha:", err)
		os.Exit(1)
	}
}
