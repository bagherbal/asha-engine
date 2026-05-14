package matrixinvariantkoideaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func MatrixInvariantKoideTypeTracePolynomialAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-MATRIX-INVARIANT-KOIDE-TYPE-TRACE-POLYNOMIAL-AUDIT"
	const name = "Matrix Invariant / Koide-Type Trace Polynomial Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 351 matrix invariant audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "span inherits Gate 350 without fitting data", Passed: a.Span.InheritedGate == 350 && !a.Span.AddsEmpiricalFit, Detail: FormatSpan(a.Span)},
			{Name: "Koide/root-trace invariant is formalized", Passed: a.Koide.Formalized && a.Koide.Target > 0, Detail: FormatKoide(a.Koide)},
			{Name: "empirical charged-lepton Koide alignment is quarantined", Passed: len(a.EmpiricalSpectra) >= 3 && a.EmpiricalSpectra[0].Quarantined && a.EmpiricalSpectra[0].KoideK > 0.666 && a.EmpiricalSpectra[0].KoideK < 0.667, Detail: FormatSpectrum(a.EmpiricalSpectra[0])},
			{Name: "triality and B-gap invariants do not mandate K=2/3", Passed: a.Triality.Executed && !a.Triality.NativeTwoThirdsMandated && a.Triality.MagnitudeSquaredKoide < 0.5, Detail: FormatTriality(a.Triality)},
			{Name: "characteristic polynomial program requires a root-trace operator", Passed: a.Characteristic.Audited && a.Characteristic.OneConstraintCapacity && a.Characteristic.RequiresRootTraceOperator && !a.Characteristic.CharacteristicPolynomialLocked, Detail: FormatCharacteristic(a.Characteristic)},
			{Name: "parameter count remains fifteen", Passed: a.Reduction.StartingVacuumInputs == 15 && a.Reduction.ReductionProved == 0 && a.Reduction.RemainingVacuumInputs == 15 && !a.Reduction.SevenSealTargetReached, Detail: FormatReduction(a.Reduction)},
			{Name: "summary preserves empirical vacuum quarantine", Passed: a.Summary.Executed && !a.Summary.AnyInvariantPromoted && a.Summary.RemainingVacuumInputs == 15, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 351 catalogs Koide as a genuine matrix-invariant research lane, but rejects promotion because ASHA has not derived the required root-trace operator or characteristic-polynomial constraint."}}
	}}
}
