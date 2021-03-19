- [Dell CSI Operator](#dell-csi-operator)
  - [Support](#support)
  - [Supported Platforms](#supported-platforms)
  - [Installation](#installation)
  - [Upgrading Dell CSI Operator](#upgrading-dell-csi-operator)
  - [Install CSI Drivers](#install-csi-drivers)
  - [Uninstall CSI Drivers](#uninstall-csi-drivers)

# Dell CSI Operator
Dell CSI Operator is a Kubernetes native application which helps in installing and managing CSI Drivers provided by Dell EMC for its various storage platforms. 
Dell CSI Operator uses Kubernetes CRDs (Custom Resource Definitions) to define a manifest that describes the deployment specifications for each driver to be deployed. Multiple CSI drivers provided by Dell EMC and multiple instances of each driver can be deployed by the operator by defining a manifest for each deployment.

Dell CSI Operator is built using the [operator framework](https://github.com/operator-framework) and runs custom Kubernetes controllers to manage the driver installations. These controllers listen for any create/update/delete request for the respective CRDs and try to reconcile the request.

Currently, the Dell CSI Operator can be used to deploy the following CSI drivers provided by Dell EMC

* CSI Driver for Dell EMC PowerMax
* CSI Driver for Dell EMC PowerScale
* CSI Driver for Dell EMC Unity
* CSI Driver for Dell EMC PowerFlex (formerly VxFlex OS)
* CSI Driver for Dell EMC PowerStore

Additionally, the Dell CSI Operator can also deploy Storage Classes and Volume Snapshot Classes as part of the driver deployment.
The Dell CSI Operator is itself installed as a Kubernetes deployment.

**NOTE**: You can refer to additional information about the Dell CSI Operator on the new documentation website [here](https://dell.github.io/storage-plugin-docs/docs/installation/operator/)

## Support
The Dell CSI Operator image is available on Dockerhub and is officially supported by Dell EMC.
For any CSI operator and driver issues, questions or feedback, join the [Dell EMC Container community](https://www.dell.com/community/Containers/bd-p/Containers).

## Supported Platforms
Dell CSI Operator has been tested and qualified with 

    * Upstream Kubernetes cluster v1.18, v1.19, v1.20
    * OpenShift Clusters 4.6, 4.7 with RHEL 7.x & RHCOS worker nodes

## Installation
To install Dell CSI Operator please refer the steps given here at [https://dell.github.io/storage-plugin-docs/docs/installation/operator/](https://dell.github.io/storage-plugin-docs/docs/installation/operator/)

## Upgrading Dell CSI Operator
To upgrade the driver to the latest version (across supported Kubernetes/OpenShift versions), please refer [https://dell.github.io/storage-plugin-docs/docs/upgradation/drivers/operator/](https://dell.github.io/storage-plugin-docs/docs/upgradation/drivers/operator/)

## Install CSI Drivers
To install CSI drivers using operator please refer here at [https://dell.github.io/storage-plugin-docs/docs/installation/operator/#install-csi-driver](https://dell.github.io/storage-plugin-docs/docs/installation/operator/#install-csi-driver)

## Uninstall CSI Drivers
To uninstall CSI drivers installed using operator please refer here at [https://dell.github.io/storage-plugin-docs/docs/uninstall/#uninstall-a-dell-csi-driver-installed-via-dell-csi-operator](https://dell.github.io/storage-plugin-docs/docs/uninstall/#uninstall-a-dell-csi-driver-installed-via-dell-csi-operator)