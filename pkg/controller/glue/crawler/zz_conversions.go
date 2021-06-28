/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package crawler

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/glue"

	svcapitypes "github.com/crossplane/provider-aws/apis/glue/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetCrawlerInput returns input for read
// operation.
func GenerateGetCrawlerInput(cr *svcapitypes.Crawler) *svcsdk.GetCrawlerInput {
	res := &svcsdk.GetCrawlerInput{}

	if cr.Spec.ForProvider.Name != nil {
		res.SetName(*cr.Spec.ForProvider.Name)
	}

	return res
}

// GenerateCrawler returns the current state in the form of *svcapitypes.Crawler.
func GenerateCrawler(resp *svcsdk.GetCrawlerOutput) *svcapitypes.Crawler {
	cr := &svcapitypes.Crawler{}

	return cr
}

// GenerateCreateCrawlerInput returns a create input.
func GenerateCreateCrawlerInput(cr *svcapitypes.Crawler) *svcsdk.CreateCrawlerInput {
	res := &svcsdk.CreateCrawlerInput{}

	if cr.Spec.ForProvider.Classifiers != nil {
		f0 := []*string{}
		for _, f0iter := range cr.Spec.ForProvider.Classifiers {
			var f0elem string
			f0elem = *f0iter
			f0 = append(f0, &f0elem)
		}
		res.SetClassifiers(f0)
	}
	if cr.Spec.ForProvider.Configuration != nil {
		res.SetConfiguration(*cr.Spec.ForProvider.Configuration)
	}
	if cr.Spec.ForProvider.CrawlerSecurityConfiguration != nil {
		res.SetCrawlerSecurityConfiguration(*cr.Spec.ForProvider.CrawlerSecurityConfiguration)
	}
	if cr.Spec.ForProvider.DatabaseName != nil {
		res.SetDatabaseName(*cr.Spec.ForProvider.DatabaseName)
	}
	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Spec.ForProvider.LineageConfiguration != nil {
		f5 := &svcsdk.LineageConfiguration{}
		if cr.Spec.ForProvider.LineageConfiguration.CrawlerLineageSettings != nil {
			f5.SetCrawlerLineageSettings(*cr.Spec.ForProvider.LineageConfiguration.CrawlerLineageSettings)
		}
		res.SetLineageConfiguration(f5)
	}
	if cr.Spec.ForProvider.Name != nil {
		res.SetName(*cr.Spec.ForProvider.Name)
	}
	if cr.Spec.ForProvider.RecrawlPolicy != nil {
		f7 := &svcsdk.RecrawlPolicy{}
		if cr.Spec.ForProvider.RecrawlPolicy.RecrawlBehavior != nil {
			f7.SetRecrawlBehavior(*cr.Spec.ForProvider.RecrawlPolicy.RecrawlBehavior)
		}
		res.SetRecrawlPolicy(f7)
	}
	if cr.Spec.ForProvider.Role != nil {
		res.SetRole(*cr.Spec.ForProvider.Role)
	}
	if cr.Spec.ForProvider.Schedule != nil {
		res.SetSchedule(*cr.Spec.ForProvider.Schedule)
	}
	if cr.Spec.ForProvider.SchemaChangePolicy != nil {
		f10 := &svcsdk.SchemaChangePolicy{}
		if cr.Spec.ForProvider.SchemaChangePolicy.DeleteBehavior != nil {
			f10.SetDeleteBehavior(*cr.Spec.ForProvider.SchemaChangePolicy.DeleteBehavior)
		}
		if cr.Spec.ForProvider.SchemaChangePolicy.UpdateBehavior != nil {
			f10.SetUpdateBehavior(*cr.Spec.ForProvider.SchemaChangePolicy.UpdateBehavior)
		}
		res.SetSchemaChangePolicy(f10)
	}
	if cr.Spec.ForProvider.TablePrefix != nil {
		res.SetTablePrefix(*cr.Spec.ForProvider.TablePrefix)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f12 := map[string]*string{}
		for f12key, f12valiter := range cr.Spec.ForProvider.Tags {
			var f12val string
			f12val = *f12valiter
			f12[f12key] = &f12val
		}
		res.SetTags(f12)
	}
	if cr.Spec.ForProvider.Targets != nil {
		f13 := &svcsdk.CrawlerTargets{}
		if cr.Spec.ForProvider.Targets.CatalogTargets != nil {
			f13f0 := []*svcsdk.CatalogTarget{}
			for _, f13f0iter := range cr.Spec.ForProvider.Targets.CatalogTargets {
				f13f0elem := &svcsdk.CatalogTarget{}
				if f13f0iter.DatabaseName != nil {
					f13f0elem.SetDatabaseName(*f13f0iter.DatabaseName)
				}
				if f13f0iter.Tables != nil {
					f13f0elemf1 := []*string{}
					for _, f13f0elemf1iter := range f13f0iter.Tables {
						var f13f0elemf1elem string
						f13f0elemf1elem = *f13f0elemf1iter
						f13f0elemf1 = append(f13f0elemf1, &f13f0elemf1elem)
					}
					f13f0elem.SetTables(f13f0elemf1)
				}
				f13f0 = append(f13f0, f13f0elem)
			}
			f13.SetCatalogTargets(f13f0)
		}
		if cr.Spec.ForProvider.Targets.DynamoDBTargets != nil {
			f13f1 := []*svcsdk.DynamoDBTarget{}
			for _, f13f1iter := range cr.Spec.ForProvider.Targets.DynamoDBTargets {
				f13f1elem := &svcsdk.DynamoDBTarget{}
				if f13f1iter.Path != nil {
					f13f1elem.SetPath(*f13f1iter.Path)
				}
				if f13f1iter.ScanAll != nil {
					f13f1elem.SetScanAll(*f13f1iter.ScanAll)
				}
				if f13f1iter.ScanRate != nil {
					f13f1elem.SetScanRate(*f13f1iter.ScanRate)
				}
				f13f1 = append(f13f1, f13f1elem)
			}
			f13.SetDynamoDBTargets(f13f1)
		}
		if cr.Spec.ForProvider.Targets.JdbcTargets != nil {
			f13f2 := []*svcsdk.JdbcTarget{}
			for _, f13f2iter := range cr.Spec.ForProvider.Targets.JdbcTargets {
				f13f2elem := &svcsdk.JdbcTarget{}
				if f13f2iter.ConnectionName != nil {
					f13f2elem.SetConnectionName(*f13f2iter.ConnectionName)
				}
				if f13f2iter.Exclusions != nil {
					f13f2elemf1 := []*string{}
					for _, f13f2elemf1iter := range f13f2iter.Exclusions {
						var f13f2elemf1elem string
						f13f2elemf1elem = *f13f2elemf1iter
						f13f2elemf1 = append(f13f2elemf1, &f13f2elemf1elem)
					}
					f13f2elem.SetExclusions(f13f2elemf1)
				}
				if f13f2iter.Path != nil {
					f13f2elem.SetPath(*f13f2iter.Path)
				}
				f13f2 = append(f13f2, f13f2elem)
			}
			f13.SetJdbcTargets(f13f2)
		}
		if cr.Spec.ForProvider.Targets.MongoDBTargets != nil {
			f13f3 := []*svcsdk.MongoDBTarget{}
			for _, f13f3iter := range cr.Spec.ForProvider.Targets.MongoDBTargets {
				f13f3elem := &svcsdk.MongoDBTarget{}
				if f13f3iter.ConnectionName != nil {
					f13f3elem.SetConnectionName(*f13f3iter.ConnectionName)
				}
				if f13f3iter.Path != nil {
					f13f3elem.SetPath(*f13f3iter.Path)
				}
				if f13f3iter.ScanAll != nil {
					f13f3elem.SetScanAll(*f13f3iter.ScanAll)
				}
				f13f3 = append(f13f3, f13f3elem)
			}
			f13.SetMongoDBTargets(f13f3)
		}
		if cr.Spec.ForProvider.Targets.S3Targets != nil {
			f13f4 := []*svcsdk.S3Target{}
			for _, f13f4iter := range cr.Spec.ForProvider.Targets.S3Targets {
				f13f4elem := &svcsdk.S3Target{}
				if f13f4iter.ConnectionName != nil {
					f13f4elem.SetConnectionName(*f13f4iter.ConnectionName)
				}
				if f13f4iter.Exclusions != nil {
					f13f4elemf1 := []*string{}
					for _, f13f4elemf1iter := range f13f4iter.Exclusions {
						var f13f4elemf1elem string
						f13f4elemf1elem = *f13f4elemf1iter
						f13f4elemf1 = append(f13f4elemf1, &f13f4elemf1elem)
					}
					f13f4elem.SetExclusions(f13f4elemf1)
				}
				if f13f4iter.Path != nil {
					f13f4elem.SetPath(*f13f4iter.Path)
				}
				f13f4 = append(f13f4, f13f4elem)
			}
			f13.SetS3Targets(f13f4)
		}
		res.SetTargets(f13)
	}

	return res
}

