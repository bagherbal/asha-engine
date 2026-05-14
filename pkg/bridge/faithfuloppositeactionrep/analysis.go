// Package faithfuloppositeactionrep implements Gate 270:
// Faithful Opposite-Action Representation / Non-Vacuous One-Form Calculus Audit.
//
// Gate 269 proved that the mode-level order-one condition reduces the formal
// Dirac block to M=diag(x,yI3), but in that same-side representation the
// one-form calculus is vacuous. Gate 270 tests the next necessary ingredient:
// whether a faithful doubled-S_C representation and a physical opposite action
// through J are actually available. The gate also constructs a deliberately
// limited chiral-bimodule preflight to show the mechanism by which nonzero
// commutators can appear. That preflight is useful, but it is not promoted: it
// lives on doubled mode space W_L⊕W_R, not the full doubled S_C carrier, and its
// naive opposite action fails the full order-one residual for generic M3(C)
// probes. Thus non-vacuous one-forms are exposed as a target, while the Higgs
// ratio and canonical D_F remain blocked.
package faithfuloppositeactionrep

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/canonicalfinitediracselector"
)

const (
	AuditID = "GATE270-FAITHFUL-OPPOSITE-ACTION-NONVACUOUS-ONE-FORM-CALCULUS-AUDIT"

	StatusGate269Inherited           = "CONDITIONAL_SUPPORT_GATE269_ORDER_ONE_SIEVE_INHERITED"
	StatusFaithfulLiftAudited        = "CONDITIONAL_SUPPORT_FAITHFUL_SC_REPRESENTATION_LIFT_AUDITED"
	StatusChiralPreflightConstructed = "CONDITIONAL_SUPPORT_CHIRAL_BIMODULE_PREFLIGHT_CONSTRUCTED"
	StatusNonVacuousCandidateExposed = "CONDITIONAL_SUPPORT_CANDIDATE_NONVACUOUS_ONE_FORMS_EXPOSED"
	StatusOrderOneResidualComputed   = "CONDITIONAL_SUPPORT_FULL_ORDER_ONE_RESIDUAL_COMPUTED"
	StatusMomentsRechecked           = "CONDITIONAL_SUPPORT_ORDER_ONE_FAMILY_MOMENTS_RECHECKED"
	StatusFailedFullSCRep            = "FAILED_ROUTE_FAITHFUL_TOTAL_SC_REPRESENTATION_STILL_MISSING"
	StatusFailedPhysicalJ            = "FAILED_ROUTE_PHYSICAL_J_OPPOSITE_ACTION_STILL_MISSING"
	StatusFailedCandidateOrderOne    = "FAILED_ROUTE_CANDIDATE_CHIRAL_ACTION_FAILS_FULL_ORDER_ONE"
	StatusFailedCanonicalDF          = "FAILED_ROUTE_FAITHFUL_ACTION_DOES_NOT_SELECT_CANONICAL_DF"
	StatusFailedXYRatio              = "FAILED_ROUTE_XY_RATIO_STILL_UNCONSTRAINED"
	StatusFailedHiggsRatio           = "FAILED_ROUTE_INVARIANT_HIGGS_RATIO_NOT_DERIVED"
	StatusEmpiricalSealPreserved     = "FAILED_ROUTE_EMPIRICAL_YUKAWA_SEAL_REMAINS_ACTIVE"
)

type Gate269Inheritance struct {
	OrderOneDefined       bool
	ModeLevelSieveReduced bool
	AllowedFamilyFormula  string
	SurvivingFamilyDimC   int
	NonVacuousCalculus    bool
	CanonicalDFDerived    bool
	HiggsRatioDerived     bool
	FirewallPreserved     bool
	RecommendedNextGate   string
	Verdict               string
}

