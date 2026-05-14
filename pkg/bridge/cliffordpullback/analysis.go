// Package cliffordpullback implements Gate 243:
// Clifford Action Pullback / tau_eta Endomorphism Audit.
//
// Gate 242 showed that the scalar fundamental-class signature
// tau_eta=(2,-2,1) has exactly the right selector capacities for both the
// remaining weak-plane degeneracy and the exact-triality generation problem,
// but that tau_eta is still only a scalar-bundle trace functional.  Gate 243
// audits the natural type-changing candidate: Clifford multiplication
// c: Lambda*(W) -> End(S_C).  The result is intentionally strict.  Clifford
// action is a real native map for exterior algebra elements; however tau_eta is
// not currently represented as an exterior form, a basis-blade coefficient
// vector, or an index class with a derived spinor pullback.  Therefore no
// endomorphism on S_C is constructed and no weak plane or generation texture is
// promoted.
package cliffordpullback

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/complexifiedhilbertspace"
	"github.com/bagherbal/asha-engine/pkg/bridge/tauetaspatialtagging"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

const (
	AuditID = "GATE243-CLIFFORD-ACTION-PULLBACK-TAU-ETA-ENDOMORPHISM-AUDIT"

	StatusCliffordActionAvailable      = "CONDITIONAL_SUPPORT_CLIFFORD_ACTION_MAP_AVAILABLE"
	StatusTauEtaSelectorCapacity       = "CONDITIONAL_SUPPORT_TAU_ETA_SELECTOR_CAPACITY_INHERITED"
	StatusFailedTauEtaInCliffordDomain = "FAILED_ROUTE_TAU_ETA_NOT_IN_CLIFFORD_ACTION_DOMAIN"
	StatusFailedTauEtaEndomorphism     = "FAILED_ROUTE_TAU_ETA_ENDOMORPHISM_CONSTRUCTION"
	StatusFailedSpatialSieve           = "FAILED_ROUTE_CLIFFORD_PULLBACK_WEAK_PLANE_SELECTION"
	StatusFailedTrialitySieve          = "FAILED_ROUTE_CLIFFORD_PULLBACK_GENERATION_TEXTURE"
	StatusFailedPullbackFunctor        = "FAILED_ROUTE_SCALAR_TRACE_TO_SPINOR_PULLBACK_FUNCTOR"
	StatusGlobalHStillUnselected       = "FAILED_ROUTE_GLOBAL_H_SUMMAND_STILL_UNSELECTED"
)

type CliffordActionAudit struct {
	SourceCarrier                   string
	Domain                          string
	Codomain                        string
	RealSpinorDimension             int
	ComplexSpinorDimension          int
	ExteriorBasisDimension          int
	CreationAnnihilationModel       bool
	CliffordMultiplicationAvailable bool
	RequiresExteriorElement         bool
	MapsScalarTraceFunctional       bool
	Verdict                         string
}

type TauEtaPullbackAudit struct {
	Sequence                     []int
	Magnitudes                   []int
	SourceType                   string
	ScalarTraceFunctional        bool
	ExteriorFormElement          bool
	HomogeneousGradeKnown        bool
	BasisBladeCoefficientsKnown  bool
	SpatialSlotLabelsDerived     bool
	TrialitySlotLabelsDerived    bool
	CliffordActionApplicable     bool
	AtiyahSingerIndexMapDerived  bool
	EndomorphismConstructed      bool
	HypotheticalSpatialOperator  string
	HypotheticalOperatorRejected bool
	RejectionReason              string
	Verdict                      string
}

type PullbackFunctorAudit struct {
	CliffordActionFunctorAvailable       bool
	TauEtaInFunctorDomain                bool
	TauEtaToExteriorFormDerived          bool
	TauEtaToIndexClassDerived            bool
	ScalarBundleToSpinorBundleMapDerived bool
	GaugeProjectionMapDerived            bool
	CanonicalNormalizationDerived        bool
	PullbackFunctorDerived               bool
	Verdict                              string
}

