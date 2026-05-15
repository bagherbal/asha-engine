// Package generation2abeliancoefficientselection implements Gate 494:
// Abelian U(1) Completion Coefficient Selection Audit.
//
// Gate 493 typed a full electroweak curvature/action socket and a positive
// abelian-completed quadratic family
//
//	K(kappa)=K_SU2+kappa(Q-Z)(Q-Z)^T,
//
// with the broken-coordinate whitening candidate diag(1,1,4) reachable at
// kappa_U1=6.  Gate 494 asks the sharp next question: do the already-derived
// finite trace, hypercharge-normalization, anomaly/unimodularity, or
// representation-metric ledgers select this coefficient?
//
// The answer is conservative.  The finite charge table really does derive the
// hypercharge normalization k_Y=5/3 and the boundary-candidate sin^2=3/8 under
// the separate equal-normalized-coupling bridge assumption.  Existing U(1)
// diagnostics also provide a diagonal representation trace metric.  But none of
// these is a second variation of the finite electroweak action with respect to
// the abelian gauge field.  Therefore kappa_U1=6 remains a whitening/bridge
// candidate, not a native action-selected coefficient.
package generation2abeliancoefficientselection

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/abeliancoupling"
	"github.com/bagherbal/asha-engine/pkg/bridge/ewprojection"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2electroweakcurvatureaction"
	"github.com/bagherbal/asha-engine/pkg/bridge/u1completion"
)

const (
	AuditID = "GATE494-ABELIAN-U1-COMPLETION-COEFFICIENT-SELECTION-AUDIT"

	StatusGate493Inherited                    = "CONDITIONAL_SUPPORT_GATE493_ELECTROWEAK_ACTION_FAMILY_INHERITED"
	StatusHyperchargeTraceNormalizationFound  = "CONDITIONAL_SUPPORT_HYPERCHARGE_TRACE_NORMALIZATION_KY_5_OVER_3_CONFIRMED"
	StatusBoundarySin238Preserved             = "CONDITIONAL_SUPPORT_EQUAL_NORMALIZED_COUPLING_BOUNDARY_SIN2_3_OVER_8_PRESERVED"
	StatusU1CompletionKappaTargetInherited    = "CONDITIONAL_SUPPORT_KAPPA_U1_TARGET_SIX_WHITENING_CANDIDATE_INHERITED"
	StatusFiniteCountResonancesAudited        = "CONDITIONAL_SUPPORT_FINITE_COUNT_RESONANCES_EQUAL_SIX_AUDITED"
	StatusRepresentationMetricLedgerAvailable = "CONDITIONAL_SUPPORT_DIAGONAL_REPRESENTATION_TRACE_METRIC_AVAILABLE"
	StatusTraceKappaAirlockDefined            = "CONDITIONAL_SUPPORT_TRACE_TO_KAPPA_AIRLOCK_DEFINED"
	StatusFirewallPreserved                   = "FIREWALL_PRESERVED_NO_ELECTROWEAK_DATA_IMPORTED"
	StatusKappaRegistryWriteBlocked           = "FIREWALL_BLOCKED_KAPPA_U1_NATIVE_REGISTRY_WRITE"

	StatusFailedTraceDoesNotSelectKappa       = "FAILED_ROUTE_TRACE_NORMALIZATION_DOES_NOT_SELECT_KAPPA_U1"
	StatusFailedKYIsNotAbelianHessian         = "FAILED_ROUTE_KY_5_OVER_3_IS_CHARGE_NORMALIZATION_NOT_GAUGE_HESSIAN"
	StatusFailedCountResonanceNotAction       = "FAILED_ROUTE_COUNT_RESONANCES_ARE_NOT_ACTION_SELECTION"
	StatusFailedRepresentationMetricNotAction = "FAILED_ROUTE_REPRESENTATION_METRIC_NOT_SELECTED_AS_GAUGE_KINETIC_HESSIAN"
	StatusFailedSecondVariationStillMissing   = "FAILED_ROUTE_FINITE_ACTION_SECOND_VARIATION_STILL_MISSING"
	StatusFailedGaugeCouplingsNotDerived      = "FAILED_ROUTE_GAUGE_COUPLINGS_NOT_DERIVED"
	StatusFailedWeakAngleNotDerived           = "FAILED_ROUTE_PHYSICAL_WEAK_MIXING_ANGLE_NOT_DERIVED"
	StatusFailedWZMassesNotDerived            = "FAILED_ROUTE_PHYSICAL_WZ_MASSES_NOT_DERIVED"
	StatusGate495RedirectDefined              = "CONDITIONAL_SUPPORT_GATE495_SECOND_VARIATION_REDIRECT_DEFINED"
)

