// Package contactcharpoly implements Gate 148: exact contact overlap
// characteristic polynomial / symbolic number-field construction attempt.
//
// Gate 147 refused to treat the non-rational-looking contact overlap rows as
// exact algebraic data because the project only had numerical eigenvalues. Gate
// 148 makes the strongest safe next move: rationally reconstruct a symbolic
// characteristic-polynomial candidate for the seven partial contact rows and
// test it against the finite contact spectrum.
//
// The candidate is strong and useful:
//
//	P_partial(x) = (2x-1)(3x-2)(3x-1)
//	               (3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271) / 58320
//
// It covers all seven partial-overlap rows numerically and isolates the four
// non-rational-looking rows as roots of a quartic candidate. However this gate
// deliberately does not certify the polynomial as exact. The current project
// still computes the overlap matrix with float64 arithmetic and does not derive
// an exact determinant over a rational/number-field matrix. Therefore the
// candidate opens a new symbolic path, but no charge, representation,
// hypercharge, local-field, mass, decoupling, or beta-permission theorem is
// unlocked.
package contactcharpoly

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactalgebraic"
)

type Rational struct {
	Num int64
	Den int64
}

type RowPolynomialAudit struct {
	Name                         string
	Value                        float64
	RationalDiagnostic           bool
	RationalRootFactor           string
	QuarticCandidateRoot         bool
	CandidateMinimalPolynomial   string
	PartialPolynomialResidual    float64
	QuarticResidual              float64
	CandidatePolynomialCoversRow bool
	CertifiedExactRow            bool
	ChargeSemantic               bool
	RepresentationSemantic       bool
	OpensBetaPermission          bool
	Verdict                      string
}

type CharacteristicPolynomialCandidate struct {
	PartialDegree                    int
	FullDegree                       int
	RationalDiagnosticFactors        []string
	QuarticFactor                    string
	PartialPolynomial                string
	FullPolynomialFactorization      string
	CandidateNumberFieldDegree       int
	RowsCovered                      int
	MaxPartialResidual               float64
	MaxQuarticResidual               float64
	CandidateFactorizationRecognized bool
	ExactMatrixOverNumberField       bool
	ExactDeterminantComputed         bool
	ExactCharacteristicCertified     bool
	RowMinimalPolynomialsCertified   bool
	Verdict                          string
}

type SymbolicConstructionRequirements struct {
	RationalReconstructionCandidate bool
	ResidualCheckPassed             bool
	ExactOverlapMatrix              bool
	ExactDeterminant                bool
	IndependentCertificate          bool
	RootIsolationCertificate        bool
	RowwiseRootAssignmentProof      bool
	AlgebraicRowSemantics           bool
	ChargeOperatorSelected          bool
	RepresentationRowsSelected      bool
	ObservedInputFree               bool
	AllSatisfied                    bool
	Verdict                         string
}

type Summary struct {
	ContactRows                int
	CandidateCoveredRows       int
	RationalFactorRows         int
	QuarticCandidateRows       int
	CandidateNumberFieldDegree int
	ExactCertifiedRows         int
	ExactCharacteristicProofs  int
	ChargeSemanticRows         int
	RepresentationCompleteRows int
	RepresentationOpenRows     int
	ContactBetaRowsAllowed     int
	ContactZeroRowsProved      int
	ResidualS6Choices          int
	ResidualNullityBefore      int
	ResidualNullityAfter       int
}

