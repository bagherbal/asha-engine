// Package tauetargtexture implements Gate 355:
// τ_eta Diagonal Texture RG Evolution / Mass Hierarchy from Topological Seed.
package tauetargtexture

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE355-TAU-ETA-DIAGONAL-TEXTURE-RG-EVOLUTION-MASS-HIERARCHY-AUDIT"

	StatusSeedFormalized        = "CONDITIONAL_SUPPORT_TAU_ETA_DIAGONAL_TEXTURE_SEED_FORMALIZED"
	StatusNormalizationAudited  = "CONDITIONAL_SUPPORT_RPLUS_SECTOR_NORMALIZATION_AUDITED"
	StatusRGTextureExecuted     = "CONDITIONAL_SUPPORT_DIAGONAL_TEXTURE_RG_EVOLUTION_EXECUTED"
	StatusDegeneracyPreserved   = "CONDITIONAL_SUPPORT_FIRST_SECOND_GENERATION_DEGENERACY_PRESERVED"
	StatusHierarchyCompared     = "CONDITIONAL_SUPPORT_MASS_HIERARCHY_COMPARISON_EXECUTED"
	StatusCensusUpdated         = "CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED"
	StatusSpiralCapacityAudited = "CONDITIONAL_SUPPORT_SPIRAL_SEED_CAPACITY_AUDITED"

	StatusTensionNoAmplitudeScale     = "CONDITIONAL_TENSION_OVERALL_YUKAWA_SCALE_X_NOT_DERIVED"
	StatusTensionNoHierarchyAmplified = "CONDITIONAL_TENSION_RG_DOES_NOT_AMPLIFY_DIAGONAL_2_2_1_SEED_TO_OBSERVED_HIERARCHY"
	StatusTensionSignsInvisible       = "CONDITIONAL_TENSION_TAU_ETA_SIGNS_ARE_INVISIBLE_TO_DIAGONAL_YUKAWA_SINGULAR_VALUE_RG"
	StatusTensionTextureRequired      = "CONDITIONAL_TENSION_NON_DIAGONAL_FLAVOR_TEXTURE_OPERATOR_REQUIRED"

	StatusFailedHierarchyNotGenerated = "FAILED_ROUTE_TAU_ETA_DIAGONAL_TEXTURE_DOES_NOT_GENERATE_MASS_HIERARCHY"
	StatusFailedFirstSecondNotSplit   = "FAILED_ROUTE_FIRST_SECOND_GENERATION_SPLITTING_NOT_DERIVED"
	StatusFailedOrderingNotInverted   = "FAILED_ROUTE_THIRD_GENERATION_ENHANCEMENT_NOT_DERIVED"
	StatusFailedNormalizationMissing  = "FAILED_ROUTE_SPECTRAL_YUKAWA_NORMALIZATION_X_NOT_DERIVED"
	StatusFailedCKMTextureMissing     = "FAILED_ROUTE_SIGN_DEPENDENT_CKM_TEXTURE_NOT_DERIVED"
	StatusFailedNoReduction           = "FAILED_ROUTE_NO_ADDITIONAL_PARAMETER_REDUCTION_PROVED"
	StatusFailedSevenNotProved        = "FAILED_ROUTE_SEVEN_VACUUM_COORDINATES_NOT_PROVED"
)

const (
	inheritedGate        = 354
	startingVacuumInputs = 15
	sevenSealTarget      = 7

	rPlus               = 1.645
	gStarSquared        = 0.5
	gutScaleGeV         = 2.40099519719e15
	thresholdScaleGeV   = 1.46774973718e6
	electroweakScaleGeV = 246.22

	integrationHighSteps = 12000
	integrationLowSteps  = 6000
	perturbativeLimitSq  = 16.0 * math.Pi * math.Pi
)

var tauAbs = [3]float64{2, 2, 1}
var tauSigned = [3]float64{2, -2, 1}

type betaCoefficients struct{ B1GUT, B2, B3 float64 }

var highBeta = betaCoefficients{B1GUT: 41.0/10.0 + 7.78628724237, B2: -19.0/6.0 + 9.65295390904, B3: -7.0 + 8.98628724237}
var lowBeta = betaCoefficients{B1GUT: 41.0 / 10.0, B2: -19.0 / 6.0, B3: -7.0}

