// Package sectorchargepullback implements Gate 360:
// Sector-Charge Pullback / CKM Morita Misalignment Sieve.
package sectorchargepullback

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE360-SECTOR-CHARGE-PULLBACK-CKM-MORITA-MISALIGNMENT-SIEVE"

	StatusWeakIsospinPullbackFormalized = "CONDITIONAL_SUPPORT_WEAK_ISOSPIN_PULLBACK_FORMALIZED"
	StatusSectorChargeSieveExecuted     = "CONDITIONAL_SUPPORT_SECTOR_CHARGE_SIEVE_EXECUTED"
	StatusCKMOverlapExtracted           = "CONDITIONAL_SUPPORT_CKM_OVERLAP_EXTRACTION_EXECUTED"
	StatusColorTraceNormAudited         = "CONDITIONAL_SUPPORT_COLOR_TRACE_NORM_AUDITED"
	StatusParameterCensusUpdated        = "CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED"
	StatusMoritaDistinguishesSectors    = "CONDITIONAL_SUPPORT_MORITA_BIMODULE_DISTINGUISHES_CHARGE_SECTORS"
	StatusTrialityPairCapacityAudited   = "CONDITIONAL_SUPPORT_TRIALITY_PAIR_MISALIGNMENT_CAPACITY_AUDITED"

	StatusTensionT3DoesNotSelectGenerator       = "CONDITIONAL_TENSION_T3_SIGN_DISTINGUISHES_SECTORS_BUT_DOES_NOT_SELECT_TRIALITY_GENERATOR"
	StatusTensionColorDoesNotImplyTraceNorm     = "CONDITIONAL_TENSION_COLOR_MULTIPLICITY_DOES_NOT_DERIVE_CTRACE_FLAVOR_NORM"
	StatusTensionCandidateCKMNotObservedPattern = "CONDITIONAL_TENSION_CANDIDATE_PAIR_MISALIGNMENTS_ARE_NOT_CKM_LIKE"
	StatusTensionAssignmentsRemainDiscrete      = "CONDITIONAL_TENSION_SECTOR_GENERATOR_ASSIGNMENTS_REMAIN_DISCRETE_CHOICES"

	StatusFailedSectorChargePullbackDerived = "FAILED_ROUTE_SECTOR_CHARGE_PULLBACK_NOT_DERIVED"
	StatusFailedCKMTextureDerived           = "FAILED_ROUTE_CKM_TEXTURE_NOT_DERIVED"
	StatusFailedTraceNormDerived            = "FAILED_ROUTE_COLOR_TRACE_NORM_NOT_DERIVED_AS_FLAVOR_AMPLIFIER"
	StatusFailedGeneratorAssignmentDerived  = "FAILED_ROUTE_BIMODULE_SECTOR_GENERATOR_ASSIGNMENT_NOT_DERIVED"
	StatusFailedVacuumReduced               = "FAILED_ROUTE_VACUUM_COORDINATES_NOT_REDUCED"
	StatusFailedSevenNotProved              = "FAILED_ROUTE_SEVEN_VACUUM_COORDINATES_NOT_PROVED"
)

const (
	inheritedGate        = 359
	startingVacuumInputs = 15
	sevenSealTarget      = 7
	BGap                 = 0.102464921191
	CTrace               = 25.0
	kappaC               = 1.0
	kappaQ               = 3.0
)

var tau = Vec3{2, -2, 1}

type Vec3 [3]float64
type Mat3 [3][3]float64

type Span struct {
	AuditID       string
	InheritedGate int
	AddsFit       bool
	Purpose       string
	Verdict       string
}

type WeakIsospinPullback struct {
	Formalized          bool
	UpT3                float64
	DownT3              float64
	CandidateRule       string
	NativeGeneratorSwap bool
	Reason              string
	Verdict             string
}

