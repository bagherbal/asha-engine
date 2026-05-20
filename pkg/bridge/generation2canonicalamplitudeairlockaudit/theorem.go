package generation2canonicalamplitudeairlockaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2CanonicalAmplitudeAirlockForBoundaryWeightedDeficitClosureAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 666 — Canonical Amplitude Airlock for BoundaryWeightedDeficitClosure Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate666 amplitude-airlock audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate665 amplitude coordinate seal", Passed: a.Inherited.CoordinateSealInherited && a.Inherited.AmplitudeNatural && !a.Inherited.RGNativeInverseNatural && !a.Inherited.CoordinateRobust && math.Abs(a.Inherited.AmplitudeWBestMinus7Over72) < 1e-6 && math.Abs(a.Inherited.InverseWBestMinus7Over72) > 0.01 && a.Inherited.NoNativeDualRoot && a.Inherited.NoNativeSevenOver72 && a.Inherited.NoNativeTransport && a.Inherited.NoBoundaryStress, Detail: FormatInherited(a.Inherited)},
			{Name: "audit coordinate stack", Passed: len(a.CoordinateStack.Rows) == 5 && a.CoordinateStack.AmplitudeLayerPasses && !a.CoordinateStack.InverseKineticLayerPasses && !a.CoordinateStack.StrengthLayerPasses && !a.CoordinateStack.LogLayerPasses, Detail: FormatCoordinateStack(a.CoordinateStack)},
			{Name: "audit kinetic-to-amplitude nonlinearity", Passed: a.KineticToAmplitude.AmplitudeResidual > 0.04 && a.KineticToAmplitude.AmplitudeResidual < 0.06 && a.KineticToAmplitude.InverseOverAmplitude > 1.8 && a.KineticToAmplitude.InverseOverAmplitude < 2.1 && a.KineticToAmplitude.AmplitudeScalarScaleGap < a.KineticToAmplitude.InverseScalarScaleGap, Detail: FormatKineticToAmplitude(a.KineticToAmplitude)},
			{Name: "audit recurring root/amplitude pattern", Passed: len(a.Pattern.Rows) == 5 && strings.Contains(a.Pattern.Pattern, "root") && strings.Contains(a.Pattern.Verdict, StatusRootAmplitudePatternSupported), Detail: FormatPattern(a.Pattern)},
			{Name: "define amplitude airlock theorem target", Passed: strings.Contains(a.Target.MissingMap, "inverse-kinetic") && strings.Contains(a.Target.CandidateTheorem, "CanonicalAmplitudeAirlockTheorem") && strings.Contains(a.Target.Verdict, StatusNoNativeAmplitudeAirlockTheorem), Detail: FormatTarget(a.Target)},
			{Name: "classify source type", Passed: a.Source.Classification == "BoundaryWeightedDeficitClosureAmplitudeSeal" && len(a.Source.Statements) == 4 && strings.Contains(a.Source.Verdict, StatusAmplitudeLayerSupported) && strings.Contains(a.Source.Verdict, StatusInverseKineticFails), Detail: FormatSource(a.Source)},
			{Name: "preserve firewalls", Passed: !a.Discipline.ClaimsNativeAmplitudeAirlockTheorem && !a.Discipline.ClaimsNativeDualRootTheorem && !a.Discipline.ClaimsNativeSevenOver72Theorem && !a.Discipline.ClaimsNativeTransportTheorem && !a.Discipline.ClaimsBoundaryStressDerivation && !a.Discipline.ClaimsHiggsPrediction && !a.Discipline.ClaimsScalarStability && !a.Discipline.ClaimsFlavorDerivation && !a.Discipline.ClaimsGaugeUnification && !a.Discipline.ClaimsCKMPMNSDerivation && a.Discipline.Verdict == StatusGate666Boundary, Detail: FormatDiscipline(a.Discipline)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