type FaithfulLiftAudit struct {
	TargetCarrier                 string
	TargetComplexDimension        int
	AvailableCarrier              string
	AvailableComplexDimension     int
	ModePreflightCarrier          string
	ModePreflightComplexDimension int
	FullSCRepresentationDerived   bool
	ChiralGradingRespected        bool
	AlgebraFaithfulOnCandidate    bool
	ImportedConnesRepresentation  bool
	Verdict                       string
}

type OppositeActionAudit struct {
	JFormula                      string
	AntiLinearJDerived            bool
	ParticleAntiparticleSemantics bool
	OppositeActionDerived         bool
	CandidateSwapActionAvailable  bool
	CandidateIsPhysical           bool
	Verdict                       string
}

type ChiralActionPreflight struct {
	Carrier                string
	Algebra                string
	LeftAction             string
	RightAction            string
	CenterCharacter        string
	DiracBlock             string
	LeftRightActionsDiffer bool
	NonVacuousPossible     bool
	FullSCPhysical         bool
	Verdict                string
}

type Probe struct {
	Name      string
	Lambda    float64
	BDiag     [3]float64
	Character float64
}

type OneFormAudit struct {
	ProbeA               Probe
	X                    float64
	Y                    float64
	SpatialOneFormDiag   [3]float64
	FrobeniusNormSq      float64
	NonZero              bool
	CentralProbeVanishes bool
	PhysicalOneForm      bool
	Verdict              string
}

type OrderOneResidualAudit struct {
	ProbeA             Probe
	ProbeB             Probe
	ResidualDiag       [3]float64
	FrobeniusNormSq    float64
	ResidualZero       bool
	CandidatePasses    bool
	FullOrderOneProved bool
	Verdict            string
}

type MomentRow struct {
	Name    string
	X       float64
	Y       float64
	TraceD2 float64
	TraceD4 float64
	Ratio   float64
	Comment string
}

type InvariantRatioAudit struct {
	Rows                        []MomentRow
	FamilyFormula               string
	XToYSelected                bool
	RatioStableAcrossFamily     bool
	DependsOnXY                 bool
	GaugeProjectionDerived      bool
	ScalarFluctuationMapDerived bool
	HiggsRatioDerived           bool
	Verdict                     string
}

type FirewallAudit struct {
	EmpiricalYukawaSealPreserved    bool
	SpontaneousCarrierSealPreserved bool
	NoObservedMassInserted          bool
	NoVEVInserted                   bool
	NoCutoffScaleInserted           bool
	NoConnesRepresentationImported  bool
	CandidateNotPromoted            bool
	NoHiggsPredictionClaim          bool
	FiniteCorePolluted              bool
	Verdict                         string
}

type FutureObligation struct {
	Name      string
	Required  bool
	Satisfied bool
	Detail    string
}

type FutureMap struct {
	Obligations              []FutureObligation
	NeedFullSCRepresentation bool
	NeedPhysicalJ            bool
	NeedOrderOnePassingBimod bool
	NeedNonVacuousCalculus   bool
	NeedCanonicalXYSelector  bool
	NeedHeatKernelProjection bool
	RecommendedNextGate      string
	Verdict                  string
}

type Summary struct {
	Gate269Inherited         bool
	FaithfulLiftAudited      bool
	FullSCRepresentation     bool
	PhysicalOppositeAction   bool
	CandidateOneFormsNonzero bool
	CandidateOrderOnePasses  bool
	CanonicalDFDerived       bool
	RatioStable              bool
	HiggsRatioDerived        bool
	FirewallPreserved        bool
	Status                   string
	NextGate                 string
	Comment                  string
}

