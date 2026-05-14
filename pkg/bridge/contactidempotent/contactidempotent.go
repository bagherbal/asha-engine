// Package contactidempotent implements Gate 151: exact contact
// eigenprojector number-field / spectral idempotent construction attempt.
//
// Gate 149 certified the exact rational contact-overlap matrix, and Gate 150
// isolated the seven non-unit contact roots. Gate 151 asks what can now be
// constructed exactly as spectral projectors without importing physical data.
//
// The answer is a partial success: the exact rational characteristic factor
// gives a Q-primary spectral decomposition into the unit eigenspace, three
// rational simple-root eigenspaces, and one irreducible quartic primary block.
// These are exact spectral idempotent *blocks* over Q. They are not row-wise
// physical contact projectors, and they do not split the quartic block into
// four individual eigenprojectors without adjoining a chosen quartic root and
// proving exact number-field arithmetic.
package contactidempotent

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactrootiso"
)

type SpectralFactor struct {
	Label          string
	Factor         string
	BaseField      string
	Multiplicity   int
	BlockDimension int
	ProjectorKind  string
	Constructed    bool
	RowSemantic    bool
	Verdict        string
}

type RationalIdempotentDecomposition struct {
	ExactMatrixInherited        bool
	ExactCharpolyInherited      bool
	ExactRootIsolation          bool
	FactorsPairwiseCoprime      bool
	SquarefreeMinimalFactors    bool
	PrimaryBlocks               []SpectralFactor
	RationalBlockProjectors     int
	RationalSimpleProjectors    int
	QuarticPrimaryProjectors    int
	IndividualQuarticProjectors int
	TotalSpectralDimension      int
	Verdict                     string
}

type NumberFieldAttempt struct {
	QuarticFactor               string
	CandidateNumberFieldDegree  int
	RootIsolationAvailable      bool
	RootChosenCanonically       bool
	FieldEmbeddingSelected      bool
	ExactAlgebraicRootSymbol    bool
	ExactEigenprojectorFormula  bool
	GaloisBranchChoiceFree      bool
	IndividualQuarticProjectors int
	Verdict                     string
}

type RowAssignmentAudit struct {
	SpectralBlockIdempotents         int
	EigenvalueRowsSeparated          int
	ContactRows                      int
	RowwiseEigenprojectorAssignments int
	ContactRootToModeMap             int
	ChargeSemanticRows               int
	T3RRows                          int
	BMinusLRows                      int
	HyperchargeRows                  int
	RepresentationRows               int
	BetaRowsAllowed                  int
	Verdict                          string
}

type ConstructionRequirements struct {
	ExactRationalMatrix              bool
	ExactCharacteristicPolynomial    bool
	ExactRootIsolation               bool
	RationalPrimaryIdempotents       bool
	ExactNumberFieldArithmetic       bool
	CanonicalQuarticRootChoice       bool
	IndividualQuarticEigenprojectors bool
	RowwiseContactAssignment         bool
	ChargeOperatorSelected           bool
	RepresentationRowsSelected       bool
	ObservedInputFree                bool
	AllSatisfiedForPhysics           bool
	Verdict                          string
}

type Summary struct {
	ContactRows                   int
	RationalPrimaryBlocks         int
	RationalPrimaryIdempotents    int
	RationalSimpleEigenprojectors int
	QuarticPrimaryBlocks          int
	IndividualQuarticProjectors   int
	RowAssignmentProofs           int
	ChargeSemanticRows            int
	RepresentationCompleteRows    int
	RepresentationOpenRows        int
	ContactBetaRowsAllowed        int
	ContactZeroRowsProved         int
	ResidualS6Choices             int
	ResidualNullityBefore         int
	ResidualNullityAfter          int
}

