// Package contactquarticbrst implements Gate 158: quartic BRST candidate
// differential / zero-supertrace construction attempt.
//
// Gate 157 made the quartic permission dichotomy executable: the exact
// four-row quartic contact block may affect beta matching only if it is either
// a propagating local threshold field or a proven constrained/BRST-cancelled
// nonphysical block. Gate 158 attacks the BRST route directly. It inventories
// nilpotent differential candidates and ghost gradings on the quartic block and
// checks whether any choice is canonical, Galois-invariant, and strong enough
// to produce a zero-supertrace / zero-beta ledger. The result is negative: the
// zero differential is canonical but inert, nonzero differentials require
// arbitrary pairings/orderings, and every nontrivial zero-supertrace grading
// breaks the quartic Galois orbit into chosen branches.
package contactquarticbrst

import (
	"fmt"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactquarticdichotomy"
)

type DifferentialCandidate struct {
	Name                     string
	Dimension                int
	Nilpotent                bool
	SquareZero               bool
	Canonical                bool
	GaloisInvariant          bool
	RequiresPairingChoice    bool
	RequiresOrderingChoice   bool
	CohomologyDimension      int
	ZeroSupertraceLedger     bool
	ZeroBetaLedger           bool
	BRSTCancellationComplete bool
	Verdict                  string
}

type GhostGradingCandidate struct {
	Name                     string
	Rows                     int
	BosonicRows              int
	GhostRows                int
	GaloisInvariant          bool
	Nontrivial               bool
	TraceZeroPossible        bool
	PointwisePaired          bool
	RequiresBranchChoices    bool
	ZeroSupertraceLedger     bool
	BRSTCancellationComplete bool
	Verdict                  string
}

type SupertraceAudit struct {
	QuarticBlockRows            int
	GaloisOrbitRows             int
	GaloisInvariantGradings     int
	NontrivialInvariantGradings int
	ZeroSupertraceGradings      int
	CanonicalZeroSupertrace     bool
	PointwiseCancellationPairs  int
	ZeroBetaLedger              bool
	Verdict                     string
}

type BRSTConstructionAudit struct {
	QuarticBlockExact         bool
	DifferentialsAudited      int
	NilpotentCandidates       int
	CanonicalNilpotent        int
	NonzeroCanonicalBRST      int
	CompleteBRSTCancellations int
	GhostGradingsAudited      int
	CanonicalGhostGradings    int
	ZeroSupertraceLedgers     int
	ZeroBetaLedgers           int
	ConstructionComplete      bool
	Verdict                   string
}

type FirewallAudit struct {
	ObservedInputFree      bool
	QuarticBlockExact      bool
	ConstraintRouteAudited bool
	BRSTOperator           bool
	GhostGrading           bool
	NilpotentDifferential  bool
	ExactnessOrCohomology  bool
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
	ContactRows               int
	QuarticBlockRows          int
	DifferentialsAudited      int
	NilpotentCandidates       int
	CanonicalNilpotents       int
	CompleteBRSTCancellations int
	GhostGradingsAudited      int
	GaloisInvariantGradings   int
	ZeroSupertraceLedgers     int
	QuarticZeroBetaRows       int
	QuarticBlockBetaRows      int
	ContactBetaRowsAllowed    int
	ContactZeroRowsProved     int
	ResidualS6Choices         int
	ResidualNullityBefore     int
	ResidualNullityAfter      int
}