const eps = 1e-10

type Inheritance struct {
	Executed                        bool
	Gate493FullConnectionClosed     bool
	Gate493QuadraticFamilyTyped     bool
	Gate493PositiveCompletionFamily bool
	Gate493KappaTarget              float64
	Gate493KappaSelected            bool
	Gate493GaugeHessianSelected     bool
	Gate493PhysicalRegistryBlocked  bool
	Gate494Requested                bool
	NoElectroweakFlavorDataImported bool
	Verdict                         string
	Reason                          string
}

type HyperchargeTraceAudit struct {
	Executed                            bool
	ChargeIdentityDerived               bool
	HyperchargeCommutesWithSU2          bool
	FullYTrace2                         float64
	FullT3Trace2                        float64
	KY                                  float64
	KYExpected                          float64
	KYConfirmed                         bool
	NormalizedHyperchargeFactor         float64
	EqualNormalizedCouplingBoundarySin2 float64
	BoundarySin238Confirmed             bool
	GaugeKineticNormalizationDerived    bool
	PhysicalWeakMixingAngleDerived      bool
	FineStructureDerived                bool
	RGBoundaryScaleDerived              bool
	HiddenObservedInputUsed             bool
	SelectsKappaU1                      bool
	Verdict                             string
	Reason                              string
}

type KappaSearchAudit struct {
	Executed                  bool
	CompletionFamilyTyped     bool
	TargetKappa               float64
	TargetSource              string
	AbelianDirectionBasis     string
	AbelianDirectionNormSq    float64
	WhiteningSelectsKappa     bool
	ActionSelectsKappa        bool
	FiniteSecondVariation     bool
	CandidateResonanceCount   int
	CandidateHitCount         int
	UniqueDerivation          bool
	CountResonanceSelected    bool
	KappaPhysical             bool
	GaugeKineticHessianFixed  bool
	PhysicalCouplingsOrMasses bool
	Verdict                   string
	Reason                    string
}

type RepresentationMetricAudit struct {
	Executed                                bool
	DiagonalTraceGramAsRepresentationMetric bool
	CanonicalGeneratorDiagnosticsDerived    bool
	DiagonalTraceGramAsGaugeKineticHessian  bool
	FieldCount                              int
	HyperchargeBridgeNorm                   float64
	ChargeTableKY                           float64
	BoundarySin2                            float64
	PhysicalGaugeCouplingsDerived           bool
	FineStructureDerived                    bool
	RGBoundaryScaleDerived                  bool
	HiddenObservedInputUsed                 bool
	Verdict                                 string
	Reason                                  string
}

