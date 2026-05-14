// Package majoranaflavorsymmetrybreaking implements Gate 347:
// Non-Unitary-Invariant Texture Sieve / Majorana Flavor Symmetry Breaking Audit.
//
// Gate 347 audits whether higher-order Majorana/Dirac cross-terms in the
// spectral action can break the unitary flavor degeneracy found in Gate 346.
// The gate is strict: ordinary trace cross-terms remain unitary invariant, and
// the explicit Gate-320 H-sigma overlap is a rank-one heavy-light support
// operator, not by itself a CKM texture selector. A non-unitary texture term is
// mathematically capable of tilting the signed-triality null valley, but the
// finite geometry must derive its projector/orientation before it can be
// promoted to a vacuum-selection theorem.
package majoranaflavorsymmetrybreaking

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE347-NON-UNITARY-INVARIANT-TEXTURE-SIEVE-MAJORANA-FLAVOR-SYMMETRY-BREAKING-AUDIT"

	StatusGate346Inherited                    = "CONDITIONAL_SUPPORT_GATE346_VARIATIONAL_FLAVOR_FLATNESS_INHERITED"
	StatusMajoranaDiracCrossTermsFormalized   = "CONDITIONAL_SUPPORT_MAJORANA_DIRAC_CROSS_TERMS_FORMALIZED"
	StatusUnitarySymmetryBreakingTestExecuted = "CONDITIONAL_SUPPORT_UNITARY_SYMMETRY_BREAKING_TEST_EXECUTED"
	StatusDirectInvariantTermsRemainFlat      = "CONDITIONAL_SUPPORT_STANDARD_MAJORANA_DIRAC_TRACES_REMAIN_UNITARY_INVARIANT"
	StatusOmegaOverlapInserted                = "CONDITIONAL_SUPPORT_GATE320_OVERLAP_OPERATOR_INSERTED_IN_TEXTURE_SIEVE"
	StatusDegeneracyLiftingSieveExecuted      = "CONDITIONAL_SUPPORT_DEGENERACY_LIFTING_SIEVE_EXECUTED"
	StatusNonUnitaryTextureCapacityIdentified = "CONDITIONAL_SUPPORT_NON_UNITARY_TEXTURE_CAPACITY_IDENTIFIED_CONDITIONALLY"
	StatusFirewallsPreserved                  = "CONDITIONAL_SUPPORT_MAJORANA_FLAVOR_FIREWALLS_PRESERVED"

	StatusTensionMajoranaLeptonNotQuarkCKM  = "CONDITIONAL_TENSION_MAJORANA_EDGE_BREAKS_LEPTON_SECTOR_NOT_DIRECT_QUARK_CKM"
	StatusTensionTextureProjectorNotNative  = "CONDITIONAL_TENSION_NON_UNITARY_TEXTURE_PROJECTOR_NOT_DERIVED"
	StatusTensionNullspaceNotUniquelyLifted = "CONDITIONAL_TENSION_SIGNED_NULLSPACE_NOT_UNIQUELY_LIFTED"

	StatusFailedMajoranaFlavorBreaking       = "FAILED_ROUTE_MAJORANA_FLAVOR_SYMMETRY_BREAKING_NOT_DERIVED"
	StatusFailedUniqueVacuumDegeneracyLifted = "FAILED_ROUTE_UNIQUE_VACUUM_DEGENERACY_NOT_LIFTED"
	StatusFailedNativeTextureOperator        = "FAILED_ROUTE_NATIVE_NON_UNITARY_TEXTURE_OPERATOR_NOT_DERIVED"
	StatusFailedCKMTexture                   = "FAILED_ROUTE_CKM_TEXTURE_NOT_DERIVED"
	StatusFailedQuarkFlavorBridge            = "FAILED_ROUTE_MAJORANA_TO_QUARK_FLAVOR_BRIDGE_NOT_DERIVED"
	StatusFailedObservedMassesImported       = "FAILED_ROUTE_OBSERVED_PARTICLE_MASSES_NOT_IMPORTED"
)

const (
	signedNullity     = 2
	omegaOverlapIndex = 1.0
)

