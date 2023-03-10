// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package db_model

import (
	"errors"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type event_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var EventBinding = event_EntityInfo{
	Entity: objectbox.Entity{
		Id: 1,
	},
	Uid: 6113939647246508274,
}

// Event_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Event_ = struct {
	Id           *objectbox.PropertyUint64
	Title        *objectbox.PropertyString
	Description  *objectbox.PropertyString
	LocationLat  *objectbox.PropertyFloat64
	LocationLon  *objectbox.PropertyFloat64
	LocationName *objectbox.PropertyString
	LocationAddr *objectbox.PropertyString
	Start        *objectbox.PropertyInt64
	End          *objectbox.PropertyInt64
	DateCreated  *objectbox.PropertyInt64
	Calendar     *objectbox.RelationToOne
}{
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &EventBinding.Entity,
		},
	},
	Title: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &EventBinding.Entity,
		},
	},
	Description: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &EventBinding.Entity,
		},
	},
	LocationLat: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &EventBinding.Entity,
		},
	},
	LocationLon: &objectbox.PropertyFloat64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &EventBinding.Entity,
		},
	},
	LocationName: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &EventBinding.Entity,
		},
	},
	LocationAddr: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &EventBinding.Entity,
		},
	},
	Start: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     8,
			Entity: &EventBinding.Entity,
		},
	},
	End: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     9,
			Entity: &EventBinding.Entity,
		},
	},
	DateCreated: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     10,
			Entity: &EventBinding.Entity,
		},
	},
	Calendar: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     11,
			Entity: &EventBinding.Entity,
		},
		Target: &CalendarBinding.Entity,
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (event_EntityInfo) GeneratorVersion() int {
	return 6
}

// AddToModel is called by ObjectBox during model build
func (event_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Event", 1, 6113939647246508274)
	model.Property("Id", 6, 1, 7937430401595985583)
	model.PropertyFlags(1)
	model.Property("Title", 9, 2, 2847455842587548770)
	model.Property("Description", 9, 3, 5544060193748436840)
	model.Property("LocationLat", 8, 4, 8561829208578403567)
	model.Property("LocationLon", 8, 5, 5820983911997034002)
	model.Property("LocationName", 9, 6, 9087047129183039073)
	model.Property("LocationAddr", 9, 7, 6244235122388051146)
	model.Property("Start", 10, 8, 8364048063538203126)
	model.PropertyFlags(8)
	model.PropertyIndex(2, 6550041169932512031)
	model.Property("End", 10, 9, 4430804333783655166)
	model.PropertyFlags(8)
	model.PropertyIndex(3, 2695485223480358088)
	model.Property("DateCreated", 10, 10, 931980621875750656)
	model.Property("Calendar", 11, 11, 7850404311720641575)
	model.PropertyFlags(520)
	model.PropertyRelation("Calendar", 1, 5038424931262362087)
	model.EntityLastPropertyId(11, 7850404311720641575)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (event_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*Event).Id, nil
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (event_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*Event).Id = id
	return nil
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (event_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	if rel := object.(*Event).Calendar; rel != nil {
		if rId, err := CalendarBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForCalendar(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (event_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	obj := object.(*Event)
	var propStart int64
	{
		var err error
		propStart, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.Start)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on Event.Start: " + err.Error())
		}
	}

	var propEnd int64
	{
		var err error
		propEnd, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.End)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on Event.End: " + err.Error())
		}
	}

	var propDateCreated int64
	{
		var err error
		propDateCreated, err = objectbox.TimeInt64ConvertToDatabaseValue(obj.DateCreated)
		if err != nil {
			return errors.New("converter objectbox.TimeInt64ConvertToDatabaseValue() failed on Event.DateCreated: " + err.Error())
		}
	}

	var offsetTitle = fbutils.CreateStringOffset(fbb, obj.Title)
	var offsetDescription = fbutils.CreateStringOffset(fbb, obj.Description)
	var offsetLocationName = fbutils.CreateStringOffset(fbb, obj.LocationName)
	var offsetLocationAddr = fbutils.CreateStringOffset(fbb, obj.LocationAddr)

	var rIdCalendar uint64
	if rel := obj.Calendar; rel != nil {
		if rId, err := CalendarBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdCalendar = rId
		}
	}

	// build the FlatBuffers object
	fbb.StartObject(11)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetUOffsetTSlot(fbb, 1, offsetTitle)
	fbutils.SetUOffsetTSlot(fbb, 2, offsetDescription)
	if obj.Calendar != nil {
		fbutils.SetUint64Slot(fbb, 10, rIdCalendar)
	}
	fbutils.SetFloat64Slot(fbb, 3, obj.LocationLat)
	fbutils.SetFloat64Slot(fbb, 4, obj.LocationLon)
	fbutils.SetUOffsetTSlot(fbb, 5, offsetLocationName)
	fbutils.SetUOffsetTSlot(fbb, 6, offsetLocationAddr)
	fbutils.SetInt64Slot(fbb, 7, propStart)
	fbutils.SetInt64Slot(fbb, 8, propEnd)
	fbutils.SetInt64Slot(fbb, 9, propDateCreated)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (event_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'Event' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propId = table.GetUint64Slot(4, 0)

	propStart, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 18))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on Event.Start: " + err.Error())
	}

	propEnd, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 20))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on Event.End: " + err.Error())
	}

	propDateCreated, err := objectbox.TimeInt64ConvertToEntityProperty(fbutils.GetInt64Slot(table, 22))
	if err != nil {
		return nil, errors.New("converter objectbox.TimeInt64ConvertToEntityProperty() failed on Event.DateCreated: " + err.Error())
	}

	var relCalendar *Calendar
	if rId := fbutils.GetUint64PtrSlot(table, 24); rId != nil && *rId > 0 {
		if rObject, err := BoxForCalendar(ob).Get(*rId); err != nil {
			return nil, err
		} else {
			relCalendar = rObject
		}
	}

	return &Event{
		Id:           propId,
		Title:        fbutils.GetStringSlot(table, 6),
		Description:  fbutils.GetStringSlot(table, 8),
		Calendar:     relCalendar,
		LocationLat:  fbutils.GetFloat64Slot(table, 10),
		LocationLon:  fbutils.GetFloat64Slot(table, 12),
		LocationName: fbutils.GetStringSlot(table, 14),
		LocationAddr: fbutils.GetStringSlot(table, 16),
		Start:        propStart,
		End:          propEnd,
		DateCreated:  propDateCreated,
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (event_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*Event, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (event_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]*Event), nil)
	}
	return append(slice.([]*Event), object.(*Event))
}

