package handlers

import (
	"io"
	"net/http"
	"simple-webservice/internal/services"
)

func (h *RouterHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(2)

	file, _, err := r.FormFile("routerData")

	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "failed to read content", http.StatusInternalServerError)
		return
	}

	if err := services.ProcessUploadedJSON(content); err != nil {
		http.Error(w, "failed to process json", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("uploaded and merged successfully"))

	// fmt.Fprintf(w, "Uploaded File: %s\n", handler.Filename)
	// fmt.Fprintf(w, "File Size: %d\n", handler.Size)
	// fmt.Fprintf(w, "MIME Header: %v\n", handler.Header)
}
