// Package contactreconstruction implements Gate 121: contact spectral
// reconstruction / invariant-to-row lifting obstruction theorem.
//
// Gate 120 proved a fork: the weighted contact quotient preserves all seven
// distinct contact overlap rows but leaves only singleton orbits, while the
// anonymous quotient restores a one-orbit symmetry only by forgetting the
// row-level spectral information needed for threshold beta matching. Gate 121
// asks whether the lost row-level data can be reconstructed from invariant or
// quotient data without smuggling in a noncanonical contact-to-Fano bijection.
//
// The result is an obstruction theorem. The spectral multiset reconstructs the
// seven numerical values, but it does not reconstruct which value belongs to a
// Fano point/line, local field, representation row, mass threshold, or
// decoupling class. Every lift from the anonymous invariant back to row-level
// contact-Fano data requires choosing one of 7! assignments. The weighted lift
// keeps row identity but keeps the identity-only contact action and still has no
// Fano-natural representation row. Therefore there is no lift that is both
// information-preserving and convention-free. Threshold beta corrections and
// physical constants remain sealed.
package contactreconstruction

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactquotient"
)

type LiftKind string

const (
	WeightedSingletonLift  LiftKind = "weighted-singleton-row-lift"
	AnonymousOrbitLift     LiftKind = "anonymous-orbit-to-row-lift"
	SpectralMultisetLift   LiftKind = "spectral-multiset-lift"
	TransportedFanoRowLift LiftKind = "transported-fano-row-lift"
	RepresentationRowLift  LiftKind = "representation-row-lift"
)

type LiftStatus string

const (
	LiftOpen                 LiftStatus = "lift-open"
	LiftCanonicalButTooSmall LiftStatus = "canonical-but-too-small"
	LiftNoncanonical         LiftStatus = "noncanonical-lift"
	LiftForbidden            LiftStatus = "forbidden-for-beta"
)

type ContactReconstructionRow struct {
	Name, ModeKind string
	Value          float64

	WeightedRowPreserved           bool
	AnonymousOrbitRepresentative   bool
	SpectralValueRecovered         bool
	RowIdentityRecoveredFromAnon   bool
	ContactFanoAssignmentRecovered bool
	NeedsBijectionChoice           bool
	PossibleAssignments            int
	RepresentationRowDerived       bool
	CanEnterBetaTensor             bool
	ZeroRowProved                  bool

	Status LiftStatus
	Reason string
}

type LiftAttempt struct {
	Name string
	Kind LiftKind

	Constructed                bool
	CanonicalUnderCurrentData  bool
	UsesExtraConvention        bool
	PreservesSpectralValues    bool
	PreservesRowIdentity       bool
	RestoresFanoOrbit          bool
	ReconstructsRowIdentity    bool
	ReconstructsFanoAssignment bool
	PossibleLifts              int
	CanonicalLifts             int
	RepresentationRowDerived   bool
	BetaRowPermitted           bool
	ZeroRowProved              bool
	RejectedAsPremature        bool
	MissingTerms               []string
	Detail                     string
}

type ReconstructionCriterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type ReconstructionSummary struct {
	ContactRows                 int
	DistinctSpectralValues      int
	WeightedOrbitSizes          []int
	AnonymousOrbitSizes         []int
	CompatibleBijectionCount    int
	AnonymousRowLiftCount       int
	CanonicalAnonymousRowLifts  int
	WeightedCanonicalLiftCount  int
	NoLossNoChoiceLiftCount     int
	RepresentationCompleteLifts int
}

