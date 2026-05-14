// Package flavororientationoperator implements Gate 324:
// Flavor Orientation Operator / Triality-to-Mass-Eigenstate Texture Audit.
//
// Gate 323 proved that pulling tau_eta=(2,-2,1) directly onto the physical top
// slot gives nonzero UV top-Yukawa fractions (4/9 or 1/9), and every nonzero
// tested fraction spoils the Gate-322 near-125 GeV transport in the one-loop
// lane.  Gate 324 therefore audits the missing flavor-orientation operator that
// should rotate the native geometric generation basis into the physical mass
// eigenbasis.
//
// The result is intentionally strict.  The algebra has a real capacity to
// suppress the top boundary: tau_eta has a two-dimensional nullspace, so a
// physical top vector placed in that nullspace has exactly zero overlap with the
// triality source.  However, the finite geometry supplied so far does not derive
// a unique unitary U_flavor or CKM-like texture selecting one null vector as the
// physical top.  Therefore the gate formalizes the mechanism and the nullspace
// capacity, but it does not authorize the successful Gate-322 gauge-only lane as
// the physical Standard Model top sector.
package flavororientationoperator

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE324-FLAVOR-ORIENTATION-OPERATOR-TRIALITY-TO-MASS-EIGENSTATE-TEXTURE-AUDIT"

	StatusBasisDistinctionFormalized    = "CONDITIONAL_SUPPORT_GEOMETRIC_BASIS_VS_MASS_BASIS_FORMALIZED"
	StatusFlavorMatrixSieveFormalized   = "CONDITIONAL_SUPPORT_FLAVOR_ORIENTATION_MATRIX_SIEVE_FORMALIZED"
	StatusTopNullspaceCapacityProved    = "CONDITIONAL_SUPPORT_TOP_ORTHOGONALITY_CAPACITY_PROVED"
	StatusRGCompatibilityAudited        = "CONDITIONAL_SUPPORT_RG_COMPATIBILITY_AUDITED"
	StatusNullRotationRecoversGate322   = "CONDITIONAL_SUPPORT_NULL_TOP_ROTATION_REPRODUCES_GATE322_ENVELOPE"
	StatusNativeOperatorMissing         = "CONDITIONAL_TENSION_NATIVE_FLAVOR_ORIENTATION_OPERATOR_NOT_SELECTED"
	StatusNullspaceDegenerate           = "CONDITIONAL_TENSION_TAU_ETA_NULLSPACE_TWO_DIMENSIONAL"
	StatusTopSuppressionRequiresTexture = "CONDITIONAL_TENSION_TOP_BOUNDARY_SUPPRESSION_REQUIRES_EXTRA_TEXTURE"

	StatusFailedFlavorOperatorNotDerived   = "FAILED_ROUTE_FLAVOR_ORIENTATION_OPERATOR_NOT_DERIVED"
	StatusFailedTopSuppressionNotJustified = "FAILED_ROUTE_TOP_BOUNDARY_SUPPRESSION_NOT_JUSTIFIED"
	StatusFailedCKMTextureNotDerived       = "FAILED_ROUTE_CKM_TEXTURE_NOT_DERIVED"
	StatusFailedMassBasisNotNative         = "FAILED_ROUTE_MASS_EIGENSTATE_BASIS_NOT_DERIVED"
	StatusFailedGate322StillDiagnostic     = "FAILED_ROUTE_GATE322_GAUGE_ONLY_LANE_STILL_DIAGNOSTIC"
	StatusFailedPoleMassNotExecuted        = "FAILED_ROUTE_POLE_MASS_CONVERSION_NOT_EXECUTED"
	StatusFailedTwoLoopNotExecuted         = "FAILED_ROUTE_TWO_LOOP_RG_NOT_EXECUTED"
	StatusFailedColliderMassNotClaimed     = "FAILED_ROUTE_FINAL_COLLIDER_HIGGS_MASS_NOT_CLAIMED"
)

const (
	gate322RunningMassGeV        = 124.9766199157
	gate323UniqueLowMassGeV      = 258.687 // inherited preflight rounded in audit text
	gate323HighSlotMassGeV       = 317.115 // inherited preflight rounded in audit text
	gate323GaugeOnlyTopFraction  = 0.0
	nearObservedTolerancePercent = 1.0
)

