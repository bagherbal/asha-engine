// Package orientationtruechirality implements Gate 239:
// Orientation Operator (chi) / True Chirality Derivation Audit.
//
// Gate 238 proved that raw occupation parity gamma=(-1)^N is a valid finite
// grading but not the Standard Model chiral selector: every candidate weak
// plane has four even and four odd doublet dimensions.  Gate 239 therefore
// audits the next possible selector: a finite orientation operator chi derived
// from the Clifford volume element or from the scalar fundamental class tau_eta.
//
// The result is intentionally conservative.  In the current Fock realization,
// the Clifford-volume/chirality candidate on Lambda*(W) is proportional to the
// same occupation parity gamma already audited in Gate 238.  It is a legitimate
// orientation grading, but it has the same eigenspaces and therefore cannot
// improve the weak-plane sieve.  The tau_eta fundamental class supplies signed
// scalar-bundle trace data, not an endomorphism of the complexified Fock spinor
// S_C; no canonical pullback from tau_eta to a 32-real-dimensional chirality
// operator is derived.  Thus physical chirality remains unconstructed.
package orientationtruechirality

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/chiralweakselector"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

const (
	AuditID = "GATE239-ORIENTATION-TRUE-CHIRALITY-DERIVATION-AUDIT"

	StatusVolumeOrientationPreflight       = "CONDITIONAL_SUPPORT_CLIFFORD_VOLUME_ORIENTATION_PREFLIGHT"
	StatusTauEtaOrientationFunctional      = "CONDITIONAL_SUPPORT_TAU_ETA_ORIENTATION_FUNCTIONAL_INHERITED"
	StatusFailedDistinctChi                = "FAILED_ROUTE_DISTINCT_ORIENTATION_CHI_DERIVATION"
	StatusFailedTauEtaPullback             = "FAILED_ROUTE_TAU_ETA_TO_SC_OPERATOR_PULLBACK"
	StatusFailedTrueChiralPlaneSelection   = "FAILED_ROUTE_TRUE_CHIRALITY_PLANE_SELECTION"
	StatusFailedLeftHandedWeakDerivation   = "FAILED_ROUTE_LEFT_HANDED_WEAK_ACTION_DERIVATION"
	StatusFailedGlobalHSummandStillBlocked = "FAILED_ROUTE_GLOBAL_H_SUMMAND_STILL_UNSELECTED"
)

type OrientationOperatorAudit struct {
	CandidateName          string
	Source                 string
	ActsOnSC               bool
	DimensionComplex       int
	DimensionReal          int
	VolumeElementAvailable bool
	EquivalentToGamma      bool
	CommutesWithGamma      bool
	AntiCommutesWithGamma  bool
	DistinctFromGamma      bool
	DistinctEigenspaces    bool
	EvenEigenDimC          int
	OddEigenDimC           int
	ManualSignAdjusted     bool
	Verdict                string
}

type TauEtaPullbackAudit struct {
	TauEtaDegrees             []int
	OrientationFlippedDegrees []int
	FunctionalOnScalarBundle  bool
	EndomorphismOnSC          bool
	CanonicalPullbackDerived  bool
	GaugeProjectionMapDerived bool
	CanActAsChiralityOperator bool
	Verdict                   string
}

type PlaneChiAudit struct {
	Plane               string
	ModeIndices         []int
	PlaneClass          string
	DoubletPlusDimC     int
	DoubletMinusDimC    int
	SingletPlusDimC     int
	SingletMinusDimC    int
	DoubletsUniformChi  bool
	SingletsUniformChi  bool
	SU2PreservesChi     bool
	SU2ActsOnlyOnOneChi bool
	SelectedByChi       bool
	Verdict             string
}

type ChiPlaneSieveAudit struct {
	CandidatePlanes        int
	UniformDoubletPlanes   int
	UniformSingletPlanes   int
	SelectedPlanes         []string
	ChiBreaksDegeneracy    bool
	AllPlanesSameCounts    bool
	TemporalSpatialClasses bool
	Verdict                string
}

type GradingComparisonAudit struct {
	GammaName                string
	ChiName                  string
	GammaEvenDimC            int
	GammaOddDimC             int
	ChiPlusDimC              int
	ChiMinusDimC             int
	SameSpectrum             bool
	SameEigenspaces          bool
	OperatorsCommute         bool
	OperatorsAntiCommute     bool
	PhysicalChiralityDerived bool
	Verdict                  string
}

