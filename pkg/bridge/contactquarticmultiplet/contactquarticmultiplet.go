// Package contactquarticmultiplet implements Gate 155: quartic block multiplet
// representation / beta-index obstruction theorem.
//
// Gate 154 produced an exact branch-free four-row quartic spectral block with
// rational symmetric invariants. Gate 155 asks whether that block can be
// promoted to a physical multiplet carrying a gauge representation and Dynkin
// index, so that it may contribute a threshold beta row. The answer remains no:
// dimension-four matches exist, but no SU(3)xSU(2)xU(1) representation, spin
// statistics, local field map, mass activation, or decoupling theorem is
// derived from the finite contact data.
package contactquarticmultiplet

import (
	"fmt"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactquarticcompression"
)

type MultipletCandidate struct {
	Name                string
	Dimension           int
	MatchesQuarticRows  bool
	GaugeRepresentation bool
	SpinStatistics      bool
	LocalFieldMap       bool
	MassActivation      bool
	DecouplingRule      bool
	DynkinIndex         bool
	BetaPermission      bool
	Verdict             string
}

type MultipletAudit struct {
	QuarticBlockRows       int
	Candidates             []MultipletCandidate
	DimensionMatches       int
	RepresentationComplete int
	BetaPermitted          int
	Verdict                string
}

type DynkinIndexRequirementAudit struct {
	BlockRows                int
	GroupActionRows          int
	RepresentationRows       int
	TraceNormalizationRows   int
	SpinStatisticsRows       int
	MultiplicityRows         int
	MassThresholdRows        int
	DecouplingRows           int
	DynkinIndexRows          int
	BetaIndexRows            int
	AllRequirementsSatisfied bool
	Verdict                  string
}

type BlockInvariantUseAudit struct {
	SymmetricInvariants     int
	DegreeInvariant         bool
	TraceInvariant          bool
	DeterminantInvariant    bool
	CharacteristicDataExact bool
	RepresentationSemantics bool
	DynkinIndexSemantics    bool
	ChargeSemantics         bool
	Verdict                 string
}

type BetaFirewallAudit struct {
	ObservedInputFree          bool
	QuarticBlockExact          bool
	BlockCompressionBranchFree bool
	MultipletRepresentation    bool
	DynkinIndex                bool
	LocalFieldMap              bool
	MassActivation             bool
	Decoupling                 bool
	ThresholdBetaRows          bool
	PhysicalConstants          bool
	FirewallClosed             bool
	Verdict                    string
}

type Summary struct {
	ContactRows                 int
	QuarticBlockRows            int
	QuarticBlockInvariants      int
	MultipletCandidatesAudited  int
	DimensionMatchingCandidates int
	RepresentationCompleteRows  int
	QuarticBlockBetaRows        int
	ContactBetaRowsAllowed      int
	ContactZeroRowsProved       int
	DynkinIndexRows             int
	ResidualS6Choices           int
	ResidualNullityBefore       int
	ResidualNullityAfter        int
}

