package moonraker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/sebasptsch/discraker/moonraker/structs"
)

// https://moonraker.readthedocs.io/en/latest/external_api/file_manager/#list-available-files
func (s *Session) ServerFilesList(params structs.ServerFilesListParams) (structs.ServerFilesList, error) {
	return rpc[structs.ServerFilesList](s, "server.files.list", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/file_manager/#list-registered-roots
func (s *Session) ServerFilesRoots() (structs.ServerFilesRoots, error) {
	return rpc[structs.ServerFilesRoots](s, "server.files.roots", nil)
}

// https://moonraker.readthedocs.io/en/latest/external_api/file_manager/#get-gcode-metadata
func (s *Session) ServerFilesMetadata(params structs.ServerFilesMetadataParams) (structs.ServerFilesMetadata, error) {
	return rpc[structs.ServerFilesMetadata](s, "server.files.metadata", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/file_manager/#scan-gcode-metadata
func (s *Session) ServerFilesMetascan(params structs.ServerFilesMetadataParams) (structs.ServerFilesMetadata, error) {
	return rpc[structs.ServerFilesMetadata](s, "server.files.metascan", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/file_manager/#get-gcode-thumbnail-details
func (s *Session) ServerFilesThumbnails(params structs.ServerFilesMetadataParams) (structs.ServerFilesThumbnails, error) {
	return rpc[structs.ServerFilesThumbnails](s, "server.files.thumbnails", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/file_manager/#get-directory-information
func (s *Session) ServerFilesGetDirectory(params structs.ServerFilesGetDirectoryParams) (structs.ServerFilesGetDirectory, error) {
	return rpc[structs.ServerFilesGetDirectory](s, "server.files.get_directory", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/file_manager/#create-directory
func (s *Session) ServerFilesPostDirectory(params structs.ServerFilesPostDirectoryParams) (structs.ServerFilesPostDirectory, error) {
	return rpc[structs.ServerFilesPostDirectory](s, "server.files.post_directory", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/file_manager/#delete-directory
func (s *Session) ServerFilesDeleteDirectory(params structs.ServerFilesDeleteDirectoryParams) (structs.ServerFilesDeleteDirectory, error) {
	return rpc[structs.ServerFilesDeleteDirectory](s, "server.files.delete_directory", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/file_manager/#move-a-file-or-directory
func (s *Session) ServerFilesMove(params structs.ServerFilesMoveParams) (structs.ServerFilesMove, error) {
	return rpc[structs.ServerFilesMove](s, "server.files.move", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/file_manager/#copy-a-file-or-directory
func (s *Session) ServerFilesCopy(params structs.ServerFilesCopyParams) (structs.ServerFilesCopy, error) {
	return rpc[structs.ServerFilesCopy](s, "server.files.copy", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/file_manager/#create-a-zip-archive
func (s *Session) ServerFilesZip(params structs.ServerFilesZipParams) (structs.ServerFilesZip, error) {
	return rpc[structs.ServerFilesZip](s, "server.files.zip", params)
}

// https://moonraker.readthedocs.io/en/latest/external_api/file_manager/#file-download
func (s *Session) ServerFilesDownload(params structs.ServerFilesDownloadParams) (io.ReadCloser, error) {
	requestURL, err := url.JoinPath(s.ConnectionURL.String(), fmt.Sprintf("/server/files/%s/%s", params.Root, params.Filename))

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(*s.Context, "GET", requestURL, nil)

	if err != nil {
		return nil, err
	}

	response, err := s.HTTPClient.Do(request)

	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/file_manager/#file-upload
func (s *Session) ServerFilesUpload(params structs.ServerFilesUploadParams) (structs.ServerFilesUpload, error) {
	requestURL, err := url.JoinPath(s.ConnectionURL.String(), "/server/files/upload")

	if err != nil {
		return structs.ServerFilesUpload{}, err
	}

	buf := &bytes.Buffer{}

	mpw := multipart.NewWriter(buf)

	mpw.WriteField("root", params.Root)

	if params.Checksum != nil {
		mpw.WriteField("checksum", *params.Checksum)
	}

	if params.Path != nil {
		mpw.WriteField("path", *params.Path)
	}

	if params.Print != nil {
		mpw.WriteField("print", fmt.Sprintf("%t", *params.Print))
	}

	part, err := mpw.CreateFormFile("file", "cfvhgbhn.gcode")

	if err != nil {
		return structs.ServerFilesUpload{}, err
	}

	_, err = io.Copy(part, params.File)

	if err != nil {
		return structs.ServerFilesUpload{}, err
	}

	err = mpw.Close()

	if err != nil {
		return structs.ServerFilesUpload{}, err
	}

	request, err := http.NewRequestWithContext(*s.Context, "POST", requestURL, buf)

	if err != nil {
		return structs.ServerFilesUpload{}, err
	}

	uploadResponse, err := s.HTTPClient.Do(request)

	if err != nil {
		return structs.ServerFilesUpload{}, err
	}
	defer uploadResponse.Body.Close()

	var reply structs.ServerFilesUpload

	err = json.NewDecoder(uploadResponse.Body).Decode(&reply)

	if err != nil {
		return structs.ServerFilesUpload{}, err
	}

	return reply, nil
}

// https://moonraker.readthedocs.io/en/latest/external_api/file_manager/#file-delete
func (s *Session) ServerFileDelete(params structs.ServerFilesDeleteParams) (structs.ServerFilesDelete, error) {
	return rpc[structs.ServerFilesDelete](s, "server.files.delete", params)
}
