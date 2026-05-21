package generation2totalrelativetracemagnitudeoperatorandledgerconsistencyaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-829-TOTAL-RELATIVE-TRACE-MAGNITUDE-OPERATOR-LEDGER-CONSISTENCY"
	theoremName = "Gate 829 — Total Relative TraceMagnitude Operator and Ledger Consistency Audit"
)

func Generation2TotalRelativeTraceMagnitudeOperatorAndLedgerConsistencyAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 829 total operator ledger audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 828 obstruction and use alpha_B only as sealed bridge input", Passed: math.Abs(a.Ledger.AlphaB-0.0003878958469680527) < 1e-18 && containsAll(a.Operator.Verdicts, []string{StatusGate828Inherited, StatusAlphaSealedInput}) && containsAll(a.Operator.Failures, []string{FailureAlphaNotNative, FailureNoBoundaryAlphaMap}), Detail: FormatLedger(a.Ledger)},
			{Name: "assemble total relative operator from I3 top block and B-L rest transfer", Passed: len(a.Operator.TopBlock) == 3 && len(a.Operator.RestBlock) == 4 && len(a.Operator.TotalSpectrum) == 7 && a.Operator.TopBlock[0] == 1 && a.Operator.TopBlock[1] == 1 && a.Operator.TopBlock[2] == 1 && strings.Contains(a.Operator.Formula, "I_3") && containsAll(a.Operator.Supports, []string{SupportTotalOperatorGivenAlpha, SupportGate826RestTransfer}), Detail: FormatOperator(a.Operator)},
			{Name: "derive total trace and square trace from the operator", Passed: math.Abs(a.Operator.TraceTotal-(3+3*a.Ledger.AlphaB)) < 1e-15 && math.Abs(a.Operator.SquareTraceTotal-TotalSquareTrace(a.Ledger.AlphaB)) < 1e-15 && math.Abs(a.Operator.RestSquareTrace-a.Operator.ExpectedRestSquareTrace) < 1e-21 && containsAll(a.Operator.Supports, []string{SupportTraceFormula, SupportSquareTraceFormula}), Detail: FormatOperator(a.Operator)},
			{Name: "derive operator N_eff and confirm absolute T cancels", Passed: a.Operator.AbsoluteTScaleCancels && math.Abs(a.Operator.OperatorNEffFromTrace-a.Operator.OperatorNEffClosedForm) < 1e-15 && math.Abs(a.Operator.OperatorNEffFromTrace-a.Ledger.OperatorNEff) < 1e-18 && containsAll(a.Operator.Verdicts, []string{StatusAbsoluteTCancels, StatusOperatorNEffDerived}), Detail: FormatOperator(a.Operator)},
			{Name: "separate operator, BFN-truncated, and official frozen ledgers", Passed: a.Consistency.LedgerSeparationEnforced && !a.Consistency.SilentCollapseDetected && !a.Consistency.OfficialUsedAsCandidate && math.Abs(a.Consistency.OperatorNEff-a.Consistency.BFNTruncatedNEff) < 1e-15 && math.Abs(a.Consistency.OperatorNEff-a.Consistency.OfficialNEff) > 1e-10 && containsAll(a.Consistency.Supports, []string{SupportLedgerSeparation, SupportOfficialFreeze}), Detail: FormatConsistency(a.Consistency)},
			{Name: "record fifth-order operator-vs-BFN residual", Passed: math.Abs(a.Consistency.OperatorMinusBFN-a.Consistency.FifthOrderResidual) < 1e-15 && math.Abs(a.Consistency.FifthOrderResidual) < 1e-14 && strings.Contains(a.Consistency.FifthOrderResidualFormula, "-24 alpha_B^5"), Detail: FormatConsistency(a.Consistency)},
			{Name: "freeze N_eff, C_Yukawa, and C_Higgs updates", Passed: !a.Consistency.CanPromoteOperatorToOfficial && !a.Consistency.CanUpdateNEff && !a.Consistency.CanUpdateCYukawa && !a.Consistency.CanUpdateCHiggs && math.Abs(a.Ledger.OfficialCYukawa-OfficialCYukawa) < 1e-15 && math.Abs(a.Ledger.OfficialCHiggs-OfficialCHiggs) < 1e-15 && containsAll(a.Consistency.Failures, []string{FailureNoNEffSealReduction, FailureNoCYukawaUpdate}), Detail: FormatConsistency(a.Consistency)},
			{Name: "preserve alpha-source, sector-ledger, and physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.AlphaNotNative && a.Firewalls.NoBoundaryAlphaMap && a.Firewalls.NotR3SectorLedger && a.Firewalls.NotR4NativeYukawa && a.Firewalls.NoNEffSealReduction && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NoSectorAssignment && a.Firewalls.NoVariationalAlpha && a.Firewalls.NoPMNSCKM && a.Firewalls.NoHiggs && a.Firewalls.Verdict == StatusFirewallGate829, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatOperator(a.Operator), FormatConsistency(a.Consistency), FormatSource(a.Source), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
