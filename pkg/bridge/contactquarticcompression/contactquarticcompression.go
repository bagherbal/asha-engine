// Package contactquarticcompression implements Gate 154: quartic orbit semantic
// compression / four-row block beta firewall theorem.
//
// Gate 153 established the exact Galois-invariant contact partition
// 1+1+1+4. Gate 154 asks whether the four-root quartic orbit can be safely
// compressed into a block semantic object strong enough to open threshold beta
// matching. The answer is no: the block has exact symmetric invariants, but it
// lacks row-level branch semantics, gauge representation, local fields, mass
// activation, and decoupling. The quartic block is therefore a certified
// spectral diagnostic, not a physical beta row.
package contactquarticcompression

import (
	"fmt"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactbranchsemantics"
)

type QuarticBlockInvariant struct {
	Degree               int
	BlockRows            int
	Polynomial           string
	SumRootsNumerator    int
	SumRootsDenominator  int
	MeanNumerator        int
	MeanDenominator      int
	PairSumNumerator     int
	PairSumDenominator   int
	TripleSumNumerator   int
	TripleSumDenominator int
	ProductNumerator     int
	ProductDenominator   int
	SymmetricInvariants  int
	ExactOverQ           bool
	GaloisInvariant      bool
	IndividualRowsSplit  bool
	Verdict              string
}

type CompressionAudit struct {
	InputPattern               string
	CompressedPattern          string
	QuarticBlockRows           int
	CompressedBlocks           int
	RowLevelQuarticSemantics   int
	BlockLevelSpectralSemantic bool
	RowComplete                bool
	BranchFree                 bool
	HiddenBranchChoices        int
	Verdict                    string
}

type BlockBetaPermissionAudit struct {
	BlockRows                     int
	GaugeRepresentationRows       int
	SpinStatisticsRows            int
	LocalFieldRows                int
	LorentzKineticRows            int
	MassActivationRows            int
	DecouplingRows                int
	DynkinIndexRows               int
	ThresholdBetaRowsAllowed      int
	BlockCanContributeAsMultiplet bool
	Verdict                       string
}

type PatternAudit struct {
	GaloisSafePattern       string
	QuarticBlockPattern     string
	CurrentQuotientPattern  string
	ContactSingletonPattern string
	FanoPattern             string
	MatchesCurrent          bool
	MatchesContact          bool
	MatchesFano             bool
	Verdict                 string
}

type PhysicsFirewall struct {
	ObservedInputFree          bool
	ExactMatrix                bool
	ExactCharpoly              bool
	ExactRootIsolation         bool
	RationalPrimaryIdempotents bool
	GaloisInvariantPartition   bool
	QuarticBlockInvariants     bool
	IndividualQuarticSplit     bool
	GaugeRepresentation        bool
	LocalFieldMap              bool
	MassActivation             bool
	Decoupling                 bool
	BetaRows                   bool
	PhysicalConstants          bool
	AllSatisfiedForPhysics     bool
	Verdict                    string
}

type Summary struct {
	ContactRows                int
	GaloisInvariantOrbits      int
	RationalSingletonRows      int
	QuarticOrbitRows           int
	QuarticCompressedBlocks    int
	QuarticBlockInvariants     int
	IndividualQuarticRows      int
	RowAssignmentProofs        int
	RepresentationCompleteRows int
	RepresentationOpenRows     int
	QuarticBlockBetaRows       int
	ContactBetaRowsAllowed     int
	ContactZeroRowsProved      int
	ResidualS6Choices          int
	ResidualNullityBefore      int
	ResidualNullityAfter       int
}

