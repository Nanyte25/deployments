// +build !ignore_autogenerated

/*

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIGatewaySpec) DeepCopyInto(out *APIGatewaySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIGatewaySpec.
func (in *APIGatewaySpec) DeepCopy() *APIGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(APIGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompactSpec) DeepCopyInto(out *CompactSpec) {
	*out = *in
	in.VolumeClaimTemplate.DeepCopyInto(&out.VolumeClaimTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompactSpec.
func (in *CompactSpec) DeepCopy() *CompactSpec {
	if in == nil {
		return nil
	}
	out := new(CompactSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hashring) DeepCopyInto(out *Hashring) {
	*out = *in
	if in.Tenants != nil {
		in, out := &in.Tenants, &out.Tenants
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hashring.
func (in *Hashring) DeepCopy() *Hashring {
	if in == nil {
		return nil
	}
	out := new(Hashring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorageConfig) DeepCopyInto(out *ObjectStorageConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorageConfig.
func (in *ObjectStorageConfig) DeepCopy() *ObjectStorageConfig {
	if in == nil {
		return nil
	}
	out := new(ObjectStorageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Observatorium) DeepCopyInto(out *Observatorium) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Observatorium.
func (in *Observatorium) DeepCopy() *Observatorium {
	if in == nil {
		return nil
	}
	out := new(Observatorium)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Observatorium) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservatoriumList) DeepCopyInto(out *ObservatoriumList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Observatorium, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservatoriumList.
func (in *ObservatoriumList) DeepCopy() *ObservatoriumList {
	if in == nil {
		return nil
	}
	out := new(ObservatoriumList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObservatoriumList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservatoriumSpec) DeepCopyInto(out *ObservatoriumSpec) {
	*out = *in
	out.ObjectStorageConfig = in.ObjectStorageConfig
	if in.Hashrings != nil {
		in, out := &in.Hashrings, &out.Hashrings
		*out = make([]*Hashring, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Hashring)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	in.Compact.DeepCopyInto(&out.Compact)
	out.ReceiveController = in.ReceiveController
	in.Receivers.DeepCopyInto(&out.Receivers)
	in.QueryCache.DeepCopyInto(&out.QueryCache)
	in.Store.DeepCopyInto(&out.Store)
	in.Rule.DeepCopyInto(&out.Rule)
	out.APIGateway = in.APIGateway
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservatoriumSpec.
func (in *ObservatoriumSpec) DeepCopy() *ObservatoriumSpec {
	if in == nil {
		return nil
	}
	out := new(ObservatoriumSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservatoriumStatus) DeepCopyInto(out *ObservatoriumStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservatoriumStatus.
func (in *ObservatoriumStatus) DeepCopy() *ObservatoriumStatus {
	if in == nil {
		return nil
	}
	out := new(ObservatoriumStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryCacheSpec) DeepCopyInto(out *QueryCacheSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryCacheSpec.
func (in *QueryCacheSpec) DeepCopy() *QueryCacheSpec {
	if in == nil {
		return nil
	}
	out := new(QueryCacheSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReceiveControllerSpec) DeepCopyInto(out *ReceiveControllerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReceiveControllerSpec.
func (in *ReceiveControllerSpec) DeepCopy() *ReceiveControllerSpec {
	if in == nil {
		return nil
	}
	out := new(ReceiveControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReceiversSpec) DeepCopyInto(out *ReceiversSpec) {
	*out = *in
	in.VolumeClaimTemplate.DeepCopyInto(&out.VolumeClaimTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReceiversSpec.
func (in *ReceiversSpec) DeepCopy() *ReceiversSpec {
	if in == nil {
		return nil
	}
	out := new(ReceiversSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpec) DeepCopyInto(out *RuleSpec) {
	*out = *in
	in.VolumeClaimTemplate.DeepCopyInto(&out.VolumeClaimTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpec.
func (in *RuleSpec) DeepCopy() *RuleSpec {
	if in == nil {
		return nil
	}
	out := new(RuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreSpec) DeepCopyInto(out *StoreSpec) {
	*out = *in
	in.VolumeClaimTemplate.DeepCopyInto(&out.VolumeClaimTemplate)
	if in.Shards != nil {
		in, out := &in.Shards, &out.Shards
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreSpec.
func (in *StoreSpec) DeepCopy() *StoreSpec {
	if in == nil {
		return nil
	}
	out := new(StoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeClaimTemplate) DeepCopyInto(out *VolumeClaimTemplate) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeClaimTemplate.
func (in *VolumeClaimTemplate) DeepCopy() *VolumeClaimTemplate {
	if in == nil {
		return nil
	}
	out := new(VolumeClaimTemplate)
	in.DeepCopyInto(out)
	return out
}