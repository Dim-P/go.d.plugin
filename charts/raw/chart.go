package raw

import (
	"errors"
	"fmt"
	"strings"
)

const (
	idxChartTitle = iota
	idxChartUnits
	idxChartFamily
	idxChartContext
	idxChartType
	idxChartOverrideID
)

const (
	Line    = "line"
	Area    = "area"
	Stacked = "stacked"
)

var (
	defaultChartType = Line
)

// NewChart creates new Chart from id, options and optionally dimensions.
func NewChart(id string, options Options, dims ...Dimension) Chart {
	newChart := Chart{ID: id, Options: options}
	for _, dim := range dims {
		newChart.AddDim(dim)
	}
	return newChart
}

type (
	Chart struct {
		ID         string
		Options    Options
		Dimensions Dimensions
		Variables  Variables
	}

	Options    [6]string
	Dimensions []Dimension
	Variables  []Variable
)

func (c *Chart) IsValid() error {
	switch {
	case c.ID == "":
		return errors.New("id not specified")
	case c.Title() == "":
		return errors.New("title not specified")
	case c.Units() == "":
		return errors.New("units not specified")
	case c.Family() == "":
		return errors.New("family not specified")
	default:
		return nil
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// FIELD GETTER

// Title returns 0 element of Options.
func (c *Chart) Title() string {
	return c.Options[idxChartTitle]
}

// Units returns 1 element of Options.
func (c *Chart) Units() string {
	return c.Options[idxChartUnits]
}

// Family returns 2 element of Options converted to lower case.
func (c *Chart) Family() string {
	return strings.ToLower(c.Options[idxChartFamily])
}

// Context returns 3 element of Options.
func (c *Chart) Context() string {
	return c.Options[idxChartContext]
}

// ChartType returns 4 element of Options.
func (c *Chart) ChartType() string {
	if ValidChartType(c.Options[idxChartType]) {
		return c.Options[idxChartType]
	}
	return defaultChartType
}

// OverrideID returns 5 element of Options.
func (c *Chart) OverrideID() string {
	return c.Options[idxChartOverrideID]
}

// ---------------------------------------------------------------------------------------------------------------------

// FIELD SETTER

// SetTitle sets 0 element of Options.
func (c *Chart) SetTitle(t string) *Chart {
	c.Options[idxChartTitle] = t
	return c
}

// SetUnits sets 1 element of Options.
func (c *Chart) SetUnits(u string) *Chart {
	c.Options[idxChartUnits] = u
	return c
}

// SetFamily sets 2 element of Options.
func (c *Chart) SetFamily(f string) *Chart {
	c.Options[idxChartFamily] = f
	return c
}

// SetContext sets 3 element of Options.
func (c *Chart) SetContext(ctx string) *Chart {
	c.Options[idxChartContext] = ctx
	return c
}

// SetChartType sets 4 element of Options.
func (c *Chart) SetChartType(t string) *Chart {
	c.Options[idxChartType] = t
	return c
}

// SetOverrideID sets 5 element of Options.
func (c *Chart) SetOverrideID(id string) *Chart {
	c.Options[idxChartOverrideID] = id
	return c
}

// ---------------------------------------------------------------------------------------------------------------------

// DIM/VAR GETTER

// GetDimByID returns dimension by id.
func (c *Chart) GetDimByID(dimID string) *Dimension {
	if idx, ok := c.indexDim(dimID); ok {
		return &c.Dimensions[idx]
	}
	return nil
}

// GetDimByIndex returns dimension by index.
func (c *Chart) GetDimByIndex(idx int) *Dimension {
	if idx >= 0 && idx < len(c.Dimensions) {
		return &c.Dimensions[idx]
	}
	return nil
}

// GetVarByID returns variable by id.
func (c *Chart) GetVarByID(varID string) *Variable {
	if idx, ok := c.indexVar(varID); ok {
		return &c.Variables[idx]
	}
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// DIM/VAR DELETER

// DeleteDimByID deletes dimension by id.
func (c *Chart) DeleteDimByID(dimID string) error {
	if idx, ok := c.indexDim(dimID); ok {
		c.Dimensions = append(c.Dimensions[:idx], c.Dimensions[idx+1:]...)
		return nil
	}
	return errors.New("nonexistent dimension")
}

// DeleteDimByIndex deletes dimension by index.
func (c *Chart) DeleteDimByIndex(idx int) error {
	if idx >= 0 && idx < len(c.Dimensions) {
		c.Dimensions = append(c.Dimensions[:idx], c.Dimensions[idx+1:]...)
		return nil
	}
	return errors.New("nonexistent dimension")
}

// DeleteVarByID deletes variable by id.
func (c *Chart) DeleteVarByID(varID string) error {
	if idx, ok := c.indexVar(varID); ok {
		c.Variables = append(c.Variables[:idx], c.Variables[idx+1:]...)
		return nil
	}
	return errors.New("nonexistent variable")
}

// ---------------------------------------------------------------------------------------------------------------------

// DIM/VAR ADDER

// AddDim adds valid non duplicate dimension.
func (c *Chart) AddDim(d Dimension) error {
	if err := d.IsValid(); err != nil {
		return fmt.Errorf("chart '%s': invalid dimension (%s)", c.ID, err)
	}
	if _, ok := c.indexDim(d.ID()); ok {
		return fmt.Errorf("chart '%s': duplicate dimension %s", c.ID, d.ID())
	}
	c.Dimensions = append(c.Dimensions, d)
	return nil
}

// AddVar adds valid non duplicate variable.
func (c *Chart) AddVar(v Variable) error {
	if err := v.IsValid(); err != nil {
		return fmt.Errorf("chart '%s': invalid variable (%s)", c.ID, err)
	}
	if _, ok := c.indexVar(v.ID()); ok {
		return fmt.Errorf("chart '%s': duplicate variable %s", c.ID, v.ID())
	}
	c.Variables = append(c.Variables, v)
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

func (c *Chart) indexDim(dimID string) (int, bool) {
	for idx, dim := range c.Dimensions {
		if dim.ID() == dimID {
			return idx, true
		}
	}
	return 0, false
}

func (c *Chart) indexVar(varID string) (int, bool) {
	for idx, v := range c.Variables {
		if v.ID() == varID {
			return idx, true
		}
	}
	return 0, false
}

func (c *Chart) copy() Chart {
	newChart := Chart{ID: c.ID, Options: c.Options}
	for _, d := range c.Dimensions {
		newChart.AddDim(d)
	}
	for _, v := range c.Variables {
		newChart.AddVar(v)
	}
	return newChart
}

// ValidChartType returns whether the chart type is valid.
// Valid chart types: "line", "area", "stacked".
func ValidChartType(t string) bool {
	switch t {
	case Line, Area, Stacked:
		return true
	}
	return false
}
