package scalarorientationsource

import "github.com/bagherbal/asha-engine/pkg/theorem"

func EtaOddScalarOrientationSourceMatterPullbackSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-ETA-ODD-SCALAR-ORIENTATION-SOURCE-MATTER-PULLBACK-SEARCH-AUDIT"
	const name = "eta-odd scalar-orientation source / matter-pullback search audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build eta-odd scalar orientation-source audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "Gate 189 compatibility is inherited but eta orientation remains unresolved", Passed: a.Summary.Gate189CompatibilityInherited && !a.PreviousGate189.Firewall.CanonicalEtaOrientationDerived && !a.PreviousGate189.Firewall.PhysicalScalarBundleDerived, Detail: scalarbundleDetail(a)},
			{Name: "weak isospin and scalar hypercharge do not provide a gauge-invariant eta-odd source", Passed: a.WeakGauge.T3LAvailable && a.WeakGauge.ScalarHyperchargeAvailable && a.WeakGauge.MatterSU2LAvailable && a.WeakGauge.T3CommutesWithHighLowProjectors && a.WeakGauge.YCommutesWithHighLowProjectors && a.WeakGauge.T1MixesHighLowPlanes && a.WeakGauge.T2MixesHighLowPlanes && a.WeakGauge.WeylReflectionExchangesPlanes && !a.WeakGauge.GaugeActionSelectsOrientation && !a.WeakGauge.EtaOddGaugeInvariantSource, Detail: FormatWeakGauge(a.WeakGauge)},
			{Name: "charge conjugation exchanges orientations rather than selecting eta", Passed: a.Conjugation.ContactChargeConjugationAvailable && a.Conjugation.ContactChargeConjugationInvolution && a.Conjugation.ContactChargeConjugationExchanges && !a.Conjugation.ContactChargeConjugationSelects && a.Conjugation.HiggsConjugateCollapseRejected && a.Conjugation.HiggsBranchUniquenessByKind && a.Conjugation.MirrorsEtaInvolution && !a.Conjugation.SelectsEtaOrientation, Detail: FormatConjugation(a.Conjugation)},
			{Name: "broken-sector diagnostics cannot retroactively select eta orientation", Passed: a.BrokenSector.CovariantDerivativeTemplate && !a.BrokenSector.VacuumOrientationChosen && a.BrokenSector.DimensionlessWZPhotonSignature && a.BrokenSector.GoldstoneImageDiagnostic && !a.BrokenSector.FiniteGaugeEatingDerived && !a.BrokenSector.BrokenMetricPhysicalPrediction && a.BrokenSector.GaugeNormalizationArtifact && !a.BrokenSector.BrokenSectorEtaOddForce && !a.BrokenSector.SelectsEtaOrientation, Detail: FormatBrokenSector(a.BrokenSector)},
			{Name: "complete source search finds no available gauge-invariant eta-odd selector", Passed: a.SourceSearch.AvailableCandidates >= 7 && a.SourceSearch.EtaOddCandidates == 0 && a.SourceSearch.GaugeInvariantEtaOddSources == 0 && a.SourceSearch.CandidatesSelectingEta == 0 && a.SourceSearch.ExchangeNotSelectorCount >= 2 && !a.SourceSearch.EtaOddSourceFound && !a.SourceSearch.GaugeInvariantSourceFound && !a.SourceSearch.CanonicalOrientationDerived, Detail: FormatSources(a.SourceSearch)},
			{Name: "eta orientation is classified as spontaneous/gauge data", Passed: a.Spontaneous.EtaInvolutionPreserved && a.Spontaneous.EtaInvolutionEquivalentToPlaneSwap && a.Spontaneous.GaugeSymmetryExplainsNonselection && a.Spontaneous.SpontaneousOrientationDataRequired && !a.Spontaneous.FiniteObservableCanSelect && a.Spontaneous.PhysicalScalarBundleStillUnfixed && a.Spontaneous.OrientationInsertionPointIsolated, Detail: FormatSpontaneous(a.Spontaneous)},
			{Name: "summary records positive obstruction and no physical scalar-bundle promotion", Passed: a.Summary.TestsAudited == 6 && a.Summary.WeakHyperchargeAudited && a.Summary.ChargeConjugationAudited && a.Summary.BrokenSectorAudited && a.Summary.ContactSignedSourcesAudited && !a.Summary.EtaOddSourceFound && !a.Summary.GaugeInvariantEtaOddSourceFound && !a.Summary.CanonicalEtaOrientationDerived && a.Summary.EtaOrientationClassifiedSpontaneous && !a.Summary.PhysicalScalarBundleDerived, Detail: FormatSummary(a.Summary)},
			{Name: "firewall keeps physical bundle, Chern-Weil, heat-kernel, thresholds, and constants sealed", Passed: !a.Firewall.UsesObservedInputForDerivation && !a.Firewall.UsesNumericRootApproximation && !a.Firewall.UsesIndividualRootDiagonalization && !a.Firewall.UsesArbitraryEtaHighLowAssignment && a.Firewall.Gate189CompatibilityInherited && a.Firewall.WeakHyperchargeSourceAudited && a.Firewall.ChargeConjugationSourceAudited && a.Firewall.BrokenSectorSourceAudited && a.Firewall.ContactSignedSourceAudited && !a.Firewall.EtaOddFiniteSourceFound && !a.Firewall.GaugeInvariantEtaOddSourceFound && !a.Firewall.CanonicalEtaOrientationDerived && a.Firewall.EtaOrientationClassifiedSpontaneous && !a.Firewall.PhysicalScalarBundleDerived && !a.Firewall.ChernWeilCarrierDerived && !a.Firewall.HeatKernelMatchingDerived && !a.Firewall.ThresholdCorrectedBetaDerived && !a.Firewall.AbsoluteCouplingPromoted && !a.Firewall.PhysicalConstantsDerived && a.Firewall.StrictNullityBefore == 3 && a.Firewall.StrictNullityAfter == 3 && a.Firewall.ConditionalNullityBefore == 1 && a.Firewall.ConditionalNullityAfter == 0, Detail: FormatFirewall(a.Firewall)},
		}, Notes: []string{
			a.TruthStatement,
			"Gate 190 is labelled FAILED_ROUTE only as a constructive eta-source search; as an obstruction theorem it succeeds and localizes the spontaneous orientation insertion point.",
			"The next lawful move is a controlled spontaneous-orientation seal, not another hidden selector search or a direct constants promotion.",
		}}
	}}
}

func scalarbundleDetail(a Analysis) string {
	return "Gate 189 dimension compatibility and conditional maps inherited; eta-to-high/low and physical scalar bundle remain unselected before Gate 190 audit"
}
