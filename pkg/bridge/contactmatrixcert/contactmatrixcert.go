// Package contactmatrixcert implements Gate 149: exact rational
// contact-overlap matrix lift / determinant certificate search.
//
// Gate 148 reconstructed a rational characteristic-polynomial candidate from
// numerical contact overlap eigenvalues. Gate 149 removes the largest remaining
// numerical weakness: the contact overlap matrix itself is lifted to an exact
// rational 14×14 matrix.
//
// The lift uses only finite data already present in the project:
//
//   - the Boolean incidence matrix M between Λ³ and Λ⁴;
//   - the closed-form rational inverse of the Boolean Gram matrix
//     G=MᵀM, where G^{-1}_{AB} depends only on |A∩B|;
//   - the integer G₂ raw calibration columns R with RᵀR=4I, so Q_G=R/2.
//
// Thus
//
//	Ω_exact = Q_Gᵀ P_B Q_G
//	        = 1/4 · (MᵀR)ᵀ · (MᵀM)^{-1} · (MᵀR).
//
// Gate 149 then computes the characteristic polynomial by exact rational
// Faddeev--LeVerrier arithmetic, verifies the determinant and trace, and checks
// that the Gate-148 candidate polynomial is the exact characteristic
// polynomial. This is a genuine exact-matrix and determinant/charpoly
// certificate. It still does not produce contact charge semantics, local fields,
// mass activation, decoupling, or threshold beta permission.
package contactmatrixcert

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactcharpoly"
	"github.com/bagherbal/asha-engine/pkg/combinatorics"
	geomcontact "github.com/bagherbal/asha-engine/pkg/geometry/contact"
	"github.com/bagherbal/asha-engine/pkg/geometry/g2"
)

type GramInverseCoefficient struct {
	Intersection int
	Coefficient  string
	Verified     bool
}

type ExactMatrixLift struct {
	Rows                     int
	Cols                     int
	RawG2Rows                int
	RawG2Cols                int
	BooleanLowerRows         int
	BooleanUpperRows         int
	RawG2GramIsFourIdentity  bool
	BooleanGramInverseClosed bool
	RationalMatrixBuilt      bool
	Symmetric                bool
	Denominators             []int64
	MaxFloatResidual         float64
	Trace                    string
	Determinant              string
	RankOmegaMinusIdentity   int
	UnitEigenspaceDimension  int
	Verdict                  string
}

type CharacteristicCertificate struct {
	ExactCharpolyComputed                bool
	CandidateCharpolyMatches             bool
	CandidatePolynomialAnnihilatesMatrix bool
	Degree                               int
	Coefficients                         []string
	TraceMatches                         bool
	DeterminantMatches                   bool
	UnitEigenMultiplicity                int
	PartialDegree                        int
	RationalFactors                      int
	QuarticFactorCertified               bool
	Verdict                              string
}

type ConstructionRequirements struct {
	ExactBooleanProjectorFormula  bool
	ExactG2RawColumns             bool
	ExactRationalOverlapMatrix    bool
	ExactCharacteristicPolynomial bool
	ExactDeterminantCertificate   bool
	ExactAnnihilationCertificate  bool
	RootIsolationCertificate      bool
	RowwiseRootAssignmentProof    bool
	ChargeOperatorSelected        bool
	RepresentationRowsSelected    bool
	ObservedInputFree             bool
	AllSatisfiedForPhysics        bool
	Verdict                       string
}

type Summary struct {
	ContactRows                  int
	ExactMatrixDimension         int
	ExactMatrixBuilt             bool
	ExactCharpolyCertificates    int
	ExactDeterminantCertificates int
	UnitEigenMultiplicity        int
	PartialDegree                int
	CandidateNumberFieldDegree   int
	RootIsolationCertificates    int
	RowAssignmentProofs          int
	ChargeSemanticRows           int
	RepresentationCompleteRows   int
	RepresentationOpenRows       int
	ContactBetaRowsAllowed       int
	ContactZeroRowsProved        int
	ResidualS6Choices            int
	ResidualNullityBefore        int
	ResidualNullityAfter         int
}

