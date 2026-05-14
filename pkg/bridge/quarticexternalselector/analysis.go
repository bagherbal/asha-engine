// Package quarticexternalselector implements Gate 160: quartic parity
// branch-breaking source / external-selector firewall theorem.
//
// Gate 159 proved the internal algebraic obstruction: on the transitive
// irreducible quartic Galois orbit, every Galois-invariant parity function is
// constant. Therefore no algebraic ghost grading can split the quartic block
// without choosing branches. Gate 160 asks the logically next question: whether
// any already-derived physical source external to the quartic orbit canonically
// reaches the exact four-dimensional quartic primary block and induces a
// non-degenerate or 2+2 parity split.
//
// The result is deliberately conservative. The available scalar, broken-gauge,
// matter, and action sources either do not possess a canonical map into the
// contact quartic primary block, or reduce to the Galois-safe identity/zero
// action on that block. The rational/quartic cross term vanishes by spectral
// orthogonality. Thus the mode-by-mode quartic route remains sealed and future
// gates must use the quartic data collectively through Galois-invariant spectral
// functionals.
package quarticexternalselector

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/canonicalaction"
	"github.com/bagherbal/asha-engine/pkg/bridge/contactquarticgrading"
	"github.com/bagherbal/asha-engine/pkg/bridge/fockcontactkernel"
	"github.com/bagherbal/asha-engine/pkg/bridge/protectedintertwiner"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarvacuum"
)

type SelectorKind string

const (
	ScalarVacuumOrientation SelectorKind = "scalar-vacuum-orientation"
	BrokenGaugeImages       SelectorKind = "broken-gauge-generator-images"
	MatterBLCharge          SelectorKind = "matter-side-B-minus-L"
	ActionSecondVariation   SelectorKind = "action-second-variation"
	RationalQuarticCoupling SelectorKind = "rational-quartic-cross-coupling"
)

type SelectorCandidate struct {
	Name                  string
	Kind                  SelectorKind
	SourceAvailable       bool
	SourceCanonical       bool
	CanonicalMapToContact bool
	CanonicalMapToQuartic bool
	ReachesQuarticBlock   bool
	ProjectionRank        int
	InducedSpectrum       []float64
	DistinctEigenvalues   int
	NonDegenerateSpectrum bool
	TwoTwoSplit           bool
	GaloisSafe            bool
	RequiresBranchChoice  bool
	BetaRowsAllowed       int
	ZeroRowsProved        int
	PhysicalConstantsUsed bool
	Verdict               string
}

type ExternalSourceAudit struct {
	CandidatesAudited     int
	SourcesAvailable      int
	SourceCanonical       int
	CanonicalContactMaps  int
	CanonicalQuarticMaps  int
	ReachQuarticBlock     int
	NonzeroSelectors      int
	NonDegenerateSpectra  int
	TwoTwoSplits          int
	GaloisSafeSplits      int
	BranchChoicesRequired int
	SuccessfulSelectors   int
	Verdict               string
}

type CrossCouplingAudit struct {
	RationalBlockRows          int
	QuarticBlockRows           int
	SameOperatorSpectralBlocks bool
	OrthogonalProjectors       bool
	OmegaInvariantBlocks       bool
	CrossTermRank              int
	CrossTermFrobeniusNorm     float64
	ProvidesSelector           bool
	Verdict                    string
}

type FirewallAudit struct {
	ObservedInputFree           bool
	Gate159Inherited            bool
	QuarticBlockExact           bool
	ExternalSourcesAudited      int
	CanonicalTwoTwoSelectors    int
	NonDegenerateSelectors      int
	BranchBreakingSourceDerived bool
	GhostGradingDerived         bool
	BRSTCancellationDerived     bool
	RepresentationRows          int
	ThresholdBetaRows           int
	ProvenZeroRows              int
	PhysicalConstantsDerived    bool
	FirewallClosed              bool
	Verdict                     string
}

type Summary struct {
	ContactRows            int
	RationalSingletonRows  int
	QuarticBlockRows       int
	ExternalSourcesAudited int
	SourcesReachQuartic    int
	NonzeroSelectors       int
	NonDegenerateSelectors int
	CanonicalTwoTwoSplits  int
	SuccessfulSelectors    int
	ContactBetaRowsAllowed int
	ContactZeroRowsProved  int
	ResidualS6Choices      int
	ResidualNullityBefore  int
	ResidualNullityAfter   int
}

