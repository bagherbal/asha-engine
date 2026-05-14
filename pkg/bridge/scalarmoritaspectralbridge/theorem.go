package scalarmoritaspectralbridge

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ScalarMoritaSpectralShapeBridgeBranchSelectorHeatKernelNormalizationAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-SCALAR-MORITA-SPECTRAL-SHAPE-BRIDGE-BRANCH-SELECTOR-HEAT-KERNEL-NORMALIZATION-AUDIT"
	const name = "Scalar-Morita Spectral Shape Bridge / Branch Selector and Heat-Kernel Normalization Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 276 scalar-Morita spectral bridge audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 275 two-branch constraint is inherited", Passed: a.Inheritance.ScalarMoritaSolved && a.Inheritance.TwoBranchXYConstrained && !a.Inheritance.UniqueXYLocked && !a.Inheritance.A2A4Derived && a.Inheritance.InheritedBranchCount == 2, Detail: FormatInheritance(a.Inheritance)},
			{Name: "scalar-Morita bridge is formalized as scale-free shape", Passed: a.Bridge.LambdaNumerator == 1197 && a.Bridge.LambdaDenominator == 4624 && a.Bridge.KappaC == 1 && a.Bridge.KappaQ == 3 && a.Bridge.CrossTowerBridge && a.Bridge.ScaleFreeShapeOnly && !a.Bridge.EquivalentToA2A4, Detail: FormatBridge(a.Bridge)},
			{Name: "both branches are carried and reproduce lambda", Passed: BranchResidualsOK(a.Branches), Detail: FormatBranches(a.Branches)},
			{Name: "branch selector audit leaves two-fold ambiguity", Passed: a.Selector.UpperBranchAllowed && a.Selector.LowerBranchAllowed && !a.Selector.UniqueBranch && !a.Selector.FiniteCoreSelector && a.Selector.RequiresFutureInput, Detail: FormatSelector(a.Selector)},
			{Name: "formal heat-kernel obligations are defined but not derived", Passed: a.HeatKernel.FormalExpansion != "" && !a.HeatKernel.CutoffMomentsSpecified && !a.HeatKernel.SubtractionSchemeDerived && !a.HeatKernel.GaugeKineticProjection && !a.HeatKernel.ScalarFluctuationMap && !a.HeatKernel.CanMapRawTracesToA2A4, Detail: FormatHeatKernel(a.HeatKernel)},
			{Name: "raw moments are not promoted to a2/a4 or Higgs ratio", Passed: !a.HiggsRatio.UsesSelectedBranch && !a.HiggsRatio.UsesAbsoluteDFScale && !a.HiggsRatio.UsesHeatKernelMap && !a.HiggsRatio.InvariantA2A4Computed && !a.HiggsRatio.HiggsMassRatioComputed, Detail: FormatHiggs(a.HiggsRatio)},
			{Name: "firewalls preserve empirical and spectral boundaries", Passed: a.Firewall.NoObservedMassInserted && a.Firewall.NoVEVInserted && a.Firewall.NoCKMPMNSInserted && a.Firewall.NoEmpiricalYukawaAmplitudeInserted && a.Firewall.RawTraceShapeNotPromoted && a.Firewall.CandidateBranchesNotPredictions && a.Firewall.EmpiricalYukawaSealPreserved && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "future map lists all remaining projection obligations", Passed: a.Future.NeedBranchSelector && a.Future.NeedAbsoluteScale && a.Future.NeedPhysicalJ && a.Future.NeedHypercharge && a.Future.NeedHeatKernelProjection && a.Future.NeedFieldNormalization && len(a.Future.Criteria) >= 7, Detail: FormatFuture(a.Future)},
			{Name: "summary records bridge support but no Higgs prediction", Passed: a.Summary.Gate275Inherited && a.Summary.BridgeFormalized && a.Summary.TwoBranchesCarried && !a.Summary.UniqueBranchSelected && a.Summary.HeatKernelFormalized && !a.Summary.HeatKernelDerived && !a.Summary.A2A4Derived && !a.Summary.HiggsRatioClaimed && a.Summary.FirewallPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{
			a.TruthStatement,
			"Gate 276 confirms the scalar-Morita bridge as a finite scale-free shape constraint, but it does not select r_+/r_- or derive Seeley-de Witt coefficients.",
			"A physical Higgs ratio still requires branch selection, heat-kernel projection, field normalization, and physical J/hypercharge completion.",
		}}
	}}
}
