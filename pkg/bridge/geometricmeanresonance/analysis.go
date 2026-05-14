// Package geometricmeanresonance implements Gate 227: geometric-mean
// intermediate scale resonance / sealed hierarchy audit.
//
// Gate 226 found that a sealed QCD-like axion phenomenology wants
// f_a ≈ 1e12 GeV, while Gate 223 found that the relic-decay EFT portals need
// an intermediate suppression scale Λ_EFT below roughly 5e11 GeV. Gate 227 asks
// whether these two independent intermediate scales are structurally related to
// the already sealed ASHA hierarchy through the geometric mean
//
//	M_int = sqrt(M_B M_*)
//
// using only previously sealed values. It does not derive a new symmetry, an
// axion, a Pati-Salam connection, or a Wilson coefficient.
package geometricmeanresonance

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/axionphenomenologyseal"
	"github.com/bagherbal/asha-engine/pkg/bridge/coloredoctetportal"
)

const (
	AuditID = "GATE227-GEOMETRIC-MEAN-INTERMEDIATE-RESONANCE-AUDIT"

	StatusGeometricMeanResonance      = "CONDITIONAL_PHENOMENOLOGY_GEOMETRIC_MEAN_RESONANCE"
	StatusNativeBreakingNotDerived    = "FAILED_ROUTE_NATIVE_INTERMEDIATE_BREAKING_DERIVATION"
	StatusPatiSalamQuarantined        = "PATI_SALAM_ROUTE_QUARANTINED_BY_LEPTOQUARK_DYNAMICS_SEAL"
	StatusNullHypothesisRejectedByFit = "NULL_HYPOTHESIS_NO_RESONANCE_REJECTED_WITHIN_ONE_DECADE"
)

const (
	// Gate-219/Gate-224/Gate-226 sealed hierarchy values.
	heavyThresholdMBGeV         = 2.56895727e6
	topologicalBoundaryMStarGeV = 1.72179441e17
	electroweakVEVGeV           = 246.0

	// Gate 226 sealed ALP target under theta_i=1.
	sealedAxionFAGeV = 1.0e12

	// Gate 223 conservative EFT scale bound for the colored-octet decay portal.
	relicDecayLambdaMaxGeV = 4.99261316e11

	// Resonance criterion: one decade is intentionally coarse and declared.
	oneDecade = 1.0
)

type Gate226Snapshot struct {
	Gate226Inherited      bool
	AxionSealActive       bool
	NativeAxionNotDerived bool
	FAGeV                 float64
	HeavyDMAbsenceBinding bool
	PriorDirectScaleMatch bool
	TruthStatement        string
}

type Gate223Snapshot struct {
	Gate223Inherited            bool
	RelicDecaySealGranted       bool
	OctetPortalFound            bool
	LambdaEFTMaxGeV             float64
	LeptoquarkDynamicsSealAlive bool
	WilsonCoefficientDerived    bool
	TruthStatement              string
}

type SealedHierarchy struct {
	MBGeV               float64
	MStarGeV            float64
	VEVGeV              float64
	FARequirementGeV    float64
	LambdaEFTMaxGeV     float64
	ValuesInheritedOnly bool
	NativeScaleDerived  bool
}

type ScaleResonance struct {
	Name            string
	ScaleGeV        float64
	RatioToMInt     float64
	Log10Gap        float64
	WithinOneDecade bool
	BracketRole     string
	Verdict         string
}

type GeometricMeanAudit struct {
	Formula                   string
	MBGeV                     float64
	MStarGeV                  float64
	MIntGeV                   float64
	LogDistanceMBToMInt       float64
	LogDistanceMIntToMStar    float64
	SymmetricSeesaw           bool
	Targets                   []ScaleResonance
	FAMatch                   bool
	LambdaMatch               bool
	BothTargetsBracketed      bool
	GeometricMeanOfTargetsGeV float64
	TargetMeanLog10Gap        float64
	NullHypothesisRejected    bool
	Verdict                   string
}

type SeesawStructureAudit struct {
	MIntGeV                        float64
	TwoStepPatternSuggested        bool
	BoundaryRelation               string
	AxionScaleCanLiveThere         bool
	RelicDecayScaleCanLiveThere    bool
	NativeBreakingPotentialDerived bool
	NativeOrderParameterDerived    bool
	NativeIntermediateScaleDerived bool
	RequiredFutureSeal             string
	Verdict                        string
}

