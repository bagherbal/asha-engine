// Package betamatching attempts the next bridge after finite filtration:
// representation-complete threshold beta matching.
//
// Gate 108 asks whether the current finite engine can construct the map
//
//	finite mode -> SU(3)c × SU(2)L × U(1)Y representation -> Δb_i row.
//
// The answer is partial. The scalar/contact active sector supplies one
// representation-complete sector-level baseline row. The B-sector gap and
// contact partial-overlap modes remain representation-incomplete and cannot
// correct beta coefficients.
package betamatching

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/filtration"
	"github.com/bagherbal/asha-engine/pkg/bridge/threshold"
	"github.com/bagherbal/asha-engine/pkg/bridge/thresholdactivation"
	"github.com/bagherbal/asha-engine/pkg/bridge/thresholdrep"
)

type RowKind string

const (
	SectorBaselineRow RowKind = "sector-baseline"
	IndividualModeRow RowKind = "individual-mode"
	OpenModeRow       RowKind = "open-mode"
	VacuumOnlyRow     RowKind = "vacuum-only"
)

type BetaVector struct{ B1, B2, B3 float64 }

func (v BetaVector) Zero(eps float64) bool {
	return math.Abs(v.B1) <= eps && math.Abs(v.B2) <= eps && math.Abs(v.B3) <= eps
}

type MatchingRow struct {
	Name, ModeKind string
	Kind           RowKind
	Status         thresholdactivation.ActivationStatus
	Assignment     thresholdrep.AssignmentStatus
	Rep            thresholdrep.GaugeRepresentation

	RepresentationComplete, DeltaBDerived, SectorLevel, IndividualModeLevel bool
	ContinuumBaseline, HeavyThreshold, DecouplingDerived, CanCorrectBeta    bool
	ExcludedFromThresholds                                                  bool
	DeltaB                                                                  BetaVector
	Reason                                                                  string
}

type CompletionAttempt struct {
	Name                                                                                                string
	Constructed, RepresentationComplete, ActivationDerived, DecouplingRuleDerived, PhysicalScaleDerived bool
	DeltaBRowDerived, CanCorrectBeta, RejectedAsThresholdPhysics                                        bool
	Detail                                                                                              string
}

type AmbiguityWitness struct {
	Name                     string
	CompatibleWithFiniteData bool
	DeltaB                   BetaVector
	Reason                   string
}

type Analysis struct {
	Filtration filtration.Analysis

	Rows                       []MatchingRow
	CompletionAttempts         []CompletionAttempt
	AmbiguityWitnesses         []AmbiguityWitness
	RepresentationOpenModes    []string
	RepresentationCompleteRows int
	IncompleteOpenRows         int
	SectorBaselineRows         int
	IndividualThresholdRows    int
	VacuumRows                 int
	DeltaBRowsDerived          int
	BetaCorrectionRowsAllowed  int

	ScalarSectorRowConstructed                                                                     bool
	ScalarSectorDeltaB                                                                             BetaVector
	ScalarSectorMatchesBaseline                                                                    bool
	ScalarSectorIsThresholdCorrection                                                              bool
	BGapRepresentationCompleted                                                                    bool
	ContactOverlapRepresentationCompleted                                                          bool
	AllOpenModesRepresentationComplete                                                             bool
	ActivationRuleDerived                                                                          bool
	DecouplingMatchingRuleDerived                                                                  bool
	PhysicalScaleDerived                                                                           bool
	ThresholdCorrectedBetaDerived                                                                  bool
	FullFiniteBetaMatchingTensorDerived                                                            bool
	ResidualNullityBefore, ResidualNullityAfter                                                    int
	ResidualSymmetryBroken                                                                         bool
	PhysicalWeakAngleDerived, FineStructureDerived, PhysicalMassesDerived, HiddenObservedInputUsed bool
	TruthStatement                                                                                 string
	RejectedClaims, RemainingUnknowns                                                              []string
	RecommendedNextGate                                                                            string
}

