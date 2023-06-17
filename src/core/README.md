# Core Server Source Code

## Auth

## Configuration


## Functions
These are the build in functions made available through the customized configuration language.

The customized configuration language is based on [HCL](https://github.com/hashicorp/hcl/tree/main) and edited to only expose the primitive functions built into the application within the function's folder, and specifically handle the linked list data structure of the steps.

### Primitives
- Add
- Sub
- Mult
- Divide

### Strings
- ToUpper
- ToLower
- ToCamelCase
- Concat
- Slice

### DataStructure
- ArrayLoop
- ArraySlice
- 

### Network
- http
- smb
- TCP
- UDP

## API
