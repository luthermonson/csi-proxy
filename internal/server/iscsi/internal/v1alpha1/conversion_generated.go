// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	v1alpha1 "github.com/kubernetes-csi/csi-proxy/client/api/iscsi/v1alpha1"
	internal "github.com/kubernetes-csi/csi-proxy/internal/server/iscsi/internal"
)

func autoConvert_v1alpha1_AddTargetPortalRequest_To_internal_AddTargetPortalRequest(in *v1alpha1.AddTargetPortalRequest, out *internal.AddTargetPortalRequest) error {
	if in.TargetPortal != nil {
		in, out := &in.TargetPortal, &out.TargetPortal
		*out = new(internal.TargetPortal)
		if err := Convert_v1alpha1_TargetPortal_To_internal_TargetPortal(*in, *out); err != nil {
			return err
		}
	} else {
		out.TargetPortal = nil
	}
	return nil
}

// Convert_v1alpha1_AddTargetPortalRequest_To_internal_AddTargetPortalRequest is an autogenerated conversion function.
func Convert_v1alpha1_AddTargetPortalRequest_To_internal_AddTargetPortalRequest(in *v1alpha1.AddTargetPortalRequest, out *internal.AddTargetPortalRequest) error {
	return autoConvert_v1alpha1_AddTargetPortalRequest_To_internal_AddTargetPortalRequest(in, out)
}

func autoConvert_internal_AddTargetPortalRequest_To_v1alpha1_AddTargetPortalRequest(in *internal.AddTargetPortalRequest, out *v1alpha1.AddTargetPortalRequest) error {
	if in.TargetPortal != nil {
		in, out := &in.TargetPortal, &out.TargetPortal
		*out = new(v1alpha1.TargetPortal)
		if err := Convert_internal_TargetPortal_To_v1alpha1_TargetPortal(*in, *out); err != nil {
			return err
		}
	} else {
		out.TargetPortal = nil
	}
	return nil
}

// Convert_internal_AddTargetPortalRequest_To_v1alpha1_AddTargetPortalRequest is an autogenerated conversion function.
func Convert_internal_AddTargetPortalRequest_To_v1alpha1_AddTargetPortalRequest(in *internal.AddTargetPortalRequest, out *v1alpha1.AddTargetPortalRequest) error {
	return autoConvert_internal_AddTargetPortalRequest_To_v1alpha1_AddTargetPortalRequest(in, out)
}

func autoConvert_v1alpha1_AddTargetPortalResponse_To_internal_AddTargetPortalResponse(in *v1alpha1.AddTargetPortalResponse, out *internal.AddTargetPortalResponse) error {
	return nil
}

// Convert_v1alpha1_AddTargetPortalResponse_To_internal_AddTargetPortalResponse is an autogenerated conversion function.
func Convert_v1alpha1_AddTargetPortalResponse_To_internal_AddTargetPortalResponse(in *v1alpha1.AddTargetPortalResponse, out *internal.AddTargetPortalResponse) error {
	return autoConvert_v1alpha1_AddTargetPortalResponse_To_internal_AddTargetPortalResponse(in, out)
}

func autoConvert_internal_AddTargetPortalResponse_To_v1alpha1_AddTargetPortalResponse(in *internal.AddTargetPortalResponse, out *v1alpha1.AddTargetPortalResponse) error {
	return nil
}

// Convert_internal_AddTargetPortalResponse_To_v1alpha1_AddTargetPortalResponse is an autogenerated conversion function.
func Convert_internal_AddTargetPortalResponse_To_v1alpha1_AddTargetPortalResponse(in *internal.AddTargetPortalResponse, out *v1alpha1.AddTargetPortalResponse) error {
	return autoConvert_internal_AddTargetPortalResponse_To_v1alpha1_AddTargetPortalResponse(in, out)
}

func autoConvert_v1alpha1_ConnectTargetRequest_To_internal_ConnectTargetRequest(in *v1alpha1.ConnectTargetRequest, out *internal.ConnectTargetRequest) error {
	if in.TargetPortal != nil {
		in, out := &in.TargetPortal, &out.TargetPortal
		*out = new(internal.TargetPortal)
		if err := Convert_v1alpha1_TargetPortal_To_internal_TargetPortal(*in, *out); err != nil {
			return err
		}
	} else {
		out.TargetPortal = nil
	}
	out.Iqn = in.Iqn
	out.AuthType = internal.AuthenticationType(in.AuthType)
	out.ChapUsername = in.ChapUsername
	out.ChapSecret = in.ChapSecret
	return nil
}

