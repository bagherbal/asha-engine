// Package generation2scalarkineticvacuumprovenance implements Gate 496:
// Scalar Kinetic Metric Provenance and Vacuum Orientation Closure Audit.
//
// Gate 495 found the strongest current electroweak Hessian clue: the legacy
// canonical finite-action candidate computes a dimensionless broken Hessian
// diag(1,1,4), selects kappa_U1=6 inside that candidate, and gives a positive
// rank-four full electroweak Hessian.  It refused native promotion because the
// scalar kinetic metric I4, vacuum orientation, scalar SU(2)L action, and DΦ
// were still provenance inputs.
//
// Gate 496 audits those inputs directly.  The Hilbert-Schmidt scalar trace does
// prove a positive scalar kinetic *class* and blocks ghost signs, but it does
// not select the active-frame Euclidean I4 metric with unit normalization.  The
// scalar vacuum audit selects a lower two-plane at fixed radius, but leaves an
// S1 phase/gauge freedom inside that plane.  The scalar SU(2) audit supplies an
// abstract doublet representation, but the finite scalar response only selects
// a T3-like pair rotation; full SU(2)L is not selected by scalar data alone.
// Thus Gate 496 partially closes the provenance ledger while still blocking the
// native electroweak Hessian and W/Z registry write.
package generation2scalarkineticvacuumprovenance

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2finiteactionsecondvariation"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarkinetictracepositivity"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarsu2"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarvacuum"
)

const (
	AuditID = "GATE496-SCALAR-KINETIC-METRIC-PROVENANCE-AND-VACUUM-ORIENTATION-CLOSURE-AUDIT"

	StatusGate495Inherited               = "CONDITIONAL_SUPPORT_GATE495_HESSIAN_CANDIDATE_INHERITED"
	StatusHilbertSchmidtMetricClassFound = "CONDITIONAL_SUPPORT_HILBERT_SCHMIDT_SCALAR_METRIC_CLASS_FOUND"
	StatusGhostFreeScalarMetricPreserved = "CONDITIONAL_SUPPORT_GHOST_FREE_SCALAR_KINETIC_METRIC_PRESERVED"
	StatusLowerPairVacuumPlaneSelected   = "CONDITIONAL_SUPPORT_LOWER_PAIR_VACUUM_PLANE_SELECTED"
	StatusDiagnosticVacuumIsMinimizer    = "CONDITIONAL_SUPPORT_DIAGNOSTIC_UNITARY_GAUGE_VECTOR_IS_VALID_MINIMIZER"
	StatusAbstractScalarDoubletAvailable = "CONDITIONAL_SUPPORT_ABSTRACT_SCALAR_SU2_DOUBLET_REPRESENTATION_AVAILABLE"
	StatusProvenanceBoundarySharpened    = "CONDITIONAL_SUPPORT_METRIC_VACUUM_PROVENANCE_BOUNDARY_SHARPENED"
	StatusFirewallPreserved              = "FIREWALL_PRESERVED_NO_ELECTROWEAK_OR_FLAVOR_DATA_IMPORTED"
	StatusNativeRegistryWriteBlocked     = "FIREWALL_BLOCKED_NATIVE_I4_VACUUM_AND_WZ_REGISTRY_WRITE"

	StatusFailedI4MetricNotSelected         = "FAILED_ROUTE_HILBERT_SCHMIDT_TRACE_DOES_NOT_SELECT_ACTIVE_I4_UNIT_METRIC"
	StatusFailedTraceNormalizationSealed    = "FAILED_ROUTE_SCALAR_TRACE_NORMALIZATION_AND_ZH_VALUE_STILL_SEALED"
	StatusFailedVacuumVectorNotSelected     = "FAILED_ROUTE_SCALAR_VACUUM_VECTOR_NOT_NATIVE_SELECTED"
	StatusFailedResidualS1PhaseUnquotiented = "FAILED_ROUTE_RESIDUAL_S1_VACUUM_PHASE_NOT_YET_PROVEN_PURE_GAUGE"
	StatusFailedFullSU2NotSelected          = "FAILED_ROUTE_FULL_SCALAR_SU2_ACTION_NOT_SELECTED_BY_SCALAR_RESPONSE"
	StatusFailedDphiStillUnclosed           = "FAILED_ROUTE_NATIVE_DPHI_PROVENANCE_STILL_UNCLOSED"
	StatusFailedKappaStillBridge            = "FAILED_ROUTE_KAPPA_U1_SIX_REMAINS_BRIDGE_CANDIDATE"
	StatusFailedWZMassStillBlocked          = "FAILED_ROUTE_PHYSICAL_WZ_MASS_MATRIX_STILL_BLOCKED"
	StatusGate497RedirectDefined            = "CONDITIONAL_SUPPORT_GATE497_VACUUM_GAUGE_ORBIT_QUOTIENT_REDIRECT_DEFINED"
)