type SpatialEndomorphismSieve struct {
	SpatialAxes                  []string
	CandidatePureSpatialPlanes   []string
	TauMagnitudeSelectorCapacity bool
	EndomorphismAvailable        bool
	ProjectedToSpatialModes      bool
	MatrixSpectrumAvailable      bool
	UniqueAxisConditionallySeen  string
	ComplementPlaneConditionally string
	UniqueAxisDerived            string
	WeakPlaneDerived             bool
	S3DegeneracyBroken           bool
	Verdict                      string
}

type TrialityEndomorphismSieve struct {
	GenerationCarrierDimension     int
	TauSignedSpectrum              []int
	DistinctEigenvalueCapacity     bool
	EndomorphismAvailable          bool
	ProjectedToTrialitySectors     bool
	DiagonalGenerationOperator     bool
	GenerationTextureDerived       bool
	NonCommutingTexturePairDerived bool
	CKMPMNSDerived                 bool
	Verdict                        string
}

type FirewallAudit struct {
	ForcedTauAsExteriorForm      bool
	ForcedSpatialSlotMap         bool
	ForcedTrialitySlotMap        bool
	InventedCliffordEndomorphism bool
	ImportedWeakPlane            bool
	ImportedGenerationTexture    bool
	ClaimedGlobalH               bool
	ClaimedPhysicalChirality     bool
	ClaimedCKMPMNS               bool
	ClaimedMasses                bool
	FiniteCorePolluted           bool
	Verdict                      string
}

type Summary struct {
	CliffordActionAvailable    bool
	TauEtaSelectorCapacity     bool
	TauEtaInCliffordDomain     bool
	EndomorphismConstructed    bool
	WeakPlaneConditionallySeen bool
	WeakPlaneDerived           bool
	GenerationBreakingCapacity bool
	GenerationTextureDerived   bool
	PullbackFunctorDerived     bool
	GlobalHDerived             bool
	Status                     string
	NextGate                   string
	Comment                    string
}

type Analysis struct {
	PreviousTau    tauetaspatialtagging.Analysis
	ComplexCarrier complexifiedhilbertspace.Analysis
	CliffordAction CliffordActionAudit
	TauEtaPullback TauEtaPullbackAudit
	Functor        PullbackFunctorAudit
	Spatial        SpatialEndomorphismSieve
	Triality       TrialityEndomorphismSieve
	Firewall       FirewallAudit
	Summary        Summary
	TruthStatement string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := tauetaspatialtagging.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 242 predecessor: %w", err)
			return
		}
		carrier, err := complexifiedhilbertspace.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 235 complexified carrier: %w", err)
			return
		}
		f, err := spinor.NewCovariantPhaseFockSpace(4)
		if err != nil {
			defaultErr = fmt.Errorf("construct native Fock space: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev, carrier, f)
	})
	return defaultA, defaultErr
}

func Build(prev tauetaspatialtagging.Analysis, carrier complexifiedhilbertspace.Analysis, f spinor.FockSpace) (Analysis, error) {
	if f.ModeCount() != 4 || f.StateCount() != 16 {
		return Analysis{}, fmt.Errorf("Gate 243 requires native four-mode 16-state Fock carrier, got modes=%d states=%d", f.ModeCount(), f.StateCount())
	}
	cl := auditCliffordAction(carrier, f)
	tau := auditTauEtaPullback(prev)
	fun := auditFunctor(cl, tau)
	sp := auditSpatial(prev, tau, fun)
	tr := auditTriality(prev, tau, fun)
	fw := auditFirewall()
	sum := summarize(cl, tau, fun, sp, tr)
	truth := buildTruth(cl, tau, fun, sp, tr)
	return Analysis{PreviousTau: prev, ComplexCarrier: carrier, CliffordAction: cl, TauEtaPullback: tau, Functor: fun, Spatial: sp, Triality: tr, Firewall: fw, Summary: sum, TruthStatement: truth}, nil
}

