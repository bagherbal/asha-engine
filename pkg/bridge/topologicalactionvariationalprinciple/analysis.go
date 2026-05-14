// Package topologicalactionvariationalprinciple implements Gate 287:
// Topological Action Variational Principle / S_top Boundary Selector Audit.
//
// Gate 286 showed that finite NCG inner fluctuations are the correct algebraic
// replacement for continuum Hopf/Chern-Simons machinery, but a local diagnostic
// did not generate a non-trivial instanton saddle or an inverse-B_gap action.
// Gate 287 audits the proposed dynamical correction: use the exact topological
// action S_top=8π² as a variational boundary constraint on finite spectral
// moments.  The gate formalizes the constraint and derives its stationarity
// equations, but it refuses to promote S_top into a selector: with free cutoff
// moments, absolute Dirac scale, and missing physical J/hypercharge data, the
// variational problem remains underdetermined and does not select r_+ vs r_-.
package topologicalactionvariationalprinciple

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/finitencginstantonaction"
)

const (
	AuditID = "GATE287-TOPOLOGICAL-ACTION-VARIATIONAL-PRINCIPLE-S-TOP-BOUNDARY-SELECTOR-AUDIT"

	StatusGate286Inherited          = "CONDITIONAL_SUPPORT_GATE286_NCG_SADDLE_BARRIER_INHERITED"
	StatusSTopConstraintFormalized  = "CONDITIONAL_SUPPORT_S_TOP_BOUNDARY_ACTION_CONSTRAINT_FORMALIZED"
	StatusMomentModelInherited      = "CONDITIONAL_SUPPORT_SCALAR_MORITA_MOMENT_MODEL_INHERITED"
	StatusVariationalEquations      = "CONDITIONAL_SUPPORT_VARIATIONAL_EQUATIONS_DERIVED"
	StatusShapeExtremumAudited      = "CONDITIONAL_SUPPORT_SHAPE_EXTREMUM_AUDIT_COMPLETED"
	StatusCutoffUnderdetermined     = "CONDITIONAL_SUPPORT_CUTOFF_MOMENT_UNDERDETERMINATION_PROVED"
	StatusPathBCConvergenceRecorded = "CONDITIONAL_SUPPORT_PATH_B_C_CONVERGENCE_RECORDED"
	StatusFirewallsPreserved        = "CONDITIONAL_SUPPORT_TOPOLOGICAL_ACTION_VARIATIONAL_FIREWALLS_PRESERVED"

	StatusFailedRBranchNotSelected       = "FAILED_ROUTE_S_TOP_VARIATION_DOES_NOT_SELECT_R_BRANCH"
	StatusFailedFreeCutoffMoments        = "FAILED_ROUTE_VARIATIONAL_PRINCIPLE_UNDERDETERMINED_WITH_FREE_CUTOFF_MOMENTS"
	StatusFailedPhysicalJNotDerived      = "FAILED_ROUTE_PHYSICAL_J_NOT_DERIVED_AS_EXTREMUM_SYMMETRY"
	StatusFailedCutoffRatiosNotExtracted = "FAILED_ROUTE_CUTOFF_MOMENT_RATIOS_NOT_EXTRACTED"
	StatusFailedFourPiInstNotDerived     = "FAILED_ROUTE_FOUR_OVER_PI_INSTANTON_NOT_DERIVED_BY_VARIATION"
	StatusFailedHiggsBGapFirewalled      = "FAILED_ROUTE_HIGGS_AND_BGAP_DYNAMICS_STILL_FIREWALLED"
)

const (
	sTop       = 8 * math.Pi * math.Pi
	lambdaNum  = 1197
	lambdaDen  = 4624
	lambda     = float64(lambdaNum) / float64(lambdaDen)
	kappaC     = 1
	kappaQ     = 3
	bGap       = 0.1024649212
	fourOverPi = 4 / math.Pi
)

type Gate286Inheritance struct {
	NCGCalculusFormalized    bool
	InnerFluctuationBuilt    bool
	NontrivialSaddleDerived  bool
	InverseBGapActionDerived bool
	FourOverPiGenerated      bool
	IntermediateSealGranted  bool
	Verdict                  string
}

