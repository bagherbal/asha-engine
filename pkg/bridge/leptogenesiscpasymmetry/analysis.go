// Package leptogenesiscpasymmetry implements Gate 354:
// Leptogenesis Decay & CP-Asymmetry / B-Gap Majorana Cosmogenesis Audit.
//
// Gate 353 showed that RG-time attractors alone do not reduce the 15 minimal
// vacuum coordinates. Gate 354 audits a different time-selection mechanism:
// early-universe Majorana decays. It asks whether the B-gap/seesaw sector has
// enough native topological capacity to generate the observed baryon asymmetry
// through leptogenesis, and whether that capacity actually derives a CP-odd
// operator or only defines a target/witness.
package leptogenesiscpasymmetry

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE354-LEPTOGENESIS-DECAY-CP-ASYMMETRY-BGAP-MAJORANA-COSMOGENESIS-AUDIT"

	StatusDecayFormalized            = "CONDITIONAL_SUPPORT_MAJORANA_DECAY_CHANNEL_FORMALIZED"
	StatusSakharovLedgerFormalized   = "CONDITIONAL_SUPPORT_SAKHAROV_LEDGER_FORMALIZED"
	StatusCPAsymmetryTargetExtracted = "CONDITIONAL_SUPPORT_CP_ASYMMETRY_TARGET_EXTRACTED"
	StatusTopologyCapacityAudited    = "CONDITIONAL_SUPPORT_BGAP_TOPOLOGICAL_CP_CAPACITY_AUDITED"
	StatusLeptogenesisSieveExecuted  = "CONDITIONAL_SUPPORT_LEPTOGENESIS_VIABILITY_SIEVE_EXECUTED"
	StatusBaryogenesisFormalized     = "CONDITIONAL_SUPPORT_BARYOGENESIS_CONSTRAINT_FORMALIZED"
	StatusCensusUpdated              = "CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED"

	StatusTensionCapacityNearTarget      = "CONDITIONAL_TENSION_BGAP_INSTANTON_OVERLAP_CAPACITY_NEAR_LEPTOGENESIS_TARGET"
	StatusTensionEfficiencyRequired      = "CONDITIONAL_TENSION_WASHOUT_EFFICIENCY_REQUIRED_BUT_NOT_DERIVED"
	StatusTensionCPPhaseRequired         = "CONDITIONAL_TENSION_CP_ODD_PHASE_REQUIRED_BUT_NOT_DERIVED"
	StatusTensionCKMShadowNotEstablished = "CONDITIONAL_TENSION_CKM_SHADOW_OF_MAJORANA_CP_NOT_ESTABLISHED"
	StatusTensionNoParameterReduction    = "CONDITIONAL_TENSION_NO_VACUUM_PARAMETER_REDUCTION_PROVED"

	StatusFailedCPAsymmetryOperatorNotDerived = "FAILED_ROUTE_CP_ASYMMETRY_OPERATOR_NOT_DERIVED"
	StatusFailedMajoranaPhaseNotDerived       = "FAILED_ROUTE_MAJORANA_CP_PHASE_NOT_DERIVED"
	StatusFailedWashoutNotDerived             = "FAILED_ROUTE_LEPTOGENESIS_EFFICIENCY_WASHOUT_NOT_DERIVED"
	StatusFailedBoltzmannSystemNotSolved      = "FAILED_ROUTE_BOLTZMANN_TRANSPORT_NOT_EXECUTED"
	StatusFailedCKMPhaseNotDerived            = "FAILED_ROUTE_CKM_PHASE_NOT_DERIVED_FROM_LEPTOGENESIS"
	StatusFailedNoParameterReduction          = "FAILED_ROUTE_NO_ADDITIONAL_PARAMETER_REDUCTION_PROVED"
	StatusFailedSevenCoordinatesNotProved     = "FAILED_ROUTE_SEVEN_VACUUM_COORDINATES_NOT_PROVED"
)

const (
	inheritedGate        = 353
	startingVacuumInputs = 15
	observedEtaB         = 6.12e-10
	sphaleronFactor      = 28.0 / 79.0
	entropyPhotonFactor  = 7.04
	gStarSM              = 106.75
	bGap                 = 0.102464921191
	kappaQ               = 3.0
	topologicalResonance = 4.0 / math.Pi
)

type Span struct {
	AuditID       string
	InheritedGate int
	AddsFit       bool
	Purpose       string
	Verdict       string
}