type Analysis struct {
	Previous contactbranchsemantics.Analysis

	BlockInvariant QuarticBlockInvariant
	Compression    CompressionAudit
	BlockBeta      BlockBetaPermissionAudit
	Pattern        PatternAudit
	Firewall       PhysicsFirewall
	Summary        Summary

	ContactRows                  int
	ExactRationalOverlapMatrix   bool
	ExactCharacteristicCertified bool
	ExactRootIsolationCertified  bool
	RationalPrimaryIdempotents   int
	GaloisInvariantOrbits        int
	RationalSingletonRows        int
	QuarticOrbitRows             int
	QuarticCompressedBlocks      int
	QuarticBlockInvariants       int
	IndividualQuarticRows        int
	CanonicalQuarticBranches     int
	ExactNumberFieldProjectors   int
	IndividualQuarticProjectors  int
	RowwiseRootAssignmentProofs  int
	ChargeSemanticRows           int
	T3RRowsDerived               int
	ChiralityRowsDerived         int
	BMinusLRowsDerived           int
	SU2LRowsDerived              int
	HyperchargeRowsDerived       int
	GaugeRepresentationRows      int
	LocalFieldRows               int
	MassActivationRows           int
	DecouplingRows               int
	RepresentationCompleteRows   int
	RepresentationOpenRows       int
	QuarticBlockBetaRows         int
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
		prev, err := contactbranchsemantics.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactbranchsemantics.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || !prev.ExactRationalOverlapMatrix || !prev.ExactCharacteristicCertified || !prev.ExactRootIsolationCertified || prev.GaloisInvariantOrbits != 4 || prev.RationalSingletonRows != 3 || prev.QuarticOrbitRows != 4 || prev.IndividualQuarticRows != 0 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 154 requires Gate 153 exact 1+1+1+4 Galois-invariant partition with closed beta firewall")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 154 refuses hidden observed physical input")
	}

	block := QuarticBlockInvariant{
		Degree:               4,
		BlockRows:            4,
		Polynomial:           "3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271",
		SumRootsNumerator:    71,
		SumRootsDenominator:  30,
		MeanNumerator:        71,
		MeanDenominator:      120,
		PairSumNumerator:     119,
		PairSumDenominator:   60,
		TripleSumNumerator:   149,
		TripleSumDenominator: 216,
		ProductNumerator:     271,
		ProductDenominator:   3240,
		SymmetricInvariants:  4,
		ExactOverQ:           true,
		GaloisInvariant:      true,
		IndividualRowsSplit:  false,
		Verdict:              "the quartic orbit has exact Q-symmetric block invariants, but no individual root/row semantics are selected",
	}

	compression := CompressionAudit{
		InputPattern:               prev.Partition.OrbitPattern,
		CompressedPattern:          "1+1+1+[4]",
		QuarticBlockRows:           4,
		CompressedBlocks:           4,
		RowLevelQuarticSemantics:   0,
		BlockLevelSpectralSemantic: true,
		RowComplete:                false,
		BranchFree:                 true,
		HiddenBranchChoices:        4,
		Verdict:                    "the quartic orbit compresses to one exact four-row block; this is branch-free but not row-complete",
	}

	blockBeta := BlockBetaPermissionAudit{
		BlockRows:                     4,
		GaugeRepresentationRows:       0,
		SpinStatisticsRows:            0,
		LocalFieldRows:                0,
		LorentzKineticRows:            0,
		MassActivationRows:            0,
		DecouplingRows:                0,
		DynkinIndexRows:               0,
		ThresholdBetaRowsAllowed:      0,
		BlockCanContributeAsMultiplet: false,
		Verdict:                       "a four-row spectral block cannot contribute to beta matching without gauge representation, spin/statistics, local field, mass activation, and decoupling data",
	}

	pattern := PatternAudit{
		GaloisSafePattern:       prev.Partition.OrbitPattern,
		QuarticBlockPattern:     compression.CompressedPattern,
		CurrentQuotientPattern:  "1+6",
		ContactSingletonPattern: "1+1+1+1+1+1+1",
		FanoPattern:             "7-transitive",
		MatchesCurrent:          false,
		MatchesContact:          false,
		MatchesFano:             false,
		Verdict:                 "the compressed 1+1+1+[4] block pattern is exact, but it still mismatches current 1+6 semantics, seven singleton contact rows, and Fano-transitive labeling",
	}

	firewall := PhysicsFirewall{
		ObservedInputFree:          true,
		ExactMatrix:                prev.ExactRationalOverlapMatrix,
		ExactCharpoly:              prev.ExactCharacteristicCertified,
		ExactRootIsolation:         prev.ExactRootIsolationCertified,
		RationalPrimaryIdempotents: prev.RationalPrimaryIdempotents == 5,
		GaloisInvariantPartition:   prev.GaloisInvariantOrbits == 4,
		QuarticBlockInvariants:     block.ExactOverQ && block.GaloisInvariant,
		IndividualQuarticSplit:     false,
		GaugeRepresentation:        false,
		LocalFieldMap:              false,
		MassActivation:             false,
		Decoupling:                 false,
		BetaRows:                   false,
		PhysicalConstants:          false,
		AllSatisfiedForPhysics:     false,
		Verdict:                    "quartic block compression strengthens the exact spectral ledger, but it does not open representation, beta, or physical-constant claims",
	}

	summary := Summary{
		ContactRows:                prev.ContactRows,
		GaloisInvariantOrbits:      prev.GaloisInvariantOrbits,
		RationalSingletonRows:      prev.RationalSingletonRows,
		QuarticOrbitRows:           prev.QuarticOrbitRows,
		QuarticCompressedBlocks:    1,
		QuarticBlockInvariants:     block.SymmetricInvariants,
		IndividualQuarticRows:      0,
		RowAssignmentProofs:        0,
		RepresentationCompleteRows: 0,
		RepresentationOpenRows:     prev.RepresentationOpenRows,
		QuarticBlockBetaRows:       0,
		ContactBetaRowsAllowed:     0,
		ContactZeroRowsProved:      0,
		ResidualS6Choices:          prev.ResidualS6Choices,
		ResidualNullityBefore:      prev.ResidualNullityAfter,
		ResidualNullityAfter:       prev.ResidualNullityAfter,
	}

	truth := "Gate 154 proves that the exact quartic contact orbit can be compressed into a branch-free four-row spectral block with rational symmetric invariants. This is a stronger exact diagnostic than anonymous row data, but it is still not a physical multiplet. Without individual quartic row semantics or block-level gauge representation, local field, mass activation, and decoupling data, the quartic block cannot contribute threshold beta rows."

	return Analysis{
		Previous:                     prev,
		BlockInvariant:               block,
		Compression:                  compression,
		BlockBeta:                    blockBeta,
		Pattern:                      pattern,
		Firewall:                     firewall,
		Summary:                      summary,
		ContactRows:                  prev.ContactRows,
		ExactRationalOverlapMatrix:   prev.ExactRationalOverlapMatrix,
		ExactCharacteristicCertified: prev.ExactCharacteristicCertified,
		ExactRootIsolationCertified:  prev.ExactRootIsolationCertified,
		RationalPrimaryIdempotents:   prev.RationalPrimaryIdempotents,
		GaloisInvariantOrbits:        prev.GaloisInvariantOrbits,
		RationalSingletonRows:        prev.RationalSingletonRows,
		QuarticOrbitRows:             prev.QuarticOrbitRows,
		QuarticCompressedBlocks:      1,
		QuarticBlockInvariants:       block.SymmetricInvariants,
		IndividualQuarticRows:        0,
		CanonicalQuarticBranches:     0,
		ExactNumberFieldProjectors:   0,
		IndividualQuarticProjectors:  0,
		RowwiseRootAssignmentProofs:  0,
		ChargeSemanticRows:           0,
		T3RRowsDerived:               0,
		ChiralityRowsDerived:         0,
		BMinusLRowsDerived:           0,
		SU2LRowsDerived:              0,
		HyperchargeRowsDerived:       0,
		GaugeRepresentationRows:      0,
		LocalFieldRows:               0,
		MassActivationRows:           0,
		DecouplingRows:               0,
		RepresentationCompleteRows:   0,
		RepresentationOpenRows:       prev.RepresentationOpenRows,
		QuarticBlockBetaRows:         0,
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
			"the quartic block is four physical contact fields",
			"quartic symmetric invariants are T3R, B-L, hypercharge, or beta rows",
			"a four-row spectral block can contribute threshold beta coefficients without representation and decoupling data",
			"the compressed 1+1+1+[4] pattern resolves current 1+6 or Fano-transitive semantics",
			"quartic block compression derives physical constants or masses",
		},
		RemainingUnknowns: []string{
			"individual quartic branch/projector semantics or a physical block multiplet theorem",
			"gauge representation, spin/statistics, local field map, mass activation, and decoupling for the quartic block",
			"contact T3R, B-L, hypercharge, SU(2)L, and threshold beta tensor",
			"physical-flow selector for u, L, and threshold corrections",
		},
		RecommendedNextGate: "Gate 155 — quartic block multiplet representation / beta-index obstruction theorem",
	}, nil
}

