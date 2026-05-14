// Package rgfirewall implements Gate 103: the finite RG flow and boundary-scale
// selection firewall.
//
// Gate 102 produced embedded finite boundary data
//
//	K_* = diag(1,1,1,5/3),     sin^2_* = 3/8,
//
// without inserting observed couplings.  This package performs the next honest
// operation: it couples that boundary seed to the finite-spectrum beta audit and
// writes the formal one-loop flow family.  The result is deliberately a firewall
// theorem.  It proves what the current engine can seed and what it still cannot
// select: the boundary coupling g_*^2, the logarithmic scale interval
// L=ln(M*/mu), threshold/decoupling corrections, and therefore physical
// low-energy alpha, theta_W, g2, gY, W/Z masses, and the Higgs scale.
package rgfirewall

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/betacoeff"
	"github.com/bagherbal/asha-engine/pkg/bridge/contactembedding"
	"github.com/bagherbal/asha-engine/pkg/bridge/thresholdactivation"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type RunningSample struct {
	Name          string
	UInverseGStar float64
	LogInterval   float64
	A1Inverse     float64
	A2Inverse     float64
	Sin2          float64
	InverseAlpha  float64
	Physical      bool
	Detail        string
}

type MissingVariable struct {
	Name    string
	Symbol  string
	Derived bool
	Detail  string
}

type FormalFlow struct {
	Convention           string
	U1InverseExpression  string
	SU2InverseExpression string
	EMInverseExpression  string
	Sin2Expression       string
	AlphaExpression      string
}

