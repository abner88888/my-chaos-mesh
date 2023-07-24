// Copyright Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by chaos-builder. DO NOT EDIT.

package v1alpha1

import (
	"reflect"
	"testing"

	"github.com/bxcodec/faker"
	. "github.com/onsi/gomega"
)

func TestAWSChaosIsDeleted(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &AWSChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsDeleted()
}

func TestAWSChaosIsIsPaused(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &AWSChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsPaused()
}

func TestAWSChaosGetDuration(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &AWSChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.Spec.GetDuration()
}

func TestAWSChaosGetStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &AWSChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.GetStatus()
}

func TestAWSChaosGetSpecAndMetaString(t *testing.T) {
	g := NewGomegaWithT(t)
	chaos := &AWSChaos{}
	err := faker.FakeData(chaos)
	g.Expect(err).To(BeNil())
	chaos.GetSpecAndMetaString()
}

func TestAWSChaosListChaos(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &AWSChaosList{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.ListChaos()
}

func TestAzureChaosIsDeleted(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &AzureChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsDeleted()
}

func TestAzureChaosIsIsPaused(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &AzureChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsPaused()
}

func TestAzureChaosGetDuration(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &AzureChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.Spec.GetDuration()
}

func TestAzureChaosGetStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &AzureChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.GetStatus()
}

func TestAzureChaosGetSpecAndMetaString(t *testing.T) {
	g := NewGomegaWithT(t)
	chaos := &AzureChaos{}
	err := faker.FakeData(chaos)
	g.Expect(err).To(BeNil())
	chaos.GetSpecAndMetaString()
}

func TestAzureChaosListChaos(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &AzureChaosList{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.ListChaos()
}

func TestBlockChaosIsDeleted(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &BlockChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsDeleted()
}

func TestBlockChaosIsIsPaused(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &BlockChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsPaused()
}

func TestBlockChaosGetDuration(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &BlockChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.Spec.GetDuration()
}

func TestBlockChaosGetStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &BlockChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.GetStatus()
}

func TestBlockChaosGetSpecAndMetaString(t *testing.T) {
	g := NewGomegaWithT(t)
	chaos := &BlockChaos{}
	err := faker.FakeData(chaos)
	g.Expect(err).To(BeNil())
	chaos.GetSpecAndMetaString()
}

func TestBlockChaosListChaos(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &BlockChaosList{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.ListChaos()
}

func TestDNSChaosIsDeleted(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &DNSChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsDeleted()
}

func TestDNSChaosIsIsPaused(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &DNSChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsPaused()
}

func TestDNSChaosGetDuration(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &DNSChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.Spec.GetDuration()
}

func TestDNSChaosGetStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &DNSChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.GetStatus()
}

func TestDNSChaosGetSpecAndMetaString(t *testing.T) {
	g := NewGomegaWithT(t)
	chaos := &DNSChaos{}
	err := faker.FakeData(chaos)
	g.Expect(err).To(BeNil())
	chaos.GetSpecAndMetaString()
}

func TestDNSChaosListChaos(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &DNSChaosList{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.ListChaos()
}

func TestGCPChaosIsDeleted(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &GCPChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsDeleted()
}

func TestGCPChaosIsIsPaused(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &GCPChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsPaused()
}

func TestGCPChaosGetDuration(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &GCPChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.Spec.GetDuration()
}

func TestGCPChaosGetStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &GCPChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.GetStatus()
}

func TestGCPChaosGetSpecAndMetaString(t *testing.T) {
	g := NewGomegaWithT(t)
	chaos := &GCPChaos{}
	err := faker.FakeData(chaos)
	g.Expect(err).To(BeNil())
	chaos.GetSpecAndMetaString()
}

func TestGCPChaosListChaos(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &GCPChaosList{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.ListChaos()
}

func TestHTTPChaosIsDeleted(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &HTTPChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsDeleted()
}

func TestHTTPChaosIsIsPaused(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &HTTPChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsPaused()
}

func TestHTTPChaosGetDuration(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &HTTPChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.Spec.GetDuration()
}

func TestHTTPChaosGetStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &HTTPChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.GetStatus()
}

func TestHTTPChaosGetSpecAndMetaString(t *testing.T) {
	g := NewGomegaWithT(t)
	chaos := &HTTPChaos{}
	err := faker.FakeData(chaos)
	g.Expect(err).To(BeNil())
	chaos.GetSpecAndMetaString()
}

func TestHTTPChaosListChaos(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &HTTPChaosList{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.ListChaos()
}

func TestIOChaosIsDeleted(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &IOChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsDeleted()
}

func TestIOChaosIsIsPaused(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &IOChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsPaused()
}

func TestIOChaosGetDuration(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &IOChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.Spec.GetDuration()
}

func TestIOChaosGetStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &IOChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.GetStatus()
}

func TestIOChaosGetSpecAndMetaString(t *testing.T) {
	g := NewGomegaWithT(t)
	chaos := &IOChaos{}
	err := faker.FakeData(chaos)
	g.Expect(err).To(BeNil())
	chaos.GetSpecAndMetaString()
}

func TestIOChaosListChaos(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &IOChaosList{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.ListChaos()
}

func TestJVMChaosIsDeleted(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &JVMChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsDeleted()
}

func TestJVMChaosIsIsPaused(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &JVMChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsPaused()
}

func TestJVMChaosGetDuration(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &JVMChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.Spec.GetDuration()
}

func TestJVMChaosGetStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &JVMChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.GetStatus()
}

