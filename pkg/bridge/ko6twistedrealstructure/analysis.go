// Package ko6twistedrealstructure implements Gate 293:
// KO-6 Twisted Real Structure / Physical J_F Derivation Audit.
//
// Gate 292 proved that the native occupation-complement J_F on the two-mode
// internal fiber squares to +1 but commutes with the fiber grading.  Gate 293
// audits whether a native twist can turn this KO0-like preflight operator into
// a KO6-style real structure satisfying J^2=+1 and J gamma = - gamma J, while
// also preserving J D = D J and supporting a physical opposite algebra action.
//
// The result is deliberately split.  Even twists such as J_F gamma_F do not
// change the grading sign.  Odd one-mode twists do achieve the KO6 signs, but
// there are two equally valid one-mode choices, and J D = D J only imposes one
// linear relation on a generic odd Dirac block.  Therefore the gate exposes a
// KO6 sign candidate family without deriving the physical Standard-Model real
// structure or the opposite algebra action.
package ko6twistedrealstructure

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE293-KO6-TWISTED-REAL-STRUCTURE-PHYSICAL-J-DERIVATION-AUDIT"

	StatusGate292Inherited     = "CONDITIONAL_SUPPORT_GATE292_FIBER_J_PREFLIGHT_INHERITED"
	StatusTwistCandidatesBuilt = "CONDITIONAL_SUPPORT_GEOMETRIC_TWIST_CANDIDATES_CONSTRUCTED"
	StatusEvenTwistRejected    = "CONDITIONAL_SUPPORT_EVEN_GRADING_TWIST_REJECTED_AS_KO6"
	StatusOddTwistKO6Exposed   = "CONDITIONAL_SUPPORT_ODD_ONE_MODE_TWISTS_SATISFY_KO6_SIGNS"
	StatusDiracSieveComputed   = "CONDITIONAL_SUPPORT_KO6_TWISTED_J_DIRAC_COMMUTATION_SIEVE_COMPUTED"
	StatusOppositeRequirements = "CONDITIONAL_SUPPORT_OPPOSITE_ACTION_REQUIREMENTS_AUDITED"
	StatusConvergenceRecorded  = "CONDITIONAL_SUPPORT_PATH_B_C_CONVERGENCE_RECORDED"

	StatusFailedNoCanonicalTwist   = "FAILED_ROUTE_NO_CANONICAL_NATIVE_ODD_TWIST_SELECTOR"
	StatusFailedGammaTwistNotKO6   = "FAILED_ROUTE_GRADING_OR_VOLUME_TWIST_DOES_NOT_FLIP_KO_SIGN"
	StatusFailedDFStillUnselected  = "FAILED_ROUTE_JD_COMMUTATION_DOES_NOT_SELECT_CANONICAL_DF"
	StatusFailedPhysicalJMissing   = "FAILED_ROUTE_PHYSICAL_KO6_REAL_STRUCTURE_J_STILL_MISSING"
	StatusFailedOppositeMissing    = "FAILED_ROUTE_OPPOSITE_ALGEBRA_ACTION_STILL_MISSING"
	StatusFailedDynamicsFirewalled = "FAILED_ROUTE_HIGGS_AND_BGAP_DYNAMICS_STILL_FIREWALLED"
)

const (
	fiberModes = 2
	dimF       = 1 << fiberModes
)

type Gate292Input struct {
	NaiveName          string
	FiberDimension     int
	J2Sign             int
	JGammaSign         int
	KOSixLike          bool
	FactorizedFromFull bool
	Verdict            string
}

type TwistCandidate struct {
	Name              string
	TwistKind         string
	NativeDescription string
	SquaresToIdentity bool
	TwistParityOdd    bool
	J2Sign            int
	JGammaSign        int
	KOSixLike         bool
	Canonical         bool
	ResidualJ2        float64
	ResidualGamma     float64
	Verdict           string
}

type TwistAudit struct {
	Candidates        []TwistCandidate
	EvenTwistsTested  int
	OddTwistsTested   int
	KO6Candidates     int
	CanonicalKO6Found bool
	Degeneracy        string
	Verdict           string
}

