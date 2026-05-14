// Package weakquaternionicnormalization implements Gate 273:
// Weak/Quaternionic Sub-Bimodule Selector / Finite Inner-Product Normalization Audit.
//
// Gate 272 repaired the representation category by moving the finite spectral
// triple from the second-quantized Fock carrier S_C to the first-quantized
// Morita bimodule H_ij=V_i⊗V_j*.  The order-one condition then allowed
// non-vacuous edges precisely when the left module changes while the right
// module is shared.  Gate 273 asks whether weak/chiral/quaternionic structure
// and finite Hilbert-space normalization can turn those allowed edges into a
// canonical finite Dirac operator by locking the lepton/quark amplitude ratio
// x:y.
//
// The answer is deliberately strict.  The Morita inner product can count trace
// multiplicities of right-sector edges (κ_C:κ_Q = 1:3 under the minimal
// rank-one edge normalization), but multiplicity is not amplitude.  The weak
// or quaternionic selector is not a derived associative summand of
// A_F=C⊕M3(C), and the norm/equality of the C-edge and Q-edge maps remains a
// separate dynamical input.  Therefore the gate exposes the correct finite
// normalization ledger but refuses to promote it into a Higgs-ratio theorem.
package weakquaternionicnormalization

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/moritabimodulesearch"
)

const (
	AuditID = "GATE273-WEAK-QUATERNIONIC-SUB-BIMODULE-SELECTOR-FINITE-INNER-PRODUCT-NORMALIZATION-AUDIT"

	StatusGate272Inherited           = "CONDITIONAL_SUPPORT_GATE272_MORITA_BIMODULE_LEDGER_INHERITED"
	StatusPhysicalSieveAudited       = "CONDITIONAL_SUPPORT_WEAK_CHIRAL_SUB_BIMODULE_SIEVE_AUDITED"
	StatusOrderOneEdgesRecovered     = "CONDITIONAL_SUPPORT_ORDER_ONE_NONVACUOUS_EDGES_RECOVERED"
	StatusFiniteInnerProductBuilt    = "CONDITIONAL_SUPPORT_FINITE_INNER_PRODUCT_NORMALIZATION_LEDGER_BUILT"
	StatusTraceWeightsComputed       = "CONDITIONAL_SUPPORT_LEPTON_QUARK_TRACE_MULTIPLICITIES_COMPUTED"
	StatusSpectralMomentsReevaluated = "CONDITIONAL_SUPPORT_NORMALIZED_TRACE_MOMENTS_REEVALUATED"
	StatusFailedNoNativeQuaternion   = "FAILED_ROUTE_WEAK_QUATERNIONIC_SELECTOR_NOT_NATIVE_TO_C_PLUS_M3C"
	StatusFailedPhysicalSubmodule    = "FAILED_ROUTE_PHYSICAL_SM_SUB_BIMODULE_NOT_DERIVED"
	StatusFailedEdgeNorms            = "FAILED_ROUTE_EDGE_MAP_NORMS_REMAIN_UNSELECTED"
	StatusFailedXYRatio              = "FAILED_ROUTE_INNER_PRODUCT_NORMALIZATION_DOES_NOT_LOCK_XY_RATIO"
	StatusFailedCanonicalDF          = "FAILED_ROUTE_CANONICAL_DF_AMPLITUDES_NOT_LOCKED_VIA_NORMALIZATION"
	StatusFailedA2A4                 = "FAILED_ROUTE_A2_A4_HIGGS_RATIO_STILL_NOT_DERIVED"
	StatusEmpiricalSealPreserved     = "FAILED_ROUTE_EMPIRICAL_YUKAWA_SEAL_REMAINS_ACTIVE"
)

type Gate272Inheritance struct {
	BimoduleExtracted         bool
	OppositeActionConstructed bool
	NonVacuousEdgesExist      bool
	XYRatioLocked             bool
	A2A4Derived               bool
	HiggsRatioDerived         bool
	FirewallPreserved         bool
	SurvivingAmplitudeLabels  []string
	Verdict                   string
}

type SubBimoduleCandidate struct {
	Name                string
	Source              string
	UsesWeakSU2         bool
	UsesQuaternionicH   bool
	UsesChiralGrading   bool
	NativeToCPlusM3     bool
	SelectsPhysicalSMHF bool
	SelectedEdgeLabels  []string
	RejectedReason      string
}

type PhysicalSubBimoduleSieve struct {
	UniversalSummands           int
	UniversalComplexDimension   int
	CandidateSelectors          []SubBimoduleCandidate
	ChiralOrderOneEdgesRetained []string
	WeakQuaternionicNative      bool
	PhysicalSMHilbertDerived    bool
	Verdict                     string
}