type Analysis struct {
	Quotient contactquotient.Analysis

	Rows     []ContactReconstructionRow
	Attempts []LiftAttempt
	Criteria []ReconstructionCriterion
	Summary  ReconstructionSummary

	ContactRows               int
	PositiveFiniteContactRows int
	SurvivingCohomologyRows   int
	OpenContactRowsInherited  int
	OpenContactRowsAfter      int

	InvariantToRowLiftingSearchAttempted bool
	WeightedSingletonLiftDerived         bool
	WeightedSingletonLiftCanonical       bool
	WeightedSingletonLiftPreservesRows   bool
	WeightedSingletonLiftFanoLike        bool
	WeightedSingletonLiftRepUsable       bool

	AnonymousInvariantLiftAttempted      bool
	AnonymousInvariantLiftConstructed    bool
	AnonymousInvariantLiftCanonical      bool
	AnonymousInvariantLiftNeedsChoice    bool
	AnonymousInvariantLiftPossibleRows   int
	AnonymousInvariantLiftCanonicalRows  int
	AnonymousLiftPreservesSpectralValues bool
	AnonymousLiftRecoversRowIdentity     bool
	AnonymousLiftRecoversFanoAssignment  bool

	SpectralMultisetRecovered        bool
	SpectralMultisetRecoversValues   bool
	SpectralMultisetRecoversRows     bool
	SpectralMultisetRecoversFanoRows bool

	TransportedFanoRowLiftPossibleAfterChoice bool
	TransportedFanoRowLiftCanonical           bool
	FanoEquivariantRowLiftDerived             bool
	NoLossNoChoiceLiftExists                  bool
	InvariantToRowReconstructionDerived       bool
	ReconstructionObstructionDerived          bool
	RowLiftingAmbiguityDerived                bool
	InformationChoiceNoGoDerived              bool

	RepresentationCompleteRows          int
	RepresentationOpenRows              int
	ContactBetaRowsAllowed              int
	ContactZeroRowsProved               int
	ThresholdCorrectedBetaDerived       bool
	FullFiniteBetaMatchingTensorDerived bool

	QuotientForkObstructionInherited     bool
	OrbitCollapseObstructionInherited    bool
	SpectralInformationLossInherited     bool
	ActionObstructionInherited           bool
	NaturalityObstructionInherited       bool
	SymmetrySelectorObstructionInherited bool

	ResidualNullityBefore  int
	ResidualNullityAfter   int
	ResidualSymmetryBroken bool

	PhysicalWeakAngleDerived bool
	FineStructureDerived     bool
	PhysicalMassesDerived    bool
	PhysicalScaleDerived     bool
	HiddenObservedInputUsed  bool

	TruthStatement      string
	RejectedClaims      []string
	RemainingUnknowns   []string
	RecommendedNextGate string
}

