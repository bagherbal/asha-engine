// Package gaugecouplingboundaryseal implements Gate 199: gauge-coupling
// boundary seal / symbolic RG evaluation firewall audit.
//
// Gate 198 built the exact rational fermion threshold beta-row scaffold under
// three explicit seals: empirical Yukawa texture, empirical VEV scale, and a
// continuum decoupling-scheme convention.  Gate 199 asks what additional data
// are needed to turn that symbolic threshold tree into an evaluable RG
// trajectory.  It admits UV boundary data only as quarantined seals, constructs
// the closed-form symbolic one-loop expression, and separates top-down
// prediction scaffolding from bottom-up phenomenological convergence audits.
//
// No observed coupling, boundary scale, topological u=1 branch, W/Z threshold,
// finite matching correction, or physical RG prediction is derived here.
package gaugecouplingboundaryseal

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/conditionalthresholdbeta"
)

type Rational = conditionalthresholdbeta.Rational

func R(num, den int64) Rational { return conditionalthresholdbeta.NewRational(num, den) }

type BetaVector struct {
	U1Y  Rational
	SU2L Rational
	SU3C Rational
}

func (v BetaVector) String() string {
	return fmt.Sprintf("(%s,%s,%s)", v.U1Y, v.SU2L, v.SU3C)
}

func (v BetaVector) Equal(w BetaVector) bool {
	return v.U1Y.Equal(w.U1Y) && v.SU2L.Equal(w.SU2L) && v.SU3C.Equal(w.SU3C)
}

type BoundaryScaleSeal struct {
	Name                         string
	AxiomID                      string
	Symbol                       string
	Dimension                    string
	ExplicitBoundaryData         bool
	Quarantined                  bool
	RequiredForTopDownEvaluation bool
	DerivedFromFiniteAlgebra     bool
	DerivedFromTopologicalSeal   bool
	ObservedValueInserted        bool
	DownstreamMustDeclareSeal    bool
	Verdict                      string
}

type AbsoluteCouplingSeal struct {
	Name                         string
	AxiomID                      string
	Symbol                       string
	Definition                   string
	Dimensionless                bool
	ExplicitBoundaryData         bool
	Quarantined                  bool
	RequiredForTopDownEvaluation bool
	DerivedFromFiniteAlgebra     bool
	DerivedFromEightPiSquared    bool
	UnitTopologicalBranchAssumed bool
	ObservedValueInserted        bool
	DownstreamMustDeclareSeal    bool
	Verdict                      string
}

type BoundaryAxiomAudit struct {
	Gate198ThresholdTreeInherited                  bool
	BoundaryScaleSealInserted                      bool
	AbsoluteCouplingSealInserted                   bool
	TopDownUVParametersAccepted                    bool
	BottomUpIRInputsAllowedForAudit                bool
	BottomUpIRInputsDerived                        bool
	IRCouplingsQuarantinedIfUsed                   bool
	TopologicalUOneBranchAvailableOnlyAsComparison bool
	BoundaryDataDerivedFromFinite                  bool
	ObservedInputsUsedInDefaultBuild               bool
	Verdict                                        string
}

type SymbolicRGTrajectory struct {
	BaselineBetaVector              BetaVector
	FermionContributionVector       BetaVector
	ThresholdRows                   int
	Expression                      string
	ThresholdSum                    string
	LogKernel                       string
	TopDownInputs                   []string
	ConditionalInputs               []string
	PiecewiseClosedFormBuilt        bool
	TreeLevelContinuityInherited    bool
	FiniteMatchingCorrectionsSealed bool
	ThresholdOrderingKnown          bool
	EvaluatedNumerically            bool
	PhysicalPredictionMade          bool
	UsesObservedCouplings           bool
	Verdict                         string
}

type BottomUpConvergenceAudit struct {
	AllowedAsPhenomenologicalAudit     bool
	RequiresIRCouplingSeal             bool
	IRBoundaryScaleSymbol              string
	IRInputSymbols                     []string
	InvertibilityEquationsBuilt        bool
	PairwiseDifferenceEquationsBuilt   bool
	CanSolveFormalPairwiseLogIntervals bool
	SingleUVIntersectionCondition      string
	CanTestUEqualsOneIfInputsProvided  bool
	TopologicalUOneDerived             bool
	IRInputsDerived                    bool
	ThresholdOrderingKnown             bool
	NumericalConvergenceDetermined     bool
	ReducesStrictNullity               bool
	Verdict                            string
}

