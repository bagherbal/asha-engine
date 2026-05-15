package generation2coefficientledger

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2SectorCoefficientSourceLedgerAmplitudeFirewallClosureTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 sector-coefficient source ledger and amplitude firewall closure"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate447 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate444/Gate445/Gate446 structural boundary", Passed: a.Inheritance.Executed && a.Inheritance.Gate444KGenForced && a.Inheritance.Gate444Generation2BareZero && a.Inheritance.Gate445XSupportForced && a.Inheritance.Gate445AmplitudeSealed && a.Inheritance.Gate446SignedCycleSealed && a.Inheritance.Gate446ComplexPhaseSealed && a.Inheritance.Gate446YGenQuarantined && a.Inheritance.NoEmpiricalInputsImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "K/X/Y charged coefficient arena formalized", Passed: a.Arena.Executed && a.Arena.KAxisGeometricallyForced && a.Arena.XSupportGeometricallyForced && !a.Arena.YQuadratureNative && a.Arena.HermitianFamilySources && a.Arena.TraceNeutralBasis && a.Arena.TotalSymbolicCoefficients == KXYCoeffDim, Detail: FormatArena(a.Arena)},
			{Name: "all native boundaries applied but none selects amplitudes", Passed: len(a.Boundaries) == 5 && boundariesPassWithoutSelection(a.Boundaries), Detail: FormatBoundary(a.Boundaries[0]) + " | " + FormatBoundary(a.Boundaries[1]) + " | " + FormatBoundary(a.Boundaries[2]) + " | " + FormatBoundary(a.Boundaries[3]) + " | " + FormatBoundary(a.Boundaries[4])},
			{Name: "functional selector sieve finds no native coefficient rule", Passed: len(a.Functionals) == 5 && functionalsDoNotPredict(a.Functionals), Detail: FormatFunctional(a.Functionals[0]) + " | " + FormatFunctional(a.Functionals[1]) + " | " + FormatFunctional(a.Functionals[2])},
			{Name: "counter-ledger witnesses prove underdetermination", Passed: a.Sieve.Executed && !a.Sieve.UniqueCoefficientLedger && a.Sieve.SurvivingLedgers >= 3 && a.Sieve.DistinctSurvivors >= 3 && !a.Sieve.ForcesUniversalSectorRay && !a.Sieve.ForcesKCoefficientValues && !a.Sieve.ForcesXCoefficientValues && !a.Sieve.ForcesYCoefficientValues, Detail: FormatSieve(a.Sieve)},
			{Name: "nine charged K/X/Y coefficients remain quarantined", Passed: a.Ledger.Executed && a.Ledger.TotalSymbols == KXYCoeffDim && a.Ledger.NativeCoefficientValues == 0 && a.Ledger.QuarantinedCoefficientDim == KXYCoeffDim && a.Ledger.KAxisForced && a.Ledger.XSupportForced && a.Ledger.YQuadratureQuarantined && a.Ledger.AmplitudeValuesSealed && !a.Ledger.PhysicalMassesPredicted && !a.Ledger.CKMPredicted && !a.Ledger.PMNSPredicted, Detail: FormatCoefficientLedger(a.Ledger)},
			{Name: "amplitude firewall formally closed", Passed: a.Closure.Executed && a.Closure.NativeFlavorDimAfter == NativeFlavorDim && a.Closure.KXYCoeffDimAfter == KXYCoeffDim && !a.Closure.NativeReductionBelow13 && !a.Closure.CoefficientReductionBelow9 && a.Closure.KGenStructuralAxiomPreserved && a.Closure.XSupportStructuralPreserved && !a.Closure.YGenPromotedNative && !a.Closure.AnyCoefficientAxiomPromoted && a.Closure.AmplitudeFirewallClosed, Detail: FormatClosure(a.Closure)},
			{Name: "empirical flavor firewall preserved", Passed: a.Firewall.Executed && a.Firewall.NoObservedMuonMassImported && a.Firewall.NoObservedCharmMassImported && a.Firewall.NoObservedYukawaImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoPoleMassFit && a.Firewall.NoCurveFit && a.Firewall.KGenNative && a.Firewall.XSupportNative && a.Firewall.YGenQuarantined && a.Firewall.CoefficientsQuarantined, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate reconciles publication atlas", Passed: a.Next.Gate == 448 && a.Next.Title == "Post-444 Flavor Frontier Atlas Reconciliation", Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: checks, Notes: []string{StatusFailedMultipleLedgersSurvive, StatusNineCoefficientsRemainQuarantined, StatusCoefficientFirewallClosed, a.Truth}}
	}}
}

func boundariesPassWithoutSelection(xs []Boundary) bool {
	if len(xs) == 0 {
		return false
	}
	for _, x := range xs {
		if !x.Applied || !x.Passed || x.SelectsCoefficientValues {
			return false
		}
	}
	return true
}

func functionalsDoNotPredict(xs []FunctionalAudit) bool {
	if len(xs) == 0 {
		return false
	}
	for _, x := range xs {
		if !x.Executed || x.PredictsMassValues || x.PredictsMixingAngles {
			return false
		}
	}
	return true
}
