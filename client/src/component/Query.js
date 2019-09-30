// @flow strict

import type {Node} from 'react';

import {useFetchInitialData} from '../api/useFetchInitialData';

type Props<T> = {
  apiPath: string,
  errorRenderer: Error => Node,
  loadingComponent: Node,
  successRenderer: T => Node,
};

function Query<T>({
  apiPath,
  errorRenderer,
  loadingComponent,
  successRenderer,
}: Props<T>) {
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

export default Query;