type EvaluationDomainAudit struct {
	FormalFermionThresholdsAvailable bool
	FermionThresholdOrderingKnown    bool
	GaugeBosonThresholdsAvailable    bool
	WZThresholdsBlocked              bool
	RunToMZAllowed                   bool
	DeepInfraredFlowDefined          bool
	OutOfBoundsRule                  string
	AllowedDomainStatement           string
	Verdict                          string
}

type FirewallAudit struct {
	Gate198Inherited                           bool
	BoundaryScaleDerivedStrict                 bool
	AbsoluteCouplingDerivedStrict              bool
	GaugeCouplingsDerived                      bool
	TopologicalEightPiSquaredImported          bool
	FiniteToContinuumScaleDerived              bool
	ObservedInputsImported                     bool
	PhysicalRGPredictionMade                   bool
	NumericalTrajectoryEvaluated               bool
	WZThresholdsDerived                        bool
	ThresholdOrderingDerived                   bool
	FiniteMatchingCorrectionsDerived           bool
	BottomUpAuditCanDeriveFiniteTheorem        bool
	StrictNullityBefore                        int
	StrictNullityAfter                         int
	ConditionalBoundarySealNullityBefore       int
	ConditionalBoundarySealNullityAfter        int
	ConditionalSymbolicEvaluationNullityBefore int
	ConditionalSymbolicEvaluationNullityAfter  int
	PhysicalPredictionNullityBefore            int
	PhysicalPredictionNullityAfter             int
	OpenRequirements                           []string
	RecommendedNextGate                        string
	Verdict                                    string
}

type Summary struct {
	TestsAudited                      int
	UVBoundarySealsRecorded           bool
	SymbolicTopDownTrajectoryBuilt    bool
	BottomUpViabilityAuditSeparated   bool
	LowEnergyDomainFirewallPreserved  bool
	AbsoluteCouplingFirewallPreserved bool
	NoObservedInputsInDefaultBuild    bool
	Comment                           string
}

type Analysis struct {
	PreviousGate198 conditionalthresholdbeta.Analysis
	ScaleSeal       BoundaryScaleSeal
	CouplingSeal    AbsoluteCouplingSeal
	BoundaryAudit   BoundaryAxiomAudit
	Trajectory      SymbolicRGTrajectory
	BottomUp        BottomUpConvergenceAudit
	Domain          EvaluationDomainAudit
	Firewall        FirewallAudit
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := conditionalthresholdbeta.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 198 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev)
	})
	return defaultA, defaultErr
}

func Build(prev conditionalthresholdbeta.Analysis) (Analysis, error) {
	if !prev.Summary.PiecewiseRGScaffoldBuilt || !prev.Firewall.ConditionalBetaRowsActivated || prev.Firewall.PhysicalRGFlowEvaluated {
		return Analysis{}, fmt.Errorf("Gate 199 requires Gate 198 symbolic threshold scaffolding with physical RG evaluation still sealed")
	}
	if prev.Firewall.AbsoluteBoundaryScaleDerived || prev.Firewall.AbsoluteBoundaryCouplingDerived || prev.Firewall.GaugeCouplingsDerived || prev.Firewall.ObservedInputsImported || prev.Firewall.TopologicalEightPiSquaredImported {
		return Analysis{}, fmt.Errorf("Gate 199 refuses leaked boundary scale, boundary coupling, gauge coupling, observed input, or 8π² import")
	}

	scale := buildBoundaryScaleSeal()
	coupling := buildAbsoluteCouplingSeal()
	boundary := auditBoundaryAxioms(prev, scale, coupling)
	traj := buildSymbolicTrajectory(prev)
	bottom := auditBottomUpConvergence(prev, traj)
	domain := auditDomain(prev)
	fw := auditFirewall(prev, scale, coupling, boundary, traj, bottom, domain)
	summary := Summary{
		TestsAudited:                      7,
		UVBoundarySealsRecorded:           scale.ExplicitBoundaryData && coupling.ExplicitBoundaryData && scale.Quarantined && coupling.Quarantined,
		SymbolicTopDownTrajectoryBuilt:    traj.PiecewiseClosedFormBuilt && !traj.EvaluatedNumerically && !traj.PhysicalPredictionMade,
		BottomUpViabilityAuditSeparated:   bottom.AllowedAsPhenomenologicalAudit && bottom.RequiresIRCouplingSeal && !bottom.ReducesStrictNullity,
		LowEnergyDomainFirewallPreserved:  domain.WZThresholdsBlocked && !domain.RunToMZAllowed && !domain.DeepInfraredFlowDefined,
		AbsoluteCouplingFirewallPreserved: !fw.AbsoluteCouplingDerivedStrict && !fw.GaugeCouplingsDerived && !fw.TopologicalEightPiSquaredImported && fw.StrictNullityBefore == fw.StrictNullityAfter,
		NoObservedInputsInDefaultBuild:    !fw.ObservedInputsImported && !boundary.ObservedInputsUsedInDefaultBuild && !traj.UsesObservedCouplings,
		Comment:                           "Gate 199 records UV boundary data as explicit seals, builds the symbolic threshold-corrected RG trajectory, and permits bottom-up convergence only as a quarantined viability audit. It does not derive M*, u*, observed couplings, W/Z thresholds, finite matching corrections, or a physical prediction.",
	}
	truth := "Gate 199 converts Gate 198's threshold beta-row scaffold into an evaluable symbolic RG form only after two new quarantined boundary inputs: M* and u*=1/g_*². Top-down UV parameters are accepted as seals for formal IR trajectories. Bottom-up IR coupling inputs may be used only in a separate phenomenological convergence audit that checks symbolic intersection conditions; it does not derive the UV boundary. The finite algebra still does not supply M*, u*, W/Z thresholds, threshold ordering, finite matching corrections, 8π² normalization, physical gauge couplings, or observed low-energy values."
	return Analysis{PreviousGate198: prev, ScaleSeal: scale, CouplingSeal: coupling, BoundaryAudit: boundary, Trajectory: traj, BottomUp: bottom, Domain: domain, Firewall: fw, Summary: summary, TruthStatement: truth}, nil
}

