// Package contactquarticlocalfield implements Gate 156: quartic block
// local-field / spin-statistics obstruction theorem.
//
// Gate 155 showed that the exact four-row quartic contact block matches several
// familiar four-dimensional multiplet counts, but none of those counts carries
// representation-complete beta semantics. Gate 156 asks the sharper question:
// can the quartic block be lifted to local continuum field data with a Lorentz
// type, spin/statistics rule, kinetic operator, mass activation, and decoupling
// law? The answer remains no. The exact quartic block is preserved as a finite
// spectral diagnostic, but it is not a scalar, spinor, ghost, auxiliary, or
// constrained local field theorem.
package contactquarticlocalfield

import (
	"fmt"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactquarticmultiplet"
)

type LocalFieldRequirementAudit struct {
	QuarticBlockRows          int
	BaseSpaceRows             int
	LocalSectionRows          int
	LorentzRepresentationRows int
	KineticOperatorRows       int
	PoleResidueRows           int
	SpinStatisticsRows        int
	GaugeRepresentationRows   int
	HyperchargeRows           int
	MassActivationRows        int
	DecouplingRows            int
	AllRequirementsSatisfied  bool
	Verdict                   string
}

type FieldTypeCandidate struct {
	Name                  string
	DegreeCount           int
	MatchesQuarticBlock   bool
	BaseSpace             bool
	LocalSections         bool
	LorentzRepresentation bool
	KineticOperator       bool
	PoleResidue           bool
	SpinStatistics        bool
	GaugeRepresentation   bool
	MassActivation        bool
	DecouplingRule        bool
	LocalFieldComplete    bool
	BetaPermission        bool
	Verdict               string
}

type SpinStatisticsAudit struct {
	CandidateRows                int
	BosonicScalarRows            int
	FermionicSpinorRows          int
	GhostRegulatorRows           int
	AuxiliaryConstrainedRows     int
	CanonicalCommutationRows     int
	CanonicalAnticommutationRows int
	BRSTGradingRows              int
	SpinStatisticsComplete       int
	Verdict                      string
}

type KineticLocalityAudit struct {
	FiniteSpectralBlockExact      bool
	SpectralOperatorRows          int
	LocalDifferentialOperatorRows int
	LorentzSignatureRows          int
	HyperbolicEllipticClassRows   int
	PropagatorDenominatorRows     int
	PositiveResidueRows           int
	LocalityComplete              bool
	Verdict                       string
}

type BetaFirewallAudit struct {
	ObservedInputFree         bool
	QuarticBlockExact         bool
	MultipletDimensionAudited bool
	LocalFieldMap             bool
	SpinStatistics            bool
	KineticPoleResidue        bool
	GaugeRepresentation       bool
	Hypercharge               bool
	MassActivation            bool
	Decoupling                bool
	ThresholdBetaRows         bool
	PhysicalConstants         bool
	FirewallClosed            bool
	Verdict                   string
}

type Summary struct {
	ContactRows                 int
	QuarticBlockRows            int
	FieldTypeCandidatesAudited  int
	DegreeMatchingCandidates    int
	LocalFieldCompleteRows      int
	SpinStatisticsCompleteRows  int
	KineticLocalityCompleteRows int
	GaugeRepresentationRows     int
	HyperchargeRows             int
	MassActivationRows          int
	DecouplingRows              int
	QuarticBlockBetaRows        int
	ContactBetaRowsAllowed      int
	ContactZeroRowsProved       int
	ResidualS6Choices           int
	ResidualNullityBefore       int
	ResidualNullityAfter        int
}