// Convert_v1alpha1_ConnectTargetRequest_To_internal_ConnectTargetRequest is an autogenerated conversion function.
func Convert_v1alpha1_ConnectTargetRequest_To_internal_ConnectTargetRequest(in *v1alpha1.ConnectTargetRequest, out *internal.ConnectTargetRequest) error {
	return autoConvert_v1alpha1_ConnectTargetRequest_To_internal_ConnectTargetRequest(in, out)
}

func autoConvert_internal_ConnectTargetRequest_To_v1alpha1_ConnectTargetRequest(in *internal.ConnectTargetRequest, out *v1alpha1.ConnectTargetRequest) error {
	if in.TargetPortal != nil {
		in, out := &in.TargetPortal, &out.TargetPortal
		*out = new(v1alpha1.TargetPortal)
		if err := Convert_internal_TargetPortal_To_v1alpha1_TargetPortal(*in, *out); err != nil {
			return err
		}
	} else {
		out.TargetPortal = nil
	}
	out.Iqn = in.Iqn
	out.AuthType = v1alpha1.AuthenticationType(in.AuthType)
	out.ChapUsername = in.ChapUsername
	out.ChapSecret = in.ChapSecret
	return nil
}

// Convert_internal_ConnectTargetRequest_To_v1alpha1_ConnectTargetRequest is an autogenerated conversion function.
func Convert_internal_ConnectTargetRequest_To_v1alpha1_ConnectTargetRequest(in *internal.ConnectTargetRequest, out *v1alpha1.ConnectTargetRequest) error {
	return autoConvert_internal_ConnectTargetRequest_To_v1alpha1_ConnectTargetRequest(in, out)
}

func autoConvert_v1alpha1_ConnectTargetResponse_To_internal_ConnectTargetResponse(in *v1alpha1.ConnectTargetResponse, out *internal.ConnectTargetResponse) error {
	return nil
}

// Convert_v1alpha1_ConnectTargetResponse_To_internal_ConnectTargetResponse is an autogenerated conversion function.
func Convert_v1alpha1_ConnectTargetResponse_To_internal_ConnectTargetResponse(in *v1alpha1.ConnectTargetResponse, out *internal.ConnectTargetResponse) error {
	return autoConvert_v1alpha1_ConnectTargetResponse_To_internal_ConnectTargetResponse(in, out)
}

func autoConvert_internal_ConnectTargetResponse_To_v1alpha1_ConnectTargetResponse(in *internal.ConnectTargetResponse, out *v1alpha1.ConnectTargetResponse) error {
	return nil
}

// Convert_internal_ConnectTargetResponse_To_v1alpha1_ConnectTargetResponse is an autogenerated conversion function.
func Convert_internal_ConnectTargetResponse_To_v1alpha1_ConnectTargetResponse(in *internal.ConnectTargetResponse, out *v1alpha1.ConnectTargetResponse) error {
	return autoConvert_internal_ConnectTargetResponse_To_v1alpha1_ConnectTargetResponse(in, out)
}

func autoConvert_v1alpha1_DisconnectTargetRequest_To_internal_DisconnectTargetRequest(in *v1alpha1.DisconnectTargetRequest, out *internal.DisconnectTargetRequest) error {
	if in.TargetPortal != nil {
		in, out := &in.TargetPortal, &out.TargetPortal
		*out = new(internal.TargetPortal)
		if err := Convert_v1alpha1_TargetPortal_To_internal_TargetPortal(*in, *out); err != nil {
			return err
		}
	} else {
		out.TargetPortal = nil
	}
	out.Iqn = in.Iqn
	return nil
}

// Convert_v1alpha1_DisconnectTargetRequest_To_internal_DisconnectTargetRequest is an autogenerated conversion function.
func Convert_v1alpha1_DisconnectTargetRequest_To_internal_DisconnectTargetRequest(in *v1alpha1.DisconnectTargetRequest, out *internal.DisconnectTargetRequest) error {
	return autoConvert_v1alpha1_DisconnectTargetRequest_To_internal_DisconnectTargetRequest(in, out)
}

