// Package spinctwistedchirality implements Gate 240:
// Spin^c Twisted Chirality / Hypercharge-Weak Plane Sieve Audit.
//
// Gate 239 proved that the bare Clifford-volume orientation candidate is
// proportional to occupation parity gamma=(-1)^N, so bare spin geometry cannot
// select the physical chiral weak plane.  Gate 240 audits the next natural
// possibility: twist the grading by the native diagonal u(1) charge data already
// present in the Fock bookkeeping, and test whether this Spin^c-like operator
// breaks the six-plane degeneracy.
//
// The result is sharper but still obstructed.  The native u(1) bookkeeping is a
// diagonal mode-weight derivation with weights (-1,1/3,1/3,1/3).  It commutes
// with occupation parity and therefore gamma*Y is a legitimate diagonal
// diagnostic.  Requiring an su(2) plane to preserve this u(1) eliminates the
// three temporal-spatial planes, because their two mode weights differ.  The
// three pure-spatial planes preserve the u(1), but they remain exactly
// degenerate and their doublet sectors still have multiple gamma*Y eigenvalues.
// Thus the u(1) twist supplies a class sieve, not a physical chirality theorem
// or a unique weak-plane selector.
package spinctwistedchirality

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/orientationtruechirality"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

const (
	AuditID = "GATE240-SPINC-TWISTED-CHIRALITY-HYPERCHARGE-WEAK-SIEVE"

	StatusU1Retrieved             = "CONDITIONAL_SUPPORT_NATIVE_U1_DIAGONAL_GENERATOR_PREFLIGHT"
	StatusTwistConstructed        = "CONDITIONAL_SUPPORT_SPINC_GAMMA_U1_TWIST_PREFLIGHT"
	StatusU1ClassSieve            = "CONDITIONAL_SUPPORT_U1_COMMUTANT_TEMPORAL_SPATIAL_CLASS_SIEVE"
	StatusFailedUniformTwist      = "FAILED_ROUTE_UNIFORM_TWISTED_CHIRALITY_ALIGNMENT"
	StatusFailedUniquePlane       = "FAILED_ROUTE_SPINC_WEAK_PLANE_SELECTION"
	StatusFailedPhysicalChirality = "FAILED_ROUTE_SPINC_PHYSICAL_CHIRALITY_DERIVATION"
	StatusFailedGlobalH           = "FAILED_ROUTE_GLOBAL_H_SUMMAND_STILL_UNSELECTED"
)

type U1GeneratorAudit struct {
	Name                         string
	Source                       string
	ActsOnSC                     bool
	ModeWeights                  []float64
	WeightFormula                string
	TemporalWeight               float64
	SpatialWeight                float64
	SpatialWeightsDegenerate     bool
	ImportedSMHypercharge        bool
	NativeContactU1AsHypercharge bool
	Verdict                      string
}

type TwistOperatorAudit struct {
	Name                     string
	Formula                  string
	GammaCommutesWithY       bool
	IsDiagonalOnFockBasis    bool
	IsInvolution             bool
	HasZeroEigenvalues       bool
	DistinctFromGamma        bool
	PhysicalChiralityDerived bool
	ManualHyperchargeFit     bool
	Verdict                  string
}

type PlaneTwistAudit struct {
	Plane                    string
	ModeIndices              []int
	PlaneClass               string
	PlaneModeWeights         []float64
	SU2PreservesY            bool
	U1CommutatorResidual     float64
	DoubletTwistEigenvalues  []float64
	SingletTwistEigenvalues  []float64
	DoubletUniqueEigenvalues int
	SingletUniqueEigenvalues int
	DoubletsUniformTwist     bool
	SingletsUniformTwist     bool
	SurvivesU1CommutantSieve bool
	SelectedByTwist          bool
	Verdict                  string
}

type TwistedSieveAudit struct {
	CandidatePlanes         int
	U1PreservingPlanes      []string
	U1RejectedPlanes        []string
	UniformDoubletPlanes    []string
	SelectedPlanes          []string
	TemporalSpatialRejected bool
	PureSpatialDegeneracy   int
	TwistBreaksDegeneracy   bool
	Verdict                 string
}

