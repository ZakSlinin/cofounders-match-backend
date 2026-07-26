package vision_service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type VisionService struct {
	apiKey   string
	folderID string
}

func NewVisionService(apiKey, folderID string) *VisionService {
	return &VisionService{apiKey: apiKey, folderID: folderID}
}

type ImageModerator interface {
	CheckImage(ctx context.Context, imageURL string) (bool, error)
}

func (s *VisionService) CheckImage(ctx context.Context, imageURL string) (bool, error) {
	body := map[string]any{
		"folderId": s.folderID,
		"analyzeSpecs": []map[string]any{
			{
				"source": map[string]string{
					"url": imageURL,
				},
				"features": []map[string]any{
					{"type": "MODERATION"},
				},
			},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return false, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST",
		"https://vision.api.cloud.yandex.net/vision/v1/batchAnalyze",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return false, err
	}

	req.Header.Set("Authorization", "Api-Key "+s.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	results, ok := result["results"].([]any)
	if !ok {
		return false, fmt.Errorf("unexpected response format")
	}

	for _, r := range results {
		res := r.(map[string]any)
		innerResults, ok := res["results"].([]any)
		if !ok {
			continue
		}
		for _, feature := range innerResults {
			f := feature.(map[string]any)
			moderation, ok := f["moderationAnnotation"].(map[string]any)
			if !ok {
				continue
			}
			categories, ok := moderation["moderationCategories"].([]any)
			if !ok {
				continue
			}
			for _, item := range categories {
				category := item.(map[string]any)
				if confident, ok := category["confident"].(bool); ok && confident {
					return false, nil
				}
			}
		}
	}

	return true, nil
}
