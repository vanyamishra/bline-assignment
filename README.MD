Clone the project using the appropriate command. For example, if you want to use HTTPS, the following command would work:

`git clone https://github.com/vanyamishra/bline-assignment`

Once the project repository has been cloned, ensure that you have navigated to the folder gin-poc. If you are in the parent folder, you can run the following command `cd gin-poc`.

This folder will contain the file `nasa_apis.go`, which contains the required URL mappings for calling the NASA APIs. You can run this file using the following command:
`go run nasa_apis.go`

You should see the following log messages on the terminal,
```
[GIN-debug] GET    /apod                     --> main.createExternalGetAPIMapping.func1 (3 handlers)
[GIN-debug] GET    /mars-rover-photos/:earth_date --> main.createExternalGetAPIMappingWithParameter.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
```

Now, you can run the following commands from a different terminal window.

`curl http://localhost:8080/apod`

You should see the required JSON output on the terminal.

The above URL does not require any request parameters. However, the URL to view the mars rover photos requires a request parameter. The URL for viewing the mars rover photos must have the following format,

`curl http://localhost:8080/mars-rover-photos/<earth_date>`

The `earth_date` must be in the `DDDD-MM-YY` format. For example,

`curl http://localhost:8080/mars-rover-photos/2023-11-30`

You should see the required JSON output on the terminal.

Alternatively, you can skip the `curl` command and copy-paste the same commands on a browser window to view the same results.