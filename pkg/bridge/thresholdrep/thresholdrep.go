// Package thresholdrep audits whether the dimensionless threshold anchors from
// the finite source engine have representation assignments under
// SU(3)c × SU(2)L × U(1)Y.
//
// This is deliberately stricter than the beta-coefficient package.  A spectral
// number cannot correct a beta coefficient merely because it exists.  To become
// a threshold contribution it needs at least:
//
//  1. a finite-to-continuum activation rule,
//  2. a physical mass scale or boundary unit,
//  3. a gauge representation assignment, and
//  4. a decoupling/matching prescription.
//
// Gate 45 only audits item 3.  The scalar/contact active sector has a
// previously derived doublet representation at the sector level.  The B-sector
// gap and contact partial-overlap modes do not yet have gauge representations.
// Therefore threshold-corrected beta coefficients remain forbidden.
package thresholdrep

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/threshold"
)

type AssignmentStatus string

const (
	AssignedSectorLevel AssignmentStatus = "assigned-sector-level"
	AssignedBridgeOnly  AssignmentStatus = "assigned-bridge-only"
	Unassigned          AssignmentStatus = "unassigned"
	NotAThreshold       AssignmentStatus = "not-a-threshold"
)

type GaugeRepresentation struct {
	SU3         string
	SU2         string
	Hypercharge string
	Detail      string
}

type CandidateAssignment struct {
	Candidate threshold.Candidate
	Status    AssignmentStatus
	Rep       GaugeRepresentation
	Reason    string

	IndividualModeAssignment bool
	ContinuumActiveDerived   bool
}

type Analysis struct {
	Threshold threshold.Analysis

	CandidateCount int
	Assignments    []CandidateAssignment

	SectorAssignedCount        int
	IndividualAssignedCount    int
	UnassignedCount            int
	NonThresholdCount          int
	ContinuumActiveCount       int
	ScalarDoubletSectorDerived bool

	BGapRepresentationDerived            bool
	ContactOverlapRepresentationsDerived bool
	LeakageModeRepresentationDerived     bool
	ThresholdCorrectionsAllowed          bool
	ThresholdCorrectedBetaDerived        bool

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
		th, err := threshold.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(th)
	})
	return defaultValue, defaultErr
}

func Build(th threshold.Analysis) (Analysis, error) {
	if len(th.Candidates) == 0 {
		return Analysis{}, fmt.Errorf("threshold candidate audit requires candidate anchors")
	}
	assignments := make([]CandidateAssignment, 0, len(th.Candidates))
	for _, c := range th.Candidates {
		assignments = append(assignments, classify(c))
	}
	sort.Slice(assignments, func(i, j int) bool {
		if assignments[i].Status == assignments[j].Status {
			return assignments[i].Candidate.Name < assignments[j].Candidate.Name
		}
		return assignments[i].Status < assignments[j].Status
	})

	sector, individual, unassigned, nonthreshold, active := 0, 0, 0, 0, 0
	scalarDoublet := false
	for _, a := range assignments {
		switch a.Status {
		case AssignedSectorLevel, AssignedBridgeOnly:
			sector++
		case Unassigned:
			unassigned++
		case NotAThreshold:
			nonthreshold++
		}
		if a.IndividualModeAssignment {
			individual++
		}
		if a.ContinuumActiveDerived {
			active++
		}
		if a.Candidate.Kind == threshold.ScalarActiveCandidate && a.Status == AssignedSectorLevel {
			scalarDoublet = true
		}
	}

	return Analysis{
		Threshold:                            th,
		CandidateCount:                       len(th.Candidates),
		Assignments:                          assignments,
		SectorAssignedCount:                  sector,
		IndividualAssignedCount:              individual,
		UnassignedCount:                      unassigned,
		NonThresholdCount:                    nonthreshold,
		ContinuumActiveCount:                 active,
		ScalarDoubletSectorDerived:           scalarDoublet,
		BGapRepresentationDerived:            false,
		ContactOverlapRepresentationsDerived: false,
		LeakageModeRepresentationDerived:     false,
		ThresholdCorrectionsAllowed:          false,
		ThresholdCorrectedBetaDerived:        false,
		TruthStatement:                       "The scalar/contact active sector has a previously derived scalar-doublet representation at sector level, but the B-sector gap, contact partial-overlap modes, and leakage/frustration invariant do not yet have gauge representation assignments or activation rules. Therefore no threshold correction to b1,b2,b3 is allowed yet.",
		MinimumMissingData: []string{
			"derive a representation for the B-sector positive modes under SU(3)c, SU(2)L, and U(1)Y",
			"derive whether the seven contact partial-overlap modes are physical fields, regulator modes, or vacuum-frustration modes",
			"derive individual scalar active charge/orientation if the real modes are to be threshold-split rather than treated as one doublet sector",
			"derive a continuum activation and decoupling rule for every assigned threshold mode",
			"derive the physical mass unit or boundary scale before any threshold mass enters RG flow",
		},
	}, nil
}

