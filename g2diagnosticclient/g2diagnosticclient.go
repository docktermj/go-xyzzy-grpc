/*
 *
 */

// Package main implements a client for the service.
package g2diagnosticclient

import (
	"context"
	pb "github.com/docktermj/go-xyzzy-grpc/g2diagnostic"
)

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// CheckDBPerf performs inserts to determine rate of insertion.
func (client *G2diagnosticClient) CheckDBPerf(ctx context.Context, secondsToRun int) (string, error) {
	request := pb.CheckDBPerfRequest{
		SecondsToRun: int32(secondsToRun),
	}
	response, err := client.G2DiagnosticGrpcClient.CheckDBPerf(ctx, &request)
	result := response.GetResult()
	return result, err
}

// TODO: Document.
func (client *G2diagnosticClient) ClearLastException(ctx context.Context) error {
	request := pb.ClearLastExceptionRequest{}
	_, err := client.G2DiagnosticGrpcClient.ClearLastException(ctx, &request)
	return err
}

//// TODO: Document.
//func (client *G2diagnosticClient) CloseEntityListBySize(ctx context.Context, entityListBySizeHandle interface{}) error {
//	//  _DLEXPORT int G2Diagnostic_closeEntityListBySize(EntityListBySizeHandle entityListBySizeHandle);
//	var err error = nil
//	result := C.G2Diagnostic_closeEntityListBySize(C.EntityListBySizeHandle(&entityListBySizeHandle))
//	if result != 0 {
//		err = g2diagnostic.getError(ctx, 2)
//	}
//	return err
//}
//
//// TODO: Document.
//func (client *G2diagnosticClient) Destroy(ctx context.Context) error {
//	//  _DLEXPORT int G2Diagnostic_destroy();
//	var err error = nil
//	result := C.G2Diagnostic_destroy()
//	if result != 0 {
//		err = g2diagnostic.getError(ctx, 3)
//	}
//	return err
//}
//
//// TODO: Document.
//func (client *G2diagnosticClient) FetchNextEntityBySize(ctx context.Context, entityListBySizeHandle interface{}) (string, error) {
//	//  _DLEXPORT int G2Diagnostic_fetchNextEntityBySize(EntityListBySizeHandle entityListBySizeHandle, char *responseBuf, const size_t bufSize);
//	var err error = nil
//	stringBuffer := g2diagnostic.getByteArray(initialByteArraySize)
//	result := C.G2Diagnostic_fetchNextEntityBySize(C.EntityListBySizeHandle(&entityListBySizeHandle), (*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
//	if result != 0 {
//		err = g2diagnostic.getError(ctx, 4)
//	}
//	stringBuffer = bytes.Trim(stringBuffer, "\x00")
//	return string(stringBuffer), err
//}
//
//// TODO: Document.
//func (client *G2diagnosticClient) FindEntitiesByFeatureIDs(ctx context.Context, features string) (string, error) {
//	//  _DLEXPORT int G2Diagnostic_findEntitiesByFeatureIDs(const char *features, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
//	var err error = nil
//	featuresForC := C.CString(features)
//	defer C.free(unsafe.Pointer(featuresForC))
//	stringBuffer := C.GoString(C.G2Diagnostic_findEntitiesByFeatureIDs_local(featuresForC))
//	if len(stringBuffer) == 0 {
//		err = g2diagnostic.getError(ctx, 5, features)
//	}
//	return stringBuffer, err
//}
//
//// GetAvailableMemory returns the available memory, in bytes, on the host system.
//func (client *G2diagnosticClient) GetAvailableMemory(ctx context.Context) (int64, error) {
//	// _DLEXPORT long long G2Diagnostic_getAvailableMemory();
//	var err error = nil
//	result := C.G2Diagnostic_getAvailableMemory()
//	return int64(result), err
//}
//
//// TODO: Document.
//func (client *G2diagnosticClient) GetDataSourceCounts(ctx context.Context) (string, error) {
//	//  _DLEXPORT int G2Diagnostic_getDataSourceCounts(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
//	var err error = nil
//	stringBuffer := C.GoString(C.G2Diagnostic_getDataSourceCounts_local())
//	if len(stringBuffer) == 0 {
//		err = g2diagnostic.getError(ctx, 6)
//	}
//	return stringBuffer, err
//}
//
//// GetDBInfo returns information about the database connection.
//func (client *G2diagnosticClient) GetDBInfo(ctx context.Context) (string, error) {
//	// _DLEXPORT int G2Diagnostic_getDBInfo(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
//	var err error = nil
//	stringBuffer := C.GoString(C.G2Diagnostic_getDBInfo_local())
//	if len(stringBuffer) == 0 {
//		err = g2diagnostic.getError(ctx, 7)
//	}
//	return stringBuffer, err
//}
//
//// TODO: Document.
//func (client *G2diagnosticClient) GetEntityDetails(ctx context.Context, entityID int64, includeInternalFeatures int) (string, error) {
//	//  _DLEXPORT int G2Diagnostic_getEntityDetails(const long long entityID, const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
//	var err error = nil
//	stringBuffer := C.GoString(C.G2Diagnostic_getEntityDetails_local(C.longlong(entityID), C.int(includeInternalFeatures)))
//	if len(stringBuffer) == 0 {
//		err = g2diagnostic.getError(ctx, 8, strconv.FormatInt(entityID, 10), strconv.Itoa(includeInternalFeatures))
//	}
//	return stringBuffer, err
//}
//
//// TODO: Document.
//func (client *G2diagnosticClient) GetEntityListBySize(ctx context.Context, entitySize int) (interface{}, error) {
//	//  _DLEXPORT int G2Diagnostic_getEntityListBySize(const size_t entitySize, EntityListBySizeHandle* entityListBySizeHandle);
//	var err error = nil
//	var entityListBySizeHandle unsafe.Pointer
//	result := C.G2Diagnostic_getEntityListBySize(C.size_t(entitySize), (*C.EntityListBySizeHandle)(&entityListBySizeHandle))
//	if result != 0 {
//		err = g2diagnostic.getError(ctx, 9, strconv.Itoa(entitySize))
//	}
//	return entityListBySizeHandle, err
//}
//
//// TODO: Document.
//func (client *G2diagnosticClient) GetEntityResume(ctx context.Context, entityID int64) (string, error) {
//	//  _DLEXPORT int G2Diagnostic_getEntityResume(const long long entityID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
//	var err error = nil
//	stringBuffer := C.GoString(C.G2Diagnostic_getEntityResume_local(C.longlong(entityID)))
//	if len(stringBuffer) == 0 {
//		err = g2diagnostic.getError(ctx, 10, strconv.FormatInt(entityID, 10))
//	}
//	return stringBuffer, err
//}
//
//// TODO: Document.
//func (client *G2diagnosticClient) GetEntitySizeBreakdown(ctx context.Context, minimumEntitySize int, includeInternalFeatures int) (string, error) {
//	//  _DLEXPORT int G2Diagnostic_getEntitySizeBreakdown(const size_t minimumEntitySize, const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
//	var err error = nil
//	stringBuffer := C.GoString(C.G2Diagnostic_getEntitySizeBreakdown_local(C.size_t(minimumEntitySize), C.int(includeInternalFeatures)))
//	if len(stringBuffer) == 0 {
//		err = g2diagnostic.getError(ctx, 11, strconv.Itoa(minimumEntitySize), strconv.Itoa(includeInternalFeatures))
//	}
//	return stringBuffer, err
//}
//
//// TODO: Document.
//func (client *G2diagnosticClient) GetFeature(ctx context.Context, libFeatID int64) (string, error) {
//	//  _DLEXPORT int G2Diagnostic_getFeature(const long long libFeatID, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize));
//	var err error = nil
//	stringBuffer := C.GoString(C.G2Diagnostic_getFeature_local(C.longlong(libFeatID)))
//	if len(stringBuffer) == 0 {
//		err = g2diagnostic.getError(ctx, 12, strconv.FormatInt(libFeatID, 10))
//	}
//	return stringBuffer, err
//}
//
//// TODO: Document.
//func (client *G2diagnosticClient) GetGenericFeatures(ctx context.Context, featureType string, maximumEstimatedCount int) (string, error) {
//	//  _DLEXPORT int G2Diagnostic_getGenericFeatures(const char* featureType, const size_t maximumEstimatedCount, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
//	var err error = nil
//	featureTypeForC := C.CString(featureType)
//	defer C.free(unsafe.Pointer(featureTypeForC))
//	stringBuffer := C.GoString(C.G2Diagnostic_getGenericFeatures_local(featureTypeForC, C.size_t(maximumEstimatedCount)))
//	if len(stringBuffer) == 0 {
//		err = g2diagnostic.getError(ctx, 13, featureType, strconv.Itoa(maximumEstimatedCount))
//	}
//	return stringBuffer, err
//}

