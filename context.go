package jin

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

type H map[string]interface{}

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
	Params     map[string]string
	handlers   []HandlerFunc
	index      int
	engine     *Engine
	sameSite   http.SameSite
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, value ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, value...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) XML(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/xml")
	c.Status(code)
	encoder := xml.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(500, err.Error())
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

func (c *Context) Param(key string) string {
	value := c.Params[key]
	return value
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Cookie(name string) (string, error) {
	cookie, err := c.Req.Cookie(name)
	if err != nil {
		return "", err
	}
	value, _ := url.QueryUnescape(cookie.Value)
	return value, nil
}

func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) {
	if path == "" {
		path = "/"
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		SameSite: c.sameSite,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
}

func (c *Context) FormFile(name string) (*multipart.FileHeader, error) {
	if nil == c.Req.MultipartForm {
		if err := c.Req.ParseMultipartForm(c.engine.MaxMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := c.Req.FormFile(name)
	if err != nil {
		return nil, err
	}
	f.Close()
	return fh, err
}

func (c *Context) MultipartForm() (*multipart.Form, error) {
	err := c.Req.ParseMultipartForm(c.engine.MaxMultipartMemory)
	return c.Req.MultipartForm, err
}

func (c *Context) SaveUploadFile(file *multipart.FileHeader, dst string, FileName string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst + FileName)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
