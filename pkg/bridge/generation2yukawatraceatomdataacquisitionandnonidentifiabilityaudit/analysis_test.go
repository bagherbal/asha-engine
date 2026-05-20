package generation2yukawatraceatomdataacquisitionandnonidentifiabilityaudit

import (
	"strings"
	"testing"
)

func TestGate795AcquisitionAndNonIdentifiability(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate794.Inherited {
		t.Fatalf("bad inheritance: %+v", a.Gate794)
	}
	if !a.Hierarchy.Defined || !strings.Contains(a.Hierarchy.HighestAvailable, "aggregate") || !containsAll(a.Hierarchy.Priority, []string{"native Yukawa", "singular-value", "sector", "atom", "aggregate"}) {
		t.Fatalf("bad hierarchy: %+v", a.Hierarchy)
	}
	if !a.Acquisition.Required || a.Acquisition.CanPopulate || !hasRow(a.Acquisition.Rows, "a_u", StatusMissing) || !hasRow(a.Acquisition.Rows, "trace atoms", StatusMissing) || !hasRow(a.Acquisition.Rows, "top channel", StatusMissing) || !hasRow(a.Acquisition.Rows, "neutrino", StatusAmbiguous) {
		t.Fatalf("bad acquisition: %s", FormatAcquisition(a.Acquisition))
	}
	if !a.NonIdentifiability.Proved || a.NonIdentifiability.ConstraintCount != 2 || !a.NonIdentifiability.InfiniteFamilies || !containsAll(a.NonIdentifiability.CannotIdentify, []string{"sector", "top", "neutrino", "D4"}) {
		t.Fatalf("bad non-identifiability: %s", FormatNonIdentifiability(a.NonIdentifiability))
	}
}

func TestGate795MinimumAtomsTopBoundsAndRestPressure(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.MinimumAtom.Completed || a.MinimumAtom.MinimumNonzeroAtoms != 4 || !a.MinimumAtom.RequiresRestBeyondThree || !strings.Contains(a.MinimumAtom.CompatibleReading, "top-color") {
		t.Fatalf("bad minimum atom audit: %s", FormatMinimumAtom(a.MinimumAtom))
	}
	if !a.TopBounds.Computed || !closeAbs(a.TopBounds.AOverThree, 0.9474698380779695, 1e-16) || !closeAbs(a.TopBounds.SqrtBOverThree, 0.9471025365183062, 1e-16) || !closeAbs(a.TopBounds.UpperBoundT, 0.9471025365183062, 1e-16) || a.TopBounds.DeterminesT {
		t.Fatalf("bad top bounds: %s", FormatTopBounds(a.TopBounds))
	}
	if !a.LinearizedRest.Recorded || !closeAbs(a.LinearizedRest.AlphaEstimate, 0.0003875905593995199, 5e-16) || a.LinearizedRest.IsTheorem {
		t.Fatalf("bad rest pressure: %s", FormatLinearized(a.LinearizedRest))
	}
}

func TestGate795ConventionsImpactBranchAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Validation.ProtocolExecuted || a.Validation.DataExists || a.Validation.Validated || !containsAll(a.Validation.Rules, []string{"sum sectors", "a^2/b", "sum atoms"}) {
		t.Fatalf("bad validation: %+v", a.Validation)
	}
	if !a.TopRest.ExecutedIfTExists || a.TopRest.TypedTFound || a.TopRest.AlphaBetaComputed || !strings.Contains(a.TopRest.FormulaRatio, "beta") {
		t.Fatalf("bad top/rest: %+v", a.TopRest)
	}
	if !a.Neutrino.Required || !a.Neutrino.Implicit || a.Neutrino.Status != "Y_nu unknown" {
		t.Fatalf("bad neutrino: %+v", a.Neutrino)
	}
	if !a.Scale.Required || a.Scale.Scale != "M_Z" || a.Scale.MultiScaleLedger || a.Scale.ScaleStabilityCertified {
		t.Fatalf("bad scale: %+v", a.Scale)
	}
	if !a.Impact.Recorded || !a.Impact.ValidatedAtomLedgerWouldImprove || !a.Impact.NEffAggregateSealed || a.Impact.CHiggsLevelC {
		t.Fatalf("bad impact: %+v", a.Impact)
	}
	if !a.Branch.Recorded || !strings.Contains(a.Branch.Recommended, "External Yukawa Ledger") || !containsAll(a.Branch.Alternatives, []string{"Sector Contribution", "Native Yukawa"}) {
		t.Fatalf("bad branch: %s", FormatBranch(a.Branch))
	}
	if !a.Firewalls.Enforced || a.Firewalls.AggregateIsAtomLedger || a.Firewalls.TopDominanceTopYukawaTheorem || a.Firewalls.MinimumAtomGenerationTheorem || a.Firewalls.NEffD4Triality || a.Firewalls.SectorDataNativeYukawa || a.Firewalls.ValidatedAtomsPMNSCKM || a.Firewalls.SingleScaleStable || a.Firewalls.CHiggsLevelC || a.Firewalls.TreeProxyPoleMass {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate795TheoremStatusesAndFinalStatement(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.FinalStatement, "does not yet expose") || !strings.Contains(a.FinalStatement, "a,b prove inverse participation") || !strings.Contains(a.FinalStatement, "top-color dominance") {
		t.Fatalf("bad final statement: %s", a.FinalStatement)
	}
	res := Generation2YukawaTraceAtomDataAcquisitionAndNonIdentifiabilityAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
