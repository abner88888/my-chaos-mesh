// +build !ignore_autogenerated

// Copyright Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthSpec) DeepCopyInto(out *BandwidthSpec) {
	*out = *in
	if in.Peakrate != nil {
		in, out := &in.Peakrate, &out.Peakrate
		*out = new(uint64)
		**out = **in
	}
	if in.Minburst != nil {
		in, out := &in.Minburst, &out.Minburst
		*out = new(uint32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthSpec.
func (in *BandwidthSpec) DeepCopy() *BandwidthSpec {
	if in == nil {
		return nil
	}
	out := new(BandwidthSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUStressor) DeepCopyInto(out *CPUStressor) {
	*out = *in
	out.Stressor = in.Stressor
	if in.Load != nil {
		in, out := &in.Load, &out.Load
		*out = new(int)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUStressor.
func (in *CPUStressor) DeepCopy() *CPUStressor {
	if in == nil {
		return nil
	}
	out := new(CPUStressor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosStatus) DeepCopyInto(out *ChaosStatus) {
	*out = *in
	in.Scheduler.DeepCopyInto(&out.Scheduler)
	in.Experiment.DeepCopyInto(&out.Experiment)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosStatus.
func (in *ChaosStatus) DeepCopy() *ChaosStatus {
	if in == nil {
		return nil
	}
	out := new(ChaosStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CorruptSpec) DeepCopyInto(out *CorruptSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CorruptSpec.
func (in *CorruptSpec) DeepCopy() *CorruptSpec {
	if in == nil {
		return nil
	}
	out := new(CorruptSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DelaySpec) DeepCopyInto(out *DelaySpec) {
	*out = *in
	if in.Reorder != nil {
		in, out := &in.Reorder, &out.Reorder
		*out = new(ReorderSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DelaySpec.
func (in *DelaySpec) DeepCopy() *DelaySpec {
	if in == nil {
		return nil
	}
	out := new(DelaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DuplicateSpec) DeepCopyInto(out *DuplicateSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DuplicateSpec.
func (in *DuplicateSpec) DeepCopy() *DuplicateSpec {
	if in == nil {
		return nil
	}
	out := new(DuplicateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentStatus) DeepCopyInto(out *ExperimentStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = (*in).DeepCopy()
	}
	if in.PodRecords != nil {
		in, out := &in.PodRecords, &out.PodRecords
		*out = make([]PodStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentStatus.
func (in *ExperimentStatus) DeepCopy() *ExperimentStatus {
	if in == nil {
		return nil
	}
	out := new(ExperimentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailKernRequest) DeepCopyInto(out *FailKernRequest) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Callchain != nil {
		in, out := &in.Callchain, &out.Callchain
		*out = make([]Frame, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailKernRequest.
func (in *FailKernRequest) DeepCopy() *FailKernRequest {
	if in == nil {
		return nil
	}
	out := new(FailKernRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Frame) DeepCopyInto(out *Frame) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Frame.
func (in *Frame) DeepCopy() *Frame {
	if in == nil {
		return nil
	}
	out := new(Frame)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IoChaos) DeepCopyInto(out *IoChaos) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IoChaos.
func (in *IoChaos) DeepCopy() *IoChaos {
	if in == nil {
		return nil
	}
	out := new(IoChaos)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IoChaos) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IoChaosList) DeepCopyInto(out *IoChaosList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IoChaos, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IoChaosList.
func (in *IoChaosList) DeepCopy() *IoChaosList {
	if in == nil {
		return nil
	}
	out := new(IoChaosList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IoChaosList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IoChaosSpec) DeepCopyInto(out *IoChaosSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.Scheduler != nil {
		in, out := &in.Scheduler, &out.Scheduler
		*out = new(SchedulerSpec)
		**out = **in
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IoChaosSpec.
func (in *IoChaosSpec) DeepCopy() *IoChaosSpec {
	if in == nil {
		return nil
	}
	out := new(IoChaosSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IoChaosStatus) DeepCopyInto(out *IoChaosStatus) {
	*out = *in
	in.ChaosStatus.DeepCopyInto(&out.ChaosStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IoChaosStatus.
func (in *IoChaosStatus) DeepCopy() *IoChaosStatus {
	if in == nil {
		return nil
	}
	out := new(IoChaosStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelChaos) DeepCopyInto(out *KernelChaos) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelChaos.
func (in *KernelChaos) DeepCopy() *KernelChaos {
	if in == nil {
		return nil
	}
	out := new(KernelChaos)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KernelChaos) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelChaosList) DeepCopyInto(out *KernelChaosList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KernelChaos, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelChaosList.
func (in *KernelChaosList) DeepCopy() *KernelChaosList {
	if in == nil {
		return nil
	}
	out := new(KernelChaosList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KernelChaosList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelChaosSpec) DeepCopyInto(out *KernelChaosSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	in.FailKernRequest.DeepCopyInto(&out.FailKernRequest)
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.Scheduler != nil {
		in, out := &in.Scheduler, &out.Scheduler
		*out = new(SchedulerSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelChaosSpec.
func (in *KernelChaosSpec) DeepCopy() *KernelChaosSpec {
	if in == nil {
		return nil
	}
	out := new(KernelChaosSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelChaosStatus) DeepCopyInto(out *KernelChaosStatus) {
	*out = *in
	in.ChaosStatus.DeepCopyInto(&out.ChaosStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelChaosStatus.
func (in *KernelChaosStatus) DeepCopy() *KernelChaosStatus {
	if in == nil {
		return nil
	}
	out := new(KernelChaosStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LossSpec) DeepCopyInto(out *LossSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LossSpec.
func (in *LossSpec) DeepCopy() *LossSpec {
	if in == nil {
		return nil
	}
	out := new(LossSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryStressor) DeepCopyInto(out *MemoryStressor) {
	*out = *in
	out.Stressor = in.Stressor
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryStressor.
func (in *MemoryStressor) DeepCopy() *MemoryStressor {
	if in == nil {
		return nil
	}
	out := new(MemoryStressor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkChaos) DeepCopyInto(out *NetworkChaos) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkChaos.
func (in *NetworkChaos) DeepCopy() *NetworkChaos {
	if in == nil {
		return nil
	}
	out := new(NetworkChaos)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkChaos) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkChaosList) DeepCopyInto(out *NetworkChaosList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkChaos, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkChaosList.
func (in *NetworkChaosList) DeepCopy() *NetworkChaosList {
	if in == nil {
		return nil
	}
	out := new(NetworkChaosList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkChaosList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkChaosSpec) DeepCopyInto(out *NetworkChaosSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.Scheduler != nil {
		in, out := &in.Scheduler, &out.Scheduler
		*out = new(SchedulerSpec)
		**out = **in
	}
	if in.Delay != nil {
		in, out := &in.Delay, &out.Delay
		*out = new(DelaySpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Loss != nil {
		in, out := &in.Loss, &out.Loss
		*out = new(LossSpec)
		**out = **in
	}
	if in.Duplicate != nil {
		in, out := &in.Duplicate, &out.Duplicate
		*out = new(DuplicateSpec)
		**out = **in
	}
	if in.Corrupt != nil {
		in, out := &in.Corrupt, &out.Corrupt
		*out = new(CorruptSpec)
		**out = **in
	}
	if in.Bandwidth != nil {
		in, out := &in.Bandwidth, &out.Bandwidth
		*out = new(BandwidthSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(Target)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalTargets != nil {
		in, out := &in.ExternalTargets, &out.ExternalTargets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkChaosSpec.
func (in *NetworkChaosSpec) DeepCopy() *NetworkChaosSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkChaosSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkChaosStatus) DeepCopyInto(out *NetworkChaosStatus) {
	*out = *in
	in.ChaosStatus.DeepCopyInto(&out.ChaosStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkChaosStatus.
func (in *NetworkChaosStatus) DeepCopy() *NetworkChaosStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkChaosStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodChaos) DeepCopyInto(out *PodChaos) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodChaos.
func (in *PodChaos) DeepCopy() *PodChaos {
	if in == nil {
		return nil
	}
	out := new(PodChaos)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodChaos) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodChaosList) DeepCopyInto(out *PodChaosList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodChaos, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodChaosList.
func (in *PodChaosList) DeepCopy() *PodChaosList {
	if in == nil {
		return nil
	}
	out := new(PodChaosList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodChaosList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodChaosSpec) DeepCopyInto(out *PodChaosSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.Scheduler != nil {
		in, out := &in.Scheduler, &out.Scheduler
		*out = new(SchedulerSpec)
		**out = **in
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodChaosSpec.
func (in *PodChaosSpec) DeepCopy() *PodChaosSpec {
	if in == nil {
		return nil
	}
	out := new(PodChaosSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodChaosStatus) DeepCopyInto(out *PodChaosStatus) {
	*out = *in
	in.ChaosStatus.DeepCopyInto(&out.ChaosStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodChaosStatus.
func (in *PodChaosStatus) DeepCopy() *PodChaosStatus {
	if in == nil {
		return nil
	}
	out := new(PodChaosStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodStatus) DeepCopyInto(out *PodStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodStatus.
func (in *PodStatus) DeepCopy() *PodStatus {
	if in == nil {
		return nil
	}
	out := new(PodStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReorderSpec) DeepCopyInto(out *ReorderSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReorderSpec.
func (in *ReorderSpec) DeepCopy() *ReorderSpec {
	if in == nil {
		return nil
	}
	out := new(ReorderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleStatus) DeepCopyInto(out *ScheduleStatus) {
	*out = *in
	if in.NextStart != nil {
		in, out := &in.NextStart, &out.NextStart
		*out = (*in).DeepCopy()
	}
	if in.NextRecover != nil {
		in, out := &in.NextRecover, &out.NextRecover
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleStatus.
func (in *ScheduleStatus) DeepCopy() *ScheduleStatus {
	if in == nil {
		return nil
	}
	out := new(ScheduleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerSpec) DeepCopyInto(out *SchedulerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerSpec.
func (in *SchedulerSpec) DeepCopy() *SchedulerSpec {
	if in == nil {
		return nil
	}
	out := new(SchedulerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectorSpec) DeepCopyInto(out *SelectorSpec) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.NodeSelectors != nil {
		in, out := &in.NodeSelectors, &out.NodeSelectors
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FieldSelectors != nil {
		in, out := &in.FieldSelectors, &out.FieldSelectors
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LabelSelectors != nil {
		in, out := &in.LabelSelectors, &out.LabelSelectors
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AnnotationSelectors != nil {
		in, out := &in.AnnotationSelectors, &out.AnnotationSelectors
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodPhaseSelectors != nil {
		in, out := &in.PodPhaseSelectors, &out.PodPhaseSelectors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectorSpec.
func (in *SelectorSpec) DeepCopy() *SelectorSpec {
	if in == nil {
		return nil
	}
	out := new(SelectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StressChaos) DeepCopyInto(out *StressChaos) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StressChaos.
func (in *StressChaos) DeepCopy() *StressChaos {
	if in == nil {
		return nil
	}
	out := new(StressChaos)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StressChaos) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StressChaosList) DeepCopyInto(out *StressChaosList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StressChaos, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StressChaosList.
func (in *StressChaosList) DeepCopy() *StressChaosList {
	if in == nil {
		return nil
	}
	out := new(StressChaosList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StressChaosList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StressChaosSpec) DeepCopyInto(out *StressChaosSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.Stressors != nil {
		in, out := &in.Stressors, &out.Stressors
		*out = new(Stressors)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerName != nil {
		in, out := &in.ContainerName, &out.ContainerName
		*out = new(string)
		**out = **in
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.Scheduler != nil {
		in, out := &in.Scheduler, &out.Scheduler
		*out = new(SchedulerSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StressChaosSpec.
func (in *StressChaosSpec) DeepCopy() *StressChaosSpec {
	if in == nil {
		return nil
	}
	out := new(StressChaosSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StressChaosStatus) DeepCopyInto(out *StressChaosStatus) {
	*out = *in
	in.ChaosStatus.DeepCopyInto(&out.ChaosStatus)
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make(map[string]StressInstance, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StressChaosStatus.
func (in *StressChaosStatus) DeepCopy() *StressChaosStatus {
	if in == nil {
		return nil
	}
	out := new(StressChaosStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StressInstance) DeepCopyInto(out *StressInstance) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StressInstance.
func (in *StressInstance) DeepCopy() *StressInstance {
	if in == nil {
		return nil
	}
	out := new(StressInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stressor) DeepCopyInto(out *Stressor) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stressor.
func (in *Stressor) DeepCopy() *Stressor {
	if in == nil {
		return nil
	}
	out := new(Stressor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stressors) DeepCopyInto(out *Stressors) {
	*out = *in
	if in.MemoryStressor != nil {
		in, out := &in.MemoryStressor, &out.MemoryStressor
		*out = new(MemoryStressor)
		(*in).DeepCopyInto(*out)
	}
	if in.CPUStressor != nil {
		in, out := &in.CPUStressor, &out.CPUStressor
		*out = new(CPUStressor)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stressors.
func (in *Stressors) DeepCopy() *Stressors {
	if in == nil {
		return nil
	}
	out := new(Stressors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Target) DeepCopyInto(out *Target) {
	*out = *in
	in.TargetSelector.DeepCopyInto(&out.TargetSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Target.
func (in *Target) DeepCopy() *Target {
	if in == nil {
		return nil
	}
	out := new(Target)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeChaos) DeepCopyInto(out *TimeChaos) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeChaos.
func (in *TimeChaos) DeepCopy() *TimeChaos {
	if in == nil {
		return nil
	}
	out := new(TimeChaos)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TimeChaos) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeChaosList) DeepCopyInto(out *TimeChaosList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TimeChaos, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeChaosList.
func (in *TimeChaosList) DeepCopy() *TimeChaosList {
	if in == nil {
		return nil
	}
	out := new(TimeChaosList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TimeChaosList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeChaosSpec) DeepCopyInto(out *TimeChaosSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.ClockIds != nil {
		in, out := &in.ClockIds, &out.ClockIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ContainerNames != nil {
		in, out := &in.ContainerNames, &out.ContainerNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.Scheduler != nil {
		in, out := &in.Scheduler, &out.Scheduler
		*out = new(SchedulerSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeChaosSpec.
func (in *TimeChaosSpec) DeepCopy() *TimeChaosSpec {
	if in == nil {
		return nil
	}
	out := new(TimeChaosSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeChaosStatus) DeepCopyInto(out *TimeChaosStatus) {
	*out = *in
	in.ChaosStatus.DeepCopyInto(&out.ChaosStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeChaosStatus.
func (in *TimeChaosStatus) DeepCopy() *TimeChaosStatus {
	if in == nil {
		return nil
	}
	out := new(TimeChaosStatus)
	in.DeepCopyInto(out)
	return out
}

func (h HttpFaultChaos) DeepCopyObject() runtime.Object {
	panic("implement me")
}

func (h HttpFaultChaosList) DeepCopyObject() runtime.Object {
	panic("implement me")
}