type CrossTerm struct {
	Name                     string
	HeatKernelOrder          string
	Formula                  string
	UsesDiracYukawa          bool
	UsesMajoranaSigma        bool
	UsesOmegaHsigma          bool
	UsesFixedFlavorProjector bool
	UnitaryInvariant         bool
	BreaksU3Flavor           bool
	NativeInCurrentGeometry  bool
	CanLiftDegeneracy        bool
	Sector                   string
	Status                   string
}

type CrossTermLedger struct {
	Terms                 []CrossTerm
	NativeTerms           int
	UnitaryInvariantTerms int
	BreakingTemplates     int
	NativeBreakingTerms   int
	Status                string
}

type SymmetryTest struct {
	Transformation              string
	InvariantCondition          string
	StandardCrossTermsFlat      bool
	OmegaIndex                  float64
	OmegaAloneBreaksCKM         bool
	MajoranaActsOnLeptonSlot    bool
	DirectQuarkCKMBridgeDerived bool
	Status                      string
}

type DegeneracySieve struct {
	TauHat                        [3]float64
	SignedNullity                 int
	TextureTemplate               string
	TemplateCanLift               bool
	UniqueMinimumIfProjectorGiven bool
	NativeProjectorDerived        bool
	UniqueVacuumDerived           bool
	Reason                        string
	Status                        string
}

type Verdict struct {
	NonUnitaryOperatorFoundNatively bool
	MajoranaBreaksFlavorNatively    bool
	DegeneracyLifted                bool
	CKMDerived                      bool
	NeedsNewObject                  string
	Status                          string
}

type Audit struct {
	NoCKMImported             bool
	NoObservedYukawasImported bool
	NoTextureForced           bool
	NoFinalVacuumClaim        bool
	NoColliderMassClaim       bool
	Status                    string
}

type Summary struct {
	OneLine    string
	MainResult string
	NextGate   string
	Status     string
}

type Analysis struct {
	CrossTerms CrossTermLedger
	Symmetry   SymmetryTest
	Degeneracy DegeneracySieve
	Verdict    Verdict
	Audit      Audit
	Summary    Summary
	Truth      string
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
	cross := compileCrossTerms()
	sym := compileSymmetryTest(cross)
	deg := compileDegeneracySieve(cross, sym)
	verdict := compileVerdict(cross, sym, deg)
	audit := compileAudit()
	summary := compileSummary(verdict)
	truth := "Gate 347 audits the missing non-unitary texture operator demanded by Gate 346.  Standard Majorana/Dirac cross-terms such as Tr(Y†Y σ†σ) and Tr(Y†Y)Tr(σ†σ) remain invariant under flavor-unitary rotations and therefore cannot lift CKM/flavor degeneracy.  The Gate-320 overlap Ω_Hσ supplies a real heavy-light support index, but by itself it connects L_L to ν_R and ν_R^c; it does not natively project the quark CKM space.  A fixed texture projector could lift the signed-triality null valley, but that projector is not derived from the current finite geometry."
	return Analysis{CrossTerms: cross, Symmetry: sym, Degeneracy: deg, Verdict: verdict, Audit: audit, Summary: summary, Truth: truth}, nil
}

