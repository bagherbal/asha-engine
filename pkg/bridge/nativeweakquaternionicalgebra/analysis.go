// Package nativeweakquaternionicalgebra implements Gate 274:
// Native Weak Quaternionic Algebra / Physical Finite Hilbert Space Reconstruction Audit.
//
// Gate 273 derived the Morita trace multiplicities κ_C:κ_Q=1:3 but correctly
// refused to promote multiplicity into a finite Dirac amplitude theorem.  The
// next proposed selector is the missing weak/quaternionic structure.  Gate 274
// therefore audits whether the native Clifford/Fock machinery supplies a true
// quaternionic algebra H, whether it completes A_F=C⊕H⊕M3(C), and whether the
// completed algebra locks the lepton/quark edge norms x:y.
//
// The result is intentionally two-tiered.  A selected two-mode weak plane has
// an exact local pseudo-real doublet representation of H; the quaternion units
// close with zero residual in a 2×2 complex representation.  But the local H is
// still conditional on a weak-plane/vacuum selector and does not by itself
// provide the physical finite Hilbert sub-bimodule, physical opposite action J,
// hypercharge/chirality attachments, or a finite action rule for Dirac edge
// norms.  Therefore Gate 274 upgrades the local quaternionic algebra support,
// but it does not derive the full Connes algebra or lock x:y.
package nativeweakquaternionicalgebra

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/weakquaternionicnormalization"
)

const (
	AuditID = "GATE274-NATIVE-WEAK-QUATERNIONIC-ALGEBRA-PHYSICAL-HILBERT-SPACE-RECONSTRUCTION-AUDIT"

	StatusGate273Inherited              = "CONDITIONAL_SUPPORT_GATE273_INNER_PRODUCT_MULTIPLICITY_LEDGER_INHERITED"
	StatusLocalHExtracted               = "CONDITIONAL_SUPPORT_LOCAL_WEAK_QUATERNIONIC_H_EXTRACTED_ON_SELECTED_DOUBLET"
	StatusQuaternionClosureVerified     = "CONDITIONAL_SUPPORT_QUATERNIONIC_CLOSURE_TABLE_VERIFIED"
	StatusCandidateFullAlgebraAssembled = "CONDITIONAL_SUPPORT_CANDIDATE_C_PLUS_H_PLUS_M3C_ASSEMBLED_UNDER_SELECTOR"
	StatusHilbertSpaceAudited           = "CONDITIONAL_SUPPORT_PHYSICAL_HILBERT_SPACE_REQUIREMENTS_AUDITED"
	StatusAmplitudeReauditCompleted     = "CONDITIONAL_SUPPORT_QUATERNIONIC_AMPLITUDE_LOCKING_REAUDIT_COMPLETED"

	StatusFailedGlobalHNative    = "FAILED_ROUTE_NATIVE_GLOBAL_QUATERNIONIC_H_SUMMAND_NOT_DERIVED"
	StatusFailedExactSMAlgebra   = "FAILED_ROUTE_EXACT_C_PLUS_H_PLUS_M3C_ALGEBRA_NOT_DERIVED"
	StatusFailedPhysicalHF       = "FAILED_ROUTE_PHYSICAL_FINITE_HILBERT_SPACE_NOT_DERIVED"
	StatusFailedPhysicalJ        = "FAILED_ROUTE_PHYSICAL_OPPOSITE_ACTION_J_STILL_MISSING"
	StatusFailedAmplitudeLock    = "FAILED_ROUTE_QUATERNIONIC_STRUCTURE_DOES_NOT_LOCK_EDGE_AMPLITUDES"
	StatusFailedXYRatio          = "FAILED_ROUTE_XY_RATIO_REMAINS_UNCONSTRAINED"
	StatusFailedA2A4             = "FAILED_ROUTE_A2_A4_HIGGS_RATIO_STILL_NOT_DERIVED"
	StatusEmpiricalSealPreserved = "FAILED_ROUTE_EMPIRICAL_YUKAWA_SEAL_REMAINS_ACTIVE"
)

type Gate273Inheritance struct {
	InnerProductBuilt    bool
	TraceWeightsComputed bool
	KappaC               float64
	KappaQ               float64
	EdgeNormsDerived     bool
	XYRatioLocked        bool
	A2A4Derived          bool
	FirewallPreserved    bool
	Verdict              string
}

