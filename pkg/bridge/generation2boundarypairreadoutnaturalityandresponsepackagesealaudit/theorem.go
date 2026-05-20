package generation2boundarypairreadoutnaturalityandresponsepackagesealaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2BoundaryPairReadoutNaturalityAndResponsePackageSealAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 786 — Boundary Pair Readout Naturality and Response-Package Seal Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate785 response package obstruction", Passed: a.Gate785.Inherited && strings.Contains(a.Gate785.ConditionalPackage, "Theta_ext") && strings.Contains(a.Gate785.ConditionalPackage, "chi_ext") && !a.Gate785.PriorNative, Detail: a.Gate785.CurrentBottleneck},
			{Name: "record boundary pair data inventory", Passed: a.Inventory.Recorded && strings.Contains(a.Inventory.BoundaryCarrier, "B_boundary") && containsAll(a.Inventory.BoundaryAxes, []string{"b_lambda", "b_R"}) && containsAll(a.Inventory.BoundaryReadouts, []string{"s=", "xi_boundary"}) && closeRel(a.Inventory.K7EventWeight, 7.0/72.0, 1e-15) && strings.Contains(a.Inventory.FlavorScalar, "kappa_e_red") && containsAll(a.Inventory.ExteriorPackage, []string{"Theta_ext", "chi_ext", "ordered boundary orientation"}) && a.Inventory.RolesSeparated && !a.Inventory.AutoDefinesPackage, Detail: FormatInventory(a.Inventory)},
			{Name: "audit labelled basis versus native basis", Passed: a.Basis.Audited && a.Basis.LabelledBridgeBasis && !a.Basis.NativeInvariantBasis && containsAll(a.Basis.SourceTypes, []string{"scalar wall depth", "gauge/boundary stress"}) && a.Basis.BridgeBasisVerdict == StatusLabelledBoundaryAxesDefineBridgeBasis && a.Basis.NativeVerdict == StatusLabelledBridgeBasisNotNativeInvariantBasis, Detail: strings.Join(a.Basis.SourceTypes, "; ")},
			{Name: "audit degree-one axis/readout candidates", Passed: a.DegreeOne.Audited && a.DegreeOne.SplitAxisCandidate && strings.Contains(a.DegreeOne.SplitAxisSource, "b_R-b_lambda") && !a.DegreeOne.SplitAxisSourcesKappaE && a.DegreeOne.MidpointAxisCandidate && strings.Contains(a.DegreeOne.MidpointAxisSource, "b_lambda+b_R") && !a.DegreeOne.MidpointSourcesKappaE && a.DegreeOne.FlavorReadoutCoefficient && !a.DegreeOne.KappaESourcedByBoundary && !a.DegreeOne.NativeReadoutTheorem, Detail: FormatDegreeOne(a.DegreeOne)},
			{Name: "audit degree-two orientation and sign", Passed: a.DegreeTwo.Audited && a.DegreeTwo.Lambda2Exists && strings.Contains(a.DegreeTwo.VolumeForm, "omega_B") && a.DegreeTwo.RequiresOrderedBasis && closeRel(a.DegreeTwo.TwoP, 7.0/36.0, 1e-15) && strings.Contains(a.DegreeTwo.MagnitudeSource, "dim(B_boundary)") && strings.Contains(a.DegreeTwo.NegativeSign, "negative") && !a.DegreeTwo.NativeOrderedOrientation && !a.DegreeTwo.NativeNegativeStressPull, Detail: a.DegreeTwo.NegativeSign},
			{Name: "audit scalar readout chi_ext", Passed: a.ChiExt.Audited && a.ChiExt.DegreeZeroCanonical && closeRel(a.ChiExt.Chi0, 1, 1e-15) && a.ChiExt.DegreeOneRequiresFlavor && closeRel(a.ChiExt.Chi1, kappaERedSnapshot, 1e-15) && a.ChiExt.DegreeTwoRequiresK7AndSign && closeRel(a.ChiExt.Chi2, -2*pK7Snapshot, 1e-15) && !a.ChiExt.NativeFromBoundaryPairAlone, Detail: a.ChiExt.Verdict},
			{Name: "audit boundary symmetry naturality", Passed: a.Naturality.Audited && a.Naturality.AbstractGL2Freedom && !a.Naturality.CanonicalBetaUnderGL2 && !a.Naturality.CanonicalOmegaSignUnderGL2 && !a.Naturality.CanonicalChiUnderGL2 && a.Naturality.LabelledPairReducesSymmetry && a.Naturality.LabelledPackageBridgeSealed, Detail: a.Naturality.Verdict},
			{Name: "define minimal response-package seal", Passed: a.Seal.Defined && a.Seal.Name == "BoundaryExteriorResponsePackageSeal" && containsAll(a.Seal.Components, []string{"Theta_ext", "chi_ext", "negative stress-pull sign"}) && containsAll(a.Seal.Supplies, []string{"Theta_ext(M_n)", "chi_ext(1_B)=1", "chi_ext(beta_B)=kappa_e_red", "chi_ext(omega_B)=-2p", "Theta_ext(M_n>=4)=0"}) && a.Seal.Minimal && !a.Seal.Native && a.Seal.MatchesFWall, Detail: FormatSeal(a.Seal)},
			{Name: "audit impact on F_wall status", Passed: a.Impact.Audited && a.Impact.RepresentableBySeal && !a.Impact.NativeGeneratingFunction && strings.Contains(a.Impact.Status, "sealed exterior-response"), Detail: a.Impact.Status},
			{Name: "record status propagation", Passed: a.Propagation.Recorded && strings.Contains(a.Propagation.FWall3, "Level B+") && strings.Contains(a.Propagation.KappaLambda, "Level B") && strings.Contains(a.Propagation.CHistory, "Level B") && strings.Contains(a.Propagation.CHiggs, "not Level C"), Detail: FormatPropagation(a.Propagation)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.LabelledAxesNative && !a.Firewalls.SplitAxisKappaETheorem && !a.Firewalls.MidpointKappaETheorem && !a.Firewalls.OmegaSignTheorem && !a.Firewalls.TwoPMagnitudeFullTheorem && !a.Firewalls.SealNativeGeneratingFunc && !a.Firewalls.FWallNative && !a.Firewalls.KappaLambdaNative && !a.Firewalls.CHistoryIndependent && !a.Firewalls.TreeProxyPoleMass && !a.Firewalls.YukawaNative && a.Firewalls.Verdict == StatusFirewallPreservedGate786, Detail: a.Firewalls.Verdict},
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