type Inheritance struct {
	Executed                        bool
	Gate495AuditDefined             bool
	HessianCandidateAccepted        bool
	Diag114BridgeCandidate          bool
	KappaSixBridgeCandidate         bool
	FullHessianBridgeCandidate      bool
	Gate495MetricProvenanceOpen     bool
	Gate495VacuumProvenanceOpen     bool
	Gate495DphiProvenanceOpen       bool
	NoElectroweakFlavorDataImported bool
	Verdict                         string
	Reason                          string
}

type KineticMetricAudit struct {
	Executed                        bool
	TraceFunctionalFormalized       bool
	UsesHilbertSchmidtNorm          bool
	PositiveSemidefinite            bool
	GhostRiskEliminatedStructurally bool
	StrictPositiveConditional       bool
	StrictPositiveProved            bool
	NumericalZHComputed             bool
	AmplitudeValuesSealed           bool
	TraceConventionExplicit         bool
	ActiveRealDimension             int
	MetricClassAvailable            bool
	ActiveI4UnitMetricSelected      bool
	CanonicalNormalizationSelected  bool
	PhysicalKineticScaleDerived     bool
	Verdict                         string
	Reason                          string
}

type VacuumAudit struct {
	Executed                       bool
	ActiveRealDimension            int
	RadialNormalFormSelectsRadius  bool
	RadialNormalFormSelectsVector  bool
	LowPairSelected                bool
	LowPairDimension               int
	HighPairDimension              int
	DiagnosticVacuumIsMinimizer    bool
	ResidualPhaseFreedomDimension  int
	UnitaryGaugeVectorSelected     bool
	CanonicalPhaseSelected         bool
	FiniteVacuumOrientationDerived bool
	GaugeEatingDiagnosticAvailable bool
	FullGaugeEatingTheoremDerived  bool
	VacuumOrbitCanonical           bool
	Verdict                        string
	Reason                         string
}

type ScalarSU2Audit struct {
	Executed                       bool
	ActiveRealDimension            int
	ComplexDoubletDimension        int
	AbstractDoubletRepresentation  bool
	PairDegenerate                 bool
	PairSplitNonzero               bool
	FullSU2SelectedByScalarData    bool
	U1PairRotationSelected         bool
	CanonicalComplexStructure      bool
	CovariantDerivativeDerived     bool
	GaugeEatingTheoremDerived      bool
	PairRotationCommutantDimension int
	Verdict                        string
	Reason                         string
}

type ProvenanceBoundary struct {
	Executed                        bool
	MetricClassNative               bool
	ActiveI4MetricNative            bool
	VacuumPlaneNative               bool
	VacuumVectorNative              bool
	ResidualPhaseQuotientNative     bool
	FullScalarSU2Native             bool
	NativeDphiClosed                bool
	CanonicalActionProvenanceClosed bool
	KappaU1NativeSelected           bool
	GaugeHessianNativeSelected      bool
	WZMassMatrixDerived             bool
	Verdict                         string
	Reason                          string
}

