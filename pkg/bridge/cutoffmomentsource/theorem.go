package cutoffmomentsource

import "github.com/bagherbal/asha-engine/pkg/theorem"

func CutoffMomentSourcePositiveF0TestFunctionClassAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-CUTOFF-MOMENT-SOURCE-POSITIVE-F0-TEST-FUNCTION-CLASS-AUDIT"
	const name = "Cutoff Moment Source / Positive f0 Test-Function Class Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 303 cutoff moment source audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 302 positive f0 obligation is inherited without numerical f0 or ZH claim", Passed: a.Input.PositivePrefactorLedgerFormalized && a.Input.KRawPositiveCarrierInherited && a.Input.N4PositiveClassAvailable && a.Input.F0PositiveRequired && !a.Input.F0NumericalValueDerived && !a.Input.CutoffGateActivated && !a.Input.NumericalZHComputed && !a.Input.YukawaNumbersInserted, Detail: FormatGate302Inheritance(a.Input)},
			{Name: "generic positive spectral-action test-function class guarantees f0 sign but not value", Passed: a.Generic.GuaranteesF0Positive && !a.Generic.FixesNumericalF0 && !a.Generic.ObservedInputUsed && a.Generic.PositivityCondition != "", Detail: FormatGeneric(a.Generic)},
			{Name: "contact-spectral candidate zeta_contact(0)=7 is exact, internal, positive, and only preflight-sealed", Passed: a.Contact.Gate162LedgerAvailable && a.Contact.Gate288CutoffIdentificationAudited && a.Contact.IntegerValue == 7 && a.Contact.StrictlyPositive && a.Contact.InternalAlgebraicSource && !a.Contact.ObservedInputUsed && a.Contact.MayBeActivatedAsSeal && !a.Contact.ActivatedAsFinalSource && !a.Contact.HeatKernelEqualityDerived && a.Contact.SatisfiesGate302SignRequirement && a.Contact.DoesNotDeriveHiggsPrediction, Detail: FormatContact(a.Contact)},
			{Name: "free phenomenological f0 preserves stability only by external positive-domain restriction and loses internal prediction", Passed: a.Free.GuaranteesF0PositiveIfImposed && !a.Free.FixesNumericalF0 && a.Free.InternalPredictionLost && a.Free.ExternalExperimentNeeded && a.Free.AdmissibleForStabilityOnly && len(a.Free.PredictiveLosses) >= 4, Detail: FormatFree(a.Free)},
			{Name: "three source lanes are compared without forcing a unique final cutoff source", Passed: len(a.Comparison.Candidates) == 3 && a.Comparison.AnyPositiveLaneAvailable && a.Comparison.ContactLaneSatisfiesSign && a.Comparison.GenericLaneSatisfiesSign && a.Comparison.FreeLaneSatisfiesSign && !a.Comparison.UniqueFinalSourceSelected && a.Comparison.NoObservedInputRequired, Detail: FormatComparison(a.Comparison)},
			{Name: "positive f0 test-function class is formalized while final numerical f0, ZH, and Higgs prediction remain unclaimed", Passed: a.Sieve.StrictPositivityCanBeEnsured && a.Sieve.ContactValueCleanlySatisfies && !a.Sieve.FinalNumericalF0Claimed && !a.Sieve.NumericalZHClaimed && !a.Sieve.HiggsPredictionClaimed, Detail: FormatSieve(a.Sieve)},
			{Name: "empirical, final-source, Yukawa, ZH, Higgs, gauge-coupling, and B-gap firewalls are preserved", Passed: a.Firewalls.NoObservedF0Inserted && a.Firewalls.NoFinalCutoffSourceForced && a.Firewalls.NoYukawaNumbersInserted && a.Firewalls.NoNumericalZHComputed && a.Firewalls.NoHiggsMassQuarticClaimed && a.Firewalls.NoGaugeCouplingAbsoluteClaimed && a.Firewalls.NoBGapInstantonClaimed && a.Firewalls.ContactValueUsedOnlyAsSealedPreflight && !a.Firewalls.FiniteCorePolluted, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "summary records positive f0 source classes without declaring physical dynamics", Passed: a.Summary.Gate302Inherited && a.Summary.GenericClassFormalized && a.Summary.ContactSealPositive && a.Summary.FreeF0SieveCompleted && a.Summary.PositiveF0ClassFormalized && !a.Summary.UniqueFinalSourceSelected && !a.Summary.NumericalF0Locked && !a.Summary.NumericalZHComputed && !a.Summary.PhysicalDynamicsDerived && a.Summary.FirewallPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 303 proves f0 positivity can be supplied by admissible positive test-function classes and that the sealed contact value ζ_contact(0)=7 satisfies the sign obligation. It does not activate the contact cutoff as the final physical heat-kernel source.", "The remaining hard gate is source promotion: construct a canonical positive cutoff profile/spectral measure whose a4 moment is the contact value without breaking existing firewalls."}}
	}}
}
