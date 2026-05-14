// Package fockcontactkernel implements Gate 138: Fock-contact kernel
// selection / operator-intertwining obstruction theorem.
//
// Gate 137 showed that generic maps H_Fock -> R7_contact exist, but they
// require choosing a 9-dimensional kernel. Gate 138 asks whether that kernel
// becomes canonical when the map is required to intertwine the matter-side
// operators that would have to pull back to contact rows: temporal T3R,
// chirality/grading, B-L, and SU(2)L/current semantics.
//
// The result is a sharper obstruction. Operator-intertwining first demands an
// invariant kernel. T3R invariance reduces arbitrary Gr(9,16) choices to a
// family of spectral split patterns; adding chirality refines it to joint
// 4+4+4+4 split patterns. But neither refinement selects a unique kernel, nor
// does it construct a contact-row target operator. Thus the Fock-to-contact map
// remains underived and the contact beta firewall remains closed.
package fockcontactkernel

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactt3rpullback"
)

type KernelSearchKind string

const (
	UnconstrainedRankSevenKernel KernelSearchKind = "unconstrained-rank-seven-kernel"
	T3RInvariantKernel           KernelSearchKind = "T3R-invariant-kernel"
	T3RChiralityInvariantKernel  KernelSearchKind = "T3R-chirality-invariant-kernel"
	FullOperatorIntertwiner      KernelSearchKind = "full-operator-intertwiner"
	S6LeptoquarkAssignmentKernel KernelSearchKind = "S6-leptoquark-assignment-kernel"
	ContactSpectralKernel        KernelSearchKind = "contact-spectral-kernel"
)

type KernelCandidate struct {
	Name                           string
	Kind                           KernelSearchKind
	Domain                         string
	Codomain                       string
	DomainDim                      int
	CodomainDim                    int
	KernelDim                      int
	GenericExists                  bool
	InvariantUnderT3R              bool
	InvariantUnderChirality        bool
	InvariantUnderBMinusL          bool
	InvariantUnderSU2L             bool
	IntertwinesT3R                 bool
	IntertwinesChirality           bool
	IntertwinesBMinusL             bool
	IntertwinesSU2L                bool
	TargetContactOperatorDerived   bool
	CanonicalKernelDerived         bool
	RepresentationRowsDerived      int
	ContactBetaRowsAllowed         int
	SplitPatternCount              int
	ResidualContinuousDimensionMin int
	ResidualContinuousDimensionMax int
	HiddenDiscreteChoices          int
	Obstruction                    string
}

type SplitAudit struct {
	Name                           string
	SectorDims                     []int
	QuotientDim                    int
	KernelDim                      int
	PatternCount                   int
	ResidualContinuousDimensionMin int
	ResidualContinuousDimensionMax int
	CanonicalPatternSelected       bool
	TargetOperatorSelected         bool
	Obstruction                    string
}

type Summary struct {
	MatterDimension                    int
	ContactRows                        int
	RequiredKernelDim                  int
	UnconstrainedGrassmannDimension    int
	T3RInvariantSplitPatterns          int
	T3RInvariantResidualDimMin         int
	T3RInvariantResidualDimMax         int
	T3RChiralityInvariantSplitPatterns int
	T3RChiralityResidualDimMin         int
	T3RChiralityResidualDimMax         int
	CanonicalKernelCandidates          int
	FullOperatorIntertwinersDerived    int
	TargetContactOperatorsDerived      int
	T3RPullbackRowsDerived             int
	ChiralityPullbackRowsDerived       int
	BMinusLPullbackRowsDerived         int
	SU2LPullbackRowsDerived            int
	HyperchargeRowsDerived             int
	RepresentationCompleteRows         int
	ContactBetaRowsAllowed             int
	ContactZeroRowsProved              int
	ResidualS6Choices                  int
	ResidualNullityBefore              int
	ResidualNullityAfter               int
}

