// Package tauetaspatialtagging implements Gate 242:
// Scalar Fundamental Class (tau_eta) Spatial Tagging and Generation Breaking Audit.
//
// Gate 241 showed that the contact space K is available but that no native
// contact one-form/Reeb-vector/projection package exists to tag one of the
// three surviving pure-spatial weak-plane candidates. Gate 242 audits the next
// available finite datum with exactly three signed components: the scalar
// fundamental-class signature tau_eta=(2,-2,1) inherited from Gate 193.
//
// The result is deliberately split.  The tau_eta signature has exactly the
// right algebraic capacities: its magnitudes |2|,|2|,|1| would isolate one
// spatial axis and select the complementary two-plane, while its signed values
// 2,-2,1 form a 1+1+1 diagonal spectrum capable of breaking three generation
// labels.  However tau_eta is still a scalar-bundle trace functional, not yet a
// derived operator on the Fock generator carrier W or on the triality generation
// carrier.  Therefore this gate records real conditional support and real
// obstructions without promoting the weak plane or generation texture.
package tauetaspatialtagging

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/reebweakselection"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarfundamentalclass"
)

const (
	AuditID = "GATE242-TAU-ETA-SPATIAL-TAGGING-GENERATION-BREAKING-AUDIT"

	StatusTauEtaRetrieved              = "CONDITIONAL_SUPPORT_TAU_ETA_SEQUENCE_RETRIEVED"
	StatusTauEtaMagnitudeSelector      = "CONDITIONAL_SUPPORT_TAU_ETA_MAGNITUDE_2PLUS1_SELECTOR_CAPACITY"
	StatusTauEtaGenerationCapacity     = "CONDITIONAL_SUPPORT_TAU_ETA_GENERATION_BREAKING_CAPACITY"
	StatusFailedTauToSpatialPullback   = "FAILED_ROUTE_TAU_ETA_TO_FOCK_SPATIAL_PULLBACK"
	StatusFailedWeakPlaneSelection     = "FAILED_ROUTE_TAU_ETA_WEAK_PLANE_SELECTION"
	StatusFailedTauToTrialityPullback  = "FAILED_ROUTE_TAU_ETA_TO_TRIALITY_GENERATION_PULLBACK"
	StatusFailedGenerationTextureClaim = "FAILED_ROUTE_TAU_ETA_GENERATION_TEXTURE_DERIVATION"
	StatusGlobalHStillUnselected       = "FAILED_ROUTE_GLOBAL_H_SUMMAND_STILL_UNSELECTED"
)

type TauEtaSequenceAudit struct {
	SourceGate                      string
	SourceExpression                string
	Sequence                        []int
	Magnitudes                      []int
	StableNativeDegrees             bool
	ScalarTraceFunctionalOnly       bool
	ThreeComponentSignature         bool
	MagnitudePattern                string
	SignedPattern                   string
	TwoPlusOneMagnitudeSelector     bool
	OnePlusOnePlusOneSignedSpectrum bool
	Verdict                         string
}

type SpatialMappingAudit struct {
	SpatialAxes                []string
	CandidatePureSpatialPlanes []string
	DimensionCompatible        bool
	TauEtaActsOnScalarBundle   bool
	TauEtaActsOnFockW          bool
	NativePullbackDerived      bool
	ManualAxisAssignment       bool
	Magnitudes                 []int
	UniqueMagnitudeValue       int
	UniqueMagnitudeIndex       int
	UniqueAxisIfMapped         string
	ComplementPlaneIfMapped    string
	WeakPlaneConditionallySeen bool
	WeakPlaneDerived           bool
	S3DegeneracyWouldBreak     bool
	S3DegeneracyActuallyBroken bool
	Verdict                    string
}

type PlaneTauAudit struct {
	Plane                string
	ComplementAxis       string
	InheritedFromGate240 bool
	SurvivesU1Twist      bool
	SelectedIfTauMapped  bool
	SelectedNatively     bool
	SelectionReason      string
	Verdict              string
}

type GenerationBreakingAudit struct {
	TrialityCarrierDimension         int
	TauValues                        []int
	DistinctEigenvalueCount          int
	SignedSpectrumBreaksAllThree     bool
	ExactTrialityKnownTooSymmetric   bool
	TauEtaToGenerationPullback       bool
	CanonicalTrialityOperatorDerived bool
	MixingOperatorDerived            bool
	CKMPMNSDerived                   bool
	CapacitySupported                bool
	TextureDerived                   bool
	Verdict                          string
}