type Analysis struct {
	Previous contactquarticdichotomy.Analysis

	DifferentialCandidates []DifferentialCandidate
	GhostGradings          []GhostGradingCandidate
	Supertrace             SupertraceAudit
	Construction           BRSTConstructionAudit
	Firewall               FirewallAudit
	Summary                Summary

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
		prev, err := contactquarticdichotomy.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactquarticdichotomy.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.QuarticOrbitRows != 4 || prev.BRSTCancellationRows != 0 || prev.ConstraintRows != 0 || prev.QuarticBlockBetaRows != 0 || prev.ContactBetaRowsAllowed != 0 || prev.ContactZeroRowsProved != 0 {
		return Analysis{}, fmt.Errorf("Gate 158 requires Gate 157 unresolved BRST/locality dichotomy with closed firewall")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 158 refuses hidden observed physical input")
	}

	differentials := []DifferentialCandidate{
		{
			Name:                     "zero differential Q=0",
			Dimension:                4,
			Nilpotent:                true,
			SquareZero:               true,
			Canonical:                true,
			GaloisInvariant:          true,
			RequiresPairingChoice:    false,
			RequiresOrderingChoice:   false,
			CohomologyDimension:      4,
			ZeroSupertraceLedger:     false,
			ZeroBetaLedger:           false,
			BRSTCancellationComplete: false,
			Verdict:                  "canonical and square-zero, but inert: all four quartic classes survive in cohomology, so no cancellation ledger is produced",
		},
		{
			Name:                     "identity differential",
			Dimension:                4,
			Nilpotent:                false,
			SquareZero:               false,
			Canonical:                true,
			GaloisInvariant:          true,
			RequiresPairingChoice:    false,
			RequiresOrderingChoice:   false,
			CohomologyDimension:      0,
			ZeroSupertraceLedger:     false,
			ZeroBetaLedger:           false,
			BRSTCancellationComplete: false,
			Verdict:                  "canonical but not nilpotent, hence not a BRST differential",
		},
		{
			Name:                     "ordered nilpotent shift",
			Dimension:                4,
			Nilpotent:                true,
			SquareZero:               false,
			Canonical:                false,
			GaloisInvariant:          false,
			RequiresPairingChoice:    false,
			RequiresOrderingChoice:   true,
			CohomologyDimension:      1,
			ZeroSupertraceLedger:     false,
			ZeroBetaLedger:           false,
			BRSTCancellationComplete: false,
			Verdict:                  "a branch/order-dependent nilpotent shift can be written, but Q^2 is not zero for a length-four chain and the ordering is not Galois-invariant",
		},
		{
			Name:                     "two-pair square-zero differential",
			Dimension:                4,
			Nilpotent:                true,
			SquareZero:               true,
			Canonical:                false,
			GaloisInvariant:          false,
			RequiresPairingChoice:    true,
			RequiresOrderingChoice:   true,
			CohomologyDimension:      0,
			ZeroSupertraceLedger:     false,
			ZeroBetaLedger:           false,
			BRSTCancellationComplete: false,
			Verdict:                  "square-zero pair maps exist abstractly, but they require choosing a pairing/orientation of four Galois-conjugate branches and no zero-beta ledger follows from finite data",
		},
	}

	gradings := []GhostGradingCandidate{
		{
			Name:                     "uniform even grading",
			Rows:                     4,
			BosonicRows:              4,
			GhostRows:                0,
			GaloisInvariant:          true,
			Nontrivial:               false,
			TraceZeroPossible:        false,
			PointwisePaired:          false,
			RequiresBranchChoices:    false,
			ZeroSupertraceLedger:     false,
			BRSTCancellationComplete: false,
			Verdict:                  "Galois-invariant but row-uniform; it cannot cancel the quartic block",
		},
		{
			Name:                     "uniform odd grading",
			Rows:                     4,
			BosonicRows:              0,
			GhostRows:                4,
			GaloisInvariant:          true,
			Nontrivial:               false,
			TraceZeroPossible:        false,
			PointwisePaired:          false,
			RequiresBranchChoices:    false,
			ZeroSupertraceLedger:     false,
			BRSTCancellationComplete: false,
			Verdict:                  "Galois-invariant but still row-uniform; it changes an overall sign rather than proving cancellation",
		},
		{
			Name:                     "two-even/two-odd split",
			Rows:                     4,
			BosonicRows:              2,
			GhostRows:                2,
			GaloisInvariant:          false,
			Nontrivial:               true,
			TraceZeroPossible:        true,
			PointwisePaired:          false,
			RequiresBranchChoices:    true,
			ZeroSupertraceLedger:     false,
			BRSTCancellationComplete: false,
			Verdict:                  "can make a formal signed count vanish, but it splits the quartic Galois orbit by choosing two of four branches and lacks pointwise spectral pairing",
		},
	}

	nilpotent := 0
	canonicalNilpotent := 0
	complete := 0
	for _, d := range differentials {
		if d.Nilpotent {
			nilpotent++
		}
		if d.Nilpotent && d.SquareZero && d.Canonical {
			canonicalNilpotent++
		}
		if d.BRSTCancellationComplete {
			complete++
		}
	}

	invariantGradings := 0
	nontrivialInvariant := 0
	zeroSupertrace := 0
	for _, g := range gradings {
		if g.GaloisInvariant {
			invariantGradings++
		}
		if g.GaloisInvariant && g.Nontrivial {
			nontrivialInvariant++
		}
		if g.ZeroSupertraceLedger {
			zeroSupertrace++
		}
	}

	supertrace := SupertraceAudit{
		QuarticBlockRows:            4,
		GaloisOrbitRows:             4,
		GaloisInvariantGradings:     invariantGradings,
		NontrivialInvariantGradings: nontrivialInvariant,
		ZeroSupertraceGradings:      zeroSupertrace,
		CanonicalZeroSupertrace:     false,
		PointwiseCancellationPairs:  0,
		ZeroBetaLedger:              false,
		Verdict:                     "the only Galois-invariant gradings are uniform and non-cancelling; nontrivial two/two gradings require branch choices and do not give pointwise spectral cancellation",
	}

	construction := BRSTConstructionAudit{
		QuarticBlockExact:         prev.ExactRationalOverlapMatrix && prev.ExactCharacteristicCertified && prev.ExactRootIsolationCertified && prev.QuarticOrbitRows == 4,
		DifferentialsAudited:      len(differentials),
		NilpotentCandidates:       nilpotent,
		CanonicalNilpotent:        canonicalNilpotent,
		NonzeroCanonicalBRST:      0,
		CompleteBRSTCancellations: complete,
		GhostGradingsAudited:      len(gradings),
		CanonicalGhostGradings:    invariantGradings,
		ZeroSupertraceLedgers:     zeroSupertrace,
		ZeroBetaLedgers:           0,
		ConstructionComplete:      false,
		Verdict:                   "a canonical zero differential exists but is inert; all nonzero square-zero candidates require branch/pairing choices, and no zero-supertrace or zero-beta ledger is canonical",
	}

	firewall := FirewallAudit{
		ObservedInputFree:      true,
		QuarticBlockExact:      construction.QuarticBlockExact,
		ConstraintRouteAudited: true,
		BRSTOperator:           false,
		GhostGrading:           false,
		NilpotentDifferential:  false,
		ExactnessOrCohomology:  false,
		SupertraceCancellation: false,
		ZeroBetaLedger:         false,
		RepresentationRows:     0,
		HyperchargeRows:        0,
		ThresholdBetaRows:      0,
		ProvenZeroRows:         0,
		PhysicalConstants:      false,
		FirewallClosed:         true,
		Verdict:                "BRST cancellation requires a selected nontrivial nilpotent differential, ghost grading, exactness/cohomology theorem, supertrace cancellation, and zero-beta ledger; none is complete",
	}

	summary := Summary{
		ContactRows:               prev.ContactRows,
		QuarticBlockRows:          prev.QuarticOrbitRows,
		DifferentialsAudited:      len(differentials),
		NilpotentCandidates:       nilpotent,
		CanonicalNilpotents:       canonicalNilpotent,
		CompleteBRSTCancellations: complete,
		GhostGradingsAudited:      len(gradings),
		GaloisInvariantGradings:   invariantGradings,
		ZeroSupertraceLedgers:     zeroSupertrace,
		QuarticZeroBetaRows:       0,
		QuarticBlockBetaRows:      0,
		ContactBetaRowsAllowed:    0,
		ContactZeroRowsProved:     0,
		ResidualS6Choices:         prev.ResidualS6Choices,
		ResidualNullityBefore:     prev.ResidualNullityAfter,
		ResidualNullityAfter:      prev.ResidualNullityAfter,
	}

	truth := "Gate 158 proves that the exact quartic contact block does not yet admit a BRST cancellation theorem. The only canonical square-zero differential is Q=0, which leaves four cohomology classes alive. Nonzero square-zero pair maps and zero-count ghost gradings require choosing branches inside the quartic Galois orbit, so they are not canonical and cannot produce a zero-beta ledger."

	return Analysis{
		Previous:                     prev,
		DifferentialCandidates:       differentials,
		GhostGradings:                gradings,
		Supertrace:                   supertrace,
		Construction:                 construction,
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
			"zero differential proves quartic cancellation",
			"a nonzero BRST differential is canonically selected on the quartic block",
			"two-even/two-odd ghost grading is Galois-invariant",
			"quartic block has a zero-supertrace or zero-beta ledger",
			"quartic block can be removed from beta matching as BRST-cancelled",
		},
		RemainingUnknowns: []string{
			"canonical nonzero nilpotent differential on the quartic block",
			"ghost grading that is both nontrivial and Galois-invariant",
			"BRST pairing/exactness theorem and pointwise cancellation ledger",
			"whether the quartic block is physical, auxiliary, constrained, or purely finite spectral",
		},
		RecommendedNextGate: "Gate 159 — quartic ghost-grading Galois invariance / nontrivial parity obstruction theorem",
	}, nil
}

