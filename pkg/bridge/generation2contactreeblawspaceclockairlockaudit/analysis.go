// Package generation2contactreeblawspaceclockairlockaudit implements Gate 566:
// Contact/Reeb Law-Space Clock and Product-Time Airlock Audit.
//
// This gate audits whether the finite contact vacuum K_7 already carries the
// explicit contact package (alpha, d alpha, Reeb vector, contact volume) needed
// for a native law-space clock-flow. It deliberately separates such a possible
// law-space flow from continuum time in the M factor of M x F, OS/Wick/Hilbert
// reconstruction, RG scale, and observed history.
package generation2contactreeblawspaceclockairlockaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
)

const (
	AuditID = "GATE566-CONTACT-REEB-LAW-SPACE-CLOCK-PRODUCT-TIME-AIRLOCK-AUDIT"

	StatusContactK7Certified         = "PASS_CONTACT_K7_PROJECTOR_AND_DIMENSION_CERTIFIED"
	StatusNoContactForm              = "FAILED_ROUTE_NO_EXPLICIT_CONTACT_FORM_ON_K7"
	StatusNoReebVector               = "FAILED_ROUTE_NO_REEB_VECTOR_CERTIFICATE"
	StatusNoContactOrientationVolume = "FAILED_ROUTE_NO_CONTACT_ALPHA_WEDGE_DALPHA_VOLUME_CERTIFICATE"
	StatusBooleanG2OrientationOnly   = "CONDITIONAL_SUPPORT_BOOLEAN_OCTONIONIC_CONTACT_PROJECTOR_ORIENTATION_DATA_ONLY"
	StatusE0ReebSeparated            = "FAILED_ROUTE_NO_CANONICAL_E0_TO_REEB_RELATION"
	StatusQ4ContactSpectralOnly      = "CONDITIONAL_SUPPORT_Q4_REMAINS_INDEPENDENT_CONTACT_SPECTRAL_DATA"
	StatusQ4NotReebDynamics          = "FAILED_ROUTE_Q4_NOT_CERTIFIED_AS_REEB_FLOW_OR_CONTACT_RETURN_SPECTRUM"
	StatusProductTimeAirlockBlocked  = "FAILED_ROUTE_NO_CONTACT_TO_PHYSICAL_TIME_AIRLOCK"
	StatusModularStillNeedsState     = "FAILED_ROUTE_CONTACT_REEB_AUDIT_DOES_NOT_SOLVE_TRACIAL_MODULAR_TIME_OBSTRUCTION"
	StatusNoRGScale                  = "FAILED_ROUTE_REEB_CONTACT_FLOW_DOES_NOT_DERIVE_RG_SCALE_OR_CUTOFF"
	StatusEWBridgeStillSealed        = "FIREWALL_PRESERVED_GATE564_GATE565_ELECTROWEAK_BRIDGE_LEVEL"
	StatusPhysicalTimeFirewall       = "FIREWALL_PRESERVED_GATE566_CONTACT_LAW_SPACE_CLOCK_PRODUCT_TIME_BOUNDARY"
)

type ContactPackageAudit struct {
	K7Dimension                  int
	ExpectedDimension            int
	AmbientDimension             int
	ContactIndex                 float64
	FrameIsometryResidual        float64
	BooleanContainmentResidual   float64
	G2ContainmentResidual        float64
	ProjectorExists              bool
	AlphaAvailable               bool
	DAlphaAvailable              bool
	ContactVolumeComputable      bool
	AlphaWedgeDAlphaCubedNonzero bool
	Verdict                      string
}

type ReebVectorAudit struct {
	Definition              string
	AlphaOfRCondition       string
	ContractionCondition    string
	AlphaAndDAlphaAvailable bool
	ReebVectorAvailable     bool
	ReebUnique              bool
	Split7As1Plus6          bool
	Verdict                 string
}

type OrientationVolumeAudit struct {
	AlphaVolumeAvailable              bool
	NativeContactOrientationFromAlpha bool
	BooleanOctonionicProjectorData    bool
	PhysicalSpacetimeOrientationClaim bool
	Verdict                           string
}

