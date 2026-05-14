// Package realstructurekofactorization implements Gate 292:
// Paths B & C Convergence / Real Structure J Fiber Factorization KO-Dimension Audit.
//
// The gate starts from the Gate-234 occupation-complement candidate J_c on the
// four-Witt-mode Fock carrier and asks whether the Gate-3 spacetime/fiber split
// can factor it into J_M ⊗ J_F.  It then computes the finite KO signs of the
// restricted fiber candidate directly, instead of assuming the Standard Model
// KO-dimension.  The result is a useful no-go: factorization exists, but the
// native occupation-complement restriction on the two internal/fiber Witt modes
// commutes with the fiber parity grading.  It therefore does not yield the
// KO-6 sign J_F gamma_F = - gamma_F J_F required by the physical SM finite
// spectral triple.
package realstructurekofactorization

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE292-REAL-STRUCTURE-J-FIBER-FACTORIZATION-KO-DIMENSION-AUDIT"

	StatusGate234Inherited       = "CONDITIONAL_SUPPORT_GATE234_OCCUPATION_COMPLEMENT_J_INHERITED"
	StatusGate3SplitInherited    = "CONDITIONAL_SUPPORT_GATE3_SPACETIME_FIBER_SPLIT_INHERITED"
	StatusFactorizationVerified  = "CONDITIONAL_SUPPORT_FULL_J_FACTORIZATION_VERIFIED"
	StatusFiberSignsComputed     = "CONDITIONAL_SUPPORT_FIBER_J_KO_SIGNS_COMPUTED"
	StatusDRealitySieveAvailable = "CONDITIONAL_SUPPORT_FIBER_J_REALITY_SIEVE_AVAILABLE"
	StatusConvergenceRecorded    = "CONDITIONAL_SUPPORT_PATH_B_C_CONVERGENCE_RECORDED"

	StatusFailedKO6NotFound       = "FAILED_ROUTE_FIBER_OCCUPATION_COMPLEMENT_J_NOT_KO6"
	StatusFailedPhysicalJMissing  = "FAILED_ROUTE_PHYSICAL_REAL_STRUCTURE_J_F_STILL_MISSING"
	StatusFailedOppositeMissing   = "FAILED_ROUTE_OPPOSITE_ALGEBRA_ACTION_STILL_MISSING"
	StatusFailedDFStillUnselected = "FAILED_ROUTE_CANONICAL_DF_STILL_UNSELECTED"
	StatusFailedDynamicsBlocked   = "FAILED_ROUTE_HIGGS_AND_BGAP_DYNAMICS_STILL_FIREWALLED"
)

const (
	totalModes     = 4
	spacetimeModes = 2
	fiberModes     = 2
	dimTotal       = 1 << totalModes
	dimM           = 1 << spacetimeModes
	dimF           = 1 << fiberModes
)

type Gate234Candidate struct {
	Dimension                 int
	Definition                string
	J2Sign                    int
	JGammaSign                int
	KODescription             string
	PhysicalChargeConjugation bool
	AntiLinearPartDerived     bool
	Verdict                   string
}

type Split struct {
	Source                  string
	SpacetimeRealDirections []int
	FiberRealDirections     []int
	SpacetimeWittModes      []int
	FiberWittModes          []int
	FullIndexFormula        string
	DimensionFactorization  string
	Verdict                 string
}

type Factorization struct {
	FullDimension               int
	SpacetimeDimension          int
	FiberDimension              int
	FullComplementMatchesTensor bool
	Residual                    float64
	FullJ2Sign                  int
	SpacetimeJ2Sign             int
	FiberJ2Sign                 int
	Verdict                     string
}

type KOSigns struct {
	CandidateName     string
	ModesComplemented int
	J2Sign            int
	JGammaSign        int
	JDSignConditional int
	KOZeroLike        bool
	KOSixLike         bool
	ComputedTuple     string
	RequiredSMTuple   string
	Verdict           string
}

