package contactspectralcutoffpromotion

import "testing"

func TestGate303Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Input.PositiveF0ClassFormalized || !a.Input.ContactCandidateAvailable || a.Input.ContactCandidateValue != 7 || !a.Input.ContactCandidatePositive || a.Input.ContactCandidateObservedInput {
		t.Fatalf("bad Gate 303 inheritance: %s", FormatGate303Inheritance(a.Input))
	}
	if a.Input.FinalSourcePreviouslySelected || a.Input.PromotionTheoremPreviouslyDerived || a.Input.NumericalZHPreviouslyComputed || a.Input.PhysicalDynamicsPreviouslyDerived {
		t.Fatalf("Gate 304 must inherit unresolved source/dynamics state: %s", FormatGate303Inheritance(a.Input))
	}
}

func TestContinuousPositiveProfileFormalization(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Profile.PreservesGate302Sign || !a.Profile.MomentFunctionalPositive || a.Profile.ObservedInputUsed || a.Profile.MomentFunctional == "" || a.Profile.AbstractFunctionalRule == "" || a.Profile.AdmissibleBaseProfileCondition == "" {
		t.Fatalf("bad continuous profile formalization: %s", FormatProfile(a.Profile))
	}
}

func TestCanonicalPositiveProfileConstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Construction.ProfileNonNegative || !a.Construction.ProfileContinuous || !a.Construction.ProfileSmooth || !a.Construction.MomentFinite || !a.Construction.MomentEqualsContactZeta0 || !a.Construction.PreservesPositivity || !a.Construction.CanonicalNormalizationRule {
		t.Fatalf("profile construction did not establish the coefficient bridge: %s", FormatConstruction(a.Construction))
	}
	if a.Construction.UniqueProfileShapeDerived || a.Construction.VariationalPreferenceDerived || a.Construction.UsesObservedInput {
		t.Fatalf("profile construction overclaimed uniqueness/observed input: %s", FormatConstruction(a.Construction))
	}
}

func TestDiscreteContinuousMap(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Map.DiscreteValue != 7 || !a.Map.MapIsAlgebraicallyExact || !a.Map.MapPreservesSign || !a.Map.MapUsesPositiveLinearFunctional || !a.Map.MapRequiresNonZeroBaseMoment || !a.Map.MapIsUniqueAtCoefficientLevel {
		t.Fatalf("bad discrete-to-continuous map: %s", FormatMap(a.Map))
	}
	if a.Map.MapIsUniqueAtProfileShapeLevel || a.Map.ImportsEmpiricalData || a.Map.LocksHigherMoments {
		t.Fatalf("discrete-to-continuous map overclaimed: %s", FormatMap(a.Map))
	}
}

func TestPromotionSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Seal.ActivatedConditionally || a.Seal.PromotedF0Value != 7 || !a.Seal.F0Positive || !a.Seal.F0SourceSelectedBySeal || !a.Seal.ContactSpectrumPromoted || !a.Seal.HeatKernelCoefficientPromoted {
		t.Fatalf("promotion seal did not activate correctly: %s", FormatSeal(a.Seal))
	}
	if a.Seal.ProfileShapePromoted || a.Seal.HigherMomentsPromoted || a.Seal.NumericalZHComputed || a.Seal.HiggsPredictionClaimed || a.Seal.GaugeCouplingAbsoluteClaimed || a.Seal.BGapInstantonClaimed {
		t.Fatalf("promotion seal overclaimed dynamics: %s", FormatSeal(a.Seal))
	}
}

func TestFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoObservedInputInserted || !a.Firewalls.NoYukawaNumbersInserted || !a.Firewalls.NoNumericalZHComputed || !a.Firewalls.NoHiggsMassQuarticClaimed || !a.Firewalls.NoAbsoluteGaugeCouplingsClaimed || !a.Firewalls.NoBGapInstantonClaimed || !a.Firewalls.NoHeatKernelSubtractionClaimed || !a.Firewalls.NoUniqueProfileShapeClaimed || !a.Firewalls.NoHigherMomentLockClaimed || !a.Firewalls.ContactF0PromotedOnlyAsSealedSource || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failure: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestTheoremPassesWithBridgeRequiredStatus(t *testing.T) {
	res := ContactSpectralCutoffPromotionCanonicalPositiveTestProfileConstructionAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
