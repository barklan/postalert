package client

type Client struct {
	Namespace string
	Addr      string
}

func New(addr string, namespace string) *Client {
	return &Client{
		Addr:      addr,
		Namespace: namespace,
	}
}

func (c *Client) Info() {
}

func (c *Client) Warn() {
}

func (c *Client) Error() {
}

func (c *Client) Crit() {
}
