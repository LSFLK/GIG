package models

import (
	"GIG-Scripts/crawlers/utils"
	"GIG/commons"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

type Entity struct {
	id         bson.ObjectId
	title      string
	imageURL   string
	source     string
	attributes []Attribute
	links      []string
	categories []string
	createdAt  time.Time
	updatedAt  time.Time
	snippet    string
}

func (e Entity) NewEntity() Entity {
	e.id = bson.NewObjectId()
	e.createdAt = time.Now()
	e.updatedAt = time.Now()
	return e
}

func (e Entity) GetId() bson.ObjectId {
	return e.id
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

	if e.GetTitle() != title {
		e.title = title
		e.attributes = e.SetAttribute("title", titleValue).attributes

		//TODO: sort all titles by date and set the last as current title
	}
	return e
}

func (e Entity) GetTitle() string {
	return e.title
}

func (e Entity) SetImageURL(value string) Entity {
	e.imageURL = value
	e.updatedAt = time.Now()

	return e
}

func (e Entity) GetImageURL() string {
	return e.imageURL
}

func (e Entity) SetSource(value string) Entity {
	e.source = value
	e.updatedAt = time.Now()

	return e
}

func (e Entity) GetSource() string {
	return e.source
}

/**
Add or update an existing attribute with a new value
 */
func (e Entity) SetAttribute(attributeName string, value Value) Entity {
	//iterate through all attributes
	var attributes []Attribute
	attributeFound := false
	value.updatedAt = time.Now()
	for _, attribute := range e.attributes {
		if attribute.GetName() == attributeName { //if attribute name matches an existing attribute
			valueExists := false
			for _, existingValue := range attribute.GetValues() {
				if existingValue.GetValueString() == value.GetValueString() && existingValue.GetDate() == value.GetDate() {
					valueExists = true
				}
			}
			fmt.Println(attribute.GetValue().GetValueString(), value.GetValueString())
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
	e.attributes = attributes
	e.updatedAt = time.Now()
	return e
}

/**
Get an attribute
 */
func (e Entity) GetAttribute(attributeName string) (Attribute, error) {
	for _, attribute := range e.attributes {
		if attribute.GetName() == attributeName {
			return attribute, nil
		}
	}
	return Attribute{}, errors.New("Attribute not found.")
}

func (e Entity) GetAttributes() map[string]map[string]Value {
	result := make(map[string]map[string]Value)
	for _, attribute := range e.attributes {
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
		e.links = append(e.GetLinks(), title)
		e.updatedAt = time.Now()
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

func (e Entity) GetLinks() []string {
	return e.links
}

/**
Create snippet for the entity
 */
func (e Entity) SetSnippet() Entity {
	if e.snippet == "" {
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
		e.snippet = snippet
	}
	return e
}

func (e Entity) GetSnippet() string {
	return e.snippet
}

/**
Compare if a given entity is equal to this entity
 */
func (e Entity) IsEqualTo(otherEntity Entity) bool {
	return e.GetTitle() == otherEntity.GetTitle()
}

/**
Check if the entity has data
 */
func (e Entity) HasContent() bool {
	if len(e.links) != 0 {
		return true
	}
	if len(e.categories) != 0 {
		return true
	}
	if len(e.attributes) != 0 {
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
	e.categories = append(e.GetCategories(), category)
	e.updatedAt = time.Now()
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

func (e Entity) GetCategories() []string {
	return e.categories
}

func (e Entity) GetCreatedDate() time.Time {
	return e.createdAt
}

func (e Entity) GetUpdatedDate() time.Time {
	return e.updatedAt
}
