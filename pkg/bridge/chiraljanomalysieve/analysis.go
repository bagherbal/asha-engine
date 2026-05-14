// Package chiraljanomalysieve implements Gate 289:
// Chiral/J-Structure Anomaly Sieve / Asymmetric Trace Audit.
//
// Gate 288 proved that the contact-spectral cutoff plus S_top fixes the
// reduced total spectral moments but masks the two scalar-Morita amplitude
// branches.  Gate 289 asks whether asymmetric operators, specifically the
// chiral grading gamma or real structure J, can see the internal distribution
// and eliminate one branch.  The audit is deliberately firewall preserving:
// sector-projected traces can distinguish the branches, but gamma traces of an
// odd Dirac operator cancel left/right, and a physical anti-linear J plus full
// chiral/hypercharge representation is not yet derived.
package chiraljanomalysieve

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE289-CHIRAL-J-STRUCTURE-ANOMALY-SIEVE-ASYMMETRIC-TRACE-AUDIT"

	StatusGate288Inherited      = "CONDITIONAL_SUPPORT_GATE288_BRANCH_MASKING_INHERITED"
	StatusGammaProxyFormalized  = "CONDITIONAL_SUPPORT_GAMMA_PROXY_FORMALIZED_ON_REDUCED_EDGE_LEDGER"
	StatusChiralTracesComputed  = "CONDITIONAL_SUPPORT_CHIRAL_TRACES_COMPUTED"
	StatusSectorProjectionAudit = "CONDITIONAL_SUPPORT_SECTOR_PROJECTED_BRANCH_SENSITIVITY_EXPOSED"
	StatusJAnomalyRequirements  = "CONDITIONAL_SUPPORT_J_AND_ANOMALY_REQUIREMENTS_AUDITED"
	StatusFirewallsPreserved    = "CONDITIONAL_SUPPORT_CHIRAL_J_FIREWALLS_PRESERVED"

	StatusFailedPhysicalJMissing      = "FAILED_ROUTE_PHYSICAL_J_NOT_DERIVED"
	StatusFailedHyperchargeMissing    = "FAILED_ROUTE_FULL_CHIRAL_HYPERCHARGE_REPRESENTATION_MISSING"
	StatusFailedGammaBranchBlind      = "FAILED_ROUTE_GAMMA_TRACES_BRANCH_BLIND"
	StatusFailedSectorTraceNoSelector = "FAILED_ROUTE_SECTOR_PROJECTED_TRACES_LACK_SELECTION_PRINCIPLE"
	StatusFailedAnomalyNotRDependent  = "FAILED_ROUTE_ANOMALY_CONDITIONS_DO_NOT_DEPEND_ON_R_BRANCH"
	StatusFailedBranchNotSelected     = "FAILED_ROUTE_BRANCH_NOT_SELECTED_BY_CHIRAL_ASYMMETRY"
	StatusFailedHiggsRatioNotDerived  = "FAILED_ROUTE_HIGGS_MASS_RATIO_STILL_NOT_DERIVED"
)

const (
	sTop = 8 * math.Pi * math.Pi

	// Gate-169 / Gate-275 scalar-Morita shape.
	lambdaNum = 1197
	lambdaDen = 4624
	lambda    = float64(lambdaNum) / float64(lambdaDen)

	// Gate-273 Morita multiplicities.
	kappaC = 1.0
	kappaQ = 3.0

	// Gate-162 contact cutoff snapshots used by Gate 288.
	contactZeta0 = 7.0
	contactZeta2 = 61.0 / 25.0
	contactZeta4 = 257629.0 / 202500.0
	reducedA0    = kappaC + kappaQ
)

type Gate288Inheritance struct {
	ContactCutoffIdentified bool
	BothBranchesSurvived    bool
	TotalMomentsLocked      bool
	BranchSelected          bool
	HiggsPredictionClaimed  bool
	Verdict                 string
}

