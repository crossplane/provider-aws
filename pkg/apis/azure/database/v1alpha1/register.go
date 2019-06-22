/*
Copyright 2018 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/runtime/scheme"
)

// Package type metadata.
const (
	Group   = "database.azure.crossplane.io"
	Version = "v1alpha1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

// MysqlServer type metadata.
var (
	MysqlServerKind             = reflect.TypeOf(MysqlServer{}).Name()
	MysqlServerKindAPIVersion   = MysqlServerKind + "." + SchemeGroupVersion.String()
	MysqlServerGroupVersionKind = SchemeGroupVersion.WithKind(MysqlServerKind)
)

// PostgresqlServer type metadata.
var (
	PostgresqlServerKind             = reflect.TypeOf(PostgresqlServer{}).Name()
	PostgresqlServerKindAPIVersion   = PostgresqlServerKind + "." + SchemeGroupVersion.String()
	PostgresqlServerGroupVersionKind = SchemeGroupVersion.WithKind(PostgresqlServerKind)
)

func init() {
	SchemeBuilder.Register(&MysqlServer{}, &MysqlServerList{})
	SchemeBuilder.Register(&PostgresqlServer{}, &PostgresqlServerList{})
}