type Complex2Matrix struct {
	Name string
	A00  complex128
	A01  complex128
	A10  complex128
	A11  complex128
}

type QuaternionClosureAudit struct {
	SelectedPlane            string
	SelectionAuthority       string
	Source                   string
	Identity                 Complex2Matrix
	I                        Complex2Matrix
	J                        Complex2Matrix
	K                        Complex2Matrix
	ISquareResidual          float64
	JSquareResidual          float64
	KSquareResidual          float64
	IJMinusKResidual         float64
	JIMinusNegativeKResidual float64
	AntiCommutatorResidual   float64
	LocalHExtracted          bool
	NativeToSelectedDoublet  bool
	SelectionUnsealed        bool
	GlobalHSummandDerived    bool
	Verdict                  string
}

type FullAlgebraAudit struct {
	ComplexSummandInherited     bool
	ColorM3Inherited            bool
	LocalQuaternionicH          bool
	CandidateAlgebra            string
	CandidateRealDimension      int
	CandidateComplexEnvelopeDim int
	AssembledOnlyUnderSelector  bool
	ExactSMFiniteAlgebraDerived bool
	FaithfulRepresentationReady bool
	OppositeActionReady         bool
	OrderOneReady               bool
	Verdict                     string
}

type HilbertSector struct {
	Label           string
	LeftAction      string
	RightAction     string
	ComplexDim      int
	PhysicalMeaning string
	Derived         bool
	Conditional     bool
	Missing         string
}

type PhysicalHilbertAudit struct {
	UniversalMoritaLedgerInherited bool
	CandidateSectors               []HilbertSector
	LeftDoubletHActionAvailable    bool
	RightSingletCActionAvailable   bool
	ColorActionAvailable           bool
	ChiralGradingPhysical          bool
	HyperchargeAttachmentDerived   bool
	OppositeActionJDerived         bool
	ExactPhysicalHFDerived         bool
	Verdict                        string
}

type AmplitudeLockingAudit struct {
	KappaC                        float64
	KappaQ                        float64
	QuaternionicLeftDoubletFactor float64
	MultiplicityWeightsUpdated    bool
	EdgeNormCSelected             bool
	EdgeNormQSelected             bool
	XOverYLocked                  bool
	CandidateIfEqualEdgeNorms     string
	EqualEdgeNormsDerived         bool
	Reason                        string
	Verdict                       string
}

type SpectralTraceCandidate struct {
	Name    string
	X       float64
	Y       float64
	TraceD2 float64
	TraceD4 float64
	Ratio   float64
}

type SpectralTraceAudit struct {
	FormulaD2            string
	FormulaD4            string
	Candidates           []SpectralTraceCandidate
	RatioDependsOnXOverY bool
	StableInvariant      bool
	A2A4Derived          bool
	HiggsRatioDerived    bool
	Verdict              string
}

type FirewallAudit struct {
	NoConnesAlgebraImportedAsTheorem bool
	NoWeakPlaneUnsealed              bool
	NoObservedMassInserted           bool
	NoYukawaAmplitudeInserted        bool
	NoVEVInserted                    bool
	NoHiggsPredictionClaimed         bool
	LocalHNotPromotedToGlobalH       bool
	MultiplicityNotAmplitude         bool
	FiniteCorePolluted               bool
	Verdict                          string
}

type FutureCriterion struct {
	Name      string
	Required  bool
	Satisfied bool
	Detail    string
}

type FutureMap struct {
	Criteria                 []FutureCriterion
	NeedUnsealedWeakPlane    bool
	NeedPhysicalFiniteHF     bool
	NeedPhysicalJ            bool
	NeedEdgeNormAction       bool
	NeedHeatKernelProjection bool
	RecommendedNextGate      string
	Verdict                  string
}

type Summary struct {
	Gate273Inherited       bool
	LocalHExtracted        bool
	QuaternionClosureExact bool
	CandidateAlgebraBuilt  bool
	ExactSMAlgebraDerived  bool
	PhysicalHFDerived      bool
	PhysicalJDerived       bool
	EdgeAmplitudesLocked   bool
	XYRatioLocked          bool
	A2A4Derived            bool
	HiggsRatioDerived      bool
	FirewallPreserved      bool
	Status                 string
	NextGate               string
	Comment                string
}

