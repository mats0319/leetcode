package mario

import "strconv"

const urlPrefix = "http://tinyurl.com/"

type Codec struct {
	count int
	data  map[string]string // short url - long url
}

func Constructor() Codec {
	return Codec{
		data: make(map[string]string),
	}
}

// Encodes a URL to a shortened URL.
func (c *Codec) encode(longUrl string) string {
	shortURL := urlPrefix + strconv.Itoa(c.count)
	c.data[shortURL] = longUrl

	return shortURL
}

// Decodes a shortened URL to its original URL.
func (c *Codec) decode(shortUrl string) string {
	return c.data[shortUrl]
}
