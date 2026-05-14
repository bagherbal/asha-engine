// Package fourfermion audits whether the native x∧p / u(4) phase-space gauge
// sector is already sufficient to derive an NJL four-fermion kernel.
//
// Gates 52-55 built the composite-Higgs direction up to a criticality ledger:
// the finite engine has a Yukawa-incidence operator, three-color amplification,
// and a generation-pressure spurion, but the attractive four-fermion strength
// G_hat is still missing.  This package tests the natural next source: the
// cross-term x∧p sector, whose one-generation Fock action has u(4) structure.
//
// The result is intentionally conservative.  The u(4) inventory supplies a
// standard Pati-Salam-shaped current algebra and a formal current-current
// exchange template.  It does not by itself fix a Fierz projection into the
// scalar left-right channel, the sign of the attractive channel, a propagator
// normalization, or a regulator.  Therefore it does not yet derive G_hat, a
// nonzero NJL gap, or a Higgs scale.
package fourfermion

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/gapledger"
)

type Sector struct {
	Name        string
	Dimension   int
	Role        string
	Derived     bool
	NJLRelevant bool
}

type KernelCondition struct {
	Name        string
	Expression  string
	KnownPart   string
	MissingPart string
	Derived     bool
}

type Analysis struct {
	Gap gapledger.Analysis

	FockModes             int
	U4Dimension           int
	SU4AdjointDimension   int
	ColorSU3Dimension     int
	BLDimension           int
	LeptoquarkDimension   int
	CentralU1Dimension    int
	DecompositionComplete bool

	Sectors          []Sector
	KernelConditions []KernelCondition

	CurrentAlgebraAvailable             bool
	CurrentCurrentTemplateAvailable     bool
	ScalarLRChannelAvailable            bool
	FierzProjectionDerived              bool
	AttractiveChannelSignDerived        bool
	GaugePropagatorNormalizationDerived bool
	FourFermionStrengthDerived          bool
	NativeNJLKernelDerived              bool
	UpDownSplittingDerived              bool
	RegulatorDerived                    bool
	CriticalityClosed                   bool
	HiddenObservedCouplingsUsed         bool
	HiddenMassScaleUsed                 bool

	FormalKernelExpression        string
	KnownFiniteSkeleton           string
	MissingKernelData             string
	MaxFinitePressureSkeleton     float64
	RequiredCouplingStillSymbolic bool

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
		g, err := gapledger.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(g)
	})
	return defaultValue, defaultErr
}