func buildBoundaryScaleSeal() BoundaryScaleSeal {
	return BoundaryScaleSeal{
		Name:                         "BoundaryScaleSeal",
		AxiomID:                      "CONDITIONAL_ON_BOUNDARY_SCALE",
		Symbol:                       "M_*",
		Dimension:                    "energy",
		ExplicitBoundaryData:         true,
		Quarantined:                  true,
		RequiredForTopDownEvaluation: true,
		DerivedFromFiniteAlgebra:     false,
		DerivedFromTopologicalSeal:   false,
		ObservedValueInserted:        false,
		DownstreamMustDeclareSeal:    true,
		Verdict:                      "M_* is a dimensional boundary parameter required to evaluate logs; finite topology supplies no energy ruler",
	}
}

func buildAbsoluteCouplingSeal() AbsoluteCouplingSeal {
	return AbsoluteCouplingSeal{
		Name:                         "AbsoluteCouplingSeal",
		AxiomID:                      "CONDITIONAL_ON_ABSOLUTE_BOUNDARY_COUPLING",
		Symbol:                       "u_*",
		Definition:                   "u_* = 1/g_*^2",
		Dimensionless:                true,
		ExplicitBoundaryData:         true,
		Quarantined:                  true,
		RequiredForTopDownEvaluation: true,
		DerivedFromFiniteAlgebra:     false,
		DerivedFromEightPiSquared:    false,
		UnitTopologicalBranchAssumed: false,
		ObservedValueInserted:        false,
		DownstreamMustDeclareSeal:    true,
		Verdict:                      "u_* is the absolute intercept of the RG trajectory; the engine may carry it symbolically but does not derive or set it to one",
	}
}

func auditBoundaryAxioms(prev conditionalthresholdbeta.Analysis, scale BoundaryScaleSeal, coupling AbsoluteCouplingSeal) BoundaryAxiomAudit {
	return BoundaryAxiomAudit{
		Gate198ThresholdTreeInherited:                  prev.Summary.PiecewiseRGScaffoldBuilt && prev.Firewall.ConditionalBetaRowsActivated,
		BoundaryScaleSealInserted:                      scale.ExplicitBoundaryData && scale.Quarantined,
		AbsoluteCouplingSealInserted:                   coupling.ExplicitBoundaryData && coupling.Quarantined,
		TopDownUVParametersAccepted:                    true,
		BottomUpIRInputsAllowedForAudit:                true,
		BottomUpIRInputsDerived:                        false,
		IRCouplingsQuarantinedIfUsed:                   true,
		TopologicalUOneBranchAvailableOnlyAsComparison: true,
		BoundaryDataDerivedFromFinite:                  false,
		ObservedInputsUsedInDefaultBuild:               false,
		Verdict:                                        "top-down M*,u* seals are allowed for symbolic evaluation; bottom-up IR data are allowed only as quarantined comparison/audit inputs",
	}
}

