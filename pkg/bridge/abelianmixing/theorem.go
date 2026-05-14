package abelianmixing

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func AbelianMixingHyperchargeNormalizationTheorem() theorem.Theorem {
	const id = "BRIDGE-ABELIAN-MIXING-HYPERCHARGE-NORMALIZATION"
	const name = "abelian mixing / hypercharge coupling normalization search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build abelian mixing audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 73 abelian bridge input", Passed: a.AbelianBridgeDimension == 2, Detail: fmt.Sprintf("central/contact-u1 and B-L/contact-u1 expose a %dD abelian bridge domain", a.AbelianBridgeDimension)},
			{Name: "right-singlet hypercharge trace", Passed: abs(a.RightSingletTrace) < 1e-9, Detail: fmt.Sprintf("selected odd-branch right-singlet/conjugate table has Tr(Y)=%.3e", a.RightSingletTrace)},
			{Name: "central u(1) rejected for hypercharge", Passed: a.CentralRejectedForHypercharge && a.Central.Selected, Detail: fmt.Sprintf("%s=%.1f; unit central shift changes Tr(Y) by %.1f, so central coefficient is fixed to 0 for hypercharge", a.Central.Name, a.Central.Value, a.CentralTraceShiftPerUnit)},
			{Name: "B-L half-weight selected", Passed: a.BMinusLCoefficientSelected, Detail: fmt.Sprintf("%s=%.10f in Y=T3_R+(B-L)/2 at charge-table level", a.BMinusL.Name, a.BMinusL.Value)},
			{Name: "charge-level abelian hypercharge bridge", Passed: a.ChargeLevelHyperchargeBridge, Detail: "Y direction uses B-L half-weight plus matter T3_R; central u(1) is not part of hypercharge"},
			{Name: "hypercharge normalization inherited", Passed: abs(a.KY-5.0/3.0) < 1e-9, Detail: fmt.Sprintf("k_Y=%.10f, normalized factor=%.10f, boundary candidate sin²=%.10f", a.KY, a.NormalizedHyperchargeFactor, a.BoundarySin2)},
			{Name: "U(1) kinetic normalization", Passed: a.GaugeKineticNormalizationDerived, Detail: "not derived; charge-level coefficients do not fix gauge kinetic terms or g_Y"},
			{Name: "contact-u1 kinetic mixing", Passed: a.ContactU1KineticMixingDerived, Detail: "not derived; the contact-u1 field still lacks a finite kinetic Hessian with B-L/central sectors"},
			{Name: "physical U(1) coupling and alpha", Passed: a.PhysicalU1CouplingDerived && a.FineStructureDerived, Detail: "not derived; no physical alpha_em or low-energy weak angle is computed"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed alpha, theta_W, v, or measured coupling was inserted"},
		}, Notes: []string{a.TruthStatement, "Next: " + a.RecommendedNextGate}}
	}}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