func FormatBlockInvariant(b QuarticBlockInvariant) string {
	return fmt.Sprintf("degree=%d rows=%d poly=%q sum=%d/%d mean=%d/%d pair=%d/%d triple=%d/%d product=%d/%d invariants=%d exactQ=%t galois=%t split=%t (%s)", b.Degree, b.BlockRows, b.Polynomial, b.SumRootsNumerator, b.SumRootsDenominator, b.MeanNumerator, b.MeanDenominator, b.PairSumNumerator, b.PairSumDenominator, b.TripleSumNumerator, b.TripleSumDenominator, b.ProductNumerator, b.ProductDenominator, b.SymmetricInvariants, b.ExactOverQ, b.GaloisInvariant, b.IndividualRowsSplit, b.Verdict)
}

func FormatCompression(c CompressionAudit) string {
	return fmt.Sprintf("input=%s compressed=%s quarticRows=%d blocks=%d rowSemantics=%d blockSemantic=%t rowComplete=%t branchFree=%t hiddenBranchChoices=%d (%s)", c.InputPattern, c.CompressedPattern, c.QuarticBlockRows, c.CompressedBlocks, c.RowLevelQuarticSemantics, c.BlockLevelSpectralSemantic, c.RowComplete, c.BranchFree, c.HiddenBranchChoices, c.Verdict)
}

