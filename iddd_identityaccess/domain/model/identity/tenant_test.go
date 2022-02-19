package identity

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestNewTenant(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId, err := NewTenantId(uu)
		if err != nil {
			t.Fatal(err)
		}

		name := "TenantName"
		tenant, err := NewTenant(*tenantId, name)
		if err != nil {
			t.Fatal(tenant)
		}

		if !reflect.DeepEqual(tenant.tenantId.tenantId, uu) {
			t.Errorf("tenant.tenantId %v should be equal to uu %v", tenant.tenantId, uu)
		}

		if !reflect.DeepEqual(tenant.name, name) {
			t.Errorf("tenant.name %v should be equal to name %v", tenant.name, name)
		}
	})
	t.Run("fail empty name", func(t *testing.T) {
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId, err := NewTenantId(uu)
		if err != nil {
			t.Fatal(err)
		}

		name := ""
		tenant, err := NewTenant(*tenantId, name)
		want := fmt.Sprintf("tenant.NewTenant(%v, %s): The tenant name is required.", *tenantId, name)
		if got := err.Error(); want != got {
			t.Errorf("got %s, want %s", got, want)
		}
		if tenant != nil {
			t.Errorf("tenant should be nil, but %v", tenant)
		}
	})
}
