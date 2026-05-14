// Package texture searches for the first mathematically honest generation-
// breaking Yukawa texture operator after the triality extension.
//
// Gate 26 established that charge rules allow 3x3 flavor texture blocks for
// each fermion kind.  This package asks a sharper question: does the finite data
// already select numerical 3x3 matrices?
//
// The answer is deliberately conservative. Exact triality symmetry can replicate
// the one-generation channels, but it cannot produce three distinct generations
// or a CKM/PMNS matrix by itself.  A fully triality-invariant real symmetric
// texture has only two parameters and a 1+2 eigenvalue pattern.  Therefore a
// genuine hierarchy requires an additional finite generation-breaking operator,
// not a fitted mass table.
package texture

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/matter/trialityyukawa"
	"github.com/bagherbal/asha-engine/pkg/matter/yukawaintertwiner"
)

type SymmetryClass string

const (
	GeneralRealMatrix      SymmetryClass = "general real 3x3"
	RealSymmetricMatrix    SymmetryClass = "real symmetric 3x3"
	CyclicInvariantMatrix  SymmetryClass = "C3-invariant symmetric"
	FullTrialityInvariant  SymmetryClass = "S3/triality-invariant symmetric"
	DiagonalBreakingMatrix SymmetryClass = "diagonal generation-breaking candidate"
)

type TextureSpace struct {
	Name             SymmetryClass
	Dimension        int
	EigenPattern     string
	CanBreakAllThree bool
	CanProduceMixing bool
	RequiresNewInput bool
}

type KindTextureSummary struct {
	Kind                  yukawaintertwiner.FermionKind
	AllowedFlavorEntries  int
	GeneralRealDim        int
	SymmetricDim          int
	TrialityInvariantDim  int
	TrialityEigenPattern  string
	FullHierarchySelected bool
	MixingSelected        bool
}

type Analysis struct {
	Triality trialityyukawa.Analysis

	GenerationDimension int
	FermionKinds        int

	TextureSpaces []TextureSpace
	KindSummaries []KindTextureSummary

	TrialityInvariantTextureDim int
	SymmetricTextureDim         int
	GeneralTextureDim           int

	ExactTrialitySelectsTexture     bool
	ExactTrialityCanBreakAllThree   bool
	GenerationBreakingOperatorFound bool
	CouplingsDerived                bool
	CKMDerived                      bool
	PMNSDerived                     bool

	CandidateBreakingOperatorName string
	CandidateBreakingOperatorDim  int
	CandidateEigenPattern         string
	CandidateIsCanonical          bool

	NoGoStatement     string
	RemainingUnknowns []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		defaultValue, defaultErr = buildDefaultUncached()
	})
	return defaultValue, defaultErr
}

func buildDefaultUncached() (Analysis, error) {
	t, err := trialityyukawa.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(t)
}

func Build(t trialityyukawa.Analysis) (Analysis, error) {
	if !t.GenerationMixingAllowedByCharges || t.GenerationCount != 3 {
		return Analysis{}, fmt.Errorf("Gate 28 requires Gate 26 three-generation Yukawa selection space")
	}

	spaces := DefaultTextureSpaces()
	summaries := makeKindSummaries(t)

	return Analysis{
		Triality:                        t,
		GenerationDimension:             t.GenerationCount,
		FermionKinds:                    len(summaries),
		TextureSpaces:                   spaces,
		KindSummaries:                   summaries,
		TrialityInvariantTextureDim:     2,
		SymmetricTextureDim:             6,
		GeneralTextureDim:               9,
		ExactTrialitySelectsTexture:     false,
		ExactTrialityCanBreakAllThree:   false,
		GenerationBreakingOperatorFound: false,
		CouplingsDerived:                false,
		CKMDerived:                      false,
		PMNSDerived:                     false,
		CandidateBreakingOperatorName:   "none selected by current finite gates",
		CandidateBreakingOperatorDim:    0,
		CandidateEigenPattern:           "not available; exact triality gives at most 1+2",
		CandidateIsCanonical:            false,
		NoGoStatement:                   "Exact triality symmetry can copy channels into three generations, but by itself it cannot select a full 3x3 Yukawa texture or three distinct generation eigenvalues.",
		RemainingUnknowns: []string{
			"U-16A-GENERATION-BREAKING-OPERATOR: find a finite operator not commuting with exact triality that is still compatible with contact/BF/Higgs geometry",
			"U-16B-TEXTURE-SELECTION: derive the entries of Y_u, Y_d, Y_nu, and Y_e from finite spectra or index data",
			"U-16C-CKM-PMNS: derive mixing matrices only after at least two non-commuting finite texture blocks exist",
			"U-16D-TRIALITY-SPURION: decide whether generation breaking comes from contact curvature, Higgs anisotropy, boundary conditions, or a new finite source operator",
		},
	}, nil
}

