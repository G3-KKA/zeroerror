# Idea.
 Packgae idea is to separate static errors,
 that are used to determine behavior if error happened.

 From context-of-error, like dynamic data, trace and/or messages,
 that, on the other hand, used for logging and debugging.

 This is [github.com/rs/zerolog] specific package,
 so it is not suitable to use with any other logger,
 or for different purpose.

 ---

# Usage
 Every exported bit are documented very clearly.


    return nil, zeroerror.WithMsg(ErrSomethingBadHappened,"details")

    zerologger.Debug().Func(TryInsert(err)).Send()

 ---
# Compatability.

 Fully compatable with [errors] package.  
 All compatable logic covered with succesful tests.  

 ---

# Test,Linter,Coverage.
 
    make lint
    make test
    make coverage

 Coverage 90+%.

 ---

# Hint.

 Any mention of a "context" in this package means messages or values,
 that stored inside [ZeroError], and not [context.Context].