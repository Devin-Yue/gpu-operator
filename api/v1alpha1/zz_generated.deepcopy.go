//go:build !ignore_autogenerated

/*
Copyright 2021.

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
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerProbeSpec) DeepCopyInto(out *ContainerProbeSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerProbeSpec.
func (in *ContainerProbeSpec) DeepCopy() *ContainerProbeSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerProbeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverCertConfigSpec) DeepCopyInto(out *DriverCertConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverCertConfigSpec.
func (in *DriverCertConfigSpec) DeepCopy() *DriverCertConfigSpec {
	if in == nil {
		return nil
	}
	out := new(DriverCertConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverLicensingConfigSpec) DeepCopyInto(out *DriverLicensingConfigSpec) {
	*out = *in
	if in.NLSEnabled != nil {
		in, out := &in.NLSEnabled, &out.NLSEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverLicensingConfigSpec.
func (in *DriverLicensingConfigSpec) DeepCopy() *DriverLicensingConfigSpec {
	if in == nil {
		return nil
	}
	out := new(DriverLicensingConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverManagerSpec) DeepCopyInto(out *DriverManagerSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]EnvVar, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverManagerSpec.
func (in *DriverManagerSpec) DeepCopy() *DriverManagerSpec {
	if in == nil {
		return nil
	}
	out := new(DriverManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverRepoConfigSpec) DeepCopyInto(out *DriverRepoConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverRepoConfigSpec.
func (in *DriverRepoConfigSpec) DeepCopy() *DriverRepoConfigSpec {
	if in == nil {
		return nil
	}
	out := new(DriverRepoConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVar) DeepCopyInto(out *EnvVar) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVar.
func (in *EnvVar) DeepCopy() *EnvVar {
	if in == nil {
		return nil
	}
	out := new(EnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GPUDirectRDMASpec) DeepCopyInto(out *GPUDirectRDMASpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.UseHostMOFED != nil {
		in, out := &in.UseHostMOFED, &out.UseHostMOFED
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GPUDirectRDMASpec.
func (in *GPUDirectRDMASpec) DeepCopy() *GPUDirectRDMASpec {
	if in == nil {
		return nil
	}
	out := new(GPUDirectRDMASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GPUDirectStorageSpec) DeepCopyInto(out *GPUDirectStorageSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]EnvVar, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GPUDirectStorageSpec.
func (in *GPUDirectStorageSpec) DeepCopy() *GPUDirectStorageSpec {
	if in == nil {
		return nil
	}
	out := new(GPUDirectStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelModuleConfigSpec) DeepCopyInto(out *KernelModuleConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelModuleConfigSpec.
func (in *KernelModuleConfigSpec) DeepCopy() *KernelModuleConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KernelModuleConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NVIDIADriver) DeepCopyInto(out *NVIDIADriver) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NVIDIADriver.
func (in *NVIDIADriver) DeepCopy() *NVIDIADriver {
	if in == nil {
		return nil
	}
	out := new(NVIDIADriver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NVIDIADriver) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NVIDIADriverList) DeepCopyInto(out *NVIDIADriverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NVIDIADriver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NVIDIADriverList.
func (in *NVIDIADriverList) DeepCopy() *NVIDIADriverList {
	if in == nil {
		return nil
	}
	out := new(NVIDIADriverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NVIDIADriverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NVIDIADriverSpec) DeepCopyInto(out *NVIDIADriverSpec) {
	*out = *in
	if in.UsePrecompiled != nil {
		in, out := &in.UsePrecompiled, &out.UsePrecompiled
		*out = new(bool)
		**out = **in
	}
	if in.StartupProbe != nil {
		in, out := &in.StartupProbe, &out.StartupProbe
		*out = new(ContainerProbeSpec)
		**out = **in
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(ContainerProbeSpec)
		**out = **in
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(ContainerProbeSpec)
		**out = **in
	}
	if in.GPUDirectRDMA != nil {
		in, out := &in.GPUDirectRDMA, &out.GPUDirectRDMA
		*out = new(GPUDirectRDMASpec)
		(*in).DeepCopyInto(*out)
	}
	if in.GPUDirectStorage != nil {
		in, out := &in.GPUDirectStorage, &out.GPUDirectStorage
		*out = new(GPUDirectStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Manager.DeepCopyInto(&out.Manager)
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]EnvVar, len(*in))
		copy(*out, *in)
	}
	if in.RepoConfig != nil {
		in, out := &in.RepoConfig, &out.RepoConfig
		*out = new(DriverRepoConfigSpec)
		**out = **in
	}
	if in.CertConfig != nil {
		in, out := &in.CertConfig, &out.CertConfig
		*out = new(DriverCertConfigSpec)
		**out = **in
	}
	if in.LicensingConfig != nil {
		in, out := &in.LicensingConfig, &out.LicensingConfig
		*out = new(DriverLicensingConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.VirtualTopologyConfig != nil {
		in, out := &in.VirtualTopologyConfig, &out.VirtualTopologyConfig
		*out = new(VirtualTopologyConfigSpec)
		**out = **in
	}
	if in.KernelModuleConfig != nil {
		in, out := &in.KernelModuleConfig, &out.KernelModuleConfig
		*out = new(KernelModuleConfigSpec)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		*out = new(v1.NodeAffinity)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NVIDIADriverSpec.
func (in *NVIDIADriverSpec) DeepCopy() *NVIDIADriverSpec {
	if in == nil {
		return nil
	}
	out := new(NVIDIADriverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NVIDIADriverStatus) DeepCopyInto(out *NVIDIADriverStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NVIDIADriverStatus.
func (in *NVIDIADriverStatus) DeepCopy() *NVIDIADriverStatus {
	if in == nil {
		return nil
	}
	out := new(NVIDIADriverStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequirements) DeepCopyInto(out *ResourceRequirements) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequirements.
func (in *ResourceRequirements) DeepCopy() *ResourceRequirements {
	if in == nil {
		return nil
	}
	out := new(ResourceRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualTopologyConfigSpec) DeepCopyInto(out *VirtualTopologyConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualTopologyConfigSpec.
func (in *VirtualTopologyConfigSpec) DeepCopy() *VirtualTopologyConfigSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualTopologyConfigSpec)
	in.DeepCopyInto(out)
	return out
}
