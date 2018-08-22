package blobclient

import "code.uber.internal/infra/kraken/lib/hostlist"

// Provider defines an interface for creating Client scoped to an origin addr.
type Provider interface {
	Provide(addr string) Client
}

// HTTPProvider provides HTTPClients.
type HTTPProvider struct{}

// NewProvider returns a new HTTPProvider.
func NewProvider() HTTPProvider {
	return HTTPProvider{}
}

// Provide implements ClientProvider's Provide.
// TODO(codyg): Make this return error.
func (p HTTPProvider) Provide(addr string) Client {
	return New(addr)
}

// ClusterProvider creates ClusterClients from dns records.
type ClusterProvider interface {
	Provide(dns string) (ClusterClient, error)
}

// HTTPClusterProvider provides ClusterClients backed by HTTP. Does not include
// health checks.
type HTTPClusterProvider struct{}

// NewClusterProvider returns a new HTTPClusterProvider.
func NewClusterProvider() HTTPClusterProvider {
	return HTTPClusterProvider{}
}

// Provide creates a new ClusterClient.
func (p HTTPClusterProvider) Provide(dns string) (ClusterClient, error) {
	hosts, err := hostlist.New(hostlist.Config{DNS: dns})
	if err != nil {
		return nil, err
	}
	return NewClusterClient(NewClientResolver(NewProvider(), hosts)), nil
}
