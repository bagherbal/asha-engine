package su2spinorlift

import "github.com/bagherbal/asha-engine/pkg/theorem"

func SU2SpinorLiftQuaternionicClosureAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-SU2-SPINOR-LIFT-QUATERNIONIC-CLOSURE-AUDIT"
	const name = "Explicit su(2) spinor lift / quaternionic closure audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 237 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 236 native algebra obstruction is inherited", Passed: a.Previous.Summary.CPlusM3Preflight && !a.Previous.Summary.QuaternionicHDerived && !a.Previous.Summary.ExactSMAlgebraDerived, Detail: a.Previous.TruthStatement},
			{Name: "candidate exterior su(2) lifts exist on every two-mode plane", Passed: a.Lift.CandidateWedgeLiftsComputed && a.Lift.CandidatePlaneCount == 6 && a.Lift.ClosureResidual == 0 && !a.Lift.ExplicitContactMatricesOnSC && !a.Lift.NativeIdentifiesContactSU2 && !a.Lift.NativeWeakPlaneSelected, Detail: FormatLift(a.Lift) + " :: " + FormatPlanes(a.Planes)},
			{Name: "candidate lifts contain the correct doublet/singlet dimensional pattern", Passed: a.Doublets.CandidatePlanes == 6 && a.Doublets.DoubletCopiesPerPlane == 4 && a.Doublets.DoubletStateDimCPerPlane == 8 && a.Doublets.SingletStateDimCPerPlane == 8 && a.Doublets.DimensionalMatchToQLPlusLL && !a.Doublets.HyperchargeAssignmentDerived && !a.Doublets.PhysicalLeftDoubletProjection, Detail: FormatDoublets(a.Doublets)},
			{Name: "local pseudo-real doublet structure supports H only after choosing a plane", Passed: a.Quaternionic.FundamentalDoubletPseudoReal && a.Quaternionic.LocalQuaternionicStructureOnDoublet && a.Quaternionic.PlaneSelectionRequired && !a.Quaternionic.GlobalHSummandDerived && !a.Quaternionic.OppositeActionDerived && !a.Quaternionic.OrderOneReady, Detail: FormatQuaternionic(a.Quaternionic)},
			{Name: "full native finite algebra remains incomplete", Passed: a.Algebra.PreviousCPlusM3Preflight && a.Algebra.U1ComplexPreflight && a.Algebra.LocalHPreflight && !a.Algebra.NativeHGlobalSummand && !a.Algebra.ExactCPlusHPlusM3Derived && !a.Algebra.FaithfulRepresentationOnSC && !a.Algebra.FullOrderOneCalculusReady, Detail: FormatAlgebra(a.Algebra)},
			{Name: "firewall blocks forced weak algebra", Passed: !a.Firewall.PauliMatricesImportedAsAnswer && !a.Firewall.ConnesAlgebraImported && !a.Firewall.WeakPlaneForced && !a.Firewall.HyperchargeForced && !a.Firewall.SMGaugeGroupInserted && !a.Firewall.BGapPromotedToMass && !a.Firewall.ClaimedExactH && !a.Firewall.ClaimedOrderOne && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "summary records support but not derivation", Passed: a.Summary.CandidateSU2Lifts && a.Summary.DoubletDimensionalSupport && a.Summary.PseudoRealLocalHSupport && !a.Summary.NativeContactLiftDerived && !a.Summary.CanonicalWeakPlane && !a.Summary.GlobalHDerived && !a.Summary.ExactSMAlgebraDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}
		notes := []string{
			"Gate 237 shows that every selected two-mode plane U⊂W gives a mathematically valid exterior su(2) lift on S_C=Λ*W and decomposes into four doublets plus eight singlets.",
			"The doublet sector has the correct complex dimension for one SM generation's Q_L⊕L_L and is pseudo-real, so a local quaternionic module is available after a plane is selected.",
			"The decisive selector remains missing: the finite core has not identified the contact-preserving su(2) with one canonical plane, nor derived hypercharge/color attachment, opposite action, or order-one calculus. The H summand is therefore supported but not derived.",
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
