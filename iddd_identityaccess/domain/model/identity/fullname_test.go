package identity

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewFullName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got, err := NewFullName("FirstName", "lastName")
		if err != nil {
			log.Fatal(err)
		}

		want := &FullName{firstName: "FirstName", lastName: "lastName"}

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(FullName{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail firstName empty", func(t *testing.T) {
		firstName, lastName := "", "lastName"
		_, err := NewFullName(firstName, lastName)
		if reflect.TypeOf(err) != reflect.TypeOf(&NameIsRequiredError{}) {
			t.Errorf("err type:%v, expect type:%v", reflect.TypeOf(err), reflect.TypeOf(&TenantIdParseError{}))
		}
	})
	t.Run("fail fistName is over 50 characters", func(t *testing.T) {
		firstName, lastName := RandString(51), "lastName"
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): First name must be 50 characters or less.", firstName, lastName)
		if got := err.Error(); got != want {
			log.Fatal(err)
		}
	})
	t.Run("fail first Name not start with a capital letter", func(t *testing.T) {
		firstName, lastName := "firstName", "lastName"
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): First name must be at least one character in length, starting with a capital letter.", firstName, lastName)
		if got := err.Error(); got != want {
			log.Fatal(err)
		}
	})
	t.Run("fail Last name must be 50 characters or less.", func(t *testing.T) {
		firstName, lastName := "FirstName", ""
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): Last name is required.", firstName, lastName)
		if got := err.Error(); got != want {
			log.Fatal(err)
		}
	})
	t.Run("fail Last name must be 50 characters or less.", func(t *testing.T) {
		firstName, lastName := "FirstName", RandString(51)
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): Last name must be 50 characters or less.", firstName, lastName)
		if got := err.Error(); got != want {
			log.Fatal(err)
		}
	})
	t.Run("fail Last name must be at least one character in length.", func(t *testing.T) {
		firstName, lastName := "FirstName", "#"
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): Last name must be at least one character in length.", firstName, lastName)
		if got := err.Error(); got != want {
			log.Fatal(err)
		}
	})
}

func TestAsFormattedName(t *testing.T) {
	firstName, lastName := "FirstName", "lastName"
	fullName, err := NewFullName(firstName, lastName)
	if err != nil {
		log.Fatal(err)
	}
	got := fullName.AsFormattedName()
	want := "FirstName lastName"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestFullNameEquals(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fullName, err := NewFullName("FirstName", "lastName")
		if err != nil {
			log.Fatal(err)
		}

		otherFullName := &FullName{firstName: "FirstName", lastName: "lastName"}

		if !fullName.Equal(otherFullName) {
			t.Errorf("fullName: %v must be equal to otherFullName %v", fullName, otherFullName)
		}
	})
	t.Run("not equal", func(t *testing.T) {
		fullName, err := NewFullName("FirstName", "lastName")
		if err != nil {
			log.Fatal(err)
		}

		otherFullName := &FullName{firstName: "FirstName", lastName: "LastName"}

		if fullName.Equal(otherFullName) {
			t.Errorf("fullName: %v must not be equal to otherFullName %v", fullName, otherFullName)
		}
	})
}

func TestFistName(t *testing.T) {
	fullName, err := NewFullName("FirstName", "lastName")
	if err != nil {
		log.Fatal(err)
	}
	got := fullName.FirstName()
	want := "FirstName"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestLastName(t *testing.T) {
	fullName, err := NewFullName("FirstName", "lastName")
	if err != nil {
		log.Fatal(err)
	}
	got := fullName.LastName()
	want := "lastName"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestWithChangedFirstName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fullName, err := NewFullName("FirstName", "lastName")
		if err != nil {
			log.Fatal(err)
		}
		changedFirstName := "ChangedFirstName"
		got, err := fullName.WithChangedFirstName(changedFirstName)
		if err != nil {
			log.Fatal(err)
		}
		want := &FullName{firstName: changedFirstName, lastName: "lastName"}

		if diff := cmp.Diff(want, got, cmp.AllowUnexported(FullName{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail", func(t *testing.T) {
		firstName, lastName := "FirstName", "lastName"
		fullName, err := NewFullName(firstName, lastName)
		if err != nil {
			log.Fatal(err)
		}
		changedFirstName := "changedFirstName"
		_, err = fullName.WithChangedFirstName(changedFirstName)
		want := fmt.Sprintf("fullname.WithChangedFirstName(%s): fullname.NewFullName(%v, %v): First name must be at least one character in length, starting with a capital letter.", changedFirstName, changedFirstName, lastName)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestWithChangedLastName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fullName, err := NewFullName("FirstName", "lastName")
		if err != nil {
			log.Fatal(err)
		}
		changedLastName := "ChangedLastName"
		got, err := fullName.WithChangedLastName(changedLastName)
		if err != nil {
			log.Fatal(err)
		}
		want := &FullName{firstName: "FirstName", lastName: changedLastName}

		if diff := cmp.Diff(want, got, cmp.AllowUnexported(FullName{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail", func(t *testing.T) {
		firstName, lastName := "FirstName", "lastName"
		fullName, err := NewFullName(firstName, lastName)
		if err != nil {
			log.Fatal(err)
		}
		changedLastName := ""
		_, err = fullName.WithChangedLastName(changedLastName)
		want := fmt.Sprintf("fullname.WithChangedLastName(%s): fullname.NewFullName(%v, %v): Last name is required.", changedLastName, firstName, changedLastName)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestFullNameString(t *testing.T) {
	firstName, lastName := "FirstName", "lastName"
	fullName, err := NewFullName(firstName, lastName)
	if err != nil {
		log.Fatal(err)
	}
	got := fmt.Sprint(fullName)
	want := fmt.Sprintf("FullName [firstName=" + firstName + ", lastName=" + lastName + "]")
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