type TopologicalActionConstraint struct {
	STop                           float64
	ExactForm                      string
	SpectralActionFormula          string
	ConstraintEquation             string
	TreatedAsDerivedFiniteDatum    bool
	TreatedAsCompleteDynamics      bool
	RequiresCutoffMoments          bool
	RequiresPhysicalSpectralTriple bool
	Verdict                        string
}

type AmplitudeBranch struct {
	Name             string
	Sign             string
	R                float64
	AbsYOverX        float64
	ShapeLambda      float64
	ShapeResidualAbs float64
	D2ForXEqualsOne  float64
	D4ForXEqualsOne  float64
}

type SpectralMomentModel struct {
	KappaC             int
	KappaQ             int
	LambdaContact      float64
	D2Formula          string
	D4Formula          string
	ShapeEquation      string
	Quadratic          string
	Branches           []AmplitudeBranch
	HasTwoBranches     bool
	AbsoluteScaleKnown bool
	Verdict            string
}

type VariationalEquationAudit struct {
	Variables                            []string
	SFormula                             string
	DerivativeWRT_R                      string
	StationaryRFormula                   string
	PositiveCutoffMomentsSelectPositiveR bool
	ArbitrarySignedMomentsCanFitAnyR     bool
	ShapeDerivative                      string
	ShapeExtremumR                       float64
	ShapeExtremumLambda                  float64
	BranchesAreShapeExtrema              bool
	SelectsUpperBranch                   bool
	SelectsLowerBranch                   bool
	UniqueBranchSelected                 bool
	Verdict                              string
}

type ConstraintRankAudit struct {
	Unknowns                    []string
	NativeEquations             []string
	UnknownCount                int
	EquationCount               int
	Underdetermined             bool
	CutoffMomentRatiosExtracted bool
	AbsoluteScaleExtracted      bool
	JExtractedAsSymmetry        bool
	Verdict                     string
}

type JAsConsequenceAudit struct {
	ProposedLogic             string
	VacuumExtremumSelected    bool
	CandidateJCanBeTested     bool
	PhysicalJDerived          bool
	JInvarianceEquation       string
	KOAxiomsVerified          bool
	OppositeActionConstructed bool
	Verdict                   string
}

type CutoffMomentExtractionAudit struct {
	Constraint                       string
	LinearInCutoffCombinations       bool
	InfiniteCutoffSolutions          bool
	F0F2F4RatiosExtracted            bool
	HeatKernelSubtractionDerived     bool
	ScalarGaugeNormalizationDerived  bool
	DimensionlessObservableSpecified bool
	Verdict                          string
}

type FourOverPiTest struct {
	TargetInstanton              float64
	CandidateEquation            string
	STopCanEncodeFourOverPi      bool
	BGapAsInverseCouplingDerived bool
	NonPerturbativeSectorDerived bool
	ProducesInstantonLaw         bool
	Verdict                      string
}

type Firewalls struct {
	DoesNotTreatSTopAsFullAction      bool
	DoesNotSelectRBranch              bool
	DoesNotInventCutoffMoments        bool
	DoesNotInventPhysicalJ            bool
	DoesNotClaimHiggsPrediction       bool
	DoesNotClaimBGapInstanton         bool
	IntermediateBreakingSealPreserved bool
	FiniteCorePolluted                bool
	Verdict                           string
}

type Summary struct {
	Gate286Inherited            bool
	STopConstraintFormalized    bool
	VariationalEquationsDerived bool
	BranchSelected              bool
	CutoffMomentsExtracted      bool
	PhysicalJDerived            bool
	FourPiInstantonDerived      bool
	HiggsPredictionDerived      bool
	IntermediateSealGranted     bool
	FirewallPreserved           bool
	Status                      string
	DirectAnswer                string
	NextGate                    string
}

