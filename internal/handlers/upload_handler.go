package handlers

import (
	"io"
	"net/http"
)

func (h *RouterHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(2)

	file, _, err := r.FormFile("routerData")

	if err != nil {
		http.Error(w, "error retrieving the file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "failed to read content", http.StatusInternalServerError)
		return
	}

	if err := h.Service.ProcessUploadedJSON(content); err != nil {
		http.Error(w, "failed to process json", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("uploaded and merged successfully"))
}
