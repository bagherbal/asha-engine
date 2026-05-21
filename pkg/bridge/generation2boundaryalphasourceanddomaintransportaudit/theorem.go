package generation2boundaryalphasourceanddomaintransportaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-827-BOUNDARY-ALPHA-SOURCE-AND-DOMAIN-TRANSPORT"
	theoremName = "Gate 827 — BoundaryAlpha Source and Domain-Transport Audit"
)

func Generation2BoundaryAlphaSourceAndDomainTransportAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 827 boundary-alpha audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 826 pressure point and boundary split coordinate", Passed: math.Abs(a.Ledger.S-SBoundary) < 1e-18 && containsAll(a.Alpha.Supports, []string{SupportGate826Pressure, SupportR2PlusSharpened}) && containsAll(a.Alpha.Verdicts, []string{StatusBoundarySplitInherited}), Detail: FormatLedger(a.Ledger)},
			{Name: "verify linear coefficient 3/10 as triplet over vector-plus-boundary chamber", Passed: a.Coefficients.TripletRank == 3 && a.Coefficients.VectorDim == 8 && a.Coefficients.BoundaryDim == 2 && a.Coefficients.VectorBoundaryDim == 10 && math.Abs(a.Coefficients.LinearCoeff-0.3) < 1e-15 && containsAll(a.Coefficients.Supports, []string{SupportLinearCoeffSource}), Detail: FormatCoefficients(a.Coefficients)},
			{Name: "verify quadratic coefficient 7/72 as K7 over augmented Lambda4 chamber", Passed: a.Coefficients.K7Dim == 7 && a.Coefficients.Lambda4V8Dim == 70 && a.Coefficients.H72Dim == 72 && math.Abs(a.Coefficients.QuadraticCoeff-(7.0/72.0)) < 1e-15 && containsAll(a.Coefficients.Supports, []string{SupportQuadraticCoeffSource}), Detail: FormatCoefficients(a.Coefficients)},
			{Name: "reconstruct alpha_B from two-domain coefficient decomposition", Passed: math.Abs(a.Alpha.Alpha-a.Ledger.ExpectedAlphaB) < 1e-18 && math.Abs(a.Alpha.LinearContribution+a.Alpha.QuadraticContribution-a.Alpha.Alpha) < 1e-21 && a.Alpha.PowersSeparated && containsAll(a.Alpha.Verdicts, []string{StatusAlphaDecomposed, StatusPowersSeparated}), Detail: FormatAlpha(a.Alpha)},
			{Name: "distinguish dimension-ratio source candidates from a transport theorem", Passed: a.Transport.DimensionRatiosVerified && !a.Transport.LinearTransportCertified && !a.Transport.QuadraticTransportCertified && !a.Transport.UnifiedTransportCertified && !a.Transport.NativeAlphaTheoremCertified && containsAll(a.Transport.Failures, []string{FailureDimensionRatioOnly, FailureNoActivationMap, FailureLinearNoTransport, FailureQuadraticNoTransport}), Detail: FormatTransport(a.Transport)},
			{Name: "enforce noncircular direction from s to alpha to rest operator to N_eff", Passed: a.NonCircular.ComputesAlphaBeforeNEff && !a.NonCircular.UsesNEffToDefineAlpha && !a.NonCircular.UsesObservedYukawas && !a.NonCircular.UsesHiggsMass && strings.Contains(a.NonCircular.Direction, "s -> alpha_B") && containsAll(a.NonCircular.Verdicts, []string{StatusNonCircularity, StatusNoBackfit}), Detail: FormatNonCircularity(a.NonCircular)},
			{Name: "preserve C_Yukawa and C_Higgs freeze", Passed: !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && math.Abs(a.Impact.OfficialCYukawa-CYukawa) < 1e-15 && math.Abs(a.Impact.OfficialCHiggs-CHiggs) < 1e-15 && containsAll(a.Impact.Failures, []string{FailureNoCYukawaUpdate, FailureCHiggsLevelB, FailureNoTraceReadout, FailureNoSectorLedger}), Detail: FormatImpact(a.Impact)},
			{Name: "enforce physical and theorem-layer firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoActivationMap && a.Firewalls.DimensionRatioOnly && a.Firewalls.NoBackfit && a.Firewalls.NoTraceReadout && a.Firewalls.NoSectorLedger && a.Firewalls.NotR4 && a.Firewalls.NoCYukawaUpdate && a.Firewalls.CHiggsLevelB && a.Firewalls.NoPMNSCKM && a.Firewalls.NoHiggs && a.Firewalls.Verdict == StatusFirewallGate827, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatCoefficients(a.Coefficients), FormatAlpha(a.Alpha), FormatTransport(a.Transport), FormatNonCircularity(a.NonCircular), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