type WeakOutcomeAudit struct {
	InheritsGate241ReebFailure       bool
	TauMagnitudeCanSelectAxis        bool
	TauToSpatialPullbackDerived      bool
	UniqueWeakPlaneConditionallySeen bool
	UniqueWeakPlaneDerived           bool
	PhysicalLeftHandedDerived        bool
	GlobalHSummandDerived            bool
	OrderOneReady                    bool
	Verdict                          string
}

type FirewallAudit struct {
	ForcedTauToSpatialMap       bool
	ForcedAxisAssignment        bool
	ImportedSMWeakPlane         bool
	ImportedGenerationTexture   bool
	PromotedTraceToSpinorMatrix bool
	ClaimedPhysicalChirality    bool
	ClaimedGlobalH              bool
	ClaimedCKMPMNS              bool
	ClaimedMasses               bool
	FiniteCorePolluted          bool
	Verdict                     string
}

type Summary struct {
	TauEtaRetrieved            bool
	MagnitudeSelectorCapacity  bool
	SpatialPullbackDerived     bool
	WeakPlaneConditionallySeen bool
	WeakPlaneDerived           bool
	GenerationBreakingCapacity bool
	GenerationTextureDerived   bool
	GlobalHDerived             bool
	Status                     string
	NextGate                   string
	Comment                    string
}

type Analysis struct {
	PreviousReeb   reebweakselection.Analysis
	Fundamental    scalarfundamentalclass.Analysis
	TauEta         TauEtaSequenceAudit
	Spatial        SpatialMappingAudit
	Planes         []PlaneTauAudit
	Generation     GenerationBreakingAudit
	Weak           WeakOutcomeAudit
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
		prev, err := reebweakselection.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 241 predecessor: %w", err)
			return
		}
		sf, err := scalarfundamentalclass.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build scalar fundamental-class input: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev, sf)
	})
	return defaultA, defaultErr
}

func Build(prev reebweakselection.Analysis, sf scalarfundamentalclass.Analysis) (Analysis, error) {
	tau, err := retrieveTauEta(sf)
	if err != nil {
		return Analysis{}, err
	}
	spatial := auditSpatialMapping(prev, tau)
	planes := auditPlanes(prev, spatial)
	gen := auditGeneration(tau)
	weak := auditWeak(prev, tau, spatial)
	fw := auditFirewall()
	sum := summarize(tau, spatial, gen, weak)
	truth := buildTruth(tau, spatial, gen, weak)
	return Analysis{PreviousReeb: prev, Fundamental: sf, TauEta: tau, Spatial: spatial, Planes: planes, Generation: gen, Weak: weak, Firewall: fw, Summary: sum, TruthStatement: truth}, nil
}

func retrieveTauEta(sf scalarfundamentalclass.Analysis) (TauEtaSequenceAudit, error) {
	if !sf.Normalization.StableQuantizedInvariants {
		return TauEtaSequenceAudit{}, fmt.Errorf("Gate 242 requires stable Gate 193 tau_eta degrees")
	}
	vals := []int{
		roundInt(sf.Normalization.NeutralQNativeDegree),
		roundInt(sf.Normalization.NeutralZNativeDegree),
		roundInt(sf.Normalization.NeutralMixedNativeDegree),
	}
	expected := []int{2, -2, 1}
	for i := range expected {
		if vals[i] != expected[i] {
			return TauEtaSequenceAudit{}, fmt.Errorf("unexpected tau_eta sequence %v, expected %v", vals, expected)
		}
	}
	mags := []int{abs(vals[0]), abs(vals[1]), abs(vals[2])}
	distinctSigned := distinctIntCount(vals) == 3
	magCounts := countInts(mags)
	twoPlusOne := len(magCounts) == 2 && magCounts[2] == 2 && magCounts[1] == 1
	return TauEtaSequenceAudit{
		SourceGate:                      "Gate 193 scalar fundamental-class finite eta-graded trace",
		SourceExpression:                "(tau_eta(Q^TQ), tau_eta(Z^TZ), tau_eta(T3^T Y_phi))",
		Sequence:                        vals,
		Magnitudes:                      mags,
		StableNativeDegrees:             true,
		ScalarTraceFunctionalOnly:       true,
		ThreeComponentSignature:         len(vals) == 3,
		MagnitudePattern:                "|tau_eta| = (2,2,1) = 2+1",
		SignedPattern:                   "tau_eta = (2,-2,1) = 1+1+1 signed spectrum",
		TwoPlusOneMagnitudeSelector:     twoPlusOne,
		OnePlusOnePlusOneSignedSpectrum: distinctSigned,
		Verdict:                         "The scalar fundamental class supplies an exact three-component signed integer signature. Its magnitudes have selector capacity and its signs have generation-breaking capacity, but the datum is still a scalar-bundle trace functional until a pullback is derived.",
	}, nil
}

