package cutoffmomentsource

import "testing"

func TestGate302Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Input.PositivePrefactorLedgerFormalized || !a.Input.KRawPositiveCarrierInherited || !a.Input.N4PositiveClassAvailable || !a.Input.F0PositiveRequired {
		t.Fatalf("bad Gate 302 inheritance: %s", FormatGate302Inheritance(a.Input))
	}
	if a.Input.F0NumericalValueDerived || a.Input.CutoffGateActivated || a.Input.NumericalZHComputed || a.Input.YukawaNumbersInserted {
		t.Fatalf("Gate 303 must inherit numerical firewalls: %s", FormatGate302Inheritance(a.Input))
	}
}

func TestGenericTestFunctionAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Generic.GuaranteesF0Positive || a.Generic.FixesNumericalF0 || a.Generic.ObservedInputUsed || a.Generic.PositivityCondition == "" || a.Generic.MomentDefinition == "" {
		t.Fatalf("bad generic audit: %s", FormatGeneric(a.Generic))
	}
}

func TestContactSpectralCutoffPreflight(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Contact.Gate162LedgerAvailable || !a.Contact.Gate288CutoffIdentificationAudited || a.Contact.IntegerValue != 7 || !a.Contact.StrictlyPositive || !a.Contact.InternalAlgebraicSource || a.Contact.ObservedInputUsed {
		t.Fatalf("bad contact preflight: %s", FormatContact(a.Contact))
	}
	if !a.Contact.MayBeActivatedAsSeal || a.Contact.ActivatedAsFinalSource || a.Contact.HeatKernelEqualityDerived || !a.Contact.SatisfiesGate302SignRequirement || !a.Contact.DoesNotDeriveHiggsPrediction {
		t.Fatalf("contact preflight overclaim or missing sign: %s", FormatContact(a.Contact))
	}
}

func TestFreePhenomenologicalF0Sieve(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Free.GuaranteesF0PositiveIfImposed || a.Free.FixesNumericalF0 || !a.Free.InternalPredictionLost || !a.Free.ExternalExperimentNeeded || !a.Free.AdmissibleForStabilityOnly || len(a.Free.PredictiveLosses) < 4 {
		t.Fatalf("bad free f0 sieve: %s", FormatFree(a.Free))
	}
}

func TestSourceComparison(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if len(a.Comparison.Candidates) != 3 || !a.Comparison.AnyPositiveLaneAvailable || !a.Comparison.ContactLaneSatisfiesSign || !a.Comparison.GenericLaneSatisfiesSign || !a.Comparison.FreeLaneSatisfiesSign || a.Comparison.UniqueFinalSourceSelected || !a.Comparison.NoObservedInputRequired {
		t.Fatalf("bad source comparison: %s", FormatComparison(a.Comparison))
	}
}

func TestPositiveF0ClassSieve(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Sieve.StrictPositivityCanBeEnsured || !a.Sieve.ContactValueCleanlySatisfies || a.Sieve.FinalNumericalF0Claimed || a.Sieve.NumericalZHClaimed || a.Sieve.HiggsPredictionClaimed {
		t.Fatalf("bad positive f0 sieve: %s", FormatSieve(a.Sieve))
	}
}

func TestFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoObservedF0Inserted || !a.Firewalls.NoFinalCutoffSourceForced || !a.Firewalls.NoYukawaNumbersInserted || !a.Firewalls.NoNumericalZHComputed || !a.Firewalls.NoHiggsMassQuarticClaimed || !a.Firewalls.NoGaugeCouplingAbsoluteClaimed || !a.Firewalls.NoBGapInstantonClaimed || !a.Firewalls.ContactValueUsedOnlyAsSealedPreflight || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failure: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestTheoremPassesWithBridgeRequiredStatus(t *testing.T) {
	res := CutoffMomentSourcePositiveF0TestFunctionClassAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
