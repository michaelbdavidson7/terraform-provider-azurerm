package client

import (
	"github.com/Azure/azure-sdk-for-go/services/datashare/mgmt/2019-11-01/datashare"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/common"
)

type Client struct {
	AccountClient         *datashare.AccountsClient
	SharesClient          *datashare.SharesClient
	SynchronizationClient *datashare.SynchronizationSettingsClient
}

func NewClient(o *common.ClientOptions) *Client {
	accountClient := datashare.NewAccountsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&accountClient.Client, o.ResourceManagerAuthorizer)

	sharesClient := datashare.NewSharesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&sharesClient.Client, o.ResourceManagerAuthorizer)

	synchronizationSettingsClient := datashare.NewSynchronizationSettingsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&synchronizationSettingsClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		AccountClient:         &accountClient,
		SharesClient:          &sharesClient,
		SynchronizationClient: &synchronizationSettingsClient,
	}
}
