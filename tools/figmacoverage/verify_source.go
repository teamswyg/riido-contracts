package main

import "errors"

func verifyFigmaSource(source figmaSource) error {
	if blank(source.FileKey) || blank(source.FileName) || blank(source.PageID) || blank(source.PageName) {
		return errors.New("figma file_key, file_name, page_id, and page_name are required")
	}
	if blank(source.InspectedAt) || blank(source.InspectionSource) {
		return errors.New("figma inspected_at and inspection_source are required")
	}
	return nil
}

func verifyPages(pages []page) error {
	if len(pages) == 0 {
		return errors.New("expected_pages are required")
	}
	for _, page := range pages {
		if blank(page.NodeID) || blank(page.Name) || page.ChildCount <= 0 {
			return errors.New("expected_pages require node_id, name, and positive child_count")
		}
	}
	return nil
}
