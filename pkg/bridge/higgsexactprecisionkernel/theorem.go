package higgsexactprecisionkernel

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ExactNativeHiggsPredictionArbitraryPrecisionNumericalKernelAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-EXACT-NATIVE-HIGGS-PREDICTION-ARBITRARY-PRECISION-NUMERICAL-KERNEL-AUDIT"
	const name = "Exact Native Higgs Prediction / Arbitrary-Precision Numerical Kernel Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 335 exact precision audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 334 inherited and exact rational inputs installed", Passed: a.Inputs.HighestInheritedGate == inheritedHighestGate && a.Inputs.ContactShapeRational.Cmp(rat(1197, 4624)) == 0 && a.Inputs.LambdaHRational.Cmp(rat(1197, 9248)) == 0, Detail: FormatInputs(a.Inputs)},
			{Name: "high-precision Machin pi and alpha inverse computed", Passed: a.Pi.PrecisionBits == precisionBits && nearlyFloat(a.Pi.AlphaInverse, 25.132741228718345, 1e-14), Detail: FormatPi(a.Pi)},
			{Name: "native closed-form Higgs proxy computed without float64", Passed: nearlyFloat(a.Native.MassGeV, 125.27415714969897, 1e-12) && a.Native.LambdaH.Cmp(rat(1197, 9248)) == 0, Detail: FormatNative(a.Native)},
			{Name: "exact precision gap and self-energy target computed", Passed: nearlyRat(a.Gap.RequiredRePiGeV2, 43.60444956747405, 1e-12) && a.Gap.RequiredRePiGeV2.RatString() == "504067437/11560000", Detail: FormatGap(a.Gap)},
			{Name: "efficiency ledger preserves deterministic exact native branch", Passed: !a.Efficiency.UsesFloat64ForNativeBranch && a.Efficiency.Deterministic && !a.Efficiency.FullPVContractionExecuted, Detail: FormatEfficiency(a.Efficiency)},
			{Name: "full-pole firewalls preserved", Passed: a.Firewalls.NoPVContraction && a.Firewalls.NoCounterterms && a.Firewalls.NoTwoLoop && a.Firewalls.NoColliderPoleClaim, Detail: FormatFirewalls(a.Firewalls)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary), FormatStatuses(Statuses(a))}}
	}}
}