func DefaultTextureSpaces() []TextureSpace {
	return []TextureSpace{
		{
			Name:             GeneralRealMatrix,
			Dimension:        9,
			EigenPattern:     "unconstrained 3x3 map; can fit anything if inserted by hand",
			CanBreakAllThree: true,
			CanProduceMixing: true,
			RequiresNewInput: true,
		},
		{
			Name:             RealSymmetricMatrix,
			Dimension:        6,
			EigenPattern:     "three real eigenvalues possible, but entries are not selected by current finite data",
			CanBreakAllThree: true,
			CanProduceMixing: true,
			RequiresNewInput: true,
		},
		{
			Name:             CyclicInvariantMatrix,
			Dimension:        2,
			EigenPattern:     "1+2 degeneracy for real symmetric C3-invariant textures",
			CanBreakAllThree: false,
			CanProduceMixing: false,
			RequiresNewInput: false,
		},
		{
			Name:             FullTrialityInvariant,
			Dimension:        2,
			EigenPattern:     "singlet plus doublet: λ_s=a+2b, λ_d=a-b with multiplicity 2",
			CanBreakAllThree: false,
			CanProduceMixing: false,
			RequiresNewInput: false,
		},
		{
			Name:             DiagonalBreakingMatrix,
			Dimension:        3,
			EigenPattern:     "three diagonal generation weights possible, but no mixing unless another non-commuting operator appears",
			CanBreakAllThree: true,
			CanProduceMixing: false,
			RequiresNewInput: true,
		},
	}
}

func makeKindSummaries(t trialityyukawa.Analysis) []KindTextureSummary {
	out := make([]KindTextureSummary, 0, len(t.KindSummaries))
	for _, k := range t.KindSummaries {
		out = append(out, KindTextureSummary{
			Kind:                  k.Kind,
			AllowedFlavorEntries:  k.FlavorMatrixDim,
			GeneralRealDim:        9,
			SymmetricDim:          6,
			TrialityInvariantDim:  2,
			TrialityEigenPattern:  "1+2",
			FullHierarchySelected: false,
			MixingSelected:        false,
		})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Kind < out[j].Kind })
	return out
}

func TrialityInvariantEigenvalues(a, b float64) []float64 {
	return []float64{a + 2*b, a - b, a - b}
}

func TrialityDegeneracyResidual(a, b float64) float64 {
	eig := TrialityInvariantEigenvalues(a, b)
	return math.Abs(eig[1] - eig[2])
}

func FormatTextureSpaces(spaces []TextureSpace) string {
	xs := make([]string, 0, len(spaces))
	for _, s := range spaces {
		flags := []string{}
		if s.CanBreakAllThree {
			flags = append(flags, "breaks3")
		}
		if s.CanProduceMixing {
			flags = append(flags, "mixing")
		}
		if s.RequiresNewInput {
			flags = append(flags, "needs-new-input")
		}
		xs = append(xs, fmt.Sprintf("%s dim=%d pattern=%s flags=%s", s.Name, s.Dimension, s.EigenPattern, strings.Join(flags, ",")))
	}
	return "[" + strings.Join(xs, "; ") + "]"
}

func FormatKindSummaries(summaries []KindTextureSummary) string {
	xs := make([]string, 0, len(summaries))
	for _, s := range summaries {
		xs = append(xs, fmt.Sprintf("%s allowed=%d general=%d symmetric=%d trialityInvariant=%d pattern=%s selected=%v", s.Kind, s.AllowedFlavorEntries, s.GeneralRealDim, s.SymmetricDim, s.TrialityInvariantDim, s.TrialityEigenPattern, s.FullHierarchySelected))
	}
	return "[" + strings.Join(xs, "; ") + "]"
}

func FormatUnknowns(unknowns []string) string {
	return "[" + strings.Join(unknowns, "; ") + "]"
}
