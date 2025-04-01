package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tbytes404/clipboard/store"
	"github.com/tbytes404/clipboard/view"
)

type BlobsHandler struct {
	store *store.BlobsStore
}

func NewBlobsHandler(store *store.BlobsStore) *BlobsHandler {
	return &BlobsHandler{store}
}

func (h *BlobsHandler) HomePage(c echo.Context) error {
	blobs, err := h.store.FilterBlobs()
	if err != nil {
		return err
	}
	return Render(c, http.StatusOK, view.HomePage(blobs))
}

func (h *BlobsHandler) AddBlob(c echo.Context) error {
	text := c.FormValue("blob")
	if text == "" {
		return echo.ErrBadRequest
	}
	blob, err := h.store.AddBlob(text)
	if err != nil {
		return err
	}
	return Render(c, http.StatusOK, view.Blob(blob))
}

func (h *BlobsHandler) DeleteBlob(c echo.Context) error {
	id, err := strconv.ParseInt(c.FormValue("blob_id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	if err := h.store.DeleteBlob(id); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (h *BlobsHandler) FetchBlobs(c echo.Context) error {
	blobs, err := h.store.FilterBlobs()
	if err != nil {
		return err
	}
	return Render(c, http.StatusOK, view.Blobs(blobs))
}