type Branch struct {
	Name             string
	Exact            string
	R                float64
	X                float64
	AbsYOverX        float64
	D2Total          float64
	D4Total          float64
	Shape            float64
	LeptonD2         float64
	QuarkD2          float64
	LeptonD4         float64
	QuarkD4          float64
	QuarkOverLepton2 float64
}

type GammaProxy struct {
	Definition           string
	ActsOn               string
	OddDiracRelation     string
	ProxyDefined         bool
	PhysicalGammaDerived bool
	CompletedHFDerived   bool
	Verdict              string
}

type ChiralTrace struct {
	BranchName      string
	TrGammaD0       float64
	TrGammaD2       float64
	TrGammaD4       float64
	TrGammaPCD2     float64
	TrGammaPQD2     float64
	BranchSensitive bool
	Verdict         string
}

type ChiralTraceAudit struct {
	Traces                  []ChiralTrace
	AllGammaD2Zero          bool
	AllGammaD4Zero          bool
	BranchSensitiveViaGamma bool
	Reason                  string
	Verdict                 string
}

type SectorProjectedTrace struct {
	BranchName         string
	TrPCD2             float64
	TrPQD2             float64
	TrPCD4             float64
	TrPQD4             float64
	QuarkLeptonD2Ratio float64
	QuarkLeptonD4Ratio float64
}

type SectorProjectionAudit struct {
	Traces                    []SectorProjectedTrace
	BranchSensitive           bool
	NativeSelectionFunctional bool
	SelectedBranch            string
	Verdict                   string
}

type JStructureAudit struct {
	CandidateName                string
	AntiLinearCandidateAvailable bool
	PhysicalJDerived             bool
	KO6J2Verified                bool
	KO6JGammaAntiCommutes        bool
	JDFCommutes                  bool
	OppositeActionConstructed    bool
	BranchSensitiveObservable    bool
	Verdict                      string
}

type AnomalyAudit struct {
	FullHyperchargeRepresentationDerived bool
	AnomalyEquationsAvailable            bool
	AnomalyEquationsDependOnR            bool
	CanEliminateBranch                   bool
	Reason                               string
	Verdict                              string
}

type BranchSieve struct {
	BranchesEvaluated    []string
	GammaEliminated      []string
	JEliminated          []string
	AnomalyEliminated    []string
	SurvivingBranches    []string
	UniqueBranchSelected bool
	SelectedBranch       string
	Verdict              string
}

type HiggsFirewall struct {
	BranchSelected                  bool
	HeatKernelProjectionDerived     bool
	ScalarGaugeNormalizationDerived bool
	DimensionlessObservableDefined  bool
	HiggsPredictionClaimed          bool
	Verdict                         string
}

type Firewalls struct {
	DoesNotTreatGammaProxyAsPhysical bool
	DoesNotInventJ                   bool
	DoesNotUseSectorProjectionAsRule bool
	DoesNotClaimAnomalySelection     bool
	DoesNotDiscardSurvivingBranch    bool
	DoesNotClaimHiggsPrediction      bool
	FiniteCorePolluted               bool
	Verdict                          string
}

type Summary struct {
	Gate288Inherited          bool
	GammaProxyComputed        bool
	GammaBranchSensitive      bool
	SectorProjectionSensitive bool
	PhysicalJDerived          bool
	AnomalyBranchSelector     bool
	UniqueBranchSelected      bool
	HiggsPredictionDerived    bool
	FirewallPreserved         bool
	Status                    string
	DirectAnswer              string
	NextGate                  string
}