// GetLastException returns the last exception encountered in the Senzing Engine.
func (client *G2diagnosticClient) GetLastException(ctx context.Context) (string, error) {
	request := pb.GetLastExceptionRequest{}
	response, err := client.G2DiagnosticGrpcClient.GetLastException(ctx, &request)
	result := response.GetResult()
	return result, err
}

//// TODO: Document.
//func (client *G2diagnosticClient) GetLastExceptionCode(ctx context.Context) (string, error) {
//	//  _DLEXPORT int G2Diagnostic_getLastException(char *buffer, const size_t bufSize);
//	var err error = nil
//	stringBuffer := g2diagnostic.getByteArray(initialByteArraySize)
//	result := C.G2Diagnostic_getLastException((*C.char)(unsafe.Pointer(&stringBuffer[0])), C.ulong(len(stringBuffer)))
//	if result != 0 {
//		err = g2diagnostic.getError(ctx, 14)
//	}
//	stringBuffer = bytes.Trim(stringBuffer, "\x00")
//	return string(stringBuffer), err
//}
//
//// GetLogicalCores returns the number of logical cores on the host system.
//func (client *G2diagnosticClient) GetLogicalCores(ctx context.Context) (int, error) {
//	// _DLEXPORT int G2Diagnostic_getLogicalCores();
//	var err error = nil
//	result := C.G2Diagnostic_getLogicalCores()
//	return int(result), err
//}
//
//// TODO: Document.
//func (client *G2diagnosticClient) GetMappingStatistics(ctx context.Context, includeInternalFeatures int) (string, error) {
//	//  _DLEXPORT int G2Diagnostic_getMappingStatistics(const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
//	var err error = nil
//	stringBuffer := C.GoString(C.G2Diagnostic_getMappingStatistics_local(C.int(includeInternalFeatures)))
//	if len(stringBuffer) == 0 {
//		err = g2diagnostic.getError(ctx, 15, strconv.Itoa(includeInternalFeatures))
//	}
//	return stringBuffer, err
//}
//
//// GetPhysicalCores returns the number of physical cores on the host system.
//func (client *G2diagnosticClient) GetPhysicalCores(ctx context.Context) (int, error) {
//	// _DLEXPORT int G2Diagnostic_getPhysicalCores();
//	var err error = nil
//	result := C.G2Diagnostic_getPhysicalCores()
//	return int(result), err
//}
//
//// TODO: Document.
//func (client *G2diagnosticClient) GetRelationshipDetails(ctx context.Context, relationshipID int64, includeInternalFeatures int) (string, error) {
//	//  _DLEXPORT int G2Diagnostic_getRelationshipDetails(const long long relationshipID, const int includeInternalFeatures, char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
//	var err error = nil
//	stringBuffer := C.GoString(C.G2Diagnostic_getRelationshipDetails_local(C.longlong(relationshipID), C.int(includeInternalFeatures)))
//	if len(stringBuffer) == 0 {
//		err = g2diagnostic.getError(ctx, 16, strconv.FormatInt(relationshipID, 10), strconv.Itoa(includeInternalFeatures))
//	}
//	return stringBuffer, err
//}
//
//// TODO: Document.
//func (client *G2diagnosticClient) GetResolutionStatistics(ctx context.Context) (string, error) {
//	//  _DLEXPORT int G2Diagnostic_getResolutionStatistics(char **responseBuf, size_t *bufSize, void *(*resizeFunc)(void *ptr, size_t newSize) );
//	var err error = nil
//	stringBuffer := C.GoString(C.G2Diagnostic_getResolutionStatistics_local())
//	if len(stringBuffer) == 0 {
//		err = g2diagnostic.getError(ctx, 17)
//	}
//	return stringBuffer, err
//}
//
//// GetTotalSystemMemory returns the total memory, in bytes, on the host system.
//func (client *G2diagnosticClient) GetTotalSystemMemory(ctx context.Context) (int64, error) {
//	// _DLEXPORT long long G2Diagnostic_getTotalSystemMemory();
//	var err error = nil
//	result := C.G2Diagnostic_getTotalSystemMemory()
//	return int64(result), err
//}
//

