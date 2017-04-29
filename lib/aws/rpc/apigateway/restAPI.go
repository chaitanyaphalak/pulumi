// *** WARNING: this file was generated by the Coconut IDL Compiler (CIDLC).  ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
    "errors"

    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/coconut/pkg/resource"
    "github.com/pulumi/coconut/pkg/tokens"
    "github.com/pulumi/coconut/pkg/util/contract"
    "github.com/pulumi/coconut/pkg/util/mapper"
    "github.com/pulumi/coconut/sdk/go/pkg/cocorpc"
)

/* RPC stubs for RestAPI resource provider */

// RestAPIToken is the type token corresponding to the RestAPI package type.
const RestAPIToken = tokens.Type("aws:apigateway/restAPI:RestAPI")

// RestAPIProviderOps is a pluggable interface for RestAPI-related management functionality.
type RestAPIProviderOps interface {
    Check(ctx context.Context, obj *RestAPI) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *RestAPI) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*RestAPI, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *RestAPI, new *RestAPI, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *RestAPI, new *RestAPI, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// RestAPIProvider is a dynamic gRPC-based plugin for managing RestAPI resources.
type RestAPIProvider struct {
    ops RestAPIProviderOps
}

// NewRestAPIProvider allocates a resource provider that delegates to a ops instance.
func NewRestAPIProvider(ops RestAPIProviderOps) cocorpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &RestAPIProvider{ops: ops}
}

func (p *RestAPIProvider) Check(
    ctx context.Context, req *cocorpc.CheckRequest) (*cocorpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(RestAPIToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr == nil || len(decerr.Failures()) == 0 {
        failures, err := p.ops.Check(ctx, obj)
        if err != nil {
            return nil, err
        }
        if len(failures) > 0 {
            decerr = mapper.NewDecodeErr(failures)
        }
    }
    return resource.NewCheckResponse(decerr), nil
}

func (p *RestAPIProvider) Name(
    ctx context.Context, req *cocorpc.NameRequest) (*cocorpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(RestAPIToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == "" {
        return nil, errors.New("Name property cannot be empty")
    }
    return &cocorpc.NameResponse{Name: obj.Name}, nil
}

func (p *RestAPIProvider) Create(
    ctx context.Context, req *cocorpc.CreateRequest) (*cocorpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(RestAPIToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    id, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &cocorpc.CreateResponse{
        Id:   string(id),
    }, nil
}

func (p *RestAPIProvider) Get(
    ctx context.Context, req *cocorpc.GetRequest) (*cocorpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(RestAPIToken))
    id := resource.ID(req.GetId())
    obj, err := p.ops.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return &cocorpc.GetResponse{
        Properties: resource.MarshalProperties(
            nil, resource.NewPropertyMap(obj), resource.MarshalOptions{}),
    }, nil
}

func (p *RestAPIProvider) InspectChange(
    ctx context.Context, req *cocorpc.ChangeRequest) (*cocorpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(RestAPIToken))
    old, oldprops, decerr := p.Unmarshal(req.GetOlds())
    if decerr != nil {
        return nil, decerr
    }
    new, newprops, decerr := p.Unmarshal(req.GetNews())
    if decerr != nil {
        return nil, decerr
    }
    diff := oldprops.Diff(newprops)
    var replaces []string
    if diff.Changed("name") {
        replaces = append(replaces, "name")
    }
    id := resource.ID(req.GetId())
    more, err := p.ops.InspectChange(ctx, id, old, new, diff)
    if err != nil {
        return nil, err
    }
    return &cocorpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *RestAPIProvider) Update(
    ctx context.Context, req *cocorpc.ChangeRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(RestAPIToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, id, old, new, diff); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *RestAPIProvider) Delete(
    ctx context.Context, req *cocorpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(RestAPIToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *RestAPIProvider) Unmarshal(
    v *pbstruct.Struct) (*RestAPI, resource.PropertyMap, mapper.DecodeError) {
    var obj RestAPI
    props := resource.UnmarshalProperties(v)
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable RestAPI structure(s) */

// RestAPI is a marshalable representation of its corresponding IDL type.
type RestAPI struct {
    Name string `json:"name"`
    Body *interface{} `json:"body,omitempty"`
    BodyS3Location *S3Location `json:"bodyS3Location,omitempty"`
    CloneFrom *resource.ID `json:"cloneFrom,omitempty"`
    Description *string `json:"description,omitempty"`
    FailOnWarnings *bool `json:"failOnWarnings,omitempty"`
    APIName *string `json:"apiName,omitempty"`
    Parameters *[]string `json:"parameters,omitempty"`
}

// RestAPI's properties have constants to make dealing with diffs and property bags easier.
const (
    RestAPI_Name = "name"
    RestAPI_Body = "body"
    RestAPI_BodyS3Location = "bodyS3Location"
    RestAPI_CloneFrom = "cloneFrom"
    RestAPI_Description = "description"
    RestAPI_FailOnWarnings = "failOnWarnings"
    RestAPI_APIName = "apiName"
    RestAPI_Parameters = "parameters"
)

/* Marshalable S3Location structure(s) */

// S3Location is a marshalable representation of its corresponding IDL type.
type S3Location struct {
    Object resource.ID `json:"object"`
    ETag *string `json:"etag,omitempty"`
    Version *string `json:"version,omitempty"`
}

// S3Location's properties have constants to make dealing with diffs and property bags easier.
const (
    S3Location_Object = "object"
    S3Location_ETag = "etag"
    S3Location_Version = "version"
)


