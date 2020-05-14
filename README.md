![CI](https://github.com/jseris/go-site24x7/workflows/CI/badge.svg) [![codecov](https://codecov.io/gh/jseris/go-site24x7/branch/master/graph/badge.svg)](https://codecov.io/gh/jseris/go-site24x7) ![status-badge](https://goreportcard.com/badge/github.com/jseris/go-site24x7) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jseris/go-site24x7)

# go-site24x7
Go client library for the Site24x7 API.  
At this moment this library is a huge WIP and many site24x7 API-resources are still missing. 
  
## Usage
### Prerequisites
In order to do API-requests to site24x7, the following is required:
* Client ID
* Client Secret
* Refresh Token

Follow the steps described in https://www.site24x7.com/help/api/index.html#authentication to obtain these.  

### Getting a site24x7 client
Construct a Site24x7 client requires two steps. 
First get an authenticator:  
  
    auth, err := oauth.NewAuthenticator("my clientid", "my clientsecret", "my_refreshtoken", "https://accounts.zoho.eu")  
   
For the `tokenDomain`, make sure you use the same domain as when you registered the client application.  
  
Next, get the client itself:  
  
    client := site24x7.NewClient(auth, "my_api_domain")  
  
Optionally, but recommended, you can construct a site24x7 client with a custom HTTP client:  
  
    client := site24x7.NewClient(auth, "my_api_domain", http.Client{Timeout: time.Second * 10})  
  
For a complete example, see the [examples directory](https://github.com/jseris/go-site24x7/tree/master/examples).  
  
### Invoking the API
E.g. to get a list of users:  
  
    users, err := client.Users.List()  

See the [examples directory](https://github.com/jseris/go-site24x7/tree/master/examples) for more examples.  




