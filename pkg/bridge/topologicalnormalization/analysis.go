// Package topologicalnormalization implements Gate 174: spectral-action
// normalization from the topological action seal.
//
// Gates 166-167 closed the relative gauge-kinetic ratio by the Fock
// representation trace, while Gate 173 sealed the mass-generation route at the
// current finite-data stage. Gate 174 therefore tests the next independent
// possibility: whether the finite topological action seal S_top=8*pi^2 fixes
// the absolute spectral-action prefactor f0 and hence u=1/g_*^2.
//
// The result is deliberately strict. The gate computes the conditional
// instanton-matching branch: if the finite index is identified with the
// continuum Yang--Mills topological charge and the finite trace normalization is
// identified with the continuum kinetic normalization, then S_top=8*pi^2 gives
// u=1, g_*^2=1, and the Fock trace K=(2,2,2,10/3) gives the same boundary
// ratio diag(1,1,1,5/3). But those two identifications are exactly the missing
// finite-to-continuum normalization bridge. Therefore the strict theorem ledger
// does not reduce nullity; it records only a conditional nullity reduction.
package topologicalnormalization

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/couplingnorm"
	"github.com/bagherbal/asha-engine/pkg/bridge/fockrepresentationtrace"
	"github.com/bagherbal/asha-engine/pkg/bridge/noncommutingtexturepair"
)

type InputAudit struct {
	GaugeRatioClosed         bool
	WeakAngleSeedClosed      bool
	MassGenerationSealed     bool
	TopologicalSealAvailable bool
	ContactIndex             float64
	TopologicalActionSeal    float64
	RepresentationTraceSU2   string
	RepresentationTraceY     string
	NormalizedBoundary       string
	WeakAngleSeed            string
	UsesObservedInput        bool
	Verdict                  string
}

type InstantonMatchingAudit struct {
	ContinuumFormula               string
	FiniteFormula                  string
	DirectEquation                 string
	ConditionalUInverseGStar       float64
	ConditionalGStarSquared        float64
	ConditionalGStar               float64
	RequiresContinuumIndexBridge   bool
	RequiresTraceKineticBridge     bool
	ContinuumIndexBridgeDerived    bool
	TraceKineticBridgeDerived      bool
	CanonicalStrictMatchingDerived bool
	ConditionalMatchingAvailable   bool
	TopologicalSealAloneSufficient bool
	Verdict                        string
}

type SpectralPrefactorConvention struct {
	Name                  string
	Formula               string
	Multiplier            float64
	KSU2                  string
	KY                    string
	ConditionalF0         float64
	ConditionalInverseSU2 float64
	ConditionalInverseY   float64
	ConditionalSin2       float64
	BoundaryRatio         string
	ConventionDependentF0 bool
	SameBoundaryPhysics   bool
	Verdict               string
}

type NormalizationFirewall struct {
	RelativeGaugeRatioClosed      bool
	TopologicalSealDerived        bool
	ConditionalAbsoluteUAvailable bool
	StrictAbsoluteUDerived        bool
	StrictF0Derived               bool
	BoundaryScaleDerived          bool
	ThresholdCorrectionsDerived   bool
	PhysicalCouplingsDerived      bool
	PhysicalFineStructureDerived  bool
	PhysicalMassesDerived         bool
	HiddenObservedInputUsed       bool
	StrictNullityBefore           int
	StrictNullityAfter            int
	ConditionalNullityAfter       int
	RemainingStrictUnknowns       []string
	ConditionalRemainingUnknowns  []string
	RecommendedNextGate           string
	Verdict                       string
}