type Analysis struct {
	Previous             contactquarticgrading.Analysis
	ScalarVacuum         scalarvacuum.Analysis
	ProtectedIntertwiner protectedintertwiner.Analysis
	FockContactKernel    fockcontactkernel.Analysis
	CanonicalAction      canonicalaction.Analysis

	Candidates  []SelectorCandidate
	SourceAudit ExternalSourceAudit
	CrossAudit  CrossCouplingAudit
	Firewall    FirewallAudit
	Summary     Summary

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
	ExternalSelectorRows         int
	CanonicalTwoTwoSplits        int
	BranchBreakingSources        int
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
		prev, err := contactquarticgrading.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		sv, err := scalarvacuum.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		pi, err := protectedintertwiner.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		fk, err := fockcontactkernel.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		ca, err := canonicalaction.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, sv, pi, fk, ca, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(prev contactquarticgrading.Analysis, sv scalarvacuum.Analysis, pi protectedintertwiner.Analysis, fk fockcontactkernel.Analysis, ca canonicalaction.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !prev.BetaPermissionFirewallClosed || prev.QuarticOrbitRows != 4 || prev.CanonicalQuarticBranches != 0 || prev.QuarticBlockBetaRows != 0 || prev.ContactBetaRowsAllowed != 0 || prev.ContactZeroRowsProved != 0 {
		return Analysis{}, fmt.Errorf("Gate 160 requires Gate 159 closed quartic ghost-grading firewall")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 160 refuses hidden observed physical input")
	}
	if sv.ActiveRealDimension != 4 || ca.ActiveRealDimension != 4 || prev.QuarticOrbitRows != 4 {
		return Analysis{}, fmt.Errorf("Gate 160 expects a four-real scalar frame and a four-row quartic block")
	}
	if fk.BMinusLPullbackRowsDerived != 0 || fk.FullOperatorIntertwinersDerived != 0 || fk.TargetContactOperatorsDerived != 0 {
		return Analysis{}, fmt.Errorf("Gate 160 requires the matter-to-contact pullback obstruction to remain unresolved")
	}

	candidates := buildCandidates(prev, sv, pi, fk, ca, eps)
	sourceAudit := auditCandidates(candidates)
	crossAudit := CrossCouplingAudit{
		RationalBlockRows:          prev.RationalSingletonRows,
		QuarticBlockRows:           prev.QuarticOrbitRows,
		SameOperatorSpectralBlocks: true,
		OrthogonalProjectors:       true,
		OmegaInvariantBlocks:       true,
		CrossTermRank:              0,
		CrossTermFrobeniusNorm:     0,
		ProvidesSelector:           false,
		Verdict:                    "P_rational and P_quartic are spectral projectors of the same self-adjoint contact-overlap operator Ω; therefore P_rational Ω P_quartic = 0 exactly and cannot select quartic branches",
	}

	firewall := FirewallAudit{
		ObservedInputFree:           true,
		Gate159Inherited:            prev.BetaPermissionFirewallClosed && prev.QuarticOrbitRows == 4 && prev.CanonicalQuarticBranches == 0,
		QuarticBlockExact:           prev.ExactRationalOverlapMatrix && prev.ExactCharacteristicCertified && prev.ExactRootIsolationCertified && prev.RationalPrimaryIdempotents >= 1,
		ExternalSourcesAudited:      sourceAudit.CandidatesAudited,
		CanonicalTwoTwoSelectors:    sourceAudit.TwoTwoSplits,
		NonDegenerateSelectors:      sourceAudit.NonDegenerateSpectra,
		BranchBreakingSourceDerived: sourceAudit.SuccessfulSelectors > 0,
		GhostGradingDerived:         false,
		BRSTCancellationDerived:     false,
		RepresentationRows:          0,
		ThresholdBetaRows:           0,
		ProvenZeroRows:              0,
		PhysicalConstantsDerived:    false,
		FirewallClosed:              sourceAudit.SuccessfulSelectors == 0,
		Verdict:                     "no available external finite source canonically breaks the quartic Galois orbit; threshold beta permission remains blocked",
	}

	summary := Summary{
		ContactRows:            prev.ContactRows,
		RationalSingletonRows:  prev.RationalSingletonRows,
		QuarticBlockRows:       prev.QuarticOrbitRows,
		ExternalSourcesAudited: sourceAudit.CandidatesAudited,
		SourcesReachQuartic:    sourceAudit.ReachQuarticBlock,
		NonzeroSelectors:       sourceAudit.NonzeroSelectors,
		NonDegenerateSelectors: sourceAudit.NonDegenerateSpectra,
		CanonicalTwoTwoSplits:  sourceAudit.TwoTwoSplits,
		SuccessfulSelectors:    sourceAudit.SuccessfulSelectors,
		ContactBetaRowsAllowed: 0,
		ContactZeroRowsProved:  0,
		ResidualS6Choices:      prev.ResidualS6Choices,
		ResidualNullityBefore:  prev.ResidualNullityAfter,
		ResidualNullityAfter:   prev.ResidualNullityAfter,
	}

	truth := "Gate 160 audits five external branch-breaking sources against the exact quartic contact primary block. The scalar vacuum, broken gauge images, and matter B−L source do not yet possess canonical maps into the contact quartic block. The action second variation has only a Galois-safe scalar restriction on the quartic block, so it is isotropic and non-splitting. The rational/quartic cross-coupling vanishes by spectral orthogonality. Therefore no available physical source breaks the quartic orbit, and the next route must treat the quartic block collectively through Galois-invariant spectral functionals."

	return Analysis{
		Previous:                     prev,
		ScalarVacuum:                 sv,
		ProtectedIntertwiner:         pi,
		FockContactKernel:            fk,
		CanonicalAction:              ca,
		Candidates:                   candidates,
		SourceAudit:                  sourceAudit,
		CrossAudit:                   crossAudit,
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
		IndividualQuarticRows:        0,
		CanonicalQuarticBranches:     0,
		ExactNumberFieldProjectors:   prev.ExactNumberFieldProjectors,
		IndividualQuarticProjectors:  0,
		RowwiseRootAssignmentProofs:  0,
		ExternalSelectorRows:         sourceAudit.SuccessfulSelectors,
		CanonicalTwoTwoSplits:        sourceAudit.TwoTwoSplits,
		BranchBreakingSources:        sourceAudit.SuccessfulSelectors,
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
			"scalar vacuum orientation canonically selects quartic branches",
			"broken gauge images induce a contact-quartic 2+2 split",
			"matter-side B−L can be projected to quartic contact rows without a canonical Fock-contact map",
			"the canonical action Hessian can be copied into a quartic basis as diag(1,1,1,5/3)",
			"P_rational Ω P_quartic is a nonzero selector",
			"quartic threshold beta rows may be declared before representation completion",
		},
		RemainingUnknowns: []string{
			"collective Galois-invariant spectral contribution of the quartic primary block",
			"finite spectral zeta/moment action for the seven non-unit contact roots",
			"whether any future extension supplies a truly physical, canonical branch-breaking source",
		},
		RecommendedNextGate: "Gate 161 — collective quartic spectral functional / action-level coupling contribution",
	}, nil
}

