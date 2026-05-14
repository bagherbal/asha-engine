package betacoeff

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteSpectrumBetaCoefficientAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-SPECTRUM-BETA-AUDIT"
	const name = "finite spectrum and beta-coefficient bridge audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct beta audit", Passed: false, Detail: err.Error()}}}
		}
		inv := a.Inventory
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "finite field inventory", Passed: inv.Generations == 3 && inv.WeylStatesPerGeneration == 16, Detail: fmt.Sprintf("generations=%d, Weyl states/gen=%d, total Weyl states=%d", inv.Generations, inv.WeylStatesPerGeneration, inv.WeylStatesTotal)},
			{Name: "left doublet inventory", Passed: inv.QuarkDoubletsPerGeneration == 3 && inv.LeptonDoubletsPerGeneration == 1, Detail: fmt.Sprintf("quark doublets/gen=%d, lepton doublets/gen=%d, SU(2) Weyl index/gen=%.10f", inv.QuarkDoubletsPerGeneration, inv.LeptonDoubletsPerGeneration, inv.WeakDoubletWeylIndexPerGeneration)},
			{Name: "color and hypercharge inventory", Passed: inv.ColorTripletWeylIndexPerGeneration > 0 && inv.HyperchargeY2PerGeneration > 0, Detail: fmt.Sprintf("SU(3) Weyl index/gen=%.10f, ΣY²/gen=%.10f", inv.ColorTripletWeylIndexPerGeneration, inv.HyperchargeY2PerGeneration)},
			{Name: "finite scalar inventory", Passed: inv.ComplexScalarDoublets == 1 && inv.ScalarRealDirections == 4, Detail: fmt.Sprintf("complex scalar doublets=%d, real active directions=%d, scalar ΣY²=%.10f", inv.ComplexScalarDoublets, inv.ScalarRealDirections, inv.ScalarHyperchargeY2Sum)},
			{Name: "continuum one-loop formula assumption", Passed: a.ContinuumOneLoopFormulaUsed, Detail: "uses β(g)=b g³/(16π²), Weyl contribution (2/3)T(R), complex-scalar contribution (1/3)T(R), and gauge contribution −(11/3)C2(G)"},
			{Name: "candidate beta coefficients", Passed: a.DerivedFromFiniteInventory, Detail: fmt.Sprintf("b1=%.10f, b2=%.10f, b3=%.10f; components %s", a.B1GUTNormalized, a.B2, a.B3, FormatComponents(a.Components))},
			{Name: "imported Standard Model beta table", Passed: !a.ImportedSMBetaTable, Detail: "coefficients are reconstructed from the finite representation inventory plus stated continuum assumptions, not imported as a lookup table"},
			{Name: "finite beta theorem", Passed: a.FiniteBetaTheoremDerived, Detail: "not derived; current result still assumes the continuum one-loop field-theory formula"},
			{Name: "threshold spectrum", Passed: a.ThresholdSpectrumDerived, Detail: "not derived; finite heavy modes and matching rules remain open"},
			{Name: "gauge kinetic normalization", Passed: a.GaugeKineticNormalizationDerived, Detail: "not derived; beta coefficients alone do not fix initial couplings"},
			{Name: "physical running and alpha", Passed: a.PhysicalRunningDetermined && a.FineStructureDerived, Detail: "not derived; requires boundary scale, threshold matching, and normalized boundary coupling"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedCouplingsUsed, Detail: "no observed α_em, θ_W, or measured couplings were inserted"},
		}, Notes: []string{a.TruthStatement, fmt.Sprintf("minimum missing data: %v", a.MinimumMissingData)}}
	}}
}