func auditSpatialMapping(prev reebweakselection.Analysis, tau TauEtaSequenceAudit) SpatialMappingAudit {
	axes := []string{"a†_1", "a†_2", "a†_3"}
	planes := append([]string(nil), prev.Sieve.InheritedGate240Candidates...)
	if len(planes) == 0 {
		planes = []string{"U={a†_1,a†_2}", "U={a†_1,a†_3}", "U={a†_2,a†_3}"}
	}
	uniqueIdx := uniqueMagnitudeIndex(tau.Magnitudes)
	uniqueAxis := ""
	complement := ""
	uniqueMag := 0
	if uniqueIdx >= 0 && uniqueIdx < len(axes) {
		uniqueAxis = axes[uniqueIdx]
		complement = complementPlane(uniqueIdx + 1)
		uniqueMag = tau.Magnitudes[uniqueIdx]
	}
	return SpatialMappingAudit{
		SpatialAxes:                axes,
		CandidatePureSpatialPlanes: planes,
		DimensionCompatible:        len(tau.Sequence) == len(axes),
		TauEtaActsOnScalarBundle:   true,
		TauEtaActsOnFockW:          false,
		NativePullbackDerived:      false,
		ManualAxisAssignment:       false,
		Magnitudes:                 append([]int(nil), tau.Magnitudes...),
		UniqueMagnitudeValue:       uniqueMag,
		UniqueMagnitudeIndex:       uniqueIdx + 1,
		UniqueAxisIfMapped:         uniqueAxis,
		ComplementPlaneIfMapped:    complement,
		WeakPlaneConditionallySeen: uniqueIdx >= 0 && complement != "",
		WeakPlaneDerived:           false,
		S3DegeneracyWouldBreak:     uniqueIdx >= 0,
		S3DegeneracyActuallyBroken: false,
		Verdict:                    "The 3-component tau_eta signature is dimension-compatible with the three spatial Fock modes and its magnitudes would isolate the |1| axis. However no native tau_eta -> W_spatial pullback is derived, so the weak-plane selection is conditional only.",
	}
}

func auditPlanes(prev reebweakselection.Analysis, spatial SpatialMappingAudit) []PlaneTauAudit {
	planes := make([]PlaneTauAudit, 0, len(spatial.CandidatePureSpatialPlanes))
	for _, p := range spatial.CandidatePureSpatialPlanes {
		comp := complementForPlane(p)
		selectedIf := spatial.ComplementPlaneIfMapped == p
		planes = append(planes, PlaneTauAudit{
			Plane:                p,
			ComplementAxis:       comp,
			InheritedFromGate240: true,
			SurvivesU1Twist:      true,
			SelectedIfTauMapped:  selectedIf,
			SelectedNatively:     false,
			SelectionReason:      fmt.Sprintf("conditional only: selected iff tau_eta magnitudes are lawfully pulled back to spatial axes and tag %s", spatial.UniqueAxisIfMapped),
			Verdict:              "Pure-spatial plane remains unselected natively; tau_eta provides only a conditional complement rule until a spatial pullback exists.",
		})
	}
	sort.Slice(planes, func(i, j int) bool { return planes[i].Plane < planes[j].Plane })
	_ = prev
	return planes
}

func auditGeneration(tau TauEtaSequenceAudit) GenerationBreakingAudit {
	return GenerationBreakingAudit{
		TrialityCarrierDimension:         3,
		TauValues:                        append([]int(nil), tau.Sequence...),
		DistinctEigenvalueCount:          distinctIntCount(tau.Sequence),
		SignedSpectrumBreaksAllThree:     distinctIntCount(tau.Sequence) == 3,
		ExactTrialityKnownTooSymmetric:   true,
		TauEtaToGenerationPullback:       false,
		CanonicalTrialityOperatorDerived: false,
		MixingOperatorDerived:            false,
		CKMPMNSDerived:                   false,
		CapacitySupported:                distinctIntCount(tau.Sequence) == 3,
		TextureDerived:                   false,
		Verdict:                          "As a hypothetical diagonal operator on the 3D generation carrier, tau_eta=(2,-2,1) would break the exact-triality 1+2 degeneracy into 1+1+1. Gate 242 records this capacity, but no tau_eta -> triality-generation pullback or non-commuting texture pair is derived.",
	}
}

