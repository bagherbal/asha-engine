// Package generation2ckmcommutatorpolynomial implements Gate 487:
// CKM Rephasing-Invariant Polynomial Constraint Search.
//
// Gate 486 blocked the CKM 4->2 native registry write: a two-coordinate
// null-mirror chart is not a theorem about the physical CKM quotient. Gate 487
// audits the proposed rescue route: maybe the shared null-C3 baseline suppresses
// the up/down commutator algebra and therefore supplies the two missing
// rephasing-invariant polynomial constraints.
//
// The result is again fail-closed. A null-C3 baseline constrains a spectral
// shadow/eigenvalue shape. The CKM/Jarlskog sector depends on relative
// eigenvectors and rephasing-invariant commutator polynomials. Synthetic
// operators with the same null-C3 spectrum can have zero, rank-two, or rank-three
// commutators depending only on bridge-chosen relative eigenbasis. Therefore the
// shared null boundary does not itself force the two required invariant
// constraints.
package generation2ckmcommutatorpolynomial

import (
	"fmt"
	"math"
	"math/cmplx"
	"strings"
	"sync"
)

const (
	AuditID = "GATE487-CKM-REPHASING-INVARIANT-POLYNOMIAL-CONSTRAINT-SEARCH"

	StatusGate486Inherited                          = "CONDITIONAL_SUPPORT_GATE486_CKM_FIREWALL_INHERITED"
	StatusNullC3SpectrumAnsatzConstructed           = "CONDITIONAL_SUPPORT_NULL_C3_SPECTRUM_ANSATZ_CONSTRUCTED"
	StatusCommutatorSieveExecuted                   = "CONDITIONAL_SUPPORT_COMMUTATOR_SIEVE_EXECUTED_ON_SYNTHETIC_NULL_SPECTRA"
	StatusCommutatorRankNotSuppressed               = "FAILED_ROUTE_SHARED_NULL_CONE_DOES_NOT_SUPPRESS_UP_DOWN_COMMUTATOR_RANK"
	StatusNoJarlskogPolynomialDerived               = "FAILED_ROUTE_JARLSKOG_COMMUTATOR_POLYNOMIAL_NOT_DERIVED"
	StatusNoRephasingInvariantConstraintsDerived    = "FAILED_ROUTE_TWO_REPHASING_INVARIANT_CKM_CONSTRAINTS_NOT_DERIVED"
	StatusNativeUpDownOperatorsStillAbsent          = "FAILED_ROUTE_NATIVE_UP_DOWN_CLIFFORD_OPERATORS_STILL_ABSENT"
	StatusNullSpectrumDoesNotDetermineEigenbasis    = "FAILED_ROUTE_NULL_SPECTRUM_SHAPE_DOES_NOT_DETERMINE_EIGENBASIS_MISALIGNMENT"
	StatusFirewallBlockedCKMPolynomialRegistryWrite = "FIREWALL_BLOCKED_CKM_POLYNOMIAL_CONSTRAINT_NATIVE_WRITE"
	StatusEmpiricalCKMFitRejected                   = "FAILED_ROUTE_EMPIRICAL_CKM_WOLFENSTEIN_QUARK_MASS_INPUT_REJECTED"
	StatusGate488NativeOperatorSourceSearchDefined  = "CONDITIONAL_SUPPORT_GATE488_NATIVE_UP_DOWN_OPERATOR_SOURCE_SEARCH_DEFINED"
)

const (
	NativeFlavorDim                 = 13
	KXYCoeffDim                     = 9
	CKMPhysicalParameterDim         = 4
	ProposedCompressedDim           = 2
	RequiredConstraintsForFourToTwo = 2
	DerivedInvariantConstraintsNow  = 0
	NullSpectrumS                   = 1.0
	NullSpectrumPsi                 = 0.2
	rankTolerance                   = 1e-9
)

type matrix3 [3][3]complex128

type Inheritance struct {
	Executed                          bool
	Gate485KoideBaselineInherited     bool
	Gate486NullMirrorBridgeOnly       bool
	Gate486NativeCKMTheoremBlocked    bool
	Gate486RequiredInvariantEquations int
	Gate486DerivedInvariantEquations  int
	NoObservedCKMImported             bool
	NativeRegistryClean               bool
	Verdict                           string
	Reason                            string
}

type NullSpectrumAnsatz struct {
	Executed                    bool
	S                           float64
	R                           float64
	Psi                         float64
	EigenShadow                 [3]float64
	RatioRoverS                 float64
	MinkowskiResidual           float64
	SameSpectrumOnly            bool
	EigenvectorsSpecifiedNative bool
	ObservedMassesImported      bool
	Verdict                     string
	Reason                      string
}

