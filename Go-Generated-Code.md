#Packages

If a .proto file contains a package declaration, the generated code uses the proto's package as its Go package name, converting . characters into _ first. For example, a proto package name of example.high_score results in a Go package name of example_high_score.

You can override the default generated package for a particular .proto using the option go_package in your .proto file. For example, a .proto file containing

```
package example.high_score;
option go_package = "hs";

```