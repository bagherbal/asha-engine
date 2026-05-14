package loopoperator

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteLoopOperatorConstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-LOOP-OPERATOR-CONSTRUCTION"
	const name = "finite Fock/Yukawa loop-operator construction"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct finite loop operator", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "finite Yukawa incidence operator", Passed: a.FiniteYukawaIncidenceOperatorDerived, Detail: fmt.Sprintf("Y:%d×%d from H_left⊗H_Φ (%d domain states) to H_right (%d states); allowed scalar-fiber entries=%d", a.RightDimension, a.DomainDimension, a.DomainDimension, a.RightDimension, a.AllowedFiberEntries)},
			{Name: "operator rank and saturation", Passed: a.Rank == a.RightDimension && a.UnusedDomainEntries == 16, Detail: fmt.Sprintf("rank(Y)=%d, unused domain entries=%d, max column occupancy=%.0f", a.Rank, a.UnusedDomainEntries, a.MaxColumnOccupancy)},
			{Name: "right-channel Gram trace", Passed: a.NativeLoopTraceSkeletonDerived, Detail: fmt.Sprintf("Tr(YYᵀ)=%.10f, Tr(YᵀY)=%.10f, row norms² range=[%.10f, %.10f]", a.RightTrace, a.DomainTrace, a.MinRightRowNormSquared, a.MaxRightRowNormSquared)},
			{Name: "unit fermion-loop pressure skeleton", Passed: a.UnitFermionLoopPressure < 0, Detail: fmt.Sprintf("−Tr(YYᵀ)=%.10f as unit-incidence fermion-loop pressure", a.UnitFermionLoopPressure)},
			{Name: "top-like color skeleton", Passed: a.UpTypeFiberEntries == 6 && a.UnitTopLikeSkeleton == -6, Detail: fmt.Sprintf("up-type color fiber entries=%d, skeleton=%.0f·y_top-like²", a.UpTypeFiberEntries, a.UnitTopLikeSkeleton)},
			{Name: "top dominance selected", Passed: a.TopDominanceSelected, Detail: "open; unit-incidence operator gives equal row norm² to all right channels and does not select the top channel"},
			{Name: "Yukawa strengths", Passed: a.TopLikeYukawaStrengthDerived, Detail: "open; no observed y_t or fitted coupling appears in the operator"},
			{Name: "bosonic counter-operators", Passed: a.BosonicCounterOperatorDerived, Detail: "open; gauge/scalar positive loop operators still require kinetic normalization"},
			{Name: "regulator and renormalization", Passed: a.RegulatorOrRenormalizationDerived, Detail: "open; finite spectral regulator/cutoff not yet derived"},
			{Name: "native scalar mass sign", Passed: a.MuSquaredSignDerived && a.NativeEffectivePotentialComputed, Detail: "open; this constructs the incidence trace skeleton, not μ²_eff<0"},
			{Name: "hidden observed couplings", Passed: !a.HiddenObservedCouplingsUsed, Detail: "no observed y_t, g, g′, λ, v, or Higgs mass was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("kind pressures: %s", FormatKindTraces(a.KindTraces)),
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
