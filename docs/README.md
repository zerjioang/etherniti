# Documentation generation

To create a documentation, first make sure you dont have any active containers:

```bash
docker rm cc_mkdocs_container -f
```

After that, execute for hot reloading execution:

```bash
./container_build.sh && ./container_run.sh && ./docs_serve.sh
```