type DRealitySieve struct {
	FiberDimension        int
	EvenFiberStates       []int
	OddFiberStates        []int
	GenericOddBlockParams int
	JRealityFreeParams    int
	JDEqualsDJAutomatic   bool
	ConstraintDescription string
	CanonicalDFSelected   bool
	Verdict               string
}

type OppositeActionAudit struct {
	PhysicalFiniteAlgebraAvailable bool
	PhysicalJAvailable             bool
	OppositeActionConstructed      bool
	OrderOneReevaluated            bool
	HeatKernelUnblocked            bool
	BGapInstantonUnblocked         bool
	Missing                        []string
	Verdict                        string
}

type Firewalls struct {
	DoesNotCallFiberJPhysical      bool
	DoesNotClaimKO6                bool
	DoesNotConstructOppositeAction bool
	DoesNotUnlockHiggs             bool
	DoesNotUnlockBGap              bool
	FiniteCorePolluted             bool
	Verdict                        string
}

type Summary struct {
	JFactorized               bool
	FiberSignsComputed        bool
	FiberJKO6                 bool
	PhysicalJDerived          bool
	OppositeActionConstructed bool
	DynamicsUnblocked         bool
	FirewallPreserved         bool
	Status                    string
	DirectAnswer              string
	NextGate                  string
}

type Analysis struct {
	Inherited Gate234Candidate
	Split     Split
	Factor    Factorization
	KO        KOSigns
	DReality  DRealitySieve
	Opposite  OppositeActionAudit
	Firewalls Firewalls
	Summary   Summary
	Truth     string
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
	inherited := inheritGate234()
	split := inheritGate3Split()
	factor, err := auditFactorization()
	if err != nil {
		return Analysis{}, err
	}
	ko := computeFiberKOSigns()
	drel := auditDRealitySieve()
	opp := auditOppositeAction(ko, drel)
	fw := auditFirewalls(ko, opp)
	summary := buildSummary(factor, ko, opp, fw)
	truth := "Gate 292 confirms that the Gate-234 occupation-complement candidate J_c factorizes across the Gate-3 spacetime/fiber Witt-mode split.  However, the restricted fiber complement acts on two internal Witt modes, so it commutes with the fiber parity grading: J_F γ_F = +γ_F J_F.  The computed fiber sign tuple is therefore not the KO-6 Standard-Model tuple.  The physical internal real structure, opposite algebra action, canonical D_F, heat-kernel Higgs route, and B-gap instanton route remain firewalled."
	return Analysis{Inherited: inherited, Split: split, Factor: factor, KO: ko, DReality: drel, Opposite: opp, Firewalls: fw, Summary: summary, Truth: truth}, nil
}

func inheritGate234() Gate234Candidate {
	return Gate234Candidate{Dimension: dimTotal, Definition: "J_c|n0 n1 n2 n3> = |1-n0,1-n1,1-n2,1-n3>", J2Sign: +1, JGammaSign: +1, KODescription: "Gate-234 preflight sign tuple (+,+,+), KO0-like only as bookkeeping", PhysicalChargeConjugation: false, AntiLinearPartDerived: false, Verdict: StatusGate234Inherited}
}

func inheritGate3Split() Split {
	return Split{Source: "Gate-3 covariant phase-space split, read in Witt pairs", SpacetimeRealDirections: []int{0, 1, 2, 3}, FiberRealDirections: []int{4, 5, 6, 7}, SpacetimeWittModes: []int{0, 1}, FiberWittModes: []int{2, 3}, FullIndexFormula: "index = n0 + 2n1 + 4n2 + 8n3 = i_M + 4 i_F", DimensionFactorization: "S_C(4 Witt modes) = S_M(2 modes) ⊗ S_F(2 modes), 16 = 4 × 4", Verdict: StatusGate3SplitInherited}
}

type matrix [][]float64

