package generation2flavorboundaryreadoutandstresspullorientationsealfactorizationaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2FlavorBoundaryReadoutAndStressPullOrientationSealFactorizationAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 787 — Flavor-Boundary Readout and Stress-Pull Orientation Seal Factorization Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate786 boundary-pair response-package seal", Passed: a.Gate786.Inherited && strings.Contains(a.Gate786.PackageSeal, "BoundaryExteriorResponsePackageSeal") && !a.Gate786.PriorNative, Detail: a.Gate786.CurrentProblem},
			{Name: "factor response package into subseals", Passed: a.Factorization.Recorded && a.Factorization.CompositeSeal == "BoundaryExteriorResponsePackageSeal" && containsAll(a.Factorization.Subseals, []string{"DegreeRuleSeal", "FlavorBoundaryReadoutSeal", "BoundaryStressPullOrientationSeal"}) && a.Factorization.DegreeZeroCanonical && a.Factorization.ThreeNontrivialSubobjects, Detail: FormatFactorization(a.Factorization)},
			{Name: "audit degree rule", Passed: a.DegreeRule.Audited && strings.Contains(a.DegreeRule.Rule, "Lambda^(n-1)B_boundary") && a.DegreeRule.ExplainsCubicStop && !a.DegreeRule.Native && a.DegreeRule.ThetaExtSealed && a.DegreeRule.ProjectorPowersContinue, Detail: a.DegreeRule.Rule},
			{Name: "audit degree-one flavor-boundary readout", Passed: a.FlavorReadout.Audited && closeRel(a.FlavorReadout.KappaERed, kappaERedSnapshot, 1e-15) && closeRel(a.FlavorReadout.KappaOrient+a.FlavorReadout.KappaBoundary, kappaERedSnapshot, 1e-15) && containsAll(a.FlavorReadout.TermTypes, []string{"PMNS", "CKM", "hypercharge", "boundary-stress"}) && a.FlavorReadout.MixedReadout && !a.FlavorReadout.NativeFromBoundary && !a.FlavorReadout.NativeFlavorTheorem, Detail: FormatFlavorReadout(a.FlavorReadout)},
			{Name: "audit boundary-only degree-one candidates", Passed: a.Axis.Audited && a.Axis.SplitAxisCandidate && strings.Contains(a.Axis.SplitAxis, "b_R - b_lambda") && !a.Axis.SplitAxisSourcesKappaE && a.Axis.MidpointAxisCandidate && strings.Contains(a.Axis.MidpointAxis, "b_lambda + b_R") && !a.Axis.MidpointAxisSourcesKappaE && !a.Axis.BoundaryAxesReplaceFlavor, Detail: a.Axis.Verdict},
			{Name: "audit boundary part of kappa_e_red", Passed: a.KappaBoundary.Audited && closeRel(a.KappaBoundary.KappaBoundary, (-5.0/3.0+xiBoundarySnapshot*pK7Snapshot)*sSplitSnapshot*sSplitSnapshot, 1e-18) && closeRel(a.KappaBoundary.HyperchargeFactor, 5.0/3.0, 1e-15) && closeRel(a.KappaBoundary.BoundaryStressTerm, xiBoundarySnapshot*pK7Snapshot, 1e-15) && a.KappaBoundary.StrongSourceType && !a.KappaBoundary.FullKappaESource && strings.Contains(a.KappaBoundary.MainNonNativePart, "sin^2(theta13)/4 - J_CKM") && !a.KappaBoundary.NativeCoupling, Detail: FormatKappaBoundary(a.KappaBoundary)},
			{Name: "audit degree-two stress-pull orientation", Passed: a.StressPull.Audited && closeRel(a.StressPull.Magnitude, 2*pK7Snapshot, 1e-15) && strings.Contains(a.StressPull.MagnitudeSource, "dim(B_boundary)") && a.StressPull.NegativeSign && containsAll(a.StressPull.CandidateSignSources, []string{"ordered boundary orientation", "restorative stress-pull"}) && !a.StressPull.NativeNegativeSign && !a.StressPull.MatchingSignNative, Detail: FormatStressPull(a.StressPull)},
			{Name: "record minimal factor classification", Passed: a.Minimal.Recorded && containsAll(a.Minimal.Seals, []string{"DegreeRuleSeal", "FlavorBoundaryReadoutSeal", "BoundaryStressPullOrientationSeal"}) && strings.Contains(a.Minimal.Mapping["FlavorBoundaryReadoutSeal"], "kappa_e_red") && strings.Contains(a.Minimal.Mapping["BoundaryStressPullOrientationSeal"], "negative sign") && !a.Minimal.Native, Detail: a.Minimal.Verdict},
			{Name: "audit formula-level runtime target absence", Passed: a.Runtime.Audited && containsAll(a.Runtime.ForbiddenDirectVariables, []string{"lambda_runtime", "m_H_pole", "C_Higgs", "G_F", "v"}) && !a.Runtime.ContainsForbidden && a.Runtime.FormulaLevelIndependent && !a.Runtime.TheoremLevelIndependent, Detail: a.Runtime.Verdict},
			{Name: "record status propagation", Passed: a.Propagation.Recorded && strings.Contains(a.Propagation.FWall3, "Level B+") && strings.Contains(a.Propagation.KappaLambda, "Level B") && strings.Contains(a.Propagation.CHistory, "Level B") && strings.Contains(a.Propagation.CHiggs, "not Level C"), Detail: FormatPropagation(a.Propagation)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.KappaENativeFlavorTheorem && !a.Firewalls.KappaBoundaryFullFlavorTheorem && !a.Firewalls.KappaOrientPMNSCKMTheorem && !a.Firewalls.SplitMidpointAxesKappaETheorem && !a.Firewalls.TwoPMagnitudeFullSignTheorem && !a.Firewalls.NegativeCubicSignNative && !a.Firewalls.ResponsePackageNativeGenerating && !a.Firewalls.FWallNative && !a.Firewalls.KappaLambdaNative && !a.Firewalls.CHistoryIndependent && !a.Firewalls.TreeProxyPoleMass && a.Firewalls.Verdict == StatusFirewallPreservedGate787, Detail: a.Firewalls.Verdict},
		}
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
		notes := append([]string{a.Truth, a.FinalStatement}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