type Span struct {
	AuditID       string
	InheritedGate int
	AddsFit       bool
	Purpose       string
	Verdict       string
}

type SeedFormalization struct {
	Formalized    bool
	SignedTau     [3]float64
	AbsoluteTau   [3]float64
	Texture       string
	SignsAffectRG bool
	SignComment   string
	Verdict       string
}

type NormalizationAudit struct {
	Formalized       bool
	Constraint       string
	RPlus            float64
	ExampleYe0       float64
	ExampleYu0       float64
	ExampleYd0       float64
	OverallScaleFree bool
	XDerived         bool
	Verdict          string
}

type TextureRun struct {
	Name                   string
	Ye0                    float64
	Yu0                    float64
	Yd0                    float64
	Perturbative           bool
	UpIR                   [3]float64
	DownIR                 [3]float64
	LeptonIR               [3]float64
	UpHighLowRatio         float64
	DownHighLowRatio       float64
	LeptonHighLowRatio     float64
	UpFirstSecondSplit     float64
	DownFirstSecondSplit   float64
	LeptonFirstSecondSplit float64
	MaxObservedLikeRatio   float64
	GeneratedObservedScale bool
	FailureReason          string
}

type RGTextureAudit struct {
	Executed                bool
	HighScaleGeV            float64
	ThresholdGeV            float64
	LowScaleGeV             float64
	BoundaryGStarSq         float64
	UsesPeVThreshold        bool
	Runs                    []TextureRun
	DegeneracyPreserved     bool
	OrderingInverted        bool
	BestHighLowRatio        float64
	ObservedTopCharm        float64
	ObservedBottomStrange   float64
	ObservedTauMuon         float64
	MatchesOrderOfMagnitude bool
	ParameterReduction      int
	ReductionProved         bool
	Verdict                 string
}

type SignTextureAudit struct {
	Formalized              bool
	SignedSeed              [3]float64
	SingularValueSeed       [3]float64
	SignVisibleInDiagonalRG bool
	NeedsOffDiagonalTexture bool
	CKMReductionProved      bool
	Verdict                 string
}

type Census struct {
	StartingVacuumInputs int
	YukawaReduction      int
	CKMReduction         int
	TotalReduction       int
	RemainingInputs      int
	SevenSealTarget      int
	SevenSealReached     bool
	Verdict              string
}

type Summary struct {
	Executed           bool
	SeedPlanted        bool
	HierarchyGenerated bool
	AnyReductionProved bool
	RemainingInputs    int
	Status             string
	DirectAnswer       string
	NextGate           string
}

type Analysis struct {
	Span          Span
	Seed          SeedFormalization
	Normalization NormalizationAudit
	RG            RGTextureAudit
	SignTexture   SignTextureAudit
	Census        Census
	Summary       Summary
	Truth         string
}

type rgState struct {
	G1Sq, G2Sq, G3Sq float64
	U, D, E          [3]float64
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	span := compileSpan()
	seed := formalizeSeed()
	norm := auditNormalization()
	rg := executeTextureRG(norm)
	sign := auditSignTexture()
	census := updateCensus(rg, sign)
	summary := buildSummary(seed, rg, census)
	truth := "Gate 355 plants the tau_eta diagonal seed in the RG spiral.  The seed is transported, but diagonal one-loop RG preserves the first/second generation degeneracy and does not invert the 2:2:1 ordering into a steep third-generation hierarchy.  The signs of tau_eta are invisible to diagonal singular-value running; they can matter only through a non-diagonal flavor texture or CKM-like orientation operator, which remains undeveloped.  Therefore no additional reduction of the 15 minimal vacuum coordinates is proven."
	return Analysis{Span: span, Seed: seed, Normalization: norm, RG: rg, SignTexture: sign, Census: census, Summary: summary, Truth: truth}, nil
}

func compileSpan() Span {
	return Span{AuditID: AuditID, InheritedGate: inheritedGate, AddsFit: false, Purpose: "plant the tau_eta generation seed in the RG spiral and test whether one-loop flow amplifies it into observed fermion hierarchies", Verdict: StatusSpiralCapacityAudited}
}