type Analysis struct {
	Previous contactcharpoly.Analysis

	GramInverse  []GramInverseCoefficient
	MatrixLift   ExactMatrixLift
	Certificate  CharacteristicCertificate
	Requirements ConstructionRequirements
	Summary      Summary

	ContactRows                   int
	ExactRationalOverlapMatrix    bool
	ExactDeterminantComputed      bool
	ExactCharacteristicCertified  bool
	ExactAnnihilationCertified    bool
	ExactNumberFieldLiftCertified bool
	RootIsolationCertificates     int
	RowwiseRootAssignmentProofs   int
	ChargeSemanticRows            int
	T3RRowsDerived                int
	ChiralityRowsDerived          int
	BMinusLRowsDerived            int
	SU2LRowsDerived               int
	HyperchargeRowsDerived        int
	RepresentationCompleteRows    int
	RepresentationOpenRows        int
	ContactBetaRowsAllowed        int
	ContactZeroRowsProved         int
	BetaPermissionFirewallClosed  bool
	ThresholdCorrectedBeta        bool
	FullBetaMatchingTensor        bool
	ResidualS6Choices             int
	ResidualNullityBefore         int
	ResidualNullityAfter          int
	HiddenObservedInputUsed       bool
	PhysicalWeakAngleDerived      bool
	FineStructureDerived          bool
	PhysicalMassesDerived         bool
	PhysicalScaleDerived          bool

	TruthStatement      string
	RejectedClaims      []string
	RemainingUnknowns   []string
	RecommendedNextGate string
}

var defaultOnce sync.Once
var defaultValue Analysis
var defaultErr error

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := contactcharpoly.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		contact, err := geomcontact.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		g, err := g2.BuildCalibrationSupport()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, contact, g)
	})
	return defaultValue, defaultErr
}

