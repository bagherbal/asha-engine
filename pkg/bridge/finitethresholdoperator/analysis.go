// Package finitethresholdoperator implements Gate 178: finite threshold operator /
// decoupling spectrum search.
//
// Gate 177 showed that non-universal Δb_i deformations can repair the external
// M_Z comparison ledger only as an underived fit family. Gate 178 asks whether
// the engine already contains a finite object that can provide such a threshold
// deformation lawfully. The required chain is strict:
//
//	finite mode -> physical activation/decoupling predicate -> gauge representation
//	-> beta-index row -> non-universal Δb_i contribution.
//
// The gate audits every currently derived finite threshold candidate and records
// whether it has all pieces at once. Observed comparison data inherited from
// Gate 176/177 remains quarantined and is never used as finite-core input.
package finitethresholdoperator

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/betamatching"
	"github.com/bagherbal/asha-engine/pkg/bridge/normalizationthresholdaudit"
)

type CandidateClass string

const (
	BaselineFieldClass        CandidateClass = "baseline-continuum-field"
	FiniteSpectrumClass       CandidateClass = "finite-spectrum-anchor"
	CollectiveSpectralClass   CandidateClass = "collective-spectral-functional"
	YukawaMassClass           CandidateClass = "yukawa-mass-data"
	PhenomenologicalFitClass  CandidateClass = "phenomenological-fit-vector"
	TopologicalNormalizeClass CandidateClass = "topological-normalization-seal"
)

type ThresholdCandidate struct {
	Name                    string
	SourceGate              string
	Class                   CandidateClass
	FiniteSpectrum          bool
	ExactFiniteData         bool
	PhysicalMassUnit        bool
	ActivationPredicate     bool
	GaugeRepresentation     bool
	DecouplingRule          bool
	BetaIndexRow            bool
	NonUniversalSectorRow   bool
	BaselineAlreadyCounted  bool
	HeavyThresholdCandidate bool
	DerivedFromFinite       bool
	UsesObservedComparison  bool
	CanCorrectBeta          bool
	CompleteThresholdOp     bool
	Verdict                 string
}

type RequirementAudit struct {
	RequiredPieces             []string
	CandidatesAudited          int
	WithFiniteSpectrum         int
	WithExactFiniteData        int
	WithPhysicalMassUnit       int
	WithActivationPredicate    int
	WithGaugeRepresentation    int
	WithDecouplingRule         int
	WithBetaIndexRow           int
	WithNonUniversalSectorRow  int
	CompleteThresholdOperators int
	FiniteDerivedThresholdOps  int
	ObservedFitVectors         int
	BaselineRowsAlreadyCounted int
	OpenFiniteSpectrumAnchors  int
	NoCandidateHasAllPieces    bool
	NeededForGate177Repair     string
}

type CombinationAttempt struct {
	Name                      string
	Inputs                    []string
	Constructed               bool
	FiniteSpectrum            bool
	ActivationPredicate       bool
	GaugeRepresentation       bool
	DecouplingRule            bool
	BetaIndexRow              bool
	NonUniversalSectorRow     bool
	UsesObservedComparison    bool
	AdmissibleAsFiniteTheorem bool
	CanRepairGate177Strictly  bool
	Reason                    string
}

type DeltaBWitnessAudit struct {
	Gate177NonUniversalFitExists      bool
	Gate177FiniteThresholdDerived     bool
	MinimumNormWitnessUsesExternalFit bool
	MinimumNormDeltaB                 [3]float64
	MinimumNormL                      float64
	MinimumNormU                      float64
	SignPatternPreserved              bool
	CanBePromotedToFiniteOperator     bool
	Reason                            string
}

type FirewallAudit struct {
	ThresholdOperatorDerived         bool
	PhysicalMassSpectrumDerived      bool
	ActivationPredicateDerived       bool
	DecouplingRuleDerived            bool
	GaugeRepresentationRowsCompleted bool
	NonUniversalDeltaBDerived        bool
	ThresholdCorrectedBetaDerived    bool
	UsesObservedInputForDerivation   bool
	Gate177RepairPromoted            bool
	StrictNullityBefore              int
	StrictNullityAfter               int
	ConditionalNullityBefore         int
	ConditionalNullityAfter          int
	PhysicalConstantsDerived         bool
	BoundaryScaleDerived             bool
	RemainingMissingObjects          []string
	RecommendedNextGate              string
	Verdict                          string
}

