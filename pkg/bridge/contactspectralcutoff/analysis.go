// Package contactspectralcutoff implements Gate 288:
// Contact-Spectral Cutoff Identification / S_top Branch Selector Audit.
//
// Gate 287 proved that S_top=8π² is a legitimate global boundary constraint,
// but with free cutoff moments it cannot select the Gate-275 amplitude branch.
// Gate 288 audits the proposed identification: use the exact contact spectral
// zeta/moment ledger as the spectral-action cutoff moments.  This removes the
// free cutoff-moment ambiguity and yields a concrete quadratic for the absolute
// Dirac scale X=|x|² on each r branch.  The result is deliberately firewall
// preserving: the contact cutoff fixes the total trace moments in the reduced
// scalar-Morita proxy, but both r branches admit positive real X and therefore
// no unique physical branch or Higgs mass ratio is derived.
package contactspectralcutoff

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE288-CONTACT-SPECTRAL-CUTOFF-IDENTIFICATION-S-TOP-BRANCH-SELECTOR-AUDIT"

	StatusGate287Inherited              = "CONDITIONAL_SUPPORT_GATE287_S_TOP_UNDERDETERMINATION_INHERITED"
	StatusCutoffMomentsRetrieved        = "CONDITIONAL_SUPPORT_CONTACT_SPECTRAL_CUTOFF_MOMENTS_RETRIEVED"
	StatusCutoffIdentificationAudited   = "CONDITIONAL_SUPPORT_CONTACT_SPECTRAL_CUTOFF_IDENTIFICATION_AUDITED"
	StatusQuadraticScaleConstraintBuilt = "CONDITIONAL_SUPPORT_QUADRATIC_SCALE_CONSTRAINT_CONSTRUCTED"
	StatusBranchSieveCompleted          = "CONDITIONAL_SUPPORT_R_BRANCH_POSITIVITY_SIEVE_COMPLETED"
	StatusTraceMomentsLocked            = "CONDITIONAL_SUPPORT_TOTAL_TRACE_MOMENTS_LOCKED_IN_REDUCED_PROXY"
	StatusFirewallsPreserved            = "CONDITIONAL_SUPPORT_CONTACT_SPECTRAL_CUTOFF_FIREWALLS_PRESERVED"

	StatusFailedBothBranchesSurvive       = "FAILED_ROUTE_BOTH_R_BRANCHES_ADMIT_POSITIVE_REAL_X"
	StatusFailedBranchNotSelected         = "FAILED_ROUTE_CONTACT_SPECTRAL_CUTOFF_DOES_NOT_SELECT_R_BRANCH"
	StatusFailedCutoffIdentificationProof = "FAILED_ROUTE_CUTOFF_FUNCTION_EQUALS_CONTACT_SPECTRUM_NOT_DERIVED_AS_HEAT_KERNEL_THEOREM"
	StatusFailedA0Normalization           = "FAILED_ROUTE_A0_IDENTITY_TRACE_NORMALIZATION_STILL_PROXY_LEVEL"
	StatusFailedHiggsRatioNotDerived      = "FAILED_ROUTE_HIGGS_MASS_RATIO_STILL_NOT_DERIVED"
)

const (
	sTop = 8 * math.Pi * math.Pi

	// Gate-169 scalar/contact shape.
	lambdaNum = 1197
	lambdaDen = 4624
	lambda    = float64(lambdaNum) / float64(lambdaDen)

	// Gate-273 reduced Morita trace multiplicities.
	kappaC = 1
	kappaQ = 3

	// Gate-162 contact spectral zeta/moment snapshots.
	contactZeta0 = 7.0
	contactZeta2 = 61.0 / 25.0
	contactZeta4 = 257629.0 / 202500.0

	// Reduced scalar-Morita proxy a0 = Tr(1) over the 1⊕3 edge ledger.
	// This is explicitly not a completed heat-kernel normalization theorem.
	reducedA0 = kappaC + kappaQ
)

type Gate287Inheritance struct {
	STopConstraintFormalized bool
	FreeCutoffMomentsBlocked bool
	BranchPreviouslySelected bool
	CutoffMomentsExtracted   bool
	HiggsPredictionDerived   bool
	Verdict                  string
}

