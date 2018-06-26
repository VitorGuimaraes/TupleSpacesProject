// Put: adds a tuple to a space.

// Query: blocks until a tuple is found which matches a given template. It then returns the matched tuple.
// QueryP: searches for a tuple atching a template. It then returns the matched tuple (if any).
// QueryAll: returns all tuples matching a template.

// Get: like Query but also removes the found tuple.
// GetP: like QueryP but also removes the found tuple (if any).
// GetAll: returns all tuples matching a template and removes them from the space.