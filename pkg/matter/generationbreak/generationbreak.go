// Package generationbreak searches for finite generation-breaking sources after
// the triality/Yukawa texture no-go.
//
// Gate 28 established a hard limitation: exact triality gives three copies and
// allows 3x3 texture spaces, but it does not select three distinct generation
// eigenvalues or mixing matrices. This package therefore asks a sharper and
// more physical question: which already-computed finite objects can act as a
// generation-breaking source without inserting observed masses?
//
// The result is deliberately conservative. The Higgs/contact anisotropy yields
// a natural diagonal spurion with three distinct weights once the protected
// contact sector is used as a 3D generation carrier. However, no canonical
// finite operator has yet been derived that both lives intrinsically on that
// 3D carrier and produces mixing. Therefore this gate exposes the first viable
// spurion and the remaining missing theorem, rather than pretending to derive
// CKM/PMNS or fermion masses.
package generationbreak

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/dynamics/higgspotential"
	"github.com/bagherbal/asha-engine/pkg/gauge/connection"
	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
	"github.com/bagherbal/asha-engine/pkg/matter/texture"
)

type SourceKind string

const (
	SourceHiggsAnisotropy   SourceKind = "Higgs/contact spectral anisotropy"
	SourceProtectedContact  SourceKind = "protected contact carrier"
	SourceSecondFundamental SourceKind = "projected-connection second fundamental curvature"
	SourceContactLeakage    SourceKind = "bare contact leakage / partial overlap"
	SourceBFCurvature       SourceKind = "finite BF boundary curvature"
)

type Candidate struct {
	Name           string
	Source         SourceKind
	Dimension      int
	Eigenvalues    []float64
	EigenPattern   string
	BreaksAllThree bool
	ProducesMixing bool
	Canonical      bool
	RequiresBridge bool
	Detail         string
}

type Analysis struct {
	Texture    texture.Analysis
	Potential  higgspotential.Analysis
	Connection connection.Analysis
	Contact    contact.Space

	GenerationCarrierDimension int
	ProtectedContactDimension  int

	HiggsAnisotropy        float64
	SecondFundamentalSize  float64
	ContactLeakageNorm     float64
	PartialOverlapSpectrum []float64

	Candidates    []Candidate
	BestCandidate Candidate

	DiagonalSpurionFound   bool
	MixingOperatorFound    bool
	CanonicalOperatorFound bool
	CouplingsDerived       bool
	CKMDerived             bool
	PMNSDerived            bool
	FullHierarchySelected  bool

	TruthStatement    string
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
	tx, err := texture.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	hp, err := higgspotential.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	cn, err := connection.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	k, err := contact.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(tx, hp, cn, k, 1e-8)
}

func Build(tx texture.Analysis, hp higgspotential.Analysis, cn connection.Analysis, k contact.Space, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-8
	}
	if tx.GenerationDimension != 3 {
		return Analysis{}, fmt.Errorf("generation-breaking search requires three triality sectors, got %d", tx.GenerationDimension)
	}
	if hp.ProtectedContactDimension != tx.GenerationDimension {
		return Analysis{}, fmt.Errorf("protected contact dimension %d does not match generation dimension %d", hp.ProtectedContactDimension, tx.GenerationDimension)
	}
	if len(hp.ActiveContactSpectrum) < 2 {
		return Analysis{}, fmt.Errorf("Higgs/contact spectrum has no anisotropy data")
	}

	partial := partialOverlap(k.OverlapEigenvalues, eps)
	higgsSpurion := higgsDiagonalSpurion(hp, tx.GenerationDimension, eps)
	protected := Candidate{
		Name:           "protected K-residual generation carrier",
		Source:         SourceProtectedContact,
		Dimension:      hp.ProtectedContactDimension,
		Eigenvalues:    zeros(hp.ProtectedContactDimension),
		EigenPattern:   fmt.Sprintf("carrier dimension=%d; no splitting operator yet", hp.ProtectedContactDimension),
		BreaksAllThree: false,
		ProducesMixing: false,
		Canonical:      true,
		RequiresBridge: true,
		Detail:         "The finite Higgs/contact sector leaves exactly three protected contact directions; this is a natural carrier for generation labels, not yet a texture operator.",
	}
	second := Candidate{
		Name:           "second-fundamental curvature scalar source",
		Source:         SourceSecondFundamental,
		Dimension:      1,
		Eigenvalues:    []float64{cn.MaxSecondFundamentalNorm},
		EigenPattern:   "nonzero scalar curvature source; not a 3x3 generation operator",
		BreaksAllThree: false,
		ProducesMixing: false,
		Canonical:      true,
		RequiresBridge: true,
		Detail:         "The projected-connection curvature is nonzero and physically important, but the current implementation has not restricted it to a canonical 3D generation carrier.",
	}
	leakage := Candidate{
		Name:           "raw contact partial-overlap leakage spectrum",
		Source:         SourceContactLeakage,
		Dimension:      len(partial),
		Eigenvalues:    partial,
		EigenPattern:   fmt.Sprintf("%d partial-overlap modes; not canonically reduced to three generations", len(partial)),
		BreaksAllThree: false,
		ProducesMixing: false,
		Canonical:      false,
		RequiresBridge: true,
		Detail:         "The contact leakage spectrum is real finite data, but selecting three of its partial-overlap modes would currently be arbitrary.",
	}
	bf := Candidate{
		Name:           "finite BF boundary-curvature generator",
		Source:         SourceBFCurvature,
		Dimension:      0,
		Eigenvalues:    nil,
		EigenPattern:   "not implemented in current engine",
		BreaksAllThree: false,
		ProducesMixing: false,
		Canonical:      false,
		RequiresBridge: true,
		Detail:         "The framework predicts this as a serious source, but the engine has not yet implemented the finite BF curvature operator on the generation carrier.",
	}

	candidates := []Candidate{protected, higgsSpurion, second, leakage, bf}
	best := higgsSpurion
	diagonalFound := best.BreaksAllThree && !best.ProducesMixing

	return Analysis{
		Texture:                    tx,
		Potential:                  hp,
		Connection:                 cn,
		Contact:                    k,
		GenerationCarrierDimension: tx.GenerationDimension,
		ProtectedContactDimension:  hp.ProtectedContactDimension,
		HiggsAnisotropy:            hp.SpectralAnisotropy,
		SecondFundamentalSize:      cn.MaxSecondFundamentalNorm,
		ContactLeakageNorm:         k.BareLeakageNorm(),
		PartialOverlapSpectrum:     partial,
		Candidates:                 candidates,
		BestCandidate:              best,
		DiagonalSpurionFound:       diagonalFound,
		MixingOperatorFound:        false,
		CanonicalOperatorFound:     false,
		CouplingsDerived:           false,
		CKMDerived:                 false,
		PMNSDerived:                false,
		FullHierarchySelected:      false,
		TruthStatement:             "The finite geometry now exposes a natural diagonal generation-breaking spurion from Higgs/contact anisotropy, but not yet a canonical 3x3 operator with mixing or numerical Yukawa couplings.",
		RemainingUnknowns: []string{
			"U-16A-GENERATION-CARRIER-MAP: construct a canonical map from the three protected contact directions to the three triality sectors",
			"U-16B-CURVATURE-ON-GENERATIONS: restrict second-fundamental/BF curvature to the 3D generation carrier",
			"U-16C-NONCOMMUTING-TEXTURES: derive at least two non-commuting finite generation operators before CKM/PMNS can be claimed",
			"U-16D-YUKAWA-SCALE-BRIDGE: convert dimensionless finite texture spectra into physical Yukawa couplings without fitting masses",
		},
	}, nil
}