type Analysis struct {
	PreviousGate177 normalizationthresholdaudit.Analysis
	BetaMatching    betamatching.Analysis

	Candidates     []ThresholdCandidate
	Requirements   RequirementAudit
	Combinations   []CombinationAttempt
	DeltaBWitness  DeltaBWitnessAudit
	Firewall       FirewallAudit
	TruthStatement string
	RejectedClaims []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := normalizationthresholdaudit.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		bm, err := betamatching.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, bm)
	})
	return defaultValue, defaultErr
}

func Build(prev normalizationthresholdaudit.Analysis, bm betamatching.Analysis) (Analysis, error) {
	if prev.Firewall.StrictNullityAfter != 3 || prev.Firewall.ConditionalNullityAfter != 2 {
		return Analysis{}, fmt.Errorf("Gate 178 requires Gate 177 to leave strict nullity 3 and conditional nullity 2")
	}
	if !prev.Firewall.NonUniversalThresholdCanFitByConstruction || prev.Firewall.NonUniversalThresholdDerived || prev.Firewall.ThresholdCorrectionsDerived {
		return Analysis{}, fmt.Errorf("Gate 178 requires Gate 177's non-universal fit family to remain underived")
	}
	if prev.Firewall.HiddenObservedInputUsedForDerivation || !prev.Firewall.UsesObservedInputOnlyForComparison {
		return Analysis{}, fmt.Errorf("Gate 178 refuses non-quarantined observed input")
	}
	if bm.ThresholdCorrectedBetaDerived || bm.FullFiniteBetaMatchingTensorDerived || bm.BetaCorrectionRowsAllowed != 0 {
		return Analysis{}, fmt.Errorf("Gate 178 requires previous finite beta-matching firewall to remain closed")
	}

	candidates := buildCandidates(prev, bm)
	req := auditRequirements(candidates)
	combos := buildCombinationAttempts()
	delta := DeltaBWitnessAudit{
		Gate177NonUniversalFitExists:      prev.Thresholds.CanRepairPhenomenologyByFit,
		Gate177FiniteThresholdDerived:     prev.Thresholds.FiniteThresholdOperatorDerived,
		MinimumNormWitnessUsesExternalFit: !prev.Thresholds.MinimumNormForUOne.FiniteDerived,
		MinimumNormDeltaB: [3]float64{
			prev.Thresholds.MinimumNormForUOne.DeltaB1,
			prev.Thresholds.MinimumNormForUOne.DeltaB2,
			prev.Thresholds.MinimumNormForUOne.DeltaB3,
		},
		MinimumNormL:                  prev.Thresholds.MinimumNormForUOne.LogIntervalL,
		MinimumNormU:                  prev.Thresholds.MinimumNormForUOne.UInverseGStar,
		SignPatternPreserved:          prev.Thresholds.MinimumNormForUOne.SignPatternPreserved,
		CanBePromotedToFiniteOperator: false,
		Reason:                        "Gate 177's Δb vector is selected from external comparison data plus a Euclidean criterion; it has no finite spectrum, no activation predicate, and no mode-to-representation map.",
	}
	firewall := FirewallAudit{
		ThresholdOperatorDerived:         req.CompleteThresholdOperators > 0 && req.FiniteDerivedThresholdOps > 0,
		PhysicalMassSpectrumDerived:      req.WithPhysicalMassUnit > 0,
		ActivationPredicateDerived:       req.WithActivationPredicate > 0,
		DecouplingRuleDerived:            req.WithDecouplingRule > 0,
		GaugeRepresentationRowsCompleted: bm.AllOpenModesRepresentationComplete,
		NonUniversalDeltaBDerived:        false,
		ThresholdCorrectedBetaDerived:    false,
		UsesObservedInputForDerivation:   false,
		Gate177RepairPromoted:            false,
		StrictNullityBefore:              prev.Firewall.StrictNullityAfter,
		StrictNullityAfter:               prev.Firewall.StrictNullityAfter,
		ConditionalNullityBefore:         prev.Firewall.ConditionalNullityAfter,
		ConditionalNullityAfter:          prev.Firewall.ConditionalNullityAfter,
		PhysicalConstantsDerived:         false,
		BoundaryScaleDerived:             false,
		RemainingMissingObjects: []string{
			"finite physical mass unit or boundary scale for threshold anchors",
			"activation predicate distinguishing baseline fields, heavy fields, regulator modes, and constrained finite modes",
			"representation rows for B-sector/contact threshold candidates",
			"decoupling/matching law assigning each active heavy mode a beta-index contribution",
			"non-universal finite operator producing Δb_i without observed comparison fitting",
		},
		RecommendedNextGate: "Gate 179 — threshold-origin dichotomy / new-sector versus continuum-decoupling bridge audit",
		Verdict:             "No finite threshold operator or decoupling spectrum is derived; Gate 177's non-universal repair remains an external fit family.",
	}
	truth := "Gate 178 audits the full threshold chain after the Gate-177 deformation audit. The engine has exact finite spectra, a baseline scalar representation row, collective contact spectral invariants, and a phenomenological Δb witness, but no single finite object has spectrum, physical activation, gauge representation, decoupling law, and non-universal beta-index row at once. Therefore non-universal thresholds remain underived and cannot reduce nullity."
	return Analysis{PreviousGate177: prev, BetaMatching: bm, Candidates: candidates, Requirements: req, Combinations: combos, DeltaBWitness: delta, Firewall: firewall, TruthStatement: truth, RejectedClaims: []string{
		"contact spectral multiplicities alone define threshold beta rows",
		"the scalar/contact active doublet baseline row is a heavy threshold correction",
		"Gate 177's fitted Δb vector is a finite operator",
		"collective zeta invariants can be inserted as non-universal beta rows without representation and decoupling data",
		"Fock Dirac Yukawa amplitudes supply threshold corrections before the amplitudes and physical masses are derived",
	}}, nil
}