type DecayChannel struct {
	Formalized      bool
	HeavyState      string
	DecayMatter     string
	DecayAntimatter string
	CPAsymmetry     string
	MajoranaEdge    string
	BGap            float64
	Verdict         string
}

type SakharovLedger struct {
	Formalized              bool
	BaryonViolation         string
	CPViolation             string
	OutOfEquilibrium        string
	ASHAHasMajoranaCapacity bool
	ASHAHasCPPhaseOperator  bool
	ASHASolvesBoltzmann     bool
	Verdict                 string
}

type TargetAudit struct {
	ObservedEtaB       float64
	SphaleronFactor    float64
	EntropyPhoton      float64
	GStar              float64
	ConversionFactor   float64
	RequiredEpsKappa   float64
	RequiredEpsKappa1  float64
	RequiredEpsKappa01 float64
	RequiredEpsKappa02 float64
	Verdict            string
}

type TopologicalCapacity struct {
	BGap                         float64
	KappaQ                       float64
	FourOverPi                   float64
	SInst                        float64
	ExpMinusSInst                float64
	PortalOverlap                float64
	InstantonOverlapEps          float64
	RequiredEfficiencyAtMaxPhase float64
	ViableEfficiencyWindow       bool
	Verdict                      string
}

type StandardLeptogenesis struct {
	Formalized           bool
	Formula              string
	Requires             []string
	DerivedCPInvariant   bool
	DerivedHeavySpectrum bool
	DerivedEfficiency    bool
	Verdict              string
}

type CKMShadowAudit struct {
	Formalized                        bool
	Hypothesis                        string
	MajoranaSectorCanSourceLeptonicCP bool
	QuarkCKMBridgeDerived             bool
	ConsumesCKMPhase                  bool
	ConsumesPMNSPhase                 bool
	ParameterReduction                int
	ReductionProved                   bool
	Verdict                           string
}

type Census struct {
	StartingVacuumInputs  int
	LeptogenesisReduction int
	RemainingInputs       int
	SevenSealTarget       int
	SevenSealReached      bool
	Verdict               string
}

type Summary struct {
	Executed                 bool
	HasTopologicalCapacity   bool
	CPAsymmetryDerived       bool
	BaryonAsymmetryPredicted bool
	AnyReductionProved       bool
	RemainingInputs          int
	Status                   string
	DirectAnswer             string
	NextGate                 string
}

type Analysis struct {
	Span         Span
	Decay        DecayChannel
	Sakharov     SakharovLedger
	Target       TargetAudit
	Capacity     TopologicalCapacity
	Leptogenesis StandardLeptogenesis
	CKMShadow    CKMShadowAudit
	Census       Census
	Summary      Summary
	Truth        string
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
	span := compileSpan()
	decay := formalizeDecay()
	sakharov := formalizeSakharov()
	target := extractTarget()
	capacity := auditCapacity(target)
	lepto := formalizeStandardLeptogenesis()
	ckm := auditCKMShadow()
	census := updateCensus(ckm)
	summary := compileSummary(capacity, lepto, ckm, census)
	truth := "Gate 354 finds a serious topological capacity witness for leptogenesis: kappa_Q*(4/pi)*B_gap*exp(-S_inst) is of the right CP-asymmetry scale if the washout efficiency is about 1.7% and the CP-odd phase is maximal.  This is a capacity result, not a derivation.  The finite core still lacks the CP-odd invariant, heavy-neutrino hierarchy, Boltzmann/washout solution, and bridge from Majorana CP to CKM or PMNS phases.  Therefore baryogenesis is formalized as a Phase-III dynamical constraint but no vacuum parameter is removed."
	return Analysis{Span: span, Decay: decay, Sakharov: sakharov, Target: target, Capacity: capacity, Leptogenesis: lepto, CKMShadow: ckm, Census: census, Summary: summary, Truth: truth}, nil
}

func compileSpan() Span {
	return Span{AuditID: AuditID, InheritedGate: inheritedGate, AddsFit: false, Purpose: "audit whether B-gap Majorana decays can dynamically select CP phases through leptogenesis", Verdict: StatusBaryogenesisFormalized}
}

func formalizeDecay() DecayChannel {
	return DecayChannel{Formalized: true, HeavyState: "N1 / B-gap Majorana carrier", DecayMatter: "N1 -> H + L", DecayAntimatter: "N1 -> H* + Lbar", CPAsymmetry: "epsilon1 = [Gamma(N1->HL)-Gamma(N1->H*Lbar)]/[Gamma(N1->HL)+Gamma(N1->H*Lbar)]", MajoranaEdge: "nu_R <-> nu_R^c via J_swap / B_gap", BGap: bGap, Verdict: StatusDecayFormalized}
}