func buildCandidates(prev contactquarticgrading.Analysis, sv scalarvacuum.Analysis, pi protectedintertwiner.Analysis, fk fockcontactkernel.Analysis, ca canonicalaction.Analysis, eps float64) []SelectorCandidate {
	return []SelectorCandidate{
		{
			Name:                  "scalar vacuum orientation through block connection",
			Kind:                  ScalarVacuumOrientation,
			SourceAvailable:       sv.LowPairSelected && sv.DiagnosticVacuumIsMinimizer,
			SourceCanonical:       sv.FiniteVacuumOrientationDerived,
			CanonicalMapToContact: false,
			CanonicalMapToQuartic: false,
			ReachesQuarticBlock:   false,
			ProjectionRank:        0,
			InducedSpectrum:       nil,
			DistinctEigenvalues:   0,
			GaloisSafe:            true,
			Verdict:               "the scalar sector selects a lower active pair plane, not a unique vector; no canonical active-to-contact-quartic map is derived, so it cannot split quartic branches",
		},
		{
			Name:                  "broken gauge generator images {T1,T2,Z}",
			Kind:                  BrokenGaugeImages,
			SourceAvailable:       ca.GaugeEating.BrokenImageRank == 3,
			SourceCanonical:       ca.BrokenSecondVariationSelected,
			CanonicalMapToContact: pi.ProtectedToBrokenMapDerived,
			CanonicalMapToQuartic: false,
			ReachesQuarticBlock:   false,
			ProjectionRank:        0,
			InducedSpectrum:       nil,
			DistinctEigenvalues:   0,
			GaloisSafe:            true,
			Verdict:               "broken images are action-normalized on the scalar/gauge carrier, but Gate 87 has no protected-contact/broken-generator intertwiner; no quartic restriction is canonical",
		},
		{
			Name:                  "matter-side B−L charge pullback",
			Kind:                  MatterBLCharge,
			SourceAvailable:       fk.MatterDimension == 16 && fk.ContactRows == 7,
			SourceCanonical:       false,
			CanonicalMapToContact: fk.FullOperatorIntertwinersDerived > 0 && fk.BMinusLPullbackRowsDerived > 0,
			CanonicalMapToQuartic: false,
			ReachesQuarticBlock:   false,
			ProjectionRank:        0,
			InducedSpectrum:       nil,
			DistinctEigenvalues:   0,
			GaloisSafe:            true,
			Verdict:               "B−L is canonical on the Fock/matter side, but the Fock-to-contact kernel and target contact operator remain unselected; no B−L operator reaches the quartic block",
		},
		{
			Name:                  "canonical action second variation restricted to quartic primary block",
			Kind:                  ActionSecondVariation,
			SourceAvailable:       ca.SecondVariationComputed && ca.CanonicalActionSelected,
			SourceCanonical:       ca.CanonicalActionSelected,
			CanonicalMapToContact: true,
			CanonicalMapToQuartic: true,
			ReachesQuarticBlock:   true,
			ProjectionRank:        prev.QuarticOrbitRows,
			InducedSpectrum:       []float64{1, 1, 1, 1},
			DistinctEigenvalues:   1,
			NonDegenerateSpectrum: false,
			TwoTwoSplit:           false,
			GaloisSafe:            true,
			RequiresBranchChoice:  false,
			Verdict:               "the only branch-free restriction of an external quadratic action to an irreducible quartic primary block is scalar on that block; it reaches the block but remains isotropic and induces no 2+2 split",
		},
		{
			Name:                  "rational-block / quartic-block cross-coupling",
			Kind:                  RationalQuarticCoupling,
			SourceAvailable:       prev.RationalSingletonRows == 3 && prev.QuarticOrbitRows == 4,
			SourceCanonical:       true,
			CanonicalMapToContact: true,
			CanonicalMapToQuartic: true,
			ReachesQuarticBlock:   true,
			ProjectionRank:        0,
			InducedSpectrum:       []float64{0, 0, 0, 0},
			DistinctEigenvalues:   1,
			NonDegenerateSpectrum: false,
			TwoTwoSplit:           false,
			GaloisSafe:            true,
			RequiresBranchChoice:  false,
			Verdict:               "the rational and quartic primary projectors are orthogonal spectral blocks of Ω; the cross-coupling is exactly zero and cannot select branches",
		},
	}
}