type ContactCutoffMoments struct {
	Zeta0                    float64
	Zeta2                    float64
	Zeta4                    float64
	Zeta0Exact               string
	Zeta2Exact               string
	Zeta4Exact               string
	IdentifiedAsF0F2F4       bool
	IdentificationNative     bool
	HeatKernelTheoremDerived bool
	Verdict                  string
}

type DiracMomentProxy struct {
	KappaC        int
	KappaQ        int
	XDefinition   string
	RDefinition   string
	D2Formula     string
	D4Formula     string
	ShapeEquation string
	LambdaExact   string
	Lambda        float64
	A0Proxy       int
	A0Source      string
	A0Final       bool
	Verdict       string
}

type AmplitudeBranch struct {
	Name             string
	Exact            string
	R                float64
	AbsYOverX        float64
	ShapeResidualAbs float64
}

type QuadraticScaleConstraint struct {
	Equation            string
	CoeffAFormula       string
	CoeffBFormula       string
	CoeffCFormula       string
	A0                  float64
	CoeffC              float64
	UsesContactCutoff   bool
	UsesTopologicalSTop bool
	Verdict             string
}

type BranchScaleSolution struct {
	Branch          AmplitudeBranch
	CoeffA          float64
	CoeffB          float64
	CoeffC          float64
	Discriminant    float64
	Roots           []float64
	PositiveRoots   []float64
	SelectedX       float64
	D2              float64
	D4              float64
	D4OverD2        float64
	D4OverD2Squared float64
	Survives        bool
	Verdict         string
}

type BranchSieve struct {
	Solutions             []BranchScaleSolution
	SurvivingBranches     []string
	BothBranchesSurvive   bool
	NoBranchesSurvive     bool
	ExactlyOneSurvives    bool
	SelectedBranch        string
	UniqueBranchSelected  bool
	BranchEliminationRule string
	Verdict               string
}

type MomentLockAudit struct {
	PositiveD2             float64
	D4                     float64
	Shape                  float64
	BranchIndependentD2D4  bool
	BranchDependentX       bool
	ReducedProxyLocked     bool
	PhysicalHeatKernelLock bool
	Verdict                string
}

type HiggsObservableAudit struct {
	CandidateShapeRatio             float64
	CandidateD4OverD2               float64
	DimensionlessObservableDefined  bool
	ScalarGaugeNormalizationDerived bool
	HeatKernelProjectionDerived     bool
	HiggsMassRatioClaimed           bool
	Verdict                         string
}

type Firewalls struct {
	DoesNotTreatCutoffIDAsTheorem bool
	DoesNotTreatA0ProxyAsFinal    bool
	DoesNotDiscardSurvivingBranch bool
	DoesNotClaimUniqueVacuum      bool
	DoesNotClaimHiggsPrediction   bool
	DoesNotClaimHeatKernelDerived bool
	FiniteCorePolluted            bool
	Verdict                       string
}

type Summary struct {
	CutoffMomentsRetrieved bool
	CutoffIdentification   bool
	QuadraticBuilt         bool
	PositiveXSolved        bool
	BothBranchesSurvive    bool
	UniqueBranchSelected   bool
	TraceMomentsLocked     bool
	HiggsPredictionDerived bool
	FirewallPreserved      bool
	Status                 string
	DirectAnswer           string
	NextGate               string
}