func autoConvert_internal_DisconnectTargetRequest_To_v1alpha1_DisconnectTargetRequest(in *internal.DisconnectTargetRequest, out *v1alpha1.DisconnectTargetRequest) error {
	if in.TargetPortal != nil {
		in, out := &in.TargetPortal, &out.TargetPortal
		*out = new(v1alpha1.TargetPortal)
		if err := Convert_internal_TargetPortal_To_v1alpha1_TargetPortal(*in, *out); err != nil {
			return err
		}
	} else {
		out.TargetPortal = nil
	}
	out.Iqn = in.Iqn
	return nil
}

// Convert_internal_DisconnectTargetRequest_To_v1alpha1_DisconnectTargetRequest is an autogenerated conversion function.
func Convert_internal_DisconnectTargetRequest_To_v1alpha1_DisconnectTargetRequest(in *internal.DisconnectTargetRequest, out *v1alpha1.DisconnectTargetRequest) error {
	return autoConvert_internal_DisconnectTargetRequest_To_v1alpha1_DisconnectTargetRequest(in, out)
}

func autoConvert_v1alpha1_DisconnectTargetResponse_To_internal_DisconnectTargetResponse(in *v1alpha1.DisconnectTargetResponse, out *internal.DisconnectTargetResponse) error {
	return nil
}

// Convert_v1alpha1_DisconnectTargetResponse_To_internal_DisconnectTargetResponse is an autogenerated conversion function.
func Convert_v1alpha1_DisconnectTargetResponse_To_internal_DisconnectTargetResponse(in *v1alpha1.DisconnectTargetResponse, out *internal.DisconnectTargetResponse) error {
	return autoConvert_v1alpha1_DisconnectTargetResponse_To_internal_DisconnectTargetResponse(in, out)
}

func autoConvert_internal_DisconnectTargetResponse_To_v1alpha1_DisconnectTargetResponse(in *internal.DisconnectTargetResponse, out *v1alpha1.DisconnectTargetResponse) error {
	return nil
}

// Convert_internal_DisconnectTargetResponse_To_v1alpha1_DisconnectTargetResponse is an autogenerated conversion function.
func Convert_internal_DisconnectTargetResponse_To_v1alpha1_DisconnectTargetResponse(in *internal.DisconnectTargetResponse, out *v1alpha1.DisconnectTargetResponse) error {
	return autoConvert_internal_DisconnectTargetResponse_To_v1alpha1_DisconnectTargetResponse(in, out)
}

func autoConvert_v1alpha1_DiscoverTargetPortalRequest_To_internal_DiscoverTargetPortalRequest(in *v1alpha1.DiscoverTargetPortalRequest, out *internal.DiscoverTargetPortalRequest) error {
	if in.TargetPortal != nil {
		in, out := &in.TargetPortal, &out.TargetPortal
		*out = new(internal.TargetPortal)
		if err := Convert_v1alpha1_TargetPortal_To_internal_TargetPortal(*in, *out); err != nil {
			return err
		}
	} else {
		out.TargetPortal = nil
	}
	return nil
}

// Convert_v1alpha1_DiscoverTargetPortalRequest_To_internal_DiscoverTargetPortalRequest is an autogenerated conversion function.
func Convert_v1alpha1_DiscoverTargetPortalRequest_To_internal_DiscoverTargetPortalRequest(in *v1alpha1.DiscoverTargetPortalRequest, out *internal.DiscoverTargetPortalRequest) error {
	return autoConvert_v1alpha1_DiscoverTargetPortalRequest_To_internal_DiscoverTargetPortalRequest(in, out)
}

func autoConvert_internal_DiscoverTargetPortalRequest_To_v1alpha1_DiscoverTargetPortalRequest(in *internal.DiscoverTargetPortalRequest, out *v1alpha1.DiscoverTargetPortalRequest) error {
	if in.TargetPortal != nil {
		in, out := &in.TargetPortal, &out.TargetPortal
		*out = new(v1alpha1.TargetPortal)
		if err := Convert_internal_TargetPortal_To_v1alpha1_TargetPortal(*in, *out); err != nil {
			return err
		}
	} else {
		out.TargetPortal = nil
	}
	return nil
}

// Convert_internal_DiscoverTargetPortalRequest_To_v1alpha1_DiscoverTargetPortalRequest is an autogenerated conversion function.
func Convert_internal_DiscoverTargetPortalRequest_To_v1alpha1_DiscoverTargetPortalRequest(in *internal.DiscoverTargetPortalRequest, out *v1alpha1.DiscoverTargetPortalRequest) error {
	return autoConvert_internal_DiscoverTargetPortalRequest_To_v1alpha1_DiscoverTargetPortalRequest(in, out)
}

