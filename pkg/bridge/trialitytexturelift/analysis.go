// Package trialitytexturelift implements Gate 172: the triality-lifted Yukawa
// texture operator search.
//
// Gates 168-171 turned the scalar-sector mismatch into a precise finite
// Yukawa moment target, but the one-generation kind assignment remained a
// branch choice. Gate 172 moves the question into the triality layer.  The
// Yukawa data after Gate 26 are four fermion-kind texture blocks
//
//	Y_u, Y_d, Y_ν, Y_e ∈ Mat_3,
//
// one 3x3 generation matrix per kind.  This gate audits whether existing
// finite data select these matrices, or at least a canonical texture operator
// tying the Gate-169 contact scalar-shape weights to generation/triality.
//
// The result is deliberately conservative. Exact triality is canonical but
// kind-blind and gives only a 1+2 eigenpattern. The Higgs/contact anisotropy
// spurion can split three generation weights, but it is diagonal, not a
// canonical total Yukawa operator, and produces no mixing. Separable ansätze
// combining the contact kind weights with a diagonal generation spurion remain
// branch choices and keep all kind matrices simultaneously diagonal. Therefore
// no mass or CKM/PMNS theorem is claimed.
package trialitytexturelift

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactkindassignment"
	"github.com/bagherbal/asha-engine/pkg/matter/generationbreak"
	"github.com/bagherbal/asha-engine/pkg/matter/texture"
	"github.com/bagherbal/asha-engine/pkg/matter/trialityyukawa"
)

type TextureCandidateKind string

const (
	CandidateExactTriality       TextureCandidateKind = "exact triality-invariant texture"
	CandidateDiagonalSpurion     TextureCandidateKind = "Higgs/contact diagonal generation spurion"
	CandidateContactKindIdentity TextureCandidateKind = "contact four-kind weights times generation identity"
	CandidateSeparableTexture    TextureCandidateKind = "separable kind-contact times generation-spurion texture"
	CandidateGeneralTexture      TextureCandidateKind = "general four-matrix Yukawa texture"
	CandidateNoncommutingPair    TextureCandidateKind = "missing non-commuting finite texture pair"
)

type TextureOperatorCandidate struct {
	Name                         string
	Kind                         TextureCandidateKind
	Source                       string
	KindAxis                     string
	GenerationAxis               string
	MatrixBlocks                 int
	Parameters                   int
	EigenPattern                 string
	BreaksAllThreeGenerations    bool
	ProducesMixing               bool
	Canonical                    bool
	RequiresBridge               bool
	RequiresBranchChoice         bool
	KindSensitive                bool
	GenerationSensitive          bool
	TiesContactWeightsToKinds    bool
	MatchesScalarShapeCondition  bool
	SelectsYukawaAmplitudes      bool
	SelectsMassEigenvalues       bool
	SelectsCKMPMNS               bool
	NonCommutingWithOtherTexture bool
	Verdict                      string
}

type TrialityLiftAudit struct {
	GenerationCount              int
	FermionKindBlocks            int
	YukawaTextureMatrices        int
	GeneralEntriesPerMatrix      int
	TotalGeneralRealEntries      int
	SymmetricEntriesPerMatrix    int
	TotalSymmetricRealEntries    int
	TrialityInvariantDimPerKind  int
	TotalTrialityInvariantParams int
	FullMixingMaps               int
	DiagonalTrialityChannels     int
	FullMixingChannels           int
	ScalarMomentConstraints      int
	TextureUnderdetermined       bool
	Verdict                      string
}

type OperatorSearchAudit struct {
	CandidateCount                     int
	CanonicalOperatorsFound            int
	CanonicalBreakingOperatorsFound    int
	GenerationSplittingCandidatesFound int
	MixingOperatorsFound               int
	NonCommutingPairsFound             int
	ScalarShapeConditionalCandidates   int
	BranchChoiceCandidates             int
	UniqueTextureSelected              bool
	Verdict                            string
}

type AxisCouplingAudit struct {
	ContactKindAssignmentsSurvive int
	CanonicalKindPartitions       int
	TrialityInvariantPattern      string
	DiagonalGenerationSpurionSeen bool
	DiagonalSpurionCanonical      bool
	KindGenerationCouplingDerived bool
	SeparableAnsatzOnly           bool
	AllKindMatricesAligned        bool
	Verdict                       string
}

