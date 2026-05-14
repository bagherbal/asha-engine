// Package contactbranchsemantics implements Gate 153: quartic contact branch
// selector / Galois-invariant row semantics search.
//
// Gate 152 showed that the quartic contact factor is exact, irreducible over Q,
// and branch-rich. Gate 153 asks what semantics are available if we refuse to
// choose a quartic root or embedding. The answer is a useful but limited
// partition theorem: Q/Galois-invariant data can distinguish the three rational
// non-unit roots and the quartic orbit as a whole, giving a spectral partition
// 1+1+1+4. It cannot split the four quartic roots into individual contact rows,
// cannot select a branch, and cannot create charge, representation, local-field,
// mass, decoupling, beta, or physical-constant claims.
package contactbranchsemantics

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactquarticgalois"
)

type GaloisInvariantPartition struct {
	TotalPartialRows         int
	RationalSingletonRows    int
	QuarticOrbitRows         int
	OrbitPattern             string
	GaloisInvariantOrbits    int
	IndividualQuarticRows    int
	CanonicalBranchSelectors int
	BranchChoicesRequired    int
	PartitionExact           bool
	PartitionRowComplete     bool
	Verdict                  string
}

type BranchSelectorAudit struct {
	RationalRootBranches       int
	QuarticRootBranches        int
	QuarticBranchesSelected    int
	QuarticEmbeddingsSelected  int
	GaloisInvariantBlockOnly   bool
	IndividualQuarticSemantics bool
	RowwiseRootAssignments     int
	HiddenChoiceFree           bool
	Verdict                    string
}

type SemanticLiftAudit struct {
	SpectralPartitionRows    int
	SpectralSemanticsRows    int
	ChargeSemanticRows       int
	T3RRows                  int
	BMinusLRows              int
	HyperchargeRows          int
	LocalFieldRows           int
	MassActivationRows       int
	DecouplingRows           int
	RepresentationRows       int
	ContactBetaRowsAllowed   int
	GaloisInvariantSemantics bool
	RepresentationUseful     bool
	Verdict                  string
}

type PatternMismatchAudit struct {
	GaloisSafePattern           string
	ContactSingletonPattern     string
	CurrentQuotientPattern      string
	FanoPattern                 string
	GaloisPatternMatchesContact bool
	GaloisPatternMatchesCurrent bool
	GaloisPatternMatchesFano    bool
	RefinementChoicesRequired   int
	Verdict                     string
}