func higgsDiagonalSpurion(hp higgspotential.Analysis, generationDim int, eps float64) Candidate {
	values := append([]float64(nil), hp.ActiveContactSpectrum...)
	sort.Slice(values, func(i, j int) bool { return values[i] > values[j] })
	high := values[0]
	low := values[len(values)-1]
	mid := hp.MeanActiveEigenvalue
	eigen := []float64{high, mid, low}
	distinct := distinctCount(eigen, eps)
	return Candidate{
		Name:           "Higgs/contact anisotropy diagonal spurion",
		Source:         SourceHiggsAnisotropy,
		Dimension:      generationDim,
		Eigenvalues:    eigen,
		EigenPattern:   fmt.Sprintf("%d distinct diagonal weights from {λmax, mean, λmin}", distinct),
		BreaksAllThree: distinct == generationDim,
		ProducesMixing: false,
		Canonical:      false,
		RequiresBridge: true,
		Detail:         "This is the first finite 1+1+1 generation-breaking seed. It is diagonal and scale-free, so it cannot by itself produce CKM/PMNS mixing.",
	}
}

func partialOverlap(values []float64, eps float64) []float64 {
	out := make([]float64, 0)
	for _, v := range values {
		if math.Abs(v-1) <= eps || math.Abs(v) <= eps {
			continue
		}
		out = append(out, v)
	}
	sort.Slice(out, func(i, j int) bool { return out[i] > out[j] })
	return out
}

func zeros(n int) []float64 {
	out := make([]float64, n)
	return out
}

func distinctCount(values []float64, eps float64) int {
	if len(values) == 0 {
		return 0
	}
	sorted := append([]float64(nil), values...)
	sort.Float64s(sorted)
	count := 1
	last := sorted[0]
	for _, v := range sorted[1:] {
		if math.Abs(v-last) > eps {
			count++
			last = v
		}
	}
	return count
}

func FormatCandidate(c Candidate) string {
	return fmt.Sprintf("%s source=%s dim=%d eigen=%s pattern=%s breaks3=%t mixing=%t canonical=%t", c.Name, c.Source, c.Dimension, FormatFloatSlice(c.Eigenvalues), c.EigenPattern, c.BreaksAllThree, c.ProducesMixing, c.Canonical)
}

func FormatCandidates(candidates []Candidate) string {
	xs := make([]string, 0, len(candidates))
	for _, c := range candidates {
		xs = append(xs, FormatCandidate(c))
	}
	return "[" + strings.Join(xs, "; ") + "]"
}

func FormatFloatSlice(values []float64) string {
	if len(values) == 0 {
		return "[]"
	}
	xs := make([]string, len(values))
	for i, v := range values {
		xs[i] = fmt.Sprintf("%.10g", v)
	}
	return "[" + strings.Join(xs, ", ") + "]"
}