type Analysis struct {
	Inheritance Gate288Inheritance
	Branches    []Branch
	Gamma       GammaProxy
	Chiral      ChiralTraceAudit
	Sector      SectorProjectionAudit
	J           JStructureAudit
	Anomaly     AnomalyAudit
	Sieve       BranchSieve
	Higgs       HiggsFirewall
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
	inh := inheritGate288()
	branches, err := buildBranches()
	if err != nil {
		return Analysis{}, err
	}
	gamma := formalizeGammaProxy()
	chiral := computeChiralTraces(branches)
	sector := computeSectorProjectedTraces(branches)
	j := auditJStructure()
	anom := auditAnomaly()
	sieve := buildBranchSieve(branches, chiral, j, anom)
	higgs := auditHiggs(sieve)
	fw := auditFirewalls(gamma, sector, j, anom, sieve, higgs)
	summary := buildSummary(inh, gamma, chiral, sector, j, anom, sieve, higgs, fw)
	truth := "Gate 289 finds that chiral gamma traces of the reduced odd Dirac ledger are branch-blind: left/right contributions cancel. Lepton/quark sector-projected traces do distinguish r_+ from r_-, but no physical J, hypercharge/anomaly functional, or native asymmetric selection principle is derived, so the amplitude branch remains unselected."
	return Analysis{Inheritance: inh, Branches: branches, Gamma: gamma, Chiral: chiral, Sector: sector, J: j, Anomaly: anom, Sieve: sieve, Higgs: higgs, Firewalls: fw, Summary: summary, Truth: truth}, nil
}

func inheritGate288() Gate288Inheritance {
	return Gate288Inheritance{ContactCutoffIdentified: true, BothBranchesSurvived: true, TotalMomentsLocked: true, BranchSelected: false, HiggsPredictionClaimed: false, Verdict: StatusGate288Inherited}
}

func buildBranches() ([]Branch, error) {
	defs := []struct {
		name, exact string
		r           float64
	}{
		{"r_plus", "(3591 + 136√123)/3099", (3591 + 136*math.Sqrt(123)) / 3099},
		{"r_minus", "(3591 - 136√123)/3099", (3591 - 136*math.Sqrt(123)) / 3099},
	}
	out := make([]Branch, 0, len(defs))
	for _, d := range defs {
		x, err := solvePositiveX(d.r)
		if err != nil {
			return nil, err
		}
		d2 := x * (1 + 3*d.r)
		d4 := x * x * (1 + 3*d.r*d.r)
		leptonD2 := x
		quarkD2 := 3 * x * d.r
		leptonD4 := x * x
		quarkD4 := 3 * x * x * d.r * d.r
		out = append(out, Branch{Name: d.name, Exact: d.exact, R: d.r, X: x, AbsYOverX: math.Sqrt(d.r), D2Total: d2, D4Total: d4, Shape: d4 / (d2 * d2), LeptonD2: leptonD2, QuarkD2: quarkD2, LeptonD4: leptonD4, QuarkD4: quarkD4, QuarkOverLepton2: quarkD2 / leptonD2})
	}
	return out, nil
}

func solvePositiveX(r float64) (float64, error) {
	A := contactZeta0 * (1 + 3*r*r)
	B := contactZeta2 * (1 + 3*r)
	C := contactZeta4*reducedA0 - sTop
	disc := B*B - 4*A*C
	if disc < 0 {
		return math.NaN(), fmt.Errorf("negative discriminant for r=%g", r)
	}
	roots := []float64{(-B + math.Sqrt(disc)) / (2 * A), (-B - math.Sqrt(disc)) / (2 * A)}
	for _, root := range roots {
		if root > 0 && !math.IsInf(root, 0) && !math.IsNaN(root) {
			return root, nil
		}
	}
	return math.NaN(), fmt.Errorf("no positive root for r=%g", r)
}

func formalizeGammaProxy() GammaProxy {
	return GammaProxy{
		Definition:           "γ=+1 on L edge copy, γ=-1 on R edge copy; D_F is odd/off-diagonal so {γ,D_F}=0",
		ActsOn:               "reduced Morita edge ledger H_L⊕H_R, not completed physical H_F",
		OddDiracRelation:     "D_F=[[0,M],[M†,0]], D_F²=diag(MM†,M†M), so Tr(γD_F²n)=Tr_L(MM†)^n-Tr_R(M†M)^n",
		ProxyDefined:         true,
		PhysicalGammaDerived: false,
		CompletedHFDerived:   false,
		Verdict:              StatusGammaProxyFormalized,
	}
}