func auditCliffordAction(carrier complexifiedhilbertspace.Analysis, f spinor.FockSpace) CliffordActionAudit {
	return CliffordActionAudit{
		SourceCarrier:                   "S_C = S ⊗_R C = Λ*(W) over C",
		Domain:                          "Λ*(W) exterior/Clifford elements with explicit grade and basis-blade coefficients",
		Codomain:                        "End_C(S_C)",
		RealSpinorDimension:             carrier.Complexification.RealDimensionAfter,
		ComplexSpinorDimension:          carrier.Complexification.ComplexDimensionAfter,
		ExteriorBasisDimension:          f.StateCount(),
		CreationAnnihilationModel:       true,
		CliffordMultiplicationAvailable: true,
		RequiresExteriorElement:         true,
		MapsScalarTraceFunctional:       false,
		Verdict:                         "The native Clifford/Fock machinery supplies c:Λ*(W)->End(S_C) for actual exterior algebra elements. It does not automatically act on scalar trace functionals that have not been represented as forms or index classes.",
	}
}

func auditTauEtaPullback(prev tauetaspatialtagging.Analysis) TauEtaPullbackAudit {
	return TauEtaPullbackAudit{
		Sequence:                     append([]int(nil), prev.TauEta.Sequence...),
		Magnitudes:                   append([]int(nil), prev.TauEta.Magnitudes...),
		SourceType:                   "finite scalar-bundle eta-graded trace functional: (tau_eta(Q^TQ), tau_eta(Z^TZ), tau_eta(T3^T Y_phi))",
		ScalarTraceFunctional:        true,
		ExteriorFormElement:          false,
		HomogeneousGradeKnown:        false,
		BasisBladeCoefficientsKnown:  false,
		SpatialSlotLabelsDerived:     false,
		TrialitySlotLabelsDerived:    false,
		CliffordActionApplicable:     false,
		AtiyahSingerIndexMapDerived:  false,
		EndomorphismConstructed:      false,
		HypotheticalSpatialOperator:  "T_tau? = 2 N_1 - 2 N_2 + 1 N_3 or diag(2,-2,1) on chosen slots",
		HypotheticalOperatorRejected: true,
		RejectionReason:              "This would choose the tau_eta components as spatial/triality basis labels by hand. No finite theorem identifies the scalar trace slots with Fock modes, basis blades, or generation sectors.",
		Verdict:                      "tau_eta has the right three-component signature, but it is not in the Clifford-action domain. Without a tau_eta->form or tau_eta->index-class theorem, no spinor endomorphism is constructed.",
	}
}

func auditFunctor(cl CliffordActionAudit, tau TauEtaPullbackAudit) PullbackFunctorAudit {
	return PullbackFunctorAudit{
		CliffordActionFunctorAvailable:       cl.CliffordMultiplicationAvailable,
		TauEtaInFunctorDomain:                tau.ExteriorFormElement || tau.BasisBladeCoefficientsKnown,
		TauEtaToExteriorFormDerived:          false,
		TauEtaToIndexClassDerived:            false,
		ScalarBundleToSpinorBundleMapDerived: false,
		GaugeProjectionMapDerived:            false,
		CanonicalNormalizationDerived:        false,
		PullbackFunctorDerived:               false,
		Verdict:                              "The Clifford action functor exists, but tau_eta is not an object in its current domain. The missing theorem is a scalar-bundle/index-class to spinor-endomorphism pullback with slot labels and normalization.",
	}
}

func auditSpatial(prev tauetaspatialtagging.Analysis, tau TauEtaPullbackAudit, fun PullbackFunctorAudit) SpatialEndomorphismSieve {
	return SpatialEndomorphismSieve{
		SpatialAxes:                  []string{"a†_1", "a†_2", "a†_3"},
		CandidatePureSpatialPlanes:   append([]string(nil), prev.Spatial.CandidatePureSpatialPlanes...),
		TauMagnitudeSelectorCapacity: prev.Spatial.WeakPlaneConditionallySeen,
		EndomorphismAvailable:        tau.EndomorphismConstructed && fun.PullbackFunctorDerived,
		ProjectedToSpatialModes:      false,
		MatrixSpectrumAvailable:      false,
		UniqueAxisConditionallySeen:  prev.Spatial.UniqueAxisIfMapped,
		ComplementPlaneConditionally: prev.Spatial.ComplementPlaneIfMapped,
		UniqueAxisDerived:            "",
		WeakPlaneDerived:             false,
		S3DegeneracyBroken:           false,
		Verdict:                      "If tau_eta were lawfully pulled back, |tau_eta|=(2,2,1) would tag a†_3 and conditionally select U={a†_1,a†_2}. Gate 243 does not derive the endomorphism, so the spatial S3 degeneracy is not broken natively.",
	}
}

