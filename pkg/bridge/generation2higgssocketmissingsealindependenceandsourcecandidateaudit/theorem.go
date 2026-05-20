package generation2higgssocketmissingsealindependenceandsourcecandidateaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HiggsSocketMissingSealIndependenceAndSourceCandidateAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: AuditID, Name: "Gate 720 — Higgs Socket Missing-Seal Independence and Source-Candidate Audit", Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: AuditID, Name: "Gate 720 — Higgs Socket Missing-Seal Independence and Source-Candidate Audit", Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate719 conditional Higgs socket", Passed: a.Gate719.Inherited && a.Gate719.SocketAssembled && a.Gate719.RequiresN && a.Gate719.RequiresQ && a.Gate719.RepresentationCompatible && !a.Gate719.NativeTwistorSelector && !a.Gate719.NativeHyperchargeNorm && !a.Gate719.CanonicalThetaH && !a.Gate719.PhysicalHiggsDoubletTheorem && !a.Gate719.HiggsMassOrRuntime && !a.Gate719.YukawaOperatorOrEigenvalue && a.Gate719.Verdict == StatusGate719ConditionalHiggsSocketInherited, Detail: FormatGate719(a.Gate719)},
			{Name: "audit n selector source candidates", Passed: a.NSelector.LivesInS2K7Minus && a.NSelector.SelectsJH && a.NSelector.SelectsPhaseLine && a.NSelector.SelectsComplexCarrier && len(a.NSelector.Candidates) == nSourceCandidateCount && !a.NSelector.NativeSelectorFound && a.NSelector.RequiresSelectorSeal && strings.Contains(a.NSelector.Verdict, StatusNoNativeTwistorSelectorN), Detail: FormatNSelector(a.NSelector)},
			{Name: "audit q normalization source candidates", Passed: a.QNorm.LivesInRNonzero && a.QNorm.NormalizesPhaseGenerator && len(a.QNorm.Candidates) == qSourceCandidateCount && a.QNorm.CanMatchTargetConvention && !a.QNorm.NativeQDerived && a.QNorm.RequiresNormalizationSeal && strings.Contains(a.QNorm.Verdict, StatusNoNativeHyperchargeNormalizationQ), Detail: FormatQNorm(a.QNorm)},
			{Name: "audit n and q type distinction", Passed: strings.Contains(a.Types.NType, "S^2") && strings.Contains(a.Types.QType, "R^×") && a.Types.TypeDistinct && a.Types.ChangingNChangesLine && a.Types.ChangingQRescalesLine && !a.Types.NCanDetermineQ && !a.Types.QCanDetermineN && a.Types.IndependentAtLevel && strings.Contains(a.Types.Verdict, StatusNAndQTypeDistinctMissingSeals), Detail: FormatTypes(a.Types)},
			{Name: "audit forbidden shortcuts", Passed: len(a.Shortcuts.Shortcuts) == forbiddenShortcutCount && !a.Shortcuts.ScalarQuantitiesSelectN && !a.Shortcuts.EventProbabilityFixesQ && a.Shortcuts.AllShortcutsRejected && strings.Contains(a.Shortcuts.Verdict, StatusScalarBridgeDataDoNotSelectN) && strings.Contains(a.Shortcuts.Verdict, StatusK7EventProbabilityDoesNotFixQ), Detail: FormatShortcuts(a.Shortcuts)},
			{Name: "classify missing seals", Passed: strings.Contains(a.Seals.TwistorSelectorSeal, "TwistorSelectorSeal") && strings.Contains(a.Seals.HyperchargeNormalizationSeal, "HyperchargeNormalizationSeal") && len(a.Seals.Seals) == sealCount && a.Seals.ConditionalSocketRemains && !a.Seals.DerivedNative && strings.Contains(a.Seals.Verdict, StatusConditionalHiggsSocketReadyButNotNative), Detail: FormatSeals(a.Seals)},
			{Name: "enforce physical firewall", Passed: !a.Physical.ConditionalSocketPhysicalHiggsTheorem && !a.Physical.MatchedQDerivedHypercharge && !a.Physical.ChosenNDerivedVacuumOrientation && !a.Physical.K7MinusSelectorFlavorHierarchy && !a.Physical.K7PlusPhysicalHiggsMassTheorem && !a.Physical.ScalarPotential && !a.Physical.QuarticRuntimeLambda && !a.Physical.HiggsPoleMass && !a.Physical.YukawaOperators && !a.Physical.FlavorHierarchy && !a.Physical.CKMPMNS && strings.Contains(a.Physical.Verdict, StatusGate720MissingSealIndependenceBoundary), Detail: FormatPhysical(a.Physical)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: AuditID, Name: "Gate 720 — Higgs Socket Missing-Seal Independence and Source-Candidate Audit", Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
