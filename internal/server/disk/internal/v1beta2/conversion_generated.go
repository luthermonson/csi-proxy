// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1beta2

import (
	v1beta2 "github.com/kubernetes-csi/csi-proxy/client/api/disk/v1beta2"
	internal "github.com/kubernetes-csi/csi-proxy/internal/server/disk/internal"
)

func autoConvert_v1beta2_DiskIDs_To_internal_DiskIDs(in *v1beta2.DiskIDs, out *internal.DiskIDs) error {
	return nil
}

// Convert_v1beta2_DiskIDs_To_internal_DiskIDs is an autogenerated conversion function.
func Convert_v1beta2_DiskIDs_To_internal_DiskIDs(in *v1beta2.DiskIDs, out *internal.DiskIDs) error {
	return autoConvert_v1beta2_DiskIDs_To_internal_DiskIDs(in, out)
}

func autoConvert_internal_DiskIDs_To_v1beta2_DiskIDs(in *internal.DiskIDs, out *v1beta2.DiskIDs) error {
	return nil
}

// Convert_internal_DiskIDs_To_v1beta2_DiskIDs is an autogenerated conversion function.
func Convert_internal_DiskIDs_To_v1beta2_DiskIDs(in *internal.DiskIDs, out *v1beta2.DiskIDs) error {
	return autoConvert_internal_DiskIDs_To_v1beta2_DiskIDs(in, out)
}

func autoConvert_v1beta2_DiskLocation_To_internal_DiskLocation(in *v1beta2.DiskLocation, out *internal.DiskLocation) error {
	out.Adapter = in.Adapter
	out.Bus = in.Bus
	out.Target = in.Target
	out.LUNID = in.LUNID
	return nil
}

// Convert_v1beta2_DiskLocation_To_internal_DiskLocation is an autogenerated conversion function.
func Convert_v1beta2_DiskLocation_To_internal_DiskLocation(in *v1beta2.DiskLocation, out *internal.DiskLocation) error {
	return autoConvert_v1beta2_DiskLocation_To_internal_DiskLocation(in, out)
}

func autoConvert_internal_DiskLocation_To_v1beta2_DiskLocation(in *internal.DiskLocation, out *v1beta2.DiskLocation) error {
	out.Adapter = in.Adapter
	out.Bus = in.Bus
	out.Target = in.Target
	out.LUNID = in.LUNID
	return nil
}

// Convert_internal_DiskLocation_To_v1beta2_DiskLocation is an autogenerated conversion function.
func Convert_internal_DiskLocation_To_v1beta2_DiskLocation(in *internal.DiskLocation, out *v1beta2.DiskLocation) error {
	return autoConvert_internal_DiskLocation_To_v1beta2_DiskLocation(in, out)
}

func autoConvert_v1beta2_DiskStatsRequest_To_internal_DiskStatsRequest(in *v1beta2.DiskStatsRequest, out *internal.DiskStatsRequest) error {
	out.DiskID = in.DiskID
	return nil
}

// Convert_v1beta2_DiskStatsRequest_To_internal_DiskStatsRequest is an autogenerated conversion function.
func Convert_v1beta2_DiskStatsRequest_To_internal_DiskStatsRequest(in *v1beta2.DiskStatsRequest, out *internal.DiskStatsRequest) error {
	return autoConvert_v1beta2_DiskStatsRequest_To_internal_DiskStatsRequest(in, out)
}

func autoConvert_internal_DiskStatsRequest_To_v1beta2_DiskStatsRequest(in *internal.DiskStatsRequest, out *v1beta2.DiskStatsRequest) error {
	out.DiskID = in.DiskID
	return nil
}

// Convert_internal_DiskStatsRequest_To_v1beta2_DiskStatsRequest is an autogenerated conversion function.
func Convert_internal_DiskStatsRequest_To_v1beta2_DiskStatsRequest(in *internal.DiskStatsRequest, out *v1beta2.DiskStatsRequest) error {
	return autoConvert_internal_DiskStatsRequest_To_v1beta2_DiskStatsRequest(in, out)
}

func autoConvert_v1beta2_DiskStatsResponse_To_internal_DiskStatsResponse(in *v1beta2.DiskStatsResponse, out *internal.DiskStatsResponse) error {
	out.DiskSize = in.DiskSize
	return nil
}

// Convert_v1beta2_DiskStatsResponse_To_internal_DiskStatsResponse is an autogenerated conversion function.
func Convert_v1beta2_DiskStatsResponse_To_internal_DiskStatsResponse(in *v1beta2.DiskStatsResponse, out *internal.DiskStatsResponse) error {
	return autoConvert_v1beta2_DiskStatsResponse_To_internal_DiskStatsResponse(in, out)
}

