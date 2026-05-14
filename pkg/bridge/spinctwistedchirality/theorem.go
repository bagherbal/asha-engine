package spinctwistedchirality

import "github.com/bagherbal/asha-engine/pkg/theorem"

func SpinCTwistedChiralityHyperchargeWeakSieveAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-SPINC-TWISTED-CHIRALITY-HYPERCHARGE-WEAK-SIEVE"
	const name = "Spin^c twisted chirality and hypercharge weak-plane sieve audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 240 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 239 orientation route is inherited as failed", Passed: a.Previous.Summary.VolumeOrientationAvailable && !a.Previous.Summary.ChiSelectsPlane && !a.Previous.Summary.PhysicalChiralityDerived, Detail: a.Previous.TruthStatement},
			{Name: "native diagonal u(1) bookkeeping is retrieved without importing SM hypercharge", Passed: a.U1.ActsOnSC && a.U1.SpatialWeightsDegenerate && !a.U1.ImportedSMHypercharge && !a.U1.NativeContactU1AsHypercharge, Detail: FormatU1(a.U1)},
			{Name: "Spin^c-like gamma times u(1) twist is constructed as a diagnostic, not a chirality theorem", Passed: a.Twist.GammaCommutesWithY && a.Twist.IsDiagonalOnFockBasis && !a.Twist.IsInvolution && a.Twist.DistinctFromGamma && !a.Twist.PhysicalChiralityDerived && !a.Twist.ManualHyperchargeFit, Detail: FormatTwist(a.Twist)},
			{Name: "u(1) commutant sieve rejects temporal-spatial planes but leaves pure-spatial degeneracy", Passed: len(a.Planes) == 6 && len(a.Sieve.U1RejectedPlanes) == 3 && len(a.Sieve.U1PreservingPlanes) == 3 && a.Sieve.TemporalSpatialRejected && a.Sieve.PureSpatialDegeneracy == 3, Detail: FormatSieve(a.Sieve) + " :: " + FormatPlanes(a.Planes)},
			{Name: "twisted chirality still gives no uniform doublet alignment or unique weak plane", Passed: len(a.Sieve.UniformDoubletPlanes) == 0 && len(a.Sieve.SelectedPlanes) == 0 && !a.Sieve.TwistBreaksDegeneracy && !a.Summary.UniqueWeakPlaneDerived, Detail: FormatSieve(a.Sieve)},
			{Name: "physical left-handed weak action and global H summand remain unselected", Passed: a.Weak.Gate239ChiFailed && a.Weak.U1TwistImprovesClassSieve && !a.Weak.UniqueWeakPlaneSelected && !a.Weak.PhysicalLeftHandedDerived && !a.Weak.GlobalHSummandDerived && !a.Weak.OrderOneReady, Detail: FormatWeak(a.Weak)},
			{Name: "firewall preserved: no tuned hypercharge, forced plane, or imported Spin^c structure", Passed: !a.Firewall.ImportedSMHypercharge && !a.Firewall.TunedU1Weights && !a.Firewall.ForcedWeakPlane && !a.Firewall.ForcedLeftHandedAction && !a.Firewall.ImportedSpinCStructure && !a.Firewall.ClaimedGlobalH && !a.Firewall.ClaimedOrderOne && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "summary records class-sieve progress but not physical chirality", Passed: a.Summary.NativeU1Available && a.Summary.TwistConstructed && a.Summary.U1RejectsTemporalPlanes && a.Summary.PureSpatialPlanesRemain == 3 && !a.Summary.UniformTwistedDoublets && !a.Summary.PhysicalChiralityDerived && !a.Summary.GlobalHDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}
		notes := []string{
			"Gate 240 tests the Spin^c intuition directly using only native diagonal u(1) charge bookkeeping. It does not import Standard Model hypercharge assignments.",
			"The u(1) twist is useful: an su(2) plane can preserve the diagonal u(1) only when its two generator modes carry equal weights. This kills temporal-spatial planes and leaves the three pure-spatial planes.",
			"The remaining pure-spatial planes have multiple gamma*Y eigenvalues in their doublet sectors, so the twist is not a physical chirality operator and does not complete the global H summand.",
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