type Analysis struct {
	Embedding  contactembedding.Analysis
	Beta       betacoeff.Analysis
	Thresholds thresholdactivation.Analysis

	EmbeddedBoundaryHessian linear.Matrix
	BoundaryKY              float64
	BoundarySin2            float64
	BoundaryKineticPositive bool
	BoundaryDataDerived     bool

	B1                      float64
	B2                      float64
	B3                      float64
	B1Slope                 float64
	B2Slope                 float64
	B3Slope                 float64
	BetaDiagnosticAvailable bool
	BetaThresholdCorrected  bool

	Flow                        FormalFlow
	MissingVariables            []MissingVariable
	FreeVariableCount           int
	FormalRGFamilyConstructed   bool
	TwoParameterUnderdetermined bool
	BoundaryScaleDimensionNoGo  bool
	ThresholdFirewallClosed     bool

	SampleA                RunningSample
	SampleB                RunningSample
	NonUniquenessWitnessed bool

	BoundaryCouplingDerived  bool
	BoundaryScaleDerived     bool
	ThresholdRuleDerived     bool
	FiniteRGTheoremDerived   bool
	PhysicalWeakAngleDerived bool
	FineStructureDerived     bool
	PhysicalMassesDerived    bool
	HiddenObservedInputUsed  bool

	TruthStatement      string
	RejectedClaims      []string
	RemainingUnknowns   []string
	RecommendedNextGate string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		embedding, err := contactembedding.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		beta, err := betacoeff.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		thresholds, err := thresholdactivation.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(embedding, beta, thresholds, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(embedding contactembedding.Analysis, beta betacoeff.Analysis, thresholds thresholdactivation.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !embedding.ContactToMatterEmbeddingSelected || !embedding.EmbeddedMatterHessianSelected {
		return Analysis{}, fmt.Errorf("Gate 103 requires the Gate 102 embedded matter boundary Hessian")
	}
	if !beta.DerivedFromFiniteInventory || beta.ImportedSMBetaTable || beta.HiddenObservedCouplingsUsed {
		return Analysis{}, fmt.Errorf("Gate 103 requires beta candidates reconstructed from finite inventory without observed couplings")
	}
	if embedding.EmbeddedMatterHessian.Rows() != 4 || embedding.EmbeddedMatterHessian.Cols() != 4 {
		return Analysis{}, fmt.Errorf("Gate 103 expects a 4x4 embedded boundary Hessian")
	}
	kY := embedding.MatterHyperchargeKY
	if kY <= eps {
		return Analysis{}, fmt.Errorf("Gate 103 requires positive k_Y")
	}
	boundarySin2 := 1 / (1 + kY)
	h := embedding.EmbeddedMatterHessian
	positive := h.At(0, 0) > eps && h.At(1, 1) > eps && h.At(2, 2) > eps && h.At(3, 3) > eps

	b1, b2, b3 := beta.B1GUTNormalized, beta.B2, beta.B3
	s1, s2, s3 := b1/(8*math.Pi*math.Pi), b2/(8*math.Pi*math.Pi), b3/(8*math.Pi*math.Pi)
	flow := FormalFlow{
		Convention:           "beta(g_i)=b_i g_i^3/(16π²), L=ln(M*/μ), u=1/g_*²",
		U1InverseExpression:  fmt.Sprintf("1/g_Y²(μ)=k_Y·u + (b1/8π²)L = %.10f·u + %.10f·L", kY, s1),
		SU2InverseExpression: fmt.Sprintf("1/g_2²(μ)=u + (b2/8π²)L = u %+ .10f·L", s2),
		EMInverseExpression:  fmt.Sprintf("1/e²(μ)=(1+k_Y)u + ((b1+b2)/8π²)L = %.10f·u %+ .10f·L", 1+kY, s1+s2),
		Sin2Expression:       fmt.Sprintf("sin²θ(μ)=(u+(b2/8π²)L)/((1+k_Y)u+((b1+b2)/8π²)L) with k_Y=%.10f", kY),
		AlphaExpression:      "α(μ)=1/(4π[(1+k_Y)u+((b1+b2)/8π²)L])",
	}

	missing := []MissingVariable{
		{Name: "normalized boundary coupling", Symbol: "g_*² or u=1/g_*²", Derived: false, Detail: "finite Hessian fixes relative kinetic normalization, not the absolute coupling unit"},
		{Name: "boundary scale", Symbol: "M*", Derived: false, Detail: "current finite data are dimensionless; no length/energy unit has been selected"},
		{Name: "RG interval", Symbol: "L=ln(M*/μ)", Derived: false, Detail: "without M* and a physical evaluation scale μ the log interval is a free parameter"},
		{Name: "threshold activation rule", Symbol: "Δb_i(L)", Derived: false, Detail: "threshold candidates exist, but decoupling/matching rules are still open"},
		{Name: "continuum matching prescription", Symbol: "finite fields → continuum fields", Derived: false, Detail: "the beta numbers use an explicit continuum one-loop assumption rather than a native finite RG theorem"},
	}

	// Two purely mathematical sample points. They are not physical inputs; they are
	// witnesses that the current formal family has more than one output until u
	// and L are selected by a later theorem.
	sampleA := runningSample("boundary/no-running witness", kY, s1, s2, 1, 0)
	sampleB := runningSample("shifted-log witness", kY, s1, s2, 1, 1)
	nonUnique := math.Abs(sampleA.Sin2-sampleB.Sin2) > eps && math.Abs(sampleA.InverseAlpha-sampleB.InverseAlpha) > eps

	thresholdClosed := !thresholds.DecouplingRuleDerived && !thresholds.PhysicalMassUnitDerived && !thresholds.ThresholdCorrectedBetaDerived && !thresholds.HiddenScaleInserted
	truth := "Gate 103 constructs the formal finite-seeded RG family from the Gate 102 embedded Hessian and the finite-spectrum beta diagnostic.  The boundary data are real and useful: K_*=diag(1,1,1,5/3), sin²_*=3/8, and beta candidates b1=41/10, b2=-19/6, b3=-7 are available under the stated continuum one-loop assumption.  But the flow still has free u=1/g_*² and L=ln(M*/μ), and the threshold activation map is not derived.  Therefore the engine proves a firewall: physical alpha, theta_W, couplings, masses, and a boundary scale cannot yet be claimed."

	return Analysis{
		Embedding:                   embedding,
		Beta:                        beta,
		Thresholds:                  thresholds,
		EmbeddedBoundaryHessian:     h,
		BoundaryKY:                  kY,
		BoundarySin2:                boundarySin2,
		BoundaryKineticPositive:     positive,
		BoundaryDataDerived:         positive && close(boundarySin2, 3.0/8.0, 1e-8),
		B1:                          b1,
		B2:                          b2,
		B3:                          b3,
		B1Slope:                     s1,
		B2Slope:                     s2,
		B3Slope:                     s3,
		BetaDiagnosticAvailable:     close(b1, 41.0/10.0, 1e-8) && close(b2, -19.0/6.0, 1e-8) && close(b3, -7, 1e-8),
		BetaThresholdCorrected:      false,
		Flow:                        flow,
		MissingVariables:            missing,
		FreeVariableCount:           len(missing),
		FormalRGFamilyConstructed:   true,
		TwoParameterUnderdetermined: true,
		BoundaryScaleDimensionNoGo:  true,
		ThresholdFirewallClosed:     thresholdClosed,
		SampleA:                     sampleA,
		SampleB:                     sampleB,
		NonUniquenessWitnessed:      nonUnique,
		BoundaryCouplingDerived:     false,
		BoundaryScaleDerived:        false,
		ThresholdRuleDerived:        false,
		FiniteRGTheoremDerived:      false,
		PhysicalWeakAngleDerived:    false,
		FineStructureDerived:        false,
		PhysicalMassesDerived:       false,
		HiddenObservedInputUsed:     false,
		TruthStatement:              truth,
		RejectedClaims: []string{
			"sin²_*=3/8 is the measured low-energy weak angle",
			"the finite boundary Hessian fixes alpha or g_*²",
			"the continuum one-loop beta diagnostic is already a native finite RG theorem",
			"a physical GUT/boundary scale M* follows from dimensionless finite normalization data",
			"threshold candidates can correct beta coefficients before activation and decoupling rules are derived",
		},
		RemainingUnknowns: []string{
			"U-23A-BOUNDARY-SCALE-OPERATOR: derive a dimensionful or scale-selecting finite invariant for M*",
			"U-23B-ABSOLUTE-COUPLING-UNIT: derive g_*² or an equivalent action normalization beyond relative Hessian entries",
			"U-23C-NATIVE-FINITE-RG: replace the continuum one-loop assumption with a finite coarse-graining/flow theorem",
			"U-23D-THRESHOLD-DECOUPLING: derive Δb_i activation and matching rules for finite heavy/contact modes",
			"U-23E-PHYSICAL-EVALUATION: only after M*, thresholds, and g_*² exist may alpha, thetaW, g2, gY, and masses be computed",
		},
		RecommendedNextGate: "Gate 104 — boundary scale operator / absolute coupling unit search",
	}, nil
}

func runningSample(name string, kY, s1, s2, u, L float64) RunningSample {
	a1 := kY*u + s1*L
	a2 := u + s2*L
	emInv := a1 + a2
	sin2 := a2 / emInv
	invAlpha := 4 * math.Pi * emInv
	return RunningSample{
		Name:          name,
		UInverseGStar: u,
		LogInterval:   L,
		A1Inverse:     a1,
		A2Inverse:     a2,
		Sin2:          sin2,
		InverseAlpha:  invAlpha,
		Physical:      false,
		Detail:        "mathematical witness only; not an observed or fitted value",
	}
}

func close(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatMatrix(m linear.Matrix) string {
	rows := make([]string, 0, m.Rows())
	for r := 0; r < m.Rows(); r++ {
		cols := make([]string, 0, m.Cols())
		for c := 0; c < m.Cols(); c++ {
			v := m.At(r, c)
			if math.Abs(v) < 1e-12 {
				v = 0
			}
			cols = append(cols, fmt.Sprintf("%.10f", v))
		}
		rows = append(rows, "["+strings.Join(cols, ", ")+"]")
	}
	return strings.Join(rows, " ")
}

func FormatMissing(xs []MissingVariable) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		state := "open"
		if x.Derived {
			state = "derived"
		}
		parts = append(parts, fmt.Sprintf("%s[%s]: %s", x.Symbol, state, x.Detail))
	}
	return strings.Join(parts, "; ")
}

func FormatSample(s RunningSample) string {
	return fmt.Sprintf("%s: u=%.10f, L=%.10f, A1=%.10f, A2=%.10f, sin²=%.10f, α⁻¹=%.10f, physical=%t", s.Name, s.UInverseGStar, s.LogInterval, s.A1Inverse, s.A2Inverse, s.Sin2, s.InverseAlpha, s.Physical)
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
