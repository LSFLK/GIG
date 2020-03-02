package models

import (
	"GIG-Scripts/crawlers/utils"
	"GIG/commons"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

/**
It is recommended to use get,set functions to access values of the entity.
Directly modify attributes only if you know what you are doing.
 */
type Entity struct {
	Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title      string        `json:"title" bson:"title"`
	ImageURL   string        `json:"image_url" bson:"image_url"`
	Source     string        `json:"source" bson:"source"`
	SourceDate time.Time     `json:"source_date" bson:"source_date"`
	Attributes []Attribute   `json:"attributes" bson:"attributes"`
	Links      []string      `json:"links" bson:"links"`
	Categories []string      `json:"categories" bson:"categories"`
	CreatedAt  time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at" bson:"updated_at"`
	Snippet    string        `json:"snippet" bson:"snippet"`
}

func (e Entity) NewEntity() Entity {
	e.Id = bson.NewObjectId()
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return e
}

func (e Entity) GetId() bson.ObjectId {
	return e.Id
}

func (e Entity) SetTitle(titleValue Value) Entity {
	// preprocess title
	title := titleValue.GetValueString()
	title = strings.TrimSpace(strings.NewReplacer(
		"%", "",
		"/", "-",
		"~", "2",
		"?", "",
	).Replace(title))

	e.Attributes = e.SetAttribute("titles", titleValue).Attributes
	if titleAttribute, err := e.GetAttribute("titles"); err == nil {
		e.Title = titleAttribute.GetValue().GetValueString()
	}

	return e
}

func (e Entity) GetTitle() string {
	return e.Title
}

func (e Entity) SetImageURL(value string) Entity {
	e.ImageURL = value
	e.UpdatedAt = time.Now()

	return e
}

func (e Entity) GetImageURL() string {
	return e.ImageURL
}

func (e Entity) SetSource(value string) Entity {
	e.Source = value
	e.UpdatedAt = time.Now()

	return e
}

func (e Entity) GetSource() string {
	return e.Source
}

func (e Entity) SetSourceDate(value time.Time) Entity {
	e.SourceDate = value
	e.UpdatedAt = time.Now()

	return e
}

func (e Entity) GetSourceDate() time.Time {
	return e.SourceDate
}

/**
Add or update an existing attribute with a new value
 */
func (e Entity) SetAttribute(attributeName string, value Value) Entity {
	//iterate through all attributes
	var attributes []Attribute
	attributeFound := false
	value.UpdatedAt = time.Now()
	for _, attribute := range e.Attributes {
		if attribute.GetName() == attributeName { //if attribute name matches an existing attribute
			valueExists := false
			for _, existingValue := range attribute.GetValues() {
				if existingValue.GetValueString() == value.GetValueString() && existingValue.GetDate() == value.GetDate() {
					valueExists = true
					break
				}
			}
			if !valueExists && attribute.GetValue().GetValueString() != value.GetValueString() { // if the new value doesn't exist already
				attribute = attribute.SetValue(value) // append new value to the attribute
			}

			attributeFound = true
		}
		attributes = append(attributes, attribute)
	}
	if !attributeFound { //else create new attribute and append value

		attribute := Attribute{}.SetName(attributeName).SetValue(value)
		attributes = append(attributes, attribute)
	}
	e.Attributes = attributes
	e.UpdatedAt = time.Now()
	return e
}

/**
Get an attribute
 */
func (e Entity) GetAttribute(attributeName string) (Attribute, error) {
	for _, attribute := range e.Attributes {
		if attribute.GetName() == attributeName {
			return attribute, nil
		}
	}
	return Attribute{}, errors.New("Attribute not found.")
}

func (e Entity) GetAttributes() map[string][]Value {
	result := make(map[string][]Value)
	for _, attribute := range e.Attributes {
		result[attribute.GetName()] = attribute.GetValues()
	}

	return result
}

/**
Add new link to entity
 */
func (e Entity) AddLink(title string) Entity {
	if commons.StringInSlice(e.GetLinks(), title) {
		return e
	}
	if title != "" {
		e.Links = append(e.GetLinks(), title)
		e.UpdatedAt = time.Now()
	}
	return e
}

/**
Add new links to entity
 */
func (e Entity) AddLinks(titles []string) Entity {
	entity := e
	for _, title := range titles {
		entity = e.AddLink(title)
	}
	return entity
}

func (e Entity) GetLinks() []string {
	return e.Links
}

/**
Create snippet for the entity
 */
func (e Entity) SetSnippet() Entity {
	if e.Snippet == "" {
		contentAttr, err := e.GetAttribute("")
		snippet := ""
		if err == nil { // if content attribute found
			switch contentAttr.GetValue().GetType() {
			case "html":
				newsDoc, _ := utils.HTMLStringToDoc(contentAttr.GetValue().GetValueString())
				snippet = strings.Replace(newsDoc.Text(), "  ", "", -1)
			default:
				snippet = contentAttr.GetValue().GetValueString()
			}
		}
		if len(snippet) > 300 {
			snippet = snippet[0:300] + "..."
		}
		e.Snippet = snippet
	}
	return e
}

func (e Entity) GetSnippet() string {
	return e.Snippet
}

/**
Compare if a given entity is equal to this entity on a given date
 */
func (e Entity) IsSameAs(otherEntity Entity, date time.Time) bool {
	//TODO
	return false
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
	if e.GetTitle() != "" {
		return false
	}
	return !e.HasContent()
}

/**
Add new category to entity
 */
func (e Entity) AddCategory(category string) Entity {
	if commons.StringInSlice(e.GetCategories(), category) {
		return e
	}
	e.Categories = append(e.GetCategories(), category)
	e.UpdatedAt = time.Now()
	return e
}

/**
Add new categories to entity
 */
func (e Entity) AddCategories(categories []string) Entity {
	entity := e
	for _, category := range categories {
		entity = e.AddCategory(category)
	}
	return entity
}

func (e Entity) GetCategories() []string {
	return e.Categories
}

func (e Entity) GetCreatedDate() time.Time {
	return e.CreatedAt
}

func (e Entity) GetUpdatedDate() time.Time {
	return e.UpdatedAt
}