func formalizeSeed() SeedFormalization {
	return SeedFormalization{Formalized: true, SignedTau: tauSigned, AbsoluteTau: tauAbs, Texture: "Y_s(Lambda_GUT)=y_s0*diag(|tau_eta|)=y_s0*diag(2,2,1) for s in {u,d,e}", SignsAffectRG: false, SignComment: "diagonal Yukawa singular-value RG depends on Y^dagger Y, so signs +2 and -2 are erased unless an off-diagonal orientation/texture operator is present", Verdict: strings.Join([]string{StatusSeedFormalized, StatusTensionSignsInvisible, StatusFailedCKMTextureMissing}, ";")}
}

func auditNormalization() NormalizationAudit {
	ye0 := 0.1
	yu0 := math.Sqrt(rPlus/2.0) * ye0
	yd0 := yu0
	return NormalizationAudit{Formalized: true, Constraint: "(y_u0^2+y_d0^2)/y_e0^2 = r_+; symmetric witness y_u0=y_d0=sqrt(r_+/2)y_e0", RPlus: rPlus, ExampleYe0: ye0, ExampleYu0: yu0, ExampleYd0: yd0, OverallScaleFree: true, XDerived: false, Verdict: strings.Join([]string{StatusNormalizationAudited, StatusTensionNoAmplitudeScale, StatusFailedNormalizationMissing}, ";")}
}

func executeTextureRG(norm NormalizationAudit) RGTextureAudit {
	yeLanes := []struct {
		Name string
		Ye0  float64
	}{{"small_linear_seed", 0.01}, {"moderate_calibrated_witness", norm.ExampleYe0}, {"large_near_attractor_witness", 0.7}}
	runs := make([]TextureRun, 0, len(yeLanes))
	best := 0.0
	degenerate := true
	inverted := false
	matchOrder := false
	for _, lane := range yeLanes {
		run := runTextureLane(lane.Name, lane.Ye0)
		runs = append(runs, run)
		best = math.Max(best, math.Max(run.UpHighLowRatio, math.Max(run.DownHighLowRatio, run.LeptonHighLowRatio)))
		degenerate = degenerate && math.Abs(run.UpFirstSecondSplit) < 1e-10 && math.Abs(run.DownFirstSecondSplit) < 1e-10 && math.Abs(run.LeptonFirstSecondSplit) < 1e-10
		inverted = inverted || run.UpIR[2] > run.UpIR[1] || run.DownIR[2] > run.DownIR[1] || run.LeptonIR[2] > run.LeptonIR[1]
		matchOrder = matchOrder || run.GeneratedObservedScale
	}
	verdict := strings.Join([]string{StatusRGTextureExecuted, StatusDegeneracyPreserved, StatusHierarchyCompared, StatusTensionNoHierarchyAmplified, StatusFailedHierarchyNotGenerated, StatusFailedFirstSecondNotSplit, StatusFailedOrderingNotInverted}, ";")
	return RGTextureAudit{Executed: true, HighScaleGeV: gutScaleGeV, ThresholdGeV: thresholdScaleGeV, LowScaleGeV: electroweakScaleGeV, BoundaryGStarSq: gStarSquared, UsesPeVThreshold: true, Runs: runs, DegeneracyPreserved: degenerate, OrderingInverted: inverted, BestHighLowRatio: best, ObservedTopCharm: 136, ObservedBottomStrange: 44, ObservedTauMuon: 17, MatchesOrderOfMagnitude: matchOrder, ParameterReduction: 0, ReductionProved: false, Verdict: verdict}
}