func buildCandidates(prev normalizationthresholdaudit.Analysis, bm betamatching.Analysis) []ThresholdCandidate {
	return []ThresholdCandidate{
		{
			Name: "scalar/contact active aggregate doublet", SourceGate: "Gate 108 / Gate 167", Class: BaselineFieldClass,
			FiniteSpectrum: false, ExactFiniteData: true, GaugeRepresentation: bm.ScalarSectorRowConstructed, BetaIndexRow: bm.ScalarSectorRowConstructed,
			BaselineAlreadyCounted: true, DerivedFromFinite: true, CanCorrectBeta: false,
			Verdict: "representation-complete baseline scalar row Δb=(1/10,1/6,0); not a heavy threshold deformation",
		},
		{
			Name: "individual scalar active real modes", SourceGate: "Gate 37 / Gate 42", Class: FiniteSpectrumClass,
			FiniteSpectrum: true, ExactFiniteData: true, PhysicalMassUnit: false, ActivationPredicate: false, GaugeRepresentation: false, DecouplingRule: false,
			DerivedFromFinite: true, CanCorrectBeta: false, HeavyThresholdCandidate: false,
			Verdict: "finite active eigenvalues exist, but individual real modes are not separately oriented as heavy threshold fields",
		},
		{
			Name: "radial scalar response", SourceGate: "Gate 37", Class: FiniteSpectrumClass,
			FiniteSpectrum: true, ExactFiniteData: true, GaugeRepresentation: true, BetaIndexRow: true, DerivedFromFinite: true,
			Verdict: "bridge-level singlet/radial response gives zero gauge row and no heavy activation law",
		},
		{
			Name: "B-sector first spectral gap", SourceGate: "Gate 6 / Gate 42", Class: FiniteSpectrumClass,
			FiniteSpectrum: true, ExactFiniteData: true, DerivedFromFinite: true, HeavyThresholdCandidate: true,
			Verdict: "finite action gap exists, but gauge representation, activation, physical mass unit, and decoupling rule are absent",
		},
		{
			Name: "seven contact partial-overlap modes", SourceGate: "Gates 149-162", Class: FiniteSpectrumClass,
			FiniteSpectrum: true, ExactFiniteData: true, DerivedFromFinite: true, HeavyThresholdCandidate: true,
			Verdict: "exact contact spectrum exists, but local field class, representation rows, activation, and decoupling are not derived",
		},
		{
			Name: "quartic collective contact zeta block", SourceGate: "Gates 161-162", Class: CollectiveSpectralClass,
			FiniteSpectrum: true, ExactFiniteData: true, DerivedFromFinite: true,
			Verdict: "collective Galois-invariant spectral functional is branch-free, but it is not a rowwise threshold operator",
		},
		{
			Name: "Fock Dirac / Yukawa amplitude spectrum", SourceGate: "Gates 166-173", Class: YukawaMassClass,
			FiniteSpectrum: false, ExactFiniteData: false, GaugeRepresentation: true, DerivedFromFinite: false,
			Verdict: "identifies mass/Yukawa texture arena, but amplitudes and physical masses are not derived and do not supply threshold beta rows",
		},
		{
			Name: "Gate 177 non-universal Δb witness", SourceGate: "Gate 177", Class: PhenomenologicalFitClass,
			BetaIndexRow: true, NonUniversalSectorRow: true, UsesObservedComparison: true, CanCorrectBeta: true,
			DerivedFromFinite: false, CompleteThresholdOp: false,
			Verdict: fmt.Sprintf("fit vector Δb=(%.4f,%.4f,%.4f) repairs comparison by construction but has no finite operator", prev.Thresholds.MinimumNormForUOne.DeltaB1, prev.Thresholds.MinimumNormForUOne.DeltaB2, prev.Thresholds.MinimumNormForUOne.DeltaB3),
		},
		{
			Name: "topological normalization seal", SourceGate: "Gates 174-175", Class: TopologicalNormalizeClass,
			ExactFiniteData: true, DerivedFromFinite: true,
			Verdict: "conditional absolute-normalization branch, not a threshold spectrum or decoupling operator",
		},
	}
}

