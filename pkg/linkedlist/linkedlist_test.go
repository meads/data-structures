package linkedlist

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_New_Returns_Default_Instance(t *testing.T) {
	sut := New()

	if sut.(*LinkedList).Head != nil {
		t.Errorf("expected '<nil>' got '%+v", sut.(*LinkedList).Head)
	}
}

func Test_InsertFront_Makes_New_Head_Node(t *testing.T) {
	sut := LinkedList{}
	sut.InsertFront("Testing")
	sut.InsertFront("One")
	sut.InsertFront("Two")

	if (*sut.Head).Data != "Two" {
		t.Errorf("expected 'Two' got '%s'", (*sut.Head).Data)
		MarshalAndPrint(sut)
	}
	if (*sut.Head).Next.Data != "One" {
		t.Errorf("expected 'One' got '%s'", (*sut.Head).Next.Data)
		MarshalAndPrint(sut)
	}
	if (*sut.Head).Next.Next.Data != "Testing" {
		t.Errorf("expected 'Testing' got '%s'", (*sut.Head).Next.Next.Data)
		MarshalAndPrint(sut)
	}
	if (*sut.Head).Next.Next.Next != nil {
		t.Errorf("expected '<nil>' got '%+v'", (*sut.Head).Next.Next.Next)
		MarshalAndPrint(sut)
	}
}

func Test_InsertLast(t *testing.T) {
	sut := LinkedList{}
	sut.InsertLast("Testing")
	sut.InsertLast("One")
	sut.InsertLast("Two")

	if (*sut.Head).Data != "Testing" {
		t.Errorf("expected 'Testing' got '%s'", (*sut.Head).Data)
		MarshalAndPrint(sut)
	}
	if (*sut.Head).Next.Data != "One" {
		t.Errorf("expected 'One' got '%s'", (*sut.Head).Next.Data)
		MarshalAndPrint(sut)
	}
	if (*sut.Head).Next.Next.Data != "Two" {
		t.Errorf("expected 'Two' got '%s'", (*sut.Head).Next.Next.Data)
		MarshalAndPrint(sut)
	}
	if (*sut.Head).Next.Next.Next != nil {
		t.Errorf("expected '<nil>' got '%+v'", (*sut.Head).Next.Next.Next)
		MarshalAndPrint(sut)
	}
}

func Test_InsertAfter(t *testing.T) {
	sut := LinkedList{}
	sut.InsertLast("Testing")
	sut.InsertLast("One")
	sut.InsertLast("Two")

	if (*sut.Head).Next.Data != "One" {
		t.Fatalf("expected Data for element to equal 'One' got '%s'", (*sut.Head).Next.Data)
		MarshalAndPrint(sut)
	}

	sut.InsertAfter((*sut.Head).Next, "One and a half...")
	if (*sut.Head).Next.Next.Data != "One and a half..." {
		t.Errorf("expected 'One and a half...' got '%s'", (*sut.Head).Next.Next.Data)
		MarshalAndPrint(sut)
	}
}

func Test_InsertAfter_Given_Nil_Value_Input_Is_Ignored(t *testing.T) {
	sut := LinkedList{}
	sut.InsertLast("Testing")
	sut.InsertLast("One")
	sut.InsertLast("Two")

	sut.InsertAfter(nil, "test")

	if (*sut.Head).Data != "Testing" {
		t.Errorf("expected 'Testing' got '%s'", (*sut.Head).Data)
		MarshalAndPrint(sut)
	}
	if (*sut.Head).Next.Data != "One" {
		t.Errorf("expected 'One' got '%s'", (*sut.Head).Next.Data)
		MarshalAndPrint(sut)
	}
	if (*sut.Head).Next.Next.Data != "Two" {
		t.Errorf("expected 'Two' got '%s'", (*sut.Head).Next.Next.Data)
		MarshalAndPrint(sut)
	}
	if (*sut.Head).Next.Next.Next != nil {
		t.Errorf("expected '<nil>' got '%+v'", (*sut.Head).Next.Next.Next)
		MarshalAndPrint(sut)
	}
}

