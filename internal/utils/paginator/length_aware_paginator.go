package paginator

import (
	"strconv"
)

type TLengthAwarePaginator struct {
	TPaginator
	onEachSide int
}

type TElement struct {
	Show   bool
	IsDots bool
	Items  map[int]string
}

func NewLengthAwarePaginator(items interface{}, total int, perPage int, currentPage int, options map[string]string, manualSlice bool) *TLengthAwarePaginator {
	slicedData := items
	if manualSlice {
		startIndex := (currentPage - 1) * perPage
		endIndex := startIndex + perPage
		if endIndex > total {
			endIndex = total
		}

		slicedData = items.([]map[string]interface{})[startIndex:endIndex]
	}
	paginator := NewPaginator(slicedData, total, perPage, currentPage, options)

	lengthAwarePaginator := &TLengthAwarePaginator{TPaginator: *paginator}

	if options != nil {
		for optionKey, optionValue := range options {
			if optionKey == "onEachSide" {
				i, _ := strconv.Atoi(optionValue)
				lengthAwarePaginator.onEachSide = i
			}
		}
	}

	if lengthAwarePaginator.onEachSide == 0 {
		lengthAwarePaginator.onEachSide = 3
	}

	return lengthAwarePaginator
}

func (p *TLengthAwarePaginator) getSmallSlider() map[string]map[int]string {
	return map[string]map[int]string{
		"first":  p.getUrlRange(1, p.LastPage),
		"slider": nil,
		"last":   nil,
	}
}

func (p *TLengthAwarePaginator) getUrlSlider(onEachSide int) map[string]map[int]string {
	window := onEachSide * 2

	if !p.HasPage() {
		return map[string]map[int]string{"first": nil, "slider": nil, "last": nil}
	}

	if p.CurrentPage <= window {
		return p.getSliderTooCloseToBeginning(window)
	} else if p.CurrentPage > (p.LastPage - window) {
		return p.getSliderTooCloseToEnding(window)
	}

	return p.getFullSlider(onEachSide)
}

func (p *TLengthAwarePaginator) getSliderTooCloseToBeginning(window int) map[string]map[int]string {
	return map[string]map[int]string{
		"first":  p.getUrlRange(1, window+2),
		"slider": nil,
		"last":   p.getFinish(),
	}
}

func (p *TLengthAwarePaginator) getSliderTooCloseToEnding(window int) map[string]map[int]string {
	last := p.getUrlRange(p.LastPage-(window+2), p.LastPage)

	return map[string]map[int]string{"first": p.getStart(), "slider": nil, "last": last}
}

func (p *TLengthAwarePaginator) getFullSlider(onEachSide int) map[string]map[int]string {
	return map[string]map[int]string{
		"first":  p.getStart(),
		"slider": p.getAdjacentUrlRange(onEachSide),
		"last":   p.getFinish(),
	}
}

func (p *TLengthAwarePaginator) getAdjacentUrlRange(onEachSide int) map[int]string {
	return p.getUrlRange(p.CurrentPage-onEachSide, p.CurrentPage+onEachSide)
}

func (p *TLengthAwarePaginator) getStart() map[int]string {
	return p.getUrlRange(1, 2)
}

func (p *TLengthAwarePaginator) getFinish() map[int]string {
	return p.getUrlRange(p.LastPage-1, p.LastPage)
}

func (p *TLengthAwarePaginator) Elements() []TElement {
	window := p.makeWindow()
	sliderDots := TElement{false, false, nil}
	lastDots := TElement{false, false, nil}

	if window["slider"] != nil {
		sliderDots = TElement{true, true, nil}
	}

	if window["last"] != nil {
		lastDots = TElement{true, true, nil}
	}

	return []TElement{
		{true, false, window["first"]},
		sliderDots,
		{true, false, window["slider"]},
		lastDots,
		{true, false, window["last"]},
	}
}

func (p *TLengthAwarePaginator) makeWindow() map[string]map[int]string {
	if p.LastPage < (p.onEachSide*2)+6 {
		return p.getSmallSlider()
	}

	return p.getUrlSlider(p.onEachSide)
}

func (p *TLengthAwarePaginator) GetStringMap() map[string]interface{} {
	return map[string]interface{}{
		"current_page":   p.CurrentPage,
		"data":           p.Data,
		"first_page_url": p.url(1),
		"from":           p.firstItem(),
		"last_page":      p.lastPage(),
		"last_page_url":  p.url(p.lastPage()),
		"next_page_url":  p.NextPageUrl(),
		"path":           p.Path,
		"per_page":       p.PerPage,
		"prev_page_url":  p.PreviousPageUrl(),
		"to":             p.lastItem(),
		"total":          p.Total,
	}
}
