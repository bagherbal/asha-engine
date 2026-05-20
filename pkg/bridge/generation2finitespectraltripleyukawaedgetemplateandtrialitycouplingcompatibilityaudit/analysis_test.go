package generation2finitespectraltripleyukawaedgetemplateandtrialitycouplingcompatibilityaudit

import (
	"strings"
	"testing"
)

func TestGate804InheritanceAndFiniteTripleEdges(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inheritance.TD4IsLedger || a.Inheritance.HasYukawaReadout || !a.Inheritance.FiniteTripleNextHost {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.EdgeTemplate.Recorded || len(a.EdgeTemplate.Edges) != 4 || !containsAll(a.EdgeTemplate.Edges, []string{"Q_L -> u_R", "Q_L -> d_R", "L_L -> e_R", "L_L -> nu_R"}) {
		t.Fatalf("bad edge template: %s", FormatEdgeTemplate(a.EdgeTemplate))
	}
	if !containsAll(a.EdgeTemplate.Failures, []string{StatusEdgeNoEigenvalues, StatusEdgeNoMixing}) {
		t.Fatalf("edge firewall missing: %s", FormatEdgeTemplate(a.EdgeTemplate))
	}
}

func TestGate804CompatibilityAndSlotFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Target.Defined || a.Target.SealName != "EdgeTrialityKernelCompatibilitySeal" || !containsAll(a.Target.Failures, []string{StatusKernelNotReadout}) {
		t.Fatalf("bad target: %s", FormatTarget(a.Target))
	}
	if !a.SlotMatching.Audited || !containsAll(a.SlotMatching.Supports, []string{StatusTD4CorrectArity}) || !containsAll(a.SlotMatching.Failures, []string{StatusArityNoEmbedding, StatusNoHLREmbedding}) {
		t.Fatalf("bad slot matching: %s", FormatSlot(a.SlotMatching))
	}
	if !a.FourSector.Audited || len(a.FourSector.Sectors) != 4 || len(a.FourSector.D4Slots) != 3 || !containsAll(a.FourSector.Failures, []string{StatusThreeSlotsNotFour, StatusTD4NoSectorReplacement}) {
		t.Fatalf("bad four-sector firewall: %s", FormatFour(a.FourSector))
	}
}

func TestGate804GaugeHiggsTraceAndPackageStatus(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !containsAll(a.Gauge.Supports, []string{StatusFSTPartialGaugeAssignment}) || !containsAll(a.Gauge.Failures, []string{StatusD4NoGaugeAssignment, StatusNoEdgeToD4Theorem}) {
		t.Fatalf("bad gauge: %s", FormatGauge(a.Gauge))
	}
	if !containsAll(a.Higgs.Supports, []string{StatusHiggsOneFormCandidate}) || !containsAll(a.Higgs.Failures, []string{StatusNoHiggsSlotEmbedding, StatusK7PlusNotD4Vector}) {
		t.Fatalf("bad higgs: %s", FormatHiggs(a.Higgs))
	}
	if !containsAll(a.Trace.Supports, []string{StatusTD4PreTraceKernel}) || !containsAll(a.Trace.Failures, []string{StatusTD4NoTraceInputs, StatusTD4NoABNEffUpdate}) {
		t.Fatalf("bad trace: %s", FormatTrace(a.Trace))
	}
	if !a.Package.Updated || !containsAll(a.Package.PartiallySupplied, []string{"GaugeRepresentationAssignmentSeal", "SectorAssignmentSeal"}) || !containsAll(a.Package.NotSupplied, []string{"TraceAtomExtractionSeal", "RealDescentSeal"}) {
		t.Fatalf("bad package: %s", FormatPackage(a.Package))
	}
}

func TestGate804TheoremAndStatusLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.CHiggs.Preserved || !containsAll(a.CHiggs.Failures, []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}) {
		t.Fatalf("bad C_Higgs firewall: %s", FormatCHiggs(a.CHiggs))
	}
	if !strings.Contains(a.Branch.Next, "Gate 805") || a.Branch.Seal != "EdgeTrialityEmbeddingSeal" {
		t.Fatalf("bad branch: %+v", a.Branch)
	}
	res := Generation2FiniteSpectralTripleYukawaEdgeTemplateAndTrialityCouplingCompatibilityAuditTheorem().Verify()
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
