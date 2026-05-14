package scalarfockspectralpotential

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FockDiracScalarSpectralActionContactShapeComparisonTheorem() theorem.Theorem {
	const id = "BRIDGE-FOCK-DIRAC-SCALAR-SPECTRAL-ACTION-CONTACT-SHAPE-COMPARISON"
	const name = "Fock Dirac spectral-action scalar potential and contact quartic-shape comparison theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Fock/contact scalar-shape comparison", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "gauge-ratio theorem is inherited as closed input", Passed: a.Firewall.GaugeRatioAlreadyClosed && a.Previous.TraceAudit.BoundaryDiagMatched && a.Previous.TraceAudit.WeakAngleSeedMatched, Detail: "Gate167 boundary diag=(" + a.Previous.TraceAudit.NormalizedT1.String() + "," + a.Previous.TraceAudit.NormalizedT2.String() + "," + a.Previous.TraceAudit.NormalizedT3.String() + "," + a.Previous.TraceAudit.NormalizedY.String() + "), sin2=" + a.Previous.TraceAudit.WeakAngleSeed.String()},
			{Name: "Fock Dirac scalar spectral moments are computed", Passed: a.UnitYukawa.ChannelCount == 8 && a.UnitYukawa.QuadraticMomentA == 8 && a.UnitYukawa.QuarticMomentB == 8 && a.UnitYukawa.FullDiracTraceD2 == 16 && a.UnitYukawa.FullDiracTraceD4 == 16 && !a.UnitYukawa.UsesObservedInput, Detail: FormatYukawaMoment(a.UnitYukawa)},
			{Name: "scalar spectral action uses Yukawa moments, not a representation trace", Passed: a.SpectralAction.CutoffMomentsRequired && !a.SpectralAction.CutoffMomentsDerived && a.SpectralAction.ScalarShapeComparable && a.SpectralAction.ScalarShapeAmplitudeDependent && !a.SpectralAction.GaugeLikeRepresentationRigidity, Detail: FormatSpectralAction(a.SpectralAction)},
			{Name: "Gate37 contact/Higgs quartic shape is available independently", Passed: a.ContactShape.ActiveRealDimension == 4 && a.ContactShape.ProtectedDirectionCount == 3 && a.ContactShape.LambdaShape > 0 && a.ContactShape.PairDegenerate && a.ContactShape.ShiftedNormalForm && !a.ContactShape.ElectroweakScaleDerived && !a.ContactShape.HiggsMassDerived, Detail: FormatContactShape(a.ContactShape)},
			{Name: "unit-incidence scalar shape does not match Gate37 contact shape", Passed: !a.Comparison.UnitIncidenceMatchesContact && !a.Comparison.ConvergenceClosed, Detail: FormatComparison(a.Comparison)},
			{Name: "Gate37 scalar shape lies in the allowed Yukawa-moment range", Passed: a.Comparison.ContactWithinYukawaShapeRange && a.Comparison.ConstraintOpened && a.Comparison.ContactEffectiveSlots > 1 && a.Comparison.ContactEffectiveSlots < float64(a.UnitYukawa.ChannelCount), Detail: fmt.Sprintf("%s :: effective participation %.12g requires non-uniform finite Yukawa amplitudes", FormatComparison(a.Comparison), a.Comparison.ContactEffectiveSlots)},
			{Name: "amplitude sensitivity is explicit", Passed: a.DeformedYukawa.ChiralShape != a.UnitYukawa.ChiralShape && !a.DeformedYukawa.AmplitudeIndependent, Detail: FormatYukawaMoment(a.DeformedYukawa)},
			{Name: "scalar comparison opens the mass-texture constraint without deriving physical constants", Passed: !a.Firewall.ScalarConvergenceClosed && a.Firewall.ScalarAmplitudeConstraintOpened && !a.Firewall.YukawaAmplitudesDerived && !a.Firewall.FermionMassesDerived && !a.Firewall.CKMPMNSDerived && !a.Firewall.ElectroweakScaleDerived && !a.Firewall.HiggsMassDerived && !a.Firewall.ThresholdCorrectionsDerived && !a.Firewall.RGRunningDerived && !a.Firewall.PhysicalConstantsDerived && a.Firewall.ResidualNullityBefore == 3 && a.Firewall.ResidualNullityAfter == 3, Detail: FormatFirewall(a.Firewall) + " :: " + a.TruthStatement},
		}, Notes: []string{
			"Gate 168 tests the scalar sector after the gauge-ratio convergence of Gates 166-167.",
			"The gauge ratio is amplitude-independent; the scalar potential is not. Its finite spectral-action shape is B/A^2 with A=Σ|y_i|² and B=Σ|y_i|⁴.",
			"Unit incidence gives B/A^2=1/8, while Gate37 gives λ_shape≈0.258866782. This is a mismatch, not a second closure theorem.",
			"Because Gate37 lies inside the allowed range [1/8,1], the comparison becomes a concrete amplitude-texture target for the next gate rather than an observed-constant fit.",
		}}
	}}
}