type Analysis struct {
	PreviousGate286 finitencginstantonaction.Analysis
	Inheritance     Gate286Inheritance
	Constraint      TopologicalActionConstraint
	MomentModel     SpectralMomentModel
	Variation       VariationalEquationAudit
	Rank            ConstraintRankAudit
	J               JAsConsequenceAudit
	Cutoff          CutoffMomentExtractionAudit
	FourPi          FourOverPiTest
	Firewalls       Firewalls
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	prev, err := finitencginstantonaction.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 286 predecessor: %w", err)
	}
	inh := inheritGate286(prev)
	constraint := formalizeSTopConstraint()
	moments := buildMomentModel()
	variation := auditVariation(moments)
	rank := auditConstraintRank(constraint, moments, variation)
	j := auditJAsConsequence(variation)
	cutoff := auditCutoffExtraction(constraint, rank)
	fourpi := auditFourOverPi(constraint, variation, cutoff)
	fw := auditFirewalls(constraint, variation, rank, j, cutoff, fourpi)
	summary := buildSummary(inh, constraint, variation, rank, j, cutoff, fourpi, fw)
	truth := "Gate 287 promotes S_top=8π² from a diagnostic into a proposed variational boundary constraint and derives the corresponding finite spectral moment equations. The proposal is structurally important, but one scalar action constraint plus free cutoff moments and a free absolute D_F scale cannot select the Gate-275 r_+/r_- branch, derive physical J, extract f0:f2:f4, or generate (4/π)/B_gap. S_top is a lawful boundary datum, not yet a complete dynamical principle."
	return Analysis{PreviousGate286: prev, Inheritance: inh, Constraint: constraint, MomentModel: moments, Variation: variation, Rank: rank, J: j, Cutoff: cutoff, FourPi: fourpi, Firewalls: fw, Summary: summary, TruthStatement: truth}, nil
}

func inheritGate286(prev finitencginstantonaction.Analysis) Gate286Inheritance {
	return Gate286Inheritance{
		NCGCalculusFormalized:    prev.Summary.NCGCalculusFormalized,
		InnerFluctuationBuilt:    prev.Summary.InnerFluctuationBuilt,
		NontrivialSaddleDerived:  prev.Summary.NontrivialSaddleDerived,
		InverseBGapActionDerived: prev.Summary.InverseBGapActionDerived,
		FourOverPiGenerated:      prev.Summary.FourOverPiGenerated,
		IntermediateSealGranted:  prev.Summary.IntermediateSealGranted,
		Verdict:                  StatusGate286Inherited,
	}
}

func formalizeSTopConstraint() TopologicalActionConstraint {
	return TopologicalActionConstraint{
		STop:                           sTop,
		ExactForm:                      "S_top = 8π²",
		SpectralActionFormula:          "S_total = F4 a0(D_F) + F2 a2(D_F) + F0 a4(D_F), where F4=f4Λ⁴, F2=f2Λ², F0=f0",
		ConstraintEquation:             "F4 a0 + F2 Tr(D_F²) + F0 Tr(D_F⁴) = 8π²",
		TreatedAsDerivedFiniteDatum:    true,
		TreatedAsCompleteDynamics:      false,
		RequiresCutoffMoments:          true,
		RequiresPhysicalSpectralTriple: true,
		Verdict:                        StatusSTopConstraintFormalized,
	}
}

func buildMomentModel() SpectralMomentModel {
	branches := []AmplitudeBranch{
		branch("upper_branch", "+", (3591+136*math.Sqrt(123))/3099),
		branch("lower_branch", "-", (3591-136*math.Sqrt(123))/3099),
	}
	return SpectralMomentModel{
		KappaC:             kappaC,
		KappaQ:             kappaQ,
		LambdaContact:      lambda,
		D2Formula:          "a2 proxy / Tr(D_F²) = X(1+3r), with X=|x|²",
		D4Formula:          "a4 proxy / Tr(D_F⁴) = X²(1+3r²)",
		ShapeEquation:      "(1+3r²)/(1+3r)² = 1197/4624",
		Quadratic:          "3099r² - 7182r + 3427 = 0",
		Branches:           branches,
		HasTwoBranches:     len(branches) == 2,
		AbsoluteScaleKnown: false,
		Verdict:            StatusMomentModelInherited,
	}
}

func branch(name, sign string, r float64) AmplitudeBranch {
	d2 := 1 + 3*r
	d4 := 1 + 3*r*r
	l := d4 / (d2 * d2)
	return AmplitudeBranch{Name: name, Sign: sign, R: r, AbsYOverX: math.Sqrt(r), ShapeLambda: l, ShapeResidualAbs: math.Abs(l - lambda), D2ForXEqualsOne: d2, D4ForXEqualsOne: d4}
}

