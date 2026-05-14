package canonicalaction

import (
	"fmt"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CanonicalFiniteVariationalActionTheorem() theorem.Theorem {
	const id = "BRIDGE-CANONICAL-FINITE-VARIATIONAL-ACTION"
	const name = "canonical finite variational action and second-variation selection"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build canonical finite variational action", Passed: false, Detail: err.Error()}}}
		}
		terms := make([]string, 0, len(a.Terms))
		for _, t := range a.Terms {
			state := "derived"
			if !t.Derived {
				state = "not-derived"
			}
			terms = append(terms, fmt.Sprintf("%s [%s]: %s", t.Name, state, t.Formula))
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "canonical action assembled", Passed: len(a.Terms) == 4, Detail: a.ActionFormula},
			{Name: "scalar kinetic normalization selected", Passed: a.ScalarKineticSelected, Detail: fmt.Sprintf("K_phi=I_%d, trace=%.10f, rank=%d", a.ActiveRealDimension, a.ScalarKineticTrace, a.ScalarKineticRank)},
			{Name: "broken-orbit second variation computed", Passed: a.BrokenSecondVariationSelected, Detail: fmt.Sprintf("raw diag=%s; charged unit=%.10f; selected diag=%s", FormatFloatSlice(a.BrokenRawDiagonal), a.BrokenChargedUnit, FormatFloatSlice(a.BrokenSelectedDiagonal))},
			{Name: "full gauge kinetic Hessian selected", Passed: a.FullGaugeHessianSelected, Detail: fmt.Sprintf("K_EW=%s; eigenvalues=%s; rank=%d", FormatMatrix(a.FullGaugeHessian), FormatFloatSlice(a.FullGaugeHessianEigenvalues), a.FullGaugeHessianRank)},
			{Name: "U(1) completion coefficient selected", Passed: a.U1CompletionCoefficientSelected, Detail: fmt.Sprintf("%s; %s", a.KappaSelectionEquation.Equation, a.KappaSelectionEquation.Detail)},
			{Name: "generation-breaking source map selected", Passed: a.GenerationSourceSelected, Detail: fmt.Sprintf("%s; eig=%s; traceless=%s; rank=%d; mixing=%t", a.GenerationSource.Name, FormatFloatSlice(a.GenerationSource.Eigenvalues), FormatFloatSlice(a.GenerationSource.TracelessEigenvalues), a.GenerationSource.Rank, a.GenerationSource.ProducesMixing)},
			{Name: "active-to-generation mixing remains sealed", Passed: !a.ActiveToGenerationMixingSelected && a.SourceAction.NaturalSelectsZero, Detail: "the old 3x4 Hom(H_active,H_generation) source action still selects zero; the new source is an intrinsic diagonal generation-spectrum map, not a fitted mixing tensor"},
			{Name: "canonical finite action selected", Passed: a.CanonicalActionSelected, Detail: "dimensionless variational selection succeeds without observed constants"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no alpha, thetaW, v, Higgs mass, fermion masses, CKM, or PMNS values were inserted"},
			{Name: "physical constants remain unclaimed", Passed: !a.PhysicalCouplingsDerived && !a.PhysicalMassesDerived && !a.CKMPMNSDerived, Detail: "selected dimensionless action data is not yet a physical scale/RG theorem"},
		}, Notes: []string{
			a.TruthStatement,
			"action terms: " + strings.Join(terms, " | "),
			"rejected claims: " + Join(a.RejectedClaims),
			"remaining unknowns: " + Join(a.RemainingUnknowns),
			"Next: " + a.RecommendedNextGate,
		}}
	}}
}
