// Package trialityyukawa lifts the one-generation Yukawa/intertwiner audit
// into a three-generation triality layer.
//
// The package is deliberately conservative.  Triality is treated here as the
// finite reason to carry three copies of the one-generation channel pattern.
// It does not by itself assign numerical Yukawa couplings, a CKM matrix, a PMNS
// matrix, or a hierarchy.  The gate therefore distinguishes three facts:
//
//  1. the one-generation gauge-compatible channels replicate into three
//     triality sectors;
//  2. charge rules allow arbitrary generation mixing inside each fermion kind;
//  3. the present finite data do not yet select the actual 3x3 texture.
package trialityyukawa

import (
	"sync"

	"fmt"
	"sort"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/matter/yukawaintertwiner"
)

type TrialitySector struct {
	Index int
	Name  string
	Role  string
}

type GenerationChannel struct {
	Name string

	LeftGeneration  TrialitySector
	RightGeneration TrialitySector
	Base            yukawaintertwiner.Channel

	GenerationDiagonal bool
}

type KindSummary struct {
	Kind             yukawaintertwiner.FermionKind
	BaseChannels     int
	DiagonalChannels int
	FullMixingMaps   int
	FlavorMatrixDim  int
}

type Analysis struct {
	OneGeneration yukawaintertwiner.Analysis

	TrialitySectors []TrialitySector
	GenerationCount int

	OneGenerationChannels int
	DiagonalChannels      []GenerationChannel
	FullMixingChannels    []GenerationChannel

	DiagonalChannelCount int
	FullMixingMapCount   int

	OneGenerationFiberEntries int
	DiagonalFiberEntries      int
	FullMixingFiberEntries    int

	KindSummaries []KindSummary

	TrialityCopiesChannelPattern     bool
	GenerationMixingAllowedByCharges bool
	TextureSelectedByFiniteData      bool
	CouplingsDerived                 bool
	CKMDerived                       bool
	PMNSDerived                      bool

	RemainingUnknowns []string
}

var (
	trialityYukawaDefaultOnce  sync.Once
	trialityYukawaDefaultValue Analysis
	trialityYukawaDefaultErr   error
)

func BuildDefault() (Analysis, error) {
	trialityYukawaDefaultOnce.Do(func() {
		trialityYukawaDefaultValue, trialityYukawaDefaultErr = buildTrialityYukawaDefaultUncached()
	})
	return trialityYukawaDefaultValue, trialityYukawaDefaultErr
}

func buildTrialityYukawaDefaultUncached() (Analysis, error) {
	one, err := yukawaintertwiner.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(one)
}

