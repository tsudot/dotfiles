Vim�UnDo� ���a�k'7�%&�:��w(���XU7N6{�              	                       U�X�    _�                             ����                                                                                                                                                                                                                                                                                                                                                             U�X�    �              
   // Code generated by ModelQ   �// plivo_routing_qualityoutgoingsmsrouter.go contains model for the database table [public.plivo_routing_qualityoutgoingsmsrouter]       package postgres       2type PlivoRoutingQualityoutgoingsmsrouter struct {   	Id       int `json:"id"`   	RouteId  int `json:"route_id"`   	Priority int `json:"priority"`   }�   
            �   	            �      
          	"encoding/json"�      	          	"encoding/gob"�                	"database/sql"�                import (�                 �                package postgres�                 �                �// plivo_routing_qualityoutgoingsmsrouter.go contains model for the database table [public.plivo_routing_qualityoutgoingsmsrouter]�                 // Code generated by ModelQ�   	   
       
   	"fmt"   	"github.com/mijia/modelq/gmq"   
	"strings"   )       2type PlivoRoutingQualityoutgoingsmsrouter struct {   	Id       int `json:"id"`   	RouteId  int `json:"route_id"`   	Priority int `json:"priority"`   }�                   :// Start of the PlivoRoutingQualityoutgoingsmsrouter APIs.       Afunc (obj PlivoRoutingQualityoutgoingsmsrouter) String() string {   0	if data, err := json.Marshal(obj); err != nil {   L		return fmt.Sprintf("<PlivoRoutingQualityoutgoingsmsrouter Id=%d>", obj.Id)   		} else {   		return string(data)   	}   }       ufunc (obj PlivoRoutingQualityoutgoingsmsrouter) Insert(dbtx gmq.DbTx) (PlivoRoutingQualityoutgoingsmsrouter, error) {   _	if result, err := PlivoRoutingQualityoutgoingsmsrouterObjs.Insert(obj).Run(dbtx); err != nil {   		return obj, err   		} else {   &		if dbtx.DriverName() != "postgres" {   4			if id, err := result.LastInsertId(); err != nil {   				return obj, err   			} else {   				obj.Id = int(id)   				return obj, err   			}   		}   		return obj, nil   	}   }       Vfunc (obj PlivoRoutingQualityoutgoingsmsrouter) Update(dbtx gmq.DbTx) (int64, error) {   *	fields := []string{"RouteId", "Priority"}   I	filter := PlivoRoutingQualityoutgoingsmsrouterObjs.FilterId("=", obj.Id)   x	if result, err := PlivoRoutingQualityoutgoingsmsrouterObjs.Update(obj, fields...).Where(filter).Run(dbtx); err != nil {   		return 0, err   		} else {   		return result.RowsAffected()   	}   }       Vfunc (obj PlivoRoutingQualityoutgoingsmsrouter) Delete(dbtx gmq.DbTx) (int64, error) {   I	filter := PlivoRoutingQualityoutgoingsmsrouterObjs.FilterId("=", obj.Id)   j	if result, err := PlivoRoutingQualityoutgoingsmsrouterObjs.Delete().Where(filter).Run(dbtx); err != nil {   		return 0, err   		} else {   		return result.RowsAffected()   	}   }       // Start of the inner Query Api       8type _PlivoRoutingQualityoutgoingsmsrouterQuery struct {   
	gmq.Query   }       tfunc (q _PlivoRoutingQualityoutgoingsmsrouterQuery) Where(f gmq.Filter) _PlivoRoutingQualityoutgoingsmsrouterQuery {   	q.Query = q.Query.Where(f)   		return q   }       vfunc (q _PlivoRoutingQualityoutgoingsmsrouterQuery) OrderBy(by ...string) _PlivoRoutingQualityoutgoingsmsrouterQuery {   "	tBy := make([]string, 0, len(by))   	for _, b := range by {   		sortDir := ""   !		if b[0] == '-' || b[0] == '+' {   			sortDir = string(b[0])   			b = b[1:]   		}   G		if col, ok := PlivoRoutingQualityoutgoingsmsrouterObjs.fcMap[b]; ok {   !			tBy = append(tBy, sortDir+col)   		}   	}   "	q.Query = q.Query.OrderBy(tBy...)   		return q   }       vfunc (q _PlivoRoutingQualityoutgoingsmsrouterQuery) GroupBy(by ...string) _PlivoRoutingQualityoutgoingsmsrouterQuery {   "	tBy := make([]string, 0, len(by))   	for _, b := range by {   G		if col, ok := PlivoRoutingQualityoutgoingsmsrouterObjs.fcMap[b]; ok {   			tBy = append(tBy, col)   		}   	}   "	q.Query = q.Query.GroupBy(tBy...)   		return q   }       xfunc (q _PlivoRoutingQualityoutgoingsmsrouterQuery) Limit(offsets ...int64) _PlivoRoutingQualityoutgoingsmsrouterQuery {   $	q.Query = q.Query.Limit(offsets...)   		return q   }       wfunc (q _PlivoRoutingQualityoutgoingsmsrouterQuery) Page(number, size int) _PlivoRoutingQualityoutgoingsmsrouterQuery {   %	q.Query = q.Query.Page(number, size)   		return q   }       \func (q _PlivoRoutingQualityoutgoingsmsrouterQuery) Run(dbtx gmq.DbTx) (sql.Result, error) {   	return q.Query.Exec(dbtx)   }       gtype PlivoRoutingQualityoutgoingsmsrouterRowVisitor func(obj PlivoRoutingQualityoutgoingsmsrouter) bool       �func (q _PlivoRoutingQualityoutgoingsmsrouterQuery) Iterate(dbtx gmq.DbTx, functor PlivoRoutingQualityoutgoingsmsrouterRowVisitor) error {   U	return q.Query.SelectList(dbtx, func(columns []gmq.Column, rb []sql.RawBytes) bool {   e		obj := PlivoRoutingQualityoutgoingsmsrouterObjs.toPlivoRoutingQualityoutgoingsmsrouter(columns, rb)   		return functor(obj)   	})   }       vfunc (q _PlivoRoutingQualityoutgoingsmsrouterQuery) One(dbtx gmq.DbTx) (PlivoRoutingQualityoutgoingsmsrouter, error) {   -	var obj PlivoRoutingQualityoutgoingsmsrouter   T	err := q.Query.SelectOne(dbtx, func(columns []gmq.Column, rb []sql.RawBytes) bool {   d		obj = PlivoRoutingQualityoutgoingsmsrouterObjs.toPlivoRoutingQualityoutgoingsmsrouter(columns, rb)   		return true   	})   	return obj, err   }       yfunc (q _PlivoRoutingQualityoutgoingsmsrouterQuery) List(dbtx gmq.DbTx) ([]PlivoRoutingQualityoutgoingsmsrouter, error) {   >	result := make([]PlivoRoutingQualityoutgoingsmsrouter, 0, 10)   U	err := q.Query.SelectList(dbtx, func(columns []gmq.Column, rb []sql.RawBytes) bool {   e		obj := PlivoRoutingQualityoutgoingsmsrouterObjs.toPlivoRoutingQualityoutgoingsmsrouter(columns, rb)   		result = append(result, obj)   		return true   	})   	return result, err   }       "// Start of the model facade Apis.       7type _PlivoRoutingQualityoutgoingsmsrouterObjs struct {   	fcMap map[string]string   }       Xfunc (o _PlivoRoutingQualityoutgoingsmsrouterObjs) Names() (schema, tbl, alias string) {   b	return "public", "plivo_routing_qualityoutgoingsmsrouter", "PlivoRoutingQualityoutgoingsmsrouter"   }       xfunc (o _PlivoRoutingQualityoutgoingsmsrouterObjs) Select(fields ...string) _PlivoRoutingQualityoutgoingsmsrouterQuery {   2	q := _PlivoRoutingQualityoutgoingsmsrouterQuery{}   	if len(fields) == 0 {   0		fields = []string{"Id", "RouteId", "Priority"}   	}   .	q.Query = gmq.Select(o, o.columns(fields...))   		return q   }       �func (o _PlivoRoutingQualityoutgoingsmsrouterObjs) Insert(obj PlivoRoutingQualityoutgoingsmsrouter) _PlivoRoutingQualityoutgoingsmsrouterQuery {   2	q := _PlivoRoutingQualityoutgoingsmsrouterQuery{}   G	q.Query = gmq.Insert(o, o.columnsWithData(obj, "RouteId", "Priority"))   		return q   }       �func (o _PlivoRoutingQualityoutgoingsmsrouterObjs) Update(obj PlivoRoutingQualityoutgoingsmsrouter, fields ...string) _PlivoRoutingQualityoutgoingsmsrouterQuery {   2	q := _PlivoRoutingQualityoutgoingsmsrouterQuery{}   ;	q.Query = gmq.Update(o, o.columnsWithData(obj, fields...))   		return q   }       hfunc (o _PlivoRoutingQualityoutgoingsmsrouterObjs) Delete() _PlivoRoutingQualityoutgoingsmsrouterQuery {   2	q := _PlivoRoutingQualityoutgoingsmsrouterQuery{}   	q.Query = gmq.Delete(o)   		return q   }       (///// Managed Objects Filters definition       efunc (o _PlivoRoutingQualityoutgoingsmsrouterObjs) FilterId(op string, p int, ps ...int) gmq.Filter {   )	params := make([]interface{}, 1+len(ps))   	params[0] = p   	for i := range ps {   		params[i+1] = ps[i]   	}   (	return o.newFilter("id", op, params...)   }       jfunc (o _PlivoRoutingQualityoutgoingsmsrouterObjs) FilterRouteId(op string, p int, ps ...int) gmq.Filter {   )	params := make([]interface{}, 1+len(ps))   	params[0] = p   	for i := range ps {   		params[i+1] = ps[i]   	}   .	return o.newFilter("route_id", op, params...)   }       kfunc (o _PlivoRoutingQualityoutgoingsmsrouterObjs) FilterPriority(op string, p int, ps ...int) gmq.Filter {   )	params := make([]interface{}, 1+len(ps))   	params[0] = p   	for i := range ps {   		params[i+1] = ps[i]   	}   .	return o.newFilter("priority", op, params...)   }       (///// Managed Objects Columns definition       Rfunc (o _PlivoRoutingQualityoutgoingsmsrouterObjs) ColumnId(p ...int) gmq.Column {   	var value interface{}   	if len(p) > 0 {   		value = p[0]   	}   	return gmq.Column{"id", value}   }       Wfunc (o _PlivoRoutingQualityoutgoingsmsrouterObjs) ColumnRouteId(p ...int) gmq.Column {   	var value interface{}   	if len(p) > 0 {   		value = p[0]   	}   %	return gmq.Column{"route_id", value}   }       Xfunc (o _PlivoRoutingQualityoutgoingsmsrouterObjs) ColumnPriority(p ...int) gmq.Column {   	var value interface{}   	if len(p) > 0 {   		value = p[0]   	}   %	return gmq.Column{"priority", value}   }       ////// Internal helper funcs       qfunc (o _PlivoRoutingQualityoutgoingsmsrouterObjs) newFilter(name, op string, params ...interface{}) gmq.Filter {   !	if strings.ToUpper(op) == "IN" {   #		return gmq.InFilter(name, params)   	}   +	return gmq.UnitFilter(name, op, params[0])   }       �func (o _PlivoRoutingQualityoutgoingsmsrouterObjs) toPlivoRoutingQualityoutgoingsmsrouter(columns []gmq.Column, rb []sql.RawBytes) PlivoRoutingQualityoutgoingsmsrouter {   .	obj := PlivoRoutingQualityoutgoingsmsrouter{}   	if len(columns) == len(rb) {   		for i := range columns {   			switch columns[i].Name {   			case "id":   				obj.Id = gmq.AsInt(rb[i])   			case "route_id":   "				obj.RouteId = gmq.AsInt(rb[i])   			case "priority":   #				obj.Priority = gmq.AsInt(rb[i])   			}   		}   	}   	return obj   }       [func (o _PlivoRoutingQualityoutgoingsmsrouterObjs) columns(fields ...string) []gmq.Column {   +	data := make([]gmq.Column, 0, len(fields))   	for _, f := range fields {   		switch f {   		case "Id":   $			data = append(data, o.ColumnId())   		case "RouteId":   )			data = append(data, o.ColumnRouteId())   		case "Priority":   *			data = append(data, o.ColumnPriority())   		}   	}   	return data   }       �func (o _PlivoRoutingQualityoutgoingsmsrouterObjs) columnsWithData(obj PlivoRoutingQualityoutgoingsmsrouter, fields ...string) []gmq.Column {   +	data := make([]gmq.Column, 0, len(fields))   	for _, f := range fields {   		switch f {   		case "Id":   *			data = append(data, o.ColumnId(obj.Id))   		case "RouteId":   4			data = append(data, o.ColumnRouteId(obj.RouteId))   		case "Priority":   6			data = append(data, o.ColumnPriority(obj.Priority))   		}   	}   	return data   }       Vvar PlivoRoutingQualityoutgoingsmsrouterObjs _PlivoRoutingQualityoutgoingsmsrouterObjs       func init() {   D	PlivoRoutingQualityoutgoingsmsrouterObjs.fcMap = map[string]string{   		"Id":       "id",   		"RouteId":  "route_id",   		"Priority": "priority",   	}   5	gob.Register(PlivoRoutingQualityoutgoingsmsrouter{})   }5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             U�X�    �              
   // Code generated by ModelQ   �// plivo_routing_qualityoutgoingsmsrouter.go contains model for the database table [public.plivo_routing_qualityoutgoingsmsrouter]       package model       2type PlivoRoutingQualityoutgoingsmsrouter struct {   	Id       int `json:"id"`   	RouteId  int `json:"route_id"`   	Priority int `json:"priority"`   }�   
            �   	             }�      
          	Priority int `json:"priority"`�      	          	RouteId  int `json:"route_id"`�                	Id       int `json:"id"`�                2type PlivoRoutingQualityoutgoingsmsrouter struct {�                 �         
      package postgres�                 �                �// plivo_routing_qualityoutgoingsmsrouter.go contains model for the database table [public.plivo_routing_qualityoutgoingsmsrouter]�                 // Code generated by ModelQ5�_�                           ����                                                                                                                                                                                                                                                                                                                                      	          ���    U�X�     �              
   // Code generated by ModelQ   �// plivo_routing_qualityoutgoingsmsrouter.go contains model for the database table [public.plivo_routing_qualityoutgoingsmsrouter]       package model       2type PlivoRoutingQualityoutgoingsmsrouter struct {   	Id       int   	RouteId  int   	Priority int   }�   
            �   	             }�      
          	Priority int �      	          	RouteId  int �                	Id       int �                2type PlivoRoutingQualityoutgoingsmsrouter struct {�                 �                package model�                 �                �// plivo_routing_qualityoutgoingsmsrouter.go contains model for the database table [public.plivo_routing_qualityoutgoingsmsrouter]�                 // Code generated by ModelQ�      
   
      	Id       int `json:"id"`   	RouteId  int `json:"route_id"`   	Priority int `json:"priority"`5�_�                           ����                                                                                                                                                                                                                                                                                                                                      	          ���    U�X�     �         
      2type PlivoRoutingQualityoutgoingsmsrouter struct {5�_�                           ����                                                                                                                                                                                                                                                                                                                                      	          ���    U�X�     �         
      &type Qualityoutgoingsmsrouter struct {5�_�                           ����                                                                                                                                                                                                                                                                                                                                      	          ���    U�X�    �         
      &type QualityOutgoingsmsrouter struct {5�_�                           ����                                                                                                                                                                                                                                                                                                                                      	          ���    U�X�     �              
   // Code generated by ModelQ   �// plivo_routing_qualityoutgoingsmsrouter.go contains model for the database table [public.plivo_routing_qualityoutgoingsmsrouter]       package model       &type QualityOutgoingSMSRouter struct {   	Id       int   	RouteId  int   	Priority int   }�   
            �   	             }�      
          	Priority int�      	          	RouteId  int�                	Id       int�                &type QualityOutgoingSMSRouter struct {�                 �                package model�                 �                �// plivo_routing_qualityoutgoingsmsrouter.go contains model for the database table [public.plivo_routing_qualityoutgoingsmsrouter]�                 // Code generated by ModelQ�              
   // Code generated by ModelQ   �// plivo_routing_qualityoutgoingsmsrouter.go contains model for the database table [public.plivo_routing_qualityoutgoingsmsrouter]       package model       &type QualityOutgoingSMSRouter struct {   	Id       int   	RouteId  int   	Priority int   }�   
            �   	             }�      
          	Priority int�      	          	RouteId  int�                	Id       int�         
      &type QualityOutgoingSMSrouter struct {�                 �                package model�                 �                �// plivo_routing_qualityoutgoingsmsrouter.go contains model for the database table [public.plivo_routing_qualityoutgoingsmsrouter]�                 // Code generated by ModelQ5�_�      	              
        ����                                                                                                                                                                                                                                                                                                                                      	          ���    U�X�     �   
            5�_�      
           	           ����                                                                                                                                                                                                                                                                                                                                      	          ���    U�X�     �               �               5�_�   	              
          ����                                                                                                                                                                                                                                                                                                                                      	          ���    U�X�     �               )func (p ProductPlan) TableName() string {5�_�   
                        ����                                                                                                                                                                                                                                                                                                                                      	          ���    U�X�     �               )func (q ProductPlan) TableName() string {5�_�                       	    ����                                                                                                                                                                                                                                                                                                                                      	          ���    U�X�    �               #	return "plivo_product_productplan"5�_�                        	    ����                                                                                                                                                                                                                                                                                                                                      	          ���    U�X�     �                 // Code generated by ModelQ   �// plivo_routing_qualityoutgoingsmsrouter.go contains model for the database table [public.plivo_routing_qualityoutgoingsmsrouter]       package model       &type QualityOutgoingSMSRouter struct {   	Id       int   	RouteId  int   	Priority int   }       6func (q QualityOutgoingSMSRouter) TableName() string {   0	return "plivo_routing_qualityoutgoingsmsrouter"   }�               �                }�                0	return "plivo_routing_qualityoutgoingsmsrouter"�                6func (q QualityOutgoingSMSRouter) TableName() string {�   
              �   	             }�      
          	Priority int�      	          	RouteId  int�                	Id       int�                &type QualityOutgoingSMSRouter struct {�                 �                package model�                 �                �// plivo_routing_qualityoutgoingsmsrouter.go contains model for the database table [public.plivo_routing_qualityoutgoingsmsrouter]�                 // Code generated by ModelQ�                 // Code generated by ModelQ   �// plivo_routing_qualityoutgoingsmsrouter.go contains model for the database table [public.plivo_routing_qualityoutgoingsmsrouter]       package model       &type QualityOutgoingSMSRouter struct {   	Id       int   	RouteId  int   	Priority int   }       6func (q QualityOutgoingSMSRouter) TableName() string {   0	return "plivo_routing_qualityoutgoingsmsrouter"   }�               �                }�                0	return "plivo_routing_qualityoutgoingsmsrouter"�                6func (q QualityOutgoingSMSRouter) TableName() string {�   
              �   	             }�      
          	Priority int�      	          	RouteId  int�                	Id       int�                &type QualityOutgoingSMSRouter struct {�                 �                package model�                 �                �// plivo_routing_qualityoutgoingsmsrouter.go contains model for the database table [public.plivo_routing_qualityoutgoingsmsrouter]�                 // Code generated by ModelQ�               
	return ""5��