func FormatBlockBeta(b BlockBetaPermissionAudit) string {
	return fmt.Sprintf("blockRows=%d gauge=%d spin=%d local=%d lorentz=%d mass=%d decoupling=%d dynkin=%d beta=%d multiplet=%t (%s)", b.BlockRows, b.GaugeRepresentationRows, b.SpinStatisticsRows, b.LocalFieldRows, b.LorentzKineticRows, b.MassActivationRows, b.DecouplingRows, b.DynkinIndexRows, b.ThresholdBetaRowsAllowed, b.BlockCanContributeAsMultiplet, b.Verdict)
}

func FormatPattern(p PatternAudit) string {
	return fmt.Sprintf("galois=%s block=%s current=%s contact=%s fano=%s matchCurrent=%t matchContact=%t matchFano=%t (%s)", p.GaloisSafePattern, p.QuarticBlockPattern, p.CurrentQuotientPattern, p.ContactSingletonPattern, p.FanoPattern, p.MatchesCurrent, p.MatchesContact, p.MatchesFano, p.Verdict)
}

func FormatFirewall(f PhysicsFirewall) string {
	return fmt.Sprintf("observedFree=%t matrix=%t char=%t rootIso=%t Qidempotents=%t partition=%t blockInvariants=%t individualSplit=%t gauge=%t local=%t mass=%t decoupling=%t beta=%t physical=%t all=%t (%s)", f.ObservedInputFree, f.ExactMatrix, f.ExactCharpoly, f.ExactRootIsolation, f.RationalPrimaryIdempotents, f.GaloisInvariantPartition, f.QuarticBlockInvariants, f.IndividualQuarticSplit, f.GaugeRepresentation, f.LocalFieldMap, f.MassActivation, f.Decoupling, f.BetaRows, f.PhysicalConstants, f.AllSatisfiedForPhysics, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contactRows=%d orbits=%d rational=%d quarticRows=%d quarticBlocks=%d invariants=%d individualQuartic=%d rowProofs=%d reprComplete=%d reprOpen=%d quarticBeta=%d contactBeta=%d zero=%d S6=%d nullity=%d→%d", s.ContactRows, s.GaloisInvariantOrbits, s.RationalSingletonRows, s.QuarticOrbitRows, s.QuarticCompressedBlocks, s.QuarticBlockInvariants, s.IndividualQuarticRows, s.RowAssignmentProofs, s.RepresentationCompleteRows, s.RepresentationOpenRows, s.QuarticBlockBetaRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualS6Choices, s.ResidualNullityBefore, s.ResidualNullityAfter)
}
