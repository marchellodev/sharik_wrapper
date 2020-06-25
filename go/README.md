# sharik_wrapper

This Go package implements the host-side of the Flutter [sharik_wrapper](https://github.com/marchellodev/sharik_wrapper) plugin.

## Usage

Import as:

```go
import sharik_wrapper "github.com/marchellodev/sharik_wrapper/go"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&sharik_wrapper.SharikWrapperPlugin{}),
```
