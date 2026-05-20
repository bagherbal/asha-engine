package generation2realformtrialityairlockandnativestatusfirewallaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-801-REAL-FORM-TRIALITY-AIRLOCK-NATIVE-STATUS-FIREWALL"
	theoremName = "Gate 801 — Real-Form Triality Airlock and Native-Status Firewall Audit"
)

func Generation2RealFormTrialityAirlockAndNativeStatusFirewallAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 801 analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusFirewallPreservedGate801}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 800 complex-only real-form result", Passed: a.Inheritance.ClAlgebra == "Cl(1,7) ≅ Mat(16,R)" && a.Inheritance.VolumeSquare == -1 && !a.Inheritance.RealChiralityCertified && strings.Contains(a.Inheritance.Outcome, "Outcome C") && containsAll(a.Inheritance.Verdicts, []string{StatusGate800Inherited, StatusComplexOnlyInherited, StatusNoFullNativeCL17}), Detail: a.Inheritance.Outcome},
			{Name: "define triality native-status ladder", Passed: a.StatusLevels.Defined && strings.Contains(a.StatusLevels.CurrentStatus, "T1") && a.StatusLevels.NotNative && a.StatusLevels.NotYukawa && containsAll(a.StatusLevels.Supports, []string{StatusCurrentT1ComplexOnly}) && containsAll(a.StatusLevels.Failures, []string{StatusT1NotNative, StatusT1NotYukawa}), Detail: FormatLevels(a.StatusLevels)},
			{Name: "define ComplexD4TrialityAirlock", Passed: a.ComplexAirlock.Defined && a.ComplexAirlock.Name == "ComplexD4TrialityAirlock" && containsAll(a.ComplexAirlock.Payload, []string{"so(8,C) D4 structure", "S3 outer automorphism", "real-descent obstruction ledger"}) && containsAll(a.ComplexAirlock.Supports, []string{StatusComplexAuxSearch}) && containsAll(a.ComplexAirlock.Failures, []string{StatusComplexAirlockNotNative}), Detail: FormatAirlock(a.ComplexAirlock)},
			{Name: "define compact and split real-form airlocks", Passed: a.CompactAirlock.Defined && a.SplitAirlock.Defined && containsAll(a.CompactAirlock.Failures, []string{StatusCompactNotNative, StatusNoWickTransport}) && containsAll(a.SplitAirlock.Supports, []string{StatusSplitUsefulSearch}) && containsAll(a.SplitAirlock.Failures, []string{StatusSplitNotNative}), Detail: FormatAirlock(a.CompactAirlock) + " | " + FormatAirlock(a.SplitAirlock)},
			{Name: "define real descent obstruction", Passed: a.Descent.Defined && strings.Contains(a.Descent.Map, "auxiliary triality object") && !a.Descent.NativeImport && containsAll(a.Descent.MustPreserve, []string{"real structure", "bilinear signatures", "trace/readout target"}) && containsAll(a.Descent.Failures, []string{StatusNoNativeImport, StatusAuxCannotBeNative}), Detail: FormatDescent(a.Descent)},
			{Name: "refine trilinear invariant status", Passed: a.Trilinear.Refined && strings.Contains(a.Trilinear.Formula, "γ") && a.Trilinear.NotYukawa && a.Trilinear.NoReadout && containsAll(a.Trilinear.MissingPackage, []string{"map to sector operators Y_u,Y_d,Y_e,Y_nu", "trace atom extraction x_i=y_i^2", "top-dominance breaking operator"}) && containsAll(a.Trilinear.Failures, []string{StatusTrilinearNotYukawa, StatusNoTrialityYukawaReadout}), Detail: FormatTrilinear(a.Trilinear)},
			{Name: "preserve N_eff firewall", Passed: a.NEff.Preserved && a.NEff.CertifiedSource == "color-tripled top dominance" && !a.NEff.AirlockChangesC && !a.NEff.ExplainsDelta && containsAll(a.NEff.Supports, []string{StatusAirlockFutureNEff}) && containsAll(a.NEff.Failures, []string{StatusAirlockNoNEff, StatusAirlockNoNEffMinus3}), Detail: FormatNEff(a.NEff)},
			{Name: "separate triality, GJ, SU3/A2, K7, and motif lanes", Passed: a.Lanes.Recorded && len(a.Lanes.Lanes) >= 5 && containsAll(a.Lanes.Blocked, []string{"GJ Clebsch three = D4 triality theorem", "A2 hexagon = D4 triality theorem", "K7 4|3 = triality carrier"}) && containsAll(a.Lanes.Failures, []string{StatusMotifNotTheorem}), Detail: strings.Join(a.Lanes.Lanes, "; ")},
			{Name: "record methodological status", Passed: a.Methodology.Recorded && strings.Contains(a.Methodology.HonestUse, "airlocked auxiliary") && strings.Contains(a.Methodology.InvalidUse, "direct explanation") && containsAll(a.Methodology.Supports, []string{StatusD4BranchUsefulAirlock}) && containsAll(a.Methodology.Failures, []string{StatusD4CannotAdvanceNative}), Detail: a.Methodology.HonestUse + "; blocked=" + a.Methodology.InvalidUse},
			{Name: "record branch decision", Passed: a.Branch.Recorded && strings.Contains(a.Branch.NextNative, "Complex D4 Trilinear") && strings.Contains(a.Branch.AlternativeTest, "External Yukawa") && a.Branch.Support == StatusNextTrilinearObstruct, Detail: a.Branch.NextNative},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoYukawa && a.Firewalls.NoEigenvalues && a.Firewalls.NoPMNSCKM && a.Firewalls.NoNEff && a.Firewalls.NoGJ && a.Firewalls.NoScalar && a.Firewalls.NoPoleMass && a.Firewalls.NoVEVGF && a.Firewalls.NoHistoryLoop && a.Firewalls.Verdict == StatusFirewallPreservedGate801, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLevels(a.StatusLevels), FormatAirlock(a.ComplexAirlock), FormatAirlock(a.CompactAirlock), FormatAirlock(a.SplitAirlock), FormatDescent(a.Descent), FormatTrilinear(a.Trilinear), FormatNEff(a.NEff), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