// GenerateUpdateCrawlerInput returns an update input.
func GenerateUpdateCrawlerInput(cr *svcapitypes.Crawler) *svcsdk.UpdateCrawlerInput {
	res := &svcsdk.UpdateCrawlerInput{}

	if cr.Spec.ForProvider.Classifiers != nil {
		f0 := []*string{}
		for _, f0iter := range cr.Spec.ForProvider.Classifiers {
			var f0elem string
			f0elem = *f0iter
			f0 = append(f0, &f0elem)
		}
		res.SetClassifiers(f0)
	}
	if cr.Spec.ForProvider.Configuration != nil {
		res.SetConfiguration(*cr.Spec.ForProvider.Configuration)
	}
	if cr.Spec.ForProvider.CrawlerSecurityConfiguration != nil {
		res.SetCrawlerSecurityConfiguration(*cr.Spec.ForProvider.CrawlerSecurityConfiguration)
	}
	if cr.Spec.ForProvider.DatabaseName != nil {
		res.SetDatabaseName(*cr.Spec.ForProvider.DatabaseName)
	}
	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Spec.ForProvider.LineageConfiguration != nil {
		f5 := &svcsdk.LineageConfiguration{}
		if cr.Spec.ForProvider.LineageConfiguration.CrawlerLineageSettings != nil {
			f5.SetCrawlerLineageSettings(*cr.Spec.ForProvider.LineageConfiguration.CrawlerLineageSettings)
		}
		res.SetLineageConfiguration(f5)
	}
	if cr.Spec.ForProvider.Name != nil {
		res.SetName(*cr.Spec.ForProvider.Name)
	}
	if cr.Spec.ForProvider.RecrawlPolicy != nil {
		f7 := &svcsdk.RecrawlPolicy{}
		if cr.Spec.ForProvider.RecrawlPolicy.RecrawlBehavior != nil {
			f7.SetRecrawlBehavior(*cr.Spec.ForProvider.RecrawlPolicy.RecrawlBehavior)
		}
		res.SetRecrawlPolicy(f7)
	}
	if cr.Spec.ForProvider.Role != nil {
		res.SetRole(*cr.Spec.ForProvider.Role)
	}
	if cr.Spec.ForProvider.Schedule != nil {
		res.SetSchedule(*cr.Spec.ForProvider.Schedule)
	}
	if cr.Spec.ForProvider.SchemaChangePolicy != nil {
		f10 := &svcsdk.SchemaChangePolicy{}
		if cr.Spec.ForProvider.SchemaChangePolicy.DeleteBehavior != nil {
			f10.SetDeleteBehavior(*cr.Spec.ForProvider.SchemaChangePolicy.DeleteBehavior)
		}
		if cr.Spec.ForProvider.SchemaChangePolicy.UpdateBehavior != nil {
			f10.SetUpdateBehavior(*cr.Spec.ForProvider.SchemaChangePolicy.UpdateBehavior)
		}
		res.SetSchemaChangePolicy(f10)
	}
	if cr.Spec.ForProvider.TablePrefix != nil {
		res.SetTablePrefix(*cr.Spec.ForProvider.TablePrefix)
	}
	if cr.Spec.ForProvider.Targets != nil {
		f12 := &svcsdk.CrawlerTargets{}
		if cr.Spec.ForProvider.Targets.CatalogTargets != nil {
			f12f0 := []*svcsdk.CatalogTarget{}
			for _, f12f0iter := range cr.Spec.ForProvider.Targets.CatalogTargets {
				f12f0elem := &svcsdk.CatalogTarget{}
				if f12f0iter.DatabaseName != nil {
					f12f0elem.SetDatabaseName(*f12f0iter.DatabaseName)
				}
				if f12f0iter.Tables != nil {
					f12f0elemf1 := []*string{}
					for _, f12f0elemf1iter := range f12f0iter.Tables {
						var f12f0elemf1elem string
						f12f0elemf1elem = *f12f0elemf1iter
						f12f0elemf1 = append(f12f0elemf1, &f12f0elemf1elem)
					}
					f12f0elem.SetTables(f12f0elemf1)
				}
				f12f0 = append(f12f0, f12f0elem)
			}
			f12.SetCatalogTargets(f12f0)
		}
		if cr.Spec.ForProvider.Targets.DynamoDBTargets != nil {
			f12f1 := []*svcsdk.DynamoDBTarget{}
			for _, f12f1iter := range cr.Spec.ForProvider.Targets.DynamoDBTargets {
				f12f1elem := &svcsdk.DynamoDBTarget{}
				if f12f1iter.Path != nil {
					f12f1elem.SetPath(*f12f1iter.Path)
				}
				if f12f1iter.ScanAll != nil {
					f12f1elem.SetScanAll(*f12f1iter.ScanAll)
				}
				if f12f1iter.ScanRate != nil {
					f12f1elem.SetScanRate(*f12f1iter.ScanRate)
				}
				f12f1 = append(f12f1, f12f1elem)
			}
			f12.SetDynamoDBTargets(f12f1)
		}
		if cr.Spec.ForProvider.Targets.JdbcTargets != nil {
			f12f2 := []*svcsdk.JdbcTarget{}
			for _, f12f2iter := range cr.Spec.ForProvider.Targets.JdbcTargets {
				f12f2elem := &svcsdk.JdbcTarget{}
				if f12f2iter.ConnectionName != nil {
					f12f2elem.SetConnectionName(*f12f2iter.ConnectionName)
				}
				if f12f2iter.Exclusions != nil {
					f12f2elemf1 := []*string{}
					for _, f12f2elemf1iter := range f12f2iter.Exclusions {
						var f12f2elemf1elem string
						f12f2elemf1elem = *f12f2elemf1iter
						f12f2elemf1 = append(f12f2elemf1, &f12f2elemf1elem)
					}
					f12f2elem.SetExclusions(f12f2elemf1)
				}
				if f12f2iter.Path != nil {
					f12f2elem.SetPath(*f12f2iter.Path)
				}
				f12f2 = append(f12f2, f12f2elem)
			}
			f12.SetJdbcTargets(f12f2)
		}
		if cr.Spec.ForProvider.Targets.MongoDBTargets != nil {
			f12f3 := []*svcsdk.MongoDBTarget{}
			for _, f12f3iter := range cr.Spec.ForProvider.Targets.MongoDBTargets {
				f12f3elem := &svcsdk.MongoDBTarget{}
				if f12f3iter.ConnectionName != nil {
					f12f3elem.SetConnectionName(*f12f3iter.ConnectionName)
				}
				if f12f3iter.Path != nil {
					f12f3elem.SetPath(*f12f3iter.Path)
				}
				if f12f3iter.ScanAll != nil {
					f12f3elem.SetScanAll(*f12f3iter.ScanAll)
				}
				f12f3 = append(f12f3, f12f3elem)
			}
			f12.SetMongoDBTargets(f12f3)
		}
		if cr.Spec.ForProvider.Targets.S3Targets != nil {
			f12f4 := []*svcsdk.S3Target{}
			for _, f12f4iter := range cr.Spec.ForProvider.Targets.S3Targets {
				f12f4elem := &svcsdk.S3Target{}
				if f12f4iter.ConnectionName != nil {
					f12f4elem.SetConnectionName(*f12f4iter.ConnectionName)
				}
				if f12f4iter.Exclusions != nil {
					f12f4elemf1 := []*string{}
					for _, f12f4elemf1iter := range f12f4iter.Exclusions {
						var f12f4elemf1elem string
						f12f4elemf1elem = *f12f4elemf1iter
						f12f4elemf1 = append(f12f4elemf1, &f12f4elemf1elem)
					}
					f12f4elem.SetExclusions(f12f4elemf1)
				}
				if f12f4iter.Path != nil {
					f12f4elem.SetPath(*f12f4iter.Path)
				}
				f12f4 = append(f12f4, f12f4elem)
			}
			f12.SetS3Targets(f12f4)
		}
		res.SetTargets(f12)
	}

	return res
}

// GenerateDeleteCrawlerInput returns a deletion input.
func GenerateDeleteCrawlerInput(cr *svcapitypes.Crawler) *svcsdk.DeleteCrawlerInput {
	res := &svcsdk.DeleteCrawlerInput{}

	if cr.Spec.ForProvider.Name != nil {
		res.SetName(*cr.Spec.ForProvider.Name)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "EntityNotFoundException"
}