type MassTextureAudit struct {
	FourYukawaMatricesRecognized      bool
	MassesAreSingularValues           bool
	MixingNeedsRelativeLeftEigenbasis bool
	AtLeastTwoNoncommutingNeeded      bool
	AtLeastTwoNoncommutingFound       bool
	ScalarShapeIsOneMomentConstraint  bool
	YukawaMatricesDerived             bool
	FermionMassesDerived              bool
	CKMPMNSDerived                    bool
	Verdict                           string
}

type FirewallAudit struct {
	GaugeRatioClosed                 bool
	ScalarShapeTargetAvailable       bool
	ContactKindAssignmentDerived     bool
	TrialityLiftPerformed            bool
	CanonicalTextureOperatorSelected bool
	YukawaAmplitudesDerived          bool
	GenerationHierarchyDerived       bool
	FermionMassesDerived             bool
	CKMPMNSDerived                   bool
	PhysicalConstantsDerived         bool
	ResidualNullityBefore            int
	ResidualNullityAfter             int
	Verdict                          string
}

type Analysis struct {
	ContactKind contactkindassignment.Analysis
	Triality    trialityyukawa.Analysis
	Texture     texture.Analysis
	Generation  generationbreak.Analysis

	LiftAudit      TrialityLiftAudit
	Candidates     []TextureOperatorCandidate
	OperatorAudit  OperatorSearchAudit
	AxisAudit      AxisCouplingAudit
	MassAudit      MassTextureAudit
	Firewall       FirewallAudit
	TruthStatement string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		ck, err := contactkindassignment.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		tr, err := trialityyukawa.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		tx, err := texture.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		gb, err := generationbreak.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(ck, tr, tx, gb)
	})
	return defaultValue, defaultErr
}

func Build(ck contactkindassignment.Analysis, tr trialityyukawa.Analysis, tx texture.Analysis, gb generationbreak.Analysis) (Analysis, error) {
	if len(ck.KindSignatures) != 4 || ck.AssignmentAudit.SurvivingBranchChoices != 6 {
		return Analysis{}, fmt.Errorf("Gate 172 requires Gate 171 unresolved four-kind assignment with six branches")
	}
	if tr.GenerationCount != 3 || !tr.GenerationMixingAllowedByCharges {
		return Analysis{}, fmt.Errorf("Gate 172 requires Gate 26 three-generation triality Yukawa lift")
	}
	if tx.GenerationDimension != 3 || tx.FermionKinds != 4 {
		return Analysis{}, fmt.Errorf("Gate 172 requires Gate 28 four-kind, three-generation texture space")
	}
	if gb.GenerationCarrierDimension != 3 {
		return Analysis{}, fmt.Errorf("Gate 172 requires Gate 29 generation carrier dimension 3")
	}

	lift := buildLiftAudit(tr, tx)
	candidates := buildCandidates(ck, tr, tx, gb)
	op := auditOperators(candidates)
	axis := auditAxes(ck, tx, gb)
	mass := MassTextureAudit{
		FourYukawaMatricesRecognized:      lift.YukawaTextureMatrices == 4,
		MassesAreSingularValues:           true,
		MixingNeedsRelativeLeftEigenbasis: true,
		AtLeastTwoNoncommutingNeeded:      true,
		AtLeastTwoNoncommutingFound:       op.NonCommutingPairsFound >= 1,
		ScalarShapeIsOneMomentConstraint:  lift.ScalarMomentConstraints == 1,
		YukawaMatricesDerived:             op.UniqueTextureSelected,
		FermionMassesDerived:              false,
		CKMPMNSDerived:                    false,
		Verdict:                           "after triality the mass problem is four 3x3 finite Dirac/Yukawa matrices; the scalar shape supplies one moment constraint, not the matrices, phases, singular values, or relative eigenbases",
	}
	fw := FirewallAudit{
		GaugeRatioClosed:                 ck.Firewall.GaugeRatioClosed,
		ScalarShapeTargetAvailable:       ck.Firewall.ScalarShapeTargetAvailable,
		ContactKindAssignmentDerived:     ck.Firewall.ContactKindAssignmentDerived,
		TrialityLiftPerformed:            true,
		CanonicalTextureOperatorSelected: op.UniqueTextureSelected,
		YukawaAmplitudesDerived:          false,
		GenerationHierarchyDerived:       false,
		FermionMassesDerived:             false,
		CKMPMNSDerived:                   false,
		PhysicalConstantsDerived:         false,
		ResidualNullityBefore:            3,
		ResidualNullityAfter:             3,
		Verdict:                          "triality exposes the correct finite matrix arena, but no canonical non-commuting Yukawa texture operator is selected; the mass and mixing firewall remains closed",
	}
	return Analysis{
		ContactKind:    ck,
		Triality:       tr,
		Texture:        tx,
		Generation:     gb,
		LiftAudit:      lift,
		Candidates:     candidates,
		OperatorAudit:  op,
		AxisAudit:      axis,
		MassAudit:      mass,
		Firewall:       fw,
		TruthStatement: "Gate 172 lifts the scalar-shape target into the triality Yukawa arena and proves that current finite data still select no canonical mass texture: exact triality is kind-blind, contact weights are kind-sensitive, diagonal generation spurions are aligned, and CKM/PMNS require at least two non-commuting finite texture operators.",
	}, nil
}