// Box provides CRUD access to Event objects
type EventBox struct {
	*objectbox.Box
}

// BoxForEvent opens a box of Event objects
func BoxForEvent(ob *objectbox.ObjectBox) *EventBox {
	return &EventBox{
		Box: ob.InternalBox(1),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Event.Id property on the passed object will be assigned the new ID as well.
func (box *EventBox) Put(object *Event) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Event.Id property on the passed object will be assigned the new ID as well.
func (box *EventBox) Insert(object *Event) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *EventBox) Update(object *Event) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *EventBox) PutAsync(object *Event) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Event.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Event.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *EventBox) PutMany(objects []*Event) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *EventBox) Get(id uint64) (*Event, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Event), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is nil
func (box *EventBox) GetMany(ids ...uint64) ([]*Event, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Event), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *EventBox) GetManyExisting(ids ...uint64) ([]*Event, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]*Event), nil
}

// GetAll reads all stored objects
func (box *EventBox) GetAll() ([]*Event, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*Event), nil
}

// Remove deletes a single object
func (box *EventBox) Remove(object *Event) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *EventBox) RemoveMany(objects ...*Event) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = object.Id
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the Event_ struct to create conditions.
// Keep the *EventQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *EventBox) Query(conditions ...objectbox.Condition) *EventQuery {
	return &EventQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Event_ struct to create conditions.
// Keep the *EventQuery if you intend to execute the query multiple times.
func (box *EventBox) QueryOrError(conditions ...objectbox.Condition) (*EventQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &EventQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See EventAsyncBox for more information.
func (box *EventBox) Async() *EventAsyncBox {
	return &EventAsyncBox{AsyncBox: box.Box.Async()}
}

// EventAsyncBox provides asynchronous operations on Event objects.
//
// Asynchronous operations are executed on a separate internal thread for better performance.
//
// There are two main use cases:
//
// 1) "execute & forget:" you gain faster put/remove operations as you don't have to wait for the transaction to finish.
//
// 2) Many small transactions: if your write load is typically a lot of individual puts that happen in parallel,
// this will merge small transactions into bigger ones. This results in a significant gain in overall throughput.
//
// In situations with (extremely) high async load, an async method may be throttled (~1ms) or delayed up to 1 second.
// In the unlikely event that the object could still not be enqueued (full queue), an error will be returned.
//
// Note that async methods do not give you hard durability guarantees like the synchronous Box provides.
// There is a small time window in which the data may not have been committed durably yet.
type EventAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForEvent creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use EventBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForEvent(ob *objectbox.ObjectBox, timeoutMs uint64) *EventAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 1, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 1: %s" + err.Error())
	}
	return &EventAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *EventAsyncBox) Put(object *Event) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *EventAsyncBox) Insert(object *Event) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *EventAsyncBox) Update(object *Event) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *EventAsyncBox) Remove(object *Event) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all Event which Id is either 42 or 47:
//
//	box.Query(Event_.Id.In(42, 47)).Find()
type EventQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *EventQuery) Find() ([]*Event, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*Event), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *EventQuery) Offset(offset uint64) *EventQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *EventQuery) Limit(limit uint64) *EventQuery {
	query.Query.Limit(limit)
	return query
}
