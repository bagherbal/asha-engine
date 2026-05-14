// Package fierz audits the missing finite Fierz projection needed to turn the
// native x∧p/u(4) current inventory into an NJL scalar-channel kernel.
//
// Gate 56 established the available current algebra and the formal template
//
//	J_A J_A -> c_A (Ψ̄_R Ψ_L)(Ψ̄_L Ψ_R) + ...
//
// but it deliberately did not assume the scalar-channel coefficients c_A.  This
// package makes that obstruction explicit.  It constructs the finite bookkeeping
// domain for a Fierz projection, verifies that the target scalar left-right
// channels exist, and classifies the tensors still required before an attractive
// NJL kernel can be claimed.
//
// The result is a disciplined no-go/audit: generator counts and Yukawa incidence
// are insufficient to determine attraction.  A real finite Fierz theorem needs a
// chiral bilinear form, Lorentz/Clifford contraction rules, current-generator
// normalization, and an explicit projection onto the scalar LR channel.
package fierz

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/fourfermion"
)

type FierzTensorRequirement struct {
	Name        string
	Role        string
	Available   bool
	MissingPart string
}

type CurrentProjectionSlot struct {
	Sector          string
	Dimension       int
	TargetCandidate string
	ProjectionKnown bool
	SignKnown       bool
	Coefficient     string
}

type Analysis struct {
	FourFermion fourfermion.Analysis

	CurrentSectorCount int
	U4Dimension        int

	ScalarLRTargetAvailable bool
	ScalarChannelCount      int
	ChannelKinds            int
	GenerationCount         int

	ProjectionSlots []CurrentProjectionSlot
	Requirements    []FierzTensorRequirement

	ChiralBilinearMetricDerived       bool
	CliffordTraceRulesDerived         bool
	GeneratorNormalizationDerived     bool
	ScalarProjectionCoefficientsKnown bool
	AttractiveSignDerived             bool
	NativeFierzProjectionComplete     bool
	FourFermionStrengthDerived        bool
	UpDownSplittingDerived            bool
	HiddenObservedInputUsed           bool

	FormalProjectionExpression string
	TruthStatement             string
	RecommendedNextGate        string
	RemainingUnknowns          []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		ff, err := fourfermion.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(ff)
	})
	return defaultValue, defaultErr
}

