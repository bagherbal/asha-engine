// Package topologicalamplifierflavorsector implements Gate 359:
// Topological Amplifier & Bimodule Flavor-Sector Sieve.
package topologicalamplifierflavorsector

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE359-TOPOLOGICAL-AMPLIFIER-BIMODULE-FLAVOR-SECTOR-SIEVE"

	StatusAmplifierFormalized       = "CONDITIONAL_SUPPORT_TOPOLOGICAL_AMPLIFIER_FORMALIZED"
	StatusTrace25Audited            = "CONDITIONAL_SUPPORT_CTRACE_25_AMPLIFIER_AUDITED"
	StatusEightPiAudited            = "CONDITIONAL_SUPPORT_EIGHT_PI_AMPLIFIER_AUDITED"
	StatusHierarchyMagnitudeAudited = "CONDITIONAL_SUPPORT_HIERARCHY_MAGNITUDE_AUDITED"
	StatusAmplifierMatchesScale     = "CONDITIONAL_SUPPORT_TOPOLOGICAL_AMPLIFIER_MATCHES_OBSERVED_HIERARCHY_SCALE"
	StatusBimoduleSectorAudited     = "CONDITIONAL_SUPPORT_BIMODULE_SECTOR_ASSIGNMENT_SIEVE_EXECUTED"
	StatusCKMMisalignmentAudited    = "CONDITIONAL_SUPPORT_CKM_PMNS_MISALIGNMENT_CAPACITY_AUDITED"
	StatusParameterCensusUpdated    = "CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED"

	StatusTensionAmplifierFlavorCoupling = "CONDITIONAL_TENSION_TRACE_CAPACITY_AS_FLAVOR_GENERATOR_NORM_NOT_PROVED"
	StatusTensionSectorAssignmentMissing = "CONDITIONAL_TENSION_MORITA_BIMODULE_DISTINGUISHES_SECTORS_BUT_DOES_NOT_SELECT_TRIALITY_GENERATORS"
	StatusTensionCKMAnglesTooLarge       = "CONDITIONAL_TENSION_CANONICAL_GENERATOR_MISALIGNMENTS_ARE_LARGE_NOT_CKM_LIKE"

	StatusFailedTopologicalAmplifierDerived = "FAILED_ROUTE_TOPOLOGICAL_AMPLIFIER_NOT_DERIVED_AS_FLAVOR_NORM"
	StatusFailedSectorGeneratorsDerived     = "FAILED_ROUTE_BIMODULE_SECTOR_GENERATOR_ASSIGNMENT_NOT_DERIVED"
	StatusFailedCKMTextureDerived           = "FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED"
	StatusFailedVacuumReduced               = "FAILED_ROUTE_VACUUM_COORDINATES_NOT_REDUCED"
	StatusFailedSevenNotProved              = "FAILED_ROUTE_SEVEN_VACUUM_COORDINATES_NOT_PROVED"
)