func autoConvert_v1alpha1_DiscoverTargetPortalResponse_To_internal_DiscoverTargetPortalResponse(in *v1alpha1.DiscoverTargetPortalResponse, out *internal.DiscoverTargetPortalResponse) error {
	out.Iqns = *(*[]string)(unsafe.Pointer(&in.Iqns))
	return nil
}

// Convert_v1alpha1_DiscoverTargetPortalResponse_To_internal_DiscoverTargetPortalResponse is an autogenerated conversion function.
func Convert_v1alpha1_DiscoverTargetPortalResponse_To_internal_DiscoverTargetPortalResponse(in *v1alpha1.DiscoverTargetPortalResponse, out *internal.DiscoverTargetPortalResponse) error {
	return autoConvert_v1alpha1_DiscoverTargetPortalResponse_To_internal_DiscoverTargetPortalResponse(in, out)
}

func autoConvert_internal_DiscoverTargetPortalResponse_To_v1alpha1_DiscoverTargetPortalResponse(in *internal.DiscoverTargetPortalResponse, out *v1alpha1.DiscoverTargetPortalResponse) error {
	out.Iqns = *(*[]string)(unsafe.Pointer(&in.Iqns))
	return nil
}

// Convert_internal_DiscoverTargetPortalResponse_To_v1alpha1_DiscoverTargetPortalResponse is an autogenerated conversion function.
func Convert_internal_DiscoverTargetPortalResponse_To_v1alpha1_DiscoverTargetPortalResponse(in *internal.DiscoverTargetPortalResponse, out *v1alpha1.DiscoverTargetPortalResponse) error {
	return autoConvert_internal_DiscoverTargetPortalResponse_To_v1alpha1_DiscoverTargetPortalResponse(in, out)
}

func autoConvert_v1alpha1_GetTargetDisksRequest_To_internal_GetTargetDisksRequest(in *v1alpha1.GetTargetDisksRequest, out *internal.GetTargetDisksRequest) error {
	if in.TargetPortal != nil {
		in, out := &in.TargetPortal, &out.TargetPortal
		*out = new(internal.TargetPortal)
		if err := Convert_v1alpha1_TargetPortal_To_internal_TargetPortal(*in, *out); err != nil {
			return err
		}
	} else {
		out.TargetPortal = nil
	}
	out.Iqn = in.Iqn
	return nil
}

// Convert_v1alpha1_GetTargetDisksRequest_To_internal_GetTargetDisksRequest is an autogenerated conversion function.
func Convert_v1alpha1_GetTargetDisksRequest_To_internal_GetTargetDisksRequest(in *v1alpha1.GetTargetDisksRequest, out *internal.GetTargetDisksRequest) error {
	return autoConvert_v1alpha1_GetTargetDisksRequest_To_internal_GetTargetDisksRequest(in, out)
}

func autoConvert_internal_GetTargetDisksRequest_To_v1alpha1_GetTargetDisksRequest(in *internal.GetTargetDisksRequest, out *v1alpha1.GetTargetDisksRequest) error {
	if in.TargetPortal != nil {
		in, out := &in.TargetPortal, &out.TargetPortal
		*out = new(v1alpha1.TargetPortal)
		if err := Convert_internal_TargetPortal_To_v1alpha1_TargetPortal(*in, *out); err != nil {
			return err
		}
	} else {
		out.TargetPortal = nil
	}
	out.Iqn = in.Iqn
	return nil
}

// Convert_internal_GetTargetDisksRequest_To_v1alpha1_GetTargetDisksRequest is an autogenerated conversion function.
func Convert_internal_GetTargetDisksRequest_To_v1alpha1_GetTargetDisksRequest(in *internal.GetTargetDisksRequest, out *v1alpha1.GetTargetDisksRequest) error {
	return autoConvert_internal_GetTargetDisksRequest_To_v1alpha1_GetTargetDisksRequest(in, out)
}

func autoConvert_v1alpha1_GetTargetDisksResponse_To_internal_GetTargetDisksResponse(in *v1alpha1.GetTargetDisksResponse, out *internal.GetTargetDisksResponse) error {
	out.DiskIDs = *(*[]string)(unsafe.Pointer(&in.DiskIDs))
	return nil
}