func autoConvert_internal_DiskStatsResponse_To_v1beta2_DiskStatsResponse(in *internal.DiskStatsResponse, out *v1beta2.DiskStatsResponse) error {
	out.DiskSize = in.DiskSize
	return nil
}

// Convert_internal_DiskStatsResponse_To_v1beta2_DiskStatsResponse is an autogenerated conversion function.
func Convert_internal_DiskStatsResponse_To_v1beta2_DiskStatsResponse(in *internal.DiskStatsResponse, out *v1beta2.DiskStatsResponse) error {
	return autoConvert_internal_DiskStatsResponse_To_v1beta2_DiskStatsResponse(in, out)
}

func autoConvert_v1beta2_GetAttachStateRequest_To_internal_GetAttachStateRequest(in *v1beta2.GetAttachStateRequest, out *internal.GetAttachStateRequest) error {
	out.DiskID = in.DiskID
	return nil
}

// Convert_v1beta2_GetAttachStateRequest_To_internal_GetAttachStateRequest is an autogenerated conversion function.
func Convert_v1beta2_GetAttachStateRequest_To_internal_GetAttachStateRequest(in *v1beta2.GetAttachStateRequest, out *internal.GetAttachStateRequest) error {
	return autoConvert_v1beta2_GetAttachStateRequest_To_internal_GetAttachStateRequest(in, out)
}

func autoConvert_internal_GetAttachStateRequest_To_v1beta2_GetAttachStateRequest(in *internal.GetAttachStateRequest, out *v1beta2.GetAttachStateRequest) error {
	out.DiskID = in.DiskID
	return nil
}

// Convert_internal_GetAttachStateRequest_To_v1beta2_GetAttachStateRequest is an autogenerated conversion function.
func Convert_internal_GetAttachStateRequest_To_v1beta2_GetAttachStateRequest(in *internal.GetAttachStateRequest, out *v1beta2.GetAttachStateRequest) error {
	return autoConvert_internal_GetAttachStateRequest_To_v1beta2_GetAttachStateRequest(in, out)
}

func autoConvert_v1beta2_GetAttachStateResponse_To_internal_GetAttachStateResponse(in *v1beta2.GetAttachStateResponse, out *internal.GetAttachStateResponse) error {
	out.IsOnline = in.IsOnline
	return nil
}

// Convert_v1beta2_GetAttachStateResponse_To_internal_GetAttachStateResponse is an autogenerated conversion function.
func Convert_v1beta2_GetAttachStateResponse_To_internal_GetAttachStateResponse(in *v1beta2.GetAttachStateResponse, out *internal.GetAttachStateResponse) error {
	return autoConvert_v1beta2_GetAttachStateResponse_To_internal_GetAttachStateResponse(in, out)
}

func autoConvert_internal_GetAttachStateResponse_To_v1beta2_GetAttachStateResponse(in *internal.GetAttachStateResponse, out *v1beta2.GetAttachStateResponse) error {
	out.IsOnline = in.IsOnline
	return nil
}

// Convert_internal_GetAttachStateResponse_To_v1beta2_GetAttachStateResponse is an autogenerated conversion function.
func Convert_internal_GetAttachStateResponse_To_v1beta2_GetAttachStateResponse(in *internal.GetAttachStateResponse, out *v1beta2.GetAttachStateResponse) error {
	return autoConvert_internal_GetAttachStateResponse_To_v1beta2_GetAttachStateResponse(in, out)
}

func autoConvert_v1beta2_ListDiskIDsRequest_To_internal_ListDiskIDsRequest(in *v1beta2.ListDiskIDsRequest, out *internal.ListDiskIDsRequest) error {
	return nil
}

// Convert_v1beta2_ListDiskIDsRequest_To_internal_ListDiskIDsRequest is an autogenerated conversion function.
func Convert_v1beta2_ListDiskIDsRequest_To_internal_ListDiskIDsRequest(in *v1beta2.ListDiskIDsRequest, out *internal.ListDiskIDsRequest) error {
	return autoConvert_v1beta2_ListDiskIDsRequest_To_internal_ListDiskIDsRequest(in, out)
}

func autoConvert_internal_ListDiskIDsRequest_To_v1beta2_ListDiskIDsRequest(in *internal.ListDiskIDsRequest, out *v1beta2.ListDiskIDsRequest) error {
	return nil
}

// Convert_internal_ListDiskIDsRequest_To_v1beta2_ListDiskIDsRequest is an autogenerated conversion function.
func Convert_internal_ListDiskIDsRequest_To_v1beta2_ListDiskIDsRequest(in *internal.ListDiskIDsRequest, out *v1beta2.ListDiskIDsRequest) error {
	return autoConvert_internal_ListDiskIDsRequest_To_v1beta2_ListDiskIDsRequest(in, out)
}

