// Package yukawaintertwiner audits the first explicit gauge-compatible Yukawa
// selection channels after the finite SU(2)_L doublet representation has been
// constructed.
//
// The package deliberately does not fit masses and does not assign numerical
// Yukawa couplings.  It constructs only the charge- and representation-level
// selection rule
//
//	Y_R = Y_L + Y_Φ
//
// using the left-doublet table from Gate 24 and the finite scalar charge bridge
// from Gate 20.  The result is a finite list of allowed intertwiners from
// left doublets times scalar branches into right singlets.
package yukawaintertwiner

import (
	"sync"

	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/matter/su2lgauge"
)

type FermionKind string

const (
	UpType       FermionKind = "up"
	DownType     FermionKind = "down"
	NeutrinoType FermionKind = "neutrino"
	ElectronType FermionKind = "electron"
)

type ScalarBranch struct {
	Name         string
	Hypercharge  float64
	Multiplicity int
}

type RightSinglet struct {
	Name        string
	Kind        FermionKind
	Color       int
	Hypercharge float64
}

type Channel struct {
	Name string

	Left   su2lgauge.LeftDoubletState
	Scalar ScalarBranch
	Right  RightSinglet

	HyperchargeResidual float64
	ColorPreserving     bool
	LeptonPreserving    bool
}

type ChannelSummary struct {
	Kind               FermionKind
	Count              int
	ScalarMultiplicity int
	TotalFiberEntries  int
}

type Analysis struct {
	SU2L su2lgauge.Analysis

	LeftDimension  int
	RightDimension int
	ScalarBranches []ScalarBranch
	RightStates    []RightSinglet

	Channels  []Channel
	Summaries []ChannelSummary

	HyperchargeResidualMax float64
	ColorPreserving        bool
	LeptonPreserving       bool

	MinimalChannelCount int
	FiberEntryCount     int

	UpChannels       int
	DownChannels     int
	NeutrinoChannels int
	ElectronChannels int

	ChargeCompatibleYukawaChannelsDerived  bool
	GaugeInvariantCouplingConstantsDerived bool
	MassMatrixDerived                      bool
	FlavorMixingDerived                    bool
	RemainingUnknowns                      []string
}

var (
	yukawaIntertwinerDefaultOnce  sync.Once
	yukawaIntertwinerDefaultValue Analysis
	yukawaIntertwinerDefaultErr   error
)

func BuildDefault() (Analysis, error) {
	yukawaIntertwinerDefaultOnce.Do(func() {
		yukawaIntertwinerDefaultValue, yukawaIntertwinerDefaultErr = buildYukawaIntertwinerDefaultUncached()
	})
	return yukawaIntertwinerDefaultValue, yukawaIntertwinerDefaultErr
}

func buildYukawaIntertwinerDefaultUncached() (Analysis, error) {
	a, err := su2lgauge.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(a, 1e-10)
}

func Build(a su2lgauge.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !a.NonabelianSU2LGeneratorsDerived {
		return Analysis{}, fmt.Errorf("Gate 25 requires Gate 24 SU(2)_L generator audit")
	}

	scalarBranches := []ScalarBranch{
		{Name: "Φ_+", Hypercharge: 0.5, Multiplicity: 2},
		{Name: "Φ_-", Hypercharge: -0.5, Multiplicity: 2},
	}
	right := buildRightSinglets()
	channels := make([]Channel, 0)
	for _, l := range a.States {
		for _, s := range scalarBranches {
			target := canon(l.Hypercharge+s.Hypercharge, eps)
			for _, r := range right {
				if !compatibleKind(l, r) {
					continue
				}
				if !sameColorOrLepton(l, r) {
					continue
				}
				if math.Abs(canon(r.Hypercharge, eps)-target) > eps {
					continue
				}
				ch := Channel{
					Name:                fmt.Sprintf("%s ⊗ %s → %s", l.Name, s.Name, r.Name),
					Left:                l,
					Scalar:              s,
					Right:               r,
					HyperchargeResidual: canon(r.Hypercharge-(l.Hypercharge+s.Hypercharge), eps),
					ColorPreserving:     colorPreserving(l, r),
					LeptonPreserving:    leptonPreserving(l, r),
				}
				channels = append(channels, ch)
			}
		}
	}

	maxResidual := 0.0
	colorOK := true
	leptonOK := true
	up, down, nu, electron := 0, 0, 0, 0
	for _, ch := range channels {
		if math.Abs(ch.HyperchargeResidual) > maxResidual {
			maxResidual = math.Abs(ch.HyperchargeResidual)
		}
		colorOK = colorOK && ch.ColorPreserving
		leptonOK = leptonOK && ch.LeptonPreserving
		switch ch.Right.Kind {
		case UpType:
			up++
		case DownType:
			down++
		case NeutrinoType:
			nu++
		case ElectronType:
			electron++
		}
	}

	summaries := summarize(channels)
	fiberEntries := 0
	for _, s := range summaries {
		fiberEntries += s.TotalFiberEntries
	}

	complete := up == 3 && down == 3 && nu == 1 && electron == 1 && maxResidual < eps && colorOK && leptonOK
	return Analysis{
		SU2L:                                   a,
		LeftDimension:                          len(a.States),
		RightDimension:                         len(right),
		ScalarBranches:                         scalarBranches,
		RightStates:                            right,
		Channels:                               channels,
		Summaries:                              summaries,
		HyperchargeResidualMax:                 maxResidual,
		ColorPreserving:                        colorOK,
		LeptonPreserving:                       leptonOK,
		MinimalChannelCount:                    len(channels),
		FiberEntryCount:                        fiberEntries,
		UpChannels:                             up,
		DownChannels:                           down,
		NeutrinoChannels:                       nu,
		ElectronChannels:                       electron,
		ChargeCompatibleYukawaChannelsDerived:  complete,
		GaugeInvariantCouplingConstantsDerived: false,
		MassMatrixDerived:                      false,
		FlavorMixingDerived:                    false,
		RemainingUnknowns: []string{
			"U-07A-YUKAWA-COUPLINGS: derive actual coupling strengths from finite dynamics rather than inserting parameters",
			"U-07B-FLAVOR-MIXING: extend the one-generation channel audit to triality/generation mixing",
			"U-13B-SU2L-FINITE-GEOMETRIC-ORIGIN: derive the SU(2)_L representation directly from Boolean/contact geometry",
			"U-14-CONJUGATION-CONVENTION: select particle vs conjugate orientation from finite reality/CPT structure",
		},
	}, nil
}

