package <%= entityCap %>Response

import (
  "<%= appName %>/helpers/page"
  "<%= appName %>/models"
)

// ListResponseDTO godoc
type ListResponseDTO struct {
  Items      []ListItemResponseDTO `json:"items"`
  TotalItems int64                     `json:"totalItems"`
}

// ListItemResponseDTO godoc
type ListItemResponseDTO struct {
  ID uint `json:"id"`
}

// CreateListResponseDTOFromPage create page from page models.<%= entityCap %>
func CreateListResponseDTOFromPage(page page.Page) *ListResponseDTO {
  result := make([]ListItemResponseDTO, page.GetTotalElements())
  for i, v := range page.GetContent() {
    entity := v.(models.<%= entityCap %>)

    result[i] = ListItemResponseDTO{
      ID: entity.ID,
    }
  }
  return &ListResponseDTO{
    Items:      result,
    TotalItems: page.GetTotalItems(),
  }
}