func auditRequirements(candidates []ThresholdCandidate) RequirementAudit {
	r := RequirementAudit{RequiredPieces: []string{"finite spectrum or mode carrier", "physical mass unit / activation predicate", "gauge representation row", "decoupling/matching law", "beta-index contribution"}, CandidatesAudited: len(candidates), NeededForGate177Repair: "a non-universal finite Δb_i operator, not merely a fitted vector"}
	for _, c := range candidates {
		if c.FiniteSpectrum {
			r.WithFiniteSpectrum++
		}
		if c.ExactFiniteData {
			r.WithExactFiniteData++
		}
		if c.PhysicalMassUnit {
			r.WithPhysicalMassUnit++
		}
		if c.ActivationPredicate {
			r.WithActivationPredicate++
		}
		if c.GaugeRepresentation {
			r.WithGaugeRepresentation++
		}
		if c.DecouplingRule {
			r.WithDecouplingRule++
		}
		if c.BetaIndexRow {
			r.WithBetaIndexRow++
		}
		if c.NonUniversalSectorRow {
			r.WithNonUniversalSectorRow++
		}
		if c.CompleteThresholdOp {
			r.CompleteThresholdOperators++
		}
		if c.CompleteThresholdOp && c.DerivedFromFinite && !c.UsesObservedComparison {
			r.FiniteDerivedThresholdOps++
		}
		if c.UsesObservedComparison {
			r.ObservedFitVectors++
		}
		if c.BaselineAlreadyCounted {
			r.BaselineRowsAlreadyCounted++
		}
		if c.FiniteSpectrum && !c.CompleteThresholdOp {
			r.OpenFiniteSpectrumAnchors++
		}
	}
	r.NoCandidateHasAllPieces = r.CompleteThresholdOperators == 0 && r.FiniteDerivedThresholdOps == 0
	return r
}

func buildCombinationAttempts() []CombinationAttempt {
	return []CombinationAttempt{
		{
			Name: "contact spectrum × scalar representation row", Inputs: []string{"seven contact modes", "baseline scalar doublet row"}, Constructed: true,
			FiniteSpectrum: true, GaugeRepresentation: true, BetaIndexRow: true, ActivationPredicate: false, DecouplingRule: false, NonUniversalSectorRow: false,
			Reason: "multiplying open contact spectra by the scalar row is a branch choice; no contact-to-doublet field map or activation law selects it",
		},
		{
			Name: "Gate 177 Δb vector × finite spectral anchors", Inputs: []string{"fitted Δb", "B/contact/scalar spectra"}, Constructed: true,
			FiniteSpectrum: true, BetaIndexRow: true, NonUniversalSectorRow: true, UsesObservedComparison: true,
			Reason: "can label finite anchors after the fact, but the Δb values come from external comparison, not from the finite spectra",
		},
		{
			Name: "quartic zeta functional as threshold row", Inputs: []string{"quartic zeta ledger"}, Constructed: true,
			FiniteSpectrum: true, NonUniversalSectorRow: false, ActivationPredicate: false, GaugeRepresentation: false, DecouplingRule: false,
			Reason: "collective Galois-invariant data is not a rowwise local field representation or decoupling spectrum",
		},
		{
			Name: "Fock Dirac eigenvalues as heavy thresholds", Inputs: []string{"Yukawa amplitudes", "Fock representation"}, Constructed: true,
			GaugeRepresentation: true, ActivationPredicate: false, DecouplingRule: false, FiniteSpectrum: false,
			Reason: "Yukawa amplitudes are the open mass-generation problem; using them as thresholds would import undetermined masses",
		},
		{
			Name: "topological normalization as threshold shift", Inputs: []string{"S_top=8π²"}, Constructed: true,
			Reason: "topological seal concerns absolute action normalization, not sector-specific decoupling or beta-index rows",
		},
	}
}