func runTextureLane(name string, ye0 float64) TextureRun {
	yu0 := math.Sqrt(rPlus/2.0) * ye0
	yd0 := yu0
	state := initialState(yu0, yd0, ye0)
	high, okHigh := integrate(state, gutScaleGeV, thresholdScaleGeV, highBeta, integrationHighSteps)
	if !okHigh {
		return TextureRun{Name: name, Ye0: ye0, Yu0: yu0, Yd0: yd0, Perturbative: false, FailureReason: "high-scale segment became nonperturbative"}
	}
	low, okLow := integrate(high, thresholdScaleGeV, electroweakScaleGeV, lowBeta, integrationLowSteps)
	if !okLow {
		return TextureRun{Name: name, Ye0: ye0, Yu0: yu0, Yd0: yd0, Perturbative: false, FailureReason: "low-scale segment became nonperturbative"}
	}
	upRatio := safeRatio(low.U[0], low.U[2])
	downRatio := safeRatio(low.D[0], low.D[2])
	leptonRatio := safeRatio(low.E[0], low.E[2])
	maxRatio := math.Max(upRatio, math.Max(downRatio, leptonRatio))
	observedScale := maxRatio > 10
	return TextureRun{Name: name, Ye0: ye0, Yu0: yu0, Yd0: yd0, Perturbative: true, UpIR: low.U, DownIR: low.D, LeptonIR: low.E, UpHighLowRatio: upRatio, DownHighLowRatio: downRatio, LeptonHighLowRatio: leptonRatio, UpFirstSecondSplit: low.U[0] - low.U[1], DownFirstSecondSplit: low.D[0] - low.D[1], LeptonFirstSecondSplit: low.E[0] - low.E[1], MaxObservedLikeRatio: maxRatio, GeneratedObservedScale: observedScale}
}

func initialState(yu0, yd0, ye0 float64) rgState {
	return rgState{G1Sq: gStarSquared, G2Sq: gStarSquared, G3Sq: gStarSquared, U: [3]float64{2 * yu0, 2 * yu0, yu0}, D: [3]float64{2 * yd0, 2 * yd0, yd0}, E: [3]float64{2 * ye0, 2 * ye0, ye0}}
}

func integrate(s rgState, from, to float64, beta betaCoefficients, steps int) (rgState, bool) {
	t0, t1 := math.Log(from), math.Log(to)
	h := (t1 - t0) / float64(steps)
	for i := 0; i < steps; i++ {
		k1 := deriv(s, beta)
		k2 := deriv(addScaled(s, k1, h/2), beta)
		k3 := deriv(addScaled(s, k2, h/2), beta)
		k4 := deriv(addScaled(s, k3, h), beta)
		s = combineRK4(s, k1, k2, k3, k4, h)
		if !validState(s) {
			return s, false
		}
	}
	return s, true
}

func deriv(s rgState, beta betaCoefficients) rgState {
	var out rgState
	out.G1Sq = beta.B1GUT / (8 * math.Pi * math.Pi) * s.G1Sq * s.G1Sq
	out.G2Sq = beta.B2 / (8 * math.Pi * math.Pi) * s.G2Sq * s.G2Sq
	out.G3Sq = beta.B3 / (8 * math.Pi * math.Pi) * s.G3Sq * s.G3Sq
	trace := 3*sumSquares(s.U) + 3*sumSquares(s.D) + sumSquares(s.E)
	for i := 0; i < 3; i++ {
		u2, d2, e2 := s.U[i]*s.U[i], s.D[i]*s.D[i], s.E[i]*s.E[i]
		out.U[i] = s.U[i] / (16 * math.Pi * math.Pi) * (1.5*(u2-d2) + trace - (17.0/20.0*s.G1Sq + 9.0/4.0*s.G2Sq + 8.0*s.G3Sq))
		out.D[i] = s.D[i] / (16 * math.Pi * math.Pi) * (1.5*(d2-u2) + trace - (1.0/4.0*s.G1Sq + 9.0/4.0*s.G2Sq + 8.0*s.G3Sq))
		out.E[i] = s.E[i] / (16 * math.Pi * math.Pi) * (1.5*e2 + trace - (9.0/4.0*s.G1Sq + 9.0/4.0*s.G2Sq))
	}
	return out
}

func addScaled(s, k rgState, h float64) rgState {
	r := s
	r.G1Sq += h * k.G1Sq
	r.G2Sq += h * k.G2Sq
	r.G3Sq += h * k.G3Sq
	for i := 0; i < 3; i++ {
		r.U[i] += h * k.U[i]
		r.D[i] += h * k.D[i]
		r.E[i] += h * k.E[i]
	}
	return r
}