type Firewall struct {
	Executed                  bool
	ObservedWMassImported     bool
	ObservedZMassImported     bool
	ObservedHiggsMassImported bool
	FermiConstantImported     bool
	WeakAngleImported         bool
	FineStructureImported     bool
	GaugeCouplingImported     bool
	YukawaImported            bool
	CKMPMNSImported           bool
	NativeI4MetricWritten     bool
	NativeVacuumVectorWritten bool
	NativeSU2ActionWritten    bool
	NativeDphiWritten         bool
	NativeKappaWritten        bool
	NativeGaugeHessianWritten bool
	NativeWZMassWritten       bool
	Verdict                   string
	Reason                    string
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
	Metric      KineticMetricAudit
	Vacuum      VacuumAudit
	ScalarSU2   ScalarSU2Audit
	Boundary    ProvenanceBoundary
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
	g495, err := generation2finiteactionsecondvariation.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate495 finite action second-variation audit: %w", err)
	}
	kt, err := scalarkinetictracepositivity.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit scalar kinetic trace positivity audit: %w", err)
	}
	sv, err := scalarvacuum.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit scalar vacuum orientation audit: %w", err)
	}
	su2, err := scalarsu2.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit scalar SU(2) audit: %w", err)
	}

	a := Analysis{}
	a.Inheritance = buildInheritance(g495)
	a.Metric = buildMetricAudit(kt, g495)
	a.Vacuum = buildVacuumAudit(sv)
	a.ScalarSU2 = buildScalarSU2Audit(su2)
	a.Boundary = buildBoundary(a.Metric, a.Vacuum, a.ScalarSU2, g495)
	a.Firewall = buildFirewall()
	a.Registry = buildRegistryUpdate(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g generation2finiteactionsecondvariation.Analysis) Inheritance {
	return Inheritance{
		Executed:                        true,
		Gate495AuditDefined:             true,
		HessianCandidateAccepted:        g.Boundary.DimensionlessSecondVariationCandidate,
		Diag114BridgeCandidate:          g.Boundary.Diag114AcceptedAsBridgeCandidate,
		KappaSixBridgeCandidate:         g.Boundary.KappaSixAcceptedAsBridgeCandidate,
		FullHessianBridgeCandidate:      g.Boundary.FullHessianAcceptedAsBridgeCandidate,
		Gate495MetricProvenanceOpen:     !g.Provenance.Gate492ScalarKineticMetricNative,
		Gate495VacuumProvenanceOpen:     !g.Provenance.Gate492ScalarVacuumOrientationNative,
		Gate495DphiProvenanceOpen:       !g.Provenance.Gate492NativeDphiDerived,
		NoElectroweakFlavorDataImported: g.Firewall.Executed && !g.Firewall.ObservedWMassImported && !g.Firewall.ObservedZMassImported && !g.Firewall.WeakAngleImported && !g.Firewall.GaugeCouplingImported && !g.Firewall.YukawaImported && !g.Firewall.CKMPMNSImported,
		Verdict:                         StatusGate495Inherited,
		Reason:                          "Gate495 supplies the dimensionless Hessian candidate but explicitly leaves the scalar I4 metric, vacuum orientation, and DΦ provenance open.",
	}
}

func buildMetricAudit(k scalarkinetictracepositivity.Analysis, g generation2finiteactionsecondvariation.Analysis) KineticMetricAudit {
	metricClass := k.Trace.Formalized && k.Trace.UsesHilbertSchmidtNorm && k.Positivity.PositiveSemidefinite && k.Positivity.GhostRiskEliminatedStructurally
	return KineticMetricAudit{
		Executed:                        true,
		TraceFunctionalFormalized:       k.Trace.Formalized,
		UsesHilbertSchmidtNorm:          k.Trace.UsesHilbertSchmidtNorm,
		PositiveSemidefinite:            k.Positivity.PositiveSemidefinite,
		GhostRiskEliminatedStructurally: k.Positivity.GhostRiskEliminatedStructurally,
		StrictPositiveConditional:       k.Positivity.StrictPositiveConditional,
		StrictPositiveProved:            k.Positivity.StrictPositiveProved,
		NumericalZHComputed:             k.ZH.NumericalZHComputed,
		AmplitudeValuesSealed:           k.Seals.AllNumericalValuesSealed,
		TraceConventionExplicit:         true,
		ActiveRealDimension:             g.Candidate.ActiveRealDimension,
		MetricClassAvailable:            metricClass,
		ActiveI4UnitMetricSelected:      false,
		CanonicalNormalizationSelected:  false,
		PhysicalKineticScaleDerived:     false,
		Verdict:                         StatusHilbertSchmidtMetricClassFound,
		Reason:                          "The scalar Hilbert-Schmidt trace proves a positive kinetic inner-product class, but it is an amplitude/convention-dependent trace carrier; it does not by itself select the active-frame unit metric I4 used by the canonical Hessian candidate.",
	}
}