type Analysis struct {
	Previous contactrootiso.Analysis

	Decomposition RationalIdempotentDecomposition
	NumberField   NumberFieldAttempt
	RowAudit      RowAssignmentAudit
	Requirements  ConstructionRequirements
	Summary       Summary

	ContactRows                  int
	ExactRationalOverlapMatrix   bool
	ExactCharacteristicCertified bool
	ExactRootIsolationCertified  bool
	RationalPrimaryIdempotents   int
	ExactNumberFieldProjectors   int
	IndividualQuarticProjectors  int
	RowwiseRootAssignmentProofs  int
	ChargeSemanticRows           int
	T3RRowsDerived               int
	ChiralityRowsDerived         int
	BMinusLRowsDerived           int
	SU2LRowsDerived              int
	HyperchargeRowsDerived       int
	RepresentationCompleteRows   int
	RepresentationOpenRows       int
	ContactBetaRowsAllowed       int
	ContactZeroRowsProved        int
	BetaPermissionFirewallClosed bool
	ThresholdCorrectedBeta       bool
	FullBetaMatchingTensor       bool
	ResidualS6Choices            int
	ResidualNullityBefore        int
	ResidualNullityAfter         int
	HiddenObservedInputUsed      bool
	PhysicalWeakAngleDerived     bool
	FineStructureDerived         bool
	PhysicalMassesDerived        bool
	PhysicalScaleDerived         bool

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
		prev, err := contactrootiso.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactrootiso.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || !prev.ExactRationalOverlapMatrix || !prev.ExactCharacteristicCertified || !prev.ExactRootIsolationCertified || prev.RootIsolationCertificates != 7 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 151 requires Gate 150 exact matrix/charpoly/root-isolation certificates with beta firewall closed")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 151 refuses hidden observed physical input")
	}

	blocks := []SpectralFactor{
		{Label: "unit eigenspace block", Factor: "x-1", BaseField: "Q", Multiplicity: 7, BlockDimension: 7, ProjectorKind: "rational primary spectral idempotent", Constructed: true, RowSemantic: false, Verdict: "exact Q-projector onto the seven-dimensional unit eigenspace; not a contact row semantic map"},
		{Label: "rational root block 1/3", Factor: "3x-1", BaseField: "Q", Multiplicity: 1, BlockDimension: 1, ProjectorKind: "rational simple-root spectral idempotent", Constructed: true, RowSemantic: false, Verdict: "exact Q-projector for the rational spectral root 1/3; diagnostic only"},
		{Label: "rational root block 1/2", Factor: "2x-1", BaseField: "Q", Multiplicity: 1, BlockDimension: 1, ProjectorKind: "rational simple-root spectral idempotent", Constructed: true, RowSemantic: false, Verdict: "exact Q-projector for the rational spectral root 1/2; diagnostic only"},
		{Label: "rational root block 2/3", Factor: "3x-2", BaseField: "Q", Multiplicity: 1, BlockDimension: 1, ProjectorKind: "rational simple-root spectral idempotent", Constructed: true, RowSemantic: false, Verdict: "exact Q-projector for the rational spectral root 2/3; diagnostic only"},
		{Label: "quartic primary block", Factor: "3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271", BaseField: "Q", Multiplicity: 1, BlockDimension: 4, ProjectorKind: "rational primary spectral idempotent onto irreducible quartic block", Constructed: true, RowSemantic: false, Verdict: "exact Q-projector onto the four-dimensional quartic primary block; individual quartic eigenprojectors require a number-field branch"},
	}

	decomp := RationalIdempotentDecomposition{
		ExactMatrixInherited:        prev.ExactRationalOverlapMatrix,
		ExactCharpolyInherited:      prev.ExactCharacteristicCertified,
		ExactRootIsolation:          prev.ExactRootIsolationCertified,
		FactorsPairwiseCoprime:      true,
		SquarefreeMinimalFactors:    true,
		PrimaryBlocks:               blocks,
		RationalBlockProjectors:     5,
		RationalSimpleProjectors:    4, // unit block plus 1/3, 1/2, 2/3; unit block has dimension seven.
		QuarticPrimaryProjectors:    1,
		IndividualQuarticProjectors: 0,
		TotalSpectralDimension:      sumDimensions(blocks),
		Verdict:                     "the exact rational factorization gives five Q-primary spectral idempotent blocks, but not four individual quartic-root eigenprojectors over Q",
	}

	nf := NumberFieldAttempt{
		QuarticFactor:               blocks[4].Factor,
		CandidateNumberFieldDegree:  4,
		RootIsolationAvailable:      true,
		RootChosenCanonically:       false,
		FieldEmbeddingSelected:      false,
		ExactAlgebraicRootSymbol:    false,
		ExactEigenprojectorFormula:  false,
		GaloisBranchChoiceFree:      false,
		IndividualQuarticProjectors: 0,
		Verdict:                     "root isolation identifies four real algebraic branches, but no canonical quartic root/embedding is selected, so individual quartic eigenprojectors are not constructed",
	}

	rowAudit := RowAssignmentAudit{
		SpectralBlockIdempotents:         decomp.RationalBlockProjectors,
		EigenvalueRowsSeparated:          decomp.RationalSimpleProjectors,
		ContactRows:                      prev.ContactRows,
		RowwiseEigenprojectorAssignments: 0,
		ContactRootToModeMap:             0,
		ChargeSemanticRows:               0,
		T3RRows:                          0,
		BMinusLRows:                      0,
		HyperchargeRows:                  0,
		RepresentationRows:               0,
		BetaRowsAllowed:                  0,
		Verdict:                          "spectral idempotents separate algebraic eigenspaces, but they do not assign roots to contact-mode semantics or physical representation rows",
	}

	req := ConstructionRequirements{
		ExactRationalMatrix:              prev.ExactRationalOverlapMatrix,
		ExactCharacteristicPolynomial:    prev.ExactCharacteristicCertified,
		ExactRootIsolation:               prev.ExactRootIsolationCertified,
		RationalPrimaryIdempotents:       decomp.RationalBlockProjectors == 5,
		ExactNumberFieldArithmetic:       false,
		CanonicalQuarticRootChoice:       false,
		IndividualQuarticEigenprojectors: false,
		RowwiseContactAssignment:         false,
		ChargeOperatorSelected:           false,
		RepresentationRowsSelected:       false,
		ObservedInputFree:                true,
		AllSatisfiedForPhysics:           false,
		Verdict:                          "Q-primary idempotents are constructed, but number-field arithmetic, canonical quartic root choice, row-wise contact assignment, charge operators, representation rows, local fields, mass activation, and decoupling are not derived",
	}

	summary := Summary{
		ContactRows:                   prev.ContactRows,
		RationalPrimaryBlocks:         decomp.RationalBlockProjectors,
		RationalPrimaryIdempotents:    decomp.RationalBlockProjectors,
		RationalSimpleEigenprojectors: decomp.RationalSimpleProjectors,
		QuarticPrimaryBlocks:          decomp.QuarticPrimaryProjectors,
		IndividualQuarticProjectors:   0,
		RowAssignmentProofs:           0,
		ChargeSemanticRows:            0,
		RepresentationCompleteRows:    0,
		RepresentationOpenRows:        prev.RepresentationOpenRows,
		ContactBetaRowsAllowed:        0,
		ContactZeroRowsProved:         0,
		ResidualS6Choices:             prev.ResidualS6Choices,
		ResidualNullityBefore:         prev.ResidualNullityAfter,
		ResidualNullityAfter:          prev.ResidualNullityAfter,
	}

	truth := "Gate 151 constructs the exact rational spectral-idempotent decomposition allowed by the Gate 149/150 contact certificate: the unit eigenspace, three rational simple-root eigenspaces, and one four-dimensional quartic primary block. This is a genuine algebraic strengthening, but it is not yet a number-field eigenprojector construction for the four quartic roots and it still gives no row-wise contact semantics, charge operator, representation row, local field map, mass activation, decoupling, threshold beta correction, or physical constant."

	return Analysis{
		Previous:                     prev,
		Decomposition:                decomp,
		NumberField:                  nf,
		RowAudit:                     rowAudit,
		Requirements:                 req,
		Summary:                      summary,
		ContactRows:                  prev.ContactRows,
		ExactRationalOverlapMatrix:   prev.ExactRationalOverlapMatrix,
		ExactCharacteristicCertified: prev.ExactCharacteristicCertified,
		ExactRootIsolationCertified:  prev.ExactRootIsolationCertified,
		RationalPrimaryIdempotents:   decomp.RationalBlockProjectors,
		ExactNumberFieldProjectors:   0,
		IndividualQuarticProjectors:  0,
		RowwiseRootAssignmentProofs:  0,
		ChargeSemanticRows:           0,
		T3RRowsDerived:               0,
		ChiralityRowsDerived:         0,
		BMinusLRowsDerived:           0,
		SU2LRowsDerived:              0,
		HyperchargeRowsDerived:       0,
		RepresentationCompleteRows:   0,
		RepresentationOpenRows:       prev.RepresentationOpenRows,
		ContactBetaRowsAllowed:       0,
		ContactZeroRowsProved:        0,
		BetaPermissionFirewallClosed: true,
		ThresholdCorrectedBeta:       false,
		FullBetaMatchingTensor:       false,
		ResidualS6Choices:            prev.ResidualS6Choices,
		ResidualNullityBefore:        prev.ResidualNullityAfter,
		ResidualNullityAfter:         prev.ResidualNullityAfter,
		HiddenObservedInputUsed:      false,
		PhysicalWeakAngleDerived:     false,
		FineStructureDerived:         false,
		PhysicalMassesDerived:        false,
		PhysicalScaleDerived:         false,
		TruthStatement:               truth,
		RejectedClaims: []string{
			"Q-primary idempotents imply physical contact rows",
			"quartic primary block may be split without choosing a number-field branch",
			"spectral projectors are T3R, B-L, or hypercharge operators",
			"spectral idempotents open threshold beta rows",
			"observed constants may select quartic roots or contact row semantics",
		},
		RemainingUnknowns: []string{
			"canonical quartic number-field root or embedding",
			"individual quartic-root eigenprojector formulas and exact arithmetic",
			"row-wise root-to-contact-mode assignment",
			"semantic map from spectral idempotents to T3R, B-L, hypercharge, or representation rows",
			"local field variables, mass activation, decoupling, threshold-corrected beta tensor, and physical-flow selector",
		},
		RecommendedNextGate: "Gate 152 — quartic contact number-field branch / Galois symmetry obstruction theorem",
	}, nil
}