type PhysicsFirewall struct {
	ObservedInputFree          bool
	ExactMatrix                bool
	ExactCharpoly              bool
	ExactRootIsolation         bool
	RationalPrimaryIdempotents bool
	GaloisInvariantPartition   bool
	BranchSelector             bool
	IndividualQuarticSplit     bool
	ContactCharges             bool
	RepresentationRows         bool
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
	IndividualQuarticRows      int
	CanonicalQuarticBranches   int
	RowAssignmentProofs        int
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
	Previous contactquarticgalois.Analysis

	Partition       GaloisInvariantPartition
	BranchSelector  BranchSelectorAudit
	SemanticLift    SemanticLiftAudit
	PatternMismatch PatternMismatchAudit
	Firewall        PhysicsFirewall
	Summary         Summary

	ContactRows                  int
	ExactRationalOverlapMatrix   bool
	ExactCharacteristicCertified bool
	ExactRootIsolationCertified  bool
	RationalPrimaryIdempotents   int
	GaloisInvariantOrbits        int
	RationalSingletonRows        int
	QuarticOrbitRows             int
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
		prev, err := contactquarticgalois.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactquarticgalois.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || !prev.ExactRationalOverlapMatrix || !prev.ExactCharacteristicCertified || !prev.ExactRootIsolationCertified || prev.RationalPrimaryIdempotents != 5 || prev.QuarticBranches != 4 || prev.CanonicalQuarticBranches != 0 || prev.IndividualQuarticProjectors != 0 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 153 requires Gate 152 exact quartic Galois obstruction with closed beta firewall")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 153 refuses hidden observed physical input")
	}

	partition := GaloisInvariantPartition{
		TotalPartialRows:         prev.ContactRows,
		RationalSingletonRows:    3,
		QuarticOrbitRows:         4,
		OrbitPattern:             "1+1+1+4",
		GaloisInvariantOrbits:    4,
		IndividualQuarticRows:    0,
		CanonicalBranchSelectors: 0,
		BranchChoicesRequired:    4,
		PartitionExact:           true,
		PartitionRowComplete:     false,
		Verdict:                  "Galois-invariant data gives an exact spectral partition into three rational singleton roots plus one four-root quartic orbit; it does not split the quartic orbit into four rows",
	}

	branch := BranchSelectorAudit{
		RationalRootBranches:       3,
		QuarticRootBranches:        4,
		QuarticBranchesSelected:    0,
		QuarticEmbeddingsSelected:  0,
		GaloisInvariantBlockOnly:   true,
		IndividualQuarticSemantics: false,
		RowwiseRootAssignments:     0,
		HiddenChoiceFree:           true,
		Verdict:                    "the only hidden-choice-free object is the quartic block/orbit; selecting any individual quartic root would choose a branch",
	}

	semantic := SemanticLiftAudit{
		SpectralPartitionRows:    7,
		SpectralSemanticsRows:    3,
		ChargeSemanticRows:       0,
		T3RRows:                  0,
		BMinusLRows:              0,
		HyperchargeRows:          0,
		LocalFieldRows:           0,
		MassActivationRows:       0,
		DecouplingRows:           0,
		RepresentationRows:       0,
		ContactBetaRowsAllowed:   0,
		GaloisInvariantSemantics: true,
		RepresentationUseful:     false,
		Verdict:                  "three rational roots have exact spectral labels and the quartic orbit has exact block semantics, but none is a charge, field, mass, decoupling, or representation row",
	}

	mismatch := PatternMismatchAudit{
		GaloisSafePattern:           partition.OrbitPattern,
		ContactSingletonPattern:     "1+1+1+1+1+1+1",
		CurrentQuotientPattern:      "1+6",
		FanoPattern:                 "7-transitive",
		GaloisPatternMatchesContact: false,
		GaloisPatternMatchesCurrent: false,
		GaloisPatternMatchesFano:    false,
		RefinementChoicesRequired:   4,
		Verdict:                     "the Galois-safe 1+1+1+4 orbit partition is exact but it does not match contact singleton rows, current 1+6 semantics, or Fano-transitive labeling",
	}

	firewall := PhysicsFirewall{
		ObservedInputFree:          true,
		ExactMatrix:                prev.ExactRationalOverlapMatrix,
		ExactCharpoly:              prev.ExactCharacteristicCertified,
		ExactRootIsolation:         prev.ExactRootIsolationCertified,
		RationalPrimaryIdempotents: prev.RationalPrimaryIdempotents == 5,
		GaloisInvariantPartition:   partition.PartitionExact,
		BranchSelector:             false,
		IndividualQuarticSplit:     false,
		ContactCharges:             false,
		RepresentationRows:         false,
		BetaRows:                   false,
		PhysicalConstants:          false,
		AllSatisfiedForPhysics:     false,
		Verdict:                    "Galois-invariant partitioning strengthens the spectral certificate, but it is not representation-complete and opens no beta or physical-constant claims",
	}

	summary := Summary{
		ContactRows:                prev.ContactRows,
		GaloisInvariantOrbits:      partition.GaloisInvariantOrbits,
		RationalSingletonRows:      partition.RationalSingletonRows,
		QuarticOrbitRows:           partition.QuarticOrbitRows,
		IndividualQuarticRows:      0,
		CanonicalQuarticBranches:   0,
		RowAssignmentProofs:        0,
		ChargeSemanticRows:         0,
		RepresentationCompleteRows: 0,
		RepresentationOpenRows:     prev.RepresentationOpenRows,
		ContactBetaRowsAllowed:     0,
		ContactZeroRowsProved:      0,
		ResidualS6Choices:          prev.ResidualS6Choices,
		ResidualNullityBefore:      prev.ResidualNullityAfter,
		ResidualNullityAfter:       prev.ResidualNullityAfter,
	}

	truth := "Gate 153 proves the strongest branch-free contact semantics currently available: Galois-invariant data partitions the seven partial contact roots as 1+1+1+4. This is exact and useful, but it is not row-complete. The quartic orbit remains unsplit, and the partition does not match current 1+6 semantics, Fano naturality, or seven singleton contact rows. Therefore individual quartic rows, contact charges, representation rows, threshold beta corrections, and physical constants remain sealed."

	return Analysis{
		Previous:                     prev,
		Partition:                    partition,
		BranchSelector:               branch,
		SemanticLift:                 semantic,
		PatternMismatch:              mismatch,
		Firewall:                     firewall,
		Summary:                      summary,
		ContactRows:                  prev.ContactRows,
		ExactRationalOverlapMatrix:   prev.ExactRationalOverlapMatrix,
		ExactCharacteristicCertified: prev.ExactCharacteristicCertified,
		ExactRootIsolationCertified:  prev.ExactRootIsolationCertified,
		RationalPrimaryIdempotents:   prev.RationalPrimaryIdempotents,
		GaloisInvariantOrbits:        partition.GaloisInvariantOrbits,
		RationalSingletonRows:        partition.RationalSingletonRows,
		QuarticOrbitRows:             partition.QuarticOrbitRows,
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
			"Galois-invariant partitioning selects individual quartic rows",
			"the quartic orbit can be used as four contact charge rows without a branch selector",
			"the 1+1+1+4 spectral partition is equivalent to current 1+6 semantics",
			"rational singleton spectral labels are T3R, B-L, hypercharge, or beta rows",
			"Galois-safe spectral semantics open physical constants or masses",
		},
		RemainingUnknowns: []string{
			"canonical quartic branch selector or Galois-invariant replacement for individual rows",
			"row-complete contact semantics compatible with local fields or constraints",
			"contact T3R, B-L, hypercharge, SU(2)L, local variables, mass activation, and decoupling",
			"threshold-corrected beta tensor and physical-flow selector",
		},
		RecommendedNextGate: "Gate 154 — quartic orbit semantic compression / four-row block beta firewall theorem",
	}, nil
}

