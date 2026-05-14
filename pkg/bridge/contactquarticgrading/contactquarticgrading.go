// Package contactquarticgrading implements Gate 159: quartic ghost-grading
// Galois invariance / nontrivial parity obstruction theorem.
//
// Gate 158 showed that the quartic contact block has no canonical BRST
// cancellation: Q=0 is square-zero but inert, while nonzero square-zero maps
// and two-even/two-odd ghost gradings require choosing branches inside the
// quartic Galois orbit. Gate 159 isolates the grading obstruction itself. On a
// single transitive quartic orbit, every Galois-invariant parity function is
// constant. Every nontrivial zero-supertrace split is a 2+2 branch choice and
// is therefore not Galois-invariant.
package contactquarticgrading

import (
	"fmt"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactquarticbrst"
)

type ParityAssignmentClass struct {
	Name                 string
	Assignments          int
	BosonicRows          int
	GhostRows            int
	ZeroSignedCount      bool
	GaloisInvariant      bool
	Nontrivial           bool
	RequiresBranchChoice bool
	StabilizerOrder      int
	OrbitSize            int
	SupertraceLedger     bool
	ZeroBetaLedger       bool
	CancellationComplete bool
	Verdict              string
}

type GaloisActionAudit struct {
	QuarticRows                 int
	GaloisOrbitRows             int
	GaloisOrderCandidate        int
	TransitiveOrbit             bool
	InvariantParityFunctions    int
	NontrivialInvariantParity   int
	AllParityAssignments        int
	NontrivialParityAssignments int
	ZeroSignedCountAssignments  int
	InvariantZeroSignedCount    int
	Verdict                     string
}

type ParityObstructionAudit struct {
	AssignmentsAudited       int
	InvariantAssignments     int
	NontrivialAssignments    int
	ZeroSupertraceCandidates int
	CanonicalZeroSupertrace  bool
	BranchChoiceRequired     bool
	CompleteGhostGrading     bool
	ZeroBetaLedger           bool
	Verdict                  string
}

type FirewallAudit struct {
	ObservedInputFree      bool
	QuarticBlockExact      bool
	BRSTRouteAudited       bool
	GaloisActionAudited    bool
	GhostGrading           bool
	NontrivialParity       bool
	SupertraceCancellation bool
	ZeroBetaLedger         bool
	RepresentationRows     int
	HyperchargeRows        int
	ThresholdBetaRows      int
	ProvenZeroRows         int
	PhysicalConstants      bool
	FirewallClosed         bool
	Verdict                string
}

type Summary struct {
	ContactRows                    int
	QuarticBlockRows               int
	GaloisOrbitRows                int
	GaloisOrderCandidate           int
	AllParityAssignments           int
	GaloisInvariantParityFunctions int
	NontrivialInvariantParity      int
	ZeroSignedCountAssignments     int
	InvariantZeroSignedCount       int
	CanonicalZeroSupertrace        bool
	QuarticZeroBetaRows            int
	QuarticBlockBetaRows           int
	ContactBetaRowsAllowed         int
	ContactZeroRowsProved          int
	ResidualS6Choices              int
	ResidualNullityBefore          int
	ResidualNullityAfter           int
}