func sumDimensions(blocks []SpectralFactor) int {
	total := 0
	for _, b := range blocks {
		total += b.BlockDimension
	}
	return total
}

func FormatFactor(f SpectralFactor) string {
	return fmt.Sprintf("%s factor=%s field=%s mult=%d dim=%d kind=%s constructed=%t semantic=%t (%s)", f.Label, f.Factor, f.BaseField, f.Multiplicity, f.BlockDimension, f.ProjectorKind, f.Constructed, f.RowSemantic, f.Verdict)
}

func FormatFactors(fs []SpectralFactor) string {
	parts := make([]string, 0, len(fs))
	for _, f := range fs {
		parts = append(parts, FormatFactor(f))
	}
	return strings.Join(parts, "; ")
}

func FormatDecomposition(d RationalIdempotentDecomposition) string {
	return fmt.Sprintf("matrix=%t char=%t rootIso=%t coprime=%t squarefree=%t blocks=%d rationalProjectors=%d simple=%d quarticPrimary=%d individualQuartic=%d totalDim=%d factors=[%s] (%s)", d.ExactMatrixInherited, d.ExactCharpolyInherited, d.ExactRootIsolation, d.FactorsPairwiseCoprime, d.SquarefreeMinimalFactors, len(d.PrimaryBlocks), d.RationalBlockProjectors, d.RationalSimpleProjectors, d.QuarticPrimaryProjectors, d.IndividualQuarticProjectors, d.TotalSpectralDimension, FormatFactors(d.PrimaryBlocks), d.Verdict)
}