func buildSymbolicTrajectory(prev conditionalthresholdbeta.Analysis) SymbolicRGTrajectory {
	baseline := BetaVector{U1Y: R(41, 10), SU2L: R(-19, 6), SU3C: R(-7, 1)}
	fermion := BetaVector{U1Y: prev.Ledger.FermionContributionU1Y, SU2L: prev.Ledger.FermionContributionSU2L, SU3C: prev.Ledger.FermionContributionSU3C}
	thresholdSum := "Σ_{f,g | M_{f,g}>μ} Δb_{i,f,g} log(M_{f,g}/μ)"
	expr := "A_i(μ)=u_* + (b_i/8π²) log(M_*/μ) + (1/8π²) " + thresholdSum + ", with A_i=1/g_i² and Δb below threshold equal to the negative row from Gate 198"
	return SymbolicRGTrajectory{
		BaselineBetaVector:              baseline,
		FermionContributionVector:       fermion,
		ThresholdRows:                   prev.Ledger.ThresholdRows,
		Expression:                      expr,
		ThresholdSum:                    thresholdSum,
		LogKernel:                       "log(high scale / low scale)",
		TopDownInputs:                   []string{"M_*", "u_*"},
		ConditionalInputs:               []string{"Yukawa singular values σ_f,g", "VEV v", "decoupling scheme Θ", "formal masses M_f,g=(v/√2)σ_f,g"},
		PiecewiseClosedFormBuilt:        true,
		TreeLevelContinuityInherited:    prev.SchemeSeal.TreeLevelContinuityEnforced,
		FiniteMatchingCorrectionsSealed: !prev.SchemeSeal.FiniteMatchingCorrectionsDerived,
		ThresholdOrderingKnown:          false,
		EvaluatedNumerically:            false,
		PhysicalPredictionMade:          false,
		UsesObservedCouplings:           false,
		Verdict:                         "closed-form symbolic one-loop threshold tree is assembled, but all dimensional boundaries and threshold orderings remain symbolic",
	}
}

func auditBottomUpConvergence(prev conditionalthresholdbeta.Analysis, traj SymbolicRGTrajectory) BottomUpConvergenceAudit {
	return BottomUpConvergenceAudit{
		AllowedAsPhenomenologicalAudit:     true,
		RequiresIRCouplingSeal:             true,
		IRBoundaryScaleSymbol:              "μ_IR",
		IRInputSymbols:                     []string{"A_1(μ_IR)", "A_2(μ_IR)", "A_3(μ_IR)"},
		InvertibilityEquationsBuilt:        traj.PiecewiseClosedFormBuilt,
		PairwiseDifferenceEquationsBuilt:   true,
		CanSolveFormalPairwiseLogIntervals: true,
		SingleUVIntersectionCondition:      "L_12(thresholds,IR)=L_13(thresholds,IR)=L_23(thresholds,IR) with one common M_* and one common u_*",
		CanTestUEqualsOneIfInputsProvided:  true,
		TopologicalUOneDerived:             false,
		IRInputsDerived:                    false,
		ThresholdOrderingKnown:             prev.Firewall.ThresholdOrderingKnown,
		NumericalConvergenceDetermined:     false,
		ReducesStrictNullity:               false,
		Verdict:                            "bottom-up running is permitted only as a quarantined viability test; without IR input values and threshold ordering the engine builds equations but determines no convergence",
	}
}

func auditDomain(prev conditionalthresholdbeta.Analysis) EvaluationDomainAudit {
	return EvaluationDomainAudit{
		FormalFermionThresholdsAvailable: prev.Firewall.FormalMassThresholdsAvailable && prev.Ledger.ThresholdRows == 12,
		FermionThresholdOrderingKnown:    prev.Firewall.ThresholdOrderingKnown,
		GaugeBosonThresholdsAvailable:    prev.Domain.GaugeBosonThresholdsAvailable,
		WZThresholdsBlocked:              prev.Domain.WZThresholdsBlocked,
		RunToMZAllowed:                   prev.Domain.RunToMZAllowed,
		DeepInfraredFlowDefined:          prev.Domain.DeepInfraredFlowDefined,
		OutOfBoundsRule:                  "reject numerical evaluation in domains requiring W/Z thresholds, deep-IR electromagnetic matching, or an undeclared threshold ordering",
		AllowedDomainStatement:           "symbolic evaluation is allowed only on declared broken-phase intervals whose active fermion thresholds and continuum scheme are explicitly sealed",
		Verdict:                          "the symbolic tree exists, but evaluation below the electroweak gauge-boson threshold domain remains out of bounds",
	}
}