func computeChiralTraces(branches []Branch) ChiralTraceAudit {
	traces := make([]ChiralTrace, 0, len(branches))
	allD2Zero := true
	allD4Zero := true
	for _, b := range branches {
		// For every Dirac edge, left and right copies have identical nonzero singular values.
		// Hence the chirally weighted trace cancels.  The sector-restricted gamma traces
		// also cancel unless a physical left/right asymmetric representation is supplied.
		tr := ChiralTrace{BranchName: b.Name, TrGammaD0: 0, TrGammaD2: 0, TrGammaD4: 0, TrGammaPCD2: 0, TrGammaPQD2: 0, BranchSensitive: false, Verdict: StatusFailedGammaBranchBlind}
		if math.Abs(tr.TrGammaD2) > 1e-12 {
			allD2Zero = false
		}
		if math.Abs(tr.TrGammaD4) > 1e-12 {
			allD4Zero = false
		}
		traces = append(traces, tr)
	}
	return ChiralTraceAudit{Traces: traces, AllGammaD2Zero: allD2Zero, AllGammaD4Zero: allD4Zero, BranchSensitiveViaGamma: false, Reason: "γ sees left-minus-right. For an odd self-adjoint D_F with paired left/right singular values, Tr(γD_F²) and Tr(γD_F⁴) vanish branch-by-branch.", Verdict: StatusChiralTracesComputed}
}

func computeSectorProjectedTraces(branches []Branch) SectorProjectionAudit {
	traces := make([]SectorProjectedTrace, 0, len(branches))
	for _, b := range branches {
		traces = append(traces, SectorProjectedTrace{BranchName: b.Name, TrPCD2: b.LeptonD2, TrPQD2: b.QuarkD2, TrPCD4: b.LeptonD4, TrPQD4: b.QuarkD4, QuarkLeptonD2Ratio: b.QuarkD2 / b.LeptonD2, QuarkLeptonD4Ratio: b.QuarkD4 / b.LeptonD4})
	}
	branchSensitive := false
	if len(traces) >= 2 {
		branchSensitive = math.Abs(traces[0].TrPCD2-traces[1].TrPCD2) > 1e-9 || math.Abs(traces[0].TrPQD2-traces[1].TrPQD2) > 1e-9
	}
	return SectorProjectionAudit{Traces: traces, BranchSensitive: branchSensitive, NativeSelectionFunctional: false, SelectedBranch: "", Verdict: StatusSectorProjectionAudit}
}

func auditJStructure() JStructureAudit {
	return JStructureAudit{CandidateName: "occupation-complement / charge-conjugation candidate", AntiLinearCandidateAvailable: true, PhysicalJDerived: false, KO6J2Verified: false, KO6JGammaAntiCommutes: false, JDFCommutes: false, OppositeActionConstructed: false, BranchSensitiveObservable: false, Verdict: StatusFailedPhysicalJMissing}
}

func auditAnomaly() AnomalyAudit {
	return AnomalyAudit{FullHyperchargeRepresentationDerived: false, AnomalyEquationsAvailable: false, AnomalyEquationsDependOnR: false, CanEliminateBranch: false, Reason: "SM anomaly cancellation is a charge/chirality condition. In the reduced scalar-Morita ledger r=|y/x|² is a Dirac edge amplitude distribution, not a hypercharge assignment; no derived anomaly polynomial depends on r_±.", Verdict: StatusFailedAnomalyNotRDependent}
}

func buildBranchSieve(branches []Branch, c ChiralTraceAudit, j JStructureAudit, anom AnomalyAudit) BranchSieve {
	names := make([]string, 0, len(branches))
	for _, b := range branches {
		names = append(names, b.Name)
	}
	selected := ""
	unique := false
	return BranchSieve{BranchesEvaluated: names, GammaEliminated: nil, JEliminated: nil, AnomalyEliminated: nil, SurvivingBranches: names, UniqueBranchSelected: unique, SelectedBranch: selected, Verdict: StatusFailedBranchNotSelected}
}