func buildVacuumAudit(v scalarvacuum.Analysis) VacuumAudit {
	return VacuumAudit{
		Executed:                       true,
		ActiveRealDimension:            v.ActiveRealDimension,
		RadialNormalFormSelectsRadius:  v.RadialNormalFormSelectsRadius,
		RadialNormalFormSelectsVector:  v.RadialNormalFormSelectsVector,
		LowPairSelected:                v.LowPairSelected,
		LowPairDimension:               v.LowPairDimension,
		HighPairDimension:              v.HighPairDimension,
		DiagnosticVacuumIsMinimizer:    v.DiagnosticVacuumIsMinimizer,
		ResidualPhaseFreedomDimension:  v.ResidualPhaseFreedomDimension,
		UnitaryGaugeVectorSelected:     v.UnitaryGaugeVectorSelected,
		CanonicalPhaseSelected:         v.CanonicalPhaseSelected,
		FiniteVacuumOrientationDerived: v.FiniteVacuumOrientationDerived,
		GaugeEatingDiagnosticAvailable: v.GaugeEatingDiagnosticAvailable,
		FullGaugeEatingTheoremDerived:  v.FullGaugeEatingTheoremDerived,
		VacuumOrbitCanonical:           v.LowPairSelected && v.ResidualPhaseFreedomDimension == 1 && !v.FiniteVacuumOrientationDerived,
		Verdict:                        StatusLowerPairVacuumPlaneSelected,
		Reason:                         "The finite scalar/contact response selects the lower two-plane at fixed radius, and the unitary-gauge vector is a valid minimizer representative.  The particular vector and phase remain unselected until the residual S1 is proven to be pure gauge or quotiented natively.",
	}
}

func buildScalarSU2Audit(s scalarsu2.Analysis) ScalarSU2Audit {
	return ScalarSU2Audit{
		Executed:                       true,
		ActiveRealDimension:            s.ActiveRealDimension,
		ComplexDoubletDimension:        s.ComplexDoubletDimension,
		AbstractDoubletRepresentation:  s.AbstractDoubletRepresentation,
		PairDegenerate:                 s.PairDegenerate,
		PairSplitNonzero:               s.PairSplit > 0,
		FullSU2SelectedByScalarData:    s.FullSU2SelectedByScalarData,
		U1PairRotationSelected:         s.U1PairRotationSelected,
		CanonicalComplexStructure:      s.CanonicalComplexStructure,
		CovariantDerivativeDerived:     s.CovariantDerivativeDerived,
		GaugeEatingTheoremDerived:      s.GaugeEatingTheoremDerived,
		PairRotationCommutantDimension: s.PairRotationCommutantDimension,
		Verdict:                        StatusAbstractScalarDoubletAvailable,
		Reason:                         "A four-real active scalar frame supports the abstract realification of a complex SU(2) doublet.  But the finite scalar response with split pair spectrum selects only a T3-like pair rotation; full SU(2)L and the scalar complex structure are not selected by scalar data alone.",
	}
}

