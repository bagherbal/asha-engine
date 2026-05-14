// Package contactrowsemantics implements Gate 122: contact row semantics /
// local variable reconstruction from incidence-weighted spectrum.
//
// Gate 121 showed that invariant/quotient data can recover the seven contact
// spectral values, but cannot lift them to row-level contact/Fano semantics
// without choosing one of 7! labelings. Gate 122 asks whether adding the
// incidence weights already present in the Fano/contact geometry changes this:
// can an incidence-weighted spectrum construct local variables, constraint
// semantics, or representation rows?
//
// The answer is still no. The Fano incidence degrees are uniform: each point
// lies on three lines, and each line contains three points. Multiplying the
// contact spectrum by this canonical incidence degree preserves the seven rows
// and their distinct numerical values, but it adds no new row semantics. Signed
// incidence or line/point attachment would require a contact-to-Fano labeling,
// again one of 7! choices. Therefore incidence-weighted spectral data remains a
// diagnostic, not a local-variable reconstruction or beta-matching permission.
package contactrowsemantics

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactreconstruction"
)

type SemanticStatus string

const (
	RowSemanticOpen SemanticStatus = "row-semantics-open"
)

type SemanticsAttemptKind string

const (
	UniformIncidenceWeightAttempt SemanticsAttemptKind = "uniform-incidence-weighted-spectrum"
	SignedIncidenceWeightAttempt  SemanticsAttemptKind = "signed-incidence-weighted-spectrum"
	MomentReconstructionAttempt   SemanticsAttemptKind = "incidence-spectral-moment-reconstruction"
	LocalVariableAttempt          SemanticsAttemptKind = "local-variable-reconstruction"
	ConstraintSemanticAttempt     SemanticsAttemptKind = "constraint-semantic-map"
	RepresentationRuleAttempt     SemanticsAttemptKind = "representation-row-rule"
)

type ContactSemanticRow struct {
	Name, ModeKind string
	Value          float64

	IncidenceDegree        int
	IncidenceWeightedValue float64
	SpectralValueRecovered bool
	RowIdentityPreserved   bool

	FanoPointOrLineAssigned     bool
	NeedsContactFanoBijection   bool
	PossibleAssignments         int
	LocalVariableDerived        bool
	ConstraintSemanticDerived   bool
	GaugeRepresentationDerived  bool
	LorentzKineticDerived       bool
	MassActivationDerived       bool
	DecouplingRuleDerived       bool
	RepresentationOrZeroDerived bool
	CanEnterBetaTensor          bool
	ZeroRowProved               bool

	Status SemanticStatus
	Reason string
}

type SemanticsAttempt struct {
	Name string
	Kind SemanticsAttemptKind

	Constructed               bool
	CanonicalUnderCurrentData bool
	UsesExtraConvention       bool
	PreservesSpectralValues   bool
	PreservesRowIdentity      bool
	AddsRowSemantics          bool
	RequiresBijection         bool
	PossibleAssignments       int
	LocalVariablesDerived     bool
	ConstraintMapDerived      bool
	RepresentationRowDerived  bool
	BetaRowPermitted          bool
	ZeroRowProved             bool
	RejectedAsPremature       bool
	MissingTerms              []string
	Detail                    string
}

type SemanticsCriterion struct {
	Name     string
	Required bool
	Derived  bool
	Detail   string
}

type SemanticsSummary struct {
	ContactRows                int
	IncidenceDegree            int
	WeightedValuesDistinct     int
	UniformWeightPreservesRows bool
	UniformWeightAddsSemantics bool
	SignedIncidenceChoices     int
	CanonicalSignedAssignments int
	LocalVariableRows          int
	ConstraintSemanticRows     int
	RepresentationCompleteRows int
	ContactBetaRowsAllowed     int
}