type Analysis struct {
	Previous contactquarticcompression.Analysis

	MultipletAudit MultipletAudit
	DynkinAudit    DynkinIndexRequirementAudit
	InvariantUse   BlockInvariantUseAudit
	Firewall       BetaFirewallAudit
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
	SpinStatisticsRows           int
	LocalFieldRows               int
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
		prev, err := contactquarticcompression.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactquarticcompression.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.QuarticCompressedBlocks != 1 || prev.QuarticBlockInvariants != 4 || prev.QuarticOrbitRows != 4 || prev.QuarticBlockBetaRows != 0 || prev.ContactBetaRowsAllowed != 0 || prev.RepresentationCompleteRows != 0 {
		return Analysis{}, fmt.Errorf("Gate 155 requires Gate 154 exact quartic block compression with closed beta firewall")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 155 refuses hidden observed physical input")
	}

	candidates := []MultipletCandidate{
		{
			Name:                "real scalar 4-vector",
			Dimension:           4,
			MatchesQuarticRows:  true,
			GaugeRepresentation: false,
			SpinStatistics:      false,
			LocalFieldMap:       false,
			MassActivation:      false,
			DecouplingRule:      false,
			DynkinIndex:         false,
			BetaPermission:      false,
			Verdict:             "dimension matches the quartic block, but no gauge representation or local scalar field map is derived",
		},
		{
			Name:                "complex scalar SU(2) doublet candidate",
			Dimension:           4,
			MatchesQuarticRows:  true,
			GaugeRepresentation: false,
			SpinStatistics:      false,
			LocalFieldMap:       false,
			MassActivation:      false,
			DecouplingRule:      false,
			DynkinIndex:         false,
			BetaPermission:      false,
			Verdict:             "four real degrees match a Higgs-like doublet count, but SU(2)L action and hypercharge are not derived on the quartic contact block",
		},
		{
			Name:                "four singlet scalar thresholds",
			Dimension:           4,
			MatchesQuarticRows:  true,
			GaugeRepresentation: false,
			SpinStatistics:      false,
			LocalFieldMap:       false,
			MassActivation:      false,
			DecouplingRule:      false,
			DynkinIndex:         false,
			BetaPermission:      false,
			Verdict:             "splitting the quartic orbit into four singlets would require individual branch choices and local field variables",
		},
		{
			Name:                "fermionic Dirac-like block",
			Dimension:           4,
			MatchesQuarticRows:  true,
			GaugeRepresentation: false,
			SpinStatistics:      false,
			LocalFieldMap:       false,
			MassActivation:      false,
			DecouplingRule:      false,
			DynkinIndex:         false,
			BetaPermission:      false,
			Verdict:             "dimension alone cannot supply fermion statistics, Lorentz spinor structure, or gauge index",
		},
	}

	dimensionMatches := 0
	representationComplete := 0
	betaPermitted := 0
	for _, c := range candidates {
		if c.MatchesQuarticRows {
			dimensionMatches++
		}
		if c.GaugeRepresentation && c.SpinStatistics && c.LocalFieldMap && c.MassActivation && c.DecouplingRule && c.DynkinIndex {
			representationComplete++
		}
		if c.BetaPermission {
			betaPermitted++
		}
	}

	multiplet := MultipletAudit{
		QuarticBlockRows:       4,
		Candidates:             candidates,
		DimensionMatches:       dimensionMatches,
		RepresentationComplete: representationComplete,
		BetaPermitted:          betaPermitted,
		Verdict:                "four-dimensional multiplet interpretations exist by dimension count, but none is representation-complete or beta-permitted",
	}

	dynkin := DynkinIndexRequirementAudit{
		BlockRows:                4,
		GroupActionRows:          0,
		RepresentationRows:       0,
		TraceNormalizationRows:   0,
		SpinStatisticsRows:       0,
		MultiplicityRows:         0,
		MassThresholdRows:        0,
		DecouplingRows:           0,
		DynkinIndexRows:          0,
		BetaIndexRows:            0,
		AllRequirementsSatisfied: false,
		Verdict:                  "Dynkin/beta indices require a gauge action, representation, trace normalization, spin/statistics, multiplicity, threshold mass, and decoupling rule; the quartic block supplies none of these",
	}

	invariantUse := BlockInvariantUseAudit{
		SymmetricInvariants:     prev.QuarticBlockInvariants,
		DegreeInvariant:         true,
		TraceInvariant:          true,
		DeterminantInvariant:    true,
		CharacteristicDataExact: prev.ExactCharacteristicCertified,
		RepresentationSemantics: false,
		DynkinIndexSemantics:    false,
		ChargeSemantics:         false,
		Verdict:                 "quartic symmetric invariants are exact spectral invariants, not representation, charge, or beta-index semantics",
	}

	firewall := BetaFirewallAudit{
		ObservedInputFree:          true,
		QuarticBlockExact:          prev.ExactRationalOverlapMatrix && prev.ExactCharacteristicCertified && prev.ExactRootIsolationCertified && prev.QuarticBlockInvariants == 4,
		BlockCompressionBranchFree: prev.QuarticCompressedBlocks == 1,
		MultipletRepresentation:    false,
		DynkinIndex:                false,
		LocalFieldMap:              false,
		MassActivation:             false,
		Decoupling:                 false,
		ThresholdBetaRows:          false,
		PhysicalConstants:          false,
		FirewallClosed:             true,
		Verdict:                    "exact quartic block data remains below the threshold required for a physical multiplet or beta-index row",
	}

	summary := Summary{
		ContactRows:                 prev.ContactRows,
		QuarticBlockRows:            prev.QuarticOrbitRows,
		QuarticBlockInvariants:      prev.QuarticBlockInvariants,
		MultipletCandidatesAudited:  len(candidates),
		DimensionMatchingCandidates: dimensionMatches,
		RepresentationCompleteRows:  0,
		QuarticBlockBetaRows:        0,
		ContactBetaRowsAllowed:      0,
		ContactZeroRowsProved:       0,
		DynkinIndexRows:             0,
		ResidualS6Choices:           prev.ResidualS6Choices,
		ResidualNullityBefore:       prev.ResidualNullityAfter,
		ResidualNullityAfter:        prev.ResidualNullityAfter,
	}

	truth := "Gate 155 proves that the exact four-row quartic contact block can be audited as a possible multiplet only at the level of dimension count. No gauge action, representation, spin/statistics, local field map, mass activation, decoupling rule, or Dynkin index is selected. Therefore the quartic block remains an exact spectral diagnostic and cannot contribute a threshold beta row."

	return Analysis{
		Previous:                     prev,
		MultipletAudit:               multiplet,
		DynkinAudit:                  dynkin,
		InvariantUse:                 invariantUse,
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
			"quartic block dimension four proves a physical Higgs-like doublet",
			"quartic block symmetric invariants determine a Dynkin index",
			"four contact spectral rows can be counted as four physical thresholds",
			"the quartic block contributes threshold beta coefficients without gauge representation and decoupling data",
			"multiplet dimension matching derives physical constants or masses",
		},
		RemainingUnknowns: []string{
			"gauge action and representation carried by the quartic block",
			"spin/statistics and local field map for any quartic multiplet interpretation",
			"mass activation and decoupling rule for threshold matching",
			"Dynkin index and beta row derived from finite contact data",
			"physical-flow selector for u, L, and threshold corrections",
		},
		RecommendedNextGate: "Gate 156 — quartic block local-field/spin-statistics obstruction theorem",
	}, nil
}

