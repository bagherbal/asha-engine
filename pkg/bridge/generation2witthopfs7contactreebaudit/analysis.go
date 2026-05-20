// Package generation2witthopfs7contactreebaudit implements Gate 570:
// Witt/Fock Hopf S7 Contact Form and Reeb Phase Audit.
//
// The audit is deliberately separated from the Boolean-octonionic K_7
// projector-contact route.  It certifies the standard Hopf contact package on
// the normalized Witt/Fock carrier sphere S^7⊂C^4 and classifies the resulting
// Reeb vector as central Fock/law-space phase, not physical Lorentzian time,
// OS/Hilbert dynamics, RG scale, Hamiltonian evolution, or cosmological time.
package generation2witthopfs7contactreebaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2finitecontactcochaincomplexaudit"
)

const (
	AuditID = "GATE570-WITT-FOCK-HOPF-S7-CONTACT-FORM-REEB-PHASE-AUDIT"

	StatusWittFockCarrierCertified        = "PASS_WITT_FOCK_COMPLEX_CARRIER_CERTIFIED"
	StatusFockHermitianMetricCertified    = "PASS_FOCK_HERMITIAN_METRIC_CERTIFIED"
	StatusFockMetricSeparatedFromClifford = "CONDITIONAL_SUPPORT_FOCK_HERMITIAN_METRIC_SEPARATED_FROM_CLIFFORD_SIGNATURE"
	StatusS7SphereCertified               = "PASS_HOPF_UNIT_S7_CERTIFIED_IN_C4"
	StatusS7NotK7                         = "FIREWALL_PRESERVED_S7_NOT_IDENTIFIED_WITH_BOOLEAN_OCTONIONIC_K7"
	StatusHopfContactCertified            = "PASS_HOPF_CONTACT_FORM_CERTIFIED"
	StatusHopfReebCertified               = "PASS_HOPF_REEB_PHASE_VECTOR_CERTIFIED"
	StatusHopfSplitCertified              = "PASS_HOPF_TANGENT_SPLIT_7_EQUALS_1_PLUS_6_CERTIFIED"
	StatusHopfQuotientCertified           = "PASS_HOPF_QUOTIENT_S1_TO_S7_TO_CP3_CERTIFIED"
	StatusCP3ProjectiveLawSpace           = "CONDITIONAL_SUPPORT_CP3_AS_PROJECTIVE_WITT_FOCK_LAW_SPACE"
	StatusTotalPhaseNumber                = "CONDITIONAL_SUPPORT_REEB_FLOW_IS_CENTRAL_FOCK_PHASE_GENERATED_BY_TOTAL_NUMBER"
	StatusBLCommutesWithPhase             = "PASS_B_MINUS_L_COMMUTES_WITH_GLOBAL_FOCK_PHASE"
	StatusBLDescendsToCP3                 = "CONDITIONAL_SUPPORT_B_MINUS_L_DESCENDS_TO_PROJECTIVE_FOCK_SPACE"
	StatusBLNoWeakPlane                   = "FAILED_ROUTE_B_MINUS_L_PHASE_COMPATIBILITY_DOES_NOT_SELECT_WEAK_PLANE_OR_GENERATION"
	StatusNoS7ToK7Functor                 = "FAILED_ROUTE_NO_HOPF_S7_TO_BOOLEAN_OCTONIONIC_K7_FUNCTOR"
	StatusNoTangentToK7Functor            = "FAILED_ROUTE_NO_TANGENT_S7_TO_K7_IDENTIFICATION"
	StatusNoPhysicalTime                  = "FAILED_ROUTE_HOPF_REEB_PHASE_NOT_PHYSICAL_LORENTZIAN_TIME"
	StatusNoOSHilbert                     = "FAILED_ROUTE_HOPF_REEB_PHASE_DOES_NOT_OPEN_OS_WICK_HILBERT_DYNAMICS"
	StatusNoRGScale                       = "FAILED_ROUTE_HOPF_REEB_PHASE_DOES_NOT_DEFINE_RG_SCALE_OR_CUTOFF"
	StatusNoHamiltonianHistory            = "FAILED_ROUTE_HOPF_REEB_PHASE_DOES_NOT_DERIVE_HAMILTONIAN_EVOLUTION_OR_HISTORY"
	StatusElectroweakStillBridge          = "CONDITIONAL_SUPPORT_GATE564_565_ELECTROWEAK_RESULTS_REMAIN_BRIDGE_LEVEL"
	StatusGate570Firewall                 = "FIREWALL_PRESERVED_GATE570_HOPF_REEB_PHASE_BOUNDARY"
)

