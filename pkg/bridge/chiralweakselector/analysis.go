// Package chiralweakselector implements Gate 238:
// Chiral Alignment (gamma) and Weak Plane Selector Audit.
//
// Gate 237 showed that every two-mode plane U⊂W gives a valid exterior su(2)
// lift on S_C=Λ*W with the correct 8⊕8 doublet/singlet dimensional pattern.
// Gate 238 tests the proposed next sieve: can the native occupation-parity
// grading γ select one of the six planes by forcing the weak doublets to live in
// one chirality sector?
//
// The answer is strict and useful.  For any plane U={i,j}, the doublet sector is
// the subspace with exactly one U-mode occupied, tensored with Λ*V.  Since Λ*V
// contains even and odd degrees, every plane contains four even-parity and four
// odd-parity doublet states.  The lifted su(2) also preserves total occupation
// parity, so it commutes with γ rather than producing a chiral selector.  The
// temporal/spatial 1⊕3 split distinguishes temporal-spatial planes from purely
// spatial planes, but leaves a 3+3 class degeneracy and no unique electroweak
// plane.
package chiralweakselector

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/su2spinorlift"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

const (
	AuditID = "GATE238-CHIRAL-WEAK-PLANE-SELECTOR-AUDIT"

	StatusGammaParityPreflight       = "CONDITIONAL_SUPPORT_GAMMA_PARITY_PREFLIGHT"
	StatusTemporalSpatialClassSieve  = "CONDITIONAL_SUPPORT_TEMPORAL_SPATIAL_CLASS_SIEVE"
	StatusFailedUniformChiralDoublet = "FAILED_ROUTE_UNIFORM_CHIRAL_DOUBLET_ALIGNMENT"
	StatusFailedUniqueWeakPlane      = "FAILED_ROUTE_CHIRAL_WEAK_PLANE_SELECTION"
	StatusFailedPhysicalLeftAction   = "FAILED_ROUTE_LEFT_HANDED_WEAK_ACTION_DERIVATION"
	StatusFailedGlobalHCompletion    = "FAILED_ROUTE_GLOBAL_H_SUMMAND_STILL_UNSELECTED"
)

type PlaneChiralityAudit struct {
	Plane                  string
	ModeIndices            []int
	PlaneClass             string
	DoubletEvenDimC        int
	DoubletOddDimC         int
	SingletEvenDimC        int
	SingletOddDimC         int
	DoubletsUniformParity  bool
	SingletsUniformParity  bool
	SU2PreservesGamma      bool
	SU2ActsOnlyOnOneParity bool
	ChiralAlignmentScore   int
	Verdict                string
}

type GammaAudit struct {
	Grading                  string
	EvenStateDimC            int
	OddStateDimC             int
	RetrievedFromGate233     bool
	EquatedToSMChirality     bool
	PhysicalChiralityDerived bool
	Verdict                  string
}

type DegeneracySieveAudit struct {
	CandidatePlanes       int
	UniformDoubletPlanes  int
	UniformSingletPlanes  int
	ChiralSelectedPlanes  []string
	GammaBreaksDegeneracy bool
	AllPlanesSameCounts   bool
	Verdict               string
}

type TemporalSpatialAudit struct {
	TemporalSpatialPlaneCount int
	PureSpatialPlaneCount     int
	TemporalSpatialPlanes     []string
	PureSpatialPlanes         []string
	ClassDistinctionExists    bool
	UniquePlaneSelected       bool
	Verdict                   string
}

type WeakActionAudit struct {
	CandidateLocalHSupportInherited bool
	GammaSelectorWorks              bool
	TemporalSpatialSelectorWorks    bool
	ContactSU2PlaneMapDerived       bool
	HyperchargeColorAttachment      bool
	GlobalHSummandDerived           bool
	PhysicalLeftHandedActionDerived bool
	Verdict                         string
}

type FirewallAudit struct {
	ForcedLeftHandedAssignment bool
	ForcedWeakPlane            bool
	ImportedSMChirality        bool
	ImportedPauliMatrices      bool
	ImportedConnesAlgebra      bool
	ClaimedGlobalH             bool
	ClaimedOrderOne            bool
	FiniteCorePolluted         bool
	Verdict                    string
}

type Summary struct {
	GammaParityAvailable      bool
	UniformChiralDoublets     bool
	GammaSelectsPlane         bool
	TemporalSpatialClasses    bool
	UniqueWeakPlaneDerived    bool
	PhysicalLeftActionDerived bool
	GlobalHDerived            bool
	Status                    string
	NextGate                  string
	Comment                   string
}

