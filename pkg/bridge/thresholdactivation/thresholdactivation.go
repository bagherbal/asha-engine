// Package thresholdactivation audits whether finite threshold anchors have a
// derived activation / decoupling rule.
//
// A finite spectral number is not a physical threshold by itself.  Even after a
// sector-level gauge representation is known, the engine still needs a rule
// saying whether that sector remains continuum-active, is integrated out above a
// matching scale, or is only a vacuum-frustration invariant.  Gate 46 audits
// that missing rule without inserting a physical mass scale or observed
// threshold.
package thresholdactivation

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/threshold"
	"github.com/bagherbal/asha-engine/pkg/bridge/thresholdrep"
)

type ActivationStatus string

const (
	ContinuumFieldCandidate ActivationStatus = "continuum-field-candidate"
	ThresholdOpen           ActivationStatus = "threshold-open"
	IntegratedOutOpen       ActivationStatus = "integrated-out-open"
	VacuumFrustrationOnly   ActivationStatus = "vacuum-frustration-only"
	Unclassified            ActivationStatus = "unclassified"
)

type ActivationDecision struct {
	Assignment thresholdrep.CandidateAssignment
	Status     ActivationStatus
	Reason     string

	ContinuumActiveDerived bool
	IntegratedOutDerived   bool
	HeavyThresholdDerived  bool
	DecouplingRuleDerived  bool
	CanCorrectBeta         bool
}

type Analysis struct {
	Representation thresholdrep.Analysis

	CandidateCount int
	Decisions      []ActivationDecision

	ContinuumFieldCandidateCount int
	ContinuumActiveDerivedCount  int
	IntegratedOutDerivedCount    int
	HeavyThresholdDerivedCount   int
	VacuumOnlyCount              int
	UnclassifiedCount            int
	BetaCorrectionAllowedCount   int

	ScalarSectorRemainsContinuumCandidate bool
	ScalarHeavyThresholdDerived           bool
	BGapActivationDerived                 bool
	ContactOverlapActivationDerived       bool
	LeakageClassifiedAsVacuumOnly         bool
	PhysicalMassUnitDerived               bool
	DecouplingRuleDerived                 bool
	ThresholdCorrectedBetaDerived         bool
	HiddenScaleInserted                   bool

	TruthStatement     string
	MinimumMissingData []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		rep, err := thresholdrep.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(rep)
	})
	return defaultValue, defaultErr
}