func auditVariation(m SpectralMomentModel) VariationalEquationAudit {
	shapeAt1 := (1 + 3*1.0*1.0) / math.Pow(1+3*1.0, 2)
	return VariationalEquationAudit{
		Variables:                            []string{"r=|y/x|²", "X=|x|²", "F0=f0", "F2=f2Λ²", "F4=f4Λ⁴", "physical J", "vacuum orientation", "B_gap coupling map"},
		SFormula:                             "S(r,X)=F4 a0 + F2 X(1+3r) + F0 X²(1+3r²)",
		DerivativeWRT_R:                      "∂S/∂r = 3F2 X + 6F0 X² r",
		StationaryRFormula:                   "r_* = -F2/(2F0 X), if F0·X ≠ 0",
		PositiveCutoffMomentsSelectPositiveR: false,
		ArbitrarySignedMomentsCanFitAnyR:     true,
		ShapeDerivative:                      "d/dr[(1+3r²)/(1+3r)²] = 6(r-1)/(1+3r)³",
		ShapeExtremumR:                       1,
		ShapeExtremumLambda:                  shapeAt1,
		BranchesAreShapeExtrema:              false,
		SelectsUpperBranch:                   false,
		SelectsLowerBranch:                   false,
		UniqueBranchSelected:                 false,
		Verdict:                              StatusVariationalEquations,
	}
}

func auditConstraintRank(c TopologicalActionConstraint, m SpectralMomentModel, v VariationalEquationAudit) ConstraintRankAudit {
	unknowns := []string{"r", "X=|x|²", "F0", "F2", "F4", "a0 normalization", "physical J", "chiral/hypercharge representation", "B_gap insertion/coupling"}
	equations := []string{c.ConstraintEquation, m.ShapeEquation, v.DerivativeWRT_R + " = 0 (only after choosing admissible cutoff signs/scale)"}
	return ConstraintRankAudit{
		Unknowns:                    unknowns,
		NativeEquations:             equations,
		UnknownCount:                len(unknowns),
		EquationCount:               len(equations),
		Underdetermined:             true,
		CutoffMomentRatiosExtracted: false,
		AbsoluteScaleExtracted:      false,
		JExtractedAsSymmetry:        false,
		Verdict:                     StatusCutoffUnderdetermined,
	}
}

func auditJAsConsequence(v VariationalEquationAudit) JAsConsequenceAudit {
	return JAsConsequenceAudit{
		ProposedLogic:             "derive the physical J as the antiunitary symmetry preserving the selected variational extremum",
		VacuumExtremumSelected:    v.UniqueBranchSelected,
		CandidateJCanBeTested:     false,
		PhysicalJDerived:          false,
		JInvarianceEquation:       "J D_F(r_*) = D_F(r_*) J and Jγ=-γJ, after r_* and physical H_F are derived",
		KOAxiomsVerified:          false,
		OppositeActionConstructed: false,
		Verdict:                   StatusFailedPhysicalJNotDerived,
	}
}

func auditCutoffExtraction(c TopologicalActionConstraint, r ConstraintRankAudit) CutoffMomentExtractionAudit {
	return CutoffMomentExtractionAudit{
		Constraint:                       c.ConstraintEquation,
		LinearInCutoffCombinations:       true,
		InfiniteCutoffSolutions:          true,
		F0F2F4RatiosExtracted:            false,
		HeatKernelSubtractionDerived:     false,
		ScalarGaugeNormalizationDerived:  false,
		DimensionlessObservableSpecified: false,
		Verdict:                          StatusFailedCutoffRatiosNotExtracted,
	}
}

func auditFourOverPi(c TopologicalActionConstraint, v VariationalEquationAudit, hk CutoffMomentExtractionAudit) FourOverPiTest {
	return FourOverPiTest{
		TargetInstanton:              fourOverPi / bGap,
		CandidateEquation:            "S_inst ?= ΔS_finite at a non-trivial constrained extremum ?= (4/π)/B_gap",
		STopCanEncodeFourOverPi:      true,
		BGapAsInverseCouplingDerived: false,
		NonPerturbativeSectorDerived: false,
		ProducesInstantonLaw:         false,
		Verdict:                      StatusFailedFourPiInstNotDerived,
	}
}