func Build(prev contactcharpoly.Analysis, contact geomcontact.Space, g g2.CalibrationSupport) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.ContactRows != 7 || prev.CandidateCoveredRows != 7 || !prev.Candidate.CandidateFactorizationRecognized {
		return Analysis{}, fmt.Errorf("Gate 149 requires Gate 148 closed-firewall seven-row charpoly candidate")
	}
	if prev.ExactCharacteristicCertified || prev.ExactNumberFieldLiftCertified || prev.ContactBetaRowsAllowed != 0 || prev.ChargeSemanticRows != 0 || prev.RepresentationCompleteRows != 0 {
		return Analysis{}, fmt.Errorf("Gate 149 expects Gate 148 to have no exact certificate or contact charge semantics yet")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 149 refuses hidden observed physical input")
	}

	omega, gramCoeffs, rawGramOK, gramInverseOK, err := buildExactOverlap(g)
	if err != nil {
		return Analysis{}, err
	}
	symmetric := matrixEqual(omega, transpose(omega))
	dens := denominators(omega)
	floatResidual := maxResidualAgainstFloat(omega, contact.Overlap)
	trace := matrixTrace(omega)
	det := determinant(omega)
	rankMinusI := rank(matrixSub(omega, identity(len(omega))))
	unitMultiplicity := len(omega) - rankMinusI

	charCoeffs := charpolyFaddeevLeVerrier(omega)
	candidateCoeffs := candidateFullPolynomialCoefficients()
	candidateMatches := ratSliceEqual(charCoeffs, candidateCoeffs)
	annihilates := isZeroMatrix(evalPolynomialMatrix(candidateCoeffs, omega))
	traceMatches := ratEqual(trace, ratNeg(charCoeffs[1]))
	determinantMatches := ratEqual(det, charCoeffs[len(charCoeffs)-1])

	matrixLift := ExactMatrixLift{
		Rows:                     len(omega),
		Cols:                     len(omega),
		RawG2Rows:                g.RawColumns.Rows(),
		RawG2Cols:                g.RawColumns.Cols(),
		BooleanLowerRows:         56,
		BooleanUpperRows:         70,
		RawG2GramIsFourIdentity:  rawGramOK,
		BooleanGramInverseClosed: gramInverseOK,
		RationalMatrixBuilt:      true,
		Symmetric:                symmetric,
		Denominators:             dens,
		MaxFloatResidual:         floatResidual,
		Trace:                    ratString(trace),
		Determinant:              ratString(det),
		RankOmegaMinusIdentity:   rankMinusI,
		UnitEigenspaceDimension:  unitMultiplicity,
		Verdict:                  "exact rational Ω=1/4*(M^T R)^T*(M^T M)^-1*(M^T R) is constructed without numerical eigensolver input",
	}

	cert := CharacteristicCertificate{
		ExactCharpolyComputed:                true,
		CandidateCharpolyMatches:             candidateMatches,
		CandidatePolynomialAnnihilatesMatrix: annihilates,
		Degree:                               len(charCoeffs) - 1,
		Coefficients:                         ratStrings(charCoeffs),
		TraceMatches:                         traceMatches,
		DeterminantMatches:                   determinantMatches,
		UnitEigenMultiplicity:                unitMultiplicity,
		PartialDegree:                        7,
		RationalFactors:                      3,
		QuarticFactorCertified:               candidateMatches && annihilates,
		Verdict:                              "exact rational Faddeev-LeVerrier characteristic polynomial equals the Gate-148 candidate and annihilates Ω exactly",
	}

	exactOK := matrixLift.RationalMatrixBuilt && matrixLift.Symmetric && matrixLift.RawG2GramIsFourIdentity && matrixLift.BooleanGramInverseClosed && cert.ExactCharpolyComputed && cert.CandidateCharpolyMatches && cert.CandidatePolynomialAnnihilatesMatrix && cert.TraceMatches && cert.DeterminantMatches && cert.UnitEigenMultiplicity == 7 && cert.Degree == 14

	req := ConstructionRequirements{
		ExactBooleanProjectorFormula:  gramInverseOK,
		ExactG2RawColumns:             rawGramOK,
		ExactRationalOverlapMatrix:    matrixLift.RationalMatrixBuilt && matrixLift.Symmetric,
		ExactCharacteristicPolynomial: cert.ExactCharpolyComputed && cert.CandidateCharpolyMatches,
		ExactDeterminantCertificate:   cert.DeterminantMatches,
		ExactAnnihilationCertificate:  cert.CandidatePolynomialAnnihilatesMatrix,
		RootIsolationCertificate:      false,
		RowwiseRootAssignmentProof:    false,
		ChargeOperatorSelected:        false,
		RepresentationRowsSelected:    false,
		ObservedInputFree:             true,
		AllSatisfiedForPhysics:        false,
		Verdict:                       "exact matrix and determinant/characteristic certificates now exist, but root isolation, row-wise assignment, charge semantics, representation rows, local fields, mass activation, and decoupling are still missing",
	}

	summary := Summary{
		ContactRows:                  prev.ContactRows,
		ExactMatrixDimension:         matrixLift.Rows,
		ExactMatrixBuilt:             matrixLift.RationalMatrixBuilt,
		ExactCharpolyCertificates:    boolInt(cert.ExactCharpolyComputed && cert.CandidateCharpolyMatches),
		ExactDeterminantCertificates: boolInt(cert.DeterminantMatches),
		UnitEigenMultiplicity:        unitMultiplicity,
		PartialDegree:                cert.PartialDegree,
		CandidateNumberFieldDegree:   prev.CandidateNumberFieldDegree,
		RootIsolationCertificates:    0,
		RowAssignmentProofs:          0,
		ChargeSemanticRows:           0,
		RepresentationCompleteRows:   0,
		RepresentationOpenRows:       prev.RepresentationOpenRows,
		ContactBetaRowsAllowed:       0,
		ContactZeroRowsProved:        0,
		ResidualS6Choices:            prev.ResidualS6Choices,
		ResidualNullityBefore:        prev.ResidualNullityAfter,
		ResidualNullityAfter:         prev.ResidualNullityAfter,
	}

	truth := "Gate 149 upgrades the contact spectrum from a numerical/rationally reconstructed polynomial candidate to an exact rational matrix certificate. The Boolean projector has a closed rational inverse, the G2 calibration columns are exact integer columns with Gram 4I, and Ω=Q_G^T P_B Q_G is built as an exact rational 14×14 matrix. Exact rational Faddeev-LeVerrier arithmetic gives the same characteristic polynomial as Gate 148, including seven unit eigenvalues and the degree-seven partial factor. This certifies the determinant/charpoly candidate, but it still does not provide root isolation, row-wise eigenprojector assignment, charge semantics, local fields, mass activation, decoupling, threshold beta rows, or physical constants."

	return Analysis{
		Previous:                      prev,
		GramInverse:                   gramCoeffs,
		MatrixLift:                    matrixLift,
		Certificate:                   cert,
		Requirements:                  req,
		Summary:                       summary,
		ContactRows:                   prev.ContactRows,
		ExactRationalOverlapMatrix:    exactOK,
		ExactDeterminantComputed:      cert.DeterminantMatches,
		ExactCharacteristicCertified:  cert.ExactCharpolyComputed && cert.CandidateCharpolyMatches,
		ExactAnnihilationCertified:    cert.CandidatePolynomialAnnihilatesMatrix,
		ExactNumberFieldLiftCertified: exactOK,
		RootIsolationCertificates:     0,
		RowwiseRootAssignmentProofs:   0,
		ChargeSemanticRows:            0,
		T3RRowsDerived:                0,
		ChiralityRowsDerived:          0,
		BMinusLRowsDerived:            0,
		SU2LRowsDerived:               0,
		HyperchargeRowsDerived:        0,
		RepresentationCompleteRows:    0,
		RepresentationOpenRows:        prev.RepresentationOpenRows,
		ContactBetaRowsAllowed:        0,
		ContactZeroRowsProved:         0,
		BetaPermissionFirewallClosed:  true,
		ThresholdCorrectedBeta:        false,
		FullBetaMatchingTensor:        false,
		ResidualS6Choices:             prev.ResidualS6Choices,
		ResidualNullityBefore:         prev.ResidualNullityAfter,
		ResidualNullityAfter:          prev.ResidualNullityAfter,
		HiddenObservedInputUsed:       false,
		PhysicalWeakAngleDerived:      false,
		FineStructureDerived:          false,
		PhysicalMassesDerived:         false,
		PhysicalScaleDerived:          false,
		TruthStatement:                truth,
		RejectedClaims: []string{
			"exact characteristic polynomial implies contact charge semantics",
			"quartic number-field roots may be used as hypercharge rows",
			"root labels may be assigned to physical rows without root-isolation/eigenprojector proof",
			"threshold beta rows open once exact spectral algebra is certified",
			"observed constants may be used to orient or normalize the contact spectrum",
		},
		RemainingUnknowns: []string{
			"root-isolation certificate for the quartic factor",
			"row-wise exact eigenprojector/minimal-polynomial assignment",
			"semantic map from exact roots to contact charge or representation rows",
			"local field variables, mass activation, and decoupling",
			"threshold-corrected beta tensor and physical-flow selector",
		},
		RecommendedNextGate: "Gate 150 — exact contact root-isolation / row-wise eigenprojector assignment theorem",
	}, nil
}