func auditFirewall(prev conditionalthresholdbeta.Analysis, scale BoundaryScaleSeal, coupling AbsoluteCouplingSeal, boundary BoundaryAxiomAudit, traj SymbolicRGTrajectory, bottom BottomUpConvergenceAudit, domain EvaluationDomainAudit) FirewallAudit {
	return FirewallAudit{
		Gate198Inherited:                           boundary.Gate198ThresholdTreeInherited,
		BoundaryScaleDerivedStrict:                 scale.DerivedFromFiniteAlgebra || scale.DerivedFromTopologicalSeal,
		AbsoluteCouplingDerivedStrict:              coupling.DerivedFromFiniteAlgebra || coupling.DerivedFromEightPiSquared || coupling.UnitTopologicalBranchAssumed,
		GaugeCouplingsDerived:                      false,
		TopologicalEightPiSquaredImported:          false,
		FiniteToContinuumScaleDerived:              false,
		ObservedInputsImported:                     scale.ObservedValueInserted || coupling.ObservedValueInserted || boundary.ObservedInputsUsedInDefaultBuild || traj.UsesObservedCouplings,
		PhysicalRGPredictionMade:                   traj.PhysicalPredictionMade || bottom.NumericalConvergenceDetermined,
		NumericalTrajectoryEvaluated:               traj.EvaluatedNumerically,
		WZThresholdsDerived:                        domain.GaugeBosonThresholdsAvailable,
		ThresholdOrderingDerived:                   domain.FermionThresholdOrderingKnown || bottom.ThresholdOrderingKnown,
		FiniteMatchingCorrectionsDerived:           !traj.FiniteMatchingCorrectionsSealed,
		BottomUpAuditCanDeriveFiniteTheorem:        bottom.ReducesStrictNullity,
		StrictNullityBefore:                        prev.Firewall.StrictNullityAfter,
		StrictNullityAfter:                         prev.Firewall.StrictNullityAfter,
		ConditionalBoundarySealNullityBefore:       1,
		ConditionalBoundarySealNullityAfter:        0,
		ConditionalSymbolicEvaluationNullityBefore: 1,
		ConditionalSymbolicEvaluationNullityAfter:  0,
		PhysicalPredictionNullityBefore:            1,
		PhysicalPredictionNullityAfter:             1,
		OpenRequirements: []string{
			"derive or seal numerical boundary scale M_*",
			"derive or seal absolute boundary intercept u_*",
			"derive W/Z thresholds or declare a low-energy electroweak matching seal",
			"derive threshold ordering from numerical masses before evaluating piecewise logs",
			"derive finite matching corrections or declare a continuum matching convention",
			"derive finite-to-continuum normalization before importing any 8π² action scale",
		},
		RecommendedNextGate: "Gate 200 — topological boundary viability / bottom-up convergence comparison audit",
		Verdict:             "boundary seals and symbolic RG evaluation are conditionally available, but no strict finite-to-continuum or physical-coupling theorem is closed",
	}
}

func FormatScaleSeal(s BoundaryScaleSeal) string {
	return fmt.Sprintf("%s[%s]: symbol=%s dimension=%s explicit=%t quarantined=%t finite-derived=%t topological-derived=%t observed=%t", s.Name, s.AxiomID, s.Symbol, s.Dimension, s.ExplicitBoundaryData, s.Quarantined, s.DerivedFromFiniteAlgebra, s.DerivedFromTopologicalSeal, s.ObservedValueInserted)
}

func FormatCouplingSeal(s AbsoluteCouplingSeal) string {
	return fmt.Sprintf("%s[%s]: symbol=%s definition=%s explicit=%t quarantined=%t finite-derived=%t 8pi2-derived=%t unit-branch-assumed=%t observed=%t", s.Name, s.AxiomID, s.Symbol, s.Definition, s.ExplicitBoundaryData, s.Quarantined, s.DerivedFromFiniteAlgebra, s.DerivedFromEightPiSquared, s.UnitTopologicalBranchAssumed, s.ObservedValueInserted)
}

