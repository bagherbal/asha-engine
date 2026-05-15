// Package generation2finiteactionsecondvariation implements Gate 495:
// Finite Electroweak Action Second Variation Source Audit.
//
// Gate 494 proved that hypercharge trace normalization, anomaly cancellation,
// finite count resonances, and representation trace metrics do not by
// themselves select the abelian electroweak completion coefficient kappa_U1.
// Gate 495 therefore inspects the only remaining internal candidate: the older
// canonical finite scalar/gauge variational action package, whose broken-orbit
// second variation reproduces diag(1,1,4) and hence kappa_U1=6.
//
// The result is deliberately two-layered. The canonical action candidate is a
// coherent dimensionless variational bridge: it computes the rank-three broken
// orbit, the normalized diag(1,1,4) Hessian, and a positive rank-four full
// electroweak Hessian.  But the Generation-2 firewall refuses to promote this
// into a native theorem until the inputs of that action are themselves derived:
// the scalar covariant derivative, scalar kinetic metric I4, vacuum orientation,
// and full scalar SU(2) action are still diagnostic/bridge objects in Gates
// 492-494.  Thus Gate 495 records a strong second-variation candidate while
// blocking the native kappa/gauge-Hessian/physical-mass registry write.
package generation2finiteactionsecondvariation

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/canonicalaction"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2abeliancoefficientselection"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2electroweakcurvatureaction"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2scalarcovariantintertwiner"
)

const (
	AuditID = "GATE495-FINITE-ELECTROWEAK-ACTION-SECOND-VARIATION-SOURCE-AUDIT"

	StatusGate494Inherited                   = "CONDITIONAL_SUPPORT_GATE494_KAPPA_OBSTRUCTION_INHERITED"
	StatusCanonicalActionCandidateFound      = "CONDITIONAL_SUPPORT_LEGACY_CANONICAL_ACTION_SECOND_VARIATION_CANDIDATE_FOUND"
	StatusBrokenOrbitDiag114Reproduced       = "CONDITIONAL_SUPPORT_BROKEN_ORBIT_SECOND_VARIATION_DIAG114_REPRODUCED"
	StatusKappaSixSelectedInsideCandidate    = "CONDITIONAL_SUPPORT_KAPPA_U1_SIX_SELECTED_INSIDE_CANONICAL_ACTION_CANDIDATE"
	StatusFullHessianCandidatePositive       = "CONDITIONAL_SUPPORT_FULL_ELECTROWEAK_HESSIAN_CANDIDATE_POSITIVE_RANK_FOUR"
	StatusDimensionlessHessianBridgeAccepted = "CONDITIONAL_SUPPORT_DIMENSIONLESS_ELECTROWEAK_HESSIAN_BRIDGE_CANDIDATE_ACCEPTED"
	StatusProvenanceAirlockDefined           = "CONDITIONAL_SUPPORT_SECOND_VARIATION_PROVENANCE_AIRLOCK_DEFINED"
	StatusFirewallPreserved                  = "FIREWALL_PRESERVED_NO_ELECTROWEAK_OR_FLAVOR_DATA_IMPORTED"
	StatusNativeRegistryWriteBlocked         = "FIREWALL_BLOCKED_NATIVE_KAPPA_AND_WZ_REGISTRY_WRITE"

	StatusFailedCanonicalActionNotNativeClosed = "FAILED_ROUTE_CANONICAL_ACTION_PROVENANCE_NOT_NATIVE_CLOSED"
	StatusFailedNativeDphiStillAbstract        = "FAILED_ROUTE_NATIVE_DPHI_STILL_ABSTRACT_TEMPLATE"
	StatusFailedScalarKineticMetricNotNative   = "FAILED_ROUTE_SCALAR_KINETIC_METRIC_I4_NOT_NATIVE_DERIVED"
	StatusFailedVacuumOrientationNotNative     = "FAILED_ROUTE_SCALAR_VACUUM_ORIENTATION_NOT_NATIVE"
	StatusFailedFullScalarSU2ActionNotSelected = "FAILED_ROUTE_FULL_SCALAR_SU2_ACTION_NOT_NATIVE_SELECTED"
	StatusFailedGaugeHessianNativeNotSelected  = "FAILED_ROUTE_GAUGE_HESSIAN_NATIVE_SELECTION_NOT_CLOSED"
	StatusFailedPhysicalCouplingsNotDerived    = "FAILED_ROUTE_PHYSICAL_GAUGE_COUPLINGS_NOT_DERIVED"
	StatusFailedWeakAngleNotDerived            = "FAILED_ROUTE_WEAK_MIXING_ANGLE_NOT_DERIVED"
	StatusFailedWZMassesNotDerived             = "FAILED_ROUTE_PHYSICAL_WZ_MASSES_NOT_DERIVED"
	StatusGate496RedirectDefined               = "CONDITIONAL_SUPPORT_GATE496_SCALAR_KINETIC_PROVENANCE_REDIRECT_DEFINED"
)