type Analysis struct {
	MassSeal       noncommutingtexturepair.Analysis
	Trace          fockrepresentationtrace.Analysis
	Coupling       couplingnorm.Analysis
	Input          InputAudit
	Matching       InstantonMatchingAudit
	Conventions    []SpectralPrefactorConvention
	Firewall       NormalizationFirewall
	TruthStatement string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		mass, err := noncommutingtexturepair.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		trace, err := fockrepresentationtrace.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		coupling, err := couplingnorm.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(mass, trace, coupling, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(mass noncommutingtexturepair.Analysis, trace fockrepresentationtrace.Analysis, coupling couplingnorm.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !trace.Firewall.BoundaryGaugeRatioClosed || !trace.Firewall.BoundaryWeakAngleSeedClosed {
		return Analysis{}, fmt.Errorf("Gate 174 requires the Gate 167 representation-trace gauge ratio to be closed")
	}
	if !mass.NoGo.MassGenerationSealedAtCurrentStage || mass.Inventory.QualifiedNonCommutingPairs != 0 {
		return Analysis{}, fmt.Errorf("Gate 174 requires Gate 173 to seal the mass-generation route at the current stage")
	}
	if coupling.TopologicalActionSeal <= eps || coupling.ContactIndex <= eps {
		return Analysis{}, fmt.Errorf("Gate 174 requires positive topological action data")
	}
	if coupling.HiddenObservedCouplingUsed || trace.TraceAudit.UsesObservedInput {
		return Analysis{}, fmt.Errorf("Gate 174 refuses hidden observed coupling input")
	}

	ksu2 := trace.TraceAudit.KSU2T3.Float64()
	ky := trace.TraceAudit.KU1Y.Float64()
	if ksu2 <= eps || ky <= eps {
		return Analysis{}, fmt.Errorf("Gate 174 requires positive representation traces")
	}

	input := InputAudit{
		GaugeRatioClosed:         trace.Firewall.BoundaryGaugeRatioClosed,
		WeakAngleSeedClosed:      trace.Firewall.BoundaryWeakAngleSeedClosed,
		MassGenerationSealed:     mass.NoGo.MassGenerationSealedAtCurrentStage,
		TopologicalSealAvailable: coupling.Action.DimensionlessActionDerived,
		ContactIndex:             coupling.ContactIndex,
		TopologicalActionSeal:    coupling.TopologicalActionSeal,
		RepresentationTraceSU2:   trace.TraceAudit.KSU2T3.String(),
		RepresentationTraceY:     trace.TraceAudit.KU1Y.String(),
		NormalizedBoundary:       fmt.Sprintf("diag(%s,%s,%s,%s)", trace.TraceAudit.NormalizedT1, trace.TraceAudit.NormalizedT2, trace.TraceAudit.NormalizedT3, trace.TraceAudit.NormalizedY),
		WeakAngleSeed:            trace.TraceAudit.WeakAngleSeed.String(),
		UsesObservedInput:        coupling.HiddenObservedCouplingUsed || trace.TraceAudit.UsesObservedInput,
		Verdict:                  "relative gauge data and the topological action seal are simultaneously available; the only question is absolute normalization",
	}

	matching := buildMatching(coupling)
	conventions := buildConventions(ksu2, ky, matching.ConditionalUInverseGStar)
	firewall := NormalizationFirewall{
		RelativeGaugeRatioClosed:      trace.Firewall.BoundaryGaugeRatioClosed,
		TopologicalSealDerived:        coupling.TopologicalActionSeal > eps && coupling.ContactIndex > eps,
		ConditionalAbsoluteUAvailable: matching.ConditionalMatchingAvailable,
		StrictAbsoluteUDerived:        matching.CanonicalStrictMatchingDerived,
		StrictF0Derived:               false,
		BoundaryScaleDerived:          false,
		ThresholdCorrectionsDerived:   false,
		PhysicalCouplingsDerived:      false,
		PhysicalFineStructureDerived:  false,
		PhysicalMassesDerived:         false,
		HiddenObservedInputUsed:       false,
		StrictNullityBefore:           3,
		StrictNullityAfter:            3,
		ConditionalNullityAfter:       2,
		RemainingStrictUnknowns: []string{
			"u=1/g_*^2: still strict-open until finite index and finite trace are matched to the continuum Yang--Mills normalization",
			"L=ln(M*/mu): boundary scale/evaluation scale remains dimensionless-open",
			"Δb_i: threshold activation/decoupling remains open",
		},
		ConditionalRemainingUnknowns: []string{
			"L=ln(M*/mu): boundary scale/evaluation scale",
			"Δb_i: threshold activation/decoupling",
		},
		RecommendedNextGate: "Gate 175 — finite-to-continuum instanton trace-normalization bridge",
		Verdict:             "S_top=8π² yields a clean conditional u=1 branch, but not a strict theorem-level nullity reduction until the continuum index bridge and trace/kinetic normalization are derived",
	}

	return Analysis{
		MassSeal:       mass,
		Trace:          trace,
		Coupling:       coupling,
		Input:          input,
		Matching:       matching,
		Conventions:    conventions,
		Firewall:       firewall,
		TruthStatement: "Gate 174 shows that the topological action seal can normalize the absolute coupling only conditionally. If S_top=8π² is identified with the unit Yang--Mills instanton action, then u=1/g_*²=1 and the spectral-action prefactor is fixed up to the conventional multiplier in front of Tr_rep(T²). But the current finite engine has not derived the finite-to-continuum index map or trace/kinetic normalization, so strict nullity remains 3 rather than dropping to 2.",
	}, nil
}

func buildMatching(c couplingnorm.Analysis) InstantonMatchingAudit {
	conditionalU := 1.0 / c.ContactIndex
	return InstantonMatchingAudit{
		ContinuumFormula:               "S_YM(k=I_BG)=8π² I_BG / g_*²",
		FiniteFormula:                  "S_top=8π² I_BG",
		DirectEquation:                 "8π² I_BG = 8π² I_BG / g_*² ⇒ u=1/g_*²=1 when I_BG≠0",
		ConditionalUInverseGStar:       conditionalU,
		ConditionalGStarSquared:        1.0 / conditionalU,
		ConditionalGStar:               math.Sqrt(1.0 / conditionalU),
		RequiresContinuumIndexBridge:   true,
		RequiresTraceKineticBridge:     true,
		ContinuumIndexBridgeDerived:    c.ContinuumIndexBridgeDerived,
		TraceKineticBridgeDerived:      c.TraceNormalizationDerived,
		CanonicalStrictMatchingDerived: c.ContinuumIndexBridgeDerived && c.TraceNormalizationDerived,
		ConditionalMatchingAvailable:   c.ContactIndex > 0 && c.TopologicalActionSeal > 0,
		TopologicalSealAloneSufficient: false,
		Verdict:                        "the algebra supplies the exact 8π² seal; equating it with a continuum instanton action is a conditional matching rule, not yet an internal theorem",
	}
}

func buildConventions(ksu2, ky, u float64) []SpectralPrefactorConvention {
	return []SpectralPrefactorConvention{
		prefactorConvention("single-trace convention", "1/g_a² = f0 · Tr_rep(T_a²)", 1, ksu2, ky, u),
		prefactorConvention("two-sided/chiral convention", "1/g_a² = 2 f0 · Tr_rep(T_a²)", 2, ksu2, ky, u),
	}
}

func prefactorConvention(name, formula string, multiplier, ksu2, ky, u float64) SpectralPrefactorConvention {
	f0 := u / (multiplier * ksu2)
	invSU2 := multiplier * f0 * ksu2
	invY := multiplier * f0 * ky
	sin2 := invSU2 / (invSU2 + invY)
	return SpectralPrefactorConvention{
		Name:                  name,
		Formula:               formula,
		Multiplier:            multiplier,
		KSU2:                  rationalLike(ksu2),
		KY:                    rationalLike(ky),
		ConditionalF0:         f0,
		ConditionalInverseSU2: invSU2,
		ConditionalInverseY:   invY,
		ConditionalSin2:       sin2,
		BoundaryRatio:         fmt.Sprintf("1 : %.12g", invY/invSU2),
		ConventionDependentF0: true,
		SameBoundaryPhysics:   close(invSU2, u, 1e-10) && close(sin2, 3.0/8.0, 1e-10),
		Verdict:               "f0 changes with action-normalization convention, while u and the 5/3 ratio are unchanged on the conditional instanton-matching branch",
	}
}

func rationalLike(x float64) string {
	if close(x, 2, 1e-10) {
		return "2"
	}
	if close(x, 10.0/3.0, 1e-10) {
		return "10/3"
	}
	return fmt.Sprintf("%.12g", x)
}

func close(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatInput(a InputAudit) string {
	return fmt.Sprintf("ratioClosed=%t weakSeedClosed=%t massSealed=%t I_BG=%.12g S_top=%.12g K_SU2=%s K_Y=%s boundary=%s sin2=%s observed=%t",
		a.GaugeRatioClosed, a.WeakAngleSeedClosed, a.MassGenerationSealed, a.ContactIndex, a.TopologicalActionSeal, a.RepresentationTraceSU2, a.RepresentationTraceY, a.NormalizedBoundary, a.WeakAngleSeed, a.UsesObservedInput)
}

func FormatMatching(a InstantonMatchingAudit) string {
	return fmt.Sprintf("%s; %s; %s; conditional u=%.12g g2=%.12g g=%.12g; requiresIndexBridge=%t derived=%t requiresTraceBridge=%t derived=%t strict=%t conditional=%t",
		a.ContinuumFormula, a.FiniteFormula, a.DirectEquation, a.ConditionalUInverseGStar, a.ConditionalGStarSquared, a.ConditionalGStar,
		a.RequiresContinuumIndexBridge, a.ContinuumIndexBridgeDerived, a.RequiresTraceKineticBridge, a.TraceKineticBridgeDerived, a.CanonicalStrictMatchingDerived, a.ConditionalMatchingAvailable)
}

func FormatConventions(xs []SpectralPrefactorConvention) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%s: %s, f0=%.12g, inv=(SU2 %.12g,Y %.12g), ratio=%s, sin2=%.12g, conventionDependent=%t",
			x.Name, x.Formula, x.ConditionalF0, x.ConditionalInverseSU2, x.ConditionalInverseY, x.BoundaryRatio, x.ConditionalSin2, x.ConventionDependentF0)
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatFirewall(a NormalizationFirewall) string {
	return fmt.Sprintf("relativeRatioClosed=%t topologicalSeal=%t conditionalU=%t strictU=%t strictF0=%t boundaryScale=%t thresholds=%t physicalCouplings=%t alpha=%t masses=%t observed=%t nullityStrict=%d->%d conditional->%d",
		a.RelativeGaugeRatioClosed, a.TopologicalSealDerived, a.ConditionalAbsoluteUAvailable, a.StrictAbsoluteUDerived, a.StrictF0Derived, a.BoundaryScaleDerived, a.ThresholdCorrectionsDerived, a.PhysicalCouplingsDerived, a.PhysicalFineStructureDerived, a.PhysicalMassesDerived, a.HiddenObservedInputUsed, a.StrictNullityBefore, a.StrictNullityAfter, a.ConditionalNullityAfter)
}
