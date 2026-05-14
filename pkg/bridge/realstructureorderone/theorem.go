package realstructureorderone

import "github.com/bagherbal/asha-engine/pkg/theorem"

func RealStructureKOOrderOneCalculusAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-REAL-STRUCTURE-KO-ORDER-ONE-CALCULUS-AUDIT"
	const name = "Real Structure (J) integration / KO-Dimension and Order-One Calculus audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 234 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 233 finite D_F scaffold is inherited", Passed: a.Previous.Summary.DFAnsatzAvailable && !a.Previous.Summary.CanonicalDFDerived && !a.Previous.Summary.BGapEmbeddingDerived, Detail: a.Previous.TruthStatement},
			{Name: "occupation-complement J candidate is involutive and gamma-even", Passed: a.J.Constructed && a.J.J2Residual == 0 && a.J.J2Sign == 1 && a.J.CommutesWithOccupationGamma && a.J.JGammaSign == 1 && a.J.CandidateOnly && !a.J.PhysicalChargeConjugation, Detail: FormatJ(a.J)},
			{Name: "candidate KO signs are computable but not promoted", Passed: a.KO.J2Epsilon == 1 && a.KO.JGammaEpsilon == 1 && a.KO.JDSignIfImposed == 1 && !a.KO.JDCommutesForGenericM && a.KO.JDRequiresBlockSymmetry && a.KO.FreeParametersBefore == 64 && a.KO.FreeParametersAfterJ == 32 && !a.KO.PromotedKOTheorem, Detail: FormatKO(a.KO)},
			{Name: "J-reality sieve reduces the D_F block without selecting it", Passed: a.JReality.InitialParameters == 64 && a.JReality.ParametersAfterReality == 32 && !a.JReality.CanonicalBlockSelected && !a.JReality.ColorWeakSubblocksDerived && !a.JReality.PhysicalChiralityDerived, Detail: FormatJReality(a.JReality)},
			{Name: "finite algebra representation for non-vacuous order-one calculus remains missing", Passed: !a.Algebra.FaithfulTotalRepresentation && !a.Algebra.PhysicalSMAlgebraDerived && !a.Algebra.NontrivialOneFormsDerived, Detail: FormatAlgebra(a.Algebra)},
			{Name: "order-one condition is not verified and does not derive color/weak subblocks", Passed: !a.OrderOne.TestableWithCurrentData && !a.OrderOne.FullAlgebraRepresentation && !a.OrderOne.OrderOneVerified && !a.OrderOne.SplitsColorWeakSubblocks && !a.OrderOne.PromotableFiniteDirac, Detail: FormatOrderOne(a.OrderOne)},
			{Name: "B-gap is not forced into a right-handed Majorana slot", Passed: a.BGap.BGapAvailable && a.BGap.SterileVacuumCandidates >= 1 && !a.BGap.RightHandedNeutrinoSlotDerived && !a.BGap.MajoranaBilinearSpaceAvailable && !a.BGap.BGapCanonicalMajoranaEntry && !a.BGap.BGapPromotedToMajoranaMass && a.BGap.RequiresBroaderHilbertSpace, Detail: FormatBGap(a.BGap)},
			{Name: "firewall preserves finite-core status", Passed: !a.Firewall.ContinuumMassInserted && !a.Firewall.VEVInserted && !a.Firewall.MBInserted && !a.Firewall.MStarInserted && !a.Firewall.ObservedFermionMassInserted && !a.Firewall.BGapPromotedToMass && !a.Firewall.DFChosenByFit && !a.Firewall.KOClaimedAsTheorem && !a.Firewall.OrderOneClaimed && !a.Firewall.PMNSOrYukawaClaimed && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "Gate 234 summary records conditional J support and failed spectral-triple axiom route", Passed: a.Summary.CandidateJAvailable && a.Summary.CandidateKOSignsComputed && a.Summary.JRealityReducesParameters && !a.Summary.OrderOneDerived && !a.Summary.BGapMajoranaPlacement && !a.Summary.CanonicalDFDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}
		notes := []string{
			"The candidate occupation-complement J is useful finite bookkeeping, not yet physical charge conjugation.",
			"A non-vacuous order-one theorem still requires a faithful finite-algebra representation on the total Hilbert space and likely particle/antiparticle doubling.",
			"B_gap remains a dimensionless spectral datum; Gate 234 does not promote it to a Majorana mass or seesaw input.",
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
