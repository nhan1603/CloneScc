package cctv

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/nhan1603/CloneScc/api/internal/appconfig/httpserver"
)

var camURLMapping = map[string]string{
	"cctv_cam1":  os.Getenv("FILE_URL") + "cctv_cam1.mp4",
	"cctv_cam2":  os.Getenv("FILE_URL") + "cctv_cam2.mp4",
	"cctv_cam3":  os.Getenv("FILE_URL") + "cctv_cam3.mp4",
	"cctv_cam4":  os.Getenv("FILE_URL") + "cctv_cam4.mp4",
	"cctv_cam5":  os.Getenv("FILE_URL") + "cctv_cam5.mp4",
	"cctv_cam6":  os.Getenv("FILE_URL") + "cctv_cam1.mp4",
	"cctv_cam30": os.Getenv("FILE_URL") + "cctv_cam2.mp4",
	"cctv_cam31": os.Getenv("FILE_URL") + "cctv_cam3.mp4",
	"cctv_cam32": os.Getenv("FILE_URL") + "cctv_cam4.mp4",
	"cctv_cam33": os.Getenv("FILE_URL") + "cctv_cam5.mp4",
	"cctv_cam34": os.Getenv("FILE_URL") + "cctv_cam1.mp4",
	"cctv_cam35": os.Getenv("FILE_URL") + "cctv_cam2.mp4",
	"cctv_cam36": os.Getenv("FILE_URL") + "cctv_cam3.mp4",
	"cctv_cam40": os.Getenv("FILE_URL") + "cctv_cam4.mp4",
	"cctv_cam41": os.Getenv("FILE_URL") + "cctv_cam5.mp4",
	"cctv_cam42": os.Getenv("FILE_URL") + "cctv_cam1.mp4",
	"cctv_cam43": os.Getenv("FILE_URL") + "cctv_cam2.mp4",
	"cctv_cam44": os.Getenv("FILE_URL") + "cctv_cam3.mp4",
	"cctv_cam45": os.Getenv("FILE_URL") + "cctv_cam4.mp4",
	"cctv_cam46": os.Getenv("FILE_URL") + "cctv_cam5.mp4",
	"cctv_cam47": os.Getenv("FILE_URL") + "cctv_cam1.mp4",
	"cctv_cam48": os.Getenv("FILE_URL") + "cctv_cam2.mp4",
	"cctv_cam49": os.Getenv("FILE_URL") + "cctv_cam3.mp4",
}

// StartFFmpeg starts the FFmpeg process and returns an io.Reader to read the video stream
func StartFFmpeg(camURL string) (io.Reader, *exec.Cmd, error) {
	ffmpegPath, err := exec.LookPath("ffmpeg")
	if err != nil {
		return nil, nil, err
	}

	cmd := exec.Command(ffmpegPath, "-stream_loop", "-1", "-i", camURL, "-movflags", "frag_keyframe+empty_moov", "-c:v", "copy", "-c:a", "copy", "-f", "mp4", "pipe:1")

	// Start FFmpeg and connect its stdout to a pipe
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, nil, err
	}

	return stdoutPipe, cmd, nil
}

// StreamingVideo returns a http.HandlerFunc that streams video from a camera
func (h Handler) StreamingVideo() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		deviceID := strings.TrimSpace(chi.URLParam(r, "deviceCode"))
		if deviceID == "" {
			return webErrorDeviceCodeIsRequired
		}

		camURL, ok := camURLMapping[deviceID]
		if !ok {
			return webErrorNotFoundCam
		}

		log.Printf("Start streaming for device_id: %s\n", deviceID)

		// Get the FFmpeg reader and command for the video stream
		ffmpegReader, ffmpegCmd, err := StartFFmpeg(camURL)
		if err != nil {
			log.Printf("Error starting ffmpeg: %v\n", err)
			return err
		}
		defer ffmpegCmd.Process.Kill()

		// Set the appropriate response headers for video streaming
		w.Header().Set("Content-Type", "video/mp4")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Copy the FFmpeg output to the HTTP response writer
		_, err = io.Copy(w, ffmpegReader)
		if err != nil && err != io.EOF {
			log.Printf("Error copying video stream: %v\n", err)
		}

		log.Printf("Streaming ended for device_id: %s\n", deviceID)
		return nil
	})
}