const eps = 1e-8

type Inheritance struct {
	Executed                             bool
	Gate494AuditDefined                  bool
	TraceNormalizationDoesNotSelectKappa bool
	RepresentationMetricNotGaugeHessian  bool
	KappaTargetSixWhiteningCandidate     bool
	FiniteActionSecondVariationRequested bool
	NoElectroweakFlavorDataImported      bool
	Verdict                              string
	Reason                               string
}

type CanonicalCandidateAudit struct {
	Executed                            bool
	CanonicalActionPackageAvailable     bool
	CanonicalActionSelectedInCandidate  bool
	SecondVariationComputed             bool
	ScalarKineticSelectedInCandidate    bool
	ActiveRealDimension                 int
	BrokenImageRank                     int
	BrokenRawDiagonal                   []float64
	BrokenSelectedDiagonal              []float64
	BrokenDiag114                       bool
	KappaU1                             float64
	KappaSixSelectedInCandidate         bool
	FullGaugeHessianSelectedInCandidate bool
	FullGaugeHessianRank                int
	FullGaugeHessianPositive            bool
	BrokenRestrictionMatches            bool
	PhysicalCouplingsDerived            bool
	PhysicalMassesDerived               bool
	CKMPMNSDerived                      bool
	HiddenObservedInputUsed             bool
	Verdict                             string
	Reason                              string
}

type ProvenanceAudit struct {
	Executed                                  bool
	Gate492NativeDphiDerived                  bool
	Gate492CanonicalGoldstoneIntertwiner      bool
	Gate492FullScalarSU2ActionNativeSelected  bool
	Gate492ScalarVacuumOrientationNative      bool
	Gate492ScalarKineticMetricNative          bool
	Gate493SecondVariationComputed            bool
	Gate493AbelianCoefficientSelected         bool
	Gate493GaugeHessianActionSelected         bool
	CanonicalCandidateUsesDiagnosticDphi      bool
	CanonicalCandidateUsesI4Metric            bool
	CanonicalCandidateUsesChosenVacuum        bool
	CanonicalCandidateUsesMinimalActionChoice bool
	NativeActionProvenanceClosed              bool
	NativeKappaSelectionClosed                bool
	NativeGaugeHessianSelectionClosed         bool
	Verdict                                   string
	Reason                                    string
}

type SelectionBoundary struct {
	Executed                              bool
	DimensionlessSecondVariationCandidate bool
	Diag114AcceptedAsBridgeCandidate      bool
	KappaSixAcceptedAsBridgeCandidate     bool
	FullHessianAcceptedAsBridgeCandidate  bool
	NativeKappaSelected                   bool
	NativeGaugeHessianSelected            bool
	NativeWeakAngleDerived                bool
	NativeGaugeCouplingsDerived           bool
	NativeWZMassesDerived                 bool
	NativeHiggsVEVDerived                 bool
	PhysicalElectroweakRegistryWrite      bool
	Verdict                               string
	Reason                                string
}