type Analysis struct {
	Previous contactt3rpullback.Analysis

	Candidates []KernelCandidate
	Splits     []SplitAudit
	Summary    Summary

	MatterDimension                 int
	ContactRows                     int
	RequiredKernelDim               int
	UnconstrainedGrassmannDimension int
	T3RInvariantSplitPatterns       int
	T3RChiralitySplitPatterns       int
	CanonicalKernelCandidates       int
	FullOperatorIntertwinersDerived int
	TargetContactOperatorsDerived   int
	T3RPullbackRowsDerived          int
	ChiralityPullbackRowsDerived    int
	BMinusLPullbackRowsDerived      int
	SU2LPullbackRowsDerived         int
	HyperchargeRowsDerived          int
	RepresentationCompleteRows      int
	RepresentationOpenRows          int
	ContactBetaRowsAllowed          int
	ContactZeroRowsProved           int
	FullBetaMatchingTensorDerived   bool
	ThresholdCorrectedBetaDerived   bool
	BetaPermissionFirewallClosed    bool

	ResidualS6Choices        int
	ResidualNullityBefore    int
	ResidualNullityAfter     int
	HiddenObservedInputUsed  bool
	PhysicalWeakAngleDerived bool
	FineStructureDerived     bool
	PhysicalMassesDerived    bool
	PhysicalScaleDerived     bool

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
		prev, err := contactt3rpullback.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactt3rpullback.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.MatterDimension != 16 || prev.ContactRows != 7 || prev.FockToContactIntertwinersDerived != 0 || prev.ContactBetaRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 138 requires Gate 137 Fock-to-contact firewall to remain closed")
	}
	if prev.T3RPullbackRowsDerived != 0 || prev.ChiralityPullbackRowsDerived != 0 || prev.BMinusLPullbackRowsDerived != 0 || prev.SU2LPullbackRowsDerived != 0 || prev.HyperchargeRowsDerived != 0 {
		return Analysis{}, fmt.Errorf("Gate 138 requires no contact pullback rows from Gate 137")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 138 refuses hidden physical input from Gate 137")
	}

	matterDim := prev.MatterDimension
	contactRows := prev.ContactRows
	kernelDim := matterDim - contactRows
	if kernelDim != 9 {
		return Analysis{}, fmt.Errorf("unexpected Fock/contact kernel dimension: %d", kernelDim)
	}

	unconstrainedGrassmann := kernelDim * contactRows
	t3rSplit := auditSplit("T3R spectral split", []int{8, 8}, contactRows)
	jointSplit := auditSplit("T3R/chirality joint split", []int{4, 4, 4, 4}, contactRows)

	candidates := buildCandidates(prev, unconstrainedGrassmann, t3rSplit, jointSplit)
	canonicalKernels := count(candidates, func(c KernelCandidate) bool { return c.CanonicalKernelDerived })
	fullIntertwiners := count(candidates, func(c KernelCandidate) bool { return c.Kind == FullOperatorIntertwiner && c.CanonicalKernelDerived })
	targetOps := count(candidates, func(c KernelCandidate) bool { return c.TargetContactOperatorDerived })

	summary := Summary{
		MatterDimension:                    matterDim,
		ContactRows:                        contactRows,
		RequiredKernelDim:                  kernelDim,
		UnconstrainedGrassmannDimension:    unconstrainedGrassmann,
		T3RInvariantSplitPatterns:          t3rSplit.PatternCount,
		T3RInvariantResidualDimMin:         t3rSplit.ResidualContinuousDimensionMin,
		T3RInvariantResidualDimMax:         t3rSplit.ResidualContinuousDimensionMax,
		T3RChiralityInvariantSplitPatterns: jointSplit.PatternCount,
		T3RChiralityResidualDimMin:         jointSplit.ResidualContinuousDimensionMin,
		T3RChiralityResidualDimMax:         jointSplit.ResidualContinuousDimensionMax,
		CanonicalKernelCandidates:          canonicalKernels,
		FullOperatorIntertwinersDerived:    fullIntertwiners,
		TargetContactOperatorsDerived:      targetOps,
		T3RPullbackRowsDerived:             0,
		ChiralityPullbackRowsDerived:       0,
		BMinusLPullbackRowsDerived:         0,
		SU2LPullbackRowsDerived:            0,
		HyperchargeRowsDerived:             0,
		RepresentationCompleteRows:         0,
		ContactBetaRowsAllowed:             0,
		ContactZeroRowsProved:              0,
		ResidualS6Choices:                  prev.ResidualS6Choices,
		ResidualNullityBefore:              prev.ResidualNullityAfter,
		ResidualNullityAfter:               prev.ResidualNullityAfter,
	}

	truth := "Gate 138 upgrades the Gate 137 obstruction from generic map-counting to operator-intertwining. A Fock-to-contact quotient must choose a 9-dimensional kernel in H_Fock and that kernel must be invariant under the operators it is supposed to transport. T3R invariance leaves 8 spectral split patterns with continuous families of kernels; adding chirality leaves 80 joint split patterns. No target contact operator, canonical kernel, or full intertwiner is selected, so matter-side T3R/chirality still cannot become contact hypercharge rows."

	return Analysis{
		Previous:                        prev,
		Candidates:                      candidates,
		Splits:                          []SplitAudit{t3rSplit, jointSplit},
		Summary:                         summary,
		MatterDimension:                 matterDim,
		ContactRows:                     contactRows,
		RequiredKernelDim:               kernelDim,
		UnconstrainedGrassmannDimension: unconstrainedGrassmann,
		T3RInvariantSplitPatterns:       t3rSplit.PatternCount,
		T3RChiralitySplitPatterns:       jointSplit.PatternCount,
		CanonicalKernelCandidates:       canonicalKernels,
		FullOperatorIntertwinersDerived: fullIntertwiners,
		TargetContactOperatorsDerived:   targetOps,
		T3RPullbackRowsDerived:          0,
		ChiralityPullbackRowsDerived:    0,
		BMinusLPullbackRowsDerived:      0,
		SU2LPullbackRowsDerived:         0,
		HyperchargeRowsDerived:          0,
		RepresentationCompleteRows:      0,
		RepresentationOpenRows:          contactRows,
		ContactBetaRowsAllowed:          0,
		ContactZeroRowsProved:           0,
		FullBetaMatchingTensorDerived:   false,
		ThresholdCorrectedBetaDerived:   false,
		BetaPermissionFirewallClosed:    true,
		ResidualS6Choices:               prev.ResidualS6Choices,
		ResidualNullityBefore:           prev.ResidualNullityAfter,
		ResidualNullityAfter:            prev.ResidualNullityAfter,
		HiddenObservedInputUsed:         false,
		PhysicalWeakAngleDerived:        false,
		FineStructureDerived:            false,
		PhysicalMassesDerived:           false,
		PhysicalScaleDerived:            false,
		TruthStatement:                  truth,
		RejectedClaims: []string{
			"kernel invariance under T3R uniquely selects a Fock-to-contact quotient",
			"T3R/chirality joint eigenspace splitting supplies contact row semantics",
			"a quotient operator on contact rows exists without a selected target contact operator",
			"B-L or SU(2)L intertwining can be asserted before the kernel is chosen",
			"the leptoquark S6 assignment is solved by kernel dimension alone",
			"contact spectral identity is a Fock-to-contact intertwiner",
		},
		RemainingUnknowns: []string{
			"canonical 9-dimensional Fock kernel or quotient relation",
			"target contact operators for T3R, chirality, B-L, and SU(2)L",
			"operator equations P A = B P for the contact carrier",
			"S6 assignment of leptoquark current slots to contact rows",
			"local field map, hypercharge row, mass activation, and decoupling rule",
		},
		RecommendedNextGate: "Gate 139 — contact target-operator reconstruction / quotient-side T3R spectrum search",
	}, nil
}