// Convert_v1alpha1_GetTargetDisksResponse_To_internal_GetTargetDisksResponse is an autogenerated conversion function.
func Convert_v1alpha1_GetTargetDisksResponse_To_internal_GetTargetDisksResponse(in *v1alpha1.GetTargetDisksResponse, out *internal.GetTargetDisksResponse) error {
	return autoConvert_v1alpha1_GetTargetDisksResponse_To_internal_GetTargetDisksResponse(in, out)
}

func autoConvert_internal_GetTargetDisksResponse_To_v1alpha1_GetTargetDisksResponse(in *internal.GetTargetDisksResponse, out *v1alpha1.GetTargetDisksResponse) error {
	out.DiskIDs = *(*[]string)(unsafe.Pointer(&in.DiskIDs))
	return nil
}

// Convert_internal_GetTargetDisksResponse_To_v1alpha1_GetTargetDisksResponse is an autogenerated conversion function.
func Convert_internal_GetTargetDisksResponse_To_v1alpha1_GetTargetDisksResponse(in *internal.GetTargetDisksResponse, out *v1alpha1.GetTargetDisksResponse) error {
	return autoConvert_internal_GetTargetDisksResponse_To_v1alpha1_GetTargetDisksResponse(in, out)
}

func autoConvert_v1alpha1_ListTargetPortalsRequest_To_internal_ListTargetPortalsRequest(in *v1alpha1.ListTargetPortalsRequest, out *internal.ListTargetPortalsRequest) error {
	return nil
}

// Convert_v1alpha1_ListTargetPortalsRequest_To_internal_ListTargetPortalsRequest is an autogenerated conversion function.
func Convert_v1alpha1_ListTargetPortalsRequest_To_internal_ListTargetPortalsRequest(in *v1alpha1.ListTargetPortalsRequest, out *internal.ListTargetPortalsRequest) error {
	return autoConvert_v1alpha1_ListTargetPortalsRequest_To_internal_ListTargetPortalsRequest(in, out)
}

func autoConvert_internal_ListTargetPortalsRequest_To_v1alpha1_ListTargetPortalsRequest(in *internal.ListTargetPortalsRequest, out *v1alpha1.ListTargetPortalsRequest) error {
	return nil
}

// Convert_internal_ListTargetPortalsRequest_To_v1alpha1_ListTargetPortalsRequest is an autogenerated conversion function.
func Convert_internal_ListTargetPortalsRequest_To_v1alpha1_ListTargetPortalsRequest(in *internal.ListTargetPortalsRequest, out *v1alpha1.ListTargetPortalsRequest) error {
	return autoConvert_internal_ListTargetPortalsRequest_To_v1alpha1_ListTargetPortalsRequest(in, out)
}

func autoConvert_v1alpha1_ListTargetPortalsResponse_To_internal_ListTargetPortalsResponse(in *v1alpha1.ListTargetPortalsResponse, out *internal.ListTargetPortalsResponse) error {
	if in.TargetPortals != nil {
		in, out := &in.TargetPortals, &out.TargetPortals
		*out = make([]*internal.TargetPortal, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_TargetPortal_To_internal_TargetPortal(*&(*in)[i], *&(*out)[i]); err != nil {
				return err
			}
		}
	} else {
		out.TargetPortals = nil
	}
	return nil
}

// Convert_v1alpha1_ListTargetPortalsResponse_To_internal_ListTargetPortalsResponse is an autogenerated conversion function.
func Convert_v1alpha1_ListTargetPortalsResponse_To_internal_ListTargetPortalsResponse(in *v1alpha1.ListTargetPortalsResponse, out *internal.ListTargetPortalsResponse) error {
	return autoConvert_v1alpha1_ListTargetPortalsResponse_To_internal_ListTargetPortalsResponse(in, out)
}

// detected external conversion function
// Convert_internal_ListTargetPortalsResponse_To_v1alpha1_ListTargetPortalsResponse(in *internal.ListTargetPortalsResponse, out *v1alpha1.ListTargetPortalsResponse) error
// skipping generation of the auto function

func autoConvert_v1alpha1_RemoveTargetPortalRequest_To_internal_RemoveTargetPortalRequest(in *v1alpha1.RemoveTargetPortalRequest, out *internal.RemoveTargetPortalRequest) error {
	if in.TargetPortal != nil {
		in, out := &in.TargetPortal, &out.TargetPortal
		*out = new(internal.TargetPortal)
		if err := Convert_v1alpha1_TargetPortal_To_internal_TargetPortal(*in, *out); err != nil {
			return err
		}
	} else {
		out.TargetPortal = nil
	}
	return nil
}