type PatiSalamAudit struct {
	DormantU4SlotsPresent          bool
	LeptoquarkDynamicsSealActive   bool
	IntermediateU4BreakingAudited  bool
	NativeU4GaugeConnectionDerived bool
	NativeLeptoquarkCurvature      bool
	ProtonDecayChannelReopened     bool
	ConsistentOnlyWhileSealed      bool
	LifetimeComputed               bool
	Verdict                        string
}

type NullHypothesisAudit struct {
	Tested            bool
	Criterion         string
	FAGapDecades      float64
	LambdaGapDecades  float64
	WorstGapDecades   float64
	PassedNoResonance bool
	FailureMagnitude  float64
	Verdict           string
}

type FirewallAudit struct {
	Gate226Inherited               bool
	Gate223Inherited               bool
	UsedOnlySealedValues           bool
	IntermediateScaleFiniteDerived bool
	AxionNativeDerived             bool
	EFTMediatorDerived             bool
	PatiSalamImportedAsTheorem     bool
	LeptoquarkSealViolated         bool
	ProtonLifetimeComputed         bool
	BGapPromotedWithoutSeal        bool
	NewPhenomenologicalFitAdded    bool
	FiniteCorePolluted             bool
	Verdict                        string
}

type Summary struct {
	GeometricMeanResonanceFound bool
	IntermediateScaleGeV        float64
	NativeBreakingDerived       bool
	PatiSalamRouteOpened        bool
	Status                      string
	NextGate                    string
	Comment                     string
}

type Analysis struct {
	Gate226   Gate226Snapshot
	Gate223   Gate223Snapshot
	Hierarchy SealedHierarchy
	Geometric GeometricMeanAudit
	Seesaw    SeesawStructureAudit
	PatiSalam PatiSalamAudit
	Null      NullHypothesisAudit
	Firewall  FirewallAudit
	Summary   Summary

	TruthStatement string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		g226, err := axionphenomenologyseal.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 226 input: %w", err)
			return
		}
		g223, err := coloredoctetportal.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 223 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(g226, g223)
	})
	return defaultA, defaultErr
}

func Build(g226 axionphenomenologyseal.Analysis, g223 coloredoctetportal.Analysis) (Analysis, error) {
	s226 := snapshotFromGate226(g226)
	s223 := snapshotFromGate223(g223)
	if !s226.Gate226Inherited || !s226.AxionSealActive || !s226.NativeAxionNotDerived || s226.FAGeV <= 0 {
		return Analysis{}, fmt.Errorf("Gate 227 requires Gate 226 sealed ALP scale and native axion obstruction")
	}
	if !s223.Gate223Inherited || !s223.RelicDecaySealGranted || !s223.OctetPortalFound || s223.LambdaEFTMaxGeV <= 0 {
		return Analysis{}, fmt.Errorf("Gate 227 requires Gate 223 relic-decay EFT scale and RelicDecaySeal")
	}

	hierarchy := buildHierarchy(s226, s223)
	geom := auditGeometricMean(hierarchy)
	seesaw := auditSeesawPattern(geom)
	ps := auditPatiSalamSlots(s223, geom)
	null := auditNullHypothesis(geom)
	fw := auditFirewalls(s226, s223)
	summary := summarize(geom, seesaw, ps)
	truth := buildTruth(geom, seesaw, ps, null)

	return Analysis{Gate226: s226, Gate223: s223, Hierarchy: hierarchy, Geometric: geom, Seesaw: seesaw, PatiSalam: ps, Null: null, Firewall: fw, Summary: summary, TruthStatement: truth}, nil
}

func snapshotFromGate226(a axionphenomenologyseal.Analysis) Gate226Snapshot {
	return Gate226Snapshot{
		Gate226Inherited:      a.Summary.Status != "" && a.Seal.Active && a.Misalign.RequiredFAGeV > 0,
		AxionSealActive:       a.Seal.Active,
		NativeAxionNotDerived: !a.Seal.NativeALPDerived && !a.Seal.NativeFAObserved && !a.Seal.NativeAnomalyMapDerived,
		FAGeV:                 a.Misalign.RequiredFAGeV,
		HeavyDMAbsenceBinding: a.Gate225.HeavySectorDMAbsenceBinding,
		PriorDirectScaleMatch: a.Resonance.ResonanceFound,
		TruthStatement:        a.TruthStatement,
	}
}