func auditFirewalls(c TopologicalActionConstraint, v VariationalEquationAudit, r ConstraintRankAudit, j JAsConsequenceAudit, hk CutoffMomentExtractionAudit, fp FourOverPiTest) Firewalls {
	return Firewalls{
		DoesNotTreatSTopAsFullAction:      !c.TreatedAsCompleteDynamics,
		DoesNotSelectRBranch:              !v.UniqueBranchSelected,
		DoesNotInventCutoffMoments:        !hk.F0F2F4RatiosExtracted,
		DoesNotInventPhysicalJ:            !j.PhysicalJDerived,
		DoesNotClaimHiggsPrediction:       true,
		DoesNotClaimBGapInstanton:         !fp.ProducesInstantonLaw,
		IntermediateBreakingSealPreserved: true,
		FiniteCorePolluted:                false,
		Verdict:                           StatusFirewallsPreserved,
	}
}

func buildSummary(inh Gate286Inheritance, c TopologicalActionConstraint, v VariationalEquationAudit, r ConstraintRankAudit, j JAsConsequenceAudit, hk CutoffMomentExtractionAudit, fp FourOverPiTest, fw Firewalls) Summary {
	statuses := []string{StatusGate286Inherited, StatusSTopConstraintFormalized, StatusMomentModelInherited, StatusVariationalEquations, StatusShapeExtremumAudited, StatusCutoffUnderdetermined, StatusPathBCConvergenceRecorded, StatusFirewallsPreserved, StatusFailedRBranchNotSelected, StatusFailedFreeCutoffMoments, StatusFailedPhysicalJNotDerived, StatusFailedCutoffRatiosNotExtracted, StatusFailedFourPiInstNotDerived, StatusFailedHiggsBGapFirewalled}
	return Summary{
		Gate286Inherited:            inh.NCGCalculusFormalized,
		STopConstraintFormalized:    c.TreatedAsDerivedFiniteDatum,
		VariationalEquationsDerived: true,
		BranchSelected:              v.UniqueBranchSelected,
		CutoffMomentsExtracted:      hk.F0F2F4RatiosExtracted,
		PhysicalJDerived:            j.PhysicalJDerived,
		FourPiInstantonDerived:      fp.ProducesInstantonLaw,
		HiggsPredictionDerived:      false,
		IntermediateSealGranted:     false,
		FirewallPreserved:           !fw.FiniteCorePolluted && fw.DoesNotTreatSTopAsFullAction && fw.DoesNotInventCutoffMoments && fw.DoesNotInventPhysicalJ && fw.DoesNotClaimBGapInstanton,
		Status:                      strings.Join(statuses, ";"),
		DirectAnswer:                "S_top=8π² can be formalized as a global finite spectral-action boundary constraint, but it does not by itself select r_+ or r_-, derive J, extract cutoff-moment ratios, or produce (4/π)/B_gap.",
		NextGate:                    "Derive a physical finite spectral triple and admissible cutoff-moment/normalization scheme, or formally introduce a TopologicalActionMomentSeal before using S_top as a dynamical selector.",
	}
}

