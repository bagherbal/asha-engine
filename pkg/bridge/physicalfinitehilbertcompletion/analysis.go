// Package physicalfinitehilbertcompletion implements Gate 275:
// Physical Finite Hilbert Space / Chiral Hypercharge Opposite-Action Completion Audit.
//
// Gate 274 extracted exact local quaternionic closure on a selected weak
// doublet, but kept the Standard-Model finite Hilbert space, physical real
// structure J, hypercharge attachment, and Dirac edge amplitudes un-derived.
// Gate 275 adds a preliminary scalar-Morita bridge suggested by the previous
// audits: the Gate-169 contact/Higgs scalar shape λ=1197/4624 can be compared
// to the Gate-273 Morita trace multiplicities κ_C:κ_Q=1:3.  This yields an
// exact quadratic for r=|y/x|² and therefore a two-branch candidate amplitude
// ratio.
//
// The gate is deliberately strict.  The scalar-Morita shape bridge is recorded
// as a conditional cross-tower constraint, not as a completed spectral-action
// theorem.  It fixes r only up to a two-fold ambiguity and it does not supply
// the heat-kernel normalization, physical J, full hypercharge assignment, or
// Seeley-de Witt projection required to claim a Higgs mass prediction.
package physicalfinitehilbertcompletion

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/nativeweakquaternionicalgebra"
)

const (
	AuditID = "GATE275-PHYSICAL-FINITE-HILBERT-SPACE-CHIRAL-HYPERCHARGE-OPPOSITE-ACTION-COMPLETION-AUDIT"

	StatusGate274Inherited         = "CONDITIONAL_SUPPORT_GATE274_LOCAL_QUATERNIONIC_LEDGER_INHERITED"
	StatusScalarShapeRetrieved     = "CONDITIONAL_SUPPORT_GATE169_SCALAR_SHAPE_RETRIEVED"
	StatusMoritaMultiplicity       = "CONDITIONAL_SUPPORT_GATE273_MORITA_MULTIPLICITY_RETRIEVED"
	StatusScalarMoritaBridgeSolved = "CONDITIONAL_SUPPORT_SCALAR_MORITA_AMPLITUDE_SHAPE_BRIDGE_SOLVED"
	StatusTwoBranchXY              = "CONDITIONAL_SUPPORT_TWO_BRANCH_XY_RATIO_CONSTRAINED"
	StatusCandidateMomentsComputed = "CONDITIONAL_SUPPORT_CANDIDATE_SPECTRAL_MOMENTS_COMPUTED"
	StatusPhysicalJAudited         = "CONDITIONAL_SUPPORT_PHYSICAL_J_AND_HYPERCHARGE_AUDITED"
	StatusFailedBridgeConditional  = "FAILED_ROUTE_SCALAR_MORITA_IDENTIFICATION_REQUIRES_BRIDGE_THEOREM"
	StatusFailedTwoBranch          = "FAILED_ROUTE_TWO_BRANCH_XY_AMBIGUITY_REMAINS"
	StatusFailedPhysicalJ          = "FAILED_ROUTE_PHYSICAL_CHARGE_CONJUGATION_J_NOT_DERIVED"
	StatusFailedHypercharge        = "FAILED_ROUTE_FULL_CHIRAL_HYPERCHARGE_ASSIGNMENT_NOT_DERIVED"
	StatusFailedOppositeAction     = "FAILED_ROUTE_FULL_C_PLUS_H_PLUS_M3C_OPPOSITE_ACTION_NOT_DERIVED"
	StatusFailedA2A4               = "FAILED_ROUTE_SEELEY_DE_WITT_A2_A4_STILL_NOT_DERIVED"
	StatusFailedHiggs              = "FAILED_ROUTE_HIGGS_MASS_RATIO_NOT_CLAIMED"
)

type Gate274Inheritance struct {
	LocalHExtracted       bool
	CandidateAlgebraBuilt bool
	ExactSMAlgebraDerived bool
	PhysicalHFDerived     bool
	PhysicalJDerived      bool
	XYRatioLocked         bool
	A2A4Derived           bool
	HiggsRatioDerived     bool
	FirewallPreserved     bool
	Verdict               string
}

