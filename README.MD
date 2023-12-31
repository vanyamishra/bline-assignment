# Project Structure

This project contains the following files (in addition to any boilerplate code required by any Go project):
`nasa_apis.go`
`nasa_apis_test.go`

The `nasa_apis.go` contains endpoint mappings for two external endpoints provided by NASA. These include the following:
- https://api.nasa.gov/planetary/apod
- https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos

Functions have been added in this file in such a way that it becomes easier to add any similar mappings in the future. The main() function performs some checks on the values, but mostly supplies objects to these functions. The above endpoints require authentication using an `api_key` that has been added as a global member variable in the file.

The `nasa_apis_test.go` file includes unit tests that validate the functions defined in the `nasa_apis.go` file. Although the unit tests do not provide 100% code coverage at this time, the code has been manually tested to account for multiple scenarios, which include:
- The happy path scenario, where we receive a 200 response code with appropriate content in the response body.
- Submitting a request with incorrect `api_key`, which results in a 403 (forbidden) response code.
- Submitting `earth_date` in an incorrect format, which results in a 400 (bad request) response code.

# Steps to clone and run this project

Clone the project using the appropriate command. For example, if you want to use HTTPS, the following command would work:

`git clone https://github.com/vanyamishra/bline-assignment`

Once the project repository has been cloned, ensure that you have navigated to the folder `bline-assignment`. If you are in the parent folder, you can run the following command to do so,

`cd bline-assignment`.

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

`curl http://localhost:8080/nasa/apod`

If no errors occur, you should see the required JSON output on the terminal. A sample output is as follows:
```
{"result":"{\"copyright\":\"\\nGiulio Cobianchi\\n\",\"date\":\"2023-12-12\",\"explanation\":\"What are these two giant arches across the sky? Perhaps the more familiar one, on the left, is the central band of our Milky Way Galaxy. This grand disk of stars and nebulas here appears to encircle much of the southern sky. Visible below the stellar arch is the rusty-orange planet Mars and the extended Andromeda galaxy. But this night had more!  For a few minutes during this cold arctic night, a second giant arch appeared  encircling part of the northern sky: an aurora. Auroras are much closer than stars as they are composed of glowing air high in Earth's atmosphere.  Visible outside the green auroral arch is the group of stars popularly known as the Big Dipper.  The featured digital composite of 20 images was captured in mid-November 2022 over the Lofoten Islands in Norway.    APOD Year in Review (2023): RJN's Night Sky Network Lecture\",\"hdurl\":\"https://apod.nasa.gov/apod/image/2312/ArcticNight_Cobianchi_2048.jpg\",\"media_type\":\"image\",\"service_version\":\"v1\",\"title\":\"Aurora and Milky Way over Norway\",\"url\":\"https://apod.nasa.gov/apod/image/2312/ArcticNight_Cobianchi_1080.jpg\"}\n"}
```

The above URL does not require any request parameters. However, the URL to view the Mars rover photos requires a request parameter. The URL for viewing the Mars rover photos must have the following format,

`curl http://localhost:8080/mars-rover-photos/<earth_date>`

The `earth_date` must be in the `YYYY-MM-DD` format. For example,

`curl http://localhost:8080/nasa/mars-rover-photos/2022-09-09`

If no errors occur, you should see the required JSON output on the terminal. A sample output is as follows:
```
{"result":"{\"photos\":[]}"}
```

Alternatively, you can skip the `curl` command and copy-paste the same commands on a browser window to view the same results.

Please note that you will see a different response in case of errors.

# Running unit tests

In order to run unit tests, please navigate to your project folder (`bline-assignment` in this case) and run the following command

`go test ./...`