type Firewall struct {
	Executed                   bool
	ObservedWMassImported      bool
	ObservedZMassImported      bool
	ObservedHiggsMassImported  bool
	FermiConstantImported      bool
	WeakAngleImported          bool
	FineStructureImported      bool
	GaugeCouplingImported      bool
	YukawaImported             bool
	CKMPMNSImported            bool
	NativeKappaWritten         bool
	NativeGaugeHessianWritten  bool
	NativeWeakAngleWritten     bool
	NativeGaugeCouplingWritten bool
	NativeWZMassWritten        bool
	NativeHiggsVEVWritten      bool
	Verdict                    string
	Reason                     string
}

type RegistryUpdate struct {
	NativeEntries        []string
	BridgeEntries        []string
	EnvironmentalEntries []string
	FailedRoutes         []string
	OpenTheorems         []string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Candidate   CanonicalCandidateAudit
	Provenance  ProvenanceAudit
	Boundary    SelectionBoundary
	Firewall    Firewall
	Registry    RegistryUpdate
	Next        NextStep
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g494, err := generation2abeliancoefficientselection.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate494 abelian coefficient audit: %w", err)
	}
	g492, err := generation2scalarcovariantintertwiner.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate492 scalar covariant audit: %w", err)
	}
	g493, err := generation2electroweakcurvatureaction.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate493 electroweak curvature audit: %w", err)
	}
	ca, err := canonicalaction.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not build canonical action second-variation candidate: %w", err)
	}

	a := Analysis{}
	a.Inheritance = buildInheritance(g494)
	a.Candidate = buildCanonicalCandidate(ca)
	a.Provenance = buildProvenance(g492, g493, ca)
	a.Boundary = buildBoundary(a.Candidate, a.Provenance)
	a.Firewall = buildFirewall()
	a.Registry = buildRegistryUpdate(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g generation2abeliancoefficientselection.Analysis) Inheritance {
	return Inheritance{
		Executed:                             true,
		Gate494AuditDefined:                  true,
		TraceNormalizationDoesNotSelectKappa: !g.Boundary.TraceToKappaNativeMapDerived && !g.Boundary.SecondVariationSelectsKappa && !g.Boundary.NativeKappaSelected,
		RepresentationMetricNotGaugeHessian:  g.Metric.DiagonalTraceGramAsRepresentationMetric && !g.Metric.DiagonalTraceGramAsGaugeKineticHessian,
		KappaTargetSixWhiteningCandidate:     g.Kappa.TargetKappa == 6 && g.Kappa.WhiteningSelectsKappa && !g.Kappa.ActionSelectsKappa,
		FiniteActionSecondVariationRequested: g.Next.Gate == 495,
		NoElectroweakFlavorDataImported:      g.Firewall.Executed && !g.Firewall.ObservedWMassImported && !g.Firewall.ObservedZMassImported && !g.Firewall.WeakAngleImported && !g.Firewall.YukawaImported && !g.Firewall.CKMPMNSImported,
		Verdict:                              StatusGate494Inherited,
		Reason:                               "Gate494 leaves one admissible target: inspect an explicit finite action second variation rather than trace normalization, count resonance, or representation metrics.",
	}
}

func buildCanonicalCandidate(c canonicalaction.Analysis) CanonicalCandidateAudit {
	return CanonicalCandidateAudit{
		Executed:                            true,
		CanonicalActionPackageAvailable:     true,
		CanonicalActionSelectedInCandidate:  c.CanonicalActionSelected,
		SecondVariationComputed:             c.SecondVariationComputed,
		ScalarKineticSelectedInCandidate:    c.ScalarKineticSelected,
		ActiveRealDimension:                 c.ActiveRealDimension,
		BrokenImageRank:                     c.GaugeEating.BrokenImageRank,
		BrokenRawDiagonal:                   append([]float64(nil), c.BrokenRawDiagonal...),
		BrokenSelectedDiagonal:              append([]float64(nil), c.BrokenSelectedDiagonal...),
		BrokenDiag114:                       closeSlice(c.BrokenSelectedDiagonal, []float64{1, 1, 4}, 1e-8),
		KappaU1:                             c.KappaU1,
		KappaSixSelectedInCandidate:         c.U1CompletionCoefficientSelected && math.Abs(c.KappaU1-6) < eps,
		FullGaugeHessianSelectedInCandidate: c.FullGaugeHessianSelected,
		FullGaugeHessianRank:                c.FullGaugeHessianRank,
		FullGaugeHessianPositive:            c.FullGaugeHessianPositive,
		BrokenRestrictionMatches:            c.BrokenRestrictionMatches,
		PhysicalCouplingsDerived:            c.PhysicalCouplingsDerived,
		PhysicalMassesDerived:               c.PhysicalMassesDerived,
		CKMPMNSDerived:                      c.CKMPMNSDerived,
		HiddenObservedInputUsed:             c.HiddenObservedInputUsed,
		Verdict:                             StatusCanonicalActionCandidateFound,
		Reason:                              "The legacy canonical-action package computes a coherent dimensionless second-variation candidate: I4 scalar kinetic metric, rank-three broken orbit, normalized diag(1,1,4), full positive rank-four Hessian, and kappa_U1=6 inside that candidate.",
	}
}