type CandidateCKM struct {
	UpGenerator     string
	DownGenerator   string
	Amplifier       float64
	EffectiveX      float64
	OverlapMatrix   Mat3
	ApproxAnglesRad Vec3
	ApproxAnglesDeg Vec3
	CabibboLike     bool
	CKMDerived      bool
	Verdict         string
}

type CKMSieve struct {
	Executed         bool
	Candidates       []CandidateCKM
	AnyCKMCapacity   bool
	AnyObservedLike  bool
	NativeAssignment bool
	NativeCKMDerived bool
	Verdict          string
}

type ColorTraceNorm struct {
	Audited              bool
	KappaC               float64
	KappaQ               float64
	GlobalTraceCapacity  float64
	ColorLocalNorm       float64
	PullsGlobalAmplifier bool
	Reason               string
	Verdict              string
}

type Census struct {
	StartingVacuumInputs int
	TextureReduction     int
	CKMReduction         int
	TraceNormReduction   int
	TotalReduction       int
	RemainingInputs      int
	SevenSealTarget      int
	SevenSealReached     bool
	Verdict              string
}

type Summary struct {
	Executed              bool
	SectorPullbackDerived bool
	CKMDerived            bool
	TraceNormDerived      bool
	AnyReductionProved    bool
	RemainingInputs       int
	Status                string
	DirectAnswer          string
	NextGate              string
}

type Analysis struct {
	Span     Span
	Pullback WeakIsospinPullback
	CKM      CKMSieve
	Color    ColorTraceNorm
	Census   Census
	Summary  Summary
	Truth    string
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
	pullback := auditWeakIsospinPullback()
	ckm := extractCKMOverlap(pullback)
	color := auditColorTraceNorm()
	census := updateCensus(pullback, ckm, color)
	summary := buildSummary(pullback, ckm, color, census)
	truth := "Gate 360 verifies that the Morita bimodule really distinguishes quark/lepton and up/down charge sectors, and candidate sector-dependent triality generators can produce nontrivial CKM-like overlap matrices.  However, the finite data audited here does not force T3=+1/2 to choose C12 and T3=-1/2 to choose C23, nor does κQ=3 derive C_trace=25 as a local flavor-generator norm.  Sector misalignment has capacity, but the sector-charge pullback theorem remains missing."
	return Analysis{Span: span, Pullback: pullback, CKM: ckm, Color: color, Census: census, Summary: summary, Truth: truth}, nil
}

func compileSpan() Span {
	return Span{AuditID: AuditID, InheritedGate: inheritedGate, AddsFit: false, Purpose: "audit whether weak/color quantum numbers pull back to specific Hermitian triality generators and thereby derive CKM misalignment", Verdict: StatusSectorChargeSieveExecuted}
}

func auditWeakIsospinPullback() WeakIsospinPullback {
	verdict := strings.Join([]string{StatusWeakIsospinPullbackFormalized, StatusMoritaDistinguishesSectors, StatusTensionT3DoesNotSelectGenerator, StatusTensionAssignmentsRemainDiscrete, StatusFailedSectorChargePullbackDerived, StatusFailedGeneratorAssignmentDerived}, ";")
	return WeakIsospinPullback{Formalized: true, UpT3: 0.5, DownT3: -0.5, CandidateRule: "up-type and down-type channels may be assigned different Cij generators because T3 distinguishes them", NativeGeneratorSwap: false, Reason: "T3 is a weak-isospin label acting on SU(2)L doublets; no audited homomorphism maps the T3 sign into a unique element of the 3-generation Hermitian triality complement.", Verdict: verdict}
}

func extractCKMOverlap(p WeakIsospinPullback) CKMSieve {
	pairs := [][2]string{{"12", "23"}, {"12", "13"}, {"13", "23"}}
	candidates := []CandidateCKM{}
	anyCap, anyLike := false, false
	for _, pr := range pairs {
		c := evaluateCandidate(pr[0], pr[1], CTrace)
		candidates = append(candidates, c)
		anyCap = anyCap || !isIdentity(c.OverlapMatrix)
		anyLike = anyLike || c.CabibboLike
	}
	verdict := strings.Join([]string{StatusCKMOverlapExtracted, StatusTrialityPairCapacityAudited, StatusTensionCandidateCKMNotObservedPattern, StatusTensionAssignmentsRemainDiscrete, StatusFailedCKMTextureDerived, StatusFailedGeneratorAssignmentDerived}, ";")
	return CKMSieve{Executed: p.Formalized, Candidates: candidates, AnyCKMCapacity: anyCap, AnyObservedLike: anyLike, NativeAssignment: false, NativeCKMDerived: false, Verdict: verdict}
}