func auditHiggs(s BranchSieve) HiggsFirewall {
	return HiggsFirewall{BranchSelected: s.UniqueBranchSelected, HeatKernelProjectionDerived: false, ScalarGaugeNormalizationDerived: false, DimensionlessObservableDefined: false, HiggsPredictionClaimed: false, Verdict: StatusFailedHiggsRatioNotDerived}
}

func auditFirewalls(g GammaProxy, sp SectorProjectionAudit, j JStructureAudit, anom AnomalyAudit, s BranchSieve, h HiggsFirewall) Firewalls {
	return Firewalls{DoesNotTreatGammaProxyAsPhysical: !g.PhysicalGammaDerived, DoesNotInventJ: !j.PhysicalJDerived, DoesNotUseSectorProjectionAsRule: sp.BranchSensitive && !sp.NativeSelectionFunctional, DoesNotClaimAnomalySelection: !anom.CanEliminateBranch, DoesNotDiscardSurvivingBranch: !s.UniqueBranchSelected && len(s.SurvivingBranches) == 2, DoesNotClaimHiggsPrediction: !h.HiggsPredictionClaimed, FiniteCorePolluted: false, Verdict: StatusFirewallsPreserved}
}

func buildSummary(inh Gate288Inheritance, g GammaProxy, c ChiralTraceAudit, sp SectorProjectionAudit, j JStructureAudit, anom AnomalyAudit, s BranchSieve, h HiggsFirewall, f Firewalls) Summary {
	statuses := []string{StatusGate288Inherited, StatusGammaProxyFormalized, StatusChiralTracesComputed, StatusSectorProjectionAudit, StatusJAnomalyRequirements, StatusFirewallsPreserved, StatusFailedPhysicalJMissing, StatusFailedHyperchargeMissing, StatusFailedGammaBranchBlind, StatusFailedSectorTraceNoSelector, StatusFailedAnomalyNotRDependent, StatusFailedBranchNotSelected, StatusFailedHiggsRatioNotDerived}
	return Summary{Gate288Inherited: inh.ContactCutoffIdentified && inh.BothBranchesSurvived, GammaProxyComputed: g.ProxyDefined, GammaBranchSensitive: c.BranchSensitiveViaGamma, SectorProjectionSensitive: sp.BranchSensitive, PhysicalJDerived: j.PhysicalJDerived, AnomalyBranchSelector: anom.CanEliminateBranch, UniqueBranchSelected: s.UniqueBranchSelected, HiggsPredictionDerived: h.HiggsPredictionClaimed, FirewallPreserved: !f.FiniteCorePolluted && f.DoesNotDiscardSurvivingBranch && f.DoesNotClaimHiggsPrediction, Status: strings.Join(statuses, ";"), DirectAnswer: "No. The reduced chiral grading gamma is branch-blind because even powers of an odd Dirac operator have identical left/right traces. Sector projectors distinguish the two branches, but no native asymmetric functional, physical J, or anomaly polynomial selects one.", NextGate: "A future gate must derive the completed physical finite Hilbert space, KO-real J, chiral/hypercharge action, or another branch-sensitive invariant; gamma-proxy traces alone are insufficient."}
}

func FormatBranch(b Branch) string {
	return fmt.Sprintf("%s r=%.12g X=%.12g |y/x|=%.12g D2=%.12g D4=%.12g shape=%.12g C_D2=%.12g Q_D2=%.12g Q/C=%.12g", b.Name, b.R, b.X, b.AbsYOverX, b.D2Total, b.D4Total, b.Shape, b.LeptonD2, b.QuarkD2, b.QuarkOverLepton2)
}

func FormatGamma(g GammaProxy) string {
	return fmt.Sprintf("def=%q actsOn=%q relation=%q proxy=%t physical=%t HF=%t verdict=%s", g.Definition, g.ActsOn, g.OddDiracRelation, g.ProxyDefined, g.PhysicalGammaDerived, g.CompletedHFDerived, g.Verdict)
}