type OperatorAudit struct {
	Executed                         bool
	NativeUpOperatorDerived          bool
	NativeDownOperatorDerived        bool
	NativeDiagonalizersDerived       bool
	CandidateSyntheticOperatorsUsed  bool
	OperatorsShareNullSpectrum       bool
	NullBoundaryConstrainsSpectrum   bool
	NullBoundaryConstrainsEigenbasis bool
	CKMAsUuDaggerUdConstructed       bool
	Verdict                          string
	Reason                           string
}

type CommutatorCase struct {
	Name                         string
	RelativeUnitary              string
	SameNullSpectrum             bool
	CommutatorRank               int
	CommutatorFrobeniusNorm      float64
	CommutatorDeterminantAbs     float64
	RephasingInvariantConstraint bool
	ObservedDataImported         bool
	Demonstrates                 string
}

type CommutatorSieve struct {
	Executed                       bool
	Cases                          []CommutatorCase
	CaseCount                      int
	RanksObserved                  []int
	RankVariabilityObserved        bool
	ZeroCommutatorPossible         bool
	RankTwoCommutatorPossible      bool
	RankThreeCommutatorPossible    bool
	SharedNullSpectrumInEveryCase  bool
	NoObservedDataImported         bool
	CommutatorRankSuppressedByNull bool
	JarlskogDeterminantLocked      bool
	InvariantPolynomialProduced    bool
	Verdict                        string
	Reason                         string
}

type ConstraintHunt struct {
	Executed                       bool
	PhysicalCKMParameterDim        int
	ProposedCompressedDim          int
	RequiredIndependentConstraints int
	DerivedIndependentConstraints  int
	ModuliPolynomialRelations      int
	JarlskogPolynomialRelations    int
	CommutatorTraceRelations       int
	CommutatorDeterminantRelation  bool
	TwoConstraintTheoremPassed     bool
	Verdict                        string
	Reason                         string
}

type Firewall struct {
	Executed                         bool
	ObservedCKMImported              bool
	ObservedWolfensteinImported      bool
	ObservedQuarkMassesImported      bool
	ObservedCPPhaseImported          bool
	CKMMatrixNativePrediction        bool
	JarlskogNativePrediction         bool
	CKMFourToTwoNativeWritten        bool
	PolynomialConstraintsNativeWrite bool
	SyntheticCommutatorBridgeOnly    bool
	NativeRegistryWritten            bool
	NativeFlavorDimAfter             int
	KXYCoeffDimAfter                 int
	Verdict                          string
	Reason                           string
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
	Spectrum    NullSpectrumAnsatz
	Operators   OperatorAudit
	Sieve       CommutatorSieve
	Constraints ConstraintHunt
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
	a := Analysis{Inheritance: buildInheritance()}
	a.Spectrum = buildNullSpectrumAnsatz()
	a.Operators = buildOperatorAudit(a.Spectrum)
	a.Sieve = buildCommutatorSieve(a.Spectrum)
	a.Constraints = buildConstraintHunt(a.Sieve)
	a.Firewall = buildFirewall(a)
	a.Registry = buildRegistryUpdate(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                          true,
		Gate485KoideBaselineInherited:     true,
		Gate486NullMirrorBridgeOnly:       true,
		Gate486NativeCKMTheoremBlocked:    true,
		Gate486RequiredInvariantEquations: RequiredConstraintsForFourToTwo,
		Gate486DerivedInvariantEquations:  0,
		NoObservedCKMImported:             true,
		NativeRegistryClean:               true,
		Verdict:                           StatusGate486Inherited,
		Reason:                            "Gate486 left one open task: derive two physical rephasing-invariant polynomial constraints from native up/down operators, not from a coordinate chart or observed CKM data",
	}
}

func buildNullSpectrumAnsatz() NullSpectrumAnsatz {
	r := math.Sqrt2 * NullSpectrumS
	var xs [3]float64
	for i, theta := range []float64{0, 2 * math.Pi / 3, 4 * math.Pi / 3} {
		xs[i] = NullSpectrumS + r*math.Cos(theta-NullSpectrumPsi)
	}
	minkowski := 3*NullSpectrumS*NullSpectrumS - 1.5*r*r
	return NullSpectrumAnsatz{
		Executed:                    true,
		S:                           NullSpectrumS,
		R:                           r,
		Psi:                         NullSpectrumPsi,
		EigenShadow:                 xs,
		RatioRoverS:                 r / NullSpectrumS,
		MinkowskiResidual:           minkowski,
		SameSpectrumOnly:            true,
		EigenvectorsSpecifiedNative: false,
		ObservedMassesImported:      false,
		Verdict:                     StatusNullC3SpectrumAnsatzConstructed,
		Reason:                      "the synthetic spectrum obeys the Gate485 null-C3 ratio R/S=sqrt(2), but it supplies only eigenvalue-shadow shape; no native eigenbasis or quark-sector diagonalizer is encoded",
	}
}