func buildProvenance(g492 generation2scalarcovariantintertwiner.Analysis, g493 generation2electroweakcurvatureaction.Analysis, c canonicalaction.Analysis) ProvenanceAudit {
	usesDiagnosticDphi := g492.Dphi.AbstractTemplateAvailable && !g492.Boundary.NativeDphiDerived
	usesI4 := c.ScalarKineticSelected && !g492.Boundary.NativeKineticMetricDerived
	usesChosenVacuum := len(c.GaugeEating.ScalarCovariant.VacuumVector) == 4 && !g492.Boundary.NativeVacuumOrientationDerived
	minimalChoice := strings.Contains(strings.ToLower(c.ActionFormula), "s_can") || strings.Contains(strings.ToLower(c.ActionName), "canonical")
	nativeClosed := g492.Boundary.NativeDphiDerived && g492.Boundary.NativeKineticMetricDerived && g492.Boundary.NativeVacuumOrientationDerived && g492.Representation.FullSU2SelectedByScalarData && g493.Boundary.SecondVariationComputed && g493.Boundary.AbelianCoefficientSelected && g493.Boundary.GaugeHessianActionSelected
	return ProvenanceAudit{
		Executed:                                  true,
		Gate492NativeDphiDerived:                  g492.Boundary.NativeDphiDerived,
		Gate492CanonicalGoldstoneIntertwiner:      g492.Intertwiner.CanonicalProtectedToBrokenDerived,
		Gate492FullScalarSU2ActionNativeSelected:  g492.Representation.FullSU2SelectedByScalarData,
		Gate492ScalarVacuumOrientationNative:      g492.Boundary.NativeVacuumOrientationDerived,
		Gate492ScalarKineticMetricNative:          g492.Boundary.NativeKineticMetricDerived,
		Gate493SecondVariationComputed:            g493.Boundary.SecondVariationComputed,
		Gate493AbelianCoefficientSelected:         g493.Boundary.AbelianCoefficientSelected,
		Gate493GaugeHessianActionSelected:         g493.Boundary.GaugeHessianActionSelected,
		CanonicalCandidateUsesDiagnosticDphi:      usesDiagnosticDphi,
		CanonicalCandidateUsesI4Metric:            usesI4,
		CanonicalCandidateUsesChosenVacuum:        usesChosenVacuum,
		CanonicalCandidateUsesMinimalActionChoice: minimalChoice,
		NativeActionProvenanceClosed:              nativeClosed,
		NativeKappaSelectionClosed:                nativeClosed && c.U1CompletionCoefficientSelected,
		NativeGaugeHessianSelectionClosed:         nativeClosed && c.FullGaugeHessianSelected,
		Verdict:                                   StatusProvenanceAirlockDefined,
		Reason:                                    "The candidate computes the desired Hessian, but its ingredients still pass through diagnostic DΦ, I4 kinetic metric, chosen vacuum orientation, and a minimal canonical-action choice; those inputs are not yet native Generation-2 theorems.",
	}
}

