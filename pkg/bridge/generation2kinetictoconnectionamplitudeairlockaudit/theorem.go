package generation2kinetictoconnectionamplitudeairlockaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2KineticToConnectionAmplitudeAirlockSourceAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 667 — Kinetic-to-Connection Amplitude Airlock Source Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate667 kinetic-to-connection audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate666 amplitude seal", Passed: a.Inherited.AmplitudeSealInherited && a.Inherited.AmplitudeLayerPasses && !a.Inherited.InverseKineticLayerPasses && a.Inherited.InverseOverAmplitude > 1.8 && math.Abs(a.Inherited.AmplitudeWBestMinus7Over72) < 1e-6 && math.Abs(a.Inherited.InverseWBestMinus7Over72) > 0.01 && a.Inherited.MissingAirlockTheorem && a.Inherited.NoNativeSevenOver72 && a.Inherited.NoNativeDualRoot && a.Inherited.NoNativeTransport && a.Inherited.NoBoundaryStress, Detail: FormatInherited(a.Inherited)},
			{Name: "define kinetic coordinate", Passed: strings.Contains(a.Kinetic.NativeCoordinate, "1/g_i^2") && strings.Contains(a.Kinetic.Verdict, StatusKineticCoordinateDefined) && strings.Contains(a.Kinetic.Verdict, StatusInverseKineticStillFails), Detail: FormatKinetic(a.Kinetic)},
			{Name: "audit canonical field rescaling", Passed: a.Rescaling.AmplitudeCoordinateTyped && strings.Contains(a.Rescaling.AlgebraicRelation, "u_i^{-1/2}") && strings.Contains(a.Rescaling.ConnectionAmplitude, "g_i") && strings.Contains(a.Rescaling.Verdict, StatusGaugeAmplitudeSourcedByConnection), Detail: FormatRescaling(a.Rescaling)},
			{Name: "compare gauge coordinate layers", Passed: len(a.Coordinates.Rows) == 5 && a.Coordinates.AmplitudeOnlyPasses && a.Coordinates.InverseKineticFails && strings.Contains(a.Coordinates.ClosureCoordinate, "g3/gEW"), Detail: FormatCoordinates(a.Coordinates)},
			{Name: "audit electroweak Hessian socket", Passed: a.HessianSocket.CompatibleWithClosure && len(a.HessianSocket.AmplitudeObjects) == 4 && strings.Contains(a.HessianSocket.NeutralHessianShape, "g^2") && strings.Contains(a.HessianSocket.Verdict, StatusCanonicalEndpointSocketSupported), Detail: FormatHessian(a.HessianSocket)},
			{Name: "audit scalar side type", Passed: !a.ScalarSide.NativeAmplitude && strings.Contains(a.ScalarSide.Verdict, StatusScalarRuntimeShadow), Detail: FormatScalar(a.ScalarSide)},
			{Name: "audit recurring root/amplitude pattern", Passed: len(a.Pattern.Rows) == 5 && strings.Contains(a.Pattern.Verdict, StatusRootAmplitudePatternSupported), Detail: FormatPattern(a.Pattern)},
			{Name: "define kinetic-amplitude theorem target", Passed: strings.Contains(a.Target.Name, "Kinetic") && strings.Contains(a.Target.Airlock, "u_i -> g_i") && strings.Contains(a.Target.Verdict, StatusNoNativeKineticAmplitudeTheorem), Detail: FormatTarget(a.Target)},
			{Name: "classify source type", Passed: a.Source.Classification == "BoundaryWeightedDeficitClosureConnectionAmplitudeSeal" && len(a.Source.Statements) == 4 && strings.Contains(a.Source.Verdict, StatusClosureConnectionAmplitudeLayer) && strings.Contains(a.Source.Verdict, StatusNoNativeKineticAmplitudeTheorem), Detail: FormatSource(a.Source)},
			{Name: "preserve firewalls", Passed: !a.Discipline.ClaimsNativeKineticAmplitudeTheorem && !a.Discipline.ClaimsNativeSevenOver72Theorem && !a.Discipline.ClaimsNativeDualRootTheorem && !a.Discipline.ClaimsNativeTransportTheorem && !a.Discipline.ClaimsBoundaryStressDerivation && !a.Discipline.ClaimsHiggsPrediction && !a.Discipline.ClaimsScalarStability && !a.Discipline.ClaimsFlavorDerivation && !a.Discipline.ClaimsGaugeUnification && !a.Discipline.ClaimsCKMPMNSDerivation && a.Discipline.Verdict == StatusGate667Boundary, Detail: FormatDiscipline(a.Discipline)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
