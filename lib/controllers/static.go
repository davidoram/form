package controllers

import (
	"mime"
	"net/url"
	"path"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
)

type StaticAssets struct {
	Root string
	Box  *rice.Box
}

func (s *StaticAssets) GetStaticAssets(c echo.Context) error {

	if s.Root == "" {
		s.Root = "." // For security we want to restrict to CWD.
	}

	p, err := url.PathUnescape(c.Param("*"))
	if err != nil {
		return err
	}

	mt := mime.TypeByExtension(path.Ext(p))
	if mt == "" {
		mt = "application/octet-stream"
	}
	header := c.Response().Header()
	header.Set(echo.HeaderContentType, mt)

	// name := filepath.Join(s.Root, path.Clean("/"+p)) // "/"+ for security
	//c.Response().
	b, err := s.Box.Bytes(p)
	if err != nil {
		return err
	}

	_, err = c.Response().Write(b)
	if err != nil {
		return err
	}

	return nil
}