func Build(rep thresholdrep.Analysis) (Analysis, error) {
	if len(rep.Assignments) == 0 {
		return Analysis{}, fmt.Errorf("threshold activation audit requires representation assignments")
	}
	decisions := make([]ActivationDecision, 0, len(rep.Assignments))
	for _, a := range rep.Assignments {
		decisions = append(decisions, decide(a))
	}
	sort.Slice(decisions, func(i, j int) bool {
		if decisions[i].Status == decisions[j].Status {
			return decisions[i].Assignment.Candidate.Name < decisions[j].Assignment.Candidate.Name
		}
		return decisions[i].Status < decisions[j].Status
	})

	continuumCandidate, continuumDerived, integrated, heavy, vacuumOnly, unclassified, betaAllowed := 0, 0, 0, 0, 0, 0, 0
	scalarCandidate := false
	scalarHeavy := false
	bgap := false
	overlap := false
	leakageVacuum := false
	decoupling := false
	for _, d := range decisions {
		switch d.Status {
		case ContinuumFieldCandidate:
			continuumCandidate++
		case VacuumFrustrationOnly:
			vacuumOnly++
		case Unclassified, ThresholdOpen, IntegratedOutOpen:
			unclassified++
		}
		if d.ContinuumActiveDerived {
			continuumDerived++
		}
		if d.IntegratedOutDerived {
			integrated++
		}
		if d.HeavyThresholdDerived {
			heavy++
		}
		if d.CanCorrectBeta {
			betaAllowed++
		}
		if d.DecouplingRuleDerived {
			decoupling = true
		}
		switch d.Assignment.Candidate.Kind {
		case threshold.ScalarActiveCandidate, threshold.RadialCandidate:
			if d.Status == ContinuumFieldCandidate {
				scalarCandidate = true
			}
			if d.HeavyThresholdDerived {
				scalarHeavy = true
			}
		case threshold.BGapCandidate:
			if d.HeavyThresholdDerived || d.ContinuumActiveDerived || d.IntegratedOutDerived {
				bgap = true
			}
		case threshold.ContactOverlapCandidate:
			if d.HeavyThresholdDerived || d.ContinuumActiveDerived || d.IntegratedOutDerived {
				overlap = true
			}
		case threshold.LeakageCandidate:
			if d.Status == VacuumFrustrationOnly {
				leakageVacuum = true
			}
		}
	}

	return Analysis{
		Representation:                        rep,
		CandidateCount:                        len(decisions),
		Decisions:                             decisions,
		ContinuumFieldCandidateCount:          continuumCandidate,
		ContinuumActiveDerivedCount:           continuumDerived,
		IntegratedOutDerivedCount:             integrated,
		HeavyThresholdDerivedCount:            heavy,
		VacuumOnlyCount:                       vacuumOnly,
		UnclassifiedCount:                     unclassified,
		BetaCorrectionAllowedCount:            betaAllowed,
		ScalarSectorRemainsContinuumCandidate: scalarCandidate,
		ScalarHeavyThresholdDerived:           scalarHeavy,
		BGapActivationDerived:                 bgap,
		ContactOverlapActivationDerived:       overlap,
		LeakageClassifiedAsVacuumOnly:         leakageVacuum,
		PhysicalMassUnitDerived:               false,
		DecouplingRuleDerived:                 decoupling,
		ThresholdCorrectedBetaDerived:         betaAllowed > 0,
		HiddenScaleInserted:                   false,
		TruthStatement:                        "The scalar/contact doublet is a continuum field candidate at sector level, while the leakage invariant is vacuum-frustration-only. The B-sector gap and contact partial-overlap modes still lack activation and decoupling rules, so no heavy-threshold correction to beta coefficients is allowed.",
		MinimumMissingData: []string{
			"derive a physical unit or boundary scale before any dimensionless anchor becomes a mass threshold",
			"derive an activation rule distinguishing continuum fields, heavy thresholds, regulator modes, and vacuum-frustration invariants",
			"derive decoupling/matching rules for each continuum-active heavy threshold",
			"derive gauge representation assignments for the B-sector and contact partial-overlap modes before they can affect b1,b2,b3",
			"derive whether the scalar/contact doublet is low-energy active, high-scale active, or only a finite source seed",
		},
	}, nil
}

func decide(a thresholdrep.CandidateAssignment) ActivationDecision {
	c := a.Candidate
	switch c.Kind {
	case threshold.ScalarActiveCandidate:
		return ActivationDecision{
			Assignment: a,
			Status:     ContinuumFieldCandidate,
			Reason:     "the scalar/contact active sector has a derived doublet representation and is treated as a field candidate, but no heavy activation mass or decoupling rule is derived",
		}
	case threshold.RadialCandidate:
		return ActivationDecision{
			Assignment: a,
			Status:     ContinuumFieldCandidate,
			Reason:     "the radial scalar response belongs to the scalar effective-potential bridge; it is not yet a separate heavy threshold field",
		}
	case threshold.LeakageCandidate:
		return ActivationDecision{
			Assignment: a,
			Status:     VacuumFrustrationOnly,
			Reason:     "bare contact leakage is a vacuum-frustration invariant, not a continuum field or decoupling threshold",
		}
	case threshold.BGapCandidate:
		return ActivationDecision{
			Assignment: a,
			Status:     ThresholdOpen,
			Reason:     "the B-sector spectral gap is a finite action eigenvalue, but no gauge representation, physical scale, or activation rule is derived",
		}
	case threshold.ContactOverlapCandidate:
		return ActivationDecision{
			Assignment: a,
			Status:     ThresholdOpen,
			Reason:     "contact partial-overlap modes are real finite overlap modes, but it is still open whether they are physical thresholds, regulator modes, or frustration modes",
		}
	default:
		return ActivationDecision{Assignment: a, Status: Unclassified, Reason: "candidate kind has no activation rule"}
	}
}

func FormatDecisions(ds []ActivationDecision, max int) string {
	if max <= 0 || max > len(ds) {
		max = len(ds)
	}
	parts := make([]string, 0, max)
	for i := 0; i < max; i++ {
		d := ds[i]
		parts = append(parts, fmt.Sprintf("%s:%s", d.Assignment.Candidate.Name, d.Status))
	}
	if max < len(ds) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(ds)-max))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}