type BasisAudit struct {
	Formalized       bool
	GeometricBasis   []string
	MassBasis        []string
	TrialityTensor   []float64
	TrialityNorm     float64
	NormalizedSource []float64
	Distinction      string
	Verdict          string
}

type OperatorSourceAudit struct {
	SieveFormalized             bool
	JSwapActsOnFlavor           bool
	DoubledSpaceActsOnFlavor    bool
	BimoduleOverlapActsOnFlavor bool
	InstalledNativeUnitary      bool
	CandidateAllowedByUnitarity bool
	WhyNativeMissing            string
	Verdict                     string
}

type FlavorCandidate struct {
	Name                  string
	Description           string
	TopVector             []float64
	UnitNorm              bool
	TauOverlap            float64
	TopFraction           float64
	Native                bool
	Unique                bool
	SuppressesTopBoundary bool
	NearGate322           bool
	InheritedMassGeV      float64
	Verdict               string
}

type NullspaceAudit struct {
	Computed                  bool
	Dimension                 int
	Basis                     [][]float64
	AllBasisVectorsOrthogonal bool
	TopSuppressionPossible    bool
	UniquePhysicalTopVector   bool
	Verdict                   string
}

type RGCompatibilityAudit struct {
	Audited                      bool
	Gate322GaugeOnlyMassGeV      float64
	IdentityLowSlotMassGeV       float64
	IdentityHighSlotMassGeV      float64
	NullTopMassGeV               float64
	NullRotationPreservesGate322 bool
	NativeJustificationInstalled bool
	PhysicalLaneAuthorized       bool
	Verdict                      string
}

type FirewallAudit struct {
	NoCKMImported             bool
	NoObservedTopMassInserted bool
	NoFlavorTextureInvented   bool
	NoPoleMassClaimed         bool
	NoTwoLoopClaimed          bool
	NoColliderMassClaimed     bool
	FiniteCorePolluted        bool
	Verdict                   string
}

type Summary struct {
	BasisFormalized                 bool
	FlavorSieveFormalized           bool
	NullspaceCapacityProved         bool
	NativeFlavorOperatorDerived     bool
	TopBoundarySuppressionJustified bool
	Gate322PhysicalLaneAuthorized   bool
	FirewallsPreserved              bool
	FinalMassClaimed                bool
	Status                          string
	DirectAnswer                    string
	NextGate                        string
}