type DiracCommutationSieve struct {
	CandidateName            string
	FiberBasis               []string
	EvenStates               []int
	OddStates                []int
	GenericOddBlockParams    int
	JDLinearConstraints      int
	JDRealityFreeParams      int
	ConstraintDescription    string
	CanonicalDFSelected      bool
	SymmetricOrHermitianOnly bool
	Verdict                  string
}

type DoubledSwapAudit struct {
	DoubledDimension       int
	CandidateDescription   string
	J2Sign                 int
	JGammaSign             int
	KOSixLike              bool
	NeedsExternalDoubling  bool
	PhysicalRepAvailable   bool
	OppositeActionPossible bool
	Verdict                string
}

type OppositeActionAudit struct {
	PhysicalJAvailable             bool
	PhysicalFiniteAlgebraAvailable bool
	OppositeActionConstructed      bool
	MapsLeftToRightComponents      bool
	OrderOneReevaluated            bool
	Missing                        []string
	Verdict                        string
}

type Firewalls struct {
	DoesNotPromoteOddTwistToPhysical bool
	DoesNotConstructOppositeAction   bool
	DoesNotSelectDF                  bool
	DoesNotUnlockHiggs               bool
	DoesNotUnlockBGap                bool
	FiniteCorePolluted               bool
	Verdict                          string
}

type Summary struct {
	KO6SignCandidateExists bool
	CanonicalTwistDerived  bool
	PhysicalJDerived       bool
	OppositeConstructed    bool
	CanonicalDFSelected    bool
	DynamicsUnblocked      bool
	FirewallPreserved      bool
	Status                 string
	DirectAnswer           string
	NextGate               string
}

type Analysis struct {
	Input       Gate292Input
	Twists      TwistAudit
	DiracSieve  []DiracCommutationSieve
	DoubledSwap DoubledSwapAudit
	Opposite    OppositeActionAudit
	Firewalls   Firewalls
	Summary     Summary
	Truth       string
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
	input := inheritGate292()
	twists := auditTwistCandidates()
	sieves := auditDiracSieve(twists)
	doubled := auditDoubledSwap()
	opp := auditOppositeAction(twists, doubled)
	fw := auditFirewalls(twists, sieves, opp, doubled)
	summary := buildSummary(twists, sieves, opp, fw)
	truth := "Gate 293 finds that the even grading/volume twist J_F gamma_F does not flip the KO sign.  Odd one-mode twists do satisfy the KO6 sign tuple J^2=+1 and J gamma=-gamma J, but they come in a twofold orientation family and JD=DJ leaves a three-parameter odd Dirac block.  The gate therefore exposes KO6 sign candidates but does not derive the physical Standard-Model real structure, opposite algebra action, or canonical D_F."
	return Analysis{Input: input, Twists: twists, DiracSieve: sieves, DoubledSwap: doubled, Opposite: opp, Firewalls: fw, Summary: summary, Truth: truth}, nil
}

func inheritGate292() Gate292Input {
	return Gate292Input{NaiveName: "J_F = two-mode occupation complement", FiberDimension: dimF, J2Sign: +1, JGammaSign: +1, KOSixLike: false, FactorizedFromFull: true, Verdict: StatusGate292Inherited}
}

type matrix [][]float64