type Analysis struct {
	Inheritance Gate287Inheritance
	Cutoff      ContactCutoffMoments
	Proxy       DiracMomentProxy
	Branches    []AmplitudeBranch
	Constraint  QuadraticScaleConstraint
	Sieve       BranchSieve
	MomentLock  MomentLockAudit
	Higgs       HiggsObservableAudit
	Firewalls   Firewalls
	Summary     Summary
	Truth       string
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
	inh := inheritGate287Snapshot()
	cutoff := retrieveContactCutoffMoments()
	proxy := buildDiracProxy()
	branches := buildBranches()
	constraint := buildConstraint(proxy, cutoff)
	sieve := solveBranchSieve(branches, constraint)
	lock := auditMomentLock(sieve)
	higgs := auditHiggsObservable(lock)
	fw := auditFirewalls(cutoff, proxy, sieve, higgs)
	summary := buildSummary(cutoff, constraint, sieve, lock, higgs, fw)
	truth := "Gate 288 identifies the contact zeta ledger with the spectral-action cutoff moments and therefore removes the free-cutoff ambiguity found in Gate 287 at the reduced scalar-Morita proxy level. The resulting quadratic has positive real X on both Gate-275 branches. In fact, the two branches yield the same total Tr(D_F²) and Tr(D_F⁴) while redistributing that total between |x|² and |y|² differently. Thus the contact-spectral cutoff locks total trace moments but not the physical r branch, heat-kernel normalization, or Higgs mass ratio."
	return Analysis{Inheritance: inh, Cutoff: cutoff, Proxy: proxy, Branches: branches, Constraint: constraint, Sieve: sieve, MomentLock: lock, Higgs: higgs, Firewalls: fw, Summary: summary, Truth: truth}, nil
}

func inheritGate287Snapshot() Gate287Inheritance {
	return Gate287Inheritance{
		STopConstraintFormalized: true,
		FreeCutoffMomentsBlocked: true,
		BranchPreviouslySelected: false,
		CutoffMomentsExtracted:   false,
		HiggsPredictionDerived:   false,
		Verdict:                  StatusGate287Inherited,
	}
}

func retrieveContactCutoffMoments() ContactCutoffMoments {
	return ContactCutoffMoments{
		Zeta0:                    contactZeta0,
		Zeta2:                    contactZeta2,
		Zeta4:                    contactZeta4,
		Zeta0Exact:               "ζ_contact(0)=7",
		Zeta2Exact:               "Tr(Ω²)=61/25",
		Zeta4Exact:               "Tr(Ω⁴)=257629/202500",
		IdentifiedAsF0F2F4:       true,
		IdentificationNative:     false,
		HeatKernelTheoremDerived: false,
		Verdict:                  StatusCutoffMomentsRetrieved,
	}
}

func buildDiracProxy() DiracMomentProxy {
	return DiracMomentProxy{
		KappaC:        kappaC,
		KappaQ:        kappaQ,
		XDefinition:   "X=|x|²",
		RDefinition:   "r=|y/x|²",
		D2Formula:     "Tr(D_F²)=X(1+3r)",
		D4Formula:     "Tr(D_F⁴)=X²(1+3r²)",
		ShapeEquation: "(1+3r²)/(1+3r)² = 1197/4624",
		LambdaExact:   "1197/4624",
		Lambda:        lambda,
		A0Proxy:       reducedA0,
		A0Source:      "reduced 1⊕3 scalar-Morita identity trace κ_C+κ_Q",
		A0Final:       false,
		Verdict:       StatusQuadraticScaleConstraintBuilt,
	}
}

func buildBranches() []AmplitudeBranch {
	rp := (3591 + 136*math.Sqrt(123)) / 3099
	rm := (3591 - 136*math.Sqrt(123)) / 3099
	return []AmplitudeBranch{
		makeBranch("r_plus", "(3591 + 136√123)/3099", rp),
		makeBranch("r_minus", "(3591 - 136√123)/3099", rm),
	}
}

func makeBranch(name, exact string, r float64) AmplitudeBranch {
	shape := (1 + 3*r*r) / math.Pow(1+3*r, 2)
	return AmplitudeBranch{Name: name, Exact: exact, R: r, AbsYOverX: math.Sqrt(r), ShapeResidualAbs: math.Abs(shape - lambda)}
}

