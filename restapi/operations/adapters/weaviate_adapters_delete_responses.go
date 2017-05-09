/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package adapters




import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// WeaviateAdaptersDeleteNoContentCode is the HTTP code returned for type WeaviateAdaptersDeleteNoContent
const WeaviateAdaptersDeleteNoContentCode int = 204

/*WeaviateAdaptersDeleteNoContent Successful deleted.

swagger:response weaviateAdaptersDeleteNoContent
*/
type WeaviateAdaptersDeleteNoContent struct {
}

// NewWeaviateAdaptersDeleteNoContent creates WeaviateAdaptersDeleteNoContent with default headers values
func NewWeaviateAdaptersDeleteNoContent() *WeaviateAdaptersDeleteNoContent {
	return &WeaviateAdaptersDeleteNoContent{}
}

// WriteResponse to the client
func (o *WeaviateAdaptersDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// WeaviateAdaptersDeleteNotFoundCode is the HTTP code returned for type WeaviateAdaptersDeleteNotFound
const WeaviateAdaptersDeleteNotFoundCode int = 404

/*WeaviateAdaptersDeleteNotFound Successful query result but no resource was found.

swagger:response weaviateAdaptersDeleteNotFound
*/
type WeaviateAdaptersDeleteNotFound struct {
}

// NewWeaviateAdaptersDeleteNotFound creates WeaviateAdaptersDeleteNotFound with default headers values
func NewWeaviateAdaptersDeleteNotFound() *WeaviateAdaptersDeleteNotFound {
	return &WeaviateAdaptersDeleteNotFound{}
}

// WriteResponse to the client
func (o *WeaviateAdaptersDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// WeaviateAdaptersDeleteNotImplementedCode is the HTTP code returned for type WeaviateAdaptersDeleteNotImplemented
const WeaviateAdaptersDeleteNotImplementedCode int = 501

/*WeaviateAdaptersDeleteNotImplemented Not (yet) implemented.

swagger:response weaviateAdaptersDeleteNotImplemented
*/
type WeaviateAdaptersDeleteNotImplemented struct {
}

// NewWeaviateAdaptersDeleteNotImplemented creates WeaviateAdaptersDeleteNotImplemented with default headers values
func NewWeaviateAdaptersDeleteNotImplemented() *WeaviateAdaptersDeleteNotImplemented {
	return &WeaviateAdaptersDeleteNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateAdaptersDeleteNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
