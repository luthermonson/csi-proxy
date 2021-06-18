// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package iscsi

import (
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"github.com/kubernetes-csi/csi-proxy/pkg/server/iscsi/impl"
	"github.com/kubernetes-csi/csi-proxy/pkg/server/iscsi/impl/v1alpha1"
	"github.com/kubernetes-csi/csi-proxy/pkg/server/iscsi/impl/v1alpha2"
	srvtypes "github.com/kubernetes-csi/csi-proxy/pkg/server/types"
)

const name = "iscsi"

// ensure the server defines all the required methods
var _ impl.ServerInterface = &Server{}

func (s *Server) VersionedAPIs() []*srvtypes.VersionedAPI {
	v1alpha1Server := v1alpha1.NewVersionedServer(s)
	v1alpha2Server := v1alpha2.NewVersionedServer(s)

	return []*srvtypes.VersionedAPI{
		{
			Group:      name,
			Version:    apiversion.NewVersionOrPanic("v1alpha1"),
			Registrant: v1alpha1Server.Register,
		},
		{
			Group:      name,
			Version:    apiversion.NewVersionOrPanic("v1alpha2"),
			Registrant: v1alpha2Server.Register,
		},
	}
}