func buildOperatorAudit(s NullSpectrumAnsatz) OperatorAudit {
	return OperatorAudit{
		Executed:                         true,
		NativeUpOperatorDerived:          false,
		NativeDownOperatorDerived:        false,
		NativeDiagonalizersDerived:       false,
		CandidateSyntheticOperatorsUsed:  true,
		OperatorsShareNullSpectrum:       s.Executed && s.SameSpectrumOnly,
		NullBoundaryConstrainsSpectrum:   true,
		NullBoundaryConstrainsEigenbasis: false,
		CKMAsUuDaggerUdConstructed:       false,
		Verdict:                          StatusNativeUpDownOperatorsStillAbsent,
		Reason:                           "Gate487 can test synthetic operators sharing the null-C3 spectrum, but ASHA still lacks native Clifford up/down operators whose diagonalizers would define a physical CKM quotient",
	}
}

func buildCommutatorSieve(s NullSpectrumAnsatz) CommutatorSieve {
	d := diag(s.EigenShadow)
	cases := []CommutatorCase{
		commutatorCase("aligned null spectra", "I", d, identity(), "same null spectrum can commute exactly; rank is not forced away from zero"),
		commutatorCase("real 1-2 bridge rotation", "R12(0.4)", d, rotation12(0.4), "same null spectrum can yield a rank-two commutator under a bridge-chosen real eigenbasis tilt"),
		commutatorCase("complex Fourier bridge frame", "F3", d, fourier(), "same null spectrum can yield a full-rank commutator under a bridge-chosen complex eigenbasis frame"),
	}
	ranks := make([]int, 0, len(cases))
	seen := map[int]bool{}
	allSameSpectrum := true
	noObserved := true
	zero, rankTwo, rankThree := false, false, false
	for _, c := range cases {
		if !seen[c.CommutatorRank] {
			ranks = append(ranks, c.CommutatorRank)
			seen[c.CommutatorRank] = true
		}
		allSameSpectrum = allSameSpectrum && c.SameNullSpectrum
		noObserved = noObserved && !c.ObservedDataImported
		switch c.CommutatorRank {
		case 0:
			zero = true
		case 2:
			rankTwo = true
		case 3:
			rankThree = true
		}
	}
	return CommutatorSieve{
		Executed:                       true,
		Cases:                          cases,
		CaseCount:                      len(cases),
		RanksObserved:                  ranks,
		RankVariabilityObserved:        len(ranks) > 1,
		ZeroCommutatorPossible:         zero,
		RankTwoCommutatorPossible:      rankTwo,
		RankThreeCommutatorPossible:    rankThree,
		SharedNullSpectrumInEveryCase:  allSameSpectrum,
		NoObservedDataImported:         noObserved,
		CommutatorRankSuppressedByNull: false,
		JarlskogDeterminantLocked:      false,
		InvariantPolynomialProduced:    false,
		Verdict:                        StatusCommutatorRankNotSuppressed,
		Reason:                         "holding the exact same null-C3 spectrum fixed does not fix the relative eigenbasis: the commutator rank varies across synthetic bridge frames, so the shared null boundary cannot by itself impose physical CKM invariant constraints",
	}
}

func commutatorCase(name, unitaryLabel string, d matrix3, u matrix3, demonstrates string) CommutatorCase {
	od := conjugate(d, u)
	c := commutator(d, od)
	return CommutatorCase{
		Name:                         name,
		RelativeUnitary:              unitaryLabel,
		SameNullSpectrum:             true,
		CommutatorRank:               rank(c),
		CommutatorFrobeniusNorm:      frobenius(c),
		CommutatorDeterminantAbs:     cmplx.Abs(det(c)),
		RephasingInvariantConstraint: false,
		ObservedDataImported:         false,
		Demonstrates:                 demonstrates,
	}
}