// detected external conversion function
// Convert_v1beta2_ListDiskIDsResponse_To_internal_ListDiskIDsResponse(in *v1beta2.ListDiskIDsResponse, out *internal.ListDiskIDsResponse) error
// skipping generation of the auto function

// detected external conversion function
// Convert_internal_ListDiskIDsResponse_To_v1beta2_ListDiskIDsResponse(in *internal.ListDiskIDsResponse, out *v1beta2.ListDiskIDsResponse) error
// skipping generation of the auto function

func autoConvert_v1beta2_ListDiskLocationsRequest_To_internal_ListDiskLocationsRequest(in *v1beta2.ListDiskLocationsRequest, out *internal.ListDiskLocationsRequest) error {
	return nil
}

// Convert_v1beta2_ListDiskLocationsRequest_To_internal_ListDiskLocationsRequest is an autogenerated conversion function.
func Convert_v1beta2_ListDiskLocationsRequest_To_internal_ListDiskLocationsRequest(in *v1beta2.ListDiskLocationsRequest, out *internal.ListDiskLocationsRequest) error {
	return autoConvert_v1beta2_ListDiskLocationsRequest_To_internal_ListDiskLocationsRequest(in, out)
}

func autoConvert_internal_ListDiskLocationsRequest_To_v1beta2_ListDiskLocationsRequest(in *internal.ListDiskLocationsRequest, out *v1beta2.ListDiskLocationsRequest) error {
	return nil
}

// Convert_internal_ListDiskLocationsRequest_To_v1beta2_ListDiskLocationsRequest is an autogenerated conversion function.
func Convert_internal_ListDiskLocationsRequest_To_v1beta2_ListDiskLocationsRequest(in *internal.ListDiskLocationsRequest, out *v1beta2.ListDiskLocationsRequest) error {
	return autoConvert_internal_ListDiskLocationsRequest_To_v1beta2_ListDiskLocationsRequest(in, out)
}

// detected external conversion function
// Convert_v1beta2_ListDiskLocationsResponse_To_internal_ListDiskLocationsResponse(in *v1beta2.ListDiskLocationsResponse, out *internal.ListDiskLocationsResponse) error
// skipping generation of the auto function

// detected external conversion function
// Convert_internal_ListDiskLocationsResponse_To_v1beta2_ListDiskLocationsResponse(in *internal.ListDiskLocationsResponse, out *v1beta2.ListDiskLocationsResponse) error
// skipping generation of the auto function

// detected external conversion function
// Convert_v1beta2_PartitionDiskRequest_To_internal_PartitionDiskRequest(in *v1beta2.PartitionDiskRequest, out *internal.PartitionDiskRequest) error
// skipping generation of the auto function

func autoConvert_internal_PartitionDiskRequest_To_v1beta2_PartitionDiskRequest(in *internal.PartitionDiskRequest, out *v1beta2.PartitionDiskRequest) error {
	return nil
}

// Convert_internal_PartitionDiskRequest_To_v1beta2_PartitionDiskRequest is an autogenerated conversion function.
func Convert_internal_PartitionDiskRequest_To_v1beta2_PartitionDiskRequest(in *internal.PartitionDiskRequest, out *v1beta2.PartitionDiskRequest) error {
	return autoConvert_internal_PartitionDiskRequest_To_v1beta2_PartitionDiskRequest(in, out)
}

func autoConvert_v1beta2_PartitionDiskResponse_To_internal_PartitionDiskResponse(in *v1beta2.PartitionDiskResponse, out *internal.PartitionDiskResponse) error {
	return nil
}

// Convert_v1beta2_PartitionDiskResponse_To_internal_PartitionDiskResponse is an autogenerated conversion function.
func Convert_v1beta2_PartitionDiskResponse_To_internal_PartitionDiskResponse(in *v1beta2.PartitionDiskResponse, out *internal.PartitionDiskResponse) error {
	return autoConvert_v1beta2_PartitionDiskResponse_To_internal_PartitionDiskResponse(in, out)
}

func autoConvert_internal_PartitionDiskResponse_To_v1beta2_PartitionDiskResponse(in *internal.PartitionDiskResponse, out *v1beta2.PartitionDiskResponse) error {
	return nil
}

// Convert_internal_PartitionDiskResponse_To_v1beta2_PartitionDiskResponse is an autogenerated conversion function.
func Convert_internal_PartitionDiskResponse_To_v1beta2_PartitionDiskResponse(in *internal.PartitionDiskResponse, out *v1beta2.PartitionDiskResponse) error {
	return autoConvert_internal_PartitionDiskResponse_To_v1beta2_PartitionDiskResponse(in, out)
}

func autoConvert_v1beta2_RescanRequest_To_internal_RescanRequest(in *v1beta2.RescanRequest, out *internal.RescanRequest) error {
	return nil
}

