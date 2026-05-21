package generation2symbolicfinitediracmatrixsupportfirstorderfirewallaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_SYMBOLIC_FINITE_DIRAC_MATRIX_SUPPORT_FIRST_ORDER_FIREWALL_AUDIT"
	theoremName = "Gate 848 — Symbolic Finite-Dirac Matrix Support and First-Order Firewall Audit"
)

func Generation2SymbolicFiniteDiracMatrixSupportFirstOrderFirewallAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		checks := []theorem.Check{}
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks = append(checks,
			theorem.Check{Name: "inherit Gate 847 minimal right and weak-socket support", Passed: a.Impact.Gate847Inherited && a.Y.DomainRank == HRMinRank && a.Y.TargetRank == HLRank && containsAll(a.Y.Supports, []string{SupportThreeSymbolicEdgeFamilies}), Detail: FormatY(a.Y)},
			theorem.Check{Name: "construct symbolic Y support matrix with three active families", Passed: a.Y.SupportOnly && a.Y.ActiveFamilies == 3 && a.Y.ExpectedFamilies == 3 && len(a.Y.Edges) == 3 && allActiveEdgesSymbolicNoMagnitude(a.Y.Edges) && a.Y.PreservesLeptoColor && containsAll(a.Y.Supports, []string{SupportThreeSymbolicEdgeFamilies, SupportLeptoColorPreservingMatrix}), Detail: FormatY(a.Y)},
			theorem.Check{Name: "set puncture edge coefficient to zero", Passed: a.Y.MissingEdge.Puncture && !a.Y.MissingEdge.Present && a.Y.MissingEdge.Coefficient == "y_+1=0" && a.Y.PunctureCoefficient == "y_+1" && a.Y.PunctureCoefficientZero && containsAll(a.Y.Supports, []string{SupportPunctureZeroCoefficient, SupportNeutralSingletonNullEdge}), Detail: FormatEdge(a.Y.MissingEdge)},
			theorem.Check{Name: "build chiral symbolic D_F support matrix", Passed: a.Dirac.ExplicitSupportMatrix && a.Dirac.Expression == "D_F^sym = [[0, Y_supp^dagger], [Y_supp, 0]]" && a.Dirac.LeftRank == HLRank && a.Dirac.RightRank == HRMinRank && a.Dirac.TotalRank == ChiralTotalDim && a.Dirac.BlockRows == ChiralTotalDim && a.Dirac.BlockCols == ChiralTotalDim && containsAll(a.Dirac.Supports, []string{SupportSymbolicDFSupportMatrix}), Detail: FormatDirac(a.Dirac)},
			theorem.Check{Name: "certify self-adjointness by adjoint block inclusion", Passed: a.Dirac.UsesAdjointBlock && a.Dirac.SelfAdjointByConstruction && containsAll(a.Dirac.Supports, []string{SupportSelfAdjointByBlockForm}), Detail: FormatDirac(a.Dirac)},
			theorem.Check{Name: "certify chirality oddness by left/right block form", Passed: a.Dirac.ChiralOddByConstruction && a.Dirac.GammaConvention != "" && containsAll(a.Dirac.Supports, []string{SupportChiralityOddByBlockForm}), Detail: FormatDirac(a.Dirac)},
			theorem.Check{Name: "audit first-order, bimodule, and J/opposite firewalls", Passed: !a.Dirac.NativeDFMatrix && !a.Dirac.NumericalDFMatrix && !a.Dirac.FirstOrderCertified && !a.Dirac.BimoduleCommutantProof && !a.Dirac.JOppositeCompatibilityProof && containsAll(a.Dirac.Failures, []string{FailureDFSupportMatrixSealNotNative, FailureNoNumericalDFMatrix, FailureNoFullRhoFJFGammaFPackage, FailureNoFirstOrderProof, FailureNoBimoduleCommutantProof, FailureNoJOppositeCompatibility}), Detail: FormatDirac(a.Dirac)},
			theorem.Check{Name: "preserve symbolic-variable, magnitude, alpha, and particle firewalls", Passed: !a.Y.HasNumericalValues && allActiveEdgesSymbolicNoMagnitude(a.Y.Edges) && containsAll(a.Y.Failures, []string{FailureYSymbolsNotYukawaValues, FailureEdgeSupportNotTraceMagnitude, FailureNoNumericalYukawaValues, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureNoPhysicalParticleAssign, FailureNoRightNeutrinoTheorem, FailureNoThreeGenerationTheorem}) && a.Impact.AlphaStillSealed && a.Impact.MagnitudesStillMissing, Detail: FormatY(a.Y) + " | " + FormatImpact(a.Impact)},
			theorem.Check{Name: "preserve official ledger freeze and no R3/R4 promotion", Passed: a.Ledger.OfficialFrozen && a.Ledger.R2PlusPlusPlusPlusPlus && !a.Ledger.R3 && !a.Ledger.R4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			theorem.Check{Name: "preserve Gate 848 firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.DFSupportMatrixSealNotNative && a.Firewalls.NoNumericalDFMatrix && a.Firewalls.NoFullRhoFJFGammaFPackage && a.Firewalls.NoFirstOrderProof && a.Firewalls.NoBimoduleCommutantProof && a.Firewalls.NoJOppositeCompatibilityProof && a.Firewalls.YSymbolsNotYukawaValues && a.Firewalls.EdgeSupportNotTraceMagnitude && a.Firewalls.NoNumericalYukawaValues && a.Firewalls.AlphaStillSealed && a.Firewalls.NoTraceMagnitudeReadout && a.Firewalls.PunctureNullEdgeOnlySeal && a.Firewalls.NoNativeNullEdgeTheorem && a.Firewalls.NoPhysicalParticleAssignment && a.Firewalls.NoRightNeutrinoTheorem && a.Firewalls.NoThreeGenerationTheorem && a.Firewalls.NoNEffUpdate && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.Verdict == StatusFirewallGate848, Detail: FormatFirewalls(a.Firewalls)},
		)
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatY(a.Y), FormatDirac(a.Dirac), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func allActiveEdgesSymbolicNoMagnitude(edges []SymbolicEdge) bool {
	for _, e := range edges {
		if !e.Present || e.Puncture || e.HasMagnitude || e.Coefficient == "" || e.DomainRank <= 0 || e.TargetRank <= 0 || e.ValueSource != "symbolic support variable only" {
			return false
		}
		if e.LeptoColor != "P_3 -> P_3" && e.LeptoColor != "P_1 -> P_1" {
			return false
		}
	}
	return true
}
