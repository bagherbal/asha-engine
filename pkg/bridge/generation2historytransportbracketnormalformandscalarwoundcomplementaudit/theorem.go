package generation2historytransportbracketnormalformandscalarwoundcomplementaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HistoryTransportBracketNormalFormAndScalarWoundComplementAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 759 — History Transport Bracket Normal Form and Scalar-Wound Complement Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate759 history transport bracket audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate758 one-eighth factorization", Passed: a.Gate758.Inherited && a.Gate758.ThreeFactorNormalFormAvailable && !a.Gate758.IndependentScalarRuntimeTheorem && math.Abs(a.Gate758.CYukawa-cYukawaMZ) < 1e-15 && math.Abs(a.Gate758.CHistory-cHistoryMZ) < 1e-12 && math.Abs(a.Gate758.LambdaRuntimeEff-lambdaRuntimeEffMZ) < 1e-15 && strings.Contains(a.Gate758.FactorizationFormula, "(1/8) C_Yukawa C_History"), Detail: FormatGate758(a.Gate758)},
			{Name: "define and compute History transport bracket", Passed: a.Bracket.BracketDefined && a.Bracket.OmegaComputed && !a.Bracket.PhysicalTimeOrRGScale && math.Abs(a.Bracket.LHopf-lHopf) < 1e-18 && math.Abs(a.Bracket.CHistory-cHistoryMZ) < 1e-12 && math.Abs(a.Bracket.OmegaHistory-0.9556769569304386) < 1e-12 && math.Abs(a.Bracket.OmegaResidual) < 1e-18 && strings.Contains(a.Bracket.OmegaFormula, "F_wall_3_red") && strings.Contains(a.Bracket.CHistoryFormula, "L_Hopf Omega_History"), Detail: FormatBracket(a.Bracket)},
			{Name: "define reduced scalar matching deficit and complement", Passed: a.Deficit.Defined && a.Deficit.RewrittenAsComplement && !a.Deficit.NativeScalarTheorem && math.Abs(a.Deficit.KappaLambdaRed-0.04432304306956136) < 1e-12 && math.Abs(a.Deficit.Complement-a.Bracket.OmegaHistory) < 1e-15 && math.Abs(a.Deficit.ComplementResidual) < 1e-18 && strings.Contains(a.Deficit.Definition, "F_wall_3_red") && strings.Contains(a.Deficit.Symbol, "kappa_lambda_red"), Detail: FormatDeficit(a.Deficit)},
			{Name: "write C_History and full scalar-Higgs normal form", Passed: a.NormalForm.CHistoryWritten && a.NormalForm.FullFormRewritten && a.NormalForm.ThreeFactorNormalForm && !a.NormalForm.IndependentRuntimeTheorem && math.Abs(a.NormalForm.CHistory-cHistoryMZ) < 1e-12 && math.Abs(a.NormalForm.LambdaRuntimeFromNormalForm-lambdaRuntimeEffMZ) < 1e-15 && math.Abs(a.NormalForm.NormalFormResidual) < 1e-15 && strings.Contains(a.NormalForm.CHistoryReducedFormula, "1-kappa_lambda_red") && strings.Contains(a.NormalForm.FullScalarHiggsFormula, "[1+L_Hopf(1-kappa_lambda_red)]"), Detail: FormatNormalForm(a.NormalForm)},
			{Name: "record source-type interpretation", Passed: a.Interpretation.Recorded && !a.Interpretation.NativeHistoryLoopTheorem && !a.Interpretation.NativeTransportTheorem && strings.Contains(a.Interpretation.KappaLambdaRedSourceType, "scalar matching deficit") && strings.Contains(a.Interpretation.OmegaHistorySourceType, "complement") && strings.Contains(a.Interpretation.CHistorySourceType, "HistoryLoop"), Detail: FormatInterpretation(a.Interpretation)},
			{Name: "audit layer separation", Passed: a.Layers.LayerSeparationAudited && a.Layers.FactorsMultiplyAfterScalarCollapse && !a.Layers.OperatorsOnSameNativeBoard && strings.Contains(a.Layers.CYukawaLayer, "Yukawa") && strings.Contains(a.Layers.KappaLambdaRedLayer, "boundary") && strings.Contains(a.Layers.LHopfLayer, "Radial-Hopf"), Detail: FormatLayers(a.Layers)},
			{Name: "reject illegal term identifications", Passed: a.Illegal.Audited && !a.Illegal.KappaLambdaRedNativeScalarTheorem && !a.Illegal.OmegaHistoryPhysicalTimeOrRGScale && !a.Illegal.LHopfBoundaryEventProbability && !a.Illegal.CHistoryNativeHistoryLoopTheorem && !a.Illegal.LambdaRuntimeEffIndependentPrediction && !a.Illegal.TreeProxyPoleMassPrediction && !a.Illegal.ClaimsYukawaEigenvaluesDerived && !a.Illegal.ClaimsHiggsMassOrPoleMassTheorem, Detail: FormatIllegal(a.Illegal)},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := append([]string{a.Truth}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
