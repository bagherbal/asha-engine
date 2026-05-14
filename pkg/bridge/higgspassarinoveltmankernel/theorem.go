package higgspassarinoveltmankernel

import "github.com/bagherbal/asha-engine/pkg/theorem"

func HiggsPassarinoVeltmanPoleKernelFiniteIntegralInstallationAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-HIGGS-PASSARINO-VELTMAN-POLE-KERNEL-FINITE-INTEGRAL-INSTALLATION-AUDIT"
	const name = "Higgs Passarino-Veltman Pole Kernel / Finite One-Loop Integral Installation Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 334 PV kernel audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 333 component ledger inherited", Passed: a.Inputs.HighestInheritedGate == inheritedHighestGate && a.Inputs.NativeRunMassGeV > 125 && a.Inputs.MuGeV == a.Inputs.NativeRunMassGeV, Detail: FormatInputs(a.Inputs)},
			{Name: "finite PV basis formalized", Passed: a.Basis.EqualMassLane && a.Basis.BelowThresholdOnly && a.Basis.A0FiniteDefinition != "" && a.Basis.B0FiniteDefinition != "", Detail: FormatBasis(a.Basis)},
			{Name: "finite A0/B0 values computed for top W Z H", Passed: len(a.PV.Values) == 4 && allFiniteBelowThreshold(a.PV), Detail: FormatPVLedger(a.PV)},
			{Name: "on-shell kernel slots installed but not closed", Passed: a.Slots.AllBasisAvailable && !a.Slots.FullKernelClosed && noFullCoefficients(a.Slots), Detail: FormatSlots(a.Slots)},
			{Name: "firewalls preserve no exact pole-mass claim", Passed: a.Firewalls.NoCoefficientTable && a.Firewalls.NoCounterterms && a.Firewalls.NoGaugeScheme && a.Firewalls.NoExactPoleClaim, Detail: FormatFirewalls(a.Firewalls)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary), FormatStatuses(Statuses(a))}}
	}}
}

func allFiniteBelowThreshold(pv PVLedger) bool {
	for _, v := range pv.Values {
		if v.ThresholdRatio <= 1 || v.B0Finite != v.B0Finite || v.A0FiniteGeV2 != v.A0FiniteGeV2 {
			return false
		}
	}
	return true
}

func noFullCoefficients(slots KernelSlots) bool {
	for _, s := range slots.Slots {
		if s.FullCoefficientKnown || s.FiniteCountertermKnown || !s.InstalledBasisAvailable {
			return false
		}
	}
	return true
}