func buildConstraintHunt(s CommutatorSieve) ConstraintHunt {
	return ConstraintHunt{
		Executed:                       true,
		PhysicalCKMParameterDim:        CKMPhysicalParameterDim,
		ProposedCompressedDim:          ProposedCompressedDim,
		RequiredIndependentConstraints: RequiredConstraintsForFourToTwo,
		DerivedIndependentConstraints:  DerivedInvariantConstraintsNow,
		ModuliPolynomialRelations:      0,
		JarlskogPolynomialRelations:    0,
		CommutatorTraceRelations:       0,
		CommutatorDeterminantRelation:  false,
		TwoConstraintTheoremPassed:     false,
		Verdict:                        StatusNoRephasingInvariantConstraintsDerived,
		Reason:                         fmt.Sprintf("the commutator sieve executed (%d cases), but variable synthetic ranks and absent native up/down operators produce zero physical invariant equations; the required count is two", s.CaseCount),
	}
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                         true,
		ObservedCKMImported:              false,
		ObservedWolfensteinImported:      false,
		ObservedQuarkMassesImported:      false,
		ObservedCPPhaseImported:          false,
		CKMMatrixNativePrediction:        false,
		JarlskogNativePrediction:         false,
		CKMFourToTwoNativeWritten:        false,
		PolynomialConstraintsNativeWrite: false,
		SyntheticCommutatorBridgeOnly:    a.Sieve.Executed,
		NativeRegistryWritten:            false,
		NativeFlavorDimAfter:             NativeFlavorDim,
		KXYCoeffDimAfter:                 KXYCoeffDim,
		Verdict:                          StatusFirewallBlockedCKMPolynomialRegistryWrite,
		Reason:                           "Gate487 uses synthetic null-spectrum commutator probes only as a theorem sieve; it imports no CKM/Wolfenstein/quark-mass data and writes no CKM polynomial constraint to the native registry",
	}
}

func buildRegistryUpdate(_ Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"no new native CKM, Jarlskog, or 4->2 theorem",
			"Gate485 null-C3 Koide baseline remains a spectrum-shape theorem only",
		},
		BridgeEntries: []string{
			StatusCommutatorSieveExecuted,
			"synthetic same-null-spectrum commutator probes may be used to test future native operator candidates",
		},
		EnvironmentalEntries: []string{
			"observed CKM matrix, Wolfenstein parameters, quark masses, and CP phase remain forbidden theorem inputs",
			"physical Jarlskog value remains external comparator data",
		},
		FailedRoutes: []string{
			StatusCommutatorRankNotSuppressed,
			StatusNoJarlskogPolynomialDerived,
			StatusNoRephasingInvariantConstraintsDerived,
			StatusNativeUpDownOperatorsStillAbsent,
			StatusNullSpectrumDoesNotDetermineEigenbasis,
			StatusEmpiricalCKMFitRejected,
		},
		OpenTheorems: []string{
			StatusGate488NativeOperatorSourceSearchDefined,
			"find a native finite operator source that distinguishes up/down sectors before any CKM invariant polynomial can be claimed",
		},
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        488,
		Title:       "Native Up/Down Operator Source Search",
		Reason:      "Gate487 proves that null-C3 spectrum shape alone cannot constrain the CKM commutator; the missing object is a native source for sector-specific operators and eigenvectors.",
		PrimaryTask: "search the finite Clifford/spectral data for a native up/down operator pair O_u,O_d, or prove that all available native structures remain generation-blind before the empirical airlock",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate485KoideBaselineInherited || !a.Inheritance.Gate486NullMirrorBridgeOnly || !a.Inheritance.Gate486NativeCKMTheoremBlocked || a.Inheritance.Gate486RequiredInvariantEquations != RequiredConstraintsForFourToTwo || a.Inheritance.Gate486DerivedInvariantEquations != 0 || !a.Inheritance.NoObservedCKMImported || !a.Inheritance.NativeRegistryClean {
		return fmt.Errorf("Gate487 inheritance invalid: %+v", a.Inheritance)
	}
	if !a.Spectrum.Executed || math.Abs(a.Spectrum.RatioRoverS-math.Sqrt2) > rankTolerance || math.Abs(a.Spectrum.MinkowskiResidual) > rankTolerance || !a.Spectrum.SameSpectrumOnly || a.Spectrum.EigenvectorsSpecifiedNative || a.Spectrum.ObservedMassesImported {
		return fmt.Errorf("Gate487 null spectrum invalid: %+v", a.Spectrum)
	}
	if !a.Operators.Executed || a.Operators.NativeUpOperatorDerived || a.Operators.NativeDownOperatorDerived || a.Operators.NativeDiagonalizersDerived || !a.Operators.CandidateSyntheticOperatorsUsed || !a.Operators.OperatorsShareNullSpectrum || !a.Operators.NullBoundaryConstrainsSpectrum || a.Operators.NullBoundaryConstrainsEigenbasis || a.Operators.CKMAsUuDaggerUdConstructed {
		return fmt.Errorf("Gate487 operator audit invalid: %+v", a.Operators)
	}
	if !a.Sieve.Executed || a.Sieve.CaseCount != 3 || !a.Sieve.RankVariabilityObserved || !a.Sieve.ZeroCommutatorPossible || !a.Sieve.RankTwoCommutatorPossible || !a.Sieve.RankThreeCommutatorPossible || !a.Sieve.SharedNullSpectrumInEveryCase || !a.Sieve.NoObservedDataImported || a.Sieve.CommutatorRankSuppressedByNull || a.Sieve.JarlskogDeterminantLocked || a.Sieve.InvariantPolynomialProduced {
		return fmt.Errorf("Gate487 commutator sieve invalid: %+v", a.Sieve)
	}
	if !a.Constraints.Executed || a.Constraints.PhysicalCKMParameterDim != CKMPhysicalParameterDim || a.Constraints.RequiredIndependentConstraints != RequiredConstraintsForFourToTwo || a.Constraints.DerivedIndependentConstraints != 0 || a.Constraints.ModuliPolynomialRelations != 0 || a.Constraints.JarlskogPolynomialRelations != 0 || a.Constraints.CommutatorTraceRelations != 0 || a.Constraints.CommutatorDeterminantRelation || a.Constraints.TwoConstraintTheoremPassed {
		return fmt.Errorf("Gate487 constraint hunt invalid: %+v", a.Constraints)
	}
	if !a.Firewall.Executed || a.Firewall.ObservedCKMImported || a.Firewall.ObservedWolfensteinImported || a.Firewall.ObservedQuarkMassesImported || a.Firewall.ObservedCPPhaseImported || a.Firewall.CKMMatrixNativePrediction || a.Firewall.JarlskogNativePrediction || a.Firewall.CKMFourToTwoNativeWritten || a.Firewall.PolynomialConstraintsNativeWrite || !a.Firewall.SyntheticCommutatorBridgeOnly || a.Firewall.NativeRegistryWritten || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate487 firewall invalid: %+v", a.Firewall)
	}
	return nil
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate487 result: the shared null-C3 boundary constrains spectrum shape, not relative eigenbasis. Synthetic up/down operators with the same null spectrum realize commutator ranks %v, so the null cone does not suppress [O_u,O_d] into two physical CKM constraints. The required invariant constraints remain %d, the derived constraints remain %d, and the CKM/Jarlskog native registry write is blocked.", a.Sieve.RanksObserved, a.Constraints.RequiredIndependentConstraints, a.Constraints.DerivedIndependentConstraints)
}