func auditTriality(prev tauetaspatialtagging.Analysis, tau TauEtaPullbackAudit, fun PullbackFunctorAudit) TrialityEndomorphismSieve {
	return TrialityEndomorphismSieve{
		GenerationCarrierDimension:     prev.Generation.TrialityCarrierDimension,
		TauSignedSpectrum:              append([]int(nil), tau.Sequence...),
		DistinctEigenvalueCapacity:     prev.Generation.CapacitySupported,
		EndomorphismAvailable:          tau.EndomorphismConstructed && fun.PullbackFunctorDerived,
		ProjectedToTrialitySectors:     false,
		DiagonalGenerationOperator:     false,
		GenerationTextureDerived:       false,
		NonCommutingTexturePairDerived: false,
		CKMPMNSDerived:                 false,
		Verdict:                        "The signed spectrum (2,-2,1) would be a 1+1+1 diagonal generation splitter if a triality pullback existed. No such endomorphism or non-commuting texture pair is derived.",
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		ForcedTauAsExteriorForm:      false,
		ForcedSpatialSlotMap:         false,
		ForcedTrialitySlotMap:        false,
		InventedCliffordEndomorphism: false,
		ImportedWeakPlane:            false,
		ImportedGenerationTexture:    false,
		ClaimedGlobalH:               false,
		ClaimedPhysicalChirality:     false,
		ClaimedCKMPMNS:               false,
		ClaimedMasses:                false,
		FiniteCorePolluted:           false,
		Verdict:                      "Gate 243 refuses to turn a scalar trace signature into a spinor/Fock/generation matrix without a derived form, index map, or pullback functor.",
	}
}

func summarize(cl CliffordActionAudit, tau TauEtaPullbackAudit, fun PullbackFunctorAudit, sp SpatialEndomorphismSieve, tr TrialityEndomorphismSieve) Summary {
	return Summary{
		CliffordActionAvailable:    cl.CliffordMultiplicationAvailable,
		TauEtaSelectorCapacity:     len(tau.Sequence) == 3 && len(tau.Magnitudes) == 3,
		TauEtaInCliffordDomain:     fun.TauEtaInFunctorDomain,
		EndomorphismConstructed:    tau.EndomorphismConstructed,
		WeakPlaneConditionallySeen: sp.TauMagnitudeSelectorCapacity,
		WeakPlaneDerived:           sp.WeakPlaneDerived,
		GenerationBreakingCapacity: tr.DistinctEigenvalueCapacity,
		GenerationTextureDerived:   tr.GenerationTextureDerived,
		PullbackFunctorDerived:     fun.PullbackFunctorDerived,
		GlobalHDerived:             false,
		Status: strings.Join([]string{
			StatusCliffordActionAvailable,
			StatusTauEtaSelectorCapacity,
			StatusFailedTauEtaInCliffordDomain,
			StatusFailedTauEtaEndomorphism,
			StatusFailedSpatialSieve,
			StatusFailedTrialitySieve,
			StatusFailedPullbackFunctor,
			StatusGlobalHStillUnselected,
		}, "; "),
		NextGate: "Gate 244 — scalar fundamental-class carrier theorem / tau_eta slot-label derivation audit",
		Comment:  "The Clifford action exists for exterior forms, but tau_eta is a scalar trace functional. The missing bridge is not Clifford multiplication itself; it is the theorem that represents tau_eta as a form, index class, or carrier-labelled operator.",
	}
}

func buildTruth(cl CliffordActionAudit, tau TauEtaPullbackAudit, fun PullbackFunctorAudit, sp SpatialEndomorphismSieve, tr TrialityEndomorphismSieve) string {
	return fmt.Sprintf("Gate 243 confirms the native Clifford action %s -> %s on the complexified spinor carrier, but tau_eta=%v remains a %s rather than an exterior form or index class. Therefore no End(S_C) matrix is constructed. The conditional weak-plane selector %s and generation spectrum %v remain visible as capacities, but the pullback functor is not derived.", cl.Domain, cl.Codomain, tau.Sequence, tau.SourceType, sp.ComplementPlaneConditionally, tr.TauSignedSpectrum)
}