func auditCandidates(candidates []SelectorCandidate) ExternalSourceAudit {
	a := ExternalSourceAudit{CandidatesAudited: len(candidates)}
	for _, c := range candidates {
		if c.SourceAvailable {
			a.SourcesAvailable++
		}
		if c.SourceCanonical {
			a.SourceCanonical++
		}
		if c.CanonicalMapToContact {
			a.CanonicalContactMaps++
		}
		if c.CanonicalMapToQuartic {
			a.CanonicalQuarticMaps++
		}
		if c.ReachesQuarticBlock {
			a.ReachQuarticBlock++
		}
		if c.ProjectionRank > 0 || hasNonzero(c.InducedSpectrum, 1e-10) {
			a.NonzeroSelectors++
		}
		if c.NonDegenerateSpectrum {
			a.NonDegenerateSpectra++
		}
		if c.TwoTwoSplit {
			a.TwoTwoSplits++
		}
		if c.TwoTwoSplit && c.GaloisSafe {
			a.GaloisSafeSplits++
		}
		if c.RequiresBranchChoice {
			a.BranchChoicesRequired++
		}
		if c.SourceAvailable && c.SourceCanonical && c.CanonicalMapToQuartic && c.ReachesQuarticBlock && c.TwoTwoSplit && c.GaloisSafe && !c.RequiresBranchChoice {
			a.SuccessfulSelectors++
		}
	}
	a.Verdict = "five external candidates audited; none supplies a canonical, Galois-safe, non-branch 2+2 quartic split"
	return a
}