func buildLiftAudit(tr trialityyukawa.Analysis, tx texture.Analysis) TrialityLiftAudit {
	gen := tr.GenerationCount
	kinds := tx.FermionKinds
	return TrialityLiftAudit{
		GenerationCount:              gen,
		FermionKindBlocks:            kinds,
		YukawaTextureMatrices:        kinds,
		GeneralEntriesPerMatrix:      gen * gen,
		TotalGeneralRealEntries:      kinds * gen * gen,
		SymmetricEntriesPerMatrix:    6,
		TotalSymmetricRealEntries:    kinds * 6,
		TrialityInvariantDimPerKind:  tx.TrialityInvariantTextureDim,
		TotalTrialityInvariantParams: kinds * tx.TrialityInvariantTextureDim,
		FullMixingMaps:               tr.FullMixingMapCount,
		DiagonalTrialityChannels:     tr.DiagonalChannelCount,
		FullMixingChannels:           tr.FullMixingMapCount,
		ScalarMomentConstraints:      1,
		TextureUnderdetermined:       true,
		Verdict:                      "the triality lift turns one-generation amplitudes into four 3x3 Yukawa texture blocks; one scalar-shape moment is far below the data needed to select the matrices",
	}
}

func buildCandidates(ck contactkindassignment.Analysis, tr trialityyukawa.Analysis, tx texture.Analysis, gb generationbreak.Analysis) []TextureOperatorCandidate {
	kinds := len(ck.KindSignatures)
	gen := tr.GenerationCount
	generalParams := kinds * gen * gen
	return []TextureOperatorCandidate{
		{
			Name:                         "per-kind exact triality-invariant Yukawa texture",
			Kind:                         CandidateExactTriality,
			Source:                       "Gate 26/Gate 28 exact triality symmetry",
			KindAxis:                     "copied independently to each fermion kind; no contact high/low kind assignment",
			GenerationAxis:               "span{I, democratic/all-ones}; eigenpattern singlet plus doublet",
			MatrixBlocks:                 kinds,
			Parameters:                   kinds * tx.TrialityInvariantTextureDim,
			EigenPattern:                 "1+2 in generation space",
			BreaksAllThreeGenerations:    false,
			ProducesMixing:               false,
			Canonical:                    true,
			RequiresBridge:               false,
			RequiresBranchChoice:         false,
			KindSensitive:                false,
			GenerationSensitive:          true,
			TiesContactWeightsToKinds:    false,
			MatchesScalarShapeCondition:  false,
			SelectsYukawaAmplitudes:      false,
			SelectsMassEigenvalues:       false,
			SelectsCKMPMNS:               false,
			NonCommutingWithOtherTexture: false,
			Verdict:                      "canonical but too symmetric: exact triality cannot produce three distinct generation eigenvalues or choose the contact kind weights",
		},
		{
			Name:                         "Higgs/contact diagonal generation spurion",
			Kind:                         CandidateDiagonalSpurion,
			Source:                       "Gate 29 finite generation-breaking search",
			KindAxis:                     "universal over fermion kinds unless extra kind coupling is added",
			GenerationAxis:               "diagonal weights from {λmax, mean, λmin}: " + generationbreak.FormatFloatSlice(gb.BestCandidate.Eigenvalues),
			MatrixBlocks:                 kinds,
			Parameters:                   gen,
			EigenPattern:                 gb.BestCandidate.EigenPattern,
			BreaksAllThreeGenerations:    gb.DiagonalSpurionFound,
			ProducesMixing:               false,
			Canonical:                    gb.BestCandidate.Canonical,
			RequiresBridge:               gb.BestCandidate.RequiresBridge,
			RequiresBranchChoice:         true,
			KindSensitive:                false,
			GenerationSensitive:          true,
			TiesContactWeightsToKinds:    false,
			MatchesScalarShapeCondition:  false,
			SelectsYukawaAmplitudes:      false,
			SelectsMassEigenvalues:       false,
			SelectsCKMPMNS:               false,
			NonCommutingWithOtherTexture: false,
			Verdict:                      "promising generation-splitting seed, but it is diagonal, not canonical as a total Yukawa operator, and has no relative eigenbasis/mixing data",
		},
		{
			Name:                         "contact scalar-shape four-kind weights with generation identity",
			Kind:                         CandidateContactKindIdentity,
			Source:                       "Gate 169 scalar-shape target plus Gate 171 unresolved kind assignment",
			KindAxis:                     "two high and two low contact weights; six oriented kind assignments survive",
			GenerationAxis:               "identity on three triality sectors",
			MatrixBlocks:                 kinds,
			Parameters:                   kinds,
			EigenPattern:                 "kind anisotropy only; threefold generation degeneracy inside each kind",
			BreaksAllThreeGenerations:    false,
			ProducesMixing:               false,
			Canonical:                    false,
			RequiresBridge:               true,
			RequiresBranchChoice:         ck.AssignmentAudit.SurvivingBranchChoices > 0,
			KindSensitive:                true,
			GenerationSensitive:          false,
			TiesContactWeightsToKinds:    false,
			MatchesScalarShapeCondition:  ck.Consequence.ConditionalShapeStillValid,
			SelectsYukawaAmplitudes:      false,
			SelectsMassEigenvalues:       false,
			SelectsCKMPMNS:               false,
			NonCommutingWithOtherTexture: false,
			Verdict:                      "matches the scalar-shape moment only conditionally; it leaves generation masses degenerate and keeps the Gate-171 kind branch choice",
		},
		{
			Name:                         "separable contact-kind × diagonal-generation texture",
			Kind:                         CandidateSeparableTexture,
			Source:                       "formal product of Gate 169 contact-kind weights and Gate 29 diagonal generation spurion",
			KindAxis:                     "branch-chosen contact high/low weights",
			GenerationAxis:               "one shared diagonal generation basis for every fermion kind",
			MatrixBlocks:                 kinds,
			Parameters:                   kinds * gen,
			EigenPattern:                 "rank-separated kind scale times generation scale; all kind matrices commute",
			BreaksAllThreeGenerations:    gb.DiagonalSpurionFound,
			ProducesMixing:               false,
			Canonical:                    false,
			RequiresBridge:               true,
			RequiresBranchChoice:         true,
			KindSensitive:                true,
			GenerationSensitive:          true,
			TiesContactWeightsToKinds:    false,
			MatchesScalarShapeCondition:  ck.Consequence.ConditionalShapeStillValid,
			SelectsYukawaAmplitudes:      false,
			SelectsMassEigenvalues:       false,
			SelectsCKMPMNS:               false,
			NonCommutingWithOtherTexture: false,
			Verdict:                      "the first plausible architecture, but it is separable and simultaneously diagonal, so it cannot derive CKM/PMNS and remains branch-selected",
		},
		{
			Name:                         "unconstrained four 3x3 Yukawa matrices",
			Kind:                         CandidateGeneralTexture,
			Source:                       "charge-allowed Gate 26 full mixing space",
			KindAxis:                     "four independent fermion kinds",
			GenerationAxis:               "general 3x3 matrix per kind",
			MatrixBlocks:                 kinds,
			Parameters:                   generalParams,
			EigenPattern:                 "unconstrained; can fit hierarchy and mixing only if inserted by hand",
			BreaksAllThreeGenerations:    true,
			ProducesMixing:               true,
			Canonical:                    false,
			RequiresBridge:               true,
			RequiresBranchChoice:         true,
			KindSensitive:                true,
			GenerationSensitive:          true,
			TiesContactWeightsToKinds:    false,
			MatchesScalarShapeCondition:  false,
			SelectsYukawaAmplitudes:      false,
			SelectsMassEigenvalues:       false,
			SelectsCKMPMNS:               false,
			NonCommutingWithOtherTexture: true,
			Verdict:                      "large enough to fit the mass problem, but not derived; using it would violate the no-observed-input/no-hand-fit firewall",
		},
		{
			Name:                         "required non-commuting finite texture pair",
			Kind:                         CandidateNoncommutingPair,
			Source:                       "structural CKM/PMNS precondition exposed by Gate 172",
			KindAxis:                     "must distinguish at least quark up/down and lepton ν/e sectors",
			GenerationAxis:               "two or more non-commuting operators on the 3D generation carrier",
			MatrixBlocks:                 kinds,
			Parameters:                   0,
			EigenPattern:                 "not found in current finite data",
			BreaksAllThreeGenerations:    false,
			ProducesMixing:               false,
			Canonical:                    false,
			RequiresBridge:               true,
			RequiresBranchChoice:         false,
			KindSensitive:                true,
			GenerationSensitive:          true,
			TiesContactWeightsToKinds:    false,
			MatchesScalarShapeCondition:  false,
			SelectsYukawaAmplitudes:      false,
			SelectsMassEigenvalues:       false,
			SelectsCKMPMNS:               false,
			NonCommutingWithOtherTexture: false,
			Verdict:                      "this is the missing object, not a result of the current gate",
		},
	}
}