func zero(n int) matrix {
	m := make(matrix, n)
	for i := range m {
		m[i] = make([]float64, n)
	}
	return m
}

func complementMatrix(modes int) matrix {
	n := 1 << modes
	m := zero(n)
	mask := n - 1
	for i := 0; i < n; i++ {
		m[mask^i][i] = 1
	}
	return m
}

func parityMatrix(modes int) matrix {
	n := 1 << modes
	m := zero(n)
	for i := 0; i < n; i++ {
		if bits(i)%2 == 0 {
			m[i][i] = 1
		} else {
			m[i][i] = -1
		}
	}
	return m
}

func bits(x int) int {
	c := 0
	for x > 0 {
		c += x & 1
		x >>= 1
	}
	return c
}

// tensorIndex follows index = iM + dimM*iF.  The matrix returned is J_M ⊗ J_F in that basis.
func tensorComplementMF() matrix {
	jm := complementMatrix(spacetimeModes)
	jf := complementMatrix(fiberModes)
	out := zero(dimTotal)
	for colF := 0; colF < dimF; colF++ {
		for colM := 0; colM < dimM; colM++ {
			col := colM + dimM*colF
			for rowF := 0; rowF < dimF; rowF++ {
				for rowM := 0; rowM < dimM; rowM++ {
					row := rowM + dimM*rowF
					out[row][col] = jm[rowM][colM] * jf[rowF][colF]
				}
			}
		}
	}
	return out
}

func auditFactorization() (Factorization, error) {
	full := complementMatrix(totalModes)
	tensor := tensorComplementMF()
	residual := frobDiff(full, tensor)
	if residual > 1e-12 {
		return Factorization{}, fmt.Errorf("unexpected J factorization residual %.3g", residual)
	}
	return Factorization{FullDimension: dimTotal, SpacetimeDimension: dimM, FiberDimension: dimF, FullComplementMatchesTensor: true, Residual: residual, FullJ2Sign: signJ2(full), SpacetimeJ2Sign: signJ2(complementMatrix(spacetimeModes)), FiberJ2Sign: signJ2(complementMatrix(fiberModes)), Verdict: StatusFactorizationVerified}, nil
}