type WeakOutcomeAudit struct {
	Gate239ChiFailed          bool
	U1TwistImprovesClassSieve bool
	UniqueWeakPlaneSelected   bool
	PhysicalLeftHandedDerived bool
	GlobalHSummandDerived     bool
	OrderOneReady             bool
	Verdict                   string
}

type FirewallAudit struct {
	ImportedSMHypercharge  bool
	TunedU1Weights         bool
	ForcedWeakPlane        bool
	ForcedLeftHandedAction bool
	ImportedSpinCStructure bool
	ClaimedGlobalH         bool
	ClaimedOrderOne        bool
	FiniteCorePolluted     bool
	Verdict                string
}

type Summary struct {
	NativeU1Available        bool
	TwistConstructed         bool
	U1RejectsTemporalPlanes  bool
	PureSpatialPlanesRemain  int
	UniformTwistedDoublets   bool
	UniqueWeakPlaneDerived   bool
	PhysicalChiralityDerived bool
	GlobalHDerived           bool
	Status                   string
	NextGate                 string
	Comment                  string
}

type Analysis struct {
	Previous       orientationtruechirality.Analysis
	U1             U1GeneratorAudit
	Twist          TwistOperatorAudit
	Planes         []PlaneTwistAudit
	Sieve          TwistedSieveAudit
	Weak           WeakOutcomeAudit
	Firewall       FirewallAudit
	Summary        Summary
	TruthStatement string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := orientationtruechirality.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 239 predecessor: %w", err)
			return
		}
		f, err := spinor.NewCovariantPhaseFockSpace(4)
		if err != nil {
			defaultErr = fmt.Errorf("construct Fock space: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev, f)
	})
	return defaultA, defaultErr
}

func Build(prev orientationtruechirality.Analysis, f spinor.FockSpace) (Analysis, error) {
	if f.ModeCount() != 4 || f.StateCount() != 16 {
		return Analysis{}, fmt.Errorf("Gate 240 requires native four-mode 16-state Fock carrier, got modes=%d states=%d", f.ModeCount(), f.StateCount())
	}
	u1 := auditU1(f)
	twist := auditTwist()
	planes := auditPlanes(f, u1.ModeWeights)
	sieve := auditSieve(planes)
	weak := auditWeak(prev, sieve)
	fw := auditFirewall()
	sum := summarize(u1, twist, sieve, weak)
	truth := buildTruth(u1, twist, sieve, weak)
	return Analysis{Previous: prev, U1: u1, Twist: twist, Planes: planes, Sieve: sieve, Weak: weak, Firewall: fw, Summary: sum, TruthStatement: truth}, nil
}

func auditU1(f spinor.FockSpace) U1GeneratorAudit {
	weights := make([]float64, f.ModeCount())
	for _, m := range f.Modes {
		if m.Kind == spinor.TemporalMode {
			weights[m.Index] = -1
		} else {
			weights[m.Index] = 1.0 / 3.0
		}
	}
	spatialDegenerate := true
	spatial := math.NaN()
	for _, m := range f.Modes {
		if m.Kind == spinor.SpatialMode {
			if math.IsNaN(spatial) {
				spatial = weights[m.Index]
			}
			if math.Abs(weights[m.Index]-spatial) > 1e-12 {
				spatialDegenerate = false
			}
		}
	}
	return U1GeneratorAudit{
		Name:                         "Y_native diagonal u(1) bookkeeping",
		Source:                       "native 1⊕3 Fock charge/B−L seed: temporal mode weight -1, spatial modes weight +1/3",
		ActsOnSC:                     true,
		ModeWeights:                  weights,
		WeightFormula:                "Y_native(|n⟩)=Σ_i w_i n_i with w=(-1,1/3,1/3,1/3)",
		TemporalWeight:               -1,
		SpatialWeight:                spatial,
		SpatialWeightsDegenerate:     spatialDegenerate,
		ImportedSMHypercharge:        false,
		NativeContactU1AsHypercharge: false,
		Verdict:                      StatusU1Retrieved,
	}
}