type WeakOutcomeAudit struct {
	Gate238GammaSelectorFailed      bool
	VolumeChiImprovesGate238        bool
	TauEtaSuppliesOperator          bool
	UniqueWeakPlaneSelected         bool
	PhysicalLeftHandedActionDerived bool
	GlobalHSummandDerived           bool
	Verdict                         string
}

type FirewallAudit struct {
	AdjustedChiSignsToFit    bool
	ImportedSMGamma5         bool
	ImportedConnesChirality  bool
	ForcedWeakPlane          bool
	PromotedTauEtaToOperator bool
	ClaimedLeftHandedAction  bool
	ClaimedGlobalH           bool
	FiniteCorePolluted       bool
	Verdict                  string
}

type Summary struct {
	VolumeOrientationAvailable bool
	ChiDistinctFromGamma       bool
	TauEtaPullbackDerived      bool
	UniformChiDoublets         bool
	ChiSelectsPlane            bool
	PhysicalChiralityDerived   bool
	GlobalHDerived             bool
	Status                     string
	NextGate                   string
	Comment                    string
}

type Analysis struct {
	Previous       chiralweakselector.Analysis
	Orientation    OrientationOperatorAudit
	TauEta         TauEtaPullbackAudit
	Comparison     GradingComparisonAudit
	Planes         []PlaneChiAudit
	Sieve          ChiPlaneSieveAudit
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
		prev, err := chiralweakselector.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 238 predecessor: %w", err)
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

func Build(prev chiralweakselector.Analysis, f spinor.FockSpace) (Analysis, error) {
	if f.ModeCount() != 4 || f.StateCount() != 16 {
		return Analysis{}, fmt.Errorf("Gate 239 requires native four-mode 16-state carrier, got modes=%d states=%d", f.ModeCount(), f.StateCount())
	}
	orientation := auditVolumeOrientation(f)
	tau := auditTauEtaPullback()
	comparison := compareGradings(f, orientation)
	planes := auditPlanesWithChi(f, orientation)
	sieve := auditChiSieve(planes)
	weak := auditWeak(prev, orientation, tau, sieve)
	fw := auditFirewall()
	sum := summarize(orientation, tau, sieve, weak)
	truth := buildTruth(orientation, tau, comparison, planes, sieve, weak)
	return Analysis{Previous: prev, Orientation: orientation, TauEta: tau, Comparison: comparison, Planes: planes, Sieve: sieve, Weak: weak, Firewall: fw, Summary: sum, TruthStatement: truth}, nil
}

func auditVolumeOrientation(f spinor.FockSpace) OrientationOperatorAudit {
	plus, minus := 0, 0
	for _, s := range f.States {
		// In the exterior/Fock realization of the complexified spinor, the finite
		// volume/chirality candidate is proportional to the parity operator on
		// Lambda*(W).  The overall phase/sign is conventional; it does not change
		// eigenspaces.  We choose + on even degree and - on odd degree.
		if orientationEigenvalue(s) > 0 {
			plus++
		} else {
			minus++
		}
	}
	return OrientationOperatorAudit{
		CandidateName:          "χ_vol = Clifford-volume candidate on Λ*(W)",
		Source:                 "Cl(1,7) volume-element preflight represented on the four-mode exterior spinor",
		ActsOnSC:               true,
		DimensionComplex:       f.StateCount(),
		DimensionReal:          2 * f.StateCount(),
		VolumeElementAvailable: true,
		EquivalentToGamma:      true,
		CommutesWithGamma:      true,
		AntiCommutesWithGamma:  false,
		DistinctFromGamma:      false,
		DistinctEigenspaces:    false,
		EvenEigenDimC:          plus,
		OddEigenDimC:           minus,
		ManualSignAdjusted:     false,
		Verdict:                StatusVolumeOrientationPreflight + "; " + StatusFailedDistinctChi,
	}
}

func auditTauEtaPullback() TauEtaPullbackAudit {
	return TauEtaPullbackAudit{
		TauEtaDegrees:             []int{2, -2, 1},
		OrientationFlippedDegrees: []int{-2, 2, -1},
		FunctionalOnScalarBundle:  true,
		EndomorphismOnSC:          false,
		CanonicalPullbackDerived:  false,
		GaugeProjectionMapDerived: false,
		CanActAsChiralityOperator: false,
		Verdict:                   StatusTauEtaOrientationFunctional + "; " + StatusFailedTauEtaPullback,
	}
}

func compareGradings(f spinor.FockSpace, orientation OrientationOperatorAudit) GradingComparisonAudit {
	even, odd := 0, 0
	for _, s := range f.States {
		if s.ExcitationNumber()%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return GradingComparisonAudit{
		GammaName:                "γ=(-1)^N occupation parity",
		ChiName:                  orientation.CandidateName,
		GammaEvenDimC:            even,
		GammaOddDimC:             odd,
		ChiPlusDimC:              orientation.EvenEigenDimC,
		ChiMinusDimC:             orientation.OddEigenDimC,
		SameSpectrum:             even == orientation.EvenEigenDimC && odd == orientation.OddEigenDimC,
		SameEigenspaces:          orientation.EquivalentToGamma,
		OperatorsCommute:         orientation.CommutesWithGamma,
		OperatorsAntiCommute:     orientation.AntiCommutesWithGamma,
		PhysicalChiralityDerived: false,
		Verdict:                  StatusFailedDistinctChi,
	}
}

func orientationEigenvalue(s spinor.FockState) int {
	if s.ExcitationNumber()%2 == 0 {
		return +1
	}
	return -1
}

func auditPlanesWithChi(f spinor.FockSpace, orientation OrientationOperatorAudit) []PlaneChiAudit {
	out := []PlaneChiAudit{}
	for i := 0; i < f.ModeCount(); i++ {
		for j := i + 1; j < f.ModeCount(); j++ {
			dp, dm, sp, sm := 0, 0, 0, 0
			for _, s := range f.States {
				occupiedInPlane := 0
				if s.Occupation[i] {
					occupiedInPlane++
				}
				if s.Occupation[j] {
					occupiedInPlane++
				}
				plus := orientationEigenvalue(s) > 0
				switch occupiedInPlane {
				case 1:
					if plus {
						dp++
					} else {
						dm++
					}
				case 0, 2:
					if plus {
						sp++
					} else {
						sm++
					}
				}
			}
			class := "pure-spatial"
			if f.Modes[i].Kind == spinor.TemporalMode || f.Modes[j].Kind == spinor.TemporalMode {
				class = "temporal-spatial"
			}
			uniformD := dp == 0 || dm == 0
			uniformS := sp == 0 || sm == 0
			actsOne := uniformD && (dp+dm == 8)
			plane := fmt.Sprintf("U={%s,%s}", f.Modes[i].Name, f.Modes[j].Name)
			out = append(out, PlaneChiAudit{
				Plane:               plane,
				ModeIndices:         []int{i, j},
				PlaneClass:          class,
				DoubletPlusDimC:     dp,
				DoubletMinusDimC:    dm,
				SingletPlusDimC:     sp,
				SingletMinusDimC:    sm,
				DoubletsUniformChi:  uniformD,
				SingletsUniformChi:  uniformS,
				SU2PreservesChi:     orientation.EquivalentToGamma,
				SU2ActsOnlyOnOneChi: actsOne,
				SelectedByChi:       actsOne,
				Verdict:             StatusFailedTrueChiralPlaneSelection,
			})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Plane < out[j].Plane })
	return out
}

func auditChiSieve(planes []PlaneChiAudit) ChiPlaneSieveAudit {
	selected := []string{}
	uniformD, uniformS := 0, 0
	same := true
	var dp, dm, sp, sm int
	if len(planes) > 0 {
		dp, dm, sp, sm = planes[0].DoubletPlusDimC, planes[0].DoubletMinusDimC, planes[0].SingletPlusDimC, planes[0].SingletMinusDimC
	}
	temporalSpatial, pureSpatial := 0, 0
	for _, p := range planes {
		if p.DoubletsUniformChi {
			uniformD++
		}
		if p.SingletsUniformChi {
			uniformS++
		}
		if p.SelectedByChi {
			selected = append(selected, p.Plane)
		}
		if p.DoubletPlusDimC != dp || p.DoubletMinusDimC != dm || p.SingletPlusDimC != sp || p.SingletMinusDimC != sm {
			same = false
		}
		if p.PlaneClass == "temporal-spatial" {
			temporalSpatial++
		} else {
			pureSpatial++
		}
	}
	return ChiPlaneSieveAudit{
		CandidatePlanes:        len(planes),
		UniformDoubletPlanes:   uniformD,
		UniformSingletPlanes:   uniformS,
		SelectedPlanes:         selected,
		ChiBreaksDegeneracy:    len(selected) == 1,
		AllPlanesSameCounts:    same,
		TemporalSpatialClasses: temporalSpatial == 3 && pureSpatial == 3,
		Verdict:                StatusFailedTrueChiralPlaneSelection,
	}
}

func auditWeak(prev chiralweakselector.Analysis, orientation OrientationOperatorAudit, tau TauEtaPullbackAudit, sieve ChiPlaneSieveAudit) WeakOutcomeAudit {
	return WeakOutcomeAudit{
		Gate238GammaSelectorFailed:      !prev.Summary.GammaSelectsPlane && !prev.Summary.PhysicalLeftActionDerived,
		VolumeChiImprovesGate238:        orientation.DistinctFromGamma && sieve.ChiBreaksDegeneracy,
		TauEtaSuppliesOperator:          tau.CanActAsChiralityOperator,
		UniqueWeakPlaneSelected:         sieve.ChiBreaksDegeneracy,
		PhysicalLeftHandedActionDerived: false,
		GlobalHSummandDerived:           false,
		Verdict:                         StatusFailedLeftHandedWeakDerivation + "; " + StatusFailedGlobalHSummandStillBlocked,
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		AdjustedChiSignsToFit:    false,
		ImportedSMGamma5:         false,
		ImportedConnesChirality:  false,
		ForcedWeakPlane:          false,
		PromotedTauEtaToOperator: false,
		ClaimedLeftHandedAction:  false,
		ClaimedGlobalH:           false,
		FiniteCorePolluted:       false,
		Verdict:                  "FIREWALL_PRESERVED_NO_FORCED_TRUE_CHIRALITY_OPERATOR",
	}
}

func summarize(o OrientationOperatorAudit, tau TauEtaPullbackAudit, sieve ChiPlaneSieveAudit, weak WeakOutcomeAudit) Summary {
	status := strings.Join([]string{StatusVolumeOrientationPreflight, StatusTauEtaOrientationFunctional, StatusFailedDistinctChi, StatusFailedTauEtaPullback, StatusFailedTrueChiralPlaneSelection, StatusFailedLeftHandedWeakDerivation, StatusFailedGlobalHSummandStillBlocked}, ";")
	return Summary{
		VolumeOrientationAvailable: o.VolumeElementAvailable && o.ActsOnSC,
		ChiDistinctFromGamma:       o.DistinctFromGamma,
		TauEtaPullbackDerived:      tau.CanonicalPullbackDerived,
		UniformChiDoublets:         sieve.UniformDoubletPlanes > 0,
		ChiSelectsPlane:            sieve.ChiBreaksDegeneracy,
		PhysicalChiralityDerived:   weak.PhysicalLeftHandedActionDerived,
		GlobalHDerived:             weak.GlobalHSummandDerived,
		Status:                     status,
		NextGate:                   "derive a nontrivial orientation pullback/intertwiner from contact eta data to S_C, or construct the faithful finite algebra/order-one calculus that distinguishes physical chirality from Fock parity",
		Comment:                    "The only current orientation endomorphism on S_C is the Clifford-volume candidate, which is proportional to occupation parity. Tau_eta is signed orientation data but not an S_C operator. The weak-plane degeneracy therefore survives.",
	}
}

func buildTruth(o OrientationOperatorAudit, tau TauEtaPullbackAudit, cmp GradingComparisonAudit, planes []PlaneChiAudit, sieve ChiPlaneSieveAudit, weak WeakOutcomeAudit) string {
	return fmt.Sprintf("Gate 239 audits orientation-derived chirality after Gate 238. The Clifford-volume candidate %q acts on S_C and has %d/%d eigenspaces, but it is equivalent to gamma=(-1)^N and has the same eigenspaces. Re-running the six-plane weak sieve gives uniform doublet planes=%d and selected planes=%v. The tau_eta signed degrees %v are inherited as scalar-bundle orientation data, but no canonical pullback to an endomorphism of S_C is derived. Therefore true Standard Model chirality, a unique weak plane, and the global H summand remain unselected.", o.CandidateName, cmp.ChiPlusDimC, cmp.ChiMinusDimC, sieve.UniformDoubletPlanes, sieve.SelectedPlanes, tau.TauEtaDegrees)
}