type Analysis struct {
	Reconstruction contactreconstruction.Analysis

	Rows     []ContactSemanticRow
	Attempts []SemanticsAttempt
	Criteria []SemanticsCriterion
	Summary  SemanticsSummary

	ContactRows               int
	PositiveFiniteContactRows int
	SurvivingCohomologyRows   int
	OpenContactRowsInherited  int
	OpenContactRowsAfter      int

	IncidenceWeightedSpectrumSearchAttempted bool
	UniformFanoIncidenceDegreeAvailable      bool
	FanoPointDegree                          int
	FanoLineSize                             int
	IncidenceWeightCanonical                 bool
	IncidenceWeightedValuesDistinct          bool
	IncidenceWeightingPreservesRows          bool
	IncidenceWeightingAddsRowSemantics       bool

	SignedIncidenceAttempted   bool
	SignedIncidenceCanonical   bool
	SignedIncidenceNeedsChoice bool
	SignedIncidenceChoiceCount int

	IncidenceMomentReconstructionAttempted bool
	IncidenceMomentsRecoverSpectrum        bool
	IncidenceMomentsRecoverRowIdentity     bool
	IncidenceMomentsRecoverFanoSemantics   bool

	LocalVariableReconstructionAttempted bool
	LocalVariablesDerived                bool
	ConstraintSemanticMapAttempted       bool
	ConstraintSemanticMapDerived         bool
	RepresentationRowRuleAttempted       bool
	RepresentationRowRuleDerived         bool
	RowSemanticsDerived                  bool
	RowSemanticsObstructionDerived       bool
	IncidenceWeightedNoGoDerived         bool

	ReconstructionObstructionInherited bool
	QuotientForkObstructionInherited   bool
	NaturalityObstructionInherited     bool
	SymmetryObstructionInherited       bool
	ContactFanoAssignmentDerived       bool

	RepresentationCompleteRows          int
	RepresentationOpenRows              int
	ContactBetaRowsAllowed              int
	ContactZeroRowsProved               int
	ThresholdCorrectedBetaDerived       bool
	FullFiniteBetaMatchingTensorDerived bool

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
		rec, err := contactreconstruction.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(rec, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(rec contactreconstruction.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !rec.ReconstructionObstructionDerived || !rec.RowLiftingAmbiguityDerived || !rec.InformationChoiceNoGoDerived || rec.NoLossNoChoiceLiftExists {
		return Analysis{}, fmt.Errorf("Gate 122 requires Gate 121 invariant-to-row lifting obstruction")
	}
	if rec.ContactRows != 7 || rec.OpenContactRowsAfter != 7 || rec.RepresentationCompleteRows != 0 || rec.ContactBetaRowsAllowed != 0 || rec.ContactZeroRowsProved != 0 || rec.ThresholdCorrectedBetaDerived {
		return Analysis{}, fmt.Errorf("Gate 122 requires seven reconstruction-open contact rows and closed beta firewall")
	}
	if rec.AnonymousInvariantLiftPossibleRows != 5040 || rec.AnonymousInvariantLiftCanonicalRows != 0 || rec.TransportedFanoRowLiftCanonical || rec.FanoEquivariantRowLiftDerived {
		return Analysis{}, fmt.Errorf("Gate 122 requires 7! noncanonical contact-Fano row lifts and no equivariant row lift")
	}
	if rec.ResidualNullityAfter != 3 || rec.HiddenObservedInputUsed || rec.PhysicalWeakAngleDerived || rec.FineStructureDerived || rec.PhysicalMassesDerived || rec.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 122 refuses hidden physical input or changed residual nullity")
	}

	const degree = 3
	rows := buildRows(rec.Rows, degree, rec.AnonymousInvariantLiftPossibleRows)
	distinctWeighted := distinctWeightedValues(rows, eps)
	summary := SemanticsSummary{
		ContactRows:                len(rows),
		IncidenceDegree:            degree,
		WeightedValuesDistinct:     distinctWeighted,
		UniformWeightPreservesRows: true,
		UniformWeightAddsSemantics: false,
		SignedIncidenceChoices:     rec.AnonymousInvariantLiftPossibleRows,
		CanonicalSignedAssignments: 0,
		LocalVariableRows:          0,
		ConstraintSemanticRows:     0,
		RepresentationCompleteRows: 0,
		ContactBetaRowsAllowed:     0,
	}
	attempts := buildAttempts(summary)
	criteria := buildCriteria(summary)
	counts := countRows(rows)

	truth := "Gate 122 tests whether incidence-weighted spectral data can reconstruct row semantics for the seven unresolved contact modes. The only canonical incidence weight currently available is uniform Fano degree three: multiplying every contact overlap value by three preserves the seven distinct rows, but adds no Fano point/line identity, local section variable, constraint semantic map, gauge representation row, Lorentz kinetic row, mass activation, or decoupling rule. Signed or line-specific incidence data would require first choosing one of 7! contact-to-Fano assignments. Therefore incidence-weighted spectrum is a useful finite diagnostic, but it is not a row-semantics theorem and cannot open threshold beta matching."

	return Analysis{
		Reconstruction: rec,
		Rows:           rows,
		Attempts:       attempts,
		Criteria:       criteria,
		Summary:        summary,

		ContactRows:               counts.contact,
		PositiveFiniteContactRows: counts.positive,
		SurvivingCohomologyRows:   counts.surviving,
		OpenContactRowsInherited:  rec.OpenContactRowsAfter,
		OpenContactRowsAfter:      counts.open,

		IncidenceWeightedSpectrumSearchAttempted: true,
		UniformFanoIncidenceDegreeAvailable:      true,
		FanoPointDegree:                          degree,
		FanoLineSize:                             degree,
		IncidenceWeightCanonical:                 true,
		IncidenceWeightedValuesDistinct:          distinctWeighted == 7,
		IncidenceWeightingPreservesRows:          true,
		IncidenceWeightingAddsRowSemantics:       false,

		SignedIncidenceAttempted:   true,
		SignedIncidenceCanonical:   false,
		SignedIncidenceNeedsChoice: true,
		SignedIncidenceChoiceCount: rec.AnonymousInvariantLiftPossibleRows,

		IncidenceMomentReconstructionAttempted: true,
		IncidenceMomentsRecoverSpectrum:        true,
		IncidenceMomentsRecoverRowIdentity:     false,
		IncidenceMomentsRecoverFanoSemantics:   false,

		LocalVariableReconstructionAttempted: true,
		LocalVariablesDerived:                false,
		ConstraintSemanticMapAttempted:       true,
		ConstraintSemanticMapDerived:         false,
		RepresentationRowRuleAttempted:       true,
		RepresentationRowRuleDerived:         false,
		RowSemanticsDerived:                  false,
		RowSemanticsObstructionDerived:       true,
		IncidenceWeightedNoGoDerived:         true,

		ReconstructionObstructionInherited: rec.ReconstructionObstructionDerived,
		QuotientForkObstructionInherited:   rec.QuotientForkObstructionInherited,
		NaturalityObstructionInherited:     rec.NaturalityObstructionInherited,
		SymmetryObstructionInherited:       rec.SymmetrySelectorObstructionInherited,
		ContactFanoAssignmentDerived:       false,

		RepresentationCompleteRows:          0,
		RepresentationOpenRows:              counts.open,
		ContactBetaRowsAllowed:              0,
		ContactZeroRowsProved:               0,
		ThresholdCorrectedBetaDerived:       false,
		FullFiniteBetaMatchingTensorDerived: false,

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
			"uniform Fano incidence degree supplies contact row semantics",
			"incidence-weighted spectral values identify Fano points or lines",
			"signed incidence can be attached to contact rows without a contact-to-Fano bijection",
			"spectral moments reconstruct local variables or constraint semantics",
			"incidence-weighted spectrum derives contact representation rows or beta corrections",
			"Gate 122 derives alpha, physical thetaW, M*, g_*, threshold masses, or W/Z/Higgs/fermion masses",
		},
		RemainingUnknowns: []string{
			"semantic source or observable that labels contact rows without a hidden bijection",
			"local variable reconstruction map for contact modes",
			"constraint semantic map stronger than the zero differential",
			"representation-row rule for contact modes",
			"mass activation and decoupling rule",
			"threshold-corrected beta tensor",
		},
		RecommendedNextGate: "Gate 123 — contact semantic source-coupling / observable selector search",
	}, nil
}