func formalizeSakharov() SakharovLedger {
	return SakharovLedger{Formalized: true, BaryonViolation: "electroweak sphalerons convert B-L into B", CPViolation: "requires Im[(Y_N^dagger Y_N)_{1j}^2] or equivalent CP-odd invariant", OutOfEquilibrium: "requires Gamma_N1 < H(T=M1) and Boltzmann washout efficiency", ASHAHasMajoranaCapacity: true, ASHAHasCPPhaseOperator: false, ASHASolvesBoltzmann: false, Verdict: StatusSakharovLedgerFormalized}
}

func extractTarget() TargetAudit {
	conversion := entropyPhotonFactor * sphaleronFactor / gStarSM
	req := observedEtaB / conversion
	return TargetAudit{ObservedEtaB: observedEtaB, SphaleronFactor: sphaleronFactor, EntropyPhoton: entropyPhotonFactor, GStar: gStarSM, ConversionFactor: conversion, RequiredEpsKappa: req, RequiredEpsKappa1: req, RequiredEpsKappa01: req / 0.1, RequiredEpsKappa02: req / 0.02, Verdict: StatusCPAsymmetryTargetExtracted}
}

func auditCapacity(target TargetAudit) TopologicalCapacity {
	sInst := topologicalResonance / bGap
	expS := math.Exp(-sInst)
	portal := kappaQ * topologicalResonance * bGap
	epsWitness := portal * expS
	eff := target.RequiredEpsKappa / epsWitness
	return TopologicalCapacity{BGap: bGap, KappaQ: kappaQ, FourOverPi: topologicalResonance, SInst: sInst, ExpMinusSInst: expS, PortalOverlap: portal, InstantonOverlapEps: epsWitness, RequiredEfficiencyAtMaxPhase: eff, ViableEfficiencyWindow: eff > 0.001 && eff < 0.1, Verdict: StatusTopologyCapacityAudited}
}

func formalizeStandardLeptogenesis() StandardLeptogenesis {
	return StandardLeptogenesis{Formalized: true, Formula: "epsilon_1 = [1/(8*pi*(Y_N^dagger Y_N)11)] sum_{j!=1} Im[(Y_N^dagger Y_N)1j^2] F(M_j^2/M_1^2)", Requires: []string{"complex heavy-neutrino Yukawa matrix", "at least two heavy Majorana states", "mass hierarchy / loop function", "washout efficiency", "Boltzmann transport"}, DerivedCPInvariant: false, DerivedHeavySpectrum: false, DerivedEfficiency: false, Verdict: StatusFailedCPAsymmetryOperatorNotDerived}
}

func auditCKMShadow() CKMShadowAudit {
	return CKMShadowAudit{Formalized: true, Hypothesis: "low-energy CKM/PMNS phases may be shadows of high-scale Majorana CP violation", MajoranaSectorCanSourceLeptonicCP: true, QuarkCKMBridgeDerived: false, ConsumesCKMPhase: false, ConsumesPMNSPhase: false, ParameterReduction: 0, ReductionProved: false, Verdict: StatusTensionCKMShadowNotEstablished}
}

func updateCensus(ckm CKMShadowAudit) Census {
	remaining := startingVacuumInputs - ckm.ParameterReduction
	return Census{StartingVacuumInputs: startingVacuumInputs, LeptogenesisReduction: ckm.ParameterReduction, RemainingInputs: remaining, SevenSealTarget: 7, SevenSealReached: remaining <= 7, Verdict: StatusFailedNoParameterReduction}
}

func compileSummary(capacity TopologicalCapacity, lepto StandardLeptogenesis, ckm CKMShadowAudit, census Census) Summary {
	derived := lepto.DerivedCPInvariant && lepto.DerivedEfficiency && ckm.ReductionProved
	status := StatusFailedCPAsymmetryOperatorNotDerived
	if capacity.ViableEfficiencyWindow {
		status = StatusTensionCapacityNearTarget
	}
	return Summary{Executed: true, HasTopologicalCapacity: capacity.ViableEfficiencyWindow, CPAsymmetryDerived: lepto.DerivedCPInvariant, BaryonAsymmetryPredicted: derived, AnyReductionProved: census.LeptogenesisReduction > 0, RemainingInputs: census.RemainingInputs, Status: status, DirectAnswer: "B-gap topology has the right magnitude for leptogenesis after instanton-overlap suppression, but CP phase, washout, and Boltzmann dynamics are not derived; no CKM/PMNS parameter is consumed.", NextGate: "derive a CP-odd Majorana invariant / heavy-neutrino hierarchy, or quarantine leptogenesis phases as Phase-III vacuum data"}
}