func auditOperators(cs []TextureOperatorCandidate) OperatorSearchAudit {
	canonical := 0
	canonicalBreaking := 0
	generationSplitting := 0
	mixing := 0
	noncommutingPairs := 0
	scalarConditional := 0
	branch := 0
	for _, c := range cs {
		if c.Canonical {
			canonical++
		}
		if c.Canonical && c.BreaksAllThreeGenerations {
			canonicalBreaking++
		}
		if c.BreaksAllThreeGenerations {
			generationSplitting++
		}
		if c.ProducesMixing {
			mixing++
		}
		if c.NonCommutingWithOtherTexture && c.Canonical {
			noncommutingPairs++
		}
		if c.MatchesScalarShapeCondition {
			scalarConditional++
		}
		if c.RequiresBranchChoice {
			branch++
		}
	}
	unique := canonicalBreaking > 0 && noncommutingPairs > 0
	return OperatorSearchAudit{
		CandidateCount:                     len(cs),
		CanonicalOperatorsFound:            canonical,
		CanonicalBreakingOperatorsFound:    canonicalBreaking,
		GenerationSplittingCandidatesFound: generationSplitting,
		MixingOperatorsFound:               mixing,
		NonCommutingPairsFound:             noncommutingPairs,
		ScalarShapeConditionalCandidates:   scalarConditional,
		BranchChoiceCandidates:             branch,
		UniqueTextureSelected:              unique,
		Verdict:                            "no canonical triality-lifted Yukawa texture is selected: the canonical exact-triality operator is too symmetric, while the useful scalar/generation candidates require branch choices and remain commuting/aligned",
	}
}