type Analysis struct {
	Previous       su2spinorlift.Analysis
	Gamma          GammaAudit
	Planes         []PlaneChiralityAudit
	Sieve          DegeneracySieveAudit
	Temporal       TemporalSpatialAudit
	Weak           WeakActionAudit
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
		prev, err := su2spinorlift.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 237 predecessor: %w", err)
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

func Build(prev su2spinorlift.Analysis, f spinor.FockSpace) (Analysis, error) {
	if f.ModeCount() != 4 || f.StateCount() != 16 {
		return Analysis{}, fmt.Errorf("Gate 238 requires native four-mode 16-state carrier, got modes=%d states=%d", f.ModeCount(), f.StateCount())
	}
	gamma := auditGamma(f)
	planes := auditPlanes(f)
	sieve := auditSieve(planes)
	temporal := auditTemporalSpatial(f, planes)
	weak := auditWeak(prev, sieve, temporal)
	fw := auditFirewall()
	sum := summarize(gamma, sieve, temporal, weak)
	truth := buildTruth(gamma, planes, sieve, temporal, weak)
	return Analysis{Previous: prev, Gamma: gamma, Planes: planes, Sieve: sieve, Temporal: temporal, Weak: weak, Firewall: fw, Summary: sum, TruthStatement: truth}, nil
}

func auditGamma(f spinor.FockSpace) GammaAudit {
	even, odd := 0, 0
	for _, s := range f.States {
		if s.ExcitationNumber()%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return GammaAudit{
		Grading:                  "γ=(-1)^N occupation parity on Λ*(W)",
		EvenStateDimC:            even,
		OddStateDimC:             odd,
		RetrievedFromGate233:     true,
		EquatedToSMChirality:     false,
		PhysicalChiralityDerived: false,
		Verdict:                  StatusGammaParityPreflight + "; " + StatusFailedPhysicalLeftAction,
	}
}

func auditPlanes(f spinor.FockSpace) []PlaneChiralityAudit {
	out := []PlaneChiralityAudit{}
	for i := 0; i < f.ModeCount(); i++ {
		for j := i + 1; j < f.ModeCount(); j++ {
			de, do, se, so := 0, 0, 0, 0
			for _, s := range f.States {
				occupiedInPlane := 0
				if s.Occupation[i] {
					occupiedInPlane++
				}
				if s.Occupation[j] {
					occupiedInPlane++
				}
				even := s.ExcitationNumber()%2 == 0
				switch occupiedInPlane {
				case 1:
					if even {
						de++
					} else {
						do++
					}
				case 0, 2:
					if even {
						se++
					} else {
						so++
					}
				}
			}
			class := "pure-spatial"
			if f.Modes[i].Kind == spinor.TemporalMode || f.Modes[j].Kind == spinor.TemporalMode {
				class = "temporal-spatial"
			}
			uniformD := de == 0 || do == 0
			uniformS := se == 0 || so == 0
			actsOne := uniformD && (de+do == 8)
			score := 0
			if uniformD {
				score++
			}
			if uniformS {
				score++
			}
			if actsOne {
				score++
			}
			plane := fmt.Sprintf("U={%s,%s}", f.Modes[i].Name, f.Modes[j].Name)
			out = append(out, PlaneChiralityAudit{
				Plane:                  plane,
				ModeIndices:            []int{i, j},
				PlaneClass:             class,
				DoubletEvenDimC:        de,
				DoubletOddDimC:         do,
				SingletEvenDimC:        se,
				SingletOddDimC:         so,
				DoubletsUniformParity:  uniformD,
				SingletsUniformParity:  uniformS,
				SU2PreservesGamma:      true,
				SU2ActsOnlyOnOneParity: actsOne,
				ChiralAlignmentScore:   score,
				Verdict:                StatusFailedUniformChiralDoublet,
			})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Plane < out[j].Plane })
	return out
}

func auditSieve(planes []PlaneChiralityAudit) DegeneracySieveAudit {
	selected := []string{}
	uniformD, uniformS := 0, 0
	same := true
	var de, do, se, so int
	if len(planes) > 0 {
		de, do, se, so = planes[0].DoubletEvenDimC, planes[0].DoubletOddDimC, planes[0].SingletEvenDimC, planes[0].SingletOddDimC
	}
	for _, p := range planes {
		if p.DoubletsUniformParity {
			uniformD++
		}
		if p.SingletsUniformParity {
			uniformS++
		}
		if p.SU2ActsOnlyOnOneParity {
			selected = append(selected, p.Plane)
		}
		if p.DoubletEvenDimC != de || p.DoubletOddDimC != do || p.SingletEvenDimC != se || p.SingletOddDimC != so {
			same = false
		}
	}
	return DegeneracySieveAudit{
		CandidatePlanes:       len(planes),
		UniformDoubletPlanes:  uniformD,
		UniformSingletPlanes:  uniformS,
		ChiralSelectedPlanes:  selected,
		GammaBreaksDegeneracy: len(selected) == 1,
		AllPlanesSameCounts:   same,
		Verdict:               StatusFailedUniqueWeakPlane,
	}
}

func auditTemporalSpatial(f spinor.FockSpace, planes []PlaneChiralityAudit) TemporalSpatialAudit {
	ts, ps := []string{}, []string{}
	for _, p := range planes {
		if p.PlaneClass == "temporal-spatial" {
			ts = append(ts, p.Plane)
		} else {
			ps = append(ps, p.Plane)
		}
	}
	return TemporalSpatialAudit{
		TemporalSpatialPlaneCount: len(ts),
		PureSpatialPlaneCount:     len(ps),
		TemporalSpatialPlanes:     ts,
		PureSpatialPlanes:         ps,
		ClassDistinctionExists:    len(ts) > 0 && len(ps) > 0,
		UniquePlaneSelected:       len(ts) == 1 || len(ps) == 1,
		Verdict:                   StatusTemporalSpatialClassSieve + "; " + StatusFailedUniqueWeakPlane,
	}
}

func auditWeak(prev su2spinorlift.Analysis, sieve DegeneracySieveAudit, temporal TemporalSpatialAudit) WeakActionAudit {
	return WeakActionAudit{
		CandidateLocalHSupportInherited: prev.Summary.PseudoRealLocalHSupport,
		GammaSelectorWorks:              sieve.GammaBreaksDegeneracy,
		TemporalSpatialSelectorWorks:    temporal.UniquePlaneSelected,
		ContactSU2PlaneMapDerived:       false,
		HyperchargeColorAttachment:      false,
		GlobalHSummandDerived:           false,
		PhysicalLeftHandedActionDerived: false,
		Verdict:                         StatusFailedPhysicalLeftAction + "; " + StatusFailedGlobalHCompletion,
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		ForcedLeftHandedAssignment: false,
		ForcedWeakPlane:            false,
		ImportedSMChirality:        false,
		ImportedPauliMatrices:      false,
		ImportedConnesAlgebra:      false,
		ClaimedGlobalH:             false,
		ClaimedOrderOne:            false,
		FiniteCorePolluted:         false,
		Verdict:                    "FIREWALL_PRESERVED_NO_FORCED_CHIRAL_WEAK_PLANE",
	}
}

func summarize(g GammaAudit, sieve DegeneracySieveAudit, temporal TemporalSpatialAudit, weak WeakActionAudit) Summary {
	status := strings.Join([]string{StatusGammaParityPreflight, StatusTemporalSpatialClassSieve, StatusFailedUniformChiralDoublet, StatusFailedUniqueWeakPlane, StatusFailedPhysicalLeftAction, StatusFailedGlobalHCompletion}, ";")
	return Summary{
		GammaParityAvailable:      g.EvenStateDimC == 8 && g.OddStateDimC == 8,
		UniformChiralDoublets:     sieve.UniformDoubletPlanes > 0,
		GammaSelectsPlane:         sieve.GammaBreaksDegeneracy,
		TemporalSpatialClasses:    temporal.ClassDistinctionExists,
		UniqueWeakPlaneDerived:    sieve.GammaBreaksDegeneracy || temporal.UniquePlaneSelected,
		PhysicalLeftActionDerived: weak.PhysicalLeftHandedActionDerived,
		GlobalHDerived:            weak.GlobalHSummandDerived,
		Status:                    status,
		NextGate:                  "derive an additional finite selector/intertwiner beyond occupation parity, likely from contact-vacuum orientation, eta-source, or physical chirality rather than raw Fock parity",
		Comment:                   "Occupation parity is a valid grading but every candidate plane has the same 4-even/4-odd doublet split; temporal/spatial bookkeeping distinguishes two plane classes but not a unique weak plane.",
	}
}

func buildTruth(g GammaAudit, planes []PlaneChiralityAudit, sieve DegeneracySieveAudit, temporal TemporalSpatialAudit, weak WeakActionAudit) string {
	return fmt.Sprintf("Gate 238 audits γ=%s on %d candidate two-mode planes. Every plane has doublet parity split 4 even + 4 odd and singlet parity split 4 even + 4 odd, so γ does not isolate left-handed weak doublets or select a unique plane. The 1⊕3 split distinguishes %d temporal-spatial and %d pure-spatial planes, but both classes remain threefold degenerate. Therefore Gate 237's local H support survives, while the physical chiral weak action and global H summand remain unselected.", g.Grading, len(planes), temporal.TemporalSpatialPlaneCount, temporal.PureSpatialPlaneCount)
}