func diag(xs [3]float64) matrix3 {
	var m matrix3
	for i, x := range xs {
		m[i][i] = complex(x, 0)
	}
	return m
}

func identity() matrix3 {
	return matrix3{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
}

func rotation12(theta float64) matrix3 {
	c, s := math.Cos(theta), math.Sin(theta)
	return matrix3{{complex(c, 0), complex(s, 0), 0}, {complex(-s, 0), complex(c, 0), 0}, {0, 0, 1}}
}

func fourier() matrix3 {
	omega := cmplx.Exp(complex(0, 2*math.Pi/3))
	n := complex(1/math.Sqrt(3), 0)
	return matrix3{
		{n, n, n},
		{n, n * omega, n * omega * omega},
		{n, n * omega * omega, n * omega},
	}
}

func conjugate(a, u matrix3) matrix3 { return mul(mul(u, a), dagger(u)) }

func dagger(a matrix3) matrix3 {
	var out matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out[i][j] = cmplx.Conj(a[j][i])
		}
	}
	return out
}

func mul(a, b matrix3) matrix3 {
	var out matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var s complex128
			for k := 0; k < 3; k++ {
				s += a[i][k] * b[k][j]
			}
			out[i][j] = s
		}
	}
	return out
}

func commutator(a, b matrix3) matrix3 {
	ab, ba := mul(a, b), mul(b, a)
	var out matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out[i][j] = ab[i][j] - ba[i][j]
		}
	}
	return out
}

func frobenius(a matrix3) float64 {
	var s float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			x := cmplx.Abs(a[i][j])
			s += x * x
		}
	}
	return math.Sqrt(s)
}

func det(a matrix3) complex128 {
	return a[0][0]*(a[1][1]*a[2][2]-a[1][2]*a[2][1]) - a[0][1]*(a[1][0]*a[2][2]-a[1][2]*a[2][0]) + a[0][2]*(a[1][0]*a[2][1]-a[1][1]*a[2][0])
}