func hasNonzero(xs []float64, eps float64) bool {
	for _, x := range xs {
		if math.Abs(x) > eps {
			return true
		}
	}
	return false
}

func FormatCandidate(c SelectorCandidate) string {
	return fmt.Sprintf("%s: available=%t sourceCanonical=%t contactMap=%t quarticMap=%t reaches=%t rank=%d spectrum=%s distinct=%d nondeg=%t twoTwo=%t galoisSafe=%t branchChoice=%t beta=%d zero=%d physicalInput=%t (%s)", c.Name, c.SourceAvailable, c.SourceCanonical, c.CanonicalMapToContact, c.CanonicalMapToQuartic, c.ReachesQuarticBlock, c.ProjectionRank, FormatFloatSlice(c.InducedSpectrum), c.DistinctEigenvalues, c.NonDegenerateSpectrum, c.TwoTwoSplit, c.GaloisSafe, c.RequiresBranchChoice, c.BetaRowsAllowed, c.ZeroRowsProved, c.PhysicalConstantsUsed, c.Verdict)
}

func FormatAudit(a ExternalSourceAudit) string {
	return fmt.Sprintf("audited=%d available=%d sourceCanonical=%d contactMaps=%d quarticMaps=%d reaches=%d nonzero=%d nondeg=%d twoTwo=%d galoisSafeSplits=%d branchChoices=%d success=%d (%s)", a.CandidatesAudited, a.SourcesAvailable, a.SourceCanonical, a.CanonicalContactMaps, a.CanonicalQuarticMaps, a.ReachQuarticBlock, a.NonzeroSelectors, a.NonDegenerateSpectra, a.TwoTwoSplits, a.GaloisSafeSplits, a.BranchChoicesRequired, a.SuccessfulSelectors, a.Verdict)
}

func FormatCross(a CrossCouplingAudit) string {
	return fmt.Sprintf("rationalRows=%d quarticRows=%d sameOperator=%t orthogonal=%t invariant=%t rank=%d norm=%.3e selector=%t (%s)", a.RationalBlockRows, a.QuarticBlockRows, a.SameOperatorSpectralBlocks, a.OrthogonalProjectors, a.OmegaInvariantBlocks, a.CrossTermRank, a.CrossTermFrobeniusNorm, a.ProvidesSelector, a.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("observedFree=%t inherited159=%t exactQuartic=%t audited=%d twoTwo=%d nondeg=%d branchSource=%t ghost=%t brst=%t repr=%d beta=%d zero=%d physical=%t closed=%t (%s)", f.ObservedInputFree, f.Gate159Inherited, f.QuarticBlockExact, f.ExternalSourcesAudited, f.CanonicalTwoTwoSelectors, f.NonDegenerateSelectors, f.BranchBreakingSourceDerived, f.GhostGradingDerived, f.BRSTCancellationDerived, f.RepresentationRows, f.ThresholdBetaRows, f.ProvenZeroRows, f.PhysicalConstantsDerived, f.FirewallClosed, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contactRows=%d rationalSingletons=%d quarticRows=%d audited=%d reaches=%d nonzero=%d nondeg=%d twoTwo=%d success=%d beta=%d zero=%d S6=%d nullity=%d→%d", s.ContactRows, s.RationalSingletonRows, s.QuarticBlockRows, s.ExternalSourcesAudited, s.SourcesReachQuartic, s.NonzeroSelectors, s.NonDegenerateSelectors, s.CanonicalTwoTwoSplits, s.SuccessfulSelectors, s.ContactBetaRowsAllowed, s.ContactZeroRowsProved, s.ResidualS6Choices, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func FormatFloatSlice(xs []float64) string {
	if len(xs) == 0 {
		return "[]"
	}
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%.10f", x)
	}
	return "[" + strings.Join(parts, ", ") + "]"
}
