// Code generated by goctl. DO NOT EDIT.
package types

type File struct {
	Id         int64  `json:"id"`         // 文件编号
	Name       string `json:"name"`       // 文件名
	Type       string `json:"type"`       // 文件类型
	Path       string `json:"path"`       // 文件路径
	Size       string `json:"size"`       // 文件大小
	ShareLink  string `json:"shareLink"`  // 分享链接
	ModifyTime int64  `json:"modifyTime"` //文件最后修改时间
}

type FileUploadReq struct {
	Name       string `json:"name"`                       // 文件名
	Type       string `json:"type"`                       // 文件类型
	FileSize   string `json:"fileSize"`                   // 文件大小
	SourcePath string `json:"sourcePath"`                 // 源文件本地路径
	Size       int    `json:"size, default = 100 * 1024"` // 分片大小
	Routines   int    `json:"routines, optional"`         // 并发数
}

type FileUploadResp struct {
}

type FileDownloadReq struct {
	Id   int64  `json:"id"`   // 文件编号
	Name string `json:"name"` // 文件名
	Type string `json:"type"` // 文件类型
}

type FileDownloadResp struct {
}

type FileListReq struct {
	Id   int64 `json:"id"`   // 文件编号
	Page int64 `json:"page"` // 页码
	Size int64 `json:"size"` // 每页展示记录条数
}

type FileListResp struct {
	List  []*File `json:"list"`  // 文件列表
	Count int64   `json:"count"` // 记录总数
}

type FileNameUpdateReq struct {
	Id   int64  `json:"id"`   // 文件编号
	Name string `json:"name"` // 文件名
}

type FileNameUpdateResp struct {
}

type FileCreateReq struct {
	ParentId int64  `json:"parentId"` // 父文件夹编号
	Name     string `json:"name"`     // 文件名
}

type FileCreateResp struct {
}

type FileDeletionReq struct {
	Id int64 `json:"id"` // 文件编号
}

type FileDeletionResp struct {
}

type FileMoveReq struct {
	LastParentId int64 `json:"lastParentId"` // 原父文件夹编号
	PreParentId  int64 `json:"preParentId"`  // 目标文件夹编号
	Id           int64 `json:"id"`           // 文件编号
}

type FileMoveResp struct {
}

type FileShareReq struct {
	Id int64 `json:"id"` // 文件编号
}

type FileShareResp struct {
	ShareLink string `json:"shareLink"` // 分享链接
}

type FileShareSaveReq struct {
	ParentId  int64  `json:"parentId"`  // 父文件夹编号
	ShareLink string `json:"shareLink"` // 分享链接
	Name      string `json:"name"`      // 文件名
}

type FileShareSaveResp struct {
	Id int64 `json:"id"` // 文件编号
}

type FileDetailsReq struct {
	Id int64 `json:"id"` // 文件编号
}

type FileDetailsResp struct {
	Name string `json:"name"` // 文件名
	Type string `json:"type"` // 文件类型
	Path string `json:"path"` // 文件路径
	Size string `json:"size"` // 文件大小
	Time int64  `json:"time"` // 文件创建时间
}