var defaultOnce sync.Once
var defaultValue Analysis
var defaultErr error

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		fil, err := filtration.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(fil, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(fil filtration.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !fil.NonUniqueFiltrationWitnessed || fil.ResidualNullityAfter != 3 {
		return Analysis{}, fmt.Errorf("Gate 108 requires Gate 107 filtration obstruction with residual nullity 3")
	}
	if fil.ThresholdCorrectedBetaDerived || fil.DerivedDecouplingMatchingRule || fil.HiddenObservedInputUsed {
		return Analysis{}, fmt.Errorf("Gate 108 refuses pre-derived threshold corrections or hidden observed input")
	}
	decisions := fil.Shell.Coarse.Thresholds.Decisions
	if len(decisions) == 0 {
		return Analysis{}, fmt.Errorf("Gate 108 requires threshold activation decisions")
	}

	rows := buildRows(decisions)
	attempts := buildCompletionAttempts(rows)
	witnesses := buildAmbiguityWitnesses()
	repComplete, incompleteOpen, sectorRows, individualThresholds, vacuumRows, deltaRows, betaAllowed := countRows(rows)
	openNames := openIncompleteNames(rows)
	scalar := scalarSectorVector()
	truth := "Gate 108 constructs the representation-to-beta matching ledger. One representation-complete row exists at sector level: the scalar/contact active carrier contributes the complex scalar doublet beta row Δb=(1/10,1/6,0), already counted in the baseline finite inventory. It is not a heavy-threshold correction. The B-sector gap and contact partial-overlap modes still have no SU(3)c×SU(2)L×U(1)Y representation rows, no activation predicate, and no decoupling/matching scale. Therefore no threshold-corrected beta tensor is derived."

	return Analysis{
		Filtration: fil, Rows: rows, CompletionAttempts: attempts, AmbiguityWitnesses: witnesses,
		RepresentationOpenModes: openNames, RepresentationCompleteRows: repComplete, IncompleteOpenRows: incompleteOpen,
		SectorBaselineRows: sectorRows, IndividualThresholdRows: individualThresholds, VacuumRows: vacuumRows,
		DeltaBRowsDerived: deltaRows, BetaCorrectionRowsAllowed: betaAllowed,
		ScalarSectorRowConstructed: sectorRows == 1, ScalarSectorDeltaB: scalar,
		ScalarSectorMatchesBaseline:       closeVec(scalar, BetaVector{B1: 0.1, B2: 1.0 / 6.0}, eps),
		ScalarSectorIsThresholdCorrection: false, BGapRepresentationCompleted: false, ContactOverlapRepresentationCompleted: false,
		AllOpenModesRepresentationComplete: incompleteOpen == 0, ActivationRuleDerived: false, DecouplingMatchingRuleDerived: false,
		PhysicalScaleDerived: false, ThresholdCorrectedBetaDerived: betaAllowed > 0, FullFiniteBetaMatchingTensorDerived: false,
		ResidualNullityBefore: fil.ResidualNullityAfter, ResidualNullityAfter: fil.ResidualNullityAfter, ResidualSymmetryBroken: false,
		PhysicalWeakAngleDerived: false, FineStructureDerived: false, PhysicalMassesDerived: false, HiddenObservedInputUsed: false,
		TruthStatement: truth,
		RejectedClaims: []string{
			"a sector-level scalar beta row is a heavy-threshold correction",
			"unassigned B-sector/contact modes may enter Δb_i by multiplicity alone",
			"a finite spectral value is sufficient to decide decoupling",
			"the baseline one-loop beta diagnostic becomes threshold-corrected without a matching tensor",
			"representation completion alone would determine alpha, thetaW, or mass scales",
		},
		RemainingUnknowns: []string{
			"U-28A-REP-B: derive SU(3)c×SU(2)L×U(1)Y representation rows for the B-sector gap modes",
			"U-28B-REP-C: classify contact partial-overlap modes as physical, regulator, or vacuum-frustration modes",
			"U-28C-ACTIVATION: derive a finite activation predicate, not an arbitrary spectral cutoff",
			"U-28D-DECOUPLING: derive whether an activated mode contributes as scalar, Weyl, vector, or constrained finite mode",
			"U-28E-SCALE: derive the matching scale or shell log before Δb_i can affect RG flow",
			"U-28F-TENSOR: complete a full mode×gauge-factor beta-matching tensor",
		},
		RecommendedNextGate: "Gate 109 — finite mass/activation class classifier for B-sector and contact-overlap modes",
	}, nil
}

func buildRows(decisions []thresholdactivation.ActivationDecision) []MatchingRow {
	rows := []MatchingRow{{Name: "scalar/contact active sector aggregate", Kind: SectorBaselineRow, ModeKind: string(threshold.ScalarActiveCandidate), Status: thresholdactivation.ContinuumFieldCandidate, Assignment: thresholdrep.AssignedSectorLevel, Rep: thresholdrep.GaugeRepresentation{SU3: "1", SU2: "2", Hypercharge: "±1/2 sector", Detail: "one complex scalar doublet assembled from four real active contact directions"}, RepresentationComplete: true, DeltaBDerived: true, SectorLevel: true, ContinuumBaseline: true, DeltaB: scalarSectorVector(), Reason: "baseline complex scalar doublet contribution already counted in the finite inventory; not a heavy threshold row"}}
	for _, d := range decisions {
		c := d.Assignment.Candidate
		row := MatchingRow{Name: c.Name, Kind: IndividualModeRow, ModeKind: string(c.Kind), Status: d.Status, Assignment: d.Assignment.Status, Rep: d.Assignment.Rep, HeavyThreshold: d.HeavyThresholdDerived, DecouplingDerived: d.DecouplingRuleDerived, CanCorrectBeta: d.CanCorrectBeta, Reason: d.Reason}
		switch c.Kind {
		case threshold.ScalarActiveCandidate:
			row.RepresentationComplete = false
			row.Reason = "individual real scalar/contact modes are not separately oriented as threshold fields; only the aggregate complex doublet row is derived"
		case threshold.RadialCandidate:
			row.RepresentationComplete = true
			row.DeltaBDerived = true
			row.SectorLevel = true
			row.DeltaB = BetaVector{}
			row.Reason = "bridge-level singlet radial response has zero gauge beta row, but no separate heavy-threshold activation is derived"
		case threshold.BGapCandidate, threshold.ContactOverlapCandidate:
			row.Kind = OpenModeRow
		case threshold.LeakageCandidate:
			row.Kind = VacuumOnlyRow
			row.ExcludedFromThresholds = true
			row.Reason = "vacuum-frustration invariant excluded from beta matching"
		}
		rows = append(rows, row)
	}
	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i].Kind == rows[j].Kind {
			return rows[i].Name < rows[j].Name
		}
		return rows[i].Kind < rows[j].Kind
	})
	return rows
}