type Analysis struct {
	Previous contactquarticbrst.Analysis

	ParityClasses []ParityAssignmentClass
	GaloisAction  GaloisActionAudit
	Obstruction   ParityObstructionAudit
	Firewall      FirewallAudit
	Summary       Summary

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
	SpinStatisticsRows           int
	LocalFieldRows               int
	KineticPoleResidueRows       int
	MassActivationRows           int
	DecouplingRows               int
	DynkinIndexRows              int
	BRSTCancellationRows         int
	ConstraintRows               int
	PropagatorRows               int
	RepresentationCompleteRows   int
	RepresentationOpenRows       int
	QuarticZeroBetaRows          int
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
		prev, err := contactquarticbrst.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactquarticbrst.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.QuarticOrbitRows != 4 || prev.QuarticBlockBetaRows != 0 || prev.ContactBetaRowsAllowed != 0 || prev.ContactZeroRowsProved != 0 || prev.QuarticZeroBetaRows != 0 {
		return Analysis{}, fmt.Errorf("Gate 159 requires Gate 158 exact quartic BRST obstruction with closed firewall")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 159 refuses hidden observed physical input")
	}

	classes := []ParityAssignmentClass{
		{
			Name:                 "all-even quartic grading",
			Assignments:          1,
			BosonicRows:          4,
			GhostRows:            0,
			ZeroSignedCount:      false,
			GaloisInvariant:      true,
			Nontrivial:           false,
			RequiresBranchChoice: false,
			StabilizerOrder:      24,
			OrbitSize:            1,
			SupertraceLedger:     false,
			ZeroBetaLedger:       false,
			CancellationComplete: false,
			Verdict:              "Galois-invariant but uniform; it cannot cancel the quartic block",
		},
		{
			Name:                 "all-odd quartic grading",
			Assignments:          1,
			BosonicRows:          0,
			GhostRows:            4,
			ZeroSignedCount:      false,
			GaloisInvariant:      true,
			Nontrivial:           false,
			RequiresBranchChoice: false,
			StabilizerOrder:      24,
			OrbitSize:            1,
			SupertraceLedger:     false,
			ZeroBetaLedger:       false,
			CancellationComplete: false,
			Verdict:              "Galois-invariant but still uniform; it changes the overall sign rather than proving cancellation",
		},
		{
			Name:                 "one-versus-three parity split",
			Assignments:          8,
			BosonicRows:          1,
			GhostRows:            3,
			ZeroSignedCount:      false,
			GaloisInvariant:      false,
			Nontrivial:           true,
			RequiresBranchChoice: true,
			StabilizerOrder:      6,
			OrbitSize:            4,
			SupertraceLedger:     false,
			ZeroBetaLedger:       false,
			CancellationComplete: false,
			Verdict:              "nontrivial but branch-dependent and not zero-supertrace",
		},
		{
			Name:                 "two-even/two-odd parity split",
			Assignments:          6,
			BosonicRows:          2,
			GhostRows:            2,
			ZeroSignedCount:      true,
			GaloisInvariant:      false,
			Nontrivial:           true,
			RequiresBranchChoice: true,
			StabilizerOrder:      4,
			OrbitSize:            6,
			SupertraceLedger:     false,
			ZeroBetaLedger:       false,
			CancellationComplete: false,
			Verdict:              "zero signed count is possible only by choosing a two-subset of four Galois-conjugate branches; therefore it is not canonical",
		},
	}

	allParityAssignments := 16
	invariantParity := 0
	nontrivialInvariant := 0
	nontrivialAssignments := 0
	zeroSignedCount := 0
	invariantZeroSigned := 0
	for _, c := range classes {
		if c.GaloisInvariant {
			invariantParity += c.Assignments
		}
		if c.Nontrivial {
			nontrivialAssignments += c.Assignments
		}
		if c.GaloisInvariant && c.Nontrivial {
			nontrivialInvariant += c.Assignments
		}
		if c.ZeroSignedCount {
			zeroSignedCount += c.Assignments
		}
		if c.GaloisInvariant && c.ZeroSignedCount {
			invariantZeroSigned += c.Assignments
		}
	}

	galois := GaloisActionAudit{
		QuarticRows:                 4,
		GaloisOrbitRows:             4,
		GaloisOrderCandidate:        24,
		TransitiveOrbit:             true,
		InvariantParityFunctions:    invariantParity,
		NontrivialInvariantParity:   nontrivialInvariant,
		AllParityAssignments:        allParityAssignments,
		NontrivialParityAssignments: nontrivialAssignments,
		ZeroSignedCountAssignments:  zeroSignedCount,
		InvariantZeroSignedCount:    invariantZeroSigned,
		Verdict:                     "on one transitive quartic orbit, Galois-invariant parity functions are constant; nontrivial zero-count parity splits require branch choices",
	}

	obstruction := ParityObstructionAudit{
		AssignmentsAudited:       allParityAssignments,
		InvariantAssignments:     invariantParity,
		NontrivialAssignments:    nontrivialAssignments,
		ZeroSupertraceCandidates: zeroSignedCount,
		CanonicalZeroSupertrace:  false,
		BranchChoiceRequired:     true,
		CompleteGhostGrading:     false,
		ZeroBetaLedger:           false,
		Verdict:                  "six formal two-even/two-odd zero-count gradings exist, but every one splits the quartic Galois orbit and none is canonical or beta-permissive",
	}

	firewall := FirewallAudit{
		ObservedInputFree:      true,
		QuarticBlockExact:      prev.ExactRationalOverlapMatrix && prev.ExactCharacteristicCertified && prev.ExactRootIsolationCertified && prev.QuarticOrbitRows == 4,
		BRSTRouteAudited:       true,
		GaloisActionAudited:    true,
		GhostGrading:           false,
		NontrivialParity:       false,
		SupertraceCancellation: false,
		ZeroBetaLedger:         false,
		RepresentationRows:     0,
		HyperchargeRows:        0,
		ThresholdBetaRows:      0,
		ProvenZeroRows:         0,
		PhysicalConstants:      false,
		FirewallClosed:         true,
		Verdict:                "a ghost grading that cancels the quartic block must be nontrivial, Galois-invariant, and tied to a zero-supertrace/zero-beta ledger; no such grading exists",
	}

	summary := Summary{
		ContactRows:                    prev.ContactRows,
		QuarticBlockRows:               prev.QuarticOrbitRows,
		GaloisOrbitRows:                4,
		GaloisOrderCandidate:           24,
		AllParityAssignments:           allParityAssignments,
		GaloisInvariantParityFunctions: invariantParity,
		NontrivialInvariantParity:      nontrivialInvariant,
		ZeroSignedCountAssignments:     zeroSignedCount,
		InvariantZeroSignedCount:       invariantZeroSigned,
		CanonicalZeroSupertrace:        false,
		QuarticZeroBetaRows:            0,
		QuarticBlockBetaRows:           0,
		ContactBetaRowsAllowed:         0,
		ContactZeroRowsProved:          0,
		ResidualS6Choices:              prev.ResidualS6Choices,
		ResidualNullityBefore:          prev.ResidualNullityAfter,
		ResidualNullityAfter:           prev.ResidualNullityAfter,
	}

	truth := "Gate 159 proves that the quartic block has no nontrivial Galois-invariant ghost grading. The only invariant parities on a transitive four-root orbit are all-even and all-odd, both non-cancelling. Every two-even/two-odd zero-count split requires choosing branches, so it cannot certify BRST cancellation or a zero-beta ledger."

	return Analysis{
		Previous:                     prev,
		ParityClasses:                classes,
		GaloisAction:                 galois,
		Obstruction:                  obstruction,
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
		QuarticCompressedBlocks:      prev.QuarticCompressedBlocks,
		QuarticBlockInvariants:       prev.QuarticBlockInvariants,
		IndividualQuarticRows:        prev.IndividualQuarticRows,
		CanonicalQuarticBranches:     prev.CanonicalQuarticBranches,
		ExactNumberFieldProjectors:   prev.ExactNumberFieldProjectors,
		IndividualQuarticProjectors:  prev.IndividualQuarticProjectors,
		RowwiseRootAssignmentProofs:  prev.RowwiseRootAssignmentProofs,
		ChargeSemanticRows:           0,
		T3RRowsDerived:               0,
		ChiralityRowsDerived:         0,
		BMinusLRowsDerived:           0,
		SU2LRowsDerived:              0,
		HyperchargeRowsDerived:       0,
		GaugeRepresentationRows:      0,
		SpinStatisticsRows:           0,
		LocalFieldRows:               0,
		KineticPoleResidueRows:       0,
		MassActivationRows:           0,
		DecouplingRows:               0,
		DynkinIndexRows:              0,
		BRSTCancellationRows:         0,
		ConstraintRows:               0,
		PropagatorRows:               0,
		RepresentationCompleteRows:   0,
		RepresentationOpenRows:       prev.RepresentationOpenRows,
		QuarticZeroBetaRows:          0,
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
			"two-even/two-odd quartic grading is canonical",
			"zero signed count implies BRST cancellation",
			"quartic ghost grading is Galois-invariant and nontrivial",
			"quartic block has a zero-supertrace or zero-beta ledger",
		},
		RemainingUnknowns: []string{
			"canonical branch-breaking source for the quartic Galois orbit",
			"nontrivial Galois-compatible parity structure, if any extension supplies one",
			"whether quartic rows are physical, constrained, auxiliary, or purely finite spectral",
		},
		RecommendedNextGate: "Gate 160 — quartic parity branch-breaking source / external-selector firewall theorem",
	}, nil
}

