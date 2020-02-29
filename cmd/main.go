package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	yaml2 "sigs.k8s.io/yaml"

	"github.com/crossplane/crossplane/apis/workload/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	kyaml "k8s.io/apimachinery/pkg/util/yaml"
)

func main() {
	dir := "resources"

	resources, err := readResources(dir)
	if err != nil {
		panic(err.Error())
	}
	kapp := &v1alpha1.KubernetesApplication{
		ObjectMeta: v1.ObjectMeta{
			Name: "kapp-name",
		},
		Spec: v1alpha1.KubernetesApplicationSpec{
			ResourceTemplates: resources,
		},
	}
	data, err := yaml2.Marshal(kapp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Print(string(data))
}

func readResources(dir string) ([]v1alpha1.KubernetesApplicationResourceTemplate, error) {
	var result []v1alpha1.KubernetesApplicationResourceTemplate
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		d := kyaml.NewYAMLOrJSONDecoder(file, 4096)
		for {
			obj := &unstructured.Unstructured{}
			if err := d.Decode(obj); err != nil {
				if err == io.EOF {
					// we reached the end of the job output
					break
				}
				return err
			}
			kart := v1alpha1.KubernetesApplicationResourceTemplate{
				ObjectMeta: v1.ObjectMeta{
					Name: fmt.Sprintf("local-%s", obj.GetName()),
				},
				Spec: v1alpha1.KubernetesApplicationResourceSpec{
					Template: obj,
				},
			}
			result = append(result, kart)
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