func FormatSpan(s Span) string {
	return fmt.Sprintf("%s inherits Gate %d; adds_fit=%v; verdict=%s", s.AuditID, s.InheritedGate, s.AddsFit, s.Verdict)
}
func FormatDecay(d DecayChannel) string {
	return fmt.Sprintf("%s; %s vs %s; B_gap=%.12f; verdict=%s", d.CPAsymmetry, d.DecayMatter, d.DecayAntimatter, d.BGap, d.Verdict)
}
func FormatSakharov(s SakharovLedger) string {
	return fmt.Sprintf("B=%s; CP=%s; out-of-eq=%s; CP_operator=%v; Boltzmann=%v", s.BaryonViolation, s.CPViolation, s.OutOfEquilibrium, s.ASHAHasCPPhaseOperator, s.ASHASolvesBoltzmann)
}
func FormatTarget(t TargetAudit) string {
	return fmt.Sprintf("eta_B=%.3e; conversion=%.12f; required epsilon*kappa=%.12e; eps(kappa=.1)=%.12e; eps(kappa=.02)=%.12e", t.ObservedEtaB, t.ConversionFactor, t.RequiredEpsKappa, t.RequiredEpsKappa01, t.RequiredEpsKappa02)
}
func FormatCapacity(c TopologicalCapacity) string {
	return fmt.Sprintf("S_inst=%.12f; exp(-S)=%.12e; C_portal=%.12f; eps_witness=%.12e; kappa_required(max phase)=%.12f; viable=%v", c.SInst, c.ExpMinusSInst, c.PortalOverlap, c.InstantonOverlapEps, c.RequiredEfficiencyAtMaxPhase, c.ViableEfficiencyWindow)
}
func FormatLeptogenesis(l StandardLeptogenesis) string {
	return fmt.Sprintf("formula=%s; requires=[%s]; CP_invariant=%v; heavy_spectrum=%v; efficiency=%v", l.Formula, strings.Join(l.Requires, "; "), l.DerivedCPInvariant, l.DerivedHeavySpectrum, l.DerivedEfficiency)
}
func FormatCKMShadow(c CKMShadowAudit) string {
	return fmt.Sprintf("%s; quark_CKM_bridge=%v; consumes_CKM=%v; consumes_PMNS=%v; reduction=%d", c.Hypothesis, c.QuarkCKMBridgeDerived, c.ConsumesCKMPhase, c.ConsumesPMNSPhase, c.ParameterReduction)
}
func FormatCensus(c Census) string {
	return fmt.Sprintf("start=%d; leptogenesis_reduction=%d; remaining=%d; target=%d; reached=%v", c.StartingVacuumInputs, c.LeptogenesisReduction, c.RemainingInputs, c.SevenSealTarget, c.SevenSealReached)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("executed=%v; capacity=%v; cp_derived=%v; eta_predicted=%v; reduction=%v; remaining=%d; status=%s", s.Executed, s.HasTopologicalCapacity, s.CPAsymmetryDerived, s.BaryonAsymmetryPredicted, s.AnyReductionProved, s.RemainingInputs, s.Status)
}

func StatusLedger() []string {
	return []string{StatusDecayFormalized, StatusSakharovLedgerFormalized, StatusCPAsymmetryTargetExtracted, StatusTopologyCapacityAudited, StatusLeptogenesisSieveExecuted, StatusBaryogenesisFormalized, StatusCensusUpdated, StatusTensionCapacityNearTarget, StatusTensionEfficiencyRequired, StatusTensionCPPhaseRequired, StatusTensionCKMShadowNotEstablished, StatusTensionNoParameterReduction, StatusFailedCPAsymmetryOperatorNotDerived, StatusFailedMajoranaPhaseNotDerived, StatusFailedWashoutNotDerived, StatusFailedBoltzmannSystemNotSolved, StatusFailedCKMPhaseNotDerived, StatusFailedNoParameterReduction, StatusFailedSevenCoordinatesNotProved}
}