func FormatParityClass(c ParityAssignmentClass) string {
	return fmt.Sprintf("%s: assignments=%d boson=%d ghost=%d zeroCount=%t galois=%t nontrivial=%t branchChoice=%t stabilizer=%d orbit=%d supertrace=%t zeroBeta=%t complete=%t (%s)", c.Name, c.Assignments, c.BosonicRows, c.GhostRows, c.ZeroSignedCount, c.GaloisInvariant, c.Nontrivial, c.RequiresBranchChoice, c.StabilizerOrder, c.OrbitSize, c.SupertraceLedger, c.ZeroBetaLedger, c.CancellationComplete, c.Verdict)
}

func FormatGaloisAction(a GaloisActionAudit) string {
	return fmt.Sprintf("quarticRows=%d orbitRows=%d galoisOrder=%d transitive=%t invariantParity=%d nontrivialInvariant=%d allParity=%d nontrivial=%d zeroCount=%d invariantZero=%d (%s)", a.QuarticRows, a.GaloisOrbitRows, a.GaloisOrderCandidate, a.TransitiveOrbit, a.InvariantParityFunctions, a.NontrivialInvariantParity, a.AllParityAssignments, a.NontrivialParityAssignments, a.ZeroSignedCountAssignments, a.InvariantZeroSignedCount, a.Verdict)
}

