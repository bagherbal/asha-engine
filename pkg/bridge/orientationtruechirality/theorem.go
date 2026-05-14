package orientationtruechirality

import "github.com/bagherbal/asha-engine/pkg/theorem"

func OrientationOperatorTrueChiralityDerivationAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-ORIENTATION-TRUE-CHIRALITY-DERIVATION-AUDIT"
	const name = "Orientation operator chi and true chirality derivation audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 239 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 238 failure is inherited as the correct starting point", Passed: a.Previous.Summary.GammaParityAvailable && !a.Previous.Summary.GammaSelectsPlane && !a.Previous.Summary.PhysicalLeftActionDerived, Detail: a.Previous.TruthStatement},
			{Name: "Clifford-volume orientation candidate acts on S_C but is not distinct from gamma", Passed: a.Orientation.ActsOnSC && a.Orientation.VolumeElementAvailable && a.Orientation.EquivalentToGamma && !a.Orientation.DistinctFromGamma && !a.Orientation.ManualSignAdjusted, Detail: FormatOrientation(a.Orientation)},
			{Name: "tau_eta is inherited as orientation data but not as an S_C chirality operator", Passed: a.TauEta.FunctionalOnScalarBundle && !a.TauEta.EndomorphismOnSC && !a.TauEta.CanonicalPullbackDerived && !a.TauEta.CanActAsChiralityOperator, Detail: FormatTauEta(a.TauEta)},
			{Name: "orientation candidate has same eigenspaces as occupation parity", Passed: a.Comparison.SameSpectrum && a.Comparison.SameEigenspaces && a.Comparison.OperatorsCommute && !a.Comparison.OperatorsAntiCommute && !a.Comparison.PhysicalChiralityDerived, Detail: FormatComparison(a.Comparison)},
			{Name: "all six weak planes still have mixed chi-doublet sectors", Passed: len(a.Planes) == 6 && a.Sieve.UniformDoubletPlanes == 0 && a.Sieve.UniformSingletPlanes == 0 && a.Sieve.AllPlanesSameCounts && !a.Sieve.ChiBreaksDegeneracy, Detail: FormatSieve(a.Sieve) + " :: " + FormatPlanes(a.Planes)},
			{Name: "true left-handed weak action and global H summand remain unselected", Passed: a.Weak.Gate238GammaSelectorFailed && !a.Weak.VolumeChiImprovesGate238 && !a.Weak.TauEtaSuppliesOperator && !a.Weak.UniqueWeakPlaneSelected && !a.Weak.PhysicalLeftHandedActionDerived && !a.Weak.GlobalHSummandDerived, Detail: FormatWeak(a.Weak)},
			{Name: "firewall preserved: no forced chi signs, weak plane, or imported chirality", Passed: !a.Firewall.AdjustedChiSignsToFit && !a.Firewall.ImportedSMGamma5 && !a.Firewall.ImportedConnesChirality && !a.Firewall.ForcedWeakPlane && !a.Firewall.PromotedTauEtaToOperator && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "summary records failure without killing local quaternionic support", Passed: a.Summary.VolumeOrientationAvailable && !a.Summary.ChiDistinctFromGamma && !a.Summary.TauEtaPullbackDerived && !a.Summary.UniformChiDoublets && !a.Summary.ChiSelectsPlane && !a.Summary.PhysicalChiralityDerived && !a.Summary.GlobalHDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}
		notes := []string{
			"Gate 239 tests the proposed orientation-selector route directly. The Clifford-volume candidate is legitimate finite orientation data, but in the current exterior/Fock realization it is proportional to occupation parity.",
			"The scalar fundamental class tau_eta contains signed orientation information, but it is a trace functional on the scalar bundle, not a derived endomorphism of the complexified Fock spinor S_C.",
			"Therefore the physical Standard Model chirality operator, unique weak plane, and global quaternionic H summand remain future work; no signs or planes were tuned to force a result.",
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