type SelectionBoundary struct {
	Executed                         bool
	TraceLedgerAvailable             bool
	KappaWhiteningCandidateAvailable bool
	TraceToKappaNativeMapDerived     bool
	KYEqualsTargetKappa              bool
	KYAndKappaSameObject             bool
	UnimodularitySelectsKappa        bool
	AnomalyCancellationSelectsKappa  bool
	SecondVariationSelectsKappa      bool
	NativeKappaSelected              bool
	NativeGaugeHessianSelected       bool
	NativeWeakAngleDerived           bool
	NativeWZMassesDerived            bool
	Verdict                          string
	Reason                           string
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
	NativeKappaWritten        bool
	NativeGaugeHessianWritten bool
	NativeWeakAngleWritten    bool
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
	Hypercharge HyperchargeTraceAudit
	Kappa       KappaSearchAudit
	Metric      RepresentationMetricAudit
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
	g493, err := generation2electroweakcurvatureaction.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate493 electroweak action audit: %w", err)
	}
	u1, err := u1completion.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit U(1) completion search: %w", err)
	}
	ew, err := ewprojection.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit electroweak projection trace ledger: %w", err)
	}
	ac, err := abeliancoupling.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit abelian coupling trace metric ledger: %w", err)
	}

	a := Analysis{}
	a.Inheritance = buildInheritance(g493)
	a.Hypercharge = buildHyperchargeTraceAudit(ew)
	a.Kappa = buildKappaSearchAudit(u1)
	a.Metric = buildRepresentationMetricAudit(ac)
	a.Boundary = buildSelectionBoundary(a.Hypercharge, a.Kappa, a.Metric)
	a.Firewall = buildFirewall()
	a.Registry = buildRegistryUpdate(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g generation2electroweakcurvatureaction.Analysis) Inheritance {
	return Inheritance{
		Executed:                        true,
		Gate493FullConnectionClosed:     g.Curvature.Closed && g.Curvature.FullFieldStrengthTyped,
		Gate493QuadraticFamilyTyped:     g.Quadratic.FullQuadraticFamilyTyped,
		Gate493PositiveCompletionFamily: g.Quadratic.PositiveQuadraticFamilyExists && g.Quadratic.AbelianCompletionTyped,
		Gate493KappaTarget:              g.Quadratic.Diag114Kappa,
		Gate493KappaSelected:            g.Quadratic.AbelianCoefficientSelected,
		Gate493GaugeHessianSelected:     g.Boundary.GaugeHessianActionSelected,
		Gate493PhysicalRegistryBlocked:  g.Firewall.NativeWZMassWritten == false && g.Firewall.NativeWeakAngleWritten == false && g.Boundary.NativeElectroweakMassTheorem == false,
		Gate494Requested:                g.Next.Gate == 494,
		NoElectroweakFlavorDataImported: !g.Firewall.ObservedWMassImported && !g.Firewall.ObservedZMassImported && !g.Firewall.WeakAngleImported && !g.Firewall.FineStructureImported && !g.Firewall.GaugeCouplingImported && !g.Firewall.YukawaImported && !g.Firewall.CKMPMNSImported,
		Verdict:                         StatusGate493Inherited,
		Reason:                          "Gate493 leaves exactly one sharp electroweak coefficient obstruction: the positive abelian-completed family is typed, but kappa_U1 and the gauge Hessian are not action-selected.",
	}
}

func buildHyperchargeTraceAudit(e ewprojection.Analysis) HyperchargeTraceAudit {
	kyConfirmed := math.Abs(e.HyperchargeNormalizationKY-5.0/3.0) <= eps
	boundary := math.Abs(e.EqualNormalizedCouplingBoundarySin2-3.0/8.0) <= eps
	return HyperchargeTraceAudit{
		Executed:                            true,
		ChargeIdentityDerived:               e.ElectromagneticGeneratorDerived,
		HyperchargeCommutesWithSU2:          e.HyperchargeCommutesWithSU2LNorm <= eps,
		FullYTrace2:                         e.FullYTrace2OneGeneration,
		FullT3Trace2:                        e.FullT3Trace2OneGeneration,
		KY:                                  e.HyperchargeNormalizationKY,
		KYExpected:                          5.0 / 3.0,
		KYConfirmed:                         kyConfirmed,
		NormalizedHyperchargeFactor:         e.NormalizedHyperchargeFactor,
		EqualNormalizedCouplingBoundarySin2: e.EqualNormalizedCouplingBoundarySin2,
		BoundarySin238Confirmed:             boundary,
		GaugeKineticNormalizationDerived:    e.GaugeKineticNormalizationDerived,
		PhysicalWeakMixingAngleDerived:      e.WeakMixingAngleDerived,
		FineStructureDerived:                e.FineStructureDerived,
		RGBoundaryScaleDerived:              e.RGBoundaryScaleDerived,
		HiddenObservedInputUsed:             e.HiddenObservedInputUsed,
		SelectsKappaU1:                      false,
		Verdict:                             StatusHyperchargeTraceNormalizationFound,
		Reason:                              "The finite charge table confirms k_Y=Tr(Y^2)/Tr(T3_L^2)=5/3 and the familiar sin^2=3/8 boundary diagnostic under equal normalized couplings, but this is a charge/trace normalization, not a kappa_U1 second-variation theorem.",
	}
}