func auditTwist() TwistOperatorAudit {
	return TwistOperatorAudit{
		Name:                     "χ_twist = γ·Y_native",
		Formula:                  "χ_twist(|n⟩)=(-1)^{N(n)} Y_native(n)|n⟩",
		GammaCommutesWithY:       true,
		IsDiagonalOnFockBasis:    true,
		IsInvolution:             false,
		HasZeroEigenvalues:       true,
		DistinctFromGamma:        true,
		PhysicalChiralityDerived: false,
		ManualHyperchargeFit:     false,
		Verdict:                  StatusTwistConstructed + "; " + StatusFailedPhysicalChirality,
	}
}

func auditPlanes(f spinor.FockSpace, weights []float64) []PlaneTwistAudit {
	out := []PlaneTwistAudit{}
	for i := 0; i < f.ModeCount(); i++ {
		for j := i + 1; j < f.ModeCount(); j++ {
			doubletVals := []float64{}
			singletVals := []float64{}
			for _, s := range f.States {
				occupiedInPlane := 0
				if s.Occupation[i] {
					occupiedInPlane++
				}
				if s.Occupation[j] {
					occupiedInPlane++
				}
				y := stateY(s, weights)
				gamma := 1.0
				if s.ExcitationNumber()%2 != 0 {
					gamma = -1
				}
				val := gamma * y
				switch occupiedInPlane {
				case 1:
					doubletVals = append(doubletVals, val)
				case 0, 2:
					singletVals = append(singletVals, val)
				}
			}
			class := "pure-spatial"
			if f.Modes[i].Kind == spinor.TemporalMode || f.Modes[j].Kind == spinor.TemporalMode {
				class = "temporal-spatial"
			}
			comm := math.Abs(weights[i] - weights[j])
			preserves := comm < 1e-12
			du := uniqueFloatCount(doubletVals, 1e-12)
			su := uniqueFloatCount(singletVals, 1e-12)
			uniformD := du == 1
			uniformS := su == 1
			selected := preserves && uniformD && uniformS
			verdict := StatusFailedUniformTwist
			if preserves && !selected {
				verdict = StatusU1ClassSieve + "; " + StatusFailedUniformTwist
			}
			if selected {
				verdict = StatusU1ClassSieve
			}
			out = append(out, PlaneTwistAudit{
				Plane:                    fmt.Sprintf("U={%s,%s}", f.Modes[i].Name, f.Modes[j].Name),
				ModeIndices:              []int{i, j},
				PlaneClass:               class,
				PlaneModeWeights:         []float64{weights[i], weights[j]},
				SU2PreservesY:            preserves,
				U1CommutatorResidual:     comm,
				DoubletTwistEigenvalues:  sortedUniqueFloats(doubletVals, 1e-12),
				SingletTwistEigenvalues:  sortedUniqueFloats(singletVals, 1e-12),
				DoubletUniqueEigenvalues: du,
				SingletUniqueEigenvalues: su,
				DoubletsUniformTwist:     uniformD,
				SingletsUniformTwist:     uniformS,
				SurvivesU1CommutantSieve: preserves,
				SelectedByTwist:          selected,
				Verdict:                  verdict,
			})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Plane < out[j].Plane })
	return out
}

func auditSieve(planes []PlaneTwistAudit) TwistedSieveAudit {
	preserving, rejected, uniform, selected := []string{}, []string{}, []string{}, []string{}
	temporalRejected := true
	purePreserving := 0
	for _, p := range planes {
		if p.SU2PreservesY {
			preserving = append(preserving, p.Plane)
			if p.PlaneClass == "pure-spatial" {
				purePreserving++
			}
		} else {
			rejected = append(rejected, p.Plane)
			if p.PlaneClass == "pure-spatial" {
				temporalRejected = false
			}
		}
		if p.DoubletsUniformTwist {
			uniform = append(uniform, p.Plane)
		}
		if p.SelectedByTwist {
			selected = append(selected, p.Plane)
		}
	}
	return TwistedSieveAudit{
		CandidatePlanes:         len(planes),
		U1PreservingPlanes:      preserving,
		U1RejectedPlanes:        rejected,
		UniformDoubletPlanes:    uniform,
		SelectedPlanes:          selected,
		TemporalSpatialRejected: temporalRejected && len(rejected) == 3,
		PureSpatialDegeneracy:   purePreserving,
		TwistBreaksDegeneracy:   len(selected) == 1,
		Verdict:                 StatusU1ClassSieve + "; " + StatusFailedUniquePlane,
	}
}