type rowCounts struct{ contact, positive, surviving, open int }

func buildRows(in []contactreconstruction.ContactReconstructionRow, degree int, possibleAssignments int) []ContactSemanticRow {
	rows := make([]ContactSemanticRow, 0, len(in))
	for _, r := range in {
		rows = append(rows, ContactSemanticRow{
			Name:                        r.Name,
			ModeKind:                    r.ModeKind,
			Value:                       r.Value,
			IncidenceDegree:             degree,
			IncidenceWeightedValue:      float64(degree) * r.Value,
			SpectralValueRecovered:      r.SpectralValueRecovered,
			RowIdentityPreserved:        r.WeightedRowPreserved,
			FanoPointOrLineAssigned:     false,
			NeedsContactFanoBijection:   true,
			PossibleAssignments:         possibleAssignments,
			LocalVariableDerived:        false,
			ConstraintSemanticDerived:   false,
			GaugeRepresentationDerived:  false,
			LorentzKineticDerived:       false,
			MassActivationDerived:       false,
			DecouplingRuleDerived:       false,
			RepresentationOrZeroDerived: false,
			CanEnterBetaTensor:          false,
			ZeroRowProved:               false,
			Status:                      RowSemanticOpen,
			Reason:                      "incidence degree is uniform and preserves spectral rows, but supplies no contact-Fano identity, local variable, constraint semantics, or representation row",
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

func buildAttempts(s SemanticsSummary) []SemanticsAttempt {
	return []SemanticsAttempt{
		{
			Name:                      "uniform Fano-degree incidence weighting",
			Kind:                      UniformIncidenceWeightAttempt,
			Constructed:               true,
			CanonicalUnderCurrentData: true,
			PreservesSpectralValues:   true,
			PreservesRowIdentity:      true,
			AddsRowSemantics:          false,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"Fano point/line identity", "local variables", "representation row"},
			Detail:                    fmt.Sprintf("all Fano points and lines have degree %d, so the weighting only rescales every contact row", s.IncidenceDegree),
		},
		{
			Name:                      "signed line incidence weighting",
			Kind:                      SignedIncidenceWeightAttempt,
			Constructed:               true,
			CanonicalUnderCurrentData: false,
			UsesExtraConvention:       true,
			PreservesSpectralValues:   false,
			RequiresBijection:         true,
			PossibleAssignments:       s.SignedIncidenceChoices,
			RejectedAsPremature:       true,
			MissingTerms:              []string{"canonical contact-to-Fano assignment", "orientation selector", "row semantics"},
			Detail:                    fmt.Sprintf("signed incidence can be transported only after selecting one of %d contact-Fano labelings", s.SignedIncidenceChoices),
		},
		{
			Name:                    "incidence-weighted spectral moments",
			Kind:                    MomentReconstructionAttempt,
			Constructed:             true,
			PreservesSpectralValues: true,
			PreservesRowIdentity:    false,
			AddsRowSemantics:        false,
			RejectedAsPremature:     true,
			MissingTerms:            []string{"row identity", "Fano semantics", "field variable map"},
			Detail:                  "moments summarize the weighted spectrum but do not reconstruct row-level local variables or representation data",
		},
		{
			Name:                  "local variable reconstruction",
			Kind:                  LocalVariableAttempt,
			Constructed:           true,
			LocalVariablesDerived: false,
			RejectedAsPremature:   true,
			MissingTerms:          []string{"base support", "section variables", "Lorentz kinetic row", "pole/residue theorem"},
			Detail:                "no local continuum variable is constructed from the incidence-weighted spectrum",
		},
		{
			Name:                 "constraint semantic map",
			Kind:                 ConstraintSemanticAttempt,
			Constructed:          true,
			ConstraintMapDerived: false,
			ZeroRowProved:        false,
			RejectedAsPremature:  true,
			MissingTerms:         []string{"nonzero nilpotent differential", "pairing", "exactness/cancellation ledger"},
			Detail:               "Gate 114 already showed the zero differential leaves all seven classes alive; incidence weights do not supply a stronger constraint complex",
		},
		{
			Name:                     "representation-row rule",
			Kind:                     RepresentationRuleAttempt,
			Constructed:              true,
			RepresentationRowDerived: false,
			BetaRowPermitted:         false,
			RejectedAsPremature:      true,
			MissingTerms:             []string{"SU(3)c×SU(2)L×U(1)Y row", "hypercharge", "mass activation", "decoupling rule"},
			Detail:                   "no beta-matching row is produced by incidence weighting",
		},
	}
}

func buildCriteria(s SemanticsSummary) []SemanticsCriterion {
	return []SemanticsCriterion{
		{Name: "uniform incidence degree", Required: true, Derived: s.IncidenceDegree == 3, Detail: "Fano points and lines have uniform degree/size three"},
		{Name: "weighted spectrum remains distinct", Required: true, Derived: s.WeightedValuesDistinct == 7, Detail: "degree-three weighting preserves the seven distinct overlap values"},
		{Name: "uniform weight adds row semantics", Required: true, Derived: s.UniformWeightAddsSemantics, Detail: "must be false: uniform scaling supplies no point/line identity"},
		{Name: "signed incidence canonical assignment", Required: true, Derived: s.CanonicalSignedAssignments > 0, Detail: "must be false: signed attachment requires a hidden contact-Fano labeling"},
		{Name: "local variable or constraint semantics", Required: true, Derived: s.LocalVariableRows > 0 || s.ConstraintSemanticRows > 0, Detail: "must be false under current data"},
		{Name: "representation-complete contact beta rows", Required: true, Derived: s.RepresentationCompleteRows > 0 || s.ContactBetaRowsAllowed > 0, Detail: "must be false: beta firewall remains closed"},
	}
}

func countRows(rows []ContactSemanticRow) rowCounts {
	var c rowCounts
	for _, r := range rows {
		c.contact++
		if r.Value > 0 {
			c.positive++
		}
		if r.SpectralValueRecovered {
			c.surviving++
		}
		if r.Status == RowSemanticOpen {
			c.open++
		}
	}
	return c
}

func distinctWeightedValues(rows []ContactSemanticRow, eps float64) int {
	if len(rows) == 0 {
		return 0
	}
	values := make([]float64, 0, len(rows))
	for _, r := range rows {
		values = append(values, r.IncidenceWeightedValue)
	}
	sort.Float64s(values)
	count := 1
	last := values[0]
	for _, v := range values[1:] {
		if math.Abs(v-last) > eps {
			count++
			last = v
		}
	}
	return count
}

func FormatRows(rows []ContactSemanticRow, limit int) string {
	if limit <= 0 || limit > len(rows) {
		limit = len(rows)
	}
	parts := make([]string, 0, limit)
	for i := 0; i < limit; i++ {
		r := rows[i]
		parts = append(parts, fmt.Sprintf("%s[value=%.10f weighted=%.10f local=%t constraint=%t rep=%t beta=%t]", r.Name, r.Value, r.IncidenceWeightedValue, r.LocalVariableDerived, r.ConstraintSemanticDerived, r.GaugeRepresentationDerived, r.CanEnterBetaTensor))
	}
	if limit < len(rows) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(rows)-limit))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatAttempts(xs []SemanticsAttempt) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s(kind=%s canonical=%t semantics=%t rep=%t beta=%t choices=%d)", x.Name, x.Kind, x.CanonicalUnderCurrentData, x.AddsRowSemantics || x.LocalVariablesDerived || x.ConstraintMapDerived, x.RepresentationRowDerived, x.BetaRowPermitted, x.PossibleAssignments))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCriteria(xs []SemanticsCriterion) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s=%t", x.Name, x.Derived))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatSummary(s SemanticsSummary) string {
	return fmt.Sprintf("rows=%d degree=%d distinctWeighted=%d uniformAddsSemantics=%t signedChoices=%d canonicalSigned=%d repRows=%d betaRows=%d", s.ContactRows, s.IncidenceDegree, s.WeightedValuesDistinct, s.UniformWeightAddsSemantics, s.SignedIncidenceChoices, s.CanonicalSignedAssignments, s.RepresentationCompleteRows, s.ContactBetaRowsAllowed)
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