func buildBoundary(c CanonicalCandidateAudit, p ProvenanceAudit) SelectionBoundary {
	bridgeOK := c.SecondVariationComputed && c.BrokenDiag114 && c.KappaSixSelectedInCandidate && c.FullGaugeHessianPositive && c.FullGaugeHessianRank == 4 && !c.PhysicalCouplingsDerived && !c.PhysicalMassesDerived
	return SelectionBoundary{
		Executed:                              true,
		DimensionlessSecondVariationCandidate: bridgeOK,
		Diag114AcceptedAsBridgeCandidate:      bridgeOK,
		KappaSixAcceptedAsBridgeCandidate:     bridgeOK,
		FullHessianAcceptedAsBridgeCandidate:  bridgeOK,
		NativeKappaSelected:                   p.NativeKappaSelectionClosed,
		NativeGaugeHessianSelected:            p.NativeGaugeHessianSelectionClosed,
		NativeWeakAngleDerived:                false,
		NativeGaugeCouplingsDerived:           false,
		NativeWZMassesDerived:                 false,
		NativeHiggsVEVDerived:                 false,
		PhysicalElectroweakRegistryWrite:      false,
		Verdict:                               StatusDimensionlessHessianBridgeAccepted,
		Reason:                                "Gate495 accepts the canonical second-variation object as a dimensionless bridge candidate, but blocks native promotion until the DΦ/metric/vacuum/action provenance chain closes.",
	}
}

func buildFirewall() Firewall {
	return Firewall{
		Executed:                   true,
		ObservedWMassImported:      false,
		ObservedZMassImported:      false,
		ObservedHiggsMassImported:  false,
		FermiConstantImported:      false,
		WeakAngleImported:          false,
		FineStructureImported:      false,
		GaugeCouplingImported:      false,
		YukawaImported:             false,
		CKMPMNSImported:            false,
		NativeKappaWritten:         false,
		NativeGaugeHessianWritten:  false,
		NativeWeakAngleWritten:     false,
		NativeGaugeCouplingWritten: false,
		NativeWZMassWritten:        false,
		NativeHiggsVEVWritten:      false,
		Verdict:                    StatusFirewallPreserved,
		Reason:                     "No observed electroweak, Higgs, gauge-coupling, Yukawa, CKM, or PMNS datum is imported; the candidate Hessian does not become a physical mass or coupling theorem.",
	}
}