func Test_GetLastNode(t *testing.T) {
	sut := LinkedList{}
	sut.InsertLast("Testing")
	sut.InsertLast("One")

	actual := sut.GetLastNode()
	if actual.Data != "One" {
		t.Errorf("expected 'One' got '%s'", actual.Data)
		MarshalAndPrint(sut)
	}
}

func Test_DeleteNodeByKey_Head_Gets_Deleted(t *testing.T) {
	sut := LinkedList{}
	sut.InsertLast("Testing")
	sut.InsertLast("One")
	sut.InsertLast("Two")
	sut.DeleteNodeByKey("Testing")
	if (*sut.Head).Data != "One" {
		t.Errorf("expected 'One' got '%s'", (*sut.Head).Data)
		MarshalAndPrint(sut)
	}
	if (*sut.Head).Next.Data != "Two" {
		t.Errorf("expected 'Two' got '%s'", (*sut.Head).Next.Data)
		MarshalAndPrint(sut)
	}
}

func Test_DeleteNodeByKey_Nth_Node_Gets_Deleted(t *testing.T) {
	sut := LinkedList{}
	sut.InsertLast("Testing")
	sut.InsertLast("One")
	sut.InsertLast("Two")
	sut.DeleteNodeByKey("One")
	if (*sut.Head).Data != "Testing" {
		t.Errorf("expected 'Testing' got '%s'", (*sut.Head).Data)
		MarshalAndPrint(sut)
	}
	if (*sut.Head).Next.Data != "Two" {
		t.Errorf("expected 'Two' got '%s'", (*sut.Head).Next.Data)
		MarshalAndPrint(sut)
	}
}

func Test_DeleteNodeByKey_Last_Node_Gets_Deleted(t *testing.T) {
	sut := LinkedList{}
	sut.InsertLast("Testing")
	sut.InsertLast("One")
	sut.InsertLast("Two")
	sut.DeleteNodeByKey("Two")
	if (*sut.Head).Data != "Testing" {
		t.Errorf("expected 'Testing' got '%s'", (*sut.Head).Data)
		MarshalAndPrint(sut)
	}
	if (*sut.Head).Next.Data != "One" {
		t.Errorf("expected 'One' got '%s'", (*sut.Head).Next.Data)
		MarshalAndPrint(sut)
	}
}

func Test_DeleteNodeByKey_Invalid_Key_Gets_Ignored(t *testing.T) {
	sut := LinkedList{}
	sut.InsertLast("Testing")
	sut.InsertLast("One")
	sut.InsertLast("Two")
	sut.DeleteNodeByKey("Invalid")
	if (*sut.Head).Data != "Testing" {
		t.Errorf("expected 'Testing' got '%s'", (*sut.Head).Data)
		MarshalAndPrint(sut)
	}
	if (*sut.Head).Next.Data != "One" {
		t.Errorf("expected 'One' got '%s'", (*sut.Head).Next.Data)
		MarshalAndPrint(sut)
	}
	if (*sut.Head).Next.Next.Data != "Two" {
		t.Errorf("expected 'Two' got '%s'", (*sut.Head).Next.Next.Data)
		MarshalAndPrint(sut)
	}
}

func Test_Reverse(t *testing.T) {
	sut := LinkedList{}
	sut.InsertLast("Testing")
	sut.InsertLast("One")
	sut.InsertLast("Two")

	sut.Reverse()

	if (*sut.Head).Data != "Two" {
		t.Errorf("expected 'Two' got '%s'", (*sut.Head).Data)
		MarshalAndPrint(sut)
	}
	if (*sut.Head).Next.Data != "One" {
		t.Errorf("expected 'One' got '%s'", (*sut.Head).Next.Data)
		MarshalAndPrint(sut)
	}
	if (*sut.Head).Next.Next.Data != "Testing" {
		t.Errorf("expected 'Testing' got '%s'", (*sut.Head).Next.Next.Data)
		MarshalAndPrint(sut)
	}
	if (*sut.Head).Next.Next.Next != nil {
		t.Errorf("expected '<nil>' got '%+v'", (*sut.Head).Next.Next.Next)
		MarshalAndPrint(sut)
	}
}

func MarshalAndPrint(l LinkedList) {
	b, err := json.MarshalIndent(l, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
