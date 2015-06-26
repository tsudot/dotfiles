Vim�UnDo� ���뼄� ����V�:���n�0�W!��   H                                  UTvs    _�                     !        ����                                                                                                                                                                                                                                                                                                                                                             U@��    �              J   package handler       import (   	"net/http"   
	"strconv"       	"github.com/gorilla/context"   	"github.com/gorilla/mux"   	"github.com/jinzhu/gorm"   !	"github.com/plivo/gotham/gotham"   &	"github.com/plivo/gotham/gotham/util"   .	"github.com/plivo/gotham/gotham/util/blunder"   )       var (   6	// Account allows access to all the account handlers.   	Account *accountHandlerSet   )       func init() {   !	Account = new(accountHandlerSet)   }       type accountHandlerSet struct{}        type accountAPIResponse struct {   "	AccountID int `json:"account_id"`   	MPS       int `json:"MPS"`   }       @// accountHandlerSet handles the GET request for the Account API   Tfunc (a *accountHandlerSet) AccountHandler(w http.ResponseWriter, r *http.Request) {   :	ctx := context.Get(r, gotham.ContextKey).(gotham.Context)   3	w.Header().Set("Content-Type", "application/json")   	vars := mux.Vars(r)   3	accountID, err := strconv.Atoi(vars["account_id"])       	if err != nil {   		ctx.Logger.Error(   0			"Account ID passed: '%d', should be integer",   			accountID,   		)   $		w.WriteHeader(http.StatusNotFound)   7		w.Write(util.Response.Error(blunder.AccountNotFound))   		return   	}       :	account, err := gotham.Database.GetAccountByID(accountID)   	if err != nil {   !		if err == gorm.RecordNotFound {   %			w.WriteHeader(http.StatusNotFound)   8			w.Write(util.Response.Error(blunder.AccountNotFound))   				return   		}   		ctx.Logger.Error(   4			"error while creating message detail record: %s",   			err.Error(),   		)   /		w.WriteHeader(http.StatusInternalServerError)   5		w.Write(util.Response.Error(blunder.InternalError))   		return   	}   	ctx.Logger.Info(   !		"MPS for account ID: %d is %d",   !		account.Id, account.MpsAllowed,   	)        	var response accountAPIResponse    	response.AccountID = account.Id   "	response.MPS = account.MpsAllowed       -	w.Write(util.Response.Success(response, ""))   	return   }�   J            �              J   package handler       import (   	"net/http"   
	"strconv"       	"github.com/gorilla/context"   	"github.com/gorilla/mux"   	"github.com/jinzhu/gorm"   !	"github.com/plivo/gotham/gotham"   &	"github.com/plivo/gotham/gotham/util"   .	"github.com/plivo/gotham/gotham/util/blunder"   )       var (   6	// Account allows access to all the account handlers.   	Account *accountHandlerSet   )       func init() {   !	Account = new(accountHandlerSet)   }       type accountHandlerSet struct{}        type accountAPIResponse struct {   "	AccountID int `json:"account_id"`   	MPS       int `json:"MPS"`   }       @// accountHandlerSet handles the GET request for the Account API   Tfunc (a *accountHandlerSet) AccountHandler(w http.ResponseWriter, r *http.Request) {   :	ctx := context.Get(r, gotham.ContextKey).(gotham.Context)   3	w.Header().Set("Content-Type", "application/json")   	vars := mux.Vars(r)   3	accountID, err := strconv.Atoi(vars["account_id"])       	if err != nil {   		ctx.Logger.Error(   0			"Account ID passed: '%d', should be integer",   			accountID,   		)   $		w.WriteHeader(http.StatusNotFound)   7		w.Write(util.Response.Error(blunder.AccountNotFound))   		return   	}       :	account, err := gotham.Database.GetAccountByID(accountID)   	if err != nil {   !		if err == gorm.RecordNotFound {   %			w.WriteHeader(http.StatusNotFound)   8			w.Write(util.Response.Error(blunder.AccountNotFound))   				return   		}   		ctx.Logger.Error(   4			"error while creating message detail record: %s",   			err.Error(),   		)   /		w.WriteHeader(http.StatusInternalServerError)   5		w.Write(util.Response.Error(blunder.InternalError))   		return   	}   	ctx.Logger.Info(   !		"MPS for account ID: %d is %d",   !		account.Id, account.MpsAllowed,   	)        	var response accountAPIResponse    	response.AccountID = account.Id   "	response.MPS = account.MpsAllowed       -	w.Write(util.Response.Success(response, ""))   	return   }�   J            �       "          F	ctx := context.Get(r, gotham.GothamContextKey).(gotham.GothamContext)5�_�                    H       ����                                                                                                                                                                                                                                                                                                                                                             UA˕     �              J   package handler       import (   	"net/http"   
	"strconv"       	"github.com/gorilla/context"   	"github.com/gorilla/mux"   	"github.com/jinzhu/gorm"   !	"github.com/plivo/gotham/gotham"   &	"github.com/plivo/gotham/gotham/util"   .	"github.com/plivo/gotham/gotham/util/blunder"   )       var (   6	// Account allows access to all the account handlers.   	Account *accountHandlerSet   )       func init() {   !	Account = new(accountHandlerSet)   }       type accountHandlerSet struct{}        type accountAPIResponse struct {   "	AccountID int `json:"account_id"`   	MPS       int `json:"MPS"`   }       @// accountHandlerSet handles the GET request for the Account API   Tfunc (a *accountHandlerSet) AccountHandler(w http.ResponseWriter, r *http.Request) {   :	ctx := context.Get(r, gotham.ContextKey).(gotham.Context)   3	w.Header().Set("Content-Type", "application/json")   	vars := mux.Vars(r)   3	accountID, err := strconv.Atoi(vars["account_id"])       	if err != nil {   		ctx.Logger.Error(   0			"Account ID passed: '%d', should be integer",   			accountID,   		)   $		w.WriteHeader(http.StatusNotFound)   7		w.Write(util.Response.Error(blunder.AccountNotFound))   		return   	}       :	account, err := gotham.Database.GetAccountByID(accountID)   	if err != nil {   !		if err == gorm.RecordNotFound {   %			w.WriteHeader(http.StatusNotFound)   8			w.Write(util.Response.Error(blunder.AccountNotFound))   				return   		}   		ctx.Logger.Error(   4			"error while creating message detail record: %s",   			err.Error(),   		)   /		w.WriteHeader(http.StatusInternalServerError)   5		w.Write(util.Response.Error(blunder.InternalError))   		return   	}   	ctx.Logger.Info(   !		"MPS for account ID: %d is %d",   !		account.Id, account.MpsAllowed,   	)        	var response accountAPIResponse    	response.AccountID = account.Id   "	response.MPS = account.MpsAllowed       2	w.Write(util.Response.Success(ctx, response, ""))   	return   }�   J            �              J   package handler       import (   	"net/http"   
	"strconv"       	"github.com/gorilla/context"   	"github.com/gorilla/mux"   	"github.com/jinzhu/gorm"   !	"github.com/plivo/gotham/gotham"   &	"github.com/plivo/gotham/gotham/util"   .	"github.com/plivo/gotham/gotham/util/blunder"   )       var (   6	// Account allows access to all the account handlers.   	Account *accountHandlerSet   )       func init() {   !	Account = new(accountHandlerSet)   }       type accountHandlerSet struct{}        type accountAPIResponse struct {   "	AccountID int `json:"account_id"`   	MPS       int `json:"MPS"`   }       @// accountHandlerSet handles the GET request for the Account API   Tfunc (a *accountHandlerSet) AccountHandler(w http.ResponseWriter, r *http.Request) {   :	ctx := context.Get(r, gotham.ContextKey).(gotham.Context)   3	w.Header().Set("Content-Type", "application/json")   	vars := mux.Vars(r)   3	accountID, err := strconv.Atoi(vars["account_id"])       	if err != nil {   		ctx.Logger.Error(   0			"Account ID passed: '%d', should be integer",   			accountID,   		)   $		w.WriteHeader(http.StatusNotFound)   7		w.Write(util.Response.Error(blunder.AccountNotFound))   		return   	}       :	account, err := gotham.Database.GetAccountByID(accountID)   	if err != nil {   !		if err == gorm.RecordNotFound {   %			w.WriteHeader(http.StatusNotFound)   8			w.Write(util.Response.Error(blunder.AccountNotFound))   				return   		}   		ctx.Logger.Error(   4			"error while creating message detail record: %s",   			err.Error(),   		)   /		w.WriteHeader(http.StatusInternalServerError)   5		w.Write(util.Response.Error(blunder.InternalError))   		return   	}   	ctx.Logger.Info(   !		"MPS for account ID: %d is %d",   !		account.Id, account.MpsAllowed,   	)        	var response accountAPIResponse    	response.AccountID = account.Id   "	response.MPS = account.MpsAllowed       2	w.Write(util.Response.Success(ctx, response, ""))   	return   }�   J            �   G   I   J      -	w.Write(util.Response.Success(response, ""))5�_�                    $       ����                                                                                                                                                                                                                                                                                                                                                             UTv     �   $   '   K      	�   $   &   J    5�_�                    A       ����                                                                                                                                                                                                                                                                                                                                                             UTv_     �   @   A          	ctx.Logger.Info(5�_�                    A       ����                                                                                                                                                                                                                                                                                                                                                             UTv`     �   @   A          !		"MPS for account ID: %d is %d",5�_�                    A       ����                                                                                                                                                                                                                                                                                                                                                             UTv`     �   @   A          !		account.Id, account.MpsAllowed,5�_�                    A       ����                                                                                                                                                                                                                                                                                                                                                             UTva     �   @   A          	)5�_�      	              A        ����                                                                                                                                                                                                                                                                                                                                                             UTvb     �   @   A           5�_�      
           	   @       ����                                                                                                                                                                                                                                                                                                                                                             UTvb    �   @   B   G    5�_�   	              
   &       ����                                                                                                                                                                                                                                                                                                                                                             UTvn    �   %   '   H      N	ctx.Logger.Info('GET Account request received for account ID: %d', accountID)5�_�   
                  &   B    ����                                                                                                                                                                                                                                                                                                                                                             UTvr     �              H   package handler       import (   	"net/http"   
	"strconv"       	"github.com/gorilla/context"   	"github.com/gorilla/mux"   	"github.com/jinzhu/gorm"   !	"github.com/plivo/gotham/gotham"   &	"github.com/plivo/gotham/gotham/util"   .	"github.com/plivo/gotham/gotham/util/blunder"   )       var (   6	// Account allows access to all the account handlers.   	Account *accountHandlerSet   )       func init() {   !	Account = new(accountHandlerSet)   }       type accountHandlerSet struct{}        type accountAPIResponse struct {   "	AccountID int `json:"account_id"`   	MPS       int `json:"MPS"`   }       @// accountHandlerSet handles the GET request for the Account API   Tfunc (a *accountHandlerSet) AccountHandler(w http.ResponseWriter, r *http.Request) {   :	ctx := context.Get(r, gotham.ContextKey).(gotham.Context)   3	w.Header().Set("Content-Type", "application/json")   	vars := mux.Vars(r)   3	accountID, err := strconv.Atoi(vars["account_id"])       N	ctx.Logger.Info("GET Account request received for account ID: %d", accountID)       	if err != nil {   		ctx.Logger.Error(   0			"Account ID passed: '%d', should be integer",   			accountID,   		)   $		w.WriteHeader(http.StatusNotFound)   7		w.Write(util.Response.Error(blunder.AccountNotFound))   		return   	}       :	account, err := gotham.Database.GetAccountByID(accountID)   	if err != nil {   !		if err == gorm.RecordNotFound {   %			w.WriteHeader(http.StatusNotFound)   8			w.Write(util.Response.Error(blunder.AccountNotFound))   				return   		}   		ctx.Logger.Error(   4			"error while creating message detail record: %s",   			err.Error(),   		)   /		w.WriteHeader(http.StatusInternalServerError)   5		w.Write(util.Response.Error(blunder.InternalError))   		return   	}        	var response accountAPIResponse    	response.AccountID = account.Id   "	response.MPS = account.MpsAllowed       2	w.Write(util.Response.Success(ctx, response, ""))   	return   }�   H            �              H   package handler       import (   	"net/http"   
	"strconv"       	"github.com/gorilla/context"   	"github.com/gorilla/mux"   	"github.com/jinzhu/gorm"   !	"github.com/plivo/gotham/gotham"   &	"github.com/plivo/gotham/gotham/util"   .	"github.com/plivo/gotham/gotham/util/blunder"   )       var (   6	// Account allows access to all the account handlers.   	Account *accountHandlerSet   )       func init() {   !	Account = new(accountHandlerSet)   }       type accountHandlerSet struct{}        type accountAPIResponse struct {   "	AccountID int `json:"account_id"`   	MPS       int `json:"MPS"`   }       @// accountHandlerSet handles the GET request for the Account API   Tfunc (a *accountHandlerSet) AccountHandler(w http.ResponseWriter, r *http.Request) {   :	ctx := context.Get(r, gotham.ContextKey).(gotham.Context)   3	w.Header().Set("Content-Type", "application/json")   	vars := mux.Vars(r)   3	accountID, err := strconv.Atoi(vars["account_id"])       N	ctx.Logger.Info("GET Account request received for account ID: %d", accountID)       	if err != nil {   		ctx.Logger.Error(   0			"Account ID passed: '%d', should be integer",   			accountID,   		)   $		w.WriteHeader(http.StatusNotFound)   7		w.Write(util.Response.Error(blunder.AccountNotFound))   		return   	}       :	account, err := gotham.Database.GetAccountByID(accountID)   	if err != nil {   !		if err == gorm.RecordNotFound {   %			w.WriteHeader(http.StatusNotFound)   8			w.Write(util.Response.Error(blunder.AccountNotFound))   				return   		}   		ctx.Logger.Error(   4			"error while creating message detail record: %s",   			err.Error(),   		)   /		w.WriteHeader(http.StatusInternalServerError)   5		w.Write(util.Response.Error(blunder.InternalError))   		return   	}        	var response accountAPIResponse    	response.AccountID = account.Id   "	response.MPS = account.MpsAllowed       2	w.Write(util.Response.Success(ctx, response, ""))   	return   }�   H            �   %   '   H      N	ctx.Logger.Info("GET Account request received for account ID: %d', accountID)5��