func combineRK4(s, k1, k2, k3, k4 rgState, h float64) rgState {
	r := s
	r.G1Sq += h * (k1.G1Sq + 2*k2.G1Sq + 2*k3.G1Sq + k4.G1Sq) / 6
	r.G2Sq += h * (k1.G2Sq + 2*k2.G2Sq + 2*k3.G2Sq + k4.G2Sq) / 6
	r.G3Sq += h * (k1.G3Sq + 2*k2.G3Sq + 2*k3.G3Sq + k4.G3Sq) / 6
	for i := 0; i < 3; i++ {
		r.U[i] += h * (k1.U[i] + 2*k2.U[i] + 2*k3.U[i] + k4.U[i]) / 6
		r.D[i] += h * (k1.D[i] + 2*k2.D[i] + 2*k3.D[i] + k4.D[i]) / 6
		r.E[i] += h * (k1.E[i] + 2*k2.E[i] + 2*k3.E[i] + k4.E[i]) / 6
	}
	return r
}

func validState(s rgState) bool {
	vals := []float64{s.G1Sq, s.G2Sq, s.G3Sq, s.U[0], s.U[1], s.U[2], s.D[0], s.D[1], s.D[2], s.E[0], s.E[1], s.E[2]}
	limit := math.Sqrt(perturbativeLimitSq)
	for _, v := range vals {
		if math.IsNaN(v) || math.IsInf(v, 0) || math.Abs(v) > limit {
			return false
		}
	}
	return s.G1Sq > 0 && s.G2Sq > 0 && s.G3Sq > 0
}

func sumSquares(v [3]float64) float64 { return v[0]*v[0] + v[1]*v[1] + v[2]*v[2] }
func safeRatio(a, b float64) float64 {
	if b == 0 {
		return math.Inf(1)
	}
	return math.Abs(a / b)
}

func auditSignTexture() SignTextureAudit {
	return SignTextureAudit{Formalized: true, SignedSeed: tauSigned, SingularValueSeed: tauAbs, SignVisibleInDiagonalRG: false, NeedsOffDiagonalTexture: true, CKMReductionProved: false, Verdict: strings.Join([]string{StatusTensionSignsInvisible, StatusTensionTextureRequired, StatusFailedCKMTextureMissing}, ";")}
}

func updateCensus(rg RGTextureAudit, sign SignTextureAudit) Census {
	total := rg.ParameterReduction
	if sign.CKMReductionProved {
		total++
	}
	remaining := startingVacuumInputs - total
	return Census{StartingVacuumInputs: startingVacuumInputs, YukawaReduction: rg.ParameterReduction, CKMReduction: boolToInt(sign.CKMReductionProved), TotalReduction: total, RemainingInputs: remaining, SevenSealTarget: sevenSealTarget, SevenSealReached: remaining <= sevenSealTarget, Verdict: strings.Join([]string{StatusCensusUpdated, StatusFailedNoReduction, StatusFailedSevenNotProved}, ";")}
}

func buildSummary(seed SeedFormalization, rg RGTextureAudit, census Census) Summary {
	ok := seed.Formalized && rg.Executed
	hierarchy := rg.MatchesOrderOfMagnitude && rg.OrderingInverted
	status := strings.Join([]string{StatusRGTextureExecuted, StatusFailedHierarchyNotGenerated, StatusFailedNoReduction}, ";")
	direct := "The tau_eta diagonal seed was planted in the RG spiral, but diagonal one-loop running preserves the first/second degeneracy and keeps the seed hierarchy at roughly O(2), not O(10-100).  The spiral does not turn diag(2,2,1) into the observed charged-fermion hierarchy without a non-diagonal flavor texture/operator."
	next := "Derive a non-diagonal flavor-orientation/texture operator, or keep Yukawa singular values quarantined as vacuum coordinates."
	return Summary{Executed: ok, SeedPlanted: ok, HierarchyGenerated: hierarchy, AnyReductionProved: census.TotalReduction > 0, RemainingInputs: census.RemainingInputs, Status: status, DirectAnswer: direct, NextGate: next}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Statuses(a Analysis) []string {
	statuses := []string{a.Span.Verdict, a.Seed.Verdict, a.Normalization.Verdict, a.RG.Verdict, a.SignTexture.Verdict, a.Census.Verdict, a.Summary.Status}
	return splitStatuses(statuses...)
}

func splitStatuses(chunks ...string) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, c := range chunks {
		for _, p := range strings.Split(c, ";") {
			p = strings.TrimSpace(p)
			if p == "" || seen[p] {
				continue
			}
			seen[p] = true
			out = append(out, p)
		}
	}
	return out
}