type Analysis struct {
	PreviousGate273 weakquaternionicnormalization.Analysis
	Inheritance     Gate273Inheritance
	Quaternionic    QuaternionClosureAudit
	Algebra         FullAlgebraAudit
	Hilbert         PhysicalHilbertAudit
	Amplitude       AmplitudeLockingAudit
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
		prev, err := weakquaternionicnormalization.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 273 predecessor: %w", err)
			return
		}
		defaultA = Build(prev)
	})
	return defaultA, defaultErr
}

func Build(prev weakquaternionicnormalization.Analysis) Analysis {
	inh := inheritGate273(prev)
	quat := auditQuaternionicClosure()
	alg := auditFullAlgebra(inh, quat)
	hf := auditPhysicalHilbert(inh, quat, alg)
	amp := auditAmplitudeLocking(inh, quat, hf)
	trace := auditSpectralTrace(amp)
	fw := auditFirewall(quat, alg, hf, amp, trace)
	future := defineFuture(quat, alg, hf, amp, trace)
	summary := summarize(inh, quat, alg, hf, amp, trace, fw, future)
	truth := buildTruth(quat, alg, hf, amp, trace)
	return Analysis{PreviousGate273: prev, Inheritance: inh, Quaternionic: quat, Algebra: alg, Hilbert: hf, Amplitude: amp, SpectralTrace: trace, Firewall: fw, Future: future, Summary: summary, TruthStatement: truth}
}

func inheritGate273(prev weakquaternionicnormalization.Analysis) Gate273Inheritance {
	return Gate273Inheritance{
		InnerProductBuilt:    prev.Summary.InnerProductBuilt,
		TraceWeightsComputed: prev.Summary.TraceWeightsComputed,
		KappaC:               prev.InnerProduct.KappaCRatio,
		KappaQ:               prev.InnerProduct.KappaQRatio,
		EdgeNormsDerived:     prev.Summary.EdgeNormsDerived,
		XYRatioLocked:        prev.Summary.XYRatioLocked,
		A2A4Derived:          prev.Summary.A2A4Derived,
		FirewallPreserved:    prev.Summary.FirewallPreserved,
		Verdict:              StatusGate273Inherited + "; Gate 273 supplies κ_C:κ_Q=1:3 but leaves edge norms and x:y open",
	}
}

func auditQuaternionicClosure() QuaternionClosureAudit {
	one := Complex2Matrix{Name: "1", A00: 1, A11: 1}
	qi := Complex2Matrix{Name: "I_H", A00: 1i, A11: -1i}
	qj := Complex2Matrix{Name: "J_H", A01: 1, A10: -1}
	qk := Complex2Matrix{Name: "K_H", A01: 1i, A10: 1i}
	negOne := scale(one, -1)
	negK := scale(qk, -1)
	return QuaternionClosureAudit{
		SelectedPlane:            "U12 selected only under the prior SpontaneousCarrierSeal / τ_eta spatial-tag branch, not as an unsealed finite-core theorem",
		SelectionAuthority:       "conditional Gate 259/Gate 273 weak-plane lineage; exact local H is audited on that selected doublet",
		Source:                   "pseudo-real two-dimensional weak doublet representation inside the Clifford/Fock two-mode plane",
		Identity:                 one,
		I:                        qi,
		J:                        qj,
		K:                        qk,
		ISquareResidual:          norm2(sub(mul(qi, qi), negOne)),
		JSquareResidual:          norm2(sub(mul(qj, qj), negOne)),
		KSquareResidual:          norm2(sub(mul(qk, qk), negOne)),
		IJMinusKResidual:         norm2(sub(mul(qi, qj), qk)),
		JIMinusNegativeKResidual: norm2(sub(mul(qj, qi), negK)),
		AntiCommutatorResidual:   norm2(add(mul(qi, qj), mul(qj, qi))),
		LocalHExtracted:          true,
		NativeToSelectedDoublet:  true,
		SelectionUnsealed:        false,
		GlobalHSummandDerived:    false,
		Verdict:                  strings.Join([]string{StatusLocalHExtracted, StatusQuaternionClosureVerified, StatusFailedGlobalHNative}, ";"),
	}
}

