package generation2trialitytoyukawareadoutminimalityandnogoaudit

import (
	"strings"
	"testing"
)

func TestGate803PackageDefinitionAndInheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inheritance.HasNativeRealDescent || a.Inheritance.HasSectorOperators || a.Inheritance.HasTraceAtoms || !strings.Contains(a.Inheritance.AirlockLevel, "T1") {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Package.Defined || len(a.Package.Seals) != 10 || !containsAll(a.Package.Seals, []string{"RealDescentSeal", "TraceAtomExtractionSeal", "NonCircularitySeal"}) {
		t.Fatalf("bad package: %s", FormatPackage(a.Package))
	}
}

func TestGate803SealObstructions(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	seals := []SealAudit{a.RealDescent, a.GaugeAssignment, a.SectorAssignment, a.GenerationCarrier, a.HermitianOperator, a.Hierarchy, a.TraceAtom, a.Color, a.Scale, a.NonCircularity}
	for _, s := range seals {
		if !s.Audited || len(s.Required) == 0 || len(s.Failures) == 0 {
			t.Fatalf("bad seal audit %s: %s", s.Name, FormatSeal(s))
		}
	}
	if !containsAll(a.Color.Supports, []string{StatusColorSU3TraceMultiplicity}) {
		t.Fatalf("color support missing: %s", FormatSeal(a.Color))
	}
}

func TestGate803MinimalityNoGoAndPaths(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Minimality.Audited || len(a.Minimality.RemovalFailures) != 10 || !containsAll(a.Minimality.Failures, []string{StatusCannotCompressToTD4}) {
		t.Fatalf("bad minimality: %s", FormatMinimality(a.Minimality))
	}
	if !a.NoGo.Defined || !containsAll(a.NoGo.CannotConstruct, []string{"a,b,N_eff", "PMNS/CKM", "C_Higgs update"}) {
		t.Fatalf("bad no-go: %s", FormatNoGo(a.NoGo))
	}
	if !a.Paths.Recorded || !containsAll(a.Paths.Supports, []string{StatusExternalFastest}) || !containsAll(a.Paths.Failures, []string{StatusTrialityNotReadyExternal}) {
		t.Fatalf("bad paths: %s", FormatPaths(a.Paths))
	}
}

func TestGate803CurrentASHAChainCHiggsAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Current.Audited || !containsAll(a.Current.Supports, []string{StatusFSTEdgeTemplate, StatusColorSU3TraceMultiplicity}) || !containsAll(a.Current.Failures, []string{StatusNoCurrentReadoutPackage}) {
		t.Fatalf("bad current ASHA audit: %s", FormatCurrent(a.Current))
	}
	if !a.CHiggs.Preserved || !containsAll(a.CHiggs.Failures, []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}) {
		t.Fatalf("bad C_Higgs firewall: %s", FormatCHiggs(a.CHiggs))
	}
	if !strings.Contains(a.Branch.NextNative, "Gate 804") || !strings.Contains(a.Branch.NextNative, "Finite Spectral Triple") {
		t.Fatalf("bad branch: %+v", a.Branch)
	}
	res := Generation2TrialityToYukawaReadoutMinimalityAndNoGoAuditTheorem().Verify()
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
