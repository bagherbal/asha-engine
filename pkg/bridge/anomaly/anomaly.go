// Package anomaly implements Gate 79: the anomaly / cancellation ledger for
// abelian sources.
//
// Gates 76-78 found a repeated signed cancellation in attempts to construct a
// B-L/contact-u1 kinetic source.  This package tests whether that cancellation
// is merely accidental, a finite shadow of familiar chiral anomaly
// cancellations, or a stronger no-mixing theorem in the present dual-carrier
// architecture.
//
// The gate deliberately distinguishes three statements:
//
//  1. The one-generation charge table is anomaly-balanced when right-handed
//     states are written as left-handed conjugates.
//  2. The Yukawa-incidence B-L/T_phi source cancels by up/down and
//     neutrino/electron pair balance.
//  3. A vanishing anomaly ledger is not automatically a kinetic-mixing
//     Hessian; it explains why charge-balanced sources cancel, but it does not
//     derive a physical U(1) coupling.
package anomaly

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/u1orientation"
	"github.com/bagherbal/asha-engine/pkg/matter/su2lgauge"
	"github.com/bagherbal/asha-engine/pkg/matter/yukawaintertwiner"
)

type WeylState struct {
	Name        string
	Kind        string
	Color       int
	SU2Doublet  bool
	SU3Charged  bool
	Hypercharge float64
	BMinusL     float64
}

type Moment struct {
	Name        string
	Value       float64
	Cancels     bool
	Description string
}