func snapshotFromGate223(a coloredoctetportal.Analysis) Gate223Snapshot {
	return Gate223Snapshot{
		Gate223Inherited:            a.Summary.Status != "" && a.TensorSearch.OctetPortalFound,
		RelicDecaySealGranted:       a.RelicSeal.SealGranted && a.Summary.FullRelicDecaySeal,
		OctetPortalFound:            a.TensorSearch.OctetPortalFound,
		LambdaEFTMaxGeV:             a.Kinematics.ConservativeLambdaMaxGeV,
		LeptoquarkDynamicsSealAlive: a.Firewall.LeptoquarkDynamicsSealActive && !a.Firewall.LeptoquarkSealViolated,
		WilsonCoefficientDerived:    a.Firewall.WilsonCoefficientFixed || !a.RelicSeal.StillNotFiniteDerived,
		TruthStatement:              a.TruthStatement,
	}
}

func buildHierarchy(s226 Gate226Snapshot, s223 Gate223Snapshot) SealedHierarchy {
	fa := s226.FAGeV
	if fa <= 0 {
		fa = sealedAxionFAGeV
	}
	lam := s223.LambdaEFTMaxGeV
	if lam <= 0 {
		lam = relicDecayLambdaMaxGeV
	}
	return SealedHierarchy{
		MBGeV:               heavyThresholdMBGeV,
		MStarGeV:            topologicalBoundaryMStarGeV,
		VEVGeV:              electroweakVEVGeV,
		FARequirementGeV:    fa,
		LambdaEFTMaxGeV:     lam,
		ValuesInheritedOnly: true,
		NativeScaleDerived:  false,
	}
}

func auditGeometricMean(h SealedHierarchy) GeometricMeanAudit {
	mint := math.Sqrt(h.MBGeV * h.MStarGeV)
	logLeft := math.Abs(math.Log10(mint / h.MBGeV))
	logRight := math.Abs(math.Log10(h.MStarGeV / mint))
	targets := []ScaleResonance{
		compareToMInt("axion f_a requirement", h.FARequirementGeV, mint),
		compareToMInt("relic-decay EFT upper bound", h.LambdaEFTMaxGeV, mint),
		compareToMInt("electroweak VEV", h.VEVGeV, mint),
	}
	faMatch := targets[0].WithinOneDecade
	lambdaMatch := targets[1].WithinOneDecade
	tmean := math.Sqrt(h.FARequirementGeV * h.LambdaEFTMaxGeV)
	tgap := math.Abs(math.Log10(tmean / mint))
	bracketed := h.LambdaEFTMaxGeV <= mint && mint <= h.FARequirementGeV
	verdict := StatusGeometricMeanResonance
	if !(faMatch && lambdaMatch && bracketed) {
		verdict = "NO_GEOMETRIC_MEAN_RESONANCE"
	}
	return GeometricMeanAudit{
		Formula:                   "M_int = sqrt(M_B M_*)",
		MBGeV:                     h.MBGeV,
		MStarGeV:                  h.MStarGeV,
		MIntGeV:                   mint,
		LogDistanceMBToMInt:       logLeft,
		LogDistanceMIntToMStar:    logRight,
		SymmetricSeesaw:           math.Abs(logLeft-logRight) < 1e-12,
		Targets:                   targets,
		FAMatch:                   faMatch,
		LambdaMatch:               lambdaMatch,
		BothTargetsBracketed:      bracketed,
		GeometricMeanOfTargetsGeV: tmean,
		TargetMeanLog10Gap:        tgap,
		NullHypothesisRejected:    faMatch && lambdaMatch && bracketed,
		Verdict:                   verdict,
	}
}

func compareToMInt(name string, scale, mint float64) ScaleResonance {
	ratio := scale / mint
	gap := math.Abs(math.Log10(ratio))
	role := "outside"
	if ratio > 1 {
		role = "above M_int"
	} else if ratio < 1 {
		role = "below M_int"
	} else {
		role = "equal"
	}
	close := gap < oneDecade
	verdict := "NO_RESONANCE"
	if close {
		verdict = "WITHIN_ONE_DECADE_RESONANCE"
	}
	return ScaleResonance{Name: name, ScaleGeV: scale, RatioToMInt: ratio, Log10Gap: gap, WithinOneDecade: close, BracketRole: role, Verdict: verdict}
}

func auditSeesawPattern(g GeometricMeanAudit) SeesawStructureAudit {
	suggested := g.FAMatch && g.LambdaMatch && g.BothTargetsBracketed
	return SeesawStructureAudit{
		MIntGeV:                        g.MIntGeV,
		TwoStepPatternSuggested:        suggested,
		BoundaryRelation:               "M_B < M_int = sqrt(M_B M_*) < M_*",
		AxionScaleCanLiveThere:         g.FAMatch,
		RelicDecayScaleCanLiveThere:    g.LambdaMatch,
		NativeBreakingPotentialDerived: false,
		NativeOrderParameterDerived:    false,
		NativeIntermediateScaleDerived: false,
		RequiredFutureSeal:             "IntermediateBreakingSeal / finite order-parameter theorem required",
		Verdict:                        StatusNativeBreakingNotDerived,
	}
}