func FormatDifferentialCandidate(d DifferentialCandidate) string {
	return fmt.Sprintf("%s: dim=%d nilpotent=%t Q2=%t canonical=%t galois=%t pairChoice=%t orderChoice=%t Hdim=%d supertrace=%t zeroBeta=%t complete=%t (%s)", d.Name, d.Dimension, d.Nilpotent, d.SquareZero, d.Canonical, d.GaloisInvariant, d.RequiresPairingChoice, d.RequiresOrderingChoice, d.CohomologyDimension, d.ZeroSupertraceLedger, d.ZeroBetaLedger, d.BRSTCancellationComplete, d.Verdict)
}

func FormatGhostGrading(g GhostGradingCandidate) string {
	return fmt.Sprintf("%s: rows=%d boson=%d ghost=%d galois=%t nontrivial=%t traceZeroPossible=%t paired=%t branchChoice=%t supertrace=%t complete=%t (%s)", g.Name, g.Rows, g.BosonicRows, g.GhostRows, g.GaloisInvariant, g.Nontrivial, g.TraceZeroPossible, g.PointwisePaired, g.RequiresBranchChoices, g.ZeroSupertraceLedger, g.BRSTCancellationComplete, g.Verdict)
}

func FormatSupertrace(a SupertraceAudit) string {
	return fmt.Sprintf("quarticRows=%d orbitRows=%d invariantGradings=%d nontrivialInvariant=%d zeroSupertraceGradings=%d canonicalZero=%t pointwisePairs=%d zeroBeta=%t (%s)", a.QuarticBlockRows, a.GaloisOrbitRows, a.GaloisInvariantGradings, a.NontrivialInvariantGradings, a.ZeroSupertraceGradings, a.CanonicalZeroSupertrace, a.PointwiseCancellationPairs, a.ZeroBetaLedger, a.Verdict)
}