func FormatCandidate(c MultipletCandidate) string {
	return fmt.Sprintf("%s dim=%d matches=%t gauge=%t spin=%t local=%t mass=%t decoupling=%t dynkin=%t beta=%t (%s)", c.Name, c.Dimension, c.MatchesQuarticRows, c.GaugeRepresentation, c.SpinStatistics, c.LocalFieldMap, c.MassActivation, c.DecouplingRule, c.DynkinIndex, c.BetaPermission, c.Verdict)
}

func FormatMultipletAudit(m MultipletAudit) string {
	return fmt.Sprintf("quarticRows=%d candidates=%d dimensionMatches=%d reprComplete=%d betaPermitted=%d (%s)", m.QuarticBlockRows, len(m.Candidates), m.DimensionMatches, m.RepresentationComplete, m.BetaPermitted, m.Verdict)
}

func FormatDynkinAudit(d DynkinIndexRequirementAudit) string {
	return fmt.Sprintf("blockRows=%d group=%d repr=%d traceNorm=%d spin=%d multiplicity=%d mass=%d decoupling=%d dynkin=%d betaIndex=%d all=%t (%s)", d.BlockRows, d.GroupActionRows, d.RepresentationRows, d.TraceNormalizationRows, d.SpinStatisticsRows, d.MultiplicityRows, d.MassThresholdRows, d.DecouplingRows, d.DynkinIndexRows, d.BetaIndexRows, d.AllRequirementsSatisfied, d.Verdict)
}

func FormatInvariantUse(i BlockInvariantUseAudit) string {
	return fmt.Sprintf("invariants=%d degree=%t trace=%t det=%t charExact=%t reprSemantics=%t dynkinSemantics=%t chargeSemantics=%t (%s)", i.SymmetricInvariants, i.DegreeInvariant, i.TraceInvariant, i.DeterminantInvariant, i.CharacteristicDataExact, i.RepresentationSemantics, i.DynkinIndexSemantics, i.ChargeSemantics, i.Verdict)
}

func FormatFirewall(f BetaFirewallAudit) string {
	return fmt.Sprintf("observedFree=%t quarticExact=%t branchFree=%t multiplet=%t dynkin=%t local=%t mass=%t decoupling=%t beta=%t physical=%t closed=%t (%s)", f.ObservedInputFree, f.QuarticBlockExact, f.BlockCompressionBranchFree, f.MultipletRepresentation, f.DynkinIndex, f.LocalFieldMap, f.MassActivation, f.Decoupling, f.ThresholdBetaRows, f.PhysicalConstants, f.FirewallClosed, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contactRows=%d quarticRows=%d invariants=%d candidates=%d dimensionMatches=%d reprComplete=%d quarticBeta=%d contactBeta=%d zero=%d dynkin=%d S6=%d nullity=%d→%d", s.ContactRows, s.QuarticBlockRows, s.QuarticBlockInvariants, s.MultipletCandidatesAudited, s.DimensionMatchingCandidates, s.RepresentationCompleteRows, s.QuarticBlockBetaRows, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.DynkinIndexRows, s.ResidualS6Choices, s.ResidualNullityBefore, s.ResidualNullityAfter)
}