func buildExactOverlap(g g2.CalibrationSupport) ([][]*big.Rat, []GramInverseCoefficient, bool, bool, error) {
	if g.RawColumns.Rows() != 70 || g.RawColumns.Cols() != 14 {
		return nil, nil, false, false, fmt.Errorf("unexpected G2 raw column dimensions %dx%d", g.RawColumns.Rows(), g.RawColumns.Cols())
	}
	raw := make([][]int64, g.RawColumns.Rows())
	for i := range raw {
		raw[i] = make([]int64, g.RawColumns.Cols())
		for j := range raw[i] {
			v := g.RawColumns.At(i, j)
			r := math.Round(v)
			if math.Abs(v-r) > 1e-12 {
				return nil, nil, false, false, fmt.Errorf("G2 raw column entry is not integral at %d,%d: %.16g", i, j, v)
			}
			raw[i][j] = int64(r)
		}
	}
	rawGramOK := true
	for i := 0; i < 14; i++ {
		for j := 0; j < 14; j++ {
			sum := int64(0)
			for r := 0; r < 70; r++ {
				sum += raw[r][i] * raw[r][j]
			}
			want := int64(0)
			if i == j {
				want = 4
			}
			if sum != want {
				rawGramOK = false
			}
		}
	}

	lower, err := combinatorics.Subsets(8, 3)
	if err != nil {
		return nil, nil, false, false, err
	}
	upper, err := combinatorics.Subsets(8, 4)
	if err != nil {
		return nil, nil, false, false, err
	}
	if len(lower) != 56 || len(upper) != 70 {
		return nil, nil, false, false, fmt.Errorf("unexpected Boolean basis dimensions")
	}

	// A = M^T R, shape 56×14.
	a := make([][]int64, len(lower))
	for i := range a {
		a[i] = make([]int64, 14)
	}
	for u, up := range upper {
		for l, low := range lower {
			if up.ContainsAll(low) {
				for c := 0; c < 14; c++ {
					a[l][c] += raw[u][c]
				}
			}
		}
	}

	omega := newRatMatrix(14, 14)
	for i := 0; i < 14; i++ {
		for j := 0; j < 14; j++ {
			acc := new(big.Rat)
			for l1, s1 := range lower {
				if a[l1][i] == 0 {
					continue
				}
				for l2, s2 := range lower {
					if a[l2][j] == 0 {
						continue
					}
					coeff := gramInverseByIntersection(intersectionSize(s1, s2))
					term := new(big.Rat).SetInt64(a[l1][i] * a[l2][j])
					term.Mul(term, coeff)
					acc.Add(acc, term)
				}
			}
			acc.Mul(acc, rat(1, 4))
			omega[i][j] = acc
		}
	}

	// Verify the closed inverse formula against G=M^T M without inverting.
	gramInverseOK := verifyBooleanGramInverse(lower)
	coeffs := []GramInverseCoefficient{
		{Intersection: 3, Coefficient: "77/240", Verified: gramInverseOK},
		{Intersection: 2, Coefficient: "-29/720", Verified: gramInverseOK},
		{Intersection: 1, Coefficient: "11/720", Verified: gramInverseOK},
		{Intersection: 0, Coefficient: "-1/80", Verified: gramInverseOK},
	}
	return omega, coeffs, rawGramOK, gramInverseOK, nil
}