// Convert_v1alpha1_RemoveTargetPortalRequest_To_internal_RemoveTargetPortalRequest is an autogenerated conversion function.
func Convert_v1alpha1_RemoveTargetPortalRequest_To_internal_RemoveTargetPortalRequest(in *v1alpha1.RemoveTargetPortalRequest, out *internal.RemoveTargetPortalRequest) error {
	return autoConvert_v1alpha1_RemoveTargetPortalRequest_To_internal_RemoveTargetPortalRequest(in, out)
}

func autoConvert_internal_RemoveTargetPortalRequest_To_v1alpha1_RemoveTargetPortalRequest(in *internal.RemoveTargetPortalRequest, out *v1alpha1.RemoveTargetPortalRequest) error {
	if in.TargetPortal != nil {
		in, out := &in.TargetPortal, &out.TargetPortal
		*out = new(v1alpha1.TargetPortal)
		if err := Convert_internal_TargetPortal_To_v1alpha1_TargetPortal(*in, *out); err != nil {
			return err
		}
	} else {
		out.TargetPortal = nil
	}
	return nil
}

// Convert_internal_RemoveTargetPortalRequest_To_v1alpha1_RemoveTargetPortalRequest is an autogenerated conversion function.
func Convert_internal_RemoveTargetPortalRequest_To_v1alpha1_RemoveTargetPortalRequest(in *internal.RemoveTargetPortalRequest, out *v1alpha1.RemoveTargetPortalRequest) error {
	return autoConvert_internal_RemoveTargetPortalRequest_To_v1alpha1_RemoveTargetPortalRequest(in, out)
}

func autoConvert_v1alpha1_RemoveTargetPortalResponse_To_internal_RemoveTargetPortalResponse(in *v1alpha1.RemoveTargetPortalResponse, out *internal.RemoveTargetPortalResponse) error {
	return nil
}

// Convert_v1alpha1_RemoveTargetPortalResponse_To_internal_RemoveTargetPortalResponse is an autogenerated conversion function.
func Convert_v1alpha1_RemoveTargetPortalResponse_To_internal_RemoveTargetPortalResponse(in *v1alpha1.RemoveTargetPortalResponse, out *internal.RemoveTargetPortalResponse) error {
	return autoConvert_v1alpha1_RemoveTargetPortalResponse_To_internal_RemoveTargetPortalResponse(in, out)
}

func autoConvert_internal_RemoveTargetPortalResponse_To_v1alpha1_RemoveTargetPortalResponse(in *internal.RemoveTargetPortalResponse, out *v1alpha1.RemoveTargetPortalResponse) error {
	return nil
}

// Convert_internal_RemoveTargetPortalResponse_To_v1alpha1_RemoveTargetPortalResponse is an autogenerated conversion function.
func Convert_internal_RemoveTargetPortalResponse_To_v1alpha1_RemoveTargetPortalResponse(in *internal.RemoveTargetPortalResponse, out *v1alpha1.RemoveTargetPortalResponse) error {
	return autoConvert_internal_RemoveTargetPortalResponse_To_v1alpha1_RemoveTargetPortalResponse(in, out)
}

func autoConvert_v1alpha1_TargetPortal_To_internal_TargetPortal(in *v1alpha1.TargetPortal, out *internal.TargetPortal) error {
	out.TargetAddress = in.TargetAddress
	out.TargetPort = in.TargetPort
	return nil
}

// Convert_v1alpha1_TargetPortal_To_internal_TargetPortal is an autogenerated conversion function.
func Convert_v1alpha1_TargetPortal_To_internal_TargetPortal(in *v1alpha1.TargetPortal, out *internal.TargetPortal) error {
	return autoConvert_v1alpha1_TargetPortal_To_internal_TargetPortal(in, out)
}

func autoConvert_internal_TargetPortal_To_v1alpha1_TargetPortal(in *internal.TargetPortal, out *v1alpha1.TargetPortal) error {
	out.TargetAddress = in.TargetAddress
	out.TargetPort = in.TargetPort
	return nil
}

// Convert_internal_TargetPortal_To_v1alpha1_TargetPortal is an autogenerated conversion function.
func Convert_internal_TargetPortal_To_v1alpha1_TargetPortal(in *internal.TargetPortal, out *v1alpha1.TargetPortal) error {
	return autoConvert_internal_TargetPortal_To_v1alpha1_TargetPortal(in, out)
}