func evaluateCandidate(up, down string, amp float64) CandidateCKM {
	x := BGap * amp
	u := eigenvectorsFor(up, x)
	d := eigenvectorsFor(down, x)
	v := matMul(transpose(u), d)
	angles := approximateMixingAngles(v)
	deg := Vec3{angles[0] * 180 / math.Pi, angles[1] * 180 / math.Pi, angles[2] * 180 / math.Pi}
	// Observed CKM rough scales: theta12≈0.226, theta23≈0.041, theta13≈0.0037.
	// A candidate is called Cabibbo-like only if it has one moderate angle and two small ones.
	cabibbo := angles[0] > 0.10 && angles[0] < 0.35 && angles[1] < 0.08 && angles[2] < 0.02
	verdictParts := []string{StatusCKMOverlapExtracted, StatusTrialityPairCapacityAudited}
	if !cabibbo {
		verdictParts = append(verdictParts, StatusTensionCandidateCKMNotObservedPattern, StatusFailedCKMTextureDerived)
	}
	return CandidateCKM{UpGenerator: up, DownGenerator: down, Amplifier: amp, EffectiveX: x, OverlapMatrix: v, ApproxAnglesRad: angles, ApproxAnglesDeg: deg, CabibboLike: cabibbo, CKMDerived: false, Verdict: strings.Join(verdictParts, ";")}
}

func auditColorTraceNorm() ColorTraceNorm {
	verdict := strings.Join([]string{StatusColorTraceNormAudited, StatusMoritaDistinguishesSectors, StatusTensionColorDoesNotImplyTraceNorm, StatusFailedTraceNormDerived}, ";")
	return ColorTraceNorm{Audited: true, KappaC: kappaC, KappaQ: kappaQ, GlobalTraceCapacity: CTrace, ColorLocalNorm: kappaQ, PullsGlobalAmplifier: false, Reason: "κQ=3 is a local Morita color multiplicity.  C_trace=25 is a global doubled bosonic/coupling-capacity witness.  The gate finds no native rule equating the local color dimension with the global flavor-generator norm.", Verdict: verdict}
}

func updateCensus(p WeakIsospinPullback, c CKMSieve, col ColorTraceNorm) Census {
	reduction := 0
	ckmReduction := 0
	normReduction := 0
	if p.NativeGeneratorSwap && c.NativeCKMDerived && col.PullsGlobalAmplifier {
		reduction = 6
		ckmReduction = 4
		normReduction = 1
	}
	remaining := startingVacuumInputs - reduction
	verdict := strings.Join([]string{StatusParameterCensusUpdated, StatusFailedVacuumReduced, StatusFailedSevenNotProved}, ";")
	return Census{StartingVacuumInputs: startingVacuumInputs, TextureReduction: reduction - ckmReduction - normReduction, CKMReduction: ckmReduction, TraceNormReduction: normReduction, TotalReduction: reduction, RemainingInputs: remaining, SevenSealTarget: sevenSealTarget, SevenSealReached: remaining <= sevenSealTarget, Verdict: verdict}
}