func gramInverseByIntersection(k int) *big.Rat {
	switch k {
	case 3:
		return rat(77, 240)
	case 2:
		return rat(-29, 720)
	case 1:
		return rat(11, 720)
	case 0:
		return rat(-1, 80)
	default:
		panic("invalid intersection")
	}
}

func verifyBooleanGramInverse(lower []combinatorics.Subset) bool {
	for i, si := range lower {
		for j, sj := range lower {
			acc := new(big.Rat)
			for _, sk := range lower {
				g := booleanGramEntry(si, sk)
				if g == 0 {
					continue
				}
				term := new(big.Rat).SetInt64(int64(g))
				term.Mul(term, gramInverseByIntersection(intersectionSize(sk, sj)))
				acc.Add(acc, term)
			}
			want := rat(0, 1)
			if i == j {
				want = rat(1, 1)
			}
			if acc.Cmp(want) != 0 {
				return false
			}
		}
	}
	return true
}

func booleanGramEntry(a, b combinatorics.Subset) int {
	if a.Key() == b.Key() {
		return 5
	}
	if intersectionSize(a, b) == 2 {
		return 1
	}
	return 0
}

func intersectionSize(a, b combinatorics.Subset) int {
	count := 0
	for _, x := range a {
		for _, y := range b {
			if x == y {
				count++
				break
			}
		}
	}
	return count
}

