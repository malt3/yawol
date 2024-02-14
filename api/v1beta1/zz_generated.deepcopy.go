//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancer) DeepCopyInto(out *LoadBalancer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancer.
func (in *LoadBalancer) DeepCopy() *LoadBalancer {
	if in == nil {
		return nil
	}
	out := new(LoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerAdditionalNetwork) DeepCopyInto(out *LoadBalancerAdditionalNetwork) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerAdditionalNetwork.
func (in *LoadBalancerAdditionalNetwork) DeepCopy() *LoadBalancerAdditionalNetwork {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerAdditionalNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerDebugSettings) DeepCopyInto(out *LoadBalancerDebugSettings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerDebugSettings.
func (in *LoadBalancerDebugSettings) DeepCopy() *LoadBalancerDebugSettings {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerDebugSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerDefaultNetwork) DeepCopyInto(out *LoadBalancerDefaultNetwork) {
	*out = *in
	if in.FloatingNetID != nil {
		in, out := &in.FloatingNetID, &out.FloatingNetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerDefaultNetwork.
func (in *LoadBalancerDefaultNetwork) DeepCopy() *LoadBalancerDefaultNetwork {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerDefaultNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerEndpoint) DeepCopyInto(out *LoadBalancerEndpoint) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerEndpoint.
func (in *LoadBalancerEndpoint) DeepCopy() *LoadBalancerEndpoint {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerInfrastructure) DeepCopyInto(out *LoadBalancerInfrastructure) {
	*out = *in
	if in.FloatingNetID != nil {
		in, out := &in.FloatingNetID, &out.FloatingNetID
		*out = new(string)
		**out = **in
	}
	in.DefaultNetwork.DeepCopyInto(&out.DefaultNetwork)
	if in.AdditionalNetworks != nil {
		in, out := &in.AdditionalNetworks, &out.AdditionalNetworks
		*out = make([]LoadBalancerAdditionalNetwork, len(*in))
		copy(*out, *in)
	}
	in.Flavor.DeepCopyInto(&out.Flavor)
	in.Image.DeepCopyInto(&out.Image)
	out.AuthSecretRef = in.AuthSecretRef
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerInfrastructure.
func (in *LoadBalancerInfrastructure) DeepCopy() *LoadBalancerInfrastructure {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerInfrastructure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerList) DeepCopyInto(out *LoadBalancerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoadBalancer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerList.
func (in *LoadBalancerList) DeepCopy() *LoadBalancerList {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerLogForward) DeepCopyInto(out *LoadBalancerLogForward) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerLogForward.
func (in *LoadBalancerLogForward) DeepCopy() *LoadBalancerLogForward {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerLogForward)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerMachine) DeepCopyInto(out *LoadBalancerMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerMachine.
func (in *LoadBalancerMachine) DeepCopy() *LoadBalancerMachine {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerMachineList) DeepCopyInto(out *LoadBalancerMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoadBalancerMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerMachineList.
func (in *LoadBalancerMachineList) DeepCopy() *LoadBalancerMachineList {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerMachineMetric) DeepCopyInto(out *LoadBalancerMachineMetric) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerMachineMetric.
func (in *LoadBalancerMachineMetric) DeepCopy() *LoadBalancerMachineMetric {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerMachineMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerMachineSpec) DeepCopyInto(out *LoadBalancerMachineSpec) {
	*out = *in
	in.Infrastructure.DeepCopyInto(&out.Infrastructure)
	out.LoadBalancerRef = in.LoadBalancerRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerMachineSpec.
func (in *LoadBalancerMachineSpec) DeepCopy() *LoadBalancerMachineSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerMachineStatus) DeepCopyInto(out *LoadBalancerMachineStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = new([]v1.NodeCondition)
		if **in != nil {
			in, out := *in, *out
			*out = make([]v1.NodeCondition, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new([]LoadBalancerMachineMetric)
		if **in != nil {
			in, out := *in, *out
			*out = make([]LoadBalancerMachineMetric, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.CreationTimestamp != nil {
		in, out := &in.CreationTimestamp, &out.CreationTimestamp
		*out = (*in).DeepCopy()
	}
	if in.LastOpenstackReconcile != nil {
		in, out := &in.LastOpenstackReconcile, &out.LastOpenstackReconcile
		*out = (*in).DeepCopy()
	}
	if in.ServerID != nil {
		in, out := &in.ServerID, &out.ServerID
		*out = new(string)
		**out = **in
	}
	if in.PortID != nil {
		in, out := &in.PortID, &out.PortID
		*out = new(string)
		**out = **in
	}
	if in.DefaultPortID != nil {
		in, out := &in.DefaultPortID, &out.DefaultPortID
		*out = new(string)
		**out = **in
	}
	if in.DefaultPortName != nil {
		in, out := &in.DefaultPortName, &out.DefaultPortName
		*out = new(string)
		**out = **in
	}
	if in.DefaultPortIP != nil {
		in, out := &in.DefaultPortIP, &out.DefaultPortIP
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountName != nil {
		in, out := &in.ServiceAccountName, &out.ServiceAccountName
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountSecretName != nil {
		in, out := &in.ServiceAccountSecretName, &out.ServiceAccountSecretName
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.RoleBindingName != nil {
		in, out := &in.RoleBindingName, &out.RoleBindingName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerMachineStatus.
func (in *LoadBalancerMachineStatus) DeepCopy() *LoadBalancerMachineStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerMachineTemplateSpec) DeepCopyInto(out *LoadBalancerMachineTemplateSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerMachineTemplateSpec.
func (in *LoadBalancerMachineTemplateSpec) DeepCopy() *LoadBalancerMachineTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerMachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerOptions) DeepCopyInto(out *LoadBalancerOptions) {
	*out = *in
	if in.LoadBalancerSourceRanges != nil {
		in, out := &in.LoadBalancerSourceRanges, &out.LoadBalancerSourceRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TCPProxyProtocolPortsFilter != nil {
		in, out := &in.TCPProxyProtocolPortsFilter, &out.TCPProxyProtocolPortsFilter
		*out = make([]int32, len(*in))
		copy(*out, *in)
	}
	if in.TCPIdleTimeout != nil {
		in, out := &in.TCPIdleTimeout, &out.TCPIdleTimeout
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.UDPIdleTimeout != nil {
		in, out := &in.UDPIdleTimeout, &out.UDPIdleTimeout
		*out = new(metav1.Duration)
		**out = **in
	}
	in.LogForward.DeepCopyInto(&out.LogForward)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerOptions.
func (in *LoadBalancerOptions) DeepCopy() *LoadBalancerOptions {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerRef) DeepCopyInto(out *LoadBalancerRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerRef.
func (in *LoadBalancerRef) DeepCopy() *LoadBalancerRef {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSet) DeepCopyInto(out *LoadBalancerSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSet.
func (in *LoadBalancerSet) DeepCopy() *LoadBalancerSet {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSetList) DeepCopyInto(out *LoadBalancerSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoadBalancerSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSetList.
func (in *LoadBalancerSetList) DeepCopy() *LoadBalancerSetList {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSetSpec) DeepCopyInto(out *LoadBalancerSetSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSetSpec.
func (in *LoadBalancerSetSpec) DeepCopy() *LoadBalancerSetSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSetStatus) DeepCopyInto(out *LoadBalancerSetStatus) {
	*out = *in
	if in.AvailableReplicas != nil {
		in, out := &in.AvailableReplicas, &out.AvailableReplicas
		*out = new(int)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ReadyReplicas != nil {
		in, out := &in.ReadyReplicas, &out.ReadyReplicas
		*out = new(int)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSetStatus.
func (in *LoadBalancerSetStatus) DeepCopy() *LoadBalancerSetStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSpec) DeepCopyInto(out *LoadBalancerSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.ExistingFloatingIP != nil {
		in, out := &in.ExistingFloatingIP, &out.ExistingFloatingIP
		*out = new(string)
		**out = **in
	}
	out.DebugSettings = in.DebugSettings
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]LoadBalancerEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]v1.ServicePort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Infrastructure.DeepCopyInto(&out.Infrastructure)
	in.Options.DeepCopyInto(&out.Options)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSpec.
func (in *LoadBalancerSpec) DeepCopy() *LoadBalancerSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerStatus) DeepCopyInto(out *LoadBalancerStatus) {
	*out = *in
	if in.ReadyReplicas != nil {
		in, out := &in.ReadyReplicas, &out.ReadyReplicas
		*out = new(int)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int)
		**out = **in
	}
	if in.ExternalIP != nil {
		in, out := &in.ExternalIP, &out.ExternalIP
		*out = new(string)
		**out = **in
	}
	if in.FloatingID != nil {
		in, out := &in.FloatingID, &out.FloatingID
		*out = new(string)
		**out = **in
	}
	if in.FloatingName != nil {
		in, out := &in.FloatingName, &out.FloatingName
		*out = new(string)
		**out = **in
	}
	if in.PortID != nil {
		in, out := &in.PortID, &out.PortID
		*out = new(string)
		**out = **in
	}
	if in.PortName != nil {
		in, out := &in.PortName, &out.PortName
		*out = new(string)
		**out = **in
	}
	if in.PortIP != nil {
		in, out := &in.PortIP, &out.PortIP
		*out = new(string)
		**out = **in
	}
	if in.ServerGroupID != nil {
		in, out := &in.ServerGroupID, &out.ServerGroupID
		*out = new(string)
		**out = **in
	}
	if in.ServerGroupName != nil {
		in, out := &in.ServerGroupName, &out.ServerGroupName
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupID != nil {
		in, out := &in.SecurityGroupID, &out.SecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupName != nil {
		in, out := &in.SecurityGroupName, &out.SecurityGroupName
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupIDOld != nil {
		in, out := &in.SecurityGroupIDOld, &out.SecurityGroupIDOld
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupNameOld != nil {
		in, out := &in.SecurityGroupNameOld, &out.SecurityGroupNameOld
		*out = new(string)
		**out = **in
	}
	if in.LastOpenstackReconcile != nil {
		in, out := &in.LastOpenstackReconcile, &out.LastOpenstackReconcile
		*out = (*in).DeepCopy()
	}
	if in.OpenstackReconcileHash != nil {
		in, out := &in.OpenstackReconcileHash, &out.OpenstackReconcileHash
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerStatus.
func (in *LoadBalancerStatus) DeepCopy() *LoadBalancerStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackFlavorRef) DeepCopyInto(out *OpenstackFlavorRef) {
	*out = *in
	if in.FlavorID != nil {
		in, out := &in.FlavorID, &out.FlavorID
		*out = new(string)
		**out = **in
	}
	if in.FlavorName != nil {
		in, out := &in.FlavorName, &out.FlavorName
		*out = new(string)
		**out = **in
	}
	if in.FlavorSearch != nil {
		in, out := &in.FlavorSearch, &out.FlavorSearch
		*out = new(string)
		**out = **in
	}
	if in.FlavorIDOld != nil {
		in, out := &in.FlavorIDOld, &out.FlavorIDOld
		*out = new(string)
		**out = **in
	}
	if in.FlavorNameOld != nil {
		in, out := &in.FlavorNameOld, &out.FlavorNameOld
		*out = new(string)
		**out = **in
	}
	if in.FlavorSearchOld != nil {
		in, out := &in.FlavorSearchOld, &out.FlavorSearchOld
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackFlavorRef.
func (in *OpenstackFlavorRef) DeepCopy() *OpenstackFlavorRef {
	if in == nil {
		return nil
	}
	out := new(OpenstackFlavorRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackImageRef) DeepCopyInto(out *OpenstackImageRef) {
	*out = *in
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.ImageName != nil {
		in, out := &in.ImageName, &out.ImageName
		*out = new(string)
		**out = **in
	}
	if in.ImageSearch != nil {
		in, out := &in.ImageSearch, &out.ImageSearch
		*out = new(string)
		**out = **in
	}
	if in.ImageIDOld != nil {
		in, out := &in.ImageIDOld, &out.ImageIDOld
		*out = new(string)
		**out = **in
	}
	if in.ImageNameOld != nil {
		in, out := &in.ImageNameOld, &out.ImageNameOld
		*out = new(string)
		**out = **in
	}
	if in.ImageSearchOld != nil {
		in, out := &in.ImageSearchOld, &out.ImageSearchOld
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackImageRef.
func (in *OpenstackImageRef) DeepCopy() *OpenstackImageRef {
	if in == nil {
		return nil
	}
	out := new(OpenstackImageRef)
	in.DeepCopyInto(out)
	return out
}