func buildCandidates(prev contactt3rpullback.Analysis, grassmannDim int, t3r SplitAudit, joint SplitAudit) []KernelCandidate {
	matterDim := prev.MatterDimension
	contactRows := prev.ContactRows
	kernelDim := matterDim - contactRows
	return []KernelCandidate{
		{Name: "arbitrary rank-seven quotient H_Fock/K", Kind: UnconstrainedRankSevenKernel, Domain: "H_Fock", Codomain: "R7_contact", DomainDim: matterDim, CodomainDim: contactRows, KernelDim: kernelDim, GenericExists: true, SplitPatternCount: 1, ResidualContinuousDimensionMin: grassmannDim, ResidualContinuousDimensionMax: grassmannDim, Obstruction: fmt.Sprintf("choosing K∈Gr(9,16) has %d continuous dimensions and no finite selector", grassmannDim)},
		{Name: "T3R-invariant quotient kernel", Kind: T3RInvariantKernel, Domain: "H_Fock=T+⊕T-", Codomain: "R7_contact", DomainDim: matterDim, CodomainDim: contactRows, KernelDim: kernelDim, GenericExists: true, InvariantUnderT3R: true, IntertwinesT3R: false, SplitPatternCount: t3r.PatternCount, ResidualContinuousDimensionMin: t3r.ResidualContinuousDimensionMin, ResidualContinuousDimensionMax: t3r.ResidualContinuousDimensionMax, Obstruction: t3r.Obstruction},
		{Name: "T3R/chirality joint-invariant quotient kernel", Kind: T3RChiralityInvariantKernel, Domain: "H_Fock=⊕ four 4D joint sectors", Codomain: "R7_contact", DomainDim: matterDim, CodomainDim: contactRows, KernelDim: kernelDim, GenericExists: true, InvariantUnderT3R: true, InvariantUnderChirality: true, IntertwinesT3R: false, IntertwinesChirality: false, SplitPatternCount: joint.PatternCount, ResidualContinuousDimensionMin: joint.ResidualContinuousDimensionMin, ResidualContinuousDimensionMax: joint.ResidualContinuousDimensionMax, Obstruction: joint.Obstruction},
		{Name: "full T3R/chirality/B-L/SU2L operator intertwiner", Kind: FullOperatorIntertwiner, Domain: "H_Fock", Codomain: "R7_contact", DomainDim: matterDim, CodomainDim: contactRows, KernelDim: kernelDim, GenericExists: false, InvariantUnderT3R: false, InvariantUnderChirality: false, InvariantUnderBMinusL: false, InvariantUnderSU2L: false, IntertwinesT3R: false, IntertwinesChirality: false, IntertwinesBMinusL: false, IntertwinesSU2L: false, TargetContactOperatorDerived: false, CanonicalKernelDerived: false, Obstruction: "no contact-side target operators B_T3R, B_chirality, B_BL, or B_SU2L are derived, so equations P A = B P cannot be solved canonically"},
		{Name: "leptoquark six-block kernel refinement", Kind: S6LeptoquarkAssignmentKernel, Domain: "LQ_current_six", Codomain: "contact_six", DomainDim: 6, CodomainDim: 6, KernelDim: 0, GenericExists: true, HiddenDiscreteChoices: prev.ResidualS6Choices, Obstruction: fmt.Sprintf("even after kernel restrictions, assigning six current slots to six contact rows still has S6=%d hidden permutations", prev.ResidualS6Choices)},
		{Name: "contact spectral identity kernel", Kind: ContactSpectralKernel, Domain: "R7_contact", Codomain: "R7_contact", DomainDim: contactRows, CodomainDim: contactRows, KernelDim: 0, GenericExists: true, CanonicalKernelDerived: true, TargetContactOperatorDerived: false, Obstruction: "canonical contact-side identity has zero kernel and cannot be the required 9-dimensional Fock kernel"},
	}
}

