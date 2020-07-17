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