func buildBoundary(m KineticMetricAudit, v VacuumAudit, s ScalarSU2Audit, g generation2finiteactionsecondvariation.Analysis) ProvenanceBoundary {
	metricClassNative := m.MetricClassAvailable
	vacuumPlaneNative := v.LowPairSelected && v.DiagnosticVacuumIsMinimizer
	fullClosed := metricClassNative && m.ActiveI4UnitMetricSelected && vacuumPlaneNative && v.FiniteVacuumOrientationDerived && s.FullSU2SelectedByScalarData && s.CovariantDerivativeDerived
	return ProvenanceBoundary{
		Executed:                        true,
		MetricClassNative:               metricClassNative,
		ActiveI4MetricNative:            m.ActiveI4UnitMetricSelected,
		VacuumPlaneNative:               vacuumPlaneNative,
		VacuumVectorNative:              v.FiniteVacuumOrientationDerived,
		ResidualPhaseQuotientNative:     false,
		FullScalarSU2Native:             s.FullSU2SelectedByScalarData,
		NativeDphiClosed:                s.CovariantDerivativeDerived,
		CanonicalActionProvenanceClosed: fullClosed,
		KappaU1NativeSelected:           fullClosed && g.Candidate.KappaSixSelectedInCandidate,
		GaugeHessianNativeSelected:      fullClosed && g.Candidate.FullGaugeHessianSelectedInCandidate,
		WZMassMatrixDerived:             false,
		Verdict:                         StatusProvenanceBoundarySharpened,
		Reason:                          "Gate496 upgrades the scalar provenance ledger from 'unknown' to a precise split: positive metric class and low-pair vacuum plane are supported; unit I4 normalization, vector phase, full scalar SU(2), DΦ, kappa promotion, and W/Z masses remain blocked.",
	}
}

func buildFirewall() Firewall {
	return Firewall{
		Executed:                  true,
		ObservedWMassImported:     false,
		ObservedZMassImported:     false,
		ObservedHiggsMassImported: false,
		FermiConstantImported:     false,
		WeakAngleImported:         false,
		FineStructureImported:     false,
		GaugeCouplingImported:     false,
		YukawaImported:            false,
		CKMPMNSImported:           false,
		NativeI4MetricWritten:     false,
		NativeVacuumVectorWritten: false,
		NativeSU2ActionWritten:    false,
		NativeDphiWritten:         false,
		NativeKappaWritten:        false,
		NativeGaugeHessianWritten: false,
		NativeWZMassWritten:       false,
		Verdict:                   StatusFirewallPreserved,
		Reason:                    "No observed electroweak, Higgs, gauge-coupling, Yukawa, CKM, or PMNS datum is imported; supported scalar structures remain dimensionless provenance data only.",
	}
}

