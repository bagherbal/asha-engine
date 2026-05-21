package generation2alphavariationaltraceactionsourceobstructionaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-830-ALPHA-VARIATIONAL-TRACE-ACTION-SOURCE-OBSTRUCTION"
	theoremName = "Gate 830 — Alpha Variational / Trace-Action Source Obstruction Audit"
)

func Generation2AlphaVariationalTraceActionSourceObstructionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 830 alpha trace-action source audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 829 and select S_split -> alpha_B as the live wound", Passed: a.Ledger.AlphaSealedBridgeResponse && math.Abs(a.Ledger.AlphaB-0.0003878958469680527) < 1e-18 && containsAll(a.Trace.Verdicts, []string{StatusGate829Inherited, StatusLiveWoundSelected}), Detail: FormatLedger(a.Ledger)},
			{Name: "build formal support-trace expansion and reconstruct alpha_B", Passed: a.Trace.ReconstructsAlpha && math.Abs(a.Trace.AlphaFromTrace-a.Ledger.AlphaB) < 1e-18 && math.Abs(a.Trace.LinearContribution+a.Trace.QuadraticContribution-a.Ledger.AlphaB) < 1e-18 && containsAll(a.Trace.Supports, []string{SupportTraceRuleReconstructsAlpha, SupportTwoLaneSourceTyping}), Detail: FormatTrace(a.Trace)},
			{Name: "reject trace expansion as native source when X1 and X2 are not produced", Passed: a.Trace.ClassifiedAsRestatement && !a.Trace.X1NaturallyProduced && !a.Trace.X2NaturallyProduced && !a.Trace.TraceActionCertified && containsAll(a.Trace.Failures, []string{FailureTraceExpansionRestatesRule, FailureX1NotNative, FailureX2NotNative, FailureNoBoundaryTraceAction}), Detail: FormatTrace(a.Trace)},
			{Name: "audit linear/quadratic response-order story without certifying it", Passed: a.ResponseOrder.LinearPower == 1 && a.ResponseOrder.QuadraticPower == 2 && a.ResponseOrder.FirstSecondOrderInterpretationAllowedAsCandidate && !a.ResponseOrder.LinearOrderDerived && !a.ResponseOrder.QuadraticOrderDerived && !a.ResponseOrder.ResponseOrderTheoremCertified && containsAll(a.ResponseOrder.Failures, []string{FailureResponseOrderNotDerived, FailureNoBoundaryTraceAction}), Detail: FormatResponseOrder(a.ResponseOrder)},
			{Name: "test formal variational stationarity and reject formal repackaging", Passed: a.Variational.StationarityWorksFormally && a.Variational.WeightsAllTraceSourced && !a.Variational.PowersTypedByResponseOrder && !a.Variational.ActionNative && a.Variational.UsesInsertedAlphaRule && a.Variational.IsFormalRepackaging && !a.Variational.CertifiesAlphaTheorem && containsAll(a.Variational.Failures, []string{FailureVariationalRepackaging, FailureNoEulerLagrangeAlphaTheorem}), Detail: FormatVariational(a.Variational)},
			{Name: "enforce noncircular alpha direction", Passed: a.NonCircular.ComputesAlphaBeforeReadout && !a.NonCircular.UsesNEffToDefineAlpha && !a.NonCircular.UsesOfficialLedger && !a.NonCircular.UsesObservedYukawas && !a.NonCircular.UsesHiggsMass && strings.Contains(a.NonCircular.Direction, "never N_eff -> alpha_B"), Detail: FormatNonCircularity(a.NonCircular)},
			{Name: "block N_eff/C_Yukawa/C_Higgs updates after alpha-source obstruction", Passed: !a.Impact.CanPromoteAlpha && !a.Impact.CanPromoteOperatorNEff && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && strings.Contains(a.Impact.RecommendedAlphaStatus, "sealed bridge response") && containsAll(a.Impact.Failures, []string{FailureNoNEffSealReduction, FailureNoCYukawaUpdate, FailureNotR3SectorLedger}), Detail: FormatImpact(a.Impact)},
			{Name: "preserve physical and theorem-layer firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.TraceExpansionRestatement && a.Firewalls.X1X2NotNative && a.Firewalls.ResponseOrderOpen && a.Firewalls.NoTraceAction && a.Firewalls.VariationalRepackage && a.Firewalls.AlphaSealed && a.Firewalls.NoBoundaryAlphaMap && a.Firewalls.NoNEffSealReduction && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoSectorAssignment && a.Firewalls.NoPMNSCKM && a.Firewalls.NoHiggs && a.Firewalls.Verdict == StatusFirewallGate830, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatTrace(a.Trace), FormatResponseOrder(a.ResponseOrder), FormatVariational(a.Variational), FormatNonCircularity(a.NonCircular), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