func matMul(a, b matrix) matrix {
	n := len(a)
	m := len(b[0])
	k := len(b)
	out := zeroRect(n, m)
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
func zeroRect(n, m int) matrix {
	out := make(matrix, n)
	for i := range out {
		out[i] = make([]float64, m)
	}
	return out
}
func frobDiff(a, b matrix) float64 {
	s := 0.0
	for i := range a {
		for j := range a[i] {
			d := a[i][j] - b[i][j]
			s += d * d
		}
	}
	return math.Sqrt(s)
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
func eye(n int) matrix {
	m := zero(n)
	for i := 0; i < n; i++ {
		m[i][i] = 1
	}
	return m
}
func signJ2(j matrix) int {
	jj := matMul(j, j)
	if frobDiff(jj, eye(len(j))) < 1e-12 {
		return +1
	}
	return 0
}

func signCommute(j, g matrix) int {
	jg := matMul(j, g)
	gj := matMul(g, j)
	if frobDiff(jg, gj) < 1e-12 {
		return +1
	}
	// check anticommutation: JG + GJ = 0
	sum := zero(len(j))
	for i := range sum {
		for k := range sum[i] {
			sum[i][k] = jg[i][k] + gj[i][k]
		}
	}
	if frob(sum) < 1e-12 {
		return -1
	}
	return 0
}

func computeFiberKOSigns() KOSigns {
	jf := complementMatrix(fiberModes)
	gf := parityMatrix(fiberModes)
	j2 := signJ2(jf)
	jgamma := signCommute(jf, gf)
	ko0 := j2 == +1 && jgamma == +1
	ko6 := j2 == +1 && jgamma == -1
	verdict := StatusFiberSignsComputed
	if !ko6 {
		verdict += ";" + StatusFailedKO6NotFound
	}
	return KOSigns{CandidateName: "J_F = occupation-complement restricted to fiber modes n2,n3", ModesComplemented: fiberModes, J2Sign: j2, JGammaSign: jgamma, JDSignConditional: +1, KOZeroLike: ko0, KOSixLike: ko6, ComputedTuple: fmt.Sprintf("(%+d,%+d,+ conditional)", j2, jgamma), RequiredSMTuple: "(+1,-1,+1) for KO6-style finite SM convention", Verdict: verdict}
}

func auditDRealitySieve() DRealitySieve {
	// Fiber Fock basis 00,01,10,11. Even states are 00 and 11; odd are 01 and 10.
	return DRealitySieve{FiberDimension: dimF, EvenFiberStates: []int{0, 3}, OddFiberStates: []int{1, 2}, GenericOddBlockParams: 4, JRealityFreeParams: 2, JDEqualsDJAutomatic: false, ConstraintDescription: "For D_F=[[0,A],[A^T,0]] on the fiber parity split, JD=DJ imposes A=R_even A R_odd; it is a sieve, not a canonical D_F selector.", CanonicalDFSelected: false, Verdict: StatusDRealitySieveAvailable + ";" + StatusFailedDFStillUnselected}
}

func auditOppositeAction(ko KOSigns, d DRealitySieve) OppositeActionAudit {
	missing := []string{"KO6-compatible physical J_F", "faithful physical representation of C⊕H⊕M3(C) on H_F", "opposite algebra action ρ°(a)=Jρ(a*)J⁻¹", "canonical finite Dirac operator D_F", "heat-kernel/scalar-gauge normalization", "B_gap Majorana/inverse-coupling theorem"}
	return OppositeActionAudit{PhysicalFiniteAlgebraAvailable: false, PhysicalJAvailable: ko.KOSixLike, OppositeActionConstructed: false, OrderOneReevaluated: false, HeatKernelUnblocked: false, BGapInstantonUnblocked: false, Missing: missing, Verdict: StatusFailedPhysicalJMissing + ";" + StatusFailedOppositeMissing + ";" + StatusFailedDynamicsBlocked}
}

func auditFirewalls(ko KOSigns, opp OppositeActionAudit) Firewalls {
	return Firewalls{DoesNotCallFiberJPhysical: !ko.KOSixLike, DoesNotClaimKO6: !ko.KOSixLike, DoesNotConstructOppositeAction: !opp.OppositeActionConstructed, DoesNotUnlockHiggs: !opp.HeatKernelUnblocked, DoesNotUnlockBGap: !opp.BGapInstantonUnblocked, FiniteCorePolluted: false, Verdict: strings.Join([]string{StatusFailedKO6NotFound, StatusFailedPhysicalJMissing, StatusFailedOppositeMissing, StatusFailedDynamicsBlocked}, ";")}
}

func buildSummary(f Factorization, ko KOSigns, opp OppositeActionAudit, fw Firewalls) Summary {
	statuses := []string{StatusGate234Inherited, StatusGate3SplitInherited, StatusFactorizationVerified, StatusFiberSignsComputed, StatusDRealitySieveAvailable, StatusConvergenceRecorded, StatusFailedKO6NotFound, StatusFailedPhysicalJMissing, StatusFailedOppositeMissing, StatusFailedDFStillUnselected, StatusFailedDynamicsBlocked}
	return Summary{JFactorized: f.FullComplementMatchesTensor, FiberSignsComputed: true, FiberJKO6: ko.KOSixLike, PhysicalJDerived: ko.KOSixLike && opp.PhysicalJAvailable, OppositeActionConstructed: opp.OppositeActionConstructed, DynamicsUnblocked: opp.HeatKernelUnblocked || opp.BGapInstantonUnblocked, FirewallPreserved: !fw.FiniteCorePolluted && fw.DoesNotClaimKO6, Status: strings.Join(statuses, ";"), DirectAnswer: "No: the Gate-234 occupation-complement J factorizes cleanly, but its fiber restriction complements two Witt modes and therefore commutes with fiber parity. It is KO0-like, not the KO6 physical Standard-Model real structure.", NextGate: "A future gate must derive a twisted/internal J_F, an orientation-volume-factor modification, or a physical H_F representation whose antiunitary real structure satisfies J²=+1, Jγ=-γJ, and yields a lawful opposite algebra action."}
}

func FormatInherited(g Gate234Candidate) string {
	return fmt.Sprintf("dim=%d def=%q J2=%+d Jgamma=%+d physical=%t antiLinear=%t verdict=%s", g.Dimension, g.Definition, g.J2Sign, g.JGammaSign, g.PhysicalChargeConjugation, g.AntiLinearPartDerived, g.Verdict)
}
func FormatSplit(s Split) string {
	return fmt.Sprintf("source=%q spacetimeReal=%v fiberReal=%v spacetimeModes=%v fiberModes=%v dim=%s verdict=%s", s.Source, s.SpacetimeRealDirections, s.FiberRealDirections, s.SpacetimeWittModes, s.FiberWittModes, s.DimensionFactorization, s.Verdict)
}
func FormatFactor(f Factorization) string {
	return fmt.Sprintf("dims=(%d=%d×%d) matches=%t residual=%.3g signs(full,M,F)=(%+d,%+d,%+d) verdict=%s", f.FullDimension, f.SpacetimeDimension, f.FiberDimension, f.FullComplementMatchesTensor, f.Residual, f.FullJ2Sign, f.SpacetimeJ2Sign, f.FiberJ2Sign, f.Verdict)
}
func FormatKO(k KOSigns) string {
	return fmt.Sprintf("candidate=%q modes=%d J2=%+d Jgamma=%+d JD=%+d tuple=%s required=%s KO0like=%t KO6like=%t verdict=%s", k.CandidateName, k.ModesComplemented, k.J2Sign, k.JGammaSign, k.JDSignConditional, k.ComputedTuple, k.RequiredSMTuple, k.KOZeroLike, k.KOSixLike, k.Verdict)
}
func FormatDReality(d DRealitySieve) string {
	return fmt.Sprintf("dim=%d even=%v odd=%v params=%d->%d autoJD=%t canonical=%t constraint=%q verdict=%s", d.FiberDimension, d.EvenFiberStates, d.OddFiberStates, d.GenericOddBlockParams, d.JRealityFreeParams, d.JDEqualsDJAutomatic, d.CanonicalDFSelected, d.ConstraintDescription, d.Verdict)
}
func FormatOpposite(o OppositeActionAudit) string {
	return fmt.Sprintf("physicalAlg=%t physicalJ=%t opposite=%t orderOne=%t heatKernel=%t bGap=%t missing=%s verdict=%s", o.PhysicalFiniteAlgebraAvailable, o.PhysicalJAvailable, o.OppositeActionConstructed, o.OrderOneReevaluated, o.HeatKernelUnblocked, o.BGapInstantonUnblocked, strings.Join(o.Missing, " | "), o.Verdict)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("notPhysicalJ=%t noKO6=%t noOpposite=%t noHiggs=%t noBGap=%t polluted=%t verdict=%s", f.DoesNotCallFiberJPhysical, f.DoesNotClaimKO6, f.DoesNotConstructOppositeAction, f.DoesNotUnlockHiggs, f.DoesNotUnlockBGap, f.FiniteCorePolluted, f.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("factorized=%t fiberSigns=%t KO6=%t physicalJ=%t opposite=%t dynamics=%t firewall=%t next=%q status=%s", s.JFactorized, s.FiberSignsComputed, s.FiberJKO6, s.PhysicalJDerived, s.OppositeActionConstructed, s.DynamicsUnblocked, s.FirewallPreserved, s.NextGate, s.Status)
}
