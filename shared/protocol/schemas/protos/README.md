
## Compiling for JS

```bash
pbjs -t static-module -w commonjs -o \
    ./src/main/resources/static/protocol.js \
    ./shared/protocol/protos/api.proto
```