func FormatInheritance(g Gate286Inheritance) string {
	return fmt.Sprintf("calculus=%t inner=%t saddle=%t inverseBgap=%t fourPi=%t seal=%t verdict=%s", g.NCGCalculusFormalized, g.InnerFluctuationBuilt, g.NontrivialSaddleDerived, g.InverseBGapActionDerived, g.FourOverPiGenerated, g.IntermediateSealGranted, g.Verdict)
}
func FormatConstraint(c TopologicalActionConstraint) string {
	return fmt.Sprintf("S_top=%.12g exact=%q action=%q constraint=%q finiteDatum=%t completeDynamics=%t needsMoments=%t needsTriple=%t verdict=%s", c.STop, c.ExactForm, c.SpectralActionFormula, c.ConstraintEquation, c.TreatedAsDerivedFiniteDatum, c.TreatedAsCompleteDynamics, c.RequiresCutoffMoments, c.RequiresPhysicalSpectralTriple, c.Verdict)
}
func FormatMomentModel(m SpectralMomentModel) string {
	parts := []string{}
	for _, b := range m.Branches {
		parts = append(parts, fmt.Sprintf("%s r=%.12g |y/x|=%.12g λ=%.12g residual=%.3g", b.Name, b.R, b.AbsYOverX, b.ShapeLambda, b.ShapeResidualAbs))
	}
	return fmt.Sprintf("κ=%d:%d λ=%.12g D2=%q D4=%q shape=%q quadratic=%q branches=[%s] absScale=%t verdict=%s", m.KappaC, m.KappaQ, m.LambdaContact, m.D2Formula, m.D4Formula, m.ShapeEquation, m.Quadratic, strings.Join(parts, "; "), m.AbsoluteScaleKnown, m.Verdict)
}
func FormatVariation(v VariationalEquationAudit) string {
	return fmt.Sprintf("S=%q dSdr=%q stationary=%q positiveMomentsPositiveR=%t arbitrarySignedFit=%t shapeDerivative=%q shapeExtremumR=%.12g shapeExtremumλ=%.12g branchesExtrema=%t selected=(upper:%t lower:%t unique:%t) verdict=%s", v.SFormula, v.DerivativeWRT_R, v.StationaryRFormula, v.PositiveCutoffMomentsSelectPositiveR, v.ArbitrarySignedMomentsCanFitAnyR, v.ShapeDerivative, v.ShapeExtremumR, v.ShapeExtremumLambda, v.BranchesAreShapeExtrema, v.SelectsUpperBranch, v.SelectsLowerBranch, v.UniqueBranchSelected, v.Verdict)
}
func FormatRank(r ConstraintRankAudit) string {
	return fmt.Sprintf("unknowns=%d equations=%d underdetermined=%t cutoffRatios=%t scale=%t J=%t nativeEquations=%v verdict=%s", r.UnknownCount, r.EquationCount, r.Underdetermined, r.CutoffMomentRatiosExtracted, r.AbsoluteScaleExtracted, r.JExtractedAsSymmetry, r.NativeEquations, r.Verdict)
}
func FormatJ(j JAsConsequenceAudit) string {
	return fmt.Sprintf("logic=%q extremum=%t candidateTest=%t physicalJ=%t equation=%q KO=%t opposite=%t verdict=%s", j.ProposedLogic, j.VacuumExtremumSelected, j.CandidateJCanBeTested, j.PhysicalJDerived, j.JInvarianceEquation, j.KOAxiomsVerified, j.OppositeActionConstructed, j.Verdict)
}
func FormatCutoff(c CutoffMomentExtractionAudit) string {
	return fmt.Sprintf("constraint=%q linear=%t infiniteSolutions=%t ratios=%t subtraction=%t fieldNorm=%t observable=%t verdict=%s", c.Constraint, c.LinearInCutoffCombinations, c.InfiniteCutoffSolutions, c.F0F2F4RatiosExtracted, c.HeatKernelSubtractionDerived, c.ScalarGaugeNormalizationDerived, c.DimensionlessObservableSpecified, c.Verdict)
}
func FormatFourPi(f FourOverPiTest) string {
	return fmt.Sprintf("target=%.12g equation=%q STopCanEncode=%t inverseBgap=%t nonpert=%t produces=%t verdict=%s", f.TargetInstanton, f.CandidateEquation, f.STopCanEncodeFourOverPi, f.BGapAsInverseCouplingDerived, f.NonPerturbativeSectorDerived, f.ProducesInstantonLaw, f.Verdict)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("noSTopFullAction=%t noRSelect=%t noCutoffInvent=%t noJInvent=%t noHiggs=%t noBgapInst=%t seal=%t polluted=%t verdict=%s", f.DoesNotTreatSTopAsFullAction, f.DoesNotSelectRBranch, f.DoesNotInventCutoffMoments, f.DoesNotInventPhysicalJ, f.DoesNotClaimHiggsPrediction, f.DoesNotClaimBGapInstanton, f.IntermediateBreakingSealPreserved, f.FiniteCorePolluted, f.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("gate286=%t STop=%t variation=%t branch=%t cutoff=%t J=%t fourPi=%t higgs=%t intermediateSeal=%t firewall=%t next=%q status=%s", s.Gate286Inherited, s.STopConstraintFormalized, s.VariationalEquationsDerived, s.BranchSelected, s.CutoffMomentsExtracted, s.PhysicalJDerived, s.FourPiInstantonDerived, s.HiggsPredictionDerived, s.IntermediateSealGranted, s.FirewallPreserved, s.NextGate, s.Status)
}