func FormatBoundaryAudit(a BoundaryAxiomAudit) string {
	return fmt.Sprintf("gate198=%t topdown=%t bottomup-audit=%t ir-derived=%t ir-quarantined=%t finite-boundary-derived=%t observed-default=%t verdict=%s", a.Gate198ThresholdTreeInherited, a.TopDownUVParametersAccepted, a.BottomUpIRInputsAllowedForAudit, a.BottomUpIRInputsDerived, a.IRCouplingsQuarantinedIfUsed, a.BoundaryDataDerivedFromFinite, a.ObservedInputsUsedInDefaultBuild, a.Verdict)
}

func FormatTrajectory(t SymbolicRGTrajectory) string {
	return fmt.Sprintf("b=%s fermion=%s rows=%d built=%t evaluated=%t prediction=%t ordering-known=%t matching-corrections-sealed=%t expr=%s", t.BaselineBetaVector, t.FermionContributionVector, t.ThresholdRows, t.PiecewiseClosedFormBuilt, t.EvaluatedNumerically, t.PhysicalPredictionMade, t.ThresholdOrderingKnown, t.FiniteMatchingCorrectionsSealed, t.Expression)
}

func FormatBottomUp(b BottomUpConvergenceAudit) string {
	return fmt.Sprintf("allowed=%t requires-ir-seal=%t equations=%t pairwise=%t formal-L=%t common-intersection=%q u=1-test=%t u=1-derived=%t numerical-convergence=%t strict-reduction=%t", b.AllowedAsPhenomenologicalAudit, b.RequiresIRCouplingSeal, b.InvertibilityEquationsBuilt, b.PairwiseDifferenceEquationsBuilt, b.CanSolveFormalPairwiseLogIntervals, b.SingleUVIntersectionCondition, b.CanTestUEqualsOneIfInputsProvided, b.TopologicalUOneDerived, b.NumericalConvergenceDetermined, b.ReducesStrictNullity)
}

func FormatDomain(d EvaluationDomainAudit) string {
	return fmt.Sprintf("fermion-thresholds=%t ordering-known=%t wz-thresholds=%t wz-blocked=%t run-to-MZ=%t deep-IR=%t rule=%s", d.FormalFermionThresholdsAvailable, d.FermionThresholdOrderingKnown, d.GaugeBosonThresholdsAvailable, d.WZThresholdsBlocked, d.RunToMZAllowed, d.DeepInfraredFlowDefined, d.OutOfBoundsRule)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gate198=%t M*-derived=%t u*-derived=%t gauge-couplings=%t 8pi2=%t continuum-scale=%t observed=%t prediction=%t numeric=%t WZ=%t ordering=%t matching=%t bottomup-strict=%t strict-nullity=%d->%d boundary-nullity=%d->%d symbolic-nullity=%d->%d physical-nullity=%d->%d open=[%s] next=%s", f.Gate198Inherited, f.BoundaryScaleDerivedStrict, f.AbsoluteCouplingDerivedStrict, f.GaugeCouplingsDerived, f.TopologicalEightPiSquaredImported, f.FiniteToContinuumScaleDerived, f.ObservedInputsImported, f.PhysicalRGPredictionMade, f.NumericalTrajectoryEvaluated, f.WZThresholdsDerived, f.ThresholdOrderingDerived, f.FiniteMatchingCorrectionsDerived, f.BottomUpAuditCanDeriveFiniteTheorem, f.StrictNullityBefore, f.StrictNullityAfter, f.ConditionalBoundarySealNullityBefore, f.ConditionalBoundarySealNullityAfter, f.ConditionalSymbolicEvaluationNullityBefore, f.ConditionalSymbolicEvaluationNullityAfter, f.PhysicalPredictionNullityBefore, f.PhysicalPredictionNullityAfter, strings.Join(f.OpenRequirements, "; "), f.RecommendedNextGate)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d uv-seals=%t symbolic-topdown=%t bottomup-separated=%t low-energy-firewall=%t absolute-firewall=%t no-observed=%t comment=%s", s.TestsAudited, s.UVBoundarySealsRecorded, s.SymbolicTopDownTrajectoryBuilt, s.BottomUpViabilityAuditSeparated, s.LowEnergyDomainFirewallPreserved, s.AbsoluteCouplingFirewallPreserved, s.NoObservedInputsInDefaultBuild, s.Comment)
}