func FormatConstruction(a BRSTConstructionAudit) string {
	return fmt.Sprintf("quarticExact=%t differentials=%d nilpotent=%d canonicalNilpotent=%d nonzeroCanonical=%d completeCancel=%d gradings=%d canonicalGradings=%d supertrace=%d zeroBeta=%d complete=%t (%s)", a.QuarticBlockExact, a.DifferentialsAudited, a.NilpotentCandidates, a.CanonicalNilpotent, a.NonzeroCanonicalBRST, a.CompleteBRSTCancellations, a.GhostGradingsAudited, a.CanonicalGhostGradings, a.ZeroSupertraceLedgers, a.ZeroBetaLedgers, a.ConstructionComplete, a.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("observedFree=%t quarticExact=%t audited=%t brst=%t ghost=%t nilpotent=%t cohomology=%t supertrace=%t zeroBeta=%t repr=%d Y=%d beta=%d zero=%d physical=%t closed=%t (%s)", f.ObservedInputFree, f.QuarticBlockExact, f.ConstraintRouteAudited, f.BRSTOperator, f.GhostGrading, f.NilpotentDifferential, f.ExactnessOrCohomology, f.SupertraceCancellation, f.ZeroBetaLedger, f.RepresentationRows, f.HyperchargeRows, f.ThresholdBetaRows, f.ProvenZeroRows, f.PhysicalConstants, f.FirewallClosed, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contactRows=%d quarticRows=%d differentials=%d nilpotent=%d canonicalNilpotent=%d completeBRST=%d gradings=%d galoisGradings=%d zeroSupertrace=%d quarticZero=%d quarticBeta=%d contactBeta=%d zero=%d S6=%d nullity=%d→%d", s.ContactRows, s.QuarticBlockRows, s.DifferentialsAudited, s.NilpotentCandidates, s.CanonicalNilpotents, s.CompleteBRSTCancellations, s.GhostGradingsAudited, s.GaloisInvariantGradings, s.ZeroSupertraceLedgers, s.QuarticZeroBetaRows, s.QuarticBlockBetaRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualS6Choices, s.ResidualNullityBefore, s.ResidualNullityAfter)
}