func FormatCandidate(c ThresholdCandidate) string {
	return fmt.Sprintf("%s[%s](spec=%t,rep=%t,act=%t,dec=%t,β=%t,nonuniv=%t,finite=%t,obs=%t,complete=%t): %s", c.Name, c.Class, c.FiniteSpectrum, c.GaugeRepresentation, c.ActivationPredicate, c.DecouplingRule, c.BetaIndexRow, c.NonUniversalSectorRow, c.DerivedFromFinite, c.UsesObservedComparison, c.CompleteThresholdOp, c.Verdict)
}

func FormatCandidates(c []ThresholdCandidate) string {
	parts := make([]string, 0, len(c))
	for _, x := range c {
		parts = append(parts, FormatCandidate(x))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatRequirements(r RequirementAudit) string {
	return fmt.Sprintf("audited=%d spec=%d exact=%d mass=%d activation=%d rep=%d decoupling=%d beta=%d nonuniv=%d complete=%d finiteComplete=%d observedFit=%d baseline=%d openAnchors=%d required=[%s]", r.CandidatesAudited, r.WithFiniteSpectrum, r.WithExactFiniteData, r.WithPhysicalMassUnit, r.WithActivationPredicate, r.WithGaugeRepresentation, r.WithDecouplingRule, r.WithBetaIndexRow, r.WithNonUniversalSectorRow, r.CompleteThresholdOperators, r.FiniteDerivedThresholdOps, r.ObservedFitVectors, r.BaselineRowsAlreadyCounted, r.OpenFiniteSpectrumAnchors, strings.Join(r.RequiredPieces, "; "))
}

func FormatCombination(c CombinationAttempt) string {
	return fmt.Sprintf("%s(inputs=%s,spec=%t,act=%t,rep=%t,dec=%t,β=%t,nonuniv=%t,obs=%t,strict=%t): %s", c.Name, strings.Join(c.Inputs, "+"), c.FiniteSpectrum, c.ActivationPredicate, c.GaugeRepresentation, c.DecouplingRule, c.BetaIndexRow, c.NonUniversalSectorRow, c.UsesObservedComparison, c.AdmissibleAsFiniteTheorem, c.Reason)
}

func FormatCombinations(xs []CombinationAttempt) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatCombination(x))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatDeltaBWitness(d DeltaBWitnessAudit) string {
	return fmt.Sprintf("fitExists=%t finiteDerived=%t externalFit=%t Δb=(%.6f,%.6f,%.6f) L=%.6f u=%.6f signPreserved=%t promotable=%t: %s", d.Gate177NonUniversalFitExists, d.Gate177FiniteThresholdDerived, d.MinimumNormWitnessUsesExternalFit, d.MinimumNormDeltaB[0], d.MinimumNormDeltaB[1], d.MinimumNormDeltaB[2], d.MinimumNormL, d.MinimumNormU, d.SignPatternPreserved, d.CanBePromotedToFiniteOperator, d.Reason)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("thrOp=%t mass=%t activation=%t decoupling=%t repsComplete=%t Δb=%t beta=%t obs=%t repairPromoted=%t nullity=%d->%d conditional=%d->%d constants=%t scale=%t next=%s missing=[%s]", f.ThresholdOperatorDerived, f.PhysicalMassSpectrumDerived, f.ActivationPredicateDerived, f.DecouplingRuleDerived, f.GaugeRepresentationRowsCompleted, f.NonUniversalDeltaBDerived, f.ThresholdCorrectedBetaDerived, f.UsesObservedInputForDerivation, f.Gate177RepairPromoted, f.StrictNullityBefore, f.StrictNullityAfter, f.ConditionalNullityBefore, f.ConditionalNullityAfter, f.PhysicalConstantsDerived, f.BoundaryScaleDerived, f.RecommendedNextGate, strings.Join(f.RemainingMissingObjects, "; "))
}