func Build(ff fourfermion.Analysis) (Analysis, error) {
	if ff.U4Dimension != 16 || !ff.CurrentAlgebraAvailable {
		return Analysis{}, fmt.Errorf("Gate 57 requires the Gate 56 u(4) current inventory")
	}
	if !ff.ScalarLRChannelAvailable || len(ff.Gap.ChannelCriticalities) == 0 {
		return Analysis{}, fmt.Errorf("Gate 57 requires scalar left-right channel data from the NJL criticality ledger")
	}

	generations := ff.Gap.GenerationCount
	kinds := ff.Gap.KindCount
	scalarChannels := len(ff.Gap.ChannelCriticalities)

	slots := make([]CurrentProjectionSlot, 0, len(ff.Sectors))
	for _, s := range ff.Sectors {
		target := "not scalar-channel-selected"
		if s.NJLRelevant {
			target = "possible scalar LR channel after finite Fierz projection"
		}
		slots = append(slots, CurrentProjectionSlot{
			Sector:          s.Name,
			Dimension:       s.Dimension,
			TargetCandidate: target,
			ProjectionKnown: false,
			SignKnown:       false,
			Coefficient:     "c_" + slug(s.Name) + " open",
		})
	}

	reqs := []FierzTensorRequirement{
		{
			Name:        "chiral scalar bilinear form",
			Role:        "defines the finite target (Ψ̄_R Ψ_L)(Ψ̄_L Ψ_R) as a projector rather than a label",
			Available:   false,
			MissingPart: "explicit left/right Clifford/Fock bilinear metric",
		},
		{
			Name:        "Clifford/Lorentz contraction rules",
			Role:        "supply the spinor trace identities that replace importing continuum Fierz tables",
			Available:   false,
			MissingPart: "native gamma/pseudoscalar trace algebra for the selected representation",
		},
		{
			Name:        "u(4) generator normalization",
			Role:        "fixes the relative weights of central, color, B-L, and leptoquark current exchange",
			Available:   false,
			MissingPart: "finite kinetic/trace normalization for x∧p generators",
		},
		{
			Name:        "current-current scalar projection coefficients",
			Role:        "computes c_A in J_AJ_A → c_A(Ψ̄_RΨ_L)(Ψ̄_LΨ_R)+...",
			Available:   false,
			MissingPart: "explicit projection tensor from current-pair space to scalar LR channel",
		},
		{
			Name:        "attractive sign convention from the finite action",
			Role:        "decides whether the scalar channel lowers the energy rather than raising it",
			Available:   false,
			MissingPart: "finite propagator/action sign and Euclidean/Lorentzian continuation rule",
		},
	}

	truth := "The finite engine can now state the Fierz problem precisely: the u(4) current inventory and scalar LR target exist, but no native Fierz tensor has yet projected current-current exchange into an attractive scalar channel. Generator counts do not determine c_A or its sign. The next theorem must construct the finite chiral bilinear metric and Clifford trace rules."

	return Analysis{
		FourFermion:                       ff,
		CurrentSectorCount:                len(ff.Sectors),
		U4Dimension:                       ff.U4Dimension,
		ScalarLRTargetAvailable:           true,
		ScalarChannelCount:                scalarChannels,
		ChannelKinds:                      kinds,
		GenerationCount:                   generations,
		ProjectionSlots:                   slots,
		Requirements:                      reqs,
		ChiralBilinearMetricDerived:       false,
		CliffordTraceRulesDerived:         false,
		GeneratorNormalizationDerived:     false,
		ScalarProjectionCoefficientsKnown: false,
		AttractiveSignDerived:             false,
		NativeFierzProjectionComplete:     false,
		FourFermionStrengthDerived:        false,
		UpDownSplittingDerived:            false,
		HiddenObservedInputUsed:           false,
		FormalProjectionExpression:        "c_A = <scalar LR projector, J_A⊗J_A>_finite; G_hat = Σ_A g_A² c_A/M_A²",
		TruthStatement:                    truth,
		RecommendedNextGate:               "Gate 58 — Finite Chiral Bilinear Metric / Clifford Trace Construction",
		RemainingUnknowns: []string{
			"U-20D1A-CHIRAL-BILINEAR-METRIC: construct the finite Ψ̄_RΨ_L scalar projector",
			"U-20D1B-CLIFFORD-TRACE-RULES: compute native trace/Fierz identities instead of importing continuum tables",
			"U-20D1C-GENERATOR-NORMALIZATION: normalize x∧p/u(4) current generators by the finite kinetic trace",
			"U-20D2-ATTRACTIVE-SIGN: derive the sign of the scalar channel from the finite action",
			"U-20D4-UP-DOWN-SPLITTING: break the up/down quark tie without observed Yukawa input",
		},
	}, nil
}

func slug(s string) string {
	repl := strings.NewReplacer(" ", "_", "(", "", ")", "", "-", "_", "/", "_", "∧", "wedge")
	out := strings.ToLower(repl.Replace(s))
	out = strings.ReplaceAll(out, "__", "_")
	return out
}

func FormatSlots(xs []CurrentProjectionSlot) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		state := "open"
		if x.ProjectionKnown && x.SignKnown {
			state = "signed"
		} else if x.ProjectionKnown {
			state = "projected"
		}
		parts = append(parts, fmt.Sprintf("%s(dim=%d,%s,%s)", x.Sector, x.Dimension, state, x.Coefficient))
	}
	return strings.Join(parts, "; ")
}

func FormatRequirements(xs []FierzTensorRequirement) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		state := "missing"
		if x.Available {
			state = "available"
		}
		parts = append(parts, fmt.Sprintf("%s[%s]: %s", x.Name, state, x.MissingPart))
	}
	return strings.Join(parts, " | ")
}

func FormatUnknowns(xs []string) string { return strings.Join(xs, "; ") }