func buildRegistryUpdate(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"Positive Hilbert-Schmidt scalar kinetic metric class is preserved as a ghost-free structural carrier.",
			"The scalar/contact response selects the lower active pair plane at fixed scalar radius.",
		},
		BridgeEntries: []string{
			"The active-frame I4 metric remains the canonical candidate metric, not an independently native-selected unit metric.",
			"The unitary-gauge vacuum vector is a valid minimizer representative inside the selected lower two-plane, but its residual S1 phase is not yet quotiented natively.",
			"The abstract scalar SU(2) doublet representation is available, while full scalar SU(2)L and DΦ remain bridge/provenance objects.",
			"Gate495 diag(1,1,4), kappa_U1=6, and positive rank-four Hessian remain strong dimensionless bridge candidates only.",
		},
		EnvironmentalEntries: []string{
			"Observed W/Z masses, Higgs VEV, Fermi constant, weak mixing angle, alpha, running gauge couplings, Yukawa matrices, CKM, and PMNS remain sealed.",
		},
		FailedRoutes: []string{
			StatusFailedI4MetricNotSelected,
			StatusFailedTraceNormalizationSealed,
			StatusFailedVacuumVectorNotSelected,
			StatusFailedResidualS1PhaseUnquotiented,
			StatusFailedFullSU2NotSelected,
			StatusFailedDphiStillUnclosed,
			StatusFailedKappaStillBridge,
			StatusFailedWZMassStillBlocked,
		},
		OpenTheorems: []string{
			"Prove that the residual S1 phase inside the selected lower scalar pair is pure electroweak gauge redundancy, or keep the exact vector orientation sealed.",
			"Derive the active I4 unit metric from a finite orthonormal-frame theorem rather than choosing the Euclidean metric in the canonical action candidate.",
			"Derive full scalar SU(2)L and DΦ from finite contact/spectral data before promoting kappa_U1=6 or the electroweak Hessian to native status.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        497,
		Title:       "Vacuum Gauge-Orbit Quotient and Unitary-Gauge Representative Audit",
		Reason:      "Gate496 selects the lower vacuum plane but not a unique vector; the remaining S1 may be gauge redundancy rather than physical missing data.",
		PrimaryTask: "prove or reject that the residual lower-pair S1 phase is entirely quotiented by the electroweak gauge orbit, so the unitary-gauge representative can be used without a native vector-selector theorem and without importing v, W/Z masses, theta_W, alpha, or Yukawa data",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.HessianCandidateAccepted || !a.Inheritance.Gate495MetricProvenanceOpen || !a.Inheritance.Gate495VacuumProvenanceOpen || !a.Inheritance.Gate495DphiProvenanceOpen {
		return fmt.Errorf("Gate496 did not inherit the Gate495 provenance obstruction: %+v", a.Inheritance)
	}
	if !a.Metric.Executed || !a.Metric.MetricClassAvailable || !a.Metric.PositiveSemidefinite || !a.Metric.GhostRiskEliminatedStructurally || a.Metric.ActiveI4UnitMetricSelected || a.Metric.NumericalZHComputed || a.Metric.PhysicalKineticScaleDerived {
		return fmt.Errorf("metric audit over-promoted or failed: %+v", a.Metric)
	}
	if !a.Vacuum.Executed || !a.Vacuum.LowPairSelected || a.Vacuum.LowPairDimension != 2 || !a.Vacuum.DiagnosticVacuumIsMinimizer || a.Vacuum.ResidualPhaseFreedomDimension != 1 || a.Vacuum.FiniteVacuumOrientationDerived || a.Vacuum.CanonicalPhaseSelected {
		return fmt.Errorf("vacuum audit over-promoted or failed: %+v", a.Vacuum)
	}
	if !a.ScalarSU2.Executed || !a.ScalarSU2.AbstractDoubletRepresentation || a.ScalarSU2.FullSU2SelectedByScalarData || !a.ScalarSU2.U1PairRotationSelected || a.ScalarSU2.CanonicalComplexStructure || a.ScalarSU2.CovariantDerivativeDerived {
		return fmt.Errorf("scalar SU2 audit over-promoted or failed: %+v", a.ScalarSU2)
	}
	if !a.Boundary.Executed || !a.Boundary.MetricClassNative || a.Boundary.ActiveI4MetricNative || !a.Boundary.VacuumPlaneNative || a.Boundary.VacuumVectorNative || a.Boundary.ResidualPhaseQuotientNative || a.Boundary.FullScalarSU2Native || a.Boundary.NativeDphiClosed || a.Boundary.CanonicalActionProvenanceClosed || a.Boundary.KappaU1NativeSelected || a.Boundary.GaugeHessianNativeSelected || a.Boundary.WZMassMatrixDerived {
		return fmt.Errorf("boundary violated: %+v", a.Boundary)
	}
	if a.Firewall.ObservedWMassImported || a.Firewall.WeakAngleImported || a.Firewall.GaugeCouplingImported || a.Firewall.YukawaImported || a.Firewall.NativeI4MetricWritten || a.Firewall.NativeVacuumVectorWritten || a.Firewall.NativeKappaWritten || a.Firewall.NativeWZMassWritten {
		return fmt.Errorf("firewall leak: %+v", a.Firewall)
	}
	return nil
}

func truth(a Analysis) string {
	return "Gate496 closes the provenance ledger one layer deeper: the scalar kinetic trace gives a genuine ghost-free Hilbert-Schmidt metric class, and the finite scalar/contact response selects the lower two-dimensional vacuum plane.  But the exact active I4 unit normalization, the representative vector inside the residual S1, full scalar SU(2)L, and native DΦ are not derived.  Therefore the Gate495 electroweak Hessian shape, diag(1,1,4), kappa_U1=6, and the positive full Hessian remain powerful dimensionless bridge candidates, while native W/Z masses, weak angle, gauge couplings, Higgs VEV, and flavor data stay blocked."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate495=%v hessian_candidate=%v diag114=%v kappa6=%v full_hessian=%v metric_open=%v vacuum_open=%v dphi_open=%v no_data=%v verdict=%s reason=%s", x.Gate495AuditDefined, x.HessianCandidateAccepted, x.Diag114BridgeCandidate, x.KappaSixBridgeCandidate, x.FullHessianBridgeCandidate, x.Gate495MetricProvenanceOpen, x.Gate495VacuumProvenanceOpen, x.Gate495DphiProvenanceOpen, x.NoElectroweakFlavorDataImported, x.Verdict, x.Reason)
}