func auditAxes(ck contactkindassignment.Analysis, tx texture.Analysis, gb generationbreak.Analysis) AxisCouplingAudit {
	return AxisCouplingAudit{
		ContactKindAssignmentsSurvive: ck.AssignmentAudit.SurvivingBranchChoices,
		CanonicalKindPartitions:       ck.AssignmentAudit.CanonicalPartitionsFound,
		TrialityInvariantPattern:      fmt.Sprintf("%s with dim=%d per kind", tx.CandidateEigenPattern, tx.TrialityInvariantTextureDim),
		DiagonalGenerationSpurionSeen: gb.DiagonalSpurionFound,
		DiagonalSpurionCanonical:      gb.BestCandidate.Canonical,
		KindGenerationCouplingDerived: false,
		SeparableAnsatzOnly:           true,
		AllKindMatricesAligned:        true,
		Verdict:                       "the kind axis and generation axis remain uncoupled: contact weights need a kind assignment; triality operators act on generations; current products are separable and aligned",
	}
}

func FormatLiftAudit(a TrialityLiftAudit) string {
	return fmt.Sprintf("gen=%d kinds=%d matrices=%d generalPer=%d generalTotal=%d symmetricPer=%d symmetricTotal=%d trialityDimPer=%d trialityTotal=%d fullMaps=%d diagonalChannels=%d scalarConstraints=%d underdetermined=%t", a.GenerationCount, a.FermionKindBlocks, a.YukawaTextureMatrices, a.GeneralEntriesPerMatrix, a.TotalGeneralRealEntries, a.SymmetricEntriesPerMatrix, a.TotalSymmetricRealEntries, a.TrialityInvariantDimPerKind, a.TotalTrialityInvariantParams, a.FullMixingMaps, a.DiagonalTrialityChannels, a.ScalarMomentConstraints, a.TextureUnderdetermined)
}