// Init initializes the Senzing G2diagnosis.
func (client *G2diagnosticClient) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	request := pb.InitRequest{
		ModuleName:     moduleName,
		IniParams:      iniParams,
		VerboseLogging: int32(verboseLogging),
	}
	_, err := client.G2DiagnosticGrpcClient.Init(ctx, &request)
	return err
}

//
//// TODO: Document.
//func (client *G2diagnosticClient) InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int) error {
//	//  _DLEXPORT int G2Diagnostic_initWithConfigID(const char *moduleName, const char *iniParams, const long long initConfigID, const int verboseLogging);
//	var err error = nil
//	moduleNameForC := C.CString(moduleName)
//	defer C.free(unsafe.Pointer(moduleNameForC))
//	iniParamsForC := C.CString(iniParams)
//	defer C.free(unsafe.Pointer(iniParamsForC))
//	result := C.G2Diagnostic_initWithConfigID(moduleNameForC, iniParamsForC, C.longlong(initConfigID), C.int(verboseLogging))
//	if result != 0 {
//		err = g2diagnostic.getError(ctx, 19, moduleName, iniParams, strconv.FormatInt(initConfigID, 10), strconv.Itoa(verboseLogging))
//	}
//	return err
//}
//
//// Null shows how to report a BUG inline.
//func (client *G2diagnosticClient) Null(ctx context.Context) (int64, error) {
//	// BUG(mjd): Just an example of how to show bugs in GoDoc.
//	var err error = nil
//	return 1, err
//}
//
//// TODO: Document.
//func (client *G2diagnosticClient) Reinit(ctx context.Context, initConfigID int64) error {
//	//  _DLEXPORT int G2Diagnostic_reinit(const long long initConfigID);
//	var err error = nil
//	result := C.G2Diagnostic_reinit(C.longlong(initConfigID))
//	if result != 0 {
//		err = g2diagnostic.getError(ctx, 20, strconv.FormatInt(initConfigID, 10))
//	}
//	return err
//}