func TestJVMChaosGetSpecAndMetaString(t *testing.T) {
	g := NewGomegaWithT(t)
	chaos := &JVMChaos{}
	err := faker.FakeData(chaos)
	g.Expect(err).To(BeNil())
	chaos.GetSpecAndMetaString()
}

func TestJVMChaosListChaos(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &JVMChaosList{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.ListChaos()
}

func TestK8SChaosIsDeleted(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &K8SChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsDeleted()
}

func TestK8SChaosIsIsPaused(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &K8SChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsPaused()
}

func TestK8SChaosGetDuration(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &K8SChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.Spec.GetDuration()
}

func TestK8SChaosGetStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &K8SChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.GetStatus()
}

func TestK8SChaosGetSpecAndMetaString(t *testing.T) {
	g := NewGomegaWithT(t)
	chaos := &K8SChaos{}
	err := faker.FakeData(chaos)
	g.Expect(err).To(BeNil())
	chaos.GetSpecAndMetaString()
}

func TestK8SChaosListChaos(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &K8SChaosList{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.ListChaos()
}

func TestKernelChaosIsDeleted(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &KernelChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsDeleted()
}

func TestKernelChaosIsIsPaused(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &KernelChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsPaused()
}

func TestKernelChaosGetDuration(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &KernelChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.Spec.GetDuration()
}

func TestKernelChaosGetStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &KernelChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.GetStatus()
}

func TestKernelChaosGetSpecAndMetaString(t *testing.T) {
	g := NewGomegaWithT(t)
	chaos := &KernelChaos{}
	err := faker.FakeData(chaos)
	g.Expect(err).To(BeNil())
	chaos.GetSpecAndMetaString()
}

func TestKernelChaosListChaos(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &KernelChaosList{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.ListChaos()
}

func TestNetworkChaosIsDeleted(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &NetworkChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsDeleted()
}

func TestNetworkChaosIsIsPaused(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &NetworkChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsPaused()
}

func TestNetworkChaosGetDuration(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &NetworkChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.Spec.GetDuration()
}

func TestNetworkChaosGetStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &NetworkChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.GetStatus()
}

func TestNetworkChaosGetSpecAndMetaString(t *testing.T) {
	g := NewGomegaWithT(t)
	chaos := &NetworkChaos{}
	err := faker.FakeData(chaos)
	g.Expect(err).To(BeNil())
	chaos.GetSpecAndMetaString()
}

func TestNetworkChaosListChaos(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &NetworkChaosList{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.ListChaos()
}

func TestPhysicalMachineChaosIsDeleted(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &PhysicalMachineChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsDeleted()
}

func TestPhysicalMachineChaosIsIsPaused(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &PhysicalMachineChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsPaused()
}

func TestPhysicalMachineChaosGetDuration(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &PhysicalMachineChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.Spec.GetDuration()
}

func TestPhysicalMachineChaosGetStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &PhysicalMachineChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.GetStatus()
}

func TestPhysicalMachineChaosGetSpecAndMetaString(t *testing.T) {
	g := NewGomegaWithT(t)
	chaos := &PhysicalMachineChaos{}
	err := faker.FakeData(chaos)
	g.Expect(err).To(BeNil())
	chaos.GetSpecAndMetaString()
}

func TestPhysicalMachineChaosListChaos(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &PhysicalMachineChaosList{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.ListChaos()
}

func TestPodChaosIsDeleted(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &PodChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsDeleted()
}

func TestPodChaosIsIsPaused(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &PodChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsPaused()
}

func TestPodChaosGetDuration(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &PodChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.Spec.GetDuration()
}

func TestPodChaosGetStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &PodChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.GetStatus()
}

func TestPodChaosGetSpecAndMetaString(t *testing.T) {
	g := NewGomegaWithT(t)
	chaos := &PodChaos{}
	err := faker.FakeData(chaos)
	g.Expect(err).To(BeNil())
	chaos.GetSpecAndMetaString()
}

func TestPodChaosListChaos(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &PodChaosList{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.ListChaos()
}

func TestStressChaosIsDeleted(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &StressChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsDeleted()
}

func TestStressChaosIsIsPaused(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &StressChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsPaused()
}

func TestStressChaosGetDuration(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &StressChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.Spec.GetDuration()
}

func TestStressChaosGetStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &StressChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.GetStatus()
}

func TestStressChaosGetSpecAndMetaString(t *testing.T) {
	g := NewGomegaWithT(t)
	chaos := &StressChaos{}
	err := faker.FakeData(chaos)
	g.Expect(err).To(BeNil())
	chaos.GetSpecAndMetaString()
}

func TestStressChaosListChaos(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &StressChaosList{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.ListChaos()
}

func TestTimeChaosIsDeleted(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &TimeChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsDeleted()
}

func TestTimeChaosIsIsPaused(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &TimeChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.IsPaused()
}

func TestTimeChaosGetDuration(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &TimeChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.Spec.GetDuration()
}

func TestTimeChaosGetStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &TimeChaos{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.GetStatus()
}

func TestTimeChaosGetSpecAndMetaString(t *testing.T) {
	g := NewGomegaWithT(t)
	chaos := &TimeChaos{}
	err := faker.FakeData(chaos)
	g.Expect(err).To(BeNil())
	chaos.GetSpecAndMetaString()
}

func TestTimeChaosListChaos(t *testing.T) {
	g := NewGomegaWithT(t)

	chaos := &TimeChaosList{}
	err := faker.FakeData(chaos)

	g.Expect(err).To(BeNil())

	chaos.ListChaos()
}

func init() {
	faker.AddProvider("ioMethods", func(v reflect.Value) (interface{}, error) {
		return []IoMethod{LookUp}, nil
	})
}
