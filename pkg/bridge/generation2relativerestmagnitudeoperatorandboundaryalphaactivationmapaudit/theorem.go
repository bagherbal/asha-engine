package generation2relativerestmagnitudeoperatorandboundaryalphaactivationmapaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-825-RELATIVE-REST-MAGNITUDE-OPERATOR-BOUNDARY-ALPHA-ACTIVATION-MAP"
	theoremName = "Gate 825 — Relative RestMagnitude Operator and BoundaryAlpha Activation Map Audit"
)

func Generation2RelativeRestMagnitudeOperatorAndBoundaryAlphaActivationMapAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 824 and select BoundaryToTraceMagnitudeRestMap as live gap", Passed: math.Abs(a.Ledger.AlphaB-0.0003878958469680527) < 1e-15 && math.Abs(a.Ledger.DeltaNBFN-0.002327375081808316) < 1e-15 && containsAll(Statuses(), []string{StatusGate824Inherited, StatusLiveGapSelected}), Detail: FormatLedger(a.Ledger)},
			{Name: "define relative positive spectrum and cancel absolute T", Passed: len(a.Operator.Spectrum) == 7 && math.Abs(a.Operator.ATotalOverT-(3+3*a.Ledger.AlphaB)) < 1e-15 && containsAll(a.Operator.Supports, []string{SupportRelativeNoAbsoluteT}), Detail: FormatOperator(a.Operator)},
			{Name: "validate rest trace, beta and q_rest formulas", Passed: math.Abs(a.Operator.ARestOverT-3*a.Ledger.AlphaB) < 1e-15 && math.Abs(a.Operator.BRestOverT2-(3*a.Ledger.AlphaB*a.Ledger.AlphaB-6*math.Pow(a.Ledger.AlphaB, 3)+12*math.Pow(a.Ledger.AlphaB, 4))) < 1e-21 && math.Abs(a.Operator.Beta-(a.Ledger.AlphaB*a.Ledger.AlphaB-2*math.Pow(a.Ledger.AlphaB, 3)+4*math.Pow(a.Ledger.AlphaB, 4))) < 1e-21 && math.Abs(a.Operator.QRest-(1.0/3.0-(2.0/3.0)*a.Ledger.AlphaB+(4.0/3.0)*a.Ledger.AlphaB*a.Ledger.AlphaB)) < 1e-15, Detail: FormatOperator(a.Operator)},
			{Name: "derive N_eff operator and fifth-order residual to BFN closure", Passed: math.Abs(a.Operator.NEffOperator-a.Ledger.NEffBFN) < 1e-15 && math.Abs(a.Operator.SymbolicResidual+2.107593378826735e-16) < 1e-27 && containsAll(a.Operator.Supports, []string{SupportFourthOrderClosure}), Detail: FormatOperator(a.Operator)},
			{Name: "audit alpha_B and six coefficient sources without promotion", Passed: len(a.Coefficients) == 2 && containsAll(a.Coefficients[0].Supports, []string{SupportAlphaBSourceShape}) && containsAll(a.Coefficients[1].Supports, []string{SupportSixFromParticipation, SupportBoundaryPairSecondary}) && containsAll(a.Coefficients[1].Failures, []string{FailureSixNotNumericalClosure}), Detail: FormatCoefficients(a.Coefficients)},
			{Name: "audit projector and source lanes", Passed: len(a.Sources) == 4 && a.Sources[0].SuppliesShape && !a.Sources[0].CertifiedTraceReadout && containsAll(a.Sources[0].Failures, []string{FailureProjectiveNoReadout}) && containsAll(a.Sources[2].Failures, []string{FailureFiniteTripleNoOperator}) && containsAll(a.Sources[3].Failures, []string{FailureD4NotOperator, FailureD4NoP1P3Readout}), Detail: FormatSources(a.Sources)},
			{Name: "define BoundaryAlphaActivationMap as unsupplied missing object", Passed: !a.Activation.Certified && strings.Contains(a.Activation.RequiredObject, "BoundaryAlphaActivationMap") && containsAll(a.Activation.Failures, []string{FailureEigenvaluesNoActivation, FailureAlphaNotNative}), Detail: FormatActivation(a.Activation)},
			{Name: "execute noncircularity audit", Passed: a.NonCircular.PassesInputFirewall && a.NonCircular.OperatorWouldAvoidExternalData && strings.Contains(FormatNonCircularity(a.NonCircular), "observed Higgs mass") && containsAll(a.NonCircular.Failures, []string{FailureOperatorHypothesisIfSealed}), Detail: FormatNonCircularity(a.NonCircular)},
			{Name: "classify R2+ candidate shape without C_Yukawa update", Passed: strings.Contains(a.Status.Level, "R2+ candidate") && strings.Contains(a.Status.Outcome, "Outcome B") && !a.Status.CanUpdateCYukawa && !a.Status.HasCertifiedProjectors && !a.Status.HasBoundaryActivationMap && !a.Status.HasSectorLedger && !a.Status.NativeYukawaTheorem, Detail: a.Status.Level + " -> " + a.Status.NextGate},
			{Name: "preserve C_Yukawa and C_Higgs firewall", Passed: math.Abs(a.Impact.CYukawaCandidate-0.9992248096922658) < 1e-15 && math.Abs(a.Impact.CHiggsCandidate-1.0372205108665145) < 1e-15 && math.Abs(a.Impact.OfficialCYukawa-CYukawa) < 1e-15 && containsAll(a.Impact.Failures, []string{FailureNoCYukawaUpdate, FailureCHiggsLevelB}), Detail: FormatImpact(a.Impact)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoProjectorsNoNativeOperator && a.Firewalls.ProjectiveNeedsReadout && a.Firewalls.AlphaNeedsActivation && a.Firewalls.NoSectorAssignment && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoCYukawaUpdate && a.Firewalls.CHiggsLevelB && a.Firewalls.NoD4 && a.Firewalls.NoFiniteYukawa && a.Firewalls.Verdict == StatusFirewallGate825, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatOperator(a.Operator), FormatCoefficients(a.Coefficients), FormatSources(a.Sources), FormatActivation(a.Activation), FormatNonCircularity(a.NonCircular), a.Status.Level, a.Status.Outcome, a.Status.NextGate, FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
