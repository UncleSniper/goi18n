package goi18n

type MemoryLocalizer[MessageKeyT comparable] struct {
	MessageRoot MemoryNode[MessageKeyT, string]
	NumberFormatRoot MemoryNode[MessageKeyT, NumberFormats]
}

type MemoryNode[MessageKeyT comparable, ValueT any] interface {
	MessageKeyTypeName() string
	Lookup(*Locale, MessageKeyT) (ValueT, bool)
}

type MemoryNodeBuilder[MessageKeyT comparable, PropertyKeyT comparable, ValueT any] interface {
	MemoryNode[MessageKeyT, ValueT]
	Put(PropertyKeyT, MemoryNode[MessageKeyT, ValueT]) MemoryNodeBuilder[MessageKeyT, PropertyKeyT, ValueT]
	SetDefault(ValueT) MemoryNodeBuilder[MessageKeyT, PropertyKeyT, ValueT]
	SetDefaultAndPresence(ValueT, bool) MemoryNodeBuilder[MessageKeyT, PropertyKeyT, ValueT]
}

type extractorMemoryNode[MessageKeyT comparable, PropertyKeyT comparable, ValueT any] struct {
	extractor func(*Locale, MessageKeyT) PropertyKeyT
	postOrder bool
	children map[PropertyKeyT]MemoryNode[MessageKeyT, ValueT]
	value ValueT
	hasValue bool
}

func(node *extractorMemoryNode[MessageKeyT, PropertyKeyT, ValueT]) MessageKeyTypeName() string {
	return GetMessageKeyTypeName[MessageKeyT]()
}

func(node *extractorMemoryNode[MessageKeyT, PropertyKeyT, ValueT]) Lookup(
	locale *Locale,
	messageKey MessageKeyT,
) (value ValueT, ok bool) {
	if !node.postOrder && node.hasValue {
		value = node.value
		ok = true
		return
	}
	if node.extractor != nil {
		property := node.extractor(locale, messageKey)
		child := node.children[property]
		if child != nil {
			value, ok = child.Lookup(locale, messageKey)
			if ok {
				return
			}
		}
	}
	if node.hasValue {
		value = node.value
		ok = true
	}
	return
}

func(node *extractorMemoryNode[MessageKeyT, PropertyKeyT, ValueT]) Put(
	propertyKey PropertyKeyT,
	subNode MemoryNode[MessageKeyT, ValueT],
) MemoryNodeBuilder[MessageKeyT, PropertyKeyT, ValueT] {
	if subNode == nil {
		delete(node.children, propertyKey)
	} else {
		node.children[propertyKey] = subNode
	}
	return node
}

func(node *extractorMemoryNode[MessageKeyT, PropertyKeyT, ValueT]) SetDefault(
	defaultValue ValueT,
) MemoryNodeBuilder[MessageKeyT, PropertyKeyT, ValueT] {
	node.value = defaultValue
	node.hasValue = true
	return node
}

func(node *extractorMemoryNode[MessageKeyT, PropertyKeyT, ValueT]) SetDefaultAndPresence(
	defaultValue ValueT,
	present bool,
) MemoryNodeBuilder[MessageKeyT, PropertyKeyT, ValueT] {
	node.value = defaultValue
	node.hasValue = present
	return node
}

func SetMemoryMessageRoot[MessageKeyT comparable, PropertyKeyT comparable](
	localizer *MemoryLocalizer[MessageKeyT],
	extractor func(*Locale, MessageKeyT) PropertyKeyT,
	postOrder bool,
) (builder MemoryNodeBuilder[MessageKeyT, PropertyKeyT, string]) {
	if localizer == nil || extractor == nil {
		return
	}
	builder = &extractorMemoryNode[MessageKeyT, PropertyKeyT, string] {
		extractor: extractor,
		postOrder: postOrder,
		children: make(map[PropertyKeyT]MemoryNode[MessageKeyT, string]),
	}
	localizer.MessageRoot = builder
	return
}

func SetMemoryNumberFormatRoot[MessageKeyT comparable, PropertyKeyT comparable](
	localizer *MemoryLocalizer[MessageKeyT],
	extractor func(*Locale, MessageKeyT) PropertyKeyT,
	postOrder bool,
) (builder MemoryNodeBuilder[MessageKeyT, PropertyKeyT, NumberFormats]) {
	if localizer == nil || extractor == nil {
		return
	}
	builder = &extractorMemoryNode[MessageKeyT, PropertyKeyT, NumberFormats] {
		extractor: extractor,
		postOrder: postOrder,
		children: make(map[PropertyKeyT]MemoryNode[MessageKeyT, NumberFormats]),
	}
	localizer.NumberFormatRoot = builder
	return
}

func PutMemoryNode[MessageKeyT comparable, ParentPropertyKeyT comparable, ChildPropertyKeyT comparable, ValueT any](
	parent MemoryNodeBuilder[MessageKeyT, ParentPropertyKeyT, ValueT],
	nodeKey ParentPropertyKeyT,
	extractor func(*Locale, MessageKeyT) ChildPropertyKeyT,
	postOrder bool,
) MemoryNodeBuilder[MessageKeyT, ChildPropertyKeyT, ValueT] {
	if parent == nil || extractor == nil {
		return nil
	}
	child := &extractorMemoryNode[MessageKeyT, ChildPropertyKeyT, ValueT] {
		extractor: extractor,
		postOrder: postOrder,
		children: make(map[ChildPropertyKeyT]MemoryNode[MessageKeyT, ValueT]),
	}
	parent.Put(nodeKey, child)
	return child
}
