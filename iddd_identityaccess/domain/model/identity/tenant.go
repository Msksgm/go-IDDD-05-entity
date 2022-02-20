package identity

import (
	"fmt"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
)

type Tenant struct {
	tenantId TenantId
	name     string
	active   bool
}

func NewTenant(tenantId TenantId, name string, active bool) (_ *Tenant, err error) {
	tenant := new(Tenant)
	if err := tenant.setName(name); err != nil {
		return nil, err
	}
	if err := tenant.setTenantId(tenantId); err != nil {
		return nil, err
	}
	tenant.setActive(active)
	return tenant, nil
}

func (tenant *Tenant) setName(name string) (err error) {
	defer ierrors.Wrap(&err, "tenant.setName(%s)", name)
	if name == "" {
		return fmt.Errorf("The tenant name is required.")
	}
	if len(name) < 1 || len(name) > 100 {
		return fmt.Errorf("The tenant description must be 100 characters or less.")
	}
	tenant.name = name
	return nil
}

func (tenant *Tenant) setTenantId(tenantId TenantId) (err error) {
	defer ierrors.Wrap(&err, "tenant.setTenantId(%s)", tenantId)
	if tenantId.id == "" {
		return fmt.Errorf("TenentId is required.")
	}
	tenant.tenantId = tenantId
	return nil
}

func (tenant *Tenant) setActive(active bool) {
	tenant.active = active
}

func (tenant *Tenant) Equals(otherTenant Tenant) bool {
	return tenant.tenantId == otherTenant.tenantId
}