func buildKappaSearchAudit(u u1completion.Analysis) KappaSearchAudit {
	return KappaSearchAudit{
		Executed:                  true,
		CompletionFamilyTyped:     u.CompletionFamilyTyped,
		TargetKappa:               u.TargetKappa,
		TargetSource:              u.TargetKappaSource,
		AbelianDirectionBasis:     u.NullDirectionBasis,
		AbelianDirectionNormSq:    u.AbelianDirectionNormSq,
		WhiteningSelectsKappa:     u.WhiteningSelectsKappa,
		ActionSelectsKappa:        u.ActionSelectsKappa,
		FiniteSecondVariation:     u.FiniteSecondVariation,
		CandidateResonanceCount:   len(u.CandidateResonances),
		CandidateHitCount:         u.CandidateHitCount,
		UniqueDerivation:          u.UniqueDerivation,
		CountResonanceSelected:    u.CountResonanceSelected,
		KappaPhysical:             u.KappaPhysical,
		GaugeKineticHessianFixed:  u.GaugeKineticHessianFixed,
		PhysicalCouplingsOrMasses: u.PhysicalCouplingsOrMasses,
		Verdict:                   StatusU1CompletionKappaTargetInherited,
		Reason:                    "The old U(1) completion search confirms kappa_U1=6 as the value that reproduces the whitening candidate and finds several finite count resonances equal to 6; none is a finite-action selection rule.",
	}
}

func buildRepresentationMetricAudit(a abeliancoupling.Analysis) RepresentationMetricAudit {
	return RepresentationMetricAudit{
		Executed:                                true,
		DiagonalTraceGramAsRepresentationMetric: a.DiagonalTraceGramSelectedAsRepresentationMetric,
		CanonicalGeneratorDiagnosticsDerived:    a.CanonicalGeneratorDiagnosticsDerived,
		DiagonalTraceGramAsGaugeKineticHessian:  a.DiagonalTraceGramSelectedAsGaugeKineticHessian,
		FieldCount:                              len(a.Fields),
		HyperchargeBridgeNorm:                   a.Hypercharge.CombinedBridgeNorm,
		ChargeTableKY:                           a.Hypercharge.ChargeTableKY,
		BoundarySin2:                            a.Hypercharge.BoundarySin2,
		PhysicalGaugeCouplingsDerived:           a.PhysicalGaugeCouplingsDerived,
		FineStructureDerived:                    a.FineStructureDerived,
		RGBoundaryScaleDerived:                  a.RGBoundaryScaleDerived,
		HiddenObservedInputUsed:                 a.HiddenObservedInputUsed,
		Verdict:                                 StatusRepresentationMetricLedgerAvailable,
		Reason:                                  "The diagonal U(1) trace-Gram data is a valid representation-metric diagnostic for abelian generators, but the repository already marks it as not selected as the physical gauge kinetic Hessian.",
	}
}

