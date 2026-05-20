package generation2generationmixingoperatorsourceandflavororientationreadoutsealaudit

import (
	"strings"
	"testing"
)

func TestGate789RequiredGenerationMixingObjects(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate788.Inherited || !a.Gate788.KappaOrientFocus || a.Gate788.Formula != "sin^2(theta13)/4 - J_CKM" {
		t.Fatalf("bad Gate788 inheritance: %+v", a.Gate788)
	}
	if !a.Required.Defined || !a.Required.SectorMisalignmentNeed || a.Required.NativeUPMNSOrVCKMExists {
		t.Fatalf("bad required objects audit: %+v", a.Required)
	}
	for _, want := range []string{"generation carrier G_gen", "sector Yukawa or mass operators on G_gen", "typed diagonalization maps", "misalignment unitaries between sectors", "readout maps theta13 and J_CKM", "orientation/sign convention explaining sin^2(theta13)/4 - J_CKM"} {
		if !containsAll(a.Required.RequiredObjects, []string{want}) {
			t.Fatalf("missing required object %s in %+v", want, a.Required.RequiredObjects)
		}
	}
	if !strings.Contains(a.Required.LeptonMisalignment, "U_PMNS") || !strings.Contains(a.Required.LeptonReadout, "theta13") || !strings.Contains(a.Required.QuarkMisalignment, "V_CKM") || !strings.Contains(a.Required.QuarkReadout, "J_CKM") {
		t.Fatalf("bad sector readouts: %+v", a.Required)
	}
}

func TestGate789CandidateSourceAudits(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.YukawaTrace.Audited || !closeRel(a.YukawaTrace.NEff, nEffSnapshot, 1e-15) || !a.YukawaTrace.AggregateParticipation || a.YukawaTrace.DeterminesPMNSOrCKM || a.YukawaTrace.SuppliesEigenvectorMisalignment {
		t.Fatalf("bad Yukawa trace audit: %+v", a.YukawaTrace)
	}
	if !a.SingularLedger.Audited || !a.SingularLedger.SingularValuesCanSourceTraces || a.SingularLedger.SingularValuesDetermineMixing || a.SingularLedger.NativeEigenvectorOrientation {
		t.Fatalf("bad singular-value audit: %+v", a.SingularLedger)
	}
	if !a.FiniteTriple.Audited || !a.FiniteTriple.AllowedYukawaEdgeShapes || a.FiniteTriple.GenerationMixingOperatorSourced {
		t.Fatalf("bad finite triple audit: %+v", a.FiniteTriple)
	}
	if !a.K7Polarity.Audited || !a.K7Polarity.SelectorResonance || a.K7Polarity.DefinesGenerationMixingOperator || a.K7Polarity.QuarterWeightDerivesTheta13 {
		t.Fatalf("bad K7 polarity audit: %+v", a.K7Polarity)
	}
	if !a.FockSelector.Audited || !a.FockSelector.FutureGenerationCandidate || a.FockSelector.TypedSelectorToPMNSCKMMap {
		t.Fatalf("bad Fock/projective selector audit: %+v", a.FockSelector)
	}
	if !a.Triality.Audited || !a.Triality.ThreefoldRelevantCandidate || a.Triality.SuppliesSectorOperators || a.Triality.SuppliesRelativeOrientations || a.Triality.SuppliesPhaseData || a.Triality.SuppliesMixingReadoutMaps || a.Triality.SectorMisalignmentOperatorFound {
		t.Fatalf("bad triality audit: %+v", a.Triality)
	}
	if !a.BoundaryData.Audited || !a.BoundaryData.SmallCorrectionToReadout || a.BoundaryData.DerivesFlavorMixing {
		t.Fatalf("bad boundary data audit: %+v", a.BoundaryData)
	}
}

func TestGate789SealMinimalityRuntimePropagationAndBranch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Seal.Defined || a.Seal.Name != "GenerationMixingOperatorSeal" || a.Seal.Native || !strings.Contains(a.Seal.Readout, "FlavorOrientationReadoutSeal") {
		t.Fatalf("bad seal: %+v", a.Seal)
	}
	if !containsAll(a.Seal.Components, []string{"G_gen", "U_PMNS", "V_CKM", "readout maps theta13 and J_CKM", "orientation/sign convention"}) {
		t.Fatalf("bad seal components: %+v", a.Seal.Components)
	}
	if !a.Minimality.Audited || !a.Minimality.Minimal || !strings.Contains(a.Minimality.RemoveEffects["U_PMNS"], "theta13") || !strings.Contains(a.Minimality.RemoveEffects["V_CKM"], "J_CKM") {
		t.Fatalf("bad minimality: %+v", a.Minimality)
	}
	if !a.Runtime.Audited || a.Runtime.ContainsForbidden || !a.Runtime.RuntimeTargetIndependent || a.Runtime.TheoremLevelIndependent {
		t.Fatalf("bad runtime audit: %+v", a.Runtime)
	}
	if !a.Propagation.Recorded || !strings.Contains(a.Propagation.KappaOrient, "GenerationMixingOperatorSeal") || !strings.Contains(a.Propagation.CHistory, "Level B") || !strings.Contains(a.Propagation.CHiggs, "not Level C") {
		t.Fatalf("bad propagation: %+v", a.Propagation)
	}
	if !a.Branch.Recorded || a.Branch.Selected != "failure branch" || !strings.Contains(a.Branch.FailureBranch, "C_Higgs Dependency Freeze") {
		t.Fatalf("bad branch decision: %+v", a.Branch)
	}
}

func TestGate789FirewallsFinalStatementAndTheoremStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewalls.Enforced || a.Firewalls.NEffPMNSCKMTheorem || a.Firewalls.YukawaSingularValuesMixingTheorem || a.Firewalls.K7PolarityMixingTheorem || a.Firewalls.RadialQuarterTheta13Theorem || a.Firewalls.ProjectiveSelectorPMNSCKMTheorem || a.Firewalls.TrialityPMNSCKMTheorem || a.Firewalls.BoundaryPairFlavorMixingTheorem || a.Firewalls.KappaOrientNativeFlavorTheorem || a.Firewalls.FlavorOrientationSealNative || a.Firewalls.TreeProxyPoleMass {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	if !strings.Contains(a.FinalStatement, "does not source theta13 or J_CKM natively") || !strings.Contains(a.FinalStatement, "GenerationMixingOperatorSeal") || !strings.Contains(a.FinalStatement, "failure branch") {
		t.Fatalf("bad final statement: %s", a.FinalStatement)
	}
	res := Generation2GenerationMixingOperatorSourceAndFlavorOrientationReadoutSealAuditTheorem().Verify()
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