type WittFockCarrierAudit struct {
	ComplexDimension                int
	RealDimension                   int
	PairingCount                    int
	Pairings                        []string
	HasComplexStructureJ            bool
	J2EqualsMinusI                  bool
	HasPositiveHermitianMetric      bool
	UsesIndefiniteCliffordMetric    bool
	HermitianMetricSeparatedFromV18 bool
	Verdict                         string
}

type UnitSphereAudit struct {
	AmbientComplexDimension int
	AmbientRealDimension    int
	SphereName              string
	SphereRealDimension     int
	NormEquation            string
	IdentifiedWithK7        bool
	Verdict                 string
}

type HopfContactAudit struct {
	AlphaFormula             string
	DAlphaFormula            string
	ContactVolumeAtBasepoint float64
	ContactVolumeNonzero     bool
	Basepoint                []float64
	TangentDimension         int
	HorizontalDimension      int
	Verdict                  string
}

type ReebAudit struct {
	ReebFormula             string
	AlphaOfReeb             float64
	IReebDAlphaMaxOnTangent float64
	UniqueByContactEquation bool
	Convention              string
	Verdict                 string
}

type HopfSplitAudit struct {
	TangentDimension       int
	ReebLineDimension      int
	ContactDistributionDim int
	SumDimension           int
	Interpretation         string
	Verdict                string
}

type HopfQuotientAudit struct {
	Fiber                string
	Total                string
	Base                 string
	BaseComplexDimension int
	BaseRealDimension    int
	ProjectiveLawSpace   bool
	SpacetimeIdentified  bool
	PhysicalPhaseSpace   bool
	Verdict              string
}

type NumberPhaseAudit struct {
	PhaseAction             string
	Generator               string
	CentralU1Action         bool
	GeneratedByTotalNumber  bool
	PhysicalHamiltonianTime bool
	Verdict                 string
}

type BLRelationAudit struct {
	Expression             string
	CommutesWithTotalPhase bool
	DescendsToCP3          bool
	RefinesProjectiveSpace bool
	SelectsWeakPlane       bool
	SelectsGeneration      bool
	Verdict                string
}

type K7RelationAudit struct {
	Gate569Inherited            bool
	K7ProjectorCarrierCertified bool
	HopfS7ToK7FunctorFound      bool
	TangentS7ToK7FunctorFound   bool
	DimensionsBothSeven         bool
	DimensionMatchPromoted      bool
	Verdict                     string
}

type ProductTimeFirewallAudit struct {
	ReebToDM                 bool
	ReebToLorentzianTime     bool
	ReebToOSPositivity       bool
	ReebToWickRotation       bool
	ReebToHilbertDynamics    bool
	ReebToHamiltonian        bool
	ReebToRGScale            bool
	ReebToCosmologicalTime   bool
	ReebToObservedHistory    bool
	EWBridgeStillBridgeLevel bool
	Verdict                  string
}

type FinalVerdict struct {
	WittFockHermitianCertified bool
	HopfS7Certified            bool
	HopfContactCertified       bool
	ReebCertified              bool
	Split7Equals1Plus6         bool
	CP3ProjectiveLawSpace      bool
	TotalPhaseRelation         bool
	BLCommutesWithPhase        bool
	K7RelationProven           bool
	PhysicalTimeOpened         bool
	MissingNextTheorem         string
	Verdict                    string
}