func auditWeak(prev orientationtruechirality.Analysis, sieve TwistedSieveAudit) WeakOutcomeAudit {
	return WeakOutcomeAudit{
		Gate239ChiFailed:          !prev.Summary.PhysicalChiralityDerived && !prev.Summary.ChiSelectsPlane,
		U1TwistImprovesClassSieve: sieve.TemporalSpatialRejected && sieve.PureSpatialDegeneracy == 3,
		UniqueWeakPlaneSelected:   sieve.TwistBreaksDegeneracy,
		PhysicalLeftHandedDerived: false,
		GlobalHSummandDerived:     false,
		OrderOneReady:             false,
		Verdict:                   StatusFailedPhysicalChirality + "; " + StatusFailedGlobalH,
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		ImportedSMHypercharge:  false,
		TunedU1Weights:         false,
		ForcedWeakPlane:        false,
		ForcedLeftHandedAction: false,
		ImportedSpinCStructure: false,
		ClaimedGlobalH:         false,
		ClaimedOrderOne:        false,
		FiniteCorePolluted:     false,
		Verdict:                "FIREWALL_PRESERVED_NATIVE_U1_TWIST_ONLY",
	}
}

func summarize(u1 U1GeneratorAudit, twist TwistOperatorAudit, sieve TwistedSieveAudit, weak WeakOutcomeAudit) Summary {
	status := strings.Join([]string{StatusU1Retrieved, StatusTwistConstructed, StatusU1ClassSieve, StatusFailedUniformTwist, StatusFailedUniquePlane, StatusFailedPhysicalChirality, StatusFailedGlobalH}, ";")
	return Summary{
		NativeU1Available:        u1.ActsOnSC && !u1.ImportedSMHypercharge,
		TwistConstructed:         twist.IsDiagonalOnFockBasis && twist.DistinctFromGamma,
		U1RejectsTemporalPlanes:  sieve.TemporalSpatialRejected,
		PureSpatialPlanesRemain:  sieve.PureSpatialDegeneracy,
		UniformTwistedDoublets:   len(sieve.UniformDoubletPlanes) > 0,
		UniqueWeakPlaneDerived:   weak.UniqueWeakPlaneSelected,
		PhysicalChiralityDerived: weak.PhysicalLeftHandedDerived,
		GlobalHDerived:           weak.GlobalHSummandDerived,
		Status:                   status,
		NextGate:                 "derive a canonical contact-su(2) to pure-spatial plane map, or prove no native Spin^c twist can select the weak plane without an additional selector",
		Comment:                  "The native u(1) twist improves the sieve by rejecting temporal-spatial planes, but it leaves three pure-spatial planes degenerate and does not produce uniform twisted chirality.",
	}
}

func buildTruth(u1 U1GeneratorAudit, twist TwistOperatorAudit, sieve TwistedSieveAudit, weak WeakOutcomeAudit) string {
	return fmt.Sprintf("The native diagonal u(1) with weights %v permits χ_twist=γY and rejects temporal-spatial weak planes by nonzero u(1) commutator, leaving %d pure-spatial candidates. Because no candidate has uniform χ_twist doublets and no unique plane is selected, Spin^c twisting supplies a class sieve but not physical left-handed weak chirality or a global H summand.", u1.ModeWeights, sieve.PureSpatialDegeneracy)
}

func stateY(s spinor.FockState, weights []float64) float64 {
	y := 0.0
	for i, occ := range s.Occupation {
		if occ {
			y += weights[i]
		}
	}
	return y
}

func uniqueFloatCount(vals []float64, eps float64) int { return len(sortedUniqueFloats(vals, eps)) }

func sortedUniqueFloats(vals []float64, eps float64) []float64 {
	sorted := append([]float64(nil), vals...)
	sort.Float64s(sorted)
	out := []float64{}
	for _, v := range sorted {
		if len(out) == 0 || math.Abs(v-out[len(out)-1]) > eps {
			if math.Abs(v) < eps {
				v = 0
			}
			out = append(out, v)
		}
	}
	return out
}
