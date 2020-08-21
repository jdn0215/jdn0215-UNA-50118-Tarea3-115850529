/*
 * Books API
 *
 * This web service offers information on books
 *
 * API version: 0.1.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Book struct {

	BookId string `json:"bookId,omitempty"`

	Title string `json:"title,omitempty"`

	Edition string `json:"edition,omitempty"`

	Copyright string `json:"copyright,omitempty"`

	Language string `json:"language,omitempty"`

	Pages string `json:"pages,omitempty"`

	Author string `json:"author,omitempty"`

	Publisher string `json:"publisher,omitempty"`
}