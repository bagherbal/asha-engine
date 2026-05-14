// Package sourcemap searches for a canonical source tensor from the active
// Higgs/contact carrier into the protected three-dimensional generation carrier.
//
// Previous gates established two facts:
//  1. active Higgs/contact curvature is nonzero;
//  2. all natural raw projections into the protected generation carrier vanish.
//
// This package asks the next categorical question: can the current finite data
// select a nonzero map M : H_active -> H_generation? Such a map would be the
// missing source tensor required to contract active scalar curvature into a
// genuine 3x3 generation texture. If no such map is selected, the engine must
// keep CKM/PMNS and non-diagonal Yukawa claims open.
package sourcemap

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/matter/bfbridge"
	"github.com/bagherbal/asha-engine/pkg/matter/bfsource"
	"github.com/bagherbal/asha-engine/pkg/matter/generationbreak"
)

type CandidateKind string

const (
	CandidateConnectionCrossMap  CandidateKind = "connection cross-map GᵀAH"
	CandidateBFCurvatureCrossMap CandidateKind = "BF curvature cross-map GᵀFH"
	CandidateBFMixedSource       CandidateKind = "BF mixed source contraction"
	CandidateHiggsAnisotropy     CandidateKind = "Higgs/contact diagonal anisotropy"
	CandidateArbitraryLinearMap  CandidateKind = "arbitrary active-to-generation map space"
)

type Candidate struct {
	Name            string
	Kind            CandidateKind
	DomainDim       int
	CodomainDim     int
	SpaceDim        int
	Rank            int
	Norm            float64
	Canonical       bool
	Nonzero         bool
	ProducesTexture bool
	Detail          string
}

type Analysis struct {
	GenerationBreak generationbreak.Analysis
	Bridge          bfbridge.Analysis
	BFSource        bfsource.Analysis

	GenerationDimension int
	ActiveDimension     int
	MapSpaceDimension   int

	Candidates    []Candidate
	BestCandidate Candidate

	ExistingConnectionMapFound bool
	BFCurvatureMapFound        bool
	BFSourceMapFound           bool
	ArbitraryMapsExist         bool

	CanonicalSourceTensorFound bool
	NonDiagonalTextureFound    bool
	CouplingsDerived           bool
	CKMDerived                 bool
	PMNSDerived                bool

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
		gb, err := generationbreak.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		br, err := bfbridge.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		bs, err := bfsource.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(gb, br, bs, 1e-8)
	})
	return defaultValue, defaultErr
}

