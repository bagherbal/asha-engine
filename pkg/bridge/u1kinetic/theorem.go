package u1kinetic

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func U1KineticMixingHessianSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-U1-KINETIC-MIXING-HESSIAN-SEARCH"
	const name = "U(1) kinetic mixing / gauge coupling Hessian search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build U(1) kinetic audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 74 charge-level input", Passed: a.ChargeLevelHyperchargeSelected, Detail: "Y=T3_R+(B-L)/2 selected at charge-table level; central u(1) rejected as hypercharge component"},
			{Name: "abelian field inventory", Passed: a.AbelianFieldCount == 3, Detail: fmt.Sprintf("fields=[central u(1), B-L, contact-u1]; carrier dims=[%d matter, %d scalar]", a.MatterCarrierDimension, a.ScalarCarrierDimension)},
			{Name: "matter-carrier trace Gram", Passed: a.MatterGramDerived && a.CentralBMinusLOrthogonal, Detail: fmt.Sprintf("K_matter=[[%.10f,0],[0,%.10f]], det=%.10f", a.Central.Trace2, a.BMinusL.Trace2, a.MatterGramDeterminant)},
			{Name: "central/B-L orthogonality", Passed: a.CentralBMinusLOrthogonal, Detail: fmt.Sprintf("Tr(I·(B-L))=0; Tr(I)=%.1f, Tr(B-L)=%.1e", a.Central.Trace, a.BMinusL.Trace)},
			{Name: "contact-u1 scalar norm", Passed: a.ContactU1NormDerived && math.Abs(a.ContactU1.Trace2-1) < 1e-9, Detail: fmt.Sprintf("T_phi trace=%.1e, Tr(T_phi²)=%.10f, norm=%.10f", a.ContactU1.Trace, a.ContactU1.Trace2, a.ContactU1Norm)},
			{Name: "block-diagonal diagnostic", Passed: a.BlockDiagnosticDeterminant > 0, Detail: fmt.Sprintf("diag diagnostic det=%.10f with cross-carrier entries set to zero only diagnostically", a.BlockDiagnosticDeterminant)},
			{Name: "cross-carrier kinetic mixing", Passed: a.CrossCarrierKineticMixingDerived, Detail: "not derived; B-L/contact-u1 and central/contact-u1 Hessian entries require a finite dual-carrier action"},
			{Name: "full U(1) kinetic Hessian", Passed: a.FullU1KineticHessianDerived, Detail: "not derived; trace Gram diagnostics are not the physical gauge-field kinetic matrix"},
			{Name: "physical U(1) coupling and alpha", Passed: a.PhysicalU1CouplingDerived && a.FineStructureDerived, Detail: "not derived; no g_Y, theta_W running, or alpha_em is computed"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed alpha, theta_W, v, or measured coupling was inserted"},
		}, Notes: []string{a.TruthStatement, "Next: " + a.RecommendedNextGate}}
	}}
}