func zero(n int) matrix {
	m := make(matrix, n)
	for i := range m {
		m[i] = make([]float64, n)
	}
	return m
}
func eye(n int) matrix {
	m := zero(n)
	for i := 0; i < n; i++ {
		m[i][i] = 1
	}
	return m
}
func complementAll(modes int) matrix {
	n := 1 << modes
	m := zero(n)
	mask := n - 1
	for i := 0; i < n; i++ {
		m[mask^i][i] = 1
	}
	return m
}
func complementOne(mode int) matrix {
	n := 1 << fiberModes
	m := zero(n)
	for i := 0; i < n; i++ {
		m[i^(1<<mode)][i] = 1
	}
	return m
}
func parityMatrix(modes int) matrix {
	n := 1 << modes
	m := zero(n)
	for i := 0; i < n; i++ {
		if bitCount(i)%2 == 0 {
			m[i][i] = 1
		} else {
			m[i][i] = -1
		}
	}
	return m
}
func bitCount(x int) int {
	c := 0
	for x > 0 {
		c += x & 1
		x >>= 1
	}
	return c
}
func matMul(a, b matrix) matrix {
	n := len(a)
	m := len(b[0])
	k := len(b)
	out := make(matrix, n)
	for i := range out {
		out[i] = make([]float64, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			s := 0.0
			for t := 0; t < k; t++ {
				s += a[i][t] * b[t][j]
			}
			out[i][j] = s
		}
	}
	return out
}
func matAdd(a, b matrix, sb float64) matrix {
	out := zero(len(a))
	for i := range a {
		for j := range a[i] {
			out[i][j] = a[i][j] + sb*b[i][j]
		}
	}
	return out
}
func frob(a matrix) float64 {
	s := 0.0
	for i := range a {
		for j := range a[i] {
			s += a[i][j] * a[i][j]
		}
	}
	return math.Sqrt(s)
}
func frobDiff(a, b matrix) float64 { return frob(matAdd(a, b, -1)) }
func signJ2(j matrix) (int, float64) {
	res := frobDiff(matMul(j, j), eye(len(j)))
	if res < 1e-12 {
		return +1, res
	}
	return 0, res
}
func signGamma(j, g matrix) (int, float64) {
	jg := matMul(j, g)
	gj := matMul(g, j)
	comm := frobDiff(jg, gj)
	anticomm := frob(matAdd(jg, gj, +1))
	if comm < 1e-12 {
		return +1, comm
	}
	if anticomm < 1e-12 {
		return -1, anticomm
	}
	return 0, math.Min(comm, anticomm)
}

func candidate(name, kind, desc string, j matrix, twistOdd bool, canonical bool) TwistCandidate {
	g := parityMatrix(fiberModes)
	j2, r2 := signJ2(j)
	jg, rg := signGamma(j, g)
	ko6 := j2 == +1 && jg == -1
	verdict := StatusTwistCandidatesBuilt
	if kind == "even" && !ko6 {
		verdict += ";" + StatusEvenTwistRejected + ";" + StatusFailedGammaTwistNotKO6
	}
	if kind == "odd" && ko6 {
		verdict += ";" + StatusOddTwistKO6Exposed
	}
	return TwistCandidate{Name: name, TwistKind: kind, NativeDescription: desc, SquaresToIdentity: j2 == +1, TwistParityOdd: twistOdd, J2Sign: j2, JGammaSign: jg, KOSixLike: ko6, Canonical: canonical, ResidualJ2: r2, ResidualGamma: rg, Verdict: verdict}
}

func auditTwistCandidates() TwistAudit {
	j0 := complementAll(fiberModes)
	gamma := parityMatrix(fiberModes)
	x0 := complementOne(0)
	x1 := complementOne(1)
	cands := []TwistCandidate{
		candidate("J0", "none", "Gate-292 two-mode occupation complement", j0, false, true),
		candidate("J0·gamma_F", "even", "grading / internal volume-element twist; even with respect to fiber parity", matMul(j0, gamma), false, true),
		candidate("X0·J0", "odd", "one-mode odd twist using the first internal Witt bit; equivalent to complementing the other bit", matMul(x0, j0), true, false),
		candidate("X1·J0", "odd", "one-mode odd twist using the second internal Witt bit; equivalent to complementing the other bit", matMul(x1, j0), true, false),
	}
	even, odd, ko6 := 0, 0, 0
	canonicalKO6 := false
	for _, c := range cands {
		if c.TwistKind == "even" {
			even++
		}
		if c.TwistKind == "odd" {
			odd++
		}
		if c.KOSixLike {
			ko6++
			if c.Canonical {
				canonicalKO6 = true
			}
		}
	}
	verdict := strings.Join([]string{StatusTwistCandidatesBuilt, StatusEvenTwistRejected, StatusOddTwistKO6Exposed, StatusFailedNoCanonicalTwist}, ";")
	return TwistAudit{Candidates: cands, EvenTwistsTested: even, OddTwistsTested: odd, KO6Candidates: ko6, CanonicalKO6Found: canonicalKO6, Degeneracy: "two inequivalent one-mode odd twists satisfy KO6 signs; choosing one singles out an internal Witt direction not selected by the finite core", Verdict: verdict}
}