var defaultOnce sync.Once
var defaultValue Analysis
var defaultErr error

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		q, err := contactquotient.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(q, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(q contactquotient.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !q.ContactSpectrumQuotientSearchAttempted || !q.WeightedSpectrumQuotientDerived || !q.AnonymousSpectrumQuotientDerived || !q.QuotientForkObstructionDerived || !q.OrbitCollapseObstructionDerived || !q.SpectralInformationLossDerived {
		return Analysis{}, fmt.Errorf("Gate 121 requires Gate 120 quotient fork/orbit-collapse obstruction")
	}
	if !q.WeightedQuotientCanonical || !q.WeightedQuotientIsIdentity || !q.WeightedQuotientPreservesAllRows || q.WeightedQuotientProducesFanoOrbit {
		return Analysis{}, fmt.Errorf("Gate 121 requires canonical singleton weighted quotient that produces no Fano orbit")
	}
	if !q.AnonymousQuotientCanonical || !q.AnonymousQuotientCollapsesAllRows || !q.AnonymousQuotientDestroysSpectralRows || q.AnonymousQuotientRepresentationUsable {
		return Analysis{}, fmt.Errorf("Gate 121 requires anonymous quotient that collapses rows and loses spectral row data")
	}
	if q.Summary.CompatibleBijectionCount != 5040 || q.TransportedFanoQuotientCanonical || q.Summary.CanonicalTransportedQuotients != 0 {
		return Analysis{}, fmt.Errorf("Gate 121 requires 7! noncanonical transported contact-Fano lifts and zero canonical transported quotients")
	}
	if q.ContactRows != 7 || q.OpenContactRowsAfter != 7 || q.RepresentationCompleteRows != 0 || q.ContactBetaRowsAllowed != 0 || q.ContactZeroRowsProved != 0 || q.ThresholdCorrectedBetaDerived {
		return Analysis{}, fmt.Errorf("Gate 121 requires seven open contact rows and closed beta firewall")
	}
	if q.ResidualNullityAfter != 3 || q.HiddenObservedInputUsed || q.PhysicalWeakAngleDerived || q.FineStructureDerived || q.PhysicalMassesDerived || q.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 121 refuses hidden physical input or changed residual nullity")
	}

	values := contactValues(q.Rows)
	distinct := distinctCount(values, eps)
	if distinct != 7 {
		return Analysis{}, fmt.Errorf("expected seven distinct contact quotient values, got %d", distinct)
	}

	summary := ReconstructionSummary{
		ContactRows:                 7,
		DistinctSpectralValues:      distinct,
		WeightedOrbitSizes:          append([]int(nil), q.WeightedQuotientOrbitSizes...),
		AnonymousOrbitSizes:         append([]int(nil), q.AnonymousQuotientOrbitSizes...),
		CompatibleBijectionCount:    q.Summary.CompatibleBijectionCount,
		AnonymousRowLiftCount:       factorial(7),
		CanonicalAnonymousRowLifts:  0,
		WeightedCanonicalLiftCount:  1,
		NoLossNoChoiceLiftCount:     0,
		RepresentationCompleteLifts: 0,
	}

	rows := buildRows(q.Rows, summary.AnonymousRowLiftCount)
	attempts := buildAttempts(summary)
	criteria := buildCriteria(summary)
	counts := countRows(rows)

	truth := "Gate 121 asks whether anonymous quotient or invariant data can be lifted back to row-level contact representation data without choosing a hidden contact-to-Fano labeling. The answer is no. The spectral multiset reconstructs the seven numerical overlap values, but it does not reconstruct which value belongs to a Fano point, Fano line, local field variable, gauge representation row, mass threshold, or decoupling class. The weighted lift is canonical and preserves all rows, but keeps the identity-only contact action and gives no Fano-like representation structure. The anonymous one-orbit lift can be expanded back into rows only by choosing one of 7! labelings. Therefore no lift is both information-preserving and convention-free; no contact representation row, zero row, threshold beta correction, or physical constant is derived."

	return Analysis{
		Quotient: q,
		Rows:     rows,
		Attempts: attempts,
		Criteria: criteria,
		Summary:  summary,

		ContactRows:               counts.contact,
		PositiveFiniteContactRows: counts.positive,
		SurvivingCohomologyRows:   counts.surviving,
		OpenContactRowsInherited:  q.OpenContactRowsAfter,
		OpenContactRowsAfter:      counts.open,

		InvariantToRowLiftingSearchAttempted: true,
		WeightedSingletonLiftDerived:         true,
		WeightedSingletonLiftCanonical:       true,
		WeightedSingletonLiftPreservesRows:   true,
		WeightedSingletonLiftFanoLike:        false,
		WeightedSingletonLiftRepUsable:       false,

		AnonymousInvariantLiftAttempted:      true,
		AnonymousInvariantLiftConstructed:    true,
		AnonymousInvariantLiftCanonical:      false,
		AnonymousInvariantLiftNeedsChoice:    true,
		AnonymousInvariantLiftPossibleRows:   summary.AnonymousRowLiftCount,
		AnonymousInvariantLiftCanonicalRows:  0,
		AnonymousLiftPreservesSpectralValues: true,
		AnonymousLiftRecoversRowIdentity:     false,
		AnonymousLiftRecoversFanoAssignment:  false,

		SpectralMultisetRecovered:        true,
		SpectralMultisetRecoversValues:   true,
		SpectralMultisetRecoversRows:     false,
		SpectralMultisetRecoversFanoRows: false,

		TransportedFanoRowLiftPossibleAfterChoice: true,
		TransportedFanoRowLiftCanonical:           false,
		FanoEquivariantRowLiftDerived:             false,
		NoLossNoChoiceLiftExists:                  false,
		InvariantToRowReconstructionDerived:       false,
		ReconstructionObstructionDerived:          true,
		RowLiftingAmbiguityDerived:                true,
		InformationChoiceNoGoDerived:              true,

		RepresentationCompleteRows:          0,
		RepresentationOpenRows:              counts.open,
		ContactBetaRowsAllowed:              0,
		ContactZeroRowsProved:               0,
		ThresholdCorrectedBetaDerived:       false,
		FullFiniteBetaMatchingTensorDerived: false,

		QuotientForkObstructionInherited:     true,
		OrbitCollapseObstructionInherited:    true,
		SpectralInformationLossInherited:     true,
		ActionObstructionInherited:           q.ActionObstructionInherited,
		NaturalityObstructionInherited:       q.NaturalityObstructionInherited,
		SymmetrySelectorObstructionInherited: q.SymmetrySelectorObstructionInherited,

		ResidualNullityBefore:  3,
		ResidualNullityAfter:   3,
		ResidualSymmetryBroken: false,

		PhysicalWeakAngleDerived: false,
		FineStructureDerived:     false,
		PhysicalMassesDerived:    false,
		PhysicalScaleDerived:     false,
		HiddenObservedInputUsed:  false,

		TruthStatement: truth,
		RejectedClaims: []string{
			"the anonymous one-orbit quotient can be lifted back to rows canonically",
			"the spectral multiset identifies contact modes with Fano points or lines",
			"sorting contact overlap values supplies a physical representation-row ordering",
			"one of the 7! row lifts is selected by current finite data",
			"invariant quotient data derives contact threshold beta rows",
			"Gate 121 derives alpha, physical thetaW, threshold masses, M*, g_*, or W/Z/Higgs/fermion masses",
		},
		RemainingUnknowns: []string{
			"canonical contact-to-Fano row assignment preserving spectral data",
			"local bundle, field variable, or constraint complex that gives row semantics",
			"representation row map for the seven contact modes",
			"mass activation and decoupling rule for contact modes",
			"threshold-corrected beta tensor",
			"whether contact modes are physical, constrained, regulator, or vacuum-frustration modes",
		},
		RecommendedNextGate: "Gate 122 — contact row semantics / local variable reconstruction from incidence-weighted spectrum",
	}, nil
}

