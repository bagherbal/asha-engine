package gencurvature

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CurvatureOnGenerationCarrierTheorem() theorem.Theorem {
	const id = "MATTER-CURVATURE-GENERATION-CARRIER"
	const name = "curvature on protected generation carrier search"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerMatter,
		Status: theorem.OpenTest,
		Verify: func() theorem.Result {
			a, err := BuildDefault()
			if err != nil {
				return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct curvature-on-generation carrier", Passed: false, Detail: err.Error()}},
				}
			}
			return theorem.Result{
				ID:     id,
				Name:   name,
				Layer:  theorem.LayerMatter,
				Status: theorem.OpenTest,
				Checks: []theorem.Check{
					{Name: "protected carrier extracted", Passed: a.CarrierDimension == 3 && a.ContactKernelCount == 3, Detail: fmt.Sprintf("carrier dim=%d, contact kernel count=%d, gram residual=%.3e", a.CarrierDimension, a.ContactKernelCount, a.CarrierGramResidual)},
					{Name: "active contact carrier extracted", Passed: a.ActiveDimension == 4, Detail: fmt.Sprintf("active dim=%d, gram residual=%.3e", a.ActiveDimension, a.ActiveGramResidual)},
					{Name: "contact-side curvature operators", Passed: len(a.Operators) == 6 && len(a.ActiveOperators) == 6, Detail: fmt.Sprintf("computed %d pairwise R^K_AB operators on protected carrier and %d on active carrier", len(a.Operators), len(a.ActiveOperators))},
					{Name: "protected carrier is curvature-flat", Passed: a.NonzeroOperators == 0 && a.MaxCurvatureNorm < 1e-8, Detail: fmt.Sprintf("protected nonzero=%d/%d, max norm=%.6e", a.NonzeroOperators, len(a.Operators), a.MaxCurvatureNorm)},
					{Name: "active carrier receives curvature", Passed: a.ActiveNonzeroOperators > 0 && a.ActiveMaxCurvatureNorm > 1e-8, Detail: fmt.Sprintf("active nonzero=%d/%d, max norm=%.6e, span rank=%d", a.ActiveNonzeroOperators, len(a.ActiveOperators), a.ActiveMaxCurvatureNorm, a.ActiveOperatorSpanRank)},
					{Name: "non-diagonal generation action not selected", Passed: !a.NonDiagonalMixingFound && a.OperatorSpanRank == 0, Detail: fmt.Sprintf("protected max off-diagonal norm=%.6e, protected span rank=%d", a.MaxOffDiagonalNorm, a.OperatorSpanRank)},
					{Name: "full so(3) generation carrier theorem", Passed: a.FullSO3CarrierFound, Detail: fmt.Sprintf("fullSO3=%t; protected span rank=%d, closure relative=%.3e", a.FullSO3CarrierFound, a.OperatorSpanRank, a.ClosureRelative)},
					{Name: "symmetric mass texture still open", Passed: !a.SymmetricMassTextureFound && !a.CKMDerived && !a.PMNSDerived, Detail: "current curvature acts on active Higgs/contact directions, not the protected generation carrier; Yukawa masses require another bridge"},
				},
				Notes: []string{
					a.TruthStatement,
					"Gate 30 tests the contact-side mirror curvature R^K_AB = P_KAP_CBP_K − P_KBP_CAP_K. The result is a useful no-go: the protected 3D generation carrier is flat under this curvature, while the active 4D Higgs/contact carrier is not.",
				},
			}
		},
	}
}
