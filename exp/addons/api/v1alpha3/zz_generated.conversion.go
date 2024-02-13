//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha3

import (
	unsafe "unsafe"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1alpha3 "sigs.k8s.io/cluster-api/api/v1alpha3"
	apiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
	v1beta1 "sigs.k8s.io/cluster-api/exp/addons/api/v1beta1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ClusterResourceSet)(nil), (*v1beta1.ClusterResourceSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_ClusterResourceSet_To_v1beta1_ClusterResourceSet(a.(*ClusterResourceSet), b.(*v1beta1.ClusterResourceSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ClusterResourceSet)(nil), (*ClusterResourceSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterResourceSet_To_v1alpha3_ClusterResourceSet(a.(*v1beta1.ClusterResourceSet), b.(*ClusterResourceSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterResourceSetBinding)(nil), (*v1beta1.ClusterResourceSetBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_ClusterResourceSetBinding_To_v1beta1_ClusterResourceSetBinding(a.(*ClusterResourceSetBinding), b.(*v1beta1.ClusterResourceSetBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ClusterResourceSetBinding)(nil), (*ClusterResourceSetBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterResourceSetBinding_To_v1alpha3_ClusterResourceSetBinding(a.(*v1beta1.ClusterResourceSetBinding), b.(*ClusterResourceSetBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterResourceSetBindingList)(nil), (*v1beta1.ClusterResourceSetBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_ClusterResourceSetBindingList_To_v1beta1_ClusterResourceSetBindingList(a.(*ClusterResourceSetBindingList), b.(*v1beta1.ClusterResourceSetBindingList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ClusterResourceSetBindingList)(nil), (*ClusterResourceSetBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterResourceSetBindingList_To_v1alpha3_ClusterResourceSetBindingList(a.(*v1beta1.ClusterResourceSetBindingList), b.(*ClusterResourceSetBindingList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterResourceSetBindingSpec)(nil), (*v1beta1.ClusterResourceSetBindingSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_ClusterResourceSetBindingSpec_To_v1beta1_ClusterResourceSetBindingSpec(a.(*ClusterResourceSetBindingSpec), b.(*v1beta1.ClusterResourceSetBindingSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterResourceSetList)(nil), (*v1beta1.ClusterResourceSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_ClusterResourceSetList_To_v1beta1_ClusterResourceSetList(a.(*ClusterResourceSetList), b.(*v1beta1.ClusterResourceSetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ClusterResourceSetList)(nil), (*ClusterResourceSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterResourceSetList_To_v1alpha3_ClusterResourceSetList(a.(*v1beta1.ClusterResourceSetList), b.(*ClusterResourceSetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterResourceSetSpec)(nil), (*v1beta1.ClusterResourceSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_ClusterResourceSetSpec_To_v1beta1_ClusterResourceSetSpec(a.(*ClusterResourceSetSpec), b.(*v1beta1.ClusterResourceSetSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ClusterResourceSetSpec)(nil), (*ClusterResourceSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterResourceSetSpec_To_v1alpha3_ClusterResourceSetSpec(a.(*v1beta1.ClusterResourceSetSpec), b.(*ClusterResourceSetSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterResourceSetStatus)(nil), (*v1beta1.ClusterResourceSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_ClusterResourceSetStatus_To_v1beta1_ClusterResourceSetStatus(a.(*ClusterResourceSetStatus), b.(*v1beta1.ClusterResourceSetStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ClusterResourceSetStatus)(nil), (*ClusterResourceSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterResourceSetStatus_To_v1alpha3_ClusterResourceSetStatus(a.(*v1beta1.ClusterResourceSetStatus), b.(*ClusterResourceSetStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ResourceBinding)(nil), (*v1beta1.ResourceBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_ResourceBinding_To_v1beta1_ResourceBinding(a.(*ResourceBinding), b.(*v1beta1.ResourceBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ResourceBinding)(nil), (*ResourceBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ResourceBinding_To_v1alpha3_ResourceBinding(a.(*v1beta1.ResourceBinding), b.(*ResourceBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ResourceRef)(nil), (*v1beta1.ResourceRef)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_ResourceRef_To_v1beta1_ResourceRef(a.(*ResourceRef), b.(*v1beta1.ResourceRef), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ResourceRef)(nil), (*ResourceRef)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ResourceRef_To_v1alpha3_ResourceRef(a.(*v1beta1.ResourceRef), b.(*ResourceRef), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ResourceSetBinding)(nil), (*v1beta1.ResourceSetBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_ResourceSetBinding_To_v1beta1_ResourceSetBinding(a.(*ResourceSetBinding), b.(*v1beta1.ResourceSetBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ResourceSetBinding)(nil), (*ResourceSetBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ResourceSetBinding_To_v1alpha3_ResourceSetBinding(a.(*v1beta1.ResourceSetBinding), b.(*ResourceSetBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.ClusterResourceSetBindingSpec)(nil), (*ClusterResourceSetBindingSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterResourceSetBindingSpec_To_v1alpha3_ClusterResourceSetBindingSpec(a.(*v1beta1.ClusterResourceSetBindingSpec), b.(*ClusterResourceSetBindingSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha3_ClusterResourceSet_To_v1beta1_ClusterResourceSet(in *ClusterResourceSet, out *v1beta1.ClusterResourceSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha3_ClusterResourceSetSpec_To_v1beta1_ClusterResourceSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha3_ClusterResourceSetStatus_To_v1beta1_ClusterResourceSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_ClusterResourceSet_To_v1beta1_ClusterResourceSet is an autogenerated conversion function.
func Convert_v1alpha3_ClusterResourceSet_To_v1beta1_ClusterResourceSet(in *ClusterResourceSet, out *v1beta1.ClusterResourceSet, s conversion.Scope) error {
	return autoConvert_v1alpha3_ClusterResourceSet_To_v1beta1_ClusterResourceSet(in, out, s)
}

func autoConvert_v1beta1_ClusterResourceSet_To_v1alpha3_ClusterResourceSet(in *v1beta1.ClusterResourceSet, out *ClusterResourceSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_ClusterResourceSetSpec_To_v1alpha3_ClusterResourceSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_ClusterResourceSetStatus_To_v1alpha3_ClusterResourceSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_ClusterResourceSet_To_v1alpha3_ClusterResourceSet is an autogenerated conversion function.
func Convert_v1beta1_ClusterResourceSet_To_v1alpha3_ClusterResourceSet(in *v1beta1.ClusterResourceSet, out *ClusterResourceSet, s conversion.Scope) error {
	return autoConvert_v1beta1_ClusterResourceSet_To_v1alpha3_ClusterResourceSet(in, out, s)
}

func autoConvert_v1alpha3_ClusterResourceSetBinding_To_v1beta1_ClusterResourceSetBinding(in *ClusterResourceSetBinding, out *v1beta1.ClusterResourceSetBinding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha3_ClusterResourceSetBindingSpec_To_v1beta1_ClusterResourceSetBindingSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_ClusterResourceSetBinding_To_v1beta1_ClusterResourceSetBinding is an autogenerated conversion function.
func Convert_v1alpha3_ClusterResourceSetBinding_To_v1beta1_ClusterResourceSetBinding(in *ClusterResourceSetBinding, out *v1beta1.ClusterResourceSetBinding, s conversion.Scope) error {
	return autoConvert_v1alpha3_ClusterResourceSetBinding_To_v1beta1_ClusterResourceSetBinding(in, out, s)
}

func autoConvert_v1beta1_ClusterResourceSetBinding_To_v1alpha3_ClusterResourceSetBinding(in *v1beta1.ClusterResourceSetBinding, out *ClusterResourceSetBinding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_ClusterResourceSetBindingSpec_To_v1alpha3_ClusterResourceSetBindingSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_ClusterResourceSetBinding_To_v1alpha3_ClusterResourceSetBinding is an autogenerated conversion function.
func Convert_v1beta1_ClusterResourceSetBinding_To_v1alpha3_ClusterResourceSetBinding(in *v1beta1.ClusterResourceSetBinding, out *ClusterResourceSetBinding, s conversion.Scope) error {
	return autoConvert_v1beta1_ClusterResourceSetBinding_To_v1alpha3_ClusterResourceSetBinding(in, out, s)
}

func autoConvert_v1alpha3_ClusterResourceSetBindingList_To_v1beta1_ClusterResourceSetBindingList(in *ClusterResourceSetBindingList, out *v1beta1.ClusterResourceSetBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.ClusterResourceSetBinding, len(*in))
		for i := range *in {
			if err := Convert_v1alpha3_ClusterResourceSetBinding_To_v1beta1_ClusterResourceSetBinding(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha3_ClusterResourceSetBindingList_To_v1beta1_ClusterResourceSetBindingList is an autogenerated conversion function.
func Convert_v1alpha3_ClusterResourceSetBindingList_To_v1beta1_ClusterResourceSetBindingList(in *ClusterResourceSetBindingList, out *v1beta1.ClusterResourceSetBindingList, s conversion.Scope) error {
	return autoConvert_v1alpha3_ClusterResourceSetBindingList_To_v1beta1_ClusterResourceSetBindingList(in, out, s)
}

func autoConvert_v1beta1_ClusterResourceSetBindingList_To_v1alpha3_ClusterResourceSetBindingList(in *v1beta1.ClusterResourceSetBindingList, out *ClusterResourceSetBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterResourceSetBinding, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_ClusterResourceSetBinding_To_v1alpha3_ClusterResourceSetBinding(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_ClusterResourceSetBindingList_To_v1alpha3_ClusterResourceSetBindingList is an autogenerated conversion function.
func Convert_v1beta1_ClusterResourceSetBindingList_To_v1alpha3_ClusterResourceSetBindingList(in *v1beta1.ClusterResourceSetBindingList, out *ClusterResourceSetBindingList, s conversion.Scope) error {
	return autoConvert_v1beta1_ClusterResourceSetBindingList_To_v1alpha3_ClusterResourceSetBindingList(in, out, s)
}

func autoConvert_v1alpha3_ClusterResourceSetBindingSpec_To_v1beta1_ClusterResourceSetBindingSpec(in *ClusterResourceSetBindingSpec, out *v1beta1.ClusterResourceSetBindingSpec, s conversion.Scope) error {
	out.Bindings = *(*[]*v1beta1.ResourceSetBinding)(unsafe.Pointer(&in.Bindings))
	return nil
}

// Convert_v1alpha3_ClusterResourceSetBindingSpec_To_v1beta1_ClusterResourceSetBindingSpec is an autogenerated conversion function.
func Convert_v1alpha3_ClusterResourceSetBindingSpec_To_v1beta1_ClusterResourceSetBindingSpec(in *ClusterResourceSetBindingSpec, out *v1beta1.ClusterResourceSetBindingSpec, s conversion.Scope) error {
	return autoConvert_v1alpha3_ClusterResourceSetBindingSpec_To_v1beta1_ClusterResourceSetBindingSpec(in, out, s)
}

func autoConvert_v1beta1_ClusterResourceSetBindingSpec_To_v1alpha3_ClusterResourceSetBindingSpec(in *v1beta1.ClusterResourceSetBindingSpec, out *ClusterResourceSetBindingSpec, s conversion.Scope) error {
	out.Bindings = *(*[]*ResourceSetBinding)(unsafe.Pointer(&in.Bindings))
	// WARNING: in.ClusterName requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha3_ClusterResourceSetList_To_v1beta1_ClusterResourceSetList(in *ClusterResourceSetList, out *v1beta1.ClusterResourceSetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.ClusterResourceSet, len(*in))
		for i := range *in {
			if err := Convert_v1alpha3_ClusterResourceSet_To_v1beta1_ClusterResourceSet(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha3_ClusterResourceSetList_To_v1beta1_ClusterResourceSetList is an autogenerated conversion function.
func Convert_v1alpha3_ClusterResourceSetList_To_v1beta1_ClusterResourceSetList(in *ClusterResourceSetList, out *v1beta1.ClusterResourceSetList, s conversion.Scope) error {
	return autoConvert_v1alpha3_ClusterResourceSetList_To_v1beta1_ClusterResourceSetList(in, out, s)
}

func autoConvert_v1beta1_ClusterResourceSetList_To_v1alpha3_ClusterResourceSetList(in *v1beta1.ClusterResourceSetList, out *ClusterResourceSetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterResourceSet, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_ClusterResourceSet_To_v1alpha3_ClusterResourceSet(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_ClusterResourceSetList_To_v1alpha3_ClusterResourceSetList is an autogenerated conversion function.
func Convert_v1beta1_ClusterResourceSetList_To_v1alpha3_ClusterResourceSetList(in *v1beta1.ClusterResourceSetList, out *ClusterResourceSetList, s conversion.Scope) error {
	return autoConvert_v1beta1_ClusterResourceSetList_To_v1alpha3_ClusterResourceSetList(in, out, s)
}

func autoConvert_v1alpha3_ClusterResourceSetSpec_To_v1beta1_ClusterResourceSetSpec(in *ClusterResourceSetSpec, out *v1beta1.ClusterResourceSetSpec, s conversion.Scope) error {
	out.ClusterSelector = in.ClusterSelector
	out.Resources = *(*[]v1beta1.ResourceRef)(unsafe.Pointer(&in.Resources))
	out.Strategy = in.Strategy
	return nil
}

// Convert_v1alpha3_ClusterResourceSetSpec_To_v1beta1_ClusterResourceSetSpec is an autogenerated conversion function.
func Convert_v1alpha3_ClusterResourceSetSpec_To_v1beta1_ClusterResourceSetSpec(in *ClusterResourceSetSpec, out *v1beta1.ClusterResourceSetSpec, s conversion.Scope) error {
	return autoConvert_v1alpha3_ClusterResourceSetSpec_To_v1beta1_ClusterResourceSetSpec(in, out, s)
}

func autoConvert_v1beta1_ClusterResourceSetSpec_To_v1alpha3_ClusterResourceSetSpec(in *v1beta1.ClusterResourceSetSpec, out *ClusterResourceSetSpec, s conversion.Scope) error {
	out.ClusterSelector = in.ClusterSelector
	out.Resources = *(*[]ResourceRef)(unsafe.Pointer(&in.Resources))
	out.Strategy = in.Strategy
	return nil
}

// Convert_v1beta1_ClusterResourceSetSpec_To_v1alpha3_ClusterResourceSetSpec is an autogenerated conversion function.
func Convert_v1beta1_ClusterResourceSetSpec_To_v1alpha3_ClusterResourceSetSpec(in *v1beta1.ClusterResourceSetSpec, out *ClusterResourceSetSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_ClusterResourceSetSpec_To_v1alpha3_ClusterResourceSetSpec(in, out, s)
}

func autoConvert_v1alpha3_ClusterResourceSetStatus_To_v1beta1_ClusterResourceSetStatus(in *ClusterResourceSetStatus, out *v1beta1.ClusterResourceSetStatus, s conversion.Scope) error {
	out.ObservedGeneration = in.ObservedGeneration
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1beta1.Conditions, len(*in))
		for i := range *in {
			if err := apiv1alpha3.Convert_v1alpha3_Condition_To_v1beta1_Condition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	return nil
}

// Convert_v1alpha3_ClusterResourceSetStatus_To_v1beta1_ClusterResourceSetStatus is an autogenerated conversion function.
func Convert_v1alpha3_ClusterResourceSetStatus_To_v1beta1_ClusterResourceSetStatus(in *ClusterResourceSetStatus, out *v1beta1.ClusterResourceSetStatus, s conversion.Scope) error {
	return autoConvert_v1alpha3_ClusterResourceSetStatus_To_v1beta1_ClusterResourceSetStatus(in, out, s)
}

func autoConvert_v1beta1_ClusterResourceSetStatus_To_v1alpha3_ClusterResourceSetStatus(in *v1beta1.ClusterResourceSetStatus, out *ClusterResourceSetStatus, s conversion.Scope) error {
	out.ObservedGeneration = in.ObservedGeneration
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1alpha3.Conditions, len(*in))
		for i := range *in {
			if err := apiv1alpha3.Convert_v1beta1_Condition_To_v1alpha3_Condition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	return nil
}

// Convert_v1beta1_ClusterResourceSetStatus_To_v1alpha3_ClusterResourceSetStatus is an autogenerated conversion function.
func Convert_v1beta1_ClusterResourceSetStatus_To_v1alpha3_ClusterResourceSetStatus(in *v1beta1.ClusterResourceSetStatus, out *ClusterResourceSetStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_ClusterResourceSetStatus_To_v1alpha3_ClusterResourceSetStatus(in, out, s)
}

func autoConvert_v1alpha3_ResourceBinding_To_v1beta1_ResourceBinding(in *ResourceBinding, out *v1beta1.ResourceBinding, s conversion.Scope) error {
	if err := Convert_v1alpha3_ResourceRef_To_v1beta1_ResourceRef(&in.ResourceRef, &out.ResourceRef, s); err != nil {
		return err
	}
	out.Hash = in.Hash
	out.LastAppliedTime = (*v1.Time)(unsafe.Pointer(in.LastAppliedTime))
	out.Applied = in.Applied
	return nil
}

// Convert_v1alpha3_ResourceBinding_To_v1beta1_ResourceBinding is an autogenerated conversion function.
func Convert_v1alpha3_ResourceBinding_To_v1beta1_ResourceBinding(in *ResourceBinding, out *v1beta1.ResourceBinding, s conversion.Scope) error {
	return autoConvert_v1alpha3_ResourceBinding_To_v1beta1_ResourceBinding(in, out, s)
}

func autoConvert_v1beta1_ResourceBinding_To_v1alpha3_ResourceBinding(in *v1beta1.ResourceBinding, out *ResourceBinding, s conversion.Scope) error {
	if err := Convert_v1beta1_ResourceRef_To_v1alpha3_ResourceRef(&in.ResourceRef, &out.ResourceRef, s); err != nil {
		return err
	}
	out.Hash = in.Hash
	out.LastAppliedTime = (*v1.Time)(unsafe.Pointer(in.LastAppliedTime))
	out.Applied = in.Applied
	return nil
}

// Convert_v1beta1_ResourceBinding_To_v1alpha3_ResourceBinding is an autogenerated conversion function.
func Convert_v1beta1_ResourceBinding_To_v1alpha3_ResourceBinding(in *v1beta1.ResourceBinding, out *ResourceBinding, s conversion.Scope) error {
	return autoConvert_v1beta1_ResourceBinding_To_v1alpha3_ResourceBinding(in, out, s)
}

func autoConvert_v1alpha3_ResourceRef_To_v1beta1_ResourceRef(in *ResourceRef, out *v1beta1.ResourceRef, s conversion.Scope) error {
	out.Name = in.Name
	out.Kind = in.Kind
	return nil
}

// Convert_v1alpha3_ResourceRef_To_v1beta1_ResourceRef is an autogenerated conversion function.
func Convert_v1alpha3_ResourceRef_To_v1beta1_ResourceRef(in *ResourceRef, out *v1beta1.ResourceRef, s conversion.Scope) error {
	return autoConvert_v1alpha3_ResourceRef_To_v1beta1_ResourceRef(in, out, s)
}

func autoConvert_v1beta1_ResourceRef_To_v1alpha3_ResourceRef(in *v1beta1.ResourceRef, out *ResourceRef, s conversion.Scope) error {
	out.Name = in.Name
	out.Kind = in.Kind
	return nil
}

// Convert_v1beta1_ResourceRef_To_v1alpha3_ResourceRef is an autogenerated conversion function.
func Convert_v1beta1_ResourceRef_To_v1alpha3_ResourceRef(in *v1beta1.ResourceRef, out *ResourceRef, s conversion.Scope) error {
	return autoConvert_v1beta1_ResourceRef_To_v1alpha3_ResourceRef(in, out, s)
}

func autoConvert_v1alpha3_ResourceSetBinding_To_v1beta1_ResourceSetBinding(in *ResourceSetBinding, out *v1beta1.ResourceSetBinding, s conversion.Scope) error {
	out.ClusterResourceSetName = in.ClusterResourceSetName
	out.Resources = *(*[]v1beta1.ResourceBinding)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_v1alpha3_ResourceSetBinding_To_v1beta1_ResourceSetBinding is an autogenerated conversion function.
func Convert_v1alpha3_ResourceSetBinding_To_v1beta1_ResourceSetBinding(in *ResourceSetBinding, out *v1beta1.ResourceSetBinding, s conversion.Scope) error {
	return autoConvert_v1alpha3_ResourceSetBinding_To_v1beta1_ResourceSetBinding(in, out, s)
}

func autoConvert_v1beta1_ResourceSetBinding_To_v1alpha3_ResourceSetBinding(in *v1beta1.ResourceSetBinding, out *ResourceSetBinding, s conversion.Scope) error {
	out.ClusterResourceSetName = in.ClusterResourceSetName
	out.Resources = *(*[]ResourceBinding)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_v1beta1_ResourceSetBinding_To_v1alpha3_ResourceSetBinding is an autogenerated conversion function.
func Convert_v1beta1_ResourceSetBinding_To_v1alpha3_ResourceSetBinding(in *v1beta1.ResourceSetBinding, out *ResourceSetBinding, s conversion.Scope) error {
	return autoConvert_v1beta1_ResourceSetBinding_To_v1alpha3_ResourceSetBinding(in, out, s)
}