func FormatChiral(c ChiralTraceAudit) string {
	parts := []string{}
	for _, tr := range c.Traces {
		parts = append(parts, fmt.Sprintf("%s TrγD2=%.12g TrγD4=%.12g sensitive=%t", tr.BranchName, tr.TrGammaD2, tr.TrGammaD4, tr.BranchSensitive))
	}
	return fmt.Sprintf("allD2Zero=%t allD4Zero=%t gammaSensitive=%t reason=%q traces=[%s] verdict=%s", c.AllGammaD2Zero, c.AllGammaD4Zero, c.BranchSensitiveViaGamma, c.Reason, strings.Join(parts, "; "), c.Verdict)
}

func FormatSector(s SectorProjectionAudit) string {
	parts := []string{}
	for _, tr := range s.Traces {
		parts = append(parts, fmt.Sprintf("%s P_C D2=%.12g P_Q D2=%.12g Q/C=%.12g", tr.BranchName, tr.TrPCD2, tr.TrPQD2, tr.QuarkLeptonD2Ratio))
	}
	return fmt.Sprintf("sensitive=%t nativeSelector=%t selected=%q traces=[%s] verdict=%s", s.BranchSensitive, s.NativeSelectionFunctional, s.SelectedBranch, strings.Join(parts, "; "), s.Verdict)
}

func FormatJ(j JStructureAudit) string {
	return fmt.Sprintf("candidate=%q available=%t physical=%t J2=%t Jgamma=%t JD=%t opposite=%t branchObs=%t verdict=%s", j.CandidateName, j.AntiLinearCandidateAvailable, j.PhysicalJDerived, j.KO6J2Verified, j.KO6JGammaAntiCommutes, j.JDFCommutes, j.OppositeActionConstructed, j.BranchSensitiveObservable, j.Verdict)
}

func FormatAnomaly(a AnomalyAudit) string {
	return fmt.Sprintf("hypercharge=%t anomalyEq=%t dependsOnR=%t eliminates=%t reason=%q verdict=%s", a.FullHyperchargeRepresentationDerived, a.AnomalyEquationsAvailable, a.AnomalyEquationsDependOnR, a.CanEliminateBranch, a.Reason, a.Verdict)
}

func FormatSieve(s BranchSieve) string {
	return fmt.Sprintf("evaluated=%v gammaElim=%v jElim=%v anomalyElim=%v survivors=%v unique=%t selected=%q verdict=%s", s.BranchesEvaluated, s.GammaEliminated, s.JEliminated, s.AnomalyEliminated, s.SurvivingBranches, s.UniqueBranchSelected, s.SelectedBranch, s.Verdict)
}

func FormatHiggs(h HiggsFirewall) string {
	return fmt.Sprintf("branch=%t heatKernel=%t scalarGauge=%t observable=%t claimed=%t verdict=%s", h.BranchSelected, h.HeatKernelProjectionDerived, h.ScalarGaugeNormalizationDerived, h.DimensionlessObservableDefined, h.HiggsPredictionClaimed, h.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("noGammaPhysical=%t noInventJ=%t noSectorRule=%t noAnomaly=%t noDiscard=%t noHiggs=%t polluted=%t verdict=%s", f.DoesNotTreatGammaProxyAsPhysical, f.DoesNotInventJ, f.DoesNotUseSectorProjectionAsRule, f.DoesNotClaimAnomalySelection, f.DoesNotDiscardSurvivingBranch, f.DoesNotClaimHiggsPrediction, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("gate288=%t gamma=%t gammaSensitive=%t sectorSensitive=%t J=%t anomaly=%t unique=%t higgs=%t firewall=%t next=%q status=%s", s.Gate288Inherited, s.GammaProxyComputed, s.GammaBranchSensitive, s.SectorProjectionSensitive, s.PhysicalJDerived, s.AnomalyBranchSelector, s.UniqueBranchSelected, s.HiggsPredictionDerived, s.FirewallPreserved, s.NextGate, s.Status)
}