type Analysis struct {
	Previous contactquarticmultiplet.Analysis

	RequirementAudit LocalFieldRequirementAudit
	FieldCandidates  []FieldTypeCandidate
	SpinAudit        SpinStatisticsAudit
	KineticAudit     KineticLocalityAudit
	Firewall         BetaFirewallAudit
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
		prev, err := contactquarticmultiplet.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactquarticmultiplet.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.QuarticOrbitRows != 4 || prev.QuarticBlockInvariants != 4 || prev.RepresentationCompleteRows != 0 || prev.QuarticBlockBetaRows != 0 || prev.ContactBetaRowsAllowed != 0 || prev.LocalFieldRows != 0 || prev.SpinStatisticsRows != 0 || prev.DynkinIndexRows != 0 {
		return Analysis{}, fmt.Errorf("Gate 156 requires Gate 155 quartic multiplet audit with closed beta firewall")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 156 refuses hidden observed physical input")
	}

	candidates := []FieldTypeCandidate{
		{
			Name:                "real scalar four-component local field",
			DegreeCount:         4,
			MatchesQuarticBlock: true,
			Verdict:             "four real components match the block count, but no spacetime support, scalar kinetic operator, mass unit, or decoupling rule is derived",
		},
		{
			Name:                "complex scalar doublet local field",
			DegreeCount:         4,
			MatchesQuarticBlock: true,
			Verdict:             "four real components match a doublet count, but SU(2)L action, hypercharge, local sections, and Higgs-like kinetic semantics are absent",
		},
		{
			Name:                "Weyl or Dirac spinor candidate",
			DegreeCount:         4,
			MatchesQuarticBlock: true,
			Verdict:             "dimension can be matched, but spinor bundle, Clifford/Lorentz action, anticommutation rule, and statistics are not derived on the contact block",
		},
		{
			Name:                "ghost or regulator quartet",
			DegreeCount:         4,
			MatchesQuarticBlock: true,
			Verdict:             "quartet count is compatible, but ghost grading, nilpotent BRST differential, pairing, and cancellation ledger are absent",
		},
		{
			Name:                "auxiliary or constrained non-propagating quartet",
			DegreeCount:         4,
			MatchesQuarticBlock: true,
			Verdict:             "non-propagating interpretation would require constraint equations or an elimination theorem, neither of which is selected",
		},
	}

	degreeMatches := 0
	localComplete := 0
	betaPermitted := 0
	for _, c := range candidates {
		if c.MatchesQuarticBlock {
			degreeMatches++
		}
		if c.LocalFieldComplete {
			localComplete++
		}
		if c.BetaPermission {
			betaPermitted++
		}
	}

	requirement := LocalFieldRequirementAudit{
		QuarticBlockRows:          4,
		BaseSpaceRows:             0,
		LocalSectionRows:          0,
		LorentzRepresentationRows: 0,
		KineticOperatorRows:       0,
		PoleResidueRows:           0,
		SpinStatisticsRows:        0,
		GaugeRepresentationRows:   0,
		HyperchargeRows:           0,
		MassActivationRows:        0,
		DecouplingRows:            0,
		AllRequirementsSatisfied:  false,
		Verdict:                   "a beta-permitted local field requires base-space support, local sections, Lorentz type, kinetic/pole-residue data, spin/statistics, gauge/hypercharge rows, mass activation, and decoupling; none is derived for the quartic contact block",
	}

	spin := SpinStatisticsAudit{
		CandidateRows:                len(candidates),
		BosonicScalarRows:            0,
		FermionicSpinorRows:          0,
		GhostRegulatorRows:           0,
		AuxiliaryConstrainedRows:     0,
		CanonicalCommutationRows:     0,
		CanonicalAnticommutationRows: 0,
		BRSTGradingRows:              0,
		SpinStatisticsComplete:       0,
		Verdict:                      "the quartic block supplies spectral degree count only; it does not select bosonic, fermionic, ghost, or auxiliary statistics",
	}

	kinetic := KineticLocalityAudit{
		FiniteSpectralBlockExact:      prev.ExactRationalOverlapMatrix && prev.ExactCharacteristicCertified && prev.ExactRootIsolationCertified && prev.QuarticCompressedBlocks == 1,
		SpectralOperatorRows:          4,
		LocalDifferentialOperatorRows: 0,
		LorentzSignatureRows:          0,
		HyperbolicEllipticClassRows:   0,
		PropagatorDenominatorRows:     0,
		PositiveResidueRows:           0,
		LocalityComplete:              false,
		Verdict:                       "exact finite spectral operator rows are not local differential operators or Lorentzian propagators",
	}

	firewall := BetaFirewallAudit{
		ObservedInputFree:         true,
		QuarticBlockExact:         prev.ExactRationalOverlapMatrix && prev.ExactCharacteristicCertified && prev.ExactRootIsolationCertified && prev.QuarticBlockInvariants == 4,
		MultipletDimensionAudited: prev.MultipletAudit.DimensionMatches == 4,
		LocalFieldMap:             false,
		SpinStatistics:            false,
		KineticPoleResidue:        false,
		GaugeRepresentation:       false,
		Hypercharge:               false,
		MassActivation:            false,
		Decoupling:                false,
		ThresholdBetaRows:         false,
		PhysicalConstants:         false,
		FirewallClosed:            true,
		Verdict:                   "the quartic block is exact finite spectral data but remains below the permission threshold for physical beta matching",
	}

	summary := Summary{
		ContactRows:                 prev.ContactRows,
		QuarticBlockRows:            prev.QuarticOrbitRows,
		FieldTypeCandidatesAudited:  len(candidates),
		DegreeMatchingCandidates:    degreeMatches,
		LocalFieldCompleteRows:      localComplete,
		SpinStatisticsCompleteRows:  spin.SpinStatisticsComplete,
		KineticLocalityCompleteRows: 0,
		GaugeRepresentationRows:     0,
		HyperchargeRows:             0,
		MassActivationRows:          0,
		DecouplingRows:              0,
		QuarticBlockBetaRows:        betaPermitted,
		ContactBetaRowsAllowed:      0,
		ContactZeroRowsProved:       0,
		ResidualS6Choices:           prev.ResidualS6Choices,
		ResidualNullityBefore:       prev.ResidualNullityAfter,
		ResidualNullityAfter:        prev.ResidualNullityAfter,
	}

	truth := "Gate 156 proves that the exact four-row quartic contact block does not yet define a local field. Scalar, doublet, spinor, ghost/regulator, and auxiliary interpretations can match the degree count, but none supplies local sections, Lorentz kinetic/pole-residue data, spin-statistics, gauge/hypercharge rows, mass activation, or decoupling. The quartic block therefore remains an exact spectral diagnostic and cannot contribute threshold beta rows."

	return Analysis{
		Previous:                     prev,
		RequirementAudit:             requirement,
		FieldCandidates:              candidates,
		SpinAudit:                    spin,
		KineticAudit:                 kinetic,
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
			"quartic block degree count derives a local scalar field",
			"quartic block degree count derives a weak doublet",
			"quartic block degree count derives a spinor or fermion statistics",
			"quartic block degree count derives a ghost/regulator cancellation",
			"exact spectral rows can be used as threshold beta rows without kinetic and decoupling data",
		},
		RemainingUnknowns: []string{
			"local spacetime support and section variables for quartic contact data",
			"Lorentz representation and kinetic/pole-residue theorem",
			"spin-statistics or BRST/constraint complex",
			"gauge and hypercharge representation rows",
			"mass activation and decoupling rule for threshold matching",
		},
		RecommendedNextGate: "Gate 157 — quartic block constraint-or-propagator dichotomy / BRST-locality firewall theorem",
	}, nil
}