type Analysis struct {
	Previous contactalgebraic.Analysis

	Rows         []RowPolynomialAudit
	Candidate    CharacteristicPolynomialCandidate
	Requirements SymbolicConstructionRequirements
	Summary      Summary

	ContactRows                   int
	CandidateCoveredRows          int
	RationalFactorRows            int
	QuarticCandidateRows          int
	CandidateNumberFieldDegree    int
	ExactCertifiedRows            int
	ExactCharacteristicProofs     int
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
	ExactCharacteristicCertified  bool
	ExactNumberFieldLiftCertified bool

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
		prev, err := contactalgebraic.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactalgebraic.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.ContactRows != 7 || len(prev.PartialValues) != 7 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 148 requires Gate 147 closed-firewall seven-row contact algebraic audit")
	}
	if prev.ExactNumberFieldLifts != 0 || prev.ChargeSemanticRows != 0 || prev.RepresentationCompleteRows != 0 || prev.HyperchargeRowsDerived != 0 {
		return Analysis{}, fmt.Errorf("Gate 148 requires no prior exact number-field lift or contact charge semantics")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 148 refuses hidden observed physical input")
	}

	rows := make([]RowPolynomialAudit, 0, len(prev.PartialValues))
	maxPartialResidual := 0.0
	maxQuarticResidual := 0.0
	for i, value := range prev.PartialValues {
		row := auditRow(fmt.Sprintf("contact partial-overlap row %d", i+1), value)
		if row.PartialPolynomialResidual > maxPartialResidual {
			maxPartialResidual = row.PartialPolynomialResidual
		}
		if row.QuarticCandidateRoot && row.QuarticResidual > maxQuarticResidual {
			maxQuarticResidual = row.QuarticResidual
		}
		rows = append(rows, row)
	}

	rationalRows := count(rows, func(r RowPolynomialAudit) bool { return r.RationalDiagnostic })
	quarticRows := count(rows, func(r RowPolynomialAudit) bool { return r.QuarticCandidateRoot })
	coveredRows := count(rows, func(r RowPolynomialAudit) bool { return r.CandidatePolynomialCoversRow })
	exactRows := count(rows, func(r RowPolynomialAudit) bool { return r.CertifiedExactRow })
	semanticRows := count(rows, func(r RowPolynomialAudit) bool { return r.ChargeSemantic || r.RepresentationSemantic })

	candidate := CharacteristicPolynomialCandidate{
		PartialDegree:               7,
		FullDegree:                  14,
		RationalDiagnosticFactors:   []string{"2*x - 1", "3*x - 2", "3*x - 1"},
		QuarticFactor:               quarticPolynomialString(),
		PartialPolynomial:           partialPolynomialString(),
		FullPolynomialFactorization: "(x - 1)^7 * (2*x - 1)*(3*x - 2)*(3*x - 1)*(3240*x^4 - 7668*x^3 + 6426*x^2 - 2235*x + 271) / 58320",
		CandidateNumberFieldDegree:  4,
		RowsCovered:                 coveredRows,
		MaxPartialResidual:          maxPartialResidual,
		MaxQuarticResidual:          maxQuarticResidual,
		CandidateFactorizationRecognized: coveredRows == 7 && rationalRows == 3 && quarticRows == 4 &&
			maxPartialResidual < 1e-8 && maxQuarticResidual < 1e-8,
		ExactMatrixOverNumberField:     false,
		ExactDeterminantComputed:       false,
		ExactCharacteristicCertified:   false,
		RowMinimalPolynomialsCertified: false,
		Verdict:                        "rational reconstruction gives a strong symbolic characteristic-polynomial candidate, but the project has not computed an exact determinant over an exact contact overlap matrix",
	}

	reqs := SymbolicConstructionRequirements{
		RationalReconstructionCandidate: candidate.CandidateFactorizationRecognized,
		ResidualCheckPassed:             candidate.MaxPartialResidual < 1e-8 && candidate.MaxQuarticResidual < 1e-8,
		ExactOverlapMatrix:              false,
		ExactDeterminant:                false,
		IndependentCertificate:          false,
		RootIsolationCertificate:        false,
		RowwiseRootAssignmentProof:      false,
		AlgebraicRowSemantics:           false,
		ChargeOperatorSelected:          false,
		RepresentationRowsSelected:      false,
		ObservedInputFree:               true,
		AllSatisfied:                    false,
		Verdict:                         "symbolic candidate construction succeeds as a diagnostic, but exact matrix/determinant/certificate and row semantics are still missing",
	}

	summary := Summary{
		ContactRows:                len(prev.PartialValues),
		CandidateCoveredRows:       coveredRows,
		RationalFactorRows:         rationalRows,
		QuarticCandidateRows:       quarticRows,
		CandidateNumberFieldDegree: candidate.CandidateNumberFieldDegree,
		ExactCertifiedRows:         exactRows,
		ExactCharacteristicProofs:  0,
		ChargeSemanticRows:         semanticRows,
		RepresentationCompleteRows: 0,
		RepresentationOpenRows:     len(prev.PartialValues),
		ContactBetaRowsAllowed:     0,
		ContactZeroRowsProved:      0,
		ResidualS6Choices:          prev.ResidualS6Choices,
		ResidualNullityBefore:      prev.ResidualNullityAfter,
		ResidualNullityAfter:       prev.ResidualNullityAfter,
	}

	truth := "Gate 148 constructs a strong rational characteristic-polynomial candidate for the seven partial contact-overlap rows. The candidate factors into three rational rows and one quartic number-field candidate covering the four non-rational-looking rows. This is a real symbolic advance, but it is not yet an exact finite proof: the overlap matrix has not been lifted to an exact number field, no exact determinant/characteristic-polynomial certificate has been computed, no root-isolation proof assigns rows symbolically, and no charge or representation semantics follow. Contact T3R, B-L, hypercharge, mass activation, decoupling, and threshold beta permission remain sealed."

	return Analysis{
		Previous:                      prev,
		Rows:                          rows,
		Candidate:                     candidate,
		Requirements:                  reqs,
		Summary:                       summary,
		ContactRows:                   summary.ContactRows,
		CandidateCoveredRows:          coveredRows,
		RationalFactorRows:            rationalRows,
		QuarticCandidateRows:          quarticRows,
		CandidateNumberFieldDegree:    candidate.CandidateNumberFieldDegree,
		ExactCertifiedRows:            exactRows,
		ExactCharacteristicProofs:     0,
		ChargeSemanticRows:            0,
		T3RRowsDerived:                0,
		ChiralityRowsDerived:          0,
		BMinusLRowsDerived:            0,
		SU2LRowsDerived:               0,
		HyperchargeRowsDerived:        0,
		RepresentationCompleteRows:    0,
		RepresentationOpenRows:        summary.RepresentationOpenRows,
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
		ExactCharacteristicCertified:  false,
		ExactNumberFieldLiftCertified: false,
		TruthStatement:                truth,
		RejectedClaims: []string{
			"rational reconstruction is the same as an exact finite determinant proof",
			"the quartic candidate already gives physical contact charges",
			"the quartic roots may be mapped to hypercharge rows without local-field and representation semantics",
			"threshold beta rows are allowed once a polynomial candidate is found",
			"observed constants may be used to select root labels or charge normalization",
		},
		RemainingUnknowns: []string{
			"exact rational/number-field representation of the contact overlap matrix",
			"exact characteristic polynomial computed from that matrix",
			"root-isolation and row-wise minimal-polynomial certificates",
			"semantic map from polynomial roots to contact charge/representation rows",
			"local field variables, mass activation, and decoupling",
		},
		RecommendedNextGate: "Gate 149 — exact rational contact-overlap matrix lift / determinant certificate search",
	}, nil
}

