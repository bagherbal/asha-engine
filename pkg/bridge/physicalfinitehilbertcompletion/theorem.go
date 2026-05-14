package physicalfinitehilbertcompletion

import "github.com/bagherbal/asha-engine/pkg/theorem"

func PhysicalFiniteHilbertSpaceChiralHyperchargeOppositeActionCompletionAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-PHYSICAL-FINITE-HILBERT-SPACE-CHIRAL-HYPERCHARGE-OPPOSITE-ACTION-COMPLETION-AUDIT"
	const name = "Physical Finite Hilbert Space / Chiral Hypercharge Opposite-Action Completion Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 275 physical finite Hilbert completion audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 274 local quaternionic boundary is inherited", Passed: a.Inheritance.LocalHExtracted && a.Inheritance.CandidateAlgebraBuilt && !a.Inheritance.PhysicalJDerived && !a.Inheritance.XYRatioLocked && !a.Inheritance.A2A4Derived && a.Inheritance.FirewallPreserved, Detail: FormatInheritance(a.Inheritance)},
			{Name: "Gate 169 scalar shape is retrieved as exact finite target", Passed: a.Bridge.Shape.SourceGate == 169 && a.Bridge.Shape.ExactNumerator == 1197 && a.Bridge.Shape.ExactDenominator == 4624 && a.Bridge.Shape.DerivedFiniteCoreDatum && !a.Bridge.Shape.EmpiricalInput, Detail: FormatScalarShape(a.Bridge.Shape)},
			{Name: "Gate 273 Morita multiplicity ledger is retrieved", Passed: a.Bridge.Multiplicity.KappaC == 1 && a.Bridge.Multiplicity.KappaQ == 3 && a.Bridge.Multiplicity.MultiplicityDerived && !a.Bridge.Multiplicity.AmplitudeDerived, Detail: FormatMultiplicity(a.Bridge.Multiplicity)},
			{Name: "scalar-Morita bridge quadratic is solved exactly", Passed: a.Bridge.Quadratic.A == 3099 && a.Bridge.Quadratic.B == -7182 && a.Bridge.Quadratic.C == 3427 && a.Bridge.Quadratic.Discriminant == 9100032 && a.Bridge.Quadratic.HasTwoPositiveRoots && len(a.Bridge.Branches) == 2, Detail: FormatQuadratic(a.Bridge.Quadratic)},
			{Name: "both amplitude branches reproduce lambda_contact", Passed: branchesReproduceShape(a.Bridge), Detail: FormatBridge(a.Bridge)},
			{Name: "branch and absolute scale remain unselected", Passed: a.Bridge.RootsConstrainR && !a.Bridge.UniqueBranchSelected && !a.Bridge.AbsoluteScaleSelected && !a.Bridge.A2A4Derived && !a.Bridge.HiggsRatioDerived, Detail: FormatBridge(a.Bridge)},
			{Name: "physical J remains a candidate, not a completed charge conjugation theorem", Passed: a.J.OccupationComplementSeen && a.J.CandidateJ2 == 1 && !a.J.AntiLinearImplemented && !a.J.ParticleAntiparticleTyped && !a.J.PhysicalHFCompleted, Detail: FormatJ(a.J)},
			{Name: "hypercharge ledgers exist but full chiral assignment is not derived", Passed: a.Hypercharge.BMinusLLedgerAvailable && a.Hypercharge.T3LedgerAvailable && a.Hypercharge.CandidateHyperchargeKnown && !a.Hypercharge.FullCPlusHPlusM3Representation && !a.Hypercharge.ChiralAssignmentDerived && !a.Hypercharge.EmpiricalAssignmentsInserted, Detail: FormatHypercharge(a.Hypercharge)},
			{Name: "full opposite action and order-one recheck remain blocked", Passed: a.OppositeAction.MoritaOppositeActionInherited && a.OppositeAction.LocalHIncluded && !a.OppositeAction.PhysicalJDerived && !a.OppositeAction.FullOppositeActionDerived && !a.OppositeAction.OrderOneReevaluatedOnFullAF && !a.OppositeAction.XYRatioBranchSelectedByJ, Detail: FormatOpposite(a.OppositeAction)},
			{Name: "firewalls preserve candidate moments versus Higgs prediction", Passed: a.Firewall.NoObservedMassInserted && a.Firewall.NoVEVInserted && a.Firewall.NoCKMPMNSInserted && a.Firewall.NoEmpiricalYukawaAmplitudeInserted && a.Firewall.ScalarShapeKeptFinite && a.Firewall.ScalarMoritaBridgeMarkedConditional && a.Firewall.CandidateMomentsNotHiggsPrediction && a.Firewall.EmpiricalYukawaSealPreserved && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "future criteria identify branch, J, hypercharge, and heat-kernel obligations", Passed: a.Future.NeedScalarMoritaMapTheorem && a.Future.NeedBranchSelector && a.Future.NeedPhysicalJ && a.Future.NeedHyperchargeCompletion && a.Future.NeedHeatKernelProjection && len(a.Future.Criteria) >= 7, Detail: FormatFuture(a.Future)},
			{Name: "summary records two-branch constraint but no a2/a4 theorem", Passed: a.Summary.Gate274Inherited && a.Summary.ScalarShapeRetrieved && a.Summary.MoritaMultiplicityKnown && a.Summary.ScalarMoritaSolved && a.Summary.TwoBranchXYConstrained && !a.Summary.UniqueXYLocked && !a.Summary.PhysicalJDerived && !a.Summary.HyperchargeDerived && !a.Summary.A2A4Derived && !a.Summary.HiggsRatioClaimed && a.Summary.FirewallPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{
			a.TruthStatement,
			"Gate 275 gives a real two-branch finite amplitude-shape constraint from Gate 169 + Gate 273, but it does not select a branch or claim a Higgs mass.",
			"Physical J, chiral hypercharge assignment, full opposite action, and Seeley-de Witt projection remain required before a2/a4 can be promoted.",
		}}
	}}
}

func branchesReproduceShape(a ScalarMoritaBridgeAudit) bool {
	if len(a.Branches) != 2 {
		return false
	}
	for _, b := range a.Branches {
		if b.R <= 0 || b.AbsYOverX <= 0 || b.ShapeResidualAbs > 1e-12 {
			return false
		}
	}
	return true
}