func FormatSpan(s Span) string {
	return fmt.Sprintf("audit=%s inherited_gate=%d adds_fit=%t purpose=%q verdict=%s", s.AuditID, s.InheritedGate, s.AddsFit, s.Purpose, s.Verdict)
}
func FormatSeed(s SeedFormalization) string {
	return fmt.Sprintf("formalized=%t tau_signed=%v tau_abs=%v texture=%q signs_affect_diagonal_rg=%t verdict=%s", s.Formalized, s.SignedTau, s.AbsoluteTau, s.Texture, s.SignsAffectRG, s.Verdict)
}
func FormatNormalization(n NormalizationAudit) string {
	return fmt.Sprintf("constraint=%q r_plus=%.12f example_y=(u=%.12f,d=%.12f,e=%.12f) scale_free=%t X_derived=%t verdict=%s", n.Constraint, n.RPlus, n.ExampleYu0, n.ExampleYd0, n.ExampleYe0, n.OverallScaleFree, n.XDerived, n.Verdict)
}
func FormatRun(r TextureRun) string {
	return fmt.Sprintf("%s ye0=%.4g yu0=%.4g yd0=%.4g perturbative=%t U=%v D=%v E=%v ratios(U,D,E)=(%.6f,%.6f,%.6f) splits(U,D,E)=(%.3e,%.3e,%.3e) max_ratio=%.6f generated_observed_scale=%t failure=%q", r.Name, r.Ye0, r.Yu0, r.Yd0, r.Perturbative, r.UpIR, r.DownIR, r.LeptonIR, r.UpHighLowRatio, r.DownHighLowRatio, r.LeptonHighLowRatio, r.UpFirstSecondSplit, r.DownFirstSecondSplit, r.LeptonFirstSecondSplit, r.MaxObservedLikeRatio, r.GeneratedObservedScale, r.FailureReason)
}
func FormatRG(r RGTextureAudit) string {
	parts := make([]string, 0, len(r.Runs))
	for _, run := range r.Runs {
		parts = append(parts, FormatRun(run))
	}
	return fmt.Sprintf("executed=%t scales=(%.6e->%.6e->%.6e) gstar2=%.6f peV_threshold=%t deg_preserved=%t inverted=%t best_ratio=%.6f observed_targets=(tc=%.1f,bs=%.1f,tm=%.1f) order_match=%t reduction=%d verdict=%s runs=[%s]", r.Executed, r.HighScaleGeV, r.ThresholdGeV, r.LowScaleGeV, r.BoundaryGStarSq, r.UsesPeVThreshold, r.DegeneracyPreserved, r.OrderingInverted, r.BestHighLowRatio, r.ObservedTopCharm, r.ObservedBottomStrange, r.ObservedTauMuon, r.MatchesOrderOfMagnitude, r.ParameterReduction, r.Verdict, strings.Join(parts, " | "))
}
func FormatSignTexture(s SignTextureAudit) string {
	return fmt.Sprintf("formalized=%t signed_seed=%v singular_seed=%v sign_visible_diagonal_rg=%t needs_offdiag_texture=%t ckm_reduction=%t verdict=%s", s.Formalized, s.SignedSeed, s.SingularValueSeed, s.SignVisibleInDiagonalRG, s.NeedsOffDiagonalTexture, s.CKMReductionProved, s.Verdict)
}
func FormatCensus(c Census) string {
	return fmt.Sprintf("start=%d yukawa_reduction=%d ckm_reduction=%d total_reduction=%d remaining=%d target=%d reached=%t verdict=%s", c.StartingVacuumInputs, c.YukawaReduction, c.CKMReduction, c.TotalReduction, c.RemainingInputs, c.SevenSealTarget, c.SevenSealReached, c.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("executed=%t seed_planted=%t hierarchy_generated=%t any_reduction=%t remaining=%d status=%s answer=%q next=%q", s.Executed, s.SeedPlanted, s.HierarchyGenerated, s.AnyReductionProved, s.RemainingInputs, s.Status, s.DirectAnswer, s.NextGate)
}