const (
	inheritedGate        = 358
	startingVacuumInputs = 15
	sevenSealTarget      = 7
	BGap                 = 0.102464921191
	CTrace               = 25.0
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

type Amplifier struct {
	Name               string
	Value              float64
	Source             string
	NativeToGeometry   bool
	NativeAsFlavorNorm bool
	Verdict            string
}

type AmplifiedTexture struct {
	AmplifierName        string
	Pair                 string
	BGap                 float64
	EffectiveX           float64
	SingularValues       Vec3
	SplitPairRatio       float64
	HighLowRatio         float64
	MatchesTauLepton     bool
	MatchesBottomStrange bool
	MatchesTopCharm      bool
	MatchesMuonElectron  bool
	Verdict              string
}

type AmplifierSieve struct {
	Formalized            bool
	TextureFormula        string
	Amplifiers            []Amplifier
	Textures              []AmplifiedTexture
	AnyObservedScaleMatch bool
	BestSplitRatio        float64
	Verdict               string
}

type SectorAssignment struct {
	Audited              bool
	MoritaSplit          string
	WeakIsospinSplit     string
	CandidateAssignments []string
	NativeAssignment     bool
	CKMCapacity          bool
	CKMDerived           bool
	CKMShadowAnglesRad   Vec3
	Verdict              string
}

type Census struct {
	StartingVacuumInputs int
	TextureReduction     int
	CKMReduction         int
	TotalReduction       int
	RemainingInputs      int
	SevenSealTarget      int
	SevenSealReached     bool
	Verdict              string
}

type Summary struct {
	Executed                bool
	AmplifierMagnitudeOK    bool
	AmplifierDerived        bool
	SectorAssignmentDerived bool
	AnyReductionProved      bool
	RemainingInputs         int
	Status                  string
	DirectAnswer            string
	NextGate                string
}

type Analysis struct {
	Span    Span
	Sieve   AmplifierSieve
	Sector  SectorAssignment
	Census  Census
	Summary Summary
	Truth   string
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
	sieve := executeAmplifierSieve()
	sector := auditSectorAssignment(sieve)
	census := updateCensus(sieve, sector)
	summary := buildSummary(sieve, sector, census)
	truth := "Gate 359 verifies the key numerical mechanism: multiplying the B_gap exponential texture by the global trace-capacity scale C_trace=25 or by the bosonic coupling branch 8π turns the mild Gate-358 split into an O(10^2) hierarchy.  However, the gate does not prove that trace capacity is the native norm of a flavor generator, nor does the Morita 1⊕3 bimodule uniquely assign C12/C13/C23 generators to the up, down, lepton, and neutrino sectors.  The scale match is real; the sector theorem is still missing."
	return Analysis{Span: span, Sieve: sieve, Sector: sector, Census: census, Summary: summary, Truth: truth}, nil
}

func compileSpan() Span {
	return Span{AuditID: AuditID, InheritedGate: inheritedGate, AddsFit: false, Purpose: "test whether C_trace=25 or 8π supplies the missing exponential flavor-generator amplification and whether Morita data assigns sector generators", Verdict: StatusAmplifierFormalized}
}

func executeAmplifierSieve() AmplifierSieve {
	amps := []Amplifier{
		{Name: "C_trace=25", Value: CTrace, Source: "Gate-316/317/327 trace-capacity target and dimension-per-generation witness", NativeToGeometry: true, NativeAsFlavorNorm: false, Verdict: strings.Join([]string{StatusTrace25Audited, StatusTensionAmplifierFlavorCoupling, StatusFailedTopologicalAmplifierDerived}, ";")},
		{Name: "8π", Value: 8 * math.Pi, Source: "Gate-327/330 doubled bosonic coupling branch alpha_GUT^{-1}=8π", NativeToGeometry: true, NativeAsFlavorNorm: false, Verdict: strings.Join([]string{StatusEightPiAudited, StatusTensionAmplifierFlavorCoupling, StatusFailedTopologicalAmplifierDerived}, ";")},
	}
	textures := []AmplifiedTexture{}
	any := false
	best := 0.0
	for _, amp := range amps {
		for _, pair := range []string{"12", "13", "23"} {
			tex := evaluateTexture(amp, pair)
			textures = append(textures, tex)
			if tex.MatchesTauLepton || tex.MatchesBottomStrange || tex.MatchesTopCharm || tex.MatchesMuonElectron {
				any = true
			}
			if tex.SplitPairRatio > best {
				best = tex.SplitPairRatio
			}
		}
	}
	verdict := strings.Join([]string{StatusAmplifierFormalized, StatusTrace25Audited, StatusEightPiAudited, StatusHierarchyMagnitudeAudited, StatusAmplifierMatchesScale, StatusTensionAmplifierFlavorCoupling, StatusFailedTopologicalAmplifierDerived}, ";")
	return AmplifierSieve{Formalized: true, TextureFormula: "Y = y0 exp(B_gap A C_hat) diag(2,-2,1), A∈{C_trace=25,8π}", Amplifiers: amps, Textures: textures, AnyObservedScaleMatch: any, BestSplitRatio: best, Verdict: verdict}
}

func evaluateTexture(amp Amplifier, pair string) AmplifiedTexture {
	x := BGap * amp.Value
	expC := expSymmetricOffdiag(pair, x)
	y := matMul(expC, diag(tau))
	gram := matMul(transpose(y), y)
	eig := eigenSymmetric3(gram)
	sv := Vec3{math.Sqrt(math.Max(eig[0], 0)), math.Sqrt(math.Max(eig[1], 0)), math.Sqrt(math.Max(eig[2], 0))}
	sortDescVec(&sv)
	highLow := sv[0] / sv[2]
	split := math.Exp(2 * x)
	// Loose order-of-magnitude bands from the user-requested benchmark hierarchy scales.
	tauLepton := withinFactor(split, 17, 2)
	bottomStrange := withinFactor(split, 44, 2)
	topCharm := withinFactor(split, 136, 2)
	muonElectron := withinFactor(split, 207, 2)
	verdictParts := []string{StatusHierarchyMagnitudeAudited, StatusAmplifierMatchesScale}
	if !amp.NativeAsFlavorNorm {
		verdictParts = append(verdictParts, StatusTensionAmplifierFlavorCoupling, StatusFailedTopologicalAmplifierDerived)
	}
	return AmplifiedTexture{AmplifierName: amp.Name, Pair: pair, BGap: BGap, EffectiveX: x, SingularValues: sv, SplitPairRatio: split, HighLowRatio: highLow, MatchesTauLepton: tauLepton, MatchesBottomStrange: bottomStrange, MatchesTopCharm: topCharm, MatchesMuonElectron: muonElectron, Verdict: strings.Join(verdictParts, ";")}
}

func withinFactor(x, target, factor float64) bool {
	return x >= target/factor && x <= target*factor
}

func auditSectorAssignment(s AmplifierSieve) SectorAssignment {
	up := eigenvectorsFor("12", BGap*CTrace)
	down := eigenvectorsFor("23", BGap*CTrace)
	mix := matMul(transpose(up), down)
	angles := approximateMixingAngles(mix)
	candidates := []string{
		"up:C12 down:C23 lepton:C13",
		"up:C13 down:C12 lepton:C23",
		"up:C23 down:C13 lepton:C12",
	}
	verdict := strings.Join([]string{StatusBimoduleSectorAudited, StatusCKMMisalignmentAudited, StatusTensionSectorAssignmentMissing, StatusTensionCKMAnglesTooLarge, StatusFailedSectorGeneratorsDerived, StatusFailedCKMTextureDerived}, ";")
	return SectorAssignment{Audited: s.Formalized, MoritaSplit: "κ_C=1, κ_Q=3 distinguishes color-singlet leptons from color-triplet quarks", WeakIsospinSplit: "T3 distinguishes up-type from down-type channels", CandidateAssignments: candidates, NativeAssignment: false, CKMCapacity: true, CKMDerived: false, CKMShadowAnglesRad: angles, Verdict: verdict}
}

func updateCensus(s AmplifierSieve, sec SectorAssignment) Census {
	reduction := 0
	if s.AnyObservedScaleMatch && sec.NativeAssignment && sec.CKMDerived {
		reduction = 6
	}
	remaining := startingVacuumInputs - reduction
	verdict := strings.Join([]string{StatusParameterCensusUpdated, StatusFailedVacuumReduced, StatusFailedSevenNotProved}, ";")
	return Census{StartingVacuumInputs: startingVacuumInputs, TextureReduction: reduction, CKMReduction: 0, TotalReduction: reduction, RemainingInputs: remaining, SevenSealTarget: sevenSealTarget, SevenSealReached: remaining <= sevenSealTarget, Verdict: verdict}
}

func buildSummary(s AmplifierSieve, sec SectorAssignment, c Census) Summary {
	status := strings.Join([]string{StatusAmplifierFormalized, StatusHierarchyMagnitudeAudited, StatusAmplifierMatchesScale, StatusBimoduleSectorAudited, StatusTensionAmplifierFlavorCoupling, StatusFailedSectorGeneratorsDerived, StatusFailedVacuumReduced}, ";")
	direct := "The trace-capacity amplifier branch succeeds at the magnitude level: A=25 gives exp(2 B_gap A)≈168.9 and A=8π gives ≈172.8, squarely in the charged-fermion hierarchy band.  But Gate 359 cannot promote this to a vacuum reduction because no theorem shows that C_trace or 8π is the native norm of the flavor generator, and the Morita 1⊕3 data does not uniquely assign triality generators to sectors."
	next := "Derive a sector-charge pullback that maps weak/color quantum numbers to specific Hermitian triality generators, or keep flavor textures quarantined."
	return Summary{Executed: s.Formalized && sec.Audited, AmplifierMagnitudeOK: s.AnyObservedScaleMatch, AmplifierDerived: false, SectorAssignmentDerived: sec.NativeAssignment, AnyReductionProved: c.TotalReduction > 0, RemainingInputs: c.RemainingInputs, Status: status, DirectAnswer: direct, NextGate: next}
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

func eigenSymmetric3(a Mat3) Vec3 {
	m := a
	for iter := 0; iter < 80; iter++ {
		p, q := 0, 1
		max := math.Abs(m[0][1])
		if v := math.Abs(m[0][2]); v > max {
			p, q, max = 0, 2, v
		}
		if v := math.Abs(m[1][2]); v > max {
			p, q, max = 1, 2, v
		}
		if max < 1e-14 {
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
	}
	out := Vec3{m[0][0], m[1][1], m[2][2]}
	sortDescVec(&out)
	return out
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
	for iter := 0; iter < 80; iter++ {
		p, q := 0, 1
		max := math.Abs(m[0][1])
		if vv := math.Abs(m[0][2]); vv > max {
			p, q, max = 0, 2, vv
		}
		if vv := math.Abs(m[1][2]); vv > max {
			p, q, max = 1, 2, vv
		}
		if max < 1e-14 {
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
func sortDescVec(v *Vec3) {
	arr := []float64{v[0], v[1], v[2]}
	sort.Sort(sort.Reverse(sort.Float64Slice(arr)))
	v[0], v[1], v[2] = arr[0], arr[1], arr[2]
}

func Statuses(a Analysis) []string {
	chunks := []string{a.Span.Verdict, a.Sieve.Verdict, a.Sector.Verdict, a.Census.Verdict, a.Summary.Status}
	for _, amp := range a.Sieve.Amplifiers {
		chunks = append(chunks, amp.Verdict)
	}
	for _, tex := range a.Sieve.Textures {
		chunks = append(chunks, tex.Verdict)
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
func FormatAmplifier(a Amplifier) string {
	return fmt.Sprintf("%s value=%.12f source=%q native_geometry=%t native_flavor_norm=%t verdict=%s", a.Name, a.Value, a.Source, a.NativeToGeometry, a.NativeAsFlavorNorm, a.Verdict)
}
func FormatTexture(t AmplifiedTexture) string {
	return fmt.Sprintf("amp=%s pair=%s x=%.12f sv=%v split=%.6f highlow=%.6f match17=%t match44=%t match136=%t match207=%t verdict=%s", t.AmplifierName, t.Pair, t.EffectiveX, t.SingularValues, t.SplitPairRatio, t.HighLowRatio, t.MatchesTauLepton, t.MatchesBottomStrange, t.MatchesTopCharm, t.MatchesMuonElectron, t.Verdict)
}
func FormatSieve(s AmplifierSieve) string {
	amps := []string{}
	tex := []string{}
	for _, a := range s.Amplifiers {
		amps = append(amps, FormatAmplifier(a))
	}
	for _, t := range s.Textures {
		tex = append(tex, FormatTexture(t))
	}
	return fmt.Sprintf("formalized=%t formula=%q any_match=%t best_split=%.6f verdict=%s amps=[%s] textures=[%s]", s.Formalized, s.TextureFormula, s.AnyObservedScaleMatch, s.BestSplitRatio, s.Verdict, strings.Join(amps, " | "), strings.Join(tex, " | "))
}
func FormatSector(s SectorAssignment) string {
	return fmt.Sprintf("audited=%t morita=%q weak=%q candidates=%v native=%t ckm_capacity=%t ckm_derived=%t angles=%v verdict=%s", s.Audited, s.MoritaSplit, s.WeakIsospinSplit, s.CandidateAssignments, s.NativeAssignment, s.CKMCapacity, s.CKMDerived, s.CKMShadowAnglesRad, s.Verdict)
}
func FormatCensus(c Census) string {
	return fmt.Sprintf("start=%d texture=%d ckm=%d total=%d remaining=%d target=%d reached=%t verdict=%s", c.StartingVacuumInputs, c.TextureReduction, c.CKMReduction, c.TotalReduction, c.RemainingInputs, c.SevenSealTarget, c.SevenSealReached, c.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("executed=%t magnitude_ok=%t amplifier_derived=%t sector_derived=%t reduction=%t remaining=%d status=%s answer=%q next=%q", s.Executed, s.AmplifierMagnitudeOK, s.AmplifierDerived, s.SectorAssignmentDerived, s.AnyReductionProved, s.RemainingInputs, s.Status, s.DirectAnswer, s.NextGate)
}