func buildRegistryUpdate(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"No new physical electroweak mass, coupling, weak-angle, Higgs-scale, Yukawa, CKM, or PMNS native entry is written by Gate495.",
		},
		BridgeEntries: []string{
			"The legacy canonical-action package supplies a coherent dimensionless second-variation candidate with broken Hessian diag(1,1,4).",
			"Inside that candidate, matching the broken scalar-orbit Hessian into the closed {T1,T2,Z,Q} carrier gives kappa_U1=6.",
			"The full electroweak Hessian candidate is positive and rank four, but remains bridge-level pending provenance closure.",
		},
		EnvironmentalEntries: []string{
			"Observed W/Z masses, Higgs VEV, Fermi constant, theta_W, alpha, running gauge couplings, Yukawa matrices, CKM, and PMNS remain sealed.",
		},
		FailedRoutes: []string{
			StatusFailedCanonicalActionNotNativeClosed,
			StatusFailedNativeDphiStillAbstract,
			StatusFailedScalarKineticMetricNotNative,
			StatusFailedVacuumOrientationNotNative,
			StatusFailedFullScalarSU2ActionNotSelected,
			StatusFailedGaugeHessianNativeNotSelected,
			StatusFailedPhysicalCouplingsNotDerived,
			StatusFailedWeakAngleNotDerived,
			StatusFailedWZMassesNotDerived,
		},
		OpenTheorems: []string{
			"Derive the scalar kinetic metric I4 from a finite Hilbert-Schmidt/spectral trace theorem rather than selecting it as the active-frame Euclidean metric.",
			"Derive the scalar vacuum orientation and scalar SU(2)_L action from finite contact/spectral data rather than using a unitary-gauge diagnostic orientation.",
			"Only after DΦ, metric, and vacuum provenance close may kappa_U1=6 become a native action-selected Hessian rather than a strong dimensionless bridge candidate.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        496,
		Title:       "Scalar Kinetic Metric Provenance and Vacuum Orientation Closure Audit",
		Reason:      "Gate495 finds the right second-variation candidate but not the native provenance of the metric/vacuum/DΦ ingredients.",
		PrimaryTask: "prove or reject that the Hilbert-Schmidt finite scalar trace, scalar SU(2)_L action, and vacuum selector natively force the I4 metric and vacuum orientation used by the canonical second-variation candidate, without importing W/Z masses, theta_W, v, alpha, or Yukawa data",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.FiniteActionSecondVariationRequested {
		return fmt.Errorf("Gate495 did not inherit the Gate494 second-variation request: %+v", a.Inheritance)
	}
	if !a.Candidate.Executed || !a.Candidate.SecondVariationComputed || !a.Candidate.BrokenDiag114 || !a.Candidate.KappaSixSelectedInCandidate || !a.Candidate.FullGaugeHessianPositive || a.Candidate.FullGaugeHessianRank != 4 {
		return fmt.Errorf("canonical action candidate did not reproduce the required Hessian data: %+v", a.Candidate)
	}
	if a.Candidate.PhysicalCouplingsDerived || a.Candidate.PhysicalMassesDerived || a.Candidate.CKMPMNSDerived || a.Candidate.HiddenObservedInputUsed {
		return fmt.Errorf("canonical action candidate leaked physical data: %+v", a.Candidate)
	}
	if !a.Provenance.Executed || a.Provenance.NativeActionProvenanceClosed || a.Provenance.NativeKappaSelectionClosed || a.Provenance.NativeGaugeHessianSelectionClosed {
		return fmt.Errorf("provenance airlock over-promoted the candidate: %+v", a.Provenance)
	}
	if !a.Boundary.DimensionlessSecondVariationCandidate || a.Boundary.NativeKappaSelected || a.Boundary.NativeGaugeHessianSelected || a.Boundary.NativeWZMassesDerived || a.Boundary.PhysicalElectroweakRegistryWrite {
		return fmt.Errorf("selection boundary violated: %+v", a.Boundary)
	}
	if a.Firewall.ObservedWMassImported || a.Firewall.WeakAngleImported || a.Firewall.GaugeCouplingImported || a.Firewall.NativeKappaWritten || a.Firewall.NativeWZMassWritten {
		return fmt.Errorf("firewall leak: %+v", a.Firewall)
	}
	return nil
}

func truth(a Analysis) string {
	return "Gate495 finds the strongest electroweak action clue so far: the legacy canonical finite action computes a dimensionless second variation whose broken slice is diag(1,1,4), whose full {T1,T2,Z,Q} Hessian is positive rank four, and whose internal matching gives kappa_U1=6. But this is not yet a native physical electroweak theorem, because the scalar covariant derivative, I4 scalar kinetic metric, vacuum orientation, and scalar SU(2)_L action used by the candidate are still bridge/provenance inputs in the Generation-2 audit chain. Therefore kappa_U1=6 is accepted as a strong dimensionless bridge candidate, while native gauge Hessian, weak angle, gauge couplings, W/Z masses, Higgs VEV, and all flavor data remain blocked."
}

func closeSlice(a, b []float64, eps float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(a[i]-b[i]) > eps {
			return false
		}
	}
	return true
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate494=%v trace_not_kappa=%v metric_not_hessian=%v kappa6_candidate=%v second_variation_requested=%v no_data=%v verdict=%s reason=%s", x.Gate494AuditDefined, x.TraceNormalizationDoesNotSelectKappa, x.RepresentationMetricNotGaugeHessian, x.KappaTargetSixWhiteningCandidate, x.FiniteActionSecondVariationRequested, x.NoElectroweakFlavorDataImported, x.Verdict, x.Reason)
}

