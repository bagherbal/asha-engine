package generation2chiralitymassbridgefirewallandboundaryrestpressurerelevanceaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-812-CHIRALITY-MASS-BRIDGE-FIREWALL-BOUNDARY-RESTPRESSURE-RELEVANCE"
	theoremName = "Gate 812 — Chirality Mass-Bridge Firewall and Boundary RestPressure Relevance Audit"
)

func Generation2ChiralityMassBridgeFirewallAndBoundaryRestPressureRelevanceAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 811 boundary second-moment rest-pressure target", Passed: a.Inheritance.Gate811Inherited && a.Inheritance.BoundarySecondMomentSelected && math.Abs(a.Inheritance.M2-1.624013231638281e-7) < 1e-20 && math.Abs(a.Inheritance.C2Obs-5.8299915722461693) < 5e-13 && containsAll(a.Inheritance.Verdicts, []string{StatusGate811Inherited, StatusBoundarySecondMomentSelected}), Detail: "Gate 811 residual, M2, and c2 target inherited"},
			{Name: "audit real Cl(1,7) pseudoscalar chirality projectors", Passed: a.Pseudoscalar.Inherited && a.Pseudoscalar.Audited && a.Pseudoscalar.OmegaSquared == -1 && !a.Pseudoscalar.NaiveProjectorsIdempotent && a.Pseudoscalar.PPlusSquaredScalar == 0 && math.Abs(a.Pseudoscalar.PPlusSquaredOmega-0.5) < 1e-15 && a.Pseudoscalar.ComplexGammaSquared == 1 && containsAll(a.Pseudoscalar.Supports, []string{StatusChiralityAirlockRequired}) && containsAll(a.Pseudoscalar.Failures, []string{StatusNaiveRealProjectorsInvalid, StatusNoNativeRealChirality, StatusComplexAirlockNotNative}), Detail: FormatPseudoscalar(a.Pseudoscalar)},
			{Name: "define weak chirality restriction requirements", Passed: a.WeakRestriction.Audited && len(a.WeakRestriction.Requirements) >= 6 && containsAll(a.WeakRestriction.Supports, []string{StatusChiralRestrictionSearch}) && containsAll(a.WeakRestriction.Failures, []string{StatusNoNativeSU2L, StatusAirlockNoGaugeAssignments, StatusWeakNoTraceMagnitudes}), Detail: FormatRequirement(a.WeakRestriction)},
			{Name: "audit Higgs scalar / finite one-form mass bridge", Passed: a.HiggsMassBridge.Audited && len(a.HiggsMassBridge.Facts) >= 3 && containsAll(a.HiggsMassBridge.Verdicts, []string{StatusHiggsMassBridgeAudited, StatusFiniteTripleEdgeCompat}) && containsAll(a.HiggsMassBridge.Supports, []string{StatusMassBridgeTyped, StatusMassBridgeEdgeOnly}) && containsAll(a.HiggsMassBridge.Failures, []string{StatusHiggsScalarNoYf, StatusMassBridgeNoEigenvalues, StatusEdgeNoDeltaN, StatusMassBridgeNoTopRest}), Detail: FormatMassBridge(a.HiggsMassBridge)},
			{Name: "audit grade-zero commuting claim", Passed: a.GradeZero.Audited && containsAll(a.GradeZero.Supports, []string{StatusScalarBilinearCompat}) && containsAll(a.GradeZero.Failures, []string{StatusCommutingNoEdge, StatusGradeZeroNotUniqueYukawa, StatusScalarNotHierarchyOperator}), Detail: FormatRequirement(a.GradeZero)},
			{Name: "audit relation to TraceMagnitudeOperatorSeal", Passed: a.TraceMagnitude.Audited && containsAll(a.TraceMagnitude.Supports, []string{StatusBridgeEdgeCompatOnly}) && containsAll(a.TraceMagnitude.Failures, []string{StatusNoHermitianTraceOps, StatusNoPositiveTraceAtoms, StatusNoTopColorBlock, StatusNoRestPressureOperator}), Detail: FormatRequirement(a.TraceMagnitude)},
			{Name: "audit relation to Gate811 boundary-FN package", Passed: a.BoundaryFN.Audited && len(a.BoundaryFN.ExistingSources) == 3 && containsAll(a.BoundaryFN.Supports, []string{StatusOrthogonalToBoundaryFN}) && containsAll(a.BoundaryFN.Failures, []string{StatusChiralityNoNineFive, StatusHiggsNoSixPS2, StatusNoBoundaryTraceMap, StatusNoPositiveRestSpectrum}), Detail: FormatBoundaryFN(a.BoundaryFN)},
			{Name: "separate force hierarchy from Yukawa hierarchy", Passed: a.LagrangianHierarchy.Audited && a.LagrangianHierarchy.Separated && containsAll(a.LagrangianHierarchy.Supports, []string{StatusGaugeGravitySeparate}) && containsAll(a.LagrangianHierarchy.Failures, []string{StatusForceNotYukawa, StatusAlphaOverAlphaGNoDeltaN, StatusGravityNotRestPressure}), Detail: strings.Join(a.LagrangianHierarchy.Failures, "; ")},
			{Name: "classify chirality idea status", Passed: a.IdeaStatus.Classified && len(a.IdeaStatus.UsefulFor) >= 4 && len(a.IdeaStatus.NotUseful) >= 7 && containsAll(a.IdeaStatus.Supports, []string{StatusUsefulFirewallAudit}) && containsAll(a.IdeaStatus.Failures, []string{StatusDoesNotSolveGate811}), Detail: FormatStatus(a.IdeaStatus)},
			{Name: "preserve C_Higgs Level-B bridge", Passed: a.Impact.Preserved && math.Abs(a.Impact.NEff-NEff) < 1e-15 && math.Abs(a.Impact.CYukawa-CYukawa) < 1e-15 && math.Abs(a.Impact.CHiggs-CHiggs) < 1e-15 && containsAll(a.Impact.Failures, []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}), Detail: "C_Higgs=(3/N_eff)C_History remains unchanged"},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoNaiveRealProjectors && a.Firewalls.NoSU2Shortcut && a.Firewalls.NoYukawaValues && a.Firewalls.NoRestPressure && a.Firewalls.NoBoundaryFNReplacement && a.Firewalls.NoForceHierarchyShortcut && a.Firewalls.NoLedgerUpdate && a.Firewalls.NoPoleMass && a.Firewalls.Verdict == StatusFirewallGate812, Detail: a.Firewalls.Verdict},
			{Name: "record branch decision", Passed: a.Branch.Recorded && strings.Contains(a.Branch.Next, "Gate 813") && strings.Contains(a.Branch.Next, "Boundary Second-Moment"), Detail: a.Branch.Next},
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
		notes := []string{a.Truth, FormatPseudoscalar(a.Pseudoscalar), FormatRequirement(a.WeakRestriction), FormatMassBridge(a.HiggsMassBridge), FormatRequirement(a.GradeZero), FormatRequirement(a.TraceMagnitude), FormatBoundaryFN(a.BoundaryFN), FormatStatus(a.IdeaStatus), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
