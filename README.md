# Sample go rest api with tracing

## Tracing Stack
- opentelemetry
- jaeger



### What happens if I don't cancel a Context?

If you fail to cancel the context, the goroutine that WithCancel or WithTimeout created will be retained in memory indefinitely (until the program shuts down), causing a memory leak. If you do this a lot, your memory will balloon significantly. It's best practice to use a defer cancel() immediately after calling WithCancel() or WithTimeout()

