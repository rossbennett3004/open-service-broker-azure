package mysql

import (
	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (d *databaseManager) Bind(
	instance service.Instance,
	bp service.BindingParameters,
) (service.BindingDetails, error) {
	pdt := instance.Parent.Details.(*dbmsInstanceDetails)
	dt := instance.Details.(*databaseInstanceDetails)
	return createBinding(
		bp,
		isSSLRequired(*instance.Parent.ProvisioningParameters),
		d.sqlDatabaseDNSSuffix,
		pdt.ServerName,
		pdt.AdministratorLogin,
		string(pdt.AdministratorLoginPassword),
		pdt.FullyQualifiedDomainName,
		dt.DatabaseName,
	)
}

func (d *databaseManager) GetCredentials(
	instance service.Instance,
	binding service.Binding,
) (service.Credentials, error) {
	pdt := instance.Parent.Details.(*dbmsInstanceDetails)
	dt := instance.Details.(*databaseInstanceDetails)
	bd := binding.Details.(*bindingDetails)
	creds := createCredential(
		pdt.FullyQualifiedDomainName,
		isSSLRequired(*instance.Parent.ProvisioningParameters),
		pdt.ServerName,
		dt.DatabaseName,
		bd,
	)
	return creds, nil
}