func classify(c threshold.Candidate) CandidateAssignment {
	switch c.Kind {
	case threshold.ScalarActiveCandidate:
		return CandidateAssignment{
			Candidate: c,
			Status:    AssignedSectorLevel,
			Rep: GaugeRepresentation{
				SU3:         "1",
				SU2:         "2",
				Hypercharge: "±1/2 sector",
				Detail:      "the four real active contact directions form one complex scalar doublet at sector level",
			},
			Reason:                   "Gate 20 and Gate 37 derive the scalar/contact doublet sector; individual real-mode threshold charges are not separately selected",
			IndividualModeAssignment: false,
			ContinuumActiveDerived:   false,
		}
	case threshold.RadialCandidate:
		return CandidateAssignment{
			Candidate:                c,
			Status:                   AssignedBridgeOnly,
			Rep:                      GaugeRepresentation{SU3: "1", SU2: "1 after scalar-sector reduction", Hypercharge: "0 after scalar-sector reduction", Detail: "radial scalar response is bridge-level until gauge-eating/continuum scalar theorem is derived"},
			Reason:                   "the radial curvature is a scalar-sector invariant, not a derived heavy threshold field",
			IndividualModeAssignment: false,
			ContinuumActiveDerived:   false,
		}
	case threshold.BGapCandidate:
		return CandidateAssignment{Candidate: c, Status: Unassigned, Reason: "B-sector gap is a Boolean vacuum-action eigenvalue; no gauge representation or activation rule is derived", ContinuumActiveDerived: false}
	case threshold.ContactOverlapCandidate:
		return CandidateAssignment{Candidate: c, Status: Unassigned, Reason: "contact partial-overlap mode has no derived continuum field representation under SU(3)c×SU(2)L×U(1)Y", ContinuumActiveDerived: false}
	case threshold.LeakageCandidate:
		return CandidateAssignment{Candidate: c, Status: NotAThreshold, Reason: "bare contact leakage is a frustration invariant, not a threshold mass or field representation", ContinuumActiveDerived: false}
	default:
		return CandidateAssignment{Candidate: c, Status: Unassigned, Reason: "unknown candidate kind has no representation rule", ContinuumActiveDerived: false}
	}
}

func FormatAssignments(as []CandidateAssignment, max int) string {
	if max <= 0 || max > len(as) {
		max = len(as)
	}
	parts := make([]string, 0, max)
	for i := 0; i < max; i++ {
		a := as[i]
		rep := "no-rep"
		if a.Rep.SU3 != "" || a.Rep.SU2 != "" || a.Rep.Hypercharge != "" {
			rep = fmt.Sprintf("(%s,%s)_%s", empty(a.Rep.SU3), empty(a.Rep.SU2), empty(a.Rep.Hypercharge))
		}
		parts = append(parts, fmt.Sprintf("%s:%s:%s", a.Candidate.Name, a.Status, rep))
	}
	if max < len(as) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(as)-max))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func empty(s string) string {
	if s == "" {
		return "?"
	}
	return s
}
