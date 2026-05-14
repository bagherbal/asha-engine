package heatkernelconventionledger

import "github.com/bagherbal/asha-engine/pkg/theorem"

func HeatKernelConventionLedgerPositivePrefactorNormalizationAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-HEAT-KERNEL-CONVENTION-LEDGER-POSITIVE-PREFACTOR-NORMALIZATION-AUDIT"
	const name = "Heat-Kernel Convention Ledger / Positive Prefactor Normalization Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 302 heat-kernel convention ledger", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 301 positive scalar trace carrier inherited without numerical Z_H claim", Passed: a.Input.PositiveKRawCarrierProved && a.Input.HilbertSchmidtSumStructure && a.Input.StrictPositiveNeedsNonzeroEdge && !a.Input.NumericalYukawasInserted && !a.Input.NumericalZHComputed, Detail: FormatInput(a.Input)},
			{Name: "heat-kernel prefactor ledger isolates every scalar kinetic normalization factor", Passed: a.Ledger.AllFactorsExplicit && a.Ledger.AllEmpiricalInputsExcluded && a.Ledger.CanChoosePositiveClass && !a.Ledger.AbsoluteN4Derived && len(a.Ledger.Factors) == 6, Detail: FormatLedger(a.Ledger)},
			{Name: "all prefactor factors have explicit sign conditions and no empirical payload", Passed: allFactorsSignAudited(a.Ledger), Detail: FormatLedger(a.Ledger)},
			{Name: "Wick/sign convention maps Euclidean positivity to Lorentzian positive-energy kinetics without hidden sign ambiguity", Passed: a.Wick.PositiveEnergyMapped && !a.Wick.SignAmbiguityHidden && !a.Wick.ConventionNativeToFinite && len(a.Wick.SignLedger) >= 4, Detail: FormatWick(a.Wick)},
			{Name: "positive f0 requirement is formalized without numerical cutoff activation", Passed: a.F0.ConditionallyPositive && a.F0.CanBePositiveWithoutEmpirics && !a.F0.NumericalValueDerived && !a.F0.ContactSpectralGate288Used, Detail: FormatF0(a.F0)},
			{Name: "canonical scalar matching rule absorbs N4 f0 Kraw into Hraw to Hphys rescaling", Passed: a.Matching.RuleFormalized && !a.Matching.PhysicalZHComputed && len(a.Matching.AbsorbedFactors) == len(a.Ledger.Factors), Detail: FormatMatching(a.Matching)},
			{Name: "positive prefactor sieve proves sign-safe convention class but keeps strict ZH numerical proof conditional", Passed: a.Positivity.KRawPositiveSemidefinite && a.Positivity.PositivePrefactorAvailable && a.Positivity.StrictZHGuaranteedConditionally && !a.Positivity.NumericalStrictZHProved, Detail: FormatPositivity(a.Positivity)},
			{Name: "empirical, cutoff, subtraction, Higgs prediction, and B-gap firewalls are preserved", Passed: a.Firewalls.NoF0NumberInserted && a.Firewalls.NoCutoffGateActivated && a.Firewalls.NoYukawaNumbersInserted && a.Firewalls.NoObservedMassesInserted && a.Firewalls.NoSubtractionSchemeInvented && a.Firewalls.NoBGapInstantonClaimed && a.Firewalls.NoHiggsPredictionClaimed && !a.Firewalls.FiniteCorePolluted, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "summary records positive convention ledger without overclaiming dynamics", Passed: a.Summary.Gate301Inherited && a.Summary.PrefactorLedgerFormalized && a.Summary.WickSignRuleFormalized && a.Summary.PositiveF0ConditionRecorded && a.Summary.CanonicalMatchingFormalized && a.Summary.PositivePrefactorAvailable && !a.Summary.StrictZHNumericallyProved && !a.Summary.PhysicalDynamicsDerived && a.Summary.FirewallPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 302 formalizes a sign-safe convention class for N_4 f_0. It does not derive a numerical f_0, absolute N_4, numerical Yukawa amplitudes, Higgs mass/quartic, or B-gap instanton action.", "Strict physical Z_H positivity is now reduced to: positive convention ledger, f_0>0, and at least one nonzero sealed scalar Dirac-edge amplitude."}}
	}}
}

func allFactorsSignAudited(l PrefactorLedger) bool {
	if len(l.Factors) == 0 {
		return false
	}
	for _, f := range l.Factors {
		if f.SignCondition == "" || f.EmpiricalInput || (!f.PositiveByChoice && !f.PositiveByTheorem) {
			return false
		}
	}
	return true
}