func FormatCandidate(x CanonicalCandidateAudit) string {
	return fmt.Sprintf("package_available=%v canonical_selected=%v second_variation=%v scalar_I4_candidate=%v active_dim=%d broken_rank=%d raw_diag=%s selected_diag=%s diag114=%v kappa=%.10f kappa6_candidate=%v full_hessian_candidate=%v full_rank=%d full_positive=%v restriction_matches=%v physical_couplings=%v physical_masses=%v ckm_pmns=%v hidden_observed=%v verdict=%s reason=%s", x.CanonicalActionPackageAvailable, x.CanonicalActionSelectedInCandidate, x.SecondVariationComputed, x.ScalarKineticSelectedInCandidate, x.ActiveRealDimension, x.BrokenImageRank, formatFloatSlice(x.BrokenRawDiagonal), formatFloatSlice(x.BrokenSelectedDiagonal), x.BrokenDiag114, x.KappaU1, x.KappaSixSelectedInCandidate, x.FullGaugeHessianSelectedInCandidate, x.FullGaugeHessianRank, x.FullGaugeHessianPositive, x.BrokenRestrictionMatches, x.PhysicalCouplingsDerived, x.PhysicalMassesDerived, x.CKMPMNSDerived, x.HiddenObservedInputUsed, x.Verdict, x.Reason)
}

func FormatProvenance(x ProvenanceAudit) string {
	return fmt.Sprintf("native_Dphi=%v canonical_intertwiner=%v scalar_SU2_native=%v vacuum_native=%v kinetic_metric_native=%v gate493_second_variation=%v gate493_kappa=%v gate493_hessian=%v uses_diagnostic_Dphi=%v uses_I4=%v uses_chosen_vacuum=%v uses_minimal_action_choice=%v provenance_closed=%v native_kappa=%v native_hessian=%v verdict=%s reason=%s", x.Gate492NativeDphiDerived, x.Gate492CanonicalGoldstoneIntertwiner, x.Gate492FullScalarSU2ActionNativeSelected, x.Gate492ScalarVacuumOrientationNative, x.Gate492ScalarKineticMetricNative, x.Gate493SecondVariationComputed, x.Gate493AbelianCoefficientSelected, x.Gate493GaugeHessianActionSelected, x.CanonicalCandidateUsesDiagnosticDphi, x.CanonicalCandidateUsesI4Metric, x.CanonicalCandidateUsesChosenVacuum, x.CanonicalCandidateUsesMinimalActionChoice, x.NativeActionProvenanceClosed, x.NativeKappaSelectionClosed, x.NativeGaugeHessianSelectionClosed, x.Verdict, x.Reason)
}

