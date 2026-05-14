// Package u1orientation implements Gate 78: the chiral / orientational
// abelian source search.
//
// Gate 77 constructed a genuine non-factorized abelian source candidate: the
// Yukawa-incidence correlation between matter B-L and scalar/contact T_phi.
// The signed total cancelled between up/down and neutrino/electron branches.
//
// This gate tests whether natural finite orientation weights already present in
// the one-generation channel data can break that cancellation without arbitrary
// sign choices.  It audits weak-isospin orientation, right-handed T3_R,
// electric charge, scalar orientation, matter charge, and quark/lepton parity.
// A nonzero result from an arbitrary kind selector is deliberately reported as
// non-canonical rather than accepted as a source.
package u1orientation

import (
	"fmt"
	"math"
	"sort"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/u1nonfactor"
	"github.com/bagherbal/asha-engine/pkg/matter/su2lgauge"
	"github.com/bagherbal/asha-engine/pkg/matter/yukawaintertwiner"
)

type Probe struct {
	Name        string
	Canonical   bool
	Description string
	Signed      float64
	Absolute    float64
	Quadratic   float64
	Nonzero     bool
	Cancels     bool
}

type KindMoment struct {
	Kind     yukawaintertwiner.FermionKind
	Signed   float64
	Absolute float64
}