func auditSplit(name string, sectorDims []int, quotientDim int) SplitAudit {
	patterns := enumerateSplits(sectorDims, quotientDim)
	minDim := -1
	maxDim := -1
	for _, q := range patterns {
		d := 0
		kernelSum := 0
		for i, qi := range q {
			ki := sectorDims[i] - qi
			kernelSum += ki
			d += ki * qi
		}
		if kernelSum != sum(sectorDims)-quotientDim {
			panic("internal split enumeration produced wrong kernel dimension")
		}
		if minDim < 0 || d < minDim {
			minDim = d
		}
		if d > maxDim {
			maxDim = d
		}
	}
	return SplitAudit{
		Name:                           name,
		SectorDims:                     append([]int(nil), sectorDims...),
		QuotientDim:                    quotientDim,
		KernelDim:                      sum(sectorDims) - quotientDim,
		PatternCount:                   len(patterns),
		ResidualContinuousDimensionMin: minDim,
		ResidualContinuousDimensionMax: maxDim,
		CanonicalPatternSelected:       false,
		TargetOperatorSelected:         false,
		Obstruction:                    fmt.Sprintf("%s permits %d quotient split patterns with residual continuous kernel dimensions %d..%d; no contact target operator selects one", name, len(patterns), minDim, maxDim),
	}
}