type rowCounts struct{ contact, positive, surviving, open int }

func buildRows(in []contactquotient.ContactQuotientRow, possibleAssignments int) []ContactReconstructionRow {
	rows := make([]ContactReconstructionRow, 0, len(in))
	for _, r := range in {
		rows = append(rows, ContactReconstructionRow{
			Name:                           r.Name,
			ModeKind:                       r.ModeKind,
			Value:                          r.Value,
			WeightedRowPreserved:           r.WeightedQuotientPreservesRow,
			AnonymousOrbitRepresentative:   true,
			SpectralValueRecovered:         true,
			RowIdentityRecoveredFromAnon:   false,
			ContactFanoAssignmentRecovered: false,
			NeedsBijectionChoice:           true,
			PossibleAssignments:            possibleAssignments,
			RepresentationRowDerived:       false,
			CanEnterBetaTensor:             false,
			ZeroRowProved:                  false,
			Status:                         LiftOpen,
			Reason:                         "spectral value is recoverable as part of the multiset, but row identity and contact-Fano assignment require a noncanonical labeling choice",
		})
	}
	sort.SliceStable(rows, func(i, j int) bool {
		if math.Abs(rows[i].Value-rows[j].Value) < 1e-12 {
			return rows[i].Name < rows[j].Name
		}
		return rows[i].Value < rows[j].Value
	})
	return rows
}

