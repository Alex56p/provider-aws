//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNode) DeepCopyInto(out *ClusterNode) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNode.
func (in *ClusterNode) DeepCopy() *ClusterNode {
	if in == nil {
		return nil
	}
	out := new(ClusterNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObservation) DeepCopyInto(out *ClusterObservation) {
	*out = *in
	if in.ClusterCreateTime != nil {
		in, out := &in.ClusterCreateTime, &out.ClusterCreateTime
		*out = (*in).DeepCopy()
	}
	if in.ClusterNodes != nil {
		in, out := &in.ClusterNodes, &out.ClusterNodes
		*out = make([]ClusterNode, len(*in))
		copy(*out, *in)
	}
	if in.ClusterParameterGroups != nil {
		in, out := &in.ClusterParameterGroups, &out.ClusterParameterGroups
		*out = make([]ClusterParameterGroupStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ClusterSnapshotCopyStatus = in.ClusterSnapshotCopyStatus
	out.DataTransferProgress = in.DataTransferProgress
	if in.DeferredMaintenanceWindows != nil {
		in, out := &in.DeferredMaintenanceWindows, &out.DeferredMaintenanceWindows
		*out = make([]DeferredMaintenanceWindow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ElasticIPStatus = in.ElasticIPStatus
	out.Endpoint = in.Endpoint
	if in.ExpectedNextSnapshotScheduleTime != nil {
		in, out := &in.ExpectedNextSnapshotScheduleTime, &out.ExpectedNextSnapshotScheduleTime
		*out = (*in).DeepCopy()
	}
	out.HSMStatus = in.HSMStatus
	if in.NextMaintenanceWindowStartTime != nil {
		in, out := &in.NextMaintenanceWindowStartTime, &out.NextMaintenanceWindowStartTime
		*out = (*in).DeepCopy()
	}
	if in.PendingActions != nil {
		in, out := &in.PendingActions, &out.PendingActions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObservation.
func (in *ClusterObservation) DeepCopy() *ClusterObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameterGroupStatus) DeepCopyInto(out *ClusterParameterGroupStatus) {
	*out = *in
	if in.ClusterParameterStatusList != nil {
		in, out := &in.ClusterParameterStatusList, &out.ClusterParameterStatusList
		*out = make([]ClusterParameterStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameterGroupStatus.
func (in *ClusterParameterGroupStatus) DeepCopy() *ClusterParameterGroupStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterParameterGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameterStatus) DeepCopyInto(out *ClusterParameterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameterStatus.
func (in *ClusterParameterStatus) DeepCopy() *ClusterParameterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterParameterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameters) DeepCopyInto(out *ClusterParameters) {
	*out = *in
	if in.AllowVersionUpgrade != nil {
		in, out := &in.AllowVersionUpgrade, &out.AllowVersionUpgrade
		*out = new(bool)
		**out = **in
	}
	if in.AutomatedSnapshotRetentionPeriod != nil {
		in, out := &in.AutomatedSnapshotRetentionPeriod, &out.AutomatedSnapshotRetentionPeriod
		*out = new(int32)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.ClusterParameterGroupName != nil {
		in, out := &in.ClusterParameterGroupName, &out.ClusterParameterGroupName
		*out = new(string)
		**out = **in
	}
	if in.ClusterSecurityGroups != nil {
		in, out := &in.ClusterSecurityGroups, &out.ClusterSecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClusterSecurityGroupRefs != nil {
		in, out := &in.ClusterSecurityGroupRefs, &out.ClusterSecurityGroupRefs
		*out = make([]v1.Reference, len(*in))
		copy(*out, *in)
	}
	if in.ClusterSecurityGroupSelector != nil {
		in, out := &in.ClusterSecurityGroupSelector, &out.ClusterSecurityGroupSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterSubnetGroupName != nil {
		in, out := &in.ClusterSubnetGroupName, &out.ClusterSubnetGroupName
		*out = new(string)
		**out = **in
	}
	if in.ClusterType != nil {
		in, out := &in.ClusterType, &out.ClusterType
		*out = new(string)
		**out = **in
	}
	if in.ClusterVersion != nil {
		in, out := &in.ClusterVersion, &out.ClusterVersion
		*out = new(string)
		**out = **in
	}
	if in.DBName != nil {
		in, out := &in.DBName, &out.DBName
		*out = new(string)
		**out = **in
	}
	if in.ElasticIP != nil {
		in, out := &in.ElasticIP, &out.ElasticIP
		*out = new(string)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.EnhancedVPCRouting != nil {
		in, out := &in.EnhancedVPCRouting, &out.EnhancedVPCRouting
		*out = new(bool)
		**out = **in
	}
	if in.FinalClusterSnapshotIdentifier != nil {
		in, out := &in.FinalClusterSnapshotIdentifier, &out.FinalClusterSnapshotIdentifier
		*out = new(string)
		**out = **in
	}
	if in.FinalClusterSnapshotRetentionPeriod != nil {
		in, out := &in.FinalClusterSnapshotRetentionPeriod, &out.FinalClusterSnapshotRetentionPeriod
		*out = new(int32)
		**out = **in
	}
	if in.HSMClientCertificateIdentifier != nil {
		in, out := &in.HSMClientCertificateIdentifier, &out.HSMClientCertificateIdentifier
		*out = new(string)
		**out = **in
	}
	if in.HSMConfigurationIdentifier != nil {
		in, out := &in.HSMConfigurationIdentifier, &out.HSMConfigurationIdentifier
		*out = new(string)
		**out = **in
	}
	if in.IAMRoles != nil {
		in, out := &in.IAMRoles, &out.IAMRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IAMRoleRefs != nil {
		in, out := &in.IAMRoleRefs, &out.IAMRoleRefs
		*out = make([]v1.Reference, len(*in))
		copy(*out, *in)
	}
	if in.IAMRoleSelector != nil {
		in, out := &in.IAMRoleSelector, &out.IAMRoleSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceTrackName != nil {
		in, out := &in.MaintenanceTrackName, &out.MaintenanceTrackName
		*out = new(string)
		**out = **in
	}
	if in.ManualSnapshotRetentionPeriod != nil {
		in, out := &in.ManualSnapshotRetentionPeriod, &out.ManualSnapshotRetentionPeriod
		*out = new(int32)
		**out = **in
	}
	if in.NewMasterUserPassword != nil {
		in, out := &in.NewMasterUserPassword, &out.NewMasterUserPassword
		*out = new(string)
		**out = **in
	}
	if in.NewClusterIdentifier != nil {
		in, out := &in.NewClusterIdentifier, &out.NewClusterIdentifier
		*out = new(string)
		**out = **in
	}
	if in.NumberOfNodes != nil {
		in, out := &in.NumberOfNodes, &out.NumberOfNodes
		*out = new(int32)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.PreferredMaintenanceWindow != nil {
		in, out := &in.PreferredMaintenanceWindow, &out.PreferredMaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.PubliclyAccessible != nil {
		in, out := &in.PubliclyAccessible, &out.PubliclyAccessible
		*out = new(bool)
		**out = **in
	}
	if in.SkipFinalClusterSnapshot != nil {
		in, out := &in.SkipFinalClusterSnapshot, &out.SkipFinalClusterSnapshot
		*out = new(bool)
		**out = **in
	}
	if in.SnapshotScheduleIdentifier != nil {
		in, out := &in.SnapshotScheduleIdentifier, &out.SnapshotScheduleIdentifier
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
	if in.VPCSecurityGroupIDs != nil {
		in, out := &in.VPCSecurityGroupIDs, &out.VPCSecurityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VPCSecurityGroupIDRefs != nil {
		in, out := &in.VPCSecurityGroupIDRefs, &out.VPCSecurityGroupIDRefs
		*out = make([]v1.Reference, len(*in))
		copy(*out, *in)
	}
	if in.VPCSecurityGroupIDSelector != nil {
		in, out := &in.VPCSecurityGroupIDSelector, &out.VPCSecurityGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameters.
func (in *ClusterParameters) DeepCopy() *ClusterParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSecurityGroupMembership) DeepCopyInto(out *ClusterSecurityGroupMembership) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSecurityGroupMembership.
func (in *ClusterSecurityGroupMembership) DeepCopy() *ClusterSecurityGroupMembership {
	if in == nil {
		return nil
	}
	out := new(ClusterSecurityGroupMembership)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSnapshotCopyStatus) DeepCopyInto(out *ClusterSnapshotCopyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSnapshotCopyStatus.
func (in *ClusterSnapshotCopyStatus) DeepCopy() *ClusterSnapshotCopyStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterSnapshotCopyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataTransferProgress) DeepCopyInto(out *DataTransferProgress) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataTransferProgress.
func (in *DataTransferProgress) DeepCopy() *DataTransferProgress {
	if in == nil {
		return nil
	}
	out := new(DataTransferProgress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeferredMaintenanceWindow) DeepCopyInto(out *DeferredMaintenanceWindow) {
	*out = *in
	if in.DeferMaintenanceEndTime != nil {
		in, out := &in.DeferMaintenanceEndTime, &out.DeferMaintenanceEndTime
		*out = (*in).DeepCopy()
	}
	if in.DeferMaintenanceStartTime != nil {
		in, out := &in.DeferMaintenanceStartTime, &out.DeferMaintenanceStartTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeferredMaintenanceWindow.
func (in *DeferredMaintenanceWindow) DeepCopy() *DeferredMaintenanceWindow {
	if in == nil {
		return nil
	}
	out := new(DeferredMaintenanceWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticIPStatus) DeepCopyInto(out *ElasticIPStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticIPStatus.
func (in *ElasticIPStatus) DeepCopy() *ElasticIPStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticIPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HSMStatus) DeepCopyInto(out *HSMStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HSMStatus.
func (in *HSMStatus) DeepCopy() *HSMStatus {
	if in == nil {
		return nil
	}
	out := new(HSMStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreStatus) DeepCopyInto(out *RestoreStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreStatus.
func (in *RestoreStatus) DeepCopy() *RestoreStatus {
	if in == nil {
		return nil
	}
	out := new(RestoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCSecurityGroupMembership) DeepCopyInto(out *VPCSecurityGroupMembership) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCSecurityGroupMembership.
func (in *VPCSecurityGroupMembership) DeepCopy() *VPCSecurityGroupMembership {
	if in == nil {
		return nil
	}
	out := new(VPCSecurityGroupMembership)
	in.DeepCopyInto(out)
	return out
}