func auditFullAlgebra(inh Gate273Inheritance, q QuaternionClosureAudit) FullAlgebraAudit {
	candidate := inh.InnerProductBuilt && q.LocalHExtracted
	exact := candidate && q.GlobalHSummandDerived && q.SelectionUnsealed
	return FullAlgebraAudit{
		ComplexSummandInherited:     true,
		ColorM3Inherited:            inh.TraceWeightsComputed && inh.KappaQ == 3,
		LocalQuaternionicH:          q.LocalHExtracted,
		CandidateAlgebra:            "A_candidate = C ⊕ H_U12 ⊕ M3(C)",
		CandidateRealDimension:      2 + 4 + 18,
		CandidateComplexEnvelopeDim: 1 + 4 + 9,
		AssembledOnlyUnderSelector:  candidate && !exact,
		ExactSMFiniteAlgebraDerived: exact,
		FaithfulRepresentationReady: false,
		OppositeActionReady:         false,
		OrderOneReady:               false,
		Verdict:                     strings.Join([]string{StatusCandidateFullAlgebraAssembled, StatusFailedExactSMAlgebra}, ";"),
	}
}

func auditPhysicalHilbert(inh Gate273Inheritance, q QuaternionClosureAudit, alg FullAlgebraAudit) PhysicalHilbertAudit {
	sectors := []HilbertSector{
		{Label: "L_L candidate", LeftAction: "H_U12", RightAction: "C or M3(C) flavor/color ledger", ComplexDim: 2 * 1, PhysicalMeaning: "lepton weak doublet shape", Derived: false, Conditional: true, Missing: "physical chirality, hypercharge, and J semantics"},
		{Label: "Q_L candidate", LeftAction: "H_U12", RightAction: "M3(C)", ComplexDim: 2 * 3, PhysicalMeaning: "quark weak doublet shape with color multiplicity", Derived: false, Conditional: true, Missing: "physical chirality and color-gauge representation activation"},
		{Label: "e_R/u_R/d_R candidates", LeftAction: "C", RightAction: "C or M3(C)", ComplexDim: 1 + 3 + 3, PhysicalMeaning: "right singlet shape", Derived: false, Conditional: true, Missing: "hypercharge splitting and Yukawa edge semantics"},
	}
	return PhysicalHilbertAudit{
		UniversalMoritaLedgerInherited: inh.InnerProductBuilt,
		CandidateSectors:               sectors,
		LeftDoubletHActionAvailable:    q.LocalHExtracted,
		RightSingletCActionAvailable:   true,
		ColorActionAvailable:           alg.ColorM3Inherited,
		ChiralGradingPhysical:          false,
		HyperchargeAttachmentDerived:   false,
		OppositeActionJDerived:         false,
		ExactPhysicalHFDerived:         false,
		Verdict:                        StatusHilbertSpaceAudited + ";" + StatusFailedPhysicalHF + ";" + StatusFailedPhysicalJ,
	}
}

func auditAmplitudeLocking(inh Gate273Inheritance, q QuaternionClosureAudit, hf PhysicalHilbertAudit) AmplitudeLockingAudit {
	return AmplitudeLockingAudit{
		KappaC:                        inh.KappaC,
		KappaQ:                        inh.KappaQ,
		QuaternionicLeftDoubletFactor: 2,
		MultiplicityWeightsUpdated:    q.LocalHExtracted && hf.LeftDoubletHActionAvailable,
		EdgeNormCSelected:             false,
		EdgeNormQSelected:             false,
		XOverYLocked:                  false,
		CandidateIfEqualEdgeNorms:     "If one additionally sealed ||T_C||=||T_Q||, then x:y could be chosen equal up to multiplicity convention, but this equality is not derived by H closure.",
		EqualEdgeNormsDerived:         false,
		Reason:                        "H supplies a left doublet action and doubles left-handed weak states; it does not assign independent norms to the C-shared and Q-shared Morita Dirac edges.",
		Verdict:                       StatusAmplitudeReauditCompleted + ";" + StatusFailedAmplitudeLock + ";" + StatusFailedXYRatio,
	}
}