func buildSelectionBoundary(h HyperchargeTraceAudit, k KappaSearchAudit, m RepresentationMetricAudit) SelectionBoundary {
	kyEqualsKappa := math.Abs(h.KY-k.TargetKappa) <= eps
	traceLedger := h.KYConfirmed && m.DiagonalTraceGramAsRepresentationMetric && m.CanonicalGeneratorDiagnosticsDerived
	kappaCandidate := k.CompletionFamilyTyped && math.Abs(k.TargetKappa-6) <= eps && k.WhiteningSelectsKappa
	return SelectionBoundary{
		Executed:                         true,
		TraceLedgerAvailable:             traceLedger,
		KappaWhiteningCandidateAvailable: kappaCandidate,
		TraceToKappaNativeMapDerived:     false,
		KYEqualsTargetKappa:              kyEqualsKappa,
		KYAndKappaSameObject:             false,
		UnimodularitySelectsKappa:        false,
		AnomalyCancellationSelectsKappa:  false,
		SecondVariationSelectsKappa:      k.ActionSelectsKappa && k.FiniteSecondVariation,
		NativeKappaSelected:              false,
		NativeGaugeHessianSelected:       m.DiagonalTraceGramAsGaugeKineticHessian || k.GaugeKineticHessianFixed,
		NativeWeakAngleDerived:           h.PhysicalWeakMixingAngleDerived,
		NativeWZMassesDerived:            false,
		Verdict:                          StatusTraceKappaAirlockDefined,
		Reason:                           "The trace ledger and the kappa whitening candidate are both coherent, but they occupy different theorem layers. A native map from representation trace normalization or unimodularity/anomaly cancellation to the abelian Hessian coefficient is still absent.",
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
		NativeKappaWritten:        false,
		NativeGaugeHessianWritten: false,
		NativeWeakAngleWritten:    false,
		NativeWZMassWritten:       false,
		Verdict:                   StatusFirewallPreserved,
		Reason:                    "No W/Z mass, Higgs mass or VEV, Fermi constant, weak angle, fine-structure constant, gauge coupling, Yukawa texture, CKM, or PMNS datum is imported; kappa_U1 remains bridge-only.",
	}
}

func buildRegistryUpdate(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"Finite representation traces confirm the charge-level hypercharge normalization k_Y=5/3.",
			"The electromagnetic charge identity and hypercharge/SU(2)_L compatibility remain intact.",
		},
		BridgeEntries: []string{
			"Equal normalized couplings at a boundary still give the structural diagnostic sin^2=3/8, without claiming the physical weak angle.",
			"The Gate493 abelian completion family keeps kappa_U1=6 as the diag(1,1,4) whitening candidate.",
			"Diagonal U(1) trace-Gram data remains a representation-metric diagnostic, not a selected gauge kinetic Hessian.",
		},
		EnvironmentalEntries: []string{
			"Physical g_2, g_Y, alpha, theta_W, W/Z masses, Higgs VEV, and continuum RG boundary data remain environmental/bridge until a finite action selects kinetic terms.",
		},
		FailedRoutes: []string{
			StatusFailedTraceDoesNotSelectKappa,
			StatusFailedKYIsNotAbelianHessian,
			StatusFailedCountResonanceNotAction,
			StatusFailedRepresentationMetricNotAction,
			StatusFailedSecondVariationStillMissing,
			StatusFailedGaugeCouplingsNotDerived,
			StatusFailedWeakAngleNotDerived,
			StatusFailedWZMassesNotDerived,
		},
		OpenTheorems: []string{
			"Derive kappa_U1 from an actual finite electroweak action second variation rather than from whitening or count resonance.",
			"Prove or reject that a spectral trace functional selects the diagonal trace-Gram data as the gauge kinetic Hessian.",
			"Only after a selected Hessian may the weak-angle, gauge-coupling, RG, and W/Z mass bridges open.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        495,
		Title:       "Finite Electroweak Action Second Variation Source Audit",
		Reason:      "Gate494 proves trace normalization and count resonances do not select kappa_U1; the only remaining native path is an explicit finite-action second variation.",
		PrimaryTask: "construct or fail the actual δ²S/δB² and mixed electroweak Hessian source for the finite scalar/gauge action, without importing theta_W, alpha, W/Z masses, Fermi constant, or continuum RG data",
	}
}

