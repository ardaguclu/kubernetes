/*
Copyright 2025 The Kubernetes Authors.

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

// OWNER = sig/cli

package kubectl

import (
	"context"
	"os"
	"path/filepath"

	"github.com/onsi/ginkgo/v2"

	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/test/e2e/framework"
	admissionapi "k8s.io/pod-security-admission/api"
)

var _ = SIGDescribe("Kubectl preferences", func() {
	defer ginkgo.GinkgoRecover()
	f := framework.NewDefaultFramework("kubectl-kuberc")
	f.NamespacePodSecurityLevel = admissionapi.LevelBaseline

	tmpDir, err := os.MkdirTemp("", "test-kuberc")
	framework.ExpectNoError(err)
	defer os.Remove(tmpDir) //nolint:errcheck
	kubercFile := filepath.Join(tmpDir, "kuberc")

	var ns string
	var c clientset.Interface
	ginkgo.BeforeEach(func() {
		c = f.ClientSet
		ns = f.Namespace.Name
	})

	ginkgo.Describe("--kuberc", func() {
		ginkgo.It("should be applied on static profiles on ephemeral container", func(ctx context.Context) {

		})
	})
})
