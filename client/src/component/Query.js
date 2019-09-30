import React from 'react';
import {useFetchInitialData} from '../api/useFetchInitialData';

function Query({
  apiPath,
  errorRenderer,
  loadingComponent,
  successRenderer,
}) {
  const {
    data,
    error,
    isLoading,
  } = useFetchInitialData(apiPath);
  if (isLoading || data == null) {
    return loadingComponent;
  } else if (error != null) {
    return errorRenderer(error);
  } else {
    return successRenderer(data);
  }
};

export default React.memo(Query);