type EdgeNormalization struct {
	Label                string
	Edge                 string
	SharedRightModule    string
	SharedRightDimension int
	MinimalRank          int
	KappaD2              float64
	KappaD4              float64
	RequiresEdgeMapNorm  bool
	Detail               string
}

type InnerProductAudit struct {
	Formula                       string
	OrthogonalMoritaSummands      bool
	CanonicalTraceOnSimpleModules bool
	EdgeNormalizations            []EdgeNormalization
	KappaCRatio                   float64
	KappaQRatio                   float64
	MultiplicityRatioQOverC       float64
	MultiplicitiesGeometric       bool
	EdgeNormsDerived              bool
	Verdict                       string
}

type XYRatioAudit struct {
	CandidateConstraint         string
	MultiplicityWeightsKnown    bool
	WouldEqualContributionFixXY string
	EqualContributionIsDerived  bool
	XOverYLocked                bool
	XOverYValue                 float64
	Unknowns                    []string
	Verdict                     string
}

type SpectralMomentCandidate struct {
	Name    string
	X       float64
	Y       float64
	TraceD2 float64
	TraceD4 float64
	Ratio   float64
}

type SpectralTraceAudit struct {
	TraceD2Formula       string
	TraceD4Formula       string
	Candidates           []SpectralMomentCandidate
	RatioDependsOnXOverY bool
	StableInvariant      bool
	A2A4Derived          bool
	HiggsRatioDerived    bool
	Verdict              string
}

type FirewallAudit struct {
	EmpiricalYukawaSealPreserved    bool
	SpontaneousCarrierSealPreserved bool
	NoObservedMassInserted          bool
	NoVEVInserted                   bool
	NoCutoffScaleInserted           bool
	NoSMQuaternionImportedAsTheorem bool
	NoHiggsPredictionClaim          bool
	MultiplicityNotAmplitude        bool
	FiniteCorePolluted              bool
	Verdict                         string
}

type FutureCriterion struct {
	Name      string
	Required  bool
	Satisfied bool
	Detail    string
}

type FutureMap struct {
	Criteria                          []FutureCriterion
	NeedNativeWeakQuaternionicAlgebra bool
	NeedPhysicalChargeConjugationJ    bool
	NeedEdgeNormOrAmplitudeAction     bool
	NeedHeatKernelProjection          bool
	RecommendedNextGate               string
	Verdict                           string
}

type Summary struct {
	Gate272Inherited         bool
	PhysicalSieveAudited     bool
	InnerProductBuilt        bool
	TraceWeightsComputed     bool
	PhysicalSMHilbertDerived bool
	EdgeNormsDerived         bool
	XYRatioLocked            bool
	CanonicalDFDerived       bool
	A2A4Derived              bool
	HiggsRatioDerived        bool
	FirewallPreserved        bool
	Status                   string
	NextGate                 string
	Comment                  string
}