func FormatMetric(x KineticMetricAudit) string {
	return fmt.Sprintf("trace=%v hilbert_schmidt=%v psd=%v ghost_free=%v strict_conditional=%v strict_proved=%v numerical_ZH=%v amplitudes_sealed=%v trace_convention_explicit=%v active_dim=%d metric_class=%v I4_native=%v normalization_native=%v physical_scale=%v verdict=%s reason=%s", x.TraceFunctionalFormalized, x.UsesHilbertSchmidtNorm, x.PositiveSemidefinite, x.GhostRiskEliminatedStructurally, x.StrictPositiveConditional, x.StrictPositiveProved, x.NumericalZHComputed, x.AmplitudeValuesSealed, x.TraceConventionExplicit, x.ActiveRealDimension, x.MetricClassAvailable, x.ActiveI4UnitMetricSelected, x.CanonicalNormalizationSelected, x.PhysicalKineticScaleDerived, x.Verdict, x.Reason)
}

func FormatVacuum(x VacuumAudit) string {
	return fmt.Sprintf("active_dim=%d radius_selected=%v vector_selected_by_radial=%v low_pair=%v low_dim=%d high_dim=%d diagnostic_minimizer=%v residual_phase_dim=%d unitary_vector_selected=%v canonical_phase=%v finite_orientation=%v gauge_eating_diag=%v full_gauge_eating=%v orbit_canonical=%v verdict=%s reason=%s", x.ActiveRealDimension, x.RadialNormalFormSelectsRadius, x.RadialNormalFormSelectsVector, x.LowPairSelected, x.LowPairDimension, x.HighPairDimension, x.DiagnosticVacuumIsMinimizer, x.ResidualPhaseFreedomDimension, x.UnitaryGaugeVectorSelected, x.CanonicalPhaseSelected, x.FiniteVacuumOrientationDerived, x.GaugeEatingDiagnosticAvailable, x.FullGaugeEatingTheoremDerived, x.VacuumOrbitCanonical, x.Verdict, x.Reason)
}

func FormatScalarSU2(x ScalarSU2Audit) string {
	return fmt.Sprintf("active_dim=%d complex_dim=%d abstract_doublet=%v pair_degenerate=%v pair_split=%v full_SU2_native=%v U1_pair_rotation=%v canonical_complex=%v Dphi=%v gauge_eating=%v commutant_dim=%d verdict=%s reason=%s", x.ActiveRealDimension, x.ComplexDoubletDimension, x.AbstractDoubletRepresentation, x.PairDegenerate, x.PairSplitNonzero, x.FullSU2SelectedByScalarData, x.U1PairRotationSelected, x.CanonicalComplexStructure, x.CovariantDerivativeDerived, x.GaugeEatingTheoremDerived, x.PairRotationCommutantDimension, x.Verdict, x.Reason)
}

