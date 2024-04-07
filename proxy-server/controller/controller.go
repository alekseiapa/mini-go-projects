package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"time"
)

func ProxyServer() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
		defer cancel()

		queryURL := c.Query("url")
		if queryURL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing query parameter 'url'"})
			return
		}

		targetUrl, err := url.Parse(queryURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error parsing target URL: %s", err)})
			return
		}

		client := &http.Client{}
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetUrl.String(), c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating request: %s", err)})
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error connecting to the destination server: %s", err)})
			return
		}
		defer resp.Body.Close()

		c.Status(resp.StatusCode)
		for k, v := range resp.Header {
			c.Header(k, v[0])
		}

		if _, err := io.Copy(c.Writer, resp.Body); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error connecting to the destination server: %s", err)})
			return
		}
	}
}
