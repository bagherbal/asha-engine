package contactchargenorm

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CenteredContactSpectralCurrentChargeOperatorNormalizationObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-CENTERED-CONTACT-SPECTRAL-CURRENT-CHARGE-OPERATOR-NORMALIZATION-OBSTRUCTION"
	const name = "Centered contact spectral current / charge-operator normalization obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build centered contact charge normalization obstruction", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 144 centered trace-zero signed diagnostic is inherited", Passed: a.ContactRows == 7 && a.CenteredPositiveRows == 3 && a.CenteredNegativeRows == 4 && a.CenteredZeroRows == 0 && a.BetaPermissionFirewallClosed, Detail: fmt.Sprintf("maxAbs=%.10f frob=%.10f range=%.10f", a.MaxAbs, a.FrobeniusNorm, a.SpectralRange)},
			{Name: "canonical diagnostic normalizations exist but remain seven-eigenvalue diagnostics", Passed: a.MaxAbsAudit.CanonicalAsDiagnostic && a.FrobeniusAudit.CanonicalAsDiagnostic && a.RangeAudit.CanonicalAsDiagnostic && a.RawAudit.DistinctEigenvalues == 7 && a.MaxAbsAudit.DistinctEigenvalues == 7 && a.FrobeniusAudit.DistinctEigenvalues == 7 && a.RangeAudit.DistinctEigenvalues == 7 && !a.RawAudit.ChargeOperatorSemantic, Detail: FormatAudits([]NormalizationAudit{a.RawAudit, a.MaxAbsAudit, a.FrobeniusAudit, a.RangeAudit})},
			{Name: "binary ±1/2 and balanced trace-zero split normalizations do not become T3R or hypercharge", Passed: a.BinaryHalfAudit.TwoLevel && a.BinaryHalfAudit.UniformMagnitude && !a.BinaryHalfAudit.TraceZero && !a.BinaryHalfAudit.T3RSemantic && a.BalancedSplitAudit.TwoLevel && a.BalancedSplitAudit.TraceZero && !a.BalancedSplitAudit.UniformMagnitude && !a.BalancedSplitAudit.HyperchargeSemantic, Detail: FormatAudits([]NormalizationAudit{a.BinaryHalfAudit, a.BalancedSplitAudit})},
			{Name: "charge-operator requirements remain unsatisfied beyond trace control", Passed: a.Requirements.TraceControlled && a.Requirements.ObservedInputFree && !a.Requirements.SelectedOrientation && !a.Requirements.FiniteChargeLattice && !a.Requirements.OperatorPullback && !a.Requirements.LocalFieldMap && !a.Requirements.GaugeRepresentationRows && !a.Requirements.MassActivation && !a.Requirements.DecouplingRule && !a.Requirements.AllSatisfied, Detail: FormatRequirements(a.Requirements)},
			{Name: "normalization audit counts preserve the beta firewall", Passed: a.NormalizationsAudited == 7 && a.AvailableNormalizations == 6 && a.CanonicalDiagnosticNorms == 6 && a.TraceZeroNormalizations == 5 && a.TwoLevelNormalizations == 2 && a.ChargeSemanticNormalizations == 0 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0, Detail: FormatSummary(a.Summary)},
			{Name: "no hidden observed input or physical constants leak through the normalization search", Passed: a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: a.TruthStatement},
		}}
	}}
}
