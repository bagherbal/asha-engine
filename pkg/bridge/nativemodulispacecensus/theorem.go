package nativemodulispacecensus

import "github.com/bagherbal/asha-engine/pkg/theorem"

func NativeModuliSpaceDimensionExactDiracParameterCensusSieveTheorem() theorem.Theorem {
	const id = "BRIDGE-NATIVE-MODULI-SPACE-DIMENSION-EXACT-DIRAC-PARAMETER-CENSUS-SIEVE"
	const name = "Native Moduli Space Dimension / Exact Dirac Parameter Census Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 372 audit", Passed: false, Detail: err.Error()}}}
		}
		quark := FindScenario(a.Quotient.Scenarios, "quark Yukawa sector")
		charged := FindScenario(a.Quotient.Scenarios, "charged-lepton-only sector")
		minimal := FindScenario(a.Quotient.Scenarios, "minimal charged finite-Dirac flavor sector")
		diracExtended := FindScenario(a.Quotient.Scenarios, "quark plus Dirac-neutrino finite-Dirac sector")
		majoranaExtended := FindScenario(a.Quotient.Scenarios, "quark plus Majorana finite-Dirac sector")
		checks := []theorem.Check{
			{Name: "Gate 371 obstruction is inherited and the question is reframed", Passed: a.Inheritance.Executed && a.Inheritance.HighestInheritedGate == 371 && a.Inheritance.DynamicalIntertwinerStillOpen, Detail: FormatInheritance(a.Inheritance)},
			{Name: "general finite Dirac blocks are parameterized", Passed: a.Parameterization.Executed && len(a.Parameterization.Blocks) == 5 && a.Parameterization.MinimalChargedRawDim == 54 && a.Parameterization.DiracNeutrinoRawDim == 72 && a.Parameterization.MajoranaExtendedRawDim == 84, Detail: FormatParameterization(a.Parameterization)},
			{Name: "spectral-triple axioms impose edge shape but no hidden generation texture", Passed: a.Axioms.Executed && a.Axioms.JRealityImposesMirrorBlocks && a.Axioms.ChiralityAllowsOnlyOddEdges && a.Axioms.FirstOrderEnforcesEdgeGraph && a.Axioms.MajoranaBlockSymmetric && a.Axioms.AdditionalGenerationConstraints == 0, Detail: FormatAxioms(a.Axioms)},
			{Name: "quark quotient gives ten physical flavor moduli", Passed: quark.RawRealDim == 36 && quark.BasisGroupDim == 27 && quark.GenericStabilizerDim == 1 && quark.PhysicalDim == 10, Detail: FormatScenario(quark)},
			{Name: "charged-lepton-only quotient gives three singular values", Passed: charged.RawRealDim == 18 && charged.BasisGroupDim == 18 && charged.GenericStabilizerDim == 3 && charged.PhysicalDim == 3, Detail: FormatScenario(charged)},
			{Name: "minimal charged finite-Dirac moduli dimension is thirteen", Passed: minimal.RawRealDim == 54 && minimal.BasisGroupDim == 45 && minimal.GenericStabilizerDim == 4 && minimal.PhysicalDim == 13, Detail: FormatScenario(minimal)},
			{Name: "Dirac-neutrino extension has twenty finite-Dirac flavor moduli", Passed: diracExtended.RawRealDim == 72 && diracExtended.PhysicalDim == 20, Detail: FormatScenario(diracExtended)},
			{Name: "all-allowed Majorana finite-Dirac census has thirty-one moduli", Passed: majoranaExtended.RawRealDim == 84 && majoranaExtended.BasisGroupDim == 54 && majoranaExtended.GenericStabilizerDim == 1 && majoranaExtended.PhysicalDim == 31, Detail: FormatScenario(majoranaExtended)},
			{Name: "algebra gauge quotient is separated from generation-basis quotient", Passed: a.Quotient.Executed && a.Quotient.AlgebraGaugeRemovesGenerationDim == 0 && a.Quotient.FlavorBasisQuotientRequired && a.Quotient.MinimalChargedDFDim == 13 && a.Quotient.MajoranaSeesawDFDim == 31, Detail: FormatQuotient(a.Quotient)},
			{Name: "external fifteen is refined rather than blindly equated with dim M(D_F)", Passed: a.Native.Executed && a.Native.NPhysicalDF == 31 && a.Native.MinimalChargedDFDim == 13 && a.Native.External15 == 15 && !a.Native.NativeReductionBelow15 && a.Native.HiddenCrossSectorConstraints == 0, Detail: FormatNative(a.Native)},
			{Name: "epistemic firewalls are preserved", Passed: a.Firewall.Executed && a.Firewall.NoYukawaValuesImported && a.Firewall.NoCKMValuesImported && a.Firewall.NoPMNSValuesImported && a.Firewall.NoMassValuesImported && a.Firewall.NoVacuumDirectionForced && a.Firewall.NoGaugeFlavorConflation && a.Firewall.NoMajoranaMinimalConflation && a.Firewall.LandscapeRatiosPreserved, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 372 does not select a vacuum point; it replaces ambiguous external counting with an internal finite-Dirac census and a category-correct comparison to the minimal 15-input ledger."}}
	}}
}
