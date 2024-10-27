package thumbnails

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Run() {
	videoPath := filepath.Join(".", "videos", "video.mp4")
	// videoPath := "videos/video.mp4" // Replace with your video path
	outputDir := "thumbnails" // Directory to save thumbnails
	// Create output directory if it doesn't exist
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.Mkdir(outputDir, os.ModePerm)
		if err != nil {
			fmt.Println("Failed to create directory:", err)
			return
		}
	}

	// Generate thumbnails at every 10 seconds
	for i := 0; i < 30; i += 10 {
		thumbnailPath := filepath.Join(outputDir, fmt.Sprintf("thumbnail_%02d.png", i))
		cmd := exec.Command("ffmpeg", "-ss", fmt.Sprintf("%d", i), "-i", videoPath, "-vframes", "1", thumbnailPath)
		err := cmd.Run()
		fmt.Println(os.Getwd())
		if err != nil {
			fmt.Printf("Failed to create thumbnail for %d seconds: %v\n", i, err)
		} else {
			fmt.Printf("Created thumbnail for %d seconds at %s\n", i, thumbnailPath)
		}
	}
}