func auditRow(name string, value float64) RowPolynomialAudit {
	partialResidual := math.Abs(evalPartialPolynomial(value))
	quarticResidual := math.Abs(evalQuarticPolynomial(value))
	row := RowPolynomialAudit{
		Name:                         name,
		Value:                        value,
		PartialPolynomialResidual:    partialResidual,
		QuarticResidual:              quarticResidual,
		CandidatePolynomialCoversRow: partialResidual < 1e-8,
		CertifiedExactRow:            false,
		ChargeSemantic:               false,
		RepresentationSemantic:       false,
		OpensBetaPermission:          false,
	}
	if near(value, 0.5, 1e-8) {
		row.RationalDiagnostic = true
		row.RationalRootFactor = "2*x - 1"
		row.CandidateMinimalPolynomial = "2*x - 1"
		row.Verdict = "rational factor covers this diagnostic row, but exact row semantics are still not certified"
		return row
	}
	if near(value, 2.0/3.0, 1e-8) {
		row.RationalDiagnostic = true
		row.RationalRootFactor = "3*x - 2"
		row.CandidateMinimalPolynomial = "3*x - 2"
		row.Verdict = "rational factor covers this diagnostic row, but exact row semantics are still not certified"
		return row
	}
	if near(value, 1.0/3.0, 1e-8) {
		row.RationalDiagnostic = true
		row.RationalRootFactor = "3*x - 1"
		row.CandidateMinimalPolynomial = "3*x - 1"
		row.Verdict = "rational factor covers this diagnostic row, but exact row semantics are still not certified"
		return row
	}
	row.QuarticCandidateRoot = quarticResidual < 1e-8
	row.CandidateMinimalPolynomial = quarticPolynomialString()
	row.Verdict = "quartic factor covers this numerical row as a candidate algebraic conjugate; exact determinant/root certificate and semantics are not derived"
	return row
}

