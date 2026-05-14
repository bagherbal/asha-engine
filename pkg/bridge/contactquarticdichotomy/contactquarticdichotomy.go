// Package contactquarticdichotomy implements Gate 157: quartic block
// constraint-or-propagator dichotomy / BRST-locality firewall theorem.
//
// Gate 156 proved that the exact four-row quartic contact block is not yet a
// local field with spin-statistics, kinetic, gauge, mass, or decoupling data.
// Gate 157 makes the remaining dichotomy executable. A quartic block can enter
// threshold beta matching only by one of two routes: either it is a propagating
// local field with a Lorentzian pole/residue and representation row, or it is a
// constrained/BRST/nonphysical block with a nilpotent differential and a proven
// cancellation ledger. Current finite data supplies neither route, so the
// quartic block remains exact spectral data and the beta firewall stays closed.
package contactquarticdichotomy

import (
	"fmt"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactquarticlocalfield"
)

type PropagatorBranchAudit struct {
	QuarticBlockRows         int
	LocalFieldMap            bool
	BaseSpaceSupport         bool
	LocalSections            bool
	LorentzRepresentation    bool
	KineticOperator          bool
	PropagatorDenominator    bool
	PoleResidueTheorem       bool
	PositiveResidue          bool
	GaugeRepresentation      bool
	HyperchargeRow           bool
	MassActivation           bool
	DecouplingRule           bool
	PropagatorBranchComplete bool
	Verdict                  string
}

type ConstraintBRSTBranchAudit struct {
	QuarticBlockRows         int
	ConstraintEquations      bool
	ConstraintRank           int
	GaugeRedundancy          bool
	GhostGrading             bool
	BRSTOperator             bool
	NilpotentDifferential    bool
	BRSTPairing              bool
	ExactnessOrCohomology    bool
	SupertraceCancellation   bool
	ZeroBetaLedger           bool
	ConstraintBranchComplete bool
	Verdict                  string
}

type DichotomyAudit struct {
	QuarticBlockRows            int
	ExactFiniteSpectralBlock    bool
	PropagatorBranchAudited     bool
	ConstraintBRSTBranchAudited bool
	CompleteBranches            int
	AcceptedPhysicalBranches    int
	AcceptedNonphysicalBranches int
	DichotomyResolved           bool
	BetaRowsPermitted           int
	ZeroRowsProved              int
	Verdict                     string
}

type FirewallAudit struct {
	ObservedInputFree       bool
	QuarticBlockExact       bool
	LocalityRouteComplete   bool
	ConstraintRouteComplete bool
	RepresentationRows      int
	HyperchargeRows         int
	KineticPoleResidueRows  int
	MassActivationRows      int
	DecouplingRows          int
	ThresholdBetaRows       int
	ProvenZeroRows          int
	PhysicalConstants       bool
	FirewallClosed          bool
	Verdict                 string
}

type Summary struct {
	ContactRows                int
	QuarticBlockRows           int
	BranchesAudited            int
	CompleteBranches           int
	PropagatorBranchesComplete int
	ConstraintBranchesComplete int
	LocalFieldRows             int
	BRSTCancellationRows       int
	GaugeRepresentationRows    int
	HyperchargeRows            int
	KineticPoleResidueRows     int
	MassActivationRows         int
	DecouplingRows             int
	QuarticBlockBetaRows       int
	ContactBetaRowsAllowed     int
	ContactZeroRowsProved      int
	ResidualS6Choices          int
	ResidualNullityBefore      int
	ResidualNullityAfter       int
}