func Build(g gapledger.Analysis) (Analysis, error) {
	if g.GenerationCount != 3 || g.KindCount != 4 {
		return Analysis{}, fmt.Errorf("Gate 56 requires the Gate 55 criticality ledger with 3 generations and 4 fermion kinds")
	}
	if g.StrongestWeightedPressure <= 0 {
		return Analysis{}, fmt.Errorf("Gate 56 requires a positive finite pressure skeleton")
	}

	fockModes := 4
	u4 := fockModes * fockModes
	central := 1
	color := (fockModes-1)*(fockModes-1) - 1 // su(3): 3^2-1
	bl := 1
	leptoquark := 2 * (fockModes - 1) // 3 complex off-diagonal color-lepton generators = 6 real
	su4Adjoint := u4 - central
	decompComplete := central+color+bl+leptoquark == u4

	sectors := []Sector{
		{Name: "central u(1)", Dimension: central, Role: "overall phase/current normalization; not enough for scalar condensation", Derived: true, NJLRelevant: false},
		{Name: "su(3)c", Dimension: color, Role: "color current sector; supplies color multiplicity but not top-only dominance", Derived: true, NJLRelevant: true},
		{Name: "u(1)B-L", Dimension: bl, Role: "matter charge polarization; diagonal current-current channel", Derived: true, NJLRelevant: true},
		{Name: "leptoquark off-diagonal", Dimension: leptoquark, Role: "Pati-Salam mixing directions; scalar-channel sign and viability require Fierz/phenomenology audit", Derived: true, NJLRelevant: false},
	}

	conditions := []KernelCondition{
		{
			Name:        "u(4) current inventory",
			Expression:  "x∧p generators ≃ a†_μ a_ν, μ,ν=0..3",
			KnownPart:   fmt.Sprintf("dim u(4)=%d with decomposition %d+%d+%d+%d", u4, central, color, bl, leptoquark),
			MissingPart: "none for inventory",
			Derived:     true,
		},
		{
			Name:        "formal current-current exchange",
			Expression:  "L_eff ∼ - g_A² (J_A J_A)/M_A²",
			KnownPart:   "current algebra slots are available",
			MissingPart: "g_A, M_A, propagator sign, and continuum matching",
			Derived:     false,
		},
		{
			Name:        "scalar LR Fierz projection",
			Expression:  "J_A J_A → c_A (Ψ̄_R Ψ_L)(Ψ̄_L Ψ_R)+...",
			KnownPart:   "left/right scalar channels are known from the Yukawa incidence operator",
			MissingPart: "native Fierz coefficients c_A in the finite Clifford/Fock representation",
			Derived:     false,
		},
		{
			Name:        "NJL kernel strength",
			Expression:  "G_hat = Σ_A g_A² c_A/M_A²",
			KnownPart:   "finite pressure skeleton K_channel is known from Gate 55",
			MissingPart: "the attractive sum G_hat is not derived",
			Derived:     false,
		},
	}

	return Analysis{
		Gap:                                 g,
		FockModes:                           fockModes,
		U4Dimension:                         u4,
		SU4AdjointDimension:                 su4Adjoint,
		ColorSU3Dimension:                   color,
		BLDimension:                         bl,
		LeptoquarkDimension:                 leptoquark,
		CentralU1Dimension:                  central,
		DecompositionComplete:               decompComplete,
		Sectors:                             sectors,
		KernelConditions:                    conditions,
		CurrentAlgebraAvailable:             true,
		CurrentCurrentTemplateAvailable:     true,
		ScalarLRChannelAvailable:            len(g.ChannelCriticalities) > 0,
		FierzProjectionDerived:              false,
		AttractiveChannelSignDerived:        false,
		GaugePropagatorNormalizationDerived: false,
		FourFermionStrengthDerived:          false,
		NativeNJLKernelDerived:              false,
		UpDownSplittingDerived:              false,
		RegulatorDerived:                    false,
		CriticalityClosed:                   false,
		HiddenObservedCouplingsUsed:         false,
		HiddenMassScaleUsed:                 false,
		FormalKernelExpression:              "G_hat = Σ_A g_A² c_A/M_A², with c_A from finite Fierz projection",
		KnownFiniteSkeleton:                 fmt.Sprintf("max K_channel=%.10f from %s/G%d", g.StrongestWeightedPressure, g.StrongestWeightedChannel.Kind, g.StrongestWeightedChannel.Generation),
		MissingKernelData:                   "finite Fierz coefficients, attractive sign, propagator normalization, regulator threshold, and up/down splitting",
		MaxFinitePressureSkeleton:           g.StrongestWeightedPressure,
		RequiredCouplingStillSymbolic:       true,
		TruthStatement:                      "The native x∧p/u(4) sector supplies the correct current-algebra inventory for an NJL-style four-fermion computation, but it still does not derive the attractive scalar-channel kernel G_hat. The next truth is a finite Fierz/projection computation, not an inserted coupling.",
		RecommendedNextGate:                 "Gate 57 — Finite Fierz Projection / Scalar-Channel Sign Audit",
		RemainingUnknowns: []string{
			"U-20D1-FINITE-FIERZ: compute current-current Fierz projection into Ψ̄_RΨ_L scalar channel",
			"U-20D2-ATTRACTIVE-SIGN: derive whether the scalar channel is attractive",
			"U-20D3-PROPAGATOR-NORMALIZATION: derive g_A²/M_A² or its finite spectral analogue",
			"U-20D4-UP-DOWN-SPLITTING: break the up/down quark tie without fitting",
			"U-20D5-REGULATOR: derive the critical threshold C_reg",
		},
	}, nil
}

func FormatSectors(xs []Sector) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		flag := "non-NJL"
		if x.NJLRelevant {
			flag = "NJL-relevant"
		}
		parts = append(parts, fmt.Sprintf("%s(dim=%d,%s)", x.Name, x.Dimension, flag))
	}
	return strings.Join(parts, "; ")
}

func FormatConditions(xs []KernelCondition) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		state := "open"
		if x.Derived {
			state = "derived"
		}
		parts = append(parts, fmt.Sprintf("%s[%s]: %s", x.Name, state, x.Expression))
	}
	return strings.Join(parts, " | ")
}

func FormatUnknowns(xs []string) string {
	return strings.Join(xs, "; ")
}
