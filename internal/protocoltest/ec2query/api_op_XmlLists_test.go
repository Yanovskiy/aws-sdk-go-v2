// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2query

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/ec2query/types"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithyrand "github.com/awslabs/smithy-go/rand"
	smithytesting "github.com/awslabs/smithy-go/testing"
	smithytime "github.com/awslabs/smithy-go/time"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestClient_XmlLists_awsEc2queryDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *XmlListsOutput
	}{
		// Tests for XML list serialization
		"Ec2XmlLists": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml;charset=UTF-8"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlListsResponse xmlns="https://example.com/">
			    <stringList>
			        <member>foo</member>
			        <member>bar</member>
			    </stringList>
			    <stringSet>
			        <member>foo</member>
			        <member>bar</member>
			    </stringSet>
			    <integerList>
			        <member>1</member>
			        <member>2</member>
			    </integerList>
			    <booleanList>
			        <member>true</member>
			        <member>false</member>
			    </booleanList>
			    <timestampList>
			        <member>2014-04-29T18:30:38Z</member>
			        <member>2014-04-29T18:30:38Z</member>
			    </timestampList>
			    <enumList>
			        <member>Foo</member>
			        <member>0</member>
			    </enumList>
			    <nestedStringList>
			        <member>
			            <member>foo</member>
			            <member>bar</member>
			        </member>
			        <member>
			            <member>baz</member>
			            <member>qux</member>
			        </member>
			    </nestedStringList>
			    <renamed>
			        <item>foo</item>
			        <item>bar</item>
			    </renamed>
			    <flattenedList>hi</flattenedList>
			    <flattenedList>bye</flattenedList>
			    <customName>yep</customName>
			    <customName>nope</customName>
			    <myStructureList>
			        <item>
			            <value>1</value>
			            <other>2</other>
			        </item>
			        <item>
			            <value>3</value>
			            <other>4</other>
			        </item>
			    </myStructureList>
			    <RequestId>requestid</RequestId>
			</XmlListsResponse>
			`),
			ExpectResult: &XmlListsOutput{
				StringList: []*string{
					ptr.String("foo"),
					ptr.String("bar"),
				},
				StringSet: []*string{
					ptr.String("foo"),
					ptr.String("bar"),
				},
				IntegerList: []*int32{
					ptr.Int32(1),
					ptr.Int32(2),
				},
				BooleanList: []*bool{
					ptr.Bool(true),
					ptr.Bool(false),
				},
				TimestampList: []*time.Time{
					ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
					ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
				},
				EnumList: []types.FooEnum{
					types.FooEnum("Foo"),
					types.FooEnum("0"),
				},
				NestedStringList: [][]*string{
					{
						ptr.String("foo"),
						ptr.String("bar"),
					},
					{
						ptr.String("baz"),
						ptr.String("qux"),
					},
				},
				RenamedListMembers: []*string{
					ptr.String("foo"),
					ptr.String("bar"),
				},
				FlattenedList: []*string{
					ptr.String("hi"),
					ptr.String("bye"),
				},
				FlattenedList2: []*string{
					ptr.String("yep"),
					ptr.String("nope"),
				},
				StructureList: []*types.StructureListMember{
					{
						A: ptr.String("1"),
						B: ptr.String("2"),
					},
					{
						A: ptr.String("3"),
						B: ptr.String("4"),
					},
				},
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for k, vs := range c.Header {
					for _, v := range vs {
						w.Header().Add(k, v)
					}
				}
				if len(c.BodyMediaType) != 0 && len(w.Header().Values("Content-Type")) == 0 {
					w.Header().Set("Content-Type", c.BodyMediaType)
				}
				if len(c.Body) != 0 {
					w.Header().Set("Content-Length", strconv.Itoa(len(c.Body)))
				}
				w.WriteHeader(c.StatusCode)
				if len(c.Body) != 0 {
					if _, err := io.Copy(w, bytes.NewReader(c.Body)); err != nil {
						t.Errorf("failed to write response body, %v", err)
					}
				}
			}))
			defer server.Close()
			client := New(Options{
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options ResolverOptions) (e aws.Endpoint, err error) {
					e.URL = server.URL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               aws.NewBuildableHTTPClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params XmlListsInput
			result, err := client.XmlLists(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, cmpopts.IgnoreUnexported(middleware.Metadata{})); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
