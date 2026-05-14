// Package exponentialtauetatexture implements Gate 358:
// Exponential tau_eta Texture / B-Gap Mixing Hierarchy Audit.
package exponentialtauetatexture

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE358-EXPONENTIAL-TAU-ETA-TEXTURE-BGAP-MIXING-HIERARCHY-AUDIT"

	StatusExponentialMapFormalized  = "CONDITIONAL_SUPPORT_EXPONENTIAL_MAP_FORMALIZED"
	StatusCanonicalGeneratorAudited = "CONDITIONAL_SUPPORT_CANONICAL_TRIALITY_GENERATOR_AUDITED"
	StatusExponentialHierarchySieve = "CONDITIONAL_SUPPORT_EXPONENTIAL_HIERARCHY_SIEVE_EXECUTED"
	StatusSignInterferenceVerified  = "CONDITIONAL_SUPPORT_SIGN_INTERFERENCE_VERIFIED_IN_EXPONENTIAL_TEXTURE"
	StatusCKMShadowAudited          = "CONDITIONAL_SUPPORT_CKM_EIGENVECTOR_SHADOW_AUDITED"
	StatusParameterCensusUpdated    = "CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED"

	StatusTensionCanonicalGeneratorMild  = "CONDITIONAL_TENSION_CANONICAL_GENERATOR_SPLITS_ONLY_MILDLY"
	StatusTensionLargeGeneratorNeeded    = "CONDITIONAL_TENSION_OBSERVED_HIERARCHY_REQUIRES_LARGE_GENERATOR_NORM"
	StatusTensionSectorAssignmentMissing = "CONDITIONAL_TENSION_SECTOR_TRIALITY_OPERATOR_ASSIGNMENT_NOT_CANONICAL"
	StatusTensionCKMNotMatched           = "CONDITIONAL_TENSION_EXPONENTIAL_TEXTURE_DOES_NOT_DERIVE_CKM_PMNS"

	StatusFailedExponentialTextureDerived = "FAILED_ROUTE_EXPONENTIAL_TEXTURE_NOT_DERIVED_AS_OBSERVED_HIERARCHY"
	StatusFailedNativeGeneratorMagnitude  = "FAILED_ROUTE_REQUIRED_TRIALITY_GENERATOR_MAGNITUDE_NOT_DERIVED"
	StatusFailedHierarchyNotDerived       = "FAILED_ROUTE_HIERARCHY_DEGENERACY_NOT_BROKEN_TO_OBSERVED_SCALE"
	StatusFailedCKMNotDerived             = "FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED"
	StatusFailedNoReduction               = "FAILED_ROUTE_NO_ADDITIONAL_PARAMETER_REDUCTION_PROVED"
	StatusFailedSevenNotProved            = "FAILED_ROUTE_SEVEN_VACUUM_COORDINATES_NOT_PROVED"
)