type CliffordSignatureRelationAudit struct {
	FiniteCarrierSignature string
	E0NativeSignatureDatum bool
	ReebLawSpaceFlowDatum  bool
	CanonicalE0ToReebMap   bool
	PhysicalTimeInProductM bool
	SeparationPreserved    bool
	Verdict                string
}

type ContactQuarticRelationAudit struct {
	Q4Polynomial                 string
	ContactSectorData            bool
	ReebFlowSpectrumCertified    bool
	ContactEndomorphismSpectrum  bool
	LinearizedReturnMapCertified bool
	HiggsFlavorYukawaPromotion   bool
	Verdict                      string
}

type ProductTimeAirlockAudit struct {
	ProductGeometryAvailable       bool
	DTotalForm                     string
	ContactToDMMap                 bool
	ContactToLorentzianSignature   bool
	ContactToOSPositivity          bool
	ContactToWickRotation          bool
	ContactToHilbertReconstruction bool
	ContactToHamiltonianSpectrum   bool
	ContactToUnitaryDynamics       bool
	ContactToGlobalCausality       bool
	ContactToArrowOfTime           bool
	Verdict                        string
}

type ModularTimeComparisonAudit struct {
	PreviousModularRouteKnown         bool
	TracialStateObstructionKnown      bool
	ContactReebAvoidsObstruction      bool
	NontracialStateInserted           bool
	StillNeedsNontracialStateOrKernel bool
	Verdict                           string
}

type RGScaleFirewallAudit struct {
	ReebGivesRGScale      bool
	ReebGivesCutoffLambda bool
	ReebGivesFMoments     bool
	ReebGivesPhysicalTime bool
	Verdict               string
}

type ElectroweakRelationAudit struct {
	Gate564SymbolicHessianBridgeOnly bool
	Gate565BoundaryNormalizationOnly bool
	PhysicalWZPhotonDynamicsDerived  bool
	OSWickHilbertDynamicsDerived     bool
	ObservedDataImported             bool
	Verdict                          string
}

type FinalVerdict struct {
	ExplicitContactFormAlpha     bool
	CertifiedReebVector          bool
	K7Splits1Plus6               bool
	RRelatedToE0OrPhysicalTime   bool
	Q4PartOfReebDynamics         bool
	ContactToPhysicalTimeAirlock bool
	RGScaleOSHilbertOpened       bool
	MissingNextTheorem           string
	Verdict                      string
}