func auditSpectralTrace(a AmplitudeLockingAudit) SpectralTraceAudit {
	kC, kQ := a.KappaC, a.KappaQ
	candidates := []SpectralTraceCandidate{}
	for _, c := range []struct {
		name string
		x, y float64
	}{{"x=1,y=1", 1, 1}, {"x=2,y=1", 2, 1}, {"x=1,y=2", 1, 2}} {
		d2 := kC*c.x*c.x + kQ*c.y*c.y
		d4 := kC*math.Pow(c.x, 4) + kQ*math.Pow(c.y, 4)
		candidates = append(candidates, SpectralTraceCandidate{Name: c.name, X: c.x, Y: c.y, TraceD2: d2, TraceD4: d4, Ratio: d2 / d4})
	}
	stable := true
	for i := 1; i < len(candidates); i++ {
		if math.Abs(candidates[i].Ratio-candidates[0].Ratio) > 1e-12 {
			stable = false
		}
	}
	return SpectralTraceAudit{
		FormulaD2:            "Tr(D_F²) proxy = κ_C |x|² + κ_Q |y|² (quaternionic left-doublet factor common to the selected weak sector)",
		FormulaD4:            "Tr(D_F⁴) proxy = κ_C |x|⁴ + κ_Q |y|⁴",
		Candidates:           candidates,
		RatioDependsOnXOverY: !stable,
		StableInvariant:      stable,
		A2A4Derived:          false,
		HiggsRatioDerived:    false,
		Verdict:              StatusFailedA2A4,
	}
}

func auditFirewall(q QuaternionClosureAudit, alg FullAlgebraAudit, hf PhysicalHilbertAudit, amp AmplitudeLockingAudit, trace SpectralTraceAudit) FirewallAudit {
	polluted := alg.ExactSMFiniteAlgebraDerived || hf.ExactPhysicalHFDerived || amp.XOverYLocked || trace.HiggsRatioDerived
	return FirewallAudit{
		NoConnesAlgebraImportedAsTheorem: true,
		NoWeakPlaneUnsealed:              !q.SelectionUnsealed,
		NoObservedMassInserted:           true,
		NoYukawaAmplitudeInserted:        true,
		NoVEVInserted:                    true,
		NoHiggsPredictionClaimed:         !trace.HiggsRatioDerived,
		LocalHNotPromotedToGlobalH:       !q.GlobalHSummandDerived && !alg.ExactSMFiniteAlgebraDerived,
		MultiplicityNotAmplitude:         !amp.XOverYLocked,
		FiniteCorePolluted:               polluted,
		Verdict:                          "FIREWALL_PRESERVED_LOCAL_H_NOT_PROMOTED_TO_DYNAMICAL_AMPLITUDE",
	}
}

func defineFuture(q QuaternionClosureAudit, alg FullAlgebraAudit, hf PhysicalHilbertAudit, amp AmplitudeLockingAudit, trace SpectralTraceAudit) FutureMap {
	criteria := []FutureCriterion{
		{Name: "unsealed weak-plane selector", Required: true, Satisfied: q.SelectionUnsealed, Detail: "select H globally without relying on SpontaneousCarrierSeal/tau_eta orientation seal"},
		{Name: "exact C⊕H⊕M3(C) associative finite algebra", Required: true, Satisfied: alg.ExactSMFiniteAlgebraDerived, Detail: "local H must become a global summand with faithful representation"},
		{Name: "physical finite H_F", Required: true, Satisfied: hf.ExactPhysicalHFDerived, Detail: "derive left doublets, right singlets, hypercharge, color, chirality in one bimodule"},
		{Name: "physical anti-linear J/opposite action", Required: true, Satisfied: hf.OppositeActionJDerived, Detail: "construct charge conjugation/opposite action on H_F"},
		{Name: "edge-map norm theorem", Required: true, Satisfied: amp.EdgeNormCSelected && amp.EdgeNormQSelected && amp.XOverYLocked, Detail: "derive ||T_C|| and ||T_Q|| or an equivalent finite action"},
		{Name: "heat-kernel / Seeley-de Witt projection", Required: true, Satisfied: trace.A2A4Derived, Detail: "map finite moments to a2/a4 with normalization/subtraction scheme"},
	}
	return FutureMap{
		Criteria:                 criteria,
		NeedUnsealedWeakPlane:    !q.SelectionUnsealed,
		NeedPhysicalFiniteHF:     !hf.ExactPhysicalHFDerived,
		NeedPhysicalJ:            !hf.OppositeActionJDerived,
		NeedEdgeNormAction:       !amp.XOverYLocked,
		NeedHeatKernelProjection: !trace.A2A4Derived,
		RecommendedNextGate:      "Gate 275 — Physical Finite Hilbert Space / Chiral Hypercharge Opposite-Action Completion Audit",
		Verdict:                  "local H closure is now exact, but amplitude dynamics require a physical H_F, J, edge-norm action, and heat-kernel map",
	}
}