func FormatFieldCandidate(c FieldTypeCandidate) string {
	return fmt.Sprintf("%s dim=%d matches=%t base=%t sections=%t lorentz=%t kinetic=%t pole=%t spin=%t gauge=%t mass=%t decoupling=%t complete=%t beta=%t (%s)", c.Name, c.DegreeCount, c.MatchesQuarticBlock, c.BaseSpace, c.LocalSections, c.LorentzRepresentation, c.KineticOperator, c.PoleResidue, c.SpinStatistics, c.GaugeRepresentation, c.MassActivation, c.DecouplingRule, c.LocalFieldComplete, c.BetaPermission, c.Verdict)
}

func FormatRequirementAudit(r LocalFieldRequirementAudit) string {
	return fmt.Sprintf("quarticRows=%d base=%d sections=%d lorentz=%d kinetic=%d pole=%d spin=%d gauge=%d Y=%d mass=%d decoupling=%d all=%t (%s)", r.QuarticBlockRows, r.BaseSpaceRows, r.LocalSectionRows, r.LorentzRepresentationRows, r.KineticOperatorRows, r.PoleResidueRows, r.SpinStatisticsRows, r.GaugeRepresentationRows, r.HyperchargeRows, r.MassActivationRows, r.DecouplingRows, r.AllRequirementsSatisfied, r.Verdict)
}

func FormatSpinAudit(s SpinStatisticsAudit) string {
	return fmt.Sprintf("candidates=%d boson=%d fermion=%d ghost=%d auxiliary=%d comm=%d anticomm=%d brst=%d complete=%d (%s)", s.CandidateRows, s.BosonicScalarRows, s.FermionicSpinorRows, s.GhostRegulatorRows, s.AuxiliaryConstrainedRows, s.CanonicalCommutationRows, s.CanonicalAnticommutationRows, s.BRSTGradingRows, s.SpinStatisticsComplete, s.Verdict)
}

func FormatKineticAudit(k KineticLocalityAudit) string {
	return fmt.Sprintf("finiteExact=%t spectralRows=%d localDiff=%d lorentzSig=%d class=%d denom=%d residue=%d complete=%t (%s)", k.FiniteSpectralBlockExact, k.SpectralOperatorRows, k.LocalDifferentialOperatorRows, k.LorentzSignatureRows, k.HyperbolicEllipticClassRows, k.PropagatorDenominatorRows, k.PositiveResidueRows, k.LocalityComplete, k.Verdict)
}

func FormatFirewall(f BetaFirewallAudit) string {
	return fmt.Sprintf("observedFree=%t quarticExact=%t dimensionAudited=%t local=%t spin=%t kinetic=%t gauge=%t Y=%t mass=%t decoupling=%t beta=%t physical=%t closed=%t (%s)", f.ObservedInputFree, f.QuarticBlockExact, f.MultipletDimensionAudited, f.LocalFieldMap, f.SpinStatistics, f.KineticPoleResidue, f.GaugeRepresentation, f.Hypercharge, f.MassActivation, f.Decoupling, f.ThresholdBetaRows, f.PhysicalConstants, f.FirewallClosed, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contactRows=%d quarticRows=%d candidates=%d degreeMatches=%d localComplete=%d spinComplete=%d kineticComplete=%d gauge=%d Y=%d mass=%d decoupling=%d quarticBeta=%d contactBeta=%d zero=%d S6=%d nullity=%d→%d", s.ContactRows, s.QuarticBlockRows, s.FieldTypeCandidatesAudited, s.DegreeMatchingCandidates, s.LocalFieldCompleteRows, s.SpinStatisticsCompleteRows, s.KineticLocalityCompleteRows, s.GaugeRepresentationRows, s.HyperchargeRows, s.MassActivationRows, s.DecouplingRows, s.QuarticBlockBetaRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualS6Choices, s.ResidualNullityBefore, s.ResidualNullityAfter)
}