func charpolyFaddeevLeVerrier(a [][]*big.Rat) []*big.Rat {
	n := len(a)
	coeffs := make([]*big.Rat, n+1)
	coeffs[0] = rat(1, 1)
	b := identity(n)
	for k := 1; k <= n; k++ {
		ab := matMul(a, b)
		tr := matrixTrace(ab)
		ck := new(big.Rat).Neg(tr)
		ck.Quo(ck, rat(int64(k), 1))
		coeffs[k] = ck
		b = matAdd(ab, scalarIdentity(n, ck))
	}
	return coeffs
}

func candidateFullPolynomialCoefficients() []*big.Rat {
	// Coefficients of det(xI-Ω), high to low:
	// (x-1)^7*(2x-1)*(3x-2)*(3x-1)*(3240x^4-7668x^3+6426x^2-2235x+271)/58320.
	return []*big.Rat{
		rat(1, 1), rat(-163, 15), rat(4889, 90), rat(-22339, 135), rat(2222057, 6480),
		rat(-9933973, 19440), rat(16453811, 29160), rat(-42173, 90), rat(1713497, 5832),
		rat(-806411, 5832), rat(233417, 4860), rat(-43468, 3645), rat(116557, 58320),
		rat(-3929, 19440), rat(271, 29160),
	}
}

func evalPolynomialMatrix(coeffs []*big.Rat, a [][]*big.Rat) [][]*big.Rat {
	n := len(a)
	acc := newRatMatrix(n, n)
	for _, c := range coeffs {
		acc = matMul(acc, a)
		acc = matAdd(acc, scalarIdentity(n, c))
	}
	return acc
}

func determinant(m [][]*big.Rat) *big.Rat {
	n := len(m)
	a := cloneMatrix(m)
	det := rat(1, 1)
	sign := int64(1)
	for col := 0; col < n; col++ {
		pivot := -1
		for r := col; r < n; r++ {
			if a[r][col].Sign() != 0 {
				pivot = r
				break
			}
		}
		if pivot == -1 {
			return rat(0, 1)
		}
		if pivot != col {
			a[pivot], a[col] = a[col], a[pivot]
			sign *= -1
		}
		p := cloneRat(a[col][col])
		det.Mul(det, p)
		for r := col + 1; r < n; r++ {
			if a[r][col].Sign() == 0 {
				continue
			}
			factor := new(big.Rat).Quo(a[r][col], p)
			for c := col; c < n; c++ {
				term := new(big.Rat).Mul(factor, a[col][c])
				a[r][c].Sub(a[r][c], term)
			}
		}
	}
	if sign < 0 {
		det.Neg(det)
	}
	return det
}

func rank(m [][]*big.Rat) int {
	a := cloneMatrix(m)
	rows, cols := len(a), len(a[0])
	r := 0
	for c := 0; c < cols && r < rows; c++ {
		pivot := -1
		for i := r; i < rows; i++ {
			if a[i][c].Sign() != 0 {
				pivot = i
				break
			}
		}
		if pivot == -1 {
			continue
		}
		a[pivot], a[r] = a[r], a[pivot]
		p := cloneRat(a[r][c])
		for j := c; j < cols; j++ {
			a[r][j].Quo(a[r][j], p)
		}
		for i := 0; i < rows; i++ {
			if i == r || a[i][c].Sign() == 0 {
				continue
			}
			factor := cloneRat(a[i][c])
			for j := c; j < cols; j++ {
				term := new(big.Rat).Mul(factor, a[r][j])
				a[i][j].Sub(a[i][j], term)
			}
		}
		r++
	}
	return r
}