func buildCompletionAttempts(rows []MatchingRow) []CompletionAttempt {
	return []CompletionAttempt{
		{Name: "scalar/contact aggregate baseline row", Constructed: true, RepresentationComplete: true, ActivationDerived: true, DeltaBRowDerived: true, RejectedAsThresholdPhysics: true, Detail: "Δb=(1/10,1/6,0) is a baseline complex-scalar-doublet row, not a heavy threshold correction"},
		{Name: "B-sector gap beta row", Constructed: true, RepresentationComplete: false, RejectedAsThresholdPhysics: true, Detail: "finite gap exists, but gauge representation and decoupling class are absent"},
		{Name: "contact partial-overlap beta rows", Constructed: true, RepresentationComplete: false, RejectedAsThresholdPhysics: true, Detail: "partial overlaps are finite modes, but not yet classified as physical fields, regulators, or vacuum-frustration modes"},
		{Name: "full threshold beta-matching tensor", Constructed: true, RepresentationComplete: allRowsRepresentationComplete(rows), RejectedAsThresholdPhysics: true, Detail: "mode×gauge tensor is blocked by open rows and by missing activation/decoupling/scale map"},
	}
}

func buildAmbiguityWitnesses() []AmbiguityWitness {
	return []AmbiguityWitness{
		{Name: "treat all contact partial-overlap modes as singlet scalars", CompatibleWithFiniteData: true, DeltaB: BetaVector{}, Reason: "would give zero gauge beta row, but the singlet classification is not derived"},
		{Name: "treat contact partial-overlap modes as scalar doublets", CompatibleWithFiniteData: true, DeltaB: BetaVector{B1: 7.0 / 10.0, B2: 7.0 / 6.0}, Reason: "multiplicity-based doublet assignment gives a nonzero row, but it is an arbitrary representation choice"},
		{Name: "exclude contact partial-overlap modes as vacuum/regulator modes", CompatibleWithFiniteData: true, DeltaB: BetaVector{}, Reason: "also compatible because no physical activation theorem distinguishes the cases"},
	}
}