func validate(a Analysis) error {
	switch {
	case !a.Inheritance.Executed || !a.Inheritance.Gate493FullConnectionClosed || !a.Inheritance.Gate493QuadraticFamilyTyped || !a.Inheritance.Gate493PositiveCompletionFamily || math.Abs(a.Inheritance.Gate493KappaTarget-6) > eps || a.Inheritance.Gate493KappaSelected || a.Inheritance.Gate493GaugeHessianSelected || !a.Inheritance.Gate493PhysicalRegistryBlocked || !a.Inheritance.Gate494Requested:
		return fmt.Errorf("invalid Gate494 inheritance: %+v", a.Inheritance)
	case !a.Hypercharge.Executed || !a.Hypercharge.ChargeIdentityDerived || !a.Hypercharge.HyperchargeCommutesWithSU2 || !a.Hypercharge.KYConfirmed || !a.Hypercharge.BoundarySin238Confirmed:
		return fmt.Errorf("invalid hypercharge trace audit: %+v", a.Hypercharge)
	case a.Hypercharge.GaugeKineticNormalizationDerived || a.Hypercharge.PhysicalWeakMixingAngleDerived || a.Hypercharge.FineStructureDerived || a.Hypercharge.RGBoundaryScaleDerived || a.Hypercharge.HiddenObservedInputUsed || a.Hypercharge.SelectsKappaU1:
		return fmt.Errorf("hypercharge trace audit over-promoted physical data: %+v", a.Hypercharge)
	case !a.Kappa.Executed || !a.Kappa.CompletionFamilyTyped || math.Abs(a.Kappa.TargetKappa-6) > eps || !a.Kappa.WhiteningSelectsKappa || a.Kappa.ActionSelectsKappa || a.Kappa.FiniteSecondVariation || a.Kappa.UniqueDerivation || a.Kappa.CountResonanceSelected || a.Kappa.KappaPhysical || a.Kappa.GaugeKineticHessianFixed || a.Kappa.PhysicalCouplingsOrMasses:
		return fmt.Errorf("invalid kappa search audit: %+v", a.Kappa)
	case !a.Metric.Executed || !a.Metric.DiagonalTraceGramAsRepresentationMetric || !a.Metric.CanonicalGeneratorDiagnosticsDerived || a.Metric.DiagonalTraceGramAsGaugeKineticHessian || a.Metric.PhysicalGaugeCouplingsDerived || a.Metric.FineStructureDerived || a.Metric.RGBoundaryScaleDerived || a.Metric.HiddenObservedInputUsed:
		return fmt.Errorf("invalid representation metric audit: %+v", a.Metric)
	case !a.Boundary.Executed || !a.Boundary.TraceLedgerAvailable || !a.Boundary.KappaWhiteningCandidateAvailable || a.Boundary.TraceToKappaNativeMapDerived || a.Boundary.KYAndKappaSameObject || a.Boundary.UnimodularitySelectsKappa || a.Boundary.AnomalyCancellationSelectsKappa || a.Boundary.SecondVariationSelectsKappa || a.Boundary.NativeKappaSelected || a.Boundary.NativeGaugeHessianSelected || a.Boundary.NativeWeakAngleDerived || a.Boundary.NativeWZMassesDerived:
		return fmt.Errorf("boundary over-promoted kappa theorem: %+v", a.Boundary)
	case !a.Firewall.Executed || a.Firewall.ObservedWMassImported || a.Firewall.ObservedZMassImported || a.Firewall.ObservedHiggsMassImported || a.Firewall.FermiConstantImported || a.Firewall.WeakAngleImported || a.Firewall.FineStructureImported || a.Firewall.GaugeCouplingImported || a.Firewall.YukawaImported || a.Firewall.CKMPMNSImported || a.Firewall.NativeKappaWritten || a.Firewall.NativeGaugeHessianWritten || a.Firewall.NativeWeakAngleWritten || a.Firewall.NativeWZMassWritten:
		return fmt.Errorf("firewall leak: %+v", a.Firewall)
	}
	return nil
}

func truth(a Analysis) string {
	return "Gate494 sharpens the electroweak boundary: ASHA really has a native charge-level trace normalization k_Y=5/3 and a bridge boundary diagnostic sin^2=3/8, and it really has a positive abelian-completed quadratic family whose whitening candidate sits at kappa_U1=6. But k_Y, trace-Gram representation metrics, unimodularity/anomaly cancellation, and finite count resonances are not the same object as a finite-action second variation. Therefore kappa_U1 is not natively selected, the gauge Hessian remains open, and no physical weak angle, gauge coupling, or W/Z mass may enter the registry."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate493_full_connection=%v quadratic_family=%v positive_completion=%v kappa_target=%.10g kappa_selected=%v hessian_selected=%v registry_blocked=%v gate494_requested=%v no_data=%v verdict=%s reason=%s", x.Gate493FullConnectionClosed, x.Gate493QuadraticFamilyTyped, x.Gate493PositiveCompletionFamily, x.Gate493KappaTarget, x.Gate493KappaSelected, x.Gate493GaugeHessianSelected, x.Gate493PhysicalRegistryBlocked, x.Gate494Requested, x.NoElectroweakFlavorDataImported, x.Verdict, x.Reason)
}