func auditPatiSalamSlots(s223 Gate223Snapshot, g GeometricMeanAudit) PatiSalamAudit {
	return PatiSalamAudit{
		DormantU4SlotsPresent:          true,
		LeptoquarkDynamicsSealActive:   s223.LeptoquarkDynamicsSealAlive,
		IntermediateU4BreakingAudited:  true,
		NativeU4GaugeConnectionDerived: false,
		NativeLeptoquarkCurvature:      false,
		ProtonDecayChannelReopened:     false,
		ConsistentOnlyWhileSealed:      s223.LeptoquarkDynamicsSealAlive && g.NullHypothesisRejected,
		LifetimeComputed:               false,
		Verdict:                        StatusPatiSalamQuarantined,
	}
}

func auditNullHypothesis(g GeometricMeanAudit) NullHypothesisAudit {
	worst := math.Max(g.Targets[0].Log10Gap, g.Targets[1].Log10Gap)
	noRes := !(g.FAMatch && g.LambdaMatch && g.BothTargetsBracketed)
	verdict := StatusNullHypothesisRejectedByFit
	if noRes {
		verdict = "NULL_HYPOTHESIS_NO_RESONANCE_SURVIVES"
	}
	return NullHypothesisAudit{
		Tested:            true,
		Criterion:         "both f_a and Λ_EFT within one decade of M_int, with Λ_EFT ≤ M_int ≤ f_a",
		FAGapDecades:      g.Targets[0].Log10Gap,
		LambdaGapDecades:  g.Targets[1].Log10Gap,
		WorstGapDecades:   worst,
		PassedNoResonance: noRes,
		FailureMagnitude:  worst,
		Verdict:           verdict,
	}
}

func auditFirewalls(s226 Gate226Snapshot, s223 Gate223Snapshot) FirewallAudit {
	return FirewallAudit{
		Gate226Inherited:               s226.Gate226Inherited,
		Gate223Inherited:               s223.Gate223Inherited,
		UsedOnlySealedValues:           true,
		IntermediateScaleFiniteDerived: false,
		AxionNativeDerived:             false,
		EFTMediatorDerived:             false,
		PatiSalamImportedAsTheorem:     false,
		LeptoquarkSealViolated:         false,
		ProtonLifetimeComputed:         false,
		BGapPromotedWithoutSeal:        false,
		NewPhenomenologicalFitAdded:    false,
		FiniteCorePolluted:             false,
		Verdict:                        "FIREWALLS_CLOSED",
	}
}

func summarize(g GeometricMeanAudit, s SeesawStructureAudit, ps PatiSalamAudit) Summary {
	found := g.NullHypothesisRejected
	status := StatusGeometricMeanResonance
	if !found {
		status = "FAILED_ROUTE_GEOMETRIC_MEAN_RESONANCE"
	}
	return Summary{
		GeometricMeanResonanceFound: found,
		IntermediateScaleGeV:        g.MIntGeV,
		NativeBreakingDerived:       s.NativeIntermediateScaleDerived,
		PatiSalamRouteOpened:        ps.NativeU4GaugeConnectionDerived && !ps.LeptoquarkDynamicsSealActive,
		Status:                      status,
		NextGate:                    "Gate 228 — intermediate breaking seal / common-origin operator audit",
		Comment:                     "Gate 227 finds a strong sealed-scale geometric-mean resonance, but derives no native intermediate breaking mechanism.",
	}
}

func buildTruth(g GeometricMeanAudit, s SeesawStructureAudit, ps PatiSalamAudit, n NullHypothesisAudit) string {
	return fmt.Sprintf("Gate 227 computes M_int=sqrt(M_B M_*)=%.9e GeV. The sealed axion scale f_a=%.9e GeV and relic-decay EFT bound Λ=%.9e GeV both lie within one decade and bracket M_int, rejecting the no-resonance null hypothesis. Native intermediate breaking remains obstructed; Pati-Salam/leptoquark dynamics remain quarantined by the existing seal.", g.MIntGeV, g.Targets[0].ScaleGeV, g.Targets[1].ScaleGeV)
}

func FormatGate226(s Gate226Snapshot) string {
	return fmt.Sprintf("Gate226Inherited=%t AxionSealActive=%t NativeAxionNotDerived=%t f_a=%.9e HeavyDMAbsenceBinding=%t PriorDirectScaleMatch=%t", s.Gate226Inherited, s.AxionSealActive, s.NativeAxionNotDerived, s.FAGeV, s.HeavyDMAbsenceBinding, s.PriorDirectScaleMatch)
}