func scalarSectorVector() BetaVector { return BetaVector{B1: 1.0 / 10.0, B2: 1.0 / 6.0, B3: 0} }
func countRows(rows []MatchingRow) (repComplete, incompleteOpen, sectorRows, individualThresholds, vacuumRows, deltaRows, betaAllowed int) {
	for _, r := range rows {
		if r.RepresentationComplete {
			repComplete++
		}
		if r.Kind == OpenModeRow && !r.RepresentationComplete {
			incompleteOpen++
		}
		if r.Kind == SectorBaselineRow {
			sectorRows++
		}
		if r.IndividualModeLevel && r.HeavyThreshold {
			individualThresholds++
		}
		if r.Kind == VacuumOnlyRow {
			vacuumRows++
		}
		if r.DeltaBDerived {
			deltaRows++
		}
		if r.CanCorrectBeta {
			betaAllowed++
		}
	}
	return
}
func openIncompleteNames(rows []MatchingRow) []string {
	out := []string{}
	for _, r := range rows {
		if r.Kind == OpenModeRow && !r.RepresentationComplete {
			out = append(out, r.Name)
		}
	}
	sort.Strings(out)
	return out
}
func allRowsRepresentationComplete(rows []MatchingRow) bool {
	for _, r := range rows {
		if r.Kind == OpenModeRow && !r.RepresentationComplete {
			return false
		}
	}
	return true
}
func closeVec(a, b BetaVector, eps float64) bool {
	return math.Abs(a.B1-b.B1) <= eps && math.Abs(a.B2-b.B2) <= eps && math.Abs(a.B3-b.B3) <= eps
}
func FormatVector(v BetaVector) string {
	return fmt.Sprintf("(Δb1=%.10f, Δb2=%.10f, Δb3=%.10f)", v.B1, v.B2, v.B3)
}
func FormatRows(rows []MatchingRow, max int) string {
	if max <= 0 || max > len(rows) {
		max = len(rows)
	}
	parts := make([]string, 0, max+1)
	for i := 0; i < max; i++ {
		r := rows[i]
		rep := fmt.Sprintf("(%s,%s)_%s", empty(r.Rep.SU3), empty(r.Rep.SU2), empty(r.Rep.Hypercharge))
		flag := "open"
		if r.RepresentationComplete {
			flag = "rep"
		}
		if r.DeltaBDerived {
			flag += "+Δb"
		}
		parts = append(parts, fmt.Sprintf("%s:%s:%s:%s", r.Name, r.Kind, rep, flag))
	}
	if max < len(rows) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(rows)-max))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}
func FormatAttempts(attempts []CompletionAttempt) string {
	parts := make([]string, 0, len(attempts))
	for _, a := range attempts {
		parts = append(parts, fmt.Sprintf("%s(rep=%t,Δb=%t,correct=%t)", a.Name, a.RepresentationComplete, a.DeltaBRowDerived, a.CanCorrectBeta))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}
func FormatWitnesses(ws []AmbiguityWitness) string {
	parts := make([]string, 0, len(ws))
	for _, w := range ws {
		parts = append(parts, fmt.Sprintf("%s:%s", w.Name, FormatVector(w.DeltaB)))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}
func Join(xs []string) string { return strings.Join(xs, "; ") }
func empty(s string) string {
	if s == "" {
		return "?"
	}
	return s
}
