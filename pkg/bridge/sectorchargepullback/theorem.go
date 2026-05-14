package sectorchargepullback

import "github.com/bagherbal/asha-engine/pkg/theorem"

func SectorChargePullbackCKMMoritaMisalignmentSieveTheorem() theorem.Theorem {
	const id = "BRIDGE-SECTOR-CHARGE-PULLBACK-CKM-MORITA-MISALIGNMENT-SIEVE"
	const name = "Sector-Charge Pullback / CKM Morita Misalignment Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 360 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "span inherits Gate 359 without adding fit", Passed: a.Span.InheritedGate == 359 && !a.Span.AddsFit, Detail: FormatSpan(a.Span)},
			{Name: "weak-isospin pullback is audited but not derived", Passed: a.Pullback.Formalized && !a.Pullback.NativeGeneratorSwap, Detail: FormatPullback(a.Pullback)},
			{Name: "candidate CKM overlaps are computed", Passed: a.CKM.Executed && len(a.CKM.Candidates) >= 3 && a.CKM.AnyCKMCapacity, Detail: FormatCKM(a.CKM)},
			{Name: "CKM texture remains unpromoted", Passed: !a.CKM.NativeAssignment && !a.CKM.NativeCKMDerived, Detail: FormatCKM(a.CKM)},
			{Name: "color trace norm does not pull global amplifier", Passed: a.Color.Audited && !a.Color.PullsGlobalAmplifier, Detail: FormatColor(a.Color)},
			{Name: "parameter census remains unreduced", Passed: a.Census.TotalReduction == 0 && a.Census.RemainingInputs == 15 && !a.Census.SevenSealReached, Detail: FormatCensus(a.Census)},
		}
		passed := 0
		for _, c := range checks {
			if c.Passed {
				passed++
			}
		}
		status := theorem.BridgeRequired
		if passed != len(checks) {
			status = theorem.FailedRoute
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks}
	}}
}
