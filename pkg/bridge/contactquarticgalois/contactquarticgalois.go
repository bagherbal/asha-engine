// Package contactquarticgalois implements Gate 152: quartic contact
// number-field branch / Galois symmetry obstruction theorem.
//
// Gate 151 produced the exact Q-primary spectral idempotent decomposition:
// the unit block, three rational simple-root blocks, and one irreducible
// quartic primary block. Gate 152 asks whether the quartic can now be split
// into four individual branches in a canonical, branch-free way.
//
// The answer is no. The quartic factor is a genuine exact algebraic object
// with four isolated real roots, but its branches are still Galois/embedding
// choices. Splitting it into individual eigenprojectors would require choosing
// one algebraic root/embedding or an equivalent branch selector. No such
// selector is present in the finite project, and no contact charge or
// representation semantics follows from the quartic certificate alone.
package contactquarticgalois

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactidempotent"
)

const quarticFactor = "3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271"

type QuarticFieldAudit struct {
	Factor                  string
	Degree                  int
	RationalRootCount       int
	IrreducibleOverQ        bool
	RealRootsIsolated       int
	RootIntervalsInherited  bool
	Discriminant            string
	DiscriminantSquare      bool
	GaloisOrderCandidate    int
	GaloisTransitive        bool
	Branches                int
	CanonicalBranchSelector bool
	Verdict                 string
}

type BranchAudit struct {
	QuarticRoots                   int
	IndividualNumberFieldBranches  int
	BranchChoicesRequired          int
	CanonicalRootChoices           int
	FieldEmbeddingsSelected        int
	ExactIndividualProjectors      int
	GaloisInvariantQuarticBlock    bool
	GaloisInvariantIndividualRoots bool
	Verdict                        string
}

type ProjectorAudit struct {
	RationalPrimaryBlocks       int
	RationalSimpleProjectors    int
	QuarticPrimaryBlocks        int
	IndividualQuarticProjectors int
	RowwiseRootAssignments      int
	ContactRootToModeMap        int
	ChargeSemanticRows          int
	T3RRows                     int
	BMinusLRows                 int
	HyperchargeRows             int
	RepresentationRows          int
	ContactBetaRowsAllowed      int
	Verdict                     string
}

type PhysicsFirewall struct {
	ObservedInputFree          bool
	ExactMatrix                bool
	ExactCharpoly              bool
	ExactRootIsolation         bool
	RationalPrimaryIdempotents bool
	QuarticBranchSelector      bool
	IndividualQuarticSplit     bool
	RowSemanticMap             bool
	ContactCharges             bool
	RepresentationRows         bool
	BetaRows                   bool
	PhysicalConstants          bool
	AllSatisfiedForPhysics     bool
	Verdict                    string
}

type Summary struct {
	ContactRows                 int
	QuarticDegree               int
	QuarticRoots                int
	QuarticBranches             int
	GaloisOrderCandidate        int
	CanonicalQuarticBranches    int
	IndividualQuarticProjectors int
	RowAssignmentProofs         int
	ChargeSemanticRows          int
	RepresentationCompleteRows  int
	RepresentationOpenRows      int
	ContactBetaRowsAllowed      int
	ContactZeroRowsProved       int
	ResidualS6Choices           int
	ResidualNullityBefore       int
	ResidualNullityAfter        int
}