func evalPartialPolynomial(x float64) float64 {
	// Monic expanded candidate:
	// x^7 - 58/15 x^6 + 563/90 x^5 - 1481/270 x^4
	// + 18233/6480 x^3 - 2051/2430 x^2 + 7993/58320 x - 271/29160.
	coeffs := []float64{
		1,
		-58.0 / 15.0,
		563.0 / 90.0,
		-1481.0 / 270.0,
		18233.0 / 6480.0,
		-2051.0 / 2430.0,
		7993.0 / 58320.0,
		-271.0 / 29160.0,
	}
	return horner(coeffs, x)
}

func evalQuarticPolynomial(x float64) float64 {
	coeffs := []float64{3240, -7668, 6426, -2235, 271}
	return horner(coeffs, x)
}

func horner(coeffs []float64, x float64) float64 {
	acc := 0.0
	for _, c := range coeffs {
		acc = acc*x + c
	}
	return acc
}

func near(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func quarticPolynomialString() string {
	return "3240*x^4 - 7668*x^3 + 6426*x^2 - 2235*x + 271"
}

func partialPolynomialString() string {
	return "x^7 - 58/15*x^6 + 563/90*x^5 - 1481/270*x^4 + 18233/6480*x^3 - 2051/2430*x^2 + 7993/58320*x - 271/29160"
}

func FormatRow(r RowPolynomialAudit) string {
	factor := r.RationalRootFactor
	if factor == "" {
		factor = "quartic"
	}
	return fmt.Sprintf("%s value=%.10f factor=%s partialResidual=%.3e quarticResidual=%.3e covered=%t certified=%t semantic=%t beta=%t (%s)", r.Name, r.Value, factor, r.PartialPolynomialResidual, r.QuarticResidual, r.CandidatePolynomialCoversRow, r.CertifiedExactRow, r.ChargeSemantic || r.RepresentationSemantic, r.OpensBetaPermission, r.Verdict)
}

func FormatRows(rows []RowPolynomialAudit) string {
	parts := make([]string, 0, len(rows))
	for _, row := range rows {
		parts = append(parts, FormatRow(row))
	}
	return strings.Join(parts, "; ")
}

func FormatCandidate(c CharacteristicPolynomialCandidate) string {
	return fmt.Sprintf("partialDegree=%d fullDegree=%d factors=[%s;%s] fieldDegreeCandidate=%d rows=%d maxPartialResidual=%.3e maxQuarticResidual=%.3e exactMatrix=%t exactDet=%t exactChar=%t rowMinpoly=%t (%s)", c.PartialDegree, c.FullDegree, strings.Join(c.RationalDiagnosticFactors, ", "), c.QuarticFactor, c.CandidateNumberFieldDegree, c.RowsCovered, c.MaxPartialResidual, c.MaxQuarticResidual, c.ExactMatrixOverNumberField, c.ExactDeterminantComputed, c.ExactCharacteristicCertified, c.RowMinimalPolynomialsCertified, c.Verdict)
}

func FormatRequirements(r SymbolicConstructionRequirements) string {
	return fmt.Sprintf("candidate=%t residual=%t exactMatrix=%t exactDet=%t independentCert=%t rootIsolation=%t rowProof=%t semantics=%t charge=%t reps=%t observedFree=%t all=%t (%s)", r.RationalReconstructionCandidate, r.ResidualCheckPassed, r.ExactOverlapMatrix, r.ExactDeterminant, r.IndependentCertificate, r.RootIsolationCertificate, r.RowwiseRootAssignmentProof, r.AlgebraicRowSemantics, r.ChargeOperatorSelected, r.RepresentationRowsSelected, r.ObservedInputFree, r.AllSatisfied, r.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contact=%d covered=%d rationalFactors=%d quarticRows=%d fieldDegreeCandidate=%d exactRows=%d exactCharProofs=%d semantic=%d reps=%d/%d beta=%d nullity=%d→%d", s.ContactRows, s.CandidateCoveredRows, s.RationalFactorRows, s.QuarticCandidateRows, s.CandidateNumberFieldDegree, s.ExactCertifiedRows, s.ExactCharacteristicProofs, s.ChargeSemanticRows, s.RepresentationCompleteRows, s.RepresentationOpenRows, s.ContactBetaRowsAllowed, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func count[T any](items []T, pred func(T) bool) int {
	n := 0
	for _, item := range items {
		if pred(item) {
			n++
		}
	}
	return n
}