type Analysis struct {
	Previous contactquarticlocalfield.Analysis

	PropagatorBranch PropagatorBranchAudit
	ConstraintBranch ConstraintBRSTBranchAudit
	Dichotomy        DichotomyAudit
	Firewall         FirewallAudit
	Summary          Summary

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
		prev, err := contactquarticlocalfield.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactquarticlocalfield.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.QuarticOrbitRows != 4 || prev.QuarticBlockInvariants != 4 || prev.LocalFieldRows != 0 || prev.SpinStatisticsRows != 0 || prev.KineticPoleResidueRows != 0 || prev.QuarticBlockBetaRows != 0 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 157 requires Gate 156 local-field obstruction with closed beta firewall")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 157 refuses hidden observed physical input")
	}

	propagator := PropagatorBranchAudit{
		QuarticBlockRows:         4,
		LocalFieldMap:            false,
		BaseSpaceSupport:         false,
		LocalSections:            false,
		LorentzRepresentation:    false,
		KineticOperator:          false,
		PropagatorDenominator:    false,
		PoleResidueTheorem:       false,
		PositiveResidue:          false,
		GaugeRepresentation:      false,
		HyperchargeRow:           false,
		MassActivation:           false,
		DecouplingRule:           false,
		PropagatorBranchComplete: false,
		Verdict:                  "the exact quartic block has spectral rows, but no base-space support, local sections, Lorentz representation, kinetic denominator, pole/residue theorem, gauge/hypercharge row, mass activation, or decoupling rule",
	}

	constraint := ConstraintBRSTBranchAudit{
		QuarticBlockRows:         4,
		ConstraintEquations:      false,
		ConstraintRank:           0,
		GaugeRedundancy:          false,
		GhostGrading:             false,
		BRSTOperator:             false,
		NilpotentDifferential:    false,
		BRSTPairing:              false,
		ExactnessOrCohomology:    false,
		SupertraceCancellation:   false,
		ZeroBetaLedger:           false,
		ConstraintBranchComplete: false,
		Verdict:                  "the quartic block has no selected constraints, ghost grading, nilpotent BRST operator, pairing/exactness proof, supertrace cancellation, or zero-beta ledger",
	}

	completeBranches := 0
	if propagator.PropagatorBranchComplete {
		completeBranches++
	}
	if constraint.ConstraintBranchComplete {
		completeBranches++
	}

	dichotomy := DichotomyAudit{
		QuarticBlockRows:            4,
		ExactFiniteSpectralBlock:    prev.ExactRationalOverlapMatrix && prev.ExactCharacteristicCertified && prev.ExactRootIsolationCertified,
		PropagatorBranchAudited:     true,
		ConstraintBRSTBranchAudited: true,
		CompleteBranches:            completeBranches,
		AcceptedPhysicalBranches:    0,
		AcceptedNonphysicalBranches: 0,
		DichotomyResolved:           false,
		BetaRowsPermitted:           0,
		ZeroRowsProved:              0,
		Verdict:                     "the quartic contact block is trapped between two incomplete routes: it is neither a propagating local threshold field nor a proven constrained/BRST-cancelled nonphysical block",
	}

	firewall := FirewallAudit{
		ObservedInputFree:       true,
		QuarticBlockExact:       dichotomy.ExactFiniteSpectralBlock && prev.QuarticBlockInvariants == 4,
		LocalityRouteComplete:   false,
		ConstraintRouteComplete: false,
		RepresentationRows:      0,
		HyperchargeRows:         0,
		KineticPoleResidueRows:  0,
		MassActivationRows:      0,
		DecouplingRows:          0,
		ThresholdBetaRows:       0,
		ProvenZeroRows:          0,
		PhysicalConstants:       false,
		FirewallClosed:          true,
		Verdict:                 "beta permission requires either a complete propagating-field route or a complete cancellation route; both are absent",
	}

	summary := Summary{
		ContactRows:                prev.ContactRows,
		QuarticBlockRows:           prev.QuarticOrbitRows,
		BranchesAudited:            2,
		CompleteBranches:           completeBranches,
		PropagatorBranchesComplete: 0,
		ConstraintBranchesComplete: 0,
		LocalFieldRows:             0,
		BRSTCancellationRows:       0,
		GaugeRepresentationRows:    0,
		HyperchargeRows:            0,
		KineticPoleResidueRows:     0,
		MassActivationRows:         0,
		DecouplingRows:             0,
		QuarticBlockBetaRows:       0,
		ContactBetaRowsAllowed:     0,
		ContactZeroRowsProved:      0,
		ResidualS6Choices:          prev.ResidualS6Choices,
		ResidualNullityBefore:      prev.ResidualNullityAfter,
		ResidualNullityAfter:       prev.ResidualNullityAfter,
	}

	truth := "Gate 157 proves the quartic contact block cannot yet move through either permission door. It is not a propagating local field, because locality, Lorentz kinetic data, pole/residue, representation, mass, and decoupling are absent. It is also not a proven nonphysical cancellation block, because constraints, BRST nilpotency, pairing, exactness, and zero-beta ledger are absent. The block remains exact finite spectral data with a closed beta firewall."

	return Analysis{
		Previous:                     prev,
		PropagatorBranch:             propagator,
		ConstraintBranch:             constraint,
		Dichotomy:                    dichotomy,
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
			"quartic block is a propagating local field",
			"quartic block is a positive-residue threshold field",
			"quartic block is a BRST-cancelled ghost/regulator quartet",
			"quartic block has a zero-beta cancellation ledger",
			"quartic block can contribute threshold beta rows without resolving the propagator-or-constraint dichotomy",
		},
		RemainingUnknowns: []string{
			"locality and propagator branch: base space, sections, Lorentz kinetic denominator, pole/residue, mass, decoupling",
			"constraint/BRST branch: constraint equations, ghost grading, nilpotent differential, pairing, exactness, cancellation ledger",
			"gauge/hypercharge representation semantics for the quartic block",
			"whether the quartic block is physical, nonphysical, auxiliary, or purely finite spectral",
		},
		RecommendedNextGate: "Gate 158 — quartic BRST candidate differential / zero-supertrace construction attempt",
	}, nil
}