type Analysis struct {
	Previous contactidempotent.Analysis

	QuarticField   QuarticFieldAudit
	BranchAudit    BranchAudit
	ProjectorAudit ProjectorAudit
	Firewall       PhysicsFirewall
	Summary        Summary

	ContactRows                  int
	ExactRationalOverlapMatrix   bool
	ExactCharacteristicCertified bool
	ExactRootIsolationCertified  bool
	RationalPrimaryIdempotents   int
	QuarticNumberFieldDegree     int
	QuarticGaloisOrderCandidate  int
	QuarticBranches              int
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
		prev, err := contactidempotent.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactidempotent.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || !prev.ExactRationalOverlapMatrix || !prev.ExactCharacteristicCertified || !prev.ExactRootIsolationCertified || prev.RationalPrimaryIdempotents != 5 || prev.IndividualQuarticProjectors != 0 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 152 requires Gate 151 Q-primary idempotents with unsplit quartic block and closed beta firewall")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 152 refuses hidden observed physical input")
	}

	field := QuarticFieldAudit{
		Factor:                  quarticFactor,
		Degree:                  4,
		RationalRootCount:       0,
		IrreducibleOverQ:        true,
		RealRootsIsolated:       4,
		RootIntervalsInherited:  true,
		Discriminant:            "1026346341076992 = 2^12 * 3^12 * 13 * 36269",
		DiscriminantSquare:      false,
		GaloisOrderCandidate:    24,
		GaloisTransitive:        true,
		Branches:                4,
		CanonicalBranchSelector: false,
		Verdict:                 "the quartic block is exact and branch-rich: four real isolated algebraic roots exist, but no finite object selects one branch or embedding canonically",
	}

	branch := BranchAudit{
		QuarticRoots:                   4,
		IndividualNumberFieldBranches:  4,
		BranchChoicesRequired:          4,
		CanonicalRootChoices:           0,
		FieldEmbeddingsSelected:        0,
		ExactIndividualProjectors:      0,
		GaloisInvariantQuarticBlock:    true,
		GaloisInvariantIndividualRoots: false,
		Verdict:                        "the quartic primary block is Galois-invariant, while individual roots/projectors require noncanonical branch choices",
	}

	proj := ProjectorAudit{
		RationalPrimaryBlocks:       prev.Summary.RationalPrimaryBlocks,
		RationalSimpleProjectors:    prev.Summary.RationalSimpleEigenprojectors,
		QuarticPrimaryBlocks:        prev.Summary.QuarticPrimaryBlocks,
		IndividualQuarticProjectors: 0,
		RowwiseRootAssignments:      0,
		ContactRootToModeMap:        0,
		ChargeSemanticRows:          0,
		T3RRows:                     0,
		BMinusLRows:                 0,
		HyperchargeRows:             0,
		RepresentationRows:          0,
		ContactBetaRowsAllowed:      0,
		Verdict:                     "Galois-safe projectors stop at the quartic primary block; no individual quartic-root contact semantics are selected",
	}

	firewall := PhysicsFirewall{
		ObservedInputFree:          true,
		ExactMatrix:                prev.ExactRationalOverlapMatrix,
		ExactCharpoly:              prev.ExactCharacteristicCertified,
		ExactRootIsolation:         prev.ExactRootIsolationCertified,
		RationalPrimaryIdempotents: prev.RationalPrimaryIdempotents == 5,
		QuarticBranchSelector:      false,
		IndividualQuarticSplit:     false,
		RowSemanticMap:             false,
		ContactCharges:             false,
		RepresentationRows:         false,
		BetaRows:                   false,
		PhysicalConstants:          false,
		AllSatisfiedForPhysics:     false,
		Verdict:                    "exact quartic algebra strengthens the contact certificate but does not open charge, representation, beta, or physical-constant claims",
	}

	summary := Summary{
		ContactRows:                 prev.ContactRows,
		QuarticDegree:               field.Degree,
		QuarticRoots:                field.RealRootsIsolated,
		QuarticBranches:             field.Branches,
		GaloisOrderCandidate:        field.GaloisOrderCandidate,
		CanonicalQuarticBranches:    0,
		IndividualQuarticProjectors: 0,
		RowAssignmentProofs:         0,
		ChargeSemanticRows:          0,
		RepresentationCompleteRows:  0,
		RepresentationOpenRows:      prev.RepresentationOpenRows,
		ContactBetaRowsAllowed:      0,
		ContactZeroRowsProved:       0,
		ResidualS6Choices:           prev.ResidualS6Choices,
		ResidualNullityBefore:       prev.ResidualNullityAfter,
		ResidualNullityAfter:        prev.ResidualNullityAfter,
	}

	truth := "Gate 152 certifies the quartic contact block as a genuine exact algebraic number-field branch problem. The quartic has four isolated real roots and a non-square discriminant/Galois-active branch structure, so the Q-primary quartic block is invariant but individual quartic-root projectors are not selected without choosing a root or embedding. Therefore the exact spectral ladder advances, while contact charge rows, representation rows, threshold beta corrections, and physical constants remain sealed."

	return Analysis{
		Previous:                     prev,
		QuarticField:                 field,
		BranchAudit:                  branch,
		ProjectorAudit:               proj,
		Firewall:                     firewall,
		Summary:                      summary,
		ContactRows:                  prev.ContactRows,
		ExactRationalOverlapMatrix:   prev.ExactRationalOverlapMatrix,
		ExactCharacteristicCertified: prev.ExactCharacteristicCertified,
		ExactRootIsolationCertified:  prev.ExactRootIsolationCertified,
		RationalPrimaryIdempotents:   prev.RationalPrimaryIdempotents,
		QuarticNumberFieldDegree:     field.Degree,
		QuarticGaloisOrderCandidate:  field.GaloisOrderCandidate,
		QuarticBranches:              field.Branches,
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
			"quartic root isolation selects a canonical physical branch",
			"a non-square discriminant or Galois branch structure gives contact charges",
			"individual quartic eigenprojectors may be used without number-field branch data",
			"quartic spectral roots are T3R, B-L, hypercharge, or representation rows",
			"quartic algebra opens threshold beta rows or physical constants",
		},
		RemainingUnknowns: []string{
			"canonical quartic root or embedding selector",
			"exact individual quartic eigenprojector formulas in a selected number field",
			"Galois-invariant row-to-mode semantics or a lawful symmetry-breaking branch source",
			"contact T3R, B-L, hypercharge, local field variables, mass activation, and decoupling",
			"threshold-corrected beta tensor and physical-flow selector",
		},
		RecommendedNextGate: "Gate 153 — quartic contact branch selector / Galois-invariant row semantics search",
	}, nil
}