type Analysis struct {
	Basis      BasisAudit
	Operator   OperatorSourceAudit
	Candidates []FlavorCandidate
	Nullspace  NullspaceAudit
	RG         RGCompatibilityAudit
	Firewalls  FirewallAudit
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
	basis := formalizeBasis()
	operator := auditOperatorSources()
	nullspace := computeNullspace(basis.NormalizedSource)
	candidates := buildCandidates(basis.NormalizedSource)
	rg := auditRGCompatibility(candidates)
	firewalls := auditFirewalls(operator, rg)
	summary := buildSummary(basis, operator, nullspace, rg, firewalls)
	truth := "Gate 324 proves the exact mathematical capacity needed by Gate 322: the triality source tau_eta=(2,-2,1) has a two-dimensional nullspace, and any physical top vector placed in that nullspace has zero GUT-boundary top-Yukawa overlap.  This formally explains how a flavor orientation operator could justify the flattened-top lane.  However, the known doubled-space J_swap and seesaw overlap operators act on particle/antiparticle and chirality/heavy-light structure, not on the three-generation flavor basis.  No native CKM/flavor unitary is derived, and the nullspace is not unique.  Therefore Gate 324 formalizes the flavor-orientation mechanism and proves suppression capacity, but it preserves the firewall that the successful Gate-322 gauge-only envelope is still diagnostic rather than a fully physical top-sector transport."
	return Analysis{Basis: basis, Operator: operator, Candidates: candidates, Nullspace: nullspace, RG: rg, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func formalizeBasis() BasisAudit {
	tau := []float64{2, -2, 1}
	norm := vectorNorm(tau)
	normed := scaleVector(tau, 1.0/norm)
	return BasisAudit{
		Formalized:       true,
		GeometricBasis:   []string{"g1", "g2", "g3"},
		MassBasis:        []string{"u_phys", "c_phys", "t_phys"},
		TrialityTensor:   tau,
		TrialityNorm:     norm,
		NormalizedSource: normed,
		Distinction:      "tau_eta is diagonal in the geometric trace basis; physical Yukawa eigenstates require a unitary flavor-orientation matrix U_flavor before assigning top, charm, or up labels.",
		Verdict:          StatusBasisDistinctionFormalized,
	}
}

func auditOperatorSources() OperatorSourceAudit {
	return OperatorSourceAudit{
		SieveFormalized:             true,
		JSwapActsOnFlavor:           false,
		DoubledSpaceActsOnFlavor:    false,
		BimoduleOverlapActsOnFlavor: false,
		InstalledNativeUnitary:      false,
		CandidateAllowedByUnitarity: true,
		WhyNativeMissing:            "J_swap exchanges particle and antiparticle slots, and the Gate-320 seesaw operator connects L_L to nu_R^c; neither supplies a generation-space unitary selecting the physical CKM/mass basis.",
		Verdict:                     strings.Join([]string{StatusFlavorMatrixSieveFormalized, StatusNativeOperatorMissing, StatusFailedFlavorOperatorNotDerived, StatusFailedCKMTextureNotDerived}, ";"),
	}
}

func computeNullspace(source []float64) NullspaceAudit {
	basis := [][]float64{
		normalize([]float64{1, 1, 0}),
		normalize([]float64{1, 0, -2}),
	}
	allOrtho := true
	for _, v := range basis {
		if math.Abs(dot(source, v)) > 1e-12 {
			allOrtho = false
		}
	}
	return NullspaceAudit{
		Computed:                  true,
		Dimension:                 2,
		Basis:                     basis,
		AllBasisVectorsOrthogonal: allOrtho,
		TopSuppressionPossible:    allOrtho,
		UniquePhysicalTopVector:   false,
		Verdict:                   strings.Join([]string{StatusTopNullspaceCapacityProved, StatusNullspaceDegenerate, StatusFailedMassBasisNotNative}, ";"),
	}
}

func buildCandidates(source []float64) []FlavorCandidate {
	candidates := []FlavorCandidate{
		candidate("identity_high_slot_g1", "direct unmixed assignment to the first |tau|=2 geometric slot", []float64{1, 0, 0}, false, false, gate323HighSlotMassGeV),
		candidate("identity_high_slot_g2", "direct unmixed assignment to the second |tau|=2 geometric slot", []float64{0, 1, 0}, false, false, gate323HighSlotMassGeV),
		candidate("identity_low_slot_g3", "direct unmixed assignment to the unique |tau|=1 geometric slot", []float64{0, 0, 1}, false, false, gate323UniqueLowMassGeV),
		candidate("nullspace_symmetric_top", "constructed top vector (g1+g2)/sqrt(2), orthogonal to tau_eta", normalize([]float64{1, 1, 0}), true, false, gate322RunningMassGeV),
		candidate("nullspace_mixed_top", "constructed top vector (g1-2g3)/sqrt(5), orthogonal to tau_eta", normalize([]float64{1, 0, -2}), true, false, gate322RunningMassGeV),
	}
	for i := range candidates {
		candidates[i].TauOverlap = dot(source, candidates[i].TopVector)
		candidates[i].TopFraction = candidates[i].TauOverlap * candidates[i].TauOverlap
		candidates[i].UnitNorm = math.Abs(vectorNorm(candidates[i].TopVector)-1) < 1e-12
		candidates[i].SuppressesTopBoundary = candidates[i].TopFraction < 1e-12
		candidates[i].NearGate322 = math.Abs(candidates[i].InheritedMassGeV-gate322RunningMassGeV)/gate322RunningMassGeV*100 < nearObservedTolerancePercent
		if candidates[i].SuppressesTopBoundary {
			candidates[i].Verdict = strings.Join([]string{StatusTopNullspaceCapacityProved, StatusNullRotationRecoversGate322, StatusFailedFlavorOperatorNotDerived}, ";")
		} else {
			candidates[i].Verdict = strings.Join([]string{StatusRGCompatibilityAudited, StatusFailedTopSuppressionNotJustified}, ";")
		}
	}
	return candidates
}

func candidate(name, desc string, vector []float64, allowed bool, native bool, mass float64) FlavorCandidate {
	return FlavorCandidate{Name: name, Description: desc, TopVector: vector, Native: native, Unique: false, InheritedMassGeV: mass}
}

func auditRGCompatibility(c []FlavorCandidate) RGCompatibilityAudit {
	var nullMass, lowMass, highMass float64
	nullPreserves := false
	for _, cand := range c {
		switch cand.Name {
		case "identity_low_slot_g3":
			lowMass = cand.InheritedMassGeV
		case "identity_high_slot_g1":
			highMass = cand.InheritedMassGeV
		case "nullspace_symmetric_top":
			nullMass = cand.InheritedMassGeV
			nullPreserves = cand.NearGate322 && cand.SuppressesTopBoundary
		}
	}
	return RGCompatibilityAudit{
		Audited:                      true,
		Gate322GaugeOnlyMassGeV:      gate322RunningMassGeV,
		IdentityLowSlotMassGeV:       lowMass,
		IdentityHighSlotMassGeV:      highMass,
		NullTopMassGeV:               nullMass,
		NullRotationPreservesGate322: nullPreserves,
		NativeJustificationInstalled: false,
		PhysicalLaneAuthorized:       false,
		Verdict:                      strings.Join([]string{StatusRGCompatibilityAudited, StatusNullRotationRecoversGate322, StatusTopSuppressionRequiresTexture, StatusFailedGate322StillDiagnostic}, ";"),
	}
}

func auditFirewalls(op OperatorSourceAudit, rg RGCompatibilityAudit) FirewallAudit {
	return FirewallAudit{
		NoCKMImported:             true,
		NoObservedTopMassInserted: true,
		NoFlavorTextureInvented:   !op.InstalledNativeUnitary,
		NoPoleMassClaimed:         true,
		NoTwoLoopClaimed:          true,
		NoColliderMassClaimed:     true,
		FiniteCorePolluted:        false,
		Verdict:                   strings.Join([]string{StatusFailedFlavorOperatorNotDerived, StatusFailedPoleMassNotExecuted, StatusFailedTwoLoopNotExecuted, StatusFailedColliderMassNotClaimed}, ";"),
	}
}

func buildSummary(b BasisAudit, op OperatorSourceAudit, n NullspaceAudit, rg RGCompatibilityAudit, f FirewallAudit) Summary {
	preserved := f.NoCKMImported && f.NoObservedTopMassInserted && f.NoFlavorTextureInvented && f.NoPoleMassClaimed && f.NoTwoLoopClaimed && f.NoColliderMassClaimed && !f.FiniteCorePolluted
	return Summary{
		BasisFormalized:                 b.Formalized,
		FlavorSieveFormalized:           op.SieveFormalized,
		NullspaceCapacityProved:         n.TopSuppressionPossible,
		NativeFlavorOperatorDerived:     op.InstalledNativeUnitary,
		TopBoundarySuppressionJustified: rg.PhysicalLaneAuthorized,
		Gate322PhysicalLaneAuthorized:   rg.PhysicalLaneAuthorized,
		FirewallsPreserved:              preserved,
		FinalMassClaimed:                false,
		Status:                          strings.Join([]string{StatusBasisDistinctionFormalized, StatusFlavorMatrixSieveFormalized, StatusTopNullspaceCapacityProved, StatusNativeOperatorMissing, StatusFailedFlavorOperatorNotDerived}, ";"),
		DirectAnswer:                    "A flavor rotation can make the physical top vector exactly orthogonal to tau_eta, giving y_t(Λ_GUT)=0 and preserving the Gate-322 transport, but the finite geometry has not derived the unique U_flavor/CKM texture that selects this rotation.",
		NextGate:                        "derive the native flavor-orientation/CKM texture operator, or prove that the flattened-top lane is only a diagnostic envelope rather than the physical Standard Model sector.",
	}
}

func dot(a, b []float64) float64 {
	limit := len(a)
	if len(b) < limit {
		limit = len(b)
	}
	s := 0.0
	for i := 0; i < limit; i++ {
		s += a[i] * b[i]
	}
	return s
}

func vectorNorm(v []float64) float64 { return math.Sqrt(dot(v, v)) }

func normalize(v []float64) []float64 {
	n := vectorNorm(v)
	if n == 0 {
		return append([]float64(nil), v...)
	}
	return scaleVector(v, 1.0/n)
}

func scaleVector(v []float64, s float64) []float64 {
	out := make([]float64, len(v))
	for i, x := range v {
		out[i] = s * x
	}
	return out
}

func FormatVector(v []float64) string {
	parts := make([]string, len(v))
	for i, x := range v {
		parts[i] = fmt.Sprintf("%.12f", x)
	}
	return "[" + strings.Join(parts, ",") + "]"
}

func FormatBasis(b BasisAudit) string {
	return fmt.Sprintf("formalized=%t geometric=%v mass=%v tau=%s norm=%.12f source=%s distinction=%q verdict=%s", b.Formalized, b.GeometricBasis, b.MassBasis, FormatVector(b.TrialityTensor), b.TrialityNorm, FormatVector(b.NormalizedSource), b.Distinction, b.Verdict)
}

func FormatOperator(o OperatorSourceAudit) string {
	return fmt.Sprintf("sieve=%t JFlavor=%t doubledFlavor=%t bimoduleFlavor=%t nativeUnitary=%t allowed=%t why=%q verdict=%s", o.SieveFormalized, o.JSwapActsOnFlavor, o.DoubledSpaceActsOnFlavor, o.BimoduleOverlapActsOnFlavor, o.InstalledNativeUnitary, o.CandidateAllowedByUnitarity, o.WhyNativeMissing, o.Verdict)
}

func FormatCandidate(c FlavorCandidate) string {
	return fmt.Sprintf("name=%s vector=%s unit=%t overlap=%.12f fraction=%.12f native=%t unique=%t suppress=%t mass=%.6f nearGate322=%t verdict=%s", c.Name, FormatVector(c.TopVector), c.UnitNorm, c.TauOverlap, c.TopFraction, c.Native, c.Unique, c.SuppressesTopBoundary, c.InheritedMassGeV, c.NearGate322, c.Verdict)
}

func FormatNullspace(n NullspaceAudit) string {
	basis := make([]string, len(n.Basis))
	for i, v := range n.Basis {
		basis[i] = FormatVector(v)
	}
	return fmt.Sprintf("computed=%t dimension=%d basis=%s orthogonal=%t suppressPossible=%t uniqueTop=%t verdict=%s", n.Computed, n.Dimension, strings.Join(basis, ";"), n.AllBasisVectorsOrthogonal, n.TopSuppressionPossible, n.UniquePhysicalTopVector, n.Verdict)
}

func FormatRG(r RGCompatibilityAudit) string {
	return fmt.Sprintf("audited=%t gaugeMass=%.6f lowSlot=%.6f highSlot=%.6f nullMass=%.6f nullPreserves=%t native=%t physical=%t verdict=%s", r.Audited, r.Gate322GaugeOnlyMassGeV, r.IdentityLowSlotMassGeV, r.IdentityHighSlotMassGeV, r.NullTopMassGeV, r.NullRotationPreservesGate322, r.NativeJustificationInstalled, r.PhysicalLaneAuthorized, r.Verdict)
}

func FormatFirewalls(f FirewallAudit) string {
	return fmt.Sprintf("noCKM=%t noTopMass=%t noTexture=%t noPole=%t noTwoLoop=%t noCollider=%t polluted=%t verdict=%s", f.NoCKMImported, f.NoObservedTopMassInserted, f.NoFlavorTextureInvented, f.NoPoleMassClaimed, f.NoTwoLoopClaimed, f.NoColliderMassClaimed, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("basis=%t flavorSieve=%t null=%t nativeOperator=%t topJustified=%t gate322Physical=%t firewalls=%t finalClaim=%t status=%s answer=%q next=%q", s.BasisFormalized, s.FlavorSieveFormalized, s.NullspaceCapacityProved, s.NativeFlavorOperatorDerived, s.TopBoundarySuppressionJustified, s.Gate322PhysicalLaneAuthorized, s.FirewallsPreserved, s.FinalMassClaimed, s.Status, s.DirectAnswer, s.NextGate)
}
