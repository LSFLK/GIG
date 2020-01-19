package models

import (
	"GIG/commons"
	"GIG/scripts/crawlers/utils"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

type Entity struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title      string        `json:"title" bson:"title"`
	ImageURL   string        `json:"image_url" bson:"image_url"`
	SourceURL  string        `json:"source_url" bson:"source_url"`
	Attributes []Attribute   `json:"attributes" bson:"attributes"`
	Links      []string      `json:"links" bson:"links"`
	Categories []string      `json:"categories" bson:"categories"`
	CreatedAt  time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at" bson:"updated_at"`
	Snippet    string        `json:"snippet" bson:"snippet"`
}

func (e Entity) GetTitle() string {
	return e.Title
}

func (e Entity) SetTitle(titleValue Value) Entity {
	// preprocess title
	title := titleValue.RawValue
	title = strings.TrimSpace(strings.NewReplacer(
		"%", "",
		"/", "-",
		"~", "2",
		"?", "",
	).Replace(title))

	if e.GetTitle() != title {
		e.Title = title
		e.Attributes = e.SetAttribute("titles", titleValue).Attributes
		e.UpdatedAt = time.Now()
	}
	return e
}

/**
Create snippet for the entity
 */
func (e Entity) SetSnippet() Entity {
	if e.Snippet == "" {
		contentAttr, err := e.GetAttribute("")
		snippet := ""
		if err == nil { // if content attribute found
			switch contentAttr.GetValue().Type {
			case "html":
				newsDoc, _ := utils.HTMLStringToDoc(contentAttr.GetValue().RawValue)
				snippet = strings.Replace(newsDoc.Text(), "  ", "", -1)
			default:
				snippet = contentAttr.GetValue().RawValue
			}
		}
		if len(snippet) > 300 {
			snippet = snippet[0:300] + "..."
		}
		e.Snippet = snippet
	}
	return e
}

/**
Compare if a given entity is equal to this entity
 */
func (e Entity) IsEqualTo(otherEntity Entity) bool {
	return e.Title == otherEntity.Title
}

/**
Compare if a given entity source is equal to this entity
 */
func (e Entity) SameSource(otherEntity Entity) bool {
	return e.SourceURL != "" && e.SourceURL == otherEntity.SourceURL
}

/**
Check if the entity has data
 */
func (e Entity) HasContent() bool {
	if len(e.Links) != 0 {
		return true
	}
	if len(e.Categories) != 0 {
		return true
	}
	if len(e.Attributes) != 0 {
		return true
	}
	return false
}

/**
Check if the entity has no title, data
 */
func (e Entity) IsNil() bool {
	if e.Title != "" {
		return false
	}
	return !e.HasContent()
}

/**
Add or update an existing attribute with a new value
 */
func (e Entity) SetAttribute(attributeName string, value Value) Entity {
	//iterate through all attributes
	var attributes []Attribute
	attributeFound := false
	value.updatedAt = time.Now()
	for _, attribute := range e.Attributes {
		if attribute.Name == attributeName { //if attribute name matches an existing attribute
			valueExists := false
			for _, existingValue := range attribute.Values {
				if existingValue.RawValue == value.RawValue && existingValue.Date == value.Date {
					valueExists = true
				}
			}
			fmt.Println(attribute.GetValue().RawValue,value.RawValue)
			if !valueExists && attribute.GetValue().RawValue != value.RawValue { // if the new value doesn't exist already
				attribute = attribute.SetValue(value) // append new value to the attribute
			}

			attributeFound = true
		}
		attributes = append(attributes, attribute)
	}
	if !attributeFound { //else create new attribute and append value

		attribute := Attribute{Name: attributeName}.SetValue(value)
		attributes = append(attributes, attribute)
	}
	e.Attributes = attributes
	return e
}

/**
Get an attribute
 */
func (e Entity) GetAttribute(attributeName string) (Attribute, error) {
	for _, attribute := range e.Attributes {
		if attribute.Name == attributeName {
			return attribute, nil
		}
	}
	return Attribute{}, errors.New("Attribute not found.")
}

/**
Add new link to entity
 */
func (e Entity) AddLink(title string) Entity {
	if commons.StringInSlice(e.Links, title) {
		return e
	}
	if title != "" {
		e.Links = append(e.Links, title)
		e.UpdatedAt = time.Now()
	}
	return e
}

/**
Add new links to entity
 */
func (e Entity) AddLinks(titles []string) Entity {
	parentEntity := e
	for _, title := range titles {
		parentEntity = parentEntity.AddLink(title)
	}
	return parentEntity
}

/**
Add new category to entity
 */
func (e Entity) AddCategory(category string) Entity {
	if commons.StringInSlice(e.Categories, category) {
		return e
	}
	e.Categories = append(e.Categories, category)
	e.UpdatedAt = time.Now()
	return e
}

/**
Add new categories to entity
 */
func (e Entity) AddCategories(categories []string) Entity {
	entity := e
	for _, category := range categories {
		entity = entity.AddCategory(category)
	}
	return entity
}
