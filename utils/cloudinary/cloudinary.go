package cloudinary

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v3"
	"github.com/mcctrix/ctrix-social-go-backend/models"
)

var cld *cloudinary.Cloudinary
var ctx context.Context

func UploadImage(cld *cloudinary.Cloudinary, ctx context.Context) {
	// Upload the image.
	// Set the asset's public ID and allow overwriting the asset with new versions
	resp, err := cld.Upload.Upload(ctx, "https://cloudinary-devs.github.io/cld-docs-assets/assets/images/butterfly.jpeg", uploader.UploadParams{
		PublicID:       "quickstart_butterfly",
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true)})
	if err != nil {
		fmt.Println("error")
	}

	// Log the delivery URL
	fmt.Println("****2. Upload an image****\nDelivery URL:", resp.SecureURL)
}

func Credentials() (*cloudinary.Cloudinary, context.Context) {
	if cld != nil {
		return cld, ctx
	}
	c, _ := cloudinary.New()
	c.Config.URL.Secure = false
	cx := context.Background()
	cld = c
	ctx = cx
	return cld, ctx
}

func UploadMediaHandler(c fiber.Ctx) (models.StringArray, error) {
	cld, ctx := Credentials()
	req := c.Request()

	form, err := req.MultipartForm()
	if err != nil {
		contentType := c.Get("Content-Type")
		if !strings.HasPrefix(contentType, "multipart/form-data") {
			return nil, nil
		}
		log.Printf("Error parsing multipart form (Content-Type: %s): %v", contentType, err)
		return nil, fmt.Errorf("failed to parse form data. Ensure correct content-type and form fields, or check file size limits")
	}
	defer func() {
		if err := form.RemoveAll(); err != nil {
			log.Printf("Error removing multipart form files: %v", err)
		}
	}()

	files := form.File["files"]

	if len(files) == 0 {
		return nil, nil
	}

	var imageFiles []*multipart.FileHeader
	var videoFiles []*multipart.FileHeader
	var invalidFiles []string

	for _, file := range files {
		contentType := file.Header.Get("Content-Type")
		if strings.HasPrefix(contentType, "image/") {
			imageFiles = append(imageFiles, file)
		} else if strings.HasPrefix(contentType, "video/") {
			videoFiles = append(videoFiles, file)
		} else {
			invalidFiles = append(invalidFiles, file.Filename)
			log.Printf("Unsupported file type: %s (Content-Type: %s)", file.Filename, contentType)
		}
	}

	if len(imageFiles) > 0 && len(videoFiles) > 0 {
		return nil, fmt.Errorf("cannot upload both images and videos in the same request")
	}
	if len(invalidFiles) > 0 {
		return nil, fmt.Errorf("invalid files: %s", strings.Join(invalidFiles, ", "))
	}

	var urlList models.StringArray

	if len(imageFiles) > 0 {
		if len(imageFiles) > 3 {
			return nil, fmt.Errorf("invalid number of images provided")
		}

		var uploadedImageResults []fiber.Map
		for i, file := range imageFiles {
			fileContent, err := file.Open()
			if err != nil {
				log.Printf("Error opening image file %s: %v", file.Filename, err)
				uploadedImageResults = append(uploadedImageResults, fiber.Map{
					"filename": file.Filename,
					"status":   "failed",
					"error":    "Failed to open image file",
				})
				continue
			}
			defer fileContent.Close()

			publicID := fmt.Sprintf("ctrix-social/posts/images/%s_%d_%s", strings.TrimSuffix(file.Filename, "."), i+1, string(req.Header.UserAgent()))
			uploadResult, err := cld.Upload.Upload(ctx, fileContent, uploader.UploadParams{
				PublicID: publicID,
				Folder:   "ctrix-social/posts/images",
				Tags:     []string{"app", "image", "post"},
			})
			if err != nil {
				log.Printf("Error uploading image %s to Cloudinary: %v", file.Filename, err)
				uploadedImageResults = append(uploadedImageResults, fiber.Map{
					"filename": file.Filename,
					"status":   "failed",
					"error":    fmt.Sprintf("Cloudinary upload failed: %v", err.Error()),
				})
				continue
			}
			urlList = append(urlList, uploadResult.URL)
			uploadedImageResults = append(uploadedImageResults, fiber.Map{
				"filename":  file.Filename,
				"public_id": uploadResult.PublicID,
				"url":       uploadResult.SecureURL,
				"status":    "success",
			})
		}

		return urlList, nil
	}

	if len(videoFiles) > 0 {
		if len(videoFiles) != 1 {
			return nil, fmt.Errorf("invalid number of videos provided")
		}
		videoFile := videoFiles[0]

		fileContent, err := videoFile.Open()
		if err != nil {
			log.Printf("Error opening video file %s: %v", videoFile.Filename, err)
			return nil, fmt.Errorf("failed to open video file")
		}
		defer fileContent.Close()

		publicID := fmt.Sprintf("ctrix-social/posts/videos/%s_%s", strings.TrimSuffix(videoFile.Filename, "."), string(req.Header.UserAgent()))
		uploadResult, err := cld.Upload.Upload(ctx, fileContent, uploader.UploadParams{
			PublicID:     publicID,
			Folder:       "ctrix-social/posts/videos",
			ResourceType: api.Video,
			Tags:         []string{"app", "video", "post"},
		})
		if err != nil {
			log.Printf("Error uploading video %s to Cloudinary: %v", videoFile.Filename, err)
			return nil, fmt.Errorf("failed to upload video to Cloudinary: %v", err.Error())
		}

		urlList = append(urlList, uploadResult.URL)

		return urlList, nil
	}

	// This part should ideally not be reached if previous checks are robust
	return nil, fmt.Errorf("unexpected server error: No files processed or handled")
}