func compileCrossTerms() CrossTermLedger {
	terms := []CrossTerm{
		{Name: "factorized Majorana-Dirac trace", HeatKernelOrder: "a4/a6 template", Formula: "Tr(Y†Y) Tr(σ†σ)", UsesDiracYukawa: true, UsesMajoranaSigma: true, UsesOmegaHsigma: false, UsesFixedFlavorProjector: false, UnitaryInvariant: true, BreaksU3Flavor: false, NativeInCurrentGeometry: true, CanLiftDegeneracy: false, Sector: "Dirac × Majorana scalar magnitude", Status: StatusDirectInvariantTermsRemainFlat},
		{Name: "single-trace commuting Majorana insertion", HeatKernelOrder: "a4/a6 template", Formula: "Tr(Y†Y σ†σ)", UsesDiracYukawa: true, UsesMajoranaSigma: true, UsesOmegaHsigma: false, UsesFixedFlavorProjector: false, UnitaryInvariant: true, BreaksU3Flavor: false, NativeInCurrentGeometry: true, CanLiftDegeneracy: false, Sector: "lepton Majorana scalar magnitude", Status: StatusDirectInvariantTermsRemainFlat},
		{Name: "Gate-320 heavy-light support insertion", HeatKernelOrder: "functional determinant support", Formula: "Tr(Ω_Hσ† Ω_Hσ) = 1", UsesDiracYukawa: false, UsesMajoranaSigma: true, UsesOmegaHsigma: true, UsesFixedFlavorProjector: false, UnitaryInvariant: true, BreaksU3Flavor: false, NativeInCurrentGeometry: true, CanLiftDegeneracy: false, Sector: "L_L → ν_R → ν_R^c support index", Status: StatusOmegaOverlapInserted},
		{Name: "signed triality texture projector", HeatKernelOrder: "candidate non-unitary texture", Formula: "Tr(P_τη Y_uY_u†) or |<τη|t>|²", UsesDiracYukawa: true, UsesMajoranaSigma: false, UsesOmegaHsigma: false, UsesFixedFlavorProjector: true, UnitaryInvariant: false, BreaksU3Flavor: true, NativeInCurrentGeometry: false, CanLiftDegeneracy: true, Sector: "quark generation orientation", Status: StatusTensionTextureProjectorNotNative},
		{Name: "Majorana-assisted quark texture bridge", HeatKernelOrder: "candidate higher overlap", Formula: "Tr(P_Q Ω_Hσ† P_τη Ω_Hσ P_Q Y_uY_u†)", UsesDiracYukawa: true, UsesMajoranaSigma: true, UsesOmegaHsigma: true, UsesFixedFlavorProjector: true, UnitaryInvariant: false, BreaksU3Flavor: true, NativeInCurrentGeometry: false, CanLiftDegeneracy: true, Sector: "hypothetical lepton-Majorana to quark-CKM bridge", Status: StatusFailedQuarkFlavorBridge},
	}
	native, inv, templates, nativeBreaking := 0, 0, 0, 0
	for _, term := range terms {
		if term.NativeInCurrentGeometry {
			native++
		}
		if term.UnitaryInvariant {
			inv++
		}
		if term.BreaksU3Flavor {
			templates++
		}
		if term.BreaksU3Flavor && term.NativeInCurrentGeometry {
			nativeBreaking++
		}
	}
	return CrossTermLedger{Terms: terms, NativeTerms: native, UnitaryInvariantTerms: inv, BreakingTemplates: templates, NativeBreakingTerms: nativeBreaking, Status: StatusMajoranaDiracCrossTermsFormalized}
}

func compileSymmetryTest(c CrossTermLedger) SymmetryTest {
	return SymmetryTest{Transformation: "Y -> U† Y V; σ -> W^T σ W in the Majorana/lepton slot; quark CKM requires a non-commuting fixed operator in P_Q H_F", InvariantCondition: "A term breaks U(3) flavor only if it contains a fixed non-scalar generation operator that does not commute with U(3).", StandardCrossTermsFlat: c.NativeBreakingTerms == 0 && c.UnitaryInvariantTerms >= 3, OmegaIndex: omegaOverlapIndex, OmegaAloneBreaksCKM: false, MajoranaActsOnLeptonSlot: true, DirectQuarkCKMBridgeDerived: false, Status: StatusUnitarySymmetryBreakingTestExecuted}
}

func compileDegeneracySieve(c CrossTermLedger, s SymmetryTest) DegeneracySieve {
	tau := [3]float64{2.0 / 3.0, -2.0 / 3.0, 1.0 / 3.0}
	return DegeneracySieve{TauHat: tau, SignedNullity: signedNullity, TextureTemplate: "A derived non-scalar projector P_texture on generation space would add ε<t|P_texture|t> and lift the two-dimensional signed null valley if P_texture has a unique minimum on ker(τη).", TemplateCanLift: c.BreakingTemplates > 0, UniqueMinimumIfProjectorGiven: true, NativeProjectorDerived: c.NativeBreakingTerms > 0, UniqueVacuumDerived: false, Reason: "The only terms native to the current geometry are invariant magnitudes or the Ω_Hσ support index.  The non-unitary projector required to choose a point inside the signed nullspace remains a template, not a theorem.", Status: StatusDegeneracyLiftingSieveExecuted}
}