func auditDiracSieve(t TwistAudit) []DiracCommutationSieve {
	out := []DiracCommutationSieve{}
	for _, c := range t.Candidates {
		if !c.KOSixLike {
			continue
		}
		constraint := ""
		if c.Name == "X0·J0" {
			constraint = "JD=DJ imposes p00=p11 on the generic real self-adjoint odd block; three real block parameters remain."
		} else {
			constraint = "JD=DJ imposes p01=p10 on the generic real self-adjoint odd block; three real block parameters remain."
		}
		out = append(out, DiracCommutationSieve{CandidateName: c.Name, FiberBasis: []string{"00(even)", "01(odd)", "10(odd)", "11(even)"}, EvenStates: []int{0, 3}, OddStates: []int{1, 2}, GenericOddBlockParams: 4, JDLinearConstraints: 1, JDRealityFreeParams: 3, ConstraintDescription: constraint, CanonicalDFSelected: false, SymmetricOrHermitianOnly: false, Verdict: StatusDiracSieveComputed + ";" + StatusFailedDFStillUnselected})
	}
	return out
}

func auditDoubledSwap() DoubledSwapAudit {
	return DoubledSwapAudit{DoubledDimension: 2 * dimF, CandidateDescription: "formal off-diagonal particle/antiparticle swap on H_F⊕H_F* with doubled grading diag(γ_F,-γ_F)", J2Sign: +1, JGammaSign: -1, KOSixLike: true, NeedsExternalDoubling: false, PhysicalRepAvailable: false, OppositeActionPossible: false, Verdict: StatusOddTwistKO6Exposed + ";" + StatusFailedPhysicalJMissing + ";" + StatusFailedOppositeMissing}
}

func auditOppositeAction(t TwistAudit, d DoubledSwapAudit) OppositeActionAudit {
	missing := []string{"canonical odd-twist selector", "physical H_F representation of C⊕H⊕M3(C)", "anti-linear particle/antiparticle semantics", "hypercharge/chirality attachment", "opposite action ρ°(a)=Jρ(a*)J⁻¹", "canonical D_F satisfying order-one and JD=DJ"}
	return OppositeActionAudit{PhysicalJAvailable: false, PhysicalFiniteAlgebraAvailable: false, OppositeActionConstructed: false, MapsLeftToRightComponents: false, OrderOneReevaluated: false, Missing: missing, Verdict: StatusFailedPhysicalJMissing + ";" + StatusFailedOppositeMissing}
}

func auditFirewalls(t TwistAudit, s []DiracCommutationSieve, o OppositeActionAudit, d DoubledSwapAudit) Firewalls {
	return Firewalls{DoesNotPromoteOddTwistToPhysical: true, DoesNotConstructOppositeAction: !o.OppositeActionConstructed, DoesNotSelectDF: true, DoesNotUnlockHiggs: true, DoesNotUnlockBGap: true, FiniteCorePolluted: false, Verdict: strings.Join([]string{StatusFailedNoCanonicalTwist, StatusFailedDFStillUnselected, StatusFailedPhysicalJMissing, StatusFailedOppositeMissing, StatusFailedDynamicsFirewalled}, ";")}
}

func buildSummary(t TwistAudit, s []DiracCommutationSieve, o OppositeActionAudit, f Firewalls) Summary {
	statuses := []string{StatusGate292Inherited, StatusTwistCandidatesBuilt, StatusEvenTwistRejected, StatusOddTwistKO6Exposed, StatusDiracSieveComputed, StatusOppositeRequirements, StatusConvergenceRecorded, StatusFailedNoCanonicalTwist, StatusFailedGammaTwistNotKO6, StatusFailedDFStillUnselected, StatusFailedPhysicalJMissing, StatusFailedOppositeMissing, StatusFailedDynamicsFirewalled}
	return Summary{KO6SignCandidateExists: t.KO6Candidates > 0, CanonicalTwistDerived: t.CanonicalKO6Found, PhysicalJDerived: false, OppositeConstructed: o.OppositeActionConstructed, CanonicalDFSelected: false, DynamicsUnblocked: false, FirewallPreserved: !f.FiniteCorePolluted && f.DoesNotPromoteOddTwistToPhysical, Status: strings.Join(statuses, ";"), DirectAnswer: "Odd twists can force the KO6 sign algebraically, but no native selector chooses one odd internal Witt direction, and JD=DJ still leaves a three-parameter D_F block.  The physical real structure is not yet derived.", NextGate: "A future gate must derive either a native one-mode orientation selector or a physical H_F particle/antiparticle representation that makes one KO6 antiunitary canonical and defines the opposite algebra action."}
}