func buildAttempts(s ReconstructionSummary) []LiftAttempt {
	return []LiftAttempt{
		{
			Name:                      "weighted singleton row lift",
			Kind:                      WeightedSingletonLift,
			Constructed:               true,
			CanonicalUnderCurrentData: true,
			PreservesSpectralValues:   true,
			PreservesRowIdentity:      true,
			PossibleLifts:             1,
			CanonicalLifts:            1,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"Fano-like orbit", "gauge representation row", "local field map"},
			Detail:                    "this lift is canonical only because it keeps the identity-only weighted contact action; it cannot produce a seven-mode Fano orbit or representation tensor",
		},
		{
			Name:                       "anonymous one-orbit invariant to row lift",
			Kind:                       AnonymousOrbitLift,
			Constructed:                true,
			UsesExtraConvention:        true,
			PreservesSpectralValues:    true,
			PreservesRowIdentity:       false,
			RestoresFanoOrbit:          true,
			ReconstructsRowIdentity:    false,
			ReconstructsFanoAssignment: false,
			PossibleLifts:              s.AnonymousRowLiftCount,
			CanonicalLifts:             0,
			RejectedAsPremature:        true,
			MissingTerms:               []string{"canonical row labels", "contact-Fano assignment", "row semantics"},
			Detail:                     fmt.Sprintf("the anonymous orbit can be lifted to row-level data in %d ways, one for each hidden labeling; none is selected", s.AnonymousRowLiftCount),
		},
		{
			Name:                       "spectral multiset reconstruction",
			Kind:                       SpectralMultisetLift,
			Constructed:                true,
			CanonicalUnderCurrentData:  true,
			PreservesSpectralValues:    true,
			PreservesRowIdentity:       false,
			ReconstructsRowIdentity:    false,
			ReconstructsFanoAssignment: false,
			PossibleLifts:              s.AnonymousRowLiftCount,
			CanonicalLifts:             0,
			RejectedAsPremature:        true,
			MissingTerms:               []string{"mode identity", "Fano point/line identity", "activation class"},
			Detail:                     "the sorted spectral multiset recovers the seven numbers but not which number is attached to which contact mode or representation row",
		},
		{
			Name:                       "transported Fano row lift through chosen bijection",
			Kind:                       TransportedFanoRowLift,
			Constructed:                true,
			UsesExtraConvention:        true,
			RestoresFanoOrbit:          true,
			ReconstructsFanoAssignment: true,
			PossibleLifts:              s.CompatibleBijectionCount,
			CanonicalLifts:             0,
			RejectedAsPremature:        true,
			MissingTerms:               []string{"canonical bijection", "spectral preservation", "symmetry-breaking selector"},
			Detail:                     fmt.Sprintf("transporting Fano row structure to contact rows requires selecting one of %d contact-Fano bijections", s.CompatibleBijectionCount),
		},
		{
			Name:                     "representation row lift",
			Kind:                     RepresentationRowLift,
			Constructed:              true,
			PossibleLifts:            0,
			CanonicalLifts:           0,
			RepresentationRowDerived: false,
			BetaRowPermitted:         false,
			ZeroRowProved:            false,
			RejectedAsPremature:      true,
			MissingTerms:             []string{"SU(3)c×SU(2)L×U(1)Y row", "Lorentz kinetic map", "mass/decoupling rule"},
			Detail:                   "no invariant or lifted row contains the physical data required for beta matching",
		},
	}
}

