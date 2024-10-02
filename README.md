# Idea.
 Packgae idea is to separate static errors,
 that are used to determine behaviour if error happened.

 From context-of-error, like dynamic data, trace and/or messages,
 that, on the other hand, used for logging and debugging.

 This is [github.com/rs/zerolog] specific package,
 so it is not suitable to use with any other logger,
 or for different purpose.

 ---

# Compatability.

 Partially compatable with [errors] package, errors.As() not supported.

 ---

# Hint.

 Any mention of a "context" in this package means messages or values,
 that stored inside [ZeroError], and not [context.Context].