type Analysis struct {
	Previous u1nonfactor.Analysis

	Probes []Probe
	Kinds  []KindMoment

	NaturalProbeCount       int
	NaturalNonzeroSources   int
	ArbitraryNonzeroSources int
	BestNaturalAbsSigned    float64

	ChiralSourceDerived       bool
	OrientationSourceDerived  bool
	CanonicalSourceDerived    bool
	FullU1HessianDerived      bool
	PhysicalU1CouplingDerived bool
	HiddenObservedInputUsed   bool

	TruthStatement      string
	RecommendedNextGate string
	RemainingUnknowns   []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := u1nonfactor.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(prev u1nonfactor.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !prev.NonFactorizedActionDerived {
		return Analysis{}, fmt.Errorf("Gate 78 requires Gate 77 non-factorized source object")
	}

	y := prev.Yukawa
	probes := []Probe{
		probe(y, "left-T3", true, "weak-isospin orientation on the left doublet", func(ch yukawaintertwiner.Channel) float64 { return ch.Left.T3 }),
		probe(y, "right-T3R", true, "matter-side T3_R = Y_R - (B-L)/2", func(ch yukawaintertwiner.Channel) float64 { return ch.Right.Hypercharge - bMinusL(ch.Left)/2 }),
		probe(y, "electric-Q-left", true, "Gell-Mann--Nishijima charge Q=T3+Y on the left state", func(ch yukawaintertwiner.Channel) float64 { return ch.Left.ElectricQ }),
		probe(y, "right-hypercharge", true, "right-singlet hypercharge orientation", func(ch yukawaintertwiner.Channel) float64 { return ch.Right.Hypercharge }),
		probe(y, "scalar-Tphi", true, "scalar/contact branch orientation", func(ch yukawaintertwiner.Channel) float64 { return ch.Scalar.Hypercharge }),
		probe(y, "matter-BL", true, "matter B-L weighting", func(ch yukawaintertwiner.Channel) float64 { return bMinusL(ch.Left) }),
		probe(y, "quark-minus-lepton", true, "Pati-Salam 3+1 sector parity: quark=+1, lepton=-1", func(ch yukawaintertwiner.Channel) float64 {
			if ch.Left.Kind == su2lgauge.QuarkDoublet {
				return 1
			}
			return -1
		}),
		probe(y, "up-neutrino-vs-down-electron", true, "weak branch parity: T3-positive states versus T3-negative states", func(ch yukawaintertwiner.Channel) float64 {
			if ch.Left.T3 > 0 {
				return 1
			}
			return -1
		}),
		probe(y, "up-only-selector", false, "non-canonical selector that keeps only up-type channels; included as a fitting firewall", func(ch yukawaintertwiner.Channel) float64 {
			if ch.Right.Kind == yukawaintertwiner.UpType {
				return 1
			}
			return 0
		}),
	}

	naturalCount := 0
	naturalNonzero := 0
	arbitraryNonzero := 0
	bestNatural := 0.0
	for _, p := range probes {
		if p.Canonical {
			naturalCount++
			if p.Nonzero {
				naturalNonzero++
			}
			if math.Abs(p.Signed) > bestNatural {
				bestNatural = math.Abs(p.Signed)
			}
		} else if p.Nonzero {
			arbitraryNonzero++
		}
	}

	kindMap := map[yukawaintertwiner.FermionKind]*KindMoment{}
	for _, ch := range y.Channels {
		if _, ok := kindMap[ch.Right.Kind]; !ok {
			kindMap[ch.Right.Kind] = &KindMoment{Kind: ch.Right.Kind}
		}
		v := bMinusL(ch.Left) * ch.Scalar.Hypercharge
		for i := 0; i < ch.Scalar.Multiplicity; i++ {
			kindMap[ch.Right.Kind].Signed += v
			kindMap[ch.Right.Kind].Absolute += math.Abs(v)
		}
	}
	kinds := make([]KindMoment, 0, len(kindMap))
	for _, v := range kindMap {
		kinds = append(kinds, *v)
	}
	sort.Slice(kinds, func(i, j int) bool { return kinds[i].Kind < kinds[j].Kind })

	truth := "Natural chiral, charge, scalar-orientation, and Pati-Salam parity weights all preserve the signed cancellation of the B-L/contact-u1 source. A nonzero source can be produced only by a non-canonical selector such as up-only, which is rejected as fitting. Thus the current finite data still lacks a canonical orientational abelian kinetic source."
	return Analysis{
		Previous:                  prev,
		Probes:                    probes,
		Kinds:                     kinds,
		NaturalProbeCount:         naturalCount,
		NaturalNonzeroSources:     naturalNonzero,
		ArbitraryNonzeroSources:   arbitraryNonzero,
		BestNaturalAbsSigned:      bestNatural,
		ChiralSourceDerived:       false,
		OrientationSourceDerived:  false,
		CanonicalSourceDerived:    false,
		FullU1HessianDerived:      false,
		PhysicalU1CouplingDerived: false,
		HiddenObservedInputUsed:   false,
		TruthStatement:            truth,
		RecommendedNextGate:       "Gate 79 — Anomaly / Cancellation Ledger for Abelian Sources",
		RemainingUnknowns: []string{
			"U-20D3D1-ORIENTATION-SOURCE: no natural chiral or scalar orientation breaks the signed B-L/contact cancellation",
			"U-20D3D2-ANOMALY-INTERPRETATION: decide whether the cancellation is the finite shadow of anomaly cancellation or a stricter no-mixing theorem",
			"U-20D3D3-NONCANONICAL-SELECTOR-REJECTION: up-only/top-only selectors are nonzero but not derived",
			"U-20D3D4-U1-HESSIAN: physical U(1) coupling still needs a nonzero kinetic Hessian source",
		},
	}, nil
}

func probe(y yukawaintertwiner.Analysis, name string, canonical bool, description string, weight func(yukawaintertwiner.Channel) float64) Probe {
	signed, abs, quad := 0.0, 0.0, 0.0
	for _, ch := range y.Channels {
		base := bMinusL(ch.Left) * ch.Scalar.Hypercharge
		w := weight(ch)
		for i := 0; i < ch.Scalar.Multiplicity; i++ {
			v := base * w
			signed += v
			abs += math.Abs(v)
			quad += v * v
		}
	}
	nonzero := math.Abs(signed) > 1e-10
	return Probe{Name: name, Canonical: canonical, Description: description, Signed: signed, Absolute: abs, Quadratic: quad, Nonzero: nonzero, Cancels: !nonzero && abs > 1e-10}
}

func bMinusL(l su2lgauge.LeftDoubletState) float64 {
	if l.Kind == su2lgauge.QuarkDoublet {
		return 1.0 / 3.0
	}
	return -1.0
}