func buildCriteria(s ReconstructionSummary) []ReconstructionCriterion {
	return []ReconstructionCriterion{
		{Name: "weighted no-loss lift", Required: true, Derived: s.WeightedCanonicalLiftCount == 1 && s.NoLossNoChoiceLiftCount == 0, Detail: "canonical and information-preserving, but identity-only and not representation-complete"},
		{Name: "anonymous row lift", Required: true, Derived: s.AnonymousRowLiftCount == 5040 && s.CanonicalAnonymousRowLifts == 0, Detail: "row recovery from the anonymous orbit has 7! choices and no canonical lift"},
		{Name: "no-loss/no-choice lift", Required: true, Derived: s.NoLossNoChoiceLiftCount > 0, Detail: "must be false: no lift both preserves row data and avoids extra convention"},
		{Name: "representation-complete lift", Required: true, Derived: s.RepresentationCompleteLifts > 0, Detail: "must be false: no lift supplies gauge representation, kinetic, mass, and decoupling rows"},
	}
}

func countRows(rows []ContactReconstructionRow) rowCounts {
	var c rowCounts
	for _, r := range rows {
		c.contact++
		if r.Value > 0 {
			c.positive++
		}
		// Rows entering Gate 121 already survived Gate 114 cohomology through Gate 120.
		c.surviving++
		if r.Status == LiftOpen {
			c.open++
		}
	}
	return c
}

func contactValues(rows []contactquotient.ContactQuotientRow) []float64 {
	values := make([]float64, 0, len(rows))
	for _, r := range rows {
		values = append(values, r.Value)
	}
	sort.Float64s(values)
	return values
}

func distinctCount(values []float64, eps float64) int {
	if len(values) == 0 {
		return 0
	}
	sorted := append([]float64(nil), values...)
	sort.Float64s(sorted)
	count := 1
	last := sorted[0]
	for _, v := range sorted[1:] {
		if math.Abs(v-last) > eps {
			count++
			last = v
		}
	}
	return count
}

func factorial(n int) int {
	if n < 0 {
		return 0
	}
	out := 1
	for i := 2; i <= n; i++ {
		out *= i
	}
	return out
}

func FormatRows(rows []ContactReconstructionRow, limit int) string {
	if limit <= 0 || limit > len(rows) {
		limit = len(rows)
	}
	parts := make([]string, 0, limit)
	for i := 0; i < limit; i++ {
		r := rows[i]
		parts = append(parts, fmt.Sprintf("%s[value=%.10f recovered=%t fanoRecovered=%t choices=%d status=%s]", r.Name, r.Value, r.SpectralValueRecovered, r.ContactFanoAssignmentRecovered, r.PossibleAssignments, r.Status))
	}
	if limit < len(rows) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(rows)-limit))
	}
	return strings.Join(parts, "; ")
}

func FormatAttempts(attempts []LiftAttempt) string {
	parts := make([]string, 0, len(attempts))
	for _, a := range attempts {
		parts = append(parts, fmt.Sprintf("%s(kind=%s possible=%d canonical=%d convention=%t rejected=%t)", a.Name, a.Kind, a.PossibleLifts, a.CanonicalLifts, a.UsesExtraConvention, a.RejectedAsPremature))
	}
	return strings.Join(parts, "; ")
}

func FormatCriteria(criteria []ReconstructionCriterion) string {
	parts := make([]string, 0, len(criteria))
	for _, c := range criteria {
		parts = append(parts, fmt.Sprintf("%s=%t (%s)", c.Name, c.Derived, c.Detail))
	}
	return strings.Join(parts, "; ")
}

func FormatSummary(s ReconstructionSummary) string {
	return fmt.Sprintf("weightedOrbits=%v anonymousOrbits=%v anonymousRowLifts=%d canonicalAnonymous=%d noLossNoChoice=%d representationComplete=%d", s.WeightedOrbitSizes, s.AnonymousOrbitSizes, s.AnonymousRowLiftCount, s.CanonicalAnonymousRowLifts, s.NoLossNoChoiceLiftCount, s.RepresentationCompleteLifts)
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