type ScalarShapeLedger struct {
	SourceGate             int
	ExactNumerator         int
	ExactDenominator       int
	FloatValue             float64
	Formula                string
	DerivedFiniteCoreDatum bool
	EmpiricalInput         bool
	RequiresCrossTowerMap  bool
	Verdict                string
}

type MoritaMultiplicityLedger struct {
	SourceGate          int
	KappaC              int
	KappaQ              int
	Ratio               string
	TraceD2Formula      string
	TraceD4Formula      string
	MultiplicityDerived bool
	AmplitudeDerived    bool
	Verdict             string
}

type QuadraticLedger struct {
	Variable             string
	Equation             string
	A                    int
	B                    int
	C                    int
	Discriminant         int
	DiscriminantSqrtForm string
	HasTwoPositiveRoots  bool
	ExactRootForm        string
	Verdict              string
}

type AmplitudeBranch struct {
	Name                  string
	R                     float64
	ExactRForm            string
	AbsYOverX             float64
	TraceD2X1             float64
	TraceD4X1             float64
	ShapeLambda           float64
	ShapeResidualAbs      float64
	D4OverD2X1            float64
	D2OverD4X1            float64
	CandidateHiggsQuartic string
}

type ScalarMoritaBridgeAudit struct {
	Shape                 ScalarShapeLedger
	Multiplicity          MoritaMultiplicityLedger
	Quadratic             QuadraticLedger
	Branches              []AmplitudeBranch
	BridgeEquation        string
	RootsConstrainR       bool
	UniqueBranchSelected  bool
	AbsoluteScaleSelected bool
	HeatKernelReady       bool
	A2A4Derived           bool
	HiggsRatioDerived     bool
	Verdict               string
}

type PhysicalJAudit struct {
	CandidateName             string
	CandidateJ2               int
	CandidateKOCompatible     bool
	OccupationComplementSeen  bool
	AntiLinearImplemented     bool
	ParticleAntiparticleTyped bool
	PhysicalHFCompleted       bool
	Verdict                   string
}

type HyperchargeAudit struct {
	BMinusLLedgerAvailable         bool
	T3LedgerAvailable              bool
	CandidateHyperchargeKnown      bool
	FullCPlusHPlusM3Representation bool
	ChiralAssignmentDerived        bool
	AnomalyCancellationRechecked   bool
	EmpiricalAssignmentsInserted   bool
	Verdict                        string
}

type OppositeActionAudit struct {
	MoritaOppositeActionInherited bool
	LocalHIncluded                bool
	PhysicalJDerived              bool
	FullOppositeActionDerived     bool
	OrderOneReevaluatedOnFullAF   bool
	NonVacuousOneFormsCertified   bool
	XYRatioBranchSelectedByJ      bool
	Verdict                       string
}

type FirewallAudit struct {
	NoObservedMassInserted              bool
	NoVEVInserted                       bool
	NoCKMPMNSInserted                   bool
	NoEmpiricalYukawaAmplitudeInserted  bool
	ScalarShapeKeptFinite               bool
	ScalarMoritaBridgeMarkedConditional bool
	CandidateMomentsNotHiggsPrediction  bool
	EmpiricalYukawaSealPreserved        bool
	FiniteCorePolluted                  bool
	Verdict                             string
}

type FutureCriterion struct {
	Name      string
	Required  bool
	Satisfied bool
	Detail    string
}

type FutureMap struct {
	Criteria                   []FutureCriterion
	NeedScalarMoritaMapTheorem bool
	NeedBranchSelector         bool
	NeedPhysicalJ              bool
	NeedHyperchargeCompletion  bool
	NeedHeatKernelProjection   bool
	RecommendedNextGate        string
	Verdict                    string
}

