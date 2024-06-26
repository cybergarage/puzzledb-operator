![doc/img/logo](doc/img/logo.png)


![](https://img.shields.io/badge/status-Work%20In%20Progress-8A2BE2)![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cybergarage/puzzledb-operator) [![Go](https://github.com/cybergarage/puzzledb-operator/actions/workflows/make.yml/badge.svg)](https://github.com/cybergarage/puzzledb-operator/actions/workflows/make.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/cybergarage/puzzledb-operator.svg)](https://pkg.go.dev/github.com/cybergarage/puzzledb-operator) [![Go Report Card](https://img.shields.io/badge/go%20report-A%2B-brightgreen)](https://goreportcard.com/report/github.com/cybergarage/puzzledb-operator) 


PuzzleDB aspires to be a high-performance, distributed, cloud-native, multi-API, multi-model database. PuzzleDB is a multi-data model database capable of handling key-value, relational, and document models. To learn more about PuzzleDB, visit [the GitHub repository](https://github.com/cybergarage/puzzledb-go).

## Description

PuzzleDB Operator is a Kubernetes operator that manages the lifecycle of PuzzleDB instances. 

## Getting Started

### Prerequisites
- go version v1.20.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=cybergarage/puzzledb-operator:tag
```

**NOTE:** This image ought to be published in the personal registry you specified. 
And it is required to have access to pull the image from the working environment. 
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=cybergarage/puzzledb-operator:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin 
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## License

Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

