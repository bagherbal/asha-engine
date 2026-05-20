package generation2d4trialitycarrierpackageandcl17realformaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-800-D4-TRIALITY-CARRIER-PACKAGE-CL17-REAL-FORM"
	theoremName = "Gate 800 — D4 Triality Carrier Package Requirement and Cl(1,7) Real-Form Audit"
)

func Generation2D4TrialityCarrierPackageAndCL17RealFormAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 800 analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusFirewallPreservedGate800}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 799 D4 branch status", Passed: a.Inheritance.Gate799Inherited && a.Inheritance.D4Selected && closeAbs(a.Inheritance.NEff, 3.0023273474722147, 1e-15) && a.Inheritance.CurrentCertifiedThree == "color-tripled top dominance" && a.Inheritance.NotGeneration && a.Inheritance.NotD4, Detail: StatusGate799Inherited + "; " + StatusD4Selected},
			{Name: "define core real-form triality question", Passed: a.Board.VectorDimR == 8 && a.Board.N == 8 && a.Board.DifferenceMod8 == 2 && a.Board.VolumeSquare == -1 && !a.Board.RealChiralityProjectorsExist && a.Board.MinimalRealSpinorDim == 16 && strings.Contains(a.Board.ComplexifiedSpin, "so(8,C)"), Detail: FormatBoard(a.Board)},
			{Name: "record complex D4 candidate but separate real form", Passed: a.ComplexD4.Recorded && a.ComplexD4.LawfulComplexCandidate && !a.ComplexD4.RealNative && a.ComplexD4.OuterAutomorphism == "Out(D4)≅S3 over C" && a.ComplexD4.Support == StatusComplexD4Shape && a.ComplexD4.Failure == StatusComplexOuterNotReal, Detail: a.ComplexD4.Complexification + "; " + a.ComplexD4.OuterAutomorphism},
			{Name: "define real-form preservation test and fail native shortcut", Passed: a.RealForm.Defined && strings.Contains(a.RealForm.TestEquation, "tau sigma") && strings.Contains(a.RealForm.PreservesRealForm, "not certified") && containsAll(a.RealForm.Failures, []string{StatusNoNativeUnlessPreserves, StatusComplexNotNative}), Detail: FormatRealForm(a.RealForm)},
			{Name: "audit carrier dimensions and signatures", Passed: a.Signatures.Defined && !a.Signatures.CarrierDimensionCompatible && !a.Signatures.SignatureCompatible && strings.Contains(a.Signatures.VectorSignature, "1,7") && containsAll(a.Signatures.Failures, []string{StatusDim8NotEnough, StatusSignatureMismatch, StatusNoCarrierWithoutDims}), Detail: FormatSignatures(a.Signatures)},
			{Name: "define Clifford trilinear pre-Yukawa requirement", Passed: a.Trilinear.Defined && a.Trilinear.RequiredBeforeYukawa && a.Trilinear.NotYukawa && strings.Contains(a.Trilinear.Formula, "γ") && a.Trilinear.Support == StatusTrilinearPreYukawa && a.Trilinear.Failure == StatusTrilinearNotYukawa, Detail: a.Trilinear.Formula + "; " + a.Trilinear.RealStatus},
			{Name: "define S3 triality relation audit", Passed: a.S3Test.Defined && !a.S3Test.RealClosureCertified && containsAll(a.S3Test.Relations, []string{"tau_3cycle^3", "tau_swap^2", "tau_swap tau_3cycle"}) && a.S3Test.Failure == StatusNoS3AuditNoTriality, Detail: strings.Join(a.S3Test.Relations, "; ") + "; " + a.S3Test.Reason},
			{Name: "classify real-form outcome", Passed: a.Outcome.Defined && a.Outcome.Selected != "" && !a.Outcome.FullNativeFound && a.Outcome.OutcomeARequired && containsAll(a.Outcome.Supports, []string{StatusComplexOnlyTriality, StatusTrialityNeedsAirlock, StatusOutcomeARequired}) && containsAll(a.Outcome.Failures, []string{StatusNoFullNativeCL17D4, StatusAirlockRequired}), Detail: a.Outcome.Selected},
			{Name: "check existing ASHA objects against triality carrier role", Passed: len(a.Existing) >= 7 && containsAll([]string{FormatExisting(a.Existing)}, []string{StatusK7Not8, StatusK7HodgeNotCarrier, StatusLambda4NotModule, StatusAggregateTracesNotCarrier}), Detail: FormatExisting(a.Existing)},
			{Name: "preserve triality-to-Yukawa readout firewall", Passed: a.YukawaFirewall.Defined && a.YukawaFirewall.NotYukawaTheorem && containsAll(a.YukawaFirewall.MissingObjects, []string{"YukawaSectorAssignment", "TraceAtomReadout", "TopDominanceBreakingOperator", "GenerationMixingReadout"}) && containsAll(a.YukawaFirewall.Failures, []string{StatusTrialityCarrierNotYukawa, StatusNoTraceAtomReadout, StatusNoTopBreakingOperator, StatusNoPMNSCKMReadout}), Detail: strings.Join(a.YukawaFirewall.MissingObjects, "; ")},
			{Name: "separate GJ, SU3/A2, and D4 lanes", Passed: a.Lanes.Recorded && strings.Contains(a.Lanes.GeorgiJarlskog, "Clebsch") && strings.Contains(a.Lanes.SU3A2, "not D4") && strings.Contains(a.Lanes.D4Triality, "outer automorphism") && containsAll(a.Lanes.Failures, []string{StatusGJNotD4, StatusA2NotD4, StatusMotifNotEvidence}), Detail: a.Lanes.GeorgiJarlskog + "; " + a.Lanes.SU3A2 + "; " + a.Lanes.D4Triality},
			{Name: "define success criteria and reject generic Spin8 statement", Passed: a.Criteria.Defined && a.Criteria.GenericSpin8Rejected && containsAll(a.Criteria.AllowedResults, []string{StatusNoFullNativeCL17D4, StatusComplexOnlyTriality, StatusTrialityNeedsAirlock}) && a.Criteria.Failure == StatusGenericSpin8NotEnough, Detail: strings.Join(a.Criteria.AllowedResults, "; ")},
			{Name: "record branch decision", Passed: a.Branch.Recorded && strings.Contains(a.Branch.SelectedOutcome, "Outcome C") && strings.Contains(a.Branch.Next, "Real-Form Triality Airlock") && strings.Contains(a.Branch.AlternativeIfFull, "Trilinear") && strings.Contains(a.Branch.AlternativeIfFail, "SU3/A2"), Detail: a.Branch.Next},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoYukawa && a.Firewalls.NoPMNSCKM && a.Firewalls.NoNEff && a.Firewalls.NoGJ && a.Firewalls.NoScalarRuntime && a.Firewalls.NoPoleMass && a.Firewalls.NoHistoryLoop && a.Firewalls.Verdict == StatusFirewallPreservedGate800, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatBoard(a.Board), FormatRealForm(a.RealForm), FormatSignatures(a.Signatures), FormatExisting(a.Existing), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