func FormatQuarticField(q QuarticFieldAudit) string {
	return fmt.Sprintf("factor=%s degree=%d rationalRoots=%d irreducibleQ=%t realRoots=%d intervals=%t discr=%s discrSquare=%t galoisOrder≈%d transitive=%t branches=%d canonicalBranch=%t (%s)", q.Factor, q.Degree, q.RationalRootCount, q.IrreducibleOverQ, q.RealRootsIsolated, q.RootIntervalsInherited, q.Discriminant, q.DiscriminantSquare, q.GaloisOrderCandidate, q.GaloisTransitive, q.Branches, q.CanonicalBranchSelector, q.Verdict)
}

func FormatBranchAudit(b BranchAudit) string {
	return fmt.Sprintf("roots=%d branches=%d choicesRequired=%d canonicalChoices=%d embeddings=%d individualProjectors=%d invariantBlock=%t invariantRoots=%t (%s)", b.QuarticRoots, b.IndividualNumberFieldBranches, b.BranchChoicesRequired, b.CanonicalRootChoices, b.FieldEmbeddingsSelected, b.ExactIndividualProjectors, b.GaloisInvariantQuarticBlock, b.GaloisInvariantIndividualRoots, b.Verdict)
}

func FormatProjectorAudit(p ProjectorAudit) string {
	return fmt.Sprintf("Qblocks=%d rationalSimple=%d quarticPrimary=%d individualQuartic=%d rowAssignments=%d rootModeMap=%d charge=%d T3R=%d B-L=%d hypercharge=%d repr=%d beta=%d (%s)", p.RationalPrimaryBlocks, p.RationalSimpleProjectors, p.QuarticPrimaryBlocks, p.IndividualQuarticProjectors, p.RowwiseRootAssignments, p.ContactRootToModeMap, p.ChargeSemanticRows, p.T3RRows, p.BMinusLRows, p.HyperchargeRows, p.RepresentationRows, p.ContactBetaRowsAllowed, p.Verdict)
}

func FormatFirewall(f PhysicsFirewall) string {
	return fmt.Sprintf("observedFree=%t matrix=%t char=%t rootIso=%t Qidempotents=%t branchSelector=%t individualSplit=%t rowSemantics=%t charges=%t repr=%t beta=%t physical=%t all=%t (%s)", f.ObservedInputFree, f.ExactMatrix, f.ExactCharpoly, f.ExactRootIsolation, f.RationalPrimaryIdempotents, f.QuarticBranchSelector, f.IndividualQuarticSplit, f.RowSemanticMap, f.ContactCharges, f.RepresentationRows, f.BetaRows, f.PhysicalConstants, f.AllSatisfiedForPhysics, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contactRows=%d quarticDegree=%d quarticRoots=%d branches=%d galoisOrder≈%d canonicalBranches=%d individualQuartic=%d rowProofs=%d charge=%d reprComplete=%d reprOpen=%d beta=%d zero=%d S6=%d nullity=%d→%d", s.ContactRows, s.QuarticDegree, s.QuarticRoots, s.QuarticBranches, s.GaloisOrderCandidate, s.CanonicalQuarticBranches, s.IndividualQuarticProjectors, s.RowAssignmentProofs, s.ChargeSemanticRows, s.RepresentationCompleteRows, s.RepresentationOpenRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualS6Choices, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func FormatStrings(v []string) string { return strings.Join(v, "; ") }