func buildRightSinglets() []RightSinglet {
	out := make([]RightSinglet, 0, 8)
	for c := 1; c <= 3; c++ {
		out = append(out, RightSinglet{Name: fmt.Sprintf("u_R^%d", c), Kind: UpType, Color: c, Hypercharge: 2.0 / 3.0})
		out = append(out, RightSinglet{Name: fmt.Sprintf("d_R^%d", c), Kind: DownType, Color: c, Hypercharge: -1.0 / 3.0})
	}
	out = append(out, RightSinglet{Name: "nu_R", Kind: NeutrinoType, Color: 0, Hypercharge: 0})
	out = append(out, RightSinglet{Name: "e_R", Kind: ElectronType, Color: 0, Hypercharge: -1})
	return out
}

func compatibleKind(l su2lgauge.LeftDoubletState, r RightSinglet) bool {
	if l.Kind == su2lgauge.QuarkDoublet {
		if l.T3 > 0 {
			return r.Kind == UpType
		}
		return r.Kind == DownType
	}
	if l.Kind == su2lgauge.LeptonDoublet {
		if l.T3 > 0 {
			return r.Kind == NeutrinoType
		}
		return r.Kind == ElectronType
	}
	return false
}

func sameColorOrLepton(l su2lgauge.LeftDoubletState, r RightSinglet) bool {
	if l.Kind == su2lgauge.QuarkDoublet {
		return r.Color == l.Color
	}
	return r.Color == 0
}

func colorPreserving(l su2lgauge.LeftDoubletState, r RightSinglet) bool {
	if l.Kind != su2lgauge.QuarkDoublet {
		return true
	}
	return r.Color == l.Color && r.Color > 0
}

func leptonPreserving(l su2lgauge.LeftDoubletState, r RightSinglet) bool {
	if l.Kind != su2lgauge.LeptonDoublet {
		return true
	}
	return r.Color == 0
}

func summarize(channels []Channel) []ChannelSummary {
	m := map[FermionKind]*ChannelSummary{}
	for _, ch := range channels {
		s, ok := m[ch.Right.Kind]
		if !ok {
			s = &ChannelSummary{Kind: ch.Right.Kind, ScalarMultiplicity: ch.Scalar.Multiplicity}
			m[ch.Right.Kind] = s
		}
		s.Count++
		s.TotalFiberEntries += ch.Scalar.Multiplicity
	}
	out := make([]ChannelSummary, 0, len(m))
	for _, v := range m {
		out = append(out, *v)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Kind < out[j].Kind })
	return out
}

func canon(x float64, eps float64) float64 {
	if math.Abs(x) < eps {
		return 0
	}
	return math.Round(x*1e12) / 1e12
}

func FormatScalarBranches(branches []ScalarBranch) string {
	xs := make([]string, 0, len(branches))
	for _, b := range branches {
		xs = append(xs, fmt.Sprintf("%s:Y=%.6g×%d", b.Name, b.Hypercharge, b.Multiplicity))
	}
	return "[" + strings.Join(xs, ", ") + "]"
}

func FormatChannels(channels []Channel) string {
	xs := make([]string, 0, len(channels))
	for _, ch := range channels {
		xs = append(xs, ch.Name)
	}
	sort.Strings(xs)
	return "[" + strings.Join(xs, "; ") + "]"
}

func FormatSummaries(summaries []ChannelSummary) string {
	xs := make([]string, 0, len(summaries))
	for _, s := range summaries {
		xs = append(xs, fmt.Sprintf("%s: channels=%d, scalar-fiber entries=%d", s.Kind, s.Count, s.TotalFiberEntries))
	}
	sort.Strings(xs)
	return "[" + strings.Join(xs, "; ") + "]"
}

func FormatUnknowns(unknowns []string) string {
	xs := append([]string(nil), unknowns...)
	sort.Strings(xs)
	return "[" + strings.Join(xs, "; ") + "]"
}