func FormatCandidate(c TextureOperatorCandidate) string {
	return fmt.Sprintf("%s source=%s blocks=%d params=%d kindAxis={%s} genAxis={%s} pattern=%s breaks3=%t mixing=%t canonical=%t bridge=%t branch=%t kindSensitive=%t genSensitive=%t contactTie=%t scalarMatch=%t amplitudes=%t masses=%t CKM/PMNS=%t noncommuting=%t", c.Name, c.Source, c.MatrixBlocks, c.Parameters, c.KindAxis, c.GenerationAxis, c.EigenPattern, c.BreaksAllThreeGenerations, c.ProducesMixing, c.Canonical, c.RequiresBridge, c.RequiresBranchChoice, c.KindSensitive, c.GenerationSensitive, c.TiesContactWeightsToKinds, c.MatchesScalarShapeCondition, c.SelectsYukawaAmplitudes, c.SelectsMassEigenvalues, c.SelectsCKMPMNS, c.NonCommutingWithOtherTexture)
}

func FormatCandidates(cs []TextureOperatorCandidate) string {
	parts := make([]string, 0, len(cs))
	for _, c := range cs {
		parts = append(parts, FormatCandidate(c))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatOperatorAudit(a OperatorSearchAudit) string {
	return fmt.Sprintf("candidates=%d canonical=%d canonicalBreaking=%d splitting=%d mixing=%d noncommutingPairs=%d scalarConditional=%d branchChoice=%d unique=%t", a.CandidateCount, a.CanonicalOperatorsFound, a.CanonicalBreakingOperatorsFound, a.GenerationSplittingCandidatesFound, a.MixingOperatorsFound, a.NonCommutingPairsFound, a.ScalarShapeConditionalCandidates, a.BranchChoiceCandidates, a.UniqueTextureSelected)
}

func FormatAxisAudit(a AxisCouplingAudit) string {
	return fmt.Sprintf("kindAssignments=%d canonicalKindPartitions=%d trialityPattern=%s diagonalSpurion=%t diagonalCanonical=%t kindGenCoupling=%t separableOnly=%t aligned=%t", a.ContactKindAssignmentsSurvive, a.CanonicalKindPartitions, a.TrialityInvariantPattern, a.DiagonalGenerationSpurionSeen, a.DiagonalSpurionCanonical, a.KindGenerationCouplingDerived, a.SeparableAnsatzOnly, a.AllKindMatricesAligned)
}

func FormatMassAudit(a MassTextureAudit) string {
	return fmt.Sprintf("fourMatrices=%t singularValues=%t leftEigenbasis=%t needNoncommuting=%t foundNoncommuting=%t oneMoment=%t matricesDerived=%t masses=%t CKM/PMNS=%t", a.FourYukawaMatricesRecognized, a.MassesAreSingularValues, a.MixingNeedsRelativeLeftEigenbasis, a.AtLeastTwoNoncommutingNeeded, a.AtLeastTwoNoncommutingFound, a.ScalarShapeIsOneMomentConstraint, a.YukawaMatricesDerived, a.FermionMassesDerived, a.CKMPMNSDerived)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("gaugeClosed=%t scalarTarget=%t kindAssign=%t trialityLift=%t texture=%t amplitudes=%t hierarchy=%t masses=%t CKM/PMNS=%t constants=%t nullity=%d->%d", a.GaugeRatioClosed, a.ScalarShapeTargetAvailable, a.ContactKindAssignmentDerived, a.TrialityLiftPerformed, a.CanonicalTextureOperatorSelected, a.YukawaAmplitudesDerived, a.GenerationHierarchyDerived, a.FermionMassesDerived, a.CKMPMNSDerived, a.PhysicalConstantsDerived, a.ResidualNullityBefore, a.ResidualNullityAfter)
}
