package hphivariationalselector

import "github.com/bagherbal/asha-engine/pkg/theorem"

func HphiVariationalFunctionalCanonicalCoefficientSelectorSieveTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Hphi variational functional / canonical coefficient selector sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate408 audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "inherits Gate407 Hphi capacity boundary", Passed: a.Inheritance.Executed && a.Inheritance.Gate407FullAlgebraCapacity && a.Inheritance.Gate407NoCanonicalSelector && a.Inheritance.Gate407PairDegenerateSelectedObservables && a.Inheritance.Gate407ChargedModuliPreserved && a.Inheritance.Gate372ChargedModuliDim == Gate372ChargedModuliDim && a.Inheritance.NoEmpiricalInputsImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "native variational functional ledger audited", Passed: a.Ledger.Executed && a.Ledger.HphiDimension == HphiRealDim && a.Ledger.NativeFunctionalCount == 4 && a.Ledger.VariationalFunctionalCount == 5 && a.Ledger.ExternalSourceCount == 1 && a.Ledger.NondegenerateNativeSelectors == 0 && a.Ledger.NoObservedInputs && a.Ledger.NoYukawaInputs && a.Ledger.NoArbitrarySourcesPromoted, Detail: FormatLedger(a.Ledger)},
			{Name: "spectral Hessian selects pair-degenerate scalar response", Passed: functionalByName(a.Ledger.Functionals, "spectral-action Hessian on H_phi").SelectedElementPairDegenerate && functionalByName(a.Ledger.Functionals, "spectral-action Hessian on H_phi").SelectedMinimalDegree == 2, Detail: FormatFunctional(functionalByName(a.Ledger.Functionals, "spectral-action Hessian on H_phi"))},
			{Name: "scalar potential fixes radius but no orientation", Passed: !functionalByName(a.Ledger.Functionals, "radial scalar potential normal form").SelectedElementUnique && functionalByName(a.Ledger.Functionals, "radial scalar potential normal form").MinimizerFamilyDimension == 3, Detail: FormatFunctional(functionalByName(a.Ledger.Functionals, "radial scalar potential normal form"))},
			{Name: "one-form kinetic trace has degenerate minimizer family", Passed: !functionalByName(a.Ledger.Functionals, "one-form kinetic trace / complex-compatibility penalty").SelectedElementUnique && functionalByName(a.Ledger.Functionals, "one-form kinetic trace / complex-compatibility penalty").SelectedElementPairDegenerate, Detail: FormatFunctional(functionalByName(a.Ledger.Functionals, "one-form kinetic trace / complex-compatibility penalty"))},
			{Name: "quaternionic invariant trace is central", Passed: functionalByName(a.Ledger.Functionals, "quaternionic-invariant trace/norm functional").SelectedElementCentral && functionalByName(a.Ledger.Functionals, "quaternionic-invariant trace/norm functional").SelectedMinimalDegree == 1, Detail: FormatFunctional(functionalByName(a.Ledger.Functionals, "quaternionic-invariant trace/norm functional"))},
			{Name: "generic source functional is sealed external source", Passed: !functionalByName(a.Ledger.Functionals, "sealed generic source functional stress test").Native && functionalByName(a.Ledger.Functionals, "sealed generic source functional stress test").UsesExternalSource && functionalByName(a.Ledger.Functionals, "sealed generic source functional stress test").NondegenerateCapacity, Detail: FormatFunctional(functionalByName(a.Ledger.Functionals, "sealed generic source functional stress test"))},
			{Name: "variational outcome rejects native nondegenerate selector", Passed: !a.Outcome.NativeNondegenerateSelector && a.Outcome.OnlyCentralOrPairSelected && a.Outcome.GenericSourceWouldSelectAnyElement && !a.Outcome.GenericSourcePromoted && a.Outcome.HphiScalarLaneFlavorBlind, Detail: FormatOutcome(a.Outcome)},
			{Name: "moduli impact preserves charged flavor firewall", Passed: a.Impact.ChargedModuliStart == Gate372ChargedModuliDim && a.Impact.ChargedModuliResult == Gate372ChargedModuliDim && !a.Impact.NativeNondegenerateSelector && !a.Impact.YukawaCouplingsReduced && !a.Impact.CKMCapacityDerived && a.Impact.FlavorFirewallPreserved, Detail: FormatImpact(a.Impact)},
			{Name: "empirical/source firewalls remain clean", Passed: a.Firewall.Executed && a.Firewall.NoObservedMassesImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoYukawaAmplitudesInserted && a.Firewall.NoExternalSourcePromoted && a.Firewall.NoArbitraryCoefficientPromoted && a.Firewall.NoGenericMatrixPromoted && a.Firewall.NoFlavorModuliReductionClaimed, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate classifies Yukawa amplitude/source seal", Passed: a.Next.Gate == 409 && a.Next.Title != "", Detail: FormatNext(a.Next)},
		}, Notes: []string{a.Truth}}
	}}
}

func functionalByName(xs []Functional, name string) Functional {
	for _, x := range xs {
		if x.Name == name {
			return x
		}
	}
	return Functional{}
}
