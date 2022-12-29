// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    har, err := UnmarshalHar(bytes)
//    bytes, err = har.Marshal()

package nlibshared

import "bytes"
import "errors"
import "encoding/json"

func UnmarshalHar(data []byte) (Har, error) {
	var r Har
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Har) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Har struct {
	Log Log `json:"log"`
}

type Log struct {
	Browser *Browser `json:"browser,omitempty"`
	Comment *string  `json:"comment,omitempty"`
	Creator Creator  `json:"creator"`          
	Entries []Entry  `json:"entries"`          
	Pages   []Page   `json:"pages,omitempty"`  
	Version string   `json:"version"`          
}

type Browser struct {
	Comment *string `json:"comment,omitempty"`
	Name    string  `json:"name"`             
	Version string  `json:"version"`          
}

type Creator struct {
	Comment *string `json:"comment,omitempty"`
	Name    string  `json:"name"`             
	Version string  `json:"version"`          
}

type Entry struct {
	Cache           *CacheUnion   `json:"cache"`                    
	Comment         *string       `json:"comment,omitempty"`        
	Connection      *string       `json:"connection,omitempty"`     
	Pageref         *string       `json:"pageref,omitempty"`        
	Request         Request       `json:"request"`                  
	Response        Response      `json:"response"`                 
	ServerIPAddress *string       `json:"serverIPAddress,omitempty"`
	StartedDateTime string        `json:"startedDateTime"`          
	Time            float64       `json:"time"`                     
	Timings         *TimingsUnion `json:"timings"`                  
}

type CacheClass struct {
	AfterRequest  *AfterRequest  `json:"afterRequest"`     
	BeforeRequest *BeforeRequest `json:"beforeRequest"`    
	Comment       *string        `json:"comment,omitempty"`
}

type AfterRequest struct {
	Comment    *string `json:"comment,omitempty"`
	ETag       string  `json:"eTag"`             
	Expires    *string `json:"expires,omitempty"`
	HitCount   int64   `json:"hitCount"`         
	LastAccess string  `json:"lastAccess"`       
}

type BeforeRequest struct {
	Comment    *string `json:"comment,omitempty"`
	ETag       string  `json:"eTag"`             
	Expires    *string `json:"expires,omitempty"`
	HitCount   int64   `json:"hitCount"`         
	LastAccess string  `json:"lastAccess"`       
}

type Request struct {
	BodySize    int64     `json:"bodySize"`          
	Comment     *string   `json:"comment,omitempty"` 
	Cookies     []Cookie  `json:"cookies"`           
	Headers     []Header  `json:"headers"`           
	HeadersSize int64     `json:"headersSize"`       
	HTTPVersion string    `json:"httpVersion"`       
	Method      string    `json:"method"`            
	PostData    *PostData `json:"postData,omitempty"`
	QueryString []Query   `json:"queryString"`       
	URL         string    `json:"url"`               
}

type Cookie struct {
	Comment  *string `json:"comment,omitempty"` 
	Domain   *string `json:"domain,omitempty"`  
	Expires  *string `json:"expires"`           
	HTTPOnly *bool   `json:"httpOnly,omitempty"`
	Name     string  `json:"name"`              
	Path     *string `json:"path,omitempty"`    
	Secure   *bool   `json:"secure,omitempty"`  
	Value    string  `json:"value"`             
}

type Header struct {
	Comment *string `json:"comment,omitempty"`
	Name    string  `json:"name"`             
	Value   string  `json:"value"`            
}

type PostData struct {
	Comment  *string       `json:"comment,omitempty"`
	MIMEType string        `json:"mimeType"`         
	Params   []interface{} `json:"params,omitempty"` 
	Text     *string       `json:"text,omitempty"`   
}

type Query struct {
	Comment *string `json:"comment,omitempty"`
	Name    string  `json:"name"`             
	Value   string  `json:"value"`            
}

type Response struct {
	BodySize    int64    `json:"bodySize"`         
	Comment     *string  `json:"comment,omitempty"`
	Content     Content  `json:"content"`          
	Cookies     []Cookie `json:"cookies"`          
	Headers     []Header `json:"headers"`          
	HeadersSize int64    `json:"headersSize"`      
	HTTPVersion string   `json:"httpVersion"`      
	RedirectURL string   `json:"redirectURL"`      
	Status      int64    `json:"status"`           
	StatusText  string   `json:"statusText"`       
}

type Content struct {
	Comment     *string `json:"comment,omitempty"`    
	Compression *int64  `json:"compression,omitempty"`
	Encoding    *string `json:"encoding,omitempty"`   
	MIMEType    string  `json:"mimeType"`             
	Size        int64   `json:"size"`                 
	Text        *string `json:"text,omitempty"`       
}

type TimingsClass struct {
	Blocked *float64 `json:"blocked,omitempty"`
	Comment *string  `json:"comment,omitempty"`
	Connect *float64 `json:"connect,omitempty"`
	DNS     *float64 `json:"dns,omitempty"`    
	Receive float64  `json:"receive"`          
	Send    float64  `json:"send"`             
	SSL     *float64 `json:"ssl,omitempty"`    
	Wait    float64  `json:"wait"`             
}

type Page struct {
	Comment         *string     `json:"comment,omitempty"`
	ID              string      `json:"id"`               
	PageTimings     PageTimings `json:"pageTimings"`      
	StartedDateTime string      `json:"startedDateTime"`  
	Title           string      `json:"title"`            
}

type PageTimings struct {
	Comment       *string  `json:"comment,omitempty"`      
	OnContentLoad *float64 `json:"onContentLoad,omitempty"`
	OnLoad        *float64 `json:"onLoad,omitempty"`       
}

type CacheUnion struct {
	AnythingArray []interface{}
	Bool          *bool
	CacheClass    *CacheClass
	Double        *float64
	Integer       *int64
	String        *string
}

func (x *CacheUnion) UnmarshalJSON(data []byte) error {
	x.AnythingArray = nil
	x.CacheClass = nil
	var c CacheClass
	object, err := unmarshalUnion(data, &x.Integer, &x.Double, &x.Bool, &x.String, true, &x.AnythingArray, true, &c, false, nil, false, nil, true)
	if err != nil {
		return err
	}
	if object {
		x.CacheClass = &c
	}
	return nil
}

func (x *CacheUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, x.Double, x.Bool, x.String, x.AnythingArray != nil, x.AnythingArray, x.CacheClass != nil, x.CacheClass, false, nil, false, nil, true)
}

type TimingsUnion struct {
	AnythingArray []interface{}
	Bool          *bool
	Double        *float64
	Integer       *int64
	String        *string
	TimingsClass  *TimingsClass
}

func (x *TimingsUnion) UnmarshalJSON(data []byte) error {
	x.AnythingArray = nil
	x.TimingsClass = nil
	var c TimingsClass
	object, err := unmarshalUnion(data, &x.Integer, &x.Double, &x.Bool, &x.String, true, &x.AnythingArray, true, &c, false, nil, false, nil, true)
	if err != nil {
		return err
	}
	if object {
		x.TimingsClass = &c
	}
	return nil
}

func (x *TimingsUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, x.Double, x.Bool, x.String, x.AnythingArray != nil, x.AnythingArray, x.TimingsClass != nil, x.TimingsClass, false, nil, false, nil, true)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