// Convert_v1beta2_RescanRequest_To_internal_RescanRequest is an autogenerated conversion function.
func Convert_v1beta2_RescanRequest_To_internal_RescanRequest(in *v1beta2.RescanRequest, out *internal.RescanRequest) error {
	return autoConvert_v1beta2_RescanRequest_To_internal_RescanRequest(in, out)
}

func autoConvert_internal_RescanRequest_To_v1beta2_RescanRequest(in *internal.RescanRequest, out *v1beta2.RescanRequest) error {
	return nil
}

// Convert_internal_RescanRequest_To_v1beta2_RescanRequest is an autogenerated conversion function.
func Convert_internal_RescanRequest_To_v1beta2_RescanRequest(in *internal.RescanRequest, out *v1beta2.RescanRequest) error {
	return autoConvert_internal_RescanRequest_To_v1beta2_RescanRequest(in, out)
}

func autoConvert_v1beta2_RescanResponse_To_internal_RescanResponse(in *v1beta2.RescanResponse, out *internal.RescanResponse) error {
	return nil
}

// Convert_v1beta2_RescanResponse_To_internal_RescanResponse is an autogenerated conversion function.
func Convert_v1beta2_RescanResponse_To_internal_RescanResponse(in *v1beta2.RescanResponse, out *internal.RescanResponse) error {
	return autoConvert_v1beta2_RescanResponse_To_internal_RescanResponse(in, out)
}

func autoConvert_internal_RescanResponse_To_v1beta2_RescanResponse(in *internal.RescanResponse, out *v1beta2.RescanResponse) error {
	return nil
}

// Convert_internal_RescanResponse_To_v1beta2_RescanResponse is an autogenerated conversion function.
func Convert_internal_RescanResponse_To_v1beta2_RescanResponse(in *internal.RescanResponse, out *v1beta2.RescanResponse) error {
	return autoConvert_internal_RescanResponse_To_v1beta2_RescanResponse(in, out)
}

func autoConvert_v1beta2_SetAttachStateRequest_To_internal_SetAttachStateRequest(in *v1beta2.SetAttachStateRequest, out *internal.SetAttachStateRequest) error {
	out.DiskID = in.DiskID
	out.IsOnline = in.IsOnline
	return nil
}

// Convert_v1beta2_SetAttachStateRequest_To_internal_SetAttachStateRequest is an autogenerated conversion function.
func Convert_v1beta2_SetAttachStateRequest_To_internal_SetAttachStateRequest(in *v1beta2.SetAttachStateRequest, out *internal.SetAttachStateRequest) error {
	return autoConvert_v1beta2_SetAttachStateRequest_To_internal_SetAttachStateRequest(in, out)
}

func autoConvert_internal_SetAttachStateRequest_To_v1beta2_SetAttachStateRequest(in *internal.SetAttachStateRequest, out *v1beta2.SetAttachStateRequest) error {
	out.DiskID = in.DiskID
	out.IsOnline = in.IsOnline
	return nil
}

// Convert_internal_SetAttachStateRequest_To_v1beta2_SetAttachStateRequest is an autogenerated conversion function.
func Convert_internal_SetAttachStateRequest_To_v1beta2_SetAttachStateRequest(in *internal.SetAttachStateRequest, out *v1beta2.SetAttachStateRequest) error {
	return autoConvert_internal_SetAttachStateRequest_To_v1beta2_SetAttachStateRequest(in, out)
}

func autoConvert_v1beta2_SetAttachStateResponse_To_internal_SetAttachStateResponse(in *v1beta2.SetAttachStateResponse, out *internal.SetAttachStateResponse) error {
	return nil
}

// Convert_v1beta2_SetAttachStateResponse_To_internal_SetAttachStateResponse is an autogenerated conversion function.
func Convert_v1beta2_SetAttachStateResponse_To_internal_SetAttachStateResponse(in *v1beta2.SetAttachStateResponse, out *internal.SetAttachStateResponse) error {
	return autoConvert_v1beta2_SetAttachStateResponse_To_internal_SetAttachStateResponse(in, out)
}

func autoConvert_internal_SetAttachStateResponse_To_v1beta2_SetAttachStateResponse(in *internal.SetAttachStateResponse, out *v1beta2.SetAttachStateResponse) error {
	return nil
}

// Convert_internal_SetAttachStateResponse_To_v1beta2_SetAttachStateResponse is an autogenerated conversion function.
func Convert_internal_SetAttachStateResponse_To_v1beta2_SetAttachStateResponse(in *internal.SetAttachStateResponse, out *v1beta2.SetAttachStateResponse) error {
	return autoConvert_internal_SetAttachStateResponse_To_v1beta2_SetAttachStateResponse(in, out)
}