type Summary struct {
	Gate274Inherited        bool
	ScalarShapeRetrieved    bool
	MoritaMultiplicityKnown bool
	ScalarMoritaSolved      bool
	TwoBranchXYConstrained  bool
	UniqueXYLocked          bool
	PhysicalJDerived        bool
	HyperchargeDerived      bool
	OppositeActionDerived   bool
	A2A4Derived             bool
	HiggsRatioClaimed       bool
	FirewallPreserved       bool
	Status                  string
	NextGate                string
	Comment                 string
}

type Analysis struct {
	PreviousGate274 nativeweakquaternionicalgebra.Analysis
	Inheritance     Gate274Inheritance
	Bridge          ScalarMoritaBridgeAudit
	J               PhysicalJAudit
	Hypercharge     HyperchargeAudit
	OppositeAction  OppositeActionAudit
	Firewall        FirewallAudit
	Future          FutureMap
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
		prev, err := nativeweakquaternionicalgebra.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 274 predecessor: %w", err)
			return
		}
		inh := inheritGate274(prev)
		bridge := solveScalarMoritaBridge()
		j := auditPhysicalJ()
		y := auditHypercharge()
		op := auditOppositeAction(j, y)
		fw := auditFirewall(bridge, j, y, op)
		future := defineFuture(bridge, j, y, op)
		summary := summarize(inh, bridge, j, y, op, fw)
		truth := buildTruth(bridge, j, y, op)
		defaultA = Analysis{PreviousGate274: prev, Inheritance: inh, Bridge: bridge, J: j, Hypercharge: y, OppositeAction: op, Firewall: fw, Future: future, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate274(prev nativeweakquaternionicalgebra.Analysis) Gate274Inheritance {
	return Gate274Inheritance{
		LocalHExtracted:       prev.Summary.LocalHExtracted,
		CandidateAlgebraBuilt: prev.Summary.CandidateAlgebraBuilt,
		ExactSMAlgebraDerived: prev.Summary.ExactSMAlgebraDerived,
		PhysicalHFDerived:     prev.Summary.PhysicalHFDerived,
		PhysicalJDerived:      prev.Summary.PhysicalJDerived,
		XYRatioLocked:         prev.Summary.XYRatioLocked,
		A2A4Derived:           prev.Summary.A2A4Derived,
		HiggsRatioDerived:     prev.Summary.HiggsRatioDerived,
		FirewallPreserved:     prev.Summary.FirewallPreserved,
		Verdict:               StatusGate274Inherited + "; local H exists under selector, but physical H_F/J and amplitudes remain open",
	}
}

func solveScalarMoritaBridge() ScalarMoritaBridgeAudit {
	shape := ScalarShapeLedger{
		SourceGate:             169,
		ExactNumerator:         1197,
		ExactDenominator:       4624,
		FloatValue:             1197.0 / 4624.0,
		Formula:                "λ_contact = Tr(M_K²)/Tr(M_K)² = 1197/4624",
		DerivedFiniteCoreDatum: true,
		EmpiricalInput:         false,
		RequiresCrossTowerMap:  true,
		Verdict:                StatusScalarShapeRetrieved + "; finite scalar/contact shape available as exact rational target",
	}
	mult := MoritaMultiplicityLedger{
		SourceGate:          273,
		KappaC:              1,
		KappaQ:              3,
		Ratio:               "κ_C:κ_Q = 1:3",
		TraceD2Formula:      "Tr(D_F²) proxy = |x|² + 3|y|²",
		TraceD4Formula:      "Tr(D_F⁴) proxy = |x|⁴ + 3|y|⁴",
		MultiplicityDerived: true,
		AmplitudeDerived:    false,
		Verdict:             StatusMoritaMultiplicity + "; trace multiplicities are finite-derived while edge amplitudes are not",
	}
	q := QuadraticLedger{
		Variable:             "r = |y/x|²",
		Equation:             "(1+3r²)/(1+3r)² = 1197/4624",
		A:                    3099,
		B:                    -7182,
		C:                    3427,
		Discriminant:         9100032,
		DiscriminantSqrtForm: "272√123",
		HasTwoPositiveRoots:  true,
		ExactRootForm:        "r = (3591 ± 136√123)/3099",
		Verdict:              StatusScalarMoritaBridgeSolved + "; exact quadratic has two positive amplitude-shape branches",
	}
	roots := []struct {
		name string
		sign string
		r    float64
	}{
		{"upper_branch", "+", (3591.0 + 136.0*math.Sqrt(123)) / 3099.0},
		{"lower_branch", "-", (3591.0 - 136.0*math.Sqrt(123)) / 3099.0},
	}
	branches := make([]AmplitudeBranch, 0, len(roots))
	for _, root := range roots {
		r := root.r
		d2 := 1 + 3*r
		d4 := 1 + 3*r*r
		lam := d4 / (d2 * d2)
		branches = append(branches, AmplitudeBranch{
			Name:                  root.name,
			R:                     r,
			ExactRForm:            fmt.Sprintf("(3591 %s 136√123)/3099", root.sign),
			AbsYOverX:             math.Sqrt(r),
			TraceD2X1:             d2,
			TraceD4X1:             d4,
			ShapeLambda:           lam,
			ShapeResidualAbs:      math.Abs(lam - shape.FloatValue),
			D4OverD2X1:            d4 / d2,
			D2OverD4X1:            d2 / d4,
			CandidateHiggsQuartic: "candidate finite moment only; not a Higgs quartic until heat-kernel and normalization map are derived",
		})
	}
	return ScalarMoritaBridgeAudit{
		Shape:                 shape,
		Multiplicity:          mult,
		Quadratic:             q,
		Branches:              branches,
		BridgeEquation:        "(|x|⁴+3|y|⁴)/(|x|²+3|y|²)² = 1197/4624",
		RootsConstrainR:       true,
		UniqueBranchSelected:  false,
		AbsoluteScaleSelected: false,
		HeatKernelReady:       false,
		A2A4Derived:           false,
		HiggsRatioDerived:     false,
		Verdict:               StatusTwoBranchXY + "; scalar shape constrains |y/x|² to two branches, but does not select branch or absolute scale",
	}
}

func auditPhysicalJ() PhysicalJAudit {
	return PhysicalJAudit{
		CandidateName:             "occupation-complement plus complex conjugation candidate",
		CandidateJ2:               1,
		CandidateKOCompatible:     true,
		OccupationComplementSeen:  true,
		AntiLinearImplemented:     false,
		ParticleAntiparticleTyped: false,
		PhysicalHFCompleted:       false,
		Verdict:                   StatusPhysicalJAudited + "; preflight candidate exists, but physical anti-linear charge conjugation on H_F is not derived",
	}
}

func auditHypercharge() HyperchargeAudit {
	return HyperchargeAudit{
		BMinusLLedgerAvailable:         true,
		T3LedgerAvailable:              true,
		CandidateHyperchargeKnown:      true,
		FullCPlusHPlusM3Representation: false,
		ChiralAssignmentDerived:        false,
		AnomalyCancellationRechecked:   false,
		EmpiricalAssignmentsInserted:   false,
		Verdict:                        "CONDITIONAL_SUPPORT_NATIVE_CHARGE_LEDGERS_AVAILABLE; full chiral hypercharge assignment on C⊕H⊕M3(C) H_F remains un-derived",
	}
}

func auditOppositeAction(j PhysicalJAudit, y HyperchargeAudit) OppositeActionAudit {
	return OppositeActionAudit{
		MoritaOppositeActionInherited: true,
		LocalHIncluded:                true,
		PhysicalJDerived:              j.ParticleAntiparticleTyped && j.AntiLinearImplemented,
		FullOppositeActionDerived:     false,
		OrderOneReevaluatedOnFullAF:   false,
		NonVacuousOneFormsCertified:   false,
		XYRatioBranchSelectedByJ:      false,
		Verdict:                       StatusFailedOppositeAction + "; Morita opposite action exists abstractly, but physical J and full hypercharge representation are incomplete",
	}
}

func auditFirewall(b ScalarMoritaBridgeAudit, j PhysicalJAudit, y HyperchargeAudit, op OppositeActionAudit) FirewallAudit {
	return FirewallAudit{
		NoObservedMassInserted:              true,
		NoVEVInserted:                       true,
		NoCKMPMNSInserted:                   true,
		NoEmpiricalYukawaAmplitudeInserted:  true,
		ScalarShapeKeptFinite:               b.Shape.DerivedFiniteCoreDatum && !b.Shape.EmpiricalInput,
		ScalarMoritaBridgeMarkedConditional: b.Shape.RequiresCrossTowerMap && !b.A2A4Derived,
		CandidateMomentsNotHiggsPrediction:  !b.HiggsRatioDerived,
		EmpiricalYukawaSealPreserved:        true,
		FiniteCorePolluted:                  false,
		Verdict:                             "CONDITIONAL_SUPPORT_FIREWALLS_PRESERVED; candidate spectral moments are not promoted to physical Higgs prediction",
	}
}

func defineFuture(b ScalarMoritaBridgeAudit, j PhysicalJAudit, y HyperchargeAudit, op OppositeActionAudit) FutureMap {
	criteria := []FutureCriterion{
		{Name: "scalar-Morita bridge theorem", Required: true, Satisfied: false, Detail: "prove Gate-169 contact scalar shape and Morita D_F edge moments are the same spectral-action object"},
		{Name: "two-branch selector", Required: true, Satisfied: false, Detail: "select upper or lower r branch without empirical mass input"},
		{Name: "absolute finite Dirac scale", Required: true, Satisfied: false, Detail: "derive |x| or a scale-free normalization convention compatible with heat-kernel coefficients"},
		{Name: "physical anti-linear J", Required: true, Satisfied: false, Detail: "construct charge conjugation on physical H_F with particle/antiparticle semantics"},
		{Name: "full chiral hypercharge assignment", Required: true, Satisfied: false, Detail: "represent C⊕H⊕M3(C) with left doublets, right singlets, and anomaly-safe charge ledger"},
		{Name: "opposite action and order-one recheck", Required: true, Satisfied: false, Detail: "evaluate [[D,a],Ja*J^-1] on completed H_F"},
		{Name: "Seeley-de Witt projection", Required: true, Satisfied: false, Detail: "derive a2/a4 normalization, subtraction scheme, and scalar/gauge projection"},
	}
	return FutureMap{
		Criteria:                   criteria,
		NeedScalarMoritaMapTheorem: true,
		NeedBranchSelector:         !b.UniqueBranchSelected,
		NeedPhysicalJ:              !j.ParticleAntiparticleTyped,
		NeedHyperchargeCompletion:  !y.ChiralAssignmentDerived,
		NeedHeatKernelProjection:   !b.HeatKernelReady,
		RecommendedNextGate:        "Gate 276 — Scalar-Morita Spectral Shape Bridge / Branch Selector and Heat-Kernel Normalization Audit",
		Verdict:                    "candidate r branches are now available; future theorem must select branch and complete physical spectral triple before Higgs ratio",
	}
}

func summarize(inh Gate274Inheritance, b ScalarMoritaBridgeAudit, j PhysicalJAudit, y HyperchargeAudit, op OppositeActionAudit, fw FirewallAudit) Summary {
	return Summary{
		Gate274Inherited:        inh.LocalHExtracted && inh.CandidateAlgebraBuilt,
		ScalarShapeRetrieved:    b.Shape.DerivedFiniteCoreDatum,
		MoritaMultiplicityKnown: b.Multiplicity.MultiplicityDerived,
		ScalarMoritaSolved:      b.RootsConstrainR && len(b.Branches) == 2,
		TwoBranchXYConstrained:  b.RootsConstrainR && !b.UniqueBranchSelected,
		UniqueXYLocked:          b.UniqueBranchSelected,
		PhysicalJDerived:        j.ParticleAntiparticleTyped,
		HyperchargeDerived:      y.ChiralAssignmentDerived,
		OppositeActionDerived:   op.FullOppositeActionDerived,
		A2A4Derived:             b.A2A4Derived,
		HiggsRatioClaimed:       b.HiggsRatioDerived,
		FirewallPreserved:       !fw.FiniteCorePolluted && fw.EmpiricalYukawaSealPreserved,
		Status:                  "BRIDGE_REQUIRED_WITH_TWO_BRANCH_SCALAR_MORITA_CONSTRAINT",
		NextGate:                "Gate 276 — Scalar-Morita Spectral Shape Bridge / Branch Selector and Heat-Kernel Normalization Audit",
		Comment:                 "Gate 275 constrains |y/x|² to two finite algebraic branches using Gate 169 and Gate 273, but physical J, hypercharge, branch selection, and a2/a4 remain open.",
	}
}

func buildTruth(b ScalarMoritaBridgeAudit, j PhysicalJAudit, y HyperchargeAudit, op OppositeActionAudit) string {
	return "Gate 275 connects the finite contact scalar shape λ=1197/4624 to the Morita multiplicity trace formula κ_C:κ_Q=1:3 and solves the exact amplitude-shape equation for r=|y/x|². This produces two candidate pure-algebra branches, not a unique Higgs prediction. Physical charge conjugation J, full hypercharge/chirality assignment, opposite action, branch selection, and Seeley-de Witt normalization remain required."
}

func FormatInheritance(a Gate274Inheritance) string {
	return fmt.Sprintf("localH=%t candidateAlg=%t exactSM=%t HF=%t J=%t xy=%t a2a4=%t higgs=%t firewall=%t verdict=%s", a.LocalHExtracted, a.CandidateAlgebraBuilt, a.ExactSMAlgebraDerived, a.PhysicalHFDerived, a.PhysicalJDerived, a.XYRatioLocked, a.A2A4Derived, a.HiggsRatioDerived, a.FirewallPreserved, a.Verdict)
}

func FormatScalarShape(a ScalarShapeLedger) string {
	return fmt.Sprintf("gate=%d formula=%q exact=%d/%d value=%.15g finite=%t empirical=%t bridgeRequired=%t verdict=%s", a.SourceGate, a.Formula, a.ExactNumerator, a.ExactDenominator, a.FloatValue, a.DerivedFiniteCoreDatum, a.EmpiricalInput, a.RequiresCrossTowerMap, a.Verdict)
}

func FormatMultiplicity(a MoritaMultiplicityLedger) string {
	return fmt.Sprintf("gate=%d ratio=%s D2=%q D4=%q multiplicity=%t amplitude=%t verdict=%s", a.SourceGate, a.Ratio, a.TraceD2Formula, a.TraceD4Formula, a.MultiplicityDerived, a.AmplitudeDerived, a.Verdict)
}

func FormatQuadratic(a QuadraticLedger) string {
	return fmt.Sprintf("var=%q eq=%q coeffs=(%d,%d,%d) disc=%d sqrt=%s roots=%s positive=%t verdict=%s", a.Variable, a.Equation, a.A, a.B, a.C, a.Discriminant, a.DiscriminantSqrtForm, a.ExactRootForm, a.HasTwoPositiveRoots, a.Verdict)
}

func FormatBranch(a AmplitudeBranch) string {
	return fmt.Sprintf("%s r=%s≈%.15g |y/x|≈%.15g D2(x=1)=%.15g D4(x=1)=%.15g λ≈%.15g residual=%.3g D4/D2≈%.15g D2/D4≈%.15g note=%q", a.Name, a.ExactRForm, a.R, a.AbsYOverX, a.TraceD2X1, a.TraceD4X1, a.ShapeLambda, a.ShapeResidualAbs, a.D4OverD2X1, a.D2OverD4X1, a.CandidateHiggsQuartic)
}

func FormatBridge(a ScalarMoritaBridgeAudit) string {
	branches := []string{}
	for _, b := range a.Branches {
		branches = append(branches, FormatBranch(b))
	}
	return fmt.Sprintf("shape={%s} multiplicity={%s} quadratic={%s} bridge=%q roots=%t unique=%t scale=%t heat=%t a2a4=%t higgs=%t branches=[%s] verdict=%s", FormatScalarShape(a.Shape), FormatMultiplicity(a.Multiplicity), FormatQuadratic(a.Quadratic), a.BridgeEquation, a.RootsConstrainR, a.UniqueBranchSelected, a.AbsoluteScaleSelected, a.HeatKernelReady, a.A2A4Derived, a.HiggsRatioDerived, strings.Join(branches, "; "), a.Verdict)
}

func FormatJ(a PhysicalJAudit) string {
	return fmt.Sprintf("candidate=%q J2=%d KO=%t complement=%t antiLinear=%t particleTyped=%t HF=%t verdict=%s", a.CandidateName, a.CandidateJ2, a.CandidateKOCompatible, a.OccupationComplementSeen, a.AntiLinearImplemented, a.ParticleAntiparticleTyped, a.PhysicalHFCompleted, a.Verdict)
}

func FormatHypercharge(a HyperchargeAudit) string {
	return fmt.Sprintf("BL=%t T3=%t Ycandidate=%t fullRep=%t chiral=%t anomaly=%t empirical=%t verdict=%s", a.BMinusLLedgerAvailable, a.T3LedgerAvailable, a.CandidateHyperchargeKnown, a.FullCPlusHPlusM3Representation, a.ChiralAssignmentDerived, a.AnomalyCancellationRechecked, a.EmpiricalAssignmentsInserted, a.Verdict)
}

func FormatOpposite(a OppositeActionAudit) string {
	return fmt.Sprintf("moritaInherited=%t localH=%t physicalJ=%t fullOpposite=%t orderOne=%t oneForms=%t branchByJ=%t verdict=%s", a.MoritaOppositeActionInherited, a.LocalHIncluded, a.PhysicalJDerived, a.FullOppositeActionDerived, a.OrderOneReevaluatedOnFullAF, a.NonVacuousOneFormsCertified, a.XYRatioBranchSelectedByJ, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("noMass=%t noVEV=%t noCKM=%t noYukawaAmp=%t scalarFinite=%t bridgeConditional=%t momentsNotHiggs=%t empiricalSeal=%t polluted=%t verdict=%s", a.NoObservedMassInserted, a.NoVEVInserted, a.NoCKMPMNSInserted, a.NoEmpiricalYukawaAmplitudeInserted, a.ScalarShapeKeptFinite, a.ScalarMoritaBridgeMarkedConditional, a.CandidateMomentsNotHiggsPrediction, a.EmpiricalYukawaSealPreserved, a.FiniteCorePolluted, a.Verdict)
}

func FormatFuture(a FutureMap) string {
	missing := []string{}
	for _, c := range a.Criteria {
		if c.Required && !c.Satisfied {
			missing = append(missing, c.Name+": "+c.Detail)
		}
	}
	return fmt.Sprintf("criteria=%d missing=[%s] scalarBridge=%t branch=%t J=%t hypercharge=%t heat=%t next=%q verdict=%q", len(a.Criteria), strings.Join(missing, "; "), a.NeedScalarMoritaMapTheorem, a.NeedBranchSelector, a.NeedPhysicalJ, a.NeedHyperchargeCompletion, a.NeedHeatKernelProjection, a.RecommendedNextGate, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("inherit=%t scalar=%t multiplicity=%t solved=%t twoBranch=%t unique=%t J=%t hypercharge=%t opposite=%t a2a4=%t higgs=%t firewall=%t status=%s next=%q comment=%q", a.Gate274Inherited, a.ScalarShapeRetrieved, a.MoritaMultiplicityKnown, a.ScalarMoritaSolved, a.TwoBranchXYConstrained, a.UniqueXYLocked, a.PhysicalJDerived, a.HyperchargeDerived, a.OppositeActionDerived, a.A2A4Derived, a.HiggsRatioClaimed, a.FirewallPreserved, a.Status, a.NextGate, a.Comment)
}
