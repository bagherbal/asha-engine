package generation2masterenvironmentalhistorysealvectoraudit

import (
	"strings"
	"testing"
)

func TestGate605MasterSealTable(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.MinimalFlavorSealDefined || !a.Inherited.SigmaGaugeForBFlav {
		t.Fatalf("bad Gate604 inheritance: %+v", a.Inherited)
	}
	if !containsRow(a.MasterSealTable, "B_flav≈0") || !containsRow(a.MasterSealTable, "D_M / OS/Wick/Hilbert/time") {
		t.Fatalf("missing master rows: %s", FormatMasterSealTable(a.MasterSealTable))
	}
	if !a.Summary.BoundaryClear || a.Summary.NativeCount == 0 || a.Summary.EnvironmentalSealCount == 0 {
		t.Fatalf("bad classification summary: %+v", a.Summary)
	}
}

func TestGate605FormulaAndRanking(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !containsText(a.Formula.HistorySeals, "MinimalFlavorHistoryBranchSeal") {
		t.Fatalf("minimal flavor seal not integrated: %+v", a.Formula)
	}
	if a.Ranking[0].Path != "RG / threshold transport" {
		t.Fatalf("unexpected top ranking: %+v", a.Ranking)
	}
	if a.Firewalls.DerivesKoide || a.Firewalls.DerivesObservedEndpoint || a.Firewalls.SearchesNewConstants {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestGate605TheoremStatuses(t *testing.T) {
	res := Generation2MasterEnvironmentalHistorySealVectorAuditTheorem().Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusMasterSealVectorConstructed, StatusNativeBoundaryClarified, StatusFlavorSealIntegrated, StatusRGThresholdNextActionable, StatusNoNativeFlavorBalanceTheorem, StatusNoProductTimeAirlock, StatusGate605Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