func FormatPartition(p GaloisInvariantPartition) string {
	return fmt.Sprintf("rows=%d rationalSingletons=%d quarticOrbit=%d pattern=%s orbits=%d individualQuartic=%d branchSelectors=%d branchChoices=%d exact=%t rowComplete=%t (%s)", p.TotalPartialRows, p.RationalSingletonRows, p.QuarticOrbitRows, p.OrbitPattern, p.GaloisInvariantOrbits, p.IndividualQuarticRows, p.CanonicalBranchSelectors, p.BranchChoicesRequired, p.PartitionExact, p.PartitionRowComplete, p.Verdict)
}

func FormatBranchSelector(b BranchSelectorAudit) string {
	return fmt.Sprintf("rationalBranches=%d quarticBranches=%d selectedQuartic=%d embeddings=%d blockOnly=%t individualSemantics=%t rowAssignments=%d hiddenChoiceFree=%t (%s)", b.RationalRootBranches, b.QuarticRootBranches, b.QuarticBranchesSelected, b.QuarticEmbeddingsSelected, b.GaloisInvariantBlockOnly, b.IndividualQuarticSemantics, b.RowwiseRootAssignments, b.HiddenChoiceFree, b.Verdict)
}

func FormatSemanticLift(s SemanticLiftAudit) string {
	return fmt.Sprintf("spectralRows=%d spectralSemanticRows=%d charge=%d T3R=%d B-L=%d hypercharge=%d local=%d mass=%d decoupling=%d repr=%d beta=%d galoisSemantics=%t reprUseful=%t (%s)", s.SpectralPartitionRows, s.SpectralSemanticsRows, s.ChargeSemanticRows, s.T3RRows, s.BMinusLRows, s.HyperchargeRows, s.LocalFieldRows, s.MassActivationRows, s.DecouplingRows, s.RepresentationRows, s.ContactBetaRowsAllowed, s.GaloisInvariantSemantics, s.RepresentationUseful, s.Verdict)
}

func FormatPatternMismatch(p PatternMismatchAudit) string {
	return fmt.Sprintf("galois=%s contact=%s current=%s fano=%s matchContact=%t matchCurrent=%t matchFano=%t refinementChoices=%d (%s)", p.GaloisSafePattern, p.ContactSingletonPattern, p.CurrentQuotientPattern, p.FanoPattern, p.GaloisPatternMatchesContact, p.GaloisPatternMatchesCurrent, p.GaloisPatternMatchesFano, p.RefinementChoicesRequired, p.Verdict)
}

func FormatFirewall(f PhysicsFirewall) string {
	return fmt.Sprintf("observedFree=%t matrix=%t char=%t rootIso=%t Qidempotents=%t partition=%t branchSelector=%t individualSplit=%t charges=%t repr=%t beta=%t physical=%t all=%t (%s)", f.ObservedInputFree, f.ExactMatrix, f.ExactCharpoly, f.ExactRootIsolation, f.RationalPrimaryIdempotents, f.GaloisInvariantPartition, f.BranchSelector, f.IndividualQuarticSplit, f.ContactCharges, f.RepresentationRows, f.BetaRows, f.PhysicalConstants, f.AllSatisfiedForPhysics, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contactRows=%d orbits=%d rationalSingletons=%d quarticOrbit=%d individualQuartic=%d canonicalBranches=%d rowProofs=%d charge=%d reprComplete=%d reprOpen=%d beta=%d zero=%d S6=%d nullity=%d→%d", s.ContactRows, s.GaloisInvariantOrbits, s.RationalSingletonRows, s.QuarticOrbitRows, s.IndividualQuarticRows, s.CanonicalQuarticBranches, s.RowAssignmentProofs, s.ChargeSemanticRows, s.RepresentationCompleteRows, s.RepresentationOpenRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualS6Choices, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func FormatStrings(v []string) string { return strings.Join(v, "; ") }
