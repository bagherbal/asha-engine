package generation2lawhistorywallhierarchyandk7responsefirewallorderingaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2LawHistoryWallHierarchyAndK7ResponseFirewallOrderingAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 749 — Law-History Wall Hierarchy and K7 Response Firewall Ordering Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate749 wall hierarchy audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate748 kappa_e boundary-stress moment audit", Passed: a.Gate748.Inherited && a.Gate748.ResidualCompression && a.Gate748.BoundaryStressReappears && a.Gate748.FlavorFirewallKept && strings.Contains(a.Gate748.Verdict, StatusGate748BoundaryStressMomentInherited), Detail: FormatGate748(a.Gate748)},
			{Name: "define full law-history wall hierarchy", Passed: a.Hierarchy.Count >= 16 && containsWall(a.Hierarchy, "K7 contact wall") && containsWall(a.Hierarchy, "Boundary stress wall") && containsWall(a.Hierarchy, "Tree/pole wall") && strings.Contains(a.Hierarchy.Verdict, StatusWallHierarchyDefined), Detail: FormatHierarchy(a.Hierarchy)},
			{Name: "separate K7 support, event-weight, and forbidden promotion roles", Passed: a.K7Roles.BoundaryVectorMapBlocked && a.K7Roles.FlavorPromotionBlocked && a.K7Roles.HiggsPromotionBlocked && strings.Contains(a.K7Roles.Verdict, StatusK7ActsAsSupportAndEventWeightOnly), Detail: FormatK7Roles(a.K7Roles)},
			{Name: "construct ordered firewall ladder", Passed: a.Firewall.Count == 12 && a.Firewall.Steps[0].Order == 1 && a.Firewall.Steps[11].Order == 12 && strings.Contains(a.Firewall.Verdict, StatusFirewallOrderConstructed), Detail: FormatFirewall(a.Firewall)},
			{Name: "record Gate748 correction as law-history wall resonance", Passed: a.Resonance.UsesOrientation && a.Resonance.UsesFiveThirds && a.Resonance.UsesXiBoundary && a.Resonance.UsesK7Moment && !a.Resonance.IsFlavorTheorem && strings.Contains(a.Resonance.Verdict, StatusGate748CorrectionLawHistoryWallResonance), Detail: FormatResonance(a.Resonance)},
			{Name: "record boundary raw-moment K7 response coordinate", Passed: a.Moment.K7EventWeight > 0.09 && a.Moment.K7EventWeight < 0.1 && a.Moment.XiBoundary > 0.05 && a.Moment.M2Wall > 1e-7 && strings.Contains(a.Moment.Verdict, StatusBoundaryMomentCoordinateRecorded), Detail: FormatMoment(a.Moment)},
			{Name: "record reduction priority without chasing tiny residual", Passed: a.Reduction.DoNotChaseResidual && a.Reduction.StabilizedBeforeNext && len(a.Reduction.RecommendedTargets) == 4 && strings.Contains(a.Reduction.Verdict, StatusWallOrderingStabilizesNextReductions), Detail: FormatReduction(a.Reduction)},
			{Name: "enforce physical and theorem firewalls", Passed: a.Physical.NoBoundaryVectorMap && a.Physical.NoFlavorTheorem && a.Physical.NoHistoryLoopTheorem && a.Physical.NoScalarRuntimeTheorem && a.Physical.NoHiggsPoleMassTheorem && a.Physical.NoYukawaTheorem && strings.Contains(a.Physical.Verdict, StatusGate749Boundary), Detail: FormatPhysical(a.Physical)},
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
		notes := append([]string{a.Truth}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func containsWall(h WallHierarchy, name string) bool {
	for _, w := range h.Walls {
		if w.Name == name {
			return true
		}
	}
	return false
}