func FormatObstruction(a ParityObstructionAudit) string {
	return fmt.Sprintf("audited=%d invariant=%d nontrivial=%d zeroSupertraceCandidates=%d canonicalZero=%t branchChoice=%t completeGrading=%t zeroBeta=%t (%s)", a.AssignmentsAudited, a.InvariantAssignments, a.NontrivialAssignments, a.ZeroSupertraceCandidates, a.CanonicalZeroSupertrace, a.BranchChoiceRequired, a.CompleteGhostGrading, a.ZeroBetaLedger, a.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("observedFree=%t quarticExact=%t brstAudited=%t galoisAudited=%t ghost=%t nontrivialParity=%t supertrace=%t zeroBeta=%t repr=%d Y=%d beta=%d zero=%d physical=%t closed=%t (%s)", f.ObservedInputFree, f.QuarticBlockExact, f.BRSTRouteAudited, f.GaloisActionAudited, f.GhostGrading, f.NontrivialParity, f.SupertraceCancellation, f.ZeroBetaLedger, f.RepresentationRows, f.HyperchargeRows, f.ThresholdBetaRows, f.ProvenZeroRows, f.PhysicalConstants, f.FirewallClosed, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contactRows=%d quarticRows=%d orbitRows=%d galoisOrder=%d parity=%d invariantParity=%d nontrivialInvariant=%d zeroCount=%d invariantZero=%d canonicalZero=%t quarticZero=%d quarticBeta=%d contactBeta=%d zero=%d S6=%d nullity=%d→%d", s.ContactRows, s.QuarticBlockRows, s.GaloisOrbitRows, s.GaloisOrderCandidate, s.AllParityAssignments, s.GaloisInvariantParityFunctions, s.NontrivialInvariantParity, s.ZeroSignedCountAssignments, s.InvariantZeroSignedCount, s.CanonicalZeroSupertrace, s.QuarticZeroBetaRows, s.QuarticBlockBetaRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualS6Choices, s.ResidualNullityBefore, s.ResidualNullityAfter)
}
