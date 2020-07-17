# plantuml-converter

## Install
Install from releases
```bash
curl -L https://github.com/signavio/plantuml-converter/releases/latest/download/plantuml-converter_$(uname -s)_amd64.gz -o plantuml-converter.gz
gunzip plantuml-converter.gz && chmod +x plantuml-converter && sudo mv plantuml-converter /usr/local/bin/plantuml-converter
plantuml-converter --help
```

### Versioning
Versioning is done automatically done with [Semantics](https://github.com/stevenmatthewt/semantics).
To increase the version any commit in master branch should start with `major:`, `minor:` or `patch:`.
If you squash merge make sure your **git message aka Merge Request Title** starts with one of these.
![](images/release.png)
Once the branch was merged to master, semantics will create an incremented git tag.
Additional the cross-compiled binaries for windows, darwin and linux will be uploaded to releases.

### Class diagram

The following diagram describes the current program classes:

<!-- @startuml
class PlantUmlBlock{
  {field} lineNumber int
  {field} content string
  {field} markdownLink string
  {field} startNumber int
}

class PlantUmlFile{
  {field} filePath string
  {field} fileContent string
  {field} updatedContent string
  {field} blocks []PlantUmlBlock
}

class PlantUml{
  {field} files  []PlantUmlFile
  {field} ScanDirectory string
  {field} Pattern string
}
PlantUml "1" *-- "many" PlantUmlFile : contains
PlantUmlFile "1" *-- "many" PlantUmlBlock : contains
@enduml -->
![](https://plantuml.signavio.com/png/UDfqa35B134GXVlyYY6bS8MjfOgLo89OYSLSCgV3JYQIp277iVzTzk0WkR1jFjxBzoO8rWgJeDf7aaV9OJuxWFE1IU82maflnxId0gpMW93LI0sg5TRl1YGiSyn_-iewZxGQ5ciA5-TwWnSMQWKEB3IX_OnLAtdQiZZ-HZIActo_3gtm-TKj7tLkLIeqcTMruUy1zPaB1SlbD7uveHaLlOB5NMl0ttkus-t09zIJxqu13vTB8sjrtVj6vZAZ3Vq7003__w_-g8i=)

### Infrastructure

Current infrastructure of this application:

<!--@startuml
footer Kubernetes Plant-UML
scale max 1024 width
skinparam linetype polyline
skinparam nodesep 10
skinparam ranksep 10



' Azure
!define AzurePuml https://raw.githubusercontent.com/RicardoNiepel/Azure-PlantUML/release/2-1/dist

!includeurl AzurePuml/AzureCommon.puml
!includeurl AzurePuml/AzureSimplified.puml

!includeurl AzurePuml/Compute/AzureAppService.puml
!includeurl AzurePuml/Containers/AzureContainerRegistry.puml
!includeurl AzurePuml/Networking/AzureLoadBalancer.puml


' Kubernetes
!define KubernetesPuml https://raw.githubusercontent.com/dcasati/kubernetes-PlantUML/master/dist

!includeurl KubernetesPuml/kubernetes_Context.puml
!includeurl KubernetesPuml/kubernetes_Simplified.puml

!includeurl KubernetesPuml/OSS/KubernetesApi.puml
!includeurl KubernetesPuml/OSS/KubernetesIng.puml
!includeurl KubernetesPuml/OSS/KubernetesPod.puml
!includeurl KubernetesPuml/OSS/KubernetesSvc.puml

collections "Client" as clientalias

left to right direction

' Azure Components
AzureContainerRegistry(acr, "Argo CD", "")
AzureLoadBalancer(alb, "\nLoad\nBalancer", "Canada Central")

' Kubernetes Components
Cluster_Boundary(cluster, "Kubernetes Cluster") {
    KubernetesApi(KubernetesApi, "Kubernetes API", "")
    
    Cluster_Boundary(nsFrontEnd, "Front End") {
        KubernetesIng(ingress, "Istio Gateway", "")
    }


    Cluster_Boundary(nsBackEnd, "Back End") {
        KubernetesSvc(svc, "service", "")

        KubernetesPod(KubernetesBE1, "PlantUml Server 1", "")
        KubernetesPod(KubernetesBE2, "PlantUml Server 2", "")
    }
}

Rel(clientalias, alb, "HTTPS", "")
Rel(alb, ingress, "HTTPS", "")
Rel(nsFrontEnd,svc," ")

Rel(svc, KubernetesBE1, " ")
Rel(svc, KubernetesBE2, " ")

Rel(acr, KubernetesApi, "Sync application")
@enduml -->
![](https://plantuml.signavio.com/png/UDgKbC5kuZmGn-z-Ygal1Ibj2lfEtmsotLsqRHSrxQrIvTf3i736aUr0sLNVVHMIGWAKtLuYUsP--S-CnvvP6z31Z_857M503rCZA5mytjumBuL1oCKh3Aw6_y5Agp1dVg6f44xaO3HXM1S8XJNhQjFoaLNeiO31LSle12qQ8sFi74Q_IeViJE5C4zQxQPaRc8TG-F_Zs8dLPQR3l7mfFJffAI25Isdp-5vBuPIzqrYWYJVanIRlnzkRsA51uJ4UNWnYfNrWx4oJDAN2qfcTJ8qbDiyjNHPbRaw5fJelZ9vfL7Ne1x69pOioOCsCYY95jzGIJ_q-iHI49dJ-FQ5cVu-Pzi6jJy5t65RMBJHbDNnZXHeB8qYYQn9bvwt3tJPxP_h7ZYifl0WwNcp1NSDpuGEw8xtkghJOvwfAV0s7jNsCd3o2FUndciOxqwZGVnNg4XFAFabChVeaaIvbKuYqng0Csf87dXYD53W83tAp54OBpvZ1MO1WmUbi7a1fLnFRUmJLx5b22fuT7wIUa2u2FdAPXUGBZu3pFZiOcvumBn7m9wgCJ_Hkhi8JGK89I9224uRtkxFLJY0nPJKIpsDRaX9ktPEr8GBU1ceZxyDl1W3GER5UPzS5HzD9atw5RJu7YkI_EalXcbG4VBE4Qr9RhQxUXBAUfioXzn7mYG_QmZSHS2NMBQ4tnZxI6WkvgAMgrGcbT2bxVYaZuBv-6vh_7mcTMjLgmlXw406lxrnke7fPqC6WbTvfUdY47dQAUsFi7aslDNKHrCFm_U5XcZQnLSZ6kklNlhlL-gfK3bLzbMDJ-7v9q507pc7Zh2MhsTqVYNHD4aHH62r5THbud_q9003__y4ALHi=)


## Usage
**Quicklinks**
* [Official Documentation](https://plantuml.com)
* [Extended Diagrams](https://crashedmind.github.io/PlantUMLHitchhikersGuide)
* [Class Diagram](https://plantuml.com/class-diagram)
* [Kubernetes examples](https://github.com/dcasati/kubernetes-PlantUML)
* [AWS examples](https://crashedmind.github.io/PlantUMLHitchhikersGuide/aws/aws.html)
