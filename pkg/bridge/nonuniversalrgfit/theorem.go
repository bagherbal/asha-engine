package nonuniversalrgfit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func NonUniversalRationalLatticeRGFitAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-NONUNIVERSAL-RATIONAL-LATTICE-RG-FIT-AUDIT"
	const name = "non-universal rational lattice RG fit / sub-Planck asymptotic safety audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 210 audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "Gate 209 leptoquark dynamics seal is inherited", Passed: a.Gate209.Gate209Inherited && a.Gate209.LeptoquarkDynamicsSealActive && a.Gate209.SealedConnectionBaryonConservation && a.Gate209.ProtonLifetimeComputationObstructed, Detail: FormatGate209(a.Gate209)},
			{Name: "Gate 204 rational representation-row lattice is inherited", Passed: a.Lattice.Summary.LatticeConstructed && a.Lattice.LatticeAudit.UniqueGeneratorRows == a.GeneratorAudit.SourceUniqueRows && a.Lattice.LatticeAudit.SemigroupOnly, Detail: FormatGeneratorAudit(a.GeneratorAudit)},
			{Name: "RG inputs are quarantined and no universal row is allowed", Passed: a.Inputs.ObservedLedgerQuarantined && !a.Inputs.UsedForFiniteCoreDerivation && !a.Inputs.UniversalBetaRowAllowed && a.Inputs.SingleScaleThresholdAssumed, Detail: FormatInputs(a.Inputs)},
			{Name: "exact π-separation obstruction excludes rational single-scale closure", Passed: a.PiNoGo.LedgerCouplingsRationalDecimals && a.PiNoGo.BoundaryContainsExactPi && a.PiNoGo.DeterminantBOneANonZero && a.PiNoGo.ExactClosureRequiresDeltaOnSMBetaRay && a.PiNoGo.SMBetaRayHasNegativeComponents && a.PiNoGo.RationalLatticeNonnegativeSemigroup && a.PiNoGo.ExactClosureImpossible, Detail: FormatPiNoGo(a.PiNoGo)},
			{Name: "search basis is anomaly-safe and leptoquark-seal-compatible", Passed: a.GeneratorAudit.SafeGenerators > 0 && a.GeneratorAudit.SafeGenerators == a.GeneratorAudit.AnomalyCompatibleGenerators && a.GeneratorAudit.SafeGenerators == a.GeneratorAudit.LeptoquarkSealCompatibleRows && a.GeneratorAudit.NoUniversalBetaRowInserted && a.GeneratorAudit.NoContinuousRowCoefficients, Detail: FormatGeneratorAudit(a.GeneratorAudit)},
			{Name: "bounded Diophantine search finds no exact closure", Passed: a.Search.CombinationsAudited > 0 && a.Search.OrderedScaleCandidates > 0 && a.Search.ExactClosureCandidates == 0 && a.Search.ExactAsymptoticallySafeCandidates == 0 && a.Search.ExactNoGoProvidedByPiSeparation, Detail: FormatSearch(a.Search)},
			{Name: "asymptotic-safety filter has no exact candidate to accept", Passed: a.Safety.ExactCandidatesAudited == 0 && a.Safety.ExactCandidatesNoLandauPole == 0 && a.Safety.BestNearMissNoLandauPole && a.Safety.BestNearMissBoundaryBelowPlanck, Detail: FormatSafety(a.Safety)},
			{Name: "baryon/anomaly firewalls are preserved", Passed: a.BaryonAnomaly.LeptoquarkDynamicsSealInherited && a.BaryonAnomaly.AllSearchRowsAnomalyCompatible && a.BaryonAnomaly.AllSearchRowsSealCompatible && !a.BaryonAnomaly.ProtonDecayOperatorUsed && !a.BaryonAnomaly.ProtonLifetimeComputed, Detail: FormatBaryonAnomaly(a.BaryonAnomaly)},
			{Name: "no conditional prediction or physical fit is emitted", Passed: !a.Firewall.UniversalBetaRowInserted && !a.Firewall.ArbitraryRealRowCoefficientInserted && !a.Firewall.ExactClosureClaimed && !a.Firewall.ConditionalPredictionEmitted && !a.Firewall.ObservedLedgerUsedForFiniteCore && !a.Firewall.ProtonDecaySealViolated && !a.Firewall.AbsoluteMassPredicted && !a.Firewall.PhysicalUnificationClaimed && !a.Firewall.ThresholdCorrectedPhysicalFitClaimed && !a.Firewall.FiniteMatchingCorrectionsDerived, Detail: FormatFirewall(a.Firewall) + " :: " + FormatSummary(a.Summary)},
		}, Notes: []string{StatusFailedRouteExactSingleScale, StatusBoundedNearMissOnly, a.TruthStatement}}
	}}
}
