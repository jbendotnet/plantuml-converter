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


<!--@startuml
:Hello world;
:This is on defined on
several **lines** end of line;
@enduml-->
![](https://plantuml.signavio.com/png/UDeKny49W30C1T1xflZdZD0ku00k8EGN0o611dHzyVZ6mOZ2MpjioZXlRtYZ4iRbIKEbD1_kAw0QdcnLC0sry7FA5m00__yUtHQh)

## Usage
**Quicklinks**
* [Official Documentation](https://plantuml.com)
* [Extended Diagrams](https://crashedmind.github.io/PlantUMLHitchhikersGuide)
* [Class Diagram](https://plantuml.com/class-diagram)
* [Kubernetes examples](https://github.com/dcasati/kubernetes-PlantUML)
* [AWS examples](https://crashedmind.github.io/PlantUMLHitchhikersGuide/aws/aws.html)