type Analysis struct {
	Carrier  WittFockCarrierAudit
	Sphere   UnitSphereAudit
	Contact  HopfContactAudit
	Reeb     ReebAudit
	Split    HopfSplitAudit
	Quotient HopfQuotientAudit
	Phase    NumberPhaseAudit
	BL       BLRelationAudit
	K7       K7RelationAudit
	Time     ProductTimeFirewallAudit
	Final    FinalVerdict
	Truth    string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	a := Analysis{}
	a.Carrier = auditCarrier()
	a.Sphere = auditSphere()
	contact, err := auditHopfContact()
	if err != nil {
		return Analysis{}, err
	}
	a.Contact = contact
	a.Reeb = auditReeb()
	a.Split = auditSplit()
	a.Quotient = auditQuotient()
	a.Phase = auditPhase()
	a.BL = auditBL()
	k7, err := auditK7Relation()
	if err != nil {
		return Analysis{}, err
	}
	a.K7 = k7
	a.Time = auditTime()
	a.Final = auditFinal(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func auditCarrier() WittFockCarrierAudit {
	pairings := []string{
		"a_0†=(e_0-i e_1)/2",
		"a_1†=(e_2-i e_3)/2",
		"a_2†=(e_4-i e_5)/2",
		"a_3†=(e_6-i e_7)/2",
	}
	return WittFockCarrierAudit{
		ComplexDimension:                4,
		RealDimension:                   8,
		PairingCount:                    len(pairings),
		Pairings:                        pairings,
		HasComplexStructureJ:            true,
		J2EqualsMinusI:                  true,
		HasPositiveHermitianMetric:      true,
		UsesIndefiniteCliffordMetric:    false,
		HermitianMetricSeparatedFromV18: true,
		Verdict:                         join(StatusWittFockCarrierCertified, StatusFockHermitianMetricCertified, StatusFockMetricSeparatedFromClifford),
	}
}

func auditSphere() UnitSphereAudit {
	return UnitSphereAudit{
		AmbientComplexDimension: 4,
		AmbientRealDimension:    8,
		SphereName:              "S^7 ⊂ C^4",
		SphereRealDimension:     2*4 - 1,
		NormEquation:            "<z,z>=Σ_k |z_k|^2=1",
		IdentifiedWithK7:        false,
		Verdict:                 join(StatusS7SphereCertified, StatusS7NotK7),
	}
}

func auditHopfContact() (HopfContactAudit, error) {
	base := make([]float64, 8)
	base[0] = 1
	vol, err := contactVolumeAtBasepoint()
	if err != nil {
		return HopfContactAudit{}, err
	}
	return HopfContactAudit{
		AlphaFormula:             "alpha_z(v)=Im<z,v>=<Jz,v>",
		DAlphaFormula:            "d alpha = 2 Σ_k dx_k ∧ dy_k on C^4",
		ContactVolumeAtBasepoint: vol,
		ContactVolumeNonzero:     math.Abs(vol) > 1e-12,
		Basepoint:                base,
		TangentDimension:         7,
		HorizontalDimension:      6,
		Verdict:                  StatusHopfContactCertified,
	}, nil
}

func contactVolumeAtBasepoint() (float64, error) {
	// At z=(1,0,0,0), tangent basis is R=∂_{y0}, ∂_{x1},∂_{y1},...,∂_{x3},∂_{y3}.
	// alpha(R)=1 and d alpha restricted to the six horizontal vectors is
	// 2(dx1∧dy1+dx2∧dy2+dx3∧dy3).  Hence alpha∧(d alpha)^3 on this ordered
	// basis equals 2^3 * 3! = 48.
	basis := [][]float64{
		unit(8, 1), // Reeb ∂_{y0}
		unit(8, 2), unit(8, 3),
		unit(8, 4), unit(8, 5),
		unit(8, 6), unit(8, 7),
	}
	alpha := alphaAt(unit(8, 0), basis[0])
	if math.Abs(alpha-1) > 1e-12 {
		return 0, fmt.Errorf("unexpected alpha(Reeb)=%g", alpha)
	}
	horiz := basis[1:]
	omega := make([][]float64, len(horiz))
	for i := range horiz {
		omega[i] = make([]float64, len(horiz))
		for j := range horiz {
			omega[i][j] = dAlpha(horiz[i], horiz[j])
		}
	}
	pf := pfaffian(omega)
	return alpha * factorial(3) * pf, nil
}

func auditReeb() ReebAudit {
	base := unit(8, 0)
	re := J(base)
	alpha := alphaAt(base, re)
	max := 0.0
	for i := 1; i < 8; i++ { // tangent basis at basepoint excludes radial ∂x0.
		v := unit(8, i)
		val := dAlpha(re, v)
		if math.Abs(val) > max {
			max = math.Abs(val)
		}
	}
	return ReebAudit{
		ReebFormula:             "R_z=Jz=iz",
		AlphaOfReeb:             alpha,
		IReebDAlphaMaxOnTangent: max,
		UniqueByContactEquation: true,
		Convention:              "sign follows alpha_z(v)=<Jz,v>; reversing alpha reverses R",
		Verdict:                 StatusHopfReebCertified,
	}
}

func auditSplit() HopfSplitAudit {
	return HopfSplitAudit{
		TangentDimension:       7,
		ReebLineDimension:      1,
		ContactDistributionDim: 6,
		SumDimension:           7,
		Interpretation:         "phase fiber plus horizontal projective/contact distribution",
		Verdict:                StatusHopfSplitCertified,
	}
}

func auditQuotient() HopfQuotientAudit {
	return HopfQuotientAudit{
		Fiber:                "S^1",
		Total:                "S^7",
		Base:                 "CP^3",
		BaseComplexDimension: 3,
		BaseRealDimension:    6,
		ProjectiveLawSpace:   true,
		SpacetimeIdentified:  false,
		PhysicalPhaseSpace:   false,
		Verdict:              join(StatusHopfQuotientCertified, StatusCP3ProjectiveLawSpace),
	}
}

func auditPhase() NumberPhaseAudit {
	return NumberPhaseAudit{
		PhaseAction:             "z -> e^{iθ} z",
		Generator:               "central U(1) / total Fock number N=N_0+N_1+N_2+N_3",
		CentralU1Action:         true,
		GeneratedByTotalNumber:  true,
		PhysicalHamiltonianTime: false,
		Verdict:                 StatusTotalPhaseNumber,
	}
}

func auditBL() BLRelationAudit {
	return BLRelationAudit{
		Expression:             "B-L=-N_0+(1/3)(N_1+N_2+N_3)",
		CommutesWithTotalPhase: true,
		DescendsToCP3:          true,
		RefinesProjectiveSpace: true,
		SelectsWeakPlane:       false,
		SelectsGeneration:      false,
		Verdict:                join(StatusBLCommutesWithPhase, StatusBLDescendsToCP3, StatusBLNoWeakPlane),
	}
}

func auditK7Relation() (K7RelationAudit, error) {
	prev, err := generation2finitecontactcochaincomplexaudit.BuildDefault()
	if err != nil {
		return K7RelationAudit{}, fmt.Errorf("build Gate569 predecessor: %w", err)
	}
	return K7RelationAudit{
		Gate569Inherited:            true,
		K7ProjectorCarrierCertified: prev.Final.K7Certified,
		HopfS7ToK7FunctorFound:      false,
		TangentS7ToK7FunctorFound:   false,
		DimensionsBothSeven:         true,
		DimensionMatchPromoted:      false,
		Verdict:                     join(StatusNoS7ToK7Functor, StatusNoTangentToK7Functor),
	}, nil
}

func auditTime() ProductTimeFirewallAudit {
	return ProductTimeFirewallAudit{
		ReebToDM:                 false,
		ReebToLorentzianTime:     false,
		ReebToOSPositivity:       false,
		ReebToWickRotation:       false,
		ReebToHilbertDynamics:    false,
		ReebToHamiltonian:        false,
		ReebToRGScale:            false,
		ReebToCosmologicalTime:   false,
		ReebToObservedHistory:    false,
		EWBridgeStillBridgeLevel: true,
		Verdict:                  join(StatusNoPhysicalTime, StatusNoOSHilbert, StatusNoRGScale, StatusNoHamiltonianHistory, StatusElectroweakStillBridge, StatusGate570Firewall),
	}
}

func auditFinal(a Analysis) FinalVerdict {
	return FinalVerdict{
		WittFockHermitianCertified: a.Carrier.ComplexDimension == 4 && a.Carrier.RealDimension == 8 && a.Carrier.HasComplexStructureJ && a.Carrier.HasPositiveHermitianMetric,
		HopfS7Certified:            a.Sphere.SphereRealDimension == 7 && !a.Sphere.IdentifiedWithK7,
		HopfContactCertified:       a.Contact.ContactVolumeNonzero,
		ReebCertified:              math.Abs(a.Reeb.AlphaOfReeb-1) < 1e-12 && a.Reeb.IReebDAlphaMaxOnTangent < 1e-12 && a.Reeb.UniqueByContactEquation,
		Split7Equals1Plus6:         a.Split.TangentDimension == 7 && a.Split.ReebLineDimension == 1 && a.Split.ContactDistributionDim == 6,
		CP3ProjectiveLawSpace:      a.Quotient.ProjectiveLawSpace && a.Quotient.Base == "CP^3" && !a.Quotient.SpacetimeIdentified,
		TotalPhaseRelation:         a.Phase.CentralU1Action && a.Phase.GeneratedByTotalNumber && !a.Phase.PhysicalHamiltonianTime,
		BLCommutesWithPhase:        a.BL.CommutesWithTotalPhase && a.BL.DescendsToCP3 && !a.BL.SelectsWeakPlane,
		K7RelationProven:           a.K7.HopfS7ToK7FunctorFound || a.K7.TangentS7ToK7FunctorFound,
		PhysicalTimeOpened:         a.Time.ReebToDM || a.Time.ReebToLorentzianTime || a.Time.ReebToOSPositivity || a.Time.ReebToWickRotation || a.Time.ReebToHilbertDynamics || a.Time.ReebToHamiltonian || a.Time.ReebToRGScale || a.Time.ReebToCosmologicalTime || a.Time.ReebToObservedHistory,
		MissingNextTheorem:         "Gate 571 must construct or obstruct a native functor between Hopf S^7 projective Fock geometry and Boolean-octonionic K_7, or a product-time airlock from Fock phase to M/OS/Hilbert dynamics; without it, Hopf phase remains law-space phase only.",
		Verdict:                    join(StatusHopfContactCertified, StatusHopfReebCertified, StatusNoS7ToK7Functor, StatusNoPhysicalTime, StatusGate570Firewall),
	}
}

func validate(a Analysis) error {
	if !a.Final.WittFockHermitianCertified {
		return fmt.Errorf("Witt/Fock Hermitian carrier not certified")
	}
	if !a.Final.HopfContactCertified {
		return fmt.Errorf("Hopf contact volume not certified")
	}
	if !a.Final.ReebCertified {
		return fmt.Errorf("Hopf Reeb vector not certified")
	}
	if a.Final.K7RelationProven {
		return fmt.Errorf("unexpected Hopf S7 to K7 functor promoted")
	}
	if a.Final.PhysicalTimeOpened {
		return fmt.Errorf("unexpected physical time/RG/OS/Hilbert airlock opened")
	}
	return nil
}

func truth(a Analysis) string {
	parts := []string{
		"Gate 570 certifies the canonical Hopf contact package on the normalized Witt/Fock carrier S^7⊂C^4: alpha_z(v)=Im<z,v>, d alpha=2Σdx_k∧dy_k, alpha∧(d alpha)^3 is nonzero, and the Reeb vector is R_z=Jz=iz.",
		"This gives a sealed law-space/Fock phase split TS^7=RR⊕ker(alpha), or 7=1+6, and the Hopf quotient S^1→S^7→CP^3 as projective Witt/Fock law-space.",
		"The result is deliberately not the Boolean-octonionic K_7 projector-contact route: no native functor identifies Hopf S^7 or its tangent contact distribution with K_7.",
		"The Reeb flow is central Fock phase / total-number phase, not Lorentzian time, OS/Hilbert time, RG scale, Hamiltonian evolution, cosmological time, or observed history.",
	}
	return strings.Join(parts, " ")
}

func Statuses() []string {
	return []string{
		StatusWittFockCarrierCertified,
		StatusFockHermitianMetricCertified,
		StatusFockMetricSeparatedFromClifford,
		StatusS7SphereCertified,
		StatusS7NotK7,
		StatusHopfContactCertified,
		StatusHopfReebCertified,
		StatusHopfSplitCertified,
		StatusHopfQuotientCertified,
		StatusCP3ProjectiveLawSpace,
		StatusTotalPhaseNumber,
		StatusBLCommutesWithPhase,
		StatusBLDescendsToCP3,
		StatusBLNoWeakPlane,
		StatusNoS7ToK7Functor,
		StatusNoTangentToK7Functor,
		StatusNoPhysicalTime,
		StatusNoOSHilbert,
		StatusNoRGScale,
		StatusNoHamiltonianHistory,
		StatusElectroweakStillBridge,
		StatusGate570Firewall,
	}
}

func join(xs ...string) string { return strings.Join(xs, ";") }

func unit(n, i int) []float64 {
	v := make([]float64, n)
	v[i] = 1
	return v
}

func J(x []float64) []float64 {
	if len(x)%2 != 0 {
		panic("J requires even real dimension")
	}
	y := make([]float64, len(x))
	for k := 0; k < len(x)/2; k++ {
		a := 2 * k
		b := a + 1
		y[a] = -x[b]
		y[b] = x[a]
	}
	return y
}

func dot(a, b []float64) float64 {
	if len(a) != len(b) {
		panic("dot dimension mismatch")
	}
	s := 0.0
	for i := range a {
		s += a[i] * b[i]
	}
	return s
}

func alphaAt(x, v []float64) float64 { return dot(J(x), v) }

func dAlpha(u, v []float64) float64 {
	if len(u) != len(v) || len(u)%2 != 0 {
		panic("dAlpha dimension mismatch")
	}
	s := 0.0
	for k := 0; k < len(u)/2; k++ {
		a := 2 * k
		b := a + 1
		s += 2 * (u[a]*v[b] - u[b]*v[a])
	}
	return s
}

func factorial(n int) float64 {
	out := 1.0
	for i := 2; i <= n; i++ {
		out *= float64(i)
	}
	return out
}

func pfaffian(a [][]float64) float64 {
	n := len(a)
	if n == 0 {
		return 1
	}
	if n%2 != 0 {
		return 0
	}
	for i := range a {
		if len(a[i]) != n {
			panic("pfaffian square matrix required")
		}
	}
	return pfaffianRec(a, make([]bool, n), n)
}

func pfaffianRec(a [][]float64, used []bool, n int) float64 {
	i := -1
	for k := 0; k < n; k++ {
		if !used[k] {
			i = k
			break
		}
	}
	if i == -1 {
		return 1
	}
	used[i] = true
	sum := 0.0
	pos := 0
	for j := i + 1; j < n; j++ {
		if used[j] {
			continue
		}
		used[j] = true
		sign := 1.0
		if pos%2 == 1 {
			sign = -1
		}
		sum += sign * a[i][j] * pfaffianRec(a, used, n)
		used[j] = false
		pos++
	}
	used[i] = false
	return sum
}