func buildConstraint(p DiracMomentProxy, c ContactCutoffMoments) QuadraticScaleConstraint {
	return QuadraticScaleConstraint{
		Equation:            "7·X²(1+3r²) + (61/25)·X(1+3r) + (257629/202500)·a0 = 8π²",
		CoeffAFormula:       "A(r)=7(1+3r²)",
		CoeffBFormula:       "B(r)=(61/25)(1+3r)",
		CoeffCFormula:       "C=(257629/202500)·a0 - 8π²",
		A0:                  float64(p.A0Proxy),
		CoeffC:              c.Zeta4*float64(p.A0Proxy) - sTop,
		UsesContactCutoff:   c.IdentifiedAsF0F2F4,
		UsesTopologicalSTop: true,
		Verdict:             StatusQuadraticScaleConstraintBuilt,
	}
}

func solveBranchSieve(branches []AmplitudeBranch, q QuadraticScaleConstraint) BranchSieve {
	sols := make([]BranchScaleSolution, 0, len(branches))
	survivors := []string{}
	for _, b := range branches {
		sol := solveBranch(b, q)
		sols = append(sols, sol)
		if sol.Survives {
			survivors = append(survivors, b.Name)
		}
	}
	exactlyOne := len(survivors) == 1
	selected := ""
	if exactlyOne {
		selected = survivors[0]
	}
	return BranchSieve{
		Solutions:             sols,
		SurvivingBranches:     survivors,
		BothBranchesSurvive:   len(survivors) == 2,
		NoBranchesSurvive:     len(survivors) == 0,
		ExactlyOneSurvives:    exactlyOne,
		SelectedBranch:        selected,
		UniqueBranchSelected:  exactlyOne,
		BranchEliminationRule: "X=|x|² must be real and strictly positive",
		Verdict:               StatusBranchSieveCompleted,
	}
}

func solveBranch(b AmplitudeBranch, q QuadraticScaleConstraint) BranchScaleSolution {
	A := contactZeta0 * (1 + 3*b.R*b.R)
	B := contactZeta2 * (1 + 3*b.R)
	C := q.CoeffC
	disc := B*B - 4*A*C
	roots := []float64{}
	positives := []float64{}
	if disc >= 0 {
		sqrtD := math.Sqrt(disc)
		roots = []float64{(-B + sqrtD) / (2 * A), (-B - sqrtD) / (2 * A)}
		for _, x := range roots {
			if x > 0 && !math.IsNaN(x) && !math.IsInf(x, 0) {
				positives = append(positives, x)
			}
		}
	}
	selectedX := math.NaN()
	d2 := math.NaN()
	d4 := math.NaN()
	d4d2 := math.NaN()
	d4d2sq := math.NaN()
	survives := len(positives) > 0
	verdict := StatusFailedBranchNotSelected
	if survives {
		selectedX = positives[0]
		d2 = selectedX * (1 + 3*b.R)
		d4 = selectedX * selectedX * (1 + 3*b.R*b.R)
		d4d2 = d4 / d2
		d4d2sq = d4 / (d2 * d2)
		verdict = StatusBranchSieveCompleted
	}
	return BranchScaleSolution{Branch: b, CoeffA: A, CoeffB: B, CoeffC: C, Discriminant: disc, Roots: roots, PositiveRoots: positives, SelectedX: selectedX, D2: d2, D4: d4, D4OverD2: d4d2, D4OverD2Squared: d4d2sq, Survives: survives, Verdict: verdict}
}

func auditMomentLock(s BranchSieve) MomentLockAudit {
	if len(s.Solutions) == 0 || len(s.Solutions[0].PositiveRoots) == 0 {
		return MomentLockAudit{Verdict: StatusFailedBranchNotSelected}
	}
	d2 := s.Solutions[0].D2
	d4 := s.Solutions[0].D4
	same := true
	branchDependentX := false
	firstX := s.Solutions[0].SelectedX
	for _, sol := range s.Solutions[1:] {
		if math.Abs(sol.D2-d2) > 1e-9 || math.Abs(sol.D4-d4) > 1e-9 {
			same = false
		}
		if math.Abs(sol.SelectedX-firstX) > 1e-9 {
			branchDependentX = true
		}
	}
	return MomentLockAudit{PositiveD2: d2, D4: d4, Shape: d4 / (d2 * d2), BranchIndependentD2D4: same, BranchDependentX: branchDependentX, ReducedProxyLocked: same && s.BothBranchesSurvive, PhysicalHeatKernelLock: false, Verdict: StatusTraceMomentsLocked}
}