func maxResidualAgainstFloat(omega [][]*big.Rat, m interface{ At(int, int) float64 }) float64 {
	max := 0.0
	for i := range omega {
		for j := range omega[i] {
			f, _ := omega[i][j].Float64()
			d := math.Abs(f - m.At(i, j))
			if d > max {
				max = d
			}
		}
	}
	return max
}

func newRatMatrix(r, c int) [][]*big.Rat {
	m := make([][]*big.Rat, r)
	for i := range m {
		m[i] = make([]*big.Rat, c)
		for j := range m[i] {
			m[i][j] = rat(0, 1)
		}
	}
	return m
}

func identity(n int) [][]*big.Rat {
	m := newRatMatrix(n, n)
	for i := 0; i < n; i++ {
		m[i][i] = rat(1, 1)
	}
	return m
}

func scalarIdentity(n int, s *big.Rat) [][]*big.Rat {
	m := newRatMatrix(n, n)
	for i := 0; i < n; i++ {
		m[i][i] = cloneRat(s)
	}
	return m
}

func transpose(a [][]*big.Rat) [][]*big.Rat {
	r, c := len(a), len(a[0])
	out := newRatMatrix(c, r)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			out[j][i] = cloneRat(a[i][j])
		}
	}
	return out
}

func matMul(a, b [][]*big.Rat) [][]*big.Rat {
	r, k, c := len(a), len(a[0]), len(b[0])
	out := newRatMatrix(r, c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			acc := new(big.Rat)
			for t := 0; t < k; t++ {
				term := new(big.Rat).Mul(a[i][t], b[t][j])
				acc.Add(acc, term)
			}
			out[i][j] = acc
		}
	}
	return out
}

func matAdd(a, b [][]*big.Rat) [][]*big.Rat {
	out := newRatMatrix(len(a), len(a[0]))
	for i := range a {
		for j := range a[i] {
			out[i][j] = new(big.Rat).Add(a[i][j], b[i][j])
		}
	}
	return out
}

func matrixSub(a, b [][]*big.Rat) [][]*big.Rat {
	out := newRatMatrix(len(a), len(a[0]))
	for i := range a {
		for j := range a[i] {
			out[i][j] = new(big.Rat).Sub(a[i][j], b[i][j])
		}
	}
	return out
}

func matrixTrace(a [][]*big.Rat) *big.Rat {
	acc := new(big.Rat)
	for i := range a {
		acc.Add(acc, a[i][i])
	}
	return acc
}

func matrixEqual(a, b [][]*big.Rat) bool {
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		return false
	}
	for i := range a {
		for j := range a[i] {
			if a[i][j].Cmp(b[i][j]) != 0 {
				return false
			}
		}
	}
	return true
}

func isZeroMatrix(a [][]*big.Rat) bool {
	for i := range a {
		for j := range a[i] {
			if a[i][j].Sign() != 0 {
				return false
			}
		}
	}
	return true
}

func cloneMatrix(a [][]*big.Rat) [][]*big.Rat {
	out := newRatMatrix(len(a), len(a[0]))
	for i := range a {
		for j := range a[i] {
			out[i][j] = cloneRat(a[i][j])
		}
	}
	return out
}

func cloneRat(x *big.Rat) *big.Rat { return new(big.Rat).Set(x) }
func rat(n, d int64) *big.Rat      { return new(big.Rat).SetFrac(big.NewInt(n), big.NewInt(d)) }
func ratEqual(a, b *big.Rat) bool  { return a.Cmp(b) == 0 }
func ratNeg(a *big.Rat) *big.Rat   { return new(big.Rat).Neg(a) }

func ratSliceEqual(a, b []*big.Rat) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Cmp(b[i]) != 0 {
			return false
		}
	}
	return true
}