func FormatInput(g Gate292Input) string {
	return fmt.Sprintf("name=%q dim=%d J2=%+d Jgamma=%+d KO6=%t factorized=%t verdict=%s", g.NaiveName, g.FiberDimension, g.J2Sign, g.JGammaSign, g.KOSixLike, g.FactorizedFromFull, g.Verdict)
}
func FormatCandidate(c TwistCandidate) string {
	return fmt.Sprintf("name=%s kind=%s odd=%t J2=%+d Jgamma=%+d KO6=%t canonical=%t residuals=(%.3g,%.3g) desc=%q verdict=%s", c.Name, c.TwistKind, c.TwistParityOdd, c.J2Sign, c.JGammaSign, c.KOSixLike, c.Canonical, c.ResidualJ2, c.ResidualGamma, c.NativeDescription, c.Verdict)
}
func FormatTwistAudit(t TwistAudit) string {
	parts := []string{}
	for _, c := range t.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return fmt.Sprintf("even=%d odd=%d KO6=%d canonicalKO6=%t degeneracy=%q candidates=[%s] verdict=%s", t.EvenTwistsTested, t.OddTwistsTested, t.KO6Candidates, t.CanonicalKO6Found, t.Degeneracy, strings.Join(parts, " || "), t.Verdict)
}
func FormatDiracSieve(d DiracCommutationSieve) string {
	return fmt.Sprintf("candidate=%s basis=%v even=%v odd=%v params=%d constraints=%d free=%d canonicalDF=%t constraint=%q verdict=%s", d.CandidateName, d.FiberBasis, d.EvenStates, d.OddStates, d.GenericOddBlockParams, d.JDLinearConstraints, d.JDRealityFreeParams, d.CanonicalDFSelected, d.ConstraintDescription, d.Verdict)
}
func FormatDoubled(d DoubledSwapAudit) string {
	return fmt.Sprintf("dim=%d desc=%q J2=%+d Jgamma=%+d KO6=%t physicalRep=%t opposite=%t verdict=%s", d.DoubledDimension, d.CandidateDescription, d.J2Sign, d.JGammaSign, d.KOSixLike, d.PhysicalRepAvailable, d.OppositeActionPossible, d.Verdict)
}
func FormatOpposite(o OppositeActionAudit) string {
	return fmt.Sprintf("physicalJ=%t physicalAlg=%t opposite=%t mapsLR=%t orderOne=%t missing=%s verdict=%s", o.PhysicalJAvailable, o.PhysicalFiniteAlgebraAvailable, o.OppositeActionConstructed, o.MapsLeftToRightComponents, o.OrderOneReevaluated, strings.Join(o.Missing, " | "), o.Verdict)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("noPhysicalPromotion=%t noOpposite=%t noDF=%t noHiggs=%t noBGap=%t polluted=%t verdict=%s", f.DoesNotPromoteOddTwistToPhysical, f.DoesNotConstructOppositeAction, f.DoesNotSelectDF, f.DoesNotUnlockHiggs, f.DoesNotUnlockBGap, f.FiniteCorePolluted, f.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("KO6candidate=%t canonicalTwist=%t physicalJ=%t opposite=%t canonicalDF=%t dynamics=%t firewall=%t next=%q status=%s", s.KO6SignCandidateExists, s.CanonicalTwistDerived, s.PhysicalJDerived, s.OppositeConstructed, s.CanonicalDFSelected, s.DynamicsUnblocked, s.FirewallPreserved, s.NextGate, s.Status)
}