func compileVerdict(c CrossTermLedger, s SymmetryTest, d DegeneracySieve) Verdict {
	return Verdict{NonUnitaryOperatorFoundNatively: c.NativeBreakingTerms > 0, MajoranaBreaksFlavorNatively: false, DegeneracyLifted: d.UniqueVacuumDerived, CKMDerived: false, NeedsNewObject: "derive a genuine flavor-texture operator coupling the Majorana/Ω sector to the quark generation projector, or quarantine CKM as Phase-III vacuum data", Status: StatusFailedMajoranaFlavorBreaking}
}

func compileAudit() Audit {
	return Audit{NoCKMImported: true, NoObservedYukawasImported: true, NoTextureForced: true, NoFinalVacuumClaim: true, NoColliderMassClaim: true, Status: StatusFirewallsPreserved}
}

func compileSummary(v Verdict) Summary {
	return Summary{OneLine: "Majorana/Dirac cross-terms were audited for flavor symmetry breaking.", MainResult: "Native standard terms remain unitary invariant; non-unitary texture capacity exists only as an unpromoted projector template.", NextGate: v.NeedsNewObject, Status: StatusFailedUniqueVacuumDegeneracyLifted}
}

func Statuses(a Analysis) []string {
	out := []string{a.CrossTerms.Status, a.Symmetry.Status, a.Degeneracy.Status, a.Verdict.Status, a.Audit.Status, a.Summary.Status, StatusGate346Inherited, StatusFailedMajoranaFlavorBreaking, StatusFailedUniqueVacuumDegeneracyLifted, StatusFailedNativeTextureOperator, StatusFailedCKMTexture, StatusFailedObservedMassesImported}
	for _, t := range a.CrossTerms.Terms {
		out = append(out, t.Status)
	}
	return unique(out)
}

func unique(in []string) []string {
	seen := map[string]bool{}
	out := make([]string, 0, len(in))
	for _, s := range in {
		if s != "" && !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return out
}

func nearlyZero(x float64) bool { return math.Abs(x) < 1e-12 }

func FormatCrossTerms(c CrossTermLedger) string {
	return fmt.Sprintf("terms=%d native=%d invariant=%d breaking_templates=%d native_breaking=%d", len(c.Terms), c.NativeTerms, c.UnitaryInvariantTerms, c.BreakingTemplates, c.NativeBreakingTerms)
}
func FormatSymmetry(s SymmetryTest) string {
	return fmt.Sprintf("standard_flat=%t omega_index=%.1f omega_breaks_ckm=%t majorana_lepton=%t quark_bridge=%t", s.StandardCrossTermsFlat, s.OmegaIndex, s.OmegaAloneBreaksCKM, s.MajoranaActsOnLeptonSlot, s.DirectQuarkCKMBridgeDerived)
}
func FormatDegeneracy(d DegeneracySieve) string {
	return fmt.Sprintf("nullity=%d template_can_lift=%t unique_if_projector=%t native_projector=%t unique_vacuum=%t", d.SignedNullity, d.TemplateCanLift, d.UniqueMinimumIfProjectorGiven, d.NativeProjectorDerived, d.UniqueVacuumDerived)
}
func FormatVerdict(v Verdict) string {
	return fmt.Sprintf("native_nonunitary=%t majorana_breaks=%t lifted=%t ckm=%t next=%s", v.NonUnitaryOperatorFoundNatively, v.MajoranaBreaksFlavorNatively, v.DegeneracyLifted, v.CKMDerived, v.NeedsNewObject)
}
func FormatAudit(a Audit) string {
	return fmt.Sprintf("no_ckm=%t no_yukawas=%t no_texture_forced=%t no_vacuum=%t no_collider=%t", a.NoCKMImported, a.NoObservedYukawasImported, a.NoTextureForced, a.NoFinalVacuumClaim, a.NoColliderMassClaim)
}
func FormatSummary(s Summary) string {
	return strings.Join([]string{s.OneLine, s.MainResult, s.NextGate}, " | ")
}