func summarize(inh Gate273Inheritance, q QuaternionClosureAudit, alg FullAlgebraAudit, hf PhysicalHilbertAudit, amp AmplitudeLockingAudit, trace SpectralTraceAudit, fw FirewallAudit, fut FutureMap) Summary {
	status := strings.Join([]string{
		StatusGate273Inherited,
		StatusLocalHExtracted,
		StatusQuaternionClosureVerified,
		StatusCandidateFullAlgebraAssembled,
		StatusHilbertSpaceAudited,
		StatusAmplitudeReauditCompleted,
		StatusFailedGlobalHNative,
		StatusFailedExactSMAlgebra,
		StatusFailedPhysicalHF,
		StatusFailedPhysicalJ,
		StatusFailedAmplitudeLock,
		StatusFailedXYRatio,
		StatusFailedA2A4,
		StatusEmpiricalSealPreserved,
	}, ";")
	return Summary{
		Gate273Inherited:       inh.InnerProductBuilt && inh.TraceWeightsComputed,
		LocalHExtracted:        q.LocalHExtracted,
		QuaternionClosureExact: q.ISquareResidual == 0 && q.JSquareResidual == 0 && q.KSquareResidual == 0 && q.IJMinusKResidual == 0 && q.JIMinusNegativeKResidual == 0,
		CandidateAlgebraBuilt:  alg.AssembledOnlyUnderSelector,
		ExactSMAlgebraDerived:  alg.ExactSMFiniteAlgebraDerived,
		PhysicalHFDerived:      hf.ExactPhysicalHFDerived,
		PhysicalJDerived:       hf.OppositeActionJDerived,
		EdgeAmplitudesLocked:   amp.EdgeNormCSelected && amp.EdgeNormQSelected,
		XYRatioLocked:          amp.XOverYLocked,
		A2A4Derived:            trace.A2A4Derived,
		HiggsRatioDerived:      trace.HiggsRatioDerived,
		FirewallPreserved:      !fw.FiniteCorePolluted && fw.NoConnesAlgebraImportedAsTheorem && fw.LocalHNotPromotedToGlobalH,
		Status:                 status,
		NextGate:               fut.RecommendedNextGate,
		Comment:                "Gate 274 extracts exact local quaternionic closure on a selected weak doublet, but local H does not derive global C⊕H⊕M3(C), physical H_F, J, edge norms, or a2/a4.",
	}
}

func buildTruth(q QuaternionClosureAudit, alg FullAlgebraAudit, hf PhysicalHilbertAudit, amp AmplitudeLockingAudit, trace SpectralTraceAudit) string {
	return fmt.Sprintf("Gate 274 verifies an exact local quaternionic algebra on the selected weak doublet: I²=J²=K²=-1 and IJ=K with zero residuals. This upgrades the weak/H preflight but only under a selected plane. It does not derive a global H summand, physical finite Hilbert space, physical J, or edge-norm theorem; consequently x:y remains free and the spectral ratio remains amplitude-dependent across %d audited representatives.", len(trace.Candidates))
}

func mul(x, y Complex2Matrix) Complex2Matrix {
	return Complex2Matrix{Name: x.Name + y.Name,
		A00: x.A00*y.A00 + x.A01*y.A10,
		A01: x.A00*y.A01 + x.A01*y.A11,
		A10: x.A10*y.A00 + x.A11*y.A10,
		A11: x.A10*y.A01 + x.A11*y.A11,
	}
}

func add(x, y Complex2Matrix) Complex2Matrix {
	return Complex2Matrix{Name: x.Name + "+" + y.Name, A00: x.A00 + y.A00, A01: x.A01 + y.A01, A10: x.A10 + y.A10, A11: x.A11 + y.A11}
}

func sub(x, y Complex2Matrix) Complex2Matrix {
	return Complex2Matrix{Name: x.Name + "-" + y.Name, A00: x.A00 - y.A00, A01: x.A01 - y.A01, A10: x.A10 - y.A10, A11: x.A11 - y.A11}
}

func scale(x Complex2Matrix, s complex128) Complex2Matrix {
	return Complex2Matrix{Name: fmt.Sprintf("(%v)%s", s, x.Name), A00: s * x.A00, A01: s * x.A01, A10: s * x.A10, A11: s * x.A11}
}

func norm2(x Complex2Matrix) float64 {
	return cmag2(x.A00) + cmag2(x.A01) + cmag2(x.A10) + cmag2(x.A11)
}

func cmag2(z complex128) float64 { return real(z)*real(z) + imag(z)*imag(z) }