func Build(gb generationbreak.Analysis, br bfbridge.Analysis, bs bfsource.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-8
	}
	if gb.GenerationCarrierDimension != 3 {
		return Analysis{}, fmt.Errorf("generation carrier must be 3D, got %d", gb.GenerationCarrierDimension)
	}
	if br.ProtectedDimension != gb.GenerationCarrierDimension {
		return Analysis{}, fmt.Errorf("bridge protected dimension %d does not match generation carrier %d", br.ProtectedDimension, gb.GenerationCarrierDimension)
	}
	if br.ActiveDimension <= 0 {
		return Analysis{}, fmt.Errorf("active carrier dimension must be positive")
	}

	genDim := gb.GenerationCarrierDimension
	activeDim := br.ActiveDimension
	mapDim := genDim * activeDim

	candidates := []Candidate{
		{
			Name:            "existing compressed-connection cross maps",
			Kind:            CandidateConnectionCrossMap,
			DomainDim:       activeDim,
			CodomainDim:     genDim,
			SpaceDim:        mapDim,
			Rank:            br.CrossMapSpanRank,
			Norm:            br.MaxCrossMapNorm,
			Canonical:       true,
			Nonzero:         br.MaxCrossMapNorm > eps && br.CrossMapSpanRank > 0,
			ProducesTexture: br.CanonicalGenerationMixingFound,
			Detail:          "Uses the already-present finite maps GᵀAᵢH. These are the first natural active-to-generation tensors; they vanish in the current finite connection.",
		},
		{
			Name:            "finite Maurer-Cartan curvature cross maps",
			Kind:            CandidateBFCurvatureCrossMap,
			DomainDim:       activeDim,
			CodomainDim:     genDim,
			SpaceDim:        mapDim,
			Rank:            bs.MixedBFResponseRank,
			Norm:            bs.MixedBFMaxNorm,
			Canonical:       true,
			Nonzero:         bs.MixedBFMaxNorm > eps && bs.MixedBFResponseRank > 0,
			ProducesTexture: bs.MixedBFBridgeFound,
			Detail:          "Uses the curvature-side mixed response GᵀFH. It is the natural BF curvature bridge; it also vanishes in the current finite data.",
		},
		{
			Name:            "BF action mixed source contraction",
			Kind:            CandidateBFMixedSource,
			DomainDim:       activeDim,
			CodomainDim:     genDim,
			SpaceDim:        mapDim,
			Rank:            bs.MixedQuadratic.Rank,
			Norm:            bs.MixedQuadratic.Norm,
			Canonical:       true,
			Nonzero:         bs.MixedQuadratic.Rank > 0 && bs.MixedQuadratic.Norm > eps,
			ProducesTexture: bs.CanonicalTextureFound,
			Detail:          "Tests whether the action-level mixed quadratic source produces a 3x3 texture. It is zero here because the mixed BF response vanishes.",
		},
		{
			Name:            "Higgs/contact anisotropy diagonal spurion",
			Kind:            CandidateHiggsAnisotropy,
			DomainDim:       genDim,
			CodomainDim:     genDim,
			SpaceDim:        genDim,
			Rank:            diagonalRank(gb.BestCandidate.Eigenvalues, eps),
			Norm:            gb.HiggsAnisotropy,
			Canonical:       false,
			Nonzero:         gb.DiagonalSpurionFound,
			ProducesTexture: gb.DiagonalSpurionFound,
			Detail:          "Splits the three generation labels diagonally, but it is not an active-to-generation map and does not generate mixing.",
		},
		{
			Name:            "unconstrained source tensor space Hom(H_active,H_generation)",
			Kind:            CandidateArbitraryLinearMap,
			DomainDim:       activeDim,
			CodomainDim:     genDim,
			SpaceDim:        mapDim,
			Rank:            mapDim,
			Norm:            0,
			Canonical:       false,
			Nonzero:         true,
			ProducesTexture: false,
			Detail:          "Mathematically 12 arbitrary maps exist, but choosing one would be a fit unless selected by finite geometry, symmetry, or an action principle.",
		},
	}

	best := selectBest(candidates)
	connFound := candidates[0].Nonzero
	curvFound := candidates[1].Nonzero
	sourceFound := candidates[2].Nonzero
	canonical := (connFound || curvFound || sourceFound) && best.Canonical && best.ProducesTexture

	return Analysis{
		GenerationBreak:            gb,
		Bridge:                     br,
		BFSource:                   bs,
		GenerationDimension:        genDim,
		ActiveDimension:            activeDim,
		MapSpaceDimension:          mapDim,
		Candidates:                 candidates,
		BestCandidate:              best,
		ExistingConnectionMapFound: connFound,
		BFCurvatureMapFound:        curvFound,
		BFSourceMapFound:           sourceFound,
		ArbitraryMapsExist:         mapDim > 0,
		CanonicalSourceTensorFound: canonical,
		NonDiagonalTextureFound:    canonical,
		CouplingsDerived:           false,
		CKMDerived:                 false,
		PMNSDerived:                false,
		TruthStatement:             truth(connFound, curvFound, sourceFound, mapDim),
		RemainingUnknowns: []string{
			"U-17D-ACTIVE-GENERATION-MAP: derive a nonzero canonical map M:H_active→H_generation from finite geometry rather than choosing an arbitrary 3x4 tensor",
			"U-17E-SOURCE-TENSOR-ACTION: formulate an action whose stationary equation selects M and rejects the zero map",
			"U-16C-NONCOMMUTING-TEXTURES: obtain at least two non-commuting 3x3 texture operators before CKM/PMNS claims",
			"U-16D-YUKAWA-SCALE-BRIDGE: normalize any future texture spectra into physical Yukawa strengths without observed-mass fitting",
		},
	}, nil
}

func diagonalRank(xs []float64, eps float64) int {
	n := 0
	for _, x := range xs {
		if abs(x) > eps {
			n++
		}
	}
	return n
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func selectBest(cs []Candidate) Candidate {
	if len(cs) == 0 {
		return Candidate{}
	}
	best := cs[0]
	for _, c := range cs[1:] {
		if score(c) > score(best) {
			best = c
		}
	}
	return best
}

func score(c Candidate) int {
	s := 0
	if c.Nonzero {
		s += 10
	}
	if c.Canonical {
		s += 20
	}
	if c.ProducesTexture {
		s += 30
	}
	if c.Kind == CandidateArbitraryLinearMap {
		s -= 25
	}
	return s
}

func truth(conn, curv, source bool, mapDim int) string {
	switch {
	case conn || curv || source:
		return "A nonzero active-to-generation source tensor has appeared. This is a candidate bridge, but it still needs symmetry, action, and normalization tests before any CKM/PMNS claim."
	default:
		return fmt.Sprintf("No canonical active-to-generation map is selected by the current finite connection, curvature, or BF source data. The abstract Hom(H_active,H_generation) space has dimension %d, but choosing a tensor there would be fitting, not derivation.", mapDim)
	}
}

func FormatCandidate(c Candidate) string {
	flags := []string{}
	if c.Canonical {
		flags = append(flags, "canonical")
	} else {
		flags = append(flags, "noncanonical")
	}
	if c.Nonzero {
		flags = append(flags, "nonzero")
	} else {
		flags = append(flags, "zero")
	}
	if c.ProducesTexture {
		flags = append(flags, "texture")
	} else {
		flags = append(flags, "no-texture")
	}
	return fmt.Sprintf("%s: %dx%d spaceDim=%d rank=%d norm=%.6e [%s]", c.Name, c.CodomainDim, c.DomainDim, c.SpaceDim, c.Rank, c.Norm, strings.Join(flags, ","))
}
