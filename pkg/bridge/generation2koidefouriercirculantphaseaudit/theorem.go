package generation2koidefouriercirculantphaseaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2KoideFourierCirculantPhaseAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 Koide Fourier/circulant phase audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate582 Koide Fourier phase audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate581 Koide coordinate beta runtime", Passed: a.Runtime.Mu0GeV > 0 && a.Runtime.Lambda12GeV > a.Runtime.Mu0GeV && a.Runtime.Gate581MZDThetaDeg > 0, Detail: FormatRuntime(a.Runtime)},
			{Name: "derive Fourier/circulant Koide frame", Passed: a.Formula.Formula != "" && a.Formula.KoideEquivalence != "" && len(a.Formula.CanonicalOrder) == 3, Detail: FormatFormula(a.Formula)},
			{Name: "compute canonical M_Z Fourier phase", Passed: a.MZ.DeltaDeg > 132 && a.MZ.DeltaDeg < 133 && a.MZ.MaxReconstructionError < 1e-15 && a.MZ.PlaneAmplitudeR > 0.9999 && a.MZ.PlaneAmplitudeR < 1.0001, Detail: FormatPoint(a.MZ)},
			{Name: "compute canonical Lambda_12 Fourier phase", Passed: a.Lambda12.DeltaDeg > 132 && a.Lambda12.DeltaDeg < 133 && a.Lambda12.MaxReconstructionError < 1e-15 && a.Lambda12.PlaneAmplitudeR > 0.9999 && a.Lambda12.PlaneAmplitudeR < 1.0001, Detail: FormatPoint(a.Lambda12)},
			{Name: "certify v1 phase stability and Koide amplitude sharpening", Passed: a.Transport.PhaseStable && a.Transport.AbsDriftDeg < 3e-4 && a.Transport.AmplitudeMovesTowardOne, Detail: FormatTransport(a.Transport)},
			{Name: "audit permutation/convention dependence and block simple phase certification", Passed: !a.Permutation.UniqueWithoutOrdering && !a.Permutation.SimplePhaseCertified && a.Permutation.BestResidualDeg > a.Permutation.CertificationDeg && len(a.Permutation.Phases) == 6, Detail: FormatPermutation(a.Permutation)},
			{Name: "preserve root-trace and flavor firewalls", Passed: !a.Firewalls.DerivesLeptonMasses && !a.Firewalls.DerivesYukawaEigenvalues && !a.Firewalls.DerivesKoide && !a.Firewalls.DerivesFourierPhase && !a.Firewalls.DerivesCKM && !a.Firewalls.DerivesPMNS && !a.Firewalls.DerivesGenerationHierarchy && !a.Firewalls.AddsNewCarrier && !a.Firewalls.PromotesObservedAsNative && a.Firewalls.PreservesGate352, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "return Gate582 final seal verdict", Passed: a.Final.SealName == "ChargedLeptonKoideFourierPhaseSeal" && a.Final.PhaseStableInV1 && !a.Final.SimpleRationalCertified && !a.Final.NativeSelectorCertified && a.Final.MinimalNextRequirement != "", Detail: FormatFinal(a.Final)},
		}
		notes := append(Statuses(), a.Truth, a.Final.MinimalNextRequirement)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