func Build(one yukawaintertwiner.Analysis) (Analysis, error) {
	if !one.ChargeCompatibleYukawaChannelsDerived {
		return Analysis{}, fmt.Errorf("Gate 26 requires Gate 25 gauge-compatible one-generation Yukawa channels")
	}
	sectors := DefaultTrialitySectors()

	diagonal := make([]GenerationChannel, 0, len(one.Channels)*len(sectors))
	full := make([]GenerationChannel, 0, len(one.Channels)*len(sectors)*len(sectors))
	for _, base := range one.Channels {
		for _, gl := range sectors {
			for _, gr := range sectors {
				ch := GenerationChannel{
					Name:               fmt.Sprintf("%s[%s→%s]", base.Name, gl.Name, gr.Name),
					LeftGeneration:     gl,
					RightGeneration:    gr,
					Base:               base,
					GenerationDiagonal: gl.Index == gr.Index,
				}
				full = append(full, ch)
				if gl.Index == gr.Index {
					diagonal = append(diagonal, ch)
				}
			}
		}
	}

	diagonalFiber := 0
	for _, ch := range diagonal {
		diagonalFiber += ch.Base.Scalar.Multiplicity
	}
	fullFiber := 0
	for _, ch := range full {
		fullFiber += ch.Base.Scalar.Multiplicity
	}

	summaries := summarize(one, sectors)
	return Analysis{
		OneGeneration:                    one,
		TrialitySectors:                  sectors,
		GenerationCount:                  len(sectors),
		OneGenerationChannels:            len(one.Channels),
		DiagonalChannels:                 diagonal,
		FullMixingChannels:               full,
		DiagonalChannelCount:             len(diagonal),
		FullMixingMapCount:               len(full),
		OneGenerationFiberEntries:        one.FiberEntryCount,
		DiagonalFiberEntries:             diagonalFiber,
		FullMixingFiberEntries:           fullFiber,
		KindSummaries:                    summaries,
		TrialityCopiesChannelPattern:     len(sectors) == 3 && len(diagonal) == 3*len(one.Channels),
		GenerationMixingAllowedByCharges: len(full) == len(one.Channels)*len(sectors)*len(sectors),
		TextureSelectedByFiniteData:      false,
		CouplingsDerived:                 false,
		CKMDerived:                       false,
		PMNSDerived:                      false,
		RemainingUnknowns: []string{
			"U-07A-YUKAWA-COUPLINGS: derive numerical coupling strengths from finite dynamics rather than assigning matrices",
			"U-07B-FLAVOR-TEXTURE: derive or constrain the 3x3 generation texture for each fermion kind",
			"U-07C-CKM-PMNS: derive quark and lepton mixing matrices only after finite Yukawa textures exist",
			"U-15-TRIALITY-ACTION: implement the actual SO(8) triality automorphisms on the Fock/contact representation, not only the three-copy bookkeeping layer",
			"U-16-GENERATION-BREAKING: identify the finite operator that breaks exact generation degeneracy without fitted masses",
		},
	}, nil
}

func DefaultTrialitySectors() []TrialitySector {
	return []TrialitySector{
		{Index: 0, Name: "G1/8_v", Role: "vector triality sector"},
		{Index: 1, Name: "G2/8_s+", Role: "positive spinor triality sector"},
		{Index: 2, Name: "G3/8_s-", Role: "negative spinor triality sector"},
	}
}

func summarize(one yukawaintertwiner.Analysis, sectors []TrialitySector) []KindSummary {
	base := map[yukawaintertwiner.FermionKind]int{}
	for _, ch := range one.Channels {
		base[ch.Right.Kind]++
	}
	out := make([]KindSummary, 0, len(base))
	g := len(sectors)
	for kind, count := range base {
		out = append(out, KindSummary{
			Kind:             kind,
			BaseChannels:     count,
			DiagonalChannels: count * g,
			FullMixingMaps:   count * g * g,
			FlavorMatrixDim:  g * g,
		})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Kind < out[j].Kind })
	return out
}

func FormatSectors(sectors []TrialitySector) string {
	xs := make([]string, 0, len(sectors))
	for _, s := range sectors {
		xs = append(xs, fmt.Sprintf("%s:%s", s.Name, s.Role))
	}
	return "[" + strings.Join(xs, "; ") + "]"
}

func FormatKindSummaries(summaries []KindSummary) string {
	xs := make([]string, 0, len(summaries))
	for _, s := range summaries {
		xs = append(xs, fmt.Sprintf("%s base=%d diagonal=%d fullMixing=%d texture=%dx%d", s.Kind, s.BaseChannels, s.DiagonalChannels, s.FullMixingMaps, 3, 3))
	}
	return "[" + strings.Join(xs, "; ") + "]"
}

func FormatChannelSample(channels []GenerationChannel, limit int) string {
	if limit <= 0 || limit > len(channels) {
		limit = len(channels)
	}
	xs := make([]string, 0, limit)
	for i := 0; i < limit; i++ {
		xs = append(xs, channels[i].Name)
	}
	suffix := ""
	if limit < len(channels) {
		suffix = fmt.Sprintf("; ... +%d more", len(channels)-limit)
	}
	return "[" + strings.Join(xs, "; ") + suffix + "]"
}

func FormatUnknowns(unknowns []string) string {
	xs := append([]string(nil), unknowns...)
	sort.Strings(xs)
	return "[" + strings.Join(xs, "; ") + "]"
}