func FormatBoundary(x SelectionBoundary) string {
	return fmt.Sprintf("dimensionless_candidate=%v diag114_bridge=%v kappa6_bridge=%v full_hessian_bridge=%v native_kappa=%v native_hessian=%v native_theta=%v native_couplings=%v native_WZ=%v native_vev=%v physical_registry_write=%v verdict=%s reason=%s", x.DimensionlessSecondVariationCandidate, x.Diag114AcceptedAsBridgeCandidate, x.KappaSixAcceptedAsBridgeCandidate, x.FullHessianAcceptedAsBridgeCandidate, x.NativeKappaSelected, x.NativeGaugeHessianSelected, x.NativeWeakAngleDerived, x.NativeGaugeCouplingsDerived, x.NativeWZMassesDerived, x.NativeHiggsVEVDerived, x.PhysicalElectroweakRegistryWrite, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("observed_W=%v observed_Z=%v observed_Higgs=%v Fermi=%v theta=%v alpha=%v gauge_coupling=%v Yukawa=%v CKM_PMNS=%v native_kappa=%v native_hessian=%v native_theta=%v native_gauge=%v native_WZ=%v native_vev=%v verdict=%s reason=%s", x.ObservedWMassImported, x.ObservedZMassImported, x.ObservedHiggsMassImported, x.FermiConstantImported, x.WeakAngleImported, x.FineStructureImported, x.GaugeCouplingImported, x.YukawaImported, x.CKMPMNSImported, x.NativeKappaWritten, x.NativeGaugeHessianWritten, x.NativeWeakAngleWritten, x.NativeGaugeCouplingWritten, x.NativeWZMassWritten, x.NativeHiggsVEVWritten, x.Verdict, x.Reason)
}

func formatFloatSlice(xs []float64) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		if math.Abs(x) < 1e-12 {
			x = 0
		}
		parts = append(parts, fmt.Sprintf("%.10f", x))
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 495 Registry Audit — Finite Electroweak Action Second Variation Source Audit\n\n")
	b.WriteString("## Verdict\n\n")
	for _, v := range []string{
		StatusGate494Inherited,
		StatusCanonicalActionCandidateFound,
		StatusBrokenOrbitDiag114Reproduced,
		StatusKappaSixSelectedInsideCandidate,
		StatusFullHessianCandidatePositive,
		StatusDimensionlessHessianBridgeAccepted,
		StatusProvenanceAirlockDefined,
		StatusFailedCanonicalActionNotNativeClosed,
		StatusFailedNativeDphiStillAbstract,
		StatusFailedScalarKineticMetricNotNative,
		StatusFailedVacuumOrientationNotNative,
		StatusFailedFullScalarSU2ActionNotSelected,
		StatusFailedGaugeHessianNativeNotSelected,
		StatusFailedPhysicalCouplingsNotDerived,
		StatusFailedWeakAngleNotDerived,
		StatusFailedWZMassesNotDerived,
		StatusFirewallPreserved,
		StatusNativeRegistryWriteBlocked,
		StatusGate496RedirectDefined,
	} {
		b.WriteString("- `" + v + "`\n")
	}

	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("Gate494 proves that `k_Y = 5/3`, anomaly/unimodularity consistency, finite count resonances, and diagonal representation metrics do not select `kappa_U1`. Gate495 may therefore inspect only an explicit action second variation, and it may not import electroweak or flavor data.\n\n")

	b.WriteString("## Canonical second-variation candidate\n\n")
	b.WriteString(FormatCandidate(a.Candidate) + "\n\n")
	b.WriteString("The legacy canonical-action package produces the desired dimensionless shape: the broken scalar-orbit second variation normalizes to `diag(1,1,4)`, and matching it into the closed `{T1,T2,Z,Q}` carrier gives `kappa_U1 = 6`. The full Hessian candidate is positive and rank four.\n\n")

	b.WriteString("## Provenance airlock\n\n")
	b.WriteString(FormatProvenance(a.Provenance) + "\n\n")
	b.WriteString("The candidate is not yet promoted to a native Generation-2 theorem because the finite `DΦ`, scalar kinetic metric `I4`, scalar vacuum orientation, and full scalar `SU(2)_L` action are still diagnostic inputs rather than independently derived native objects.\n\n")

	b.WriteString("## Selection boundary\n\n")
	b.WriteString(FormatBoundary(a.Boundary) + "\n\n")
	b.WriteString("Gate495 accepts `diag(1,1,4)` and `kappa_U1 = 6` as a strong dimensionless electroweak Hessian bridge candidate. It does not write a native gauge Hessian, physical weak angle, gauge coupling, W/Z mass, Higgs VEV, or flavor observable.\n\n")

	b.WriteString("## Firewall result\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No physical electroweak number entered the native lane. The action candidate remains scale-free and dimensionless; the physical continuum interpretation remains sealed.\n\n")

	b.WriteString("## Registry update\n\n")
	b.WriteString("### Native\n\n")
	for _, x := range a.Registry.NativeEntries {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n### Bridge\n\n")
	for _, x := range a.Registry.BridgeEntries {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n### Environmental\n\n")
	for _, x := range a.Registry.EnvironmentalEntries {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n### Failed routes\n\n")
	for _, x := range a.Registry.FailedRoutes {
		b.WriteString("- `" + x + "`\n")
	}
	b.WriteString("\n### Open theorems\n\n")
	for _, x := range a.Registry.OpenTheorems {
		b.WriteString("- " + x + "\n")
	}

	b.WriteString("\n## Next step\n\n")
	b.WriteString(fmt.Sprintf("**Gate %d — %s.** %s Primary task: %s.\n\n", a.Next.Gate, a.Next.Title, a.Next.Reason, a.Next.PrimaryTask))
	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}
