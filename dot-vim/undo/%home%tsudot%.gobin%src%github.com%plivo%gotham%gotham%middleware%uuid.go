Vim�UnDo� ��΃��� k�w3+g�7P�'$a�ٵ��   &                                  U@��    _�                            ����                                                                                                                                                                                                                                                                                                                                                             U@�     �         &      ?		context.Set(r, gotham.GothamContextKey, gotham.GothamContext{5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             U@�      �         &      >		context.Set(r, gotham.othamContextKey, gotham.GothamContext{5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             U@�      �         &      =		context.Set(r, gotham.thamContextKey, gotham.GothamContext{5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             U@�      �         &      <		context.Set(r, gotham.hamContextKey, gotham.GothamContext{5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             U@�      �         &      ;		context.Set(r, gotham.amContextKey, gotham.GothamContext{5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             U@�!     �         &      :		context.Set(r, gotham.mContextKey, gotham.GothamContext{5�_�                       +    ����                                                                                                                                                                                                                                                                                                                                                             U@�#     �         &      9		context.Set(r, gotham.ContextKey, gotham.GothamContext{5�_�      	                 +    ����                                                                                                                                                                                                                                                                                                                                                             U@�$     �         &      8		context.Set(r, gotham.ContextKey, gotham.othamContext{5�_�      
           	      +    ����                                                                                                                                                                                                                                                                                                                                                             U@�$     �         &      7		context.Set(r, gotham.ContextKey, gotham.thamContext{5�_�   	              
      +    ����                                                                                                                                                                                                                                                                                                                                                             U@�$     �         &      6		context.Set(r, gotham.ContextKey, gotham.hamContext{5�_�   
                    +    ����                                                                                                                                                                                                                                                                                                                                                             U@�$    �         &      5		context.Set(r, gotham.ContextKey, gotham.amContext{5�_�                       +    ����                                                                                                                                                                                                                                                                                                                                                             U@�%     �              &   package middleware       3// X-UUID middleware to get the UUID of the request   import (   	"encoding/json"   	"net/http"       	"github.com/gorilla/context"   !	"github.com/plivo/gotham/gotham"   )       type UUIDMiddleware struct{}       cfunc (m *UUIDMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {   	// fetch the x-uuid header   	uuid := r.Header.Get("X-UUID")   	if uuid == "" {   0		// the header is not set. throw a bad reqeust.   4		w.Header().Set("Content-Type", "application/json")   &		w.WriteHeader(http.StatusBadRequest)   1		resp, _ := json.Marshal(map[string]interface{}{   &			"message": "X-UUID header missing",   		})   		w.Write(resp)   0		// don't pass the request further up the chain   		} else {   3		context.Set(r, gotham.ContextKey, gotham.Context{   			UUID:   uuid,   (			Logger: gotham.NewGothamLogger(uuid),   		})   "		// pass the request up the chain   		next(w, r)   	}   }       *func NewUUIDMiddleware() *UUIDMiddleware {   	return &UUIDMiddleware{}   }�   &            �              &   package middleware       3// X-UUID middleware to get the UUID of the request   import (   	"encoding/json"   	"net/http"       	"github.com/gorilla/context"   !	"github.com/plivo/gotham/gotham"   )       type UUIDMiddleware struct{}       cfunc (m *UUIDMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {   	// fetch the x-uuid header   	uuid := r.Header.Get("X-UUID")   	if uuid == "" {   0		// the header is not set. throw a bad reqeust.   4		w.Header().Set("Content-Type", "application/json")   &		w.WriteHeader(http.StatusBadRequest)   1		resp, _ := json.Marshal(map[string]interface{}{   &			"message": "X-UUID header missing",   		})   		w.Write(resp)   0		// don't pass the request further up the chain   		} else {   3		context.Set(r, gotham.ContextKey, gotham.Context{   			UUID:   uuid,   (			Logger: gotham.NewGothamLogger(uuid),   		})   "		// pass the request up the chain   		next(w, r)   	}   }       *func NewUUIDMiddleware() *UUIDMiddleware {   	return &UUIDMiddleware{}   }�   &            �         &      4		context.Set(r, gotham.ContextKey, gotham.mContext{5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             U@��     �         &      (			Logger: gotham.NewGothamLogger(uuid),5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             U@��     �         &      '			Logger: gotham.NewGthamLogger(uuid),5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             U@��     �         &      &			Logger: gotham.NewthamLogger(uuid),5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             U@��     �         &      %			Logger: gotham.NewhamLogger(uuid),5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             U@��    �         &      $			Logger: gotham.NewamLogger(uuid),5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             U@��     �              &   package middleware       3// X-UUID middleware to get the UUID of the request   import (   	"encoding/json"   	"net/http"       	"github.com/gorilla/context"   !	"github.com/plivo/gotham/gotham"   )       type UUIDMiddleware struct{}       cfunc (m *UUIDMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {   	// fetch the x-uuid header   	uuid := r.Header.Get("X-UUID")   	if uuid == "" {   0		// the header is not set. throw a bad reqeust.   4		w.Header().Set("Content-Type", "application/json")   &		w.WriteHeader(http.StatusBadRequest)   1		resp, _ := json.Marshal(map[string]interface{}{   &			"message": "X-UUID header missing",   		})   		w.Write(resp)   0		// don't pass the request further up the chain   		} else {   3		context.Set(r, gotham.ContextKey, gotham.Context{   			UUID:   uuid,   "			Logger: gotham.NewLogger(uuid),   		})   "		// pass the request up the chain   		next(w, r)   	}   }       *func NewUUIDMiddleware() *UUIDMiddleware {   	return &UUIDMiddleware{}   }�   &            �              &   package middleware       3// X-UUID middleware to get the UUID of the request   import (   	"encoding/json"   	"net/http"       	"github.com/gorilla/context"   !	"github.com/plivo/gotham/gotham"   )       type UUIDMiddleware struct{}       cfunc (m *UUIDMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {   	// fetch the x-uuid header   	uuid := r.Header.Get("X-UUID")   	if uuid == "" {   0		// the header is not set. throw a bad reqeust.   4		w.Header().Set("Content-Type", "application/json")   &		w.WriteHeader(http.StatusBadRequest)   1		resp, _ := json.Marshal(map[string]interface{}{   &			"message": "X-UUID header missing",   		})   		w.Write(resp)   0		// don't pass the request further up the chain   		} else {   3		context.Set(r, gotham.ContextKey, gotham.Context{   			UUID:   uuid,   "			Logger: gotham.NewLogger(uuid),   		})   "		// pass the request up the chain   		next(w, r)   	}   }       *func NewUUIDMiddleware() *UUIDMiddleware {   	return &UUIDMiddleware{}   }�   &            �         &      #			Logger: gotham.NewmLogger(uuid),5��