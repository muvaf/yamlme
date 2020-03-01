# KubernetesApplication Generation

This is a POC that I wanted to try out and see how it'd work. Tried to deploy Crossplane helm output via `KubernetesApplication` to a remote cluster.

#### Instructions

Generated helm output via:

```bash
helm2 fetch crossplane-alpha/crossplane --version v0.8.0 --untar
helm2 template --name crossplane --namespace crossplane-system ./crossplane > resources/helmoutput.yaml
```

Run the generator:

```bash
go build cmd/main.go
./main > kapp.yaml
```

```bash
kubectl create -f kapp.yaml
```

#### Results

```bash
$ kubectl get kubernetesapplication
NAME        CLUSTER                                STATUS      DESIRED   SUBMITTED
kapp-name   03c30186-4c27-4c3b-a242-4986997999eb   Submitted   34        34
```

```
$ kubectl get kubernetesapplicationresource
NAME                                                          TEMPLATE-KIND              TEMPLATE-NAME                                           CLUSTER                                STATUS
local-buckets.storage.crossplane.io                           CustomResourceDefinition   buckets.storage.crossplane.io                           03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-clusterstackinstalls.stacks.crossplane.io               CustomResourceDefinition   clusterstackinstalls.stacks.crossplane.io               03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-crossplane                                              Deployment                 crossplane                                              03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-crossplane-admin                                        ClusterRoleBinding         crossplane-admin                                        03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-crossplane-env-admin                                    ClusterRole                crossplane-env-admin                                    03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-crossplane-env-edit                                     ClusterRole                crossplane-env-edit                                     03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-crossplane-env-view                                     ClusterRole                crossplane-env-view                                     03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-crossplane-stack-manager                                Deployment                 crossplane-stack-manager                                03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-crossplane-stack-manager-env-default-admin              ClusterRole                crossplane:stack-manager:env:default:admin              03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-crossplane-stack-manager-env-default-edit               ClusterRole                crossplane:stack-manager:env:default:edit               03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-crossplane-stack-manager-env-default-view               ClusterRole                crossplane:stack-manager:env:default:view               03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-crossplane-stack-manager-ns-default-admin               ClusterRole                crossplane:stack-manager:ns:default:admin               03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-crossplane-stack-manager-ns-default-edit                ClusterRole                crossplane:stack-manager:ns:default:edit                03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-crossplane-stack-manager-ns-default-view                ClusterRole                crossplane:stack-manager:ns:default:view                03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-crossplane-system-aggregate-to-crossplane-admin         ClusterRole                crossplane:system:aggregate-to-crossplane-admin         03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-crossplane-system-ns-persona-cluster-rights             ClusterRole                crossplane:system:ns-persona-cluster-rights             03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-kubernetesapplicationresources.workload.crossplane.io   CustomResourceDefinition   kubernetesapplicationresources.workload.crossplane.io   03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-kubernetesapplications.workload.crossplane.io           CustomResourceDefinition   kubernetesapplications.workload.crossplane.io           03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-kubernetesclusters.compute.crossplane.io                CustomResourceDefinition   kubernetesclusters.compute.crossplane.io                03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-kubernetestargets.workload.crossplane.io                CustomResourceDefinition   kubernetestargets.workload.crossplane.io                03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-machineinstances.compute.crossplane.io                  CustomResourceDefinition   machineinstances.compute.crossplane.io                  03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-mysqlinstances.database.crossplane.io                   CustomResourceDefinition   mysqlinstances.database.crossplane.io                   03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-postgresqlinstances.database.crossplane.io              CustomResourceDefinition   postgresqlinstances.database.crossplane.io              03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-providers.kubernetes.crossplane.io                      CustomResourceDefinition   providers.kubernetes.crossplane.io                      03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-redisclusters.cache.crossplane.io                       CustomResourceDefinition   redisclusters.cache.crossplane.io                       03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-stack-manager                                           ServiceAccount             stack-manager                                           03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-stackdefinitions.stacks.crossplane.io                   CustomResourceDefinition   stackdefinitions.stacks.crossplane.io                   03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-stackinstalls.stacks.crossplane.io                      CustomResourceDefinition   stackinstalls.stacks.crossplane.io                      03c30186-4c27-4c3b-a242-4986997999eb   Submitted
local-stacks.stacks.crossplane.io                             CustomResourceDefinition   stacks.stacks.crossplane.io                             03c30186-4c27-4c3b-a242-4986997999eb   Submitted
```

I wrote this to generate `KubernetesApplication` from a directory of resources. All of them are deployed but Crossplane pod logged some permission issues, which are probably because I didn't configure something.