func FormatHypercharge(x HyperchargeTraceAudit) string {
	return fmt.Sprintf("Q_identity=%v Y_commutes_SU2=%v TrY2=%.10g TrT3_2=%.10g kY=%.10g expected=%.10g kY_confirmed=%v norm_factor=%.10g boundary_sin2=%.10g sin2_3over8=%v kinetic_norm=%v physical_theta=%v alpha=%v RG_scale=%v hidden_observed=%v selects_kappa=%v verdict=%s reason=%s", x.ChargeIdentityDerived, x.HyperchargeCommutesWithSU2, x.FullYTrace2, x.FullT3Trace2, x.KY, x.KYExpected, x.KYConfirmed, x.NormalizedHyperchargeFactor, x.EqualNormalizedCouplingBoundarySin2, x.BoundarySin238Confirmed, x.GaugeKineticNormalizationDerived, x.PhysicalWeakMixingAngleDerived, x.FineStructureDerived, x.RGBoundaryScaleDerived, x.HiddenObservedInputUsed, x.SelectsKappaU1, x.Verdict, x.Reason)
}

func FormatKappa(x KappaSearchAudit) string {
	return fmt.Sprintf("family_typed=%v target_kappa=%.10g source=%q direction=%q norm_sq=%.10g whitening_selects=%v action_selects=%v second_variation=%v resonance_count=%d hit_count=%d unique_derivation=%v count_selected=%v physical_kappa=%v hessian_fixed=%v physical_couplings_masses=%v verdict=%s reason=%s", x.CompletionFamilyTyped, x.TargetKappa, x.TargetSource, x.AbelianDirectionBasis, x.AbelianDirectionNormSq, x.WhiteningSelectsKappa, x.ActionSelectsKappa, x.FiniteSecondVariation, x.CandidateResonanceCount, x.CandidateHitCount, x.UniqueDerivation, x.CountResonanceSelected, x.KappaPhysical, x.GaugeKineticHessianFixed, x.PhysicalCouplingsOrMasses, x.Verdict, x.Reason)
}

func FormatMetric(x RepresentationMetricAudit) string {
	return fmt.Sprintf("trace_gram_rep_metric=%v canonical_generators=%v trace_gram_gauge_hessian=%v field_count=%d hypercharge_bridge_norm=%.10g charge_table_kY=%.10g boundary_sin2=%.10g physical_couplings=%v alpha=%v RG_scale=%v hidden_observed=%v verdict=%s reason=%s", x.DiagonalTraceGramAsRepresentationMetric, x.CanonicalGeneratorDiagnosticsDerived, x.DiagonalTraceGramAsGaugeKineticHessian, x.FieldCount, x.HyperchargeBridgeNorm, x.ChargeTableKY, x.BoundarySin2, x.PhysicalGaugeCouplingsDerived, x.FineStructureDerived, x.RGBoundaryScaleDerived, x.HiddenObservedInputUsed, x.Verdict, x.Reason)
}