const (
	inheritedGate        = 357
	startingVacuumInputs = 15
	sevenSealTarget      = 7
	BGap                 = 0.102464921191
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

type Generator struct {
	Name        string
	Formula     string
	NativeClaim string
	Pair        string
	Coefficient float64
	Matrix      Mat3
	Canonical   bool
	Verdict     string
}

type TextureResult struct {
	GeneratorName       string
	BGap                float64
	TextureFormula      string
	KineticSafe         bool
	RankPreserved       bool
	SingularValues      Vec3
	HighLowRatio        float64
	FirstSecondRatio    float64
	AnalyticTwoBlock    string
	RequiredCoeffFor17  float64
	RequiredCoeffFor44  float64
	RequiredCoeffFor136 float64
	RequiredCoeffFor207 float64
	ObservedScaleMatch  bool
	Verdict             string
}

type ExponentialMapSieve struct {
	Formalized           bool
	Seed                 Vec3
	BGap                 float64
	Generators           []Generator
	Results              []TextureResult
	AnyRankSafe          bool
	AnyObservedHierarchy bool
	BestRatio            float64
	Verdict              string
}

type CKMShadow struct {
	Audited            bool
	UpGenerator        string
	DownGenerator      string
	MixingMatrix       Mat3
	Angles             Vec3
	QualitativePattern string
	NativeSectorChoice bool
	CKMDerived         bool
	Verdict            string
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
	Executed             bool
	ExponentialTextureOK bool
	ObservedHierarchyOK  bool
	CKMTextureDerived    bool
	AnyReductionProved   bool
	RemainingInputs      int
	Status               string
	DirectAnswer         string
	NextGate             string
}

type Analysis struct {
	Span    Span
	Sieve   ExponentialMapSieve
	CKM     CKMShadow
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
	sieve := executeExponentialSieve()
	ckm := auditCKMShadow(sieve)
	census := updateCensus(sieve, ckm)
	summary := buildSummary(sieve, ckm, census)
	truth := "Gate 358 tests the correct nonlinear escape route: exponentiating a derived triality mixing operator before applying the signed tau_eta seed.  The exponential map is rank-preserving and sign interference in a 1-2 block produces exact e^{+Bc}/e^{-Bc} singular splitting.  However, with canonical normalized off-diagonal generators and B_gap≈0.102, the splitting is mild.  Charged-fermion-scale hierarchies require generator coefficients of order 14-26, or another amplification theorem.  No native sector assignment, CKM texture, or parameter reduction is derived."
	return Analysis{Span: span, Sieve: sieve, CKM: ckm, Census: census, Summary: summary, Truth: truth}, nil
}

func compileSpan() Span {
	return Span{AuditID: AuditID, InheritedGate: inheritedGate, AddsFit: false, Purpose: "audit Y=y0 exp(B_gap C) diag(2,-2,1) as a rank-safe exponential texture capable of sign-interference hierarchy generation", Verdict: StatusExponentialMapFormalized}
}

func executeExponentialSieve() ExponentialMapSieve {
	gens := []Generator{
		offdiagGenerator("C12 canonical", "C12=E12+E21", "representative Hermitian 1-2 triality mixing complement", "12", 1, true),
		offdiagGenerator("C12 amplified witness c=5", "5(E12+E21)", "not canonical; user-suggested natural group-factor witness", "12", 5, false),
		offdiagGenerator("C13 canonical", "C13=E13+E31", "representative Hermitian 1-3 triality mixing complement", "13", 1, true),
		offdiagGenerator("C23 canonical", "C23=E23+E32", "representative Hermitian 2-3 triality mixing complement", "23", 1, true),
	}
	results := make([]TextureResult, 0, len(gens))
	anyRank := false
	anyObs := false
	best := 0.0
	for _, g := range gens {
		r := evaluateTexture(g)
		results = append(results, r)
		if r.RankPreserved && r.KineticSafe {
			anyRank = true
		}
		if r.ObservedScaleMatch {
			anyObs = true
		}
		if r.HighLowRatio > best && !math.IsInf(r.HighLowRatio, 0) {
			best = r.HighLowRatio
		}
	}
	verdict := strings.Join([]string{StatusExponentialMapFormalized, StatusCanonicalGeneratorAudited, StatusExponentialHierarchySieve, StatusSignInterferenceVerified, StatusTensionCanonicalGeneratorMild, StatusTensionLargeGeneratorNeeded, StatusFailedExponentialTextureDerived}, ";")
	return ExponentialMapSieve{Formalized: true, Seed: tau, BGap: BGap, Generators: gens, Results: results, AnyRankSafe: anyRank, AnyObservedHierarchy: anyObs, BestRatio: best, Verdict: verdict}
}

func offdiagGenerator(name, formula, claim, pair string, c float64, canonical bool) Generator {
	var m Mat3
	switch pair {
	case "12":
		m[0][1] = c
		m[1][0] = c
	case "13":
		m[0][2] = c
		m[2][0] = c
	case "23":
		m[1][2] = c
		m[2][1] = c
	}
	verdict := StatusCanonicalGeneratorAudited
	if !canonical {
		verdict = strings.Join([]string{StatusExponentialHierarchySieve, StatusTensionLargeGeneratorNeeded, StatusFailedNativeGeneratorMagnitude}, ";")
	}
	return Generator{Name: name, Formula: formula, NativeClaim: claim, Pair: pair, Coefficient: c, Matrix: m, Canonical: canonical, Verdict: verdict}
}

func evaluateTexture(g Generator) TextureResult {
	expC := expSymmetricOffdiag(g.Pair, BGap*g.Coefficient)
	d := diag(tau)
	y := matMul(expC, d)
	gram := matMul(transpose(y), y)
	eig := eigenSymmetric3(gram)
	sv := Vec3{math.Sqrt(math.Max(eig[0], 0)), math.Sqrt(math.Max(eig[1], 0)), math.Sqrt(math.Max(eig[2], 0))}
	sortDescVec(&sv)
	ratio := sv[0] / sv[2]
	fs := sv[0] / sv[1]
	// In a pure 1-2 two-block with tau=(2,-2), singular ratio is exp(2 B c).
	req17 := math.Log(17) / (2 * BGap)
	req44 := math.Log(44) / (2 * BGap)
	req136 := math.Log(136) / (2 * BGap)
	req207 := math.Log(207) / (2 * BGap)
	observed := ratio >= 10
	verdictParts := []string{StatusExponentialHierarchySieve, StatusSignInterferenceVerified}
	if g.Canonical && ratio < 10 {
		verdictParts = append(verdictParts, StatusTensionCanonicalGeneratorMild, StatusFailedHierarchyNotDerived)
	}
	if !g.Canonical {
		verdictParts = append(verdictParts, StatusTensionLargeGeneratorNeeded, StatusFailedNativeGeneratorMagnitude)
	}
	if observed && !g.Canonical {
		verdictParts = append(verdictParts, StatusFailedExponentialTextureDerived)
	}
	return TextureResult{GeneratorName: g.Name, BGap: BGap, TextureFormula: "Y = exp(B_gap C) diag(2,-2,1)", KineticSafe: true, RankPreserved: minAbs(sv) > 1e-12, SingularValues: sv, HighLowRatio: ratio, FirstSecondRatio: fs, AnalyticTwoBlock: "for C12=c sigma_x, singular values in the 1-2 block are 2 exp(+B c), 2 exp(-B c), so ratio=exp(2 B c)", RequiredCoeffFor17: req17, RequiredCoeffFor44: req44, RequiredCoeffFor136: req136, RequiredCoeffFor207: req207, ObservedScaleMatch: observed && g.Canonical, Verdict: strings.Join(verdictParts, ";")}
}

func auditCKMShadow(s ExponentialMapSieve) CKMShadow {
	// Compare eigenbases of the canonical C12 and C23 Gram matrices.  This tests qualitative
	// misalignment capacity only; no charge/Morita rule selects sector assignment.
	up := evaluateGramEigenvectors(offdiagGenerator("C12 canonical", "C12", "up-sector witness", "12", 1, true))
	down := evaluateGramEigenvectors(offdiagGenerator("C23 canonical", "C23", "down-sector witness", "23", 1, true))
	mix := matMul(transpose(up), down)
	angles := approximateMixingAngles(mix)
	pattern := "canonical off-diagonal generators generate O(1) basis misalignment, closer to democratic/PMNS-like mixing than small CKM angles; sector assignment is not native"
	verdict := strings.Join([]string{StatusCKMShadowAudited, StatusTensionSectorAssignmentMissing, StatusTensionCKMNotMatched, StatusFailedCKMNotDerived}, ";")
	return CKMShadow{Audited: s.Formalized, UpGenerator: "C12 canonical", DownGenerator: "C23 canonical", MixingMatrix: mix, Angles: angles, QualitativePattern: pattern, NativeSectorChoice: false, CKMDerived: false, Verdict: verdict}
}

func updateCensus(s ExponentialMapSieve, c CKMShadow) Census {
	reduction := 0
	if s.AnyObservedHierarchy && c.CKMDerived {
		reduction = 6
	}
	remaining := startingVacuumInputs - reduction
	return Census{StartingVacuumInputs: startingVacuumInputs, TextureReduction: reduction, CKMReduction: 0, TotalReduction: reduction, RemainingInputs: remaining, SevenSealTarget: sevenSealTarget, SevenSealReached: remaining <= sevenSealTarget, Verdict: strings.Join([]string{StatusParameterCensusUpdated, StatusFailedNoReduction, StatusFailedSevenNotProved}, ";")}
}

func buildSummary(s ExponentialMapSieve, c CKMShadow, census Census) Summary {
	status := strings.Join([]string{StatusExponentialMapFormalized, StatusExponentialHierarchySieve, StatusCKMShadowAudited, StatusTensionCanonicalGeneratorMild, StatusFailedHierarchyNotDerived, StatusFailedNoReduction}, ";")
	direct := "The exponential map is the right nonlinear object: it is rank-safe and creates exact sign-interference splitting.  But canonical Gate-261-like normalized off-diagonal generators with B_gap≈0.102 produce only mild ratios around exp(2B)≈1.227.  Observed charged-fermion hierarchies require large generator coefficients not derived from the finite core, and CKM/PMNS sector assignment remains unselected."
	next := "Derive a native generator norm/amplification theorem, or preserve flavor textures as vacuum coordinates."
	return Summary{Executed: s.Formalized && c.Audited, ExponentialTextureOK: s.AnyRankSafe, ObservedHierarchyOK: s.AnyObservedHierarchy, CKMTextureDerived: c.CKMDerived, AnyReductionProved: census.TotalReduction > 0, RemainingInputs: census.RemainingInputs, Status: status, DirectAnswer: direct, NextGate: next}
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
	// Jacobi iteration for real symmetric 3x3 matrices.
	m := a
	for iter := 0; iter < 64; iter++ {
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

func eigenvectorsSymmetric3(a Mat3) Mat3 {
	m := a
	var v Mat3
	for i := 0; i < 3; i++ {
		v[i][i] = 1
	}
	for iter := 0; iter < 64; iter++ {
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

func evaluateGramEigenvectors(g Generator) Mat3 {
	expC := expSymmetricOffdiag(g.Pair, BGap*g.Coefficient)
	y := matMul(expC, diag(tau))
	gram := matMul(transpose(y), y)
	return eigenvectorsSymmetric3(gram)
}

func approximateMixingAngles(u Mat3) Vec3 {
	// Very rough CKM/PMNS-shadow angles from absolute entries; enough for structural audit.
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

func minAbs(v Vec3) float64 {
	m := math.Abs(v[0])
	for i := 1; i < 3; i++ {
		if a := math.Abs(v[i]); a < m {
			m = a
		}
	}
	return m
}

func Statuses(a Analysis) []string {
	chunks := []string{a.Span.Verdict, a.Sieve.Verdict, a.CKM.Verdict, a.Census.Verdict, a.Summary.Status}
	for _, r := range a.Sieve.Results {
		chunks = append(chunks, r.Verdict)
	}
	for _, g := range a.Sieve.Generators {
		chunks = append(chunks, g.Verdict)
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
func FormatGenerator(g Generator) string {
	return fmt.Sprintf("%s formula=%q source=%q pair=%s coeff=%.6g canonical=%t verdict=%s", g.Name, g.Formula, g.NativeClaim, g.Pair, g.Coefficient, g.Canonical, g.Verdict)
}
func FormatTexture(r TextureResult) string {
	return fmt.Sprintf("generator=%s B=%.12f formula=%q kinetic_safe=%t rank=%t sv=%v ratio=%.6f fs=%.6f req17=%.6f req44=%.6f req136=%.6f req207=%.6f observed=%t analytic=%q verdict=%s", r.GeneratorName, r.BGap, r.TextureFormula, r.KineticSafe, r.RankPreserved, r.SingularValues, r.HighLowRatio, r.FirstSecondRatio, r.RequiredCoeffFor17, r.RequiredCoeffFor44, r.RequiredCoeffFor136, r.RequiredCoeffFor207, r.ObservedScaleMatch, r.AnalyticTwoBlock, r.Verdict)
}
func FormatSieve(s ExponentialMapSieve) string {
	gs := []string{}
	rs := []string{}
	for _, g := range s.Generators {
		gs = append(gs, FormatGenerator(g))
	}
	for _, r := range s.Results {
		rs = append(rs, FormatTexture(r))
	}
	return fmt.Sprintf("formalized=%t seed=%v B=%.12f rank_safe=%t observed_hierarchy=%t best_ratio=%.6f verdict=%s generators=[%s] results=[%s]", s.Formalized, s.Seed, s.BGap, s.AnyRankSafe, s.AnyObservedHierarchy, s.BestRatio, s.Verdict, strings.Join(gs, " | "), strings.Join(rs, " | "))
}
func FormatCKM(c CKMShadow) string {
	return fmt.Sprintf("audited=%t up=%s down=%s angles(rad)=%v native_choice=%t derived=%t pattern=%q verdict=%s", c.Audited, c.UpGenerator, c.DownGenerator, c.Angles, c.NativeSectorChoice, c.CKMDerived, c.QualitativePattern, c.Verdict)
}
func FormatCensus(c Census) string {
	return fmt.Sprintf("start=%d texture_reduction=%d ckm_reduction=%d total=%d remaining=%d target=%d reached=%t verdict=%s", c.StartingVacuumInputs, c.TextureReduction, c.CKMReduction, c.TotalReduction, c.RemainingInputs, c.SevenSealTarget, c.SevenSealReached, c.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("executed=%t exponential_ok=%t observed_hierarchy=%t ckm=%t reduction=%t remaining=%d status=%s answer=%q next=%q", s.Executed, s.ExponentialTextureOK, s.ObservedHierarchyOK, s.CKMTextureDerived, s.AnyReductionProved, s.RemainingInputs, s.Status, s.DirectAnswer, s.NextGate)
}