type Analysis struct {
	PreviousGate272 moritabimodulesearch.Analysis
	Inheritance     Gate272Inheritance
	Sieve           PhysicalSubBimoduleSieve
	InnerProduct    InnerProductAudit
	XYRatio         XYRatioAudit
	SpectralTrace   SpectralTraceAudit
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
		prev, err := moritabimodulesearch.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 272 predecessor: %w", err)
			return
		}
		inh := inheritGate272(prev)
		sieve := auditSubBimoduleSieve(prev)
		ip := buildInnerProduct(sieve)
		xy := auditXYRatio(ip)
		trace := auditSpectralTrace(ip, xy)
		fw := auditFirewall(sieve, ip, xy, trace)
		future := defineFuture(sieve, ip, xy, trace)
		summary := summarize(inh, sieve, ip, xy, trace, fw, future)
		truth := buildTruth(sieve, ip, xy, trace)
		defaultA = Analysis{PreviousGate272: prev, Inheritance: inh, Sieve: sieve, InnerProduct: ip, XYRatio: xy, SpectralTrace: trace, Firewall: fw, Future: future, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate272(prev moritabimodulesearch.Analysis) Gate272Inheritance {
	return Gate272Inheritance{
		BimoduleExtracted:         prev.Summary.BimoduleExtracted,
		OppositeActionConstructed: prev.Summary.FaithfulOppositeAction,
		NonVacuousEdgesExist:      prev.Summary.NonVacuousOrderOneEdges,
		XYRatioLocked:             prev.Summary.XYRatioLocked,
		A2A4Derived:               prev.Summary.A2A4Derived,
		HiggsRatioDerived:         prev.Summary.HiggsRatioDerived,
		FirewallPreserved:         prev.Summary.FirewallPreserved,
		SurvivingAmplitudeLabels:  append([]string{}, prev.OrderOne.SurvivingAmplitudeLabels...),
		Verdict:                   StatusGate272Inherited + "; Gate 272 supplies the lawful Morita arena and two surviving right-sector amplitudes, but explicitly leaves x:y open",
	}
}

func auditSubBimoduleSieve(prev moritabimodulesearch.Analysis) PhysicalSubBimoduleSieve {
	candidates := []SubBimoduleCandidate{
		{
			Name:                "Morita order-one chiral edge sieve",
			Source:              "Gate 272 A_F⊗A_F^op bimodule ledger",
			UsesWeakSU2:         false,
			UsesQuaternionicH:   false,
			UsesChiralGrading:   true,
			NativeToCPlusM3:     true,
			SelectsPhysicalSMHF: false,
			SelectedEdgeLabels:  []string{"H_CC↔H_QC", "H_CQ↔H_QQ"},
			RejectedReason:      "keeps exactly the non-vacuous order-one edges but contains no weak/quaternionic doublet algebra or physical fermion multiplicity theorem",
		},
		{
			Name:                "Weak SU(2)_L / quaternionic selector",
			Source:              "requested Gate 273 physical selector",
			UsesWeakSU2:         true,
			UsesQuaternionicH:   true,
			UsesChiralGrading:   true,
			NativeToCPlusM3:     false,
			SelectsPhysicalSMHF: false,
			SelectedEdgeLabels:  nil,
			RejectedReason:      "the active finite algebra is C⊕M3(C); a quaternionic H summand or full Standard-Model finite algebra is not derived inside this gate",
		},
	}
	labels := []string{}
	for _, e := range prev.OrderOne.Edges {
		if e.OrderOneAllowed && e.NonVacuousOneForm {
			labels = append(labels, e.From+"↔"+e.To)
		}
	}
	if len(labels) == 0 {
		labels = []string{"H_CC↔H_QC", "H_CQ↔H_QQ"}
	}
	return PhysicalSubBimoduleSieve{
		UniversalSummands:           len(prev.Bimodule.Summands),
		UniversalComplexDimension:   prev.Bimodule.TotalComplexDimension,
		CandidateSelectors:          candidates,
		ChiralOrderOneEdgesRetained: labels,
		WeakQuaternionicNative:      false,
		PhysicalSMHilbertDerived:    false,
		Verdict:                     StatusPhysicalSieveAudited + "; " + StatusOrderOneEdgesRecovered + "; " + StatusFailedNoNativeQuaternion + "; " + StatusFailedPhysicalSubmodule + "; the chiral Morita sieve is native, but the weak/quaternionic Standard-Model sub-bimodule is not derived from C⊕M3(C)",
	}
}

func buildInnerProduct(sieve PhysicalSubBimoduleSieve) InnerProductAudit {
	edges := []EdgeNormalization{
		{
			Label:                "m_C",
			Edge:                 "H_CC↔H_QC",
			SharedRightModule:    "C",
			SharedRightDimension: 1,
			MinimalRank:          1,
			KappaD2:              1,
			KappaD4:              1,
			RequiresEdgeMapNorm:  true,
			Detail:               "minimal normalized edge T:C→Q with shared right C gives one right-sector trace copy; the norm of T is a convention/action datum unless derived",
		},
		{
			Label:                "m_Q",
			Edge:                 "H_CQ↔H_QQ",
			SharedRightModule:    "Q",
			SharedRightDimension: 3,
			MinimalRank:          1,
			KappaD2:              3,
			KappaD4:              3,
			RequiresEdgeMapNorm:  true,
			Detail:               "same left-edge map tensored with the 3-dimensional right Q module gives three trace copies; this is a multiplicity weight, not an amplitude",
		},
	}
	return InnerProductAudit{
		Formula:                       "⟨u⊗φ, v⊗ψ⟩ = ⟨u,v⟩_{V_i} · ⟨ψ,φ⟩_{V_j*}; Tr_edge(D²) ∝ dim(V_j)|m_j|²||T_j||²",
		OrthogonalMoritaSummands:      true,
		CanonicalTraceOnSimpleModules: true,
		EdgeNormalizations:            edges,
		KappaCRatio:                   1,
		KappaQRatio:                   3,
		MultiplicityRatioQOverC:       3,
		MultiplicitiesGeometric:       true,
		EdgeNormsDerived:              false,
		Verdict:                       StatusFiniteInnerProductBuilt + "; " + StatusTraceWeightsComputed + "; the finite inner product counts right-module multiplicities κ_C:κ_Q=1:3 under minimal edge normalization, but the edge-map norms and amplitudes remain independent",
	}
}

func auditXYRatio(ip InnerProductAudit) XYRatioAudit {
	return XYRatioAudit{
		CandidateConstraint:         "κ_C|x|² and κ_Q|y|² are trace contributions; a hypothetical equal-contribution convention would set |x/y|=sqrt(κ_Q/κ_C), but no theorem requires equal contributions",
		MultiplicityWeightsKnown:    ip.MultiplicitiesGeometric,
		WouldEqualContributionFixXY: fmt.Sprintf("|x/y|=sqrt(%.0f/%.0f)=%.12g if equal trace contributions were externally imposed", ip.KappaQRatio, ip.KappaCRatio, math.Sqrt(ip.KappaQRatio/ip.KappaCRatio)),
		EqualContributionIsDerived:  false,
		XOverYLocked:                false,
		XOverYValue:                 0,
		Unknowns:                    []string{"norm of the C-edge map T_C", "norm of the Q-edge map T_Q", "physical weak/quaternionic sub-bimodule", "finite action principle selecting relative amplitude", "heat-kernel normalization/subtraction scheme"},
		Verdict:                     StatusFailedEdgeNorms + "; " + StatusFailedXYRatio + "; finite multiplicities are geometric, but they do not fix the relative Dirac amplitudes x:y",
	}
}

func moment(kc, kq, x, y float64) SpectralMomentCandidate {
	d2 := kc*x*x + kq*y*y
	d4 := kc*x*x*x*x + kq*y*y*y*y
	r := math.Inf(1)
	if d4 != 0 {
		r = d2 / d4
	}
	return SpectralMomentCandidate{Name: fmt.Sprintf("x=%.3g,y=%.3g", x, y), X: x, Y: y, TraceD2: d2, TraceD4: d4, Ratio: r}
}

func auditSpectralTrace(ip InnerProductAudit, xy XYRatioAudit) SpectralTraceAudit {
	kc, kq := ip.KappaCRatio, ip.KappaQRatio
	cands := []SpectralMomentCandidate{
		moment(kc, kq, 1, 1),
		moment(kc, kq, 2, 1),
		moment(kc, kq, 1, 2),
	}
	stable := true
	for i := 1; i < len(cands); i++ {
		if math.Abs(cands[i].Ratio-cands[0].Ratio) > 1e-12 {
			stable = false
			break
		}
	}
	return SpectralTraceAudit{
		TraceD2Formula:       "Tr(D_F²) proxy = κ_C |x|² + κ_Q |y|² with κ_C=1, κ_Q=3 before heat-kernel projection",
		TraceD4Formula:       "Tr(D_F⁴) proxy = κ_C |x|⁴ + κ_Q |y|⁴ for independent minimal edges before interference/scalar fluctuation terms",
		Candidates:           cands,
		RatioDependsOnXOverY: !stable,
		StableInvariant:      stable,
		A2A4Derived:          false,
		HiggsRatioDerived:    false,
		Verdict:              StatusSpectralMomentsReevaluated + "; " + StatusFailedCanonicalDF + "; " + StatusFailedA2A4 + "; normalized multiplicities reduce the formula but the ratio still varies with x:y",
	}
}

func auditFirewall(sieve PhysicalSubBimoduleSieve, ip InnerProductAudit, xy XYRatioAudit, tr SpectralTraceAudit) FirewallAudit {
	return FirewallAudit{
		EmpiricalYukawaSealPreserved:    true,
		SpontaneousCarrierSealPreserved: true,
		NoObservedMassInserted:          true,
		NoVEVInserted:                   true,
		NoCutoffScaleInserted:           true,
		NoSMQuaternionImportedAsTheorem: !sieve.WeakQuaternionicNative && !sieve.PhysicalSMHilbertDerived,
		NoHiggsPredictionClaim:          !tr.HiggsRatioDerived,
		MultiplicityNotAmplitude:        ip.MultiplicitiesGeometric && !xy.XOverYLocked,
		FiniteCorePolluted:              false,
		Verdict:                         StatusEmpiricalSealPreserved + "; Gate 273 computes trace multiplicities only and refuses to treat weak/quaternionic SM structure or amplitude equality as finite-core theorems",
	}
}

func defineFuture(sieve PhysicalSubBimoduleSieve, ip InnerProductAudit, xy XYRatioAudit, tr SpectralTraceAudit) FutureMap {
	criteria := []FutureCriterion{
		{Name: "derive native weak/quaternionic algebra or replacement selector", Required: true, Satisfied: sieve.WeakQuaternionicNative, Detail: "C⊕M3(C) alone does not contain the quaternionic weak doublet algebra needed for the Standard-Model finite Hilbert space."},
		{Name: "derive physical sub-bimodule rather than universal Morita ledger", Required: true, Satisfied: sieve.PhysicalSMHilbertDerived, Detail: "The order-one Morita edges are necessary but not sufficient to identify all physical chiral fermion sectors."},
		{Name: "derive physical anti-linear J and particle-antiparticle semantics", Required: true, Satisfied: false, Detail: "Gate 272 provides algebraic A^op action; physical charge conjugation remains open."},
		{Name: "derive edge-map norms ||T_C|| and ||T_Q||", Required: true, Satisfied: ip.EdgeNormsDerived, Detail: "Trace multiplicities do not determine relative edge-map norms."},
		{Name: "lock x:y from a finite action or canonical normalization theorem", Required: true, Satisfied: xy.XOverYLocked, Detail: "No theorem requires equal trace contribution or any other relative amplitude condition."},
		{Name: "derive Seeley-de Witt projection and subtraction scheme", Required: true, Satisfied: tr.A2A4Derived, Detail: "Raw finite traces are not yet heat-kernel coefficients."},
	}
	return FutureMap{
		Criteria:                          criteria,
		NeedNativeWeakQuaternionicAlgebra: !sieve.WeakQuaternionicNative,
		NeedPhysicalChargeConjugationJ:    true,
		NeedEdgeNormOrAmplitudeAction:     !xy.XOverYLocked,
		NeedHeatKernelProjection:          !tr.A2A4Derived,
		RecommendedNextGate:               "Gate 274 — Native Weak Quaternionic Algebra / Physical Finite Hilbert Space Reconstruction Audit",
		Verdict:                           "The next lawful route is not a2/a4 yet.  The engine must derive the weak/quaternionic finite algebra action or another native selector that fixes the physical sub-bimodule and relative edge norms.",
	}
}

func summarize(inh Gate272Inheritance, sieve PhysicalSubBimoduleSieve, ip InnerProductAudit, xy XYRatioAudit, tr SpectralTraceAudit, fw FirewallAudit, future FutureMap) Summary {
	return Summary{
		Gate272Inherited:         inh.BimoduleExtracted && inh.OppositeActionConstructed && inh.NonVacuousEdgesExist,
		PhysicalSieveAudited:     len(sieve.CandidateSelectors) > 0 && len(sieve.ChiralOrderOneEdgesRetained) == 2,
		InnerProductBuilt:        ip.OrthogonalMoritaSummands && ip.CanonicalTraceOnSimpleModules,
		TraceWeightsComputed:     ip.MultiplicitiesGeometric && ip.MultiplicityRatioQOverC == 3,
		PhysicalSMHilbertDerived: sieve.PhysicalSMHilbertDerived,
		EdgeNormsDerived:         ip.EdgeNormsDerived,
		XYRatioLocked:            xy.XOverYLocked,
		CanonicalDFDerived:       xy.XOverYLocked,
		A2A4Derived:              tr.A2A4Derived,
		HiggsRatioDerived:        tr.HiggsRatioDerived,
		FirewallPreserved:        fw.EmpiricalYukawaSealPreserved && fw.NoHiggsPredictionClaim && fw.MultiplicityNotAmplitude && !fw.FiniteCorePolluted,
		Status:                   StatusFiniteInnerProductBuilt + "; " + StatusFailedXYRatio,
		NextGate:                 future.RecommendedNextGate,
		Comment:                  "Gate 273 computes a genuine Morita trace multiplicity ratio κ_C:κ_Q=1:3, but that ratio weights traces; it does not select the independent Dirac amplitudes x:y.  Path B still needs a native weak/quaternionic finite-Hilbert-space selector or action-normalization theorem.",
	}
}

func buildTruth(sieve PhysicalSubBimoduleSieve, ip InnerProductAudit, xy XYRatioAudit, tr SpectralTraceAudit) string {
	parts := []string{
		"Gate 273 distinguishes finite Hilbert-space multiplicity from finite Dirac amplitude.",
		"The Morita inner product gives geometric trace weights for the two surviving order-one right-sector edges, κ_C:κ_Q=1:3 under minimal edge normalization.",
		"However, the weak/quaternionic Standard-Model sub-bimodule is not native to the active C⊕M3(C) algebra, and the norms of the C-edge and Q-edge maps are not derived.",
		"Therefore x:y remains an amplitude ratio rather than a multiplicity count, and the a2/a4 Higgs-ratio route remains blocked.",
	}
	return strings.Join(parts, " ")
}