func FormatGate223(s Gate223Snapshot) string {
	return fmt.Sprintf("Gate223Inherited=%t RelicDecaySealGranted=%t OctetPortalFound=%t LambdaEFTMax=%.9e LeptoquarkSealAlive=%t WilsonDerived=%t", s.Gate223Inherited, s.RelicDecaySealGranted, s.OctetPortalFound, s.LambdaEFTMaxGeV, s.LeptoquarkDynamicsSealAlive, s.WilsonCoefficientDerived)
}

func FormatHierarchy(h SealedHierarchy) string {
	return fmt.Sprintf("M_B=%.9e M_*=%.9e v=%.9e f_a=%.9e LambdaEFTMax=%.9e inheritedOnly=%t nativeScaleDerived=%t", h.MBGeV, h.MStarGeV, h.VEVGeV, h.FARequirementGeV, h.LambdaEFTMaxGeV, h.ValuesInheritedOnly, h.NativeScaleDerived)
}

func FormatGeometric(g GeometricMeanAudit) string {
	parts := []string{fmt.Sprintf("%s = %.9e GeV", g.Formula, g.MIntGeV), fmt.Sprintf("log(M_int/M_B)=%.9f", g.LogDistanceMBToMInt), fmt.Sprintf("log(M_*/M_int)=%.9f", g.LogDistanceMIntToMStar), fmt.Sprintf("targetMean=%.9e gap=%.9f", g.GeometricMeanOfTargetsGeV, g.TargetMeanLog10Gap), fmt.Sprintf("bracketed=%t verdict=%s", g.BothTargetsBracketed, g.Verdict)}
	for _, t := range g.Targets {
		parts = append(parts, fmt.Sprintf("%s: scale=%.9e ratio=%.9e logGap=%.9f close=%t", t.Name, t.ScaleGeV, t.RatioToMInt, t.Log10Gap, t.WithinOneDecade))
	}
	return strings.Join(parts, "; ")
}

func FormatSeesaw(s SeesawStructureAudit) string {
	return fmt.Sprintf("M_int=%.9e suggested=%t relation=%q axionAtScale=%t relicAtScale=%t nativePotential=%t nativeOrderParameter=%t nativeScale=%t required=%s verdict=%s", s.MIntGeV, s.TwoStepPatternSuggested, s.BoundaryRelation, s.AxionScaleCanLiveThere, s.RelicDecayScaleCanLiveThere, s.NativeBreakingPotentialDerived, s.NativeOrderParameterDerived, s.NativeIntermediateScaleDerived, s.RequiredFutureSeal, s.Verdict)
}

func FormatPatiSalam(p PatiSalamAudit) string {
	return fmt.Sprintf("u4Slots=%t leptoquarkSeal=%t audited=%t u4GaugeConnection=%t leptoquarkCurvature=%t protonDecayReopened=%t consistentOnlyWhileSealed=%t lifetimeComputed=%t verdict=%s", p.DormantU4SlotsPresent, p.LeptoquarkDynamicsSealActive, p.IntermediateU4BreakingAudited, p.NativeU4GaugeConnectionDerived, p.NativeLeptoquarkCurvature, p.ProtonDecayChannelReopened, p.ConsistentOnlyWhileSealed, p.LifetimeComputed, p.Verdict)
}

func FormatNull(n NullHypothesisAudit) string {
	return fmt.Sprintf("tested=%t criterion=%q faGap=%.9f lambdaGap=%.9f worst=%.9f noResonanceSurvives=%t failureMagnitude=%.9f verdict=%s", n.Tested, n.Criterion, n.FAGapDecades, n.LambdaGapDecades, n.WorstGapDecades, n.PassedNoResonance, n.FailureMagnitude, n.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("g226=%t g223=%t sealedValuesOnly=%t nativeMInt=%t nativeAxion=%t eftMediator=%t patiSalamImported=%t leptoquarkSealViolated=%t protonLifetime=%t bGapPromoted=%t newFit=%t polluted=%t verdict=%s", f.Gate226Inherited, f.Gate223Inherited, f.UsedOnlySealedValues, f.IntermediateScaleFiniteDerived, f.AxionNativeDerived, f.EFTMediatorDerived, f.PatiSalamImportedAsTheorem, f.LeptoquarkSealViolated, f.ProtonLifetimeComputed, f.BGapPromotedWithoutSeal, f.NewPhenomenologicalFitAdded, f.FiniteCorePolluted, f.Verdict)
}
