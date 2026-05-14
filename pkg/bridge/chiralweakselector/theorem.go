package chiralweakselector

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ChiralAlignmentWeakPlaneSelectorAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-CHIRAL-ALIGNMENT-WEAK-PLANE-SELECTOR-AUDIT"
	const name = "Chiral alignment gamma and weak plane selector audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 238 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 237 local weak/H support is inherited but unresolved", Passed: a.Previous.Summary.PseudoRealLocalHSupport && !a.Previous.Summary.CanonicalWeakPlane && !a.Previous.Summary.GlobalHDerived, Detail: a.Previous.TruthStatement},
			{Name: "native gamma parity grading exists but is not physical chirality", Passed: a.Gamma.EvenStateDimC == 8 && a.Gamma.OddStateDimC == 8 && a.Gamma.RetrievedFromGate233 && !a.Gamma.EquatedToSMChirality && !a.Gamma.PhysicalChiralityDerived, Detail: FormatGamma(a.Gamma)},
			{Name: "all six planes have mixed-parity doublet sectors", Passed: len(a.Planes) == 6 && a.Sieve.UniformDoubletPlanes == 0 && a.Sieve.UniformSingletPlanes == 0 && a.Sieve.AllPlanesSameCounts && !a.Sieve.GammaBreaksDegeneracy, Detail: FormatSieve(a.Sieve) + " :: " + FormatPlanes(a.Planes)},
			{Name: "temporal/spatial split distinguishes classes but not a unique plane", Passed: a.Temporal.ClassDistinctionExists && a.Temporal.TemporalSpatialPlaneCount == 3 && a.Temporal.PureSpatialPlaneCount == 3 && !a.Temporal.UniquePlaneSelected, Detail: FormatTemporal(a.Temporal)},
			{Name: "physical chiral weak action remains unselected", Passed: a.Weak.CandidateLocalHSupportInherited && !a.Weak.GammaSelectorWorks && !a.Weak.TemporalSpatialSelectorWorks && !a.Weak.ContactSU2PlaneMapDerived && !a.Weak.HyperchargeColorAttachment && !a.Weak.GlobalHSummandDerived && !a.Weak.PhysicalLeftHandedActionDerived, Detail: FormatWeak(a.Weak)},
			{Name: "firewall blocks forced left-handed weak assignment", Passed: !a.Firewall.ForcedLeftHandedAssignment && !a.Firewall.ForcedWeakPlane && !a.Firewall.ImportedSMChirality && !a.Firewall.ImportedPauliMatrices && !a.Firewall.ImportedConnesAlgebra && !a.Firewall.ClaimedGlobalH && !a.Firewall.ClaimedOrderOne && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "summary records obstruction without killing local H support", Passed: a.Summary.GammaParityAvailable && !a.Summary.UniformChiralDoublets && !a.Summary.GammaSelectsPlane && a.Summary.TemporalSpatialClasses && !a.Summary.UniqueWeakPlaneDerived && !a.Summary.PhysicalLeftActionDerived && !a.Summary.GlobalHDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}
		notes := []string{
			"Gate 238 tests the proposed chirality selector directly: occupation parity γ splits S_C into 8 even and 8 odd states, but each candidate weak plane has doublet states in both parity sectors.",
			"The lifted exterior su(2) preserves γ rather than selecting one γ-sector. Therefore raw Fock parity cannot be identified with Standard Model left-handed chirality at this stage.",
			"The temporal/spatial 1⊕3 split creates two conjugacy classes of candidate planes, but each class has three members. A further selector or physical chirality theorem is still required.",
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