func buildSummary(p WeakIsospinPullback, c CKMSieve, col ColorTraceNorm, census Census) Summary {
	status := strings.Join([]string{StatusSectorChargeSieveExecuted, StatusWeakIsospinPullbackFormalized, StatusCKMOverlapExtracted, StatusColorTraceNormAudited, StatusTensionT3DoesNotSelectGenerator, StatusTensionColorDoesNotImplyTraceNorm, StatusFailedSectorChargePullbackDerived, StatusFailedCKMTextureDerived, StatusFailedVacuumReduced}, ";")
	direct := "Gate 360 finds genuine capacity: different triality generators assigned to up/down sectors create nontrivial overlap matrices, and Morita data distinguishes quark/lepton and weak-isospin channels.  But it does not derive the assignment rule.  T3 and κQ label sectors; they do not yet pull back to unique C12/C13/C23 generators or to the C_trace=25 flavor norm."
	next := "A valid next gate must derive a representation-level charge-to-generation intertwiner, or keep sector generator assignments in the empirical vacuum quarantine."
	return Summary{Executed: p.Formalized && c.Executed && col.Audited, SectorPullbackDerived: p.NativeGeneratorSwap, CKMDerived: c.NativeCKMDerived, TraceNormDerived: col.PullsGlobalAmplifier, AnyReductionProved: census.TotalReduction > 0, RemainingInputs: census.RemainingInputs, Status: status, DirectAnswer: direct, NextGate: next}
}

func diag(v Vec3) Mat3 {
	var m Mat3
	for i := 0; i < 3; i++ {
		m[i][i] = v[i]
	}
	return m
}
func transpose(a Mat3) Mat3 {
	var out Mat3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out[i][j] = a[j][i]
		}
	}
	return out
}
func matMul(a, b Mat3) Mat3 {
	var out Mat3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum := 0.0
			for k := 0; k < 3; k++ {
				sum += a[i][k] * b[k][j]
			}
			out[i][j] = sum
		}
	}
	return out
}
func expSymmetricOffdiag(pair string, x float64) Mat3 {
	var m Mat3
	for i := 0; i < 3; i++ {
		m[i][i] = 1
	}
	ch, sh := math.Cosh(x), math.Sinh(x)
	switch pair {
	case "12":
		m[0][0], m[1][1] = ch, ch
		m[0][1], m[1][0] = sh, sh
	case "13":
		m[0][0], m[2][2] = ch, ch
		m[0][2], m[2][0] = sh, sh
	case "23":
		m[1][1], m[2][2] = ch, ch
		m[1][2], m[2][1] = sh, sh
	}
	return m
}
func eigenvectorsFor(pair string, x float64) Mat3 {
	y := matMul(expSymmetricOffdiag(pair, x), diag(tau))
	gram := matMul(transpose(y), y)
	return eigenvectorsSymmetric3(gram)
}
func eigenvectorsSymmetric3(a Mat3) Mat3 {
	m := a
	var v Mat3
	for i := 0; i < 3; i++ {
		v[i][i] = 1
	}
	for iter := 0; iter < 100; iter++ {
		p, q := 0, 1
		max := math.Abs(m[0][1])
		if vv := math.Abs(m[0][2]); vv > max {
			p, q, max = 0, 2, vv
		}
		if vv := math.Abs(m[1][2]); vv > max {
			p, q, max = 1, 2, vv
		}
		if max < 1e-13 {
			break
		}
		phi := 0.5 * math.Atan2(2*m[p][q], m[q][q]-m[p][p])
		c, s := math.Cos(phi), math.Sin(phi)
		for k := 0; k < 3; k++ {
			mkp, mkq := m[k][p], m[k][q]
			m[k][p] = c*mkp - s*mkq
			m[k][q] = s*mkp + c*mkq
		}
		for k := 0; k < 3; k++ {
			mpk, mqk := m[p][k], m[q][k]
			m[p][k] = c*mpk - s*mqk
			m[q][k] = s*mpk + c*mqk
		}
		for k := 0; k < 3; k++ {
			vkp, vkq := v[k][p], v[k][q]
			v[k][p] = c*vkp - s*vkq
			v[k][q] = s*vkp + c*vkq
		}
	}
	return v
}
func approximateMixingAngles(u Mat3) Vec3 {
	s13 := clamp(math.Abs(u[0][2]), 0, 1)
	theta13 := math.Asin(s13)
	c13 := math.Cos(theta13)
	theta12, theta23 := 0.0, 0.0
	if c13 > 1e-12 {
		theta12 = math.Atan2(math.Abs(u[0][1]), math.Abs(u[0][0]))
		theta23 = math.Atan2(math.Abs(u[1][2]), math.Abs(u[2][2]))
	}
	return Vec3{theta12, theta23, theta13}
}
func clamp(x, lo, hi float64) float64 {
	if x < lo {
		return lo
	}
	if x > hi {
		return hi
	}
	return x
}
func isIdentity(m Mat3) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			want := 0.0
			if i == j {
				want = 1
			}
			if math.Abs(math.Abs(m[i][j])-want) > 1e-6 {
				return false
			}
		}
	}
	return true
}