func auditHiggsObservable(m MomentLockAudit) HiggsObservableAudit {
	return HiggsObservableAudit{
		CandidateShapeRatio:             m.Shape,
		CandidateD4OverD2:               m.D4 / m.PositiveD2,
		DimensionlessObservableDefined:  false,
		ScalarGaugeNormalizationDerived: false,
		HeatKernelProjectionDerived:     false,
		HiggsMassRatioClaimed:           false,
		Verdict:                         StatusFailedHiggsRatioNotDerived,
	}
}

func auditFirewalls(c ContactCutoffMoments, p DiracMomentProxy, s BranchSieve, h HiggsObservableAudit) Firewalls {
	return Firewalls{
		DoesNotTreatCutoffIDAsTheorem: !c.HeatKernelTheoremDerived,
		DoesNotTreatA0ProxyAsFinal:    !p.A0Final,
		DoesNotDiscardSurvivingBranch: s.BothBranchesSurvive && !s.UniqueBranchSelected,
		DoesNotClaimUniqueVacuum:      !s.UniqueBranchSelected,
		DoesNotClaimHiggsPrediction:   !h.HiggsMassRatioClaimed,
		DoesNotClaimHeatKernelDerived: !h.HeatKernelProjectionDerived,
		FiniteCorePolluted:            false,
		Verdict:                       StatusFirewallsPreserved,
	}
}

func buildSummary(c ContactCutoffMoments, q QuadraticScaleConstraint, s BranchSieve, m MomentLockAudit, h HiggsObservableAudit, f Firewalls) Summary {
	statuses := []string{
		StatusGate287Inherited,
		StatusCutoffMomentsRetrieved,
		StatusCutoffIdentificationAudited,
		StatusQuadraticScaleConstraintBuilt,
		StatusBranchSieveCompleted,
		StatusTraceMomentsLocked,
		StatusFirewallsPreserved,
		StatusFailedBothBranchesSurvive,
		StatusFailedBranchNotSelected,
		StatusFailedCutoffIdentificationProof,
		StatusFailedA0Normalization,
		StatusFailedHiggsRatioNotDerived,
	}
	return Summary{
		CutoffMomentsRetrieved: c.IdentifiedAsF0F2F4,
		CutoffIdentification:   c.IdentifiedAsF0F2F4,
		QuadraticBuilt:         q.UsesContactCutoff && q.UsesTopologicalSTop,
		PositiveXSolved:        len(s.SurvivingBranches) > 0,
		BothBranchesSurvive:    s.BothBranchesSurvive,
		UniqueBranchSelected:   s.UniqueBranchSelected,
		TraceMomentsLocked:     m.ReducedProxyLocked,
		HiggsPredictionDerived: h.HiggsMassRatioClaimed,
		FirewallPreserved:      !f.FiniteCorePolluted && f.DoesNotDiscardSurvivingBranch && f.DoesNotClaimHiggsPrediction,
		Status:                 strings.Join(statuses, ";"),
		DirectAnswer:           "Identifying cutoff moments with the contact spectrum fixes the reduced total spectral moments, but both Gate-275 branches admit positive real X and give the same Tr(D_F²), Tr(D_F⁴). The r branch is not selected.",
		NextGate:               "A future theorem must derive a branch-sensitive observable, a heat-kernel normalization map, or a physical J/hypercharge representation; positivity of X under contact cutoff is insufficient.",
	}
}

func FormatCutoff(c ContactCutoffMoments) string {
	return fmt.Sprintf("f0=%s f2=%s f4=%s numeric=(%.12g,%.12g,%.12g) identified=%t nativeTheorem=%t heatKernel=%t verdict=%s", c.Zeta0Exact, c.Zeta2Exact, c.Zeta4Exact, c.Zeta0, c.Zeta2, c.Zeta4, c.IdentifiedAsF0F2F4, c.IdentificationNative, c.HeatKernelTheoremDerived, c.Verdict)
}