func denominators(a [][]*big.Rat) []int64 {
	seen := map[int64]bool{}
	for i := range a {
		for j := range a[i] {
			seen[a[i][j].Denom().Int64()] = true
		}
	}
	out := make([]int64, 0, len(seen))
	for d := range seen {
		out = append(out, d)
	}
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			if out[j] < out[i] {
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}

func ratString(r *big.Rat) string {
	if r.Denom().Cmp(big.NewInt(1)) == 0 {
		return r.Num().String()
	}
	return r.RatString()
}

func ratStrings(rs []*big.Rat) []string {
	out := make([]string, len(rs))
	for i, r := range rs {
		out[i] = ratString(r)
	}
	return out
}

func boolInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func FormatGramInverse(coeffs []GramInverseCoefficient) string {
	parts := make([]string, 0, len(coeffs))
	for _, c := range coeffs {
		parts = append(parts, fmt.Sprintf("|A∩B|=%d -> %s verified=%t", c.Intersection, c.Coefficient, c.Verified))
	}
	return strings.Join(parts, "; ")
}

func FormatMatrixLift(m ExactMatrixLift) string {
	return fmt.Sprintf("Ω=%dx%d rawG2=%dx%d Boolean=%d→%d rawGram4I=%t gramInvClosed=%t rational=%t symmetric=%t denoms=%v floatResidual=%.3e trace=%s det=%s rank(Ω-I)=%d unitDim=%d (%s)", m.Rows, m.Cols, m.RawG2Rows, m.RawG2Cols, m.BooleanLowerRows, m.BooleanUpperRows, m.RawG2GramIsFourIdentity, m.BooleanGramInverseClosed, m.RationalMatrixBuilt, m.Symmetric, m.Denominators, m.MaxFloatResidual, m.Trace, m.Determinant, m.RankOmegaMinusIdentity, m.UnitEigenspaceDimension, m.Verdict)
}

func FormatCertificate(c CharacteristicCertificate) string {
	return fmt.Sprintf("exactChar=%t candidateMatches=%t annihilates=%t degree=%d trace=%t det=%t unitMult=%d partialDegree=%d rationalFactors=%d quarticCertified=%t coeffs=[%s] (%s)", c.ExactCharpolyComputed, c.CandidateCharpolyMatches, c.CandidatePolynomialAnnihilatesMatrix, c.Degree, c.TraceMatches, c.DeterminantMatches, c.UnitEigenMultiplicity, c.PartialDegree, c.RationalFactors, c.QuarticFactorCertified, strings.Join(c.Coefficients, ", "), c.Verdict)
}

func FormatRequirements(r ConstructionRequirements) string {
	return fmt.Sprintf("booleanFormula=%t g2Exact=%t rationalMatrix=%t exactChar=%t detCert=%t annihilation=%t rootIsolation=%t rowProof=%t charge=%t reps=%t observedFree=%t physicsAll=%t (%s)", r.ExactBooleanProjectorFormula, r.ExactG2RawColumns, r.ExactRationalOverlapMatrix, r.ExactCharacteristicPolynomial, r.ExactDeterminantCertificate, r.ExactAnnihilationCertificate, r.RootIsolationCertificate, r.RowwiseRootAssignmentProof, r.ChargeOperatorSelected, r.RepresentationRowsSelected, r.ObservedInputFree, r.AllSatisfiedForPhysics, r.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contact=%d Ωdim=%d exactMatrix=%t charCert=%d detCert=%d unitMult=%d partialDegree=%d fieldDegreeCandidate=%d rootIso=%d rowProof=%d semantic=%d reps=%d/%d beta=%d nullity=%d→%d", s.ContactRows, s.ExactMatrixDimension, s.ExactMatrixBuilt, s.ExactCharpolyCertificates, s.ExactDeterminantCertificates, s.UnitEigenMultiplicity, s.PartialDegree, s.CandidateNumberFieldDegree, s.RootIsolationCertificates, s.RowAssignmentProofs, s.ChargeSemanticRows, s.RepresentationCompleteRows, s.RepresentationOpenRows, s.ContactBetaRowsAllowed, s.ResidualNullityBefore, s.ResidualNullityAfter)
}