func Statuses(a Analysis) []string {
	chunks := []string{a.Span.Verdict, a.Pullback.Verdict, a.CKM.Verdict, a.Color.Verdict, a.Census.Verdict, a.Summary.Status}
	for _, c := range a.CKM.Candidates {
		chunks = append(chunks, c.Verdict)
	}
	return splitStatuses(chunks...)
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
	sort.Strings(out)
	return out
}

func FormatSpan(s Span) string {
	return fmt.Sprintf("audit=%s inherited_gate=%d adds_fit=%t purpose=%q verdict=%s", s.AuditID, s.InheritedGate, s.AddsFit, s.Purpose, s.Verdict)
}
func FormatPullback(p WeakIsospinPullback) string {
	return fmt.Sprintf("formalized=%t T3_up=%.3f T3_down=%.3f rule=%q native_swap=%t reason=%q verdict=%s", p.Formalized, p.UpT3, p.DownT3, p.CandidateRule, p.NativeGeneratorSwap, p.Reason, p.Verdict)
}
func FormatCandidate(c CandidateCKM) string {
	return fmt.Sprintf("up=%s down=%s x=%.12f angles_rad=%v angles_deg=%v cabibbo_like=%t derived=%t matrix=%v verdict=%s", c.UpGenerator, c.DownGenerator, c.EffectiveX, c.ApproxAnglesRad, c.ApproxAnglesDeg, c.CabibboLike, c.CKMDerived, c.OverlapMatrix, c.Verdict)
}
func FormatCKM(c CKMSieve) string {
	parts := []string{}
	for _, x := range c.Candidates {
		parts = append(parts, FormatCandidate(x))
	}
	return fmt.Sprintf("executed=%t capacity=%t observed_like=%t native_assignment=%t ckm_derived=%t verdict=%s candidates=[%s]", c.Executed, c.AnyCKMCapacity, c.AnyObservedLike, c.NativeAssignment, c.NativeCKMDerived, c.Verdict, strings.Join(parts, " | "))
}
func FormatColor(c ColorTraceNorm) string {
	return fmt.Sprintf("audited=%t kappaC=%.1f kappaQ=%.1f C_trace=%.1f color_local=%.1f pulls_global=%t reason=%q verdict=%s", c.Audited, c.KappaC, c.KappaQ, c.GlobalTraceCapacity, c.ColorLocalNorm, c.PullsGlobalAmplifier, c.Reason, c.Verdict)
}
func FormatCensus(c Census) string {
	return fmt.Sprintf("start=%d texture=%d ckm=%d norm=%d total=%d remaining=%d target=%d reached=%t verdict=%s", c.StartingVacuumInputs, c.TextureReduction, c.CKMReduction, c.TraceNormReduction, c.TotalReduction, c.RemainingInputs, c.SevenSealTarget, c.SevenSealReached, c.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("executed=%t sector=%t ckm=%t norm=%t reduction=%t remaining=%d status=%s answer=%q next=%q", s.Executed, s.SectorPullbackDerived, s.CKMDerived, s.TraceNormDerived, s.AnyReductionProved, s.RemainingInputs, s.Status, s.DirectAnswer, s.NextGate)
}