func rank(a matrix3) int {
	if cmplx.Abs(det(a)) > rankTolerance {
		return 3
	}
	maxMinor := 0.0
	for r1 := 0; r1 < 3; r1++ {
		for r2 := r1 + 1; r2 < 3; r2++ {
			for c1 := 0; c1 < 3; c1++ {
				for c2 := c1 + 1; c2 < 3; c2++ {
					minor := a[r1][c1]*a[r2][c2] - a[r1][c2]*a[r2][c1]
					if x := cmplx.Abs(minor); x > maxMinor {
						maxMinor = x
					}
				}
			}
		}
	}
	if maxMinor > rankTolerance {
		return 2
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if cmplx.Abs(a[i][j]) > rankTolerance {
				return 1
			}
		}
	}
	return 0
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate485Koide=%t Gate486BridgeOnly=%t Gate486Blocked=%t required=%d derived=%d observedCKM=%t clean=%t", x.Gate485KoideBaselineInherited, x.Gate486NullMirrorBridgeOnly, x.Gate486NativeCKMTheoremBlocked, x.Gate486RequiredInvariantEquations, x.Gate486DerivedInvariantEquations, !x.NoObservedCKMImported, x.NativeRegistryClean)
}

func FormatSpectrum(x NullSpectrumAnsatz) string {
	return fmt.Sprintf("S=%.6g R=%.6g psi=%.6g R/S=%.12g q=%.3g x=[%.6g %.6g %.6g] eigenbasis_native=%t observed_masses=%t", x.S, x.R, x.Psi, x.RatioRoverS, x.MinkowskiResidual, x.EigenShadow[0], x.EigenShadow[1], x.EigenShadow[2], x.EigenvectorsSpecifiedNative, x.ObservedMassesImported)
}

func FormatOperators(x OperatorAudit) string {
	return fmt.Sprintf("native_U=%t native_D=%t diagonalizers=%t synthetic=%t same_null_spectrum=%t spectrum=%t eigenbasis=%t CKM=%t", x.NativeUpOperatorDerived, x.NativeDownOperatorDerived, x.NativeDiagonalizersDerived, x.CandidateSyntheticOperatorsUsed, x.OperatorsShareNullSpectrum, x.NullBoundaryConstrainsSpectrum, x.NullBoundaryConstrainsEigenbasis, x.CKMAsUuDaggerUdConstructed)
}

func FormatSieve(x CommutatorSieve) string {
	return fmt.Sprintf("cases=%d ranks=%v variable=%t same_spectrum=%t rank_suppressed=%t J_locked=%t polynomial=%t", x.CaseCount, x.RanksObserved, x.RankVariabilityObserved, x.SharedNullSpectrumInEveryCase, x.CommutatorRankSuppressedByNull, x.JarlskogDeterminantLocked, x.InvariantPolynomialProduced)
}

