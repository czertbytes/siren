package siren

type Class []string

func NewClass(classes ...string) Class {
	return Class(classes)
}

type Rel []string

func NewRel(rels ...string) Rel {
	return Rel(rels)
}

type Properties map[string]interface{}

type Document struct {
	Class      Class      `json:"class,omitempty"`
	Title      string     `json:"title,omitempty"`
	Properties Properties `json:"properties,omitempty"`
	Entities   Entities   `json:"entities,omitempty"`
	Actions    Actions    `json:"actions,omitempty"`
	Links      Links      `json:"links,omitempty"`
}

func NewDocument() *Document {
	return &Document{}
}

func (self *Document) clone() *Document {
	x := *self

	return &x
}

func (self *Document) WithClass(class Class) *Document {
	self = self.clone()
	self.Class = class

	return self
}

func (self *Document) WithTitle(title string) *Document {
	self = self.clone()
	self.Title = title

	return self
}

func (self *Document) WithProperties(properties Properties) *Document {
	self = self.clone()
	self.Properties = properties

	return self
}

func (self *Document) WithEntities(entities Entities) *Document {
	self = self.clone()
	self.Entities = entities

	return self
}

func (self *Document) WithActions(actions Actions) *Document {
	self = self.clone()
	self.Actions = actions

	return self
}

func (self *Document) WithLinks(links Links) *Document {
	self = self.clone()
	self.Links = links

	return self
}

type Entity struct {
	Class      Class      `json:"class,omitempty"`
	Rel        Rel        `json:"rel"`
	Href       string     `json:"href,omitempty"`
	Properties Properties `json:"properties,omitempty"`
	Links      Links      `json:"links,omitempty"`
	Type       string     `json:"type,omitempty"`
}

type Entities []*Entity

func NewEntity(rel Rel) *Entity {
	return &Entity{
		Rel: rel,
	}
}

func NewLinkEntity(rel Rel, href string) *Entity {
	entity := NewEntity(rel)
	entity.Href = href

	return entity
}

func (self *Entity) clone() *Entity {
	x := *self

	return &x
}

func (self *Entity) WithClass(class Class) *Entity {
	self = self.clone()
	self.Class = class

	return self
}

func (self *Entity) WithProperties(properties Properties) *Entity {
	self = self.clone()
	self.Properties = properties

	return self
}

func (self *Entity) WithLinks(links Links) *Entity {
	self = self.clone()
	self.Links = links

	return self
}

func (self *Entity) WithType(typ string) *Entity {
	self = self.clone()
	self.Type = typ

	return self
}

type Action struct {
	Name        string `json:"name"`
	Class       Class  `json:"class,omitempty"`
	Title       string `json:"title,omitempty"`
	Method      string `json:"method,omitempty"`
	Href        string `json:"href"`
	ContentType string `json:"type,omitempty"`
	Fields      Fields `json:"fields,omitempty"`
}

type Actions []*Action

func NewAction(name, href string) *Action {
	return &Action{
		Name: name,
		Href: href,
	}
}

func (self *Action) WithTitle(title string) *Action {
	self = self.clone()
	self.Title = title

	return self
}

func (self *Action) WithMethod(method string) *Action {
	self = self.clone()
	self.Method = method

	return self
}

func (self *Action) WithContentType(contentType string) *Action {
	self = self.clone()
	self.ContentType = contentType

	return self
}

func (self *Action) WithFields(fields Fields) *Action {
	self = self.clone()
	self.Fields = fields

	return self
}

func (self *Action) clone() *Action {
	x := *self

	return &x
}

type Field struct {
	Name  string      `json:"name"`
	Type  string      `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Title string      `json:"title,omitempty"`
}

type Fields []*Field

func NewField(name, typ string) *Field {
	return &Field{
		Name: name,
		Type: typ,
	}
}

func NewHiddenField(name string) *Field {
	return NewField(name, "hidden")
}

func NewTextField(name string) *Field {
	return NewField(name, "text")
}

func NewNumberField(name string) *Field {
	return NewField(name, "number")
}

func (self *Field) WithValue(value interface{}) *Field {
	self = self.clone()
	self.Value = value

	return self
}

func (self *Field) WithTitle(title string) *Field {
	self = self.clone()
	self.Title = title

	return self
}

func (self *Field) clone() *Field {
	x := *self

	return &x
}

type Link struct {
	Rel   Rel    `json:"rel"`
	Href  string `json:"href"`
	Title string `json:"title,omitempty"`
	Type  string `json:"type,omitempty"`
}

type Links []*Link

func NewLink(rel Rel, href string) *Link {
	return &Link{
		Rel:  rel,
		Href: href,
	}
}

func (self *Link) WithTitle(title string) *Link {
	self = self.clone()
	self.Title = title

	return self
}

func (self *Link) WithType(typ string) *Link {
	self = self.clone()
	self.Type = typ

	return self
}

func (self *Link) clone() *Link {
	x := *self

	return &x
}