func auditWeak(prev reebweakselection.Analysis, tau TauEtaSequenceAudit, spatial SpatialMappingAudit) WeakOutcomeAudit {
	return WeakOutcomeAudit{
		InheritsGate241ReebFailure:       !prev.Summary.UniqueWeakPlaneDerived && !prev.Summary.SpatialAxisTagged,
		TauMagnitudeCanSelectAxis:        tau.TwoPlusOneMagnitudeSelector,
		TauToSpatialPullbackDerived:      spatial.NativePullbackDerived,
		UniqueWeakPlaneConditionallySeen: spatial.WeakPlaneConditionallySeen,
		UniqueWeakPlaneDerived:           spatial.WeakPlaneDerived,
		PhysicalLeftHandedDerived:        false,
		GlobalHSummandDerived:            false,
		OrderOneReady:                    false,
		Verdict:                          "tau_eta is a stronger selector-shaped datum than the missing Reeb vector because its magnitudes already contain a 2+1 split. Still, without a native spatial pullback it does not select the physical weak plane, chirality, or global H summand.",
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		ForcedTauToSpatialMap:       false,
		ForcedAxisAssignment:        false,
		ImportedSMWeakPlane:         false,
		ImportedGenerationTexture:   false,
		PromotedTraceToSpinorMatrix: false,
		ClaimedPhysicalChirality:    false,
		ClaimedGlobalH:              false,
		ClaimedCKMPMNS:              false,
		ClaimedMasses:               false,
		FiniteCorePolluted:          false,
		Verdict:                     "The gate treats tau_eta as an exact finite scalar-bundle signature and refuses to promote it into a Fock/spinor/generation operator without a derived pullback.",
	}
}

func summarize(tau TauEtaSequenceAudit, spatial SpatialMappingAudit, gen GenerationBreakingAudit, weak WeakOutcomeAudit) Summary {
	return Summary{
		TauEtaRetrieved:            tau.StableNativeDegrees,
		MagnitudeSelectorCapacity:  tau.TwoPlusOneMagnitudeSelector,
		SpatialPullbackDerived:     spatial.NativePullbackDerived,
		WeakPlaneConditionallySeen: spatial.WeakPlaneConditionallySeen,
		WeakPlaneDerived:           spatial.WeakPlaneDerived,
		GenerationBreakingCapacity: gen.CapacitySupported,
		GenerationTextureDerived:   gen.TextureDerived,
		GlobalHDerived:             weak.GlobalHSummandDerived,
		Status:                     strings.Join([]string{StatusTauEtaRetrieved, StatusTauEtaMagnitudeSelector, StatusTauEtaGenerationCapacity, StatusFailedTauToSpatialPullback, StatusFailedWeakPlaneSelection, StatusFailedTauToTrialityPullback, StatusFailedGenerationTextureClaim, StatusGlobalHStillUnselected}, "; "),
		NextGate:                   "Gate 243 — tau_eta pullback functor / scalar fundamental class to Fock-generation operator audit",
		Comment:                    "tau_eta has exactly the right signature to act as both a weak-plane selector and a generation splitter, but both uses require a still-missing pullback from the scalar-bundle functional to the Fock spatial and triality generation carriers.",
	}
}

func buildTruth(tau TauEtaSequenceAudit, spatial SpatialMappingAudit, gen GenerationBreakingAudit, weak WeakOutcomeAudit) string {
	return fmt.Sprintf("Gate 242 retrieves the exact scalar fundamental-class signature tau_eta=%v. Its magnitudes %v have 2+1 selector capacity and would pick %s with complementary plane %s if a native tau_eta->W_spatial pullback existed. Its signs provide a distinct 1+1+1 spectrum and therefore generation-breaking capacity beyond exact triality. The binding obstruction is type-theoretic: tau_eta remains a scalar-bundle trace functional, not yet an operator on Fock spatial modes or the triality generation carrier. Weak-plane selection and generation texture remain conditional, not derived.", tau.Sequence, tau.Magnitudes, spatial.UniqueAxisIfMapped, spatial.ComplementPlaneIfMapped)
}

func roundInt(x float64) int { return int(math.Round(x)) }
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distinctIntCount(xs []int) int {
	m := map[int]bool{}
	for _, x := range xs {
		m[x] = true
	}
	return len(m)
}

func countInts(xs []int) map[int]int {
	m := map[int]int{}
	for _, x := range xs {
		m[x]++
	}
	return m
}

func uniqueMagnitudeIndex(mags []int) int {
	counts := countInts(mags)
	for i, v := range mags {
		if counts[v] == 1 {
			return i
		}
	}
	return -1
}

func complementPlane(axisIndexOneBased int) string {
	switch axisIndexOneBased {
	case 1:
		return "U={a†_2,a†_3}"
	case 2:
		return "U={a†_1,a†_3}"
	case 3:
		return "U={a†_1,a†_2}"
	default:
		return ""
	}
}

func complementForPlane(plane string) string {
	switch plane {
	case "U={a†_1,a†_2}":
		return "a†_3"
	case "U={a†_1,a†_3}":
		return "a†_2"
	case "U={a†_2,a†_3}":
		return "a†_1"
	default:
		return ""
	}
}