func FormatBoundary(x ProvenanceBoundary) string {
	return fmt.Sprintf("metric_class=%v I4_native=%v vacuum_plane=%v vacuum_vector=%v phase_quotient=%v full_SU2=%v Dphi=%v action_provenance=%v kappa_native=%v hessian_native=%v WZ=%v verdict=%s reason=%s", x.MetricClassNative, x.ActiveI4MetricNative, x.VacuumPlaneNative, x.VacuumVectorNative, x.ResidualPhaseQuotientNative, x.FullScalarSU2Native, x.NativeDphiClosed, x.CanonicalActionProvenanceClosed, x.KappaU1NativeSelected, x.GaugeHessianNativeSelected, x.WZMassMatrixDerived, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("observed_W=%v observed_Z=%v observed_Higgs=%v Fermi=%v theta=%v alpha=%v gauge_coupling=%v Yukawa=%v CKM_PMNS=%v native_I4=%v native_vacuum=%v native_SU2=%v native_Dphi=%v native_kappa=%v native_hessian=%v native_WZ=%v verdict=%s reason=%s", x.ObservedWMassImported, x.ObservedZMassImported, x.ObservedHiggsMassImported, x.FermiConstantImported, x.WeakAngleImported, x.FineStructureImported, x.GaugeCouplingImported, x.YukawaImported, x.CKMPMNSImported, x.NativeI4MetricWritten, x.NativeVacuumVectorWritten, x.NativeSU2ActionWritten, x.NativeDphiWritten, x.NativeKappaWritten, x.NativeGaugeHessianWritten, x.NativeWZMassWritten, x.Verdict, x.Reason)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 496 Registry Audit — Scalar Kinetic Metric Provenance and Vacuum Orientation Closure Audit\n\n")
	b.WriteString("## Verdict\n\n")
	for _, v := range []string{
		StatusGate495Inherited,
		StatusHilbertSchmidtMetricClassFound,
		StatusGhostFreeScalarMetricPreserved,
		StatusLowerPairVacuumPlaneSelected,
		StatusDiagnosticVacuumIsMinimizer,
		StatusAbstractScalarDoubletAvailable,
		StatusProvenanceBoundarySharpened,
		StatusFailedI4MetricNotSelected,
		StatusFailedTraceNormalizationSealed,
		StatusFailedVacuumVectorNotSelected,
		StatusFailedResidualS1PhaseUnquotiented,
		StatusFailedFullSU2NotSelected,
		StatusFailedDphiStillUnclosed,
		StatusFailedKappaStillBridge,
		StatusFailedWZMassStillBlocked,
		StatusFirewallPreserved,
		StatusNativeRegistryWriteBlocked,
		StatusGate497RedirectDefined,
	} {
		b.WriteString("- `" + v + "`\n")
	}

	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("Gate495 supplies a strong dimensionless electroweak Hessian candidate, but it explicitly depends on a diagnostic `DΦ`, an active-frame `I4` metric, and a chosen vacuum representative. Gate496 audits those ingredients directly.\n\n")

	b.WriteString("## Scalar kinetic metric audit\n\n")
	b.WriteString(FormatMetric(a.Metric) + "\n\n")
	b.WriteString("The Hilbert-Schmidt scalar trace is strong enough to block ghost kinetic signs and define a positive metric class. It is not yet strong enough to select the exact active-frame unit metric `I4` or the numerical `Z_H` normalization.\n\n")

	b.WriteString("## Vacuum orientation audit\n\n")
	b.WriteString(FormatVacuum(a.Vacuum) + "\n\n")
	b.WriteString("The finite scalar/contact response selects the lower active two-plane, and the unitary-gauge vector used by the Hessian candidate is a valid minimizer. But the specific vector and phase inside the residual `S1` are not yet natively selected.\n\n")

	b.WriteString("## Scalar SU(2) action audit\n\n")
	b.WriteString(FormatScalarSU2(a.ScalarSU2) + "\n\n")
	b.WriteString("A four-real scalar frame supports an abstract complex doublet representation, but the finite scalar response does not select full `SU(2)_L`; it selects only a commuting pair-rotation direction. Therefore full scalar `SU(2)_L` and native `DΦ` remain open.\n\n")

	b.WriteString("## Provenance boundary\n\n")
	b.WriteString(FormatBoundary(a.Boundary) + "\n\n")
	b.WriteString("Gate496 partially closes the Gate495 provenance gap: positive metric class and lower vacuum plane are supported. It blocks native promotion of `I4`, vacuum vector, full scalar `SU(2)_L`, `DΦ`, `kappa_U1=6`, gauge Hessian, and W/Z mass matrix.\n\n")

	b.WriteString("## Firewall result\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No physical electroweak or flavor number entered the native lane. The audit remains a dimensionless finite-structure provenance check.\n\n")

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