type Analysis struct {
	PreviousGate269 canonicalfinitediracselector.Analysis
	Inheritance     Gate269Inheritance
	Lift            FaithfulLiftAudit
	Opposite        OppositeActionAudit
	Chiral          ChiralActionPreflight
	OneForm         OneFormAudit
	Residual        OrderOneResidualAudit
	Ratio           InvariantRatioAudit
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
		prev, err := canonicalfinitediracselector.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 269 predecessor: %w", err)
			return
		}
		inh := inheritGate269(prev)
		lift := auditFaithfulLift()
		opp := auditOppositeAction(lift)
		chiral := constructChiralPreflight(lift, opp)
		one := computeOneFormPreflight(chiral)
		resid := computeOrderOneResidual(chiral, one)
		ratio := recheckInvariantRatio(prev)
		fw := auditFirewall(lift, opp, chiral, one, resid, ratio)
		future := defineFutureMap(lift, opp, chiral, one, resid, ratio)
		summary := summarize(inh, lift, opp, chiral, one, resid, ratio, fw, future)
		truth := buildTruth(inh, lift, opp, one, resid, ratio, fw)
		defaultA = Analysis{PreviousGate269: prev, Inheritance: inh, Lift: lift, Opposite: opp, Chiral: chiral, OneForm: one, Residual: resid, Ratio: ratio, Firewall: fw, Future: future, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate269(prev canonicalfinitediracselector.Analysis) Gate269Inheritance {
	return Gate269Inheritance{
		OrderOneDefined:       prev.Summary.OrderOneDefined,
		ModeLevelSieveReduced: prev.Summary.OrderOneSieveReduced,
		AllowedFamilyFormula:  prev.Sieve.AllowedFamilyFormula,
		SurvivingFamilyDimC:   prev.Canonical.SurvivingFamilyDimensionC,
		NonVacuousCalculus:    !prev.Sieve.OneFormsVanishForAllowedFamily,
		CanonicalDFDerived:    prev.Summary.CanonicalDFDerived,
		HiggsRatioDerived:     prev.Summary.HiggsRatioDerived,
		FirewallPreserved:     prev.Summary.FirewallPreserved,
		RecommendedNextGate:   prev.Future.RecommendedNextGate,
		Verdict:               StatusGate269Inherited + "; Gate 269 supplies M=diag(x,yI3) as an order-one preflight family but not a physical spectral triple",
	}
}

func auditFaithfulLift() FaithfulLiftAudit {
	return FaithfulLiftAudit{
		TargetCarrier:                 "doubled S_C = S_C ⊕ S_C* with S_C=Λ*(C^4)",
		TargetComplexDimension:        32,
		AvailableCarrier:              "S_C kinematic carrier plus mode-level C⊕M3(C) algebra snapshots",
		AvailableComplexDimension:     16,
		ModePreflightCarrier:          "doubled mode carrier W_L ⊕ W_R, W=C⊕C^3",
		ModePreflightComplexDimension: 8,
		FullSCRepresentationDerived:   false,
		ChiralGradingRespected:        true,
		AlgebraFaithfulOnCandidate:    true,
		ImportedConnesRepresentation:  false,
		Verdict:                       StatusFaithfulLiftAudited + "; a small chiral mode-bimodule can be audited, but the faithful action on the full 32-complex-dimensional doubled S_C carrier is still missing",
	}
}

func auditOppositeAction(lift FaithfulLiftAudit) OppositeActionAudit {
	return OppositeActionAudit{
		JFormula:                      "candidate J_mode swaps W_L and W_R with complex conjugation; physical J on doubled S_C not derived",
		AntiLinearJDerived:            false,
		ParticleAntiparticleSemantics: false,
		OppositeActionDerived:         false,
		CandidateSwapActionAvailable:  lift.ModePreflightComplexDimension == 8,
		CandidateIsPhysical:           false,
		Verdict:                       StatusFailedPhysicalJ + "; a swap-conjugation candidate can be tested, but it lacks derived particle/antiparticle semantics and a full opposite representation",
	}
}

func constructChiralPreflight(lift FaithfulLiftAudit, opp OppositeActionAudit) ChiralActionPreflight {
	return ChiralActionPreflight{
		Carrier:                lift.ModePreflightCarrier,
		Algebra:                "A_F=C⊕M3(C)",
		LeftAction:             "rho_L(λ,B)=diag(λ,B)",
		RightAction:            "rho_R(λ,B)=diag(λ,χ(B)I3), χ(B)=Tr(B)/3",
		CenterCharacter:        "χ:M3(C)→C is a diagnostic center/trace character, not a finite-core theorem",
		DiracBlock:             "M_order1(x,y)=diag(x,y,y,y)",
		LeftRightActionsDiffer: true,
		NonVacuousPossible:     true,
		FullSCPhysical:         false,
		Verdict:                StatusChiralPreflightConstructed + "; the candidate separates left and right actions enough to test nonzero one-forms, but remains a mode-level diagnostic",
	}
}

func computeOneFormPreflight(chiral ChiralActionPreflight) OneFormAudit {
	probeA := Probe{Name: "traceless color probe a", Lambda: 0, BDiag: [3]float64{1, -1, 0}}
	probeA.Character = character(probeA.BDiag)
	x, y := 1.0, 1.0
	p := oneFormDiag(probeA, y)
	norm := normSq(p)
	central := Probe{Name: "central color probe", Lambda: 0, BDiag: [3]float64{1, 1, 1}, Character: 1}
	centralP := oneFormDiag(central, y)
	return OneFormAudit{
		ProbeA:               probeA,
		X:                    x,
		Y:                    y,
		SpatialOneFormDiag:   p,
		FrobeniusNormSq:      norm,
		NonZero:              norm > 0,
		CentralProbeVanishes: normSq(centralP) == 0,
		PhysicalOneForm:      false,
		Verdict:              StatusNonVacuousCandidateExposed + "; the chiral diagnostic gives ||Mρ_R(a)-ρ_L(a)M||²=2 for a traceless color probe, proving how non-vacuous one-forms can arise before physical promotion",
	}
}

func computeOrderOneResidual(chiral ChiralActionPreflight, one OneFormAudit) OrderOneResidualAudit {
	probeB := Probe{Name: "second traceless color probe b", Lambda: 0, BDiag: [3]float64{1, 0, -1}}
	probeB.Character = character(probeB.BDiag)
	res := residualDiag(one.SpatialOneFormDiag, probeB)
	norm := normSq(res)
	return OrderOneResidualAudit{
		ProbeA:             one.ProbeA,
		ProbeB:             probeB,
		ResidualDiag:       res,
		FrobeniusNormSq:    norm,
		ResidualZero:       norm == 0,
		CandidatePasses:    norm == 0,
		FullOrderOneProved: false,
		Verdict:            StatusOrderOneResidualComputed + "; " + StatusFailedCandidateOrderOne + "; the naive chiral mismatch creates one-forms but fails [[D,a],Jb*J^-1]=0 for generic M3(C) probes",
	}
}

func recheckInvariantRatio(prev canonicalfinitediracselector.Analysis) InvariantRatioAudit {
	rows := []MomentRow{
		momentRow("unit order-one family", 1, 1, "inherited allowed representative"),
		momentRow("lepton-weight family", 2, 1, "still allowed unless a new selector fixes x:y"),
		momentRow("color-weight family", 1, 2, "still allowed unless a new selector fixes x:y"),
	}
	stable := true
	for i := 1; i < len(rows); i++ {
		if !approx(rows[0].Ratio, rows[i].Ratio, 1e-12) {
			stable = false
		}
	}
	return InvariantRatioAudit{
		Rows:                        rows,
		FamilyFormula:               prev.Sieve.AllowedFamilyFormula,
		XToYSelected:                false,
		RatioStableAcrossFamily:     stable,
		DependsOnXY:                 !stable,
		GaugeProjectionDerived:      false,
		ScalarFluctuationMapDerived: false,
		HiggsRatioDerived:           false,
		Verdict:                     StatusMomentsRechecked + "; " + StatusFailedXYRatio + "; non-vacuous candidate one-forms do not by themselves select x:y or stabilize Tr(D²)/Tr(D⁴)",
	}
}

func auditFirewall(lift FaithfulLiftAudit, opp OppositeActionAudit, chiral ChiralActionPreflight, one OneFormAudit, resid OrderOneResidualAudit, ratio InvariantRatioAudit) FirewallAudit {
	return FirewallAudit{
		EmpiricalYukawaSealPreserved:    true,
		SpontaneousCarrierSealPreserved: true,
		NoObservedMassInserted:          true,
		NoVEVInserted:                   true,
		NoCutoffScaleInserted:           true,
		NoConnesRepresentationImported:  !lift.ImportedConnesRepresentation,
		CandidateNotPromoted:            !lift.FullSCRepresentationDerived && !opp.CandidateIsPhysical && !chiral.FullSCPhysical && !one.PhysicalOneForm && !resid.FullOrderOneProved,
		NoHiggsPredictionClaim:          !ratio.HiggsRatioDerived,
		FiniteCorePolluted:              false,
		Verdict:                         StatusEmpiricalSealPreserved + "; the chiral-bimodule one-form calculation is retained as a diagnostic target and not promoted into a finite-core spectral triple",
	}
}

func defineFutureMap(lift FaithfulLiftAudit, opp OppositeActionAudit, chiral ChiralActionPreflight, one OneFormAudit, resid OrderOneResidualAudit, ratio InvariantRatioAudit) FutureMap {
	obligations := []FutureObligation{
		{Name: "faithful C⊕M3(C) action on doubled S_C", Required: true, Satisfied: lift.FullSCRepresentationDerived, Detail: "The current diagnostic acts on W_L⊕W_R, not the full S_C⊕S_C* carrier."},
		{Name: "anti-linear physical J and opposite algebra action", Required: true, Satisfied: opp.AntiLinearJDerived && opp.OppositeActionDerived && opp.ParticleAntiparticleSemantics, Detail: "J must implement the derived opposite representation, not only a swap-conjugation placeholder."},
		{Name: "nonzero one-forms that also satisfy order-one", Required: true, Satisfied: one.NonZero && resid.CandidatePasses && resid.FullOrderOneProved, Detail: "Gate 270 exposes nonzero candidate one-forms, but the candidate fails the double-commutator residual."},
		{Name: "canonical x:y amplitude selector", Required: true, Satisfied: ratio.XToYSelected, Detail: "Even a non-vacuous calculus must still fix the surviving lepton/quark weight ratio."},
		{Name: "gauge/scalar fluctuation projection", Required: true, Satisfied: ratio.GaugeProjectionDerived && ratio.ScalarFluctuationMapDerived, Detail: "Higgs and gauge terms require projection maps before Seeley-de Witt coefficients can be interpreted."},
		{Name: "heat-kernel/cutoff moment normalization", Required: true, Satisfied: false, Detail: "Raw traces are not yet spectral-action coefficients without normalization and subtraction data."},
	}
	return FutureMap{
		Obligations:              obligations,
		NeedFullSCRepresentation: true,
		NeedPhysicalJ:            true,
		NeedOrderOnePassingBimod: true,
		NeedNonVacuousCalculus:   true,
		NeedCanonicalXYSelector:  true,
		NeedHeatKernelProjection: true,
		RecommendedNextGate:      "Gate 271 — Full S_C Finite Algebra Representation Search / Opposite-Action Construction Audit",
		Verdict:                  "Gate 270 identifies the missing representation theorem: derive a true bimodule/opposite action on doubled S_C where one-forms are nonzero and order-one holds simultaneously.",
	}
}

func summarize(inh Gate269Inheritance, lift FaithfulLiftAudit, opp OppositeActionAudit, chiral ChiralActionPreflight, one OneFormAudit, resid OrderOneResidualAudit, ratio InvariantRatioAudit, fw FirewallAudit, future FutureMap) Summary {
	status := strings.Join([]string{
		StatusGate269Inherited,
		StatusFaithfulLiftAudited,
		StatusChiralPreflightConstructed,
		StatusNonVacuousCandidateExposed,
		StatusOrderOneResidualComputed,
		StatusMomentsRechecked,
		StatusFailedFullSCRep,
		StatusFailedPhysicalJ,
		StatusFailedCandidateOrderOne,
		StatusFailedCanonicalDF,
		StatusFailedXYRatio,
		StatusFailedHiggsRatio,
		StatusEmpiricalSealPreserved,
	}, "; ")
	return Summary{
		Gate269Inherited:         inh.OrderOneDefined && inh.ModeLevelSieveReduced,
		FaithfulLiftAudited:      true,
		FullSCRepresentation:     lift.FullSCRepresentationDerived,
		PhysicalOppositeAction:   opp.OppositeActionDerived && opp.AntiLinearJDerived,
		CandidateOneFormsNonzero: one.NonZero,
		CandidateOrderOnePasses:  resid.CandidatePasses,
		CanonicalDFDerived:       false,
		RatioStable:              ratio.RatioStableAcrossFamily,
		HiggsRatioDerived:        ratio.HiggsRatioDerived,
		FirewallPreserved:        fw.EmpiricalYukawaSealPreserved && fw.CandidateNotPromoted && !fw.FiniteCorePolluted,
		Status:                   status,
		NextGate:                 future.RecommendedNextGate,
		Comment:                  "Gate 270 exposes a non-vacuous chiral one-form mechanism on a small diagnostic carrier, but the same candidate fails order-one and is not a faithful doubled-S_C representation; no x:y selector or Higgs ratio follows.",
	}
}

func buildTruth(inh Gate269Inheritance, lift FaithfulLiftAudit, opp OppositeActionAudit, one OneFormAudit, resid OrderOneResidualAudit, ratio InvariantRatioAudit, fw FirewallAudit) string {
	return fmt.Sprintf("Gate 270 truth: inherited269=%t targetDim=%d candidateDim=%d fullSC=%t physicalJ=%t oneFormNormSq=%.6g residualNormSq=%.6g orderOnePass=%t ratioStable=%t higgs=%t firewall=%t", inh.ModeLevelSieveReduced, lift.TargetComplexDimension, lift.ModePreflightComplexDimension, lift.FullSCRepresentationDerived, opp.OppositeActionDerived, one.FrobeniusNormSq, resid.FrobeniusNormSq, resid.CandidatePasses, ratio.RatioStableAcrossFamily, ratio.HiggsRatioDerived, fw.CandidateNotPromoted && !fw.FiniteCorePolluted)
}

func character(d [3]float64) float64 { return (d[0] + d[1] + d[2]) / 3.0 }

func oneFormDiag(a Probe, y float64) [3]float64 {
	chi := a.Character
	return [3]float64{y * (chi - a.BDiag[0]), y * (chi - a.BDiag[1]), y * (chi - a.BDiag[2])}
}

func residualDiag(p [3]float64, b Probe) [3]float64 {
	// Residual for the naive swap-opposite diagnostic:
	// R = P·B_b - χ(B_b) P on the spatial block.
	chi := b.Character
	return [3]float64{p[0] * (b.BDiag[0] - chi), p[1] * (b.BDiag[1] - chi), p[2] * (b.BDiag[2] - chi)}
}

func normSq(d [3]float64) float64 { return d[0]*d[0] + d[1]*d[1] + d[2]*d[2] }

func momentRow(name string, x, y float64, comment string) MomentRow {
	tr2 := 2 * (x*x + 3*y*y)
	tr4 := 2 * (x*x*x*x + 3*y*y*y*y)
	ratio := math.Inf(1)
	if tr4 != 0 {
		ratio = tr2 / tr4
	}
	return MomentRow{Name: name, X: x, Y: y, TraceD2: tr2, TraceD4: tr4, Ratio: ratio, Comment: comment}
}

func approx(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