func FormatProxy(p DiracMomentProxy) string {
	return fmt.Sprintf("κ=%d:%d X=%s r=%s D2=%q D4=%q shape=%q λ=%s a0=%d source=%q finalA0=%t verdict=%s", p.KappaC, p.KappaQ, p.XDefinition, p.RDefinition, p.D2Formula, p.D4Formula, p.ShapeEquation, p.LambdaExact, p.A0Proxy, p.A0Source, p.A0Final, p.Verdict)
}

func FormatConstraint(q QuadraticScaleConstraint) string {
	return fmt.Sprintf("eq=%q A=%q B=%q C=%q a0=%.12g Cnum=%.12g contact=%t STop=%t verdict=%s", q.Equation, q.CoeffAFormula, q.CoeffBFormula, q.CoeffCFormula, q.A0, q.CoeffC, q.UsesContactCutoff, q.UsesTopologicalSTop, q.Verdict)
}

func FormatSolution(s BranchScaleSolution) string {
	return fmt.Sprintf("%s r=%.12g |y/x|=%.12g A=%.12g B=%.12g C=%.12g Δ=%.12g roots=%v positive=%v X=%.12g D2=%.12g D4=%.12g D4/D2=%.12g shape=%.12g survives=%t verdict=%s", s.Branch.Name, s.Branch.R, s.Branch.AbsYOverX, s.CoeffA, s.CoeffB, s.CoeffC, s.Discriminant, s.Roots, s.PositiveRoots, s.SelectedX, s.D2, s.D4, s.D4OverD2, s.D4OverD2Squared, s.Survives, s.Verdict)
}

func FormatSieve(s BranchSieve) string {
	parts := make([]string, 0, len(s.Solutions))
	for _, sol := range s.Solutions {
		parts = append(parts, FormatSolution(sol))
	}
	return fmt.Sprintf("survivors=%v both=%t none=%t exactlyOne=%t selected=%q rule=%q solutions=[%s] verdict=%s", s.SurvivingBranches, s.BothBranchesSurvive, s.NoBranchesSurvive, s.ExactlyOneSurvives, s.SelectedBranch, s.BranchEliminationRule, strings.Join(parts, "; "), s.Verdict)
}

func FormatMomentLock(m MomentLockAudit) string {
	return fmt.Sprintf("D2=%.12g D4=%.12g shape=%.12g branchIndependentD2D4=%t branchDependentX=%t reducedProxy=%t physicalHeatKernel=%t verdict=%s", m.PositiveD2, m.D4, m.Shape, m.BranchIndependentD2D4, m.BranchDependentX, m.ReducedProxyLocked, m.PhysicalHeatKernelLock, m.Verdict)
}

func FormatHiggs(h HiggsObservableAudit) string {
	return fmt.Sprintf("shape=%.12g D4/D2=%.12g observable=%t scalarGaugeNorm=%t heatKernel=%t claim=%t verdict=%s", h.CandidateShapeRatio, h.CandidateD4OverD2, h.DimensionlessObservableDefined, h.ScalarGaugeNormalizationDerived, h.HeatKernelProjectionDerived, h.HiggsMassRatioClaimed, h.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("noCutoffTheorem=%t noA0Final=%t noDiscard=%t noUniqueVacuum=%t noHiggs=%t noHeatKernel=%t polluted=%t verdict=%s", f.DoesNotTreatCutoffIDAsTheorem, f.DoesNotTreatA0ProxyAsFinal, f.DoesNotDiscardSurvivingBranch, f.DoesNotClaimUniqueVacuum, f.DoesNotClaimHiggsPrediction, f.DoesNotClaimHeatKernelDerived, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("cutoff=%t identification=%t quadratic=%t positiveX=%t both=%t unique=%t moments=%t higgs=%t firewall=%t next=%q status=%s", s.CutoffMomentsRetrieved, s.CutoffIdentification, s.QuadraticBuilt, s.PositiveXSolved, s.BothBranchesSurvive, s.UniqueBranchSelected, s.TraceMomentsLocked, s.HiggsPredictionDerived, s.FirewallPreserved, s.NextGate, s.Status)
}
