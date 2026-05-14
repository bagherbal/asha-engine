// Package su2lgauge turns the charge-level SU(2)_L audit into explicit finite
// doublet generators.
//
// Gate 23 proved that the right-singlet/conjugate table plus the scalar doublet
// charge reconstructs the left doublet hypercharges.  This package builds the
// next object: an explicit finite weak-isospin action on the derived left
// doublet space.  It does not claim a continuum gauge field or a Yukawa matrix;
// it only verifies that the charge-level doublets really support the standard
// su(2) ladder algebra.
package su2lgauge

import (
	"sync"

	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/matter/su2l"
)

type DoubletKind string

const (
	QuarkDoublet  DoubletKind = "Q_L"
	LeptonDoublet DoubletKind = "L_L"
)

type LeftDoubletState struct {
	Index       int
	Name        string
	Kind        DoubletKind
	Color       int // 0 for leptons, 1..3 for quarks.
	T3          float64
	Hypercharge float64
	ElectricQ   float64
}

type MultipletSummary struct {
	Kind        DoubletKind
	Count       int
	Hypercharge float64
	UpCount     int
	DownCount   int
}

type Analysis struct {
	Audit su2l.Analysis

	States []LeftDoubletState

	Dimension int
	T3        linear.Matrix
	TPlus     linear.Matrix
	TMinus    linear.Matrix
	Y         linear.Matrix
	Q         linear.Matrix

	CommutatorT3TPlusNorm       float64
	CommutatorT3TMinusNorm      float64
	CommutatorTPlusTMinusNorm   float64
	CommutesWithHyperchargeNorm float64
	GellMannNishijimaNorm       float64

	QuarkColorDiagonal bool
	LeptonDoubletFound bool
	Multiplets         []MultipletSummary

	NonabelianSU2LGeneratorsDerived bool
	ContinuumGaugeFieldDerived      bool
	YukawaIntertwinerDerived        bool
	RemainingUnknowns               []string
}

var (
	su2lgaugeDefaultOnce  sync.Once
	su2lgaugeDefaultValue Analysis
	su2lgaugeDefaultErr   error
)

func BuildDefault() (Analysis, error) {
	su2lgaugeDefaultOnce.Do(func() {
		su2lgaugeDefaultValue, su2lgaugeDefaultErr = buildSu2lgaugeDefaultUncached()
	})
	return su2lgaugeDefaultValue, su2lgaugeDefaultErr
}

func buildSu2lgaugeDefaultUncached() (Analysis, error) {
	a, err := su2l.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(a, 1e-10)
}

func Build(a su2l.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !a.ChargeLevelSU2LDoubletsDerived || !a.Standard.MatchesStandardOrientation {
		return Analysis{}, fmt.Errorf("Gate 24 requires the Gate 23 standard-orientation doublet charges")
	}

	states := buildStandardStates(a.Standard)
	dim := len(states)
	t3 := linear.NewMatrix(dim, dim)
	yp := linear.NewMatrix(dim, dim)
	q := linear.NewMatrix(dim, dim)
	tp := linear.NewMatrix(dim, dim)
	tm := linear.NewMatrix(dim, dim)

	for _, s := range states {
		t3.Set(s.Index, s.Index, s.T3)
		yp.Set(s.Index, s.Index, s.Hypercharge)
		q.Set(s.Index, s.Index, s.ElectricQ)
	}

	// Build raising/lowering maps.  A +1/2 state is paired with the matching
	// -1/2 state of the same doublet kind and color.
	for _, up := range states {
		if math.Abs(up.T3-0.5) > eps {
			continue
		}
		downIdx := -1
		for _, down := range states {
			if math.Abs(down.T3+0.5) <= eps && up.Kind == down.Kind && up.Color == down.Color && math.Abs(up.Hypercharge-down.Hypercharge) <= eps {
				downIdx = down.Index
				break
			}
		}
		if downIdx < 0 {
			return Analysis{}, fmt.Errorf("could not find down partner for %s", up.Name)
		}
		tp.Set(up.Index, downIdx, 1) // T+ |down> = |up>
		tm.Set(downIdx, up.Index, 1) // T- |up> = |down>
	}

	c3p, _ := linear.Commutator(t3, tp)
	c3pMinusTp, _ := c3p.Sub(tp)
	c3m, _ := linear.Commutator(t3, tm)
	c3mPlusTm, _ := c3m.Add(tm)
	cpm, _ := linear.Commutator(tp, tm)
	twoT3 := t3.Scale(2)
	cpmMinus, _ := cpm.Sub(twoT3)
	cYp, _ := linear.Commutator(yp, tp)
	cYm, _ := linear.Commutator(yp, tm)
	cYBoth, _ := cYp.Add(cYm)

	t3PlusY, _ := t3.Add(yp)
	qResidual, _ := t3PlusY.Sub(q)

	summaries := summarize(states, eps)
	return Analysis{
		Audit:                           a,
		States:                          states,
		Dimension:                       dim,
		T3:                              t3,
		TPlus:                           tp,
		TMinus:                          tm,
		Y:                               yp,
		Q:                               q,
		CommutatorT3TPlusNorm:           c3pMinusTp.FrobeniusNorm(),
		CommutatorT3TMinusNorm:          c3mPlusTm.FrobeniusNorm(),
		CommutatorTPlusTMinusNorm:       cpmMinus.FrobeniusNorm(),
		CommutesWithHyperchargeNorm:     cYBoth.FrobeniusNorm(),
		GellMannNishijimaNorm:           qResidual.FrobeniusNorm(),
		QuarkColorDiagonal:              checkColorDiagonal(states),
		LeptonDoubletFound:              checkLeptonDoublet(states),
		Multiplets:                      summaries,
		NonabelianSU2LGeneratorsDerived: true,
		ContinuumGaugeFieldDerived:      false,
		YukawaIntertwinerDerived:        false,
		RemainingUnknowns: []string{
			"U-14-CONJUGATION-CONVENTION: select particle vs conjugate orientation from finite reality/CPT rather than using the standard branch by audit",
			"U-13B-SU2L-FINITE-GEOMETRIC-ORIGIN: derive these doublet generators directly from the Boolean-compressed block connection, not only from the charge table",
			"U-07-YUKAWA: construct explicit gauge-compatible intertwiners using the derived SU(2)_L action",
		},
	}, nil
}

