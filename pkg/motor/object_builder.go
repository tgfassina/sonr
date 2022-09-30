package motor

import (
	"encoding/json"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	objectcli "github.com/sonr-io/sonr/internal/object"
	"github.com/sonr-io/sonr/internal/schemas"
	"github.com/sonr-io/sonr/pkg/motor/x/object"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (mtr *motorNodeImpl) NewObjectBuilder(did string) (*object.ObjectBuilder, error) {
	whatIs, _, found := mtr.Resources.GetSchema(did)
	if !found {
		return nil, fmt.Errorf("could not find WhatIs with did '%s'", did)
	}

	schemaImpl := schemas.NewWithClient(mtr.GetClient(), whatIs)
	objCli := objectcli.New(schemaImpl, shell.NewShell(mtr.Cosmos.GetIPFSApiAddress()))
	return object.NewBuilder(schemaImpl, objCli), nil
}

func (mtr *motorNodeImpl) GetDocument(req mt.GetDocumentRequest) (*mt.GetDocumentResponse, error) {
	obj, err := mtr.queryDocument(req.GetCid())
	if err != nil {
		return nil, err
	}

	doc := st.NewDocumentFromMap(req.GetCid(), obj)
	return &mt.GetDocumentResponse{
		Status:   200,
		Document: doc,
		Cid:      req.GetCid(),
	}, nil
}

func (mtr *motorNodeImpl) UploadDocument(req mt.UploadDocumentRequest) (*mt.UploadDocumentResponse, error) {
	var obj map[string]interface{}
	if err := json.Unmarshal(req.GetDocument(), &obj); err != nil {
		return nil, fmt.Errorf("error decoding document JSON")
	}

	builder, err := mtr.NewObjectBuilder(req.GetSchemaDid())
	if err != nil {
		return nil, err
	}

	builder.SetLabel(req.GetLabel())
	for k, v := range obj {
		builder.Set(k, v)
	}

	resp, err := builder.Upload()
	if err != nil {
		return nil, err
	}

	return &mt.UploadDocumentResponse{
		Status:   resp.Status,
		Cid:      resp.Cid,
		Document: resp.Document,
	}, nil
}