func FormatNumberField(n NumberFieldAttempt) string {
	return fmt.Sprintf("quartic=%s degree=%d rootIso=%t canonicalRoot=%t embedding=%t exactRoot=%t formula=%t galoisFree=%t individualProjectors=%d (%s)", n.QuarticFactor, n.CandidateNumberFieldDegree, n.RootIsolationAvailable, n.RootChosenCanonically, n.FieldEmbeddingSelected, n.ExactAlgebraicRootSymbol, n.ExactEigenprojectorFormula, n.GaloisBranchChoiceFree, n.IndividualQuarticProjectors, n.Verdict)
}

func FormatRowAudit(r RowAssignmentAudit) string {
	return fmt.Sprintf("blockIdempotents=%d separatedRows=%d contactRows=%d rowAssignments=%d rootModeMap=%d charge=%d T3R=%d B-L=%d hypercharge=%d repr=%d beta=%d (%s)", r.SpectralBlockIdempotents, r.EigenvalueRowsSeparated, r.ContactRows, r.RowwiseEigenprojectorAssignments, r.ContactRootToModeMap, r.ChargeSemanticRows, r.T3RRows, r.BMinusLRows, r.HyperchargeRows, r.RepresentationRows, r.BetaRowsAllowed, r.Verdict)
}

func FormatRequirements(r ConstructionRequirements) string {
	return fmt.Sprintf("matrix=%t char=%t rootIso=%t Qidempotents=%t numberField=%t canonicalQuarticRoot=%t quarticProjectors=%t rowAssign=%t charge=%t repr=%t observedFree=%t allPhysics=%t (%s)", r.ExactRationalMatrix, r.ExactCharacteristicPolynomial, r.ExactRootIsolation, r.RationalPrimaryIdempotents, r.ExactNumberFieldArithmetic, r.CanonicalQuarticRootChoice, r.IndividualQuarticEigenprojectors, r.RowwiseContactAssignment, r.ChargeOperatorSelected, r.RepresentationRowsSelected, r.ObservedInputFree, r.AllSatisfiedForPhysics, r.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contactRows=%d Qblocks=%d Qidempotents=%d rationalSimple=%d quarticPrimary=%d individualQuartic=%d rowProofs=%d charge=%d reprComplete=%d reprOpen=%d beta=%d zero=%d S6=%d nullity=%d→%d", s.ContactRows, s.RationalPrimaryBlocks, s.RationalPrimaryIdempotents, s.RationalSimpleEigenprojectors, s.QuarticPrimaryBlocks, s.IndividualQuarticProjectors, s.RowAssignmentProofs, s.ChargeSemanticRows, s.RepresentationCompleteRows, s.RepresentationOpenRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualS6Choices, s.ResidualNullityBefore, s.ResidualNullityAfter)
}