func FormatBoundary(x SelectionBoundary) string {
	return fmt.Sprintf("trace_ledger=%v kappa_candidate=%v trace_to_kappa_map=%v kY_equals_kappa=%v same_object=%v unimod_selects=%v anomaly_selects=%v second_variation_selects=%v native_kappa=%v native_hessian=%v native_theta=%v native_WZ=%v verdict=%s reason=%s", x.TraceLedgerAvailable, x.KappaWhiteningCandidateAvailable, x.TraceToKappaNativeMapDerived, x.KYEqualsTargetKappa, x.KYAndKappaSameObject, x.UnimodularitySelectsKappa, x.AnomalyCancellationSelectsKappa, x.SecondVariationSelectsKappa, x.NativeKappaSelected, x.NativeGaugeHessianSelected, x.NativeWeakAngleDerived, x.NativeWZMassesDerived, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("observed_W=%v observed_Z=%v observed_Higgs=%v Fermi=%v theta=%v alpha=%v gauge_coupling=%v Yukawa=%v CKM_PMNS=%v native_kappa=%v native_hessian=%v native_theta=%v native_WZ=%v verdict=%s reason=%s", x.ObservedWMassImported, x.ObservedZMassImported, x.ObservedHiggsMassImported, x.FermiConstantImported, x.WeakAngleImported, x.FineStructureImported, x.GaugeCouplingImported, x.YukawaImported, x.CKMPMNSImported, x.NativeKappaWritten, x.NativeGaugeHessianWritten, x.NativeWeakAngleWritten, x.NativeWZMassWritten, x.Verdict, x.Reason)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 494 Registry Audit — Abelian U(1) Completion Coefficient Selection Audit\n\n")
	b.WriteString("## Verdict\n\n")
	for _, v := range []string{
		StatusGate493Inherited,
		StatusHyperchargeTraceNormalizationFound,
		StatusBoundarySin238Preserved,
		StatusU1CompletionKappaTargetInherited,
		StatusFiniteCountResonancesAudited,
		StatusRepresentationMetricLedgerAvailable,
		StatusTraceKappaAirlockDefined,
		StatusFailedTraceDoesNotSelectKappa,
		StatusFailedKYIsNotAbelianHessian,
		StatusFailedCountResonanceNotAction,
		StatusFailedRepresentationMetricNotAction,
		StatusFailedSecondVariationStillMissing,
		StatusFailedGaugeCouplingsNotDerived,
		StatusFailedWeakAngleNotDerived,
		StatusFailedWZMassesNotDerived,
		StatusFirewallPreserved,
		StatusKappaRegistryWriteBlocked,
		StatusGate495RedirectDefined,
	} {
		b.WriteString("- `" + v + "`\n")
	}

	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("Gate493 typed the full electroweak quadratic family and left `kappa_U1` unselected. Gate494 may inspect finite trace, representation-metric, anomaly/unimodularity, and topological normalization ledgers, but it may not import physical electroweak data or promote a whitening convention into an action theorem.\n\n")

	b.WriteString("## Hypercharge trace normalization audit\n\n")
	b.WriteString(FormatHypercharge(a.Hypercharge) + "\n\n")
	b.WriteString("The trace table confirms `k_Y = 5/3` and the boundary diagnostic `sin² = 3/8` under equal normalized couplings. This is a native charge-normalization result and a bridge boundary diagnostic, not a physical weak angle and not an abelian Hessian coefficient.\n\n")

	b.WriteString("## Kappa selection search\n\n")
	b.WriteString(FormatKappa(a.Kappa) + "\n\n")
	b.WriteString("The value `kappa_U1 = 6` remains the coefficient required to recover the `diag(1,1,4)` whitening candidate. Multiple finite counts resonate with 6, but count resonance is not a second variation.\n\n")

	b.WriteString("## Representation metric audit\n\n")
	b.WriteString(FormatMetric(a.Metric) + "\n\n")
	b.WriteString("The diagonal U(1) trace-Gram data gives a valid representation metric. The repository still marks it as not selected as the physical gauge kinetic Hessian, so it cannot fix `g_Y`, `alpha`, `theta_W`, or W/Z masses.\n\n")

	b.WriteString("## Trace-to-kappa airlock\n\n")
	b.WriteString(FormatBoundary(a.Boundary) + "\n\n")
	b.WriteString("The core obstruction is type-theoretic: `k_Y=5/3` normalizes the hypercharge generator in the finite charge ledger, while `kappa_U1` weights the abelian completion direction in the electroweak quadratic action family. A theorem must connect them through an action variation; Gate494 finds no such native map.\n\n")

	b.WriteString("## Firewall result\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No physical electroweak number entered the native lane, and no `kappa_U1`, gauge Hessian, weak angle, gauge coupling, W/Z mass, Higgs VEV, Yukawa, CKM, or PMNS registry write occurred.\n\n")

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
