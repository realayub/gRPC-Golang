### RPC (Remote Procedure Call): A Classic API Style

RPC (Remote Procedure Call) is a classic and the oldest API style currently in use. It uses procedure calls to request a service from a remote server. gRPC offers a refreshed take on the old RPC design method by making it interoperable, modern, and efficient using such technologies as Protocol Buffers and HTTP/2.

### The Success of gRPC: HTTP/2 and Protocol Buffers

gRPC owes its success to the employment of two techniques: using HTTP/2 instead of HTTP/1.1 and Protocol Buffers as an alternative to XML and JSON. HTTP/2 communication is divided into smaller messages and framed in binary format. Unlike text-based HTTP/1.1, it makes sending and receiving messages compact and efficient. While HTTP/1.1 allows for processing just one request at a time, HTTP/2 supports multiple calls via the same channel. More than that, communication is bidirectional—a single connection can send both requests and responses at the same time.

### Adoption Challenges

The maturity of specific technology can be a big hurdle to its adoption. That’s apparent with gRPC too. Compared to its peer GraphQL that has over 14k questions on StackOverflow, gRPC currently has a little under 4k. Since gRPC heavily relies on HTTP/2, you can’t call a gRPC service from a web browser directly, because no modern browsers can access HTTP/2 frames. So, you need to use a proxy, which has its limitations.

### When to Use gRPC

- Real-time communication services where you deal with streaming calls
- When efficient communication is a goal
- In multi-language environments
- For internal APIs where you don’t have to force technology choices on clients

### Conclusion: gRPC as an Alternative

gRPC is not the evolution of REST, nor is it a better way to build APIs. In a nutshell, gRPC is a way to use RPC’s lightweight structure along with HTTP with a few handy tweaks. It’s just another alternative for you to consider when you start designing a new API. As not many companies actually have the resources to transition to gRPC as Netflix did, you will probably consider it for a new microservices project, such as digital transformation—and benefit greatly.

![Screenshot (53)](https://github.com/user-attachments/assets/1f9ab210-9715-4d3a-b344-2a1bc0b6d7a2)

