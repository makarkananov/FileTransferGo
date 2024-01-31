## FileTransfer

---
<img src="https://cdn.dribbble.com/users/530731/screenshots/14727107/media/60dceabc792bdcbc515573dd7d140d22.jpg?resize=400x300&vertical=center" alt="gopher picture could be here" width="200">

---
### Description
The FileTransfer project is a file transfer application implemented in Go using gRPC for communication. The application allows clients to interact with a server to perform operations such as retrieving a list of files, obtaining file information, and fetching the content of specific files. The project includes implementation of server and client each serving a specific role in the application. 

---
### Structure
**Server Package (server)**: Implements the gRPC server responsible for handling client requests. The server interacts with a `FileUsecase`, which, in turn, communicates with a `FileRepository`. Validation and logging interceptors provided to enhance functionality.

**Client Package (client):** Provides a gRPC client for users to connect to the server. It includes methods for retrieving file lists, file information, and file content. The client also integrates interceptors for enhanced functionality.

**File Repository (repository)**
The project uses a local file repository to manage files but other implementations of `FileRepository` can be provided too. The repository is responsible for reading file lists, obtaining file information, and fetching file content from a specified storage path on the server.

---
### Usage
Basic server initialization provided in **/cmd/server/main.go**, running this will start server with `LocalFileRepository` on **:50051**

Basic client initialization provided in **/cmd/client/main.go** with tiny CLI app using [this](https://github.com/urfave/cli). This can be run with following commands:

* **List Files command**

Usage: `list` \
Aliases: `ls` \
Description: Get the list of files available on the server.

* **File Information command**

Usage: `info [filename]` \
Aliases: `i [filename]` \
Description: Get detailed information about a specific file on the server.

* **Get File Content command**

Usage: `get [filename]` \
Aliases: `g [filename]` \
Description: Retrieve the content of a specific file from the server.

* **Server address option**

Usage: `--server=[address]` \
Aliases: `-s=[address]` \
Description: Specify the address of the gRPC server. If the --server option is not specified, the client will use the default server address (default is localhost:50051).
