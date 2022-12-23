package client

// Client is the fundamental data structure that holds information
// of a client such as it's username, unique hashstring and underlying
// message pipeline.
type Client struct {
  // Username of the client.
  Username    string

  // Unique identification of a client (result of the concatenation
  // of it's custom username and unique hashstring)
  Id          string
  hashstring  string
}

func NewClient(username string) *Client {
  hashstr := newHashString(10)

  return &Client{
    Username: username,
    Id: username + hashstr,
    hashstring: hashstr,
  }
}
