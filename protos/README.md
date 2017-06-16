This folder contains examples of the protocol buffers syntax (version 3).
Note the go_package option, this is the "go path-relative" output folder for generating go stubs.

The google folder provides extra functionality that is crucial to describing RESTful APIs. It depends on google.protobuf.descriptor.proto which is bundled with your protocol buffers compiler. Make sure to include its path on greeter.proto builds with the -I flag.
