package generation2cl17boardscalarhiggstypeledgerandoperatorairlockaudit

import (
	"strings"
	"testing"
)

func TestGate750NativeHiggsBoundaryAndHistoryTyping(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate749.Inherited || !a.Gate749.WallHierarchyOrdered || !a.Gate749.K7RoleSeparated || !a.Gate749.FirewallsPreserved {
		t.Fatalf("bad Gate749 inheritance: %+v", a.Gate749)
	}
	if a.Finite.MeasurementDim != 8 || a.Finite.FourthGradeDim != 70 || a.Finite.BooleanRank != 56 || a.Finite.OctonionicRank != 14 || a.Finite.K7Dim != 7 || !strings.Contains(a.Finite.PK7LivesIn, "End") {
		t.Fatalf("bad finite board: %+v", a.Finite)
	}
	if a.Hodge.PlusDim != 4 || a.Hodge.MinusDim != 3 || !a.Hodge.NativePolarity || !a.Hodge.HiggsShadowOnly || !a.Hodge.FlavorShadowOnly {
		t.Fatalf("bad hodge board: %+v", a.Hodge)
	}
	if a.Socket.SealPackage != "(n,q,P_rad)" || !strings.Contains(a.Socket.HopfPayoffOperator, "P_rad") || !strings.Contains(a.Socket.HistoryLoopScalar, "Tr") || len(a.Socket.TypingRules) != 4 {
		t.Fatalf("bad sealed socket board: %+v", a.Socket)
	}
	if !strings.Contains(a.Boundary.Plane, "R^2") || !strings.Contains(a.Boundary.QuotientCoordinate, "S_split") || !strings.Contains(a.Boundary.LawfulAdditionReason, "same boundary quotient") {
		t.Fatalf("bad boundary board: %+v", a.Boundary)
	}
	if a.H72.Dim != 72 || !strings.Contains(a.H72.EventWeight, "7/72") || !strings.Contains(a.H72.MultiplicationMeaning, "not tensor product") || !strings.Contains(a.H72.RawMomentFormula, "p_K7") {
		t.Fatalf("bad H72 board: %+v", a.H72)
	}
	if !strings.Contains(a.History.PolynomialType, "not an operator on K7") || !strings.Contains(a.History.Verdict, StatusFWall3ScalarResponseNotOperatorGeo) {
		t.Fatalf("bad history board: %+v", a.History)
	}
}

func TestGate750RuntimeOperationsAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Runtime.OperatorMultiplicationRemains || !strings.Contains(a.Runtime.RuntimeFormula, "lambda_runtime") || !strings.Contains(a.Runtime.MultiplicationMeaning, "scalar") {
		t.Fatalf("bad runtime board: %+v", a.Runtime)
	}
	if !a.TreeProxy.PoleMassBlocked || !strings.Contains(a.TreeProxy.ProxyType, "not pole mass") {
		t.Fatalf("bad tree proxy board: %+v", a.TreeProxy)
	}
	if len(a.Operations.LawfulAdditions) != 3 || len(a.Operations.ScalarMultiplications) != 4 || len(a.Operations.OperatorCompositions) != 4 || len(a.Operations.TraceExpectations) != 3 || len(a.Operations.ForbiddenOperations) != 5 {
		t.Fatalf("bad operations audit: %+v", a.Operations)
	}
	if !strings.Contains(strings.Join(a.Operations.ForbiddenOperations, "\n"), "K7 + boundary vector") || !strings.Contains(a.Operations.Verdict, StatusHomTensorResponseNotNativeSubspace) {
		t.Fatalf("missing forbidden operation firewall: %+v", a.Operations)
	}
	if !a.Firewalls.K7BoundaryVectorBlocked || !a.Firewalls.HomTensorSubspaceBlocked || !a.Firewalls.ScalarRuntimeOperatorBlocked || !a.Firewalls.TreePoleBlocked || !a.Firewalls.NativeSealsMissing || !a.Firewalls.BoundaryGeneratingFunctionOpen || !a.Firewalls.HistoryLoopNativeOpen || !a.Firewalls.HiggsMassBlocked || !a.Firewalls.YukawaBlocked {
		t.Fatalf("bad type firewalls: %+v", a.Firewalls)
	}

	res := Generation2CL17BoardScalarHiggsTypeLedgerAndOperatorAirlockAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status note %s", want)
		}
	}
}
