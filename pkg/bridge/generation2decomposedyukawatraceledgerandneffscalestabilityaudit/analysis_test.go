package generation2decomposedyukawatraceledgerandneffscalestabilityaudit

import (
	"strings"
	"testing"
)

func TestGate793TraceAtomAndTopColorLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate792.Inherited || !a.Gate792.NEffTopNumericalLeverage {
		t.Fatalf("bad inheritance: %+v", a.Gate792)
	}
	if !a.TraceAtom.Recorded || !closeAbs(a.TraceAtom.Ratio, 0.33307493962706697, 1e-16) || !closeAbs(a.TraceAtom.NEff, 3.0023273474722147, 5e-16) {
		t.Fatalf("bad trace atom ledger: %s", FormatTraceAtom(a.TraceAtom))
	}
	if !a.TopColor.Inherited || !closeAbs(a.TopColor.DeltaRatio, -0.0002583937062663466, 1e-15) || !closeAbs(a.TopColor.NEffMinusThree, 0.0023273474722147, 1e-15) || a.TopColor.CurrentCertifiedSource == "" {
		t.Fatalf("bad top color ledger: %s", FormatTopColor(a.TopColor))
	}
}

func TestGate793MissingDecompositionAndTrialityFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sector.Defined || a.Sector.SectorTracesAvailable || a.Sector.MissingSeal != "DecomposedYukawaTraceLedgerSeal" || !containsAll(a.Sector.RequiredQuadratic, []string{"a_u", "a_d", "a_e", "a_nu"}) {
		t.Fatalf("bad sector requirement: %+v", a.Sector)
	}
	if !a.TopRest.FormulaInherited || a.TopRest.TypedTopChannelAvailable || a.TopRest.DecomposedLedgerAvailable || a.TopRest.Verdict != StatusNoAlphaBetaWithoutTopLedger {
		t.Fatalf("bad top/rest audit: %+v", a.TopRest)
	}
	if !a.Generation.Audited || a.Generation.GenerationCarrierCertified || a.Generation.GenerationResolvedTraceLedger {
		t.Fatalf("bad generation audit: %+v", a.Generation)
	}
	if !a.D4.RequirementsDefined || !a.D4.RealFormFirewallAudited || !a.D4.StrongFutureCandidate || a.D4.CurrentCertified || !containsAll(a.D4.RequiredPackage, []string{"real-form", "trace-readout", "breaking"}) {
		t.Fatalf("bad D4 audit: %+v", a.D4)
	}
}

func TestGate793ScaleImpactClassificationBranchAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Scale.RequirementsDefined || !a.Scale.DifferentialRecorded || a.Scale.Scale != "M_Z" || !strings.Contains(a.Scale.Differential, "2 d ln a - d ln b") || a.Scale.MultiScaleLedgerAvailable {
		t.Fatalf("bad scale audit: %s", FormatScale(a.Scale))
	}
	if !a.Impact.Recorded || !closeAbs(a.Impact.CHiggsTopColor, 1.038025177923625, 1e-15) || !closeAbs(a.Impact.DeltaCHiggs, 0.0008046575187645733, 1e-16) || !closeAbs(a.Impact.DeltaTreeProxy, 0.04862437568908, 5e-14) {
		t.Fatalf("bad impact audit: %s", FormatImpact(a.Impact))
	}
	if !a.SourceClassification.Completed || !a.SourceClassification.TopColorCurrent || a.SourceClassification.GenerationCertified || a.SourceClassification.D4Current || !a.SourceClassification.AggregateSealedCurrent {
		t.Fatalf("bad source classification: %+v", a.SourceClassification)
	}
	if !a.Symbolic.Audited || !a.Symbolic.D4MotivationOnly || a.Symbolic.SymbolicPatternEvidence {
		t.Fatalf("bad symbolic firewall: %+v", a.Symbolic)
	}
	if !a.Branch.Recorded || !strings.Contains(a.Branch.Recommended, "DecomposedYukawaTraceLedgerSeal") || !containsAll(a.Branch.Alternatives, []string{"Sector Contribution", "D4 Triality"}) {
		t.Fatalf("bad branch: %s", FormatBranch(a.Branch))
	}
	if !a.Firewalls.Enforced || a.Firewalls.NEffGenerationTheorem || a.Firewalls.NEffD4TrialityTheorem || a.Firewalls.TopColorGeneration || a.Firewalls.D4ResonanceReadoutTheorem || a.Firewalls.Spin8AutomaticNative || a.Firewalls.SymbolicProof || a.Firewalls.ScaleStabilityAssumed || a.Firewalls.CHiggsPoleMass || a.Firewalls.TreeProxyShiftPoleCorrection {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate793TheoremStatusesAndFinalStatement(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.FinalStatement, "does not derive N_eff natively") || !strings.Contains(a.FinalStatement, "color-tripled top dominance") || !strings.Contains(a.FinalStatement, "Gate 794") {
		t.Fatalf("bad final statement: %s", a.FinalStatement)
	}
	res := Generation2DecomposedYukawaTraceLedgerAndNEffScaleStabilityAuditTheorem().Verify()
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
			t.Fatalf("missing status note %s", want)
		}
	}
}
