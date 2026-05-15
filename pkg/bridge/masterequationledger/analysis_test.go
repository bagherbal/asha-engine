package masterequationledger

import (
	"strings"
	"testing"
)

func TestMasterEquationLedgerAudit(t *testing.T) {
	a := BuildDefault()
	if len(a.Problems) != 0 {
		t.Fatalf("unexpected validation problems: %v", a.Problems)
	}
	md := MarkdownAudit(a)
	for _, want := range []string{
		"# ASHA Master Equation Registry Audit",
		"S_{Universe}",
		StatusMasterEquationCompiled,
		StatusFailedDoesNotDeriveFlavor,
		StatusFirewallNativeWriteBlocked,
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