type Analysis struct {
	Contact     ContactPackageAudit
	Reeb        ReebVectorAudit
	Orientation OrientationVolumeAudit
	Signature   CliffordSignatureRelationAudit
	Quartic     ContactQuarticRelationAudit
	ProductTime ProductTimeAirlockAudit
	Modular     ModularTimeComparisonAudit
	RGScale     RGScaleFirewallAudit
	Electroweak ElectroweakRelationAudit
	Final       FinalVerdict
	Truth       string
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
	space, err := contact.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build contact K_7 space: %w", err)
	}
	c, err := auditContactPackage(space)
	if err != nil {
		return Analysis{}, err
	}
	a := Analysis{}
	a.Contact = c
	a.Reeb = auditReeb(c)
	a.Orientation = auditOrientation(c)
	a.Signature = auditSignatureRelation()
	a.Quartic = auditQuartic()
	a.ProductTime = auditProductTime()
	a.Modular = auditModular()
	a.RGScale = auditRGScale()
	a.Electroweak = auditElectroweak()
	a.Final = auditFinal(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func auditContactPackage(space contact.Space) (ContactPackageAudit, error) {
	frameResidual, err := space.FrameIsometryResidual()
	if err != nil {
		return ContactPackageAudit{}, fmt.Errorf("contact frame residual: %w", err)
	}
	booleanResidual, err := space.BooleanContainmentResidual()
	if err != nil {
		return ContactPackageAudit{}, fmt.Errorf("Boolean containment residual: %w", err)
	}
	g2Residual, err := space.G2ContainmentResidual()
	if err != nil {
		return ContactPackageAudit{}, fmt.Errorf("G2 containment residual: %w", err)
	}
	projectorResidual, err := space.ContactProjector.IdempotenceResidual()
	if err != nil {
		return ContactPackageAudit{}, fmt.Errorf("contact projector residual: %w", err)
	}
	ok := space.Dimension() == 7 && space.Dimension() == space.ExpectedContactDenominator() && math.Abs(space.ContactIndex()-1) < 1e-8 && projectorResidual < 1e-8 && frameResidual < 1e-8 && booleanResidual < 1e-8 && g2Residual < 1e-8
	return ContactPackageAudit{
		K7Dimension:                  space.Dimension(),
		ExpectedDimension:            space.ExpectedContactDenominator(),
		AmbientDimension:             space.AmbientDimension(),
		ContactIndex:                 space.ContactIndex(),
		FrameIsometryResidual:        frameResidual,
		BooleanContainmentResidual:   booleanResidual,
		G2ContainmentResidual:        g2Residual,
		ProjectorExists:              ok,
		AlphaAvailable:               false,
		DAlphaAvailable:              false,
		ContactVolumeComputable:      false,
		AlphaWedgeDAlphaCubedNonzero: false,
		Verdict:                      join(StatusContactK7Certified, StatusNoContactForm),
	}, nil
}

func auditReeb(c ContactPackageAudit) ReebVectorAudit {
	return ReebVectorAudit{
		Definition:              "For a contact form alpha on K_7, the Reeb vector R satisfies alpha(R)=1 and i_R d alpha=0",
		AlphaOfRCondition:       "unavailable: alpha is not present as a certified one-form on K_7",
		ContractionCondition:    "unavailable: d alpha and contraction on K_7 are not present",
		AlphaAndDAlphaAvailable: c.AlphaAvailable && c.DAlphaAvailable,
		ReebVectorAvailable:     false,
		ReebUnique:              false,
		Split7As1Plus6:          false,
		Verdict:                 join(StatusNoContactForm, StatusNoReebVector),
	}
}

func auditOrientation(c ContactPackageAudit) OrientationVolumeAudit {
	return OrientationVolumeAudit{
		AlphaVolumeAvailable:              c.ContactVolumeComputable,
		NativeContactOrientationFromAlpha: false,
		BooleanOctonionicProjectorData:    c.ProjectorExists,
		PhysicalSpacetimeOrientationClaim: false,
		Verdict:                           join(StatusNoContactOrientationVolume, StatusBooleanG2OrientationOnly),
	}
}

func auditSignatureRelation() CliffordSignatureRelationAudit {
	return CliffordSignatureRelationAudit{
		FiniteCarrierSignature: "V=R^(1,7) supplies finite Clifford signature data with an e_0 slot",
		E0NativeSignatureDatum: true,
		ReebLawSpaceFlowDatum:  false,
		CanonicalE0ToReebMap:   false,
		PhysicalTimeInProductM: true,
		SeparationPreserved:    true,
		Verdict:                join(StatusE0ReebSeparated, "e_0, a possible Reeb law-space flow, and physical time in M remain distinct typed objects"),
	}
}

func auditQuartic() ContactQuarticRelationAudit {
	return ContactQuarticRelationAudit{
		Q4Polynomial:                 "q4(x)=3240x^4-7668x^3+6426x^2-2235x+271",
		ContactSectorData:            true,
		ReebFlowSpectrumCertified:    false,
		ContactEndomorphismSpectrum:  false,
		LinearizedReturnMapCertified: false,
		HiggsFlavorYukawaPromotion:   false,
		Verdict:                      join(StatusQ4ContactSpectralOnly, StatusQ4NotReebDynamics),
	}
}

func auditProductTime() ProductTimeAirlockAudit {
	return ProductTimeAirlockAudit{
		ProductGeometryAvailable:       true,
		DTotalForm:                     "D_total = D_M ⊗ 1_F + gamma_5 ⊗ D_F",
		ContactToDMMap:                 false,
		ContactToLorentzianSignature:   false,
		ContactToOSPositivity:          false,
		ContactToWickRotation:          false,
		ContactToHilbertReconstruction: false,
		ContactToHamiltonianSpectrum:   false,
		ContactToUnitaryDynamics:       false,
		ContactToGlobalCausality:       false,
		ContactToArrowOfTime:           false,
		Verdict:                        join(StatusProductTimeAirlockBlocked, StatusPhysicalTimeFirewall),
	}
}

func auditModular() ModularTimeComparisonAudit {
	return ModularTimeComparisonAudit{
		PreviousModularRouteKnown:         true,
		TracialStateObstructionKnown:      true,
		ContactReebAvoidsObstruction:      false,
		NontracialStateInserted:           false,
		StillNeedsNontracialStateOrKernel: true,
		Verdict:                           join(StatusModularStillNeedsState, "a missing Reeb certificate cannot replace the nontracial KMS/modular kernel required by previous gates"),
	}
}

func auditRGScale() RGScaleFirewallAudit {
	return RGScaleFirewallAudit{
		ReebGivesRGScale:      false,
		ReebGivesCutoffLambda: false,
		ReebGivesFMoments:     false,
		ReebGivesPhysicalTime: false,
		Verdict:               StatusNoRGScale,
	}
}

func auditElectroweak() ElectroweakRelationAudit {
	return ElectroweakRelationAudit{
		Gate564SymbolicHessianBridgeOnly: true,
		Gate565BoundaryNormalizationOnly: true,
		PhysicalWZPhotonDynamicsDerived:  false,
		OSWickHilbertDynamicsDerived:     false,
		ObservedDataImported:             false,
		Verdict:                          StatusEWBridgeStillSealed,
	}
}

func auditFinal(a Analysis) FinalVerdict {
	return FinalVerdict{
		ExplicitContactFormAlpha:     a.Contact.AlphaAvailable,
		CertifiedReebVector:          a.Reeb.ReebVectorAvailable,
		K7Splits1Plus6:               a.Reeb.Split7As1Plus6,
		RRelatedToE0OrPhysicalTime:   a.Signature.CanonicalE0ToReebMap || a.ProductTime.ContactToDMMap || a.ProductTime.ContactToArrowOfTime,
		Q4PartOfReebDynamics:         a.Quartic.ReebFlowSpectrumCertified || a.Quartic.ContactEndomorphismSpectrum || a.Quartic.LinearizedReturnMapCertified,
		ContactToPhysicalTimeAirlock: a.ProductTime.ContactToDMMap || a.ProductTime.ContactToOSPositivity || a.ProductTime.ContactToWickRotation || a.ProductTime.ContactToHilbertReconstruction,
		RGScaleOSHilbertOpened:       a.RGScale.ReebGivesRGScale || a.RGScale.ReebGivesCutoffLambda || a.ProductTime.ContactToOSPositivity || a.ProductTime.ContactToHilbertReconstruction,
		MissingNextTheorem:           "Construct an explicit contact form alpha on K_7, compute d alpha, prove alpha∧(d alpha)^3≠0, solve the unique Reeb vector R, then separately build a product-time airlock to M/OS/Wick/Hilbert before any physical-time claim is allowed",
		Verdict:                      join(StatusNoContactForm, StatusNoReebVector, StatusProductTimeAirlockBlocked, StatusPhysicalTimeFirewall),
	}
}

func validate(a Analysis) error {
	failures := []string{}
	if !a.Contact.ProjectorExists || a.Contact.K7Dimension != 7 || a.Contact.AlphaAvailable || a.Contact.DAlphaAvailable || a.Contact.ContactVolumeComputable || a.Contact.AlphaWedgeDAlphaCubedNonzero {
		failures = append(failures, "contact package audit failed")
	}
	if a.Reeb.AlphaAndDAlphaAvailable || a.Reeb.ReebVectorAvailable || a.Reeb.ReebUnique || a.Reeb.Split7As1Plus6 {
		failures = append(failures, "Reeb audit failed")
	}
	if a.Orientation.AlphaVolumeAvailable || a.Orientation.NativeContactOrientationFromAlpha || !a.Orientation.BooleanOctonionicProjectorData || a.Orientation.PhysicalSpacetimeOrientationClaim {
		failures = append(failures, "orientation audit failed")
	}
	if !a.Signature.E0NativeSignatureDatum || a.Signature.ReebLawSpaceFlowDatum || a.Signature.CanonicalE0ToReebMap || !a.Signature.PhysicalTimeInProductM || !a.Signature.SeparationPreserved {
		failures = append(failures, "signature relation audit failed")
	}
	if !a.Quartic.ContactSectorData || a.Quartic.ReebFlowSpectrumCertified || a.Quartic.ContactEndomorphismSpectrum || a.Quartic.LinearizedReturnMapCertified || a.Quartic.HiggsFlavorYukawaPromotion {
		failures = append(failures, "quartic relation audit failed")
	}
	if !a.ProductTime.ProductGeometryAvailable || a.ProductTime.ContactToDMMap || a.ProductTime.ContactToLorentzianSignature || a.ProductTime.ContactToOSPositivity || a.ProductTime.ContactToWickRotation || a.ProductTime.ContactToHilbertReconstruction || a.ProductTime.ContactToHamiltonianSpectrum || a.ProductTime.ContactToUnitaryDynamics || a.ProductTime.ContactToGlobalCausality || a.ProductTime.ContactToArrowOfTime {
		failures = append(failures, "product-time airlock audit failed")
	}
	if !a.Modular.PreviousModularRouteKnown || !a.Modular.TracialStateObstructionKnown || a.Modular.ContactReebAvoidsObstruction || a.Modular.NontracialStateInserted || !a.Modular.StillNeedsNontracialStateOrKernel {
		failures = append(failures, "modular comparison audit failed")
	}
	if a.RGScale.ReebGivesRGScale || a.RGScale.ReebGivesCutoffLambda || a.RGScale.ReebGivesFMoments || a.RGScale.ReebGivesPhysicalTime {
		failures = append(failures, "RG/scale firewall failed")
	}
	if !a.Electroweak.Gate564SymbolicHessianBridgeOnly || !a.Electroweak.Gate565BoundaryNormalizationOnly || a.Electroweak.PhysicalWZPhotonDynamicsDerived || a.Electroweak.OSWickHilbertDynamicsDerived || a.Electroweak.ObservedDataImported {
		failures = append(failures, "electroweak relation firewall failed")
	}
	if a.Final.ExplicitContactFormAlpha || a.Final.CertifiedReebVector || a.Final.K7Splits1Plus6 || a.Final.RRelatedToE0OrPhysicalTime || a.Final.Q4PartOfReebDynamics || a.Final.ContactToPhysicalTimeAirlock || a.Final.RGScaleOSHilbertOpened {
		failures = append(failures, "final verdict failed")
	}
	if len(failures) > 0 {
		return fmt.Errorf(strings.Join(failures, "; "))
	}
	return nil
}

func Statuses() []string {
	return []string{
		StatusContactK7Certified,
		StatusNoContactForm,
		StatusNoReebVector,
		StatusNoContactOrientationVolume,
		StatusBooleanG2OrientationOnly,
		StatusE0ReebSeparated,
		StatusQ4ContactSpectralOnly,
		StatusQ4NotReebDynamics,
		StatusProductTimeAirlockBlocked,
		StatusModularStillNeedsState,
		StatusNoRGScale,
		StatusEWBridgeStillSealed,
		StatusPhysicalTimeFirewall,
	}
}

func truth(a Analysis) string {
	return join(
		"Gate 566 certifies the finite Boolean-octonionic K_7 projector and contact index, but finds no explicit contact form alpha on K_7",
		"without alpha and d alpha, no Reeb vector, no 7=1+6 Reeb/contact split, and no alpha∧(d alpha)^3 orientation theorem can be promoted",
		"the original e_0 signature datum, possible contact law-space flow, and physical continuum time in M remain separate typed objects",
		"q4 remains independent contact spectral data rather than Reeb dynamics, and no airlock connects contact flow to OS/Wick/Hilbert, RG scale, or electroweak physical dynamics",
	)
}

func join(parts ...string) string { return strings.Join(parts, "; ") }
