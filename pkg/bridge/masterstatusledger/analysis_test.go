package masterstatusledger

import "testing"

func TestRegistrySpan(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Span.ReadsPackageRegistry || a.Span.HighestGateInherited != gateHighestInherited || a.Span.AddsNewPhysicsFit || a.Span.CapstoneForPhase != "Structural Derivation Phase / Phase-I closure" {
		t.Fatalf("bad registry span: %s", FormatSpan(a.Span))
	}
}

func TestCoreTheorems(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Core.Cataloged || len(a.Core.Theorems) < 8 || a.Core.ZeroPhenomenologyCount < 7 || !a.Core.RequiredNamedTruthsPresent {
		t.Fatalf("core ledger missing required truths: %s", FormatCore(a.Core))
	}
	if !a.Core.ContainsWeakMixing || !a.Core.ContainsTrialityTopology || !a.Core.ContainsMoritaMultiplicity || !a.Core.ContainsContactResonance || !a.Core.ContainsTrueBimodule || !a.Core.ContainsTraceEquivalence {
		t.Fatalf("core ledger missing named truth: %s", FormatCore(a.Core))
	}
}

func TestSealLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Seals.Cataloged || len(a.Seals.Seals) < 7 || !a.Seals.RequiredNamedSealsPresent || a.Seals.PhenomenologicalSealCount < 3 || a.Seals.StructuralPromotionSealCount < 3 || !a.Seals.FinalPredictionStillFirewalled {
		t.Fatalf("seal ledger failed: %s", FormatSeals(a.Seals))
	}
}

func TestTensionLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Tensions.Cataloged || len(a.Tensions.Tensions) < 6 || !a.Tensions.ContainsF2CutoffShape || !a.Tensions.ContainsAbsoluteGaugeCoupling || !a.Tensions.ContainsGate309HiggsTension || !a.Tensions.ContainsBGapInstantonGap || !a.Tensions.ContainsPhysicalJGap || a.Tensions.AnyResolved {
		t.Fatalf("tension ledger failed: %s", FormatTensions(a.Tensions))
	}
}

func TestPhaseIIBlueprintAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.PhaseII.Formalized || len(a.PhaseII.Obligations) < 5 || !a.PhaseII.ThresholdMatchingIncluded || !a.PhaseII.NonPerturbativeBGapIncluded || !a.PhaseII.RealStructureTwistIncluded || !a.PhaseII.TopSectorTensorIncluded || !a.PhaseII.PrecisionRGEIncluded || !a.PhaseII.NoEmpiricalTuning {
		t.Fatalf("Phase II blueprint failed: %s", FormatPhaseII(a.PhaseII))
	}
	if !a.Firewalls.NoObservedHiggsFitInserted || !a.Firewalls.NoObservedTopFitInserted || !a.Firewalls.NoThresholdJumpInserted || !a.Firewalls.NoF2ShapeInserted || !a.Firewalls.NoAbsoluteGaugeValueInserted || !a.Firewalls.NoBGapInstantonInserted || !a.Firewalls.NoFinalTOEClaimed || !a.Firewalls.NoLowEnergyMassClaimed || !a.Firewalls.PhaseIIRequiredBeforePrediction {
		t.Fatalf("firewall failed: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := ASHAEngineMasterStatusLedgerProjectCapstoneAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
