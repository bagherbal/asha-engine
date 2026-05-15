package generation2specialbranchselector

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2TextureZeroSpecialBranchSelectorNecessaryBoundaryAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 texture-zero special-branch selector necessary boundary audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate451 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate 450 texture-zero limit", Passed: a.Inheritance.Executed && a.Inheritance.Gate444KGenForced && a.Inheritance.Gate444Generation2Zero && a.Inheritance.Gate445TriangleForced && a.Inheritance.Gate446PhaseQuarantined && a.Inheritance.Gate447CoefficientsSealed && a.Inheritance.Gate450TextureZeroSumRule && a.Inheritance.Gate450RatioSealed && a.Inheritance.NoEmpiricalInputsImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "audits edge suppression across native laws", Passed: a.EdgeAudit.Executed && len(a.EdgeAudit.Edges) == 3 && len(a.EdgeAudit.Laws) >= 5 && a.EdgeAudit.AllNativeLawsAllow13 && !a.EdgeAudit.AnyNativeLawSuppresses13 && a.EdgeAudit.FullTrianglePreserved && !a.EdgeAudit.NearestNeighborNativelyForced, Detail: FormatEdgeAudit(a.EdgeAudit)},
			{Name: "nearest-neighbor 1-3 suppression is not the native mass-lift branch", Passed: a.EdgeAudit.NearestNeighborFailsMassLift && a.EdgeAudit.FullTriangleDeterminantCoeff == 2 && a.EdgeAudit.NearestNeighborDeterminantCoeff == 0, Detail: a.EdgeAudit.FullTriangleDeterminant + "; " + a.EdgeAudit.NearestNeighborDeterminant},
			{Name: "audits phase ray without selecting one", Passed: a.PhaseAudit.Executed && a.PhaseAudit.NativeConstraintsPhaseBlind && a.PhaseAudit.SurvivingNonzeroLiftRays >= 3 && a.PhaseAudit.ContainsCZeroSurvivor && a.PhaseAudit.ContainsNonzeroCSurvivor && !a.PhaseAudit.UniqueRayForced && !a.PhaseAudit.FixesCZero && !a.PhaseAudit.FixesPiOverTwo, Detail: FormatPhaseAudit(a.PhaseAudit)},
			{Name: "GST/Fritzsch branch is not natively forced", Passed: a.GST.Executed && !a.GST.EdgeSelectorFound && !a.GST.PhaseSelectorFound && !a.GST.GSTLikeBranchNativelyForced && !a.GST.MassAngleRatiosReevaluated && a.GST.GSTFritzschEmpiricalAssumption, Detail: FormatGST(a.GST)},
			{Name: "13-moduli firewall preserved", Passed: a.Firewall.Executed && a.Firewall.NoObservedMuonMassImported && a.Firewall.NoObservedCharmMassImported && a.Firewall.NoObservedYukawaImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoCurveFit && a.Firewall.KGenStillForced && a.Firewall.Generation2ZeroStillForced && a.Firewall.XTriangleStillForced && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.GSTFritzschRelationsQuarantined && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate quarantines optional phenomenology", Passed: a.Next.Gate == 452, Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: checks, Notes: []string{StatusFailedNativeGeometryPreservesFullTriangle, StatusFailedNoNative13EdgeSuppression, StatusFailedNoNativePhaseRaySelector, StatusFailedGSTRequiresExtraTextureAssumption, StatusEmpiricalFirewallPreserved, a.Truth}}
	}}
}