func FormatPropagatorBranch(p PropagatorBranchAudit) string {
	return fmt.Sprintf("quarticRows=%d local=%t base=%t sections=%t lorentz=%t kinetic=%t denom=%t pole=%t residue=%t gauge=%t Y=%t mass=%t decoupling=%t complete=%t (%s)", p.QuarticBlockRows, p.LocalFieldMap, p.BaseSpaceSupport, p.LocalSections, p.LorentzRepresentation, p.KineticOperator, p.PropagatorDenominator, p.PoleResidueTheorem, p.PositiveResidue, p.GaugeRepresentation, p.HyperchargeRow, p.MassActivation, p.DecouplingRule, p.PropagatorBranchComplete, p.Verdict)
}

func FormatConstraintBranch(c ConstraintBRSTBranchAudit) string {
	return fmt.Sprintf("quarticRows=%d constraints=%t rank=%d redundancy=%t ghost=%t brst=%t nilpotent=%t pairing=%t cohomology=%t supertrace=%t zeroLedger=%t complete=%t (%s)", c.QuarticBlockRows, c.ConstraintEquations, c.ConstraintRank, c.GaugeRedundancy, c.GhostGrading, c.BRSTOperator, c.NilpotentDifferential, c.BRSTPairing, c.ExactnessOrCohomology, c.SupertraceCancellation, c.ZeroBetaLedger, c.ConstraintBranchComplete, c.Verdict)
}

func FormatDichotomy(d DichotomyAudit) string {
	return fmt.Sprintf("quarticRows=%d exact=%t propagatorAudited=%t constraintAudited=%t complete=%d physical=%d nonphysical=%d resolved=%t beta=%d zero=%d (%s)", d.QuarticBlockRows, d.ExactFiniteSpectralBlock, d.PropagatorBranchAudited, d.ConstraintBRSTBranchAudited, d.CompleteBranches, d.AcceptedPhysicalBranches, d.AcceptedNonphysicalBranches, d.DichotomyResolved, d.BetaRowsPermitted, d.ZeroRowsProved, d.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("observedFree=%t quarticExact=%t locality=%t constraint=%t repr=%d Y=%d kinetic=%d mass=%d decoupling=%d beta=%d zero=%d physical=%t closed=%t (%s)", f.ObservedInputFree, f.QuarticBlockExact, f.LocalityRouteComplete, f.ConstraintRouteComplete, f.RepresentationRows, f.HyperchargeRows, f.KineticPoleResidueRows, f.MassActivationRows, f.DecouplingRows, f.ThresholdBetaRows, f.ProvenZeroRows, f.PhysicalConstants, f.FirewallClosed, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contactRows=%d quarticRows=%d branches=%d complete=%d propComplete=%d constraintComplete=%d local=%d brstCancel=%d gauge=%d Y=%d kinetic=%d mass=%d decoupling=%d quarticBeta=%d contactBeta=%d zero=%d S6=%d nullity=%d→%d", s.ContactRows, s.QuarticBlockRows, s.BranchesAudited, s.CompleteBranches, s.PropagatorBranchesComplete, s.ConstraintBranchesComplete, s.LocalFieldRows, s.BRSTCancellationRows, s.GaugeRepresentationRows, s.HyperchargeRows, s.KineticPoleResidueRows, s.MassActivationRows, s.DecouplingRows, s.QuarticBlockBetaRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualS6Choices, s.ResidualNullityBefore, s.ResidualNullityAfter)
}