func FormatConstraints(x ConstraintHunt) string {
	return fmt.Sprintf("CKMdim=%d proposed=%d required=%d derived=%d moduli=%d J=%d traces=%d det_relation=%t passed=%t", x.PhysicalCKMParameterDim, x.ProposedCompressedDim, x.RequiredIndependentConstraints, x.DerivedIndependentConstraints, x.ModuliPolynomialRelations, x.JarlskogPolynomialRelations, x.CommutatorTraceRelations, x.CommutatorDeterminantRelation, x.TwoConstraintTheoremPassed)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("observed_CKM=%t wolfenstein=%t quark_masses=%t CP=%t CKM_native=%t J_native=%t constraints_write=%t synthetic_bridge=%t native_write=%t dims=(%d,%d)", x.ObservedCKMImported, x.ObservedWolfensteinImported, x.ObservedQuarkMassesImported, x.ObservedCPPhaseImported, x.CKMMatrixNativePrediction, x.JarlskogNativePrediction, x.PolynomialConstraintsNativeWrite, x.SyntheticCommutatorBridgeOnly, x.NativeRegistryWritten, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 487 Registry Audit — CKM Rephasing-Invariant Polynomial Constraint Search\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("```text\n")
	b.WriteString(StatusGate486Inherited + "\n")
	b.WriteString(StatusNullC3SpectrumAnsatzConstructed + "\n")
	b.WriteString(StatusCommutatorSieveExecuted + "\n")
	b.WriteString(StatusCommutatorRankNotSuppressed + "\n")
	b.WriteString(StatusNoJarlskogPolynomialDerived + "\n")
	b.WriteString(StatusNoRephasingInvariantConstraintsDerived + "\n")
	b.WriteString(StatusNativeUpDownOperatorsStillAbsent + "\n")
	b.WriteString(StatusNullSpectrumDoesNotDetermineEigenbasis + "\n")
	b.WriteString(StatusFirewallBlockedCKMPolynomialRegistryWrite + "\n")
	b.WriteString("```\n\n")
	b.WriteString("Gate 487 rejects the proposed native CKM commutator compression. The null-C3 boundary fixes a spectrum-shape baseline, but it does not fix the up/down eigenbasis mismatch from which physical CKM and Jarlskog invariants arise.\n\n")

	b.WriteString("## Inherited boundary\n\n")
	b.WriteString("Gate 485 remains the only accepted native shape statement in this lane:\n\n")
	b.WriteString("```text\n")
	b.WriteString("3S² - (3/2)R² = 0  ⇒  R/S = sqrt(2)  ⇒  Q = 2/3\n")
	b.WriteString("```\n\n")
	b.WriteString("Gate 486 then blocked CKM 4→2 because a null-mirror coordinate chart did not supply two rephasing-invariant CKM constraints. Gate 487 inherits that demand and tests the commutator route without CKM, Wolfenstein, quark-mass, or CP-phase inputs.\n\n")
	b.WriteString("| inherited object | status |\n|---|---:|\n")
	b.WriteString(fmt.Sprintf("| Gate 485 Koide baseline inherited | `%t` |\n", a.Inheritance.Gate485KoideBaselineInherited))
	b.WriteString(fmt.Sprintf("| Gate 486 null mirror bridge-only | `%t` |\n", a.Inheritance.Gate486NullMirrorBridgeOnly))
	b.WriteString(fmt.Sprintf("| required invariant equations from Gate 486 | `%d` |\n", a.Inheritance.Gate486RequiredInvariantEquations))
	b.WriteString(fmt.Sprintf("| inherited derived invariant equations | `%d` |\n", a.Inheritance.Gate486DerivedInvariantEquations))
	b.WriteString(fmt.Sprintf("| observed CKM imported | `%t` |\n\n", !a.Inheritance.NoObservedCKMImported))

	b.WriteString("## Algebraic commutator sieve\n\n")
	b.WriteString("The audit constructs only a synthetic theorem probe: two Hermitian operators with the same null-C3 spectrum, then changes their relative eigenbasis by a bridge-selected unitary. This is allowed as a negative sieve, not as a native physical model.\n\n")
	b.WriteString("```text\n")
	b.WriteString("O_u = diag(x_1,x_2,x_3)\n")
	b.WriteString("O_d = U diag(x_1,x_2,x_3) U†\n")
	b.WriteString("x_i = S + R cos(theta_i - psi),  R/S = sqrt(2)\n")
	b.WriteString("C_ud = [O_u,O_d]\n")
	b.WriteString("```\n\n")
	b.WriteString("| null-spectrum datum | value |\n|---|---:|\n")
	b.WriteString(fmt.Sprintf("| S | `%.12g` |\n", a.Spectrum.S))
	b.WriteString(fmt.Sprintf("| R | `%.12g` |\n", a.Spectrum.R))
	b.WriteString(fmt.Sprintf("| ψ | `%.12g` |\n", a.Spectrum.Psi))
	b.WriteString(fmt.Sprintf("| R/S | `%.12g` |\n", a.Spectrum.RatioRoverS))
	b.WriteString(fmt.Sprintf("| null residual `3S²-(3/2)R²` | `%.3g` |\n", a.Spectrum.MinkowskiResidual))
	b.WriteString(fmt.Sprintf("| native eigenbasis supplied | `%t` |\n\n", a.Spectrum.EigenvectorsSpecifiedNative))

	b.WriteString("## Commutator rank result\n\n")
	b.WriteString("All cases preserve the same null-C3 spectrum. Only the relative eigenbasis changes. The commutator rank changes anyway.\n\n")
	b.WriteString("| case | relative frame | rank `[O_u,O_d]` | Frobenius norm | abs det | result |\n|---|---|---:|---:|---:|---|\n")
	for _, c := range a.Sieve.Cases {
		b.WriteString(fmt.Sprintf("| %s | `%s` | `%d` | `%.12g` | `%.12g` | %s |\n", c.Name, c.RelativeUnitary, c.CommutatorRank, c.CommutatorFrobeniusNorm, c.CommutatorDeterminantAbs, c.Demonstrates))
	}
	b.WriteString("\n")
	b.WriteString("Therefore the shared null cone does not suppress the commutator rank. It permits commuting, rank-two, and full-rank synthetic brackets under the same spectrum. That kills the proposed native implication from null baseline to CKM polynomial compression.\n\n")

	b.WriteString("## Rephasing-invariant polynomial hunt\n\n")
	b.WriteString("A genuine CKM 4→2 theorem must produce two independent relations in the physical rephasing quotient, not merely a coordinate chart or a synthetic commutator sample. Gate 487 derives none.\n\n")
	b.WriteString("| invariant requirement | count/status |\n|---|---:|\n")
	b.WriteString(fmt.Sprintf("| physical CKM parameter dimension | `%d` |\n", a.Constraints.PhysicalCKMParameterDim))
	b.WriteString(fmt.Sprintf("| proposed compressed dimension | `%d` |\n", a.Constraints.ProposedCompressedDim))
	b.WriteString(fmt.Sprintf("| required independent polynomial constraints | `%d` |\n", a.Constraints.RequiredIndependentConstraints))
	b.WriteString(fmt.Sprintf("| derived independent polynomial constraints | `%d` |\n", a.Constraints.DerivedIndependentConstraints))
	b.WriteString(fmt.Sprintf("| moduli polynomial relations derived | `%d` |\n", a.Constraints.ModuliPolynomialRelations))
	b.WriteString(fmt.Sprintf("| Jarlskog polynomial relations derived | `%d` |\n", a.Constraints.JarlskogPolynomialRelations))
	b.WriteString(fmt.Sprintf("| commutator determinant relation derived | `%t` |\n", a.Constraints.CommutatorDeterminantRelation))
	b.WriteString(fmt.Sprintf("| 4→2 theorem passed | `%t` |\n\n", a.Constraints.TwoConstraintTheoremPassed))

	b.WriteString("## Firewall result\n\n")
	b.WriteString("```text\n")
	b.WriteString(StatusFirewallBlockedCKMPolynomialRegistryWrite + "\n")
	b.WriteString(StatusEmpiricalCKMFitRejected + "\n")
	b.WriteString("```\n\n")
	b.WriteString("No CKM matrix entries, Wolfenstein parameters, quark masses, or physical CP phase were imported. The synthetic commutator probes are bridge-only counterexamples to overclaiming; they are not native predictions.\n\n")
	b.WriteString("| firewall item | status |\n|---|---:|\n")
	b.WriteString(fmt.Sprintf("| observed CKM imported | `%t` |\n", a.Firewall.ObservedCKMImported))
	b.WriteString(fmt.Sprintf("| Wolfenstein imported | `%t` |\n", a.Firewall.ObservedWolfensteinImported))
	b.WriteString(fmt.Sprintf("| quark masses imported | `%t` |\n", a.Firewall.ObservedQuarkMassesImported))
	b.WriteString(fmt.Sprintf("| observed CP phase imported | `%t` |\n", a.Firewall.ObservedCPPhaseImported))
	b.WriteString(fmt.Sprintf("| Jarlskog native prediction | `%t` |\n", a.Firewall.JarlskogNativePrediction))
	b.WriteString(fmt.Sprintf("| CKM 4→2 native write | `%t` |\n", a.Firewall.CKMFourToTwoNativeWritten))
	b.WriteString(fmt.Sprintf("| polynomial constraints native write | `%t` |\n", a.Firewall.PolynomialConstraintsNativeWrite))
	b.WriteString(fmt.Sprintf("| synthetic commutator bridge-only | `%t` |\n", a.Firewall.SyntheticCommutatorBridgeOnly))
	b.WriteString(fmt.Sprintf("| native flavor dimension | `%d` |\n", a.Firewall.NativeFlavorDimAfter))
	b.WriteString(fmt.Sprintf("| K/X/Y charged coefficient dimension | `%d` |\n\n", a.Firewall.KXYCoeffDimAfter))

	b.WriteString("## Registry update\n\n")
	writeList := func(title string, xs []string) {
		b.WriteString(title + "\n\n")
		for _, x := range xs {
			b.WriteString("- `" + x + "`\n")
		}
		b.WriteString("\n")
	}
	writeList("Native", a.Registry.NativeEntries)
	writeList("Bridge", a.Registry.BridgeEntries)
	writeList("Environmental", a.Registry.EnvironmentalEntries)
	writeList("Failed route", a.Registry.FailedRoutes)
	writeList("Open theorem", a.Registry.OpenTheorems)

	b.WriteString("## Next step\n\n")
	b.WriteString(fmt.Sprintf("Gate %d — %s. %s\n\n", a.Next.Gate, a.Next.Title, a.Next.PrimaryTask))

	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}
