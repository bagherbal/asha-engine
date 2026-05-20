package generation2centralbaselinegaugeandscalarwallreferenceselectionaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2CentralBaselineGaugeAndScalarWallReferenceSelectionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: AuditID, Name: "Gate 707 — Central Baseline Gauge and Scalar-Wall Reference Selection Audit", Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: AuditID, Name: "Gate 707 — Central Baseline Gauge and Scalar-Wall Reference Selection Audit", Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate706 central-baseline/uplift isolation audit", Passed: a.Inherited.CentralBaselineUpliftInherited && a.Inherited.CentralBaselineIdentityShift && a.Inherited.NontrivialContentK7Uplift && a.Inherited.DBaseBaselineSubtracted && a.Inherited.BaselineDoesNotSelectK7OrRho72 && a.Inherited.NoNativeScalarBaseline && a.Inherited.NoNativeK7Uplift && a.Inherited.NoNativeUpliftTheorem && a.Inherited.NoNativeHistoryResponse && a.Inherited.NoNativeSevenOver72 && math.Abs(a.Inherited.ExpectedUplift-pK7*(lambdaLambda12+r3Minus1)) < tolerance && a.Inherited.Verdict == StatusGate706CentralBaselineUpliftInherited, Detail: FormatInheritance(a.Inherited)},
			{Name: "define central baseline gauge family", Passed: strings.Contains(a.GaugeFamily.GeneralForm, "c I_H72") && strings.Contains(a.GaugeFamily.ExpectationFormula, "65/72") && math.Abs(a.GaugeFamily.K7Correction-(r3Minus1-a.GaugeFamily.ArbitraryC)) < tolerance && math.Abs(a.GaugeFamily.ComplementCorrection-(math.Abs(lambdaLambda12)-a.GaugeFamily.ArbitraryC)) < tolerance && math.Abs(a.GaugeFamily.ExpectationAtC-a.GaugeFamily.RawExpectation) < tolerance && a.GaugeFamily.GaugeInvariant && strings.Contains(a.GaugeFamily.Verdict, StatusCentralBaselineGaugeFamilyDefined) && strings.Contains(a.GaugeFamily.Verdict, StatusTotalExpectationBaselineGaugeInvariant), Detail: FormatGaugeFamily(a.GaugeFamily)},
			{Name: "compute active scalar baseline choice", Passed: math.Abs(a.ActiveChoice.C-math.Abs(lambdaLambda12)) < tolerance && strings.Contains(a.ActiveChoice.Observable, "|lambda|I_H72") && math.Abs(a.ActiveChoice.K7Uplift-(r3Minus1-math.Abs(lambdaLambda12))) < tolerance && math.Abs(a.ActiveChoice.K7Uplift-(lambdaLambda12+r3Minus1)) < tolerance && math.Abs(a.ActiveChoice.ComplementUplift) < tolerance && a.ActiveChoice.ComplementZero && a.ActiveChoice.K7LocalResponse && a.ActiveChoice.UniqueComplementZeroChoice && strings.Contains(a.ActiveChoice.Verdict, StatusComplementZeroUpliftForAbsLambda) && strings.Contains(a.ActiveChoice.Verdict, StatusAbsLambdaUniqueBaselineForK7LocalUplift), Detail: FormatActiveChoice(a.ActiveChoice)},
			{Name: "audit typed baseline alternatives", Passed: len(a.Alternatives) == 5 && alternativeAcceptedCount(a.Alternatives) == 1 && hasAlternative(a.Alternatives, "scalar zero-wall baseline", true) && hasAlternative(a.Alternatives, "gauge baseline", false) && hasAlternative(a.Alternatives, "zero/raw two-payoff", false) && hasAlternative(a.Alternatives, "xi boundary baseline", false) && hasAlternative(a.Alternatives, "midpoint baseline", false) && allAlternativeExpectationsInvariant(a.Alternatives), Detail: FormatAlternatives(a.Alternatives)},
			{Name: "support-locality selects scalar baseline with K7 uplift", Passed: strings.Contains(a.SupportLocality.SupportLocalityCondition, "zero P_perp") && strings.Contains(a.SupportLocality.RawTwoPayoffForm, "P_perp") && strings.Contains(a.SupportLocality.ScalarBaselineForm, "S_split P_K7") && a.SupportLocality.ComplementZeroForAbs && a.SupportLocality.ForcesCAbsLambda && a.SupportLocality.Gate696Consistent && strings.Contains(a.SupportLocality.Verdict, StatusSupportLocalitySelectsScalarBaseline) && strings.Contains(a.SupportLocality.Verdict, StatusScalarBaselineSelectionSupportLocalGauge) && strings.Contains(a.SupportLocality.Verdict, StatusK7UpliftFormSharperThanRawTwoPayoff), Detail: FormatSupportLocality(a.SupportLocality)},
			{Name: "audit scalar-wall airlock compatibility", Passed: strings.Contains(a.Airlock.ScalarWallReference, "scalar zero-wall") && a.Airlock.UsesScalarZeroWallDepth && a.Airlock.GaugeBaselineAlgebraic && a.Airlock.GaugeBaselineReversesSector && a.Airlock.CompatibleWithAirlock && a.Airlock.Verdict == StatusScalarWallAirlockCompatibilityAudited, Detail: FormatAirlock(a.Airlock)},
			{Name: "classify source types", Passed: strings.Contains(a.SourceTypes.CentralGaugeFamilyRole, "cI") && strings.Contains(a.SourceTypes.AbsLambdaRole, "complement-zero") && strings.Contains(a.SourceTypes.K7UpliftRole, "K7") && strings.Contains(a.SourceTypes.GaugeBaselineRole, "P_perp") && strings.Contains(a.SourceTypes.TruthBoundary, "not native"), Detail: FormatSourceTypes(a.SourceTypes)},
			{Name: "preserve missing theorem boundary", Passed: len(a.Missing.Missing) == 5 && strings.Contains(a.Missing.Verdict, StatusBaselineChoiceNotNativeYet) && strings.Contains(a.Missing.Verdict, StatusNoNativeScalarBaselineReferenceSelection) && strings.Contains(a.Missing.Verdict, StatusNoNativeK7RatherThanComplementCarriesUplift) && strings.Contains(a.Missing.Verdict, StatusNoNativeBoundaryWoundUpliftTheorem) && strings.Contains(a.Missing.Verdict, StatusNoNativeSevenOver72Theorem), Detail: FormatMissing(a.Missing)},
			{Name: "preserve Gate707 central-baseline gauge firewall", Passed: !a.Firewall.ClaimsBaselineChoiceNative && !a.Firewall.ClaimsNativeScalarBaselineReferenceSelection && !a.Firewall.ClaimsNativeK7RatherThanComplementUplift && !a.Firewall.ClaimsNativeBoundaryWoundUpliftTheorem && !a.Firewall.ClaimsNativeSevenOver72Theorem && !a.Firewall.ClaimsBoundaryStressDerived && !a.Firewall.ClaimsScalarRGMatching && !a.Firewall.ClaimsHiggsMass && !a.Firewall.ClaimsGaugeUnification && !a.Firewall.ClaimsFlavorDerivation && !a.Firewall.ClaimsCKMPMNS && a.Firewall.Verdict == StatusGate707CentralBaselineGaugeBoundary, Detail: FormatFirewall(a.Firewall)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: AuditID, Name: "Gate 707 — Central Baseline Gauge and Scalar-Wall Reference Selection Audit", Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func alternativeAcceptedCount(xs []BaselineAlternativeAudit) int {
	count := 0
	for _, x := range xs {
		if x.Accepted {
			count++
		}
	}
	return count
}

func hasAlternative(xs []BaselineAlternativeAudit, name string, accepted bool) bool {
	for _, x := range xs {
		if x.Name == name && x.Accepted == accepted {
			return true
		}
	}
	return false
}

func allAlternativeExpectationsInvariant(xs []BaselineAlternativeAudit) bool {
	want := pK7*r3Minus1 + pComplement*math.Abs(lambdaLambda12)
	for _, x := range xs {
		if math.Abs(x.Expectation-want) > tolerance {
			return false
		}
	}
	return true
}
