package thresholdorigindichotomy

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ThresholdOriginDichotomyNewSectorContinuumBridgeAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-THRESHOLD-ORIGIN-DICHOTOMY-NEW-SECTOR-CONTINUUM-BRIDGE-AUDIT"
	const name = "threshold-origin dichotomy / new-sector versus continuum-decoupling bridge audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{{Name: "build threshold-origin dichotomy", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 178 threshold no-go is consumed as input", Passed: a.Dichotomy.PreviousGate178NoGo && a.PreviousGate178.Requirements.NoCandidateHasAllPieces && !a.PreviousGate178.Firewall.ThresholdOperatorDerived && !a.PreviousGate178.Firewall.NonUniversalDeltaBDerived, Detail: finitethresholdoperatorDetail(a)},
			{Name: "origin branches separate future programs from rejected shortcuts", Passed: len(a.Branches) == 4 && a.Dichotomy.ContinuumBranches == 1 && a.Dichotomy.NewSectorBranches == 1 && a.Dichotomy.FitBranches == 1 && a.Dichotomy.SchemeBranches == 1 && len(a.Dichotomy.SurvivingProgrammaticOrigins) == 2, Detail: FormatBranches(a.Branches)},
			{Name: "existing finite spectra require a continuum decoupling bridge", Passed: len(a.Continuum.CandidateExistingAnchors) >= 4 && a.Continuum.OrientedFourCycleRequired && a.Continuum.PrincipalBundleRequired && a.Continuum.ChernWeilNormalizationRequired && a.Continuum.ContinuumTraceConventionRequired && a.Continuum.LocalFieldMapRequired && a.Continuum.PhysicalMassUnitRequired && a.Continuum.ActivationPredicateRequired && a.Continuum.DecouplingLawRequired && a.Continuum.GaugeRepresentationRowsRequired && !a.Continuum.AllRequiredObjectsPresent && !a.Continuum.BridgeDerived && !a.Continuum.CanPromoteGate177Repair, Detail: FormatContinuum(a.Continuum)},
			{Name: "new-sector origin is open but not currently derived", Passed: a.NewSector.KnownFiniteInventoryExhausted && len(a.NewSector.RequiredFeatures) >= 6 && a.NewSector.DerivedNewSectors == 0 && a.NewSector.RepresentationCompleteHeavyMultiplets == 0 && a.NewSector.CanonicalMassSpectrumCount == 0 && !a.NewSector.CanGenerateNonUniversalDeltaB && !a.NewSector.CanPromoteGate177Repair, Detail: FormatNewSector(a.NewSector)},
			{Name: "phenomenological and scheme routes are rejected as threshold origins", Passed: len(a.Rejected) >= 3 && allRejectedOriginsSafe(a.Rejected) && a.Dichotomy.ObservedFitRejectedAsOrigin && a.Dichotomy.SchemeOnlyRejectedAsThresholdOrigin, Detail: FormatRejected(a.Rejected)},
			{Name: "dichotomy is complete at the current finite-data stage", Passed: a.Dichotomy.DichotomyCompleteAtCurrentStage && !a.Dichotomy.ThresholdOriginDerived && !a.Dichotomy.Gate177RepairPromoted, Detail: FormatDichotomy(a.Dichotomy)},
			{Name: "threshold and nullity firewall remain closed", Passed: !a.Firewall.UsesObservedInputForDerivation && !a.Firewall.NonUniversalDeltaBDerived && !a.Firewall.ThresholdOperatorDerived && !a.Firewall.ThresholdCorrectedBetaDerived && !a.Firewall.PhysicalConstantsDerived && !a.Firewall.BoundaryScaleDerived && a.Firewall.StrictNullityBefore == 3 && a.Firewall.StrictNullityAfter == 3 && a.Firewall.ConditionalNullityBefore == 2 && a.Firewall.ConditionalNullityAfter == 2, Detail: FormatFirewall(a.Firewall)},
		}, Notes: []string{
			a.TruthStatement,
			"Gate 179 does not derive thresholds; it classifies their lawful origin after all current finite candidates fail.",
			"The next bridge must either construct a continuum decoupling/heat-kernel matching map for existing finite spectra, or derive new finite heavy sectors with representation-complete beta rows.",
		}}
	}}
}

func allRejectedOriginsSafe(xs []RejectedOrigin) bool {
	for _, x := range xs {
		if x.CanServeAsTheorem || !x.PreservesFirewall {
			return false
		}
	}
	return true
}

func finitethresholdoperatorDetail(a Analysis) string {
	return "Gate178: " + a.PreviousGate178.Firewall.Verdict
}
