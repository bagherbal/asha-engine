package report

import (
	"fmt"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func RenderTerminal(results []theorem.Result) string {
	var b strings.Builder
	b.WriteString("================================================================================\n")
	b.WriteString("                          ASHA ENGINE — THEOREM LADDER                          \n")
	b.WriteString("================================================================================\n\n")

	for _, r := range results {
		mark := "✔"
		if !r.Passed() {
			mark = "✘"
		}
		b.WriteString(fmt.Sprintf("%s %-18s %-14s %s\n", mark, r.Status, r.ID, r.Name))
		for _, c := range r.Checks {
			checkMark := "  ✓"
			if !c.Passed {
				checkMark = "  ✗"
			}
			b.WriteString(fmt.Sprintf("%s %-34s %s\n", checkMark, c.Name, c.Detail))
		}
		for _, note := range r.Notes {
			b.WriteString(fmt.Sprintf("  • %s\n", note))
		}
		b.WriteString("\n")
	}

	b.WriteString("================================================================================\n")
	b.WriteString("Finite core executed. Bridge-layer constants are intentionally not computed yet.\n")
	b.WriteString("================================================================================\n")
	return b.String()
}