type Analysis struct {
	Previous u1orientation.Analysis
	Yukawa   yukawaintertwiner.Analysis

	States []WeylState

	HyperchargeMoments []Moment
	BMinusLMoments     []Moment
	MixedMoments       []Moment

	YAnomalyCancels       bool
	BMinusLAnomalyCancels bool
	MixedAbelianCancels   bool

	YukawaIncidenceCancellation bool
	AnomalyShadowSupported      bool
	StricterNoMixingTheorem     bool
	KineticMixingDerived        bool
	PhysicalU1CouplingDerived   bool
	HiddenObservedInputUsed     bool

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
		prev, err := u1orientation.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(prev u1orientation.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if prev.Previous.Yukawa.MinimalChannelCount == 0 {
		return Analysis{}, fmt.Errorf("Gate 79 requires Gate 77/78 Yukawa-incidence data")
	}
	y := prev.Previous.Yukawa
	states := buildLeftHandedWeylTable(y)

	hy := []Moment{
		moment("Tr(Y)", sum(states, func(s WeylState) float64 { return s.Hypercharge }), eps, "mixed gravitational-U(1)_Y anomaly ledger"),
		moment("Tr(Y^3)", sum(states, func(s WeylState) float64 { return cube(s.Hypercharge) }), eps, "cubic U(1)_Y anomaly ledger"),
		moment("SU(2)^2·Y", su2sq(states, func(s WeylState) float64 { return s.Hypercharge }), eps, "weak doublet anomaly ledger; common Dynkin factor suppressed"),
		moment("SU(3)^2·Y", su3sq(states, func(s WeylState) float64 { return s.Hypercharge }), eps, "color anomaly ledger; common Dynkin factor suppressed"),
	}
	bl := []Moment{
		moment("Tr(B-L)", sum(states, func(s WeylState) float64 { return s.BMinusL }), eps, "mixed gravitational-(B-L) anomaly ledger"),
		moment("Tr((B-L)^3)", sum(states, func(s WeylState) float64 { return cube(s.BMinusL) }), eps, "cubic B-L anomaly ledger with nu_R included"),
		moment("SU(2)^2·(B-L)", su2sq(states, func(s WeylState) float64 { return s.BMinusL }), eps, "weak doublet B-L anomaly ledger"),
		moment("SU(3)^2·(B-L)", su3sq(states, func(s WeylState) float64 { return s.BMinusL }), eps, "color B-L anomaly ledger"),
	}
	mixed := []Moment{
		moment("Tr(Y^2(B-L))", sum(states, func(s WeylState) float64 { return s.Hypercharge * s.Hypercharge * s.BMinusL }), eps, "mixed abelian anomaly ledger"),
		moment("Tr(Y(B-L)^2)", sum(states, func(s WeylState) float64 { return s.Hypercharge * s.BMinusL * s.BMinusL }), eps, "mixed abelian anomaly ledger"),
	}

	yCancel := allCancel(hy)
	blCancel := allCancel(bl)
	mixedCancel := allCancel(mixed)
	yukawaCancel := prev.Previous.TotalSignedCancellation && prev.NaturalNonzeroSources == 0
	shadow := yCancel && blCancel && mixedCancel && yukawaCancel

	truth := "The repeated B-L/contact-u1 source cancellation is not random: the same one-generation finite charge table is anomaly-balanced for Y, B-L, and mixed abelian moments when right-handed states are written as left-handed conjugates.  This supports reading the Yukawa-incidence cancellation as an anomaly-like finite charge-balance shadow.  It is not yet a kinetic-mixing theorem, because anomaly cancellation is a chiral consistency ledger, while a physical U(1) coupling still requires a nonzero kinetic Hessian/source action."
	return Analysis{
		Previous:                    prev,
		Yukawa:                      y,
		States:                      states,
		HyperchargeMoments:          hy,
		BMinusLMoments:              bl,
		MixedMoments:                mixed,
		YAnomalyCancels:             yCancel,
		BMinusLAnomalyCancels:       blCancel,
		MixedAbelianCancels:         mixedCancel,
		YukawaIncidenceCancellation: yukawaCancel,
		AnomalyShadowSupported:      shadow,
		StricterNoMixingTheorem:     false,
		KineticMixingDerived:        false,
		PhysicalU1CouplingDerived:   false,
		HiddenObservedInputUsed:     false,
		TruthStatement:              truth,
		RecommendedNextGate:         "Gate 80 — Anomaly-Constrained U(1) Kinetic Hessian Search",
		RemainingUnknowns: []string{
			"U-20D3E1-ANOMALY-VS-HESSIAN: anomaly cancellation explains charge-balance cancellations but does not supply a kinetic Hessian",
			"U-20D3E2-NONZERO-U1-SOURCE: derive a non-anomalous, non-factorized source term if physical U(1) kinetic mixing exists",
			"U-20D3E3-DUAL-CARRIER-ABELIAN-ACTION: construct the finite action that couples B-L/contact-u1 without violating the cancellation ledger",
			"U-20D3E4-PHYSICAL-COUPLING: g_Y and alpha remain blocked until the kinetic Hessian and RG boundary are derived",
		},
	}, nil
}

func buildLeftHandedWeylTable(y yukawaintertwiner.Analysis) []WeylState {
	states := make([]WeylState, 0, y.LeftDimension+y.RightDimension)
	for _, l := range y.SU2L.States {
		bl := -1.0
		kind := "lepton-doublet"
		su3 := false
		if l.Kind == su2lgauge.QuarkDoublet {
			bl = 1.0 / 3.0
			kind = "quark-doublet"
			su3 = true
		}
		states = append(states, WeylState{Name: l.Name, Kind: kind, Color: l.Color, SU2Doublet: true, SU3Charged: su3, Hypercharge: l.Hypercharge, BMinusL: bl})
	}
	for _, r := range y.RightStates {
		bl := rightBMinusL(r)
		// Chiral anomaly ledgers are written entirely with left-handed Weyl fields.
		// A right-handed field is therefore converted to its left-handed conjugate,
		// which flips abelian charges.  The SU(3) Dynkin index is unchanged by
		// conjugation, so SU3Charged remains true for colored conjugates.
		states = append(states, WeylState{Name: r.Name + "^c", Kind: string(r.Kind) + "-conjugate", Color: r.Color, SU2Doublet: false, SU3Charged: r.Color > 0, Hypercharge: -r.Hypercharge, BMinusL: -bl})
	}
	return states
}

func rightBMinusL(r yukawaintertwiner.RightSinglet) float64 {
	switch r.Kind {
	case yukawaintertwiner.UpType, yukawaintertwiner.DownType:
		return 1.0 / 3.0
	case yukawaintertwiner.NeutrinoType, yukawaintertwiner.ElectronType:
		return -1.0
	default:
		return 0
	}
}

func moment(name string, value, eps float64, desc string) Moment {
	if math.Abs(value) < eps {
		value = 0
	}
	return Moment{Name: name, Value: value, Cancels: math.Abs(value) <= eps, Description: desc}
}

func allCancel(xs []Moment) bool {
	for _, x := range xs {
		if !x.Cancels {
			return false
		}
	}
	return true
}

func sum(states []WeylState, f func(WeylState) float64) float64 {
	v := 0.0
	for _, s := range states {
		v += f(s)
	}
	return v
}

func su2sq(states []WeylState, charge func(WeylState) float64) float64 {
	v := 0.0
	for _, s := range states {
		if s.SU2Doublet {
			// The common fundamental Dynkin index T(2)=1/2 is included.
			v += charge(s) / 2.0
		}
	}
	return v
}

func su3sq(states []WeylState, charge func(WeylState) float64) float64 {
	v := 0.0
	for _, s := range states {
		if s.SU3Charged {
			v += charge(s) / 2.0
		}
	}
	return v
}

func cube(x float64) float64 { return x * x * x }
