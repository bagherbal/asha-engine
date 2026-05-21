package generation2bminusltracezeroresttransferfactorizationaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-826-B-MINUS-L-TRACE-ZERO-REST-TRANSFER-FACTORIZATION"
	theoremName = "Gate 826 — B-L Trace-Zero Rest-Transfer Factorization Audit"
)

func Generation2BMinusLTraceZeroRestTransferFactorizationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 826 B-L rest-transfer audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 825 alpha and relative operator target", Passed: math.Abs(a.Ledger.AlphaB-0.0003878958469680527) < 1e-15 && math.Abs(a.Ledger.NEffOperator-a.Ledger.NEffBFN) < 1e-15, Detail: FormatLedger(a.Ledger)},
			{Name: "verify P1/P3 Fock projectors and B-L selector", Passed: a.Projectors.P1Rank == 1 && a.Projectors.P3Rank == 3 && a.Projectors.Orthogonal && a.Projectors.Complete && math.Abs(a.Projectors.TraceBMinusL) < 1e-15 && containsAll(a.Projectors.Verdicts, []string{StatusSelectorInherited, StatusProjectorsVerified, StatusBMinusLVerified}), Detail: FormatProjectors(a.Projectors)},
			{Name: "construct Q_BL trace-zero transfer operator", Passed: math.Abs(a.Projectors.TraceQ) < 1e-15 && math.Abs(a.Projectors.TraceP3Q+3) < 1e-15 && math.Abs(a.Projectors.TraceQ2-12) < 1e-15 && containsAll(a.Projectors.Supports, []string{SupportTraceZeroTransferCarrier}), Detail: FormatProjectors(a.Projectors)},
			{Name: "reconstruct Gate 825 rest spectrum from alpha P3 plus alpha squared Q_BL", Passed: a.Factorization.MaxAbsResidual < 1e-18 && math.Abs(a.Factorization.FactorizedRest[0]-3*a.Ledger.AlphaB*a.Ledger.AlphaB) < 1e-21 && math.Abs(a.Factorization.FactorizedRest[1]-a.Ledger.AlphaB*(1-a.Ledger.AlphaB)) < 1e-21 && containsAll(a.Factorization.Supports, []string{SupportBMinusLRestTransfer, SupportEigenvaluesNotManual}), Detail: FormatFactorization(a.Factorization)},
			{Name: "prove quadratic B-L term is trace preserving", Passed: math.Abs(a.Factorization.TraceQuadratic) < 1e-21 && math.Abs(a.Factorization.TraceRest-3*a.Ledger.AlphaB) < 1e-18 && containsAll(a.Factorization.Supports, []string{SupportQuadraticRedistribution}), Detail: FormatFactorization(a.Factorization)},
			{Name: "source square-trace coefficients from projector traces", Passed: math.Abs(a.Factorization.SquareTrace-(3*math.Pow(a.Ledger.AlphaB, 2)-6*math.Pow(a.Ledger.AlphaB, 3)+12*math.Pow(a.Ledger.AlphaB, 4))) < 1e-21 && len(a.Factorization.TraceCoefficients) == 3 && a.Factorization.TraceCoefficients[0] == 3 && a.Factorization.TraceCoefficients[1] == -6 && a.Factorization.TraceCoefficients[2] == 12 && containsAll(a.Factorization.Supports, []string{SupportCoefficientsFromTraces}), Detail: FormatFactorization(a.Factorization)},
			{Name: "verify positivity window without sector assignment", Passed: a.Positivity.WindowCertified && a.Positivity.ActiveNonnegative && a.Positivity.SingletEigenvalue >= 0 && a.Positivity.TripletEigenvalue >= 0, Detail: FormatPositivity(a.Positivity)},
			{Name: "separate transfer factorization from alpha source and sector ledger", Passed: a.Boundary.CertifiedTransferFactorization && !a.Boundary.CertifiedAlphaSource && !a.Boundary.CertifiedTraceReadout && !a.Boundary.CertifiedSectorLedger && strings.Contains(a.Boundary.NextGate, "BoundaryAlpha Source") && containsAll(a.Boundary.Failures, []string{FailureAlphaNotSourced, FailureNoBoundaryAlphaTheorem, FailureNoSectorLedger}), Detail: FormatBoundary(a.Boundary)},
			{Name: "preserve C_Yukawa and C_Higgs freeze", Passed: !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && math.Abs(a.Impact.OfficialCYukawa-CYukawa) < 1e-15 && math.Abs(a.Impact.OfficialCHiggs-CHiggs) < 1e-15 && containsAll(a.Impact.Failures, []string{FailureNoCYukawaUpdate, FailureCHiggsLevelB}), Detail: FormatImpact(a.Impact)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.AlphaUnsourced && a.Firewalls.NoSectorLedger && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoCYukawaUpdate && a.Firewalls.CHiggsLevelB && a.Firewalls.NoD4 && a.Firewalls.NoPMNSCKM && a.Firewalls.NoHiggs && a.Firewalls.Verdict == StatusFirewallGate826, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatProjectors(a.Projectors), FormatFactorization(a.Factorization), FormatPositivity(a.Positivity), FormatBoundary(a.Boundary), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
