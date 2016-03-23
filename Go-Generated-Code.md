#Packages

If a .proto file contains a package declaration, the generated code uses the proto's package as its Go package name, converting . characters into _ first. For example, a proto package name of example.high_score results in a Go package name of example_high_score.

You can override the default generated package for a particular .proto using the option go_package in your .proto file. For example, a .proto file containing

```
package example.high_score;
option go_package = "hs";

```

##Messages

Given a simple message declaration:

```
message Foo {}

```

```
type Foo struct {
}

// Reset sets the proto's state to default values.
func (m *Foo) Reset()         { *m = Foo{} }

// String returns a string representation of the proto.
func (m *Foo) String() string { return proto.CompactTextString(m) }

// ProtoMessage acts as a tag to make sure no one accidentally implements the
// proto.Message interface.
func (*Foo) ProtoMessage()    {}
```

##Nested Types

A message can be declared inside another message. For example:

```
message Foo {
  message Bar {
  }
}

```

#Fields

even if the field name in the .proto file uses lower-case with underscores (as it should). The case-conversion works as follows:

- The first letter is capitalized for export.
- For each underscore in the name, the underscore is removed, and the following letter is capitalized.

Thus, the proto field foo_bar_baz becomes FooBarBaz in Go.

Singular Scalar Fields (proto3)

```
int32 foo = 1;

```