func buildStandardStates(o su2l.Orientation) []LeftDoubletState {
	states := make([]LeftDoubletState, 0, 8)
	add := func(name string, kind DoubletKind, color int, t3, y float64) {
		states = append(states, LeftDoubletState{
			Index:       len(states),
			Name:        name,
			Kind:        kind,
			Color:       color,
			T3:          t3,
			Hypercharge: y,
			ElectricQ:   t3 + y,
		})
	}
	for c := 1; c <= 3; c++ {
		add(fmt.Sprintf("u_L^%d", c), QuarkDoublet, c, +0.5, o.QuarkLeftHypercharge)
		add(fmt.Sprintf("d_L^%d", c), QuarkDoublet, c, -0.5, o.QuarkLeftHypercharge)
	}
	add("nu_L", LeptonDoublet, 0, +0.5, o.LeptonLeftHypercharge)
	add("e_L", LeptonDoublet, 0, -0.5, o.LeptonLeftHypercharge)
	return states
}

func checkColorDiagonal(states []LeftDoubletState) bool {
	counts := map[int]int{}
	for _, s := range states {
		if s.Kind == QuarkDoublet {
			counts[s.Color]++
		}
	}
	return counts[1] == 2 && counts[2] == 2 && counts[3] == 2 && len(counts) == 3
}

func checkLeptonDoublet(states []LeftDoubletState) bool {
	foundUp, foundDown := false, false
	for _, s := range states {
		if s.Kind != LeptonDoublet || s.Color != 0 {
			continue
		}
		if math.Abs(s.T3-0.5) < 1e-10 && math.Abs(s.Hypercharge+0.5) < 1e-10 {
			foundUp = true
		}
		if math.Abs(s.T3+0.5) < 1e-10 && math.Abs(s.Hypercharge+0.5) < 1e-10 {
			foundDown = true
		}
	}
	return foundUp && foundDown
}

func summarize(states []LeftDoubletState, eps float64) []MultipletSummary {
	type key struct {
		kind DoubletKind
		y    int
	}
	counts := map[key]*MultipletSummary{}
	for _, s := range states {
		k := key{kind: s.Kind, y: int(math.Round(s.Hypercharge * 6))}
		if _, ok := counts[k]; !ok {
			counts[k] = &MultipletSummary{Kind: s.Kind, Hypercharge: s.Hypercharge}
		}
		counts[k].Count++
		if math.Abs(s.T3-0.5) <= eps {
			counts[k].UpCount++
		}
		if math.Abs(s.T3+0.5) <= eps {
			counts[k].DownCount++
		}
	}
	out := make([]MultipletSummary, 0, len(counts))
	for _, v := range counts {
		out = append(out, *v)
	}
	sort.Slice(out, func(i, j int) bool { return string(out[i].Kind) < string(out[j].Kind) })
	return out
}

func FormatStates(states []LeftDoubletState) string {
	parts := make([]string, len(states))
	for i, s := range states {
		color := ""
		if s.Color > 0 {
			color = fmt.Sprintf(" color=%d", s.Color)
		}
		parts[i] = fmt.Sprintf("%s{%s%s,T3=%.1f,Y=%.6g,Q=%.6g}", s.Name, s.Kind, color, s.T3, s.Hypercharge, s.ElectricQ)
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatMultiplets(ms []MultipletSummary) string {
	parts := make([]string, len(ms))
	for i, m := range ms {
		parts[i] = fmt.Sprintf("%s: dim=%d, Y=%.6g, up=%d, down=%d", m.Kind, m.Count, m.Hypercharge, m.UpCount, m.DownCount)
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUnknowns(xs []string) string {
	yy := append([]string(nil), xs...)
	sort.Strings(yy)
	return "[" + strings.Join(yy, "; ") + "]"
}