func enumerateSplits(sectorDims []int, quotientDim int) [][]int {
	var out [][]int
	var rec func(i int, remaining int, cur []int)
	rec = func(i int, remaining int, cur []int) {
		if i == len(sectorDims) {
			if remaining == 0 {
				out = append(out, append([]int(nil), cur...))
			}
			return
		}
		max := sectorDims[i]
		if remaining < max {
			max = remaining
		}
		for q := 0; q <= max; q++ {
			rec(i+1, remaining-q, append(cur, q))
		}
	}
	rec(0, quotientDim, nil)
	return out
}

func FormatCandidates(c []KernelCandidate) string {
	parts := make([]string, 0, len(c))
	for _, x := range c {
		parts = append(parts, fmt.Sprintf("%s[%s dim=%d→%d kernel=%d patterns=%d cont=%d..%d canonical=%t targetOp=%t hidden=%d; %s]", x.Name, x.Kind, x.DomainDim, x.CodomainDim, x.KernelDim, x.SplitPatternCount, x.ResidualContinuousDimensionMin, x.ResidualContinuousDimensionMax, x.CanonicalKernelDerived, x.TargetContactOperatorDerived, x.HiddenDiscreteChoices, x.Obstruction))
	}
	return strings.Join(parts, "; ")
}

func FormatSplits(s []SplitAudit) string {
	parts := make([]string, 0, len(s))
	for _, x := range s {
		parts = append(parts, fmt.Sprintf("%s sectors=%v quotient=%d kernel=%d patterns=%d cont=%d..%d selected=%t; %s", x.Name, x.SectorDims, x.QuotientDim, x.KernelDim, x.PatternCount, x.ResidualContinuousDimensionMin, x.ResidualContinuousDimensionMax, x.CanonicalPatternSelected, x.Obstruction))
	}
	return strings.Join(parts, "; ")
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("matter=%d contact=%d kernel=%d GrDim=%d T3RPatterns=%d T3RCont=%d..%d jointPatterns=%d jointCont=%d..%d kernels=%d intertwiners=%d targetOps=%d beta=%d nullity=%d→%d", s.MatterDimension, s.ContactRows, s.RequiredKernelDim, s.UnconstrainedGrassmannDimension, s.T3RInvariantSplitPatterns, s.T3RInvariantResidualDimMin, s.T3RInvariantResidualDimMax, s.T3RChiralityInvariantSplitPatterns, s.T3RChiralityResidualDimMin, s.T3RChiralityResidualDimMax, s.CanonicalKernelCandidates, s.FullOperatorIntertwinersDerived, s.TargetContactOperatorsDerived, s.ContactBetaRowsAllowed, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func Join(items []string) string { return strings.Join(items, "; ") }

func count[T any](items []T, pred func(T) bool) int {
	n := 0
	for _, item := range items {
		if pred(item) {
			n++
		}
	}
	return n
}

func sum(xs []int) int {
	n := 0
	for _, x := range xs {
		n += x
	}
	return n
}
